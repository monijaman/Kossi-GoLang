package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxer250ReviewBatch19 handles seeding Suzuki Gixxer 250 product review and translation
type SuzukiGixxer250ReviewBatch19 struct {
	BaseSeeder
}

// NewSuzukiGixxer250ReviewBatch19Seeder creates a new SuzukiGixxer250ReviewBatch19
func NewSuzukiGixxer250ReviewBatch19Seeder() *SuzukiGixxer250ReviewBatch19 {
	return &SuzukiGixxer250ReviewBatch19{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer 250 Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxer250ReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer 250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gixxer 250 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding suzuki gixxer 250 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer 250 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer 250 Review Bangladesh 2024 - Sporty Commuter Reliable Partner</h1>
<p class="summary-text">The Suzuki Gixxer 250 is a 250cc liquid-cooled sports commuter combining aggressive styling with practical reliability. Priced around ৳3,85,000, it delivers responsive engine characteristics, sharp handling, efficient performance, excellent build quality, and Suzuki's dependability for enthusiasts seeking daily performance commuting with consistent reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer 250 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc Liquid-Cooled:</strong> <span class="highlight-value">Sporty performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Sports presence</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Nimble responsive</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficient:</strong> <span class="highlight-value">28-32 km/l excellent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer 250 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Sporty DNA:</strong> <span class="pro-description">250cc aggressive</span></li>
<li class="pro-item"><strong class="pro-label">Responsive Engine:</strong> <span class="pro-description">23 bhp sharp delivery</span></li>
<li class="pro-item"><strong class="pro-label">Build Quality:</strong> <span class="pro-description">Suzuki excellent</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Efficiency:</strong> <span class="pro-description">28-32 km/l exceptional</span></li>
<li class="pro-item"><strong class="pro-label">Reliability:</strong> <span class="pro-description">Suzuki proven</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer 250 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-focused seating</span></li>
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">৳3,85,000 investment</span></li>
<li class="con-item"><strong class="con-label">Storage Limited:</strong> <span class="con-description">Minimal luggage</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer 250 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,85,000 - Sports commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳10-12 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">28-32 km/l exceptional</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">250cc liquid-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">23 bhp sharp sports</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">155kg sporty lightweight</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer 250 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Sporty Character:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Aggressive DNA</span></li>
<li class="rating-item"><strong class="rating-label">Engine Responsiveness:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Sharp delivery</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Suzuki excellent</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- 28-32 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- ৳3,85,000 solid</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer 250?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,85,000, the Gixxer 250 is ideal for sporty commuters seeking aggressive 250cc styling, responsive sharp performance, excellent build quality, exceptional 28-32 km/l efficiency, and Suzuki's proven reliability for daily performance riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sporty commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.7,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating suzuki gixxer 250 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer 250 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার 250 রিভিউ বাংলাদেশ 2024 - ক্রীড়াবিদ কমিউটার নির্ভরযোগ্য অংশীদার</h1>
<p class="summary-text">সুজুকি গিক্সার 250 একটি 250cc তরল-শীতলকৃত স্পোর্টস কমিউটার যা আক্রমণাত্মক স্টাইলিংকে ব্যবহারিক নির্ভরযোগ্যতার সাথে একত্রিত করে। ৳3,85,000 টাকায় মূল্যায়িত, এটি প্রতিক্রিয়াশীল ইঞ্জিন বৈশিষ্ট্য এবং তীক্ষ্ণ হ্যান্ডলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার 250 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc তরল-শীতলকৃত:</strong> <span class="highlight-value">ক্রীড়াবিদ পারফরম্যান্স</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="highlight-value">ক্রীড়া উপস্থিতি</span></li>
<li class="highlight-item"><strong class="highlight-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="highlight-value">চতুর প্রতিক্রিয়াশীল</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষ:</strong> <span class="highlight-value">28-32 km/l চমৎকার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার 250 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ক্রীড়া DNA:</strong> <span class="pro-description">250cc আক্রমণাত্মক</span></li>
<li class="pro-item"><strong class="pro-label">প্রতিক্রিয়াশীল ইঞ্জিন:</strong> <span class="pro-description">23 bhp তীক্ষ্ণ ডেলিভারি</span></li>
<li class="pro-item"><strong class="pro-label">বিল্ড কোয়ালিটি:</strong> <span class="pro-description">সুজুকি চমৎকার</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি দক্ষতা:</strong> <span class="pro-description">28-32 km/l ব্যতিক্রমী</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্যতা:</strong> <span class="pro-description">সুজুকি প্রমাণিত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার 250 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সিট আরাম:</strong> <span class="con-description">স্পোর্ট-ফোকাসড সিটিং</span></li>
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳3,85,000 বিনিয়োগ</span></li>
<li class="con-item"><strong class="con-label">স্টোরেজ সীমিত:</strong> <span class="con-description">ন্যূনতম লাগেজ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি গিক্সার 250 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,85,000 টাকায়, গিক্সার 250 ক্রীড়াবিদ কমিউটারদের জন্য আদর্শ যারা আক্রমণাত্মক 250cc স্টাইলিং, তীক্ষ্ণ পারফরম্যান্স এবং সুজুকি প্রমাণিত নির্ভরযোগ্যতা চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্রীড়াবিদ কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer 250 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer 250\n")

	return nil
}

// GetName returns the seeder name
func (s *SuzukiGixxer250ReviewBatch19) GetName() string {
	return "SuzukiGixxer250ReviewBatch19"
}
