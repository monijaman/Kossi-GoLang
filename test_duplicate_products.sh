#!/bin/bash

# Test script for duplicate product handling
BASE_URL="http://localhost:8080"

echo "Testing duplicate product handling..."
echo "====================================="

# Test 1: Create a new product
echo -e "\n1. Creating first product 'Test Duplicate Product'..."
RESPONSE1=$(curl -s -X POST "$BASE_URL/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Duplicate Product",
    "description": "This is a test product for duplicate checking",
    "price": 99.99,
    "category_id": 1,
    "priority": 1,
    "status": true
  }')

echo "Response: $RESPONSE1"

# Test 2: Try to create the same product again (should fail with duplicate error)
echo -e "\n2. Attempting to create duplicate product with same name..."
RESPONSE2=$(curl -s -X POST "$BASE_URL/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Duplicate Product",
    "description": "This should fail due to duplicate name",
    "price": 149.99,
    "category_id": 1,
    "priority": 1,
    "status": true
  }')

echo "Response: $RESPONSE2"

# Test 3: Try to create product with explicit duplicate slug
echo -e "\n3. Attempting to create product with duplicate slug..."
RESPONSE3=$(curl -s -X POST "$BASE_URL/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Different Name But Same Slug",
    "slug": "test-duplicate-product",
    "description": "This should fail due to duplicate slug",
    "price": 199.99,
    "category_id": 1,
    "priority": 1,
    "status": true
  }')

echo "Response: $RESPONSE3"

echo -e "\n====================================="
echo "Test completed!"
