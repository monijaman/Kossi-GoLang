package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroMotocorpXPulseReview handles seeding Hero MotoCorp Xpulse product review and translation
type HeroMotocorpXPulseReview struct {
	BaseSeeder
}

// NewHeroMotocorpXPulseReviewSeeder creates a new HeroMotocorpXPulseReview
func NewHeroMotocorpXPulseReviewSeeder() *HeroMotocorpXPulseReview {
	return &HeroMotocorpXPulseReview{BaseSeeder: BaseSeeder{name: "Hero MotoCorp Xpulse Review"}}
}

// Seed implements the Seeder interface
func (s *HeroMotocorpXPulseReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Xpulse 200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero xpulse 200 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding hero xpulse 200 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero Xpulse 200 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero Xpulse 200 Review Bangladesh 2024 - Adventure Tryst Bike</h1>
<p class="summary-text">The Hero Xpulse 200 is an affordable adventure bike designed for enthusiasts seeking offroad capability with comfort. Priced around ৳3,05,000, it combines rugged design, versatile performance, and practical features for exploring diverse terrains.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Xpulse 200 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc Adventure Engine:</strong> <span class="highlight-value">Offroad ready power</span></li>
<li class="highlight-item"><strong class="highlight-label">Rugged Design:</strong> <span class="highlight-value">Adventure-oriented styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Ground Clearance:</strong> <span class="highlight-value">Superior offroad capability</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">40-45 km/l decent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Xpulse 200 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Rugged Design:</strong> <span class="pro-description">Adventure ready</span></li>
<li class="pro-item"><strong class="pro-label">Good Handling:</strong> <span class="pro-description">Offroad capable</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Budget adventure</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seating:</strong> <span class="pro-description">Long-ride friendly</span></li>
<li class="pro-item"><strong class="pro-label">Hero Reliability:</strong> <span class="pro-description">Proven performance</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Xpulse 200 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Performance:</strong> <span class="con-description">Budget adventure level</span></li>
<li class="con-item"><strong class="con-label">Braking:</strong> <span class="con-description">Could be better</span></li>
<li class="con-item"><strong class="con-label">Features:</strong> <span class="con-description">Basic equipment</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Xpulse 200 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,05,000 - Adventure tourer</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">40-45 km/l decent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Xpulse 200 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Adventure Capability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Budget adventure</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Offroad decent</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Suitable touring</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Hero proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Affordable adventure</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Xpulse 200?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,05,000, the Hero Xpulse 200 is perfect for adventure seekers wanting an affordable, reliable bike that combines daily commuting practicality with genuine offroad exploration capability and rugged durability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For adventure enthusiasts</span></p>
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
		return fmt.Errorf("error creating hero xpulse 200 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero Xpulse 200 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো Xpulse 200 রিভিউ বাংলাদেশ 2024 - অ্যাডভেঞ্চার ট্রিস্ট বাইক</h1>
<p class="summary-text">হিরো Xpulse 200 একটি সাশ্রয়ী অ্যাডভেঞ্চার বাইক যা অফরোড ক্ষমতা সহ আরাম খুঁজছেন উত্সাহীদের জন্য ডিজাইন করা হয়েছে। ৳3,05,000 টাকায় মূল্যায়িত, এটি বৈচিত্র্যময় ভূগোল অন্বেষণের জন্য কঠোর ডিজাইন, বহুমুখী কর্মক্ষমতা এবং ব্যবহারিক বৈশিষ্ট্য একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো Xpulse 200 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc অ্যাডভেঞ্চার ইঞ্জিন:</strong> <span class="highlight-value">অফরোড প্রস্তুত শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">কঠোর ডিজাইন:</strong> <span class="highlight-value">অ্যাডভেঞ্চার-ভিত্তিক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="highlight-value">উচ্চতর অফরোড ক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">40-45 km/l মানসম্পন্ন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো Xpulse 200 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">কঠোর ডিজাইন:</strong> <span class="pro-description">অ্যাডভেঞ্চার প্রস্তুত</span></li>
<li class="pro-item"><strong class="pro-label">ভাল হ্যান্ডলিং:</strong> <span class="pro-description">অফরোড সক্ষম</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">বাজেট অ্যাডভেঞ্চার</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক আসন:</strong> <span class="pro-description">দীর্ঘ-যাত্রা বান্ধব</span></li>
<li class="pro-item"><strong class="pro-label">হিরো নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত কর্মক্ষমতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো Xpulse 200 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">কর্মক্ষমতা:</strong> <span class="con-description">বাজেট অ্যাডভেঞ্চার স্তর</span></li>
<li class="con-item"><strong class="con-label">ব্রেকিং:</strong> <span class="con-description">আরও ভাল হতে পারে</span></li>
<li class="con-item"><strong class="con-label">বৈশিষ্ট্য:</strong> <span class="con-description">মৌলিক সরঞ্জাম</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো Xpulse 200 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,05,000 টাকায়, হিরো Xpulse 200 অ্যাডভেঞ্চার সন্ধানীদের জন্য নিখুঁত যারা একটি সাশ্রয়ী, নির্ভরযোগ্য বাইক খুঁজছেন যা দৈনন্দিন যাতায়াত ব্যবহারিকতা সহ খাঁটি অফরোড অন্বেষণ ক্ষমতা এবং কঠোর স্থায়িত্ব একত্রিত করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অ্যাডভেঞ্চার উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero Xpulse 200 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero Xpulse 200\n")

	return nil
}

// GetName returns the seeder name
func (s *HeroMotocorpXPulseReview) GetName() string {
	return "HeroMotocorpXPulseReview"
}
