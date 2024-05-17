# init-go-module:
# 	go mod init github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt

gen:
	protoc -I ./proto \
   	--go_out ./proto --go_opt paths=source_relative \
   	--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
   	./proto/auth/auth.proto ./proto/bio/bio.proto

gen-key:
	openssl genpkey -algorithm ED25519 -outform pem -out auth.ed
	openssl pkey -in auth.ed -pubout > auth.ed.pub

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
t := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTU5NzU5MzEsImlhdCI6MTcxNTk3NTg3MSwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAwIiwibmJmIjoxNzE1OTc1ODcxLCJyb2xlcyI6WyJhZG1pbiIsImJhc2ljIl0sInVzZXIiOiJhZG1pbiJ9.jPlIbWF1KFIrm9iBPVF1LKowNxyA0_batRDg9-dBgj8lX2Vg0h92KLSboc3j4gLkmQYUvLgk4HAiDEVCiMi5Aw

login-auth-api:
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



