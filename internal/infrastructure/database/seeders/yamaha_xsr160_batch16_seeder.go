package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaXSR160ReviewBatch16 handles seeding Yamaha XSR160 product review and translation
type YamahaXSR160ReviewBatch16 struct {
	BaseSeeder
}

// NewYamahaXSR160ReviewBatch16Seeder creates a new YamahaXSR160ReviewBatch16
func NewYamahaXSR160ReviewBatch16Seeder() *YamahaXSR160ReviewBatch16 {
	return &YamahaXSR160ReviewBatch16{BaseSeeder: BaseSeeder{name: "Yamaha XSR160 Batch16 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaXSR160ReviewBatch16) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha XSR160").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha xsr160 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding yamaha xsr160 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha XSR160 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha XSR160 Review Bangladesh 2024 - Retro Sport Soul</h1>
<p class="summary-text">The Yamaha XSR160 is a retro-styled 160cc motorcycle blending classic design with modern performance. Priced around ৳3,08,000, it delivers Yamaha's sporty DNA with old-school charm, nimble handling, and authentic retro character for riders seeking timeless style.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha XSR160 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc Retro Engine:</strong> <span class="highlight-value">Classic modern blend</span></li>
<li class="highlight-item"><strong class="highlight-label">Retro-Sport Design:</strong> <span class="highlight-value">Old-school modern</span></li>
<li class="highlight-item"><strong class="highlight-label">Nimble Handling:</strong> <span class="highlight-value">Responsive agile</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">45-50 km/l decent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha XSR160 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Retro Design:</strong> <span class="pro-description">Timeless styling</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Performance:</strong> <span class="pro-description">160cc responsive</span></li>
<li class="pro-item"><strong class="pro-label">Nimble Handling:</strong> <span class="pro-description">Agile maneuvers</span></li>
<li class="pro-item"><strong class="pro-label">Yamaha Character:</strong> <span class="pro-description">Sport DNA</span></li>
<li class="pro-item"><strong class="pro-label">Modern Reliability:</strong> <span class="pro-description">Yamaha quality</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha XSR160 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">₳3,08,000 higher</span></li>
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-oriented firm</span></li>
<li class="con-item"><strong class="con-label">Fuel Tank:</strong> <span class="con-description">Small capacity</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha XSR160 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,08,000 - Retro sport</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">45-50 km/l decent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha XSR160 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Retro timeless</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- 160cc sporty</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Nimble agile</span></li>
<li class="rating-item"><strong class="rating-label">Character:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Sport soul</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Premium retro</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha XSR160?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,08,000, the Yamaha XSR160 is ideal for riders seeking a retro-styled bike that combines old-school charm, modern performance, Yamaha's sporty character, nimble handling, and authentic timeless design that stands apart from contemporary trends.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For retro-sport enthusiasts</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating yamaha xsr160 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha XSR160 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা XSR160 রিভিউ বাংলাদেশ 2024 - রেট্রো স্পোর্ট আত্মা</h1>
<p class="summary-text">ইয়ামাহা XSR160 একটি রেট্রো-স্টাইলড 160cc মোটরসাইকেল যা ক্লাসিক ডিজাইন সহ আধুনিক কর্মক্ষমতা মিশ্রিত করে। ৳3,08,000 টাকায় মূল্যায়িত, এটি ইয়ামাহার খেলাধুলা DNA সহ পুরনো-স্কুল আকর্ষণ, চটপটে হ্যান্ডলিং এবং খাঁটি রেট্রো চরিত্র প্রদান করে রাইডারদের জন্য যারা সময়মত স্টাইল খুঁজছেন।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা XSR160 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">160cc রেট্রো ইঞ্জিন:</strong> <span class="highlight-value">ক্লাসিক আধুনিক মিশ্রণ</span></li>
<li class="highlight-item"><strong class="highlight-label">রেট্রো-স্পোর্ট ডিজাইন:</strong> <span class="highlight-value">পুরনো-স্কুল আধুনিক</span></li>
<li class="highlight-item"><strong class="highlight-label">চটপটে হ্যান্ডলিং:</strong> <span class="highlight-value">প্রতিক্রিয়াশীল চটপটে</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">45-50 km/l মানসম্পন্ন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা XSR160 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">রেট্রো ডিজাইন:</strong> <span class="pro-description">সময়মত স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">খেলাধুলা কর্মক্ষমতা:</strong> <span class="pro-description">160cc প্রতিক্রিয়াশীল</span></li>
<li class="pro-item"><strong class="pro-label">চটপটে হ্যান্ডলিং:</strong> <span class="pro-description">চটপটে কৌশল</span></li>
<li class="pro-item"><strong class="pro-label">ইয়ামাহা চরিত্র:</strong> <span class="pro-description">ক্রীড়া DNA</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক নির্ভরযোগ্যতা:</strong> <span class="pro-description">ইয়ামাহা গুণমান</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা XSR160 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳3,08,000 উচ্চতর</span></li>
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">ক্রীড়া-ফোকাসড দৃঢ়</span></li>
<li class="con-item"><strong class="con-label">জ্বালানী ট্যাঙ্ক:</strong> <span class="con-description">ছোট ক্ষমতা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা XSR160 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,08,000 টাকায়, ইয়ামাহা XSR160 রাইডারদের জন্য আদর্শ যারা একটি রেট্রো-স্টাইলড বাইক খুঁজছেন যা পুরনো-স্কুল আকর্ষণ, আধুনিক কর্মক্ষমতা, ইয়ামাহার খেলাধুলা চরিত্র, চটপটে হ্যান্ডলিং এবং খাঁটি সময়মত ডিজাইন একত্রিত করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- রেট্রো-খেলাধুলা উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha XSR160 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha XSR160\n")

	return nil
}

// GetName returns the seeder name
func (s *YamahaXSR160ReviewBatch16) GetName() string {
	return "YamahaXSR160ReviewBatch16"
}
