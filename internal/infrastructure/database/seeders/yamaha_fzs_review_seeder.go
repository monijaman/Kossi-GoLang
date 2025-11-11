package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaFZSReviewSeeder handles seeding Yamaha FZS product review and translation
type YamahaFZSReviewSeeder struct {
	BaseSeeder
}

// NewYamahaFZSReviewSeeder creates a new YamahaFZSReviewSeeder
func NewYamahaFZSReviewSeeder() *YamahaFZSReviewSeeder {
	return &YamahaFZSReviewSeeder{BaseSeeder: BaseSeeder{name: "Yamaha FZS Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaFZSReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha FZS").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha FZS product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha FZS product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha FZS already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha FZS Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha FZS is a premium 150cc motorcycle priced at ৳2,85,000, offering muscular styling, fuel injection, and Yamaha's legendary reliability. With LED lighting and digital features, it's a stylish choice for riders wanting sporty looks with everyday practicality.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha FZS Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection:</strong> <span class="highlight-value">FI technology for better efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Muscular Design:</strong> <span class="highlight-value">Aggressive sporty styling</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Full LED lights front and rear</span></li>
<li class="highlight-item"><strong class="highlight-label">Great Build Quality:</strong> <span class="highlight-value">Premium Yamaha quality</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha FZS Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Fuel Injection System:</strong> <span class="pro-description">Better efficiency and instant throttle response</span></li>
<li class="pro-item"><strong class="pro-label">Muscular Styling:</strong> <span class="pro-description">Aggressive street fighter design with tank shrouds</span></li>
<li class="pro-item"><strong class="pro-label">Full LED Lighting:</strong> <span class="pro-description">LED headlamp, taillight, and indicators</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Build Quality:</strong> <span class="pro-description">Premium fit and finish with solid feel</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Upright seating position for long rides</span></li>
<li class="pro-item"><strong class="pro-label">Good Suspension:</strong> <span class="pro-description">Telescopic front and monoshock rear</span></li>
<li class="pro-item"><strong class="pro-label">Disc Brakes:</strong> <span class="pro-description">Front and rear disc brakes with ABS option</span></li>
<li class="pro-item"><strong class="pro-label">Yamaha Reliability:</strong> <span class="pro-description">Proven Japanese engineering quality</span></li>
<li class="pro-item"><strong class="pro-label">Good Ground Clearance:</strong> <span class="pro-description">160mm clearance handles bad roads</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Fully digital LCD meter with gear indicator</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha FZS Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price:</strong> <span class="con-description">৳2,85,000 is expensive for 150cc</span></li>
<li class="con-item"><strong class="con-label">Single Channel ABS:</strong> <span class="con-description">ABS only on front wheel</span></li>
<li class="con-item"><strong class="con-label">Average Mileage:</strong> <span class="con-description">40-45 km/l not great for 150cc FI</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha centers than Honda</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">137 kg kerb weight feels heavy in traffic</span></li>
<li class="con-item"><strong class="con-label">Lower Resale Value:</strong> <span class="con-description">Yamaha bikes don't hold value like Honda</span></li>
<li class="con-item"><strong class="con-label">Firm Seat:</strong> <span class="con-description">Seat padding could be softer</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha FZS in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Style-conscious urban riders</li>
<li class="best-for-item">Those wanting premium build quality</li>
<li class="best-for-item">Riders who prioritize looks</li>
<li class="best-for-item">Daily commuters with highway use</li>
<li class="best-for-item">Yamaha brand loyalists</li>
<li class="best-for-item">Riders wanting modern features</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha FZS?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those needing maximum mileage</li>
<li class="not-recommended-item">Riders in areas with limited Yamaha service</li>
<li class="not-recommended-item">Those planning to resell soon</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha FZS Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,80,000 - ৳2,90,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳3,200 - ৳3,800 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,300-2,700/month for 30 km daily at 42 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 4 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,200 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha FZS Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent Yamaha quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Premium but expensive</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good FI power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good ergonomics</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Excellent quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- FI + LED + ABS</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Stunning looks</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha FZS Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Yamaha FZS?</h3>
<p class="faq-answer">Yamaha FZS delivers 40-45 km/l with fuel injection technology.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Yamaha FZS in Bangladesh?</h3>
<p class="faq-answer">Yamaha FZS is priced between ৳2,80,000 to ৳2,90,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Yamaha FZS have ABS?</h3>
<p class="faq-answer">Yes, Yamaha FZS has single channel ABS on the front wheel.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Yamaha FZS good for daily use?</h3>
<p class="faq-answer">Yes, it's excellent for daily use with comfortable ergonomics and reliable performance.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Yamaha FZS or Honda CB Shine?</h3>
<p class="faq-answer">FZS offers better styling and features, while CB Shine has better mileage and resale value.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Yamaha FZS?</h3>
<p class="faq-answer">Yamaha FZS can reach approximately 115-120 km/h.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Yamaha FZS have LED lights?</h3>
<p class="faq-answer">Yes, Yamaha FZS has full LED headlight, taillight, and indicators.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What engine does Yamaha FZS use?</h3>
<p class="faq-answer">Yamaha FZS uses a 149cc air-cooled fuel-injected single-cylinder engine.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha FZS in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Yamaha FZS is a premium 150cc motorcycle that delivers on styling, build quality, and features. At ৳2,85,000, it's expensive but offers muscular design, fuel injection, LED lighting, and excellent Yamaha reliability. Best for riders who prioritize looks and don't mind paying premium. However, average mileage and lower resale value make it less practical than Honda alternatives for value-focused buyers.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For style-focused premium buyers</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.4,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha FZS review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha FZS (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা এফজেডএস রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা এফজেডএস একটি প্রিমিয়াম ১৫০সিসি মোটরসাইকেল যার মূল্য ৳২,৮৫,০০০ টাকা, যা পেশীবহুল স্টাইলিং, ফুয়েল ইনজেকশন এবং ইয়ামাহার কিংবদন্তি নির্ভরযোগ্যতা প্রদান করে। এলইডি লাইটিং এবং ডিজিটাল ফিচার সহ, এটি দৈনন্দিন ব্যবহারিকতার সাথে স্পোর্টি লুক চান রাইডারদের জন্য একটি স্টাইলিশ পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা এফজেডএস এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন:</strong> <span class="highlight-value">ভালো দক্ষতার জন্য এফআই প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">পেশীবহুল ডিজাইন:</strong> <span class="highlight-value">আগ্রাসী স্পোর্টি স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">সামনে এবং পিছনে সম্পূর্ণ এলইডি লাইট</span></li>
<li class="highlight-item"><strong class="highlight-label">দুর্দান্ত বিল্ড কোয়ালিটি:</strong> <span class="highlight-value">প্রিমিয়াম ইয়ামাহা মান</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা এফজেডএস এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন সিস্টেম:</strong> <span class="pro-description">ভালো দক্ষতা এবং তাৎক্ষণিক থ্রটল প্রতিক্রিয়া</span></li>
<li class="pro-item"><strong class="pro-label">পেশীবহুল স্টাইলিং:</strong> <span class="pro-description">ট্যাঙ্ক শ্রাউড সহ আগ্রাসী স্ট্রিট ফাইটার ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">সম্পূর্ণ এলইডি লাইটিং:</strong> <span class="pro-description">এলইডি হেডল্যাম্প, টেললাইট এবং ইন্ডিকেটর</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড কোয়ালিটি:</strong> <span class="pro-description">শক্ত অনুভূতি সহ প্রিমিয়াম ফিট এবং ফিনিশ</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এর্গোনমিক্স:</strong> <span class="pro-description">লম্বা রাইডের জন্য সোজা বসার অবস্থান</span></li>
<li class="pro-item"><strong class="pro-label">ভালো সাসপেনশন:</strong> <span class="pro-description">টেলিস্কোপিক ফ্রন্ট এবং মনোশক রিয়ার</span></li>
<li class="pro-item"><strong class="pro-label">ডিস্ক ব্রেক:</strong> <span class="pro-description">এবিএস বিকল্প সহ সামনে এবং পিছনে ডিস্ক ব্রেক</span></li>
<li class="pro-item"><strong class="pro-label">ইয়ামাহা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত জাপানি ইঞ্জিনিয়ারিং মান</span></li>
<li class="pro-item"><strong class="pro-label">ভালো গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="pro-description">১৬০মিমি ক্লিয়ারেন্স খারাপ রাস্তা সামলায়</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">গিয়ার ইন্ডিকেটর সহ সম্পূর্ণ ডিজিটাল এলসিডি মিটার</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা এফজেডএস এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">১৫০সিসির জন্য ৳২,৮৫,০০০ টাকা ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="con-description">শুধুমাত্র সামনের চাকায় এবিএস</span></li>
<li class="con-item"><strong class="con-label">গড় মাইলেজ:</strong> <span class="con-description">১৫০সিসি এফআই-এর জন্য ৪০-৪৫ কিমি/লিটার ভালো নয়</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">হোন্ডার চেয়ে কম ইয়ামাহা সেন্টার</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">১৩৭ কেজি কার্ব ওজন ট্রাফিকে ভারী মনে হয়</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">ইয়ামাহা বাইক হোন্ডার মতো মূল্য ধরে রাখে না</span></li>
<li class="con-item"><strong class="con-label">শক্ত সিট:</strong> <span class="con-description">সিট প্যাডিং নরম হতে পারত</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা এফজেডএস কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Style-conscious urban riders</li>
<li class="best-for-item">Those wanting premium build quality</li>
<li class="best-for-item">Riders who prioritize looks</li>
<li class="best-for-item">Daily commuters with highway use</li>
<li class="best-for-item">Yamaha brand loyalists</li>
<li class="best-for-item">Riders wanting modern features</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা এফজেডএস কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those needing maximum mileage</li>
<li class="not-recommended-item">Riders in areas with limited Yamaha service</li>
<li class="not-recommended-item">Those planning to resell soon</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা এফজেডএস এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৮০,০০০ - ৳২,৯০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৩,২০০ - ৳৩,৮০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,৩০০-২,৭০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৪ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,২০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা এফজেডএস পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার ইয়ামাহা মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- প্রিমিয়াম কিন্তু ব্যয়বহুল</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো এফআই পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো এর্গোনমিক্স</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- চমৎকার মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- এফআই + এলইডি + এবিএস</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অত্যাশ্চর্য চেহারা</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা এফজেডএস সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা এফজেডএস এর মাইলেজ কত?</h3>
<p class="faq-answer">ইয়ামাহা এফজেডএস ফুয়েল ইনজেকশন প্রযুক্তি সহ ৪০-৪৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে ইয়ামাহা এফজেডএস এর দাম কত?</h3>
<p class="faq-answer">ইয়ামাহা এফজেডএস এর দাম ৳২,৮০,০০০ থেকে ৳২,৯০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা এফজেডএসে কি এবিএস আছে?</h3>
<p class="faq-answer">হ্যাঁ, ইয়ামাহা এফজেডএসে সামনের চাকায় সিঙ্গেল চ্যানেল এবিএস আছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য ইয়ামাহা এফজেডএস কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, এটি আরামদায়ক এর্গোনমিক্স এবং নির্ভরযোগ্য পারফরম্যান্স সহ দৈনন্দিন ব্যবহারের জন্য চমৎকার।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: ইয়ামাহা এফজেডএস নাকি হোন্ডা সিবি শাইন?</h3>
<p class="faq-answer">এফজেডএস ভালো স্টাইলিং এবং ফিচার দেয়, যেখানে সিবি শাইন ভালো মাইলেজ এবং রিসেল ভ্যালু দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা এফজেডএস এর সর্বোচ্চ গতি কত?</h3>
<p class="faq-answer">ইয়ামাহা এফজেডএস প্রায় ১১৫-১২০ কিমি/ঘন্টা পৌঁছাতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা এফজেডএসে কি এলইডি লাইট আছে?</h3>
<p class="faq-answer">হ্যাঁ, ইয়ামাহা এফজেডএসে সম্পূর্ণ এলইডি হেডলাইট, টেললাইট এবং ইন্ডিকেটর আছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা এফজেডএস কোন ইঞ্জিন ব্যবহার করে?</h3>
<p class="faq-answer">ইয়ামাহা এফজেডএস একটি ১৪৯সিসি এয়ার-কুলড ফুয়েল-ইনজেক্টেড সিঙ্গেল-সিলিন্ডার ইঞ্জিন ব্যবহার করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা এফজেডএস কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">ইয়ামাহা এফজেডএস একটি প্রিমিয়াম ১৫০সিসি মোটরসাইকেল যা স্টাইলিং, বিল্ড কোয়ালিটি এবং ফিচার প্রদান করে। ৳২,৮৫,০০০ টাকায়, এটি ব্যয়বহুল কিন্তু পেশীবহুল ডিজাইন, ফুয়েল ইনজেকশন, এলইডি লাইটিং এবং চমৎকার ইয়ামাহা নির্ভরযোগ্যতা প্রদান করে। রাইডারদের জন্য সেরা যারা চেহারাকে অগ্রাধিকার দেন এবং প্রিমিয়াম দিতে আপত্তি করেন না।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্টাইল-কেন্দ্রিক প্রিমিয়াম ক্রেতাদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha FZS already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha FZS\n")

	return nil
}
