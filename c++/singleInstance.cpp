#include "stdafx.h"  //不能写成真实的文件相对路径。
#include <Windows.h>
#include <string>
#include <iostream>
#include <sstream>
#include <iomanip>
#include <ctime>
#include "singleInstance.h"

bool gHeblockedByThis_single = false;

void hebSingleLocked(const char* key, bool &locked, bool &newLocker) {
	if (false == gHeblockedByThis_single) {
		//test for new locker
		HANDLE handle = ::CreateMutexA(NULL, TRUE, key);

		DWORD err = ::GetLastError();
		char* strErr = NULL;
		FormatMessageA(FORMAT_MESSAGE_ALLOCATE_BUFFER | FORMAT_MESSAGE_FROM_SYSTEM | FORMAT_MESSAGE_IGNORE_INSERTS,
			NULL, err,
			MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT),
			(LPSTR)&strErr, 0, NULL);

		fprintf(stdout, "CreateMutexA, handle=%d errInt=%d errStr=%s [singleInstance]\n", handle, err, strErr);
		LocalFree(strErr);

		//check return val and last err
		if (0 == err) {
			gHeblockedByThis_single = true;
			locked = true;
			newLocker = true;
		}
		else { //fail to get locker, we have to release reference count of the kernel object.
			if (INVALID_HANDLE_VALUE != handle) {
				::CloseHandle(handle);
				handle = INVALID_HANDLE_VALUE;
			}
			gHeblockedByThis_single = false;
			locked = false;
			newLocker = false;
		}

		return;
	}

	//we have keep this locker
	locked = true;
	newLocker = false;
	return;
}

int hebCurrentProcessIsSingle(const char* singleKey, const char* lockFileName) {
	if (!singleKey) {
		return -1;
	}
	int len = strlen(singleKey);
	if (len < 5 || len>20) {
		return -2;
	}

	std::string file;
	if (!lockFileName || strlen(lockFileName) == 0) {
		file = "lock.txt";
	}
	else {
		int len = strlen(lockFileName);
		if (len < 5 || len>20) {
			return -3;
		}
		file = lockFileName;
	}

	bool locked, newLocker;
	hebSingleLocked(singleKey, locked, newLocker);
	if (!locked) {
		return HEB_SINGLE__NO_SINGLE;
	}
	if (!newLocker) {
		return HEB_SINGLE__SINGLING;
	}

	//we get new locker, update time to file
	std::string exeDir(255, '\0');
	if (GetModuleFileNameA(NULL, (char*)exeDir.data(), 255) <= 0) {
		return -4;
	}
	if (strlen(exeDir.c_str()) + 5 > exeDir.size()) {
		return -5;
	}
	*(strrchr((char*)exeDir.data(), '\\')) = '\0';

	std::string path = exeDir.c_str();
	path += "\\";
	path += file;

	//create file
	FILE* pf = NULL;
	fopen_s(&pf, path.c_str(), "wb+");
	if (!pf) {
		return -6;
	}

	auto t = std::time(nullptr);
	struct tm tiM;
	localtime_s(&tiM, &t);

	std::ostringstream oss;
	oss << std::put_time(&tiM, "%d-%m-%Y %H-%M-%S");
	std::string timeStr = oss.str();

	fprintf(pf, "[%s] [pid=%d]\n", timeStr.c_str(), ::GetCurrentProcessId());
	fflush(pf);

	//do not close locker file
	return HEB_SINGLE__SINGLING;
}

