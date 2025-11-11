package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsar150TwinDiscReview handles seeding Bajaj Pulsar 150 Twin Disc product review and translation
type BajajPulsar150TwinDiscReview struct {
	BaseSeeder
}

// NewBajajPulsar150TwinDiscReviewSeeder creates a new BajajPulsar150TwinDiscReview
func NewBajajPulsar150TwinDiscReviewSeeder() *BajajPulsar150TwinDiscReview {
	return &BajajPulsar150TwinDiscReview{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar 150 Twin Disc Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsar150TwinDiscReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar 150 Twin Disc").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Pulsar 150 Twin Disc product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Pulsar 150 Twin Disc product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar 150 Twin Disc already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar 150 Twin Disc Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Pulsar 150 Twin Disc is a sporty commuter motorcycle featuring twin disc brakes for enhanced safety, 150cc DTS-i engine with digital fuel injection, stylish aggressive design, and excellent mileage. Priced at ৳1,65,000, it offers sporty performance with the confidence of twin brakes, good fuel efficiency, distinctive styling, and Bajaj reliability making it ideal for riders seeking performance and safety in a stylish commuter package.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar 150 Twin Disc Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Twin Disc Brakes:</strong> <span class="highlight-value">Front and rear disc brakes for superior stopping power</span></li>
<li class="highlight-item"><strong class="highlight-label">150cc Engine:</strong> <span class="highlight-value">Digital fuel injection for better performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Sporty Design:</strong> <span class="highlight-value">Aggressive styling with LED headlight</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">38-42 km/l fuel efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar 150 Twin Disc Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Braking:</strong> <span class="pro-description">Twin disc brakes provide superior stopping power and control</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Performance:</strong> <span class="pro-description">150cc DTS-i engine delivers good acceleration and speed</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Modern design stands out with LED headlight</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">38-42 km/l offers decent fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Brand:</strong> <span class="pro-description">Bajaj's proven track record</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar 150 Twin Disc Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">Twin disc setup increases cost</span></li>
<li class="con-item"><strong class="con-label">Heavier Build:</strong> <span class="con-description">Slightly heavier than single disc variant</span></li>
<li class="con-item"><strong class="con-label">No ABS:</strong> <span class="con-description">Despite twin discs, no ABS safety feature</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar 150 Twin Disc Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,65,000 - Premium for twin disc safety</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳5-6 per km (38-42 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar 150 Twin Disc Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good commuter performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Decent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Braking:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Excellent twin disc brakes</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Aggressive sporty design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good value for features</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar 150 Twin Disc?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,65,000, the Bajaj Pulsar 150 Twin Disc offers sporty performance with superior braking safety through dual disc brakes, good fuel economy, and aggressive modern styling. Perfect for riders who want performance and safety in an affordable commuter package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sporty performance with dual disc safety</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Bajaj Pulsar 150 Twin Disc review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar 150 Twin Disc (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার ১৫০ টুইন ডিস্ক রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ পালসার ১৫০ টুইন ডিস্ক একটি স্পোর্টি কমিউটার মোটরসাইকেল যা টুইন ডিস্ক ব্রেক, ১৫০সিসি ডিটিএস-আই ইঞ্জিন, ডিজিটাল ফুয়েল ইনজেকশন এবং আক্রমণাত্মক ডিজাইন সহ আসে। ৳১,৬৫,০০০ টাকায় মূল্যায়িত, এটি দ্বৈত ব্রেকের আত্মবিশ্বাস সহ স্পোর্টি পারফরমেন্স, ভাল জ্বালানি দক্ষতা এবং স্বতন্ত্র স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার ১৫০ টুইন ডিস্ক মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">টুইন ডিস্ক ব্রেক:</strong> <span class="highlight-value">উন্নত ব্রেকিং শক্তির জন্য সামনের এবং পিছনের ডিস্ক</span></li>
<li class="highlight-item"><strong class="highlight-label">১৫০সিসি ইঞ্জিন:</strong> <span class="highlight-value">ভাল পারফরমেন্সের জন্য ডিজিটাল জ্বালানি ইনজেকশন</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্টি ডিজাইন:</strong> <span class="highlight-value">এলইডি হেডলাইট সহ আক্রমণাত্মক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">৩৮-৪২ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার ১৫০ টুইন ডিস্ক সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার ব্রেকিং:</strong> <span class="pro-description">টুইন ডিস্ক ব্রেক উচ্চতর ব্রেকিং শক্তি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি পারফরমেন্স:</strong> <span class="pro-description">ভাল ত্বরণ এবং গতি</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">আধুনিক নকশা এলইডি হেডলাইট সহ</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">ভালো জ্বালানি দক্ষতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার ১৫০ টুইন ডিস্ক অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">টুইন ডিস্ক সেটআপ খরচ বৃদ্ধি করে</span></li>
<li class="con-item"><strong class="con-label">ভারী বিল্ড:</strong> <span class="con-description">সিঙ্গেল ডিস্ক সংস্করণের চেয়ে ভারী</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাজাজ পালসার ১৫০ টুইন ডিস্ক কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳১,৬৫,০০০ টাকায় বাজাজ পালসার ১৫০ টুইন ডিস্ক দ্বৈত ডিস্ক ব্রেকের মাধ্যমে স্পোর্টি পারফরমেন্স এবং উচ্চতর ব্রেকিং নিরাপত্তা প্রদান করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্পোর্টি পারফরমেন্স এবং দ্বৈত ডিস্ক নিরাপত্তার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar 150 Twin Disc already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar 150 Twin Disc\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsar150TwinDiscReview) GetName() string {
	return "BajajPulsar150TwinDiscReview"
}
