package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajAvenger160ReviewBatch19 handles seeding Bajaj Avenger 160 product review and translation
type BajajAvenger160ReviewBatch19 struct {
	BaseSeeder
}

// NewBajajAvenger160ReviewBatch19Seeder creates a new BajajAvenger160ReviewBatch19
func NewBajajAvenger160ReviewBatch19Seeder() *BajajAvenger160ReviewBatch19 {
	return &BajajAvenger160ReviewBatch19{BaseSeeder: BaseSeeder{name: "Bajaj Avenger 160 Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajAvenger160ReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Avenger 160").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj avenger 160 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj avenger 160 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Avenger 160 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Avenger 160 Review Bangladesh 2024 - Casual Cruiser Budget Hero</h1>
<p class="summary-text">The Bajaj Avenger 160 is a 160cc air-cooled casual cruiser combining relaxed riding with affordable pricing. Priced around ৳2,75,000, it delivers comfortable seating, lazy torque character, straightforward mechanics, excellent fuel efficiency, and the Bajaj reliability for casual riders seeking undemanding cruiser lifestyle on budget.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Avenger 160 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc Air-Cooled:</strong> <span class="highlight-value">Casual cruiser</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfort Focus:</strong> <span class="highlight-value">Relaxed riding</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Price:</strong> <span class="highlight-value">৳2,75,000 accessible</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficient:</strong> <span class="highlight-value">32-38 km/l excellent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Avenger 160 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Comfort Seating:</strong> <span class="pro-description">Relaxed positioning</span></li>
<li class="pro-item"><strong class="pro-label">Lazy Torque:</strong> <span class="pro-description">Casual character</span></li>
<li class="pro-item"><strong class="pro-label">Budget Price:</strong> <span class="pro-description">৳2,75,000 affordable</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Economy:</strong> <span class="pro-description">32-38 km/l excellent</span></li>
<li class="pro-item"><strong class="pro-label">Simple Mechanics:</strong> <span class="pro-description">Easy maintenance</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Avenger 160 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Lazy Performance:</strong> <span class="con-description">14.1 bhp modest</span></li>
<li class="con-item"><strong class="con-label">Limited Storage:</strong> <span class="con-description">Basic luggage mounting</span></li>
<li class="con-item"><strong class="con-label">Build Simple:</strong> <span class="con-description">Budget segment typical</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Avenger 160 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,75,000 - Budget cruiser</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Ultra-low - ৳6-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">32-38 km/l exceptional</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">160cc air-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">14.1 bhp lazy torque</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">142kg lightweight</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Avenger 160 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Relaxed cruiser</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- 32-38 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Affordability:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- ৳2,75,000 great</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Bajaj proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Casual excellence</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Avenger 160?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,75,000, the Avenger 160 is ideal for casual riders seeking relaxed cruiser lifestyle, exceptional 32-38 km/l efficiency, comfortable budget-friendly commuting, and Bajaj's proven reliability for undemanding daily riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For casual cruiser riders</span></p>
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
		return fmt.Errorf("error creating bajaj avenger 160 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Avenger 160 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ অ্যাভেঞ্জার 160 রিভিউ বাংলাদেশ 2024 - ক্যাজুয়াল ক্রুজার বাজেট হিরো</h1>
<p class="summary-text">বাজাজ অ্যাভেঞ্জার 160 একটি 160cc এয়ার-কুল্ড ক্যাজুয়াল ক্রুজার যা শান্ত রাইডিংকে সাশ্রয়ী মূল্যের সাথে একত্রিত করে। ৳2,75,000 টাকায় মূল্যায়িত, এটি আরামদায়ক সিটিং এবং অলস টর্ক চরিত্র প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ অ্যাভেঞ্জার 160 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc এয়ার-কুল্ড:</strong> <span class="highlight-value">ক্যাজুয়াল ক্রুজার</span></li>
<li class="highlight-item"><strong class="highlight-label">আরাম ফোকাস:</strong> <span class="highlight-value">শান্ত রাইডিং</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট মূল্য:</strong> <span class="highlight-value">৳2,75,000 অ্যাক্সেসযোগ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষ:</strong> <span class="highlight-value">32-38 km/l চমৎকার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ অ্যাভেঞ্জার 160 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আরাম সিটিং:</strong> <span class="pro-description">শান্ত পজিশনিং</span></li>
<li class="pro-item"><strong class="pro-label">অলস টর্ক:</strong> <span class="pro-description">ক্যাজুয়াল চরিত্র</span></li>
<li class="pro-item"><strong class="pro-label">বাজেট মূল্য:</strong> <span class="pro-description">৳2,75,000 সাশ্রয়ী</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি অর্থনীতি:</strong> <span class="pro-description">32-38 km/l চমৎকার</span></li>
<li class="pro-item"><strong class="pro-label">সরল মেকানিক্স:</strong> <span class="pro-description">সহজ রক্ষণাবেক্ষণ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ অ্যাভেঞ্জার 160 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অলস পারফরম্যান্স:</strong> <span class="con-description">14.1 bhp বিনম্র</span></li>
<li class="con-item"><strong class="con-label">স্টোরেজ সীমিত:</strong> <span class="con-description">মৌলিক লাগেজ মাউন্টিং</span></li>
<li class="con-item"><strong class="con-label">নির্মাণ সরল:</strong> <span class="con-description">বাজেট সেগমেন্ট সাধারণ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ অ্যাভেঞ্জার 160 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,75,000 টাকায়, অ্যাভেঞ্জার 160 ক্যাজুয়াল রাইডারদের জন্য আদর্শ যারা শান্ত ক্রুজার লাইফস্টাইল, ব্যতিক্রমী 32-38 km/l দক্ষতা এবং সাশ্রয়ী বাজেট কমিউটিং চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্যাজুয়াল ক্রুজার রাইডারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Avenger 160 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Avenger 160\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajAvenger160ReviewBatch19) GetName() string {
	return "BajajAvenger160ReviewBatch19"
}
