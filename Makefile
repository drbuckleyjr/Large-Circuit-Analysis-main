# Large Circuit Analysis - Multi-Platform Build System
# Supports Windows, Linux, macOS on amd64 and arm64 architectures
# Optimized for Apple Silicon (M1/M2) with Homebrew support

# Ensure Homebrew is in PATH (especially important for Apple Silicon)
export PATH := /opt/homebrew/bin:/opt/homebrew/sbin:$(PATH)

# Application name
APP_NAME = large-circuit-analysis
VERSION = $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build directories
BUILD_DIR = build
DIST_DIR = dist

# Go build flags with architecture-specific optimizations
LDFLAGS = -ldflags="-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.Commit=$(COMMIT) -s -w"
BUILD_FLAGS = -trimpath $(LDFLAGS)

# Architecture-specific build tags
INTEL_MAC_FLAGS = -tags=intel_mac
APPLE_SILICON_FLAGS = -tags=apple_silicon

# Target platforms and architectures
PLATFORMS = \
	windows/amd64 \
	windows/arm64 \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	freebsd/amd64 \
	openbsd/amd64

# Default target
.PHONY: all
all: clean build-all

# Help target
.PHONY: help
help:
	@echo "Large Circuit Analysis - Multi-Platform Build System"
	@echo ""
	@echo "Available targets:"
	@echo "  all          - Clean and build for all platforms"
	@echo "  build-all    - Build for all platforms"
	@echo "  build-local  - Build for current platform only"
	@echo "  clean        - Remove build artifacts"
	@echo "  test         - Run tests"
	@echo "  deps         - Download dependencies"
	@echo "  dist         - Create distribution packages"
	@echo "  release      - Build and create release packages"
	@echo ""
	@echo "Platform-specific builds:"
	@echo "  windows      - Build for Windows (amd64 + arm64)"
	@echo "  linux        - Build for Linux (amd64 + arm64)"
	@echo "  darwin       - Build for macOS (amd64 + arm64)"
	@echo "  freebsd      - Build for FreeBSD (amd64)"
	@echo "  openbsd      - Build for OpenBSD (amd64)"
	@echo ""
	@echo "Individual platform builds:"
	@$(foreach platform,$(PLATFORMS), echo "  $(subst /,-,$(platform))     - Build for $(platform)";)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR) $(DIST_DIR)
	@echo "Clean complete."

# Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies updated."

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests complete."

# Build for current platform
.PHONY: build-local
build-local: deps
	@echo "Building for current platform..."
	@mkdir -p $(BUILD_DIR)
	@go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME) main.go
	@echo "Local build complete: $(BUILD_DIR)/$(APP_NAME)"

# Build for all platforms
.PHONY: build-all
build-all: deps $(PLATFORMS)

# Platform-specific targets
.PHONY: $(PLATFORMS)
$(PLATFORMS):
	@echo "Building for $@..."
	@mkdir -p $(BUILD_DIR)/$@
	@GOOS=$(word 1,$(subst /, ,$@)) GOARCH=$(word 2,$(subst /, ,$@)) \
		go build $(BUILD_FLAGS) \
		-o $(BUILD_DIR)/$@/$(APP_NAME)$(if $(findstring windows,$@),.exe,) \
		main.go
	@echo "Build complete: $(BUILD_DIR)/$@/$(APP_NAME)$(if $(findstring windows,$@),.exe,)"

# Platform group targets
.PHONY: windows linux darwin freebsd openbsd mac-pro
windows: windows/amd64 windows/arm64
linux: linux/amd64 linux/arm64
darwin: darwin/amd64 darwin/arm64
freebsd: freebsd/amd64
openbsd: openbsd/amd64

# Special target for Intel Mac Pro (2013 'cylinder' and similar)
mac-pro: darwin/amd64
	@echo "Built specifically for Intel Mac Pro compatibility"

# Individual platform convenience targets
.PHONY: windows-amd64 windows-arm64 linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 freebsd-amd64 openbsd-amd64
windows-amd64: windows/amd64
windows-arm64: windows/arm64
linux-amd64: linux/amd64
linux-arm64: linux/arm64
darwin-amd64: darwin/amd64
darwin-arm64: darwin/arm64
freebsd-amd64: freebsd/amd64
openbsd-amd64: openbsd/amd64

