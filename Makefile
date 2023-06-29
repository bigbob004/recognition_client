generate-client:
	protoc --go_out=clients/pb --go-grpc_out=clients/pb api/clients/*.proto