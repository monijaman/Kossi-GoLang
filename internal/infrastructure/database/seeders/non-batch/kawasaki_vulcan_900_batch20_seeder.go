package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type KawasakiVulcan900ReviewBatch20 struct {
	BaseSeeder
}

func NewKawasakiVulcan900ReviewBatch20Seeder() *KawasakiVulcan900ReviewBatch20 {
	return &KawasakiVulcan900ReviewBatch20{BaseSeeder: BaseSeeder{name: "Kawasaki Vulcan 900 Batch20 Review"}}
}

func (s *KawasakiVulcan900ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Kawasaki Vulcan 900").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("kawasaki vulcan 900 product not found")
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
<h1 class="review-main-title">Kawasaki Vulcan 900 Review Bangladesh 2024 - Classic Cruiser Legend</h1>
<p class="summary-text">The Kawasaki Vulcan 900 is a 900cc liquid-cooled classic cruiser representing timeless American-style motorcycle heritage. Priced around ৳18,00,000, it delivers authentic cruiser design, relaxed riding position, distinctive V-twin character, impressive low-end torque, and heritage cruiser performance for riders seeking classic motorcycle soul.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Kawasaki Vulcan 900 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>900cc V-Twin:</strong> Cruiser authentic</li>
<li class="highlight-item"><strong>Relaxed Riding:</strong> Cruiser comfort</li>
<li class="highlight-item"><strong>50 bhp:</strong> Smooth torque</li>
<li class="highlight-item"><strong>American Heritage:</strong> Classic design</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Kawasaki Vulcan 900 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Classic Styling:</strong> Heritage design</li>
<li class="pro-item"><strong>V-Twin Engine:</strong> Character sound</li>
<li class="pro-item"><strong>Torque:</strong> 63 nm impressive</li>
<li class="pro-item"><strong>Comfort:</strong> Relaxed cruising</li>
<li class="pro-item"><strong>Build Quality:</strong> Solid Japanese</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Kawasaki Vulcan 900 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Fuel Economy:</strong> 15-18 km/l modest</li>
<li class="con-item"><strong>Heavy Weight:</strong> 240kg substantial</li>
<li class="con-item"><strong>Premium Price:</strong> ৳18,00,000</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Kawasaki Vulcan 900 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳18,00,000 - Premium cruiser</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>900cc liquid-cooled V-twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>50 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>63 nm impressive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>240kg substantial</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>15-18 km/l premium cruiser</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Kawasaki Vulcan 900 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Heritage classic</li>
<li class="rating-item"><strong>Torque:</strong> <span>4.7</span> - 63 nm smooth</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Relaxed cruising</li>
<li class="rating-item"><strong>Build Quality:</strong> <span>4.8</span> - Japanese solid</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳18,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Kawasaki Vulcan 900?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳18,00,000, the Vulcan 900 is the premium cruiser choice for riders seeking authentic American-style heritage, distinctive V-twin character, impressive low-end torque, relaxed cruiser comfort, and classic motorcycle soul.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For cruiser enthusiasts</p>
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
<h1 class="review-main-title">কাওয়াসাকি ভালকান 900 রিভিউ বাংলাদেশ 2024 - ক্লাসিক ক্রুজার লিজেন্ড</h1>
<p class="summary-text">কাওয়াসাকি ভালকান 900 একটি 900cc লিকুইড-কুল্ড ক্লাসিক ক্রুজার যা কালজয়ী আমেরিকান-শৈলী মোটরসাইকেল হেরিটেজ প্রতিনিধিত্ব করে। ৳18,00,000 টাকায় মূল্যায়িত, এটি খাঁটি ক্রুজার ডিজাইন এবং চিরন্তন আত্মা প্রদান করে।</p>
</header>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কাওয়াসাকি ভালকান 900 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳18,00,000 টাকায়, ভালকান 900 ক্রুজার উৎসাহীদের জন্য প্রিমিয়াম পছন্দ যারা খাঁটি আমেরিকান-শৈলী হেরিটেজ এবং ক্লাসিক মোটরসাইকেল চরিত্র খুঁজছেন।</p>
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

	fmt.Printf("   ✅ Created Kawasaki Vulcan 900\n")
	return nil
}

func (s *KawasakiVulcan900ReviewBatch20) GetName() string {
	return "KawasakiVulcan900ReviewBatch20"
}
