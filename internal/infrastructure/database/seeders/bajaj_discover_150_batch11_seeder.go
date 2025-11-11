package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajDiscover150Review handles seeding Bajaj Discover 150 product review and translation
type BajajDiscover150Review struct {
	BaseSeeder
}

// NewBajajDiscover150ReviewSeeder creates a new BajajDiscover150Review
func NewBajajDiscover150ReviewSeeder() *BajajDiscover150Review {
	return &BajajDiscover150Review{BaseSeeder: BaseSeeder{name: "Bajaj Discover 150 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajDiscover150Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Discover 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj discover 150 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj discover 150 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Discover 150 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Discover 150 Review Bangladesh 2024 - Reliable Workhorse</h1>
<p class="summary-text">The Bajaj Discover 150 is a robust commuter motorcycle featuring 150cc air-cooled engine, practical design, reliable performance, and exceptional fuel economy. Priced around ৳1,65,000, it's the perfect workhorse for daily commuting with proven durability and minimal maintenance requirements.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Discover 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc Air-Cooled:</strong> <span class="highlight-value">Proven reliable engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Practical Styling:</strong> <span class="highlight-value">Functional and timeless design</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Durability:</strong> <span class="highlight-value">Built to last for years</span></li>
<li class="highlight-item"><strong class="highlight-label">Low Running Cost:</strong> <span class="highlight-value">Economic commuter option</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Discover 150 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Outstanding Mileage:</strong> <span class="pro-description">40-45 km/l exceptional efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Proven Reliability:</strong> <span class="pro-description">Trusted workhorse for years</span></li>
<li class="pro-item"><strong class="pro-label">Easy Maintenance:</strong> <span class="pro-description">Simple design with affordable service</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seating:</strong> <span class="pro-description">Spacious rider and pillion area</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Great value for money</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Discover 150 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Design:</strong> <span class="con-description">Minimal modern styling features</span></li>
<li class="con-item"><strong class="con-label">No ABS:</strong> <span class="con-description">Manual brake modulation needed</span></li>
<li class="con-item"><strong class="con-label">Modest Power:</strong> <span class="con-description">Adequate but not exciting performance</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Discover 150 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,65,000 - Economical workhorse value</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳3-4 per km (40-45 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Discover 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent 40-45 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Proven durability</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Great budget option</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Spacious seating</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Adequate commuter power</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Discover 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,65,000, the Bajaj Discover 150 is the ideal choice for practical commuters seeking proven reliability, exceptional fuel economy, and minimal maintenance with outstanding value for money.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For practical daily commuters</span></p>
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
		return fmt.Errorf("error creating bajaj discover 150 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Discover 150 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ ডিসকভার १५० রিভিউ বাংলাদেশ २०२४ - নির্ভরযোগ্য कार्यहर्स</h1>
<p class="summary-text">বাজাজ ডিসকভার १५० একটি শক্তিশালী কমিউটার মোটরসাইকেল যা १५०cc এয়ার-কুল্ড ইঞ্জিন, ব্যবহারিক ডিজাইন, নির্ভরযোগ্য পারফরমেন্স এবং ব্যতিক্রমী জ্বালানি অর্থনীতি বৈশিষ্ট্যযুক্ত। ৳१,६५,००० টাকায় মূল্যায়িত, এটি প্রমাণিত স্থায়িত্ব এবং ন্যূনতম রক্ষণাবেক্ষণ প্রয়োজনীয়তা সহ দৈনন্দিন যাতায়াতের জন্য নিখুঁত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">बाजाज डिस्कवर १५० मुख्य विशेषताएं</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">१५०cc एयर-कुल्ड:</strong> <span class="highlight-value">প্রমাণিত নির্ভরযোগ্য ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">Practical Styling:</strong> <span class="highlight-value">कार्যात्मक এবং কালজয়ী ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">उत्कृष्ट स्थायिত्व:</strong> <span class="highlight-value">বছরের পর বছর স্থায়ী হওয়ার জন্য নির্মিত</span></li>
<li class="highlight-item"><strong class="highlight-label">कम चलमान খरch:</strong> <span class="highlight-value">आर्थिक कमियूटर विकल्प</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">बाजाज डिस्कवर १५० फायदे</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অসাধারণ মাইলেজ:</strong> <span class="pro-description">ব্যতিক্রমী ४०-४५ km/l দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">প্রমাণিত নির্ভরযোগ্যতা:</strong> <span class="pro-description">বছরের জন্য বিশ্বস্ত कार्यहर्स</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সাশ্রয়ী সেবা সহ সহজ ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক সিটিং:</strong> <span class="pro-description">প্রশস্ত রাইডার এবং পিলিয়ন এরিয়া</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">অর্থের জন্য দুর্দান্ত মূল্য</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">बाजाज डिस्कवर १५० अपूर्णताएं</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক ডিজাইন:</strong> <span class="con-description">न्यूनतम आधুনिक स्टाइलिंग বৈশিষ্ট্য</span></li>
<li class="con-item"><strong class="con-label">কোন ABS নেই:</strong> <span class="con-description">म्यानुअल ब्रेक মডুলেশন প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">বিনম্র শক্তি:</strong> <span class="con-description">পর্যাপ্ত কিন্তু উত্তেজনাপূর্ণ পারফরমেন्स নয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">चूड़ांत निर्णय: बाजाज डिस्कवर १५० किनें?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">४.४/५.०</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳१,६५,००० টাকায়, বাজাज ডিসকভার १५० যারা প্রমাণিত নির্ভরযোগ্যতা, ব্যতিক্রমী জ্বালানি অর্থনীতি এবং ন্যূনতম রক্ষণাবেক্ষণ সহ ব্যবহারিক অর্থনীতির জন্য খুঁজছেন তাদের জন্য আদর্শ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারिश:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্যবহারিক দৈनिक कमीউटারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Discover 150 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Discover 150\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajDiscover150Review) GetName() string {
	return "BajajDiscover150Review"
}
