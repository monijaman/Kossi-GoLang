package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiIntruder150Review handles seeding Suzuki Intruder 150 product review and translation
type SuzukiIntruder150Review struct {
	BaseSeeder
}

// NewSuzukiIntruder150ReviewSeeder creates a new SuzukiIntruder150Review
func NewSuzukiIntruder150ReviewSeeder() *SuzukiIntruder150Review {
	return &SuzukiIntruder150Review{BaseSeeder: BaseSeeder{name: "Suzuki Intruder 150 Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiIntruder150Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Intruder 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki intruder 150 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding suzuki intruder 150 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Intruder 150 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Intruder 150 Review Bangladesh 2024 - Retro Cruiser Bike</h1>
<p class="summary-text">The Suzuki Intruder 150 is a stylish 150cc retro cruiser with commanding presence, low-slung seating, smooth engine, and timeless design. Priced around ৳2,75,000, it combines classic aesthetics with modern reliability for cruiser enthusiasts.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Intruder 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc Smooth Engine:</strong> <span class="highlight-value">Refined cruiser power</span></li>
<li class="highlight-item"><strong class="highlight-label">Retro Design:</strong> <span class="highlight-value">Classic cruiser styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Low-Slung Seating:</strong> <span class="highlight-value">Cruiser riding position</span></li>
<li class="highlight-item"><strong class="highlight-label">Commanding Presence:</strong> <span class="highlight-value">Road dominating look</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Intruder 150 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">150cc refined power</span></li>
<li class="pro-item"><strong class="pro-label">Retro Styling:</strong> <span class="pro-description">Timeless cruiser design</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Low-slung seating position</span></li>
<li class="pro-item"><strong class="pro-label">Quality Build:</strong> <span class="pro-description">Suzuki reliability</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">45-50 km/l efficiency</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Intruder 150 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Weight:</strong> <span class="con-description">Heavy cruiser bike</span></li>
<li class="con-item"><strong class="con-label">Maneuverability:</strong> <span class="con-description">Less agile in traffic</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium maintenance</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Intruder 150 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,75,000 - Retro cruiser</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">45-50 km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Intruder 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Design Appeal:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Timeless retro style</span></li>
<li class="rating-item"><strong class="rating-label">Engine Smoothness:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Refined 150cc</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Low-slung cruiser</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Suzuki reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Retro features worth</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Intruder 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,75,000, the Suzuki Intruder 150 is perfect for cruiser enthusiasts seeking timeless retro design, smooth performance, comfortable seating, and reliable Suzuki engineering.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For retro cruiser lovers</span></p>
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
		return fmt.Errorf("error creating suzuki intruder 150 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Intruder 150 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি Intruder 150 রিভিউ বাংলাদেশ 2024 - রেট্রো ক্রুজার বাইক</h1>
<p class="summary-text">সুজুকি Intruder 150 একটি স্টাইলিশ 150cc রেট্রো ক্রুজার যা কমান্ডিং উপস্থিতি, নিম্ন-সঙ্গ আসন, মসৃণ ইঞ্জিন এবং নিরবধি ডিজাইন সহ আসে। ৳2,75,000 টাকায় মূল্যায়িত, এটি ক্রুজার উত্সাহীদের জন্য ক্লাসিক নন্দনতত্ত্বকে আধুনিক নির্ভরযোগ্যতার সাথে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি Intruder 150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc মসৃণ ইঞ্জিন:</strong> <span class="highlight-value">পরিমার্জিত ক্রুজার শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">রেট্রো ডিজাইন:</strong> <span class="highlight-value">ক্লাসিক ক্রুজার স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">নিম্ন-সঙ্গ আসন:</strong> <span class="highlight-value">ক্রুজার যাত্রার অবস্থান</span></li>
<li class="highlight-item"><strong class="highlight-label">কমান্ডিং উপস্থিতি:</strong> <span class="highlight-value">রোড আধিপত্য চেহারা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি Intruder 150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">150cc পরিমার্জিত শক্তি</span></li>
<li class="pro-item"><strong class="pro-label">রেট্রো স্টাইলিং:</strong> <span class="pro-description">নিরবধি ক্রুজার ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক যাত্রা:</strong> <span class="pro-description">নিম্ন-সঙ্গ আসন অবস্থান</span></li>
<li class="pro-item"><strong class="pro-label">গুণমান নির্মাণ:</strong> <span class="pro-description">সুজুকি নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">45-50 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি Intruder 150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">ওজন:</strong> <span class="con-description">ভারী ক্রুজার বাইক</span></li>
<li class="con-item"><strong class="con-label">চালনা ক্ষমতা:</strong> <span class="con-description">ট্রাফিকে কম চটপটে</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম রক্ষণাবেক্ষণ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি Intruder 150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,75,000 টাকায়, সুজুকি Intruder 150 ক্রুজার উত্সাহীদের জন্য নিখুঁত যারা নিরবধি রেট্রো ডিজাইন, মসৃণ কর্মক্ষমতা, আরামদায়ক আসন এবং নির্ভরযোগ্য সুজুকি প্রযুক্তি খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- রেট্রো ক্রুজার প্রেমীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Intruder 150 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Intruder 150\n")

	return nil
}

// GetName returns the seeder name
func (s *SuzukiIntruder150Review) GetName() string {
	return "SuzukiIntruder150Review"
}
