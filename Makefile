SOURCE := go.* *.go */*.go */*/*.go

tene3rm: $(SOURCE)
	go fmt ./...
	go vet ./...
	go test -cover ./...
	go build

cover.out: $(SOURCE)
	go test -cover ./... -coverprofile=$@

cover.html: cover.out
	go tool cover -html=cover.out -o $@

.PHONY: coverage
coverage: cover.html

.PHONY: clean
clean:
	@rm -f tene3rm
	@rm -f cover.out
	@rm -f cover.html
