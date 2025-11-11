package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSApacheRTR180ReviewBatch18 handles seeding TVS Apache RTR 180 product review and translation
type TVSApacheRTR180ReviewBatch18 struct {
	BaseSeeder
}

// NewTVSApacheRTR180ReviewBatch18Seeder creates a new TVSApacheRTR180ReviewBatch18
func NewTVSApacheRTR180ReviewBatch18Seeder() *TVSApacheRTR180ReviewBatch18 {
	return &TVSApacheRTR180ReviewBatch18{BaseSeeder: BaseSeeder{name: "TVS Apache RTR 180 Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *TVSApacheRTR180ReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 180").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs apache rtr 180 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding tvs apache rtr 180 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Apache RTR 180 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Apache RTR 180 Review Bangladesh 2024 - Performance Commuter Legendary Status</h1>
<p class="summary-text">The TVS Apache RTR 180 is a 180cc liquid-cooled performance commuter delivering track-bred DNA with affordable pricing. Priced around ৳2,95,000, it combines aggressive styling, precise handling, fuel efficiency, and the Apache RTR pedigree for enthusiasts seeking premium commuting with sport character and reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 180 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">180cc Liquid-Cooled:</strong> <span class="highlight-value">Performance commuting</span></li>
<li class="highlight-item"><strong class="highlight-label">Track Heritage:</strong> <span class="highlight-value">Racing DNA integrated</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Sports character</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficient:</strong> <span class="highlight-value">30-32 km/l excellent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 180 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">RTR Performance:</strong> <span class="pro-description">Racing pedigree proven</span></li>
<li class="pro-item"><strong class="pro-label">Handling Precision:</strong> <span class="pro-description">Cornering sharp</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Efficiency:</strong> <span class="pro-description">30-32 km/l excellent</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Design:</strong> <span class="pro-description">Sporty presence</span></li>
<li class="pro-item"><strong class="pro-label">Reliability:</strong> <span class="pro-description">TVS dependability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 180 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-focused seating</span></li>
<li class="con-item"><strong class="con-label">Storage Limited:</strong> <span class="con-description">Minimal luggage space</span></li>
<li class="con-item"><strong class="con-label">Maintenance:</strong> <span class="con-description">Regular service needs</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 180 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,95,000 - Performance commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳8-10 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">30-32 km/l excellent</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">180cc liquid-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">17 bhp racing tuned</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">148kg agile commuter</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 180 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Racing tuned</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Track precise</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- 30-32 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Aggressive sport</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- ৳2,95,000 excellent</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 180?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,95,000, the Apache RTR 180 is ideal for performance commuters seeking racing DNA with aggressive styling, precise handling, excellent fuel efficiency, and the TVS reliability badge for daily sport riding with competitive value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance commuters</span></p>
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
		return fmt.Errorf("error creating tvs apache rtr 180 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Apache RTR 180 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">টিভিএস অ্যাপাচে RTR 180 রিভিউ বাংলাদেশ 2024 - পারফরম্যান্স কমিউটার কিংবদন্তি স্ট্যাটাস</h1>
<p class="summary-text">টিভিএস অ্যাপাচে RTR 180 একটি 180cc তরল-শীতলকৃত পারফরম্যান্স কমিউটার যা ট্র্যাক-ব্রেড DNA প্রদান করে। ৳2,95,000 টাকায় মূল্যায়িত, এটি আক্রমণাত্মক স্টাইলিং, নির্ভুল হ্যান্ডলিং এবং অ্যাপাচে RTR বংশ একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস অ্যাপাচে RTR 180 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">180cc তরল-শীতলকৃত:</strong> <span class="highlight-value">পারফরম্যান্স কমিউটিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্র্যাক হেরিটেজ:</strong> <span class="highlight-value">রেসিং DNA একীভূত</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="highlight-value">স্পোর্টস চরিত্র</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষ:</strong> <span class="highlight-value">30-32 km/l চমৎকার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস অ্যাপাচে RTR 180 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">RTR পারফরম্যান্স:</strong> <span class="pro-description">রেসিং বংশ প্রমাণিত</span></li>
<li class="pro-item"><strong class="pro-label">হ্যান্ডলিং নির্ভুলতা:</strong> <span class="pro-description">কর্নারিং তীক্ষ্ণ</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি দক্ষতা:</strong> <span class="pro-description">30-32 km/l চমৎকার</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="pro-description">ক্রীড়া উপস্থিতি</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্যতা:</strong> <span class="pro-description">টিভিএস নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস অ্যাপাচে RTR 180 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সিট আরাম:</strong> <span class="con-description">স্পোর্ট-ফোকাসড সিটিং</span></li>
<li class="con-item"><strong class="con-label">স্টোরেজ সীমিত:</strong> <span class="con-description">ন্যূনতম লাগেজ স্থান</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ:</strong> <span class="con-description">নিয়মিত সেবা প্রয়োজন</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: টিভিএস অ্যাপাচে RTR 180 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,95,000 টাকায়, অ্যাপাচে RTR 180 পারফরম্যান্স কমিউটারদের জন্য আদর্শ যারা রেসিং DNA, আক্রমণাত্মক স্টাইলিং এবং চমৎকার জ্বালানি দক্ষতা চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারফরম্যান্স কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Apache RTR 180 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Apache RTR 180\n")

	return nil
}

// GetName returns the seeder name
func (s *TVSApacheRTR180ReviewBatch18) GetName() string {
	return "TVSApacheRTR180ReviewBatch18"
}
