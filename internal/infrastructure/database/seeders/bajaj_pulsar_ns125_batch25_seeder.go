package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajPulsarNS125Batch25 struct {
	BaseSeeder
}

func NewBajajPulsarNS125Batch25Seeder() *BajajPulsarNS125Batch25 {
	return &BajajPulsarNS125Batch25{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar NS125 Batch25 Review"}}
}

func (s *BajajPulsarNS125Batch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar NS125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar ns125 product not found")
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
<h1 class="review-main-title">Bajaj Pulsar NS125 Review Bangladesh 2024 - Performance Youth Commuter</h1>
<p class="summary-text">The Bajaj Pulsar NS125 is a performance-oriented 125cc youth commuter delivering thrilling daily rides. Priced around ৳2,95,000, it features 12 bhp power, sporty aggressive design, responsive handling, excellent 68+ km/l efficiency, and outstanding value for youth performance seekers.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar NS125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Youth:</strong> Performance commuter</li>
<li class="highlight-item"><strong>12 bhp Power:</strong> Thrilling responsive</li>
<li class="highlight-item"><strong>Sporty Design:</strong> Aggressive styled</li>
<li class="highlight-item"><strong>68+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar NS125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Youth Appeal:</strong> Sporty aggressive</li>
<li class="pro-item"><strong>Thrilling Power:</strong> 12 bhp exciting</li>
<li class="pro-item"><strong>Responsive Handling:</strong> Nimble capable</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 68+ km/l great</li>
<li class="pro-item"><strong>Bajaj Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar NS125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Power:</strong> 12 bhp modest</li>
<li class="con-item"><strong>Basic Comfort:</strong> Youth focused</li>
<li class="con-item"><strong>Harder Suspension:</strong> Sporty setup</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar NS125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,95,000 - Youth performance</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>12 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</strong> <span>11 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>128kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>68+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar NS125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Aggressive sporty</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Youth focused</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 12 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 68+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Bajaj proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar NS125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,95,000, the Bajaj Pulsar NS125 is perfect for youth performance seekers desiring thrilling 12 bhp power, sporty aggressive design, responsive handling, and excellent 68+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For youth performance commuters</span></p>
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
<h1 class="review-main-title">বাজাজ পালসার NS125 রিভিউ বাংলাদেশ 2024 - পারফরম্যান্স যুব যাতায়াত</h1>
<p class="summary-text">বাজাজ পালসার NS125 একটি পারফরম্যান্স-ভিত্তিক 125cc যুব যাতায়াত যা রোমাঞ্চকর দৈনিক রাইড প্রদান করে। ৳2,95,000 টাকায় মূল্যায়িত, এটি 12 bhp শক্তি এবং খেলাধুলামূলক আক্রমণাত্মক ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার NS125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc যুব:</strong> পারফরম্যান্স যাতায়াত</li>
<li class="highlight-item"><strong>12 bhp শক্তি:</strong> রোমাঞ্চকর প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>খেলাধুলামূলক ডিজাইন:</strong> আক্রমণাত্মক স্টাইল</li>
<li class="highlight-item"><strong>68+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার NS125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>যুব আবেদন:</strong> খেলাধুলা আক্রমণাত্মক</li>
<li class="pro-item"><strong>রোমাঞ্চকর শক্তি:</strong> 12 bhp উত্তেজনাপূর্ণ</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল হ্যান্ডলিং:</strong> চটপটে সক্ষম</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 68+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>বাজাজ নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার NS125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত শক্তি:</strong> 12 bhp মামুলি</li>
<li class="con-item"><strong>মৌলিক আরাম:</strong> যুব ফোকাসড</li>
<li class="con-item"><strong>কঠোর সাসপেনশন:</strong> খেলাধুলামূলক সেটআপ</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ পালসার NS125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,95,000 টাকায়, বাজাজ পালসার NS125 যুব পারফরম্যান্স খোঁজা চালকদের জন্য নিখুঁত যারা রোমাঞ্চকর 12 bhp শক্তি এবং খেলাধুলামূলক আক্রমণাত্মক ডিজাইন চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- যুব পারফরম্যান্স যাতায়াতকারীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Pulsar NS125\n")
	return nil
}

func (s *BajajPulsarNS125Batch25) GetName() string {
	return "BajajPulsarNS125Batch25"
}
