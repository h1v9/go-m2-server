@echo off
setlocal

REM Variables
set BINARY_NAME=server
set SRC_DIR=./cmd/game
set OUT_DIR=out
set CGO_ENABLED=0

REM Ensure the output directory exists
if not exist %OUT_DIR% (
    mkdir %OUT_DIR%
)

REM Check the input argument
if "%1" == "windows-amd64" goto build_windows_amd64
if "%1" == "windows-arm64" goto build_windows_arm64
if "%1" == "linux-amd64" goto build_linux_amd64
if "%1" == "linux-arm64" goto build_linux_arm64
if "%1" == "darwin-amd64" goto build_darwin_amd64
if "%1" == "darwin-arm64" goto build_darwin_arm64
if "%1" == "freebsd-amd64" goto build_freebsd_amd64
if "%1" == "freebsd-arm64" goto build_freebsd_arm64
if "%1" == "all" goto build_all
if "%1" == "clean" goto clean
goto :usage

REM Clean up the build
:clean
echo Cleaning up...
if exist %OUT_DIR% rd /S /Q %OUT_DIR%
goto :eof

REM Build for Windows AMD64
:build_windows_amd64
echo Building for Windows AMD64...
set GOOS=windows
set GOARCH=amd64
go build -o %OUT_DIR%\%BINARY_NAME%.windows-amd64.exe %SRC_DIR%
if errorlevel 1 goto :error
goto :eof

REM Build for Windows ARM64
:build_windows_arm64
echo Building for Windows ARM64...
set GOOS=windows
set GOARCH=arm64
go build -o %OUT_DIR%\%BINARY_NAME%.windows-arm64.exe %SRC_DIR%
if errorlevel 1 goto :error
goto :eof

REM Build for Linux AMD64
:build_linux_amd64
echo Building for Linux AMD64...
set GOOS=linux
set GOARCH=amd64
go build -o %OUT_DIR%\%BINARY_NAME%.linux-amd64 %SRC_DIR%
if errorlevel 1 goto :error
goto :eof

REM Build for Linux ARM64
:build_linux_arm64
echo Building for Linux ARM64...
set GOOS=linux
set GOARCH=arm64
go build -o %OUT_DIR%\%BINARY_NAME%.linux-arm64 %SRC_DIR%
if errorlevel 1 goto :error
goto :eof

REM Build for macOS AMD64
:build_darwin_amd64
echo Building for macOS AMD64...
set GOOS=darwin
set GOARCH=amd64
go build -o %OUT_DIR%\%BINARY_NAME%.darwin-amd64 %SRC_DIR%
if errorlevel 1 goto :error
goto :eof

REM Build for macOS ARM64
:build_darwin_arm64
echo Building for macOS ARM64...
set GOOS=darwin
set GOARCH=arm64
go build -o %OUT_DIR%\%BINARY_NAME%.darwin-arm64 %SRC_DIR%
if errorlevel 1 goto :error
goto :eof

REM Build for Freebsd ARM64
:build_freebsd_arm64
echo Building for FreeBSD ARM64...
set GOOS=freebsd
set GOARCH=arm64
go build -o %OUT_DIR%\%BINARY_NAME%.freebsd-arm64 %SRC_DIR%
if errorlevel 1 goto :error
goto :eof

REM Build for Freebsd AMD64
:build_freebsd_amd64
echo Building for FreeBSD AMD64...
set GOOS=freebsd
set GOARCH=amd64
go build -o %OUT_DIR%\%BINARY_NAME%.freebsd-amd64 %SRC_DIR%
if errorlevel 1 goto :error
goto :eof

REM Build all platforms
:build_all
echo Building for all platforms...

call :build_windows_amd64
call :build_windows_arm64
call :build_linux_amd64
call :build_linux_arm64
call :build_darwin_amd64
call :build_darwin_arm64
call :build_freebsd_arm64
call :build_freebsd_amd64

goto :eof

REM Error handling
:error
echo An error occurred during the build process.
exit /b 1

REM Usage information
:usage
echo Usage: %0 ^<command^>
echo Commands:
echo   build-windows-amd64     Build for Windows AMD64
echo   build-windows-arm64     Build for Windows ARM64
echo   build-linux-amd64       Build for Linux AMD64
echo   build-linux-arm64       Build for Linux ARM64
echo   build-darwin-amd64      Build for macOS AMD64
echo   build-darwin-arm64      Build for macOS ARM64
echo   build-all               Build for all platforms
echo   clean                   Clean the output directory
exit /b 1

:eof
endlocal
