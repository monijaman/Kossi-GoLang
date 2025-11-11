package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// PiaggioBeverlyReview handles seeding Piaggio Beverly product review and translation
type PiaggioBeverlyReview struct {
	BaseSeeder
}

// NewPiaggioBeverlyReviewSeeder creates a new PiaggioBeverlyReview
func NewPiaggioBeverlyReviewSeeder() *PiaggioBeverlyReview {
	return &PiaggioBeverlyReview{BaseSeeder: BaseSeeder{name: "Piaggio Beverly Review"}}
}

// Seed implements the Seeder interface
func (s *PiaggioBeverlyReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Piaggio Beverly").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("piaggio beverly product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding piaggio beverly product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Piaggio Beverly already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Piaggio Beverly Review Bangladesh 2024 - Premium Italian Maxi-Scooter</h1>
<p class="summary-text">The Piaggio Beverly is a 300cc premium maxi-scooter combining Italian elegance with practical features and spacious comfort. Priced around ৳4,20,000, it delivers sophisticated styling, comfortable two-seater design, powerful performance, and premium feel for discerning riders valuing style and utility.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Piaggio Beverly Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">300cc Maxi:</strong> <span class="highlight-value">Premium scooter</span></li>
<li class="highlight-item"><strong class="highlight-label">Italian Design:</strong> <span class="highlight-value">Elegant styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Spacious Comfort:</strong> <span class="highlight-value">Two-seater ease</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Features:</strong> <span class="highlight-value">Tech equipped</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Piaggio Beverly Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Italian Elegance:</strong> <span class="pro-description">Premium appeal</span></li>
<li class="pro-item"><strong class="pro-label">Spacious Seat:</strong> <span class="pro-description">Comfort focus</span></li>
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">300cc strong</span></li>
<li class="pro-item"><strong class="pro-label">Storage:</strong> <span class="pro-description">Practical cabin</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Ride:</strong> <span class="pro-description">Highway capable</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Piaggio Beverly Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price:</strong> <span class="con-description">৳4,20,000 premium</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Few dealers</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">Moderate efficiency</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Piaggio Beverly Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,20,000 - Premium maxi</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-10 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">28-32 km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Piaggio Beverly Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Powerful 300cc</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Spacious seating</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Italian elegant</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Stable maxi</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Premium priced</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Piaggio Beverly?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,20,000, the Beverly is perfect for riders valuing premium Italian elegance, spacious two-seater comfort, powerful 300cc performance, and practical features for stylish commuting and weekend adventures.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium scooter enthusiasts</span></p>
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
		return fmt.Errorf("error creating piaggio beverly review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Piaggio Beverly (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">পিয়াজিও বেভারলি রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম ইতালিয়ান ম্যাক্সি-স্কুটার</h1>
<p class="summary-text">পিয়াজিও বেভারলি একটি 300cc প্রিমিয়াম ম্যাক্সি-স্কুটার যা ইতালিয়ান কমনীয়তাকে ব্যবহারিক বৈশিষ্ট্য এবং প্রশস্ত আরামের সাথে একত্রিত করে। ৳4,20,000 টাকায় মূল্যায়িত, এটি পরিশীলিত স্টাইলিং, আরামদায়ক দুই-আসন ডিজাইন এবং প্রিমিয়াম অনুভূতি প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">পিয়াজিও বেভারলি মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">300cc ম্যাক্সি:</strong> <span class="highlight-value">প্রিমিয়াম স্কুটার</span></li>
<li class="highlight-item"><strong class="highlight-label">ইতালিয়ান ডিজাইন:</strong> <span class="highlight-value">মার্জিত স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রশস্ত আরাম:</strong> <span class="highlight-value">দুই-আসন সহজতা</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক বৈশিষ্ট্য:</strong> <span class="highlight-value">প্রযুক্তি সজ্জিত</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">পিয়াজিও বেভারলি সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ইতালিয়ান কমনীয়তা:</strong> <span class="pro-description">প্রিমিয়াম আবেদন</span></li>
<li class="pro-item"><strong class="pro-label">প্রশস্ত আসন:</strong> <span class="pro-description">আরাম ফোকাস</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">300cc শক্তিশালী</span></li>
<li class="pro-item"><strong class="pro-label">স্টোরেজ:</strong> <span class="pro-description">ব্যবহারিক কেবিন</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ চড়া:</strong> <span class="pro-description">হাইওয়ে সক্ষম</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">পিয়াজিও বেভারলি অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">৳4,20,000 প্রিমিয়াম</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা:</strong> <span class="con-description">কয়েকটি ডিলার</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">মধ্যম দক্ষতা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: পিয়াজিও বেভারলি কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳4,20,000 টাকায়, বেভারলি চালকদের জন্য নিখুঁত যারা প্রিমিয়াম ইতালিয়ান কমনীয়তা, প্রশস্ত দুই-আসন আরাম এবং শক্তিশালী 300cc কর্মক্ষমতা মূল্যবান করেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম স্কুটার উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Piaggio Beverly already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Piaggio Beverly\n")

	return nil
}

// GetName returns the seeder name
func (s *PiaggioBeverlyReview) GetName() string {
	return "PiaggioBeverlyReview"
}
