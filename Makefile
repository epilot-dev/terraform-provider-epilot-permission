.PHONY: all
all: docs

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s permission.yaml

docs: speakeasy
	go generate ./...

