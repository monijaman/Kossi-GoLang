package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BanglalinkReviewSeeder handles seeding product reviews for Banglalink mobile operator
type BanglalinkReviewSeeder struct {
	BaseSeeder
}

// NewBanglalinkReviewSeeder creates a new BanglalinkReviewSeeder
func NewBanglalinkReviewSeeder() *BanglalinkReviewSeeder {
	return &BanglalinkReviewSeeder{BaseSeeder: BaseSeeder{name: "Banglalink Reviews"}}
}

// Seed implements the Seeder interface
func (brs *BanglalinkReviewSeeder) Seed(db *gorm.DB) error {
	// Get the admin user (UserID = 1)
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	// Find Banglalink product
	var banglalinkProduct models.ProductModel
	if err := db.Where("slug = ?", "banglalink").First(&banglalinkProduct).Error; err != nil {
		return fmt.Errorf("Banglalink product not found, please run mobile operator product seeder first: %w", err)
	}

	// Check if review already exists for this product
	var existingReview models.ProductReviewModel
	result := db.Where("product_id = ? AND user_id = ?", banglalinkProduct.ID, adminUser.ID).First(&existingReview)

	if result.Error == gorm.ErrRecordNotFound {
		// Create comprehensive review
		reviewText := `<div class="review-content">
<h2>Banglalink Mobile Operator Bangladesh - Comprehensive Review 2024</h2>

<h3>Overview</h3>
<p>Banglalink is one of Bangladesh's leading mobile network operators, offering reliable voice, data, and digital services. As a customer-centric telecom provider, Banglalink has evolved significantly to meet the growing demands of modern mobile users in Bangladesh.</p>

<h3>Network Quality & Coverage</h3>
<ul>
<li><strong>4G LTE Coverage:</strong> Excellent 4G coverage in major cities including Dhaka, Chittagong, Sylhet, and Khulna with consistent expansion to rural areas</li>
<li><strong>Network Reliability:</strong> Generally stable network performance with competitive data speeds averaging 10-15 Mbps in urban areas</li>
<li><strong>Voice Quality:</strong> Clear voice calls with minimal drop-off rates in covered areas</li>
<li><strong>2G/3G Backup:</strong> Reliable fallback coverage for legacy devices and remote areas</li>
</ul>

<h3>Pros - Why Choose Banglalink</h3>
<ul>
<li>Affordable call rates - Competitive pricing for local and international calls</li>
<li>Data-friendly bundles - Attractive internet packages for daily browsing and streaming</li>
<li>Social media packages - Dedicated bundles for Facebook, WhatsApp, and other platforms</li>
<li>Frequent promotions - Regular special offers and bonus data deals</li>
<li>Good customer service - Multiple support channels including helpline and app</li>
<li>Banglalink App - User-friendly mobile application for balance checking and package management</li>
<li>Digital payment options - Support for mobile banking and digital wallets</li>
<li>Flexible plans - Various prepaid and postpaid options to suit different budgets</li>
<li>Roaming packages - Competitive international roaming rates</li>
<li>Emergency balance feature - Borrow balance option for emergency calls</li>
</ul>

<h3>Cons - Areas for Improvement</h3>
<ul>
<li>Network congestion during peak hours in some urban areas</li>
<li>Occasional data speed throttling on certain plans</li>
<li>Customer service wait times can be long during peak periods</li>
<li>Roaming rates still higher compared to some competitors</li>
<li>Limited coverage in some remote areas compared to larger competitors</li>
<li>SIM activation process could be simplified</li>
<li>Some packages have strict validity periods</li>
</ul>

<h3>Active Data Packages & Bundles</h3>
<ul>
<li><strong>Daily Data Packs:</strong> 100MB-500MB daily bundles at affordable rates</li>
<li><strong>Weekly Data Packs:</strong> 1GB-3GB weekly bundles perfect for regular users</li>
<li><strong>Monthly Unlimited Data:</strong> True unlimited data plans for heavy users</li>
<li><strong>Social Media Bundles:</strong> Facebook, WhatsApp, YouTube special packages</li>
<li><strong>Night Data Packages:</strong> Discounted data available during late night hours</li>
<li><strong>OTT Packages:</strong> Discounted data for streaming services like YouTube, Netflix</li>
</ul>

<h3>Active Voice & SMS Packages</h3>
<ul>
<li><strong>Local Call Bundles:</strong> Per-minute rates or daily/weekly call packages</li>
<li><strong>SMS Bundles:</strong> Affordable local and international SMS packages</li>
<li><strong>Roaming Packages:</strong> Special offers for traveling within and outside Bangladesh</li>
<li><strong>International Call Bundles:</strong> Discounted rates to popular destinations</li>
<li><strong>FnF (Friends and Family):</strong> Reduced rates to frequently called numbers</li>
</ul>

<h3>Banglalink Money & Financial Services</h3>
<ul>
<li>Mobile wallet for digital transactions</li>
<li>Money transfer services</li>
<li>Utility bill payments</li>
<li>Fast merchant payment system</li>
</ul>

<h3>Customer Support Quality</h3>
<ul>
<li>24/7 customer helpline support</li>
<li>USSD codes for quick balance and package management</li>
<li>Mobile app for self-service features</li>
<li>Social media support through Facebook and Twitter</li>
<li>Physical service centers across major cities</li>
</ul>

<h3>Pricing & Value for Money</h3>
<p>Banglalink offers competitive pricing across all segments. Their prepaid packages are value-oriented, making them accessible to budget-conscious users. However, postpaid plans are relatively higher than competitors. Overall, good value for money considering network quality and service offerings.</p>

<h3>Verdict</h3>
<p>Banglalink is a reliable choice for mobile services in Bangladesh, especially for users seeking affordable data and call packages with decent network coverage. Their frequent promotions and user-friendly app make them competitive. Ideal for daily commuters and casual users who need flexible, no-contract services.</p>

<h3>Recommendation</h3>
<p><strong>Recommended for:</strong> Budget-conscious users, daily data consumers, frequent travelers within Bangladesh, small business owners, students</p>
<p><strong>Not ideal for:</strong> Users requiring premium customer service, those in areas with limited 4G coverage, or those seeking the absolute cheapest rates</p>

<h3>Overall Rating Breakdown</h3>
<ul>
<li>Network Quality: 4/5</li>
<li>Coverage Area: 4/5</li>
<li>Call Quality: 4/5</li>
<li>Data Speed: 4/5</li>
<li>Package Value: 4.5/5</li>
<li>Customer Service: 3.5/5</li>
<li>Pricing: 4/5</li>
<li>Overall: 4.1/5</li>
</ul>

<h3>Final Thoughts</h3>
<p>Banglalink continues to be one of the best choices for mobile operators in Bangladesh. Their focus on affordable packages, regular promotions, and digital innovation through the Banglalink app makes them stand out. While there's room for improvement in customer service and coverage expansion, their commitment to customer satisfaction is evident. Whether you're a light user or a heavy data consumer, Banglalink has something to offer at competitive prices.</p>
</div>`

		additionalDetails := map[string]interface{}{
			"youtube_reviews": []string{
				"https://www.youtube.com/results?search_query=Banglalink+review+bangladesh",
				"https://www.youtube.com/results?search_query=Banglalink+packages+comparison",
			},
			"verdict":     "Best budget-friendly mobile operator with attractive data packages",
			"recommended": true,
			"use_case":    "Daily users, students, budget-conscious consumers, travelers",
			"rating_split": map[string]float32{
				"network_quality":  4.0,
				"coverage":         4.0,
				"call_quality":     4.0,
				"data_speed":       4.0,
				"package_value":    4.5,
				"customer_service": 3.5,
				"pricing":          4.0,
			},
			"seo_keywords": []string{
				"Banglalink mobile operator Bangladesh",
				"Banglalink packages 2024",
				"Banglalink data plans",
				"Banglalink reviews",
				"best mobile operator Bangladesh",
				"Banglalink 4G coverage",
				"Banglalink call rates",
				"Banglalink internet packages",
			},
			"pros": []string{
				"Affordable call rates",
				"Attractive data bundles",
				"Frequent promotions",
				"Good 4G coverage",
				"Mobile app available",
				"Digital payment support",
				"Emergency balance feature",
				"Flexible plans",
			},
			"cons": []string{
				"Peak hour congestion",
				"Limited coverage in remote areas",
				"Long customer service wait times",
				"Higher roaming rates",
				"Occasional speed throttling",
			},
		}

		additionalDetailsJSON, err := json.Marshal(additionalDetails)
		if err != nil {
			return fmt.Errorf("failed to marshal additional details: %w", err)
		}

		review := models.ProductReviewModel{
			ProductID:         banglalinkProduct.ID,
			UserID:            adminUser.ID,
			Rating:            "4.1",
			Reviews:           &reviewText,
			AdditionalDetails: additionalDetailsJSON,
			Priority:          1,
			Status:            true,
		}

		if err := db.Create(&review).Error; err != nil {
			return fmt.Errorf("failed to create Banglalink review: %w", err)
		}

		fmt.Printf("✓ Created Banglalink product review (ID: %d)\n", review.ID)
	}

	return nil
}
