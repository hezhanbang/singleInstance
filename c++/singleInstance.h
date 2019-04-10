#pragma once

#ifndef __SINGLEINSTANCE_H__
#define __SINGLEINSTANCE_H__

#define HEB_SINGLE__SINGLING 1
#define HEB_SINGLE__NO_SINGLE 0

int hebCurrentProcessIsSingle(const char* singleKey, const char* lockFileName);

#endif // !__SINGLEINSTANCE_H__

