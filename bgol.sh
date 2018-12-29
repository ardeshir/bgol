#!/bin/bash
# usage:
# ./bgol.sh main.zip 

DEP=""

if [[ -z $1 ]]; then
DEP=deployment.zip 
else 
DEP=$1
fi


echo "Building the go binary..."
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Creating ZIP file..."
zip $DEP  main 

echo "aws s3 cp deployment.zip s3://cryptogoserverless"

aws s3 cp $DEP  s3://cryptogoserverless

sleep 5

echo "Clearning up..."
rm main

