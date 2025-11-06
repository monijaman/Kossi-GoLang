package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	fmt.Printf("=== Environment Variables ===\n")
	fmt.Printf("S3_BUCKET: %s\n", bucket)
	fmt.Printf("AWS_REGION: %s\n", region)
	fmt.Printf("AWS_ACCESS_KEY_ID: %s\n", accessKey)
	fmt.Printf("AWS_SECRET_ACCESS_KEY: %s...\n", secretKey[:10])

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	fmt.Printf("\n=== AWS Config ===\n")
	fmt.Printf("Region: %s\n", cfg.Region)

	// Get credentials to verify
	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve credentials: %v", err)
	}

	fmt.Printf("Access Key ID (from config): %s\n", creds.AccessKeyID)
	fmt.Printf("Match: %v\n", creds.AccessKeyID == accessKey)

	// Create S3 client and test presigning
	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)

	// Presign a GET request
	key := "product-images/p-33/09bb5273-79c1-4576-bf67-74be03ff2c71.jpg"
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	presigned, err := presignClient.PresignGetObject(context.Background(), input)
	if err != nil {
		log.Fatalf("Failed to presign: %v", err)
	}

	fmt.Printf("\n=== Presigned URL ===\n")
	fmt.Printf("URL: %s\n", presigned.URL)

	// Check if URL contains the correct access key
	if strings.Contains(presigned.URL, accessKey) {
		fmt.Printf("✓ URL contains correct access key ID\n")
	} else {
		fmt.Printf("✗ WARNING: URL does not contain expected access key ID\n")
	}
}
