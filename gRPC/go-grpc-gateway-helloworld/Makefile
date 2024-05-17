SWAGGER_UI_VERSION:=v4.15.5

gen:	
	protoc -I ./proto \
   	--go_out ./proto --go_opt paths=source_relative \
   	--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
	./proto --openapiv2_out ./gen/openapiv2 \
   	./proto/helloworld/hello_world.proto

gen-swagger:
	protoc -I ./proto --openapiv2_out ./gen/openapiv2 \
   	./proto/helloworld/hello_world.proto 

gen-key:
	openssl genpkey -algorithm ED25519 -outform pem -out auth.ed
	openssl pkey -in auth.ed -pubout > auth.ed.pub
	
token:
	go run ./cmd/jwt-issue auth.ed 

validate:
	go run ./cmd/jwt-validate auth.ed.pub 
all:
	go run ./cmd/jwt-issue auth.ed auth.ed.pub

auth-api:
	go run ./cmd/0-auth-api auth.ed & sleep 1
	t=$(curl -s admin:pass@localhost:8081/login)
	echo "token: $t"


run-server:
	go run ./cmd/server/main.go auth.ed.pub

run-grpc-gateway:
	go run ./cmd/grpc-gateway/main.go auth.ed.pub


