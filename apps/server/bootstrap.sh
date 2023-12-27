#! /bin/bash

# go hot reload
go install github.com/cosmtrek/air@latest

# gen protobuf
brew install bufbuild/buf/buf

# protobuf code generation
go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# dependency injection
go install github.com/google/wire/cmd/wire@latest