package productreview

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type PresignGetRequest struct {
	Key string `json:"key"` // S3 object key
}

type PresignGetResponse struct {
	URL       string `json:"url"`
	ExpiresIn int64  `json:"expiresInSeconds"`
}

// PresignGetHandler returns a presigned GET URL for downloading/viewing from S3.
func PresignGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req PresignGetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if req.Key == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "key is required"})
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

	// Build GetObject input
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(req.Key),
	}

	// Presign for 1 hour
	presignCtx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	presigned, err := presignClient.PresignGetObject(presignCtx, input, func(opts *s3.PresignOptions) {
		opts.Expires = 1 * time.Hour
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to presign url"})
		return
	}

	resp := PresignGetResponse{
		URL:       presigned.URL,
		ExpiresIn: int64(60 * 60), // 1 hour in seconds
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
