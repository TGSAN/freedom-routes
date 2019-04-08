@echo off
set BIN_NAME=freedom-routes
set BIN_VERSION=1.2.0
set GOOS=linux
set GOARCH=amd64
echo building %GOOS%/%GOARCH%
xcopy routes\templates dist\%BIN_NAME%.%GOOS%.%GOARCH%\ /E
xcopy freedom-routes.etc dist\%BIN_NAME%.%GOOS%.%GOARCH%\
xcopy README.md dist\%BIN_NAME%.%GOOS%.%GOARCH%\
go build -o dist\%BIN_NAME%.%GOOS%.%GOARCH%\%BIN_NAME%.%GOOS%.%GOARCH%-%BIN_VERSION%
set GOOS=linux
set GOARCH=386
echo building %GOOS%/%GOARCH%
xcopy routes\templates dist\%BIN_NAME%.%GOOS%.%GOARCH%\ /E
xcopy freedom-routes.etc dist\%BIN_NAME%.%GOOS%.%GOARCH%\
xcopy README.md dist\%BIN_NAME%.%GOOS%.%GOARCH%\
go build -o dist\%BIN_NAME%.%GOOS%.%GOARCH%\%BIN_NAME%.%GOOS%.%GOARCH%-%BIN_VERSION%
set GOOS=homebrew
set GOARCH=amd64
echo building %GOOS%/%GOARCH%
xcopy routes\templates dist\%BIN_NAME%.%GOOS%.%GOARCH%\ /E
xcopy freedom-routes.etc dist\%BIN_NAME%.%GOOS%.%GOARCH%\
xcopy README.md dist\%BIN_NAME%.%GOOS%.%GOARCH%\
go build -o dist\%BIN_NAME%.%GOOS%.%GOARCH%\%BIN_NAME%.%GOOS%.%GOARCH%-%BIN_VERSION%
set GOOS=homebrew
set GOARCH=386
echo building %GOOS%/%GOARCH%
xcopy routes\templates dist\%BIN_NAME%.%GOOS%.%GOARCH%\ /E
xcopy freedom-routes.etc dist\%BIN_NAME%.%GOOS%.%GOARCH%\
xcopy README.md dist\%BIN_NAME%.%GOOS%.%GOARCH%\
go build -o dist\%BIN_NAME%.%GOOS%.%GOARCH%\%BIN_NAME%.%GOOS%.%GOARCH%-%BIN_VERSION%
set GOOS=windows
set GOARCH=amd64
echo building %GOOS%/%GOARCH%
xcopy routes\templates dist\%BIN_NAME%.%GOOS%.%GOARCH%\ /E
xcopy freedom-routes.etc dist\%BIN_NAME%.%GOOS%.%GOARCH%\
xcopy README.md dist\%BIN_NAME%.%GOOS%.%GOARCH%\
go build -o dist\%BIN_NAME%.%GOOS%.%GOARCH%\%BIN_NAME%.%GOOS%.%GOARCH%-%BIN_VERSION%.exe
set GOOS=windows
set GOARCH=386
echo building %GOOS%/%GOARCH%
xcopy routes\templates dist\%BIN_NAME%.%GOOS%.%GOARCH%\ /E
xcopy freedom-routes.etc dist\%BIN_NAME%.%GOOS%.%GOARCH%\
xcopy README.md dist\%BIN_NAME%.%GOOS%.%GOARCH%\
go build -o dist\%BIN_NAME%.%GOOS%.%GOARCH%\%BIN_NAME%.%GOOS%.%GOARCH%-%BIN_VERSION%.exe
pause
