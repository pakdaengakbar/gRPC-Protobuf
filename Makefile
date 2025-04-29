# Define the proto file location and the output directories
PROTO_SRC = proto
PROTO_OUT = ./generated
PROTO_BIN = $(shell pwd)/bin

# Define the protoc command and plugins
PROTOC = protoc
PROTOC_GEN_GO = $(PROTO_BIN)/protoc-gen-go
PROTOC_GEN_GRPC_GATEWAY = $(PROTO_BIN)/protoc-gen-grpc-gateway

# Check if protoc-gen-go is installed
check_protoc_gen_go:
	@if ! command -v $(PROTOC_GEN_GO) &>/dev/null; then \
		echo "protoc-gen-go not found. Please install it first."; \
		exit 1; \
	fi

# Generate Go code from protobuf files
proto:
	@$(MAKE) check_protoc_gen_go
	@echo "Generating Go code from protobuf..."
	@mkdir -p $(PROTO_OUT)
	@for file in $(PROTO_SRC)/*.proto; do \
		$(PROTOC) --proto_path=$(PROTO_SRC) \
			--go_out=$(PROTO_OUT) \
			--go-grpc_out=$(PROTO_OUT) $$file; \
	done
	@echo "Done generating Go code."

.PHONY: proto check_protoc_gen_go
