package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type PresignRequest struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	ProductID   uint   `json:"productId,omitempty"`
}

type PresignResponse struct {
	URL       string `json:"url"`
	Key       string `json:"key"`
	ObjectURL string `json:"objectUrl"`
	ExpiresIn int64  `json:"expiresInSeconds"`
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("AWS_REGION")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Testing S3 Presign Flow\n")
	fmt.Printf("Bucket: %s, Region: %s\n\n", bucket, region)

	// Step 1: Request presigned URL from server
	baseURL := fmt.Sprintf("http://localhost:%s", port)
	presignURL := baseURL + "/s3/presign"

	reqBody := PresignRequest{
		Filename:    "test-image.jpg",
		ContentType: "image/jpeg",
		ProductID:   123,
	}

	jsonData, _ := json.Marshal(reqBody)
	fmt.Printf("Step 1: Requesting presigned URL from %s\n", presignURL)
	fmt.Printf("Request body: %s\n", string(jsonData))

	resp, err := http.Post(presignURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to request presign URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Server returned status %d: %s", resp.StatusCode, string(body))
	}

	var presignResp PresignResponse
	if err := json.NewDecoder(resp.Body).Decode(&presignResp); err != nil {
		log.Fatalf("Failed to decode presign response: %v", err)
	}

	fmt.Printf("✅ Received presigned URL\n")
	fmt.Printf("   Key: %s\n", presignResp.Key)
	fmt.Printf("   Object URL: %s\n", presignResp.ObjectURL)
	fmt.Printf("   Expires in: %d seconds\n\n", presignResp.ExpiresIn)

	// Step 2: Upload file using presigned URL
	testContent := []byte("This is a test image content from presign flow - " + time.Now().Format(time.RFC3339))
	fmt.Printf("Step 2: Uploading file using presigned URL\n")
	fmt.Printf("Content size: %d bytes\n", len(testContent))

	putReq, err := http.NewRequest("PUT", presignResp.URL, bytes.NewReader(testContent))
	if err != nil {
		log.Fatalf("Failed to create PUT request: %v", err)
	}
	putReq.Header.Set("Content-Type", reqBody.ContentType)

	client := &http.Client{}
	putResp, err := client.Do(putReq)
	if err != nil {
		log.Fatalf("Failed to upload to S3: %v", err)
	}
	defer putResp.Body.Close()

	if putResp.StatusCode < 200 || putResp.StatusCode >= 300 {
		body, _ := io.ReadAll(putResp.Body)
		log.Fatalf("S3 upload failed with status %d: %s", putResp.StatusCode, string(body))
	}

	fmt.Printf("✅ File uploaded successfully to S3\n\n")

	// Step 3: Verify the file exists in S3
	fmt.Printf("Step 3: Verifying file in S3\n")
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	getResult, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &presignResp.Key,
	})
	if err != nil {
		log.Fatalf("Failed to verify file in S3: %v", err)
	}
	defer getResult.Body.Close()

	downloadedContent, _ := io.ReadAll(getResult.Body)
	fmt.Printf("✅ File verified in S3\n")
	fmt.Printf("   Content-Length: %d bytes\n", len(downloadedContent))
	fmt.Printf("   Content matches: %v\n", bytes.Equal(testContent, downloadedContent))

	// Step 4: Test register S3 images endpoint
	fmt.Printf("\nStep 4: Testing register S3 images endpoint\n")
	registerURL := baseURL + "/productimages/s3"
	registerPayload := map[string]interface{}{
		"product_id": 123,
		"files": []map[string]interface{}{
			{
				"key":  presignResp.Key,
				"url":  presignResp.ObjectURL,
				"name": reqBody.Filename,
				"size": len(testContent),
			},
		},
	}

	registerJSON, _ := json.Marshal(registerPayload)
	fmt.Printf("Request body: %s\n", string(registerJSON))

	registerResp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(registerJSON))
	if err != nil {
		log.Fatalf("Failed to register S3 images: %v", err)
	}
	defer registerResp.Body.Close()

	registerBody, _ := io.ReadAll(registerResp.Body)
	if registerResp.StatusCode != http.StatusOK {
		log.Fatalf("Register endpoint returned status %d: %s", registerResp.StatusCode, string(registerBody))
	}

	fmt.Printf("✅ Images registered successfully\n")
	fmt.Printf("Response: %s\n", string(registerBody))

	fmt.Println("\n🎉 Complete S3 presign flow working correctly!")
	fmt.Println("\nFlow summary:")
	fmt.Println("1. ✅ Requested presigned URL from server")
	fmt.Println("2. ✅ Uploaded file to S3 using presigned URL")
	fmt.Println("3. ✅ Verified file exists in S3")
	fmt.Println("4. ✅ Registered S3 image with server")
}
