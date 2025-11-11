package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSApacheRTR180ReviewSeeder handles seeding TVS Apache RTR 180 product review and translation
type TVSApacheRTR180ReviewSeeder struct {
	BaseSeeder
}

// NewTVSApacheRTR180ReviewSeeder creates a new TVSApacheRTR180ReviewSeeder
func NewTVSApacheRTR180ReviewSeeder() *TVSApacheRTR180ReviewSeeder {
	return &TVSApacheRTR180ReviewSeeder{BaseSeeder: BaseSeeder{name: "TVS Apache RTR 180 Review"}}
}

// Seed implements the Seeder interface
func (s *TVSApacheRTR180ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 180").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("TVS Apache RTR 180 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding TVS Apache RTR 180 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Apache RTR 180 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Apache RTR 180 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The TVS Apache RTR 180 is a performance-focused 180cc motorcycle priced at ৳2,35,000. With powerful engine producing 16.8 PS, race-tuned suspension, and LED lighting, it delivers thrilling performance. However, very limited service network and poor resale value make it a risky choice in Bangladesh market.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 180 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Powerful Engine:</strong> <span class="highlight-value">180cc with 16.8 PS power</span></li>
<li class="highlight-item"><strong class="highlight-label">Race Technology:</strong> <span class="highlight-value">Race-derived suspension and setup</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Full LED headlamp and DRL</span></li>
<li class="highlight-item"><strong class="highlight-label">Top Performance:</strong> <span class="highlight-value">Best power in segment</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 180 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Power:</strong> <span class="pro-description">180cc produces 16.8 PS, strongest in segment</span></li>
<li class="pro-item"><strong class="pro-label">Race-Tuned Setup:</strong> <span class="pro-description">Sharp handling with RTR race technology</span></li>
<li class="pro-item"><strong class="pro-label">LED Lighting:</strong> <span class="pro-description">Full LED headlamp with DRL</span></li>
<li class="pro-item"><strong class="pro-label">Good Value:</strong> <span class="pro-description">Feature-packed at ৳2,35,000</span></li>
<li class="pro-item"><strong class="pro-label">Dual Disc ABS:</strong> <span class="pro-description">Front and rear disc with ABS</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Advanced LCD with lap timer</span></li>
<li class="pro-item"><strong class="pro-label">Monoshock Suspension:</strong> <span class="pro-description">Race-spec rear monoshock</span></li>
<li class="pro-item"><strong class="pro-label">Top Speed:</strong> <span class="pro-description">Can reach 130+ km/h</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 180 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Limited Service:</strong> <span class="con-description">TVS service centers extremely rare</span></li>
<li class="con-item"><strong class="con-label">Worst Resale Value:</strong> <span class="con-description">TVS bikes depreciate fastest</span></li>
<li class="con-item"><strong class="con-label">Average Mileage:</strong> <span class="con-description">35-40 km/l not good for 180cc</span></li>
<li class="con-item"><strong class="con-label">Build Quality Issues:</strong> <span class="con-description">Plastic and paint quality concerns</span></li>
<li class="con-item"><strong class="con-label">Vibrations:</strong> <span class="con-description">Noticeable vibrations above 80 km/h</span></li>
<li class="con-item"><strong class="con-label">Low Brand Value:</strong> <span class="con-description">TVS not popular in Bangladesh</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy TVS Apache RTR 180 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance enthusiasts near TVS centers</li>
<li class="best-for-item">Track day and racing enthusiasts</li>
<li class="best-for-item">Those wanting maximum power on budget</li>
<li class="best-for-item">Riders not worried about resale</li>
<li class="best-for-item">Tech-savvy tuning enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy TVS Apache RTR 180?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Most buyers due to service issues</li>
<li class="not-recommended-item">Those needing good resale value</li>
<li class="not-recommended-item">Riders in remote areas</li>
<li class="not-recommended-item">Long-term ownership plans</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 180 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,30,000 - ৳2,40,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳3,500 - ৳4,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,700-3,100/month for 30 km daily at 37 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,100 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 180 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Below average</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Service issues hurt value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Outstanding power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Firm sporty setup</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Great features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Aggressive race look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">TVS Apache RTR 180 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of TVS Apache RTR 180?</h3>
<p class="faq-answer">TVS Apache RTR 180 delivers 35-40 km/l with race-tuned engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Apache RTR 180 in Bangladesh?</h3>
<p class="faq-answer">Apache RTR 180 is priced between ৳2,30,000 to ৳2,40,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Apache RTR 180 good for daily use?</h3>
<p class="faq-answer">Performance is excellent but service network issues make it risky for daily use.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is faster: RTR 160 or 180?</h3>
<p class="faq-answer">RTR 180 is significantly faster with 16.8 PS vs 160's 15.5 PS.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 180 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">TVS Apache RTR 180 offers the best performance-to-price ratio at ৳2,35,000 with powerful 180cc engine and race technology. However, it's NOT recommended for most buyers due to extremely limited service network, worst resale value in Bangladesh, and availability issues. Only suitable for performance enthusiasts living near TVS service centers who prioritize power over everything. Most riders should choose Pulsar 150 or Suzuki Gixxer instead.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- Only for enthusiasts near TVS centers</span></p>
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
		return fmt.Errorf("error creating TVS Apache RTR 180 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Apache RTR 180 (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">টিভিএস এপাচি আরটিআর ১৮০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">টিভিএস এপাচি আরটিআর ১৮০ একটি পারফরম্যান্স-কেন্দ্রিক ১৮০সিসি মোটরসাইকেল যার মূল্য ৳২,৩৫,০০০ টাকা। ১৬.৮ পিএস উৎপাদনকারী শক্তিশালী ইঞ্জিন, রেস-টিউনড সাসপেনশন এবং এলইডি লাইটিং সহ, এটি রোমাঞ্চকর পারফরম্যান্স প্রদান করে। তবে, অত্যন্ত সীমিত সার্ভিস নেটওয়ার্ক এবং খারাপ রিসেল ভ্যালু এটিকে বাংলাদেশ বাজারে একটি ঝুঁকিপূর্ণ পছন্দ করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস এপাচি আরটিআর ১৮০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">শক্তিশালী ইঞ্জিন:</strong> <span class="highlight-value">১৬.৮ পিএস পাওয়ার সহ ১৮০সিসি</span></li>
<li class="highlight-item"><strong class="highlight-label">রেস প্রযুক্তি:</strong> <span class="highlight-value">রেস-উদ্ভূত সাসপেনশন এবং সেটআপ</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">ফুল এলইডি হেডল্যাম্প এবং ডিআরএল</span></li>
<li class="highlight-item"><strong class="highlight-label">শীর্ষ পারফরম্যান্স:</strong> <span class="highlight-value">সেগমেন্টে সেরা পাওয়ার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস এপাচি আরটিআর ১৮০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার পাওয়ার:</strong> <span class="pro-description">১৮০সিসি ১৬.৮ পিএস উৎপাদন করে, সেগমেন্টে শক্তিশালীতম</span></li>
<li class="pro-item"><strong class="pro-label">রেস-টিউনড সেটআপ:</strong> <span class="pro-description">আরটিআর রেস প্রযুক্তি সহ তীক্ষ্ণ হ্যান্ডলিং</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি লাইটিং:</strong> <span class="pro-description">ডিআরএল সহ ফুল এলইডি হেডল্যাম্প</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মূল্য:</strong> <span class="pro-description">৳২,৩৫,০০০ টাকায় ফিচার-প্যাকড</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল ডিস্ক এবিএস:</strong> <span class="pro-description">এবিএস সহ সামনে এবং পিছনে ডিস্ক</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">ল্যাপ টাইমার সহ উন্নত এলসিডি</span></li>
<li class="pro-item"><strong class="pro-label">মনোশক সাসপেনশন:</strong> <span class="pro-description">রেস-স্পেক রিয়ার মনোশক</span></li>
<li class="pro-item"><strong class="pro-label">সর্বোচ্চ গতি:</strong> <span class="pro-description">১৩০+ কিমি/ঘন্টা পৌঁছাতে পারে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস এপাচি আরটিআর ১৮০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত সীমিত সার্ভিস:</strong> <span class="con-description">টিভিএস সার্ভিস সেন্টার অত্যন্ত দুর্লভ</span></li>
<li class="con-item"><strong class="con-label">সবচেয়ে খারাপ রিসেল ভ্যালু:</strong> <span class="con-description">টিভিএস বাইক দ্রুততম অবমূল্যায়ন করে</span></li>
<li class="con-item"><strong class="con-label">গড় মাইলেজ:</strong> <span class="con-description">১৮০সিসির জন্য ৩৫-৪০ কিমি/লিটার ভালো নয়</span></li>
<li class="con-item"><strong class="con-label">বিল্ড কোয়ালিটি সমস্যা:</strong> <span class="con-description">প্লাস্টিক এবং পেইন্ট মান উদ্বেগ</span></li>
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">৮০ কিমি/ঘন্টার উপরে লক্ষণীয় কম্পন</span></li>
<li class="con-item"><strong class="con-label">কম ব্র্যান্ড ভ্যালু:</strong> <span class="con-description">টিভিএস বাংলাদেশে জনপ্রিয় নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে টিভিএস এপাচি আরটিআর ১৮০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance enthusiasts near TVS centers</li>
<li class="best-for-item">Track day and racing enthusiasts</li>
<li class="best-for-item">Those wanting maximum power on budget</li>
<li class="best-for-item">Riders not worried about resale</li>
<li class="best-for-item">Tech-savvy tuning enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">টিভিএস এপাচি আরটিআর ১৮০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Most buyers due to service issues</li>
<li class="not-recommended-item">Those needing good resale value</li>
<li class="not-recommended-item">Riders in remote areas</li>
<li class="not-recommended-item">Long-term ownership plans</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">টিভিএস এপাচি আরটিআর ১৮০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৩০,০০০ - ৳২,৪০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৩,৫০০ - ৳৪,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৩৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,৭০০-৩,১০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,১০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">টিভিএস এপাচি আরটিআর ১৮০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড়ের নিচে</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- সার্ভিস সমস্যা মূল্য ক্ষতি করে</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অসাধারণ পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- শক্ত স্পোর্টি সেটআপ</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- দুর্দান্ত ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- আগ্রাসী রেস লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">টিভিএস এপাচি আরটিআর ১৮০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">টিভিএস এপাচি আরটিআর ১৮০ এর মাইলেজ কত?</h3>
<p class="faq-answer">টিভিএস এপাচি আরটিআর ১৮০ রেস-টিউনড ইঞ্জিন সহ ৩৫-৪০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে এপাচি আরটিআর ১৮০ এর দাম কত?</h3>
<p class="faq-answer">এপাচি আরটিআর ১৮০ এর দাম ৳২,৩০,০০০ থেকে ৳২,৪০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য এপাচি আরটিআর ১৮০ কি ভালো?</h3>
<p class="faq-answer">পারফরম্যান্স চমৎকার কিন্তু সার্ভিস নেটওয়ার্ক সমস্যা দৈনিক ব্যবহারের জন্য ঝুঁকিপূর্ণ করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি দ্রুত: আরটিআর ১৬০ নাকি ১৮০?</h3>
<p class="faq-answer">আরটিআর ১৬০ এর ১৫.৫ পিএস এর বিপরীতে ১৬.৮ পিএস সহ আরটিআর ১৮০ উল্লেখযোগ্যভাবে দ্রুত।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে টিভিএস এপাচি আরটিআর ১৮০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">টিভিএস এপাচি আরটিআর ১৮০ শক্তিশালী ১৮০সিসি ইঞ্জিন এবং রেস প্রযুক্তি সহ ৳২,৩৫,০০০ টাকায় সেরা পারফরম্যান্স-টু-প্রাইস অনুপাত প্রদান করে। তবে, অত্যন্ত সীমিত সার্ভিস নেটওয়ার্ক, বাংলাদেশে সবচেয়ে খারাপ রিসেল ভ্যালু এবং প্রাপ্যতা সমস্যার কারণে এটি বেশিরভাগ ক্রেতাদের জন্য সুপারিশ করা হয় না।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শুধুমাত্র টিভিএস সেন্টারের কাছাকাছি উৎসাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Apache RTR 180 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Apache RTR 180\n")

	return nil
}
