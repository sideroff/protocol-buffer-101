make sure you have:
1. make
1. go
1. protoc
1. protoc-gen-go & protoc-gen-go-grpc installed and in PATH
1. a folder named protocgen in root
1. grpcurl cli tool

run:
1. make protos
1. go run main.go
1. grpcurl --plaintext -d '{"Base": "GBP", "Destination": "BGN"}' localhost:3000 Currency.GetRate