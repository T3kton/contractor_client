all:

golang:
	rm ns_*.go
	rm service.go
	../../cinp/utils/cinp-codegen -l go -d . -s contractor http://127.0.0.1:8888/api/v1/
	sed s/NewContractor/NewContractorInt/ -i service.go
	go fmt ./...
	golint ./...
	go vet ./...


PHONY: all golang
