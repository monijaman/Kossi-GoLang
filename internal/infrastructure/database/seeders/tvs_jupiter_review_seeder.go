package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSJupiterReviewSeeder handles seeding TVS Jupiter product review and translation
type TVSJupiterReviewSeeder struct {
	BaseSeeder
}

// NewTVSJupiterReviewSeeder creates a new TVSJupiterReviewSeeder
func NewTVSJupiterReviewSeeder() *TVSJupiterReviewSeeder {
	return &TVSJupiterReviewSeeder{BaseSeeder: BaseSeeder{name: "TVS Jupiter Review"}}
}

// Seed implements the Seeder interface
func (s *TVSJupiterReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Jupiter").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("TVS Jupiter product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding TVS Jupiter product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Jupiter already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Jupiter Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The TVS Jupiter is a family-oriented 110cc scooter priced at ৳1,85,000, offering comfortable seating, large storage, and smooth ride quality. With features like external fuel filler, mobile charging port, and good build quality, it's ideal for family use and daily commuting with practical features.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Jupiter Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Family Comfort:</strong> <span class="highlight-value">Comfortable for family riding</span></li>
<li class="highlight-item"><strong class="highlight-label">Large Storage:</strong> <span class="highlight-value">Spacious under-seat storage</span></li>
<li class="highlight-item"><strong class="highlight-label">External Fuel Cap:</strong> <span class="highlight-value">Convenient fuel filling</span></li>
<li class="highlight-item"><strong class="highlight-label">Mobile Charging:</strong> <span class="highlight-value">Built-in USB charging port</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Jupiter Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Comfort:</strong> <span class="pro-description">Plush seat, smooth ride quality</span></li>
<li class="pro-item"><strong class="pro-label">Good Storage Space:</strong> <span class="pro-description">21L under-seat storage capacity</span></li>
<li class="pro-item"><strong class="pro-label">External Fuel Filler:</strong> <span class="pro-description">No need to open seat for refueling</span></li>
<li class="pro-item"><strong class="pro-label">Mobile Charging Port:</strong> <span class="pro-description">USB port for phone charging</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Solid construction and materials</span></li>
<li class="pro-item"><strong class="pro-label">Decent Mileage:</strong> <span class="pro-description">50-55 km/l fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Family Friendly:</strong> <span class="pro-description">Perfect for family use</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Jupiter Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Average Performance:</strong> <span class="con-description">110cc not very powerful</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">TVS service centers not widespread</span></li>
<li class="con-item"><strong class="con-label">Lower Resale Value:</strong> <span class="con-description">Not as good as Honda scooters</span></li>
<li class="con-item"><strong class="con-label">Plastic Quality:</strong> <span class="con-description">Some panels feel cheap</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy TVS Jupiter in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Family use and grocery shopping</li>
<li class="best-for-item">Comfortable daily commuting</li>
<li class="best-for-item">Users prioritizing practical features</li>
<li class="best-for-item">Those needing mobile charging on-the-go</li>
<li class="best-for-item">City riding with passenger</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy TVS Jupiter?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Long highway journeys</li>
<li class="not-recommended-item">Those prioritizing resale value</li>
<li class="not-recommended-item">Areas without TVS service centers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Jupiter Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,80,000 - ৳1,90,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,600 - ৳3,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,000-2,400/month for 30 km daily at 52 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳900 - ৳1,300 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Jupiter Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Reliable</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Great practical features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Simple family look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">TVS Jupiter Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of TVS Jupiter?</h3>
<p class="faq-answer">TVS Jupiter delivers 50-55 km/l in real-world conditions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of TVS Jupiter?</h3>
<p class="faq-answer">TVS Jupiter is priced between ৳1,80,000 to ৳1,90,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is TVS Jupiter good for family use?</h3>
<p class="faq-answer">Yes, excellent choice with comfortable seating and practical family features.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Jupiter or Honda Dio?</h3>
<p class="faq-answer">Jupiter offers better comfort and features, Dio has better resale value.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Jupiter in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">TVS Jupiter is an excellent family scooter at ৳1,85,000, offering outstanding comfort, practical features like external fuel cap and USB charging, and good build quality. While 110cc performance is average and TVS service network is limited compared to Honda, the comfort-first design and family-friendly features make it ideal for grocery runs, school drops, and comfortable city commuting. Best value for families prioritizing comfort over performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For comfort-focused family use</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.1,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating TVS Jupiter review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Jupiter (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">টিভিএস জুপিটার রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">টিভিএস জুপিটার একটি পরিবার-ভিত্তিক ১১০সিসি স্কুটার যার মূল্য ৳১,৮৫,০০০ টাকা, যা আরামদায়ক সিটিং, বড় স্টোরেজ এবং মসৃণ রাইড কোয়ালিটি প্রদান করে। এক্সটার্নাল ফুয়েল ফিলার, মোবাইল চার্জিং পোর্ট এবং ভালো বিল্ড কোয়ালিটির মতো ফিচার সহ, এটি ব্যবহারিক ফিচার সহ পারিবারিক ব্যবহার এবং দৈনিক যাতায়াতের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস জুপিটার এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">পারিবারিক আরাম:</strong> <span class="highlight-value">পারিবারিক রাইডিংয়ের জন্য আরামদায়ক</span></li>
<li class="highlight-item"><strong class="highlight-label">বড় স্টোরেজ:</strong> <span class="highlight-value">প্রশস্ত সিটের নিচে স্টোরেজ</span></li>
<li class="highlight-item"><strong class="highlight-label">এক্সটার্নাল ফুয়েল ক্যাপ:</strong> <span class="highlight-value">সুবিধাজনক জ্বালানি ভর্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">মোবাইল চার্জিং:</strong> <span class="highlight-value">বিল্ট-ইন ইউএসবি চার্জিং পোর্ট</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস জুপিটার এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার আরাম:</strong> <span class="pro-description">নরম সিট, মসৃণ রাইড কোয়ালিটি</span></li>
<li class="pro-item"><strong class="pro-label">ভালো স্টোরেজ স্পেস:</strong> <span class="pro-description">২১ লিটার সিটের নিচে স্টোরেজ ক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">এক্সটার্নাল ফুয়েল ফিলার:</strong> <span class="pro-description">জ্বালানি ভর্তির জন্য সিট খোলার প্রয়োজন নেই</span></li>
<li class="pro-item"><strong class="pro-label">মোবাইল চার্জিং পোর্ট:</strong> <span class="pro-description">ফোন চার্জিংয়ের জন্য ইউএসবি পোর্ট</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">শক্ত নির্মাণ এবং ম্যাটেরিয়াল</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৫০-৫৫ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">পরিবার বান্ধব:</strong> <span class="pro-description">পারিবারিক ব্যবহারের জন্য নিখুঁত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস জুপিটার এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">গড় পারফরম্যান্স:</strong> <span class="con-description">১১০সিসি খুব শক্তিশালী নয়</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">টিভিএস সার্ভিস সেন্টার ব্যাপক নয়</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">হোন্ডা স্কুটারের মতো ভালো নয়</span></li>
<li class="con-item"><strong class="con-label">প্লাস্টিকের মান:</strong> <span class="con-description">কিছু প্যানেল সস্তা মনে হয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে টিভিএস জুপিটার কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Family use and grocery shopping</li>
<li class="best-for-item">Comfortable daily commuting</li>
<li class="best-for-item">Users prioritizing practical features</li>
<li class="best-for-item">Those needing mobile charging on-the-go</li>
<li class="best-for-item">City riding with passenger</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">টিভিএস জুপিটার কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Long highway journeys</li>
<li class="not-recommended-item">Those prioritizing resale value</li>
<li class="not-recommended-item">Areas without TVS service centers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">টিভিএস জুপিটার এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৮০,০০০ - ৳১,৯০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৬০০ - ৳৩,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,০০০-২,৪০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৯০০ - ৳১,৩০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">টিভিএস জুপিটার পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- নির্ভরযোগ্য</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- দুর্দান্ত ব্যবহারিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- সাধারণ পারিবারিক লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">টিভিএস জুপিটার সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">টিভিএস জুপিটারের মাইলেজ কত?</h3>
<p class="faq-answer">টিভিএস জুপিটার বাস্তব পরিস্থিতিতে ৫০-৫৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">টিভিএস জুপিটারের দাম কত?</h3>
<p class="faq-answer">টিভিএস জুপিটারের দাম ৳১,৮০,০০০ থেকে ৳১,৯০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">পারিবারিক ব্যবহারের জন্য টিভিএস জুপিটার কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, আরামদায়ক সিটিং এবং ব্যবহারিক পারিবারিক ফিচার সহ চমৎকার পছন্দ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: জুপিটার নাকি হোন্ডা ডায়ো?</h3>
<p class="faq-answer">জুপিটার ভালো আরাম এবং ফিচার প্রদান করে, ডায়োর ভালো রিসেল ভ্যালু আছে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে টিভিএস জুপিটার কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.1/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">টিভিএস জুপিটার ৳১,৮৫,০০০ টাকায় একটি চমৎকার পারিবারিক স্কুটার, যা অসাধারণ আরাম, এক্সটার্নাল ফুয়েল ক্যাপ এবং ইউএসবি চার্জিংয়ের মতো ব্যবহারিক ফিচার এবং ভালো বিল্ড কোয়ালিটি প্রদান করে। যদিও ১১০সিসি পারফরম্যান্স গড় এবং টিভিএস সার্ভিস নেটওয়ার্ক হোন্ডার তুলনায় সীমিত, আরাম-প্রথম ডিজাইন এবং পরিবার-বান্ধব ফিচার এটিকে মুদি কেনাকাটা, স্কুল ড্রপ এবং আরামদায়ক শহর যাতায়াতের জন্য আদর্শ করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আরাম-কেন্দ্রিক পারিবারিক ব্যবহারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Jupiter already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Jupiter\n")

	return nil
}
