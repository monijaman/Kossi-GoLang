package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarN160Review handles seeding Bajaj Pulsar N160 product review and translation
type BajajPulsarN160Review struct {
	BaseSeeder
}

// NewBajajPulsarN160ReviewSeeder creates a new BajajPulsarN160Review
func NewBajajPulsarN160ReviewSeeder() *BajajPulsarN160Review {
	return &BajajPulsarN160Review{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar N160 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarN160Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar N160").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar n160 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj pulsar n160 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar N160 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar N160 Review Bangladesh 2024 - Naked Street Muscle</h1>
<p class="summary-text">The Bajaj Pulsar N160 is a bare-naked 160cc street bike with aggressive styling, powerful engine, responsive handling, and street-fighter attitude. Priced around ৳2,70,000, it delivers thrilling performance with commanding road presence.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar N160 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc Powerful Engine:</strong> <span class="highlight-value">Street fighter performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Naked Design:</strong> <span class="highlight-value">Aggressive minimalist styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Responsive maneuverability</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">45-50 km/l efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar N160 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">160cc street performance</span></li>
<li class="pro-item"><strong class="pro-label">Naked Design:</strong> <span class="pro-description">Aggressive styling</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Responsive and nimble</span></li>
<li class="pro-item"><strong class="pro-label">Bajaj Reliability:</strong> <span class="pro-description">Trusted performance</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Good value performance</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar N160 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Firm seat for long rides</span></li>
<li class="con-item"><strong class="con-label">No Fairing:</strong> <span class="con-description">No wind protection</span></li>
<li class="con-item"><strong class="con-label">Vibration:</strong> <span class="con-description">Noticeable at high revs</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar N160 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,70,000 - Street naked bike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">45-50 km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar N160 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Power:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- 160cc street</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Responsive nimble</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Aggressive naked</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Bajaj proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good performance</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar N160?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,70,000, the Bajaj Pulsar N160 is perfect for street riders seeking aggressive styling, powerful performance, sharp handling, and authentic naked bike character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For naked street riders</span></p>
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
		return fmt.Errorf("error creating bajaj pulsar n160 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar N160 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ Pulsar N160 রিভিউ বাংলাদেশ 2024 - ন্যাকেড স্ট্রিট মাসল</h1>
<p class="summary-text">বাজাজ Pulsar N160 একটি খালি ন্যাকেড 160cc স্ট্রিট বাইক যা আক্রমণাত্মক স্টাইলিং, শক্তিশালী ইঞ্জিন, প্রতিক্রিয়াশীল হ্যান্ডলিং এবং স্ট্রিট-ফাইটার মনোভাব সহ আসে। ৳2,70,000 টাকায় মূল্যায়িত, এটি কমান্ডিং রোড উপস্থিতি সহ রোমাঞ্চকর কর্মক্ষমতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ Pulsar N160 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc শক্তিশালী ইঞ্জিন:</strong> <span class="highlight-value">স্ট্রিট ফাইটার কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ন্যাকেড ডিজাইন:</strong> <span class="highlight-value">আক্রমণাত্মক মিনিমালিস্ট স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="highlight-value">প্রতিক্রিয়াশীল চালনা ক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">45-50 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ Pulsar N160 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">160cc স্ট্রিট কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">ন্যাকেড ডিজাইন:</strong> <span class="pro-description">আক্রমণাত্মক স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="pro-description">প্রতিক্রিয়াশীল এবং চটপটে</span></li>
<li class="pro-item"><strong class="pro-label">বাজাজ নির্ভরযোগ্যতা:</strong> <span class="pro-description">বিশ্বস্ত কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">ভাল মূল্য কর্মক্ষমতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ Pulsar N160 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">দীর্ঘ যাত্রার জন্য দৃঢ় আসন</span></li>
<li class="con-item"><strong class="con-label">কোন ফেয়ারিং নেই:</strong> <span class="con-description">কোন বায়ু সুরক্ষা নেই</span></li>
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">উচ্চ বিপ্লব উল্লেখযোগ্য</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ Pulsar N160 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,70,000 টাকায়, বাজাজ Pulsar N160 স্ট্রিট রাইডারদের জন্য নিখুঁত যারা আক্রমণাত্মক স্টাইলিং, শক্তিশালী কর্মক্ষমতা, তীক্ষ্ণ হ্যান্ডলিং এবং খাঁটি ন্যাকেড বাইক চরিত্র খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ন্যাকেড স্ট্রিট রাইডারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar N160 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar N160\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsarN160Review) GetName() string {
	return "BajajPulsarN160Review"
}
