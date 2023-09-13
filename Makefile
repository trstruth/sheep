protoc:
	protoc --go_out=pkg/proto --go-grpc_out=pkg/proto proto/*.proto