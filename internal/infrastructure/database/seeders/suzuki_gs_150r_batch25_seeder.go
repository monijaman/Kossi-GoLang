package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SuzukiGS150RBatch25 struct {
	BaseSeeder
}

func NewSuzukiGS150RBatch25Seeder() *SuzukiGS150RBatch25 {
	return &SuzukiGS150RBatch25{BaseSeeder: BaseSeeder{name: "Suzuki GS 150R Batch25 Review"}}
}

func (s *SuzukiGS150RBatch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki GS 150R").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gs 150r product not found")
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
<h1 class="review-main-title">Suzuki GS 150R Review Bangladesh 2024 - Balanced Sport Commuter</h1>
<p class="summary-text">The Suzuki GS 150R is a balanced 150cc sport commuter offering style and substance. Priced around ৳3,85,000, it delivers 16.1 bhp power, modern design, responsive handling, excellent 60+ km/l efficiency, and outstanding value for balanced riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki GS 150R Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Sport:</strong> Balanced commuter</li>
<li class="highlight-item"><strong>16.1 bhp Power:</strong> Responsive thrilling</li>
<li class="highlight-item"><strong>Modern Design:</strong> Contemporary styled</li>
<li class="highlight-item"><strong>60+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki GS 150R Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Modern Design:</strong> Contemporary styling</li>
<li class="pro-item"><strong>Thrilling Power:</strong> 16.1 bhp responsive</li>
<li class="pro-item"><strong>Responsive Handling:</strong> Balanced capable</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 60+ km/l great</li>
<li class="pro-item"><strong>Suzuki Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki GS 150R Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳3,85,000 mid-range</li>
<li class="con-item"><strong>Limited Service:</strong> Suzuki dealers only</li>
<li class="con-item"><strong>Moderate Features:</strong> Standard equipped</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki GS 150R Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,85,000 - Balanced sport</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>16.1 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</strong> <span>13.8 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>141kg balanced</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>60+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki GS 150R Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Modern contemporary</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Balanced capable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 16.1 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 60+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Suzuki proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki GS 150R?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,85,000, the Suzuki GS 150R is perfect for balanced riders seeking modern design, responsive 16.1 bhp power, balanced handling, and excellent 60+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For balanced sport commuter seekers</span></p>
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
<h1 class="review-main-title">সুজুকি GS 150R রিভিউ বাংলাদেশ 2024 - ভারসাম্যপূর্ণ খেলাধুলা যাতায়াত</h1>
<p class="summary-text">সুজুকি GS 150R একটি ভারসাম্যপূর্ণ 150cc খেলাধুলা যাতায়াত যা শৈলী এবং পদার্থ প্রদান করে। ৳3,85,000 টাকায় মূল্যায়িত, এটি 16.1 bhp শক্তি এবং আধুনিক ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি GS 150R মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc খেলা:</strong> ভারসাম্যপূর্ণ যাতায়াত</li>
<li class="highlight-item"><strong>16.1 bhp শক্তি:</strong> প্রতিক্রিয়াশীল রোমাঞ্চকর</li>
<li class="highlight-item"><strong>আধুনিক ডিজাইন:</strong> সমসাময়িক স্টাইল</li>
<li class="highlight-item"><strong>60+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি GS 150R সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আধুনিক ডিজাইন:</strong> সমসাময়িক স্টাইলিং</li>
<li class="pro-item"><strong>রোমাঞ্চকর শক্তি:</strong> 16.1 bhp প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল হ্যান্ডলিং:</strong> ভারসাম্যপূর্ণ সক্ষম</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 60+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>সুজুকি নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি GS 150R অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳3,85,000 মধ্য-পরিসীমা</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> সুজুকি ডিলার শুধু</li>
<li class="con-item"><strong>মধ্যম বৈশিষ্ট্য:</strong> স্ট্যান্ডার্ড সজ্জিত</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি GS 150R কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,85,000 টাকায়, সুজুকি GS 150R ভারসাম্যপূর্ণ চালকদের জন্য নিখুঁত যারা আধুনিক ডিজাইন এবং প্রতিক্রিয়াশীল 16.1 bhp শক্তি চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ভারসাম্যপূর্ণ খেলা যাতায়াত খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created Suzuki GS 150R\n")
	return nil
}

func (s *SuzukiGS150RBatch25) GetName() string {
	return "SuzukiGS150RBatch25"
}
