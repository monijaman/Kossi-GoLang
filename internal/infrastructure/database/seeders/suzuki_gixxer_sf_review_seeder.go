package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxerSFReviewSeeder handles seeding Suzuki Gixxer SF product review and translation
type SuzukiGixxerSFReviewSeeder struct {
	BaseSeeder
}

// NewSuzukiGixxerSFReviewSeeder creates a new SuzukiGixxerSFReviewSeeder
func NewSuzukiGixxerSFReviewSeeder() *SuzukiGixxerSFReviewSeeder {
	return &SuzukiGixxerSFReviewSeeder{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer SF Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxerSFReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer SF").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Gixxer SF product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Gixxer SF product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer SF already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer SF Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Gixxer SF is the faired version of the Gixxer, priced at ৳2,95,000. It offers aerodynamic full fairing, wind protection, and sporty looks while retaining the excellent handling of the naked Gixxer. Perfect for riders wanting a mini sports bike experience on Bangladesh roads.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer SF Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Full Fairing:</strong> <span class="highlight-value">Aerodynamic design with wind protection</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Same excellent dynamics as Gixxer</span></li>
<li class="highlight-item"><strong class="highlight-label">Sports Bike Look:</strong> <span class="highlight-value">Aggressive faired design</span></li>
<li class="highlight-item"><strong class="highlight-label">Powerful Engine:</strong> <span class="highlight-value">155cc with 14.1 PS power</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer SF Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Full Fairing Design:</strong> <span class="pro-description">Aerodynamic fairing provides wind protection</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Handling:</strong> <span class="pro-description">Inherits Gixxer's legendary handling</span></li>
<li class="pro-item"><strong class="pro-label">Powerful Performance:</strong> <span class="pro-description">155cc FI engine delivers strong acceleration</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Looks:</strong> <span class="pro-description">Aggressive faired sports bike styling</span></li>
<li class="pro-item"><strong class="pro-label">Dual Disc Brakes:</strong> <span class="pro-description">Front and rear disc with ABS</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Solid Japanese quality standards</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Fully digital LCD meter</span></li>
<li class="pro-item"><strong class="pro-label">Highway Stability:</strong> <span class="pro-description">Fairing provides better high-speed stability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer SF Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Expensive Price:</strong> <span class="con-description">৳2,95,000 is pricey for 155cc</span></li>
<li class="con-item"><strong class="con-label">Average Mileage:</strong> <span class="con-description">36-40 km/l lower than naked version</span></li>
<li class="con-item"><strong class="con-label">Heavier Weight:</strong> <span class="con-description">Fairing adds weight, affects city maneuverability</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Suzuki service limited compared to Honda/Yamaha</span></li>
<li class="con-item"><strong class="con-label">Expensive Fairing Repair:</strong> <span class="con-description">Fairing parts costly to replace after accidents</span></li>
<li class="con-item"><strong class="con-label">Lower Resale Value:</strong> <span class="con-description">Suzuki bikes depreciate faster</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Gixxer SF in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Sports bike enthusiasts on budget</li>
<li class="best-for-item">Highway and long-distance riders</li>
<li class="best-for-item">Those wanting wind protection</li>
<li class="best-for-item">Performance-focused riders</li>
<li class="best-for-item">Style-conscious young riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Gixxer SF?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Mileage-focused buyers</li>
<li class="not-recommended-item">City-only commuters</li>
<li class="not-recommended-item">Those worried about fairing damage</li>
<li class="not-recommended-item">Budget-constrained riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer SF Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,90,000 - ৳3,00,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳3,700 - ৳4,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,600-3,000/month for 30 km daily at 38 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 4 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,200 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer SF Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Below average</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good Suzuki quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Expensive but sporty</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Strong performance</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Sporty position</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Solid build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Fairing + FI + ABS</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Stunning sports bike</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Gixxer SF Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Suzuki Gixxer SF?</h3>
<p class="faq-answer">Suzuki Gixxer SF delivers 36-40 km/l due to the aerodynamic fairing.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Suzuki Gixxer SF in Bangladesh?</h3>
<p class="faq-answer">Suzuki Gixxer SF is priced between ৳2,90,000 to ৳3,00,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Gixxer or Gixxer SF?</h3>
<p class="faq-answer">Gixxer SF has better highway stability and looks, while naked Gixxer is better for city use.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Suzuki Gixxer SF have ABS?</h3>
<p class="faq-answer">Yes, Suzuki Gixxer SF comes with single channel ABS.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer SF in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Suzuki Gixxer SF is a mini sports bike that delivers full-faired aesthetics and excellent handling at ৳2,95,000. Best for riders wanting sports bike looks without breaking the bank. The fairing provides wind protection and highway stability, making it suitable for long rides. However, lower mileage, expensive fairing repairs, and limited service network are concerns. Great for sports enthusiasts, but practical commuters should consider naked Gixxer or Honda/Yamaha alternatives.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For mini sports bike enthusiasts</span></p>
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
		return fmt.Errorf("error creating Suzuki Gixxer SF review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer SF (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার এসএফ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি গিক্সার এসএফ গিক্সারের ফেয়ারড সংস্করণ, যার মূল্য ৳২,৯৫,০০০ টাকা। এটি এরোডাইনামিক ফুল ফেয়ারিং, উইন্ড প্রোটেকশন এবং স্পোর্টি লুক প্রদান করে যখন নেকেড গিক্সারের চমৎকার হ্যান্ডলিং বজায় রাখে। বাংলাদেশের রাস্তায় মিনি স্পোর্টস বাইক অভিজ্ঞতা চান রাইডারদের জন্য নিখুঁত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার এসএফ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুল ফেয়ারিং:</strong> <span class="highlight-value">উইন্ড প্রোটেকশন সহ এরোডাইনামিক ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="highlight-value">গিক্সারের মতো একই চমৎকার ডাইনামিক্স</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্টস বাইক লুক:</strong> <span class="highlight-value">আগ্রাসী ফেয়ারড ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">শক্তিশালী ইঞ্জিন:</strong> <span class="highlight-value">১৪.১ পিএস পাওয়ার সহ ১৫৫সিসি</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার এসএফ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ফুল ফেয়ারিং ডিজাইন:</strong> <span class="pro-description">এরোডাইনামিক ফেয়ারিং উইন্ড প্রোটেকশন প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার হ্যান্ডলিং:</strong> <span class="pro-description">গিক্সারের কিংবদন্তি হ্যান্ডলিং উত্তরাধিকার পায়</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী পারফরম্যান্স:</strong> <span class="pro-description">১৫৫সিসি এফআই ইঞ্জিন শক্তিশালী এক্সিলারেশন দেয়</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি লুক:</strong> <span class="pro-description">আগ্রাসী ফেয়ারড স্পোর্টস বাইক স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল ডিস্ক ব্রেক:</strong> <span class="pro-description">এবিএস সহ সামনে এবং পিছনে ডিস্ক</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">শক্ত জাপানি মান নির্ধারণ</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">সম্পূর্ণ ডিজিটাল এলসিডি মিটার</span></li>
<li class="pro-item"><strong class="pro-label">হাইওয়ে স্থিতিশীলতা:</strong> <span class="pro-description">ফেয়ারিং ভালো উচ্চ-গতি স্থিতিশীলতা প্রদান করে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার এসএফ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">ব্যয়বহুল মূল্য:</strong> <span class="con-description">১৫৫সিসির জন্য ৳২,৯৫,০০০ টাকা দামি</span></li>
<li class="con-item"><strong class="con-label">গড় মাইলেজ:</strong> <span class="con-description">নেকেড সংস্করণের চেয়ে ৩৬-৪০ কিমি/লিটার কম</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">ফেয়ারিং ওজন যোগ করে, শহরের চলাফেরা প্রভাবিত করে</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">হোন্ডা/ইয়ামাহার তুলনায় সুজুকি সার্ভিস সীমিত</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল ফেয়ারিং মেরামত:</strong> <span class="con-description">দুর্ঘটনার পর ফেয়ারিং পার্টস প্রতিস্থাপন ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">সুজুকি বাইক দ্রুত অবমূল্যায়ন করে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি গিক্সার এসএফ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Sports bike enthusiasts on budget</li>
<li class="best-for-item">Highway and long-distance riders</li>
<li class="best-for-item">Those wanting wind protection</li>
<li class="best-for-item">Performance-focused riders</li>
<li class="best-for-item">Style-conscious young riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি গিক্সার এসএফ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Mileage-focused buyers</li>
<li class="not-recommended-item">City-only commuters</li>
<li class="not-recommended-item">Those worried about fairing damage</li>
<li class="not-recommended-item">Budget-constrained riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি গিক্সার এসএফ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৯০,০০০ - ৳৩,০০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৩,৭০০ - ৳৪,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৩৮ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,৬০০-৩,০০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৪ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,২০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি গিক্সার এসএফ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড়ের নিচে</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো সুজুকি মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ব্যয়বহুল কিন্তু স্পোর্টি</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- শক্তিশালী পারফরম্যান্স</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- স্পোর্টি অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- শক্ত বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ফেয়ারিং + এফআই + এবিএস</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অত্যাশ্চর্য স্পোর্টস বাইক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি গিক্সার এসএফ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">সুজুকি গিক্সার এসএফ এর মাইলেজ কত?</h3>
<p class="faq-answer">এরোডাইনামিক ফেয়ারিংয়ের কারণে সুজুকি গিক্সার এসএফ ৩৬-৪০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে সুজুকি গিক্সার এসএফ এর দাম কত?</h3>
<p class="faq-answer">সুজুকি গিক্সার এসএফ এর দাম ৳২,৯০,০০০ থেকে ৳৩,০০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: গিক্সার নাকি গিক্সার এসএফ?</h3>
<p class="faq-answer">গিক্সার এসএফে ভালো হাইওয়ে স্থিতিশীলতা এবং লুক আছে, যেখানে নেকেড গিক্সার শহরের ব্যবহারের জন্য ভালো।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি গিক্সার এসএফে কি এবিএস আছে?</h3>
<p class="faq-answer">হ্যাঁ, সুজুকি গিক্সার এসএফ সিঙ্গেল চ্যানেল এবিএস সহ আসে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি গিক্সার এসএফ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">সুজুকি গিক্সার এসএফ একটি মিনি স্পোর্টস বাইক যা ৳২,৯৫,০০০ টাকায় ফুল-ফেয়ারড নান্দনিকতা এবং চমৎকার হ্যান্ডলিং প্রদান করে। রাইডারদের জন্য সেরা যারা ব্যাংক না ভেঙে স্পোর্টস বাইক লুক চান। ফেয়ারিং উইন্ড প্রোটেকশন এবং হাইওয়ে স্থিতিশীলতা প্রদান করে, যা এটিকে লম্বা রাইডের জন্য উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- মিনি স্পোর্টস বাইক উৎসাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer SF already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer SF\n")

	return nil
}
