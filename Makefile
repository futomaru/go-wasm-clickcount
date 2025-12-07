.PHONY: help wasm clean serve

GO ?= go
WASM_TARGET := web/main.wasm

help:
	@printf "Available targets:\n"
	@printf "  wasm   Build Go/Wasm bundle for the counter.\n"
	@printf "  serve  Run the Go HTTP server (cmd/server).\n"
	@printf "  clean  Remove the generated wasm binary.\n"

wasm: $(WASM_TARGET)

$(WASM_TARGET): wasm/main.go
	GOOS=js GOARCH=wasm $(GO) build -o $(WASM_TARGET) ./wasm

serve:
	$(GO) run ./cmd/server

clean:
	rm -f $(WASM_TARGET)
