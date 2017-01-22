.PHONY: fmt test

fmt:
	@gofmt -w .
	@goimports -w .

test:
	@go test -v -race

