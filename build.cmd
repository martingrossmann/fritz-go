@echo off

set GOOS=linux
set GOARCH=arm
go build

set GOOS=windows
set GOARCH=amd64
go build