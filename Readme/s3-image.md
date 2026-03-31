**S3 Image Upload — Quick README**

- **Purpose**: Steps to configure S3 for direct browser uploads using presigned PUT URLs and to debug common CORS/403 issues.

**Bucket & CORS**
- JSON to save as `cors.json` and apply with `aws s3api put-bucket-cors`:

```json
{
  "CORSRules": [
    {
      "AllowedOrigins": ["http://localhost:3000","https://your-production-domain.com"],
      "AllowedMethods": ["GET","PUT","POST","HEAD","OPTIONS"],
      "AllowedHeaders": ["*"],
      "ExposeHeaders": ["ETag","x-amz-request-id","x-amz-id-2"],
      "MaxAgeSeconds": 3000
    }
  ]
}
```

CLI:
```
aws s3api put-bucket-cors --bucket kossti-review --cors-configuration file://cors.json --region ap-south-1
```

**Bucket policy (allow backend account PutObject + optional public GetObject)**
- Replace `123456789012` with your account (example used earlier: `543927035288`):

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AllowAccountPutObject",
      "Effect": "Allow",
      "Principal": {"AWS": "arn:aws:iam::543927035288:root"},
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::kossti-review/product-images/*"
    },
    {
      "Sid": "AllowPublicGetObject",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::kossti-review/product-images/*"
    }
  ]
}
```

Note: paste this into the **Bucket Policy** editor (policy must be a JSON object starting with `{`). Do NOT paste CORS JSON into the policy editor.

**Code changes (presign + display URL)**
- File: `gocrit_server/internal/interface/handler/productreview/presign.go`
  - Ensure the `PutObjectInput` used for presigning includes `ContentType` so browser `Content-Type` matches signature. Minimal example:

```go
contentType := "application/octet-stream"
if req.ContentType != "" {
    contentType = req.ContentType
}
putInput := &s3.PutObjectInput{
    Bucket:      &bucket,
    Key:         &key,
    ContentType: &contentType,
}
// use putInput when calling presign client
```

- File: `gocrit_server/internal/interface/handler/product/handler.go`
  - Remove or update hardcoded fallback `bucket = "kossti"`; rely on `S3_BUCKET` env or set default to `kossti-review` if desired.

**Restart backend**
- After env changes (`S3_BUCKET`, `AWS_REGION`), restart the server (or `air`) so new env values are picked up.

**Diagnostics & test commands**
- Verify CORS applied:
```
aws s3api get-bucket-cors --bucket kossti-review --region ap-south-1
```
- Test OPTIONS preflight (replace PRESIGNED_URL):
```bash
curl -i -X OPTIONS "PRESIGNED_URL" \
  -H "Origin: http://localhost:3000" \
  -H "Access-Control-Request-Method: PUT" \
  -H "Access-Control-Request-Headers: content-type"
```
- PUT test (match Content-Type used when presigning):
```bash
curl -i -X PUT "PRESIGNED_URL" \
  -H "Content-Type: image/webp" \
  --data-binary @test.webp
```
- Check bucket policy and ACL:
```
aws s3api get-bucket-policy --bucket kossti-review --region ap-south-1
aws s3api get-bucket-acl --bucket kossti-review --region ap-south-1
```

**Common causes of 403 on OPTIONS / PUT**
- CORS not applied to the exact bucket/region.
- Bucket policy Deny or missing `s3:PutObject` for the IAM account that signs presigned URLs.
- Presigned PUT was signed with a `Content-Type` but browser sends a different `Content-Type` (include `ContentType` in signed request).
- Running backend still using old env or fallback `kossti` instead of `kossti-review`.

**Where to edit**
- Presign generation: `gocrit_server/internal/interface/handler/productreview/presign.go`
- Image URL generation / fallback: `gocrit_server/internal/interface/handler/product/handler.go`
- Frontend proxy: `crit_client/src/app/api/s3-presign/route.ts` (proxies to backend)
- Next.js image host whitelist: `crit_client/next.config.ts` (update host to kossti-review if needed)

**If you want me to do a patch**
- I can create a small patch to add `ContentType` in presign and remove the hardcoded fallback. Tell me to proceed and I'll apply it.

---
Generated on 2026-03-31. Thank you.
