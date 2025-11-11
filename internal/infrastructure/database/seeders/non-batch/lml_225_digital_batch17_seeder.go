package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// LML-225DigitalReviewBatch17 handles seeding LML 225 Digital product review and translation
type LML225DigitalReviewBatch17 struct {
	BaseSeeder
}

// NewLML225DigitalReviewBatch17Seeder creates a new LML225DigitalReviewBatch17
func NewLML225DigitalReviewBatch17Seeder() *LML225DigitalReviewBatch17 {
	return &LML225DigitalReviewBatch17{BaseSeeder: BaseSeeder{name: "LML 225 Digital Batch17 Review"}}
}

// Seed implements the Seeder interface
func (s *LML225DigitalReviewBatch17) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "LML 225 Digital").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("lml 225 digital product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding lml 225 digital product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for LML 225 Digital already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">LML 225 Digital Review Bangladesh 2024 - Digital Commuter Icon</h1>
<p class="summary-text">The LML 225 Digital is a 225cc digital-gauge sport commuter combining retro heritage with modern instrumentation and accessible performance. Priced around ৳2,65,000, it delivers digital sophistication, responsive engine, unique styling, and excellent value for riders appreciating innovation and digital-first features.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">LML 225 Digital Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">225cc Sport:</strong> <span class="highlight-value">Commuter blend</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Display:</strong> <span class="highlight-value">Modern tech</span></li>
<li class="highlight-item"><strong class="highlight-label">Retro Design:</strong> <span class="highlight-value">Classic appeal</span></li>
<li class="highlight-item"><strong class="highlight-label">Unique Identity:</strong> <span class="highlight-value">Distinctive look</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">LML 225 Digital Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Digital Display:</strong> <span class="pro-description">Modern instrumentation</span></li>
<li class="pro-item"><strong class="pro-label">Retro Styling:</strong> <span class="pro-description">Classic charm</span></li>
<li class="pro-item"><strong class="pro-label">Responsive Engine:</strong> <span class="pro-description">Direct power</span></li>
<li class="pro-item"><strong class="pro-label">Unique Appeal:</strong> <span class="pro-description">Distinctive character</span></li>
<li class="pro-item"><strong class="pro-label">Value Price:</strong> <span class="pro-description">Good budget</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">LML 225 Digital Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Niche Brand:</strong> <span class="con-description">Limited presence</span></li>
<li class="con-item"><strong class="con-label">Service Network:</strong> <span class="con-description">Few dealers</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">Moderate efficiency</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">LML 225 Digital Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,65,000 - Budget sport</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳7-9 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">32-36 km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">LML 225 Digital Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Responsive power</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Stable feel</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Unique retro</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Digital modern</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good budget</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy LML 225 Digital?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,65,000, the LML 225 Digital is perfect for tech-savvy riders wanting a unique 225cc sport commuter with digital sophistication, retro charm, responsive performance, and excellent budget value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For digital-tech riders</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating lml 225 digital review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for LML 225 Digital (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">LML 225 ডিজিটাল রিভিউ বাংলাদেশ 2024 - ডিজিটাল যাত্রী আইকন</h1>
<p class="summary-text">LML 225 ডিজিটাল একটি 225cc ডিজিটাল-গেজ স্পোর্ট যাত্রী যা রেট্রো ঐতিহ্যকে আধুনিক instrumentation এবং অ্যাক্সেসযোগ্য কর্মক্ষমতার সাথে একত্রিত করে। ৳2,65,000 টাকায় মূল্যায়িত, এটি ডিজিটাল পরিশীলনতা, প্রতিক্রিয়াশীল ইঞ্জিন এবং চমৎকার মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">LML 225 ডিজিটাল মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">225cc স্পোর্ট:</strong> <span class="highlight-value">যাত্রী মিশ্রণ</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল ডিসপ্লে:</strong> <span class="highlight-value">আধুনিক প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">রেট্রো ডিজাইন:</strong> <span class="highlight-value">ক্লাসিক আবেদন</span></li>
<li class="highlight-item"><strong class="highlight-label">অনন্য পরিচয়:</strong> <span class="highlight-value">স্বতন্ত্র চেহারা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">LML 225 ডিজিটাল সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ডিজিটাল ডিসপ্লে:</strong> <span class="pro-description">আধুনিক instrumentation</span></li>
<li class="pro-item"><strong class="pro-label">রেট্রো স্টাইলিং:</strong> <span class="pro-description">ক্লাসিক আকর্ষণ</span></li>
<li class="pro-item"><strong class="pro-label">প্রতিক্রিয়াশীল ইঞ্জিন:</strong> <span class="pro-description">সরাসরি শক্তি</span></li>
<li class="pro-item"><strong class="pro-label">অনন্য আবেদন:</strong> <span class="pro-description">স্বতন্ত্র চরিত্র</span></li>
<li class="pro-item"><strong class="pro-label">মূল্য দাম:</strong> <span class="pro-description">ভাল বাজেট</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">LML 225 ডিজিটাল অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">নিশ ব্র্যান্ড:</strong> <span class="con-description">সীমিত উপস্থিতি</span></li>
<li class="con-item"><strong class="con-label">সেবা নেটওয়ার্ক:</strong> <span class="con-description">কয়েকটি ডিলার</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">মধ্যমা দক্ষতা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: LML 225 ডিজিটাল কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,65,000 টাকায়, LML 225 ডিজিটাল প্রযুক্তি-সচেতন চালকদের জন্য নিখুঁত যারা একটি অনন্য 225cc স্পোর্ট যাত্রী চান ডিজিটাল পরিশীলনতা এবং রেট্রো আকর্ষণ সহ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ডিজিটাল-প্রযুক্তি চালকদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for LML 225 Digital already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for LML 225 Digital\n")

	return nil
}

// GetName returns the seeder name
func (s *LML225DigitalReviewBatch17) GetName() string {
	return "LML225DigitalReviewBatch17"
}
