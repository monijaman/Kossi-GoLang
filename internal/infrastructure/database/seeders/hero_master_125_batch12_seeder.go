package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroMaster125Review handles seeding Hero Master 125 product review and translation
type HeroMaster125Review struct {
	BaseSeeder
}

// NewHeroMaster125ReviewSeeder creates a new HeroMaster125Review
func NewHeroMaster125ReviewSeeder() *HeroMaster125Review {
	return &HeroMaster125Review{BaseSeeder: BaseSeeder{name: "Hero Master 125 Review"}}
}

// Seed implements the Seeder interface
func (s *HeroMaster125Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Master 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero master 125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding hero master 125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero Master 125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero Master 125 Review Bangladesh 2024 - Value Champion</h1>
<p class="summary-text">The Hero Master 125 is a practical commuter featuring 125cc engine, simple design, excellent reliability, and unbeatable price. Priced around ৳1,15,000, it's the perfect entry-level bike for budget-conscious first-time buyers seeking reliable Indian engineering.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Master 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc Engine:</strong> <span class="highlight-value">Reliable commuter powerplant</span></li>
<li class="highlight-item"><strong class="highlight-label">Simple Design:</strong> <span class="highlight-value">No-frills practical construction</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Price:</strong> <span class="highlight-value">Exceptional value option</span></li>
<li class="highlight-item"><strong class="highlight-label">Easy Maintenance:</strong> <span class="highlight-value">Simple repairs and service</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Master 125 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Unbeatable Price:</strong> <span class="pro-description">Exceptional value option</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">50-55 km/l excellent efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Proven 125cc platform</span></li>
<li class="pro-item"><strong class="pro-label">Easy Service:</strong> <span class="pro-description">Widely available parts and service</span></li>
<li class="pro-item"><strong class="pro-label">Spacious Seating:</strong> <span class="pro-description">Comfortable for rider and pillion</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Master 125 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Basic Design:</strong> <span class="con-description">Minimal style and features</span></li>
<li class="con-item"><strong class="con-label">No ABS/Disc:</strong> <span class="con-description">Drum brakes only</span></li>
<li class="con-item"><strong class="con-label">Lower Build Quality:</strong> <span class="con-description">Basic construction compared to Honda/Yamaha</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Master 125 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,15,000 - Absolute budget champion</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Ultra Low - ৳2-3 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳2-3 per km (50-55 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Master 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Price:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Exceptional value</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent 50-55 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good proven reliability</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Spacious seating</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Adequate commuter power</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Master 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,15,000, the Hero Master 125 is the ultimate budget option for first-time buyers and economy commuters. Excellent value with proven reliability makes it perfect for those prioritizing price over premium features.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget-conscious first-time buyers</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating hero master 125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero Master 125 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো মাস্টার 125 রিভিউ বাংলাদেশ 2024 - মূল্য চ্যাম্পিয়ন</h1>
<p class="summary-text">হিরো মাস্টার 125 একটি ব্যবহারিক কমিউটার যা 125cc ইঞ্জিন, সহজ ডিজাইন, চমৎকার নির্ভরযোগ্যতা এবং অপ্রতিদ্বন্দ্বী মূল্য বৈশিষ্ট্যযুক্ত। ৳1,15,000 টাকায় মূল্যায়িত, এটি বাজেট-সচেতন প্রথমবারের ক্রেতাদের জন্য নিখুঁত এন্ট্রি-লেভেল বাইক।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো মাস্টার 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc ইঞ্জিন:</strong> <span class="highlight-value">নির্ভরযোগ্য কমিউটার পাওয়ারপ্লান্ট</span></li>
<li class="highlight-item"><strong class="highlight-label">সহজ ডিজাইন:</strong> <span class="highlight-value">No-frills ব্যবহারিক নির্মাণ</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট মূল্য:</strong> <span class="highlight-value">ব্যতিক্রমী মূল্য বিকল্প</span></li>
<li class="highlight-item"><strong class="highlight-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="highlight-value">সহজ মেরামত এবং সেবা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো মাস্টার 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অপ্রতিদ্বন্দ্বী মূল্য:</strong> <span class="pro-description">ব্যতিক্রমী মূল্য বিকল্প</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">50-55 km/l চমৎকার দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">প্রমাণিত 125cc প্ল্যাটফর্ম</span></li>
<li class="pro-item"><strong class="pro-label">সহজ সেবা:</strong> <span class="pro-description">ব্যাপকভাবে উপলব্ধ যন্ত্রাংশ এবং সেবা</span></li>
<li class="pro-item"><strong class="pro-label">প্রশস্ত সিটিং:</strong> <span class="pro-description">রাইডার এবং পিলিয়নের জন্য আরামদায়ক</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো মাস্টার 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত মৌলিক ডিজাইন:</strong> <span class="con-description">ন্যূনতম শৈলী এবং বৈশিষ্ট্য</span></li>
<li class="con-item"><strong class="con-label">কোন ABS/Disc নেই:</strong> <span class="con-description">শুধুমাত্র ড্রাম ব্রেক</span></li>
<li class="con-item"><strong class="con-label">কম বিল্ড কোয়ালিটি:</strong> <span class="con-description">হোন্ডা/ইয়ামাহার তুলনায় মৌলিক নির্মাণ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো মাস্টার 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,15,000 টাকায়, হিরো মাস্টার 125 প্রথমবারের ক্রেতা এবং অর্থনীতি কমিউটারদের জন্য চূড়ান্ত বাজেট বিকল্প। ব্যতিক্রমী মূল্যের সাথে প্রমাণিত নির্ভরযোগ্যতা এটিকে প্রিমিয়াম বৈশিষ্ট্যের চেয়ে মূল্যকে অগ্রাধিকার দেয় তাদের জন্য নিখুঁত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট-সচেতন প্রথমবারের ক্রেতাদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero Master 125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero Master 125\n")

	return nil
}

// GetName returns the seeder name
func (s *HeroMaster125Review) GetName() string {
	return "HeroMaster125Review"
}
