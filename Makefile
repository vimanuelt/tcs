# Makefile for TCS

BINARY_NAME=tcs
BUILD_DIR=build

.PHONY: all build install uninstall clean

# Default target
all: build

# Build the application
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Install the binary to /usr/local/bin
install: build
	@install -m 0755 $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to /usr/local/bin."

# Uninstall the binary
uninstall:
	@rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "Uninstalled $(BINARY_NAME) from /usr/local/bin."

# Clean up build artifacts
clean:
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned up build artifacts."

