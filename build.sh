#!/bin/bash

# Large Circuit Analysis - Multi-Platform Build Script
# Cross-compilation build script for Windows, Linux, macOS, and BSD systems
# Optimized for Apple Silicon (M1/M2) with Homebrew support

set -e

# Ensure Homebrew is in PATH (handles both Apple Silicon and Intel Macs)
# Apple Silicon Macs (M1, M2, M4, etc.) use /opt/homebrew
# Intel Macs (including Mac Pro 2013 'cylinder') use /usr/local
if [[ -f "/opt/homebrew/bin/brew" ]]; then
    export PATH="/opt/homebrew/bin:/opt/homebrew/sbin:$PATH"
    export HOMEBREW_PREFIX="/opt/homebrew"
    export MAC_ARCH="apple_silicon"
elif [[ -f "/usr/local/bin/brew" ]]; then
    export PATH="/usr/local/bin:$PATH"
    export HOMEBREW_PREFIX="/usr/local"
    export MAC_ARCH="intel"
else
    # Try to detect if we're on macOS without Homebrew
    if [[ "$(uname -s)" == "Darwin" ]]; then
        echo "Warning: Homebrew not found. Install with:"
        if [[ "$(uname -m)" == "arm64" ]]; then
            echo "  /bin/bash -c \"\$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
        else
            echo "  /bin/bash -c \"\$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
        fi
    fi
fi

# Verify Go installation
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    echo "On Apple Silicon Mac, install with: brew install go"
    exit 1
fi

# Configuration
APP_NAME="large-circuit-analysis"
BUILD_DIR="build"
DIST_DIR="dist"

# Get version info
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build flags
LDFLAGS="-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.Commit=${COMMIT} -s -w"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Platform definitions
declare -A PLATFORMS=(
    ["windows/amd64"]="windows amd64 .exe"
    ["windows/arm64"]="windows arm64 .exe"
    ["linux/amd64"]="linux amd64 "
    ["linux/arm64"]="linux arm64 "
    ["darwin/amd64"]="darwin amd64 "
    ["darwin/arm64"]="darwin arm64 "
    ["freebsd/amd64"]="freebsd amd64 "
    ["openbsd/amd64"]="openbsd amd64 "
)

# Functions
print_header() {
    echo -e "${BLUE}=================================${NC}"
    echo -e "${BLUE} Large Circuit Analysis Builder ${NC}"
    echo -e "${BLUE}=================================${NC}"
    echo ""
}

print_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

show_usage() {
    echo "Usage: $0 [OPTIONS] [TARGETS]"
    echo ""
    echo "OPTIONS:"
    echo "  -h, --help      Show this help message"
    echo "  -c, --clean     Clean build artifacts before building"
    echo "  -d, --dist      Create distribution packages"
    echo "  -t, --test      Run tests before building"
    echo "  -v, --verbose   Verbose output"
    echo ""
    echo "TARGETS:"
    echo "  all             Build for all platforms (default)"
    echo "  local           Build for current platform only"
    echo "  windows         Build for Windows (amd64 + arm64)"
    echo "  linux           Build for Linux (amd64 + arm64)"
    echo "  darwin          Build for macOS (amd64 + arm64)"
    echo "  freebsd         Build for FreeBSD (amd64)"
    echo "  openbsd         Build for OpenBSD (amd64)"
    echo ""
    echo "Individual platforms:"
    for platform in "${!PLATFORMS[@]}"; do
        echo "  ${platform//\//-}    Build for $platform"
    done
    echo ""
    echo "Examples:"
    echo "  $0                    # Build all platforms"
    echo "  $0 --clean --dist     # Clean, build all, create packages"
    echo "  $0 linux windows      # Build only Linux and Windows"
    echo "  $0 --test local       # Test and build for current platform"
}

clean_builds() {
    print_info "Cleaning build artifacts..."
    rm -rf "$BUILD_DIR" "$DIST_DIR"
    print_info "Clean complete."
}

update_deps() {
    print_info "Updating dependencies..."
    go mod download
    go mod tidy
    print_info "Dependencies updated."
}

run_tests() {
    print_info "Running tests..."
    if go test -v ./...; then
        print_info "All tests passed."
    else
        print_error "Tests failed!"
        exit 1
    fi
}

build_platform() {
    local platform_key="$1"
    local platform_info="${PLATFORMS[$platform_key]}"
    
    if [ -z "$platform_info" ]; then
        print_error "Unknown platform: $platform_key"
        return 1
    fi
    
    read -r goos goarch ext <<< "$platform_info"
    local output_dir="$BUILD_DIR/$platform_key"
    local binary_name="$APP_NAME$ext"
    
    print_info "Building for $platform_key ($goos/$goarch)..."
    
    mkdir -p "$output_dir"
    
    if [ "$VERBOSE" = true ]; then
        GOOS="$goos" GOARCH="$goarch" go build -v \
            -trimpath \
            -ldflags="$LDFLAGS" \
            -o "$output_dir/$binary_name" .
    else
        GOOS="$goos" GOARCH="$goarch" go build \
            -trimpath \
            -ldflags="$LDFLAGS" \
            -o "$output_dir/$binary_name" . 2>/dev/null
    fi
    
    if [ $? -eq 0 ]; then
        local size=$(du -h "$output_dir/$binary_name" | cut -f1)
        print_info "✓ Built $platform_key ($size): $output_dir/$binary_name"
    else
        print_error "✗ Failed to build $platform_key"
        return 1
    fi
}

