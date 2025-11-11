package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HarleyDavidsonLowriderReviewBatch20 struct {
	BaseSeeder
}

func NewHarleyDavidsonLowriderReviewBatch20Seeder() *HarleyDavidsonLowriderReviewBatch20 {
	return &HarleyDavidsonLowriderReviewBatch20{BaseSeeder: BaseSeeder{name: "Harley-Davidson Lowrider Batch20 Review"}}
}

func (s *HarleyDavidsonLowriderReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Harley-Davidson Lowrider").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("harley-davidson lowrider product not found")
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
<h1 class="review-main-title">Harley-Davidson Lowrider Review Bangladesh 2024 - American Icon Heritage</h1>
<p class="summary-text">The Harley-Davidson Lowrider is a 1868cc liquid-cooled American legend representing iconic cruiser heritage. Priced around ৳35,00,000, it delivers legendary V-twin power, authentic American styling, commanding presence, raw character, and pure motorcycle soul for ultimate cruiser enthusiasts seeking the definitive Harley experience.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Harley-Davidson Lowrider Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>1868cc V-Twin:</strong> Legendary power</li>
<li class="highlight-item"><strong>American Icon:</strong> Heritage legend</li>
<li class="highlight-item"><strong>114 bhp:</strong> Raw power</li>
<li class="highlight-item"><strong>Commanding Presence:</strong> Iconic styling</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Harley-Davidson Lowrider Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Legendary Design:</strong> Iconic heritage</li>
<li class="pro-item"><strong>V-Twin Engine:</strong> Raw character</li>
<li class="pro-item"><strong>Power:</strong> 114 bhp commanding</li>
<li class="pro-item"><strong>Presence:</strong> Intimidating stance</li>
<li class="pro-item"><strong>Soul:</strong> Pure motorcycle spirit</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Harley-Davidson Lowrider Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Fuel Economy:</strong> 12-14 km/l low</li>
<li class="con-item"><strong>Heavy Weight:</strong> 368kg very heavy</li>
<li class="con-item"><strong>Extreme Price:</strong> ৳35,00,000 ultra-premium</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Harley-Davidson Lowrider Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳35,00,000 - Ultra-premium icon</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>1868cc liquid-cooled V-twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>114 bhp raw power</span></p>
<p class="value-item"><strong>Torque:</strong> <span>156 nm commanding</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>368kg very heavy</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>12-14 km/l premium</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Harley-Davidson Lowrider Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Legacy:</strong> <span>5.0</span> - Iconic legend</li>
<li class="rating-item"><strong>Power:</strong> <span>4.9</span> - 114 bhp raw</li>
<li class="rating-item"><strong>Presence:</strong> <span>5.0</span> - Commanding iconic</li>
<li class="rating-item"><strong>Build Quality:</strong> <span>4.9</span> - American premium</li>
<li class="rating-item"><strong>Heritage:</strong> <span>5.0</span> - Pure soul</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Harley-Davidson Lowrider?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.9/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳35,00,000, the Lowrider is the ultimate American icon for purists seeking legendary heritage, raw V-twin power, commanding presence, iconic styling, and pure motorcycle soul for the definitive cruiser experience.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For ultimate cruiser enthusiasts</p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.9,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হার্লে-ডেভিডসন লোরাইডার রিভিউ বাংলাদেশ 2024 - আমেরিকান আইকন হেরিটেজ</h1>
<p class="summary-text">হার্লে-ডেভিডসন লোরাইডার একটি 1868cc লিকুইড-কুল্ড আমেরিকান লিজেন্ড যা আইকনিক ক্রুজার হেরিটেজ প্রতিনিধিত্ব করে। ৳35,00,000 টাকায় মূল্যায়িত, এটি কিংবদন্তি V-টুইন শক্তি এবং খাঁটি আমেরিকান স্টাইলিং প্রদান করে।</p>
</header>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হার্লে-ডেভিডসন লোরাইডার কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.9/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳35,00,000 টাকায়, লোরাইডার চূড়ান্ত আমেরিকান আইকন যারা কিংবদন্তি হেরিটেজ এবং বিশুদ্ধ মোটরসাইকেল আত্মা খুঁজছেন তাদের জন্য।</p>
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

	fmt.Printf("   ✅ Created Harley-Davidson Lowrider\n")
	return nil
}

func (s *HarleyDavidsonLowriderReviewBatch20) GetName() string {
	return "HarleyDavidsonLowriderReviewBatch20"
}
