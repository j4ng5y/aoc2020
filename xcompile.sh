#!/bin/bash

if [[ $1 == "" ]]; then
    echo "Invalid argument: An argument is required.";
    echo "Usage: $0 1";
    exit 1;
else
    echo "Cross Compiling Binaries...";
    GOOS=darwin GOARCH=amd64 go build -o day$1/bin/day$1_darwin_x86-64 day$1/main.go;
    GOOS=linux GOARCH=amd64 go build -o day$1/bin/day$1_linux_x86-64 day$1/main.go;
    GOOS=windows GOARCH=amd64 go build -o day$1/bin/day$1_windows_x86-64.exe day$1/main.go;
    echo "Done";
    exit 0;
fi