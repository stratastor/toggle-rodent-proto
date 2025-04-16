# Toggle-Rodent Protocol

Shared protocol definitions for communication between Toggle and Rodent services.

## Usage

### Import the package

```go
import "github.com/stratastor/toggle-rodent-proto/proto"
```

### Using in your Go code

```go
client := proto.NewRodentServiceClient(conn)
```

## Development

### Regenerating protobuf files

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/rodent.proto
```