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
# t := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYyOTU4MjMsImlhdCI6MTcxNjA3OTgyMywiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAwIiwibmJmIjoxNzE2MDc5ODIzLCJyb2xlcyI6WyJhZG1pbiJdLCJ1c2VyIjoiYWRtaW4ifQ.KiDX753gWw1y15k64jGY4CvjQUBNZbJuksY84AT9cRd-odu4tlIiNsmCD09IjMmPLxZhWFB98d74dfKERZzZCA
# t:= eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYyOTczNTAsImlhdCI6MTcxNjA4MTM1MCwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAwIiwibmJmIjoxNzE2MDgxMzUwLCJyb2xlcyI6WyJhZG1pbiJdLCJ1c2VyIjoiYWRtaW4ifQ.kmFoeF_V4LqXH-DiLDKM4X9vk1pKD8yqJ-BuxG4rRBko5xK18s6_0qKJsUQhWTIhX43JLswpay3WLNSVjIUvBw
t:=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYyOTkzNTMsImlhdCI6MTcxNjA4MzM1MywiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAwIiwibmJmIjoxNzE2MDgzMzUzLCJyb2xlcyI6WyJhZG1pbiJdLCJ1c2VyIjoiYWRtaW4ifQ._WkugR79vcligDWebOyXnEyHKEKEFzrxNip84Egdbo-HkNGvcKiMplUYWREQPb6BO13dsj2yfB92FJLXtRRRCw
login-auth-api:
	curl -s admin:pass@localhost:5000/login-auth-api



login: 
	curl -H "Content-Type: application/json" -X POST http://localhost:8090/login \
	-d "{\"username\":\"admin\", \"password\":\"pass\"}"

# login:
# 	curl -s admin:pass http://localhost:8090/login
# # curl -s admin:pass@localhost:8090/login

hello:
	curl -s -H "Authorization: Bearer $t" \
	"Content-Type: application/json" -X POST localhost:8090/hello -d "{\"name\":\"hello \"}" \

ping:
	curl -s -H "Authorization: Bearer $t" \
	localhost:8090/ping

hi:
	curl -s -H "Authorization: Bearer $t" \
	localhost:8090/hi



