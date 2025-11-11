package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSMetroReviewSeeder handles seeding TVS Metro product review and translation
type TVSMetroReviewSeeder struct {
	BaseSeeder
}

// NewTVSMetroReviewSeeder creates a new TVSMetroReviewSeeder
func NewTVSMetroReviewSeeder() *TVSMetroReviewSeeder {
	return &TVSMetroReviewSeeder{BaseSeeder: BaseSeeder{name: "TVS Metro Review"}}
}

// Seed implements the Seeder interface
func (s *TVSMetroReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Metro").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("TVS Metro product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding TVS Metro product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Metro already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Metro Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The TVS Metro is an entry-level 100cc commuter bike priced at ৳90,000, offering basic transportation with TVS reliability. Simple design, decent fuel economy around 60 km/l, and affordable maintenance make it suitable for budget-conscious buyers needing basic daily transport.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Metro Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Very Affordable:</strong> <span class="highlight-value">One of cheapest bikes at ৳90,000</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">60-65 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Low Maintenance:</strong> <span class="highlight-value">Simple, cheap to maintain</span></li>
<li class="highlight-item"><strong class="highlight-label">Lightweight:</strong> <span class="highlight-value">Easy handling in city traffic</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Metro Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Rock Bottom Price:</strong> <span class="pro-description">At ৳90,000 it's extremely affordable</span></li>
<li class="pro-item"><strong class="pro-label">Outstanding Mileage:</strong> <span class="pro-description">60-65 km/l minimizes fuel costs</span></li>
<li class="pro-item"><strong class="pro-label">Very Light Weight:</strong> <span class="pro-description">Easy to maneuver anywhere</span></li>
<li class="pro-item"><strong class="pro-label">Simple Maintenance:</strong> <span class="pro-description">Basic design, dirt cheap parts</span></li>
<li class="pro-item"><strong class="pro-label">TVS Reliability:</strong> <span class="pro-description">Proven Indian engineering</span></li>
<li class="pro-item"><strong class="pro-label">Easy Availability:</strong> <span class="pro-description">Parts readily available</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Metro Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Basic Features:</strong> <span class="con-description">No modern features whatsoever</span></li>
<li class="con-item"><strong class="con-label">Poor Build Quality:</strong> <span class="con-description">Low-grade materials used</span></li>
<li class="con-item"><strong class="con-label">Outdated Styling:</strong> <span class="con-description">Very old-fashioned design</span></li>
<li class="con-item"><strong class="con-label">Weak Performance:</strong> <span class="con-description">100cc engine lacks power</span></li>
<li class="con-item"><strong class="con-label">Uncomfortable:</strong> <span class="con-description">Hard seat, basic suspension</span></li>
<li class="con-item"><strong class="con-label">No Resale Value:</strong> <span class="con-description">Extremely poor resale value</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy TVS Metro in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extreme budget buyers</li>
<li class="best-for-item">Short city commutes only</li>
<li class="best-for-item">Those prioritizing fuel economy above all</li>
<li class="best-for-item">Backup/second bike for household</li>
<li class="best-for-item">Young students with tight budget</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy TVS Metro?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Anyone who can afford ৳20,000 more for Honda CD 70</li>
<li class="not-recommended-item">Long distance riding</li>
<li class="not-recommended-item">Image-conscious buyers</li>
<li class="not-recommended-item">Those wanting any comfort or features</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Metro Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳85,000 - ৳95,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,800 - ৳2,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,500-1,700/month for 30 km daily at 62 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 2,500 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳400 - ৳600 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Metro Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Outstanding mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Cheap but very basic</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very weak</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Uncomfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Poor quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- No features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- Very outdated</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">TVS Metro Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of TVS Metro?</h3>
<p class="faq-answer">TVS Metro delivers excellent 60-65 km/l fuel efficiency.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of TVS Metro?</h3>
<p class="faq-answer">TVS Metro is priced around ৳85,000 to ৳95,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is TVS Metro reliable?</h3>
<p class="faq-answer">Moderately reliable but Honda CD 70 is much better for similar price.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: TVS Metro or Honda CD 70?</h3>
<p class="faq-answer">Honda CD 70 is far superior in build quality, reliability, and resale value.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Metro in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">TVS Metro is pure rock-bottom budget transportation at ৳90,000. While it offers excellent 60-65 km/l mileage and cheap maintenance, the very basic features, poor build quality, uncomfortable ride, and zero resale value make it hard to recommend. For just ৳10,000-15,000 more, Honda CD 70 offers dramatically better quality, reliability, and resale value. Only consider Metro if budget is absolutely critical.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- Only for extreme budget constraints</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating TVS Metro review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Metro (Rating: 3.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">টিভিএস মেট্রো রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">টিভিএস মেট্রো একটি এন্ট্রি-লেভেল ১০০সিসি কমিউটার বাইক যার মূল্য ৳৯০,০০০ টাকা, যা টিভিএস নির্ভরযোগ্যতা সহ মৌলিক পরিবহন প্রদান করে। সাধারণ ডিজাইন, প্রায় ৬০ কিমি/লিটার ভালো জ্বালানি অর্থনীতি এবং সাশ্রয়ী রক্ষণাবেক্ষণ এটিকে মৌলিক দৈনিক পরিবহনের প্রয়োজন বাজেট-সচেতন ক্রেতাদের জন্য উপযুক্ত করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস মেট্রো এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">খুব সাশ্রয়ী:</strong> <span class="highlight-value">৳৯০,০০০ টাকায় সবচেয়ে সস্তা বাইকগুলির একটি</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">৬০-৬৫ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">কম রক্ষণাবেক্ষণ:</strong> <span class="highlight-value">সহজ, রক্ষণাবেক্ষণ সস্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">হালকা ওজন:</strong> <span class="highlight-value">শহরের ট্রাফিকে সহজ হ্যান্ডলিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস মেট্রো এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">সর্বনিম্ন মূল্য:</strong> <span class="pro-description">৳৯০,০০০ টাকায় এটি অত্যন্ত সাশ্রয়ী</span></li>
<li class="pro-item"><strong class="pro-label">অসাধারণ মাইলেজ:</strong> <span class="pro-description">৬০-৬৫ কিমি/লিটার জ্বালানি খরচ কমিয়ে দেয়</span></li>
<li class="pro-item"><strong class="pro-label">খুব হালকা ওজন:</strong> <span class="pro-description">যেকোনো জায়গায় চালানো সহজ</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="pro-description">বেসিক ডিজাইন, অত্যন্ত সস্তা পার্টস</span></li>
<li class="pro-item"><strong class="pro-label">টিভিএস নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত ভারতীয় ইঞ্জিনিয়ারিং</span></li>
<li class="pro-item"><strong class="pro-label">সহজ প্রাপ্যতা:</strong> <span class="pro-description">পার্টস সহজলভ্য</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস মেট্রো এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">খুব বেসিক ফিচার:</strong> <span class="con-description">কোনো আধুনিক ফিচার নেই</span></li>
<li class="con-item"><strong class="con-label">খারাপ বিল্ড কোয়ালিটি:</strong> <span class="con-description">নিম্ন-মানের ম্যাটেরিয়াল ব্যবহৃত</span></li>
<li class="con-item"><strong class="con-label">পুরানো স্টাইলিং:</strong> <span class="con-description">খুব পুরানো ডিজাইন</span></li>
<li class="con-item"><strong class="con-label">দুর্বল পারফরম্যান্স:</strong> <span class="con-description">১০০সিসি ইঞ্জিন শক্তির অভাব</span></li>
<li class="con-item"><strong class="con-label">অস্বস্তিকর:</strong> <span class="con-description">শক্ত সিট, বেসিক সাসপেনশন</span></li>
<li class="con-item"><strong class="con-label">কোনো রিসেল ভ্যালু নেই:</strong> <span class="con-description">অত্যন্ত খারাপ রিসেল ভ্যালু</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে টিভিএস মেট্রো কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extreme budget buyers</li>
<li class="best-for-item">Short city commutes only</li>
<li class="best-for-item">Those prioritizing fuel economy above all</li>
<li class="best-for-item">Backup/second bike for household</li>
<li class="best-for-item">Young students with tight budget</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">টিভিএস মেট্রো কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Anyone who can afford ৳20,000 more for Honda CD 70</li>
<li class="not-recommended-item">Long distance riding</li>
<li class="not-recommended-item">Image-conscious buyers</li>
<li class="not-recommended-item">Those wanting any comfort or features</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">টিভিএস মেট্রো এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৮৫,০০০ - ৳৯৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৮০০ - ৳২,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৬২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৫০০-১,৭০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ২,৫০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৪০০ - ৳৬০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">টিভিএস মেট্রো পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অসাধারণ মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- সস্তা কিন্তু খুব বেসিক</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব দুর্বল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- অস্বস্তিকর</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খারাপ মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- কোনো ফিচার নেই</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- খুব পুরানো</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">টিভিএস মেট্রো সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">টিভিএস মেট্রোর মাইলেজ কত?</h3>
<p class="faq-answer">টিভিএস মেট্রো চমৎকার ৬০-৬৫ কিমি/লিটার জ্বালানি দক্ষতা প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">টিভিএস মেট্রোর দাম কত?</h3>
<p class="faq-answer">টিভিএস মেট্রোর দাম ৳৮৫,০০০ থেকে ৳৯৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">টিভিএস মেট্রো কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">মোটামুটি নির্ভরযোগ্য কিন্তু হোন্ডা সিডি ৭০ একই দামে অনেক ভালো।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: টিভিএস মেট্রো নাকি হোন্ডা সিডি ৭০?</h3>
<p class="faq-answer">হোন্ডা সিডি ৭০ বিল্ড কোয়ালিটি, নির্ভরযোগ্যতা এবং রিসেল ভ্যালুতে অনেক উন্নত।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে টিভিএস মেট্রো কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.5/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">টিভিএস মেট্রো হল ৳৯০,০০০ টাকায় খাঁটি সর্বনিম্ন বাজেট পরিবহন। যদিও এটি চমৎকার ৬০-৬৫ কিমি/লিটার মাইলেজ এবং সস্তা রক্ষণাবেক্ষণ প্রদান করে, খুব বেসিক ফিচার, খারাপ বিল্ড কোয়ালিটি, অস্বস্তিকর রাইড এবং শূন্য রিসেল ভ্যালু এটিকে সুপারিশ করা কঠিন করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শুধুমাত্র চরম বাজেট সীমাবদ্ধতার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Metro already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Metro\n")

	return nil
}
