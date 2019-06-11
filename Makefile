GOBIN := go
PROTOBIN := protoc
PKTCAPSULED := pktcapsuled

.PHONY: all
all: build-grpc build doc

.PHONY: build
build:
	$(GOBIN) build -o $(PKTCAPSULED) ./cmd/pktcapsuled

.PHONY: build-grpc
build-grpc:
	$(PROTOBIN) --go_out=plugins=grpc:. ./pktcapsule.proto

.PHONY: doc
doc:
	$(PROTOBIN) --doc_out=./doc --doc_opt=markdown,pktcapsule.md ./pktcapsule.proto

.PHONY: test
test:
	$(GOBIN) test -v ./...
