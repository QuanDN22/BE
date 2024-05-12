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

