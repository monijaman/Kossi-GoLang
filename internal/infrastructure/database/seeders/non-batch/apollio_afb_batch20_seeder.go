package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ApollioAFBReviewBatch20 struct {
	BaseSeeder
}

func NewApollioAFBReviewBatch20Seeder() *ApollioAFBReviewBatch20 {
	return &ApollioAFBReviewBatch20{BaseSeeder: BaseSeeder{name: "Apollio AFB Batch20 Review"}}
}

func (s *ApollioAFBReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Apollio AFB").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("apollio afb product not found")
		}
		return fmt.Errorf("error finding product: %w", err)
	}

	var existingReview models.ProductReviewModel
	if err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error; err == nil {
		fmt.Printf("   ℹ️  Review already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Apollio AFB Review Bangladesh 2024 - Stylish Urban Commuter</h1>
<p class="summary-text">The Apollio AFB is a 125cc air-cooled stylish urban commuter representing practical modern design. Priced around ৳1,95,000, it delivers modern styling, smooth operation, good fuel efficiency, user-friendly features, and attractive value for style-conscious commuters seeking contemporary practical transport.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Apollio AFB Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Air-Cooled:</strong> Urban commuter</li>
<li class="highlight-item"><strong>Modern Styling:</strong> Contemporary design</li>
<li class="highlight-item"><strong>45+ km/l:</strong> Good efficiency</li>
<li class="highlight-item"><strong>Practical Value:</strong> ৳1,95,000</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Apollio AFB Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Modern Design:</strong> Attractive styling</li>
<li class="pro-item"><strong>Smooth Operation:</strong> Refined engine</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 45+ km/l practical</li>
<li class="pro-item"><strong>Comfort:</strong> Ergonomic seating</li>
<li class="pro-item"><strong>Value:</strong> ৳1,95,000 attractive</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Apollio AFB Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Modest Power:</strong> 10.5 bhp basic</li>
<li class="con-item"><strong>Build Quality:</strong> Average finish</li>
<li class="con-item"><strong>Limited Features:</strong> Basic equipment</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Apollio AFB Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,95,000 - Mid-budget</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>10.5 bhp practical</span></p>
<p class="value-item"><strong>Torque:</strong> <span>9.5 nm smooth</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>120kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>45-50 km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Apollio AFB Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Modern attractive</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.5</span> - 45+ km/l</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Ergonomic seating</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳1,95,000</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.4</span> - Average proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Apollio AFB?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,95,000, the AFB is an attractive option for style-conscious commuters wanting modern design, smooth operation, and practical 45+ km/l efficiency for contemporary urban transport.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For style-conscious commuters</p>
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
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">অ্যাপোলিও AFB রিভিউ বাংলাদেশ 2024 - স্টাইলিশ শহুরে কমিউটার</h1>
<p class="summary-text">অ্যাপোলিও AFB একটি 125cc এয়ার-কুল্ড স্টাইলিশ শহুরে কমিউটার যা ব্যবহারিক আধুনিক ডিজাইন প্রতিনিধিত্ব করে। ৳1,95,000 টাকায় মূল্যায়িত, এটি আধুনিক স্টাইলিং এবং মসৃণ অপারেশন প্রদান করে।</p>
</header>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: অ্যাপোলিও AFB কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,95,000 টাকায়, AFB স্টাইল-সচেতন কমিউটারদের জন্য একটি আকর্ষণীয় বিকল্প যারা আধুনিক ডিজাইন এবং ব্যবহারিক দক্ষতা চান।</p>
</div>
</section>
</article>`

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Apollio AFB\n")
	return nil
}

func (s *ApollioAFBReviewBatch20) GetName() string {
	return "ApollioAFBReviewBatch20"
}
