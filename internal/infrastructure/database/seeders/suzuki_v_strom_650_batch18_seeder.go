package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiV-Strom650ReviewBatch18 handles seeding Suzuki V-Strom 650 product review and translation
type SuzukiVStrom650ReviewBatch18 struct {
	BaseSeeder
}

// NewSuzukiVStrom650ReviewBatch18Seeder creates a new SuzukiVStrom650ReviewBatch18
func NewSuzukiVStrom650ReviewBatch18Seeder() *SuzukiVStrom650ReviewBatch18 {
	return &SuzukiVStrom650ReviewBatch18{BaseSeeder: BaseSeeder{name: "Suzuki V-Strom 650 Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiVStrom650ReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki V-Strom 650").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki v-strom 650 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding suzuki v-strom 650 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki V-Strom 650 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki V-Strom 650 Review Bangladesh 2024 - Adventure Tourer Legendary Reliability</h1>
<p class="summary-text">The Suzuki V-Strom 650 is a 650cc V-twin adventure tourer combining proven reliability with adventure capability. Priced around ৳8,95,000, it delivers balanced performance, excellent fuel efficiency, versatile riding position, advanced safety features, and the Suzuki reliability reputation for global touring enthusiasts seeking dependable adventure platforms.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki V-Strom 650 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">650cc V-Twin:</strong> <span class="highlight-value">Balanced adventure power</span></li>
<li class="highlight-item"><strong class="highlight-label">Adventure Tourer:</strong> <span class="highlight-value">Global capability proven</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficiency:</strong> <span class="highlight-value">24-26 km/l excellent</span></li>
<li class="highlight-item"><strong class="highlight-label">Reliability:</strong> <span class="highlight-value">Suzuki heritage</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki V-Strom 650 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Reliability Legend:</strong> <span class="pro-description">Proven global tours</span></li>
<li class="pro-item"><strong class="pro-label">Adventure Versatile:</strong> <span class="pro-description">On/off-road capable</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Efficient:</strong> <span class="pro-description">24-26 km/l excellent</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Riding:</strong> <span class="pro-description">Upright ergonomic</span></li>
<li class="pro-item"><strong class="pro-label">Safety Features:</strong> <span class="pro-description">ABS equipped</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki V-Strom 650 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Moderate Power:</strong> <span class="con-description">645cc 51 bhp</span></li>
<li class="con-item"><strong class="con-label">Handling Weight:</strong> <span class="con-description">216kg substantial</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Regular maintenance</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki V-Strom 650 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳8,95,000 - Adventure tourer</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳10-12 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">24-26 km/l excellent</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">645cc liquid-cooled V-twin</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">51 bhp reliable torque</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">216kg adventure ready</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki V-Strom 650 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Legendary proven</span></li>
<li class="rating-item"><strong class="rating-label">Adventure Capability:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Versatile proven</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Long-distance capable</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- 24-26 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- ৳8,95,000 solid</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki V-Strom 650?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳8,95,000, the V-Strom 650 is ideal for adventure touring enthusiasts seeking legendary Suzuki reliability, versatile adventure capability, excellent fuel efficiency, and comfortable upright ergonomics for global touring adventures with proven dependability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For adventure touring seekers</span></p>
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
		return fmt.Errorf("error creating suzuki v-strom 650 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki V-Strom 650 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি ভি-স্ট্রম 650 রিভিউ বাংলাদেশ 2024 - অ্যাডভেঞ্চার ট্যুরার কিংবদন্তি নির্ভরযোগ্যতা</h1>
<p class="summary-text">সুজুকি ভি-স্ট্রম 650 একটি 650cc ভি-টুইন অ্যাডভেঞ্চার ট্যুরার যা প্রমাণিত নির্ভরযোগ্যতাকে অ্যাডভেঞ্চার ক্ষমতার সাথে একত্রিত করে। ৳8,95,000 টাকায় মূল্যায়িত, এটি ভারসাম্যপূর্ণ পারফরম্যান্স এবং চমৎকার জ্বালানি দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি ভি-স্ট্রম 650 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">650cc ভি-টুইন:</strong> <span class="highlight-value">ভারসাম্যপূর্ণ অ্যাডভেঞ্চার শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">অ্যাডভেঞ্চার ট্যুরার:</strong> <span class="highlight-value">গ্লোবাল ক্ষমতা প্রমাণিত</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষতা:</strong> <span class="highlight-value">24-26 km/l চমৎকার</span></li>
<li class="highlight-item"><strong class="highlight-label">নির্ভরযোগ্যতা:</strong> <span class="highlight-value">সুজুকি হেরিটেজ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি ভি-স্ট্রম 650 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্যতা কিংবদন্তি:</strong> <span class="pro-description">প্রমাণিত গ্লোবাল ট্যুর</span></li>
<li class="pro-item"><strong class="pro-label">অ্যাডভেঞ্চার বহুমুখী:</strong> <span class="pro-description">অন/অফ-রোড সক্ষম</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি দক্ষ:</strong> <span class="pro-description">24-26 km/l চমৎকার</span></li>
<li class="pro-item"><strong class="pro-label">আরাম রাইডিং:</strong> <span class="pro-description">সোজা এরগোনমিক</span></li>
<li class="pro-item"><strong class="pro-label">নিরাপত্তা বৈশিষ্ট্য:</strong> <span class="pro-description">ABS সজ্জিত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি ভি-স্ট্রম 650 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মধ্যম শক্তি:</strong> <span class="con-description">645cc 51 bhp</span></li>
<li class="con-item"><strong class="con-label">হ্যান্ডলিং ওজন:</strong> <span class="con-description">216kg উল্লেখযোগ্য</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">নিয়মিত রক্ষণাবেক্ষণ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি ভি-স্ট্রম 650 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳8,95,000 টাকায়, ভি-স্ট্রম 650 অ্যাডভেঞ্চার ট্যুরিং উত্সাহীদের জন্য আদর্শ যারা কিংবদন্তি সুজুকি নির্ভরযোগ্যতা এবং বহুমুখী অ্যাডভেঞ্চার ক্ষমতা চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অ্যাডভেঞ্চার ট্যুরিং সন্ধানকারীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki V-Strom 650 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki V-Strom 650\n")

	return nil
}

// GetName returns the seeder name
func (s *SuzukiVStrom650ReviewBatch18) GetName() string {
	return "SuzukiVStrom650ReviewBatch18"
}
