#!/bin/bash
cd "$(dirname "$0")"
go run -v -tags gtk_3_18 -gcflags "-N -l" ./app/main.go
