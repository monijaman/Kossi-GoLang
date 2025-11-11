package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MaybachYMX110Review handles seeding Maybach YMX110 product review and translation
type MaybachYMX110Review struct {
	BaseSeeder
}

// NewMaybachYMX110ReviewSeeder creates a new MaybachYMX110Review
func NewMaybachYMX110ReviewSeeder() *MaybachYMX110Review {
	return &MaybachYMX110Review{BaseSeeder: BaseSeeder{name: "Maybach YMX110 Review"}}
}

// Seed implements the Seeder interface
func (s *MaybachYMX110Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Maybach YMX110").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("maybach ymx110 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding maybach ymx110 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Maybach YMX110 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Maybach YMX110 Review Bangladesh 2024 - Chinese Economy Champion</h1>
<p class="summary-text">The Maybach YMX110 is a 110cc economy commuter combining practical design with reliable performance and exceptional fuel efficiency. Priced around ৳1,48,000, it delivers everyday reliability, minimal maintenance, outstanding economy, and incredible value for budget-conscious daily riders seeking pure transportation efficiency.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Maybach YMX110 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc Economy:</strong> <span class="highlight-value">Budget commuter</span></li>
<li class="highlight-item"><strong class="highlight-label">Reliability:</strong> <span class="highlight-value">Daily trusted</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Saver:</strong> <span class="highlight-value">40+ km/l</span></li>
<li class="highlight-item"><strong class="highlight-label">Simple Design:</strong> <span class="highlight-value">Easy maintenance</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Maybach YMX110 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Economy:</strong> <span class="pro-description">40-45 km/l</span></li>
<li class="pro-item"><strong class="pro-label">Low Price:</strong> <span class="pro-description">Super budget</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Daily stable</span></li>
<li class="pro-item"><strong class="pro-label">Easy Service:</strong> <span class="pro-description">Cheap maintenance</span></li>
<li class="pro-item"><strong class="pro-label">Simple Build:</strong> <span class="pro-description">Hassle-free</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Maybach YMX110 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Minimal tech</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">Economy focused</span></li>
<li class="con-item"><strong class="con-label">Build Quality:</strong> <span class="con-description">Basic level</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Maybach YMX110 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,48,000 - Ultra-budget</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Minimal - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">40-45 km/l legendary</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Maybach YMX110 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Economy:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Best class</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Daily proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Unbeatable</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Basic only</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Minimal</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Maybach YMX110?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,48,000, the YMX110 is unbeatable for budget commuters needing maximum fuel economy, minimal maintenance, and pure daily transportation without premium features, delivering legendary value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultra-budget economy riders</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating maybach ymx110 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Maybach YMX110 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">মায়বাখ YMX110 রিভিউ বাংলাদেশ 2024 - চাইনিজ ইকোনমি চ্যাম্পিয়ন</h1>
<p class="summary-text">মায়বাখ YMX110 একটি 110cc অর্থনৈতিক যাত্রী যা ব্যবহারিক ডিজাইন, নির্ভরযোগ্য কর্মক্ষমতা এবং ব্যতিক্রমী জ্বালানি দক্ষতা একত্রিত করে। ৳1,48,000 টাকায় মূল্যায়িত, এটি দৈনন্দিন নির্ভরযোগ্যতা, ন্যূনতম রক্ষণাবেক্ষণ এবং চমৎকার মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">মায়বাখ YMX110 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc অর্থনীতি:</strong> <span class="highlight-value">বাজেট যাত্রী</span></li>
<li class="highlight-item"><strong class="highlight-label">নির্ভরযোগ্যতা:</strong> <span class="highlight-value">দৈনিক বিশ্বস্ত</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি সেভার:</strong> <span class="highlight-value">40+ km/l</span></li>
<li class="highlight-item"><strong class="highlight-label">সহজ ডিজাইন:</strong> <span class="highlight-value">সহজ রক্ষণাবেক্ষণ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">মায়বাখ YMX110 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার অর্থনীতি:</strong> <span class="pro-description">40-45 km/l</span></li>
<li class="pro-item"><strong class="pro-label">কম দাম:</strong> <span class="pro-description">সুপার বাজেট</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">দৈনিক স্থিতিশীল</span></li>
<li class="pro-item"><strong class="pro-label">সহজ সেবা:</strong> <span class="pro-description">সস্তা রক্ষণাবেক্ষণ</span></li>
<li class="pro-item"><strong class="pro-label">সহজ নির্মাণ:</strong> <span class="pro-description">ঝামেলা-মুক্ত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">মায়বাখ YMX110 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতম প্রযুক্তি</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">অর্থনীতি ফোকাস</span></li>
<li class="con-item"><strong class="con-label">বিল্ড গুণমান:</strong> <span class="con-description">মৌলিক স্তর</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: মায়বাখ YMX110 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,48,000 টাকায়, YMX110 বাজেট যাত্রীদের জন্য অপ্রতিদ্বন্দ্বী যারা সর্বোচ্চ জ্বালানি অর্থনীতি, ন্যূনতম রক্ষণাবেক্ষণ এবং খাঁটি দৈনন্দিন পরিবহন চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অতি-বাজেট অর্থনৈতিক চালকদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Maybach YMX110 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Maybach YMX110\n")

	return nil
}

// GetName returns the seeder name
func (s *MaybachYMX110Review) GetName() string {
	return "MaybachYMX110Review"
}
