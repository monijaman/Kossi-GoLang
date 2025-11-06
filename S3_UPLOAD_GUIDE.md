# S3 Image Upload Configuration

## Overview

This document describes the S3 image upload functionality for product reviews.

## Architecture

The system supports two upload flows:

### 1. Direct S3 Upload (Presigned URL)

**Recommended for production** - Client uploads directly to S3, reducing server load.

**Flow:**

1. Client requests a presigned URL from `/s3/presign`
2. Server returns a presigned PUT URL valid for 15 minutes
3. Client uploads file directly to S3 using the presigned URL
4. Client registers the uploaded file with `/productimages/s3`
5. Server stores metadata in the database

### 2. Server Upload (Traditional)

Client uploads to `/productimages`, server handles S3 upload.

## Environment Configuration

Add these variables to your `.env` file:

```bash
# AWS S3 Configuration
S3_BUCKET=kossti
AWS_REGION=ap-southeast-1
AWS_ACCESS_KEY_ID=your-access-key-id
AWS_SECRET_ACCESS_KEY=your-secret-access-key
```

## S3 Bucket Setup

### Required Permissions

Your AWS IAM user needs these S3 permissions:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:PutObject",
        "s3:GetObject",
        "s3:DeleteObject",
        "s3:ListBucket"
      ],
      "Resource": ["arn:aws:s3:::kossti/*", "arn:aws:s3:::kossti"]
    }
  ]
}
```

### CORS Configuration

For browser uploads, configure CORS on your S3 bucket:

```json
[
  {
    "AllowedHeaders": ["*"],
    "AllowedMethods": ["GET", "PUT", "POST", "DELETE"],
    "AllowedOrigins": ["*"],
    "ExposeHeaders": ["ETag"],
    "MaxAgeSeconds": 3000
  }
]
```

## API Endpoints

### POST /s3/presign

Request a presigned URL for direct S3 upload.

**Request:**

```json
{
  "filename": "example.jpg",
  "contentType": "image/jpeg",
  "productId": 123
}
```

**Response:**

```json
{
  "url": "https://kossti.s3.ap-southeast-1.amazonaws.com/product-images/p-123/uuid.jpg?X-Amz-...",
  "key": "product-images/p-123/uuid.jpg",
  "objectUrl": "https://kossti.s3.ap-southeast-1.amazonaws.com/product-images/p-123/uuid.jpg",
  "expiresInSeconds": 900
}
```

### POST /productimages/s3

Register S3-uploaded images with the server.

**Request:**

```json
{
  "product_id": 123,
  "files": [
    {
      "key": "product-images/p-123/uuid.jpg",
      "url": "https://kossti.s3.ap-southeast-1.amazonaws.com/product-images/p-123/uuid.jpg",
      "name": "example.jpg",
      "size": 102400
    }
  ]
}
```

**Response:**

```json
{
  "success": true,
  "message": "S3 images registered",
  "images": [
    {
      "id": 1,
      "name": "example.jpg",
      "path": "product-images/p-123/uuid.jpg",
      "product_id": 123,
      "default_photo": true,
      "created_at": "2025-11-06T08:00:00Z",
      "updated_at": "2025-11-06T08:00:00Z"
    }
  ]
}
```

### POST /productimages

Traditional multipart upload (server handles S3).

**Request:**

```
POST /productimages
Content-Type: multipart/form-data

product_id: 123
files[0]: (binary image data)
files[1]: (binary image data)
```

## File Organization

Images are stored with this key structure:

```
product-images/
  p-{product_id}/
    {uuid}.{extension}
```

Example: `product-images/p-123/550e8400-e29b-41d4-a716-446655440000.jpg`

## Testing

### Test Direct S3 Upload

```bash
go run tools/test-s3/test_s3_upload.go
```

### Test Complete Presign Flow

```bash
# Start the server first
./bin/gocrit

# In another terminal
go run tools/test-s3/test_s3_presign_flow.go
```

## Dependencies

```bash
go get github.com/aws/aws-sdk-go-v2/config@latest
go get github.com/aws/aws-sdk-go-v2/service/s3@latest
go get github.com/google/uuid@latest
```

## Postman Collection

The Postman collection includes examples for:

- S3 Presign (Get PUT URL)
- Register S3 Uploaded Images
- Upload Product Images (traditional)

Import `postman/GO CRIT API.postman_collection.json` to test the endpoints.

## Security Notes

1. **Presigned URLs** expire after 15 minutes
2. **Product folder isolation**: Each product's images are stored in separate folders
3. **Content-Type validation**: Presign endpoint requires valid content type
4. **File extensions**: Original file extensions are preserved
5. **Never commit** AWS credentials to version control

## Troubleshooting

### Upload fails with 403 Forbidden

- Check IAM permissions on your AWS user
- Verify bucket policy allows PutObject
- Check bucket CORS configuration

### Presign endpoint returns 500

- Verify AWS credentials are set in `.env`
- Check S3_BUCKET and AWS_REGION are configured
- Ensure AWS credentials have permission to sign URLs

### Files upload but URLs don't work

- Check bucket public access settings
- Verify bucket policy for GetObject
- Confirm region in objectUrl matches bucket region

## Production Checklist

- [ ] Configure AWS IAM user with minimal required permissions
- [ ] Set up S3 bucket CORS for your frontend domain
- [ ] Enable S3 bucket encryption
- [ ] Configure bucket lifecycle rules for old test uploads
- [ ] Set up CloudFront CDN for better performance
- [ ] Implement image validation (size, type) before presigning
- [ ] Add rate limiting on presign endpoint
- [ ] Monitor S3 costs and usage
