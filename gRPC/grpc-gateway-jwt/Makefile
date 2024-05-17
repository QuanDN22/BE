init-go-module:
	go mod init github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt

gen:
	protoc -I ./proto \
   	--go_out ./proto --go_opt paths=source_relative \
   	--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
   	./proto/auth/auth.proto ./proto/bio/bio.proto

auth_server:
	go run ./cmd/auth_server/main.go auth.ed.pub