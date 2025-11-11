package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaFZ25Review handles seeding Yamaha FZ25 product review and translation
type YamahaFZ25Review struct {
	BaseSeeder
}

// NewYamahaFZ25ReviewSeeder creates a new YamahaFZ25Review
func NewYamahaFZ25ReviewSeeder() *YamahaFZ25Review {
	return &YamahaFZ25Review{BaseSeeder: BaseSeeder{name: "Yamaha FZ25 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaFZ25Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha FZ25").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha fz25 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding yamaha fz25 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha FZ25 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha FZ25 Review Bangladesh 2024 - Stylish Street Commuter</h1>
<p class="summary-text">The Yamaha FZ25 is a stylish 250cc street commuter with Yamaha's sporty DNA, fuel-injected engine, comfortable seating, and modern design. Priced around ৳3,20,000, it combines daily commuting practicality with attractive street bike styling.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha FZ25 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc Sporty Engine:</strong> <span class="highlight-value">Fuel-injected performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Stylish Design:</strong> <span class="highlight-value">Modern street looks</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seating:</strong> <span class="highlight-value">Daily commute friendly</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">38-42 km/l decent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha FZ25 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Sporty Engine:</strong> <span class="pro-description">250cc fuel-injected</span></li>
<li class="pro-item"><strong class="pro-label">Modern Design:</strong> <span class="pro-description">Stylish street looks</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Riding:</strong> <span class="pro-description">Daily commute suitable</span></li>
<li class="pro-item"><strong class="pro-label">Yamaha Quality:</strong> <span class="pro-description">Reliable performance</span></li>
<li class="pro-item"><strong class="pro-label">Good Handling:</strong> <span class="pro-description">Responsive maneuvers</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha FZ25 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Fuel Efficiency:</strong> <span class="con-description">38-42 km/l average</span></li>
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Firm for long rides</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">Not sport-oriented</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha FZ25 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,20,000 - Street commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳6-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">38-42 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha FZ25 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Performance:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Sporty 250cc</span></li>
<li class="rating-item"><strong class="rating-label">Design & Style:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Modern attractive</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Responsive good</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Yamaha proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good commuter</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha FZ25?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,20,000, the Yamaha FZ25 is perfect for riders seeking a modern street commuter with Yamaha's sporty character, fuel-injection, comfortable styling, and reliability for daily urban riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For stylish street commuting</span></p>
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
		return fmt.Errorf("error creating yamaha fz25 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha FZ25 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা FZ25 রিভিউ বাংলাদেশ 2024 - স্টাইলিশ স্ট্রিট কমিউটার</h1>
<p class="summary-text">ইয়ামাহা FZ25 একটি স্টাইলিশ 250cc স্ট্রিট কমিউটার যা ইয়ামাহার খেলাধুলা DNA, জ্বালানী-ইনজেক্টেড ইঞ্জিন, আরামদায়ক আসন এবং আধুনিক ডিজাইন বৈশিষ্ট্যযুক্ত। ৳3,20,000 টাকায় মূল্যায়িত, এটি আকর্ষণীয় স্ট্রিট বাইক স্টাইলিং সহ দৈনন্দিন যাতায়াত ব্যবহারিকতা একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা FZ25 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc খেলাধুলা ইঞ্জিন:</strong> <span class="highlight-value">জ্বালানী-ইনজেক্টেড কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">স্টাইলিশ ডিজাইন:</strong> <span class="highlight-value">আধুনিক স্ট্রিট চেহারা</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক আসন:</strong> <span class="highlight-value">দৈনিক যাত্রা বান্ধব</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">38-42 km/l মানসম্পন্ন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা FZ25 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">খেলাধুলা ইঞ্জিন:</strong> <span class="pro-description">250cc জ্বালানী-ইনজেক্টেড</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক ডিজাইন:</strong> <span class="pro-description">স্টাইলিশ স্ট্রিট চেহারা</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইডিং:</strong> <span class="pro-description">দৈনিক যাত্রা উপযুক্ত</span></li>
<li class="pro-item"><strong class="pro-label">ইয়ামাহা গুণমান:</strong> <span class="pro-description">নির্ভরযোগ্য কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">ভাল হ্যান্ডলিং:</strong> <span class="pro-description">প্রতিক্রিয়াশীল কৌশল</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা FZ25 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">জ্বালানী দক্ষতা:</strong> <span class="con-description">38-42 km/l গড়</span></li>
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">দীর্ঘ যাত্রার জন্য দৃঢ়</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">ক্রীড়া-ভিত্তিক নয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা FZ25 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,20,000 টাকায়, ইয়ামাহা FZ25 রাইডারদের জন্য নিখুঁত যারা ইয়ামাহার খেলাধুলা চরিত্র, জ্বালানী-ইনজেকশন, আরামদায়ক স্টাইলিং এবং দৈনিক শহুরে রাইডিংয়ের জন্য নির্ভরযোগ্যতা সহ আধুনিক স্ট্রিট কমিউটার খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্টাইলিশ স্ট্রিট যাত্রার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha FZ25 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha FZ25\n")

	return nil
}

// GetName returns the seeder name
func (s *YamahaFZ25Review) GetName() string {
	return "YamahaFZ25Review"
}
