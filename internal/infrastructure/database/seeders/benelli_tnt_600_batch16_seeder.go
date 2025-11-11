package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BenelliTNT600ReviewBatch16 handles seeding Benelli TNT 600 product review and translation
type BenelliTNT600ReviewBatch16 struct {
	BaseSeeder
}

// NewBenelliTNT600ReviewBatch16Seeder creates a new BenelliTNT600ReviewBatch16
func NewBenelliTNT600ReviewBatch16Seeder() *BenelliTNT600ReviewBatch16 {
	return &BenelliTNT600ReviewBatch16{BaseSeeder: BaseSeeder{name: "Benelli TNT 600 Batch16 Review"}}
}

// Seed implements the Seeder interface
func (s *BenelliTNT600ReviewBatch16) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Benelli TNT 600").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("benelli tnt 600 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding benelli tnt 600 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Benelli TNT 600 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Benelli TNT 600 Review Bangladesh 2024 - Italian Sports Fury</h1>
<p class="summary-text">The Benelli TNT 600 is a 600cc naked sportbike combining Italian heritage with aggressive styling and thrilling performance. Priced around ৳6,20,000, it delivers track-focused dynamics, premium build quality, and European engineering for riders seeking a truly distinctive middleweight sports motorcycle experience.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Benelli TNT 600 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">600cc Naked Beast:</strong> <span class="highlight-value">Italian aggression</span></li>
<li class="highlight-item"><strong class="highlight-label">Track Dynamics:</strong> <span class="highlight-value">Sporty character</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Build:</strong> <span class="highlight-value">European quality</span></li>
<li class="highlight-item"><strong class="highlight-label">Distinctive Style:</strong> <span class="highlight-value">Unique appeal</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Benelli TNT 600 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">600cc thrills</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Head-turner design</span></li>
<li class="pro-item"><strong class="pro-label">Track-Ready:</strong> <span class="pro-description">Sporty handling</span></li>
<li class="pro-item"><strong class="pro-label">Premium Quality:</strong> <span class="pro-description">European feel</span></li>
<li class="pro-item"><strong class="pro-label">Exclusive Brand:</strong> <span class="pro-description">Unique identity</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Benelli TNT 600 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Few authorized dealers</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">20-24 km/l moderate</span></li>
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">Premium positioning</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Benelli TNT 600 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳6,20,000 - Premium middleweight</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳10-14 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">20-24 km/l decent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Benelli TNT 600 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Potent 600cc</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Precise control</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Premium European</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Distinctive style</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Premium bike</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Benelli TNT 600?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳6,20,000, the TNT 600 is ideal for riders wanting a distinctive premium 600cc naked sport with Italian aggression, track-ready dynamics, and European quality that stands out from mainstream alternatives.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For discerning sports enthusiasts</span></p>
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
		return fmt.Errorf("error creating benelli tnt 600 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Benelli TNT 600 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বেনেলি TNT 600 রিভিউ বাংলাদেশ 2024 - ইতালিয়ান স্পোর্টস ফিউরি</h1>
<p class="summary-text">বেনেলি TNT 600 একটি 600cc নেকেড স্পোর্টবাইক যা ইতালিয়ান ঐতিহ্যকে আক্রমণাত্মক স্টাইলিং এবং রোমাঞ্চকর কর্মক্ষমতার সাথে একত্রিত করে। ৳6,20,000 টাকায় মূল্যায়িত, এটি ট্র্যাক-ফোকাসড গতিশীলতা, প্রিমিয়াম বিল্ড কোয়ালিটি এবং ইউরোপীয় প্রকৌশল প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বেনেলি TNT 600 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">600cc নেকেড বিস্ট:</strong> <span class="highlight-value">ইতালিয়ান আক্রমণ</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্র্যাক গতিশীলতা:</strong> <span class="highlight-value">খেলাধুলামূলক চরিত্র</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম বিল্ড:</strong> <span class="highlight-value">ইউরোপীয় গুণমান</span></li>
<li class="highlight-item"><strong class="highlight-label">স্বতন্ত্র স্টাইল:</strong> <span class="highlight-value">অনন্য আবেদন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বেনেলি TNT 600 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">600cc রোমাঞ্চ</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">মাথা-টার্নার ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">ট্র্যাক-প্রস্তুত:</strong> <span class="pro-description">খেলাধুলামূলক হ্যান্ডলিং</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম গুণমান:</strong> <span class="pro-description">ইউরোপীয় অনুভূতি</span></li>
<li class="pro-item"><strong class="pro-label">একচেটিয়া ব্র্যান্ড:</strong> <span class="pro-description">অনন্য পরিচয়</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বেনেলি TNT 600 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত সেবা:</strong> <span class="con-description">কয়েকটি অনুমোদিত ডিলার</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">20-24 km/l মধ্যমা</span></li>
<li class="con-item"><strong class="con-label">উচ্চতর মূল্য:</strong> <span class="con-description">প্রিমিয়াম অবস্থান</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বেনেলি TNT 600 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳6,20,000 টাকায়, TNT 600 চালকদের জন্য আদর্শ যারা একটি স্বতন্ত্র প্রিমিয়াম 600cc নেকেড স্পোর্ট চান যা ট্র্যাক-প্রস্তুত গতিশীলতা এবং ইউরোপীয় গুণমান সহ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বিচক্ষণ ক্রীড়া উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Benelli TNT 600 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Benelli TNT 600\n")

	return nil
}

// GetName returns the seeder name
func (s *BenelliTNT600ReviewBatch16) GetName() string {
	return "BenelliTNT600ReviewBatch16"
}
