# Upload to S3 from Postman - Step by Step Guide

## Prerequisites

1. Import the Postman collection: `postman/GO CRIT API.postman_collection.json`
2. Make sure your server is running on `http://localhost:8080`
3. Set the `BASEURL` variable in Postman to `http://localhost:8080`

---

## Method 1: Direct S3 Upload (Recommended)

This is a 3-step process where the client uploads directly to S3.

### Step 1: Get Presigned URL

**Request:** `POST {{BASEURL}}/s3/presign`

**Headers:**

```
Content-Type: application/json
Authorization: Bearer {{TOKEN}}  (optional - remove if no auth)
```

**Body (raw JSON):**

```json
{
  "filename": "my-product-image.jpg",
  "contentType": "image/jpeg",
  "productId": 1
}
```

**Response Example:**

```json
{
  "url": "https://kossti.s3.ap-southeast-1.amazonaws.com/product-images/p-1/550e8400-e29b-41d4-a716-446655440000.jpg?X-Amz-Algorithm=...",
  "key": "product-images/p-1/550e8400-e29b-41d4-a716-446655440000.jpg",
  "objectUrl": "https://kossti.s3.ap-southeast-1.amazonaws.com/product-images/p-1/550e8400-e29b-41d4-a716-446655440000.jpg",
  "expiresInSeconds": 900
}
```

**💡 Save the `url`, `key`, and `objectUrl` from the response - you'll need them!**

---

### Step 2: Upload File to S3 Using Presigned URL

**Create a NEW request in Postman:**

**Method:** `PUT`

**URL:** Paste the entire `url` value from Step 1 response

**Headers:**

```
Content-Type: image/jpeg  (must match what you specified in Step 1)
```

**Body:**

- Select **binary**
- Click **Select File**
- Choose your image file

**Click Send**

**Expected Response:**

- Status: `200 OK`
- Body: (empty or XML response from S3)

---

### Step 3: Register the Uploaded Image with Server

**Request:** `POST {{BASEURL}}/productimages/s3`

**Headers:**

```
Content-Type: application/json
Authorization: Bearer {{TOKEN}}  (optional)
```

**Body (raw JSON):**

```json
{
  "product_id": 1,
  "files": [
    {
      "key": "product-images/p-1/550e8400-e29b-41d4-a716-446655440000.jpg",
      "url": "https://kossti.s3.ap-southeast-1.amazonaws.com/product-images/p-1/550e8400-e29b-41d4-a716-446655440000.jpg",
      "name": "my-product-image.jpg",
      "size": 102400
    }
  ]
}
```

**Replace with values from Step 1 response:**

- `key` - use the `key` from Step 1
- `url` - use the `objectUrl` from Step 1
- `name` - your original filename
- `size` - file size in bytes (optional)

**Response Example:**

```json
{
  "success": true,
  "message": "S3 images registered",
  "images": [
    {
      "id": 1,
      "name": "my-product-image.jpg",
      "path": "product-images/p-1/550e8400-e29b-41d4-a716-446655440000.jpg",
      "product_id": 1,
      "default_photo": true,
      "created_at": "2025-11-06T08:00:00Z",
      "updated_at": "2025-11-06T08:00:00Z"
    }
  ]
}
```

---

## Method 2: Traditional Server Upload (Simpler but slower)

**Request:** `POST {{BASEURL}}/productimages`

**Headers:**

```
Authorization: Bearer {{TOKEN}}  (optional)
```

**Body:** Select `form-data`

| Key          | Type | Value                             |
| ------------ | ---- | --------------------------------- |
| `product_id` | Text | `1`                               |
| `files[0]`   | File | (Select your image file)          |
| `files[1]`   | File | (Select another image - optional) |

**Response Example:**

```json
{
  "success": true,
  "message": "2 image(s) uploaded successfully",
  "images": [
    {
      "id": 1,
      "name": "image1.jpg",
      "path": "/uploads/image1.jpg",
      "product_id": 1,
      "default_photo": true,
      "created_at": "2025-11-06T08:00:00Z",
      "updated_at": "2025-11-06T08:00:00Z"
    },
    {
      "id": 2,
      "name": "image2.jpg",
      "path": "/uploads/image2.jpg",
      "product_id": 1,
      "default_photo": false,
      "created_at": "2025-11-06T08:00:00Z",
      "updated_at": "2025-11-06T08:00:00Z"
    }
  ]
}
```

---

## Quick Test with Postman (No Auth Required)

### Test Direct Upload:

1. **Get presigned URL:**

   ```
   POST http://localhost:8080/s3/presign

   Body:
   {
     "filename": "test.jpg",
     "contentType": "image/jpeg",
     "productId": 1
   }
   ```

2. **Copy the `url` from response**

3. **Upload to S3:**

   ```
   PUT [paste the presigned URL here]

   Headers:
   Content-Type: image/jpeg

   Body: binary (select your image file)
   ```

4. **Register with server:**

   ```
   POST http://localhost:8080/productimages/s3

   Body:
   {
     "product_id": 1,
     "files": [
       {
         "key": "[copy key from step 1]",
         "url": "[copy objectUrl from step 1]",
         "name": "test.jpg",
         "size": 50000
       }
     ]
   }
   ```

---

## Troubleshooting

### ❌ Step 1 returns 500 error

- Make sure server is running: `./bin/gocrit`
- Check `.env` file has S3 configuration
- Verify AWS credentials are correct

### ❌ Step 2 returns 403 Forbidden

- The presigned URL expired (valid for 15 minutes)
- Get a new presigned URL and try again
- Check S3 bucket CORS settings

### ❌ Step 2 returns 400 Bad Request

- Make sure Content-Type header matches what you specified in Step 1
- Verify you're using PUT method (not POST)

### ❌ Can't find the uploaded file

- Check S3 console: https://s3.console.aws.amazon.com/s3/buckets/kossti
- Look in `product-images/p-{productId}/` folder

---

## Tips

1. **No authentication?** Just remove the `Authorization` header
2. **Testing quickly?** Use Method 2 (traditional upload) - only 1 request needed
3. **Production app?** Use Method 1 (presigned URL) - better performance
4. **Multiple files?** Add more objects to the `files` array in Step 3
5. **Want to test?** Run the automated test: `go run tools/test-s3/test_s3_presign_flow.go`

---

## Postman Collection Requests

These are already in the imported collection:

- **Images** folder:
  - `S3 Presign (Get PUT URL)` - Step 1
  - `Register S3 Uploaded Images` - Step 3
  - `Upload Product Images` - Method 2 (traditional)
  - `Get Product Images` - View all images for a product
  - `Set Default Image` - Change default image
  - `Remove Image` - Delete an image

---

## Next Steps

After uploading:

1. View images: `GET /productimages/{product_id}`
2. Set default: `POST /default-image/{image_id}`
3. Remove image: `POST /imageremove/{image_id}`
