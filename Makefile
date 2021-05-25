.PHONY: protos

protos:
	protoc -I ./protos/ ./protos/currency.proto --go-grpc_out=./protocgen --go_out=./protocgen