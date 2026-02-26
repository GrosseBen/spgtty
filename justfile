# spgtty build system
# Usage:
#   just build    - Build the spgtty binary
#   just test     - Run Go tests
#   just try      - Try spgtty in the isolated try directory
#   just clean    - Clean up the try directory

build:
    @echo "🔨 Building spgtty..."
    @go build .

test:
    @echo "🧪 Running tests..."
    @go test -v ./...

try:
    @echo "🚀 Trying spgtty in isolated workspace..."
    @echo "📋 Version info:"
    /Users/ludal/src/github.com/GrosseBen/spgtty/spgtty -v
    @echo ""
    bash -c "cd try && /Users/ludal/src/github.com/GrosseBen/spgtty/spgtty init --device shellyplus1pm-demo && /Users/ludal/src/github.com/GrosseBen/spgtty/spgtty build && echo '📄 Build output:' && cat dist/main.js"

clean:
    @echo "🧹 Cleaning try directory..."
    rm -rf ./try/dist
    rm -rf ./try/.spgtty
    @echo "✅ Try directory cleaned!"
