build:
	go build -v ./...
test:
	go test -v ./...
format:
	find  -name '*.go' | xargs gofmt -s -w
lint:
	@diff=$$(find  -name '*.go' | xargs gofmt -s -d); if [ -n "$$diff" ]; then \
		echo "Lint diff detected:"; \
		echo "$$diff"; \
		echo; \
		echo "Please run 'gofmt -s -d' on the affected files, or apply the patch above."; \
		exit 1; \
		else \
		echo "Lint OK"; \
		fi

.PHONY: build test format lint
