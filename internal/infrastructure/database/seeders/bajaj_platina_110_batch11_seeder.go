package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPlatina110Review handles seeding Bajaj Platina 110 product review and translation
type BajajPlatina110Review struct {
	BaseSeeder
}

// NewBajajPlatina110ReviewSeeder creates a new BajajPlatina110Review
func NewBajajPlatina110ReviewSeeder() *BajajPlatina110Review {
	return &BajajPlatina110Review{BaseSeeder: BaseSeeder{name: "Bajaj Platina 110 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPlatina110Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Platina 110").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj platina 110 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj platina 110 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Platina 110 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Platina 110 Review Bangladesh 2024 - Ultimate Economy King</h1>
<p class="summary-text">The Bajaj Platina 110 is the ultimate budget commuter featuring 110cc air-cooled engine, minimal design, incredible fuel efficiency, and unbeatable price. Priced around ৳70,000, it's the cheapest and most economical motorcycle for first-time buyers and daily commuters seeking pure value.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Platina 110 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc Air-Cooled:</strong> <span class="highlight-value">Simple proven engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Ultra Low Cost:</strong> <span class="highlight-value">Most economical option</span></li>
<li class="highlight-item"><strong class="highlight-label">Minimal Design:</strong> <span class="highlight-value">No-frills practicality</span></li>
<li class="highlight-item"><strong class="highlight-label">Entry-Level Perfect:</strong> <span class="highlight-value">First bike option</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Platina 110 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Unbeatable Price:</strong> <span class="pro-description">Most affordable motorcycle available</span></li>
<li class="pro-item"><strong class="pro-label">Insane Mileage:</strong> <span class="pro-description">65-70 km/l unmatched efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Zero Frills:</strong> <span class="pro-description">Simple maintenance and repairs</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Proven 110cc platform</span></li>
<li class="pro-item"><strong class="pro-label">Easy to Service:</strong> <span class="pro-description">Cheap and fast repairs</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Platina 110 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Ultra Basic:</strong> <span class="con-description">Extremely minimal features</span></li>
<li class="con-item"><strong class="con-label">No ABS/Disc:</strong> <span class="con-description">Drum brakes only</span></li>
<li class="con-item"><strong class="con-label">Very Low Power:</strong> <span class="con-description">Adequate for city only</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Platina 110 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳70,000 - Absolute budget minimum</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Ultra Low - ৳1.50-2 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳1.50-2 per km (65-70 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Platina 110 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Phenomenal 65-70 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Price:</strong> <span class="rating-score">5.0</span> <span class="rating-note">- Unbeatable value</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Simple and dependable</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Minimal comfort features</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- Adequate city riding</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Platina 110?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳70,000, the Bajaj Platina 110 is the ultimate budget motorcycle for first-time buyers, students, and economy commuters. Unbeatable price, insane mileage, and minimal running cost make it the perfect entry-level bike.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For first-time buyers and extreme budget riders</span></p>
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
		return fmt.Errorf("error creating bajaj platina 110 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Platina 110 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ প্লাটিনা ११० রিভিউ বাংলাদেশ २०२४ - চূড়ান্ত অর্থনীতি রাজা</h1>
<p class="summary-text">বাজাজ প্লাটিনা ११० চূড়ান্ত বাজেট কমিউটার যা ११०cc এয়ার-কুল্ড ইঞ્જিन, ন্यূनतम ডিজাইন, অবিশ্বাস্য জ্বালানি দক্ষতা এবং অপ্রতিদ্বন্দ্বী মূল্য বৈশিষ্ট্যযুক्त। ৳७०,००० টাকায় মূল्यায়িত, এটি প্রথমবারের ক্রেতা এবং দৈनिक कमीউটারদের জন्য সবচেয়ে সস্তা এবং সবচেয়ে অর্থনৈতिक মোটরসाइকेল।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">बाजाज पLatina ११० मुख्य विशेषताएं</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">११०cc एयर-कुल्ड:</strong> <span class="highlight-value">সহজ প्রমाণিত ইঞ્જिন</span></li>
<li class="highlight-item"><strong class="highlight-label">अत्यधिक कम लागत:</strong> <span class="highlight-value">সবচেয়ে অর্থনৈতिक বিকल्प</span></li>
<li class="highlight-item"><strong class="highlight-label">न्यूनतम ডিজাइন:</strong> <span class="highlight-value">कोई frills व्यावहारिकता नहीं</span></li>
<li class="highlight-item"><strong class="highlight-label">এন्ট्রি-লেभেल Perfect:</strong> <span class="highlight-value">প्रथम बাइক विकल्प</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">बाजाज पLatina ११० फayदे</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অপ্রতিদ্वन্দ्वী মূल্य:</strong> <span class="pro-description">সবচেয়ে সাশ्रয়ী মোটরसাइকेल উপलब्ध</span></li>
<li class="pro-item"><strong class="pro-label">पागल मीलेज:</strong> <span class="pro-description">अতুलनীয় ६५-७० km/l দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">শূন्य Frills:</strong> <span class="pro-description">সহজ রক्ষणाबेक्षण এবং মেरामत</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভरযোগ्य ইঞ्जиन:</strong> <span class="pro-description">প्রमाण्ण ११०cc প्लेटफর्ম</span></li>
<li class="pro-item"><strong class="pro-label">সেवा করা সহজ:</strong> <span class="pro-description">সস्ते और দ्রुत मरम्मत</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">बाजाज पLatina ११० असुビधा</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">अल्ट्रा Basic:</strong> <span class="con-description">অত्यন্ত न्यूनतम বৈশিষ્ट्य</span></li>
<li class="con-item"><strong class="con-label">কোন ABS/Disc নেই:</strong> <span class="con-description">শুধুमাত्र ড्রাम ব्रेक</span></li>
<li class="con-item"><strong class="con-label">অত্यন্ত कम শক्তি:</strong> <span class="con-description">শহরের জন्য पर्याप्त</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চूড়ांत সिद्धांत: बाजाज पLatina ११० किनें?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সामগ्রिक रेটिং:</strong> <span class="score-value">४.५/५.०</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳७०,००० টাকায়, बाजाज पLatina ११० প्रथमबারের ক्रेতा, पराणी और অর्थনीथिक कमीউटারদের जন्य चूड़ांত बजট मोटरसाइकेল। अप्रতिद्वंद्वी मूल्य, पागल मीलेज, और न्यूनतम चलमान लागत इसे নিখुঁত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">सुपारिश:</strong> <span class="badge badge-success">हां</span> <span class="recommendation-for">- প্রथমবারের ক্রেতা এবং চরম বাজেট রাইডারদের জন्य</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Platina 110 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Platina 110\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPlatina110Review) GetName() string {
	return "BajajPlatina110Review"
}
