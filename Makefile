'all:

golang:
	../../cinp/utils/cinp-codegen -l go -d go/autogen -s contractor http://127.0.0.1:8888/api/v1/
	cd go ; go fmt ./...
	cd go ; golint ./...
	cd go ; go vet ./...


PHONY: golang
