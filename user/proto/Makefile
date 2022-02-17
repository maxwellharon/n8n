.PHONY: proto protoc mock

PROTOC_VERSION=1.42_1

proto: protoc mock

protoc:
	docker run --rm \
		-v `pwd`:/defs/proto \
		-w /defs \
		--entrypoint="" \
		namely/protoc-all:$(PROTOC_VERSION) \
		bash -c \
		"protoc --proto_path=/defs --go_out=paths=source_relative,plugins=grpc:. --doc_out . --doc_opt=markdown,PROTO.md /defs/proto/user.proto"

mock:
	mockgen -source user.pb.go -destination mock/user.mock.pb.go -package mock_proto
