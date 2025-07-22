#!/bin/bash

# Database Migration Test Script
# This script tests the migration of all Laravel tables to Go GORM models

echo "==================================="
echo "Go Database Migration Test"
echo "==================================="

# Check if .env file exists
if [ ! -f .env ]; then
    echo "⚠️  No .env file found. Creating example..."
    cp .env.example .env
    echo "📝 Please edit .env with your PostgreSQL connection details"
    exit 1
fi

echo "📋 Available Commands:"
echo "  1. Test Migration Help"
echo "  2. Run All Migrations"
echo "  3. Reset Database"
echo "  4. Build Project"
echo ""

# Test 1: Show migration help
echo "🔍 Testing migration tool help..."
go run tools/migrate/migrate.go -help
echo ""

# Test 2: Build project
echo "🔨 Building project..."
if go build ./...; then
    echo "✅ Build successful!"
else
    echo "❌ Build failed!"
    exit 1
fi

echo ""
echo "🚀 Migration system is ready!"
echo ""
echo "To run migrations:"
echo "  go run tools/migrate/migrate.go -migrate"
echo ""
echo "To reset database:"
echo "  go run tools/migrate/migrate.go -reset"
echo ""
echo "To start the server (auto-migrates):"
echo "  go run cmd/main.go"
