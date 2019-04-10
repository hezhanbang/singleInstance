#pragma once

#ifndef __SINGLEINSTANCE_H__
#define __SINGLEINSTANCE_H__

int CurrentProcessIsSingle(const char* singleKey, const char* lockFileName);

#endif // !__SINGLEINSTANCE_H__

