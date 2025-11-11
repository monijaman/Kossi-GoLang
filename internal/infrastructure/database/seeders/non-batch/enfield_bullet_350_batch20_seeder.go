package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type EnfieldBullet350ReviewBatch20 struct {
	BaseSeeder
}

func NewEnfieldBullet350ReviewBatch20Seeder() *EnfieldBullet350ReviewBatch20 {
	return &EnfieldBullet350ReviewBatch20{BaseSeeder: BaseSeeder{name: "Enfield Bullet 350 Batch20 Review"}}
}

func (s *EnfieldBullet350ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Enfield Bullet 350").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("enfield bullet 350 product not found")
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
<h1 class="review-main-title">Royal Enfield Bullet 350 Review Bangladesh 2024 - Timeless Classic</h1>
<p class="summary-text">The Royal Enfield Bullet 350 is a 350cc air-cooled classic representing timeless British-Indian motorcycle heritage. Priced around ৳3,50,000, it delivers authentic retro styling, simple air-cooled reliability, iconic single-cylinder character, bulletproof durability, and classic soul for riders seeking timeless motorcycle authenticity.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Royal Enfield Bullet 350 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>350cc Air-Cooled:</strong> Classic authentic</li>
<li class="highlight-item"><strong>Retro Styling:</strong> Timeless design</li>
<li class="highlight-item"><strong>Single-Cylinder:</strong> Iconic character</li>
<li class="highlight-item"><strong>Bulletproof:</strong> Legendary durability</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Royal Enfield Bullet 350 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Retro Style:</strong> Timeless design</li>
<li class="pro-item"><strong>Simplicity:</strong> Easy maintenance</li>
<li class="pro-item"><strong>Reliability:</strong> Bulletproof proven</li>
<li class="pro-item"><strong>Character:</strong> Iconic single sound</li>
<li class="pro-item"><strong>Value:</strong> ৳3,50,000 classic</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Royal Enfield Bullet 350 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Performance:</strong> 20 bhp modest</li>
<li class="con-item"><strong>Vibration:</strong> Single-cylinder vibration</li>
<li class="con-item"><strong>Fuel Economy:</strong> 30-35 km/l average</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Royal Enfield Bullet 350 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,50,000 - Classic segment</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>350cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>20 bhp modest</span></p>
<p class="value-item"><strong>Torque:</strong> <span>28 nm adequate</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>202kg substantial</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>30-35 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Royal Enfield Bullet 350 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Heritage:</strong> <span>4.9</span> - Timeless iconic</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.8</span> - Bulletproof proven</li>
<li class="rating-item"><strong>Character:</strong> <span>4.7</span> - Classic single</li>
<li class="rating-item"><strong>Simplicity:</strong> <span>4.8</span> - Easy maintenance</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳3,50,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Royal Enfield Bullet 350?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,50,000, the Bullet 350 is the ultimate timeless choice for classic purists seeking authentic retro styling, simple reliability, iconic single-cylinder character, bulletproof durability, and genuine motorcycle soul for authentic riding experience.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For classic enthusiasts</p>
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
<h1 class="review-main-title">রয়াল এনফিল্ড বুলেট 350 রিভিউ বাংলাদেশ 2024 - কালজয়ী ক্লাসিক</h1>
<p class="summary-text">রয়াল এনফিল্ড বুলেট 350 একটি 350cc এয়ার-কুল্ড ক্লাসিক যা কালজয়ী ব্রিটিশ-ভারতীয় মোটরসাইকেল হেরিটেজ প্রতিনিধিত্ব করে। ৳3,50,000 টাকায় মূল্যায়িত, এটি খাঁটি রেট্রো স্টাইলিং এবং চিরন্তন মোটরসাইকেল প্রামাণিকতা প্রদান করে।</p>
</header>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: রয়াল এনফিল্ড বুলেট 350 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,50,000 টাকায়, বুলেট 350 ক্লাসিক পুরিস্তদের জন্য চূড়ান্ত কালজয়ী পছন্দ যারা খাঁটি রেট্রো স্টাইলিং এবং সত্যিকারের মোটরসাইকেল আত্মা খুঁজছেন।</p>
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

	fmt.Printf("   ✅ Created Royal Enfield Bullet 350\n")
	return nil
}

func (s *EnfieldBullet350ReviewBatch20) GetName() string {
	return "EnfieldBullet350ReviewBatch20"
}
