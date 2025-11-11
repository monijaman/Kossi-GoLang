package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// VespaLX125Review handles seeding Vespa LX125 product review and translation
type VespaLX125Review struct {
	BaseSeeder
}

// NewVespaLX125ReviewSeeder creates a new VespaLX125Review
func NewVespaLX125ReviewSeeder() *VespaLX125Review {
	return &VespaLX125Review{BaseSeeder: BaseSeeder{name: "Vespa LX125 Review"}}
}

// Seed implements the Seeder interface
func (s *VespaLX125Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Vespa LX125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("vespa lx125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding vespa lx125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Vespa LX125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Vespa LX125 Review Bangladesh 2024 - Iconic Italian Legend Scooter</h1>
<p class="summary-text">The Vespa LX125 is a 125cc iconic scooter with timeless retro design combining vintage charm with modern technology. Priced around ৳3,20,000, it delivers unmistakable Italian heritage, iconic styling, reliable performance, and unmatched street presence for fashion-forward urban riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Vespa LX125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc Legend:</strong> <span class="highlight-value">70+ years icon</span></li>
<li class="highlight-item"><strong class="highlight-label">Retro Design:</strong> <span class="highlight-value">Timeless beauty</span></li>
<li class="highlight-item"><strong class="highlight-label">Italian Heritage:</strong> <span class="highlight-value">Authentic charm</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Tech:</strong> <span class="highlight-value">Updated reliability</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Vespa LX125 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Iconic Design:</strong> <span class="pro-description">World-famous look</span></li>
<li class="pro-item"><strong class="pro-label">Premium Build:</strong> <span class="pro-description">Italian quality</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seat:</strong> <span class="pro-description">Spacious riding</span></li>
<li class="pro-item"><strong class="pro-label">Great Handling:</strong> <span class="pro-description">Nimble urban</span></li>
<li class="pro-item"><strong class="pro-label">Resale Value:</strong> <span class="pro-description">Iconic appeal</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Vespa LX125 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price:</strong> <span class="con-description">৳3,20,000 expensive</span></li>
<li class="con-item"><strong class="con-label">Service Network:</strong> <span class="con-description">Limited availability</span></li>
<li class="con-item"><strong class="con-label">Speed Limits:</strong> <span class="con-description">Not highway ready</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Vespa LX125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,20,000 - Premium classic</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">35-38 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Vespa LX125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Smooth cruiser</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Urban agile</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Roomy seat</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Iconic legend</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Premium cost</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Vespa LX125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,20,000, the Vespa LX125 is for riders valuing iconic retro design, premium Italian heritage, timeless style, and unforgettable street presence over pure performance, making every ride a fashion statement.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For retro enthusiasts and style icons</span></p>
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
		return fmt.Errorf("error creating vespa lx125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Vespa LX125 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ভেসপা LX125 রিভিউ বাংলাদেশ 2024 - আইকনিক ইতালিয়ান লেজেন্ড স্কুটার</h1>
<p class="summary-text">ভেসপা LX125 একটি 125cc আইকনিক স্কুটার যা ভিন্টেজ চার্ম এবং আধুনিক প্রযুক্তি একত্রিত করে কালজয়ী রেট্রো ডিজাইনে। ৳3,20,000 টাকায় মূল্যায়িত, এটি অস্পষ্ট ইতালিয়ান ঐতিহ্য, আইকনিক স্টাইলিং, নির্ভরযোগ্য কর্মক্ষমতা এবং অতুলনীয় রাস্তার উপস্থিতি প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ভেসপা LX125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc লেজেন্ড:</strong> <span class="highlight-value">70+ বছরের আইকন</span></li>
<li class="highlight-item"><strong class="highlight-label">রেট্রো ডিজাইন:</strong> <span class="highlight-value">কালজয়ী সৌন্দর্য</span></li>
<li class="highlight-item"><strong class="highlight-label">ইতালিয়ান ঐতিহ্য:</strong> <span class="highlight-value">খাঁটি আকর্ষণ</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক প্রযুক্তি:</strong> <span class="highlight-value">আপডেট নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ভেসপা LX125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আইকনিক ডিজাইন:</strong> <span class="pro-description">বিশ্ব-বিখ্যাত চেহারা</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বিল্ড:</strong> <span class="pro-description">ইতালিয়ান গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক আসন:</strong> <span class="pro-description">প্রশস্ত চড়া</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত হ্যান্ডলিং:</strong> <span class="pro-description">চটপটে শহুরে</span></li>
<li class="pro-item"><strong class="pro-label">রিসেল মান:</strong> <span class="pro-description">আইকনিক আবেদন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ভেসপা LX125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">৳3,20,000 ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সেবা নেটওয়ার্ক:</strong> <span class="con-description">সীমিত উপলব্ধতা</span></li>
<li class="con-item"><strong class="con-label">গতির সীমা:</strong> <span class="con-description">হাইওয়ে প্রস্তুত নয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ভেসপা LX125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,20,000 টাকায়, ভেসপা LX125 চালকদের জন্য যারা আইকনিক রেট্রো ডিজাইন, প্রিমিয়াম ইতালিয়ান ঐতিহ্য এবং কালজয়ী স্টাইল মূল্যবান করেন খাঁটি অভিজ্ঞতার জন্য।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- রেট্রো উত্সাহী এবং স্টাইল আইকন</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Vespa LX125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Vespa LX125\n")

	return nil
}

// GetName returns the seeder name
func (s *VespaLX125Review) GetName() string {
	return "VespaLX125Review"
}
