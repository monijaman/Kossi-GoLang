package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroHFDeluxeIXReview handles seeding Hero HF Deluxe IX product review and translation
type HeroHFDeluxeIXReview struct {
	BaseSeeder
}

// NewHeroHFDeluxeIXReviewSeeder creates a new HeroHFDeluxeIXReview
func NewHeroHFDeluxeIXReviewSeeder() *HeroHFDeluxeIXReview {
	return &HeroHFDeluxeIXReview{BaseSeeder: BaseSeeder{name: "Hero HF Deluxe IX Review"}}
}

// Seed implements the Seeder interface
func (s *HeroHFDeluxeIXReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero MotoCorp HF Deluxe IX").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero hf deluxe ix product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding hero hf deluxe ix product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero HF Deluxe IX already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero HF Deluxe IX Review Bangladesh 2024 - Reliable Commuter Bike</h1>
<p class="summary-text">The Hero HF Deluxe IX is a dependable 100cc commuter bike with reliable engine, fuel efficiency, comfortable seating, and practical design. Priced around ৳1,60,000, it delivers trustworthy performance with excellent mileage for daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero HF Deluxe IX Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc Reliable Engine:</strong> <span class="highlight-value">Trusted performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficient:</strong> <span class="highlight-value">55-60 km/l excellent mileage</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seating:</strong> <span class="highlight-value">Ergonomic design</span></li>
<li class="highlight-item"><strong class="highlight-label">Practical Features:</strong> <span class="highlight-value">Daily commute ready</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero HF Deluxe IX Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Reliable Performance:</strong> <span class="pro-description">100cc trustworthy engine</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">55-60 km/l efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Budget-friendly option</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Good seating position</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Simple robust design</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero HF Deluxe IX Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Limited modern tech</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">100cc limited performance</span></li>
<li class="con-item"><strong class="con-label">No ABS:</strong> <span class="con-description">Standard brakes only</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero HF Deluxe IX Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,60,000 - Affordable commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very Low - ৳3 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">55-60 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero HF Deluxe IX Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Trusted 100cc</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 55-60 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good seating</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Budget-friendly</span></li>
<li class="rating-item"><strong class="rating-label">Maintenance:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Simple design</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero HF Deluxe IX?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,60,000, the Hero HF Deluxe IX is perfect for budget-conscious commuters seeking a reliable, fuel-efficient, and low-maintenance daily commuter bike with trustworthy performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget commuters</span></p>
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
		return fmt.Errorf("error creating hero hf deluxe ix review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero HF Deluxe IX (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো মোটোকর্প HF ডিলাক্স IX রিভিউ বাংলাদেশ 2024 - নির্ভরযোগ্য কমিউটার বাইক</h1>
<p class="summary-text">হিরো HF ডিলাক্স IX একটি নির্ভরযোগ্য 100cc কমিউটার বাইক যা নির্ভরযোগ্য ইঞ্জিন, জ্বালানি দক্ষতা, আরামদায়ক আসন এবং ব্যবহারিক ডিজাইন সহ আসে। ৳1,60,000 টাকায় মূল্যায়িত, এটি দৈনন্দিন যাতায়াতের জন্য চমৎকার মাইলেজ সহ বিশ্বাসযোগ্য কর্মক্ষমতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো HF ডিলাক্স IX মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="highlight-value">বিশ্বস্ত কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষ:</strong> <span class="highlight-value">55-60 km/l চমৎকার মাইলেজ</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক আসন:</strong> <span class="highlight-value">এরগোনমিক ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যবহারিক বৈশিষ্ট্য:</strong> <span class="highlight-value">দৈনন্দিন যাতায়াত প্রস্তুত</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো HF ডিলাক্স IX সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য কর্মক্ষমতা:</strong> <span class="pro-description">100cc বিশ্বাসযোগ্য ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">55-60 km/l দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">বাজেট-বান্ধব বিকল্প</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক যাত্রা:</strong> <span class="pro-description">ভাল আসন অবস্থান</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সহজ শক্তিশালী ডিজাইন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো HF ডিলাক্স IX অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">সীমিত আধুনিক প্রযুক্তি</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">100cc সীমিত কর্মক্ষমতা</span></li>
<li class="con-item"><strong class="con-label">কোন ABS নেই:</strong> <span class="con-description">শুধুমাত্র মানক ব্রেক</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো HF ডিলাক্স IX কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,60,000 টাকায়, হিরো HF ডিলাক্স IX বাজেট-সচেতন যাত্রীদের জন্য নিখুঁত যারা নির্ভরযোগ্য, জ্বালানি-দক্ষ এবং কম রক্ষণাবেক্ষণের দৈনন্দিন কমিউটার বাইক খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero HF Deluxe IX already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero HF Deluxe IX\n")

	return nil
}

// GetName returns the seeder name
func (s *HeroHFDeluxeIXReview) GetName() string {
	return "HeroHFDeluxeIXReview"
}
