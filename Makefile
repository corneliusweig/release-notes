
release-notes: listpullreqs.go
	go build -ldflags "-s -w" -o $@

lint: listpullreqs_test.go
	hack/run-lint.sh

test: listpullreqs_test.go
	go test
