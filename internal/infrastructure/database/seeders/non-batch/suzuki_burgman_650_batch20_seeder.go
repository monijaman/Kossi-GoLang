package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SuzukiBurgman650ReviewBatch20 struct {
	BaseSeeder
}

func NewSuzukiBurgman650ReviewBatch20Seeder() *SuzukiBurgman650ReviewBatch20 {
	return &SuzukiBurgman650ReviewBatch20{BaseSeeder: BaseSeeder{name: "Suzuki Burgman 650 Batch20 Review"}}
}

func (s *SuzukiBurgman650ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Burgman 650").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki burgman 650 product not found")
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
<h1 class="review-main-title">Suzuki Burgman 650 Review Bangladesh 2024 - Premium Urban Scooter</h1>
<p class="summary-text">The Suzuki Burgman 650 is a 650cc liquid-cooled premium urban scooter representing the pinnacle of modern commuter technology. Priced around ৳12,50,000, it delivers automatic transmission, comfortable seating, advanced features, smooth highway capability, and sophisticated urban commuting for discerning riders seeking premium practical transport.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Burgman 650 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>650cc Liquid-Cooled:</strong> Premium scooter</li>
<li class="highlight-item"><strong>Automatic CVT:</strong> Smooth transmission</li>
<li class="highlight-item"><strong>40 bhp:</strong> Highway capable</li>
<li class="highlight-item"><strong>Comfort:</strong> Spacious seating</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Burgman 650 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Automatic:</strong> CVT smooth transmission</li>
<li class="pro-item"><strong>Comfort:</strong> Spacious plush seating</li>
<li class="pro-item"><strong>Power:</strong> 40 bhp highway capable</li>
<li class="pro-item"><strong>Storage:</strong> Generous under-seat</li>
<li class="pro-item"><strong>Features:</strong> Modern technology</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Burgman 650 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Fuel Economy:</strong> 20-24 km/l average</li>
<li class="con-item"><strong>Heavy Weight:</strong> 220kg substantial</li>
<li class="con-item"><strong>Premium Price:</strong> ৳12,50,000</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Burgman 650 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳12,50,000 - Premium scooter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>650cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>40 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>62 nm adequate</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>220kg substantial</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>20-24 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Burgman 650 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Comfort:</strong> <span>4.9</span> - Spacious plush</li>
<li class="rating-item"><strong>Technology:</strong> <span>4.7</span> - Modern features</li>
<li class="rating-item"><strong>Smoothness:</strong> <span>4.8</span> - CVT automatic</li>
<li class="rating-item"><strong>Build Quality:</strong> <span>4.8</span> - Japanese premium</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳12,50,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Burgman 650?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳12,50,000, the Burgman 650 is the premium choice for discerning commuters seeking sophisticated automatic transmission, comfortable seating, advanced features, smooth highway capability, and premium practical urban transport.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For premium commuters</p>
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
<h1 class="review-main-title">সুজুকি বার্গম্যান 650 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম শহুরে স্কুটার</h1>
<p class="summary-text">সুজুকি বার্গম্যান 650 একটি 650cc লিকুইড-কুল্ড প্রিমিয়াম শহুরে স্কুটার যা আধুনিক কমিউটার প্রযুক্তির শিখর প্রতিনিধিত্ব করে। ৳12,50,000 টাকায় মূল্যায়িত, এটি স্বয়ংক্রিয় প্রেরণ এবং আরামদায়ক আসন প্রদান করে।</p>
</header>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি বার্গম্যান 650 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳12,50,000 টাকায়, বার্গম্যান 650 বিচক্ষণ কমিউটারদের জন্য প্রিমিয়াম পছন্দ যারা পরিশীলিত স্বয়ংক্রিয় প্রেরণ এবং প্রিমিয়াম ব্যবহারিক পরিবহন খুঁজছেন।</p>
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

	fmt.Printf("   ✅ Created Suzuki Burgman 650\n")
	return nil
}

func (s *SuzukiBurgman650ReviewBatch20) GetName() string {
	return "SuzukiBurgman650ReviewBatch20"
}
