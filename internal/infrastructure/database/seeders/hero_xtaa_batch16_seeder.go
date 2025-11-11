package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroMotocorpExtaReview handles seeding Hero MotoCorp Xtaa product review and translation
type HeroMotocorpExtaReview struct {
	BaseSeeder
}

// NewHeroMotocorpExtaReviewSeeder creates a new HeroMotocorpExtaReview
func NewHeroMotocorpExtaReviewSeeder() *HeroMotocorpExtaReview {
	return &HeroMotocorpExtaReview{BaseSeeder: BaseSeeder{name: "Hero Xtaa Review"}}
}

// Seed implements the Seeder interface
func (s *HeroMotocorpExtaReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Xtaa").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero xtaa product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding hero xtaa product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero Xtaa already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero Xtaa Review Bangladesh 2024 - Commuter Performance Blend</h1>
<p class="summary-text">The Hero Xtaa is a practical commuter bike combining ergonomic design, efficient engine, modern styling, and daily reliability. Priced around ৳1,85,000, it offers the perfect balance of comfort, performance, and affordability for regular urban commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Xtaa Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc Efficient Engine:</strong> <span class="highlight-value">Commuter focused</span></li>
<li class="highlight-item"><strong class="highlight-label">Ergonomic Design:</strong> <span class="highlight-value">Comfort oriented</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Styling:</strong> <span class="highlight-value">Contemporary looks</span></li>
<li class="highlight-item"><strong class="highlight-label">Great Efficiency:</strong> <span class="highlight-value">55-60 km/l excellent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Xtaa Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Ergonomic Comfort:</strong> <span class="pro-description">Upright seating</span></li>
<li class="pro-item"><strong class="pro-label">Efficient Engine:</strong> <span class="pro-description">Low fuel consumption</span></li>
<li class="pro-item"><strong class="pro-label">Modern Design:</strong> <span class="pro-description">Updated styling</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Bike:</strong> <span class="pro-description">Hero proven</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Good value</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Xtaa Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Performance:</strong> <span class="con-description">100cc commuter level</span></li>
<li class="con-item"><strong class="con-label">Features:</strong> <span class="con-description">Minimal technology</span></li>
<li class="con-item"><strong class="con-label">Handling:</strong> <span class="con-description">Comfortable, not sporty</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Xtaa Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,85,000 - Budget commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very low - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">55-60 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Xtaa Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Ergonomic upright</span></li>
<li class="rating-item"><strong class="rating-label">Efficiency:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 55-60 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Modern practical</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Hero trusted</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good balance</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Xtaa?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,85,000, the Hero Xtaa is ideal for daily commuters seeking an ergonomic, efficient, modern-looking bike that combines comfort, low running costs, and reliable Hero performance for everyday urban riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For comfort-focused commuters</span></p>
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
		return fmt.Errorf("error creating hero xtaa review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero Xtaa (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো Xtaa রিভিউ বাংলাদেশ 2024 - কমিউটার পারফরম্যান্স মিশ্রণ</h1>
<p class="summary-text">হিরো Xtaa একটি ব্যবহারিক কমিউটার বাইক যা এরগনমিক ডিজাইন, দক্ষ ইঞ্জিন, আধুনিক স্টাইলিং এবং দৈনিক নির্ভরযোগ্যতা একত্রিত করে। ৳1,85,000 টাকায় মূল্যায়িত, এটি নিয়মিত শহুরে কমিউটারদের জন্য আরাম, কর্মক্ষমতা এবং সামর্থ্যের নিখুঁত ভারসাম্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো Xtaa মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc দক্ষ ইঞ্জিন:</strong> <span class="highlight-value">কমিউটার ফোকাসড</span></li>
<li class="highlight-item"><strong class="highlight-label">এরগনমিক ডিজাইন:</strong> <span class="highlight-value">আরাম-ভিত্তিক</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক স্টাইলিং:</strong> <span class="highlight-value">সমসাময়িক চেহারা</span></li>
<li class="highlight-item"><strong class="highlight-label">দুর্দান্ত দক্ষতা:</strong> <span class="highlight-value">55-60 km/l চমৎকার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো Xtaa সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">এরগনমিক আরাম:</strong> <span class="pro-description">সোজা আসন</span></li>
<li class="pro-item"><strong class="pro-label">দক্ষ ইঞ্জিন:</strong> <span class="pro-description">কম জ্বালানী খরচ</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক ডিজাইন:</strong> <span class="pro-description">আপডেট স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য বাইক:</strong> <span class="pro-description">হিরো প্রমাণিত</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">ভাল মূল্য</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো Xtaa অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক কর্মক্ষমতা:</strong> <span class="con-description">100cc কমিউটার স্তর</span></li>
<li class="con-item"><strong class="con-label">বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতম প্রযুক্তি</span></li>
<li class="con-item"><strong class="con-label">হ্যান্ডলিং:</strong> <span class="con-description">আরামদায়ক, খেলাধুলা নয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো Xtaa কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,85,000 টাকায়, হিরো Xtaa দৈনন্দিন কমিউটারদের জন্য আদর্শ যারা একটি এরগনমিক, দক্ষ, আধুনিক-দেখতে বাইক খুঁজছেন যা আরাম, কম চালানো খরচ এবং দৈনিক শহুরে রাইডিংয়ের জন্য নির্ভরযোগ্য হিরো কর্মক্ষমতা একত্রিত করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আরাম-ফোকাসড কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero Xtaa already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero Xtaa\n")

	return nil
}

// GetName returns the seeder name
func (s *HeroMotocorpExtaReview) GetName() string {
	return "HeroMotocorpExtaReview"
}
