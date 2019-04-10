// test.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include <Windows.h>
#include "../singleInstance.h"

int main()
{
	for (int i = 0; i < 10; i++) {
		int ret = hebCurrentProcessIsSingle("___myKey___", "");
		if (ret < 0) {
			printf("fail to check lock, ret=%d\n", ret);
			return -1;
		}
		if (HEB_SINGLE__SINGLING == ret) {
			printf("[%d] now we get locker, can run\n", i);
		}
		else {
			printf("[%d] another process is running, waiting it.\n", i);
		}

		::Sleep(1000 * 5);
	}
	return 0;
}

