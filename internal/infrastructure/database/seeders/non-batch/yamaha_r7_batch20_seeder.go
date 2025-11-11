package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaR7ReviewBatch20 struct {
	BaseSeeder
}

func NewYamahaR7ReviewBatch20Seeder() *YamahaR7ReviewBatch20 {
	return &YamahaR7ReviewBatch20{BaseSeeder: BaseSeeder{name: "Yamaha R7 Batch20 Review"}}
}

func (s *YamahaR7ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R7").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha r7 product not found")
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
<h1 class="review-main-title">Yamaha R7 Review Bangladesh 2024 - Modern Retro Sportbike</h1>
<p class="summary-text">The Yamaha R7 is a 689cc liquid-cooled modern retro sportbike representing contemporary performance heritage. Priced around ৳9,25,000, it delivers parallel twin engine, retro styling with modern tech, responsive handling, sporty character, and versatile capability for riders seeking modern performance with classic soul.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R7 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>689cc Twin:</strong> Modern engine</li>
<li class="highlight-item"><strong>Retro Modern:</strong> Contemporary style</li>
<li class="highlight-item"><strong>55 bhp:</strong> Responsive power</li>
<li class="highlight-item"><strong>Nimble Handling:</strong> Sporty responsive</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R7 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Retro Style:</strong> Modern classic appeal</li>
<li class="pro-item"><strong>Engine Character:</strong> Twin refinement</li>
<li class="pro-item"><strong>Handling:</strong> Nimble responsive</li>
<li class="pro-item"><strong>Technology:</strong> Modern features</li>
<li class="pro-item"><strong>Versatility:</strong> Daily ride-friendly</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R7 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Moderate Power:</strong> 55 bhp practical</li>
<li class="con-item"><strong>Fuel Economy:</strong> 22-26 km/l average</li>
<li class="con-item"><strong>Mid-Premium Price:</strong> ৳9,25,000</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R7 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳9,25,000 - Mid-premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>689cc liquid-cooled parallel twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>55 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>67 nm smooth</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>184kg nimble</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>22-26 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R7 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Retro modern</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Nimble responsive</li>
<li class="rating-item"><strong>Engine:</strong> <span>4.6</span> - Twin character</li>
<li class="rating-item"><strong>Technology:</strong> <span>4.6</span> - Modern features</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳9,25,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R7?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳9,25,000, the R7 is an excellent choice for riders seeking modern retro styling, responsive twin engine, nimble handling, contemporary features, and versatile capability combining classic soul with modern performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For retro sportbike fans</p>
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
<h1 class="review-main-title">ইয়ামাহা R7 রিভিউ বাংলাদেশ 2024 - আধুনিক রেট্রো স্পোর্টবাইক</h1>
<p class="summary-text">ইয়ামাহা R7 একটি 689cc লিকুইড-কুল্ড আধুনিক রেট্রো স্পোর্টবাইক যা সমসাময়িক পারফরম্যান্স হেরিটেজ প্রতিনিধিত্ব করে। ৳9,25,000 টাকায় মূল্যায়িত, এটি সমান্তরাল টুইন ইঞ্জিন এবং আধুনিক প্রযুক্তি প্রদান করে।</p>
</header>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা R7 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳9,25,000 টাকায়, R7 রেট্রো স্পোর্টবাইক অনুরাগীদের জন্য একটি চমৎকার পছন্দ যারা আধুনিক রেট্রো স্টাইলিং এবং প্রতিক্রিয়াশীল পারফরম্যান্স খুঁজছেন।</p>
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

	fmt.Printf("   ✅ Created Yamaha R7\n")
	return nil
}

func (s *YamahaR7ReviewBatch20) GetName() string {
	return "YamahaR7ReviewBatch20"
}
