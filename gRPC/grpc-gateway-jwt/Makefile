init-go-module:
	go mod init github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt

gen:
	protoc -I ./proto \
   	--go_out ./proto --go_opt paths=source_relative \
   	--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
   	./proto/auth/auth.proto ./proto/bio/bio.proto

gen-key:
	openssl genpkey -algorithm ED25519 -outform pem -out auth.ed
	openssl pkey -in auth.ed -pubout > auth.ed.pub

auth_server:
	go run ./cmd/auth_server/main.go auth.ed.pub

test:
	go run ./cmd/jwt-issue/main.go auth.ed auth.ed.pub


# run server
run_0-auth-api:
	go run .\cmd\0-auth-api\main.go

run_auth_server:
	go run .\cmd\auth_server\main.go auth.ed.pub

run_bio_server:
	go run .\cmd\bio_server\main.go auth.ed.pub

run_grpc_gateway:
	go run .\cmd\grpc-gateway\main.go auth.ed.pub  

# api
t := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTU5NzQzNTksImlhdCI6MTcxNTk3NDI5OSwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAwIiwibmJmIjoxNzE1OTc0Mjk5LCJyb2xlcyI6WyJhZG1pbiIsImJhc2ljIl0sInVzZXIiOiJhZG1pbiJ9.3jBt8fBJrp9cvgKoN_sd0m3gbqD2xumbIQ1DtnP_410IDc3DygVV1AjSshn17RCmloRbM_4Rq_hz0buH9yySDA

login-gen-token:
	curl -s admin:pass@localhost:5000/login-auth-api

hello:
	curl -s -H "Authorization: Bearer $t" \
	"Content-Type: application/json" -X POST localhost:8090/hello -d "{\"name\":\"hello \"}" \

ping:
	curl -s -H "Authorization: Bearer $t" \
	localhost:8090/ping

hi:
	curl -s -H "Authorization: Bearer $t" \
	localhost:8090/hi



