package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HeroPulseBatch24 struct {
	BaseSeeder
}

func NewHeroPulseBatch24Seeder() *HeroPulseBatch24 {
	return &HeroPulseBatch24{BaseSeeder: BaseSeeder{name: "Hero Pulse Batch24 Review"}}
}

func (s *HeroPulseBatch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Pulse").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero pulse product not found")
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
<h1 class="review-main-title">Hero Pulse Review Bangladesh 2024 - Youth-Oriented Performance Bike</h1>
<p class="summary-text">The Hero Pulse is a youth-oriented 200cc performance bike delivering exciting thrills. Priced around ৳3,75,000, it features 19.3 bhp power, sporty design, aggressive styling, vibrant character, and excellent 70+ km/l efficiency for young riders seeking exciting affordable performance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Pulse Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>200cc Youth:</strong> Performance focused</li>
<li class="highlight-item"><strong>19.3 bhp Power:</strong> Exciting thrills</li>
<li class="highlight-item"><strong>Sporty Design:</strong> Aggressive styled</li>
<li class="highlight-item"><strong>70+ km/l Efficiency:</strong> Outstanding economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Pulse Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Youth Appeal:</strong> Aggressive sporty</li>
<li class="pro-item"><strong>Exciting Power:</strong> 19.3 bhp thrilling</li>
<li class="pro-item"><strong>Affordable Price:</strong> ৳3,75,000 accessible</li>
<li class="pro-item"><strong>Outstanding Efficiency:</strong> 70+ km/l exceptional</li>
<li class="pro-item"><strong>Hero Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Pulse Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Features:</strong> Youth focused</li>
<li class="con-item"><strong>Limited Comfort:</strong> Performance oriented</li>
<li class="con-item"><strong>Harder Suspension:</strong> Sporty setup</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Pulse Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,75,000 - Youth accessible</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>200cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>19.3 bhp exciting</span></p>
<p class="value-item"><strong>Torque:</strong> <span>17.3 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>138kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>70+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Pulse Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Aggressive youth</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Basic sporty</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 19.3 bhp exciting</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 70+ km/l exceptional</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Hero proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Pulse?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,75,000, the Hero Pulse is ideal for young riders seeking exciting 19.3 bhp thrills, aggressive sporty styling, outstanding 70+ km/l efficiency, and affordable youth-oriented performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For young performance enthusiasts</span></p>
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
<h1 class="review-main-title">হিরো পালস রিভিউ বাংলাদেশ 2024 - যুব-ভিত্তিক পারফরম্যান্স বাইক</h1>
<p class="summary-text">হিরো পালস একটি যুব-ভিত্তিক 200cc পারফরম্যান্স বাইক যা রোমাঞ্চকর রোমাঞ্চ প্রদান করে। ৳3,75,000 টাকায় মূল্যায়িত, এটি 19.3 bhp শক্তি এবং খেলাধুলামূলক ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো পালস মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>200cc যুব:</strong> পারফরম্যান্স ফোকাসড</li>
<li class="highlight-item"><strong>19.3 bhp শক্তি:</strong> রোমাঞ্চকর রোমাঞ্চ</li>
<li class="highlight-item"><strong>খেলাধুলামূলক ডিজাইন:</strong> আক্রমণাত্মক স্টাইল</li>
<li class="highlight-item"><strong>70+ km/l দক্ষতা:</strong> অসাধারণ অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো পালস সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>যুব আবেদন:</strong> আক্রমণাত্মক খেলাধুলা</li>
<li class="pro-item"><strong>রোমাঞ্চকর শক্তি:</strong> 19.3 bhp রোমাঞ্চকর</li>
<li class="pro-item"><strong>সাশ্রয়ী মূল্য:</strong> ৳3,75,000 অ্যাক্সেসযোগ্য</li>
<li class="pro-item"><strong>অসাধারণ দক্ষতা:</strong> 70+ km/l ব্যতিক্রমী</li>
<li class="pro-item"><strong>হিরো নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো পালস অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> যুব ফোকাসড</li>
<li class="con-item"><strong>সীমিত আরাম:</strong> পারফরম্যান্স ভিত্তিক</li>
<li class="con-item"><strong>কঠোর সাসপেনশন:</strong> খেলাধুলামূলক সেটআপ</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো পালস কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,75,000 টাকায়, হিরো পালস তরুণ চালকদের জন্য আদর্শ যারা রোমাঞ্চকর 19.3 bhp রোমাঞ্চ, আক্রমণাত্মক খেলাধুলামূলক স্টাইলিং এবং অসাধারণ 70+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- তরুণ পারফরম্যান্স উত্সাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Hero Pulse\n")
	return nil
}

func (s *HeroPulseBatch24) GetName() string {
	return "HeroPulseBatch24"
}
