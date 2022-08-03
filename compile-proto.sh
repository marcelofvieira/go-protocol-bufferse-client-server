protoc -I greet/proto --go_out=. --go_opt=module=github.com/marcelofvieira/protocol-buffer3 --go-grpc_out=. --go-grpc_opt=module=github.com/marcelofvieira/protocol-buffer3 greet/proto/*.proto
# or
# protoc --proto_path=greet --go_out=greet --go_opt=paths=source_relative --go-grpc_out=greet --go-grpc_opt=paths=source_relative greet/proto/dummy.proto