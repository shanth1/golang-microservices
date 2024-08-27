gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/adder.proto

evans:
	@evans api/adder.proto -p 50051
