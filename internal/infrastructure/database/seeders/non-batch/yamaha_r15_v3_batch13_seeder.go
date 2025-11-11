package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaR15V3Review handles seeding Yamaha R15 V3 product review and translation
type YamahaR15V3Review struct {
	BaseSeeder
}

// NewYamahaR15V3ReviewSeeder creates a new YamahaR15V3Review
func NewYamahaR15V3ReviewSeeder() *YamahaR15V3Review {
	return &YamahaR15V3Review{BaseSeeder: BaseSeeder{name: "Yamaha R15 V3 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaR15V3Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R15 V3").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha r15 v3 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding yamaha r15 v3 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha R15 V3 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha R15 V3 Review Bangladesh 2024 - Premium Sport Bike</h1>
<p class="summary-text">The Yamaha R15 V3 is a powerful 150cc sport bike with liquid cooling, race-inspired design, aggressive styling, and excellent performance. Priced around ৳3,50,000, it delivers thrilling riding experience with sharp handling and dynamic performance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R15 V3 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc Liquid Cooled:</strong> <span class="highlight-value">Powerful performance engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Race-Inspired Design:</strong> <span class="highlight-value">Aggressive sport styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Excellent cornering ability</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS Braking:</strong> <span class="highlight-value">Advanced safety feature</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R15 V3 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">150cc liquid cooled performance</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Design:</strong> <span class="pro-description">Race-inspired styling</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Handling:</strong> <span class="pro-description">Sharp and responsive</span></li>
<li class="pro-item"><strong class="pro-label">Premium Features:</strong> <span class="pro-description">ABS and modern technology</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Feel:</strong> <span class="pro-description">Thrilling riding experience</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R15 V3 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Fuel Cost:</strong> <span class="con-description">Sport bike higher consumption</span></li>
<li class="con-item"><strong class="con-label">Aggressive Position:</strong> <span class="con-description">Forward riding posture</span></li>
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">More expensive option</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R15 V3 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,50,000 - Premium sport bike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳6-7 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">38-42 km/l consumption</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R15 V3 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Power:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- 150cc liquid cooled</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Sharp and responsive</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Aggressive sport design</span></li>
<li class="rating-item"><strong class="rating-label">Safety Features:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- ABS equipped</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Premium features worth</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R15 V3?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,50,000, the Yamaha R15 V3 is ideal for performance-seeking riders who want a thrilling sport bike with excellent handling, aggressive styling, and premium features for serious enthusiasts.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sport bike enthusiasts</span></p>
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
		return fmt.Errorf("error creating yamaha r15 v3 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha R15 V3 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা R15 V3 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম স্পোর্ট বাইক</h1>
<p class="summary-text">ইয়ামাহা R15 V3 একটি শক্তিশালী 150cc স্পোর্ট বাইক যা লিকুইড কুলিং, রেস-অনুপ্রাণিত ডিজাইন, আক্রমণাত্মক স্টাইলিং এবং চমৎকার কর্মক্ষমতা সহ আসে। ৳3,50,000 টাকায় মূল্যায়িত, এটি তীক্ষ্ণ হ্যান্ডলিং এবং গতিশীল কর্মক্ষমতা সহ রোমাঞ্চকর যাত্রার অভিজ্ঞতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা R15 V3 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc লিকুইড কুলড:</strong> <span class="highlight-value">শক্তিশালী কর্মক্ষমতা ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">রেস-অনুপ্রাণিত ডিজাইন:</strong> <span class="highlight-value">আক্রমণাত্মক স্পোর্ট স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="highlight-value">চমৎকার কোণকরণ ক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS ব্রেকিং:</strong> <span class="highlight-value">উন্নত নিরাপত্তা বৈশিষ্ট্য</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা R15 V3 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">150cc লিকুইড কুলড কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="pro-description">রেস-অনুপ্রাণিত স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার হ্যান্ডলিং:</strong> <span class="pro-description">তীক্ষ্ণ এবং প্রতিক্রিয়াশীল</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বৈশিষ্ট্য:</strong> <span class="pro-description">ABS এবং আধুনিক প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি অনুভূতি:</strong> <span class="pro-description">রোমাঞ্চকর যাত্রার অভিজ্ঞতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা R15 V3 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চতর জ্বালানি খরচ:</strong> <span class="con-description">স্পোর্ট বাইক উচ্চ খরচ</span></li>
<li class="con-item"><strong class="con-label">আক্রমণাত্মক অবস্থান:</strong> <span class="con-description">এগিয়ে যাওয়ার ভঙ্গি</span></li>
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">আরও ব্যয়বহুল বিকল্প</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা R15 V3 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,50,000 টাকায়, ইয়ামাহা R15 V3 কর্মক্ষমতা-চাহনীয় রাইডারদের জন্য আদর্শ যারা চমৎকার হ্যান্ডলিং, আক্রমণাত্মক স্টাইলিং এবং প্রিমিয়াম বৈশিষ্ট্য সহ রোমাঞ্চকর স্পোর্ট বাইক খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্পোর্ট বাইক উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha R15 V3 already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Yamaha R15 V3\n")

	return nil
}

// GetName returns the seeder name
func (s *YamahaR15V3Review) GetName() string {
	return "YamahaR15V3Review"
}
