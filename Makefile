GOPATH := $(HOME)/go
GOBIN  := $(GOPATH)/bin
GOMARKDOS ?= $(GOBIN)/gomarkdoc

$(GOMARKDOS): 
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest

.PHONY: docs 
docs: $(GOMARKDOS)
	mkdir -p docs/api-reference docs/guide && gomarkdoc -u --exclude-dirs ./examples/... --exclude-dirs ./internal/... --exclude-dirs ./cmd -o './docs/api-reference/{{.Dir}}.mdx' --template-file file=./docs/templates/file.gotxt  --template-file package=./docs/templates/package.gotxt ./...

.PHONY: format
format: 
	go fmt ./...

