package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaFZ300ReviewBatch18 handles seeding Yamaha FZ 300 product review and translation
type YamahaFZ300ReviewBatch18 struct {
	BaseSeeder
}

// NewYamahaFZ300ReviewBatch18Seeder creates a new YamahaFZ300ReviewBatch18
func NewYamahaFZ300ReviewBatch18Seeder() *YamahaFZ300ReviewBatch18 {
	return &YamahaFZ300ReviewBatch18{BaseSeeder: BaseSeeder{name: "Yamaha FZ 300 Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaFZ300ReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha FZ 300").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha fz 300 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding yamaha fz 300 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha FZ 300 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha FZ 300 Review Bangladesh 2024 - Street Naked Premium Sports Commuter</h1>
<p class="summary-text">The Yamaha FZ 300 is a 300cc liquid-cooled street naked combining premium positioning with accessible pricing. Priced around ৳4,75,000, it delivers sophisticated design, refined performance, excellent fuel efficiency, sophisticated braking, and the Yamaha quality promise for enthusiasts seeking premium sports commuting with Japanese reliability and refinement.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha FZ 300 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">300cc Liquid-Cooled:</strong> <span class="highlight-value">Premium sports power</span></li>
<li class="highlight-item"><strong class="highlight-label">Street Naked Design:</strong> <span class="highlight-value">Sophisticated styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficiency:</strong> <span class="highlight-value">27-30 km/l excellent</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Feel:</strong> <span class="highlight-value">Yamaha quality</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha FZ 300 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Premium Design:</strong> <span class="pro-description">Sophisticated styling</span></li>
<li class="pro-item"><strong class="pro-label">Refined Engine:</strong> <span class="pro-description">Smooth 300cc</span></li>
<li class="pro-item"><strong class="pro-label">Braking System:</strong> <span class="pro-description">Advanced ABS</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Efficiency:</strong> <span class="pro-description">27-30 km/l excellent</span></li>
<li class="pro-item"><strong class="pro-label">Japanese Quality:</strong> <span class="pro-description">Yamaha reliability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha FZ 300 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">৳4,75,000 investment</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium service pricing</span></li>
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-focused seating</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha FZ 300 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,75,000 - Premium sports</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳11-13 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">27-30 km/l excellent</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">300cc liquid-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">26 bhp refined sports</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">159kg balanced</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha FZ 300 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Design Premium:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Sophisticated styling</span></li>
<li class="rating-item"><strong class="rating-label">Engine Refinement:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Smooth 300cc</span></li>
<li class="rating-item"><strong class="rating-label">Braking Safety:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Advanced ABS</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- 27-30 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- ৳4,75,000 premium</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha FZ 300?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,75,000, the FZ 300 is ideal for premium sports commuters seeking sophisticated street naked design, refined 300cc performance, excellent fuel efficiency, advanced ABS braking, and the assurance of Yamaha quality and Japanese reliability for discerning riders.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium sports commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.7,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating yamaha fz 300 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha FZ 300 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা FZ 300 রিভিউ বাংলাদেশ 2024 - স্ট্রিট নেকেড প্রিমিয়াম স্পোর্টস কমিউটার</h1>
<p class="summary-text">ইয়ামাহা FZ 300 একটি 300cc তরল-শীতলকৃত স্ট্রিট নেকেড যা প্রিমিয়াম পজিশনিংকে সাশ্রয়ী মূল্যের সাথে একত্রিত করে। ৳4,75,000 টাকায় মূল্যায়িত, এটি পরিশীলিত ডিজাইন এবং পরিমার্জিত পারফরম্যান্স প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা FZ 300 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">300cc তরল-শীতলকৃত:</strong> <span class="highlight-value">প্রিমিয়াম স্পোর্টস শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">স্ট্রিট নেকেড ডিজাইন:</strong> <span class="highlight-value">পরিশীলিত স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষতা:</strong> <span class="highlight-value">27-30 km/l চমৎকার</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম অনুভূতি:</strong> <span class="highlight-value">ইয়ামাহা গুণমান</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা FZ 300 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম ডিজাইন:</strong> <span class="pro-description">পরিশীলিত স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">পরিমার্জিত ইঞ্জিন:</strong> <span class="pro-description">মসৃণ 300cc</span></li>
<li class="pro-item"><strong class="pro-label">ব্রেকিং সিস্টেম:</strong> <span class="pro-description">উন্নত ABS</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি দক্ষতা:</strong> <span class="pro-description">27-30 km/l চমৎকার</span></li>
<li class="pro-item"><strong class="pro-label">জাপানি গুণমান:</strong> <span class="pro-description">ইয়ামাহা নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা FZ 300 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳4,75,000 বিনিয়োগ</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম সেবা মূল্য</span></li>
<li class="con-item"><strong class="con-label">সিট আরাম:</strong> <span class="con-description">স্পোর্ট-ফোকাসড সিটিং</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা FZ 300 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳4,75,000 টাকায়, FZ 300 প্রিমিয়াম স্পোর্টস কমিউটারদের জন্য আদর্শ যারা পরিশীলিত স্ট্রিট নেকেড ডিজাইন, পরিমার্জিত পারফরম্যান্স এবং ইয়ামাহা গুণমান চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম স্পোর্টস কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha FZ 300 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha FZ 300\n")

	return nil
}

// GetName returns the seeder name
func (s *YamahaFZ300ReviewBatch18) GetName() string {
	return "YamahaFZ300ReviewBatch18"
}
