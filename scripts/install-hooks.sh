#!/bin/sh

HOOK_DIR=$(git rev-parse --show-toplevel)/.git/hooks

cat > "$HOOK_DIR/pre-commit" << 'EOF'
#!/bin/sh

echo "Running golines..."
if command -v golines > /dev/null 2>&1; then
    find . -name "*.go" -exec golines -m 180 -w {} \;
fi

echo "Running golangci-lint..."

if ! command -v golangci-lint > /dev/null 2>&1; then
    echo "Error: golangci-lint is not installed."
    echo "Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
    exit 1
fi

golangci-lint run ./...
if [ $? -ne 0 ]; then
    echo ""
    echo "Lint errors found. Please fix them before committing."
    exit 1
fi

echo "Lint passed."
exit 0
EOF

chmod +x "$HOOK_DIR/pre-commit"
echo "Pre-commit hook installed."
