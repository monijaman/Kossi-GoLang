# Image Deletion Implementation - Complete

## Summary

Implemented full image deletion functionality that removes images from both the PostgreSQL database and AWS S3 storage.

## Changes Made

### Backend (Go Server)

#### 1. Updated `handler.go` - RemoveImageHandler

**File:** `gocrit_server/internal/interface/handler/productreview/handler.go`

**New functionality:**

- Accepts `imageRepo` parameter
- Fetches image from database to get S3 key
- Deletes object from S3 using AWS SDK v2
- Deletes record from database
- Returns success response with `image_id`

**Key features:**

- ✅ Gets image metadata from DB first
- ✅ Deletes from S3 bucket using `DeleteObject`
- ✅ Continues with DB deletion even if S3 delete fails
- ✅ Proper error handling and logging
- ✅ Returns JSON response with success status

#### 2. Updated `routes.go`

**File:** `gocrit_server/internal/interface/handler/productreview/routes.go`

**Change:**

```go
// Before:
RemoveImageHandler(w, r, reviewRepo)

// After:
RemoveImageHandler(w, r, reviewRepo, imageRepo)
```

### Frontend (Next.js/React)

#### 3. Updated `uploadSlice.ts` - removeMedia Action

**File:** `crit_client/src/redux/features/upload/uploadSlice.ts`

**Old behavior:**

- Called `/api/post` proxy route (which didn't exist)
- Passed `apiUrl` in request body

**New behavior:**

- Directly calls Go API: `${NEXT_PUBLIC_API_URL}/imageremove/${productId}`
- Simple POST request with no body
- Proper error handling
- Returns `{success: true, data: responseData}`

## API Endpoint

### DELETE Image

**Endpoint:** `POST /imageremove/{id}`

**Request:**

```bash
curl -X POST http://localhost:8083/imageremove/4
```

**Response:**

```json
{
  "success": true,
  "message": "Image removed successfully",
  "image_id": 4
}
```

## Flow

1. **User clicks "Remove" in Uploader component**
2. **Frontend dispatches `removeMedia` action**
   - Calls `POST /imageremove/{imageId}`
3. **Backend receives request:**
   - Gets image from DB by ID
   - Extracts S3 key from `image.ImagePath`
   - Calls AWS S3 `DeleteObject(bucket, key)`
   - Deletes image record from DB
   - Returns success response
4. **Frontend receives response:**
   - Calls `getPhotos()` to refresh image list
   - UI updates automatically

## Testing

### Tested Successfully:

```bash
# Delete image ID 4
curl -X POST http://localhost:8083/imageremove/4

# Response:
{"image_id":4,"message":"Image removed successfully","success":true}

# Verify deletion
curl -s http://localhost:8083/productimages/33 | grep '"id"'
# Result: Only IDs 2 and 3 remain (ID 4 is gone)

# Verify S3 deletion
go run tools/test-s3/list_s3_objects.go
# Result: Only 2 objects remain in S3
```

## Database Schema

The `images` table stores:

- `id`: Primary key
- `imageable_type`: "Product"
- `imageable_id`: Product ID
- `image_path`: S3 key (e.g., "product-images/p-33/uuid.jpg")
- `status`: Active/Inactive
- `created_at`, `updated_at`: Timestamps

## Error Handling

### If S3 deletion fails:

- Logs error: `[error] Failed to delete from S3: {error}`
- **Continues with database deletion anyway**
- This prevents orphaned DB records if S3 is temporarily unavailable

### If database deletion fails:

- Returns 500 status code
- Response: `{"error": "Failed to delete image"}`
- S3 file remains (can be cleaned up later)

### If image not found:

- Returns 404 status code
- Response: `{"error": "Image not found"}`

## Frontend Integration

The Uploader component already has the UI:

```tsx
<div className="remove-file" onClick={() => handleRemoveUPloadedFile(file.id)}>
  Remove
</div>
```

When clicked:

1. Calls `dispatch(removeMedia({ productId: imageId }))`
2. Shows loading state
3. On success: refreshes photos with `getPhotos()`
4. On error: logs to console

## Server Status

- ✅ Go server running on port **8083**
- ✅ Built with updated code
- ✅ Delete endpoint tested and working
- ✅ S3 deletion confirmed working

## Next Steps

**For you to test:**

1. Go to `http://localhost:3000/admin/products/33`
2. Click "Add Photos" button
3. You should see your uploaded images
4. Click "Remove" on any image
5. Confirm it disappears from the list
6. Verify it's gone from S3

**Everything is ready to use!** 🎉
