package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaYBR125FiReview handles seeding Yamaha YBR125 Fi product review and translation
type YamahaYBR125FiReview struct {
	BaseSeeder
}

// NewYamahaYBR125FiReviewSeeder creates a new YamahaYBR125FiReview
func NewYamahaYBR125FiReviewSeeder() *YamahaYBR125FiReview {
	return &YamahaYBR125FiReview{BaseSeeder: BaseSeeder{name: "Yamaha YBR125 Fi Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaYBR125FiReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha YBR125 Fi").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha ybr125 fi product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding yamaha ybr125 fi product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha YBR125 Fi already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha YBR125 Fi Review Bangladesh 2024 - Fuel Injection Innovation</h1>
<p class="summary-text">The Yamaha YBR125 Fi is an advanced 125cc commuter featuring fuel injection technology, auto choke, responsive performance, and excellent fuel economy. Priced around ৳1,65,000, it offers modern engineering with reliable Yamaha quality for daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha YBR125 Fi Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection:</strong> <span class="highlight-value">Modern technology for efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Auto Choke:</strong> <span class="highlight-value">Easy cold start capability</span></li>
<li class="highlight-item"><strong class="highlight-label">125cc Engine:</strong> <span class="highlight-value">Reliable Yamaha powerplant</span></li>
<li class="highlight-item"><strong class="highlight-label">Responsive Performance:</strong> <span class="highlight-value">Quick acceleration</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha YBR125 Fi Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Fuel Injection Tech:</strong> <span class="pro-description">Modern technology for better efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Easy Start:</strong> <span class="pro-description">Auto choke eliminates manual adjustment</span></li>
<li class="pro-item"><strong class="pro-label">Responsive Power:</strong> <span class="pro-description">Quick acceleration response</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">45-50 km/l excellent efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Yamaha Reliability:</strong> <span class="pro-description">Trusted brand quality</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha YBR125 Fi Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Pricing:</strong> <span class="con-description">Higher than carbureted alternatives</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">FI service more expensive than carb</span></li>
<li class="con-item"><strong class="con-label">Basic Design:</strong> <span class="con-description">Minimal styling features</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha YBR125 Fi Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,65,000 - Premium efficient option</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳3-4 per km (45-50 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha YBR125 Fi Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Modern fuel injection</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Yamaha quality</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent 45-50 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good 125cc power</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good tech value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha YBR125 Fi?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,65,000, the Yamaha YBR125 Fi is excellent for buyers wanting modern fuel injection technology with proven Yamaha reliability and good fuel economy for daily commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For modern tech-oriented commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.4,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating yamaha ybr125 fi review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha YBR125 Fi (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা YBR125 Fi রিভিউ বাংলাদেশ 2024 - জ্বালানি ইঞ্জেকশন উদ্ভাবন</h1>
<p class="summary-text">ইয়ামাহা YBR125 Fi একটি উন্নত 125cc কমিউটার যা জ্বালানি ইঞ্জেকশন প্রযুক্তি, অটো চোক, প্রতিক্রিয়াশীল পারফরমেন্স এবং চমৎকার জ্বালানি অর্থনীতি বৈশিষ্ট্যযুক্ত। ৳1,65,000 টাকায় মূল্যায়িত, এটি দৈনন্দিন যাতায়াতের জন্য নির্ভরযোগ্য ইয়ামাহা গুণমানের সাথে আধুনিক প্রকৌশল প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা YBR125 Fi মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">জ্বালানি ইঞ্জেকশন:</strong> <span class="highlight-value">দক্ষতার জন্য আধুনিক প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">অটো চোক:</strong> <span class="highlight-value">সহজ শীত স্টার্ট সামর্থ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">125cc ইঞ্জিন:</strong> <span class="highlight-value">নির্ভরযোগ্য ইয়ামাহা পাওয়ারপ্লান্ট</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রতিক্রিয়াশীল পারফরমেন্স:</strong> <span class="highlight-value">দ্রুত ত্বরণ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা YBR125 Fi সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">জ্বালানি ইঞ্জেকশন প্রযুক্তি:</strong> <span class="pro-description">ভাল দক্ষতার জন্য আধুনিক প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">সহজ স্টার্ট:</strong> <span class="pro-description">অটো চোক ম্যানুয়াল সামঞ্জস্য দূর করে</span></li>
<li class="pro-item"><strong class="pro-label">প্রতিক্রিয়াশীল শক্তি:</strong> <span class="pro-description">দ্রুত ত্বরণ প্রতিক্রিয়া</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">45-50 km/l চমৎকার দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">ইয়ামাহা নির্ভরযোগ্যতা:</strong> <span class="pro-description">বিশ্বস্ত ব্র্যান্ড গুণমান</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা YBR125 Fi অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">কার্বুরেটেড বিকল্পের চেয়ে বেশি</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">FI সেবা কার্ব চেয়ে বেশি ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">মৌলিক ডিজাইন:</strong> <span class="con-description">ন্যূনতম স্টাইলিং বৈশিষ্ট্য</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা YBR125 Fi কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,65,000 টাকায়, ইয়ামাহা YBR125 Fi যারা আধুনিক জ্বালানি ইঞ্জেকশন প্রযুক্তি চান প্রমাণিত ইয়ামাহা নির্ভরযোগ্যতা এবং ভাল জ্বালানি অর্থনীতি সহ তাদের জন্য চমৎকার।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আধুনিক প্রযুক্তি-উন্মুখ কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha YBR125 Fi already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha YBR125 Fi\n")

	return nil
}

// GetName returns the seeder name
func (s *YamahaYBR125FiReview) GetName() string {
	return "YamahaYBR125FiReview"
}
