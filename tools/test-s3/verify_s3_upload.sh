#!/bin/bash

# Test script to verify if images exist in S3
# This tests the complete upload flow

echo "=== Testing S3 Upload Verification ==="

# Step 1: Get presigned PUT URL
echo -e "\n1. Getting presigned PUT URL..."
PRESIGN_RESPONSE=$(curl -s -X POST http://localhost:8083/s3/presign \
  -H "Content-Type: application/json" \
  -d '{
    "filename": "test-verify.jpg",
    "contentType": "image/jpeg",
    "productId": 33
  }')

echo "Presign Response: $PRESIGN_RESPONSE"

PRESIGN_URL=$(echo $PRESIGN_RESPONSE | grep -o '"url":"[^"]*' | cut -d'"' -f4)
S3_KEY=$(echo $PRESIGN_RESPONSE | grep -o '"key":"[^"]*' | cut -d'"' -f4)

echo "Presigned URL: $PRESIGN_URL"
echo "S3 Key: $S3_KEY"

# Step 2: Upload a test file
echo -e "\n2. Uploading test file to S3..."
echo "Test Image Content" > /tmp/test-verify.jpg

UPLOAD_RESPONSE=$(curl -s -X PUT "$PRESIGN_URL" \
  -H "Content-Type: image/jpeg" \
  --data-binary "@/tmp/test-verify.jpg" \
  -w "\nHTTP_CODE:%{http_code}")

echo "Upload Response: $UPLOAD_RESPONSE"

# Step 3: Verify file exists by getting presigned GET URL
echo -e "\n3. Getting presigned GET URL to verify..."
GET_URL_RESPONSE=$(curl -s -X POST http://localhost:8083/s3/presign-get \
  -H "Content-Type: application/json" \
  -d "{\"key\": \"$S3_KEY\"}")

echo "Get URL Response: $GET_URL_RESPONSE"

GET_URL=$(echo $GET_URL_RESPONSE | grep -o '"url":"[^"]*' | cut -d'"' -f4)
echo "Presigned GET URL: $GET_URL"

# Step 4: Try to access the file
echo -e "\n4. Testing if file is accessible..."
ACCESS_RESPONSE=$(curl -s -I "$GET_URL" 2>&1 | head -1)
echo "Access Response: $ACCESS_RESPONSE"

if [[ $ACCESS_RESPONSE == *"200"* ]]; then
    echo "✓ SUCCESS: File is accessible in S3"
elif [[ $ACCESS_RESPONSE == *"403"* ]]; then
    echo "✗ FAILED: File not found or access denied (403)"
    echo "This means the upload to S3 failed or AWS credentials are incorrect"
else
    echo "✗ UNKNOWN: Unexpected response: $ACCESS_RESPONSE"
fi

# Step 5: Register the image in database
echo -e "\n5. Registering image in database..."
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:8083/productimages/s3 \
  -H "Content-Type: application/json" \
  -d "{
    \"product_id\": 33,
    \"files\": [{
      \"key\": \"$S3_KEY\",
      \"name\": \"test-verify.jpg\",
      \"url\": \"https://kossti.s3.ap-southeast-1.amazonaws.com/$S3_KEY\",
      \"size\": 100
    }]
  }")

echo "Register Response: $REGISTER_RESPONSE"

# Cleanup
rm -f /tmp/test-verify.jpg

echo -e "\n=== Test Complete ===
