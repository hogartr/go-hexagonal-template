.PHONY: dev proto sqlc test clean

dev: proto sqlc
	go run ./cmd/server

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       proto/user.proto

sqlc:
	sqlc generate

test:
	go test ./... -v

# clean:
# 	rm -rf pb
