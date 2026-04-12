package product

import (
	"context"
	"encoding/json"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"net/http"
	"strconv"
	"strings"
)

// VideoItem represents a video entry
type VideoItem struct {
	ID         uint   `json:"id"`
	URL        string `json:"url"`
	Title      string `json:"title"`
	YoutubeURL string `json:"youtubeUrl,omitempty"`
}

// GetProductVideosHandler handles GET /product-videos/{id}
// Fetches YouTube URLs from product_reviews.additional_details
func GetProductVideosHandler(w http.ResponseWriter, r *http.Request, reviewRepo repository.ProductReviewRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/product-videos/")
	productIDStr := strings.Trim(path, "/")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	// Get reviews for this product
	reviews, err := reviewRepo.GetByProductID(context.Background(), uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch reviews"})
		return
	}

	// Extract video URLs from additional_details
	var videos []VideoItem
	for _, review := range reviews {
		if len(review.AdditionalDetails) == 0 {
			continue
		}

		// Try to parse additional_details as JSON
		var additionalData interface{}
		if err := json.Unmarshal(review.AdditionalDetails, &additionalData); err != nil {
			// If it's not JSON, treat it as a plain string (might be a direct URL)
			url := strings.TrimSpace(string(review.AdditionalDetails))
			if isYouTubeURL(url) {
				videos = append(videos, VideoItem{
					ID:         review.ID,
					URL:        url,
					YoutubeURL: url,
					Title:      getReviewTitle(review),
				})
			}
			continue
		}

		// Handle different JSON structures
		switch data := additionalData.(type) {
		case string:
			// Simple string URL
			if isYouTubeURL(data) {
				videos = append(videos, VideoItem{
					ID:         review.ID,
					URL:        data,
					YoutubeURL: data,
					Title:      getReviewTitle(review),
				})
			}

		case map[string]interface{}:
			// Object with url/youtubeUrl field
			videoURL := ""
			title := getReviewTitle(review)

			if url, ok := data["url"].(string); ok {
				videoURL = url
			} else if url, ok := data["youtubeUrl"].(string); ok {
				videoURL = url
			} else if url, ok := data["youtube_url"].(string); ok {
				videoURL = url
			}

			if titleVal, ok := data["title"].(string); ok && titleVal != "" {
				title = titleVal
			}

			if isYouTubeURL(videoURL) {
				videos = append(videos, VideoItem{
					ID:         review.ID,
					URL:        videoURL,
					YoutubeURL: videoURL,
					Title:      title,
				})
			}

		case []interface{}:
			// Array of video objects
			for _, item := range data {
				if videoObj, ok := item.(map[string]interface{}); ok {
					videoURL := ""
					title := getReviewTitle(review)

					if url, ok := videoObj["url"].(string); ok {
						videoURL = url
					} else if url, ok := videoObj["youtubeUrl"].(string); ok {
						videoURL = url
					} else if url, ok := videoObj["youtube_url"].(string); ok {
						videoURL = url
					}

					if titleVal, ok := videoObj["title"].(string); ok && titleVal != "" {
						title = titleVal
					}

					if isYouTubeURL(videoURL) {
						videos = append(videos, VideoItem{
							ID:         review.ID,
							URL:        videoURL,
							YoutubeURL: videoURL,
							Title:      title,
						})
					}
				}
			}
		}
	}

	// Return the videos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":   videos,
		"videos": videos, // Support both field names for compatibility
	})
}

// isYouTubeURL checks if a URL is a valid YouTube URL
func isYouTubeURL(url string) bool {
	if url == "" {
		return false
	}
	url = strings.ToLower(url)
	return strings.Contains(url, "youtube.com") || strings.Contains(url, "youtu.be")
}

// getReviewTitle returns a title for the review
func getReviewTitle(review *entities.ProductReview) string {
	if review.Review != nil && *review.Review != "" {
		// Use first 50 characters of review as title
		reviewText := *review.Review
		if len(reviewText) > 50 {
			return reviewText[:50] + "..."
		}
		return reviewText
	}
	return "Product Review Video"
}
