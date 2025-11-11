package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSApacheRTR160ReviewBatch15Review handles seeding TVS Apache RTR 160 (RTR variant) product review and translation
type TVSApacheRTR160ReviewBatch15Review struct {
	BaseSeeder
}

// NewTVSApacheRTR160ReviewBatch15ReviewSeeder creates a new TVSApacheRTR160ReviewBatch15Review
func NewTVSApacheRTR160ReviewBatch15ReviewSeeder() *TVSApacheRTR160ReviewBatch15Review {
	return &TVSApacheRTR160ReviewBatch15Review{BaseSeeder: BaseSeeder{name: "TVS Apache RTR 160 RTR Review"}}
}

// Seed implements the Seeder interface
func (s *TVSApacheRTR160ReviewBatch15Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 160").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs apache rtr 160 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding tvs apache rtr 160 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Apache RTR 160 RTR already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Apache RTR 160 RTR Review Bangladesh 2024 - Racing Heritage Commuter</h1>
<p class="summary-text">The TVS Apache RTR 160 (RTR variant) is a racing-heritage 160cc performance bike with track-inspired design, powerful twin-spark engine, sharp handling, and aggressive styling. Priced at ৳2,82,000, it delivers thrilling performance with authentic racing DNA for passionate riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 160 RTR Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc Twin-Spark:</strong> <span class="highlight-value">Racing heritage power</span></li>
<li class="highlight-item"><strong class="highlight-label">Track-Inspired Design:</strong> <span class="highlight-value">Aggressive racing styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Responsive performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">45-50 km/l efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 160 RTR Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Racing Heritage:</strong> <span class="pro-description">Track-inspired DNA</span></li>
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">160cc twin-spark beast</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Responsive maneuvers</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Design:</strong> <span class="pro-description">Racing styling</span></li>
<li class="pro-item"><strong class="pro-label">TVS Reliability:</strong> <span class="pro-description">Racing-tested quality</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 160 RTR Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-oriented firm</span></li>
<li class="con-item"><strong class="con-label">Aggressive Ergonomics:</strong> <span class="con-description">Not for long touring</span></li>
<li class="con-item"><strong class="con-label">Vibration:</strong> <span class="con-description">Noticeable at high RPM</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 160 RTR Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,82,000 - Performance street bike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">45-50 km/l decent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 160 RTR Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Power:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Twin-spark powerful</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Track-inspired sharp</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Aggressive racing</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- TVS proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Performance focused</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 160 RTR?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,82,000, the TVS Apache RTR 160 RTR is perfect for passionate riders seeking authentic racing heritage, powerful twin-spark performance, track-inspired handling, and aggressive styling that delivers thrilling daily riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For racing enthusiasts</span></p>
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
		return fmt.Errorf("error creating tvs apache rtr 160 rtr review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Apache RTR 160 RTR (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">TVS Apache RTR 160 RTR রিভিউ বাংলাদেশ 2024 - রেসিং হেরিটেজ কমিউটার</h1>
<p class="summary-text">TVS Apache RTR 160 (RTR ভেরিয়েন্ট) একটি রেসিং-হেরিটেজ 160cc পারফরম্যান্স বাইক যা ট্র্যাক-অনুপ্রাণিত ডিজাইন, শক্তিশালী টুইন-স্পার্ক ইঞ্জিন, তীক্ষ্ণ হ্যান্ডলিং এবং আক্রমণাত্মক স্টাইলিং বৈশিষ্ট্যযুক্ত। ৳2,82,000 টাকায় মূল্যায়িত, এটি আবেগপ্রবণ রাইডারদের জন্য খাঁটি রেসিং DNA সহ রোমাঞ্চকর কর্মক্ষমতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 160 RTR মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc টুইন-স্পার্ক:</strong> <span class="highlight-value">রেসিং হেরিটেজ শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্র্যাক-অনুপ্রাণিত ডিজাইন:</strong> <span class="highlight-value">আক্রমণাত্মক রেসিং স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="highlight-value">প্রতিক্রিয়াশীল কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">45-50 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 160 RTR সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">রেসিং হেরিটেজ:</strong> <span class="pro-description">ট্র্যাক-অনুপ্রাণিত DNA</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">160cc টুইন-স্পার্ক শক্তিশালী</span></li>
<li class="pro-item"><strong class="pro-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="pro-description">প্রতিক্রিয়াশীল কৌশল</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="pro-description">রেসিং স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">TVS নির্ভরযোগ্যতা:</strong> <span class="pro-description">রেসিং-পরীক্ষিত গুণমান</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 160 RTR অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">ক্রীড়া-ভিত্তিক দৃঢ়</span></li>
<li class="con-item"><strong class="con-label">আক্রমণাত্মক এরগনমিক্স:</strong> <span class="con-description">দীর্ঘ ট্যুরিংয়ের জন্য নয়</span></li>
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">উচ্চ RPM উল্লেখযোগ্য</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: TVS Apache RTR 160 RTR কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,82,000 টাকায়, TVS Apache RTR 160 RTR আবেগপ্রবণ রাইডারদের জন্য নিখুঁত যারা খাঁটি রেসিং হেরিটেজ, শক্তিশালী টুইন-স্পার্ক কর্মক্ষমতা, ট্র্যাক-অনুপ্রাণিত হ্যান্ডলিং এবং আক্রমণাত্মক স্টাইলিং খুঁজছেন যা রোমাঞ্চকর দৈনিক রাইডিং প্রদান করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- রেসিং উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Apache RTR 160 RTR already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Apache RTR 160 RTR\n")

	return nil
}

// GetName returns the seeder name
func (s *TVSApacheRTR160ReviewBatch15Review) GetName() string {
	return "TVSApacheRTR160ReviewBatch15Review"
}
