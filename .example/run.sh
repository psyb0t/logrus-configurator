#!/bin/bash
export LOG_LEVEL="trace"
export LOG_FORMAT="text"
export LOG_CALLER="true"

go run main.go

export LOG_LEVEL="warn"
export LOG_FORMAT="json"
export LOG_CALLER="false"

go run main.go
