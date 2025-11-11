package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HondaX1XBatch25 struct {
	BaseSeeder
}

func NewHondaX1XBatch25Seeder() *HondaX1XBatch25 {
	return &HondaX1XBatch25{BaseSeeder: BaseSeeder{name: "Honda X1X Batch25 Review"}}
}

func (s *HondaX1XBatch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda X1X").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda x1x product not found")
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
<h1 class="review-main-title">Honda X1X Review Bangladesh 2024 - Reliable Urban Commuter</h1>
<p class="summary-text">The Honda X1X is a reliable 110cc urban commuter combining durability and practicality. Priced around ৳2,15,000, it delivers 8 bhp power, robust design, comfortable seating, exceptional 75+ km/l efficiency, and outstanding value for daily urban commuting needs.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda X1X Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc Urban:</strong> Commuter focused</li>
<li class="highlight-item"><strong>8 bhp Power:</strong> Reliable responsive</li>
<li class="highlight-item"><strong>Robust Design:</strong> Durable built</li>
<li class="highlight-item"><strong>75+ km/l Efficiency:</strong> Outstanding economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda X1X Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Reliable Power:</strong> 8 bhp responsive</li>
<li class="pro-item"><strong>Robust Build:</strong> Durable construction</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Urban ready</li>
<li class="pro-item"><strong>Outstanding Efficiency:</strong> 75+ km/l exceptional</li>
<li class="pro-item"><strong>Honda Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda X1X Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Features:</strong> Commuter focused</li>
<li class="con-item"><strong>Limited Power:</strong> 8 bhp modest</li>
<li class="con-item"><strong>Minimal Styling:</strong> Functional design</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda X1X Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,15,000 - Budget commuter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>110cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>8 bhp reliable</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8.8 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>115kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>75+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda X1X Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Functional practical</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Urban comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 8 bhp reliable</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 75+ km/l exceptional</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Honda proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda X1X?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,15,000, the Honda X1X is perfect for daily urban commuters seeking reliable 8 bhp power, robust durability, comfortable seating, and exceptional 75+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For urban commuters</span></p>
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
<h1 class="review-main-title">হন্ডা X1X রিভিউ বাংলাদেশ 2024 - নির্ভরযোগ্য শহুরে যাতায়াত</h1>
<p class="summary-text">হন্ডা X1X একটি নির্ভরযোগ্য 110cc শহুরে যাতায়াত যা স্থায়িত্ব এবং ব্যবহারিকতা একত্রিত করে। ৳2,15,000 টাকায় মূল্যায়িত, এটি 8 bhp শক্তি এবং শক্তিশালী ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা X1X মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc শহুরে:</strong> যাতায়াত ফোকাসড</li>
<li class="highlight-item"><strong>8 bhp শক্তি:</strong> নির্ভরযোগ্য প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>শক্তিশালী ডিজাইন:</strong> টেকসই নির্মিত</li>
<li class="highlight-item"><strong>75+ km/l দক্ষতা:</strong> অসাধারণ অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা X1X সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>নির্ভরযোগ্য শক্তি:</strong> 8 bhp প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>শক্তিশালী বিল্ড:</strong> টেকসই নির্মাণ</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> শহুরে প্রস্তুত</li>
<li class="pro-item"><strong>অসাধারণ দক্ষতা:</strong> 75+ km/l ব্যতিক্রমী</li>
<li class="pro-item"><strong>হন্ডা নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা X1X অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> যাতায়াত ফোকাসড</li>
<li class="con-item"><strong>সীমিত শক্তি:</strong> 8 bhp মামুলি</li>
<li class="con-item"><strong>ন্যূনতম স্টাইলিং:</strong> কার্যকরী ডিজাইন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হন্ডা X1X কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,15,000 টাকায়, হন্ডা X1X দৈনিক শহুরে যাতায়াতকারীদের জন্য নিখুঁত যারা নির্ভরযোগ্য 8 bhp শক্তি এবং অসাধারণ 75+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শহুরে যাতায়াতকারীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Honda X1X\n")
	return nil
}

func (s *HondaX1XBatch25) GetName() string {
	return "HondaX1XBatch25"
}
