generate-sorm-structs:
	protoc -I proto proto/sorm.proto --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative
