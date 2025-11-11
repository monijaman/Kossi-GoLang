package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCD110DreamReview handles seeding Honda CD110 Dream product review and translation
type HondaCD110DreamReview struct {
	BaseSeeder
}

// NewHondaCD110DreamReviewSeeder creates a new HondaCD110DreamReview
func NewHondaCD110DreamReviewSeeder() *HondaCD110DreamReview {
	return &HondaCD110DreamReview{BaseSeeder: BaseSeeder{name: "Honda CD110 Dream Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCD110DreamReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CD110 Dream").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda cd110 dream product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding honda cd110 dream product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CD110 Dream already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CD110 Dream Review Bangladesh 2024 - Reliable Budget Commuter</h1>
<p class="summary-text">The Honda CD110 Dream is a reliable 110cc commuter bike with proven Honda engineering, simple design, excellent fuel efficiency, and legendary dependability. Priced at ৳1,58,000, it's the trusted workhorse for daily commuting with Honda's reputation backing it.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CD110 Dream Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Proven Honda Quality:</strong> <span class="highlight-value">Legendary reliability</span></li>
<li class="highlight-item"><strong class="highlight-label">110cc Engine:</strong> <span class="highlight-value">Smooth economical</span></li>
<li class="highlight-item"><strong class="highlight-label">Simple Design:</strong> <span class="highlight-value">Easy maintenance</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">55-60 km/l great</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CD110 Dream Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven engineering</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">110cc economical</span></li>
<li class="pro-item"><strong class="pro-label">Easy Maintenance:</strong> <span class="pro-description">Simple design</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">55-60 km/l</span></li>
<li class="pro-item"><strong class="pro-label">Good Resale:</strong> <span class="pro-description">Honda trusted brand</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CD110 Dream Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Performance:</strong> <span class="con-description">Basic commuter performance</span></li>
<li class="con-item"><strong class="con-label">Features:</strong> <span class="con-description">No modern features</span></li>
<li class="con-item"><strong class="con-label">Styling:</strong> <span class="con-description">Plain conservative design</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CD110 Dream Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,58,000 - Budget commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very low - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">55-60 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CD110 Dream Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Proven Honda</span></li>
<li class="rating-item"><strong class="rating-label">Engine Smoothness:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Quiet smooth</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 55-60 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Maintenance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Easy simple</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Trusted workhorse</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CD110 Dream?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,58,000, the Honda CD110 Dream is the perfect choice for daily commuters seeking proven reliability, excellent fuel efficiency, easy maintenance, and the peace of mind that comes with Honda's legendary quality and dependability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For reliable daily commuting</span></p>
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
		return fmt.Errorf("error creating honda cd110 dream review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CD110 Dream (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হন্ডা CD110 Dream রিভিউ বাংলাদেশ 2024 - নির্ভরযোগ্য বাজেট কমিউটার</h1>
<p class="summary-text">হন্ডা CD110 Dream একটি নির্ভরযোগ্য 110cc কমিউটার বাইক যা প্রমাণিত হন্ডা প্রকৌশল, সহজ ডিজাইন, চমৎকার জ্বালানী দক্ষতা এবং কিংবদন্তি নির্ভরযোগ্যতা সহ আসে। ৳1,58,000 টাকায় মূল্যায়িত, এটি দৈনন্দিন যাতায়াতের জন্য হন্ডার সুনাম দ্বারা সমর্থিত বিশ্বস্ত কর্মঘোড়া।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা CD110 Dream মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">প্রমাণিত হন্ডা গুণমান:</strong> <span class="highlight-value">কিংবদন্তি নির্ভরযোগ্যতা</span></li>
<li class="highlight-item"><strong class="highlight-label">110cc ইঞ্জিন:</strong> <span class="highlight-value">মসৃণ অর্থনৈতিক</span></li>
<li class="highlight-item"><strong class="highlight-label">সহজ ডিজাইন:</strong> <span class="highlight-value">সহজ রক্ষণাবেক্ষণ</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">55-60 km/l দুর্দান্ত</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা CD110 Dream সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">হন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত প্রকৌশল</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">110cc অর্থনৈতিক</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সহজ ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">55-60 km/l</span></li>
<li class="pro-item"><strong class="pro-label">ভাল পুনর্বিক্রয়:</strong> <span class="pro-description">হন্ডা বিশ্বস্ত ব্র্যান্ড</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা CD110 Dream অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">কর্মক্ষমতা:</strong> <span class="con-description">মৌলিক কমিউটার কর্মক্ষমতা</span></li>
<li class="con-item"><strong class="con-label">বৈশিষ্ট্য:</strong> <span class="con-description">কোন আধুনিক বৈশিষ্ট্য নেই</span></li>
<li class="con-item"><strong class="con-label">স্টাইলিং:</strong> <span class="con-description">সাধারণ রক্ষণশীল ডিজাইন</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হন্ডা CD110 Dream কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,58,000 টাকায়, হন্ডা CD110 Dream দৈনন্দিন কমিউটারদের জন্য নিখুঁত যারা প্রমাণিত নির্ভরযোগ্যতা, চমৎকার জ্বালানী দক্ষতা, সহজ রক্ষণাবেক্ষণ এবং হন্ডার কিংবদন্তি গুণমান এবং নির্ভরযোগ্যতার সাথে আসা শান্তি খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- নির্ভরযোগ্য দৈনন্দিন যাতায়াতের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CD110 Dream already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CD110 Dream\n")

	return nil
}

// GetName returns the seeder name
func (s *HondaCD110DreamReview) GetName() string {
	return "HondaCD110DreamReview"
}
