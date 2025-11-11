package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxer155Review handles seeding Suzuki Gixxer 155 product review and translation
type SuzukiGixxer155Review struct {
	BaseSeeder
}

// NewSuzukiGixxer155ReviewSeeder creates a new SuzukiGixxer155Review
func NewSuzukiGixxer155ReviewSeeder() *SuzukiGixxer155Review {
	return &SuzukiGixxer155Review{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer 155 Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxer155Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer 155").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gixxer 155 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding suzuki gixxer 155 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer 155 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer 155 Review Bangladesh 2024 - Street Fighter Excellence</h1>
<p class="summary-text">The Suzuki Gixxer 155 is a sporty street bike featuring 155cc engine, ABS safety, aggressive design, and excellent handling. Priced around ৳2,00,000, it delivers genuine sport performance with Suzuki reliability for riders seeking thrilling daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer 155 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">155cc Engine:</strong> <span class="highlight-value">Powerful street fighter engine</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS Safety:</strong> <span class="highlight-value">Front wheel ABS protection</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Design:</strong> <span class="highlight-value">Sporty street fighter styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Handling:</strong> <span class="highlight-value">Responsive and agile cornering</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer 155 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Performance:</strong> <span class="pro-description">Powerful 155cc with good acceleration</span></li>
<li class="pro-item"><strong class="pro-label">ABS Protection:</strong> <span class="pro-description">Front ABS for safety</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Eye-catching street fighter design</span></li>
<li class="pro-item"><strong class="pro-label">Great Handling:</strong> <span class="pro-description">Responsive and agile performance</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">40-45 km/l efficiency</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer 155 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">More expensive than budget 150cc</span></li>
<li class="con-item"><strong class="con-label">Sport Focused:</strong> <span class="con-description">Less comfortable for pillion riders</span></li>
<li class="con-item"><strong class="con-label">Maintenance Cost:</strong> <span class="con-description">Higher service expenses</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer 155 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,00,000 - Performance sport value</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳5-6 per km (40-45 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer 155 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good 155cc power</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent agility</span></li>
<li class="rating-item"><strong class="rating-label">Safety:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Front ABS equipped</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Aggressive styling</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good sport value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer 155?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,00,000, the Suzuki Gixxer 155 is ideal for young riders seeking aggressive styling, strong performance, and excellent handling with sport-focused features and ABS safety.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance-oriented sport riders</span></p>
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
		return fmt.Errorf("error creating suzuki gixxer 155 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer 155 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি Gixxer 155 রিভিউ বাংলাদেশ 2024 - স্ট্রিট ফাইটার এক্সিলেন্স</h1>
<p class="summary-text">সুজুকি Gixxer 155 একটি খেলাধুলাপূর্ণ স্ট্রিট বাইক যা 155cc ইঞ্জিন, ABS নিরাপত্তা, আক্রমণাত্মক ডিজাইন এবং চমৎকার হ্যান্ডলিং বৈশিষ্ট্যযুক্ত। ৳2,00,000 টাকায় মূল্যায়িত, এটি রোমাঞ্চকর দৈনন্দিন যাতায়াত খুঁজছেন রাইডারদের জন্য খাঁটি খেলা পারফরমেন্স সরবরাহ করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি Gixxer 155 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">155cc ইঞ্জিন:</strong> <span class="highlight-value">শক্তিশালী স্ট্রিট ফাইটার ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS নিরাপত্তা:</strong> <span class="highlight-value">সামনের চাকা ABS সুরক্ষা</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="highlight-value">খেলাধুলাপূর্ণ স্ট্রিট ফাইটার স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার হ্যান্ডলিং:</strong> <span class="highlight-value">প্রতিক্রিয়াশীল এবং চটপটে কোণ দেওয়া</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি Gixxer 155 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী পারফরমেন্স:</strong> <span class="pro-description">ভাল ত্বরণ সহ শক্তিশালী 155cc</span></li>
<li class="pro-item"><strong class="pro-label">ABS সুরক্ষা:</strong> <span class="pro-description">নিরাপত্তার জন্য সামনে ABS</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">চোখ ধাঁধানো স্ট্রিট ফাইটার ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত হ্যান্ডলিং:</strong> <span class="pro-description">প্রতিক্রিয়াশীল এবং চটপটে পারফরমেন্স</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">40-45 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি Gixxer 155 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চতর মূল্য:</strong> <span class="con-description">বাজেট 150cc এর চেয়ে বেশি ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">খেলা ফোকাসড:</strong> <span class="con-description">পিলিয়ন রাইডারদের জন্য কম আরামদায়ক</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">উচ্চতর সেবা খরচ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি Gixxer 155 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,00,000 টাকায়, সুজুকি Gixxer 155 যুব রাইডারদের জন্য আদর্শ যারা আক্রমণাত্মক স্টাইলিং, শক্তিশালী পারফরমেন্স এবং চমৎকার হ্যান্ডলিং খেলা-ফোকাস বৈশিষ্ট্য এবং ABS নিরাপত্তা সহ খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারফরমেন্স-উন্মুখ খেলা রাইডারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer 155 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer 155\n")

	return nil
}

// GetName returns the seeder name
func (s *SuzukiGixxer155Review) GetName() string {
	return "SuzukiGixxer155Review"
}
