package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajCT100Batch25 struct {
	BaseSeeder
}

func NewBajajCT100Batch25Seeder() *BajajCT100Batch25 {
	return &BajajCT100Batch25{BaseSeeder: BaseSeeder{name: "Bajaj CT 100 Batch25 Review"}}
}

func (s *BajajCT100Batch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj CT 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj ct 100 product not found")
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
<h1 class="review-main-title">Bajaj CT 100 Review Bangladesh 2024 - Ultra Budget Workhorse</h1>
<p class="summary-text">The Bajaj CT 100 is an ultra-budget workhorse offering unbeatable affordability and durability. Priced around ৳1,45,000, it delivers 7.2 bhp power, robust construction, minimal styling, exceptional 85+ km/l efficiency, and outstanding value for budget-conscious users.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj CT 100 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc Ultra-Budget:</strong> Workhorse reliable</li>
<li class="highlight-item"><strong>7.2 bhp Power:</strong> Practical sufficient</li>
<li class="highlight-item"><strong>Robust Build:</strong> Durable rugged</li>
<li class="highlight-item"><strong>85+ km/l Efficiency:</strong> Exceptional economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj CT 100 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Ultra Budget:</strong> ৳1,45,000 ultra-affordable</li>
<li class="pro-item"><strong>Robust Build:</strong> Durable construction</li>
<li class="pro-item"><strong>Workhorse Ready:</strong> Practical focused</li>
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 85+ km/l outstanding</li>
<li class="pro-item"><strong>Wide Service:</strong> Bajaj network</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj CT 100 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Power:</strong> 7.2 bhp minimal</li>
<li class="con-item"><strong>No Styling:</strong> Purely functional</li>
<li class="con-item"><strong>Basic Comfort:</strong> Minimal amenities</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj CT 100 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,45,000 - Ultra budget</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>100cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>7.2 bhp practical</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8.05 nm sufficient</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>112kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>85+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj CT 100 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.4</span> - Functional minimal</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Basic workhorse</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.5</span> - 7.2 bhp practical</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.9</span> - 85+ km/l exceptional</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Bajaj proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj CT 100?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,45,000, the Bajaj CT 100 is ideal for ultra-budget-conscious users seeking durable construction, exceptional 85+ km/l efficiency, and reliable 7.2 bhp workhorse performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultra-budget users</span></p>
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
<h1 class="review-main-title">বাজাজ CT 100 রিভিউ বাংলাদেশ 2024 - আল্ট্রা বাজেট ওয়ার্কহর্স</h1>
<p class="summary-text">বাজাজ CT 100 একটি আল্ট্রা-বাজেট ওয়ার্কহর্স যা অতুলনীয় সামর্থ্য এবং স্থায়িত্ব প্রদান করে। ৳1,45,000 টাকায় মূল্যায়িত, এটি 7.2 bhp শক্তি এবং শক্তিশালী নির্মাণ বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ CT 100 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc আল্ট্রা-বাজেট:</strong> ওয়ার্কহর্স নির্ভরযোগ্য</li>
<li class="highlight-item"><strong>7.2 bhp শক্তি:</strong> ব্যবহারিক যথেষ্ট</li>
<li class="highlight-item"><strong>শক্তিশালী বিল্ড:</strong> টেকসই রুক্ষ</li>
<li class="highlight-item"><strong>85+ km/l দক্ষতা:</strong> অসাধারণ অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ CT 100 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আল্ট্রা বাজেট:</strong> ৳1,45,000 অতি-সাশ্রয়ী</li>
<li class="pro-item"><strong>শক্তিশালী বিল্ড:</strong> টেকসই নির্মাণ</li>
<li class="pro-item"><strong>ওয়ার্কহর্স প্রস্তুত:</strong> ব্যবহারিক ফোকাসড</li>
<li class="pro-item"><strong>অসাধারণ দক্ষতা:</strong> 85+ km/l অসাধারণ</li>
<li class="pro-item"><strong>বিস্তৃত সেবা:</strong> বাজাজ নেটওয়ার্ক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ CT 100 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত শক্তি:</strong> 7.2 bhp ন্যূনতম</li>
<li class="con-item"><strong>কোন স্টাইলিং:</strong> সম্পূর্ণ কার্যকরী</li>
<li class="con-item"><strong>মৌলিক আরাম:</strong> ন্যূনতম সুবিধাবিদ</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ CT 100 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,45,000 টাকায়, বাজাজ CT 100 আল্ট্রা-বাজেট-সচেতন ব্যবহারকারীদের জন্য আদর্শ যারা টেকসই নির্মাণ এবং অসাধারণ 85+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আল্ট্রা-বাজেট ব্যবহারকারীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj CT 100\n")
	return nil
}

func (s *BajajCT100Batch25) GetName() string {
	return "BajajCT100Batch25"
}
