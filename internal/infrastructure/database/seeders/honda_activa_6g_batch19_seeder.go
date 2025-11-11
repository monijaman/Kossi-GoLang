package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaActiva6GReviewBatch19 handles seeding Honda Activa 6G product review and translation
type HondaActiva6GReviewBatch19 struct {
	BaseSeeder
}

// NewHondaActiva6GReviewBatch19Seeder creates a new HondaActiva6GReviewBatch19
func NewHondaActiva6GReviewBatch19Seeder() *HondaActiva6GReviewBatch19 {
	return &HondaActiva6GReviewBatch19{BaseSeeder: BaseSeeder{name: "Honda Activa 6G Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaActiva6GReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda Activa 6G").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda activa 6g product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding honda activa 6g product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda Activa 6G already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda Activa 6G Review Bangladesh 2024 - Practical Scooter Market Icon</h1>
<p class="summary-text">The Honda Activa 6G is a 110cc air-cooled practical scooter representing market leadership in segment. Priced around ৳2,35,000, it delivers proven reliability, excellent fuel efficiency, user-friendly operation, strong resale, and Honda quality for practical urban commuters seeking the market's most trusted scooter platform.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda Activa 6G Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc Air-Cooled:</strong> <span class="highlight-value">Practical scooter</span></li>
<li class="highlight-item"><strong class="highlight-label">Market Leadership:</strong> <span class="highlight-value">Segment icon</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficiency:</strong> <span class="highlight-value">45-50 km/l excellent</span></li>
<li class="highlight-item"><strong class="highlight-label">User Friendly:</strong> <span class="highlight-value">Simple operation</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda Activa 6G Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Market Trusted:</strong> <span class="pro-description">Proven reliability</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Economy:</strong> <span class="pro-description">45-50 km/l excellent</span></li>
<li class="pro-item"><strong class="pro-label">Easy Operation:</strong> <span class="pro-description">User-friendly platform</span></li>
<li class="pro-item"><strong class="pro-label">Strong Resale:</strong> <span class="pro-description">Market demand highest</span></li>
<li class="pro-item"><strong class="pro-label">Honda Quality:</strong> <span class="pro-description">Japanese reliability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda Activa 6G Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Modest Power:</strong> <span class="con-description">8.6 bhp basic</span></li>
<li class="con-item"><strong class="con-label">Storage Limited:</strong> <span class="con-description">Minimal under-seat</span></li>
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">Segment leader premium</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda Activa 6G Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,35,000 - Premium scooter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳9-11 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">45-50 km/l excellent</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">110cc air-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">8.6 bhp practical smooth</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">110kg lightweight</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda Activa 6G Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Market proven</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- 45-50 km/l</span></li>
<li class="rating-item"><strong class="rating-label">User Experience:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Easy operation</span></li>
<li class="rating-item"><strong class="rating-label">Resale Value:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Highest market</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- ৳2,35,000 leader</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda Activa 6G?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,35,000, the Activa 6G is the market's most trusted practical scooter, delivering proven reliability, excellent 45-50 km/l efficiency, user-friendly operation, highest resale value, and Honda quality for practical urban commuters valuing market leadership and dependability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For practical urban commuters</span></p>
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
		return fmt.Errorf("error creating honda activa 6g review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda Activa 6G (Rating: 4.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা অ্যাকটিভা 6G রিভিউ বাংলাদেশ 2024 - ব্যবহারিক স্কুটার মার্কেট আইকন</h1>
<p class="summary-text">হোন্ডা অ্যাকটিভা 6G একটি 110cc এয়ার-কুল্ড ব্যবহারিক স্কুটার যা সেগমেন্ট বাজার নেতৃত্ব প্রতিনিধিত্ব করে। ৳2,35,000 টাকায় মূল্যায়িত, এটি প্রমাণিত নির্ভরযোগ্যতা এবং চমৎকার জ্বালানি দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা অ্যাকটিভা 6G মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc এয়ার-কুল্ড:</strong> <span class="highlight-value">ব্যবহারিক স্কুটার</span></li>
<li class="highlight-item"><strong class="highlight-label">মার্কেট নেতৃত্ব:</strong> <span class="highlight-value">সেগমেন্ট আইকন</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষতা:</strong> <span class="highlight-value">45-50 km/l চমৎকার</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যবহারকারী-বান্ধব:</strong> <span class="highlight-value">সরল অপারেশন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা অ্যাকটিভা 6G সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">মার্কেট বিশ্বস্ত:</strong> <span class="pro-description">প্রমাণিত নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি অর্থনীতি:</strong> <span class="pro-description">45-50 km/l চমৎকার</span></li>
<li class="pro-item"><strong class="pro-label">সহজ অপারেশন:</strong> <span class="pro-description">ব্যবহারকারী-বান্ধব প্ল্যাটফর্ম</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী পুনঃবিক্রয়:</strong> <span class="pro-description">বাজার চাহিদা সর্বোচ্চ</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা গুণমান:</strong> <span class="pro-description">জাপানি নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা অ্যাকটিভা 6G অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বিনম্র শক্তি:</strong> <span class="con-description">8.6 bhp মৌলিক</span></li>
<li class="con-item"><strong class="con-label">স্টোরেজ সীমিত:</strong> <span class="con-description">ন্যূনতম সিট-নিচে</span></li>
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">সেগমেন্ট নেতা প্রিমিয়াম</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা অ্যাকটিভা 6G কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,35,000 টাকায়, অ্যাকটিভা 6G বাজারের সবচেয়ে বিশ্বস্ত ব্যবহারিক স্কুটার, প্রমাণিত নির্ভরযোগ্যতা, চমৎকার 45-50 km/l দক্ষতা এবং সর্বোচ্চ পুনঃবিক্রয় মূল্য প্রদান করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্যবহারিক শহুরে কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda Activa 6G already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda Activa 6G\n")

	return nil
}

// GetName returns the seeder name
func (s *HondaActiva6GReviewBatch19) GetName() string {
	return "HondaActiva6GReviewBatch19"
}
