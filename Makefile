install:
	go build -o protoc-gen-errorbuilder cmd/protoc-gen-errorbuilder/main.go
	cp protoc-gen-errorbuilder $(GOPATH)/bin

uninstall:
	rm -f $(GOPATH)/bin/protoc-gen-errorbuilder