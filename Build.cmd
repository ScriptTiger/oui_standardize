@echo off

if not exist "Release" md "Release"

set APP=oui_standardize

set GOARCH=amd64
call :Build_OS

set GOARCH=386
call :Build_OS

pause
exit /b

:Build_OS

set GOOS=windows
set EXT=.exe
call :Build

set GOOS=linux
set EXT=
call :Build

if %GOARCH% == 386 exit /b

set GOOS=darwin
set EXT=.app
call :Build

exit /b

:Build

echo Building %APP%_%GOOS%_%GOARCH%%EXT%...
go build -ldflags="-s -w" -o "Release/%APP%_%GOOS%_%GOARCH%%EXT%"
exit /b
