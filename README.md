# gRPC-workshop
Based on the official gRPC workshop - https://codelabs.developers.google.com/codelabs/cloud-grpc/index.html

Uses Go server instead of Node.js server.

## Start
To generate the Go files from the proto file we need to use the following command:

`protoc -I . books/books.proto --go_out=plugins=grpc:.`

* `-I` indicates the path of the project the proto file is in (“.” means current directory, because we run it from the directory “start”).

* `--go_out=plugins=grpc:` indicates the path of the output. “.” means current directory. This is relative to the laction of the proto file. If the proto file is in books directory then the generated file will also be in the same directory if we use “.”.