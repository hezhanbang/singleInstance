@echo off

set workDir=%cd%
set singleDir=%~dp0
cd %singleDir%

IF EXIST single.go del -F single.go
copy lib\singleWindows.go single.go

set retVal=%errorlevel%
cd /d %workDir%

if %retVal% neq 0 goto ERROR
exit /b 0

:ERROR
echo "fail to do preBuild.bat in lib [singleInstance]!"
exit /b 1
