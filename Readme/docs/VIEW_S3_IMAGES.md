# How to View/Download Images from S3

## Your bucket is private, so you need a presigned GET URL to view images.

---

## Method 1: Get Presigned GET URL (Recommended)

### Step 1: Get a viewable URL

**Request:** `POST http://localhost:8080/s3/presign-get`

**Headers:**

```
Content-Type: application/json
```

**Body (raw JSON):**

```json
{
  "key": "product-images/p-1/b4503d25-1d88-43b1-8f75-fd1b746ebf7f.jpg"
}
```

_(Use the `key` value from your upload response)_

**Response:**

```json
{
  "url": "https://kossti.s3.ap-southeast-1.amazonaws.com/product-images/p-1/...?X-Amz-Algorithm=...",
  "expiresInSeconds": 3600
}
```

### Step 2: View the image

Copy the `url` from the response and:

- **In Postman:** Create a GET request with that URL
- **In Browser:** Paste the URL directly
- **In your app:** Use this URL as the image src

The URL is valid for **1 hour**.

---

## Quick Test in Postman

1. **Get presign-get URL:**

   ```
   POST http://localhost:8080/s3/presign-get

   Body:
   {
     "key": "product-images/p-1/b4503d25-1d88-43b1-8f75-fd1b746ebf7f.jpg"
   }
   ```

2. **Copy the `url` from response**

3. **View image:**

   ```
   GET [paste the presigned URL]

   No headers needed
   ```

You should see the image in Postman's preview!

---

## Alternative: Make Bucket Public (Not Recommended for Production)

If you want images to be publicly accessible without presigned URLs:

1. Go to AWS S3 Console
2. Select your `kossti` bucket
3. Go to **Permissions** tab
4. Edit **Block public access** - uncheck "Block all public access"
5. Add this bucket policy:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PublicReadGetObject",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::kossti/product-images/*"
    }
  ]
}
```

⚠️ **Warning:** This makes all product images publicly accessible.

---

## Summary

**For private images (current setup):**

- Upload: Use presigned PUT URL (`/s3/presign`)
- View: Use presigned GET URL (`/s3/presign-get`)
- Both URLs expire after some time

**For public images:**

- Upload: Use presigned PUT URL
- View: Direct S3 URL (no presigning needed)

---

## Restart Your Server

To use the new endpoint, restart your server:

```bash
./bin/gocrit
```

Or if running in terminal:

```bash
cd gocrit_server
go run cmd/app/main.go
```
