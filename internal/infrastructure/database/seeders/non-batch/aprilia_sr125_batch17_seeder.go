package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// AppriliaSR125Review handles seeding Aprilia SR125 product review and translation
type AppriliaSR125Review struct {
	BaseSeeder
}

// NewAppriliaSR125ReviewSeeder creates a new AppriliaSR125Review
func NewAppriliaSR125ReviewSeeder() *AppriliaSR125Review {
	return &AppriliaSR125Review{BaseSeeder: BaseSeeder{name: "Aprilia SR125 Review"}}
}

// Seed implements the Seeder interface
func (s *AppriliaSR125Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Aprilia SR125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("aprilia sr125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding aprilia sr125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Aprilia SR125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Aprilia SR125 Review Bangladesh 2024 - Italian Scooter Excellence</h1>
<p class="summary-text">The Aprilia SR125 is a 125cc premium scooter combining Italian design heritage with modern convenience and aggressive styling. Priced around ৳2,95,000, it delivers sophisticated performance, comfortable ride, and distinctive European flair for urban commuters seeking style and practicality.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Aprilia SR125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc Premium:</strong> <span class="highlight-value">Italian engineering</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Design:</strong> <span class="highlight-value">Sporty aesthetic</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Ride:</strong> <span class="highlight-value">Urban ready</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Convenience:</strong> <span class="highlight-value">Premium features</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Aprilia SR125 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Italian Brand:</strong> <span class="pro-description">Premium appeal</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Head-turning design</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seat:</strong> <span class="pro-description">Long rides</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">Digital display</span></li>
<li class="pro-item"><strong class="pro-label">Good Handling:</strong> <span class="pro-description">Urban agility</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Aprilia SR125 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">৳2,95,000 expensive</span></li>
<li class="con-item"><strong class="con-label">Service Network:</strong> <span class="con-description">Limited dealers</span></li>
<li class="con-item"><strong class="con-label">Fuel Tank:</strong> <span class="con-description">Regular refills</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Aprilia SR125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,95,000 - Premium scooter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">35-40 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Aprilia SR125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Smooth power</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Urban agile</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Spacious seat</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Unique style</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Premium cost</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Aprilia SR125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,95,000, the SR125 is ideal for urban commuters wanting a premium Italian scooter with aggressive styling, comfortable ergonomics, and sophisticated European charm for daily city riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For style-conscious scooter users</span></p>
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
		return fmt.Errorf("error creating aprilia sr125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Aprilia SR125 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">অ্যাপ্রিলিয়া SR125 রিভিউ বাংলাদেশ 2024 - ইতালিয়ান স্কুটার এক্সিলেন্স</h1>
<p class="summary-text">অ্যাপ্রিলিয়া SR125 একটি 125cc প্রিমিয়াম স্কুটার যা ইতালিয়ান ডিজাইন ঐতিহ্যকে আধুনিক সুবিধা এবং আক্রমণাত্মক স্টাইলিংয়ের সাথে একত্রিত করে। ৳2,95,000 টাকায় মূল্যায়িত, এটি পরিশীলিত কর্মক্ষমতা, আরামদায়ক যাত্রা এবং স্বতন্ত্র ইউরোপীয় ফ্লেয়ার প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">অ্যাপ্রিলিয়া SR125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">125cc প্রিমিয়াম:</strong> <span class="highlight-value">ইতালিয়ান প্রকৌশল</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="highlight-value">খেলাধুলামূলক নান্দনিকতা</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক যাত্রা:</strong> <span class="highlight-value">শহুরে প্রস্তুত</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক সুবিধা:</strong> <span class="highlight-value">প্রিমিয়াম বৈশিষ্ট্য</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">অ্যাপ্রিলিয়া SR125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ইতালিয়ান ব্র্যান্ড:</strong> <span class="pro-description">প্রিমিয়াম আবেদন</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">মাথা-টার্নার ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক আসন:</strong> <span class="pro-description">দীর্ঘ যাত্রা</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক বৈশিষ্ট্য:</strong> <span class="pro-description">ডিজিটাল প্রদর্শন</span></li>
<li class="pro-item"><strong class="pro-label">ভাল হ্যান্ডলিং:</strong> <span class="pro-description">শহুরে চলনশীলতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">অ্যাপ্রিলিয়া SR125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳2,95,000 ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সেবা নেটওয়ার্ক:</strong> <span class="con-description">সীমিত ডিলার</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি ট্যাঙ্ক:</strong> <span class="con-description">নিয়মিত রিফিল</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: অ্যাপ্রিলিয়া SR125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,95,000 টাকায়, SR125 শহুরে যাত্রীদের জন্য আদর্শ যারা একটি প্রিমিয়াম ইতালিয়ান স্কুটার চান যা আক্রমণাত্মক স্টাইলিং এবং আরামদায়ক ergonomics সহ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্টাইল-সচেতন স্কুটার ব্যবহারকারীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Aprilia SR125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Aprilia SR125\n")

	return nil
}

// GetName returns the seeder name
func (s *AppriliaSR125Review) GetName() string {
	return "AppriliaSR125Review"
}
