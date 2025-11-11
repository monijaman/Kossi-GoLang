package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiSmashReview handles seeding Suzuki Smash product review and translation
type SuzukiSmashReview struct {
	BaseSeeder
}

// NewSuzukiSmashReviewSeeder creates a new SuzukiSmashReview
func NewSuzukiSmashReviewSeeder() *SuzukiSmashReview {
	return &SuzukiSmashReview{BaseSeeder: BaseSeeder{name: "Suzuki Smash Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiSmashReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Smash").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki smash product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding suzuki smash product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Smash already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Smash Review Bangladesh 2024 - Economy King 100cc</h1>
<p class="summary-text">The Suzuki Smash is a legendary 100cc economical commuter with simple design, tried-and-tested engine, excellent fuel efficiency, and affordability. Priced at ৳1,55,000, it's the budget commuter champion for cost-conscious daily riders seeking value and reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Smash Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc Budget Engine:</strong> <span class="highlight-value">Economy commuter proven</span></li>
<li class="highlight-item"><strong class="highlight-label">Simple Design:</strong> <span class="highlight-value">Easy to maintain</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">60-65 km/l exceptional</span></li>
<li class="highlight-item"><strong class="highlight-label">Affordable Price:</strong> <span class="highlight-value">Best value 100cc</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Smash Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Exceptional Mileage:</strong> <span class="pro-description">60-65 km/l amazing</span></li>
<li class="pro-item"><strong class="pro-label">Simple Design:</strong> <span class="pro-description">Easy maintenance</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Best budget option</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Reliability:</strong> <span class="pro-description">Proven quality</span></li>
<li class="pro-item"><strong class="pro-label">Low Running Cost:</strong> <span class="pro-description">₳2-3 per km</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Smash Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">No Features:</strong> <span class="con-description">Basic commuter bike</span></li>
<li class="con-item"><strong class="con-label">Performance:</strong> <span class="con-description">100cc budget performance</span></li>
<li class="con-item"><strong class="con-label">Comfort:</strong> <span class="con-description">Basic seating</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Smash Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,55,000 - Budget 100cc</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very low - ৳2-3 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">60-65 km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Smash Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- 60-65 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Affordability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Best price</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Suzuki proven</span></li>
<li class="rating-item"><strong class="rating-label">Maintenance:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Simple easy</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Exceptional budget</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Smash?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,55,000, the Suzuki Smash is the ultimate budget commuter champion for riders prioritizing exceptional fuel efficiency, low running costs, simple reliability, and maximum value for their investment.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget-conscious commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.7,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating suzuki smash review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Smash (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি স্ম্যাশ রিভিউ বাংলাদেশ 2024 - অর্থনীতি রাজা 100cc</h1>
<p class="summary-text">সুজুকি স্ম্যাশ একটি কিংবদন্তি 100cc অর্থনৈতিক কমিউটার যা সহজ ডিজাইন, পরীক্ষিত ইঞ্জিন, চমৎকার জ্বালানী দক্ষতা এবং সামর্থ্য বৈশিষ্ট্যযুক্ত। ৳1,55,000 টাকায় মূল্যায়িত, এটি খরচ-সচেতন দৈনন্দিন রাইডারদের জন্য মূল্য এবং নির্ভরযোগ্যতা খুঁজছেন বাজেট কমিউটার চ্যাম্পিয়ন।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি স্ম্যাশ মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc বাজেট ইঞ্জিন:</strong> <span class="highlight-value">অর্থনীতি কমিউটার প্রমাণিত</span></li>
<li class="highlight-item"><strong class="highlight-label">সহজ ডিজাইন:</strong> <span class="highlight-value">রক্ষণাবেক্ষণ সহজ</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">60-65 km/l ব্যতিক্রমী</span></li>
<li class="highlight-item"><strong class="highlight-label">সাশ্রয়ী মূল্য:</strong> <span class="highlight-value">সেরা মূল্য 100cc</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি স্ম্যাশ সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ব্যতিক্রমী মাইলেজ:</strong> <span class="pro-description">60-65 km/l অসাধারণ</span></li>
<li class="pro-item"><strong class="pro-label">সহজ ডিজাইন:</strong> <span class="pro-description">সহজ রক্ষণাবেক্ষণ</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">সেরা বাজেট বিকল্প</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">কম চলমান খরচ:</strong> <span class="pro-description">৳2-3 প্রতি km</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি স্ম্যাশ অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">কোন বৈশিষ্ট্য নেই:</strong> <span class="con-description">মৌলিক কমিউটার বাইক</span></li>
<li class="con-item"><strong class="con-label">কর্মক্ষমতা:</strong> <span class="con-description">100cc বাজেট কর্মক্ষমতা</span></li>
<li class="con-item"><strong class="con-label">আরাম:</strong> <span class="con-description">মৌলিক আসন</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি স্ম্যাশ কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,55,000 টাকায়, সুজুকি স্ম্যাশ চূড়ান্ত বাজেট কমিউটার চ্যাম্পিয়ন রাইডারদের জন্য যারা ব্যতিক্রমী জ্বালানী দক্ষতা, কম চলমান খরচ, সহজ নির্ভরযোগ্যতা এবং তাদের বিনিয়োগের জন্য সর্বাধিক মূল্য অগ্রাধিকার দেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট-সচেতন কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Smash already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Smash\n")

	return nil
}

// GetName returns the seeder name
func (s *SuzukiSmashReview) GetName() string {
	return "SuzukiSmashReview"
}
