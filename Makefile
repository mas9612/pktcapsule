GOBIN := go
PROTOBIN := protoc
PKTCAPSULED := pktcapsuled

.PHONY: all
all: build-grpc test build doc

.PHONY: build
build:
	CGO_ENABLED=0 $(GOBIN) build -o $(PKTCAPSULED) ./cmd/pktcapsuled

.PHONY: build-grpc
build-grpc:
	$(PROTOBIN) --go_out=plugins=grpc:. ./pktcapsule.proto

.PHONY: doc
doc:
	$(PROTOBIN) --doc_out=./doc --doc_opt=markdown,pktcapsule.md ./pktcapsule.proto

.PHONY: test
test:
	$(GOBIN) test -v ./...
