package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajPlatina100Batch25 struct {
	BaseSeeder
}

func NewBajajPlatina100Batch25Seeder() *BajajPlatina100Batch25 {
	return &BajajPlatina100Batch25{BaseSeeder: BaseSeeder{name: "Bajaj Platina 100 Batch25 Review"}}
}

func (s *BajajPlatina100Batch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Platina 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj platina 100 product not found")
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
<h1 class="review-main-title">Bajaj Platina 100 Review Bangladesh 2024 - Budget Commuter Champion</h1>
<p class="summary-text">The Bajaj Platina 100 is the budget commuter champion combining affordability and reliability. Priced around ৳1,65,000, it delivers 7.2 bhp power, sturdy construction, practical design, exceptional 80+ km/l efficiency, and outstanding value for cost-conscious commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Platina 100 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc Budget:</strong> Commuter leader</li>
<li class="highlight-item"><strong>7.2 bhp Power:</strong> Practical responsive</li>
<li class="highlight-item"><strong>Sturdy Build:</strong> Reliable durable</li>
<li class="highlight-item"><strong>80+ km/l Efficiency:</strong> Exceptional economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Platina 100 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Budget Price:</strong> ৳1,65,000 affordable</li>
<li class="pro-item"><strong>Sturdy Build:</strong> Reliable construction</li>
<li class="pro-item"><strong>Practical Design:</strong> Commuter ready</li>
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 80+ km/l outstanding</li>
<li class="pro-item"><strong>Bajaj Service:</strong> Wide network</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Platina 100 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Power:</strong> 7.2 bhp modest</li>
<li class="con-item"><strong>Basic Features:</strong> Minimal styling</li>
<li class="con-item"><strong>Harsh Ride:</strong> Budget suspension</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Platina 100 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,65,000 - Ultra budget</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>100cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>7.2 bhp practical</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8.05 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>115kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>80+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Platina 100 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.5</span> - Practical minimal</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Basic adequate</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 7.2 bhp practical</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.9</span> - 80+ km/l exceptional</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Bajaj proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Platina 100?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,65,000, the Bajaj Platina 100 is ideal for budget-conscious commuters seeking exceptional 80+ km/l efficiency, sturdy build quality, and reliable 7.2 bhp power.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget commuters</span></p>
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
<h1 class="review-main-title">বাজাজ প্ল্যাটিনা 100 রিভিউ বাংলাদেশ 2024 - বাজেট যাতায়াত চ্যাম্পিয়ন</h1>
<p class="summary-text">বাজাজ প্ল্যাটিনা 100 বাজেট যাতায়াত চ্যাম্পিয়ন যা সামর্থ্য এবং নির্ভরযোগ্যতা একত্রিত করে। ৳1,65,000 টাকায় মূল্যায়িত, এটি 7.2 bhp শক্তি এবং শক্তিশালী নির্মাণ বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ প্ল্যাটিনা 100 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc বাজেট:</strong> যাতায়াত নেতা</li>
<li class="highlight-item"><strong>7.2 bhp শক্তি:</strong> ব্যবহারিক প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>শক্তিশালী বিল্ড:</strong> নির্ভরযোগ্য টেকসই</li>
<li class="highlight-item"><strong>80+ km/l দক্ষতা:</strong> অসাধারণ অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ প্ল্যাটিনা 100 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>বাজেট মূল্য:</strong> ৳1,65,000 সাশ্রয়ী</li>
<li class="pro-item"><strong>শক্তিশালী বিল্ড:</strong> নির্ভরযোগ্য নির্মাণ</li>
<li class="pro-item"><strong>ব্যবহারিক ডিজাইন:</strong> যাতায়াত প্রস্তুত</li>
<li class="pro-item"><strong>অসাধারণ দক্ষতা:</strong> 80+ km/l অসাধারণ</li>
<li class="pro-item"><strong>বাজাজ সেবা:</strong> বিস্তৃত নেটওয়ার্ক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ প্ল্যাটিনা 100 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত শক্তি:</strong> 7.2 bhp মামুলি</li>
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> ন্যূনতম স্টাইলিং</li>
<li class="con-item"><strong>কঠোর রাইড:</strong> বাজেট সাসপেনশন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ প্ল্যাটিনা 100 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,65,000 টাকায়, বাজাজ প্ল্যাটিনা 100 বাজেট-সচেতন যাতায়াতকারীদের জন্য আদর্শ যারা অসাধারণ 80+ km/l দক্ষতা এবং শক্তিশালী বিল্ড চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট যাতায়াতকারীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Platina 100\n")
	return nil
}

func (s *BajajPlatina100Batch25) GetName() string {
	return "BajajPlatina100Batch25"
}
