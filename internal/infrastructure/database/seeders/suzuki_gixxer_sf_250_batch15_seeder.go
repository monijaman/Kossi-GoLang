package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxerSF250Batch15 handles seeding Suzuki Gixxer SF 250 product review and translation
type SuzukiGixxerSF250Batch15 struct {
	BaseSeeder
}

// NewSuzukiGixxerSF250Batch15Seeder creates a new SuzukiGixxerSF250Batch15
func NewSuzukiGixxerSF250Batch15Seeder() *SuzukiGixxerSF250Batch15 {
	return &SuzukiGixxerSF250Batch15{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer SF 250 Batch15"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxerSF250Batch15) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer SF 250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gixxer sf 250 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding suzuki gixxer sf 250 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer SF 250 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer SF 250 Review Bangladesh 2024 - Powerful Faired Sports Bike</h1>
<p class="summary-text">The Suzuki Gixxer SF 250 is an impressive 250cc faired sports bike with liquid cooling, digital displays, aggressive aerodynamic design, and excellent performance. Priced around ৳3,75,000, it offers powerful sport riding with premium Suzuki engineering.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer SF 250 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc Liquid-Cooled:</strong> <span class="highlight-value">Powerful performance engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Faired Design:</strong> <span class="highlight-value">Aerodynamic sports styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Display:</strong> <span class="highlight-value">Modern instrumentation</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Brakes:</strong> <span class="highlight-value">Safe stopping power</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer SF 250 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">250cc liquid-cooled punch</span></li>
<li class="pro-item"><strong class="pro-label">Faired Protection:</strong> <span class="pro-description">Aerodynamic wind protection</span></li>
<li class="pro-item"><strong class="pro-label">Digital Tech:</strong> <span class="pro-description">Modern instrumentation</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">32-38 km/l efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Quality:</strong> <span class="pro-description">Premium engineering</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer SF 250 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">Expensive sports bike</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium maintenance</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">Higher consumption rate</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer SF 250 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,75,000 - Premium faired sports</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Higher - ৳7-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">32-38 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer SF 250 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Power:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- 250cc liquid-cooled</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Aerodynamic fairing</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Premium Suzuki</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Thrilling delivery</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Premium features</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer SF 250?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,75,000, the Suzuki Gixxer SF 250 is ideal for performance seekers wanting a powerful faired sports bike with premium engineering, aerodynamic protection, and thrilling sport riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sports bike enthusiasts</span></p>
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
		return fmt.Errorf("error creating suzuki gixxer sf 250 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer SF 250 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি Gixxer SF 250 রিভিউ বাংলাদেশ 2024 - শক্তিশালী ফেয়ার্ড স্পোর্ট বাইক</h1>
<p class="summary-text">সুজুকি Gixxer SF 250 একটি চিত্তাকর্ষক 250cc ফেয়ার্ড স্পোর্ট বাইক যা লিকুইড কুলিং, ডিজিটাল ডিসপ্লে, আক্রমণাত্মক এয়ারোডাইনামিক ডিজাইন এবং চমৎকার কর্মক্ষমতা সহ আসে। ৳3,75,000 টাকায় মূল্যায়িত, এটি প্রিমিয়াম সুজুকি প্রযুক্তি সহ শক্তিশালী স্পোর্ট যাত্রা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি Gixxer SF 250 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc লিকুইড-কুলড:</strong> <span class="highlight-value">শক্তিশালী কর্মক্ষমতা ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">ফেয়ার্ড ডিজাইন:</strong> <span class="highlight-value">এয়ারোডাইনামিক স্পোর্ট স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল ডিসপ্লে:</strong> <span class="highlight-value">আধুনিক উপকরণ</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার ব্রেক:</strong> <span class="highlight-value">নিরাপদ থামার শক্তি</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি Gixxer SF 250 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">250cc লিকুইড-কুলড পাঞ্চ</span></li>
<li class="pro-item"><strong class="pro-label">ফেয়ার্ড সুরক্ষা:</strong> <span class="pro-description">এয়ারোডাইনামিক বায়ু সুরক্ষা</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল প্রযুক্তি:</strong> <span class="pro-description">আধুনিক উপকরণ</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">32-38 km/l দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি গুণমান:</strong> <span class="pro-description">প্রিমিয়াম প্রযুক্তি</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি Gixxer SF 250 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">ব্যয়বহুল স্পোর্ট বাইক</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">উচ্চতর খরচ হার</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি Gixxer SF 250 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,75,000 টাকায়, সুজুকি Gixxer SF 250 কর্মক্ষমতা খোঁজারদের জন্য আদর্শ যারা শক্তিশালী ফেয়ার্ড স্পোর্ট বাইক, প্রিমিয়াম প্রযুক্তি এবং রোমাঞ্চকর স্পোর্ট যাত্রা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্পোর্ট বাইক উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer SF 250 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer SF 250\n")

	return nil
}

// GetName returns the seeder name
func (s *SuzukiGixxerSF250Review) GetName() string {
	return "SuzukiGixxerSF250Review"
}
