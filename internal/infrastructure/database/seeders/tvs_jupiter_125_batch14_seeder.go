package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSJupiter125Review handles seeding TVS Jupiter 125 product review and translation
type TVSJupiter125Review struct {
	BaseSeeder
}

// NewTVSJupiter125ReviewSeeder creates a new TVSJupiter125Review
func NewTVSJupiter125ReviewSeeder() *TVSJupiter125Review {
	return &TVSJupiter125Review{BaseSeeder: BaseSeeder{name: "TVS Jupiter 125 Review"}}
}

// Seed implements the Seeder interface
func (s *TVSJupiter125Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Jupiter 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs jupiter 125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding tvs jupiter 125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Jupiter 125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Jupiter 125 Review Bangladesh 2024 - Premium Scooter</h1>
<p class="summary-text">The TVS Jupiter 125 is a premium 125cc scooter with stylish design, comfortable seating, automatic transmission, and modern features. Priced around ৳2,35,000, it offers sophisticated urban mobility with smooth performance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Jupiter 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc Reliable Engine:</strong> <span class="highlight-value">Smooth urban performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Automatic Transmission:</strong> <span class="highlight-value">Hassle-free operation</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Design:</strong> <span class="highlight-value">Stylish appearance</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seating:</strong> <span class="highlight-value">Long ride comfort</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Jupiter 125 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Reliable Performance:</strong> <span class="pro-description">125cc smooth engine</span></li>
<li class="pro-item"><strong class="pro-label">Automatic Convenience:</strong> <span class="pro-description">No clutch hassle</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">48-52 km/l efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Stylish Design:</strong> <span class="pro-description">Premium scooter aesthetics</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seating:</strong> <span class="pro-description">Well-padded seat</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Jupiter 125 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Storage Space:</strong> <span class="con-description">Limited under-seat storage</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium maintenance</span></li>
<li class="con-item"><strong class="con-label">Resale Value:</strong> <span class="con-description">Moderate depreciation</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Jupiter 125 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,35,000 - Premium scooter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳4-5 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">48-52 km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Jupiter 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Smooth 125cc</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Premium styling</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Well-padded seating</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good 48-52 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Premium features</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Jupiter 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,35,000, the TVS Jupiter 125 is ideal for urban commuters seeking a premium scooter with stylish design, reliable performance, and comfortable automatic operation.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium scooter seekers</span></p>
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
		return fmt.Errorf("error creating tvs jupiter 125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Jupiter 125 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">TVS Jupiter 125 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম স্কুটার</h1>
<p class="summary-text">TVS Jupiter 125 একটি প্রিমিয়াম 125cc স্কুটার যা স্টাইলিশ ডিজাইন, আরামদায়ক আসন, স্বয়ংক্রিয় ট্রান্সমিশন এবং আধুনিক বৈশিষ্ট্য সহ আসে। ৳2,35,000 টাকায় মূল্যায়িত, এটি মসৃণ কর্মক্ষমতা সহ পরিশীলিত শহুরে গতিশীলতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Jupiter 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="highlight-value">মসৃণ শহুরে কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">স্বয়ংক্রিয় ট্রান্সমিশন:</strong> <span class="highlight-value">ঝামেলামুক্ত পরিচালনা</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম ডিজাইন:</strong> <span class="highlight-value">স্টাইলিশ চেহারা</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক আসন:</strong> <span class="highlight-value">দীর্ঘ যাত্রার আরাম</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Jupiter 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য কর্মক্ষমতা:</strong> <span class="pro-description">125cc মসৃণ ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">স্বয়ংক্রিয় সুবিধা:</strong> <span class="pro-description">কোন ক্লাচ ঝামেলা নেই</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">48-52 km/l দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">স্টাইলিশ ডিজাইন:</strong> <span class="pro-description">প্রিমিয়াম স্কুটার নন্দনতত্ত্ব</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক আসন:</strong> <span class="pro-description">ভালো প্যাডেড সিট</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Jupiter 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">স্টোরেজ স্পেস:</strong> <span class="con-description">সীমিত আন্ডার-সিট স্টোরেজ</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">পুনর্বিক্রয় মূল্য:</strong> <span class="con-description">মধ্যপন্থী অবমূল্যায়ন</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: TVS Jupiter 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,35,000 টাকায়, TVS Jupiter 125 শহুরে যাত্রীদের জন্য আদর্শ যারা প্রিমিয়াম স্কুটার, স্টাইলিশ ডিজাইন এবং নির্ভরযোগ্য স্বয়ংক্রিয় কর্মক্ষমতা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম স্কুটার খোঁজারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Jupiter 125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Jupiter 125\n")

	return nil
}

// GetName returns the seeder name
func (s *TVSJupiter125Review) GetName() string {
	return "TVSJupiter125Review"
}
