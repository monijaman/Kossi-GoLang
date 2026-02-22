//go:build ignore
// +build ignore

// This file is intentionally ignored by the Go build system.
// It previously was empty and caused `go build ./...` to fail with
// "expected 'package', found 'EOF'". Keeping it here for historical
// reference; the build tag prevents it from being compiled.

package main

func main() {}
