package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCB350ReviewBatch16 handles seeding Honda CB350 product review and translation
type HondaCB350ReviewBatch16 struct {
	BaseSeeder
}

// NewHondaCB350ReviewBatch16Seeder creates a new HondaCB350ReviewBatch16
func NewHondaCB350ReviewBatch16Seeder() *HondaCB350ReviewBatch16 {
	return &HondaCB350ReviewBatch16{BaseSeeder: BaseSeeder{name: "Honda CB350 Batch16 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCB350ReviewBatch16) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB350").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda cb350 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding honda cb350 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CB350 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CB350 Review Bangladesh 2024 - Neo-Retro Mid-Size Cruiser</h1>
<p class="summary-text">The Honda CB350 is a neo-retro 350cc cruiser blending classic styling with modern performance. Priced around ৳3,95,000, it combines timeless design, reliable power, smooth handling, and authentic motorcycle character for riders seeking vintage vibes with contemporary reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB350 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">350cc Classic Engine:</strong> <span class="highlight-value">Smooth retro power</span></li>
<li class="highlight-item"><strong class="highlight-label">Neo-Retro Design:</strong> <span class="highlight-value">Classic modern styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Smooth Performance:</strong> <span class="highlight-value">Refined riding</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">40-45 km/l decent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB350 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Retro Design:</strong> <span class="pro-description">Classic modern blend</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">350cc refined</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Riding:</strong> <span class="pro-description">Cruiser friendly</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven quality</span></li>
<li class="pro-item"><strong class="pro-label">Versatile Bike:</strong> <span class="pro-description">Daily to touring</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB350 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">₳3,95,000 expensive</span></li>
<li class="con-item"><strong class="con-label">Performance:</strong> <span class="con-description">Not sport-oriented</span></li>
<li class="con-item"><strong class="con-label">Handling:</strong> <span class="con-description">Cruiser geometry</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB350 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,95,000 - Neo-retro cruiser</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳6-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">40-45 km/l decent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB350 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Retro classic</span></li>
<li class="rating-item"><strong class="rating-label">Engine Smoothness:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Refined 350cc</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Cruiser friendly</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Honda trusted</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Premium classic</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB350?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,95,000, the Honda CB350 is perfect for riders seeking an elegant neo-retro cruiser that blends classic motorcycle aesthetics, smooth performance, Honda's legendary reliability, and timeless design that transcends trends.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For retro-style enthusiasts</span></p>
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
		return fmt.Errorf("error creating honda cb350 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CB350 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হন্ডা CB350 রিভিউ বাংলাদেশ 2024 - নিও-রেট্রো মিড-সাইজ ক্রুজার</h1>
<p class="summary-text">হন্ডা CB350 একটি নিও-রেট্রো 350cc ক্রুজার যা ক্লাসিক স্টাইলিং সহ আধুনিক কর্মক্ষমতা মিশ্রিত করে। ৳3,95,000 টাকায় মূল্যায়িত, এটি সময়োচিত ডিজাইন, নির্ভরযোগ্য শক্তি, মসৃণ হ্যান্ডলিং এবং খাঁটি মোটরসাইকেল চরিত্র একত্রিত করে যারা সমসাময়িক নির্ভরযোগ্যতা সহ ভিন্টেজ অনুভূতি খুঁজছেন।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা CB350 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">350cc ক্লাসিক ইঞ্জিন:</strong> <span class="highlight-value">মসৃণ রেট্রো শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">নিও-রেট্রো ডিজাইন:</strong> <span class="highlight-value">ক্লাসিক আধুনিক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">মসৃণ কর্মক্ষমতা:</strong> <span class="highlight-value">পরিমার্জিত রাইডিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">40-45 km/l মানসম্পন্ন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা CB350 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">রেট্রো ডিজাইন:</strong> <span class="pro-description">ক্লাসিক আধুনিক মিশ্রণ</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">350cc পরিমার্জিত</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইডিং:</strong> <span class="pro-description">ক্রুজার বান্ধব</span></li>
<li class="pro-item"><strong class="pro-label">হন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">বহুমুখী বাইক:</strong> <span class="pro-description">দৈনিক থেকে ট্যুরিং</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা CB350 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳3,95,000 ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">কর্মক্ষমতা:</strong> <span class="con-description">ক্রীড়া-ভিত্তিক নয়</span></li>
<li class="con-item"><strong class="con-label">হ্যান্ডলিং:</strong> <span class="con-description">ক্রুজার জ্যামিতি</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হন্ডা CB350 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,95,000 টাকায়, হন্ডা CB350 রাইডারদের জন্য নিখুঁত যারা একটি মার্জিত নিও-রেট্রো ক্রুজার খুঁজছেন যা ক্লাসিক মোটরসাইকেল নান্দনিকতা, মসৃণ কর্মক্ষমতা, হন্ডার কিংবদন্তি নির্ভরযোগ্যতা এবং সময়মত ডিজাইন একত্রিত করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- রেট্রো-স্টাইল উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CB350 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CB350\n")

	return nil
}

// GetName returns the seeder name
func (s *HondaCB350ReviewBatch16) GetName() string {
	return "HondaCB350ReviewBatch16"
}
