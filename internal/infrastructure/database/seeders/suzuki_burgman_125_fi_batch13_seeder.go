package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiBurgman125FiReview handles seeding Suzuki Burgman 125 Fi product review and translation
type SuzukiBurgman125FiReview struct {
	BaseSeeder
}

// NewSuzukiBurgman125FiReviewSeeder creates a new SuzukiBurgman125FiReview
func NewSuzukiBurgman125FiReviewSeeder() *SuzukiBurgman125FiReview {
	return &SuzukiBurgman125FiReview{BaseSeeder: BaseSeeder{name: "Suzuki Burgman 125 Fi Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiBurgman125FiReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Burgman 125 FI Ride Connect").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki burgman 125 fi ride connect product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding suzuki burgman 125 fi ride connect product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Burgman 125 Fi already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Burgman 125 Fi Ride Connect Review Bangladesh 2024 - Smart Scooter</h1>
<p class="summary-text">The Suzuki Burgman 125 Fi is a sophisticated scooter featuring fuel injection technology, Ride Connect app integration, automatic transmission, and stylish design. Priced around ৳2,40,000, it offers premium scooter features with modern connectivity for urban commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Burgman 125 Fi Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection:</strong> <span class="highlight-value">Modern technology for efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Ride Connect App:</strong> <span class="highlight-value">Smart connectivity integration</span></li>
<li class="highlight-item"><strong class="highlight-label">Automatic Transmission:</strong> <span class="highlight-value">Easy hassle-free riding</span></li>
<li class="highlight-item"><strong class="highlight-label">Stylish Design:</strong> <span class="highlight-value">Premium scooter aesthetics</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Burgman 125 Fi Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Fuel Injection Tech:</strong> <span class="pro-description">Modern technology for better efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Smart Connectivity:</strong> <span class="pro-description">Ride Connect app integration</span></li>
<li class="pro-item"><strong class="pro-label">Automatic Convenience:</strong> <span class="pro-description">No clutch hassle</span></li>
<li class="pro-item"><strong class="pro-label">Premium Design:</strong> <span class="pro-description">Stylish and sophisticated look</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">50-55 km/l excellent efficiency</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Burgman 125 Fi Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">Expensive scooter option</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">FI and connectivity service expensive</span></li>
<li class="con-item"><strong class="con-label">Fuel Tank Size:</strong> <span class="con-description">Smaller capacity tank</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Burgman 125 Fi Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,40,000 - Premium smart scooter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳4-5 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳4-5 per km (50-55 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Burgman 125 Fi Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Modern FI and connectivity</span></li>
<li class="rating-item"><strong class="rating-label">Convenience:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Automatic and app integration</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Premium styling</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good 50-55 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Premium features value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Burgman 125 Fi?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,40,000, the Suzuki Burgman 125 Fi is perfect for tech-savvy urban commuters wanting premium features, smart connectivity, and effortless automatic riding with modern fuel injection efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For tech-oriented urban commuters</span></p>
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
		return fmt.Errorf("error creating suzuki burgman 125 fi review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Burgman 125 Fi (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি Burgman 125 Fi Ride Connect রিভিউ বাংলাদেশ 2024 - স্মার্ট স্কুটার</h1>
<p class="summary-text">সুজুকি Burgman 125 Fi একটি পরিশীলিত স্কুটার যা জ্বালানি ইঞ্জেকশন প্রযুক্তি, Ride Connect অ্যাপ ইন্টিগ্রেশন, স্বয়ংক্রিয় ট্রান্সমিশন এবং স্টাইলিশ ডিজাইন বৈশিষ্ট্যযুক্ত। ৳2,40,000 টাকায় মূল্যায়িত, এটি শহুরে কমিউটারদের জন্য আধুনিক সংযোগ সহ প্রিমিয়াম স্কুটার বৈশিষ্ট্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি Burgman 125 Fi মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">জ্বালানি ইঞ্জেকশন:</strong> <span class="highlight-value">দক্ষতার জন্য আধুনিক প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">Ride Connect অ্যাপ:</strong> <span class="highlight-value">স্মার্ট সংযোগ ইন্টিগ্রেশন</span></li>
<li class="highlight-item"><strong class="highlight-label">স্বয়ংক্রিয় ট্রান্সমিশন:</strong> <span class="highlight-value">সহজ ঝামেলামুক্ত যাত্রা</span></li>
<li class="highlight-item"><strong class="highlight-label">স্টাইলিশ ডিজাইন:</strong> <span class="highlight-value">প্রিমিয়াম স্কুটার নন্দনতত্ত্ব</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি Burgman 125 Fi সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">জ্বালানি ইঞ্জেকশন প্রযুক্তি:</strong> <span class="pro-description">ভাল দক্ষতার জন্য আধুনিক প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">স্মার্ট সংযোগ:</strong> <span class="pro-description">Ride Connect অ্যাপ ইন্টিগ্রেশন</span></li>
<li class="pro-item"><strong class="pro-label">স্বয়ংক্রিয় সুবিধা:</strong> <span class="pro-description">কোন ক্লাচ ঝামেলা নেই</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম ডিজাইন:</strong> <span class="pro-description">স্টাইলিশ এবং পরিশীলিত চেহারা</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">50-55 km/l চমৎকার দক্ষতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি Burgman 125 Fi অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">ব্যয়বহুল স্কুটার বিকল্প</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">FI এবং সংযোগ সেবা ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি ট্যাঙ্ক আকার:</strong> <span class="con-description">ছোট ক্ষমতা ট্যাঙ্ক</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি Burgman 125 Fi কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,40,000 টাকায়, সুজুকি Burgman 125 Fi প্রযুক্তি-সচেতন শহুরে কমিউটারদের জন্য নিখুঁত যারা প্রিমিয়াম বৈশিষ্ট্য, স্মার্ট সংযোগ এবং আধুনিক জ্বালানি ইঞ্জেকশন দক্ষতা সহ অনায়াসী স্বয়ংক্রিয় যাত্রা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রযুক্তি-উন্মুখ শহুরে কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Burgman 125 Fi already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Burgman 125 Fi\n")

	return nil
}

// GetName returns the seeder name
func (s *SuzukiBurgman125FiReview) GetName() string {
	return "SuzukiBurgman125FiReview"
}
