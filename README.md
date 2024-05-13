# House Listings REST API

Project built with GoLang and Goa Framework

Steps to build project:

-   Initialize go.mod file: `go mod init listings.com`
-   Install goa framework: `go get -u goa.design/goa/v3`
-   Install goa commandline tool: `go install goa.design/goa/v3/cmd/goa@v3`
-   To handle gRPC, install the following:
    -   `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
    -   `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
-   To generate files run: `goa gen listings.com/design`
-   Install package for sqlite3 database setup: `go get -u github.com/mattn/go-sqlite3`
