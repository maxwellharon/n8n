.PHONY: proto protoc mock

PROTOC_VERSION=1.42_1

proto: protoc mock

# The protobuf compiler *really* wants to be invoked from the repository root.
# See: https://developers.google.com/protocol-buffers/docs/proto3#importing-definitions
# Also, make sure to install protoc-gen-doc with:
# go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
protoc:
	cd ../../ && \
	docker run \
		-v `pwd`/common/proto:/defs/common/proto \
		--entrypoint="" \
		namely/protoc-all:$(PROTOC_VERSION) \
		bash -c \
		"protoc --proto_path=/defs -I /opt/include --go_out=paths=source_relative,plugins=grpc:. --doc_out /defs/common/proto --doc_opt=markdown,PROTO.md /defs/common/proto/*.proto"

mock:
	mockgen -source health.pb.go -destination mock/health.mock.pb.go -package mock_proto
