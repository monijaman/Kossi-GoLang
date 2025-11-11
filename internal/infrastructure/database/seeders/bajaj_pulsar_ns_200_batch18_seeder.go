package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarNS200ReviewBatch18 handles seeding Bajaj Pulsar NS 200 product review and translation
type BajajPulsarNS200ReviewBatch18 struct {
	BaseSeeder
}

// NewBajajPulsarNS200ReviewBatch18Seeder creates a new BajajPulsarNS200ReviewBatch18
func NewBajajPulsarNS200ReviewBatch18Seeder() *BajajPulsarNS200ReviewBatch18 {
	return &BajajPulsarNS200ReviewBatch18{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar NS 200 Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarNS200ReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar NS 200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar ns 200 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj pulsar ns 200 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar NS 200 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar NS 200 Review Bangladesh 2024 - Aggressive Sport Commuter Performance</h1>
<p class="summary-text">The Bajaj Pulsar NS 200 is a 200cc liquid-cooled sports commuter combining aggressive styling with performance capability. Priced around ৳3,25,000, it delivers the famous Pulsar DNA, sharp handling, strong torque delivery, sporty ergonomics, and the aggressive naked-sport character for enthusiasts seeking performance-oriented daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar NS 200 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc Liquid-Cooled:</strong> <span class="highlight-value">Sports performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Pulsar DNA:</strong> <span class="highlight-value">Iconic performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Naked sport character</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficient:</strong> <span class="highlight-value">28-30 km/l good</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar NS 200 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Performance:</strong> <span class="pro-description">200cc liquid-cooled</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Cornering precise</span></li>
<li class="pro-item"><strong class="pro-label">Torque Delivery:</strong> <span class="pro-description">18.3 nm strong</span></li>
<li class="pro-item"><strong class="pro-label">Sport Character:</strong> <span class="pro-description">Naked aggressive</span></li>
<li class="pro-item"><strong class="pro-label">Pulsar Legacy:</strong> <span class="pro-description">Proven reliability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar NS 200 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-focused seating</span></li>
<li class="con-item"><strong class="con-label">Fuel Tank Small:</strong> <span class="con-description">12L capacity</span></li>
<li class="con-item"><strong class="con-label">Wind Protection:</strong> <span class="con-description">Naked sport limited</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar NS 200 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,25,000 - Sports commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳9-11 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">28-30 km/l good</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">200cc liquid-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">20.8 bhp sports tuned</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">160kg agile commuter</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar NS 200 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 200cc sports tuned</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Sharp precise</span></li>
<li class="rating-item"><strong class="rating-label">Torque Character:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- 18.3 nm strong</span></li>
<li class="rating-item"><strong class="rating-label">Sport Appeal:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Aggressive styling</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- ৳3,25,000 solid</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar NS 200?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,25,000, the Pulsar NS 200 is ideal for sports commuters seeking aggressive styling, strong 200cc performance, sharp handling, the famous Pulsar DNA, and proven reliability for performance-oriented daily riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sports commuters</span></p>
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
		return fmt.Errorf("error creating bajaj pulsar ns 200 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar NS 200 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার NS 200 রিভিউ বাংলাদেশ 2024 - আক্রমণাত্মক স্পোর্ট কমিউটার পারফরম্যান্স</h1>
<p class="summary-text">বাজাজ পালসার NS 200 একটি 200cc তরল-শীতলকৃত স্পোর্টস কমিউটার যা আক্রমণাত্মক স্টাইলিংকে পারফরম্যান্স ক্ষমতার সাথে একত্রিত করে। ৳3,25,000 টাকায় মূল্যায়িত, এটি বিখ্যাত পালসার DNA, তীক্ষ্ণ হ্যান্ডলিং এবং শক্তিশালী টর্ক ডেলিভারি প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার NS 200 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc তরল-শীতলকৃত:</strong> <span class="highlight-value">স্পোর্টস পারফরম্যান্স</span></li>
<li class="highlight-item"><strong class="highlight-label">পালসার DNA:</strong> <span class="highlight-value">আইকনিক পারফরম্যান্স</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="highlight-value">নেকেড স্পোর্ট চরিত্র</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষ:</strong> <span class="highlight-value">28-30 km/l ভালো</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার NS 200 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী পারফরম্যান্স:</strong> <span class="pro-description">200cc তরল-শীতলকৃত</span></li>
<li class="pro-item"><strong class="pro-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="pro-description">কর্নারিং নির্ভুল</span></li>
<li class="pro-item"><strong class="pro-label">টর্ক ডেলিভারি:</strong> <span class="pro-description">18.3 nm শক্তিশালী</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্ট চরিত্র:</strong> <span class="pro-description">নেকেড আক্রমণাত্মক</span></li>
<li class="pro-item"><strong class="pro-label">পালসার লিগেসি:</strong> <span class="pro-description">প্রমাণিত নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার NS 200 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সিট আরাম:</strong> <span class="con-description">স্পোর্ট-ফোকাসড সিটিং</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি ট্যাংক ছোট:</strong> <span class="con-description">12L ক্ষমতা</span></li>
<li class="con-item"><strong class="con-label">বায়ু সুরক্ষা:</strong> <span class="con-description">নেকেড স্পোর্ট সীমিত</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ পালসার NS 200 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,25,000 টাকায়, পালসার NS 200 স্পোর্টস কমিউটারদের জন্য আদর্শ যারা আক্রমণাত্মক স্টাইলিং, শক্তিশালী 200cc পারফরম্যান্স এবং বিখ্যাত পালসার DNA চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্পোর্টস কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar NS 200 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar NS 200\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsarNS200ReviewBatch18) GetName() string {
	return "BajajPulsarNS200ReviewBatch18"
}
