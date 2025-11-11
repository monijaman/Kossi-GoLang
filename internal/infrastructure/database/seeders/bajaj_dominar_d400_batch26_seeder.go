package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajDominarD400Batch26 struct {
	BaseSeeder
}

func NewBajajDominarD400Batch26Seeder() *BajajDominarD400Batch26 {
	return &BajajDominarD400Batch26{BaseSeeder: BaseSeeder{name: "Bajaj Dominar D400 Batch26 Review"}}
}

func (s *BajajDominarD400Batch26) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Dominar D400").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj dominar d400 product not found")
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
<h1 class="review-main-title">Bajaj Dominar D400 Review Bangladesh 2024 - Touring Beast Master</h1>
<p class="summary-text">The Bajaj Dominar D400 is a powerful touring beast combining performance and comfort for long rides. Priced around ৳5,45,000, it delivers 40 bhp power, advanced features, comfortable seating, excellent stability, and impressive 35+ km/l highway efficiency.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Dominar D400 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>400cc Touring:</strong> Beast powerful</li>
<li class="highlight-item"><strong>40 bhp Power:</strong> Strong capable</li>
<li class="highlight-item"><strong>Advanced Features:</strong> Tech equipped modern</li>
<li class="highlight-item"><strong>Touring Comfort:</strong> Long ride ready</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Dominar D400 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Powerful Performance:</strong> 40 bhp strong</li>
<li class="pro-item"><strong>Touring Ready:</strong> Comfortable long rides</li>
<li class="pro-item"><strong>Advanced Features:</strong> Modern technology</li>
<li class="pro-item"><strong>Excellent Stability:</strong> Highway confident</li>
<li class="pro-item"><strong>Affordable Touring:</strong> Value pricing</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Dominar D400 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Heavy Weight:</strong> 193kg substantial</li>
<li class="con-item"><strong>City Maneuverability:</strong> Less agile urban</li>
<li class="con-item"><strong>Moderate Efficiency:</strong> 35+ km/l touring</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Dominar D400 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,45,000 - Touring beast</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>400cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>40 bhp powerful</span></p>
<p class="value-item"><strong>Torque:</strong> <span>35 nm strong</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>193kg substantial</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>35+ km/l highway</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Dominar D400 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Muscular touring</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Long ride ready</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - 40 bhp powerful</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.5</span> - 35+ km/l moderate</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Bajaj reliable</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Dominar D400?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,45,000, the Bajaj Dominar D400 is perfect for touring enthusiasts seeking powerful 40 bhp performance, comfortable long-ride capability, advanced features, and excellent highway stability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For touring enthusiasts</span></p>
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
<h1 class="review-main-title">বাজাজ ডমিনার D400 রিভিউ বাংলাদেশ 2024 - ট্যুরিং বিস্ট মাস্টার</h1>
<p class="summary-text">বাজাজ ডমিনার D400 একটি শক্তিশালী ট্যুরিং বিস্ট যা দীর্ঘ যাত্রার জন্য পারফরম্যান্স এবং আরাম একত্রিত করে। ৳5,45,000 টাকায় মূল্যায়িত, এটি 40 bhp শক্তি এবং উন্নত বৈশিষ্ট্য বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ ডমিনার D400 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>400cc ট্যুরিং:</strong> বিস্ট শক্তিশালী</li>
<li class="highlight-item"><strong>40 bhp শক্তি:</strong> শক্তিশালী সক্ষম</li>
<li class="highlight-item"><strong>উন্নত বৈশিষ্ট্য:</strong> প্রযুক্তি সজ্জিত আধুনিক</li>
<li class="highlight-item"><strong>ট্যুরিং আরাম:</strong> দীর্ঘ যাত্রা প্রস্তুত</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ ডমিনার D400 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>শক্তিশালী কর্মক্ষমতা:</strong> 40 bhp শক্তিশালী</li>
<li class="pro-item"><strong>ট্যুরিং প্রস্তুত:</strong> আরামদায়ক দীর্ঘ যাত্রা</li>
<li class="pro-item"><strong>উন্নত বৈশিষ্ট্য:</strong> আধুনিক প্রযুক্তি</li>
<li class="pro-item"><strong>চমৎকার স্থিতিশীলতা:</strong> হাইওয়ে আত্মবিশ্বাসী</li>
<li class="pro-item"><strong>সাশ্রয়ী ট্যুরিং:</strong> মূল্য মূল্য নির্ধারণ</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ ডমিনার D400 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>ভারী ওজন:</strong> 193kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>শহর চালচলনযোগ্যতা:</strong> কম চতুর শহুরে</li>
<li class="con-item"><strong>মধ্যম দক্ষতা:</strong> 35+ km/l ট্যুরিং</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ ডমিনার D400 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,45,000 টাকায়, বাজাজ ডমিনার D400 ট্যুরিং উৎসাহীদের জন্য নিখুঁত যারা শক্তিশালী 40 bhp পারফরম্যান্স এবং চমৎকার হাইওয়ে স্থিতিশীলতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ট্যুরিং উৎসাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Dominar D400\n")
	return nil
}

func (s *BajajDominarD400Batch26) GetName() string {
	return "BajajDominarD400Batch26"
}
