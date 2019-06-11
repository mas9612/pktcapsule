GOBIN := go
PROTOBIN := protoc

.PHONY: build-grpc
build-grpc:
	$(PROTOBIN) --go_out=plugins=grpc:. ./pktcapsule.proto

.PHONY: doc
doc:
	$(PROTOBIN) --doc_out=./doc --doc_opt=markdown,pktcapsule.md ./pktcapsule.proto
