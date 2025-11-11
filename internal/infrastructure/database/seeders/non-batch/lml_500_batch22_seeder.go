package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type LML500Batch22 struct {
	BaseSeeder
}

func NewLML500Batch22Seeder() *LML500Batch22 {
	return &LML500Batch22{BaseSeeder: BaseSeeder{name: "LML 500 Batch22 Review"}}
}

func (s *LML500Batch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "LML 500").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("lml 500 product not found")
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
<h1 class="review-main-title">LML 500 Review Bangladesh 2024 - India's Classic Retro Adventure Commuter</h1>
<p class="summary-text">The LML 500 is India's classic retro 500cc adventure commuter representing timeless heritage styling. Priced around ৳8,95,000, it delivers 25 bhp steady power, classic retro aesthetics, comfortable positioning, good 45+ km/l efficiency, and outstanding value for classic commuters seeking vintage adventure character daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">LML 500 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>500cc Single-Cylinder:</strong> Classic retro</li>
<li class="highlight-item"><strong>25 bhp Power:</strong> Steady reliable</li>
<li class="highlight-item"><strong>Retro Styling:</strong> Heritage aesthetic</li>
<li class="highlight-item"><strong>45+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">LML 500 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Classic Retro Design:</strong> Heritage character</li>
<li class="pro-item"><strong>Comfortable Position:</strong> All-day riding</li>
<li class="pro-item"><strong>Reliable 25 bhp:</strong> Steady performance</li>
<li class="pro-item"><strong>Accessible Pricing:</strong> Value proposition</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 45+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">LML 500 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Moderate Power:</strong> 25 bhp adequate</li>
<li class="con-item"><strong>Heavier Build:</strong> 400kg substantial</li>
<li class="con-item"><strong>Limited Modernity:</strong> Retro focused</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">LML 500 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳8,95,000 - Heritage value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>500cc single-cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>25 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>42 nm reliable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>400kg comfortable</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>45+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">LML 500 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Classic retro</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - All-day positioning</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.7</span> - Steady 25 bhp</li>
<li class="rating-item"><strong>Heritage:</strong> <span>4.7</span> - Retro character</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳8,95,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy LML 500?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳8,95,000, the LML 500 is perfect for classic commuters seeking India's retro heritage aesthetic, comfortable all-day positioning, steady 25 bhp performance, and vintage adventure character for comfortable daily riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For classic retro adventure commuters</span></p>
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
<h1 class="review-main-title">এলএমএল 500 রিভিউ বাংলাদেশ 2024 - ভারতের ক্লাসিক রেট্রো অ্যাডভেঞ্চার কমিউটার</h1>
<p class="summary-text">এলএমএল 500 ভারতের ক্লাসিক রেট্রো 500cc অ্যাডভেঞ্চার কমিউটার যা নিরবধি ঐতিহ্য স্টাইলিং প্রতিনিধিত্ব করে। ৳8,95,000 টাকায় মূল্যায়িত, এটি 25 bhp স্থিতিশীল শক্তি এবং ক্লাসিক রেট্রো নান্দনিকতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">এলএমএল 500 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>500cc সিঙ্গেল-সিলিন্ডার:</strong> ক্লাসিক রেট্রো</li>
<li class="highlight-item"><strong>25 bhp শক্তি:</strong> স্থিতিশীল নির্ভরযোগ্য</li>
<li class="highlight-item"><strong>রেট্রো স্টাইলিং:</strong> ঐতিহ্য নান্দনিকতা</li>
<li class="highlight-item"><strong>45+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">এলএমএল 500 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ক্লাসিক রেট্রো ডিজাইন:</strong> ঐতিহ্য চরিত্র</li>
<li class="pro-item"><strong>আরামদায়ক অবস্থান:</strong> সারাদিন চালনা</li>
<li class="pro-item"><strong>নির্ভরযোগ্য 25 bhp:</strong> স্থিতিশীল কর্মক্ষমতা</li>
<li class="pro-item"><strong>অ্যাক্সেসযোগ্য মূল্য:</strong> মূল্য প্রস্তাব</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 45+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">এলএমএল 500 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মধ্যম শক্তি:</strong> 25 bhp পর্যাপ্ত</li>
<li class="con-item"><strong>ভারী বিল্ড:</strong> 400kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>সীমিত আধুনিকতা:</strong> রেট্রো ফোকাসড</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: এলএমএল 500 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳8,95,000 টাকায়, এলএমএল 500 ক্লাসিক কমিউটারদের জন্য নিখুঁত যারা ভারতের রেট্রো ঐতিহ্য নান্দনিকতা, আরামদায়ক সারাদিন অবস্থান এবং স্থিতিশীল 25 bhp কর্মক্ষমতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্লাসিক রেট্রো অ্যাডভেঞ্চার কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created LML 500\n")
	return nil
}

func (s *LML500Batch22) GetName() string {
	return "LML500Batch22"
}
