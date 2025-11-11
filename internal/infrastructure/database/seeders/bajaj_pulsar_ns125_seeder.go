package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarNS125Review handles seeding Bajaj Pulsar NS125 product review and translation
type BajajPulsarNS125Review struct {
	BaseSeeder
}

// NewBajajPulsarNS125ReviewSeeder creates a new BajajPulsarNS125Review
func NewBajajPulsarNS125ReviewSeeder() *BajajPulsarNS125Review {
	return &BajajPulsarNS125Review{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar NS125 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarNS125Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar NS125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Pulsar NS125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Pulsar NS125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar NS125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar NS125 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Pulsar NS125 is a lightweight, entry-level street bike featuring 125cc engine, sleek modern design, good mileage, and beginner-friendly handling. Priced at ৳1,10,000, it offers an excellent introduction to premium motorcycles with Pulsar's sporty character while remaining affordable and fuel-efficient. Perfect for city commuting and new riders seeking their first premium bike.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar NS125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Lightweight Design:</strong> <span class="highlight-value">Easy to handle for new riders</span></li>
<li class="highlight-item"><strong class="highlight-label">125cc Engine:</strong> <span class="highlight-value">Good power for city commuting</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Styling:</strong> <span class="highlight-value">Sleek Pulsar DNA design</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">50-55 km/l fuel efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar NS125 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Beginner-Friendly:</strong> <span class="pro-description">Lightweight and easy to control for new riders</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Design:</strong> <span class="pro-description">Modern Pulsar styling at affordable price</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">50-55 km/l offers outstanding fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Low Price:</strong> <span class="pro-description">Most affordable Pulsar offering at ৳1,10,000</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">Adequate power for city riding</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar NS125 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Power:</strong> <span class="con-description">125cc engine not suitable for highway riding</span></li>
<li class="con-item"><strong class="con-label">No ABS:</strong> <span class="con-description">Basic braking without ABS safety</span></li>
<li class="con-item"><strong class="con-label">Compact Seat:</strong> <span class="con-description">Less comfortable for pillion passengers</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar NS125 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,10,000 - Best value entry-level premium</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very Low - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳3-4 per km (50-55 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar NS125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Good for city commuting</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Outstanding fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent for new riders</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Modern Pulsar design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Exceptional value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar NS125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At just ৳1,10,000, the Bajaj Pulsar NS125 offers exceptional value as an entry-level premium motorcycle. Perfect for new riders, city commuters, and those seeking outstanding fuel efficiency with Pulsar's sporty character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For new riders and budget-conscious commuters</span></p>
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
		return fmt.Errorf("error creating Bajaj Pulsar NS125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar NS125 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার এনএস১২৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ পালসার এনএস১২৫ একটি হালকা ওজনের এন্ট্রি-লেভেল স্ট্রিট বাইক যা ১২৫সিসি ইঞ্জিন, আধুনিক চিকন ডিজাইন এবং ভালো মাইলেজ সহ আসে। ৳১,১০,০০০ টাকায় মূল্যায়িত, এটি নতুন রাইডারদের জন্য পালসারের স্পোর্টি চরিত্র সহ প্রিমিয়াম মোটরসাইকেলে চমৎকার প্রবেশ প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার এনএস১২৫ মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">হালকা ওজন:</strong> <span class="highlight-value">নতুন রাইডারদের জন্য সহজ নিয়ন্ত্রণ</span></li>
<li class="highlight-item"><strong class="highlight-label">১২৫সিসি ইঞ্জিন:</strong> <span class="highlight-value">শহুরে যাতায়াতের জন্য ভালো শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক স্টাইলিং:</strong> <span class="highlight-value">চিকন পালসার ডিএনএ ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">৫০-৫৫ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার এনএস১২৫ সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">নতুন রাইডার-বান্ধব:</strong> <span class="pro-description">হালকা এবং নতুনদের জন্য নিয়ন্ত্রণ করা সহজ</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি ডিজাইন:</strong> <span class="pro-description">সাশ্রয়ী মূল্যে আধুনিক পালসার স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">৫০-৫৫ কিমি/লিটার অসাধারণ জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">কম মূল্য:</strong> <span class="pro-description">সবচেয়ে সাশ্রয়ী পালসার ৳১,১০,০০০ এ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার এনএস১२५ অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত শক্তি:</strong> <span class="con-description">১२५সিসি ইঞ্জিন হাইওয়ে রাইডিংয়ের জন্য উপযুক্ত নয়</span></li>
<li class="con-item"><strong class="con-label">কোন এবিএস নেই:</strong> <span class="con-description">এবিএস নিরাপত্তা ছাড়া মৌলিক ব্রেকিং</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাজাজ পালসার এনএস१२५ কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">४.३/५.०</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">মাত্র ৳१,१०,००० টাকায় বাজাজ পালসার এনএস१२५ এন্ট্রি-লেভেল প্রিমিয়াম মোটরসাইকেল হিসাবে ব্যতিক্রমী মূল্য প্রদান করে। নতুন রাইডার এবং শহুরে যাতায়াতকারীদের জন্য নিখুঁত।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- নতুন রাইডার এবং বাজেট-সচেতন যাতায়াতকারীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar NS125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar NS125\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsarNS125Review) GetName() string {
	return "BajajPulsarNS125Review"
}
