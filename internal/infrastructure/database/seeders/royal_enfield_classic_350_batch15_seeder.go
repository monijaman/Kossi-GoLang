package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RoyalEnfieldClassic350Batch15 handles seeding Royal Enfield Classic 350 product review and translation
type RoyalEnfieldClassic350Batch15 struct {
	BaseSeeder
}

// NewRoyalEnfieldClassic350Batch15Seeder creates a new RoyalEnfieldClassic350Batch15
func NewRoyalEnfieldClassic350Batch15Seeder() *RoyalEnfieldClassic350Batch15 {
	return &RoyalEnfieldClassic350Batch15{BaseSeeder: BaseSeeder{name: "Royal Enfield Classic 350 Batch15"}}
}

// Seed implements the Seeder interface
func (s *RoyalEnfieldClassic350Batch15) Seed(db *gorm.DB) error {
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
<h1 class="review-main-title">Royal Enfield Classic 350 Review Bangladesh 2024 - Iconic Cruiser Charm</h1>
<p class="summary-text">The Royal Enfield Classic 350 is an iconic retro-styled cruiser with timeless design, thumping 350cc engine, comfortable seating, and legendary reliability. Priced around ৳3,30,000, it's the classic cruiser choice for riders seeking authentic vintage character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Royal Enfield Classic 350 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Iconic Design:</strong> <span class="highlight-value">Timeless retro styling</span></li>
<li class="highlight-item"><strong class="highlight-label">350cc Thump:</strong> <span class="highlight-value">Legendary engine character</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seating:</strong> <span class="highlight-value">Long-distance friendly</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">38-42 km/l average</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Royal Enfield Classic 350 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Iconic Design:</strong> <span class="pro-description">Timeless classic cruiser</span></li>
<li class="pro-item"><strong class="pro-label">Thumping Engine:</strong> <span class="pro-description">Legendary character</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seating:</strong> <span class="pro-description">Long-ride friendly</span></li>
<li class="pro-item"><strong class="pro-label">Easy Customization:</strong> <span class="pro-description">Huge aftermarket support</span></li>
<li class="pro-item"><strong class="pro-label">Legendary Reliability:</strong> <span class="pro-description">Proven cruiser classic</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Royal Enfield Classic 350 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Fuel Efficiency:</strong> <span class="con-description">Lower mileage 38-42 km/l</span></li>
<li class="con-item"><strong class="con-label">Handling:</strong> <span class="con-description">Cruiser-oriented geometry</span></li>
<li class="con-item"><strong class="con-label">Performance:</strong> <span class="con-description">Not sport-oriented</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Royal Enfield Classic 350 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,30,000 - Iconic cruiser</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳6-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">38-42 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Royal Enfield Classic 350 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Iconic timeless</span></li>
<li class="rating-item"><strong class="rating-label">Character:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Legendary thump</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Cruiser friendly</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Proven classic</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Premium classic</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Royal Enfield Classic 350?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,30,000, the Royal Enfield Classic 350 is the choice for riders seeking timeless cruiser design, legendary thumping character, comfortable seating, and the authentic vintage motorcycle experience that transcends trends.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For classic cruiser enthusiasts</span></p>
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
<h1 class="review-main-title">রয়্যাল এনফিল্ড ক্লাসিক 350 রিভিউ বাংলাদেশ 2024 - আইকনিক ক্রুজার আকর্ষণ</h1>
<p class="summary-text">রয়্যাল এনফিল্ড ক্লাসিক 350 একটি আইকনিক রেট্রো-স্টাইলড ক্রুজার যা সময়োচিত ডিজাইন, থামপিং 350cc ইঞ্জিন, আরামদায়ক আসন এবং কিংবদন্তি নির্ভরযোগ্যতা সহ আসে। ৳3,30,000 টাকায় মূল্যায়িত, এটি খাঁটি ভিন্টেজ চরিত্র খুঁজছেন রাইডারদের জন্য ক্লাসিক ক্রুজার পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রয়্যাল এনফিল্ড ক্লাসিক 350 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">আইকনিক ডিজাইন:</strong> <span class="highlight-value">সময়োচিত রেট্রো স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">350cc থামপ:</strong> <span class="highlight-value">কিংবদন্তি ইঞ্জিন চরিত্র</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক আসন:</strong> <span class="highlight-value">দীর্ঘ-দূরত্ব বান্ধব</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">38-42 km/l গড়</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রয়্যাল এনফিল্ড ক্লাসিক 350 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আইকনিক ডিজাইন:</strong> <span class="pro-description">সময়োচিত ক্লাসিক ক্রুজার</span></li>
<li class="pro-item"><strong class="pro-label">থামপিং ইঞ্জিন:</strong> <span class="pro-description">কিংবদন্তি চরিত্র</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক আসন:</strong> <span class="pro-description">দীর্ঘ-যাত্রা বান্ধব</span></li>
<li class="pro-item"><strong class="pro-label">সহজ কাস্টমাইজেশন:</strong> <span class="pro-description">বিশাল আফটারমার্কেট সহায়তা</span></li>
<li class="pro-item"><strong class="pro-label">কিংবদন্তি নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত ক্রুজার ক্লাসিক</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রয়্যাল এনফিল্ড ক্লাসিক 350 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">জ্বালানী দক্ষতা:</strong> <span class="con-description">কম মাইলেজ 38-42 km/l</span></li>
<li class="con-item"><strong class="con-label">হ্যান্ডলিং:</strong> <span class="con-description">ক্রুজার-ভিত্তিক জ্যামিতি</span></li>
<li class="con-item"><strong class="con-label">কর্মক্ষমতা:</strong> <span class="con-description">ক্রীড়া-ভিত্তিক নয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: রয়্যাল এনফিল্ড ক্লাসিক 350 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,30,000 টাকায়, রয়্যাল এনফিল্ড ক্লাসিক 350 রাইডারদের পছন্দ যারা সময়োচিত ক্রুজার ডিজাইন, কিংবদন্তি থামপিং চরিত্র, আরামদায়ক আসন এবং খাঁটি ভিন্টেজ মোটরসাইকেল অভিজ্ঞতা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্লাসিক ক্রুজার উত্সাহীদের জন্য</span></p>
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
func (s *RoyalEnfieldClassic350Batch15) GetName() string {
	return "RoyalEnfieldClassic350Batch15"
}
