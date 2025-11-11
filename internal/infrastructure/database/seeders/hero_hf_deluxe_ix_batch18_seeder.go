package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroHFDeluxeIXReviewBatch18 handles seeding Hero HF Deluxe IX product review and translation
type HeroHFDeluxeIXReviewBatch18 struct {
	BaseSeeder
}

// NewHeroHFDeluxeIXReviewBatch18Seeder creates a new HeroHFDeluxeIXReviewBatch18
func NewHeroHFDeluxeIXReviewBatch18Seeder() *HeroHFDeluxeIXReviewBatch18 {
	return &HeroHFDeluxeIXReviewBatch18{BaseSeeder: BaseSeeder{name: "Hero HF Deluxe IX Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *HeroHFDeluxeIXReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero HF Deluxe IX").First(&product).Error; err != nil {
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
<h1 class="review-main-title">Hero HF Deluxe IX Review Bangladesh 2024 - Budget Commuter Excellence Leader</h1>
<p class="summary-text">The Hero HF Deluxe IX is a 100cc air-cooled commuter bringing budget-friendly pricing with practical features. Priced around ৳1,22,000, it delivers reliable performance, exceptional fuel efficiency, comfortable ergonomics, strong resale value, and Hero's dependability for cost-conscious daily commuters seeking value-for-money excellence.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero HF Deluxe IX Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc Air-Cooled:</strong> <span class="highlight-value">Budget commuter reliable</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficiency:</strong> <span class="highlight-value">75-85 km/l best class</span></li>
<li class="highlight-item"><strong class="highlight-label">Affordability:</strong> <span class="highlight-value">৳1,22,000 accessible</span></li>
<li class="highlight-item"><strong class="highlight-label">Resale Value:</strong> <span class="highlight-value">Strong depreciation</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero HF Deluxe IX Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Exceptional Efficiency:</strong> <span class="pro-description">75-85 km/l best</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">৳1,22,000 budget</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Hero proven</span></li>
<li class="pro-item"><strong class="pro-label">Strong Resale:</strong> <span class="pro-description">Budget bike value</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Service affordable</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero HF Deluxe IX Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Minimalist equipment</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">7.9 bhp modest</span></li>
<li class="con-item"><strong class="con-label">Comfort Limited:</strong> <span class="con-description">Budget seating</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero HF Deluxe IX Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,22,000 - Ultra-budget</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Ultra-low - ৳2-3 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">75-85 km/l exceptional</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">100cc air-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">7.9 bhp economy torque</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">108kg lightweight</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero HF Deluxe IX Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- 75-85 km/l exceptional</span></li>
<li class="rating-item"><strong class="rating-label">Affordability:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- ৳1,22,000 best</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Hero proven</span></li>
<li class="rating-item"><strong class="rating-label">Resale Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Strong demand</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Best budget bike</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero HF Deluxe IX?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,22,000, the HF Deluxe IX is the ultimate budget commuter delivering exceptional 75-85 km/l fuel efficiency, Hero reliability, strong resale value, and affordable ownership for cost-conscious daily commuters seeking maximum value in the ultra-budget segment.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget commuters</span></p>
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
		return fmt.Errorf("error creating hero hf deluxe ix review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero HF Deluxe IX (Rating: 4.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো এইচএফ ডিলাক্স IX রিভিউ বাংলাদেশ 2024 - বাজেট কমিউটার উৎকর্ষতা নেতা</h1>
<p class="summary-text">হিরো এইচএফ ডিলাক্স IX একটি 100cc এয়ার-কুল্ড কমিউটার যা বাজেট-বান্ধব মূল্য নিয়ে আসে। ৳1,22,000 টাকায় মূল্যায়িত, এটি নির্ভরযোগ্য পারফরম্যান্স, ব্যতিক্রমী জ্বালানি দক্ষতা এবং হিরো নির্ভরযোগ্যতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো এইচএফ ডিলাক্স IX মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc এয়ার-কুল্ড:</strong> <span class="highlight-value">বাজেট কমিউটার নির্ভরযোগ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষতা:</strong> <span class="highlight-value">75-85 km/l সেরা শ্রেণী</span></li>
<li class="highlight-item"><strong class="highlight-label">সামর্থ্য:</strong> <span class="highlight-value">৳1,22,000 অ্যাক্সেসযোগ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">পুনঃবিক্রয় মূল্য:</strong> <span class="highlight-value">শক্তিশালী হ্রাস</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো এইচএফ ডিলাক্স IX সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ব্যতিক্রমী দক্ষতা:</strong> <span class="pro-description">75-85 km/l সেরা</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">৳1,22,000 বাজেট</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">হিরো প্রমাণিত</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী পুনঃবিক্রয়:</strong> <span class="pro-description">বাজেট বাইক মূল্য</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সেবা সাশ্রয়ী</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো এইচএফ ডিলাক্স IX অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতমবাদী সরঞ্জাম</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">7.9 bhp বিনম্র</span></li>
<li class="con-item"><strong class="con-label">আরাম সীমিত:</strong> <span class="con-description">বাজেট সিটিং</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো এইচএফ ডিলাক্স IX কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,22,000 টাকায়, এইচএফ ডিলাক্স IX চূড়ান্ত বাজেট কমিউটার যা ব্যতিক্রমী 75-85 km/l জ্বালানি দক্ষতা এবং হিরো নির্ভরযোগ্যতা প্রদান করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট কমিউটারদের জন্য</span></p>
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
func (s *HeroHFDeluxeIXReviewBatch18) GetName() string {
	return "HeroHFDeluxeIXReviewBatch18"
}
