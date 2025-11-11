package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCB110Review handles seeding Honda CB110 product review and translation
type HondaCB110Review struct {
	BaseSeeder
}

// NewHondaCB110ReviewSeeder creates a new HondaCB110Review
func NewHondaCB110ReviewSeeder() *HondaCB110Review {
	return &HondaCB110Review{BaseSeeder: BaseSeeder{name: "Honda CB110 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCB110Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB110").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda cb110 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding honda cb110 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CB110 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CB110 Review Bangladesh 2024 - Classic Commuter</h1>
<p class="summary-text">The Honda CB110 is a timeless commuter motorcycle featuring 110cc air-cooled engine, classic design, reliable Japanese engineering, and excellent fuel economy. Priced around ৳1,20,000, it's the perfect lightweight commuter for daily urban riding with proven durability and minimal maintenance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB110 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc Air-Cooled:</strong> <span class="highlight-value">Proven reliable engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Classic Design:</strong> <span class="highlight-value">Timeless simple styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Japanese Engineering:</strong> <span class="highlight-value">Honda quality and reliability</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">60-65 km/l efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB110 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Outstanding Reliability:</strong> <span class="pro-description">Honda quality and longevity</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">60-65 km/l outstanding efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Simple and affordable service</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight:</strong> <span class="pro-description">Easy to handle in traffic</span></li>
<li class="pro-item"><strong class="pro-label">Resale Value:</strong> <span class="pro-description">Honda brand maintains value</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB110 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">Higher than Bajaj alternatives</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">Modest acceleration</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Minimal modern amenities</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB110 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,20,000 - Premium budget option</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳2-3 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳2-3 per km (60-65 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB110 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Exceptional durability</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Outstanding 60-65 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Premium construction</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Adequate for commuting</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Premium but worth it</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB110?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,20,000, the Honda CB110 is an excellent choice for buyers prioritizing reliability, fuel economy, and resale value. The premium Japanese quality justifies the higher price for serious commuters.</p>
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
		return fmt.Errorf("error creating honda cb110 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CB110 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা CB110 রিভিউ বাংলাদেশ 2024 - ক্লাসিক কমিউটার</h1>
<p class="summary-text">হোন্ডা CB110 একটি কালজয়ী কমিউটার মোটরসাইকেল যা 110cc এয়ার-কুল্ড ইঞ্জিন, ক্লাসিক ডিজাইন, নির্ভরযোগ্য জাপানি প্রকৌশল এবং চমৎকার জ্বালানি অর্থনীতি বৈশিষ্ট্যযুক্ত। ৳1,20,000 টাকায় মূল্যায়িত, এটি দৈনন্দিন শহুরে যাতায়াতের জন্য নিখুঁত হালকা কমিউটার।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা CB110 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc এয়ার-কুল্ড:</strong> <span class="highlight-value">প্রমাণিত নির্ভরযোগ্য ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">ক্লাসিক ডিজাইন:</strong> <span class="highlight-value">কালজয়ী সহজ স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">জাপানি প্রকৌশল:</strong> <span class="highlight-value">হোন্ডা গুণমান এবং নির্ভরযোগ্যতা</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">60-65 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা CB110 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অসাধারণ নির্ভরযোগ্যতা:</strong> <span class="pro-description">হোন্ডা গুণমান এবং দীর্ঘায়ু</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">60-65 km/l অসাধারণ দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সহজ এবং সাশ্রয়ী সেবা</span></li>
<li class="pro-item"><strong class="pro-label">হালকা ওজন:</strong> <span class="pro-description">ট্রাফিকে পরিচালনা করা সহজ</span></li>
<li class="pro-item"><strong class="pro-label">পুনর্বিক্রয় মূল্য:</strong> <span class="pro-description">হোন্ডা ব্র্যান্ড মূল্য বজায় রাখে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা CB110 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">বাজাজ বিকল্পের চেয়ে বেশি</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">মামুলি ত্বরণ</span></li>
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতম আধুনিক সুবিধা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা CB110 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,20,000 টাকায়, হোন্ডা CB110 যারা নির্ভরযোগ্যতা, জ্বালানি অর্থনীতি এবং পুনর্বিক্রয় মূল্যকে অগ্রাধিকার দেন তাদের জন্য চমৎকার পছন্দ। প্রিমিয়াম জাপানি গুণমান উচ্চতর মূল্যকে ন্যায্যতা দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- নির্ভরযোগ্য দৈনন্দিন যাতায়াতের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CB110 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CB110\n")

	return nil
}

// GetName returns the seeder name
func (s *HondaCB110Review) GetName() string {
	return "HondaCB110Review"
}
