package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCBTriggerReview handles seeding Honda CB Trigger product review and translation
type HondaCBTriggerReview struct {
	BaseSeeder
}

// NewHondaCBTriggerReviewSeeder creates a new HondaCBTriggerReview
func NewHondaCBTriggerReviewSeeder() *HondaCBTriggerReview {
	return &HondaCBTriggerReview{BaseSeeder: BaseSeeder{name: "Honda CB Trigger Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCBTriggerReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB Trigger").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda cb trigger product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding honda cb trigger product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CB Trigger already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CB Trigger Review Bangladesh 2024 - Reliable Street Bike</h1>
<p class="summary-text">The Honda CB Trigger is a versatile 150cc street bike with reliable engine, smooth transmission, comfortable seating, and practical features. Priced around ৳2,50,000, it offers dependable performance for daily urban commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB Trigger Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc Reliable Engine:</strong> <span class="highlight-value">Proven performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Smooth Transmission:</strong> <span class="highlight-value">Effortless operation</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seating:</strong> <span class="highlight-value">Well-padded design</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">45-50 km/l efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB Trigger Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">150cc proven performance</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Riding:</strong> <span class="pro-description">Effortless transmission</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seating:</strong> <span class="pro-description">Well-designed seat</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Trusted brand quality</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">45-50 km/l efficiency</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB Trigger Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Design:</strong> <span class="con-description">Simple styling</span></li>
<li class="con-item"><strong class="con-label">Limited Features:</strong> <span class="con-description">Minimal tech</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium maintenance</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB Trigger Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,50,000 - Practical street bike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳4-5 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">45-50 km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB Trigger Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Proven 150cc</span></li>
<li class="rating-item"><strong class="rating-label">Smoothness:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Effortless operation</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good seating</span></li>
<li class="rating-item"><strong class="rating-label">Brand Quality:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Honda trusted</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good practical features</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB Trigger?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,50,000, the Honda CB Trigger is ideal for practical commuters seeking reliable performance, smooth operation, comfortable seating, and proven Honda quality for daily urban riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For practical commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating honda cb trigger review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CB Trigger (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা CB Trigger রিভিউ বাংলাদেশ 2024 - নির্ভরযোগ্য স্ট্রিট বাইক</h1>
<p class="summary-text">হোন্ডা CB Trigger একটি বহুমুখী 150cc স্ট্রিট বাইক যা নির্ভরযোগ্য ইঞ্জিন, মসৃণ ট্রান্সমিশন, আরামদায়ক আসন এবং ব্যবহারিক বৈশিষ্ট্য সহ আসে। ৳2,50,000 টাকায় মূল্যায়িত, এটি দৈনন্দিন শহুরে যাতায়াতের জন্য নির্ভরযোগ্য কর্মক্ষমতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা CB Trigger মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="highlight-value">প্রমাণিত কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">মসৃণ ট্রান্সমিশন:</strong> <span class="highlight-value">অনায়াসী পরিচালনা</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক আসন:</strong> <span class="highlight-value">ভালো-প্যাডেড ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">45-50 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা CB Trigger সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">150cc প্রমাণিত কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ যাত্রা:</strong> <span class="pro-description">অনায়াসী ট্রান্সমিশন</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক আসন:</strong> <span class="pro-description">ভাল-ডিজাইন করা সিট</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">বিশ্বস্ত ব্র্যান্ড গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">45-50 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা CB Trigger অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক ডিজাইন:</strong> <span class="con-description">সহজ স্টাইলিং</span></li>
<li class="con-item"><strong class="con-label">সীমিত বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতম প্রযুক্তি</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম রক্ষণাবেক্ষণ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা CB Trigger কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,50,000 টাকায়, হোন্ডা CB Trigger ব্যবহারিক যাত্রীদের জন্য আদর্শ যারা নির্ভরযোগ্য কর্মক্ষমতা, মসৃণ পরিচালনা, আরামদায়ক আসন এবং প্রমাণিত হোন্ডা গুণমান খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্যবহারিক যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CB Trigger already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CB Trigger\n")

	return nil
}

// GetName returns the seeder name
func (s *HondaCBTriggerReview) GetName() string {
	return "HondaCBTriggerReview"
}
