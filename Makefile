ENTRY_POINT=cmd/app

# Генерация контрактов
proto-generate:
	protoc --proto_path=proto \
		--go_out=gen/auth --go_opt=paths=source_relative \
		--go-grpc_out=gen/auth --go-grpc_opt=paths=source_relative \
		proto/auth.proto
proto-generate-file-processing:
	protoc --proto_path=proto \
           		--go_out=gen/fileProcessing --go_opt=paths=source_relative \
           		--go-grpc_out=gen/auth/fileProcessing --go-grpc_opt=paths=source_relative \
           		proto/file_processing.proto

# Запуск
run:
	go run $(ENTRY_POINT)/main.go
