package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HeroSplendorPlusBatch25 struct {
	BaseSeeder
}

func NewHeroSplendorPlusBatch25Seeder() *HeroSplendorPlusBatch25 {
	return &HeroSplendorPlusBatch25{BaseSeeder: BaseSeeder{name: "Hero Splendor Plus Batch25 Review"}}
}

func (s *HeroSplendorPlusBatch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Splendor Plus").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero splendor plus product not found")
		}
		return fmt.Errorf("error finding product: %w", err)
	}

	var existingReview models.ProductReviewModel
	if err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error; err == nil {
		fmt.Printf("   ℹ️  Review already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero Splendor Plus Review Bangladesh 2024 - Everyday Family Commuter</h1>
<p class="summary-text">The Hero Splendor Plus is an everyday family commuter offering reliability and practicality. Priced around ৳2,55,000, it delivers 12.2 bhp power, spacious seating, comfortable ride, exceptional 80+ km/l efficiency, and outstanding value for family commuting needs.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Splendor Plus Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Family:</strong> Commuter ready</li>
<li class="highlight-item"><strong>12.2 bhp Power:</strong> Reliable responsive</li>
<li class="highlight-item"><strong>Spacious Seating:</strong> Family comfortable</li>
<li class="highlight-item"><strong>80+ km/l Efficiency:</strong> Exceptional economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Splendor Plus Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Family Friendly:</strong> Spacious seating</li>
<li class="pro-item"><strong>Reliable Power:</strong> 12.2 bhp responsive</li>
<li class="pro-item"><strong>Comfortable Ride:</strong> Smooth handling</li>
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 80+ km/l outstanding</li>
<li class="pro-item"><strong>Hero Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Splendor Plus Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Design:</strong> Family focused</li>
<li class="con-item"><strong>Limited Styling:</strong> Minimal aesthetics</li>
<li class="con-item"><strong>Standard Features:</strong> Practical only</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Splendor Plus Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,55,000 - Family commuter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>12.2 bhp reliable</span></p>
<p class="value-item"><strong>Torque:</strong> <span>12.8 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>138kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>80+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Splendor Plus Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Family practical</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Spacious comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 12.2 bhp reliable</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 80+ km/l exceptional</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Hero proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Splendor Plus?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,55,000, the Hero Splendor Plus is perfect for families seeking reliable 12.2 bhp power, spacious comfortable seating, and exceptional 80+ km/l efficiency in an affordable family commuter.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For family commuters</span></p>
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
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো স্প্লেন্ডর প্লাস রিভিউ বাংলাদেশ 2024 - দৈনন্দিন পারিবারিক যাতায়াত</h1>
<p class="summary-text">হিরো স্প্লেন্ডর প্লাস একটি দৈনন্দিন পারিবারিক যাতায়াত যা নির্ভরযোগ্যতা এবং ব্যবহারিকতা প্রদান করে। ৳2,55,000 টাকায় মূল্যায়িত, এটি 12.2 bhp শক্তি এবং প্রশস্ত আসন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো স্প্লেন্ডর প্লাস মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc পরিবার:</strong> যাতায়াত প্রস্তুত</li>
<li class="highlight-item"><strong>12.2 bhp শক্তি:</strong> নির্ভরযোগ্য প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>প্রশস্ত আসন:</strong> পারিবারিক আরামদায়ক</li>
<li class="highlight-item"><strong>80+ km/l দক্ষতা:</strong> অসাধারণ অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো স্প্লেন্ডর প্লাস সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>পরিবার বান্ধব:</strong> প্রশস্ত আসন</li>
<li class="pro-item"><strong>নির্ভরযোগ্য শক্তি:</strong> 12.2 bhp প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>আরামদায়ক রাইড:</strong> মসৃণ হ্যান্ডলিং</li>
<li class="pro-item"><strong>অসাধারণ দক্ষতা:</strong> 80+ km/l অসাধারণ</li>
<li class="pro-item"><strong>হিরো নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো স্প্লেন্ডর প্লাস অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মৌলিক ডিজাইন:</strong> পারিবারিক ফোকাসড</li>
<li class="con-item"><strong>সীমিত স্টাইলিং:</strong> ন্যূনতম নান্দনিকতা</li>
<li class="con-item"><strong>স্ট্যান্ডার্ড বৈশিষ্ট্য:</strong> ব্যবহারিক শুধু</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো স্প্লেন্ডর প্লাস কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,55,000 টাকায়, হিরো স্প্লেন্ডর প্লাস পরিবারের জন্য নিখুঁত যারা নির্ভরযোগ্য 12.2 bhp শক্তি, প্রশস্ত আরামদায়ক আসন এবং অসাধারণ 80+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারিবারিক যাতায়াতকারীদের জন্য</span></p>
</div>
</section>
</article>`

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating translation: %w", err)
	}

	fmt.Printf("   ✅ Created Hero Splendor Plus\n")
	return nil
}

func (s *HeroSplendorPlusBatch25) GetName() string {
	return "HeroSplendorPlusBatch25"
}
