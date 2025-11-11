package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiHayateReviewSeeder handles seeding Suzuki Hayate product review and translation
type SuzukiHayateReviewSeeder struct {
	BaseSeeder
}

// NewSuzukiHayateReviewSeeder creates a new SuzukiHayateReviewSeeder
func NewSuzukiHayateReviewSeeder() *SuzukiHayateReviewSeeder {
	return &SuzukiHayateReviewSeeder{BaseSeeder: BaseSeeder{name: "Suzuki Hayate Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiHayateReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Hayate").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Hayate product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Hayate product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Hayate already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Hayate Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Hayate is a budget-friendly 113cc commuter motorcycle priced at ৳1,20,000. With simple design, decent fuel efficiency, and Suzuki reliability, it's a practical choice for cost-conscious daily commuters needing basic transportation without frills.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Hayate Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Affordable Price:</strong> <span class="highlight-value">Budget-friendly at ৳1,20,000</span></li>
<li class="highlight-item"><strong class="highlight-label">Decent Mileage:</strong> <span class="highlight-value">55-60 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Simple Maintenance:</strong> <span class="highlight-value">Easy to maintain, cheap parts</span></li>
<li class="highlight-item"><strong class="highlight-label">Lightweight:</strong> <span class="highlight-value">Easy to handle in traffic</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Hayate Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Very Affordable:</strong> <span class="pro-description">At ৳1,20,000 it's one of cheapest bikes</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">55-60 km/l saves fuel costs</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight Handling:</strong> <span class="pro-description">Easy to maneuver in city traffic</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance Cost:</strong> <span class="pro-description">Simple design, cheap spare parts</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Suzuki Japanese engineering</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Convenient self-start option</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Hayate Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Basic Features:</strong> <span class="con-description">No modern features at all</span></li>
<li class="con-item"><strong class="con-label">Drum Brakes Only:</strong> <span class="con-description">No disc brake option</span></li>
<li class="con-item"><strong class="con-label">Poor Build Quality:</strong> <span class="con-description">Plastic quality is below average</span></li>
<li class="con-item"><strong class="con-label">Outdated Styling:</strong> <span class="con-description">Old-fashioned design</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Suzuki centers limited</span></li>
<li class="con-item"><strong class="con-label">Weak Performance:</strong> <span class="con-description">113cc very underpowered</span></li>
<li class="con-item"><strong class="con-label">Low Resale Value:</strong> <span class="con-description">Suzuki bikes depreciate fast</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Hayate in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extremely budget-conscious buyers</li>
<li class="best-for-item">Short-distance city commuters</li>
<li class="best-for-item">Delivery riders needing mileage</li>
<li class="best-for-item">First-time bike buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Hayate?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Highway riders</li>
<li class="not-recommended-item">Style-conscious buyers</li>
<li class="not-recommended-item">Those needing good resale value</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Hayate Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,15,000 - ৳1,25,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,700 - ৳2,100 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,300-1,600/month for 30 km daily at 57 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳500 - ৳800 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Hayate Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Excellent mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Cheap and efficient</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Very weak</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Poor quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very basic</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Outdated</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Hayate Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Suzuki Hayate?</h3>
<p class="faq-answer">Suzuki Hayate delivers excellent 55-60 km/l.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Suzuki Hayate?</h3>
<p class="faq-answer">Suzuki Hayate is priced between ৳1,15,000 to ৳1,25,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Suzuki Hayate good for daily use?</h3>
<p class="faq-answer">Yes, for basic city commuting it's adequate with excellent mileage.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Hayate or CD 70?</h3>
<p class="faq-answer">Honda CD 70 is better with superior build quality and resale value.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Hayate in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.9/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Suzuki Hayate is pure budget transportation for those who need cheapest possible running costs. At ৳1,20,000 with 55-60 km/l mileage, it minimizes ownership expenses. However, very basic features, poor build quality, and weak performance make it suitable only for short city rides. Honda CD 70 offers much better overall package for slightly more money.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For extreme budget buyers</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.9,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Suzuki Hayate review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Hayate (Rating: 3.9/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি হায়াটে রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি হায়াটে একটি বাজেট-বান্ধব ১১৩সিসি কমিউটার মোটরসাইকেল যার মূল্য ৳১,২০,০০০ টাকা। সাধারণ ডিজাইন, ভালো জ্বালানি দক্ষতা এবং সুজুকি নির্ভরযোগ্যতা সহ, এটি অতিরিক্ত ছাড়াই মৌলিক পরিবহনের প্রয়োজন খরচ-সচেতন দৈনিক যাত্রীদের জন্য একটি ব্যবহারিক পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি হায়াটে এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সাশ্রয়ী মূল্য:</strong> <span class="highlight-value">৳১,২০,০০০ টাকায় বাজেট-বান্ধব</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৫৫-৬০ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="highlight-value">রক্ষণাবেক্ষণ সহজ, সস্তা পার্টস</span></li>
<li class="highlight-item"><strong class="highlight-label">হালকা ওজন:</strong> <span class="highlight-value">ট্রাফিকে সামলানো সহজ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি হায়াটে এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">খুব সাশ্রয়ী:</strong> <span class="pro-description">৳১,২০,০০০ টাকায় এটি সবচেয়ে সস্তা বাইকগুলির একটি</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">৫৫-৬০ কিমি/লিটার জ্বালানি খরচ বাঁচায়</span></li>
<li class="pro-item"><strong class="pro-label">হালকা হ্যান্ডলিং:</strong> <span class="pro-description">শহরের ট্রাফিকে চলাফেরা সহজ</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ খরচ:</strong> <span class="pro-description">সাধারণ ডিজাইন, সস্তা স্পেয়ার পার্টস</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">সুজুকি জাপানি ইঞ্জিনিয়ারিং</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">সুবিধাজনক সেল্ফ-স্টার্ট বিকল্প</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি হায়াটে এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">খুব বেসিক ফিচার:</strong> <span class="con-description">কোন আধুনিক ফিচার নেই</span></li>
<li class="con-item"><strong class="con-label">শুধুমাত্র ড্রাম ব্রেক:</strong> <span class="con-description">কোন ডিস্ক ব্রেক বিকল্প নেই</span></li>
<li class="con-item"><strong class="con-label">খারাপ বিল্ড কোয়ালিটি:</strong> <span class="con-description">প্লাস্টিকের গুণমান গড়ের নিচে</span></li>
<li class="con-item"><strong class="con-label">পুরানো স্টাইলিং:</strong> <span class="con-description">পুরানো ফ্যাশন ডিজাইন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">সুজুকি সেন্টার সীমিত</span></li>
<li class="con-item"><strong class="con-label">দুর্বল পারফরম্যান্স:</strong> <span class="con-description">১১৩সিসি খুব শক্তিহীন</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">সুজুকি বাইক দ্রুত অবমূল্যায়ন করে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি হায়াটে কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extremely budget-conscious buyers</li>
<li class="best-for-item">Short-distance city commuters</li>
<li class="best-for-item">Delivery riders needing mileage</li>
<li class="best-for-item">First-time bike buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি হায়াটে কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Highway riders</li>
<li class="not-recommended-item">Style-conscious buyers</li>
<li class="not-recommended-item">Those needing good resale value</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি হায়াটে এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,১৫,০০০ - ৳১,২৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৭০০ - ৳২,১০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৩০০-১,৬০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৫০০ - ৳৮০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি হায়াটে পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- চমৎকার মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- সস্তা এবং দক্ষ</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- খুব দুর্বল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- বেসিক আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- খারাপ মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব বেসিক</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- পুরানো</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি হায়াটে সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">সুজুকি হায়াটের মাইলেজ কত?</h3>
<p class="faq-answer">সুজুকি হায়াটে চমৎকার ৫৫-৬০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি হায়াটের দাম কত?</h3>
<p class="faq-answer">সুজুকি হায়াটের দাম ৳১,১৫,০০০ থেকে ৳১,২৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য সুজুকি হায়াটে কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, মৌলিক শহর যাতায়াতের জন্য এটি চমৎকার মাইলেজ সহ পর্যাপ্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: হায়াটে নাকি সিডি ৭০?</h3>
<p class="faq-answer">হোন্ডা সিডি ৭০ উন্নত বিল্ড কোয়ালিটি এবং রিসেল ভ্যালু সহ ভালো।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি হায়াটে কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.9/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">সুজুকি হায়াটে সেই সকলের জন্য খাঁটি বাজেট পরিবহন যাদের সবচেয়ে সস্তা সম্ভাব্য চলমান খরচ প্রয়োজন। ৫৫-৬০ কিমি/লিটার মাইলেজ সহ ৳১,২০,০০০ টাকায়, এটি মালিকানা ব্যয় কমিয়ে দেয়। তবে, খুব বেসিক ফিচার, খারাপ বিল্ড কোয়ালিটি এবং দুর্বল পারফরম্যান্স এটিকে শুধুমাত্র ছোট শহর রাইডের জন্য উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চরম বাজেট ক্রেতাদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Hayate already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Hayate\n")

	return nil
}
