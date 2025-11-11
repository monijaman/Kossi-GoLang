package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SuzukiStorm125Batch23 struct {
	BaseSeeder
}

func NewSuzukiStorm125Batch23Seeder() *SuzukiStorm125Batch23 {
	return &SuzukiStorm125Batch23{BaseSeeder: BaseSeeder{name: "Suzuki Storm 125 Batch23 Review"}}
}

func (s *SuzukiStorm125Batch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Storm 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki storm 125 product not found")
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
<h1 class="review-main-title">Suzuki Storm 125 Review Bangladesh 2024 - Indonesia's Affordable Youth Commuter</h1>
<p class="summary-text">The Suzuki Storm 125 is Indonesia's affordable 125cc youth commuter representing practical budget performance. Priced around ৳3,25,000, it delivers 11 bhp responsive power, modern minimalist styling, good handling, excellent 54+ km/l efficiency, and outstanding value for budget youth seeking practical affordable character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Storm 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Air-Cooled:</strong> Youth commuter</li>
<li class="highlight-item"><strong>11 bhp Power:</strong> Responsive practical</li>
<li class="highlight-item"><strong>Modern Design:</strong> Minimalist character</li>
<li class="highlight-item"><strong>54+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Storm 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Responsive 11 bhp:</strong> Youth practical</li>
<li class="pro-item"><strong>Modern Styling:</strong> Contemporary appeal</li>
<li class="pro-item"><strong>Good Handling:</strong> Nimble responsive</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 54+ km/l superior</li>
<li class="pro-item"><strong>Affordability:</strong> ৳3,25,000 value</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Storm 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Modest Power:</strong> 11 bhp adequate</li>
<li class="con-item"><strong>Basic Build:</strong> Budget materials</li>
<li class="con-item"><strong>Limited Service:</strong> Suzuki dealers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Storm 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,25,000 - Affordable value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>11 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>11 nm practical</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>130kg modern</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>54+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Storm 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 11 bhp responsive</li>
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Modern minimalist</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 54+ km/l excellent</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.6</span> - Good nimble</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳3,25,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Storm 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,25,000, the Suzuki Storm 125 is perfect for budget youth seeking Indonesia's practical commuter, responsive 11 bhp power, modern minimalist design, excellent 54+ km/l efficiency, and affordable practical value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For affordable youth commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি ঝড় 125 রিভিউ বাংলাদেশ 2024 - ইন্দোনেশিয়ার সাশ্রয়ী যুব কমিউটার</h1>
<p class="summary-text">সুজুকি ঝড় 125 ইন্দোনেশিয়ার সাশ্রয়ী 125cc যুব কমিউটার যা ব্যবহারিক বাজেট পারফরম্যান্স প্রতিনিধিত্ব করে। ৳3,25,000 টাকায় মূল্যায়িত, এটি 11 bhp প্রতিক্রিয়াশীল শক্তি এবং আধুনিক ন্যূনতমবাদী স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি ঝড় 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc এয়ার-কুল্ড:</strong> যুব কমিউটার</li>
<li class="highlight-item"><strong>11 bhp শক্তি:</strong> প্রতিক্রিয়াশীল ব্যবহারিক</li>
<li class="highlight-item"><strong>আধুনিক ডিজাইন:</strong> ন্যূনতমবাদী চরিত্র</li>
<li class="highlight-item"><strong>54+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি ঝড় 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রতিক্রিয়াশীল 11 bhp:</strong> যুব ব্যবহারিক</li>
<li class="pro-item"><strong>আধুনিক স্টাইলিং:</strong> সমসাময়িক আবেদন</li>
<li class="pro-item"><strong>ভালো হ্যান্ডলিং:</strong> চটপটে প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 54+ km/l উচ্চতর</li>
<li class="pro-item"><strong>সামর্থ্য:</strong> ৳3,25,000 মূল্য</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি ঝড় 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মনোরম শক্তি:</strong> 11 bhp পর্যাপ্ত</li>
<li class="con-item"><strong>মৌলিক বিল্ড:</strong> বাজেট উপকরণ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> সুজুকি ডিলার</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি ঝড় 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,25,000 টাকায়, সুজুকি ঝড় 125 বাজেট যুবকদের জন্য নিখুঁত যারা ইন্দোনেশিয়ার ব্যবহারিক কমিউটার, প্রতিক্রিয়াশীল 11 bhp শক্তি এবং চমৎকার 54+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সাশ্রয়ী যুব কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Suzuki Storm 125\n")
	return nil
}

func (s *SuzukiStorm125Batch23) GetName() string {
	return "SuzukiStorm125Batch23"
}
