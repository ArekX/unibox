#!/bin/bash
cd "$(dirname "$0")"
echo "Building unibox..."
go build -o unibox -v -tags gtk_3_18 -gcflags "-N -l" ./app/main.go
