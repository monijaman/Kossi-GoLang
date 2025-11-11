package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSApacheRTR160TwoWheelerReview handles seeding TVS Apache RTR 160 product review and translation
type TVSApacheRTR160TwoWheelerReview struct {
	BaseSeeder
}

// NewTVSApacheRTR160TwoWheelerReviewSeeder creates a new TVSApacheRTR160TwoWheelerReview
func NewTVSApacheRTR160TwoWheelerReviewSeeder() *TVSApacheRTR160TwoWheelerReview {
	return &TVSApacheRTR160TwoWheelerReview{BaseSeeder: BaseSeeder{name: "TVS Apache RTR 160 2W Review"}}
}

// Seed implements the Seeder interface
func (s *TVSApacheRTR160TwoWheelerReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 160 2 Wheeler").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs apache rtr 160 2w product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding tvs apache rtr 160 2w product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Apache RTR 160 2W already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Apache RTR 160 2W Review Bangladesh 2024 - Performance Street Bike</h1>
<p class="summary-text">The TVS Apache RTR 160 is a sporty 160cc performance bike with responsive handling, aggressive design, twin-spark technology, and excellent maneuverability. Priced around ৳2,80,000, it offers thrilling street performance with reliable TVS engineering.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 160 2W Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc Twin-Spark:</strong> <span class="highlight-value">Responsive performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Design:</strong> <span class="highlight-value">Street fighter styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Excellent maneuverability</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">45-50 km/l efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 160 2W Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Responsive Engine:</strong> <span class="pro-description">160cc twin-spark punch</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Nimble and agile</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Street fighter design</span></li>
<li class="pro-item"><strong class="pro-label">TVS Reliability:</strong> <span class="pro-description">Trusted performance</span></li>
<li class="pro-item"><strong class="pro-label">Good Value:</strong> <span class="pro-description">Affordable performance</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 160 2W Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Hard seat for long rides</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Moderate maintenance</span></li>
<li class="con-item"><strong class="con-label">Vibration:</strong> <span class="con-description">Noticeable at high rpm</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 160 2W Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,80,000 - Performance street</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">45-50 km/l</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 160 2W Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Response:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 160cc twin-spark</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Sharp and nimble</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Aggressive street</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- TVS proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good performance</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 160 2W?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,80,000, the TVS Apache RTR 160 is perfect for performance seekers wanting responsive handling, aggressive styling, twin-spark punch, and affordable street performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For street performance riders</span></p>
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
		return fmt.Errorf("error creating tvs apache rtr 160 2w review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Apache RTR 160 2W (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">TVS Apache RTR 160 2W রিভিউ বাংলাদেশ 2024 - পারফরম্যান্স স্ট্রিট বাইক</h1>
<p class="summary-text">TVS Apache RTR 160 একটি স্পোর্টি 160cc কর্মক্ষমতা বাইক যা প্রতিক্রিয়াশীল হ্যান্ডলিং, আক্রমণাত্মক ডিজাইন, টুইন-স্পার্ক প্রযুক্তি এবং চমৎকার চালনা ক্ষমতা সহ আসে। ৳2,80,000 টাকায় মূল্যায়িত, এটি নির্ভরযোগ্য TVS প্রযুক্তি সহ রোমাঞ্চকর স্ট্রিট কর্মক্ষমতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 160 2W মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc টুইন-স্পার্ক:</strong> <span class="highlight-value">প্রতিক্রিয়াশীল কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="highlight-value">স্ট্রিট ফাইটার স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="highlight-value">চমৎকার চালনা ক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">45-50 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 160 2W সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রতিক্রিয়াশীল ইঞ্জিন:</strong> <span class="pro-description">160cc টুইন-স্পার্ক পাঞ্চ</span></li>
<li class="pro-item"><strong class="pro-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="pro-description">চটপটে এবং চমৎকার</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">স্ট্রিট ফাইটার ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">TVS নির্ভরযোগ্যতা:</strong> <span class="pro-description">বিশ্বস্ত কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মূল্য:</strong> <span class="pro-description">সাশ্রয়ী কর্মক্ষমতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 160 2W অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">দীর্ঘ যাত্রার জন্য কঠিন আসন</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">মধ্যপন্থী রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">উচ্চ rpm-এ লক্ষণীয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: TVS Apache RTR 160 2W কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,80,000 টাকায়, TVS Apache RTR 160 কর্মক্ষমতা খোঁজারদের জন্য নিখুঁত যারা প্রতিক্রিয়াশীল হ্যান্ডলিং, আক্রমণাত্মক স্টাইলিং এবং সাশ্রয়ী স্ট্রিট কর্মক্ষমতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্ট্রিট কর্মক্ষমতা রাইডারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Apache RTR 160 2W already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Apache RTR 160 2W\n")

	return nil
}

// GetName returns the seeder name
func (s *TVSApacheRTR160TwoWheelerReview) GetName() string {
	return "TVSApacheRTR160TwoWheelerReview"
}
