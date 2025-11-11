package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsar200TwinDiscReview handles seeding Bajaj Pulsar 200 Twin Disc product review and translation
type BajajPulsar200TwinDiscReview struct {
	BaseSeeder
}

// NewBajajPulsar200TwinDiscReviewSeeder creates a new BajajPulsar200TwinDiscReview
func NewBajajPulsar200TwinDiscReviewSeeder() *BajajPulsar200TwinDiscReview {
	return &BajajPulsar200TwinDiscReview{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar 200 Twin Disc Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsar200TwinDiscReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar 200 Twin Disc").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar 200 twin disc product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj pulsar 200 twin disc product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar 200 Twin Disc already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar 200 Twin Disc Review Bangladesh 2024 - Sports Performance</h1>
<p class="summary-text">The Bajaj Pulsar 200 Twin Disc is a high-performance sport commuter featuring twin disc brakes, 200cc liquid-cooled engine, aggressive styling, and DTS-i technology. Priced around ৳2,15,000, it delivers exciting performance with safety-focused braking for thrilling daily commutes and weekend rides.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar 200 Twin Disc Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Twin Disc Brakes:</strong> <span class="highlight-value">Front and rear disc safety</span></li>
<li class="highlight-item"><strong class="highlight-label">200cc Liquid-Cooled:</strong> <span class="highlight-value">High-performance DTS-i engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Sporty and muscular design</span></li>
<li class="highlight-item"><strong class="highlight-label">Performance Focused:</strong> <span class="highlight-value">Thrilling acceleration and speed</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar 200 Twin Disc Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Braking:</strong> <span class="pro-description">Twin disc brakes for confident stopping</span></li>
<li class="pro-item"><strong class="pro-label">Strong Performance:</strong> <span class="pro-description">Exciting 200cc acceleration</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Design:</strong> <span class="pro-description">Sporty and eye-catching styling</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">35-40 km/l decent efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Proven DTS-i technology</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar 200 Twin Disc Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Pricing:</strong> <span class="con-description">More expensive than smaller displacement</span></li>
<li class="con-item"><strong class="con-label">Higher Running Cost:</strong> <span class="con-description">Fuel and maintenance higher than 150cc</span></li>
<li class="con-item"><strong class="con-label">Not Budget Friendly:</strong> <span class="con-description">Premium bike for serious enthusiasts</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar 200 Twin Disc Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,15,000 - Performance sport value</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳5-6 per km (35-40 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar 200 Twin Disc Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Strong 200cc power</span></li>
<li class="rating-item"><strong class="rating-label">Braking:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Twin disc excellence</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Aggressive styling</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Sport-oriented seating</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good performance value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar 200 Twin Disc?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,15,000, the Bajaj Pulsar 200 Twin Disc is ideal for riders seeking genuine sports performance with safety-focused twin disc braking and aggressive styling for thrilling daily commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance-oriented commuters</span></p>
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
		return fmt.Errorf("error creating bajaj pulsar 200 twin disc review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar 200 Twin Disc (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার २०० টুইন ডিস্ক রিভিউ বাংলাদেশ २०२४ - স্পোর্টস পারফরমেন্স</h1>
<p class="summary-text">বাজাজ পালসার २०० টুইন ডিস্ক একটি উচ্চ-পারফরমেন্স স্পোর্ট কমিউটার যা টুইন ডিস্ক ব্রেক, २००cc তরল-শীতল ইঞ্জিন, আক্রমণাত্মক স্টাইলিং এবং DTS-i প্রযুক্তি বৈশিষ্ট্যযুক্ত। ৳२,१५,००० টাকায় মূল্যায়িত, এটি নিরাপত্তা-ফোকাস ব্রেকিং সহ উত্তেজনাপূর্ণ পারফরমেন্স প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বाजाज पালसर २०० टुইन डिस्क मुख्य विशेषताएं</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">टुইन डिस्क ব্রেক:</strong> <span class="highlight-value">সামনে এবং পিছনে ডিস্ক নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">२००cc तरल-ठंडा:</strong> <span class="highlight-value">উচ্চ-পারফরমেন্স DTS-i ইঞ្જिন</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">খেলাধুলাপূর্ণ এবং পেশীবহুল ডিজাइন</span></li>
<li class="highlight-item"><strong class="highlight-label">Performance Focused:</strong> <span class="highlight-value">রোমাঞ্চকর ত্বরণ এবং গতি</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাजाज पालसर २००० টুইন ডिস्क সুবिধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার ব্রেকিং:</strong> <span class="pro-description">আত্মবিশ্বাসী থেমে থাকার জন्য টুইन ডিস्क ব्रेक</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী পারফরমেন্স:</strong> <span class="pro-description">উত्तेजनाপूर्ण २००cc ত্বরণ</span></li>
<li class="pro-item"><strong class="pro-label">আক्रामक ডিজাइন:</strong> <span class="pro-description">খেলাধুলাপূর্ণ এবং চোখ ধাঁধানো স্টাইলिং</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">३५-४० km/l আদর্শ দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ्জिन:</strong> <span class="pro-description">প्রमाण्ण DTS-i প्रযুक्তि</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">बाजाज पालसर २००० টুইন ডिস्क असुवидhaएं</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল्य:</strong> <span class="con-description">ছোট স्थานचyuति से বেশি ব्यย়बಹುल</span></li>
<li class="con-item"><strong class="con-label">উচ্চতर চলमান খরज:</strong> <span class="con-description">१५०cc क उच्च ज्वलन और रखरखाव खरz</span></li>
<li class="con-item"><strong class="con-label">ব্যাজেট বান্ধব নয়:</strong> <span class="con-description">गम्भीर उत्साहى के लिए प्रीमियम बाइक</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">चूड़ांत निर्णय: बाजाज पालसर २००० टुइन डिस्क किনें?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ्রिक রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳२,१५,००० টाকায়, বाजाज पालसर २००० টुਇन ডिस्क যারা নিरাপत्তা-ফোকাস টুইन ডਿस्क ব्रेকिং এবং আক্रमণাত्मक स্টাইলिং সহ খाঁটি স्পोर्ट्स पারফরмेन्स খুঁজছেন তাদের জন्য আदर्श।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সुपاริш:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পारফরmens-उन्मुख कमीউटারदের के लिए</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar 200 Twin Disc already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar 200 Twin Disc\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsar200TwinDiscReview) GetName() string {
	return "BajajPulsar200TwinDiscReview"
}
