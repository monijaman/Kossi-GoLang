#!/bin/bash

# Test the presigned GET URL properly

echo "=== Testing Presigned GET URL ==="

# Get presigned URL
RESPONSE=$(curl -s "http://localhost:8083/s3/presign-get" \
  -H "Content-Type: application/json" \
  -d '{"key":"product-images/p-33/09bb5273-79c1-4576-bf67-74be03ff2c71.jpg"}')

echo "Response: $RESPONSE"

# Extract URL using Python to properly decode JSON
URL=$(echo "$RESPONSE" | python3 -c "import sys, json; print(json.load(sys.stdin)['url'])" 2>/dev/null)

if [ -z "$URL" ]; then
    echo "Failed to extract URL from JSON"
    exit 1
fi

echo -e "\nPresigned URL: $URL"

# Test accessing the URL
echo -e "\n=== Testing URL Access ==="
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" "$URL")

echo "HTTP Status Code: $HTTP_CODE"

if [ "$HTTP_CODE" == "200" ]; then
    echo "✓ SUCCESS: Image is accessible!"
    
    # Download and check file info
    curl -s "$URL" -o /tmp/test-download.jpg
    FILE_SIZE=$(wc -c < /tmp/test-download.jpg)
    echo "File downloaded: $FILE_SIZE bytes"
    rm -f /tmp/test-download.jpg
elif [ "$HTTP_CODE" == "403" ]; then
    echo "✗ FAILED: Access Denied (403)"
    echo "Possible causes:"
    echo "  - AWS credentials don't match"
    echo "  - Bucket policy denies access"
    echo "  - Signature mismatch"
else
    echo "✗ FAILED: Unexpected status code $HTTP_CODE"
fi
