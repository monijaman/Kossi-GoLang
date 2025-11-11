package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PiggiBike125Batch21 struct {
	BaseSeeder
}

func NewPiggiBike125Batch21Seeder() *PiggiBike125Batch21 {
	return &PiggiBike125Batch21{BaseSeeder: BaseSeeder{name: "Piggi Bike 125 Batch21 Review"}}
}

func (s *PiggiBike125Batch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Piggi Bike 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("piggi bike 125 product not found")
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
<h1 class="review-main-title">Piggi Bike 125 Review Bangladesh 2024 - Stylish Youth Budget Commuter</h1>
<p class="summary-text">The Piggi Bike 125 is a 125cc air-cooled stylish youth budget commuter representing casual urban attitude and accessible performance. Priced around ৳1,65,000, it delivers 11 bhp responsive power, youthful styling, good efficiency, comfortable riding, and practical value for budget-conscious youth commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Piggi Bike 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Air-Cooled:</strong> Youth commuter</li>
<li class="highlight-item"><strong>11 bhp Power:</strong> Responsive spirited</li>
<li class="highlight-item"><strong>Youthful Styling:</strong> Casual urban</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Piggi Bike 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Responsive 125cc:</strong> Spirited performance</li>
<li class="pro-item"><strong>Youthful Design:</strong> Casual urban appeal</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 50+ km/l exceptional</li>
<li class="pro-item"><strong>Comfortable Riding:</strong> Daily commute friendly</li>
<li class="pro-item"><strong>Budget-Friendly:</strong> ৳1,65,000 youth value</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Piggi Bike 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Moderate Power:</strong> 11 bhp practical only</li>
<li class="con-item"><strong>Basic Quality:</strong> Budget segment build</li>
<li class="con-item"><strong>Limited Features:</strong> Entry-level design</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Piggi Bike 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,65,000 - Youth budget value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>11 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>11 nm spirited</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>132kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Piggi Bike 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 11 bhp responsive</li>
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Youthful casual</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 50+ km/l excellent</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Daily commute</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳1,65,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Piggi Bike 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,65,000, the Piggi Bike 125 is perfect for budget-conscious youth seeking youthful styling, responsive 11 bhp power, excellent 50+ km/l efficiency, and practical daily commuting value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget youth commuters</span></p>
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
<h1 class="review-main-title">পিগি বাইক 125 রিভিউ বাংলাদেশ 2024 - স্টাইলিশ যুব বাজেট কমিউটার</h1>
<p class="summary-text">পিগি বাইক 125 একটি 125cc এয়ার-কুল্ড স্টাইলিশ যুব বাজেট কমিউটার যা নৈমিত্তিক শহুরে মনোভাব এবং সহজলভ্য কর্মক্ষমতা প্রতিনিধিত্ব করে। ৳1,65,000 টাকায় মূল্যায়িত, এটি 11 bhp প্রতিক্রিয়াশীল শক্তি এবং যুব স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">পিগি বাইক 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc এয়ার-কুল্ড:</strong> যুব কমিউটার</li>
<li class="highlight-item"><strong>11 bhp শক্তি:</strong> প্রতিক্রিয়াশীল প্রাণবন্ত</li>
<li class="highlight-item"><strong>যুব স্টাইলিং:</strong> নৈমিত্তিক শহুরে</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">পিগি বাইক 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রতিক্রিয়াশীল 125cc:</strong> প্রাণবন্ত কর্মক্ষমতা</li>
<li class="pro-item"><strong>যুব ডিজাইন:</strong> নৈমিত্তিক শহুরে আবেদন</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 50+ km/l ব্যতিক্রমী</li>
<li class="pro-item"><strong>আরামদায়ক যাত্রা:</strong> দৈনিক যাতায়াত বান্ধব</li>
<li class="pro-item"><strong>বাজেট-বান্ধব:</strong> ৳1,65,000 যুব মূল্য</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">পিগি বাইক 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মধ্যম শক্তি:</strong> 11 bhp মাত্র ব্যবহারিক</li>
<li class="con-item"><strong>মৌলিক গুণমান:</strong> বাজেট সেগমেন্ট বিল্ড</li>
<li class="con-item"><strong>সীমিত বৈশিষ্ট্য:</strong> এন্ট্রি-লেভেল ডিজাইন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: পিগি বাইক 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,65,000 টাকায়, পিগি বাইক 125 বাজেট-সচেতন তরুণদের জন্য নিখুঁত যারা যুব স্টাইলিং, প্রতিক্রিয়াশীল 11 bhp শক্তি এবং চমৎকার 50+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট যুব কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Piggi Bike 125\n")
	return nil
}

func (s *PiggiBike125Batch21) GetName() string {
	return "PiggiBike125Batch21"
}
