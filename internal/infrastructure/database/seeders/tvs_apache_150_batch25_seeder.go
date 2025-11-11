package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type TVSApache150Batch25 struct {
	BaseSeeder
}

func NewTVSApache150Batch25Seeder() *TVSApache150Batch25 {
	return &TVSApache150Batch25{BaseSeeder: BaseSeeder{name: "TVS Apache 150 Batch25 Review"}}
}

func (s *TVSApache150Batch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs apache 150 product not found")
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
<h1 class="review-main-title">TVS Apache 150 Review Bangladesh 2024 - Sport Commuter Excellence</h1>
<p class="summary-text">The TVS Apache 150 is a sport commuter offering thrilling performance and daily practicality. Priced around ৳3,35,000, it delivers 16.2 bhp power, aggressive styling, responsive handling, excellent 65+ km/l efficiency, and outstanding value for sport commuter seekers.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Sport:</strong> Commuter capable</li>
<li class="highlight-item"><strong>16.2 bhp Power:</strong> Thrilling responsive</li>
<li class="highlight-item"><strong>Aggressive Design:</strong> Sporty styled</li>
<li class="highlight-item"><strong>65+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache 150 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Thrilling Power:</strong> 16.2 bhp exciting</li>
<li class="pro-item"><strong>Aggressive Design:</strong> Sporty styling</li>
<li class="pro-item"><strong>Responsive Handling:</strong> Nimble capable</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 65+ km/l great</li>
<li class="pro-item"><strong>TVS Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache 150 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Higher Price:</strong> ৳3,35,000 mid-range</li>
<li class="con-item"><strong>Limited Comfort:</strong> Sport oriented</li>
<li class="con-item"><strong>Harder Suspension:</strong> Sporty setup</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache 150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,35,000 - Sport commuter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>16.2 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</strong> <span>13.3 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>142kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>65+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Aggressive sporty</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Sport configured</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - 16.2 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 65+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - TVS proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,35,000, the TVS Apache 150 is perfect for sport commuter enthusiasts seeking thrilling 16.2 bhp power, aggressive styling, responsive handling, and excellent 65+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sport commuter seekers</span></p>
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
<h1 class="review-main-title">টিভিএস অ্যাপাচ 150 রিভিউ বাংলাদেশ 2024 - খেলাধুলামূলক যাতায়াত শ্রেষ্ঠত্ব</h1>
<p class="summary-text">টিভিএস অ্যাপাচ 150 একটি খেলাধুলামূলক যাতায়াত যা রোমাঞ্চকর পারফরম্যান্স এবং দৈনিক ব্যবহারিকতা প্রদান করে। ৳3,35,000 টাকায় মূল্যায়িত, এটি 16.2 bhp শক্তি এবং আক্রমণাত্মক স্টাইলিং বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস অ্যাপাচ 150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc খেলাধুলা:</strong> যাতায়াত সক্ষম</li>
<li class="highlight-item"><strong>16.2 bhp শক্তি:</strong> রোমাঞ্চকর প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> খেলাধুলামূলক স্টাইল</li>
<li class="highlight-item"><strong>65+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস অ্যাপাচ 150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>রোমাঞ্চকর শক্তি:</strong> 16.2 bhp উত্তেজনাপূর্ণ</li>
<li class="pro-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> খেলাধুলামূলক স্টাইলিং</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল হ্যান্ডলিং:</strong> চটপটে সক্ষম</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 65+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>টিভিএস নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস অ্যাপাচ 150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>উচ্চতর মূল্য:</strong> ৳3,35,000 মধ্য-পরিসীমা</li>
<li class="con-item"><strong>সীমিত আরাম:</strong> খেলাধুলা ভিত্তিক</li>
<li class="con-item"><strong>কঠোর সাসপেনশন:</strong> খেলাধুলামূলক সেটআপ</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: টিভিএস অ্যাপাচ 150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,35,000 টাকায়, টিভিএস অ্যাপাচ 150 খেলাধুলামূলক যাতায়াত উত্সাহীদের জন্য নিখুঁত যারা রোমাঞ্চকর 16.2 bhp শক্তি এবং চমৎকার 65+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- খেলাধুলামূলক যাতায়াত খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created TVS Apache 150\n")
	return nil
}

func (s *TVSApache150Batch25) GetName() string {
	return "TVSApache150Batch25"
}
