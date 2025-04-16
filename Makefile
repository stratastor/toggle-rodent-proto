.PHONY: all generate clean

all: generate

# Generate protobuf files
generate:
	@echo "Generating protobuf files..."
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/rodent.proto
	@echo "Done"

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f proto/*.pb.go
	@echo "Done"

# Check if all dependencies are installed
check-deps:
	@which protoc > /dev/null || (echo "protoc is not installed. Please install it first." && exit 1)
	@which protoc-gen-go > /dev/null || (echo "protoc-gen-go is not installed. Run: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest" && exit 1)
	@which protoc-gen-go-grpc > /dev/null || (echo "protoc-gen-go-grpc is not installed. Run: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest" && exit 1)
	@echo "All dependencies are installed"