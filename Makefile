
release-notes: listpullreqs.go
	go build -trimpath -tags netgo -ldflags "-s -w -extldflags '-static'" -o $@

lint: listpullreqs_test.go
	hack/run-lint.sh

test: listpullreqs_test.go
	go test