# Create distribution packages
.PHONY: dist
dist: build-all
	@echo "Creating distribution packages..."
	@mkdir -p $(DIST_DIR)
	@for platform in $(PLATFORMS); do \
		os=$$(echo $$platform | cut -d'/' -f1); \
		arch=$$(echo $$platform | cut -d'/' -f2); \
		if [ "$$os" = "windows" ]; then \
			ext=".exe"; \
		else \
			ext=""; \
		fi; \
		package_name="$(APP_NAME)-$(VERSION)-$$os-$$arch"; \
		echo "Creating package: $$package_name"; \
		mkdir -p $(DIST_DIR)/$$package_name; \
		cp $(BUILD_DIR)/$$platform/$(APP_NAME)$$ext $(DIST_DIR)/$$package_name/; \
		cp README.md $(DIST_DIR)/$$package_name/ 2>/dev/null || echo "README.md not found, skipping"; \
		cp LICENSE $(DIST_DIR)/$$package_name/ 2>/dev/null || echo "LICENSE not found, skipping"; \
		if [ "$$os" = "windows" ]; then \
			cd $(DIST_DIR) && zip -r $$package_name.zip $$package_name/; \
		else \
			cd $(DIST_DIR) && tar -czf $$package_name.tar.gz $$package_name/; \
		fi; \
		rm -rf $(DIST_DIR)/$$package_name; \
	done
	@echo "Distribution packages created in $(DIST_DIR)/"

# Release target
.PHONY: release
release: clean test dist
	@echo "Release packages ready in $(DIST_DIR)/"
	@ls -la $(DIST_DIR)/

# Development targets
.PHONY: run
run: build-local
	@echo "Running application..."
	@./$(BUILD_DIR)/$(APP_NAME)

# Install locally built binary to GOPATH/bin
.PHONY: install
install: build-local
	@echo "Installing to $(GOPATH)/bin/$(APP_NAME)..."
	@cp $(BUILD_DIR)/$(APP_NAME) $(GOPATH)/bin/$(APP_NAME)
	@echo "Installation complete."

# Environment check (supports both Apple Silicon and Intel Macs)
.PHONY: check-env
check-env:
	@echo "Environment Check:"
	@echo "  System: $(shell uname -s) $(shell uname -m)"
	@echo "  Hardware: $(shell system_profiler SPHardwareDataType | grep 'Model Name:' | head -1 | sed 's/.*Model Name: //' || echo 'Unknown')"
	@echo "  Go Path: $(shell which go)"
	@echo "  Go Version: $(shell go version)"
	@echo "  GOPATH: $(shell go env GOPATH)"
	@echo "  GOROOT: $(shell go env GOROOT)"
	@echo "  CGO Enabled: $(shell go env CGO_ENABLED)"
ifneq ($(shell uname -s),Darwin)
	@echo "  Note: Not running on macOS"
else
	@echo "  macOS Version: $(shell sw_vers -productVersion)"
ifeq ($(shell uname -m),arm64)
	@echo "  Architecture: Apple Silicon (ARM64)"
	@echo "  Chip: $(shell system_profiler SPHardwareDataType | grep 'Chip:' | head -1 | sed 's/.*Chip: //' || echo 'Apple Silicon')"
	@echo "  Homebrew: $(shell which brew 2>/dev/null || echo 'Not found') ($(shell brew --prefix 2>/dev/null || echo 'N/A'))"
	@echo "  Optimization: Apple Silicon build paths enabled"
else
	@echo "  Architecture: Intel x86_64"
	@echo "  Processor: $(shell system_profiler SPHardwareDataType | grep 'Processor Name:' | head -1 | sed 's/.*Processor Name: //' || echo 'Intel x86_64')"
	@echo "  Homebrew: $(shell which brew 2>/dev/null || echo 'Not found') ($(shell brew --prefix 2>/dev/null || echo 'N/A'))"
	@echo "  Optimization: Intel Mac build paths enabled"
	@echo "  Note: Building for Intel Mac Pro (2013 'cylinder') compatibility"
endif
endif

# Show build information
.PHONY: info
info: check-env
	@echo ""
	@echo "Build Information:"
	@echo "  App Name: $(APP_NAME)"
	@echo "  Version:  $(VERSION)"
	@echo "  Commit:   $(COMMIT)"
	@echo "  Build Time: $(BUILD_TIME)"
	@echo "  Go Version: $(shell go version)"
	@echo ""
	@echo "Target Platforms:"
	@$(foreach platform,$(PLATFORMS), echo "  $(platform)";)