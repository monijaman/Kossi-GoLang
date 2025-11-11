package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HeroHF100Batch22 struct {
	BaseSeeder
}

func NewHeroHF100Batch22Seeder() *HeroHF100Batch22 {
	return &HeroHF100Batch22{BaseSeeder: BaseSeeder{name: "Hero HF 100 Batch22 Review"}}
}

func (s *HeroHF100Batch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero HF 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero hf 100 product not found")
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
<h1 class="review-main-title">Hero HF 100 Review Bangladesh 2024 - India's Ultimate Budget Economy King</h1>
<p class="summary-text">The Hero HF 100 is India's ultimate budget 100cc economy king representing supreme value engineering. Priced around ৳1,35,000, it delivers 8.2 bhp steady power, exceptional 80+ km/l efficiency, simple design, basic reliability, and supreme value for budget-conscious commuters seeking pure affordability daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero HF 100 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc Air-Cooled:</strong> Budget economy</li>
<li class="highlight-item"><strong>8.2 bhp Power:</strong> Steady basic</li>
<li class="highlight-item"><strong>80+ km/l Efficiency:</strong> Exceptional economy</li>
<li class="highlight-item"><strong>Ultimate Value:</strong> ৳1,35,000 supreme</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero HF 100 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 80+ km/l supreme</li>
<li class="pro-item"><strong>Ultimate Affordability:</strong> ৳1,35,000 lowest</li>
<li class="pro-item"><strong>Simple Maintenance:</strong> Budget service</li>
<li class="pro-item"><strong>Hero Reliability:</strong> Proven reputation</li>
<li class="pro-item"><strong>Lightweight:</strong> Easy maneuver</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero HF 100 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Low Power:</strong> 8.2 bhp minimal</li>
<li class="con-item"><strong>Basic Quality:</strong> Budget materials</li>
<li class="con-item"><strong>Minimal Features:</strong> No extras</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero HF 100 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,35,000 - Ultimate value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>100cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>8.2 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8.4 nm basic</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>100kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>80+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero HF 100 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Value:</strong> <span>4.9</span> - ৳1,35,000 supreme</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.9</span> - 80+ km/l exceptional</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.8</span> - Hero proven</li>
<li class="rating-item"><strong>Maintenance:</strong> <span>4.8</span> - Simple budget</li>
<li class="rating-item"><strong>Build:</strong> <span>4.5</span> - Budget basic</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero HF 100?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,35,000, the Hero HF 100 is perfect for budget-conscious commuters seeking India's ultimate economy king, exceptional 80+ km/l efficiency, simple maintenance, and supreme affordable daily value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultimate budget economy commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.8,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো এইচএফ 100 রিভিউ বাংলাদেশ 2024 - ভারতের চূড়ান্ত বাজেট অর্থনীতি রাজা</h1>
<p class="summary-text">হিরো এইচএফ 100 ভারতের চূড়ান্ত বাজেট 100cc অর্থনীতি রাজা যা সর্বোচ্চ মূল্য প্রকৌশল প্রতিনিধিত্ব করে। ৳1,35,000 টাকায় মূল্যায়িত, এটি 8.2 bhp স্থিতিশীল শক্তি এবং ব্যতিক্রমী 80+ km/l দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো এইচএফ 100 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc এয়ার-কুল্ড:</strong> বাজেট অর্থনীতি</li>
<li class="highlight-item"><strong>8.2 bhp শক্তি:</strong> স্থিতিশীল মৌলিক</li>
<li class="highlight-item"><strong>80+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
<li class="highlight-item"><strong>চূড়ান্ত মূল্য:</strong> ৳1,35,000 সর্বোচ্চ</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো এইচএফ 100 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 80+ km/l সর্বোচ্চ</li>
<li class="pro-item"><strong>চূড়ান্ত সামর্থ্য:</strong> ৳1,35,000 সবচেয়ে কম</li>
<li class="pro-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> বাজেট সেবা</li>
<li class="pro-item"><strong>হিরো নির্ভরযোগ্যতা:</strong> প্রমাণিত খ্যাতি</li>
<li class="pro-item"><strong>হালকা ওজন:</strong> সহজ কৌশল</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো এইচএফ 100 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>কম শক্তি:</strong> 8.2 bhp ন্যূনতম</li>
<li class="con-item"><strong>মৌলিক গুণমান:</strong> বাজেট উপকরণ</li>
<li class="con-item"><strong>ন্যূনতম বৈশিষ্ট্য:</strong> কোন অতিরিক্ত নেই</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো এইচএফ 100 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,35,000 টাকায়, হিরো এইচএফ 100 বাজেট-সচেতন কমিউটারদের জন্য নিখুঁত যারা ভারতের চূড়ান্ত অর্থনীতি রাজা, ব্যতিক্রমী 80+ km/l দক্ষতা এবং চূড়ান্ত সাশ্রয়ী মূল্য খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চূড়ান্ত বাজেট অর্থনীতি কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Hero HF 100\n")
	return nil
}

func (s *HeroHF100Batch22) GetName() string {
	return "HeroHF100Batch22"
}
