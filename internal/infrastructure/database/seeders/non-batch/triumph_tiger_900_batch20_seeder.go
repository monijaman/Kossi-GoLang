package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type TriumphTiger900ReviewBatch20 struct {
	BaseSeeder
}

func NewTriumphTiger900ReviewBatch20Seeder() *TriumphTiger900ReviewBatch20 {
	return &TriumphTiger900ReviewBatch20{BaseSeeder: BaseSeeder{name: "Triumph Tiger 900 Batch20 Review"}}
}

func (s *TriumphTiger900ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Triumph Tiger 900").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("triumph tiger 900 product not found")
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
<h1 class="review-main-title">Triumph Tiger 900 Review Bangladesh 2024 - Adventure Touring Beast</h1>
<p class="summary-text">The Triumph Tiger 900 is a 889cc liquid-cooled adventure touring motorcycle representing modern expedition capability. Priced around ৳19,00,000, it delivers triple-cylinder engine character, adventure-focused design, sophisticated electronics, long-range capability, and adventure touring soul for explorers seeking capable expedition transport.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Triumph Tiger 900 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>889cc Triple:</strong> Adventure engine</li>
<li class="highlight-item"><strong>Adventure Design:</strong> Expedition capable</li>
<li class="highlight-item"><strong>94 bhp:</strong> Smooth powerful</li>
<li class="highlight-item"><strong>Long Range:</strong> Adventure touring</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Triumph Tiger 900 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Triple Engine:</strong> Smooth character</li>
<li class="pro-item"><strong>Capability:</strong> Adventure-ready</li>
<li class="pro-item"><strong>Electronics:</strong> Sophisticated systems</li>
<li class="pro-item"><strong>Comfort:</strong> Long-distance capable</li>
<li class="pro-item"><strong>Build:</strong> Premium British</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Triumph Tiger 900 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Fuel Economy:</strong> 16-20 km/l average</li>
<li class="con-item"><strong>Heavy Weight:</strong> 216kg adventure</li>
<li class="con-item"><strong>Premium Price:</strong> ৳19,00,000</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Triumph Tiger 900 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳19,00,000 - Premium adventure</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>889cc liquid-cooled triple</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>94 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>87 nm adequate</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>216kg adventure</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>16-20 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Triumph Tiger 900 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Adventure Capability:</strong> <span>4.9</span> - Expedition ready</li>
<li class="rating-item"><strong>Engine Character:</strong> <span>4.8</span> - Triple smooth</li>
<li class="rating-item"><strong>Technology:</strong> <span>4.7</span> - Sophisticated</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Long-distance</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳19,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Triumph Tiger 900?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳19,00,000, the Tiger 900 is the premium adventure choice for explorers seeking capable expedition transport, sophisticated triple-engine character, adventure-focused design, long-range capability, and modern adventure touring soul.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For adventure explorers</p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.8,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ট্রাইম্ফ টাইগার 900 রিভিউ বাংলাদেশ 2024 - অ্যাডভেঞ্চার ট্যুরিং বিস্ট</h1>
<p class="summary-text">ট্রাইম্ফ টাইগার 900 একটি 889cc লিকুইড-কুল্ড অ্যাডভেঞ্চার ট্যুরিং মোটরসাইকেল যা আধুনিক অভিযান সক্ষমতা প্রতিনিধিত্ব করে। ৳19,00,000 টাকায় মূল্যায়িত, এটি ট্রিপল-সিলিন্ডার ইঞ্জিন চরিত্র এবং অ্যাডভেঞ্চার-কেন্দ্রিক ডিজাইন প্রদান করে।</p>
</header>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ট্রাইম্ফ টাইগার 900 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳19,00,000 টাকায়, টাইগার 900 অভিজ্ঞতা অন্বেষণকারীদের জন্য প্রিমিয়াম অ্যাডভেঞ্চার পছন্দ যারা সক্ষম অভিযান পরিবহন এবং আধুনিক অ্যাডভেঞ্চার ট্যুরিং আত্মা খুঁজছেন।</p>
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

	fmt.Printf("   ✅ Created Triumph Tiger 900\n")
	return nil
}

func (s *TriumphTiger900ReviewBatch20) GetName() string {
	return "TriumphTiger900ReviewBatch20"
}
