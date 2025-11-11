package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HeroSplendorPlusBatch21 struct {
	BaseSeeder
}

func NewHeroSplendorPlusBatch21Seeder() *HeroSplendorPlusBatch21 {
	return &HeroSplendorPlusBatch21{BaseSeeder: BaseSeeder{name: "Hero Splendor Plus Batch21 Review"}}
}

func (s *HeroSplendorPlusBatch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Splendor Plus").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero splendor plus product not found")
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
<h1 class="review-main-title">Hero Splendor Plus Review Bangladesh 2024 - India's Best-Selling Commuter Bike</h1>
<p class="summary-text">The Hero Splendor Plus is a 97cc air-cooled budget commuter motorcycle representing India's most trusted and best-selling commuter platform. Priced around ৳1,10,000, it delivers exceptional 60+ km/l fuel efficiency, legendary reliability, simple maintenance, comfortable commuting, and proven durability for everyday practical transportation.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Splendor Plus Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>97cc Air-Cooled:</strong> Budget commuter</li>
<li class="highlight-item"><strong>60+ km/l Efficiency:</strong> Exceptional economy</li>
<li class="highlight-item"><strong>Legendary Reliability:</strong> India's best-seller</li>
<li class="highlight-item"><strong>Simple Maintenance:</strong> Affordable upkeep</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Splendor Plus Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 60+ km/l legendary</li>
<li class="pro-item"><strong>Legendary Reliability:</strong> India's trusted bike</li>
<li class="pro-item"><strong>Affordable Maintenance:</strong> Simple parts availability</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Daily commute friendly</li>
<li class="pro-item"><strong>Budget-Friendly:</strong> ৳1,10,000 value</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Splendor Plus Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Power:</strong> 7.2 bhp minimal</li>
<li class="con-item"><strong>Simple Design:</strong> No frills styling</li>
<li class="con-item"><strong>Manual Only:</strong> No automatic option</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Splendor Plus Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,10,000 - India's budget icon</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>97cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>7.2 bhp practical</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8 nm smooth</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>110kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>60+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Splendor Plus Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Efficiency:</strong> <span>5.0</span> - 60+ km/l legendary</li>
<li class="rating-item"><strong>Reliability:</strong> <span>5.0</span> - India's best-seller</li>
<li class="rating-item"><strong>Maintenance:</strong> <span>4.9</span> - Simple affordable</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Daily commute</li>
<li class="rating-item"><strong>Value:</strong> <span>4.9</span> - ৳1,10,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Splendor Plus?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.9/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,10,000, the Splendor Plus is the ultimate budget commuter for practical everyday transportation, legendary 60+ km/l efficiency, trusted reliability, simple maintenance, and affordable value for daily city commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget commuters</span></p>
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
<h1 class="review-main-title">হিরো স্প্লেন্ডর প্লাস রিভিউ বাংলাদেশ 2024 - ভারতের সেরা-বিক্রি কমিউটার বাইক</h1>
<p class="summary-text">হিরো স্প্লেন্ডর প্লাস একটি 97cc এয়ার-কুল্ড বাজেট কমিউটার মোটরসাইকেল যা ভারতের সবচেয়ে বিশ্বস্ত এবং সেরা-বিক্রি কমিউটার প্ল্যাটফর্ম প্রতিনিধিত্ব করে। ৳1,10,000 টাকায় মূল্যায়িত, এটি ব্যতিক্রমী 60+ km/l জ্বালানি দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো স্প্লেন্ডর প্লাস মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>97cc এয়ার-কুল্ড:</strong> বাজেট কমিউটার</li>
<li class="highlight-item"><strong>60+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
<li class="highlight-item"><strong>কিংবদন্তী নির্ভরযোগ্যতা:</strong> ভারতের সেরা-বিক্রি</li>
<li class="highlight-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> সাশ্রয়ী রক্ষণাবেক্ষণ</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো স্প্লেন্ডর প্লাস সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 60+ km/l কিংবদন্তী</li>
<li class="pro-item"><strong>কিংবদন্তী নির্ভরযোগ্যতা:</strong> ভারতের বিশ্বস্ত বাইক</li>
<li class="pro-item"><strong>সাশ্রয়ী রক্ষণাবেক্ষণ:</strong> সহজ যন্ত্রাংশ প্রাপ্যতা</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> দৈনন্দিন যাতায়াত বান্ধব</li>
<li class="pro-item"><strong>বাজেট-বান্ধব:</strong> ৳1,10,000 মূল্য</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো স্প্লেন্ডর প্লাস অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মৌলিক শক্তি:</strong> 7.2 bhp ন্যূনতম</li>
<li class="con-item"><strong>সহজ ডিজাইন:</strong> কোনো ফ্রিল স্টাইলিং নেই</li>
<li class="con-item"><strong>শুধুমাত্র ম্যানুয়াল:</strong> স্বয়ংক্রিয় বিকল্প নেই</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো স্প্লেন্ডর প্লাস কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.9/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,10,000 টাকায়, স্প্লেন্ডর প্লাস ব্যবহারিক দৈনন্দিন পরিবহনের জন্য চূড়ান্ত বাজেট কমিউটার এবং দৈনন্দিন শহর যাতায়াতের জন্য সাশ্রয়ী মূল্য।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Hero Splendor Plus\n")
	return nil
}

func (s *HeroSplendorPlusBatch21) GetName() string {
	return "HeroSplendorPlusBatch21"
}
