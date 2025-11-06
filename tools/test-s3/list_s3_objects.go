package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	fmt.Printf("Loaded AWS Config Region: %s\n", cfg.Region)

	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)

	// List objects in the product-images/p-33/ prefix
	fmt.Println("\n=== Listing objects in bucket ===")
	listInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String("product-images/p-33/"),
	}

	result, err := s3Client.ListObjectsV2(context.Background(), listInput)
	if err != nil {
		log.Fatalf("Failed to list objects: %v", err)
	}

	fmt.Printf("Found %d objects:\n", len(result.Contents))
	for _, obj := range result.Contents {
		fmt.Printf("  - %s (size: %d bytes)\n", *obj.Key, obj.Size)
	}

	if len(result.Contents) == 0 {
		fmt.Println("\n⚠️  No objects found! The images were never uploaded to S3.")
		fmt.Println("This explains why the presigned GET URLs return 403.")
	}
}
