#!/bin/bash
set -e

cd "$(dirname "$0")"

echo "Running benchmarks..."
echo ""

go test -bench=. -benchtime=2s . 2>&1 | tee /tmp/bench_output.txt

echo ""
go run report.go < /tmp/bench_output.txt
