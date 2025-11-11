package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HarleyDavidsonStreet750Review handles seeding Harley-Davidson Street 750 product review and translation
type HarleyDavidsonStreet750Review struct {
	BaseSeeder
}

// NewHarleyDavidsonStreet750ReviewSeeder creates a new HarleyDavidsonStreet750Review
func NewHarleyDavidsonStreet750ReviewSeeder() *HarleyDavidsonStreet750Review {
	return &HarleyDavidsonStreet750Review{BaseSeeder: BaseSeeder{name: "Harley-Davidson Street 750 Review"}}
}

// Seed implements the Seeder interface
func (s *HarleyDavidsonStreet750Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Harley-Davidson Street 750").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("harley-davidson street 750 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding harley-davidson street 750 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Harley-Davidson Street 750 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Harley-Davidson Street 750 Review Bangladesh 2024 - American Cruiser Legend</h1>
<p class="summary-text">The Harley-Davidson Street 750 is a 750cc liquid-cooled cruiser delivering authentic American muscle, iconic styling, and commanding presence. Priced around ৳12,50,000, it offers raw character, accessible power, premium feel, and unmatched brand prestige for cruiser enthusiasts seeking legendary American heritage.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Harley-Davidson Street 750 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">750cc Cruiser:</strong> <span class="highlight-value">American muscle</span></li>
<li class="highlight-item"><strong class="highlight-label">Liquid-Cooled:</strong> <span class="highlight-value">Modern reliability</span></li>
<li class="highlight-item"><strong class="highlight-label">Iconic Design:</strong> <span class="highlight-value">Harley presence</span></li>
<li class="highlight-item"><strong class="highlight-label">Brand Prestige:</strong> <span class="highlight-value">Legend status</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Harley-Davidson Street 750 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Iconic Brand:</strong> <span class="pro-description">Legendary heritage</span></li>
<li class="pro-item"><strong class="pro-label">Distinctive Rumble:</strong> <span class="pro-description">V-Twin character</span></li>
<li class="pro-item"><strong class="pro-label">Cruiser Comfort:</strong> <span class="pro-description">Relaxed riding</span></li>
<li class="pro-item"><strong class="pro-label">Premium Build:</strong> <span class="pro-description">Quality feel</span></li>
<li class="pro-item"><strong class="pro-label">Community:</strong> <span class="pro-description">Strong worldwide</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Harley-Davidson Street 750 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price:</strong> <span class="con-description">৳12,50,000 premium</span></li>
<li class="con-item"><strong class="con-label">Maintenance:</strong> <span class="con-description">Costly upkeep</span></li>
<li class="con-item"><strong class="con-label">Fuel Economy:</strong> <span class="con-description">18-21 km/l low</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Harley-Davidson Street 750 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳12,50,000 - Premium cruiser</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳15-18 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">18-21 km/l modest</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Harley-Davidson Street 750 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- V-Twin character</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Cruiser stable</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Relaxed cruising</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Legendary icon</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Premium positioned</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Harley-Davidson Street 750?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳12,50,000, the Street 750 is for cruiser enthusiasts valuing American legend status, iconic V-Twin character, relaxed comfort, and unmatched brand prestige as ultimate cruising machine.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For American cruiser lovers</span></p>
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
		return fmt.Errorf("error creating harley-davidson street 750 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Harley-Davidson Street 750 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হার্লে-ডেভিডসন স্ট্রিট 750 রিভিউ বাংলাদেশ 2024 - আমেরিকান ক্রুজার লেজেন্ড</h1>
<p class="summary-text">হার্লে-ডেভিডসন স্ট্রিট 750 একটি 750cc লিকুইড-কুল্ড ক্রুজার যা খাঁটি আমেরিকান মাসল, আইকনিক স্টাইলিং এবং কমান্ডিং উপস্থিতি প্রদান করে। ৳12,50,000 টাকায় মূল্যায়িত, এটি কাঁচা চরিত্র, অ্যাক্সেসযোগ্য শক্তি এবং অতুলনীয় ব্র্যান্ড প্রতিপত্তি অফার করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হার্লে-ডেভিডসন স্ট্রিট 750 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">750cc ক্রুজার:</strong> <span class="highlight-value">আমেরিকান মাসল</span></li>
<li class="highlight-item"><strong class="highlight-label">লিকুইড-কুল্ড:</strong> <span class="highlight-value">আধুনিক নির্ভরযোগ্যতা</span></li>
<li class="highlight-item"><strong class="highlight-label">আইকনিক ডিজাইন:</strong> <span class="highlight-value">হার্লে উপস্থিতি</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্র্যান্ড প্রতিপত্তি:</strong> <span class="highlight-value">কিংবদন্তি অবস্থা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হার্লে-ডেভিডসন স্ট্রিট 750 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আইকনিক ব্র্যান্ড:</strong> <span class="pro-description">কিংবদন্তি ঐতিহ্য</span></li>
<li class="pro-item"><strong class="pro-label">স্বতন্ত্র রম্বল:</strong> <span class="pro-description">V-টুইন চরিত্র</span></li>
<li class="pro-item"><strong class="pro-label">ক্রুজার আরাম:</strong> <span class="pro-description">শিথিল চড়া</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বিল্ড:</strong> <span class="pro-description">গুণমান অনুভূতি</span></li>
<li class="pro-item"><strong class="pro-label">কমিউনিটি:</strong> <span class="pro-description">শক্তিশালী বিশ্বব্যাপী</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হার্লে-ডেভিডসন স্ট্রিট 750 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">৳12,50,000 প্রিমিয়াম</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ:</strong> <span class="con-description">ব্যয়বহুল রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি অর্থনীতি:</strong> <span class="con-description">18-21 km/l কম</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হার্লে-ডেভিডসন স্ট্রিট 750 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳12,50,000 টাকায়, স্ট্রিট 750 ক্রুজার উত্সাহীদের জন্য যারা আমেরিকান কিংবদন্তি অবস্থা, আইকনিক V-টুইন চরিত্র এবং অতুলনীয় ব্র্যান্ড প্রতিপত্তি মূল্যবান করেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আমেরিকান ক্রুজার প্রেমীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Harley-Davidson Street 750 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Harley-Davidson Street 750\n")

	return nil
}

// GetName returns the seeder name
func (s *HarleyDavidsonStreet750Review) GetName() string {
	return "HarleyDavidsonStreet750Review"
}
