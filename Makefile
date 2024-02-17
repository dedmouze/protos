.PHONY: gen

gen:
	make gen_main
	make gen_gateway

gen_main:
	protoc -I proto proto/sso/*.proto \
	--go_out=./gen/go/ \
	--go_opt=paths=source_relative \
	--go-grpc_out=./gen/go/ \
	--go-grpc_opt=paths=source_relative

gen_gateway:
	protoc -I proto proto/sso/*.proto \
	--grpc-gateway_out ./gen/go \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true