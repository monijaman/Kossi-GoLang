package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BenelliTNT1130ReviewBatch20 struct {
	BaseSeeder
}

func NewBenelliTNT1130ReviewBatch20Seeder() *BenelliTNT1130ReviewBatch20 {
	return &BenelliTNT1130ReviewBatch20{BaseSeeder: BaseSeeder{name: "Benelli TNT 1130 Batch20 Review"}}
}

func (s *BenelliTNT1130ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Benelli TNT 1130").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("benelli tnt 1130 product not found")
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
<h1 class="review-main-title">Benelli TNT 1130 Review Bangladesh 2024 - Italian Performance Beast</h1>
<p class="summary-text">The Benelli TNT 1130 is a 1130cc liquid-cooled Italian performance beast representing extreme street performance engineering. Priced around ৳24,00,000, it delivers 158 bhp raw power, aggressive styling, smooth triple-cylinder character, track-capable handling, and Italian performance soul for purists seeking extreme streetbike experience.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Benelli TNT 1130 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>1130cc Liquid-Cooled:</strong> Italian extreme</li>
<li class="highlight-item"><strong>158 bhp:</strong> Raw extreme power</li>
<li class="highlight-item"><strong>Triple Engine:</strong> Smooth aggressive</li>
<li class="highlight-item"><strong>Track Capable:</strong> Performance handling</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Benelli TNT 1130 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Extreme Power:</strong> 158 bhp thrilling</li>
<li class="pro-item"><strong>Triple Engine:</strong> Smooth refined character</li>
<li class="pro-item"><strong>Aggressive Style:</strong> Intimidating presence</li>
<li class="pro-item"><strong>Track Handling:</strong> Performance capable</li>
<li class="pro-item"><strong>Italian Heritage:</strong> Premium engineering</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Benelli TNT 1130 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Fuel Economy:</strong> 15-18 km/l thirsty</li>
<li class="con-item"><strong>Heavy Weight:</strong> 210kg substantial</li>
<li class="con-item"><strong>Extreme Price:</strong> ৳24,00,000 ultra-premium</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Benelli TNT 1130 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳24,00,000 - Ultra-premium extreme</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>1130cc liquid-cooled triple</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>158 bhp extreme</span></p>
<p class="value-item"><strong>Torque:</strong> <span>92 nm aggressive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>210kg moderate</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>15-18 km/l performance</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Benelli TNT 1130 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.9</span> - 158 bhp extreme</li>
<li class="rating-item"><strong>Engine Character:</strong> <span>4.9</span> - Triple smooth</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.8</span> - Track capable</li>
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Aggressive style</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳24,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Benelli TNT 1130?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳24,00,000, the TNT 1130 is the premium choice for extreme performance seekers wanting Italian engineering, 158 bhp raw power, smooth triple character, track-capable handling, and authentic Italian performance soul for the ultimate streetbike experience.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For extreme performance enthusiasts</span></p>
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
<h1 class="review-main-title">বেনেলি TNT 1130 রিভিউ বাংলাদেশ 2024 - ইতালিয়ান পারফরম্যান্স বিস্ট</h1>
<p class="summary-text">বেনেলি TNT 1130 একটি 1130cc লিকুইড-কুল্ড ইতালিয়ান পারফরম্যান্স বিস্ট যা চরম রাস্তার পারফরম্যান্স ইঞ্জিনিয়ারিং প্রতিনিধিত্ব করে। ৳24,00,000 টাকায় মূল্যায়িত, এটি 158 bhp অপ্রতিরোধ্য শক্তি এবং মসৃণ ট্রিপল-সিলিন্ডার চরিত্র প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বেনেলি TNT 1130 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>1130cc লিকুইড-কুল্ড:</strong> ইতালিয়ান চরম</li>
<li class="highlight-item"><strong>158 bhp:</strong> অপ্রতিরোধ্য চরম শক্তি</li>
<li class="highlight-item"><strong>ট্রিপল ইঞ্জিন:</strong> মসৃণ আক্রমণাত্মক</li>
<li class="highlight-item"><strong>ট্র্যাক সক্ষম:</strong> পারফরম্যান্স হ্যান্ডলিং</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বেনেলি TNT 1130 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>চরম শক্তি:</strong> 158 bhp রোমাঞ্চকর</li>
<li class="pro-item"><strong>ট্রিপল ইঞ্জিন:</strong> মসৃণ পরিমার্জিত চরিত্র</li>
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইল:</strong> ভীতিকর উপস্থিতি</li>
<li class="pro-item"><strong>ট্র্যাক হ্যান্ডলিং:</strong> পারফরম্যান্স সক্ষম</li>
<li class="pro-item"><strong>ইতালিয়ান হেরিটেজ:</strong> প্রিমিয়াম ইঞ্জিনিয়ারিং</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বেনেলি TNT 1130 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>জ্বালানি অর্থনীতি:</strong> 15-18 km/l তৃষ্ণার্ত</li>
<li class="con-item"><strong>ভারী ওজন:</strong> 210kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>চরম মূল্য:</strong> ৳24,00,000 অতি-প্রিমিয়াম</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বেনেলি TNT 1130 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳24,00,000 টাকায়, TNT 1130 চরম পারফরম্যান্স চাহনাকারীদের জন্য প্রিমিয়াম পছন্দ যারা 158 bhp অপ্রতিরোধ্য শক্তি এবং খাঁটি ইতালিয়ান পারফরম্যান্স চাহিদা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চরম পারফরম্যান্স উৎসাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Benelli TNT 1130\n")
	return nil
}

func (s *BenelliTNT1130ReviewBatch20) GetName() string {
	return "BenelliTNT1130ReviewBatch20"
}
