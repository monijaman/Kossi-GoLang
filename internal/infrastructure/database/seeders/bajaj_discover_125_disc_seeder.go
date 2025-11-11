package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajDiscover125DiscReview handles seeding Bajaj Discover 125 Disc product review and translation
type BajajDiscover125DiscReview struct {
	BaseSeeder
}

// NewBajajDiscover125DiscReviewSeeder creates a new BajajDiscover125DiscReview
func NewBajajDiscover125DiscReviewSeeder() *BajajDiscover125DiscReview {
	return &BajajDiscover125DiscReview{BaseSeeder: BaseSeeder{name: "Bajaj Discover 125 Disc Review"}}
}

// Seed implements the Seeder interface
func (s *BajajDiscover125DiscReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Discover 125 Disc").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Discover 125 Disc product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Discover 125 Disc product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Discover 125 Disc already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Discover 125 Disc Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Discover 125 Disc is a workhorse commuter motorcycle featuring 125cc engine, front disc brake for better stopping power, comfortable seating, excellent mileage, and proven reliability. Priced at ৳95,000, it remains one of the most popular budget commuters with exceptional durability, low maintenance costs, and practical design. Perfect for daily commuting and long-distance touring on a tight budget.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Discover 125 Disc Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Front Disc Brake:</strong> <span class="highlight-value">Better braking control than drum brakes</span></li>
<li class="highlight-item"><strong class="highlight-label">125cc Engine:</strong> <span class="highlight-value">Reliable and proven technology</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seat:</strong> <span class="highlight-value">Well-padded for long commutes</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">50-55 km/l outstanding efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Discover 125 Disc Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Incredibly Affordable:</strong> <span class="pro-description">One of the cheapest bikes at ৳95,000</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">50-55 km/l provides outstanding economy</span></li>
<li class="pro-item"><strong class="pro-label">Proven Reliability:</strong> <span class="pro-description">Millions sold with excellent track record</span></li>
<li class="pro-item"><strong class="pro-label">Front Disc Brake:</strong> <span class="pro-description">Better braking than standard drum brake</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Riding:</strong> <span class="pro-description">Good ergonomics for daily commuting</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Simple, durable, and cheap to maintain</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Discover 125 Disc Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Dated Design:</strong> <span class="con-description">Basic practical design, not modern styling</span></li>
<li class="con-item"><strong class="con-label">Limited Features:</strong> <span class="con-description">No ABS, no digital instruments</span></li>
<li class="con-item"><strong class="con-label">No Rear Disc:</strong> <span class="con-description">Rear drum brake limits braking performance</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Discover 125 Disc Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳95,000 - Best budget commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very Low - ৳3 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳3 per km (50-55 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Discover 125 Disc Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- Basic but adequate</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Outstanding efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Excellent proven reliability</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good for commuting</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Exceptional budget value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Discover 125 Disc?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At just ৳95,000, the Bajaj Discover 125 Disc is an absolute bargain for budget-conscious buyers. With proven reliability, outstanding mileage, and low maintenance costs, it remains a practical choice for daily commuting and long-distance touring despite its basic design.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For the most budget-friendly reliable commuter</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.2,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Bajaj Discover 125 Disc review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Discover 125 Disc (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ ডিসকভার ১२५ ডিস্ক রিভিউ বাংলাদেশ २०२४ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ ডিসকভার १२५ ডিস্ক একটি নির্ভরযোগ্য কর্মঘোড়া কমিউটার মোটরসাইকেল যা १२५সিসি ইঞ্জিন, সামনের ডিস্ক ব্রেক, আরামদায়ক সিট এবং চমৎকার মাইলেজ সহ আসে। ৳७५,००० টাকায় মূল্যায়িত, এটি সবচেয়ে জনপ্রিয় বাজেট কমিউটারগুলির মধ্যে একটি থেকে গেছে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ ডিসকভার १२५ ডিস্ক মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সামনের ডিস্ক ব্রেক:</strong> <span class="highlight-value">ড্রাম ব্রেকের চেয়ে ভাল ব্রেকিং নিয়ন্ত্রণ</span></li>
<li class="highlight-item"><strong class="highlight-label">१२५সিসি ইঞ্জিন:</strong> <span class="highlight-value">নির্ভরযোগ্য এবং প্রমাণিত প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক সিট:</strong> <span class="highlight-value">দীর্ঘ যাতায়াতের জন্য ভালো-প্যাড করা</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">५०-५५ कमी/लिटर অসাধারণ দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ ডিসকভার १२५ ডিস্ক সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অবিশ্বাস্যভাবে সাশ্রয়ী:</strong> <span class="pro-description">८५,०००ँ টাকায় সবচেয়ে সস্তা বাইকগুলির মধ্যে একটি</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">५०-५५ किमी/लिटर আউটস्टैंडिং অর্থনীতি</span></li>
<li class="pro-item"><strong class="pro-label">প্রমাণিত নির্ভরযোগ্যতা:</strong> <span class="pro-description">লক্ষ লক্ষ বিক্রি করা হয়েছে চমৎকার ট্র্যাক রেকর্ড সহ</span></li>
<li class="pro-item"><strong class="pro-label">সামনের ডিস্ক ব্রেক:</strong> <span class="pro-description">স্ট্যান্ডার্ড ড্রাম ব্রেকের চেয়ে ভাল</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ ডিসকভার १२५ ডিস্ক অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">পুরানো ডিজাইন:</strong> <span class="con-description">মৌলिक অনুশীলনী ডিজাইন, আধুनिक স্টাইলिং নয়</span></li>
<li class="con-item"><strong class="con-label">সীমিত বৈশিষ্ট্য:</strong> <span class="con-description">কোন এবিএস, কোন ডিজিটাল যন্ত্র নেই</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাজাজ ডিসকভার १२५ ডিস्क किनबेन?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ्रिक रेटिং:</strong> <span class="score-value">४.२/५.०</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">মাত্र ৳८५,००० টাকায় বাজাজ ডিসকভার १२५ ডিस्क বাজেট-সচেতন ক্রেতাদের জন্য একটি সম্পূর্ণ দর। প্রমাণিত নির্ভরযোগ্যতা এবং চমৎকার মাইলেজের সাথে, এটি দৈনন্দিন যাতায়াত এবং দীর্ঘ-দূরত্ব ট্যুরিংয়ের জন্য একটি অনুশীলনী পছন্দ থাকে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">हां</span> <span class="recommendation-for">- সবচেয়ে বাজেট-বান্ধব নির্ভরযোগ্য কমিউটারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Discover 125 Disc already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Discover 125 Disc\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajDiscover125DiscReview) GetName() string {
	return "BajajDiscover125DiscReview"
}
