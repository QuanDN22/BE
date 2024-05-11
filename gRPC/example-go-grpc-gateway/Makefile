
# Syntax Compiling Protocol Buffers: 
# protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
protoc:
	protoc -I=protos --go_out=./protogen/golang --go_opt=paths=source_relative --go-grpc_out=./protogen/golang --go-grpc_opt=paths=source_relative --grpc-gateway_out=./protogen/golang --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true orders/order.proto &
	protoc -I=protos --go_out=./protogen/golang --go_opt=paths=source_relative product/product.proto
