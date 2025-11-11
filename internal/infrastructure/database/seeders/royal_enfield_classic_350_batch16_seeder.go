package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RoyalEnfieldClassic350ReviewBatch16 handles seeding Royal Enfield Classic 350 product review and translation
type RoyalEnfieldClassic350ReviewBatch16 struct {
	BaseSeeder
}

// NewRoyalEnfieldClassic350ReviewBatch16Seeder creates a new RoyalEnfieldClassic350ReviewBatch16
func NewRoyalEnfieldClassic350ReviewBatch16Seeder() *RoyalEnfieldClassic350ReviewBatch16 {
	return &RoyalEnfieldClassic350ReviewBatch16{BaseSeeder: BaseSeeder{name: "Royal Enfield Classic 350 Batch16 Review"}}
}

// Seed implements the Seeder interface
func (s *RoyalEnfieldClassic350ReviewBatch16) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Royal Enfield Classic 350").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("royal enfield classic 350 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding royal enfield classic 350 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Royal Enfield Classic 350 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Royal Enfield Classic 350 Review Bangladesh 2024 - Timeless Cruiser Heritage</h1>
<p class="summary-text">The Royal Enfield Classic 350 is a 349cc retro-cruiser icon combining classic motorcycle heritage with modern reliability. Priced around ৳3,85,000, it's the ultimate choice for riders seeking authentic vintage styling, comfortable ergonomics, and legendary British engineering trusted across South Asia for decades.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Royal Enfield Classic 350 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">349cc Engine:</strong> <span class="highlight-value">Bullet-proof reliability</span></li>
<li class="highlight-item"><strong class="highlight-label">Retro Design:</strong> <span class="highlight-value">Vintage soul</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfort First:</strong> <span class="highlight-value">Easy riding</span></li>
<li class="highlight-item"><strong class="highlight-label">Brand Heritage:</strong> <span class="highlight-value">British legend</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Royal Enfield Classic 350 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Iconic Retro:</strong> <span class="pro-description">Timeless appeal</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Proven legendary</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable:</strong> <span class="pro-description">Easy cruising</span></li>
<li class="pro-item"><strong class="pro-label">Affordable:</strong> <span class="pro-description">Good value</span></li>
<li class="pro-item"><strong class="pro-label">Community:</strong> <span class="pro-description">Strong network</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Royal Enfield Classic 350 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Tech:</strong> <span class="con-description">Minimal electronics</span></li>
<li class="con-item"><strong class="con-label">Modest Power:</strong> <span class="con-description">Not sporty</span></li>
<li class="con-item"><strong class="con-label">Fuel Tank:</strong> <span class="con-description">Regular refills needed</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Royal Enfield Classic 350 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,85,000 - Mid-range retro</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳6-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">32-36 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Royal Enfield Classic 350 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Reliable classic</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Stable cruiser</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Easy ergonomics</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Iconic retro</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Great heritage</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Royal Enfield Classic 350?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,85,000, the Classic 350 is perfect for riders seeking vintage charm, proven reliability, comfortable cruising, and the timeless appeal of British motorcycle heritage with modern practicality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For retro enthusiasts and long tours</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating royal enfield classic 350 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Royal Enfield Classic 350 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">রয়্যাল এনফিল্ড ক্লাসিক 350 রিভিউ বাংলাদেশ 2024 - চিরন্তন ক্রুজার ঐতিহ্য</h1>
<p class="summary-text">রয়্যাল এনফিল্ড ক্লাসিক 350 একটি 349cc রেট্রো-ক্রুজার আইকন যা ক্লাসিক মোটরসাইকেল ঐতিহ্যকে আধুনিক নির্ভরযোগ্যতার সাথে একত্রিত করে। ৳3,85,000 টাকায় মূল্যায়িত, এটি ক্রেতাদের জন্য চূড়ান্ত পছন্দ যারা প্রামাণিক ভিন্টেজ স্টাইলিং, আরামদায়ক ergonomics এবং কিংবদন্তি ব্রিটিশ প্রকৌশল খুঁজছেন।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রয়্যাল এনফিল্ড ক্লাসিক 350 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">349cc ইঞ্জিন:</strong> <span class="highlight-value">বুলেট-প্রমাণ নির্ভরযোগ্যতা</span></li>
<li class="highlight-item"><strong class="highlight-label">রেট্রো ডিজাইন:</strong> <span class="highlight-value">ভিন্টেজ আত্মা</span></li>
<li class="highlight-item"><strong class="highlight-label">আরাম প্রথম:</strong> <span class="highlight-value">সহজ চড়া</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্র্যান্ড ঐতিহ্য:</strong> <span class="highlight-value">ব্রিটিশ কিংবদন্তি</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রয়্যাল এনফিল্ড ক্লাসিক 350 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আইকনিক রেট্রো:</strong> <span class="pro-description">কালজয়ী আবেদন</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">প্রমাণিত কিংবদন্তি</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক:</strong> <span class="pro-description">সহজ ক্রুজিং</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী:</strong> <span class="pro-description">ভাল মূল্য</span></li>
<li class="pro-item"><strong class="pro-label">কমিউনিটি:</strong> <span class="pro-description">শক্তিশালী নেটওয়ার্ক</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রয়্যাল এনফিল্ড ক্লাসিক 350 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক প্রযুক্তি:</strong> <span class="con-description">ন্যূনতম ইলেকট্রনিক্স</span></li>
<li class="con-item"><strong class="con-label">মডেস্ট পাওয়ার:</strong> <span class="con-description">খেলাধুলামূলক নয়</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি ট্যাঙ্ক:</strong> <span class="con-description">নিয়মিত রিফিল প্রয়োজন</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: রয়্যাল এনফিল্ড ক্লাসিক 350 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,85,000 টাকায়, ক্লাসিক 350 চালকদের জন্য নিখুঁত যারা ভিন্টেজ আকর্ষণ, প্রমাণিত নির্ভরযোগ্যতা, আরামদায়ক ক্রুজিং এবং ব্রিটিশ মোটরসাইকেল ঐতিহ্যের কালজয়ী আবেদন খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- রেট্রো উত্সাহী এবং দীর্ঘ ভ্রমণের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Royal Enfield Classic 350 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Royal Enfield Classic 350\n")

	return nil
}

// GetName returns the seeder name
func (s *RoyalEnfieldClassic350ReviewBatch16) GetName() string {
	return "RoyalEnfieldClassic350ReviewBatch16"
}
