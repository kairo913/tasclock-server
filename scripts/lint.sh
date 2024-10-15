#!/bin/bash
set -euo pipefail

main() {
    echo golangci-lint
    go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 run

    return 0
}

main "$@"
