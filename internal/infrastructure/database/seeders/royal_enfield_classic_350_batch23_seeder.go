package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type RoyalEnfieldClassic350Batch23 struct {
	BaseSeeder
}

func NewRoyalEnfieldClassic350Batch23Seeder() *RoyalEnfieldClassic350Batch23 {
	return &RoyalEnfieldClassic350Batch23{BaseSeeder: BaseSeeder{name: "Royal Enfield Classic 350 Batch23 Review"}}
}

func (s *RoyalEnfieldClassic350Batch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Royal Enfield Classic 350").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("royal enfield classic 350 product not found")
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
<h1 class="review-main-title">Royal Enfield Classic 350 Review Bangladesh 2024 - India's Iconic Retro Cruiser Legend</h1>
<p class="summary-text">The Royal Enfield Classic 350 is India's iconic 350cc retro cruiser legend representing timeless heritage character. Priced around ৳6,85,000, it delivers 19.8 bhp steady power, classic retro aesthetics, commanding road presence, good 50+ km/l efficiency, and outstanding value for retro enthusiasts seeking legendary iconic presence daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Royal Enfield Classic 350 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>350cc Single-Cylinder:</strong> Iconic legend</li>
<li class="highlight-item"><strong>19.8 bhp Power:</strong> Steady commanding</li>
<li class="highlight-item"><strong>Classic Retro Design:</strong> Heritage aesthetic</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Royal Enfield Classic 350 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Iconic Legend Status:</strong> Timeless heritage</li>
<li class="pro-item"><strong>Classic Retro Appeal:</strong> Commanding presence</li>
<li class="pro-item"><strong>Steady Power:</strong> 19.8 bhp reliable</li>
<li class="pro-item"><strong>Build Quality:</strong> Premium construction</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 50+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Royal Enfield Classic 350 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>High Price:</strong> ৳6,85,000 premium</li>
<li class="con-item"><strong>Heavier Weight:</strong> 202kg substantial</li>
<li class="con-item"><strong>Limited Modern Tech:</strong> Traditional focused</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Royal Enfield Classic 350 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳6,85,000 - Legend premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>350cc single-cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>19.8 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>28 nm commanding</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>202kg classic</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Royal Enfield Classic 350 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Heritage:</strong> <span>4.9</span> - Iconic legend</li>
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Classic retro</li>
<li class="rating-item"><strong>Build Quality:</strong> <span>4.8</span> - Premium construction</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.8</span> - Proven platform</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳6,85,000 legend</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Royal Enfield Classic 350?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳6,85,000, the Royal Enfield Classic 350 is perfect for retro enthusiasts seeking India's iconic legend, timeless heritage character, commanding road presence, premium build quality, and legendary iconic status.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For iconic retro legend seekers</span></p>
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
<h1 class="review-main-title">রয়্যাল এনফিল্ড ক্লাসিক 350 রিভিউ বাংলাদেশ 2024 - ভারতের আইকনিক রেট্রো ক্রুজার কিংবদন্তি</h1>
<p class="summary-text">রয়্যাল এনফিল্ড ক্লাসিক 350 ভারতের আইকনিক 350cc রেট্রো ক্রুজার কিংবদন্তি যা নিরবধি ঐতিহ্য চরিত্র প্রতিনিধিত্ব করে। ৳6,85,000 টাকায় মূল্যায়িত, এটি 19.8 bhp স্থিতিশীল শক্তি এবং ক্লাসিক রেট্রো নান্দনিকতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রয়্যাল এনফিল্ড ক্লাসিক 350 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>350cc সিঙ্গেল-সিলিন্ডার:</strong> আইকনিক কিংবদন্তি</li>
<li class="highlight-item"><strong>19.8 bhp শক্তি:</strong> স্থিতিশীল কমান্ডিং</li>
<li class="highlight-item"><strong>ক্লাসিক রেট্রো ডিজাইন:</strong> ঐতিহ্য নান্দনিকতা</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রয়্যাল এনফিল্ড ক্লাসিক 350 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আইকনিক কিংবদন্তি স্থিতি:</strong> নিরবধি ঐতিহ্য</li>
<li class="pro-item"><strong>ক্লাসিক রেট্রো আবেদন:</strong> কমান্ডিং উপস্থিতি</li>
<li class="pro-item"><strong>স্থিতিশীল শক্তি:</strong> 19.8 bhp নির্ভরযোগ্য</li>
<li class="pro-item"><strong>বিল্ড গুণমান:</strong> প্রিমিয়াম নির্মাণ</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 50+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রয়্যাল এনফিল্ড ক্লাসিক 350 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>উচ্চ মূল্য:</strong> ৳6,85,000 প্রিমিয়াম</li>
<li class="con-item"><strong>ভারী ওজন:</strong> 202kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>সীমিত আধুনিক প্রযুক্তি:</strong> ঐতিহ্য ফোকাসড</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: রয়্যাল এনফিল্ড ক্লাসিক 350 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳6,85,000 টাকায়, রয়্যাল এনফিল্ড ক্লাসিক 350 রেট্রো উৎসাহীদের জন্য নিখুঁত যারা ভারতের আইকনিক কিংবদন্তি, নিরবধি ঐতিহ্য চরিত্র, কমান্ডিং উপস্থিতি এবং প্রিমিয়াম বিল্ড গুণমান চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আইকনিক রেট্রো কিংবদন্তি খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created Royal Enfield Classic 350\n")
	return nil
}

func (s *RoyalEnfieldClassic350Batch23) GetName() string {
	return "RoyalEnfieldClassic350Batch23"
}
