build:
    @echo "Building project..."
    @go build .
test:
    @echo "Running tests..."
    @go test -v ./...
try:
    cd ./try
    $(pwd)/spgtty init --device shellyplus1pm-demo
    $(pwd)/spgtty build

    echo "Build output:"
    cat ./dist/main.js
clean:
    rm -rf ./try/dist
    rm -rd .spgtty
    rm -rf ./try/.spgtty
