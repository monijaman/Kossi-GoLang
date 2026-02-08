#!/bin/bash

# Product Translation Tool Runner
# This script makes it easy to run the product translator

cd "$(dirname "$0")" || exit 1
cd ../../ || exit 1

# Check if .env exists
if [ ! -f ".env" ]; then
    echo "❌ Error: .env file not found"
    echo "Please create .env with DATABASE_URL"
    exit 1
fi

# Parse command line arguments
ANALYZE_ONLY=""
GENERATE_BN=""
SHOW_ALL=""
VERBOSE=""

while [[ $# -gt 0 ]]; do
    case $1 in
        analyze)
            ANALYZE_ONLY="--analyze"
            shift
            ;;
        generate)
            GENERATE_BN="--generate-bn"
            shift
            ;;
        all)
            SHOW_ALL="--show-all"
            shift
            ;;
        verbose)
            VERBOSE="--verbose"
            shift
            ;;
        *)
            shift
            ;;
    esac
done

echo "🚀 Starting Product Translator..."
echo ""

# Run the Go program
go run cmd/product-translator/main.go $ANALYZE_ONLY $GENERATE_BN $SHOW_ALL $VERBOSE
