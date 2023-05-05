.PHONY: all docs
all: docs speakeasy

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s permission.yaml

docs:
	go generate ./...

