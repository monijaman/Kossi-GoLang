package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HoneyWellMaxBatch24 struct {
	BaseSeeder
}

func NewHoneyWellMaxBatch24Seeder() *HoneyWellMaxBatch24 {
	return &HoneyWellMaxBatch24{BaseSeeder: BaseSeeder{name: "HoneyWell Max Batch24 Review"}}
}

func (s *HoneyWellMaxBatch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "HoneyWell Max").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honeywell max product not found")
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
<h1 class="review-main-title">HoneyWell Max Review Bangladesh 2024 - Premium Quality Scooter</h1>
<p class="summary-text">The HoneyWell Max is a premium 125cc scooter emphasizing quality and comfort. Priced around ৳5,35,000, it delivers 9.2 bhp power, premium build quality, comfortable seating, excellent 60+ km/l efficiency, and outstanding value for riders prioritizing comfort and reliability daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">HoneyWell Max Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Premium:</strong> Quality focused</li>
<li class="highlight-item"><strong>9.2 bhp Power:</strong> Responsive efficient</li>
<li class="highlight-item"><strong>Premium Build:</strong> Quality assured</li>
<li class="highlight-item"><strong>60+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">HoneyWell Max Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Premium Build:</strong> Quality manufacturing</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Long ride ready</li>
<li class="pro-item"><strong>Responsive Power:</strong> 9.2 bhp smooth</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 60+ km/l great</li>
<li class="pro-item"><strong>Reliability Focus:</strong> Build quality emphasis</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">HoneyWell Max Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Pricing:</strong> ৳5,35,000 upscale</li>
<li class="con-item"><strong>Limited Brand:</strong> Emerging market</li>
<li class="con-item"><strong>Service Availability:</strong> Growing network</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">HoneyWell Max Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,35,000 - Premium quality</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc scooter single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>9.2 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>10.3 nm smooth</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>133kg quality</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>60+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">HoneyWell Max Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Premium modern</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Very comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 9.2 bhp smooth</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 60+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Premium build</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy HoneyWell Max?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,35,000, the HoneyWell Max is perfect for riders prioritizing premium quality, comfortable seating, excellent 60+ km/l efficiency, and reliability-focused engineering daily.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For comfort and quality seekers</span></p>
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
<h1 class="review-main-title">হানিওয়েল ম্যাক্স রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম গুণমান স্কুটার</h1>
<p class="summary-text">হানিওয়েল ম্যাক্স একটি প্রিমিয়াম 125cc স্কুটার যা গুণমান এবং আরামকে জোর দেয়। ৳5,35,000 টাকায় মূল্যায়িত, এটি 9.2 bhp শক্তি, প্রিমিয়াম বিল্ড গুণমান এবং আরামদায়ক আসন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হানিওয়েল ম্যাক্স মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc প্রিমিয়াম:</strong> গুণমান ফোকাসড</li>
<li class="highlight-item"><strong>9.2 bhp শক্তি:</strong> প্রতিক্রিয়াশীল দক্ষ</li>
<li class="highlight-item"><strong>প্রিমিয়াম বিল্ড:</strong> গুণমান নিশ্চিত</li>
<li class="highlight-item"><strong>60+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হানিওয়েল ম্যাক্স সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রিমিয়াম বিল্ড:</strong> গুণমান উৎপাদন</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> দীর্ঘ যাত্রা প্রস্তুত</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 9.2 bhp মসৃণ</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 60+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>নির্ভরযোগ্যতা ফোকাস:</strong> বিল্ড গুণমান জোর</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হানিওয়েল ম্যাক্স অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳5,35,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত ব্র্যান্ড:</strong> উদীয়মান বাজার</li>
<li class="con-item"><strong>সেবা উপলব্ধতা:</strong> ক্রমবর্ধমান নেটওয়ার্ক</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হানিওয়েল ম্যাক্স কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,35,000 টাকায়, হানিওয়েল ম্যাক্স চালকদের জন্য নিখুঁত যারা প্রিমিয়াম গুণমান, আরামদায়ক আসন এবং চমৎকার 60+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আরাম এবং গুণমান খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created HoneyWell Max\n")
	return nil
}

func (s *HoneyWellMaxBatch24) GetName() string {
	return "HoneyWellMaxBatch24"
}
