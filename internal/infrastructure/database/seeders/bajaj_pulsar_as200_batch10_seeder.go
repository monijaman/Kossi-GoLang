package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarAS200Review handles seeding Bajaj Pulsar AS200 product review and translation
type BajajPulsarAS200Review struct {
	BaseSeeder
}

// NewBajajPulsarAS200ReviewSeeder creates a new BajajPulsarAS200Review
func NewBajajPulsarAS200ReviewSeeder() *BajajPulsarAS200Review {
	return &BajajPulsarAS200Review{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar AS200 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarAS200Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar AS200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Pulsar AS200 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Pulsar AS200 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar AS200 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar AS200 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Pulsar AS200 is an adventurous street sport bike featuring 200cc engine, ABS braking, semi-fairing, rugged design, and all-terrain capability. Priced around ৳2,45,000, it combines sport performance with adventure styling, perfect for riders seeking versatile highway and off-road capability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar AS200 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc Engine:</strong> <span class="highlight-value">Liquid-cooled DTS-i technology</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS Safety:</strong> <span class="highlight-value">Front wheel ABS protection</span></li>
<li class="highlight-item"><strong class="highlight-label">Semi-Fairing:</strong> <span class="highlight-value">Adventure sport styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Rugged Design:</strong> <span class="highlight-value">Off-road capable stance</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar AS200 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">Strong 200cc with excellent torque</span></li>
<li class="pro-item"><strong class="pro-label">Versatile Design:</strong> <span class="pro-description">Sport and adventure styling combined</span></li>
<li class="pro-item"><strong class="pro-label">ABS Protection:</strong> <span class="pro-description">Safety-focused braking system</span></li>
<li class="pro-item"><strong class="pro-label">Off-Road Capability:</strong> <span class="pro-description">Can handle rough terrain</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">35-40 km/l efficiency</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar AS200 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Moderate Pricing:</strong> <span class="con-description">More expensive than smaller models</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">Higher running cost than 150cc</span></li>
<li class="con-item"><strong class="con-label">Not Pure Adventure:</strong> <span class="con-description">Lighter than true ADV bikes</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar AS200 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,45,000 - Versatile sport-adventure value</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳6-7 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳6-7 per km (35-40 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar AS200 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good 200cc acceleration</span></li>
<li class="rating-item"><strong class="rating-label">Versatility:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Sport and adventure capability</span></li>
<li class="rating-item"><strong class="rating-label">Safety:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- ABS protection</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Modern features included</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good versatile value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar AS200?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,45,000, the Bajaj Pulsar AS200 offers versatile sport-adventure capability with modern safety features. Excellent for riders seeking a do-it-all bike combining highway performance with off-road character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For versatile adventure sport enthusiasts</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Bajaj Pulsar AS200 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar AS200 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাज পালসার AS200 রিভিউ বাংলাদেশ 2024 - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাज পালসার AS200 একটি অ্যাডভেঞ্চারাস স্ট্রিট স্পোর্ট বাইক যা 200cc ইঞ್জিন, ABS ব্রেকিং, সেমি-ফেয়ারিং, রাগড ডিজাইন এবং সব-ভূমি সামর্থ্য বৈশিষ্ট্যযুক্ত। ৳2,45,000 টাকায় মূল্যায়িত, এটি স্পোর্ট পারফরমেন্স এবং অ্যাডভেঞ্চার স্টাইলিং সংমিশ্রণ করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার AS200 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc ইঞ্জিন:</strong> <span class="highlight-value">লিকুইড-কুল্ড DTS-i প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS নিরাপত্তা:</strong> <span class="highlight-value">সামনের চাকা ABS সুরক্ষা</span></li>
<li class="highlight-item"><strong class="highlight-label">সেমি-ফেয়ারিং:</strong> <span class="highlight-value">অ্যাডভেঞ্চার স্পোর্ট স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">রাগড ডিজাইন:</strong> <span class="highlight-value">অফ-রোড সক্ষম অবস্থান</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার AS200 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">চমৎকার টর্কের সাথে শক্তিশালী 200cc</span></li>
<li class="pro-item"><strong class="pro-label">বহুমুখী ডিজাইন:</strong> <span class="pro-description">স্পোর্ট এবং অ্যাডভেঞ্চার স্টাইলিং সংমিশ্রিত</span></li>
<li class="pro-item"><strong class="pro-label">ABS সুরক্ষা:</strong> <span class="pro-description">নিরাপত্তা-ফোকাস ব্রেকিং সিস্টেম</span></li>
<li class="pro-item"><strong class="pro-label">অফ-রোড সামর্থ্য:</strong> <span class="pro-description">রুক্ষ ভূখণ্ড পরিচালনা করতে পারে</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">35-40 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার AS200 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মধ্যম মূল্য:</strong> <span class="con-description">ছোট মডেলের চেয়ে বেশি ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">150cc এর চেয়ে বেশি চলমান খরচ</span></li>
<li class="con-item"><strong class="con-label">খাঁটি অ্যাডভেঞ্চার নয়:</strong> <span class="con-description">সত্যিকারের ADV বাইকের চেয়ে হালকা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ পালসার AS200 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,45,000 টাকায়, বাজাজ পালসার AS200 আধুনিক নিরাপত্তা বৈশিষ্ট্যের সাথে বহুমুখী স্পোর্ট-অ্যাডভেঞ্চার সামর্থ্য প্রদান করে। যারা হাইওয়ে পারফরমেন্স এবং অফ-রোড ক্যারেক্টার একত্রিত করে এমন একটি সর্বজনীন বাইক খুঁজছেন তাদের জন্য চমৎকার।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বহুমুখী অ্যাডভেঞ্চার স্পোর্ট উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar AS200 already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating Bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar AS200\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsarAS200Review) GetName() string {
	return "BajajPulsarAS200Review"
}
