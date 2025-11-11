package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HondaActiva6GBatch26 struct {
	BaseSeeder
}

func NewHondaActiva6GBatch26Seeder() *HondaActiva6GBatch26 {
	return &HondaActiva6GBatch26{BaseSeeder: BaseSeeder{name: "Honda Activa 6G Batch26 Review"}}
}

func (s *HondaActiva6GBatch26) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda Activa 6G").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda activa 6g product not found")
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
<h1 class="review-main-title">Honda Activa 6G Review Bangladesh 2024 - Family Scooter King</h1>
<p class="summary-text">The Honda Activa 6G is Bangladesh's family scooter king combining trusted reliability and convenience. Priced around ৳2,75,000, it delivers 7.8 bhp power, spacious storage, comfortable seating, smooth riding, and exceptional 60+ km/l efficiency for families.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda Activa 6G Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc Scooter:</strong> Family king trusted</li>
<li class="highlight-item"><strong>7.8 bhp Power:</strong> Smooth reliable</li>
<li class="highlight-item"><strong>Spacious Storage:</strong> Convenient under-seat</li>
<li class="highlight-item"><strong>60+ km/l Efficiency:</strong> Exceptional economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda Activa 6G Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Trusted Brand:</strong> Honda reliability legacy</li>
<li class="pro-item"><strong>Family Friendly:</strong> Spacious convenient</li>
<li class="pro-item"><strong>Smooth Riding:</strong> 7.8 bhp comfortable</li>
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 60+ km/l economical</li>
<li class="pro-item"><strong>Wide Service:</strong> Extensive network</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda Activa 6G Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Power:</strong> 7.8 bhp modest</li>
<li class="con-item"><strong>Basic Features:</strong> Minimal tech</li>
<li class="con-item"><strong>Traditional Design:</strong> Conservative styling</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda Activa 6G Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,75,000 - Family scooter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>110cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>7.8 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>9 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>109kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>60+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda Activa 6G Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Traditional family</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Spacious convenient</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 7.8 bhp smooth</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 60+ km/l exceptional</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Honda trusted</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda Activa 6G?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,75,000, the Honda Activa 6G is perfect for families seeking trusted reliability, spacious storage, smooth 7.8 bhp riding, and exceptional 60+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For families</span></p>
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
<h1 class="review-main-title">হন্ডা অ্যাক্টিভা 6G রিভিউ বাংলাদেশ 2024 - পারিবারিক স্কুটার কিং</h1>
<p class="summary-text">হন্ডা অ্যাক্টিভা 6G হল বাংলাদেশের পারিবারিক স্কুটার কিং যা বিশ্বস্ত নির্ভরযোগ্যতা এবং সুবিধা একত্রিত করে। ৳2,75,000 টাকায় মূল্যায়িত, এটি 7.8 bhp শক্তি এবং প্রশস্ত সঞ্চয়স্থান বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা অ্যাক্টিভা 6G মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc স্কুটার:</strong> পারিবারিক কিং বিশ্বস্ত</li>
<li class="highlight-item"><strong>7.8 bhp শক্তি:</strong> মসৃণ নির্ভরযোগ্য</li>
<li class="highlight-item"><strong>প্রশস্ত সঞ্চয়স্থান:</strong> সুবিধাজনক আন্ডার-সিট</li>
<li class="highlight-item"><strong>60+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা অ্যাক্টিভা 6G সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>বিশ্বস্ত ব্র্যান্ড:</strong> হন্ডা নির্ভরযোগ্যতা উত্তরাধিকার</li>
<li class="pro-item"><strong>পরিবার বান্ধব:</strong> প্রশস্ত সুবিধাজনক</li>
<li class="pro-item"><strong>মসৃণ চালনা:</strong> 7.8 bhp আরামদায়ক</li>
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 60+ km/l সাশ্রয়ী</li>
<li class="pro-item"><strong>বিস্তৃত সেবা:</strong> বিস্তৃত নেটওয়ার্ক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা অ্যাক্টিভা 6G অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত শক্তি:</strong> 7.8 bhp মধ্যম</li>
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> ন্যূনতম প্রযুক্তি</li>
<li class="con-item"><strong>ঐতিহ্যবাহী ডিজাইন:</strong> রক্ষণশীল স্টাইলিং</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হন্ডা অ্যাক্টিভা 6G কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,75,000 টাকায়, হন্ডা অ্যাক্টিভা 6G পরিবারদের জন্য নিখুঁত যারা বিশ্বস্ত নির্ভরযোগ্যতা এবং ব্যতিক্রমী 60+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পরিবারের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Honda Activa 6G\n")
	return nil
}

func (s *HondaActiva6GBatch26) GetName() string {
	return "HondaActiva6GBatch26"
}
