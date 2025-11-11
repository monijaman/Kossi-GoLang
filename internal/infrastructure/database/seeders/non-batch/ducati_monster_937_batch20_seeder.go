package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// DucatiMonster937ReviewBatch20 handles seeding Ducati Monster 937 product review
type DucatiMonster937ReviewBatch20 struct {
	BaseSeeder
}

// NewDucatiMonster937ReviewBatch20Seeder creates a new DucatiMonster937ReviewBatch20
func NewDucatiMonster937ReviewBatch20Seeder() *DucatiMonster937ReviewBatch20 {
	return &DucatiMonster937ReviewBatch20{BaseSeeder: BaseSeeder{name: "Ducati Monster 937 Batch20 Review"}}
}

// Seed implements the Seeder interface
func (s *DucatiMonster937ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Ducati Monster 937").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ducati monster 937 product not found")
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
<h1 class="review-main-title">Ducati Monster 937 Review Bangladesh 2024 - Iconic Naked Beast</h1>
<p class="summary-text">The Ducati Monster 937 is a 937cc air-cooled iconic naked roadster representing Italian engineering mastery. Priced around ৳22,50,000, it delivers legendary Ducati style, distinctive L-twin engine, impressive performance, timeless naked design, and thrilling riding dynamics for purists seeking authentic Italian motorcycle character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Ducati Monster 937 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>937cc Air-Cooled:</strong> Iconic Ducati</li>
<li class="highlight-item"><strong>L-Twin Engine:</strong> Distinctive character</li>
<li class="highlight-item"><strong>111 bhp:</strong> Thrilling performance</li>
<li class="highlight-item"><strong>Naked Design:</strong> Timeless styling</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Ducati Monster 937 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Iconic Design:</strong> Legendary styling heritage</li>
<li class="pro-item"><strong>L-Twin Engine:</strong> Distinctive character sound</li>
<li class="pro-item"><strong>Performance:</strong> 111 bhp thrilling</li>
<li class="pro-item"><strong>Handling:</strong> Nimble responsive</li>
<li class="pro-item"><strong>Premium Build:</strong> Italian craftsmanship</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Ducati Monster 937 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳22,50,000 premium</li>
<li class="con-item"><strong>Fuel Economy:</strong> 14-16 km/l modest</li>
<li class="con-item"><strong>Maintenance:</strong> Expensive upkeep required</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Ducati Monster 937 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳22,50,000 - Premium icon</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>937cc air-cooled L-twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>111 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</span> <span>86 nm aggressive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>198kg moderate</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>14-16 km/l premium</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Ducati Monster 937 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.9</span> - 111 bhp thrilling</li>
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Iconic legendary</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.8</span> - Responsive nimble</li>
<li class="rating-item"><strong>Build Quality:</strong> <span>4.9</span> - Italian premium</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - Premium ৳22,50,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Ducati Monster 937?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳22,50,000, the Monster 937 is the premium choice for purists seeking iconic Ducati character, distinctive L-twin engineering, thrilling 111 bhp performance, responsive handling, and timeless naked motorcycle styling for authentic riding experience.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span>YES</span> - For premium motorcycle enthusiasts</p>
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

	fmt.Printf("   ✅ Created English review for Ducati Monster 937 (Rating: 4.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ডুকাটি মনস্টার 937 রিভিউ বাংলাদেশ 2024 - আইকনিক নেকেড বিস্ট</h1>
<p class="summary-text">ডুকাটি মনস্টার 937 একটি 937cc এয়ার-কুল্ড আইকনিক নেকেড রোডস্টার যা ইতালিয়ান ইঞ্জিনিয়ারিং মাস্টারি প্রতিনিধিত্ব করে। ৳22,50,000 টাকায় মূল্যায়িত, এটি কিংবদন্তি ডুকাটি স্টাইল এবং আতœীয় আইকনিক ডিজাইন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ডুকাটি মনস্টার 937 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>937cc এয়ার-কুল্ড:</strong> আইকনিক ডুকাটি</li>
<li class="highlight-item"><strong>L-টুইন ইঞ্জিন:</strong> স্বতন্ত্র চরিত্র</li>
<li class="highlight-item"><strong>111 bhp:</strong> রোমাঞ্চকর পারফরম্যান্স</li>
<li class="highlight-item"><strong>নেকেড ডিজাইন:</strong> কালজয়ী স্টাইলিং</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ডুকাটি মনস্টার 937 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আইকনিক ডিজাইন:</strong> কিংবদন্তি স্টাইলিং হেরিটেজ</li>
<li class="pro-item"><strong>L-টুইন ইঞ্জিন:</strong> স্বতন্ত্র চরিত্র সাউন্ড</li>
<li class="pro-item"><strong>পারফরম্যান্স:</strong> 111 bhp রোমাঞ্চকর</li>
<li class="pro-item"><strong>হ্যান্ডলিং:</strong> নিপুণ প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>প্রিমিয়াম বিল্ড:</strong> ইতালিয়ান কারুশিল্প</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ডুকাটি মনস্টার 937 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳22,50,000 প্রিমিয়াম</li>
<li class="con-item"><strong>জ্বালানি অর্থনীতি:</strong> 14-16 km/l মামুলি</li>
<li class="con-item"><strong>রক্ষণাবেক্ষণ:</strong> ব্যয়বহুল রক্ষণাবেক্ষণ প্রয়োজন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ডুকাটি মনস্টার 937 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳22,50,000 টাকায়, মনস্টার 937 প্রিমিয়াম পছন্দ যারা কিংবদন্তি ডুকাটি চরিত্র এবং স্বতন্ত্র L-টুইন ইঞ্জিনিয়ারিং খুঁজছেন তাদের জন্য।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span>হ্যাঁ</span> - প্রিমিয়াম মোটরসাইকেল উৎসাহীদের জন্য</p>
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

	fmt.Printf("   ✅ Created Bangla translation for Ducati Monster 937\n")
	return nil
}

// GetName returns the seeder name
func (s *DucatiMonster937ReviewBatch20) GetName() string {
	return "DucatiMonster937ReviewBatch20"
}
