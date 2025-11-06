package productreview

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
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

// PresignS3Handler returns a presigned PUT URL for uploading directly to S3.
func PresignS3Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req PresignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("AWS_REGION")
	if bucket == "" || region == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "S3_BUCKET or AWS_REGION is not configured"})
		return
	}

	// Load AWS config with explicit region
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to load AWS config"})
		return
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)

	// Generate a unique key. Keep original extension if present.
	ext := filepath.Ext(req.Filename)
	uid := uuid.New().String()
	key := "product-images/"
	if req.ProductID != 0 {
		// Use fmt to format numeric product ID correctly
		key = key + fmt.Sprintf("p-%d/", req.ProductID)
	}
	key = key + uid + ext

	// Build PutObject input
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String(req.ContentType),
	}

	// Presign for 15 minutes
	presignCtx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	presigned, err := presignClient.PresignPutObject(presignCtx, input, func(opts *s3.PresignOptions) {
		opts.Expires = 15 * time.Minute
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to presign url"})
		return
	}

	// Optional public object URL (depends on bucket region and settings)
	// Use the configured region
	objectURL := "https://" + bucket + ".s3." + region + ".amazonaws.com/" + key

	resp := PresignResponse{
		URL:       presigned.URL,
		Key:       key,
		ObjectURL: objectURL,
		ExpiresIn: int64(15 * 60),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
