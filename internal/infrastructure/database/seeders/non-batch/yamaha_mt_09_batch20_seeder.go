package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaMT09ReviewBatch20 struct {
	BaseSeeder
}

func NewYamahaMT09ReviewBatch20Seeder() *YamahaMT09ReviewBatch20 {
	return &YamahaMT09ReviewBatch20{BaseSeeder: BaseSeeder{name: "Yamaha MT-09 Batch20 Review"}}
}

func (s *YamahaMT09ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha MT-09").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha mt-09 product not found")
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
<h1 class="review-main-title">Yamaha MT-09 Review Bangladesh 2024 - Hyper-Naked Street Performance Sportbike</h1>
<p class="summary-text">The Yamaha MT-09 is an 889cc liquid-cooled hyper-naked street sportbike representing aggressive aggressive street performance. Priced around ৳17,00,000, it delivers 119 bhp power, sharp aggressive handling, advanced ABS braking, modern technology, and track-ready street focus.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha MT-09 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>889cc Liquid-Cooled:</strong> Triple engine aggressive</li>
<li class="highlight-item"><strong>119 bhp Power:</strong> Hyper-naked performance</li>
<li class="highlight-item"><strong>Sharp Handling:</strong> Track-ready street</li>
<li class="highlight-item"><strong>Advanced ABS:</strong> Cornering braking</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha MT-09 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>119 bhp Triple:</strong> Aggressive 889cc engine</li>
<li class="pro-item"><strong>Sharp Handling:</strong> Track-ready street control</li>
<li class="pro-item"><strong>Advanced ABS:</strong> Cornering braking safety</li>
<li class="pro-item"><strong>Modern Technology:</strong> Electronic systems</li>
<li class="pro-item"><strong>Aggressive Styling:</strong> Hyper-naked attitude</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha MT-09 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>High Power:</strong> 119 bhp demanding</li>
<li class="con-item"><strong>Fuel Consumption:</strong> 15-18 km/l aggressive</li>
<li class="con-item"><strong>Premium Price:</strong> ৳17,00,000 performance</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha MT-09 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳17,00,000 - Hyper-naked premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>889cc liquid-cooled triple</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>119 bhp aggressive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>93 nm sharp</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>193kg nimble</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>15-18 km/l aggressive</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha MT-09 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Power:</strong> <span>4.9</span> - 119 bhp triple</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.9</span> - Sharp track-ready</li>
<li class="rating-item"><strong>Braking:</strong> <span>4.9</span> - Advanced ABS</li>
<li class="rating-item"><strong>Technology:</strong> <span>4.7</span> - Modern systems</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳17,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha MT-09?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳17,00,000, the MT-09 is the ultimate hyper-naked for aggressive street riders seeking 119 bhp triple power, sharp track-ready handling, advanced ABS braking, and aggressive modern street performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For aggressive street riders</span></p>
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
<h1 class="review-main-title">ইয়ামাহা MT-09 রিভিউ বাংলাদেশ 2024 - হাইপার-নেকেড রাস্তা কর্মক্ষমতা স্পোর্টবাইক</h1>
<p class="summary-text">ইয়ামাহা MT-09 একটি 889cc লিকুইড-কুল্ড হাইপার-নেকেড রাস্তা স্পোর্টবাইক যা আক্রমণাত্মক রাস্তা কর্মক্ষমতা প্রতিনিধিত্ব করে। ৳17,00,000 টাকায় মূল্যায়িত, এটি 119 bhp শক্তি এবং তীক্ষ্ণ আক্রমণাত্মক হ্যান্ডলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা MT-09 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>889cc লিকুইড-কুল্ড:</strong> ট্রিপল ইঞ্জিন আক্রমণাত্মক</li>
<li class="highlight-item"><strong>119 bhp শক্তি:</strong> হাইপার-নেকেড কর্মক্ষমতা</li>
<li class="highlight-item"><strong>তীক্ষ্ণ হ্যান্ডলিং:</strong> ট্র্যাক-প্রস্তুত রাস্তা</li>
<li class="highlight-item"><strong>উন্নত ABS:</strong> কোণীয় ব্রেকিং</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা MT-09 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>119 bhp ট্রিপল:</strong> আক্রমণাত্মক 889cc ইঞ্জিন</li>
<li class="pro-item"><strong>তীক্ষ্ণ হ্যান্ডলিং:</strong> ট্র্যাক-প্রস্তুত রাস্তা নিয়ন্ত্রণ</li>
<li class="pro-item"><strong>উন্নত ABS:</strong> কোণীয় ব্রেকিং নিরাপত্তা</li>
<li class="pro-item"><strong>আধুনিক প্রযুক্তি:</strong> ইলেকট্রনিক সিস্টেম</li>
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> হাইপার-নেকেড মনোভাব</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা MT-09 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>উচ্চ শক্তি:</strong> 119 bhp দাবিপূর্ণ</li>
<li class="con-item"><strong>জ্বালানি ব্যবহার:</strong> 15-18 km/l আক্রমণাত্মক</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳17,00,000 কর্মক্ষমতা</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা MT-09 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳17,00,000 টাকায়, MT-09 আক্রমণাত্মক রাস্তা চালকদের জন্য চূড়ান্ত হাইপার-নেকেড যারা 119 bhp ট্রিপল শক্তি এবং তীক্ষ্ণ ট্র্যাক-প্রস্তুত হ্যান্ডলিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আক্রমণাত্মক রাস্তা চালকদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Yamaha MT-09\n")
	return nil
}

func (s *YamahaMT09ReviewBatch20) GetName() string {
	return "YamahaMT09ReviewBatch20"
}
