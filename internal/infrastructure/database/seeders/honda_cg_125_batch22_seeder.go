package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HondaCG125Batch22 struct {
	BaseSeeder
}

func NewHondaCG125Batch22Seeder() *HondaCG125Batch22 {
	return &HondaCG125Batch22{BaseSeeder: BaseSeeder{name: "Honda CG 125 Batch22 Review"}}
}

func (s *HondaCG125Batch22) Seed(db *gorm.DB) error {
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
<h1 class="review-main-title">Honda CG 125 Review Bangladesh 2024 - Asia's Best-Selling Legendary Workhorse</h1>
<p class="summary-text">The Honda CG 125 is Asia's legendary 125cc best-selling workhorse representing supreme reliability engineering. Priced around ৳2,85,000, it delivers 12 bhp power, excellent 55+ km/l efficiency, timeless proven design, unmatched durability, and outstanding value for riders worldwide seeking legendary iconic performance daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CG 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Air-Cooled:</strong> Legendary best-selling</li>
<li class="highlight-item"><strong>12 bhp Power:</strong> Reliable proven</li>
<li class="highlight-item"><strong>55+ km/l Efficiency:</strong> Excellent economy</li>
<li class="highlight-item"><strong>Timeless Design:</strong> Iconic workhorse</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CG 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Legendary Reliability:</strong> Best-selling icon</li>
<li class="pro-item"><strong>Unmatched Durability:</strong> Proven workhorse</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 55+ km/l value</li>
<li class="pro-item"><strong>Simple Maintenance:</strong> Global parts availability</li>
<li class="pro-item"><strong>Timeless Appeal:</strong> Iconic universal design</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CG 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Dated Styling:</strong> Retro aesthetic</li>
<li class="con-item"><strong>Basic Comfort:</strong> Minimal features</li>
<li class="con-item"><strong>Higher Price:</strong> ৳2,85,000 premium</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CG 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,85,000 - Premium legendary</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>12 bhp reliable</span></p>
<p class="value-item"><strong>Torque:</strong> <span>12 nm solid</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>125kg proven</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>55+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CG 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Reliability:</strong> <span>4.9</span> - Legendary iconic</li>
<li class="rating-item"><strong>Durability:</strong> <span>4.9</span> - Unmatched workhorse</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 55+ km/l excellent</li>
<li class="rating-item"><strong>Maintenance:</strong> <span>4.9</span> - Simple global parts</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳2,85,000 premium</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CG 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,85,000, the Honda CG 125 is perfect for riders worldwide seeking Asia's best-selling legendary workhorse, unmatched reliability, timeless proven design, excellent 55+ km/l efficiency, and iconic global performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For legendary reliable workhorses</span></p>
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
<h1 class="review-main-title">হোন্ডা সিজি 125 রিভিউ বাংলাদেশ 2024 - এশিয়ার সেরা বিক্রয় করা কিংবদন্তি কর্মঘোড়া</h1>
<p class="summary-text">হোন্ডা সিজি 125 এশিয়ার কিংবদন্তি 125cc সেরা বিক্রয় করা কর্মঘোড়া যা সর্বোচ্চ নির্ভরযোগ্যতা প্রকৌশল প্রতিনিধিত্ব করে। ৳2,85,000 টাকায় মূল্যায়িত, এটি 12 bhp শক্তি এবং চমৎকার 55+ km/l দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা সিজি 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc এয়ার-কুল্ড:</strong> কিংবদন্তি সেরা বিক্রয়</li>
<li class="highlight-item"><strong>12 bhp শক্তি:</strong> নির্ভরযোগ্য প্রমাণিত</li>
<li class="highlight-item"><strong>55+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
<li class="highlight-item"><strong>নিরবধি ডিজাইন:</strong> আইকনিক কর্মঘোড়া</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা সিজি 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>কিংবদন্তি নির্ভরযোগ্যতা:</strong> সেরা বিক্রয় আইকন</li>
<li class="pro-item"><strong>অতুলনীয় স্থায়িত্ব:</strong> প্রমাণিত কর্মঘোড়া</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 55+ km/l মূল্য</li>
<li class="pro-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> বিশ্বব্যাপী যন্ত্রাংশ উপলব্ধতা</li>
<li class="pro-item"><strong>নিরবধি আবেদন:</strong> আইকনিক সার্বজনীন ডিজাইন</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা সিজি 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>পুরানো স্টাইলিং:</strong> রেট্রো নান্দনিক</li>
<li class="con-item"><strong>মৌলিক আরাম:</strong> ন্যূনতম বৈশিষ্ট্য</li>
<li class="con-item"><strong>উচ্চতর মূল্য:</strong> ৳2,85,000 প্রিমিয়াম</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা সিজি 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,85,000 টাকায়, হোন্ডা সিজি 125 বিশ্বব্যাপী চালকদের জন্য নিখুঁত যারা এশিয়ার সেরা বিক্রয় করা কিংবদন্তি কর্মঘোড়া, অতুলনীয় নির্ভরযোগ্যতা এবং চমৎকার 55+ km/l দক্ষতা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- কিংবদন্তি নির্ভরযোগ্য কর্মঘোড়াদের জন্য</span></p>
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

func (s *HondaCG125Batch22) GetName() string {
	return "HondaCG125Batch22"
}
