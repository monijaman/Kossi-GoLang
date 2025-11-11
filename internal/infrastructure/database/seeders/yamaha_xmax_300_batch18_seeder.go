package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaXMax300ReviewBatch18 handles seeding Yamaha X-MAX 300 product review and translation
type YamahaXMax300ReviewBatch18 struct {
	BaseSeeder
}

// NewYamahaXMax300ReviewBatch18Seeder creates a new YamahaXMax300ReviewBatch18
func NewYamahaXMax300ReviewBatch18Seeder() *YamahaXMax300ReviewBatch18 {
	return &YamahaXMax300ReviewBatch18{BaseSeeder: BaseSeeder{name: "Yamaha X-MAX 300 Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaXMax300ReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha X-MAX 300").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha x-max 300 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding yamaha x-max 300 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha X-MAX 300 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha X-MAX 300 Review Bangladesh 2024 - Premium Maxi-Scooter Performance Machine</h1>
<p class="summary-text">The Yamaha X-MAX 300 is a 300cc premium maxi-scooter delivering urban agility, highway-capable performance, and sophisticated technology. Priced around ৳4,85,000, it combines responsive engine, advanced electronics, luxurious comfort, and distinctive styling for commuters and weekend riders seeking practical performance and premium riding experience in maxi-scooter form.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha X-MAX 300 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">300cc Maxi Beast:</strong> <span class="highlight-value">Performance scooter</span></li>
<li class="highlight-item"><strong class="highlight-label">Liquid Cooling:</strong> <span class="highlight-value">Continuous power delivery</span></li>
<li class="highlight-item"><strong class="highlight-label">Electronic Aids:</strong> <span class="highlight-value">Traction control equipped</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Design:</strong> <span class="highlight-value">Aggressive maxi styling</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha X-MAX 300 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Performance Engine:</strong> <span class="pro-description">Responsive 300cc power</span></li>
<li class="pro-item"><strong class="pro-label">Traction Control:</strong> <span class="pro-description">Safety technology</span></li>
<li class="pro-item"><strong class="pro-label">Comfort Seating:</strong> <span class="pro-description">Premium long-ride comfort</span></li>
<li class="pro-item"><strong class="pro-label">Storage:</strong> <span class="pro-description">Practical cabin space</span></li>
<li class="pro-item"><strong class="pro-label">Highway Capable:</strong> <span class="pro-description">Sustained speed comfort</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha X-MAX 300 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">৳4,85,000 expensive</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium maintenance</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">25-30 km/l moderate</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha X-MAX 300 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,85,000 - Premium performance maxi</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳9-12 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">25-30 km/l moderate</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">300cc liquid-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">28.5 bhp responsive</span></p>
<p class="value-item"><strong class="value-label">Top Speed:</strong> <span class="value-amount">180+ km/h highway capable</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha X-MAX 300 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Responsive liquid-cooled</span></li>
<li class="rating-item"><strong class="rating-label">Acceleration:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Brisk highway</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Premium luxury</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Stable maxi</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Premium positioned</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha X-MAX 300?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,85,000, the X-MAX 300 is ideal for premium maxi-scooter seekers wanting responsive 300cc performance, advanced technology, luxurious comfort, and highway-capable practicality with sophisticated Japanese engineering.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium urban/highway riders</span></p>
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
		return fmt.Errorf("error creating yamaha x-max 300 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha X-MAX 300 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা X-MAX 300 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম ম্যাক্সি-স্কুটার পারফরম্যান্স মেশিন</h1>
<p class="summary-text">ইয়ামাহা X-MAX 300 একটি 300cc প্রিমিয়াম ম্যাক্সি-স্কুটার যা শহুরে চপলতা, হাইওয়ে-সক্ষম কর্মক্ষমতা এবং পরিশীলিত প্রযুক্তি প্রদান করে। ৳4,85,000 টাকায় মূল্যায়িত, এটি প্রতিক্রিয়াশীল ইঞ্জিন, উন্নত ইলেকট্রনিক্স এবং বিলাসবহুল আরাম একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা X-MAX 300 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">300cc ম্যাক্সি বিস্ট:</strong> <span class="highlight-value">পারফরম্যান্স স্কুটার</span></li>
<li class="highlight-item"><strong class="highlight-label">তরল শীতলন:</strong> <span class="highlight-value">ক্রমাগত শক্তি ডেলিভারি</span></li>
<li class="highlight-item"><strong class="highlight-label">ইলেকট্রনিক সহায়তা:</strong> <span class="highlight-value">ট্র্যাকশন কন্ট্রোল সজ্জিত</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম ডিজাইন:</strong> <span class="highlight-value">আক্রমণাত্মক ম্যাক্সি স্টাইলিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা X-MAX 300 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">পারফরম্যান্স ইঞ্জিন:</strong> <span class="pro-description">প্রতিক্রিয়াশীল 300cc শক্তি</span></li>
<li class="pro-item"><strong class="pro-label">ট্র্যাকশন কন্ট্রোল:</strong> <span class="pro-description">নিরাপত্তা প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">আরাম আসন:</strong> <span class="pro-description">প্রিমিয়াম দীর্ঘ-যাত্রা আরাম</span></li>
<li class="pro-item"><strong class="pro-label">স্টোরেজ:</strong> <span class="pro-description">ব্যবহারিক কেবিন স্পেস</span></li>
<li class="pro-item"><strong class="pro-label">হাইওয়ে সক্ষম:</strong> <span class="pro-description">টেকসই গতি আরাম</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা X-MAX 300 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳4,85,000 ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">25-30 km/l মধ্যমা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা X-MAX 300 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳4,85,000 টাকায়, X-MAX 300 প্রিমিয়াম ম্যাক্সি-স্কুটার চাহনাকারীদের জন্য আদর্শ যারা প্রতিক্রিয়াশীল কর্মক্ষমতা, উন্নত প্রযুক্তি এবং হাইওয়ে-সক্ষম ব্যবহারিকতা মূল্যবান করেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম শহুরে/হাইওয়ে চালকদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha X-MAX 300 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha X-MAX 300\n")

	return nil
}

// GetName returns the seeder name
func (s *YamahaXMax300ReviewBatch18) GetName() string {
	return "YamahaXMax300ReviewBatch18"
}
