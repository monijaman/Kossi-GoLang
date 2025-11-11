package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// DucatiPanigaleV4ReviewBatch16 handles seeding Ducati Panigale V4 product review and translation
type DucatiPanigaleV4ReviewBatch16 struct {
	BaseSeeder
}

// NewDucatiPanigaleV4ReviewBatch16Seeder creates a new DucatiPanigaleV4ReviewBatch16
func NewDucatiPanigaleV4ReviewBatch16Seeder() *DucatiPanigaleV4ReviewBatch16 {
	return &DucatiPanigaleV4ReviewBatch16{BaseSeeder: BaseSeeder{name: "Ducati Panigale V4 Batch16 Review"}}
}

// Seed implements the Seeder interface
func (s *DucatiPanigaleV4ReviewBatch16) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Ducati Panigale V4").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ducati panigale v4 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding ducati panigale v4 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Ducati Panigale V4 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Ducati Panigale V4 Review Bangladesh 2024 - Italian Racing Legend</h1>
<p class="summary-text">The Ducati Panigale V4 is an iconic 1103cc superbike featuring Ducati's signature V4 engine, MotoGP-derived technology, razor-sharp handling, and unmistakable Italian design. Priced around ৳18,00,000, it's the ultimate expression of Ducati racing heritage for enthusiasts seeking raw performance art.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Ducati Panigale V4 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1103cc V4 Beast:</strong> <span class="highlight-value">MotoGP DNA engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Racing Technology:</strong> <span class="highlight-value">Track-proven performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Razor Handling:</strong> <span class="highlight-value">Circuit-tuned dynamics</span></li>
<li class="highlight-item"><strong class="highlight-label">Italian Icon:</strong> <span class="highlight-value">Legendary design</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Ducati Panigale V4 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">V4 Engine:</strong> <span class="pro-description">Legendary power</span></li>
<li class="pro-item"><strong class="pro-label">Racing Tech:</strong> <span class="pro-description">MotoGP heritage</span></li>
<li class="pro-item"><strong class="pro-label">Incredible Handling:</strong> <span class="pro-description">Circuit sharp</span></li>
<li class="pro-item"><strong class="pro-label">Italian Design:</strong> <span class="pro-description">Iconic beauty</span></li>
<li class="pro-item"><strong class="pro-label">Performance Art:</strong> <span class="pro-description">Pure racing</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Ducati Panigale V4 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Extreme Price:</strong> <span class="con-description">৳18,00,000 ultra-expensive</span></li>
<li class="con-item"><strong class="con-label">High Maintenance:</strong> <span class="con-description">Very costly</span></li>
<li class="con-item"><strong class="con-label">Racing Focused:</strong> <span class="con-description">Not commuter friendly</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Ducati Panigale V4 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳18,00,000 - Ultra-premium superbike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very high - ৳15-25 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">18-22 km/l low</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Ducati Panigale V4 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- V4 legendary</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Circuit razor</span></li>
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- MotoGP proven</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Iconic Italian</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Ultra-premium</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Ducati Panigale V4?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳18,00,000, the Ducati Panigale V4 is for collectors and enthusiasts seeking the ultimate Italian racing machine, MotoGP-derived engineering, legendary V4 power, and the pure art of motorcycle performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultimate performance purists</span></p>
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
		return fmt.Errorf("error creating ducati panigale v4 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Ducati Panigale V4 (Rating: 4.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ডুকাটি পানিগালে V4 রিভিউ বাংলাদেশ 2024 - ইতালিয়ান রেসিং লেজেন্ড</h1>
<p class="summary-text">ডুকাটি পানিগালে V4 একটি আইকনিক 1103cc সুপারবাইক যা ডুকাটির স্বাক্ষর V4 ইঞ্জিন, MotoGP-ডেরাইভড প্রযুক্তি, ধারালো হ্যান্ডলিং এবং স্বতন্ত্র ইতালিয়ান ডিজাইন বৈশিষ্ট্যযুক্ত। ৳18,00,000 টাকায় মূল্যায়িত, এটি ডুকাটি রেসিং ঐতিহ্যের চূড়ান্ত অভিব্যক্তি উত্সাহীদের জন্য যারা কাঁচা কর্মক্ষমতা শিল্প খুঁজছেন।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ডুকাটি পানিগালে V4 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1103cc V4 বিস্ট:</strong> <span class="highlight-value">MotoGP DNA ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">রেসিং প্রযুক্তি:</strong> <span class="highlight-value">ট্র্যাক-প্রমাণিত কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ধারালো হ্যান্ডলিং:</strong> <span class="highlight-value">সার্কিট-টিউনড গতিশীলতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ইতালিয়ান আইকন:</strong> <span class="highlight-value">কিংবদন্তি ডিজাইন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ডুকাটি পানিগালে V4 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">V4 ইঞ্জিন:</strong> <span class="pro-description">কিংবদন্তি শক্তি</span></li>
<li class="pro-item"><strong class="pro-label">রেসিং প্রযুক্তি:</strong> <span class="pro-description">MotoGP ঐতিহ্য</span></li>
<li class="pro-item"><strong class="pro-label">অবিশ্বাস্য হ্যান্ডলিং:</strong> <span class="pro-description">সার্কিট ধারালো</span></li>
<li class="pro-item"><strong class="pro-label">ইতালিয়ান ডিজাইন:</strong> <span class="pro-description">আইকনিক সৌন্দর্য</span></li>
<li class="pro-item"><strong class="pro-label">পারফরম্যান্স শিল্প:</strong> <span class="pro-description">বিশুদ্ধ রেসিং</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ডুকাটি পানিগালে V4 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">চরম মূল্য:</strong> <span class="con-description">৳18,00,000 অতি-ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">উচ্চ রক্ষণাবেক্ষণ:</strong> <span class="con-description">অত্যন্ত ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">রেসিং ফোকাসড:</strong> <span class="con-description">কমিউটার বান্ধব নয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ডুকাটি পানিগালে V4 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳18,00,000 টাকায়, ডুকাটি পানিগালে V4 সংগ্রাহক এবং উত্সাহীদের জন্য যারা চূড়ান্ত ইতালিয়ান রেসিং মেশিন, MotoGP-ডেরাইভড প্রকৌশল, কিংবদন্তি V4 শক্তি এবং মোটরসাইকেল কর্মক্ষমতার বিশুদ্ধ শিল্প খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চূড়ান্ত কর্মক্ষমতা বিশুদ্ধতাবাদীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Ducati Panigale V4 already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Ducati Panigale V4\n")

	return nil
}

// GetName returns the seeder name
func (s *DucatiPanigaleV4ReviewBatch16) GetName() string {
	return "DucatiPanigaleV4ReviewBatch16"
}
