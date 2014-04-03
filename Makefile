.PHONY: test

test: test/test
	@/bin/bash test/test.sh

test/test: terminal.go test/cmd.go
	cd test; go build
