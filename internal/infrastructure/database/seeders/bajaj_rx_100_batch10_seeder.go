package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajRX100Review handles seeding Bajaj RX 100 product review and translation
type BajajRX100Review struct {
	BaseSeeder
}

// NewBajajRX100ReviewSeeder creates a new BajajRX100Review
func NewBajajRX100ReviewSeeder() *BajajRX100Review {
	return &BajajRX100Review{BaseSeeder: BaseSeeder{name: "Bajaj RX 100 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajRX100Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj RX 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj RX 100 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj RX 100 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj RX 100 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj RX 100 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj RX 100 is a legendary, highly affordable commuter bike featuring 100cc two-stroke engine, minimal maintenance, exceptional fuel efficiency up to 70 km/l, and legendary reliability. Priced at ৳65,000, it remains the ultimate budget commuter for rural and semi-urban riders. Perfect for those prioritizing cost-effectiveness and proven durability over modern features.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj RX 100 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Exceptional Mileage:</strong> <span class="highlight-value">70 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">100cc Two-Stroke:</strong> <span class="highlight-value">Simple and rugged engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Minimal Maintenance:</strong> <span class="highlight-value">Easy to maintain and repair</span></li>
<li class="highlight-item"><strong class="highlight-label">Legendary Reliability:</strong> <span class="highlight-value">Time-tested workhorse</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj RX 100 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Exceptional Mileage:</strong> <span class="pro-description">70 km/l incredible fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Ultra Affordable:</strong> <span class="pro-description">Starting at just ৳65,000</span></li>
<li class="pro-item"><strong class="pro-label">Legendary Reliability:</strong> <span class="pro-description">Proven workhorse for decades</span></li>
<li class="pro-item"><strong class="pro-label">Easy Maintenance:</strong> <span class="pro-description">Simple two-stroke engine design</span></li>
<li class="pro-item"><strong class="pro-label">Minimal Running Cost:</strong> <span class="pro-description">Cheapest cost per km among all bikes</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj RX 100 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Outdated Technology:</strong> <span class="con-description">No modern features whatsoever</span></li>
<li class="con-item"><strong class="con-label">Two-Stroke Emissions:</strong> <span class="con-description">Higher pollution compared to four-stroke</span></li>
<li class="con-item"><strong class="con-label">Basic Design:</strong> <span class="con-description">Very basic, utilitarian styling</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj RX 100 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳65,000 - The cheapest bike available</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Minimal - ৳2 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2 per km (70 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj RX 100 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Unmatched fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Legendary durability</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Best value available</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj RX 100?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳65,000, the Bajaj RX 100 is the absolute cheapest bike with legendary reliability and unmatched mileage. Ideal for budget-conscious rural riders who need maximum economy and proven durability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultimate budget commuting with proven reliability</span></p>
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
		return fmt.Errorf("error creating Bajaj RX 100 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj RX 100 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ আরএক্স १०० রিভিউ বাংলাদেশ २०२४ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ আরএক्स १०० একটি কিংবদন্তি, অত্যন্ত সাশ্রয়ী কমিউটার বাইক যা १००cc টু-স্ট্রোক ইঞ్జिन, न्यूनतम রক্ষণাবেক্ষণ এবং ७० किमी/लिटر ব্যতিক্রমী জ्वালानी দক्षता সহ আসে। ৳६५,००० টাকায় মূল্যায়িত, এটি গ্রামীণ এবং আধা-শহুরे রাইডারদের জন্য চূড়ান्त বাজেט কমিউটার থেকে যায়।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বाजाज आरएक्स १००মূল बৈशिष्ट्य</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">असाधारण मాइलेज:</strong> <span class="highlight-value">७० किमी/लिटर জ্বালানি দক्षता</span></li>
<li class="highlight-item"><strong class="highlight-label">१००सिcc टू-स्ट्रोक:</strong> <span class="highlight-value">সহজ এবং মজবুত ইঞ્જિन</span></li>
<li class="highlight-item"><strong class="highlight-label">न्यूनतम रक्षणाबेক्षण:</strong> <span class="highlight-value">রক্ষণাবেक্ষণ এবং মেরামত করা সহজ</span></li>
<li class="highlight-item"><strong class="highlight-label">কিंवদंति नір्भरयोग्यता:</strong> <span class="highlight-value">সময়-পরीক्षিত কর্মক्षेत्र</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বाजाज आरএक्स १००সुवিधा</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">अসाधारण mাইलेज:</strong> <span class="pro-description">७० किमी/लिटr অবিশ्वাস्য জ్వালानी দক్ષता</span></li>
<li class="pro-item"><strong class="pro-label">অতি সাশ्रয়ী:</strong> <span class="pro-description">মাত्र ৳६५,००० থেকে শুरू</span></li>
<li class="pro-item"><strong class="pro-label">किंवदंति নir্भरयోग्य:</strong> <span class="pro-description">দশক ধরে প्रमाणित কর्मक्षेत्र</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক्षणাবेक्षण:</strong> <span class="pro-description">সহজ টু-स्ट्रোक ইঞ్జिन ডিজाइন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">बाजाज आरएक्स १००अपूर्णताएं</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">पुरानо প्रযुক्তि:</strong> <span class="con-description">कोই आधुनिक বৈশिষ्ट्য नाহิ</span></li>
<li class="con-item"><strong class="con-label">টু-স्ট्রोक उत्सर्जन:</strong> <span class="con-description">चार-stroke स्वरूप एवi उच्च प्रदूষण</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চूড়ांत निर्णय: বाजाज আरএक्स १०० किनें?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামग्रिक रेটिং:</strong> <span class="score-value">४.६/५.०</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳६५,००० টাকায়, বাজाज আरএक्स १००সম্পূर্ণভাবে সবচেয়ে সস্তা বাইক যার কিंवदंti नir्भरয्याग्यता र छि। ग्रामीण রাইডারদের জন्য আদರ্শ যারা সর्वाधिक अर्थনीতি প्रয়োজन।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">सुपारिश:</strong> <span class="badge badge-success">हाँ</span> <span class="recommendation-for">- চূড়ান্त বাজেট যাতায়াতের জন्য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj RX 100 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj RX 100\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajRX100Review) GetName() string {
	return "BajajRX100Review"
}
