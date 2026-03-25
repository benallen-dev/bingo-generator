#!/usr/bin/env zsh

go run ./cmd/main.go
./scripts/collate-all.sh
echo "Generation complete"
