proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/pb/user/user.proto pkg/pb/auth/auth.proto pkg/pb/book/book.proto

server:
	go run cmd/main.go

	