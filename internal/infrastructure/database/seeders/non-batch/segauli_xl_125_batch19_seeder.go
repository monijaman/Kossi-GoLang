package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SegaXL125ReviewBatch19 handles seeding Segauli XL 125 product review and translation
type SegaXL125ReviewBatch19 struct {
	BaseSeeder
}

// NewSegaXL125ReviewBatch19Seeder creates a new SegaXL125ReviewBatch19
func NewSegaXL125ReviewBatch19Seeder() *SegaXL125ReviewBatch19 {
	return &SegaXL125ReviewBatch19{BaseSeeder: BaseSeeder{name: "Segauli XL 125 Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *SegaXL125ReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Segauli XL 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("segauli xl 125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding segauli xl 125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Segauli XL 125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Segauli XL 125 Review Bangladesh 2024 - Budget Commuter Powerful Engine</h1>
<p class="summary-text">The Segauli XL 125 is a 125cc air-cooled commuter delivering practical features with affordable pricing. Priced around ৳1,65,000, it combines modest power, comfortable seating, good fuel efficiency, sturdy construction, and accessibility for daily commuters seeking reliable transportation with acceptable performance standards.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Segauli XL 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc Air-Cooled:</strong> <span class="highlight-value">Modest commuter power</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Friendly:</strong> <span class="highlight-value">৳1,65,000 accessible</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seating:</strong> <span class="highlight-value">Ergonomic positioning</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficient:</strong> <span class="highlight-value">50-55 km/l excellent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Segauli XL 125 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Budget Price:</strong> <span class="pro-description">৳1,65,000 accessible</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Economy:</strong> <span class="pro-description">50-55 km/l excellent</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seat:</strong> <span class="pro-description">Long commutes friendly</span></li>
<li class="pro-item"><strong class="pro-label">Sturdy Build:</strong> <span class="pro-description">Durable construction</span></li>
<li class="pro-item"><strong class="pro-label">Low Cost:</strong> <span class="pro-description">Maintenance affordable</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Segauli XL 125 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Modest Power:</strong> <span class="con-description">10.5 bhp limited</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Minimal equipment</span></li>
<li class="con-item"><strong class="con-label">Build Quality:</strong> <span class="con-description">Budget segment typical</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Segauli XL 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,65,000 - Budget commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Ultra-low - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">50-55 km/l exceptional</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">125cc air-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">10.5 bhp commuter modest</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">125kg lightweight</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Segauli XL 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- 50-55 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Affordability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- ৳1,65,000 good</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Commute friendly</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Budget segment</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Budget excellent</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Segauli XL 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,65,000, the XL 125 is ideal for budget-conscious commuters seeking exceptional 50-55 km/l fuel efficiency, comfortable seating, sturdy construction, and affordable daily transportation with modest but adequate performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating segauli xl 125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Segauli XL 125 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সেগাউলি XL 125 রিভিউ বাংলাদেশ 2024 - বাজেট কমিউটার শক্তিশালী ইঞ্জিন</h1>
<p class="summary-text">সেগাউলি XL 125 একটি 125cc এয়ার-কুল্ড কমিউটার যা ব্যবহারিক বৈশিষ্ট্য সাশ্রয়ী মূল্যের সাথে প্রদান করে। ৳1,65,000 টাকায় মূল্যায়িত, এটি বিনম্র শক্তি এবং আরামদায়ক সিটিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সেগাউলি XL 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc এয়ার-কুল্ড:</strong> <span class="highlight-value">বিনম্র কমিউটার শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট বান্ধব:</strong> <span class="highlight-value">৳1,65,000 অ্যাক্সেসযোগ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">আরাম সিটিং:</strong> <span class="highlight-value">এরগোনমিক পজিশনিং</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষ:</strong> <span class="highlight-value">50-55 km/l চমৎকার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সেগাউলি XL 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">বাজেট মূল্য:</strong> <span class="pro-description">৳1,65,000 অ্যাক্সেসযোগ্য</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি অর্থনীতি:</strong> <span class="pro-description">50-55 km/l চমৎকার</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক আসন:</strong> <span class="pro-description">দীর্ঘ যাত্রা বান্ধব</span></li>
<li class="pro-item"><strong class="pro-label">মজবুত নির্মাণ:</strong> <span class="pro-description">টেকসই গঠন</span></li>
<li class="pro-item"><strong class="pro-label">কম খরচ:</strong> <span class="pro-description">রক্ষণাবেক্ষণ সাশ্রয়ী</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সেগাউলি XL 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বিনম্র শক্তি:</strong> <span class="con-description">10.5 bhp সীমিত</span></li>
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতম সরঞ্জাম</span></li>
<li class="con-item"><strong class="con-label">বিল্ড কোয়ালিটি:</strong> <span class="con-description">বাজেট সেগমেন্ট সাধারণ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সেগাউলি XL 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,65,000 টাকায়, XL 125 বাজেট-সচেতন কমিউটারদের জন্য আদর্শ যারা ব্যতিক্রমী 50-55 km/l জ্বালানি দক্ষতা এবং সাশ্রয়ী দৈনিক পরিবহন চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Segauli XL 125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Segauli XL 125\n")

	return nil
}

// GetName returns the seeder name
func (s *SegaXL125ReviewBatch19) GetName() string {
	return "SegaXL125ReviewBatch19"
}
