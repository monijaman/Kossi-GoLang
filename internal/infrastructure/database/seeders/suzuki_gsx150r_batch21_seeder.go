package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SuzukiGSX150RBatch21 struct {
	BaseSeeder
}

func NewSuzukiGSX150RBatch21Seeder() *SuzukiGSX150RBatch21 {
	return &SuzukiGSX150RBatch21{BaseSeeder: BaseSeeder{name: "Suzuki GSX150R Batch21 Review"}}
}

func (s *SuzukiGSX150RBatch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki GSX150R").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gsx150r product not found")
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
<h1 class="review-main-title">Suzuki GSX150R Review Bangladesh 2024 - Sporty Street Performance Commuter</h1>
<p class="summary-text">The Suzuki GSX150R is a 150cc liquid-cooled sporty street commuter representing Japanese sportbike heritage and performance attitude. Priced around ৳2,95,000, it delivers 14.8 bhp responsive power, sport styling, good handling, practical efficiency, and enthusiast-friendly character for dynamic riding.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki GSX150R Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Liquid-Cooled:</strong> Sporty Japanese</li>
<li class="highlight-item"><strong>14.8 bhp Power:</strong> Responsive performance</li>
<li class="highlight-item"><strong>Sport Styling:</strong> Street character</li>
<li class="highlight-item"><strong>44+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki GSX150R Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Responsive 150cc:</strong> Sporty performance</li>
<li class="pro-item"><strong>Sport Styling:</strong> Aggressive street character</li>
<li class="pro-item"><strong>Good Handling:</strong> Responsive controls</li>
<li class="pro-item"><strong>Practical Value:</strong> ৳2,95,000 sportbike</li>
<li class="pro-item"><strong>Enthusiast Appeal:</strong> Rider-focused bike</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki GSX150R Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Sport Seating:</strong> Less commute comfort</li>
<li class="con-item"><strong>Premium Segment:</strong> ৳2,95,000 pricing</li>
<li class="con-item"><strong>Aggressive Position:</strong> Sports over practicality</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki GSX150R Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,95,000 - Sporty street premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>14.8 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>14 nm lively</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>140kg sport-focused</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>44+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki GSX150R Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 14.8 bhp responsive</li>
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Sport aesthetic</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Good responsive</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 44+ km/l good</li>
<li class="rating-item"><strong>Value:</strong> <span>4.5</span> - ৳2,95,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki GSX150R?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,95,000, the GSX150R is ideal for enthusiasts seeking sporty street design, responsive 14.8 bhp performance, good handling, and enthusiast-friendly character for dynamic everyday riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sporty riders</span></p>
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
<h1 class="review-main-title">সুজুকি GSX150R রিভিউ বাংলাদেশ 2024 - খেলাধুলাপূর্ণ রাস্তা কর্মক্ষমতা কমিউটার</h1>
<p class="summary-text">সুজুকি GSX150R একটি 150cc লিকুইড-কুল্ড খেলাধুলাপূর্ণ রাস্তা কমিউটার যা জাপানি স্পোর্টবাইক ঐতিহ্য এবং কর্মক্ষমতা মনোভাব প্রতিনিধিত্ব করে। ৳2,95,000 টাকায় মূল্যায়িত, এটি 14.8 bhp প্রতিক্রিয়াশীল শক্তি এবং খেলাধুলা স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি GSX150R মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc লিকুইড-কুল্ড:</strong> খেলাধুলাপূর্ণ জাপানি</li>
<li class="highlight-item"><strong>14.8 bhp শক্তি:</strong> প্রতিক্রিয়াশীল কর্মক্ষমতা</li>
<li class="highlight-item"><strong>খেলাধুলা স্টাইলিং:</strong> রাস্তা চরিত্র</li>
<li class="highlight-item"><strong>44+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি GSX150R সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রতিক্রিয়াশীল 150cc:</strong> খেলাধুলাপূর্ণ কর্মক্ষমতা</li>
<li class="pro-item"><strong>খেলাধুলা স্টাইলিং:</strong> আক্রমণাত্মক রাস্তা চরিত্র</li>
<li class="pro-item"><strong>ভালো হ্যান্ডলিং:</strong> প্রতিক্রিয়াশীল নিয়ন্ত্রণ</li>
<li class="pro-item"><strong>ব্যবহারিক মূল্য:</strong> ৳2,95,000 স্পোর্টবাইক</li>
<li class="pro-item"><strong>উত্সাহী আবেদন:</strong> রাইডার-ফোকাসড বাইক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি GSX150R অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>খেলাধুলা আসন:</strong> কম যাতায়াত আরাম</li>
<li class="con-item"><strong>প্রিমিয়াম সেগমেন্ট:</strong> ৳2,95,000 মূল্য নির্ধারণ</li>
<li class="con-item"><strong>আক্রমণাত্মক অবস্থান:</strong> খেলাধুলা অতি ব্যবহারিকতা</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি GSX150R কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,95,000 টাকায়, GSX150R উত্সাহীদের জন্য আদর্শ যারা খেলাধুলাপূর্ণ রাস্তা ডিজাইন, প্রতিক্রিয়াশীল 14.8 bhp কর্মক্ষমতা এবং ভালো হ্যান্ডলিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- খেলাধুলাপূর্ণ চালকদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Suzuki GSX150R\n")
	return nil
}

func (s *SuzukiGSX150RBatch21) GetName() string {
	return "SuzukiGSX150RBatch21"
}
