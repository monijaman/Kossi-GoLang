package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajPulsar150Batch21 struct {
	BaseSeeder
}

func NewBajajPulsar150Batch21Seeder() *BajajPulsar150Batch21 {
	return &BajajPulsar150Batch21{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar 150 Batch21 Review"}}
}

func (s *BajajPulsar150Batch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar 150 product not found")
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
<h1 class="review-main-title">Bajaj Pulsar 150 Review Bangladesh 2024 - Youth Performance Commuter</h1>
<p class="summary-text">The Bajaj Pulsar 150 is a 150cc liquid-cooled youth-oriented performance commuter representing aggressive street attitude. Priced around ৳2,75,000, it delivers 14 bhp responsive power, sharp styling, good efficiency, youthful character, and practical performance for dynamic urban riding and college commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Liquid-Cooled:</strong> Youth performance</li>
<li class="highlight-item"><strong>14 bhp Responsive:</strong> Spirited power</li>
<li class="highlight-item"><strong>Aggressive Styling:</strong> Youth appeal</li>
<li class="highlight-item"><strong>42+ km/l Efficiency:</strong> Practical economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar 150 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Responsive 150cc:</strong> Spirited performance</li>
<li class="pro-item"><strong>Aggressive Design:</strong> Youth lifestyle appeal</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 42+ km/l practical</li>
<li class="pro-item"><strong>Affordable Price:</strong> ৳2,75,000 youth value</li>
<li class="pro-item"><strong>Comfortable Riding:</strong> Daily commute capable</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar 150 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Moderate Power:</strong> 14 bhp adequate only</li>
<li class="con-item"><strong>Build Quality:</strong> Budget segment materials</li>
<li class="con-item"><strong>Noise Levels:</strong> Typical commuter engine</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar 150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,75,000 - Youth performance value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>14 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>14 nm spirited</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>148kg practical</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>42+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 14 bhp responsive</li>
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Aggressive youth</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 42+ km/l good</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Daily commute</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳2,75,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,75,000, the Pulsar 150 is perfect for youth seeking aggressive styling, responsive 14 bhp performance, good 42+ km/l efficiency, and practical daily commuting with personality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For youth commuters</span></p>
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
<h1 class="review-main-title">বাজাজ পালসার 150 রিভিউ বাংলাদেশ 2024 - যুব কর্মক্ষমতা কমিউটার</h1>
<p class="summary-text">বাজাজ পালসার 150 একটি 150cc লিকুইড-কুল্ড যুব-ভিত্তিক কর্মক্ষমতা কমিউটার যা আক্রমণাত্মক রাস্তা মনোভাব প্রতিনিধিত্ব করে। ৳2,75,000 টাকায় মূল্যায়িত, এটি 14 bhp প্রতিক্রিয়াশীল শক্তি এবং আক্রমণাত্মক স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার 150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc লিকুইড-কুল্ড:</strong> যুব কর্মক্ষমতা</li>
<li class="highlight-item"><strong>14 bhp প্রতিক্রিয়াশীল:</strong> প্রাণবন্ত শক্তি</li>
<li class="highlight-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> যুব আবেদন</li>
<li class="highlight-item"><strong>42+ km/l দক্ষতা:</strong> ব্যবহারিক অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার 150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রতিক্রিয়াশীল 150cc:</strong> প্রাণবন্ত কর্মক্ষমতা</li>
<li class="pro-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> যুব জীবনযাত্রা আবেদন</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 42+ km/l ব্যবহারিক</li>
<li class="pro-item"><strong>সাশ্রয়ী মূল্য:</strong> ৳2,75,000 যুব মূল্য</li>
<li class="pro-item"><strong>আরামদায়ক যাত্রা:</strong> দৈনন্দিন যাতায়াত সক্ষম</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার 150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মধ্যম শক্তি:</strong> 14 bhp মাত্র পর্যাপ্ত</li>
<li class="con-item"><strong>বিল্ড গুণমান:</strong> বাজেট সেগমেন্ট উপকরণ</li>
<li class="con-item"><strong>শব্দ স্তর:</strong> সাধারণ কমিউটার ইঞ্জিন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ পালসার 150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,75,000 টাকায়, পালসার 150 তরুণদের জন্য নিখুঁত যারা আক্রমণাত্মক স্টাইলিং, প্রতিক্রিয়াশীল 14 bhp কর্মক্ষমতা এবং ব্যবহারিক দৈনন্দিন যাতায়াত চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- তরুণ কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Pulsar 150\n")
	return nil
}

func (s *BajajPulsar150Batch21) GetName() string {
	return "BajajPulsar150Batch21"
}
