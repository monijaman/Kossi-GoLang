package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type AppriliaSR150ReviewBatch20 struct {
	BaseSeeder
}

func NewAppriliaSR150ReviewBatch20Seeder() *AppriliaSR150ReviewBatch20 {
	return &AppriliaSR150ReviewBatch20{BaseSeeder: BaseSeeder{name: "Aprilia SR 150 Batch20 Review"}}
}

func (s *AppriliaSR150ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Aprilia SR 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("aprilia sr 150 product not found")
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
<h1 class="review-main-title">Aprilia SR 150 Review Bangladesh 2024 - Sporty Italian Urban Scooter</h1>
<p class="summary-text">The Aprilia SR 150 is a 150cc liquid-cooled sporty Italian urban scooter representing aggressive Italian design heritage. Priced around ৳2,50,000, it delivers responsive performance, distinctive Italian character, practical 45+ km/l efficiency, and youth-oriented sporty styling for dynamic urban commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Aprilia SR 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Liquid-Cooled:</strong> Sporty responsive</li>
<li class="highlight-item"><strong>11 bhp Performance:</strong> Urban sharp</li>
<li class="highlight-item"><strong>Italian Aggression:</strong> Design heritage</li>
<li class="highlight-item"><strong>45+ km/l Efficiency:</strong> Practical</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Aprilia SR 150 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Aggressive Design:</strong> Italian styling character</li>
<li class="pro-item"><strong>Responsive 150cc:</strong> Lively performance</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 45+ km/l practical</li>
<li class="pro-item"><strong>Youth Appeal:</strong> Sporty attitude</li>
<li class="pro-item"><strong>Italian Heritage:</strong> Piaggio group design</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Aprilia SR 150 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Compact Comfort:</strong> Minimal seating</li>
<li class="con-item"><strong>Limited Storage:</strong> Small compartment</li>
<li class="con-item"><strong>Premium Pricing:</strong> ৳2,50,000 sports scooter</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Aprilia SR 150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,50,000 - Sporty Italian premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>11 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>12 nm sharp</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>142kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>45+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Aprilia SR 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Aggressive Italian</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - Responsive 150cc</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 45+ km/l excellent</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Sharp urban agile</li>
<li class="rating-item"><strong>Value:</strong> <span>4.5</span> - ৳2,50,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Aprilia SR 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,50,000, the SR 150 is perfect for youth seeking distinctive Italian character, responsive 150cc performance, excellent 45+ km/l efficiency, and aggressive sporty design for dynamic urban commuting with personality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sporty urban youth</span></p>
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
<h1 class="review-main-title">অ্যাপ্রিলিয়া SR 150 রিভিউ বাংলাদেশ 2024 - খেলাধুলাপূর্ণ ইতালিয়ান শহুরে স্কুটার</h1>
<p class="summary-text">অ্যাপ্রিলিয়া SR 150 একটি 150cc লিকুইড-কুল্ড খেলাধুলাপূর্ণ ইতালিয়ান শহুরে স্কুটার যা আক্রমণাত্মক ইতালিয়ান ডিজাইন ঐতিহ্য প্রতিনিধিত্ব করে। ৳2,50,000 টাকায় মূল্যায়িত, এটি প্রতিক্রিয়াশীল কর্মক্ষমতা এবং স্বতন্ত্র ইতালিয়ান চরিত্র প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">অ্যাপ্রিলিয়া SR 150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc লিকুইড-কুল্ড:</strong> খেলাধুলাপূর্ণ প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>11 bhp কর্মক্ষমতা:</strong> শহুরে তীক্ষ্ণ</li>
<li class="highlight-item"><strong>ইতালিয়ান আক্রমণ:</strong> ডিজাইন ঐতিহ্য</li>
<li class="highlight-item"><strong>45+ km/l দক্ষতা:</strong> ব্যবহারিক</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">অ্যাপ্রিলিয়া SR 150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> ইতালিয়ান স্টাইলিং চরিত্র</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল 150cc:</strong> জীবন্ত কর্মক্ষমতা</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 45+ km/l ব্যবহারিক</li>
<li class="pro-item"><strong>তরুণ আবেদন:</strong> খেলাধুলাপূর্ণ মনোভাব</li>
<li class="pro-item"><strong>ইতালিয়ান ঐতিহ্য:</strong> পিয়াজিও গ্রুপ ডিজাইন</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">অ্যাপ্রিলিয়া SR 150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>কম্প্যাক্ট আরাম:</strong> ন্যূনতম আসন</li>
<li class="con-item"><strong>সীমিত স্টোরেজ:</strong> ছোট বগি</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳2,50,000 খেলাধুলাপূর্ণ</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: অ্যাপ্রিলিয়া SR 150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,50,000 টাকায়, SR 150 তরুণদের জন্য নিখুঁত যারা স্বতন্ত্র ইতালিয়ান চরিত্র, প্রতিক্রিয়াশীল 150cc কর্মক্ষমতা এবং আক্রমণাত্মক খেলাধুলাপূর্ণ ডিজাইন চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- খেলাধুলাপূর্ণ শহুরে তরুণদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Aprilia SR 150\n")
	return nil
}

func (s *AppriliaSR150ReviewBatch20) GetName() string {
	return "AppriliaSR150ReviewBatch20"
}
