package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type MitsubishiXFCReviewBatch20 struct {
	BaseSeeder
}

func NewMitsubishiXFCReviewBatch20Seeder() *MitsubishiXFCReviewBatch20 {
	return &MitsubishiXFCReviewBatch20{BaseSeeder: BaseSeeder{name: "Mitsubishi XFC Batch20 Review"}}
}

func (s *MitsubishiXFCReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Mitsubishi XFC").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("mitsubishi xfc product not found")
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
<h1 class="review-main-title">Mitsubishi XFC Review Bangladesh 2024 - Practical Budget Commuter</h1>
<p class="summary-text">The Mitsubishi XFC is a 100cc air-cooled practical budget commuter representing reliable Asian engineering. Priced around ৳1,25,000, it delivers proven dependability, exceptional fuel efficiency, user-friendly operation, strong build quality, and practical value for budget-conscious urban commuters seeking reliable daily transport.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Mitsubishi XFC Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc Air-Cooled:</strong> Budget commuter</li>
<li class="highlight-item"><strong>Reliable Build:</strong> Proven quality</li>
<li class="highlight-item"><strong>50+ km/l:</strong> Exceptional efficiency</li>
<li class="highlight-item"><strong>Low Cost:</strong> ৳1,25,000 budget</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Mitsubishi XFC Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Budget Price:</strong> ৳1,25,000 entry-level</li>
<li class="pro-item"><strong>Fuel Efficiency:</strong> 50+ km/l exceptional</li>
<li class="pro-item"><strong>Reliability:</strong> Proven dependable</li>
<li class="pro-item"><strong>Easy Maintenance:</strong> Simple service</li>
<li class="pro-item"><strong>Strong Build:</strong> Quality construction</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Mitsubishi XFC Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Features:</strong> Minimal styling</li>
<li class="con-item"><strong>Low Power:</strong> 7.5 bhp basic</li>
<li class="con-item"><strong>Modest Speed:</strong> 80-90 km/h top</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Mitsubishi XFC Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,25,000 - Budget segment</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>100cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>7.5 bhp basic</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8 nm practical</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>100kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50-55 km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Mitsubishi XFC Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Reliability:</strong> <span>4.7</span> - Proven dependable</li>
<li class="rating-item"><strong>Fuel Economy:</strong> <span>4.8</span> - 50+ km/l</li>
<li class="rating-item"><strong>Build Quality:</strong> <span>4.6</span> - Solid construction</li>
<li class="rating-item"><strong>Value:</strong> <span>4.8</span> - ৳1,25,000 excellent</li>
<li class="rating-item"><strong>Ease of Use:</strong> <span>4.7</span> - Simple operation</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Mitsubishi XFC?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,25,000, the XFC is the ultimate budget choice for commuters needing reliable, fuel-efficient daily transport with proven quality, exceptional 50+ km/l efficiency, and practical value for cost-conscious riders.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For budget commuters</p>
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
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">মিতসুবিশি XFC রিভিউ বাংলাদেশ 2024 - ব্যবহারিক বাজেট কমিউটার</h1>
<p class="summary-text">মিতসুবিশি XFC একটি 100cc এয়ার-কুল্ড ব্যবহারিক বাজেট কমিউটার যা নির্ভরযোগ্য এশিয়ান ইঞ্জিনিয়ারিং প্রতিনিধিত্ব করে। ৳1,25,000 টাকায় মূল্যায়িত, এটি প্রমাণিত নির্ভরযোগ্যতা এবং ব্যতিক্রমী জ্বালানি দক্ষতা প্রদান করে।</p>
</header>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: মিতসুবিশি XFC কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,25,000 টাকায়, XFC বাজেট কমিউটারদের জন্য চূড়ান্ত পছন্দ যারা নির্ভরযোগ্য, জ্বালানি-দক্ষ দৈনন্দিন পরিবহন প্রয়োজন।</p>
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

	fmt.Printf("   ✅ Created Bangla translation for Mitsubishi XFC\n")
	return nil
}

func (s *MitsubishiXFCReviewBatch20) GetName() string {
	return "MitsubishiXFCReviewBatch20"
}
