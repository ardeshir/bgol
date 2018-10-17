#!/bin/bash

echo "Building the go binary..."
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Creating ZIP file..."
zip deployment.zip main 

echo "Clearning up..."
rm main

