package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SuzukiSkywave400ReviewBatch20 struct {
	BaseSeeder
}

func NewSuzukiSkywave400ReviewBatch20Seeder() *SuzukiSkywave400ReviewBatch20 {
	return &SuzukiSkywave400ReviewBatch20{BaseSeeder: BaseSeeder{name: "Suzuki Skywave 400 Batch20 Review"}}
}

func (s *SuzukiSkywave400ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Skywave 400").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki skywave 400 product not found")
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
<h1 class="review-main-title">Suzuki Skywave 400 Review Bangladesh 2024 - Premium Cruise Automatic Scooter</h1>
<p class="summary-text">The Suzuki Skywave 400 is a 400cc liquid-cooled premium cruise automatic scooter representing smooth transmission technology. Priced around ৳8,00,000, it delivers smooth CVT transmission, comfortable spacious seating, excellent storage, highway cruising capability, and practical efficiency for comfortable daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Skywave 400 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>400cc Liquid-Cooled:</strong> Cruise scooter</li>
<li class="highlight-item"><strong>Smooth CVT:</strong> Automatic transmission</li>
<li class="highlight-item"><strong>32 bhp:</strong> Adequate power</li>
<li class="highlight-item"><strong>Spacious Comfort:</strong> Premium seating</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Skywave 400 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Smooth CVT:</strong> Automatic comfort</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Spacious design</li>
<li class="pro-item"><strong>Excellent Storage:</strong> Large compartment</li>
<li class="pro-item"><strong>Highway Capable:</strong> Cruise comfortable</li>
<li class="pro-item"><strong>Premium Build:</strong> Japanese quality</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Skywave 400 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Fuel Economy:</strong> 25-28 km/l average</li>
<li class="con-item"><strong>Heavy Weight:</strong> 228kg substantial</li>
<li class="con-item"><strong>Premium Price:</strong> ৳8,00,000</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Skywave 400 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳8,00,000 - Premium cruise scooter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>400cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>32 bhp adequate</span></p>
<p class="value-item"><strong>Torque:</strong> <span>37 nm smooth</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>228kg substantial</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>25-28 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Skywave 400 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Spacious seating</li>
<li class="rating-item"><strong>Transmission:</strong> <span>4.8</span> - Smooth CVT</li>
<li class="rating-item"><strong>Storage:</strong> <span>4.8</span> - Excellent capacity</li>
<li class="rating-item"><strong>Highway Cruising:</strong> <span>4.7</span> - Capable smooth</li>
<li class="rating-item"><strong>Value:</strong> <span>4.5</span> - ৳8,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Skywave 400?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳8,00,000, the Skywave 400 is the premium choice for comfort-focused commuters wanting smooth automatic transmission, spacious seating, excellent storage, highway cruising capability, and practical 25-28 km/l efficiency for comfortable daily transport.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For comfort commuters</span></p>
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
<h1 class="review-main-title">সুজুকি স্কাইওয়েভ 400 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম ক্রুজ অটোমেটিক স্কুটার</h1>
<p class="summary-text">সুজুকি স্কাইওয়েভ 400 একটি 400cc লিকুইড-কুল্ড প্রিমিয়াম ক্রুজ অটোমেটিক স্কুটার যা মসৃণ প্রেরণ প্রযুক্তি প্রতিনিধিত্ব করে। ৳8,00,000 টাকায় মূল্যায়িত, এটি মসৃণ CVT প্রেরণ এবং আরামদায়ক প্রশস্ত আসন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি স্কাইওয়েভ 400 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>400cc লিকুইড-কুল্ড:</strong> ক্রুজ স্কুটার</li>
<li class="highlight-item"><strong>মসৃণ CVT:</strong> অটোমেটিক প্রেরণ</li>
<li class="highlight-item"><strong>32 bhp:</strong> পর্যাপ্ত শক্তি</li>
<li class="highlight-item"><strong>প্রশস্ত আরাম:</strong> প্রিমিয়াম আসন</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি স্কাইওয়েভ 400 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>মসৃণ CVT:</strong> অটোমেটিক আরাম</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> প্রশস্ত ডিজাইন</li>
<li class="pro-item"><strong>চমৎকার স্টোরেজ:</strong> বৃহৎ বগি</li>
<li class="pro-item"><strong>হাইওয়ে সক্ষম:</strong> ক্রুজ আরামদায়ক</li>
<li class="pro-item"><strong>প্রিমিয়াম বিল্ড:</strong> জাপানি গুণমান</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি স্কাইওয়েভ 400 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>জ্বালানি অর্থনীতি:</strong> 25-28 km/l গড়</li>
<li class="con-item"><strong>ভারী ওজন:</strong> 228kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳8,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি স্কাইওয়েভ 400 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳8,00,000 টাকায়, স্কাইওয়েভ 400 আরাম-কেন্দ্রিক কমিউটারদের জন্য প্রিমিয়াম পছন্দ যারা মসৃণ অটোমেটিক প্রেরণ, প্রশস্ত আসন এবং আরামদায়ক দৈনন্দিন পরিবহন খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আরাম কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Suzuki Skywave 400\n")
	return nil
}

func (s *SuzukiSkywave400ReviewBatch20) GetName() string {
	return "SuzukiSkywave400ReviewBatch20"
}
