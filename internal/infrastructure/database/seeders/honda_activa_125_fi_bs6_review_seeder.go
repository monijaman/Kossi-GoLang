package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaActiva125FiBs6Review handles seeding Honda Activa 125 Fi BS6 product review and translation
type HondaActiva125FiBs6Review struct {
	BaseSeeder
}

// NewHondaActiva125FiBs6ReviewSeeder creates a new HondaActiva125FiBs6Review
func NewHondaActiva125FiBs6ReviewSeeder() *HondaActiva125FiBs6Review {
	return &HondaActiva125FiBs6Review{BaseSeeder: BaseSeeder{name: "Honda Activa 125 Fi BS6 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaActiva125FiBs6Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda Activa 125 Fi BS6").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda Activa 125 Fi BS6 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda Activa 125 Fi BS6 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda Activa 125 Fi BS6 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda Activa 125 Fi BS6 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda Activa 125 Fi BS6 is India's most popular automatic scooter featuring refined 125cc fuel-injected engine, BS6 emission compliance, spacious under-seat storage, and proven Honda reliability. Priced at ৳2,25,000, it offers unmatched automatic convenience with excellent fuel efficiency, comfortable ergonomics, practical storage solutions, modern FI technology, and legendary dependability making it the gold standard for family-friendly automatic transportation.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda Activa 125 Fi BS6 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection System:</strong> <span class="highlight-value">Modern FI for better efficiency and instant starts</span></li>
<li class="highlight-item"><strong class="highlight-label">BS6 Emission Compliance:</strong> <span class="highlight-value">Environment-friendly latest emission standards</span></li>
<li class="highlight-item"><strong class="highlight-label">Spacious Storage:</strong> <span class="highlight-value">Large under-seat storage for daily needs</span></li>
<li class="highlight-item"><strong class="highlight-label">Proven Reliability:</strong> <span class="highlight-value">India's most trusted scooter brand</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda Activa 125 Fi BS6 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Unmatched Reliability:</strong> <span class="pro-description">Proven Honda engineering with legendary dependability</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Fuel Efficiency:</strong> <span class="pro-description">Achieves 50-55 km/l with FI technology</span></li>
<li class="pro-item"><strong class="pro-label">Automatic Convenience:</strong> <span class="pro-description">CVT transmission provides effortless riding</span></li>
<li class="pro-item"><strong class="pro-label">Practical Storage:</strong> <span class="pro-description">Spacious under-seat compartment for helmet and bags</span></li>
<li class="pro-item"><strong class="pro-label">Family-Friendly:</strong> <span class="pro-description">Easy to ride, comfortable for all family members</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Honda service centers widely available</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda Activa 125 Fi BS6 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Conservative Styling:</strong> <span class="con-description">Design may feel dated compared to newer scooters</span></li>
<li class="con-item"><strong class="con-label">Limited Features:</strong> <span class="con-description">Lacks modern digital displays and connectivity</span></li>
<li class="con-item"><strong class="con-label">Average Performance:</strong> <span class="con-description">125cc provides adequate but not exciting power</span></li>
<li class="con-item"><strong class="con-label">Moderate Pricing:</strong> <span class="con-description">More expensive than some competitive scooters</span></li>
<li class="con-item"><strong class="con-label">Basic Suspension:</strong> <span class="con-description">Can feel harsh on very rough roads</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda Activa 125 Fi BS6 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Family transportation</li>
<li class="best-for-item">Daily urban commuting</li>
<li class="best-for-item">First-time scooter buyers</li>
<li class="best-for-item">Women riders seeking easy operation</li>
<li class="best-for-item">Reliability-focused buyers</li>
<li class="best-for-item">Elderly riders needing automatic convenience</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda Activa 125 Fi BS6?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Style-conscious younger riders</li>
<li class="not-recommended-item">Those seeking modern features</li>
<li class="not-recommended-item">Budget-constrained buyers</li>
<li class="not-recommended-item">Sporty riding enthusiasts</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda Activa 125 Fi BS6 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,25,000 - Premium for proven reliability</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳5-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳5-6 per km (50-55 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳4,000-6,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda Activa 125 Fi BS6 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Adequate scooter performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good daily comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Conservative design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good for reliability</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Outstanding Honda quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda Activa 125 Fi BS6 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Why is Activa so popular in India?</h3>
<p class="faq-answer">Activa's popularity stems from unmatched Honda reliability, ease of use, excellent fuel economy, wide service network, and proven track record. It's trusted by millions of families for over two decades.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for male riders?</h3>
<p class="faq-answer">Absolutely! While marketed as family-friendly, many male riders choose Activa for its reliability, convenience, and practicality. It's a sensible choice for anyone prioritizing dependability over sporty image.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What's the difference between BS4 and BS6 version?</h3>
<p class="faq-answer">The BS6 version features cleaner emission technology with fuel injection, slightly refined engine, and updated exhaust system to meet stricter emission norms while maintaining similar performance and fuel economy.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda Activa 125 Fi BS6 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,25,000, the Honda Activa 125 Fi BS6 represents the gold standard for automatic scooters with unmatched reliability, excellent fuel efficiency (50-55 km/l), automatic convenience, practical storage, and legendary Honda dependability. It excels as family transportation and daily commuter with wide service network support. However, conservative styling, limited modern features, average performance, moderate pricing compared to competitors, and basic suspension make it best suited for families, reliability-focused buyers, and those prioritizing proven dependability and ease of use over sporty styling and modern features.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For unmatched reliability and family-friendly automatic convenience</span></p>
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
		return fmt.Errorf("error creating Honda Activa 125 Fi BS6 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda Activa 125 Fi BS6 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ হল ভারতের সবচেয়ে জনপ্রিয় অটোমেটিক স্কুটার যেখানে পরিমার্জিত ১২৫সিসি ফুয়েল-ইনজেক্টেড ইঞ্জিন, বিএস৬ এমিশন কমপ্লায়েন্স, প্রশস্ত আন্ডার-সিট স্টোরেজ এবং প্রমাণিত হোন্ডা নির্ভরযোগ্যতা রয়েছে। ৳২,২৫,০০০ টাকায় মূল্যায়িত, এটি চমৎকার জ্বালানি দক্ষতা, আরামদায়ক এরগোনমিক্স, ব্যবহারিক স্টোরেজ সমাধান, আধুনিক এফআই প্রযুক্তি এবং কিংবদন্তি নির্ভরযোগ্যতা সহ অতুলনীয় অটোমেটিক সুবিধা প্রদান করে যা এটিকে পরিবার-বন্ধুত্বপূর্ণ অটোমেটিক পরিবহনের জন্য স্বর্ণ মান করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন সিস্টেম:</strong> <span class="highlight-value">ভাল দক্ষতা এবং তাৎক্ষণিক স্টার্টের জন্য আধুনিক এফআই</span></li>
<li class="highlight-item"><strong class="highlight-label">বিএস৬ এমিশন কমপ্লায়েন্স:</strong> <span class="highlight-value">পরিবেশ-বান্ধব সর্বশেষ এমিশন মান</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রশস্ত স্টোরেজ:</strong> <span class="highlight-value">দৈনিক চাহিদার জন্য বড় আন্ডার-সিট স্টোরেজ</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রমাণিত নির্ভরযোগ্যতা:</strong> <span class="highlight-value">ভারতের সবচেয়ে বিশ্বস্ত স্কুটার ব্র্যান্ড</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অতুলনীয় নির্ভরযোগ্যতা:</strong> <span class="pro-description">কিংবদন্তি নির্ভরযোগ্যতা সহ প্রমাণিত হোন্ডা ইঞ্জিনিয়ারিং</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার জ্বালানি দক্ষতা:</strong> <span class="pro-description">এফআই প্রযুক্তি দিয়ে ৫০-৫৫ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">অটোমেটিক সুবিধা:</strong> <span class="pro-description">সিভিটি ট্রান্সমিশন অনায়াস রাইডিং প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক স্টোরেজ:</strong> <span class="pro-description">হেলমেট এবং ব্যাগের জন্য প্রশস্ত আন্ডার-সিট কম্পার্টমেন্ট</span></li>
<li class="pro-item"><strong class="pro-label">পরিবার-বান্ধুত্বপূর্ণ:</strong> <span class="pro-description">চালনা সহজ, সব পরিবার সদস্যদের জন্য আরামদায়</span></li>
<li class="pro-item"><strong class="pro-label">বিস্তৃত সেবা নেটওয়ার্ক:</strong> <span class="pro-description">হোন্ডা সেবা কেন্দ্র ব্যাপকভাবে উপলব্ধ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">রক্ষণশীল স্টাইলিং:</strong> <span class="con-description">নতুন স্কুটারের তুলনায় ডিজাইন পুরাতন মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">সীমিত বৈশিষ্ট্য:</strong> <span class="con-description">আধুনিক ডিজিটাল ডিসপ্লে এবং সংযোগের অভাব</span></li>
<li class="con-item"><strong class="con-label">গড় পারফরমেন্স:</strong> <span class="con-description">১২৫সিসি পর্যাপ্ত কিন্তু উত্তেজনাপূর্ণ নয় এমন শক্তি প্রদান করে</span></li>
<li class="con-item"><strong class="con-label">মাঝারি মূল্য:</strong> <span class="con-description">কিছু প্রতিযোগী স্কুটারের চেয়ে বেশি ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">মৌলিক সাসপেনশন:</strong> <span class="con-description">অত্যন্ত রুক্ষ রাস্তায় কঠোর মনে হতে পারে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Family transportation</li>
<li class="best-for-item">Daily urban commuting</li>
<li class="best-for-item">First-time scooter buyers</li>
<li class="best-for-item">Women riders seeking easy operation</li>
<li class="best-for-item">Reliability-focused buyers</li>
<li class="best-for-item">Elderly riders needing automatic convenience</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Style-conscious younger riders</li>
<li class="not-recommended-item">Those seeking modern features</li>
<li class="not-recommended-item">Budget-constrained buyers</li>
<li class="not-recommended-item">Sporty riding enthusiasts</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,২৫,০০০ - প্রমাণিত নির্ভরযোগ্যতার জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম - ৳৫-৭ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৫-৬ প্রতি কিমি (৫০-৫৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৪,০০০-৬,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- পর্যাপ্ত স্কুটার পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- ভাল দৈনিক আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- রক্ষণশীল ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- নির্ভরযোগ্যতার জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- অসাধারণ হোন্ডা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ভারতে অ্যাক্টিভা এত জনপ্রিয় কেন?</h3>
<p class="faq-answer">অ্যাক্টিভার জনপ্রিয়তা অতুলনীয় হোন্ডা নির্ভরযোগ্যতা, ব্যবহারের সহজতা, চমৎকার জ্বালানি অর্থনীতি, বিস্তৃত সেবা নেটওয়ার্ক এবং প্রমাণিত ট্র্যাক রেকর্ড থেকে উদ্ভূত। এটি দুই দশকেরও বেশি সময় ধরে লক্ষ লক্ষ পরিবার দ্বারা বিশ্বস্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি পুরুষ রাইডারদের জন্য উপযুক্ত?</h3>
<p class="faq-answer">একেবারেই! পরিবার-বান্ধুত্বপূর্ণ হিসাবে বিপণন করা হলেও, অনেক পুরুষ রাইডার এর নির্ভরযোগ্যতা, সুবিধা এবং ব্যবহারিকতার জন্য অ্যাক্টিভা বেছে নেয়। এটি স্পোর্টি ইমেজের চেয়ে নির্ভরযোগ্যতাকে অগ্রাধিকার দেওয়া যে কারও জন্য একটি বুদ্ধিমান পছন্দ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বিএস৪ এবং বিএস৬ সংস্করণের মধ্যে পার্থক্য কী?</h3>
<p class="faq-answer">বিএস৬ সংস্করণে ফুয়েল ইনজেকশন সহ পরিষ্কার এমিশন প্রযুক্তি, সামান্য পরিমার্জিত ইঞ্জিন এবং অনুরূপ পারফরমেন্স এবং জ্বালানি অর্থনীতি বজায় রেখে কঠোর এমিশন নিয়ম মেনে চলতে আপডেটেড এক্সহস্ট সিস্টেম রয়েছে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,২৫,০০০ টাকায় হোন্ডা অ্যাক্টিভা ১২৫ এফআই বিএস৬ অতুলনীয় নির্ভরযোগ্যতা, চমৎকার জ্বালানি দক্ষতা (৫০-৫৫ কিমি/লিটার), অটোমেটিক সুবিধা, ব্যবহারিক স্টোরেজ এবং কিংবদন্তি হোন্ডা নির্ভরযোগ্যতা সহ অটোমেটিক স্কুটারের জন্য স্বর্ণ মান উপস্থাপন করে। এটি বিস্তৃত সেবা নেটওয়ার্ক সহায়তা সহ পরিবার পরিবহন এবং দৈনিক যাত্রী হিসাবে শ্রেষ্ঠ। তবে, রক্ষণশীল স্টাইলিং, সীমিত আধুনিক বৈশিষ্ট্য, গড় পারফরমেন্স, প্রতিযোগীদের তুলনায় মাঝারি মূল্য এবং মৌলিক সাসপেনশন এটিকে পরিবার, নির্ভরযোগ্যতা-ফোকাসড ক্রেতা এবং স্পোর্টি স্টাইলিং এবং আধুনিক বৈশিষ্ট্যের চেয়ে প্রমাণিত নির্ভরযোগ্যতা এবং ব্যবহারের সহজতাকে অগ্রাধিকার দেওয়াদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অতুলনীয় নির্ভরযোগ্যতা এবং পরিবার-বান্ধুত্বপূর্ণ অটোমেটিক সুবিধার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda Activa 125 Fi BS6 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda Activa 125 Fi BS6\n")

	return nil
}
