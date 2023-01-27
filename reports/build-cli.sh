#!/bin/bash

rm -rf bin/*

# Generate for Darwin arm
env GOOS=darwin GOARCH=arm64 go build -o bin/report_generator_darwin_arm64 report_generator.go
env GOOS=darwin GOARCH=amd64 go build -o bin/report_generator_darwin_amd64 report_generator.go

# Generate for Linux AMD64
env GOOS=linux GOARCH=amd64 go build -o bin/report_generator_linux_amd64 report_generator.go
env GOOS=linux GOARCH=arm64 go build -o bin/report_generator_linux_arm64 report_generator.go