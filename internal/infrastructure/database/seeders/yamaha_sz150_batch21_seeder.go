package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaSZ150Batch21 struct {
	BaseSeeder
}

func NewYamahaSZ150Batch21Seeder() *YamahaSZ150Batch21 {
	return &YamahaSZ150Batch21{BaseSeeder: BaseSeeder{name: "Yamaha SZ 150 Batch21 Review"}}
}

func (s *YamahaSZ150Batch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha SZ 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha sz 150 product not found")
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
<h1 class="review-main-title">Yamaha SZ 150 Review Bangladesh 2024 - Stylish Youth Commuter</h1>
<p class="summary-text">The Yamaha SZ 150 is a 150cc liquid-cooled stylish youth commuter representing aggressive street attitude and contemporary design. Priced around ৳2,85,000, it delivers 14 bhp spirited power, aggressive styling, good agility, practical efficiency, and youth-oriented character for dynamic city riding.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha SZ 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Liquid-Cooled:</strong> Stylish youth</li>
<li class="highlight-item"><strong>14 bhp Power:</strong> Spirited performance</li>
<li class="highlight-item"><strong>Aggressive Design:</strong> Street attitude</li>
<li class="highlight-item"><strong>43+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha SZ 150 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Spirited 150cc:</strong> Lively performance</li>
<li class="pro-item"><strong>Aggressive Styling:</strong> Street personality</li>
<li class="pro-item"><strong>Good Agility:</strong> Responsive handling</li>
<li class="pro-item"><strong>Youth Appeal:</strong> Contemporary design</li>
<li class="pro-item"><strong>Affordable Value:</strong> ৳2,85,000 young riders</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha SZ 150 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Moderate Power:</strong> 14 bhp adequate</li>
<li class="con-item"><strong>Build Finish:</strong> Entry-level quality</li>
<li class="con-item"><strong>Seat Comfort:</strong> Brief seating design</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha SZ 150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,85,000 - Youth street value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>14 bhp spirited</span></p>
<p class="value-item"><strong>Torque:</strong> <span>13 nm lively</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>145kg agile</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>43+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha SZ 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 14 bhp spirited</li>
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Aggressive style</li>
<li class="rating-item"><strong>Agility:</strong> <span>4.7</span> - Good responsive</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 43+ km/l good</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳2,85,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha SZ 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,85,000, the SZ 150 is perfect for youth seeking aggressive street style, spirited 14 bhp power, good agility, and youthful character for dynamic city commuting with personality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For stylish youth riders</span></p>
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
<h1 class="review-main-title">ইয়ামাহা SZ 150 রিভিউ বাংলাদেশ 2024 - স্টাইলিশ যুব কমিউটার</h1>
<p class="summary-text">ইয়ামাহা SZ 150 একটি 150cc লিকুইড-কুল্ড স্টাইলিশ যুব কমিউটার যা আক্রমণাত্মক রাস্তা মনোভাব এবং সমসাময়িক ডিজাইন প্রতিনিধিত্ব করে। ৳2,85,000 টাকায় মূল্যায়িত, এটি 14 bhp প্রাণবন্ত শক্তি এবং আক্রমণাত্মক স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা SZ 150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc লিকুইড-কুল্ড:</strong> স্টাইলিশ যুব</li>
<li class="highlight-item"><strong>14 bhp শক্তি:</strong> প্রাণবন্ত কর্মক্ষমতা</li>
<li class="highlight-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> রাস্তা মনোভাব</li>
<li class="highlight-item"><strong>43+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা SZ 150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রাণবন্ত 150cc:</strong> জীবন্ত কর্মক্ষমতা</li>
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> রাস্তা ব্যক্তিত্ব</li>
<li class="pro-item"><strong>ভালো চপলতা:</strong> প্রতিক্রিয়াশীল হ্যান্ডলিং</li>
<li class="pro-item"><strong>যুব আবেদন:</strong> সমসাময়িক ডিজাইন</li>
<li class="pro-item"><strong>সাশ্রয়ী মূল্য:</strong> ৳2,85,000 তরুণ চালক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা SZ 150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মধ্যম শক্তি:</strong> 14 bhp পর্যাপ্ত</li>
<li class="con-item"><strong>বিল্ড ফিনিশ:</strong> এন্ট্রি-লেভেল মান</li>
<li class="con-item"><strong>আসন আরাম:</strong> সংক্ষিপ্ত আসন ডিজাইন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা SZ 150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,85,000 টাকায়, SZ 150 তরুণদের জন্য নিখুঁত যারা আক্রমণাত্মক রাস্তা স্টাইল, প্রাণবন্ত 14 bhp শক্তি এবং যুব চরিত্র চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্টাইলিশ তরুণ চালকদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Yamaha SZ 150\n")
	return nil
}

func (s *YamahaSZ150Batch21) GetName() string {
	return "YamahaSZ150Batch21"
}
