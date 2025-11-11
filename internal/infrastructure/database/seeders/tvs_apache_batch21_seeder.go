package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type TVSApacheBatch21 struct {
	BaseSeeder
}

func NewTVSApacheBatch21Seeder() *TVSApacheBatch21 {
	return &TVSApacheBatch21{BaseSeeder: BaseSeeder{name: "TVS Apache Batch21 Review"}}
}

func (s *TVSApacheBatch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 160").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs apache rtr 160 product not found")
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
<h1 class="review-main-title">TVS Apache RTR 160 Review Bangladesh 2024 - Racing-Inspired Performance Commuter</h1>
<p class="summary-text">The TVS Apache RTR 160 is a 160cc liquid-cooled racing-inspired performance commuter representing track heritage street attitude. Priced around ৳3,15,000, it delivers 15 bhp spirited power, racing design, sport ergonomics, good handling, and practical efficiency for enthusiasts seeking performance-focused commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 160 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>160cc Liquid-Cooled:</strong> Racing heritage</li>
<li class="highlight-item"><strong>15 bhp Power:</strong> Spirited performance</li>
<li class="highlight-item"><strong>Racing Design:</strong> Track-inspired styling</li>
<li class="highlight-item"><strong>43+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 160 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Spirited 160cc:</strong> Racing-inspired power</li>
<li class="pro-item"><strong>Racing Design:</strong> Track-focused aesthetics</li>
<li class="pro-item"><strong>Sport Ergonomics:</strong> Rider-centric positioning</li>
<li class="pro-item"><strong>Good Handling:</strong> Responsive controls</li>
<li class="pro-item"><strong>Practical Value:</strong> ৳3,15,000 performance</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 160 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Racing Focus:</strong> Less commute comfort</li>
<li class="con-item"><strong>Premium Pricing:</strong> ৳3,15,000 segment</li>
<li class="con-item"><strong>Aggressive Stance:</strong> Sports over practicality</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 160 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,15,000 - Racing-inspired premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>160cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>15 bhp spirited</span></p>
<p class="value-item"><strong>Torque:</strong> <span>15 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>155kg race-focused</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>43+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 160 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 15 bhp spirited</li>
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Racing aesthetic</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Sport responsive</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 43+ km/l good</li>
<li class="rating-item"><strong>Value:</strong> <span>4.5</span> - ৳3,15,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 160?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,15,000, the Apache RTR 160 is ideal for enthusiasts seeking racing-inspired design, 15 bhp spirited performance, sport ergonomics, and good handling for performance-focused everyday riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance enthusiasts</span></p>
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
<h1 class="review-main-title">টিভিএস অ্যাপাচ RTR 160 রিভিউ বাংলাদেশ 2024 - রেসিং-অনুপ্রাণিত কর্মক্ষমতা কমিউটার</h1>
<p class="summary-text">টিভিএস অ্যাপাচ RTR 160 একটি 160cc লিকুইড-কুল্ড রেসিং-অনুপ্রাণিত কর্মক্ষমতা কমিউটার যা ট্র্যাক ঐতিহ্য রাস্তা মনোভাব প্রতিনিধিত্ব করে। ৳3,15,000 টাকায় মূল্যায়িত, এটি 15 bhp প্রাণবন্ত শক্তি এবং রেসিং ডিজাইন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস অ্যাপাচ RTR 160 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>160cc লিকুইড-কুল্ড:</strong> রেসিং ঐতিহ্য</li>
<li class="highlight-item"><strong>15 bhp শক্তি:</strong> প্রাণবন্ত কর্মক্ষমতা</li>
<li class="highlight-item"><strong>রেসিং ডিজাইন:</strong> ট্র্যাক-অনুপ্রাণিত স্টাইলিং</li>
<li class="highlight-item"><strong>43+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস অ্যাপাচ RTR 160 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রাণবন্ত 160cc:</strong> রেসিং-অনুপ্রাণিত শক্তি</li>
<li class="pro-item"><strong>রেসিং ডিজাইন:</strong> ট্র্যাক-ফোকাসড নান্দনিক</li>
<li class="pro-item"><strong>খেলাধুলা এরগনমিক্স:</strong> রাইডার-কেন্দ্রিক অবস্থান</li>
<li class="pro-item"><strong>ভালো হ্যান্ডলিং:</strong> প্রতিক্রিয়াশীল নিয়ন্ত্রণ</li>
<li class="pro-item"><strong>ব্যবহারিক মূল্য:</strong> ৳3,15,000 কর্মক্ষমতা</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস অ্যাপাচ RTR 160 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>রেসিং ফোকাস:</strong> কম যাতায়াত আরাম</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳3,15,000 সেগমেন্ট</li>
<li class="con-item"><strong>আক্রমণাত্মক মনোভাব:</strong> ক্রীড়া অতি ব্যবহারিকতা</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: টিভিএস অ্যাপাচ RTR 160 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,15,000 টাকায়, অ্যাপাচ RTR 160 উত্সাহীদের জন্য আদর্শ যারা রেসিং-অনুপ্রাণিত ডিজাইন, 15 bhp প্রাণবন্ত কর্মক্ষমতা এবং খেলাধুলা এরগনমিক্স চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- কর্মক্ষমতা উত্সাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created TVS Apache RTR 160\n")
	return nil
}

func (s *TVSApacheBatch21) GetName() string {
	return "TVSApacheBatch21"
}
