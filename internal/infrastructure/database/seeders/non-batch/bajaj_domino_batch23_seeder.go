package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajDominoBatch23 struct {
	BaseSeeder
}

func NewBajajDominoBatch23Seeder() *BajajDominoBatch23 {
	return &BajajDominoBatch23{BaseSeeder: BaseSeeder{name: "Bajaj Domino Batch23 Review"}}
}

func (s *BajajDominoBatch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Domino").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj domino product not found")
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
<h1 class="review-main-title">Bajaj Domino Review Bangladesh 2024 - India's Commuter-Friendly Automatic Scooter</h1>
<p class="summary-text">The Bajaj Domino is India's commuter-friendly 110cc automatic scooter representing practical daily utility. Priced around ৳3,15,000, it delivers 8.2 bhp steady power, automatic transmission, spacious design, exceptional 62+ km/l efficiency, and outstanding value for practical daily commuters seeking hassle-free utility.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Domino Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc Automatic:</strong> Commuter friendly</li>
<li class="highlight-item"><strong>8.2 bhp Power:</strong> Steady practical</li>
<li class="highlight-item"><strong>Automatic Transmission:</strong> Hassle-free</li>
<li class="highlight-item"><strong>62+ km/l Efficiency:</strong> Exceptional economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Domino Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Automatic Transmission:</strong> Zero clutch effort</li>
<li class="pro-item"><strong>Spacious Design:</strong> Daily commuting</li>
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 62+ km/l supreme</li>
<li class="pro-item"><strong>Affordable Pricing:</strong> ৳3,15,000 value</li>
<li class="pro-item"><strong>Bajaj Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Domino Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Styling:</strong> Conservative look</li>
<li class="con-item"><strong>Low Power:</strong> 8.2 bhp modest</li>
<li class="con-item"><strong>Limited Features:</strong> Functional only</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Domino Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,15,000 - Practical value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>110cc automatic scooter</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>8.2 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>9.5 nm practical</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>115kg spacious</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>62+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Domino Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Practicality:</strong> <span>4.7</span> - Daily commuting</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 62+ km/l exceptional</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.7</span> - Bajaj proven</li>
<li class="rating-item"><strong>Transmission:</strong> <span>4.8</span> - Automatic smooth</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳3,15,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Domino?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,15,000, the Bajaj Domino is perfect for practical daily commuters seeking India's commuter-friendly automatic scooter, spacious design, exceptional 62+ km/l efficiency, and hassle-free practical utility.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For practical daily commuters</span></p>
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
<h1 class="review-main-title">বাজাজ ডমিনো রিভিউ বাংলাদেশ 2024 - ভারতের যাত্রী-বান্ধব অটোমেটিক স্কুটার</h1>
<p class="summary-text">বাজাজ ডমিনো ভারতের যাত্রী-বান্ধব 110cc অটোমেটিক স্কুটার যা ব্যবহারিক দৈনন্দিন উপযোগিতা প্রতিনিধিত্ব করে। ৳3,15,000 টাকায় মূল্যায়িত, এটি 8.2 bhp স্থিতিশীল শক্তি এবং অটোমেটিক ট্রান্সমিশন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ ডমিনো মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc অটোমেটিক:</strong> যাত্রী-বান্ধব</li>
<li class="highlight-item"><strong>8.2 bhp শক্তি:</strong> স্থিতিশীল ব্যবহারিক</li>
<li class="highlight-item"><strong>অটোমেটিক ট্রান্সমিশন:</strong> ঝামেলা-মুক্ত</li>
<li class="highlight-item"><strong>62+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ ডমিনো সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>অটোমেটিক ট্রান্সমিশন:</strong> শূন্য ক্লাচ প্রচেষ্টা</li>
<li class="pro-item"><strong>প্রশস্ত ডিজাইন:</strong> দৈনন্দিন যাতায়াত</li>
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 62+ km/l সর্বোচ্চ</li>
<li class="pro-item"><strong>সাশ্রয়ী মূল্য:</strong> ৳3,15,000 মূল্য</li>
<li class="pro-item"><strong>বাজাজ নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ ডমিনো অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মৌলিক স্টাইলিং:</strong> রক্ষণশীল চেহারা</li>
<li class="con-item"><strong>কম শক্তি:</strong> 8.2 bhp বিনম্র</li>
<li class="con-item"><strong>সীমিত বৈশিষ্ট্য:</strong> কার্যকরী শুধুমাত্র</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ ডমিনো কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,15,000 টাকায়, বাজাজ ডমিনো ব্যবহারিক দৈনন্দিন কমিউটারদের জন্য নিখুঁত যারা ভারতের যাত্রী-বান্ধব অটোমেটিক স্কুটার, প্রশস্ত ডিজাইন এবং ব্যতিক্রমী 62+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্যবহারিক দৈনন্দিন কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Domino\n")
	return nil
}

func (s *BajajDominoBatch23) GetName() string {
	return "BajajDominoBatch23"
}
