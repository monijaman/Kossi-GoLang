package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajPlatinaBatch22 struct {
	BaseSeeder
}

func NewBajajPlatinaBatch22Seeder() *BajajPlatinaBatch22 {
	return &BajajPlatinaBatch22{BaseSeeder: BaseSeeder{name: "Bajaj Platina Batch22 Review"}}
}

func (s *BajajPlatinaBatch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Platina").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj platina product not found")
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
<h1 class="review-main-title">Bajaj Platina Review Bangladesh 2024 - India's Most Reliable Economy Commuter</h1>
<p class="summary-text">The Bajaj Platina is a 100cc air-cooled ultimate economy commuter representing India's proven reliability. Priced around ৳1,20,000, it delivers 7.8 bhp power, exceptional 75+ km/l efficiency, durable build, simple maintenance, and outstanding value for practical riders seeking hassle-free commuting daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Platina Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc Air-Cooled:</strong> Proven reliability</li>
<li class="highlight-item"><strong>7.8 bhp Power:</strong> Steady practical</li>
<li class="highlight-item"><strong>75+ km/l Efficiency:</strong> Exceptional economy</li>
<li class="highlight-item"><strong>Durable Build:</strong> Bajaj heritage</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Platina Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 75+ km/l outstanding</li>
<li class="pro-item"><strong>Proven Reliability:</strong> Bajaj reputation</li>
<li class="pro-item"><strong>Simple Maintenance:</strong> Hassle-free service</li>
<li class="pro-item"><strong>Durable Build:</strong> Heritage quality</li>
<li class="pro-item"><strong>Outstanding Value:</strong> ৳1,20,000 supreme</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Platina Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Low Power:</strong> 7.8 bhp basic</li>
<li class="con-item"><strong>Basic Styling:</strong> Minimal aesthetic</li>
<li class="con-item"><strong>No Features:</strong> Functional only</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Platina Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,20,000 - Ultimate value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>100cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>7.8 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8.1 nm practical</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>100kg economy</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>75+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Platina Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Reliability:</strong> <span>4.9</span> - Proven heritage</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.9</span> - 75+ km/l exceptional</li>
<li class="rating-item"><strong>Maintenance:</strong> <span>4.9</span> - Simple hassle-free</li>
<li class="rating-item"><strong>Durability:</strong> <span>4.8</span> - Built to last</li>
<li class="rating-item"><strong>Value:</strong> <span>4.9</span> - ৳1,20,000 supreme</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Platina?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.9/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,20,000, the Bajaj Platina is perfect for practical riders seeking India's most reliable economy commuter, exceptional 75+ km/l efficiency, simple maintenance, durable build, and supreme hassle-free daily value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultimate economy commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.9,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ প্লাটিনা রিভিউ বাংলাদেশ 2024 - ভারতের সবচেয়ে নির্ভরযোগ্য অর্থনীতি কমিউটার</h1>
<p class="summary-text">বাজাজ প্লাটিনা একটি 100cc এয়ার-কুল্ড চূড়ান্ত অর্থনীতি কমিউটার যা ভারতের প্রমাণিত নির্ভরযোগ্যতা প্রতিনিধিত্ব করে। ৳1,20,000 টাকায় মূল্যায়িত, এটি 7.8 bhp শক্তি এবং ব্যতিক্রমী 75+ km/l দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ প্লাটিনা মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc এয়ার-কুল্ড:</strong> প্রমাণিত নির্ভরযোগ্যতা</li>
<li class="highlight-item"><strong>7.8 bhp শক্তি:</strong> স্থিতিশীল ব্যবহারিক</li>
<li class="highlight-item"><strong>75+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
<li class="highlight-item"><strong>টেকসই বিল্ড:</strong> বাজাজ ঐতিহ্য</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ প্লাটিনা সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 75+ km/l অসাধারণ</li>
<li class="pro-item"><strong>প্রমাণিত নির্ভরযোগ্যতা:</strong> বাজাজ খ্যাতি</li>
<li class="pro-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> ঝামেলা-মুক্ত সেবা</li>
<li class="pro-item"><strong>টেকসই বিল্ড:</strong> ঐতিহ্য গুণমান</li>
<li class="pro-item"><strong>অসাধারণ মূল্য:</strong> ৳1,20,000 সর্বোচ্চ</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ প্লাটিনা অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>কম শক্তি:</strong> 7.8 bhp মৌলিক</li>
<li class="con-item"><strong>মৌলিক স্টাইলিং:</strong> ন্যূনতম নান্দনিক</li>
<li class="con-item"><strong>কোন বৈশিষ্ট্য নেই:</strong> কার্যকরী শুধুমাত্র</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ প্লাটিনা কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.9/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,20,000 টাকায়, বাজাজ প্লাটিনা ব্যবহারিক চালকদের জন্য নিখুঁত যারা ভারতের সবচেয়ে নির্ভরযোগ্য অর্থনীতি কমিউটার এবং ব্যতিক্রমী 75+ km/l দক্ষতা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চূড়ান্ত অর্থনীতি কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Platina\n")
	return nil
}

func (s *BajajPlatinaBatch22) GetName() string {
	return "BajajPlatinaBatch22"
}
