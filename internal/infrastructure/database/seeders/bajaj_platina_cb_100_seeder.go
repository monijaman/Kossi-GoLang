package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPlatinaCB100Review handles seeding Bajaj Platina CB 100 product review and translation
type BajajPlatinaCB100Review struct {
	BaseSeeder
}

// NewBajajPlatinaCB100ReviewSeeder creates a new BajajPlatinaCB100Review
func NewBajajPlatinaCB100ReviewSeeder() *BajajPlatinaCB100Review {
	return &BajajPlatinaCB100Review{BaseSeeder: BaseSeeder{name: "Bajaj Platina CB 100 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPlatinaCB100Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Platina CB 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Platina CB 100 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Platina CB 100 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Platina CB 100 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Platina CB 100 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Platina CB 100 is a solid, reliable commuter motorcycle combining affordability with dependability. Featuring 100cc engine, good fuel economy up to 60 km/l, practical design, and rock-solid build quality. Priced at ৳78,000, it's perfect for first-time buyers and daily commuters who value reliability over modern features.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Platina CB 100 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">Up to 60 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">100cc Engine:</strong> <span class="highlight-value">Simple and proven technology</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seat:</strong> <span class="highlight-value">Decent ergonomics for daily commute</span></li>
<li class="highlight-item"><strong class="highlight-label">Reliable Design:</strong> <span class="highlight-value">Time-tested motorcycle with proven durability</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Platina CB 100 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Great Mileage:</strong> <span class="pro-description">60 km/l excellent fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Very Affordable:</strong> <span class="pro-description">Starting at just ৳78,000</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Machine:</strong> <span class="pro-description">Proven track record of durability</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Affordable parts and service</span></li>
<li class="pro-item"><strong class="pro-label">First-Bike Friendly:</strong> <span class="pro-description">Perfect for new motorcycle owners</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Platina CB 100 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Design:</strong> <span class="con-description">Outdated styling, very conventional</span></li>
<li class="con-item"><strong class="con-label">Limited Power:</strong> <span class="con-description">100cc offers minimal performance</span></li>
<li class="con-item"><strong class="con-label">Minimal Features:</strong> <span class="con-description">No modern amenities or technology</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Platina CB 100 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳78,000 - Best budget commuter option</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very Low - ৳2.8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2.8 per km (60 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Platina CB 100 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Highly dependable</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good for commuting</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Exceptional value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Platina CB 100?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳78,000, the Bajaj Platina CB 100 is an unbeatable value for first-time motorcycle owners and budget-conscious commuters. With excellent reliability, great mileage, and affordable maintenance, it's the practical choice over style-focused alternatives.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For reliable budget commuting and first-time buyers</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.2,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Bajaj Platina CB 100 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Platina CB 100 (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ প্লাটিনা সিবি १०० রিভিউ বাংলাদেশ २०२४ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ প্লাটিনা CB १०० একটি দৃঢ়, নির্ভরযোগ্য কমিউটার মোটরসাইকেল যা সামর্থ্য এবং নির্ভরযোগ্যতা একত্রিত করে। १००সিসি ইঞ্জিন, ভাল জ্বালানি অর্থনীতি এবং অনুশীলনী ডিজাইন সহ। ৳७८,००० টাকায় মূল্যায়িত, এটি প্রথমবারের ক্রেতা এবং দৈনন্দিন যাতায়াতকারীদের জন্য নিখুঁত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ প্লাটিনা সিবি १००मূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">६० किमी/लिटor জ্বালানি দক্ষতা পর্যন্ত</span></li>
<li class="highlight-item"><strong class="highlight-label">१००सिcc ইঞ্জিন:</strong> <span class="highlight-value">সহজ এবং প্রমাণিত প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক সিট:</strong> <span class="highlight-value">দৈনন্দিন যাতায়াতের জন্য ভালো এরগোনমিক্স</span></li>
<li class="highlight-item"><strong class="highlight-label">নির্ভরযোগ্য ডিজাইন:</strong> <span class="highlight-value">সময়-পরীক্ষিত মোটরসাইকেল</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ প্লাটিনা সিবি १००सुवিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">দুর्দांत मাइलेज:</strong> <span class="pro-description">६० किमी/लिटor उत्कृष्ट ईंधन দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">অত्যন্ত সাশ্রয়ী:</strong> <span class="pro-description">মাত्र ৳७८,००० থেকে শুরু</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য মেশিন:</strong> <span class="pro-description">স্থায়িত্বের প্রমাণিত ট্র্যাক রেকর্ড</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সাশ্রয়ী যন্ত্রাংশ এবং সেবা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজाज प्लेटिना CB १००অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক ডিজাইন:</strong> <span class="con-description">পুরানো স্টাইলিং, অত্যন্ত প্রচলিত</span></li>
<li class="con-item"><strong class="con-label">সীমিত শক্তি:</strong> <span class="con-description">१००सिcc সীমিত পারফরমেন্স প্রদান করে</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত रায়: बাजाজ प्लेटिना CB १००किनबেन?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ्रिक रेटिং:</strong> <span class="score-value">४.२/५.०</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳७८,००० টাকায়, বाजाज प्लेटिना CB १००প्रথमবার मোटरसाइकिल مালिকদের এবং বাজেট-সচেतন যাতায়াতকারীদের জন্য একটি অপ্রতিরোধ্য মূল্য। চমৎকার নির্ভরযোগ্যতা এবং মাইলেজের সাথে, এটি স্টাইল-কেন্দ্রিক বিকল্পের চেয়ে অনুশীলনী পছন্দ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">सुপारिश:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- নির্ভরযোগ্য বাজেট যাতায়াত এবং প্রথমবার ক্রেতাদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Platina CB 100 already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating Bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Platina CB 100\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPlatinaCB100Review) GetName() string {
	return "BajajPlatinaCB100Review"
}
