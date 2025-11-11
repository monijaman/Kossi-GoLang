package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajPulsar150Batch26 struct {
	BaseSeeder
}

func NewBajajPulsar150Batch26Seeder() *BajajPulsar150Batch26 {
	return &BajajPulsar150Batch26{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar 150 Batch26 Review"}}
}

func (s *BajajPulsar150Batch26) Seed(db *gorm.DB) error {
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
<h1 class="review-main-title">Bajaj Pulsar 150 Review Bangladesh 2024 - Youth Performance Leader</h1>
<p class="summary-text">The Bajaj Pulsar 150 is Bangladesh's leading youth sport commuter delivering aggressive styling and responsive performance. Priced around ৳3,25,000, it combines 14 bhp power, stylish design, comfortable riding, agile handling, and thrilling acceleration for young riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Sport:</strong> Youth performance leader</li>
<li class="highlight-item"><strong>14 bhp Power:</strong> Aggressive responsive</li>
<li class="highlight-item"><strong>Stylish Design:</strong> Aggressive attractive</li>
<li class="highlight-item"><strong>Thrilling Acceleration:</strong> Exciting responsive</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar 150 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Aggressive Styling:</strong> Modern youthful</li>
<li class="pro-item"><strong>Responsive Performance:</strong> 14 bhp thrilling</li>
<li class="pro-item"><strong>Excellent Handling:</strong> Agile sporty</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Long ride ready</li>
<li class="pro-item"><strong>Wide Support:</strong> Extensive service network</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar 150 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Higher Fuel Cost:</strong> Sport bike consumption</li>
<li class="con-item"><strong>Limited Storage:</strong> Minimal carrying capacity</li>
<li class="con-item"><strong>Sport Positioning:</strong> Performance focused</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar 150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,25,000 - Youth commuter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>14 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>13 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>138kg agile</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l balanced</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Aggressive stylish</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Sport positioned</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - 14 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 50+ km/l balanced</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Bajaj proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,25,000, the Bajaj Pulsar 150 is perfect for young riders seeking aggressive styling, responsive 14 bhp performance, thrilling acceleration, and excellent handling.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For youth riders</span></p>
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
<h1 class="review-main-title">বাজাজ পালসার 150 রিভিউ বাংলাদেশ 2024 - যুব পারফরম্যান্স নেতা</h1>
<p class="summary-text">বাজাজ পালসার 150 হল বাংলাদেশের অগ্রণী যুব স্পোর্ট কমিউটার যা আক্রমণাত্মক স্টাইলিং এবং প্রতিক্রিয়াশীল কর্মক্ষমতা প্রদান করে। ৳3,25,000 টাকায় মূল্যায়িত, এটি 14 bhp শক্তি এবং স্টাইলিশ ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার 150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc স্পোর্ট:</strong> যুব পারফরম্যান্স নেতা</li>
<li class="highlight-item"><strong>14 bhp শক্তি:</strong> আক্রমণাত্মক প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>স্টাইলিশ ডিজাইন:</strong> আক্রমণাত্মক আকর্ষণীয়</li>
<li class="highlight-item"><strong>রোমাঞ্চকর ত্বরণ:</strong> উত্তেজনাপূর্ণ প্রতিক্রিয়াশীল</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার 150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> আধুনিক যুব</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল কর্মক্ষমতা:</strong> 14 bhp রোমাঞ্চকর</li>
<li class="pro-item"><strong>চমৎকার হ্যান্ডলিং:</strong> চতুর স্পোর্টি</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> দীর্ঘ যাত্রা প্রস্তুত</li>
<li class="pro-item"><strong>বিস্তৃত সহায়তা:</strong> বিস্তৃত সেবা নেটওয়ার্ক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার 150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>উচ্চতর জ্বালানি খরচ:</strong> স্পোর্ট বাইক খরচ</li>
<li class="con-item"><strong>সীমিত সঞ্চয়স্থান:</strong> ন্যূনতম বহন ক্ষমতা</li>
<li class="con-item"><strong>স্পোর্ট অবস্থান:</strong> পারফরম্যান্স দৃষ্টিনিবদ্ধ</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ পালসার 150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,25,000 টাকায়, বাজাজ পালসার 150 যুব আরোহীদের জন্য নিখুঁত যারা আক্রমণাত্মক স্টাইলিং এবং প্রতিক্রিয়াশীল 14 bhp কর্মক্ষমতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- যুব আরোহীদের জন্য</span></p>
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

func (s *BajajPulsar150Batch26) GetName() string {
	return "BajajPulsar150Batch26"
}
