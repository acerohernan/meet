default: watch

dev:
	go run cmd/server/main.go --config config.yaml

watch:
	air

protoc:
	buf generate ../../packages/proto

wire:
	cd pkg/service && wire

gen:
	go generate ./...

test:
	go test -v ./...