build_local() {
    print_info "Building for current platform..."
    mkdir -p "$BUILD_DIR"
    
    if go build -trimpath -ldflags="$LDFLAGS" -o "$BUILD_DIR/$APP_NAME" .; then
        local size=$(du -h "$BUILD_DIR/$APP_NAME" | cut -f1)
        print_info "✓ Local build complete ($size): $BUILD_DIR/$APP_NAME"
    else
        print_error "✗ Local build failed"
        exit 1
    fi
}

create_distributions() {
    print_info "Creating distribution packages..."
    mkdir -p "$DIST_DIR"
    
    for platform in "${!PLATFORMS[@]}"; do
        local platform_info="${PLATFORMS[$platform]}"
        read -r goos goarch ext <<< "$platform_info"
        
        local package_name="$APP_NAME-$VERSION-$goos-$goarch"
        local temp_dir="$DIST_DIR/$package_name"
        local binary_path="$BUILD_DIR/$platform/$APP_NAME$ext"
        
        if [ ! -f "$binary_path" ]; then
            print_warning "Binary not found for $platform, skipping..."
            continue
        fi
        
        print_info "Creating package: $package_name"
        
        mkdir -p "$temp_dir"
        cp "$binary_path" "$temp_dir/"
        
        # Copy additional files if they exist
        [ -f "README.md" ] && cp "README.md" "$temp_dir/"
        [ -f "LICENSE" ] && cp "LICENSE" "$temp_dir/"
        [ -f "CHANGELOG.md" ] && cp "CHANGELOG.md" "$temp_dir/"
        
        # Create archive based on platform
        cd "$DIST_DIR"
        if [ "$goos" = "windows" ]; then
            zip -r "$package_name.zip" "$package_name/" > /dev/null
            rm -rf "$package_name"
            print_info "✓ Created $package_name.zip"
        else
            tar -czf "$package_name.tar.gz" "$package_name/" 2>/dev/null
            rm -rf "$package_name"
            print_info "✓ Created $package_name.tar.gz"
        fi
        cd - > /dev/null
    done
}

show_build_info() {
    echo ""
    echo -e "${BLUE}Build Information:${NC}"
    echo "  App Name:     $APP_NAME"
    echo "  Version:      $VERSION"
    echo "  Commit:       $COMMIT"
    echo "  Build Time:   $BUILD_TIME"
    echo "  Go Version:   $(go version | cut -d' ' -f3-4)"
    echo ""
}

# Parse command line arguments
CLEAN=false
DIST=false
TEST=false
VERBOSE=false
TARGETS=()

while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_usage
            exit 0
            ;;
        -c|--clean)
            CLEAN=true
            shift
            ;;
        -d|--dist)
            DIST=true
            shift
            ;;
        -t|--test)
            TEST=true
            shift
            ;;
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        -*)
            print_error "Unknown option: $1"
            show_usage
            exit 1
            ;;
        *)
            TARGETS+=("$1")
            shift
            ;;
    esac
done

# Default to 'all' if no targets specified
if [ ${#TARGETS[@]} -eq 0 ]; then
    TARGETS=("all")
fi

# Main execution
print_header
show_build_info

# Clean if requested
if [ "$CLEAN" = true ]; then
    clean_builds
fi

# Update dependencies
update_deps

# Run tests if requested
if [ "$TEST" = true ]; then
    run_tests
fi

# Build targets
for target in "${TARGETS[@]}"; do
    case $target in
        all)
            print_info "Building all platforms..."
            failed_builds=()
            for platform in "${!PLATFORMS[@]}"; do
                if ! build_platform "$platform"; then
                    failed_builds+=("$platform")
                fi
            done
            
            if [ ${#failed_builds[@]} -gt 0 ]; then
                print_warning "Some builds failed: ${failed_builds[*]}"
            else
                print_info "All builds completed successfully!"
            fi
            ;;
        local)
            build_local
            ;;
        windows)
            build_platform "windows/amd64"
            build_platform "windows/arm64"
            ;;
        linux)
            build_platform "linux/amd64"
            build_platform "linux/arm64"
            ;;
        darwin)
            build_platform "darwin/amd64"
            build_platform "darwin/arm64"
            ;;
        freebsd)
            build_platform "freebsd/amd64"
            ;;
        openbsd)
            build_platform "openbsd/amd64"
            ;;
        *)
            # Check if it's a specific platform
            platform_key="${target//-//}"
            if [[ -n "${PLATFORMS[$platform_key]}" ]]; then
                build_platform "$platform_key"
            else
                print_error "Unknown target: $target"
                exit 1
            fi
            ;;
    esac
done

# Create distributions if requested
if [ "$DIST" = true ]; then
    create_distributions
    echo ""
    print_info "Distribution packages:"
    ls -lh "$DIST_DIR"/*.{zip,tar.gz} 2>/dev/null || print_warning "No distribution packages found"
fi

echo ""
print_info "Build process complete!"