# S3 Image Access Issue - 403 Forbidden

## Problem

The presigned URLs are being generated correctly, but return **403 Forbidden** when accessed. This is an **AWS IAM permissions issue**.

## Root Cause

The AWS Access Key ID `AKIAYS2NSM5EYHJ6WFWU` has permission to:

- ✅ `s3:ListBucket` (can list objects)
- ❌ `s3:GetObject` (CANNOT read objects)
- ❌ `s3:PutObject` (likely CANNOT upload either)

## Solution

You need to grant the IAM user/role associated with `AKIAYS2NSM5EYHJ6WFWU` full S3 access to the `kossti` bucket.

### Option 1: Update IAM User Policy (Recommended)

1. Go to [AWS IAM Console](https://console.aws.amazon.com/iam/)
2. Navigate to **Users** → Find your IAM user
3. Click **Add permissions** → **Attach policies directly**
4. Create a new inline policy using `iam-user-policy.json`:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "S3FullAccessToKosstiBucket",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject",
        "s3:GetObject",
        "s3:DeleteObject",
        "s3:ListBucket",
        "s3:GetBucketLocation"
      ],
      "Resource": ["arn:aws:s3:::kossti", "arn:aws:s3:::kossti/*"]
    }
  ]
}
```

### Option 2: Update S3 Bucket Policy

1. Go to [S3 Console](https://s3.console.aws.amazon.com/s3/)
2. Select `kossti` bucket → **Permissions** tab
3. Edit **Bucket Policy**
4. Replace `YOUR_ACCOUNT_ID` and `YOUR_IAM_USER` in `s3-bucket-policy.json`
5. Save the policy

### Option 3: Use AWS CLI

If you have AWS CLI configured:

```bash
# Attach policy to IAM user
aws iam put-user-policy \
  --user-name YOUR_IAM_USER \
  --policy-name KosstiS3Access \
  --policy-document file://iam-user-policy.json

# OR update bucket policy
aws s3api put-bucket-policy \
  --bucket kossti \
  --policy file://s3-bucket-policy.json
```

## Verification

After applying the policy:

1. Wait 1-2 minutes for IAM changes to propagate
2. Test the presigned URL again:

```bash
cd /i/GO/kossti/gocrit_server
bash tools/test-s3/test_presigned_get.sh
```

3. You should see **HTTP 200 OK** instead of 403

## Current Status

- ✅ Files uploaded to S3: `product-images/p-33/09bb5273-79c1-4576-bf67-74be03ff2c71.jpg`
- ✅ Presigned URLs generated correctly with valid signatures
- ❌ Access denied due to missing `s3:GetObject` permission
- ⏳ Waiting for IAM policy update

## Next Steps

1. Apply IAM permissions (see options above)
2. Test presigned URLs
3. Update frontend `.env` to use `http://localhost:8083`
4. Apply CORS configuration from `cors-config.json`
5. Test complete upload flow from frontend
