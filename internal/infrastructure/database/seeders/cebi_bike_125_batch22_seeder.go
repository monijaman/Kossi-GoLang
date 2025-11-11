package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type CebiBike125Batch22 struct {
	BaseSeeder
}

func NewCebiBike125Batch22Seeder() *CebiBike125Batch22 {
	return &CebiBike125Batch22{BaseSeeder: BaseSeeder{name: "Cebi Bike 125 Batch22 Review"}}
}

func (s *CebiBike125Batch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Cebi Bike 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("cebi bike 125 product not found")
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
<h1 class="review-main-title">Cebi Bike 125 Review Bangladesh 2024 - Affordable Youth Performance Commuter</h1>
<p class="summary-text">The Cebi Bike 125 is a 125cc air-cooled affordable youth performance commuter representing practical street character. Priced around ৳1,55,000, it delivers 11 bhp responsive power, sporty styling, good handling, excellent 50+ km/l efficiency, and practical value for budget-conscious youth riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Cebi Bike 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Air-Cooled:</strong> Youth performance</li>
<li class="highlight-item"><strong>11 bhp Power:</strong> Responsive spirited</li>
<li class="highlight-item"><strong>Sporty Design:</strong> Street character</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Cebi Bike 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Responsive 125cc:</strong> Spirited performance</li>
<li class="pro-item"><strong>Sporty Styling:</strong> Street attitude appeal</li>
<li class="pro-item"><strong>Good Handling:</strong> Responsive controls</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 50+ km/l outstanding</li>
<li class="pro-item"><strong>Affordable Value:</strong> ৳1,55,000 youth</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Cebi Bike 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Moderate Power:</strong> 11 bhp adequate</li>
<li class="con-item"><strong>Build Quality:</strong> Budget segment materials</li>
<li class="con-item"><strong>Limited Comfort:</strong> Sporty seating focus</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Cebi Bike 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,55,000 - Youth value performance</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>11 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>11 nm spirited</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>130kg sporty</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Cebi Bike 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 11 bhp responsive</li>
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Sporty street</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 50+ km/l excellent</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.6</span> - Good responsive</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳1,55,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Cebi Bike 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,55,000, the Cebi Bike 125 is perfect for budget-conscious youth seeking sporty street character, responsive 11 bhp power, excellent 50+ km/l efficiency, and practical daily commuting value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget youth performers</span></p>
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
<h1 class="review-main-title">সেবি বাইক 125 রিভিউ বাংলাদেশ 2024 - সাশ্রয়ী যুব কর্মক্ষমতা কমিউটার</h1>
<p class="summary-text">সেবি বাইক 125 একটি 125cc এয়ার-কুল্ড সাশ্রয়ী যুব কর্মক্ষমতা কমিউটার যা ব্যবহারিক রাস্তা চরিত্র প্রতিনিধিত্ব করে। ৳1,55,000 টাকায় মূল্যায়িত, এটি 11 bhp প্রতিক্রিয়াশীল শক্তি এবং খেলাধুলাপূর্ণ স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সেবি বাইক 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc এয়ার-কুল্ড:</strong> যুব কর্মক্ষমতা</li>
<li class="highlight-item"><strong>11 bhp শক্তি:</strong> প্রতিক্রিয়াশীল প্রাণবন্ত</li>
<li class="highlight-item"><strong>খেলাধুলাপূর্ণ ডিজাইন:</strong> রাস্তা চরিত্র</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সেবি বাইক 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রতিক্রিয়াশীল 125cc:</strong> প্রাণবন্ত কর্মক্ষমতা</li>
<li class="pro-item"><strong>খেলাধুলাপূর্ণ স্টাইলিং:</strong> রাস্তা মনোভাব আবেদন</li>
<li class="pro-item"><strong>ভালো হ্যান্ডলিং:</strong> প্রতিক্রিয়াশীল নিয়ন্ত্রণ</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 50+ km/l অসাধারণ</li>
<li class="pro-item"><strong>সাশ্রয়ী মূল্য:</strong> ৳1,55,000 তরুণ</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সেবি বাইক 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মধ্যম শক্তি:</strong> 11 bhp পর্যাপ্ত</li>
<li class="con-item"><strong>বিল্ড গুণমান:</strong> বাজেট সেগমেন্ট উপকরণ</li>
<li class="con-item"><strong>সীমিত আরাম:</strong> খেলাধুলাপূর্ণ আসন ফোকাস</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সেবি বাইক 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,55,000 টাকায়, সেবি বাইক 125 বাজেট-সচেতন তরুণদের জন্য নিখুঁত যারা খেলাধুলাপূর্ণ রাস্তা চরিত্র, প্রতিক্রিয়াশীল 11 bhp শক্তি এবং চমৎকার 50+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট যুব পারফরমারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Cebi Bike 125\n")
	return nil
}

func (s *CebiBike125Batch22) GetName() string {
	return "CebiBike125Batch22"
}
