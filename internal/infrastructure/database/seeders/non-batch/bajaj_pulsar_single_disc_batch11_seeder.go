package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarSingleDiscReview handles seeding Bajaj Pulsar Single Disc product review and translation
type BajajPulsarSingleDiscReview struct {
	BaseSeeder
}

// NewBajajPulsarSingleDiscReviewSeeder creates a new BajajPulsarSingleDiscReview
func NewBajajPulsarSingleDiscReviewSeeder() *BajajPulsarSingleDiscReview {
	return &BajajPulsarSingleDiscReview{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar Single Disc Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarSingleDiscReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar Single Disc").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar single disc product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj pulsar single disc product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar Single Disc already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar Single Disc Review Bangladesh 2024 - Budget Commuter</h1>
<p class="summary-text">The Bajaj Pulsar Single Disc is an economical commuter bike featuring single front disc brake, 150cc engine, practical design, and excellent mileage. Priced around ৳1,45,000, it offers reliable performance for daily commuting with impressive fuel efficiency and lower maintenance costs.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar Single Disc Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc DTS-i Engine:</strong> <span class="highlight-value">Proven reliable commuter engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Single Disc Front:</strong> <span class="highlight-value">Cost-effective braking system</span></li>
<li class="highlight-item"><strong class="highlight-label">Practical Design:</strong> <span class="highlight-value">Simple and serviceable construction</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Friendly:</strong> <span class="highlight-value">Affordable purchase price</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar Single Disc Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">42-48 km/l outstanding efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Low Purchase Price:</strong> <span class="pro-description">Budget-conscious commuter option</span></li>
<li class="pro-item"><strong class="pro-label">Easy Maintenance:</strong> <span class="pro-description">Simple design and widely available parts</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Proven 150cc DTS-i powerplant</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seating:</strong> <span class="pro-description">Spacious rider and pillion seat</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar Single Disc Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Single Front Disc:</strong> <span class="con-description">Basic braking compared to twin disc</span></li>
<li class="con-item"><strong class="con-label">No ABS:</strong> <span class="con-description">Manual brake modulation required</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Minimal modern amenities</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar Single Disc Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,45,000 - Budget commuter value</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳3-4 per km (42-48 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar Single Disc Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Exceptional 42-48 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Proven engine platform</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent budget option</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Spacious seating</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Adequate commuter power</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar Single Disc?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,45,000, the Bajaj Pulsar Single Disc is an excellent choice for budget-conscious commuters seeking reliable daily transportation with exceptional fuel economy and minimal running costs.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For economical daily commuting</span></p>
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
		return fmt.Errorf("error creating bajaj pulsar single disc review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar Single Disc (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার সিঙ্গেল ডিস্ক রিভিউ বাংলাদেশ 2024 - বাজেট কমিউটার</h1>
<p class="summary-text">বাজাজ পালসার সিঙ্গেল ডিস্ক একটি অর্থনৈতিক কমিউটার বাইক যা সিঙ্গেল ফ্রন্ট ডিস্ক ব্রেক, 150cc ইঞ্জিন, ব্যবহারিক ডিজাইন এবং চমৎকার মাইলেজ বৈশিষ্ট্যযুক্ত। ৳1,45,000 টাকায় মূল্যায়িত, এটি দৈনন্দিন যাতায়াতের জন্য নির্ভরযোগ্য পারফরমেন্স প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার সিঙ্গেল ডিস্ক মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc DTS-i ইঞ্জিন:</strong> <span class="highlight-value">প্রমাণিত নির্ভরযোগ্য কমিউটার ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">সিঙ্গেল ডিস্ক ফ্রন্ট:</strong> <span class="highlight-value">সাশ্রয়ী ব্রেকিং সিস্টেম</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যবহারিক ডিজাইন:</strong> <span class="highlight-value">সহজ এবং সেবাযোগ্য নির্মাণ</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট বান্ধব:</strong> <span class="highlight-value">সাশ্রয়ী ক্রয় মূল্য</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার সিঙ্গেল ডিস্ক সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">অসাধারণ 42-48 km/l দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">কম ক্রয় মূল্য:</strong> <span class="pro-description">বাজেট-সচেতন কমিউটার অপশন</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সহজ ডিজাইন এবং সহজলভ্য যন্ত্রাংশ</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">প্রমাণিত 150cc DTS-i পাওয়ারপ্লান্ট</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক সিটিং:</strong> <span class="pro-description">প্রশস্ত রাইডার এবং পিলিয়ন সিট</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার সিঙ্গেল ডিস্ক অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সিঙ্গেল ফ্রন্ট ডিস্ক:</strong> <span class="con-description">টুইন ডিস্কের তুলনায় মৌলিক ব্রেকিং</span></li>
<li class="con-item"><strong class="con-label">কোন ABS নেই:</strong> <span class="con-description">ম্যানুয়াল ব্রেক মডুলেশন প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতম আধুনিক সুবিধা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ পালসার সিঙ্গেল ডিস্ক কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,45,000 টাকায়, বাজাজ পালসার সিঙ্গেল ডিস্ক বাজেট-সচেতন কমিউটারদের জন্য চমৎকার পছন্দ যারা অসাধারণ জ্বালানি অর্থনীতি এবং ন্যূনতম চলমান খরচ সহ নির্ভরযোগ্য দৈনন্দিন পরিবহন খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অর্থনৈতিক দৈনন্দিন যাতায়াতের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar Single Disc already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar Single Disc\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsarSingleDiscReview) GetName() string {
	return "BajajPulsarSingleDiscReview"
}
