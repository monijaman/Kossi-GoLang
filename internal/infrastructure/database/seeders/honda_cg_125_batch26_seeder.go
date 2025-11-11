package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HondaCG125Batch26 struct {
	BaseSeeder
}

func NewHondaCG125Batch26Seeder() *HondaCG125Batch26 {
	return &HondaCG125Batch26{BaseSeeder: BaseSeeder{name: "Honda CG 125 Batch26 Review"}}
}

func (s *HondaCG125Batch26) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CG 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda cg 125 product not found")
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
<h1 class="review-main-title">Honda CG 125 Review Bangladesh 2024 - Trusted Family Workhorse</h1>
<p class="summary-text">The Honda CG 125 is a trusted 125cc family workhorse combining reliability and affordability. Priced around ৳2,85,000, it delivers 11 bhp power, spacious design, comfortable seating, exceptional 75+ km/l efficiency, and outstanding value for family commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CG 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Family:</strong> Workhorse reliable</li>
<li class="highlight-item"><strong>11 bhp Power:</strong> Responsive capable</li>
<li class="highlight-item"><strong>Spacious Design:</strong> Family comfortable</li>
<li class="highlight-item"><strong>75+ km/l Efficiency:</strong> Outstanding economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CG 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Trusted Brand:</strong> Honda heritage</li>
<li class="pro-item"><strong>Family Friendly:</strong> Spacious seating</li>
<li class="pro-item"><strong>Responsive Power:</strong> 11 bhp capable</li>
<li class="pro-item"><strong>Outstanding Efficiency:</strong> 75+ km/l exceptional</li>
<li class="pro-item"><strong>Wide Service:</strong> Extensive network</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CG 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Styling:</strong> Traditional design</li>
<li class="con-item"><strong>Basic Features:</strong> Minimal amenities</li>
<li class="con-item"><strong>Moderate Price:</strong> Not budget leader</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CG 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,85,000 - Family commuter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>11 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>11 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>130kg family</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>75+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CG 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Traditional family</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Spacious comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 11 bhp capable</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 75+ km/l exceptional</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Honda proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CG 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,85,000, the Honda CG 125 is perfect for family commuters seeking trusted reliability, spacious seating, responsive 11 bhp power, and exceptional 75+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For family commuters</span></p>
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
<h1 class="review-main-title">হন্ডা CG 125 রিভিউ বাংলাদেশ 2024 - বিশ্বস্ত পারিবারিক ওয়ার্কহর্স</h1>
<p class="summary-text">হন্ডা CG 125 একটি বিশ্বস্ত 125cc পারিবারিক ওয়ার্কহর্স যা নির্ভরযোগ্যতা এবং সামর্থ্য একত্রিত করে। ৳2,85,000 টাকায় মূল্যায়িত, এটি 11 bhp শক্তি এবং প্রশস্ত ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা CG 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc পরিবার:</strong> ওয়ার্কহর্স নির্ভরযোগ্য</li>
<li class="highlight-item"><strong>11 bhp শক্তি:</strong> প্রতিক্রিয়াশীল সক্ষম</li>
<li class="highlight-item"><strong>প্রশস্ত ডিজাইন:</strong> পারিবারিক আরামদায়ক</li>
<li class="highlight-item"><strong>75+ km/l দক্ষতা:</strong> অসাধারণ অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা CG 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>বিশ্বস্ত ব্র্যান্ড:</strong> হন্ডা ঐতিহ্য</li>
<li class="pro-item"><strong>পরিবার বান্ধব:</strong> প্রশস্ত আসন</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 11 bhp সক্ষম</li>
<li class="pro-item"><strong>অসাধারণ দক্ষতা:</strong> 75+ km/l ব্যতিক্রমী</li>
<li class="pro-item"><strong>বিস্তৃত সেবা:</strong> বিস্তৃত নেটওয়ার্ক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা CG 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত স্টাইলিং:</strong> ঐতিহ্যবাহী ডিজাইন</li>
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> ন্যূনতম সুবিধাবিদ</li>
<li class="con-item"><strong>মধ্যম মূল্য:</strong> বাজেট নেতা নয়</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হন্ডা CG 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,85,000 টাকায়, হন্ডা CG 125 পারিবারিক যাতায়াতকারীদের জন্য নিখুঁত যারা বিশ্বস্ত নির্ভরযোগ্যতা এবং অসাধারণ 75+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারিবারিক যাতায়াতকারীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Honda CG 125\n")
	return nil
}

func (s *HondaCG125Batch26) GetName() string {
	return "HondaCG125Batch26"
}
