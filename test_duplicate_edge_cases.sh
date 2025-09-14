#!/bin/bash

# Additional test cases for duplicate handling edge cases
BASE_URL="http://localhost:8080"

echo "Testing edge cases for duplicate product handling..."
echo "====================================================="

# Test 4: Case insensitive duplicate check
echo -e "\n4. Testing case-insensitive duplicate name check..."
RESPONSE4=$(curl -s -X POST "$BASE_URL/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "TEST DUPLICATE PRODUCT",
    "description": "Testing case insensitive duplicate",
    "price": 299.99,
    "category_id": 1,
    "priority": 1,
    "status": true
  }')

echo "Response: $RESPONSE4"

# Test 5: Try creating with empty name (should fail validation)
echo -e "\n5. Testing empty name validation..."
RESPONSE5=$(curl -s -X POST "$BASE_URL/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "description": "Empty name test",
    "price": 49.99,
    "category_id": 1,
    "priority": 1,
    "status": true
  }')

echo "Response: $RESPONSE5"

# Test 6: Create a product with unique name (should succeed)
echo -e "\n6. Creating product with unique name (should succeed)..."
UNIQUE_NAME="Unique Test Product $(date +%s)"
RESPONSE6=$(curl -s -X POST "$BASE_URL/products" \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"$UNIQUE_NAME\",
    \"description\": \"This should succeed with unique name\",
    \"price\": 79.99,
    \"category_id\": 1,
    \"priority\": 1,
    \"status\": true
  }")

echo "Response: $RESPONSE6"

echo -e "\n====================================================="
echo "Edge case testing completed!"
