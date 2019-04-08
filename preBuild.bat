@echo off

set workDir=%cd%
set singleDir=%~dp0
cd %singleDir%

IF EXIST single.go del -F single.go
copy lib\singleWindows.go single.go
if %errorlevel% neq 0 goto ERROR
exit /b 0

:ERROR
echo "fail to do preBuild.bat in lib [SingleInstance]!"
cd /d %workDir%
exit /b 1
