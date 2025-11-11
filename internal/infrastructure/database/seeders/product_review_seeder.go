package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// ProductReviewSeeder handles seeding product reviews for motorcycles
type ProductReviewSeeder struct {
	BaseSeeder
}

// NewProductReviewSeeder creates a new ProductReviewSeeder
func NewProductReviewSeeder() *ProductReviewSeeder {
	return &ProductReviewSeeder{BaseSeeder: BaseSeeder{name: "Product Reviews"}}
}

// Seed implements the Seeder interface
func (prs *ProductReviewSeeder) Seed(db *gorm.DB) error {
	// Get the admin user (UserID = 1)
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	reviews := []struct {
		ProductName       string
		Rating            float32
		Review            string
		AdditionalDetails map[string]interface{}
	}{
		{
			ProductName: "Honda CD 70",
			Rating:      4.2,
			Review: `<div class="review-content">
<h3>Pros</h3>
<ul>
<li>Extremely fuel efficient - delivers 60-70 km/l</li>
<li>Very affordable and easy to maintain</li>
<li>Reliable Honda engine with proven track record</li>
<li>Lightweight and easy to handle for beginners</li>
<li>Low running costs</li>
<li>Good resale value</li>
<li>Perfect for daily commuting in city traffic</li>
</ul>

<h3>Cons</h3>
<ul>
<li>Underpowered for highway riding</li>
<li>Basic design with minimal features</li>
<li>Not suitable for tall riders</li>
<li>Lacks modern styling</li>
<li>No digital instrumentation</li>
<li>Average build quality compared to newer models</li>
<li>Limited top speed</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=dQw4w9WgXcQ",
					"https://www.youtube.com/watch?v=xYZ123AbCdE",
				},
				"verdict":      "Best budget commuter bike with unmatched fuel economy",
				"recommended":  true,
				"use_case":     "Daily commute, budget-conscious riders",
				"rating_split": map[string]float32{"performance": 3.5, "fuel_economy": 5.0, "comfort": 4.0, "value": 5.0},
			},
		},
		{
			ProductName: "Honda Livo",
			Rating:      4.3,
			Review: `<div class="review-content">
<h3>Pros</h3>
<ul>
<li>Excellent fuel efficiency of 55-65 km/l</li>
<li>Comfortable riding position for daily commute</li>
<li>Modern and attractive design</li>
<li>Good build quality typical of Honda</li>
<li>Smooth and refined engine</li>
<li>Adequate power for city riding</li>
<li>Low maintenance requirements</li>
<li>Comfortable seat for both rider and pillion</li>
</ul>

<h3>Cons</h3>
<ul>
<li>Slightly underpowered compared to competition</li>
<li>Average performance on highways</li>
<li>Basic instrumentation cluster</li>
<li>Suspension could be softer</li>
<li>Limited storage options</li>
<li>Higher price compared to CD 70</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=abc123XYZ456",
					"https://www.youtube.com/watch?v=def789GHI012",
				},
				"verdict":      "Well-balanced commuter with modern looks",
				"recommended":  true,
				"use_case":     "Daily commute, style-conscious riders",
				"rating_split": map[string]float32{"performance": 4.0, "fuel_economy": 4.5, "comfort": 4.5, "value": 4.2},
			},
		},
		{
			ProductName: "Honda CB125F",
			Rating:      4.4,
			Review: `<div class="review-content">
<h3>Pros</h3>
<ul>
<li>Peppy 125cc engine with good power delivery</li>
<li>Impressive fuel economy of 50-60 km/l</li>
<li>Sporty and aggressive styling</li>
<li>Honda's legendary reliability and build quality</li>
<li>Comfortable for both city and highway riding</li>
<li>Good handling and maneuverability</li>
<li>Digital-analog instrument cluster</li>
<li>Strong braking performance</li>
</ul>

<h3>Cons</h3>
<ul>
<li>Premium pricing compared to competitors</li>
<li>Seat could be more comfortable for long rides</li>
<li>No ABS option in base variant</li>
<li>Ground clearance could be better</li>
<li>Limited under-seat storage</li>
<li>Headlight brightness could be improved</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=CB125Review01",
					"https://www.youtube.com/watch?v=CB125Test2023",
				},
				"verdict":      "Excellent all-rounder for daily riding",
				"recommended":  true,
				"use_case":     "City and highway commute, enthusiast riders",
				"rating_split": map[string]float32{"performance": 4.5, "fuel_economy": 4.3, "comfort": 4.2, "value": 4.4},
			},
		},
		{
			ProductName: "Honda CB Shine",
			Rating:      4.5,
			Review: `<div class="review-content">
<h3>Pros</h3>
<ul>
<li>Refined and powerful 125cc engine</li>
<li>Excellent build quality and premium feel</li>
<li>Very comfortable riding position</li>
<li>Great fuel efficiency of 55-65 km/l</li>
<li>Smooth gear shifts and clutch operation</li>
<li>Good suspension for Indian road conditions</li>
<li>Reliable Honda brand reputation</li>
<li>Low maintenance costs</li>
<li>Strong resale value</li>
</ul>

<h3>Cons</h3>
<ul>
<li>Conservative design may not appeal to younger riders</li>
<li>Premium pricing</li>
<li>Lacks modern features like LED lights (in base variant)</li>
<li>Not the most exciting bike to ride</li>
<li>Heavy compared to some competitors</li>
<li>Basic instrument console</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=ShineReview123",
					"https://www.youtube.com/watch?v=ShineTestRide",
				},
				"verdict":      "Premium commuter with exceptional refinement",
				"recommended":  true,
				"use_case":     "Mature riders, long-distance commute",
				"rating_split": map[string]float32{"performance": 4.4, "fuel_economy": 4.6, "comfort": 4.8, "value": 4.3},
			},
		},
		{
			ProductName: "Hero Splendor Plus",
			Rating:      4.1,
			Review: `<div class="review-content">
<h3>Pros</h3>
<ul>
<li>Outstanding fuel efficiency of 65-75 km/l</li>
<li>Most affordable 100cc motorcycle</li>
<li>Extremely low maintenance costs</li>
<li>Proven reliability over decades</li>
<li>Lightweight and easy to handle</li>
<li>Wide service network across Bangladesh</li>
<li>Good for beginners</li>
<li>Strong resale value</li>
</ul>

<h3>Cons</h3>
<ul>
<li>Very basic design and features</li>
<li>Underpowered engine for highway use</li>
<li>Average build quality</li>
<li>Outdated styling</li>
<li>Uncomfortable seat on long rides</li>
<li>Poor braking performance in wet conditions</li>
<li>Lacks modern amenities</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=SplendorReview",
					"https://www.youtube.com/watch?v=HeroSplendor2023",
				},
				"verdict":      "Best value for money commuter motorcycle",
				"recommended":  true,
				"use_case":     "Budget buyers, first-time riders, daily commute",
				"rating_split": map[string]float32{"performance": 3.5, "fuel_economy": 5.0, "comfort": 3.8, "value": 5.0},
			},
		},
		{
			ProductName: "Yamaha YZF R15 V4",
			Rating:      4.7,
			Review: `<div class="review-content">
<h3>Pros</h3>
<ul>
<li>Powerful and rev-happy 155cc engine</li>
<li>Stunning sporty design with aerodynamic fairing</li>
<li>Excellent handling and cornering ability</li>
<li>Premium build quality and attention to detail</li>
<li>Full digital LCD instrument cluster with connectivity</li>
<li>Dual-channel ABS for superior safety</li>
<li>LED headlights and taillights</li>
<li>Thrilling riding experience</li>
<li>Strong brand image</li>
</ul>

<h3>Cons</h3>
<ul>
<li>Expensive compared to competitors</li>
<li>Aggressive riding position not comfortable for daily commute</li>
<li>Lower fuel efficiency (40-45 km/l)</li>
<li>High maintenance costs</li>
<li>Not suitable for pillion comfort</li>
<li>Expensive spare parts</li>
<li>Limited under-seat storage</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=R15V4Review",
					"https://www.youtube.com/watch?v=YamahaR15Test",
				},
				"verdict":      "Best sporty motorcycle for enthusiasts",
				"recommended":  true,
				"use_case":     "Sport riding, track days, young enthusiasts",
				"rating_split": map[string]float32{"performance": 4.9, "fuel_economy": 3.8, "comfort": 3.5, "value": 4.5},
			},
		},
		{
			ProductName: "Bajaj Pulsar 150",
			Rating:      4.3,
			Review: `<div class="review-content">
<h3>Pros</h3>
<ul>
<li>Powerful 150cc engine with strong mid-range</li>
<li>Aggressive and muscular styling</li>
<li>Good value for money</li>
<li>Comfortable riding position for long rides</li>
<li>Decent fuel economy of 45-50 km/l</li>
<li>Strong and sturdy build</li>
<li>Good handling characteristics</li>
<li>Wide availability of spare parts</li>
</ul>

<h3>Cons</h3>
<ul>
<li>Vibrations at high RPMs</li>
<li>Average build quality compared to Japanese bikes</li>
<li>Outdated single-channel ABS</li>
<li>Heavy clutch operation</li>
<li>Basic instrument cluster</li>
<li>Suspension could be better tuned</li>
<li>Lower resale value</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=Pulsar150Review",
					"https://www.youtube.com/watch?v=BajajPulsar2023",
				},
				"verdict":      "Affordable performance motorcycle",
				"recommended":  true,
				"use_case":     "Budget performance, daily commute with power",
				"rating_split": map[string]float32{"performance": 4.5, "fuel_economy": 4.0, "comfort": 4.3, "value": 4.5},
			},
		},
		{
			ProductName: "Suzuki Gixxer SF",
			Rating:      4.4,
			Review: `<div class="review-content">
<h3>Pros</h3>
<ul>
<li>Refined and powerful 155cc engine</li>
<li>Attractive full-faired sporty design</li>
<li>Excellent handling and stability</li>
<li>Good build quality</li>
<li>Comfortable for both city and highway</li>
<li>Fuel injection for better efficiency (48-52 km/l)</li>
<li>Single-channel ABS for safety</li>
<li>Fully digital instrument console</li>
</ul>

<h3>Cons</h3>
<ul>
<li>Premium pricing</li>
<li>Limited service network compared to Honda/Hero</li>
<li>Headlight brightness could be better</li>
<li>Slightly cramped for tall riders</li>
<li>Average pillion comfort</li>
<li>Expensive spare parts</li>
<li>Less ground clearance</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=GixxerSFReview",
					"https://www.youtube.com/watch?v=SuzukiGixxer155",
				},
				"verdict":      "Well-balanced sport-commuter motorcycle",
				"recommended":  true,
				"use_case":     "Sporty commuting, weekend rides",
				"rating_split": map[string]float32{"performance": 4.5, "fuel_economy": 4.2, "comfort": 4.3, "value": 4.3},
			},
		},
	}

	for _, reviewData := range reviews {
		// Find the product by name
		var product models.ProductModel
		if err := db.Where("name = ?", reviewData.ProductName).First(&product).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("   ⚠️  Product '%s' not found, skipping review\n", reviewData.ProductName)
				continue
			}
			return fmt.Errorf("error finding product '%s': %w", reviewData.ProductName, err)
		}

		// Check if review already exists
		var existing models.ProductReviewModel
		err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existing).Error
		if err == nil {
			fmt.Printf("   ℹ️  Review for '%s' already exists, skipping\n", reviewData.ProductName)
			continue
		}

		// Marshal additional details to JSON
		additionalDetailsJSON, err := json.Marshal(reviewData.AdditionalDetails)
		if err != nil {
			return fmt.Errorf("error marshaling additional details for '%s': %w", reviewData.ProductName, err)
		}

		// Create new review
		review := &models.ProductReviewModel{
			ProductID:         product.ID,
			UserID:            adminUser.ID,
			Rating:            reviewData.Rating,
			Reviews:           &reviewData.Review,
			AdditionalDetails: additionalDetailsJSON,
			Priority:          1,
			Status:            true, // status = 1 means active/published
		}

		if err := db.Create(review).Error; err != nil {
			return fmt.Errorf("error creating review for '%s': %w", reviewData.ProductName, err)
		}

		fmt.Printf("   ✅ Created review for '%s' (Rating: %.1f/5.0)\n", reviewData.ProductName, reviewData.Rating)
	}

	return nil
}
