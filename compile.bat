REM This file compile this lambda for Linux environment.

echo "Starting to set environment variables..."
set GOOS=linux
set GOARCH=amd64

go build main.go
del main.zip
tar.exe -a -cf main.zip main

echo "Compilation process was done."