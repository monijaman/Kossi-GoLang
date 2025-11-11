package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// GaselliGSRX150ReviewBatch17 handles seeding Gazelli GSRX150 product review and translation
type GaselliGSRX150ReviewBatch17 struct {
	BaseSeeder
}

// NewGaselliGSRX150ReviewBatch17Seeder creates a new GaselliGSRX150ReviewBatch17
func NewGaselliGSRX150ReviewBatch17Seeder() *GaselliGSRX150ReviewBatch17 {
	return &GaselliGSRX150ReviewBatch17{BaseSeeder: BaseSeeder{name: "Gazelli GSRX150 Batch17 Review"}}
}

// Seed implements the Seeder interface
func (s *GaselliGSRX150ReviewBatch17) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Gazelli GSRX150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("gazelli gsrx150 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding gazelli gsrx150 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Gazelli GSRX150 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Gazelli GSRX150 Review Bangladesh 2024 - Youth Sport Commuter</h1>
<p class="summary-text">The Gazelli GSRX150 is a 150cc youth-oriented sport commuter combining aggressive styling with accessible performance and affordability. Priced around ৳2,15,000, it delivers sporty character, nimble handling, modern aesthetics, and outstanding value for young riders seeking style and fun on a budget.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Gazelli GSRX150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc Sport:</strong> <span class="highlight-value">Youth appeal</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Modern lines</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Price:</strong> <span class="highlight-value">Great value</span></li>
<li class="highlight-item"><strong class="highlight-label">Fun Riding:</strong> <span class="highlight-value">Sporty feel</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Gazelli GSRX150 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Aggressive Design:</strong> <span class="pro-description">Youth focused</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Engine:</strong> <span class="pro-description">Fun responsive</span></li>
<li class="pro-item"><strong class="pro-label">Great Handling:</strong> <span class="pro-description">Nimble feel</span></li>
<li class="pro-item"><strong class="pro-label">Budget Price:</strong> <span class="pro-description">Excellent value</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">Contemporary tech</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Gazelli GSRX150 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Prestige:</strong> <span class="con-description">New brand</span></li>
<li class="con-item"><strong class="con-label">Build Quality:</strong> <span class="con-description">Budget construction</span></li>
<li class="con-item"><strong class="con-label">Service Network:</strong> <span class="con-description">Developing</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Gazelli GSRX150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,15,000 - Budget sport</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳5-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">35-40 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Gazelli GSRX150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Responsive power</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Nimble agile</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Modern style</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent budget</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Modern equipped</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Gazelli GSRX150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,15,000, the GSRX150 is perfect for young riders wanting a sporty 150cc commuter with modern styling, fun performance, great handling, and exceptional value without premium costs.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget-conscious young riders</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.4,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating gazelli gsrx150 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Gazelli GSRX150 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">গ্যাজেলি GSRX150 রিভিউ বাংলাদেশ 2024 - যুব স্পোর্ট যাত্রী</h1>
<p class="summary-text">গ্যাজেলি GSRX150 একটি 150cc যুব-ভিত্তিক স্পোর্ট যাত্রী যা আক্রমণাত্মক স্টাইলিং এবং সাশ্রয়ী মূল্যকে একত্রিত করে। ৳2,15,000 টাকায় মূল্যায়িত, এটি খেলাধুলামূলক চরিত্র, নমনীয় হ্যান্ডলিং, আধুনিক নান্দনিকতা এবং চমৎকার মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">গ্যাজেলি GSRX150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc স্পোর্ট:</strong> <span class="highlight-value">যুব আবেদন</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="highlight-value">আধুনিক লাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট মূল্য:</strong> <span class="highlight-value">দুর্দান্ত মূল্য</span></li>
<li class="highlight-item"><strong class="highlight-label">মজাদার চড়া:</strong> <span class="highlight-value">খেলাধুলামূলক অনুভূতি</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">গ্যাজেলি GSRX150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="pro-description">যুব ফোকাস</span></li>
<li class="pro-item"><strong class="pro-label">খেলাধুলামূলক ইঞ্জিন:</strong> <span class="pro-description">মজাদার প্রতিক্রিয়াশীল</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত হ্যান্ডলিং:</strong> <span class="pro-description">নমনীয় অনুভূতি</span></li>
<li class="pro-item"><strong class="pro-label">বাজেট মূল্য:</strong> <span class="pro-description">চমৎকার মূল্য</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক বৈশিষ্ট্য:</strong> <span class="pro-description">সমসাময়িক প্রযুক্তি</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">গ্যাজেলি GSRX150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত প্রতিপত্তি:</strong> <span class="con-description">নতুন ব্র্যান্ড</span></li>
<li class="con-item"><strong class="con-label">বিল্ড গুণমান:</strong> <span class="con-description">বাজেট নির্মাণ</span></li>
<li class="con-item"><strong class="con-label">সেবা নেটওয়ার্ক:</strong> <span class="con-description">বিকাশশীল</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: গ্যাজেলি GSRX150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,15,000 টাকায়, GSRX150 তরুণ চালকদের জন্য নিখুঁত যারা একটি খেলাধুলামূলক 150cc যাত্রী চান আধুনিক স্টাইলিং এবং মজাদার কর্মক্ষমতা সহ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট-সচেতন তরুণ চালকদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Gazelli GSRX150 already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Gazelli GSRX150\n")

	return nil
}

// GetName returns the seeder name
func (s *GaselliGSRX150ReviewBatch17) GetName() string {
	return "GaselliGSRX150ReviewBatch17"
}
