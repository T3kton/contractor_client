all:

golang:
	rm go/ns_*.go
	rm go/service.go
	../../cinp/utils/cinp-codegen -l go -d go -s contractor http://127.0.0.1:8888/api/v1/
	sed s/NewContractor/NewContractorInt/ -i go/service.go
	cd go ; go fmt ./...
	cd go ; golint ./...
	cd go ; go vet ./...


PHONY: all golang
