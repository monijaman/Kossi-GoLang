package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroHeroesReviewBatch19 handles seeding Hero Heroes product review and translation
type HeroHeroesReviewBatch19 struct {
	BaseSeeder
}

// NewHeroHeroesReviewBatch19Seeder creates a new HeroHeroesReviewBatch19
func NewHeroHeroesReviewBatch19Seeder() *HeroHeroesReviewBatch19 {
	return &HeroHeroesReviewBatch19{BaseSeeder: BaseSeeder{name: "Hero Heroes Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *HeroHeroesReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Heroes").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero heroes product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding hero heroes product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero Heroes already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero Heroes Review Bangladesh 2024 - Family Workhorse Trusted Transport</h1>
<p class="summary-text">The Hero Heroes is a 110cc air-cooled family-oriented commuter designed for everyday practical transport. Priced around ৳1,35,000, it delivers spacious seating, reliable mechanics, exceptional fuel efficiency, strong resale, and Hero's dependability for families seeking affordable, straightforward daily commuting with trusted reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Heroes Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc Air-Cooled:</strong> <span class="highlight-value">Family commuter</span></li>
<li class="highlight-item"><strong class="highlight-label">Spacious Seating:</strong> <span class="highlight-value">Comfortable positioning</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Affordable:</strong> <span class="highlight-value">৳1,35,000 accessible</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Economy:</strong> <span class="highlight-value">65-70 km/l exceptional</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Heroes Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Spacious Comfort:</strong> <span class="pro-description">Family friendly</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Economy:</strong> <span class="pro-description">65-70 km/l exceptional</span></li>
<li class="pro-item"><strong class="pro-label">Budget Price:</strong> <span class="pro-description">৳1,35,000 accessible</span></li>
<li class="pro-item"><strong class="pro-label">Strong Resale:</strong> <span class="pro-description">Market demand high</span></li>
<li class="pro-item"><strong class="pro-label">Reliability:</strong> <span class="pro-description">Hero proven</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Heroes Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Minimalist equipment</span></li>
<li class="con-item"><strong class="con-label">Modest Power:</strong> <span class="con-description">8.36 bhp basic</span></li>
<li class="con-item"><strong class="con-label">Build Simple:</strong> <span class="con-description">Budget segment typical</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Heroes Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,35,000 - Ultra-budget family</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Ultra-low - ৳2-3 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">65-70 km/l exceptional</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">110cc air-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">8.36 bhp basic family</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">111kg lightweight</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Heroes Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Family Practicality:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Spacious friendly</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- 65-70 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Affordability:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- ৳1,35,000 best</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Hero dependable</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Family champion</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Heroes?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,35,000, the Heroes is perfect for families seeking ultimate budget commuting, exceptional 65-70 km/l efficiency, spacious comfort, strong resale value, and Hero's trusted reliability for straightforward, affordable family transportation.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For family commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.8,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating hero heroes review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero Heroes (Rating: 4.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো হিরোস রিভিউ বাংলাদেশ 2024 - ফ্যামিলি ওয়ার্কহর্স বিশ্বস্ত পরিবহন</h1>
<p class="summary-text">হিরো হিরোস একটি 110cc এয়ার-কুল্ড পারিবারিক-কেন্দ্রিক কমিউটার যা দৈনন্দিন ব্যবহারিক পরিবহনের জন্য ডিজাইন করা। ৳1,35,000 টাকায় মূল্যায়িত, এটি প্রশস্ত সিটিং এবং নির্ভরযোগ্য যান্ত্রিকতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো হিরোস মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc এয়ার-কুল্ড:</strong> <span class="highlight-value">পারিবারিক কমিউটার</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রশস্ত সিটিং:</strong> <span class="highlight-value">আরামদায়ক পজিশনিং</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট সাশ্রয়ী:</strong> <span class="highlight-value">৳1,35,000 অ্যাক্সেসযোগ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি অর্থনীতি:</strong> <span class="highlight-value">65-70 km/l ব্যতিক্রমী</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো হিরোস সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রশস্ত আরাম:</strong> <span class="pro-description">পরিবার বান্ধব</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি অর্থনীতি:</strong> <span class="pro-description">65-70 km/l ব্যতিক্রমী</span></li>
<li class="pro-item"><strong class="pro-label">বাজেট মূল্য:</strong> <span class="pro-description">৳1,35,000 অ্যাক্সেসযোগ্য</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী পুনঃবিক্রয়:</strong> <span class="pro-description">বাজার চাহিদা উচ্চ</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্যতা:</strong> <span class="pro-description">হিরো প্রমাণিত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো হিরোস অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতমবাদী সরঞ্জাম</span></li>
<li class="con-item"><strong class="con-label">বিনম্র শক্তি:</strong> <span class="con-description">8.36 bhp মৌলিক</span></li>
<li class="con-item"><strong class="con-label">নির্মাণ সরল:</strong> <span class="con-description">বাজেট সেগমেন্ট সাধারণ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো হিরোস কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,35,000 টাকায়, হিরোস পরিবারগুলির জন্য নিখুঁত যারা চূড়ান্ত বাজেট কমিউটিং, ব্যতিক্রমী 65-70 km/l দক্ষতা, প্রশস্ত আরাম এবং হিরোর বিশ্বস্ত নির্ভরযোগ্যতা চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারিবারিক কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero Heroes already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero Heroes\n")

	return nil
}

// GetName returns the seeder name
func (s *HeroHeroesReviewBatch19) GetName() string {
	return "HeroHeroesReviewBatch19"
}
