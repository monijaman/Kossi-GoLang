package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("AWS_REGION")

	if bucket == "" {
		log.Fatal("S3_BUCKET is not set")
	}
	if region == "" {
		log.Fatal("AWS_REGION is not set")
	}

	fmt.Printf("Testing S3 upload to bucket: %s in region: %s\n", bucket, region)

	// Load AWS config
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)

	// Create test file content
	testContent := fmt.Sprintf("S3 Test Upload - %s", time.Now().Format(time.RFC3339))
	testKey := fmt.Sprintf("test-uploads/test-%d.txt", time.Now().Unix())

	// Upload test file
	fmt.Printf("Uploading test file to: %s\n", testKey)
	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &testKey,
		Body:        strings.NewReader(testContent),
		ContentType: stringPtr("text/plain"),
	})

	if err != nil {
		log.Fatalf("Failed to upload to S3: %v", err)
	}

	fmt.Printf("✅ Successfully uploaded test file!\n")
	fmt.Printf("   Key: %s\n", testKey)
	fmt.Printf("   URL: https://%s.s3.%s.amazonaws.com/%s\n", bucket, region, testKey)

	// Verify the file exists by reading it back
	fmt.Println("\nVerifying upload by reading the file...")
	result, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &testKey,
	})
	if err != nil {
		log.Fatalf("Failed to read from S3: %v", err)
	}
	defer result.Body.Close()

	fmt.Println("✅ File verified successfully!")
	fmt.Printf("   Content-Type: %s\n", *result.ContentType)
	fmt.Printf("   Content-Length: %d bytes\n", result.ContentLength)

	// List objects in test-uploads folder
	fmt.Println("\nListing test-uploads folder...")
	listResult, err := s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &bucket,
		Prefix: stringPtr("test-uploads/"),
		MaxKeys: int32Ptr(10),
	})
	if err != nil {
		log.Printf("Warning: Could not list objects: %v", err)
	} else {
		fmt.Printf("Found %d object(s) in test-uploads/:\n", len(listResult.Contents))
		for _, obj := range listResult.Contents {
			fmt.Printf("  - %s (%d bytes, modified: %s)\n", 
				*obj.Key, obj.Size, obj.LastModified.Format(time.RFC3339))
		}
	}

	fmt.Println("\n🎉 S3 configuration is working correctly!")
}

func stringPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}
