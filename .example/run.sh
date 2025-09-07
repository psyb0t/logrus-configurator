#!/bin/bash

echo "=========================================="
echo "=== BASIC EXAMPLE (trace level, text) ==="
echo "=========================================="
export LOG_LEVEL="trace"
export LOG_FORMAT="text"
export LOG_CALLER="true"

go run basic.go

echo ""
echo "=========================================="
echo "=== BASIC EXAMPLE (warn level, json) ===="
echo "=========================================="
export LOG_LEVEL="warn"
export LOG_FORMAT="json"
export LOG_CALLER="false"

go run basic.go

echo ""
echo "=========================================="
echo "=== ADD HOOK EXAMPLE ===================="
echo "=========================================="
export LOG_LEVEL="info"
export LOG_FORMAT="text"
export LOG_CALLER="false"

go run addhook.go

echo ""
echo "=========================================="
echo "=== SET HOOKS EXAMPLE ==================="
echo "=========================================="
export LOG_LEVEL="debug"
export LOG_FORMAT="json"
export LOG_CALLER="false"

go run sethooks.go
