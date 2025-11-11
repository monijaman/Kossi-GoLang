package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type UniversalMotorbike75Batch22 struct {
	BaseSeeder
}

func NewUniversalMotorbike75Batch22Seeder() *UniversalMotorbike75Batch22 {
	return &UniversalMotorbike75Batch22{BaseSeeder: BaseSeeder{name: "Universal Motorbike 75 Batch22 Review"}}
}

func (s *UniversalMotorbike75Batch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Universal Motorbike 75").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("universal motorbike 75 product not found")
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
<h1 class="review-main-title">Universal Motorbike 75 Review Bangladesh 2024 - Bangladesh's Ultra-Budget Micro Commuter</h1>
<p class="summary-text">The Universal Motorbike 75 is Bangladesh's ultra-budget 75cc micro commuter representing extreme affordability. Priced around ৳75,000, it delivers 5.5 bhp steady power, exceptional 80+ km/l efficiency, minimal design, basic reliability, and supreme value for ultra-budget conscious riders seeking absolute lowest cost.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Universal Motorbike 75 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>75cc Air-Cooled:</strong> Micro commuter</li>
<li class="highlight-item"><strong>5.5 bhp Power:</strong> Steady minimal</li>
<li class="highlight-item"><strong>80+ km/l Efficiency:</strong> Exceptional economy</li>
<li class="highlight-item"><strong>৳75,000 Price:</strong> Absolute lowest</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Universal Motorbike 75 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 80+ km/l ultimate</li>
<li class="pro-item"><strong>Absolute Affordability:</strong> ৳75,000 lowest</li>
<li class="pro-item"><strong>Minimal Maintenance:</strong> Ultra budget</li>
<li class="pro-item"><strong>Lightweight:</strong> Easy operate</li>
<li class="pro-item"><strong>Simple Parts:</strong> Local availability</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Universal Motorbike 75 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Very Low Power:</strong> 5.5 bhp minimal</li>
<li class="con-item"><strong>Basic Build:</strong> Minimal materials</li>
<li class="con-item"><strong>Limited Features:</strong> No comfort</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Universal Motorbike 75 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳75,000 - Absolute lowest</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>75cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>5.5 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>6 nm minimal</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>75kg ultra-light</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>80+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Universal Motorbike 75 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Value:</strong> <span>4.9</span> - ৳75,000 absolute lowest</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.9</span> - 80+ km/l exceptional</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.6</span> - Basic steady</li>
<li class="rating-item"><strong>Maintenance:</strong> <span>4.7</span> - Ultra simple</li>
<li class="rating-item"><strong>Design:</strong> <span>4.5</span> - Minimal basic</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Universal Motorbike 75?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳75,000, the Universal Motorbike 75 is perfect for ultra-budget conscious riders seeking Bangladesh's absolute lowest cost micro commuter, exceptional 80+ km/l efficiency, and ultra-simple maintenance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultra-budget micro commuters</span></p>
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
<h1 class="review-main-title">ইউনিভার্সাল মোটরবাইক 75 রিভিউ বাংলাদেশ 2024 - বাংলাদেশের অতি-বাজেট মাইক্রো কমিউটার</h1>
<p class="summary-text">ইউনিভার্সাল মোটরবাইক 75 বাংলাদেশের অতি-বাজেট 75cc মাইক্রো কমিউটার যা চরম সামর্থ্য প্রতিনিধিত্ব করে। ৳75,000 টাকায় মূল্যায়িত, এটি 5.5 bhp স্থিতিশীল শক্তি এবং ব্যতিক্রমী 80+ km/l দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইউনিভার্সাল মোটরবাইক 75 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>75cc এয়ার-কুল্ড:</strong> মাইক্রো কমিউটার</li>
<li class="highlight-item"><strong>5.5 bhp শক্তি:</strong> স্থিতিশীল ন্যূনতম</li>
<li class="highlight-item"><strong>80+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
<li class="highlight-item"><strong>৳75,000 মূল্য:</strong> সম্পূর্ণ নিম্নতম</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইউনিভার্সাল মোটরবাইক 75 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 80+ km/l চূড়ান্ত</li>
<li class="pro-item"><strong>সম্পূর্ণ সামর্থ্য:</strong> ৳75,000 নিম্নতম</li>
<li class="pro-item"><strong>ন্যূনতম রক্ষণাবেক্ষণ:</strong> অতি বাজেট</li>
<li class="pro-item"><strong>হালকা ওজন:</strong> সহজ পরিচালনা</li>
<li class="pro-item"><strong>সহজ যন্ত্রাংশ:</strong> স্থানীয় উপলব্ধতা</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইউনিভার্সাল মোটরবাইক 75 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>অত্যন্ত কম শক্তি:</strong> 5.5 bhp ন্যূনতম</li>
<li class="con-item"><strong>মৌলিক বিল্ড:</strong> ন্যূনতম উপকরণ</li>
<li class="con-item"><strong>সীমিত বৈশিষ্ট্য:</strong> কোন আরাম নেই</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইউনিভার্সাল মোটরবাইক 75 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳75,000 টাকায়, ইউনিভার্সাল মোটরবাইক 75 অতি-বাজেট-সচেতন চালকদের জন্য নিখুঁত যারা বাংলাদেশের সম্পূর্ণ সর্বনিম্ন খরচ মাইক্রো কমিউটার এবং ব্যতিক্রমী 80+ km/l দক্ষতা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অতি-বাজেট মাইক্রো কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Universal Motorbike 75\n")
	return nil
}

func (s *UniversalMotorbike75Batch22) GetName() string {
	return "UniversalMotorbike75Batch22"
}
