#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -o medco-loader *.go