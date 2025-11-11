package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RunnerBoxerReviewSeeder handles seeding Runner Boxer product review and translation
type RunnerBoxerReviewSeeder struct {
	BaseSeeder
}

// NewRunnerBoxerReviewSeeder creates a new RunnerBoxerReviewSeeder
func NewRunnerBoxerReviewSeeder() *RunnerBoxerReviewSeeder {
	return &RunnerBoxerReviewSeeder{BaseSeeder: BaseSeeder{name: "Runner Boxer Review"}}
}

// Seed implements the Seeder interface
func (s *RunnerBoxerReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Runner Boxer").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Runner Boxer product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Runner Boxer product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Runner Boxer already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Runner Boxer Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Runner Boxer is the cheapest local Bangladeshi bike at ৳88,000, offering absolute rock-bottom pricing. With extremely poor build quality, highly questionable reliability, and zero brand value, it's NOT RECOMMENDED even for extreme budget buyers. The ৳7-12,000 savings over better bikes is lost many times over in repair costs and reliability issues.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Runner Boxer Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Cheapest Bike:</strong> <span class="highlight-value">Absolute lowest price at ৳88,000</span></li>
<li class="highlight-item"><strong class="highlight-label">Lightweight:</strong> <span class="highlight-value">Very light and easy to handle</span></li>
<li class="highlight-item"><strong class="highlight-label">Wide Availability:</strong> <span class="highlight-value">Available across Bangladesh</span></li>
<li class="highlight-item"><strong class="highlight-label">Cheap Parts:</strong> <span class="highlight-value">Low-cost spare parts</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Runner Boxer Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Absolute Lowest Price:</strong> <span class="pro-description">Cheapest motorcycle in Bangladesh</span></li>
<li class="pro-item"><strong class="pro-label">Very Lightweight:</strong> <span class="pro-description">Easy to maneuver anywhere</span></li>
<li class="pro-item"><strong class="pro-label">Cheap Spare Parts:</strong> <span class="pro-description">Parts are very inexpensive</span></li>
<li class="pro-item"><strong class="pro-label">Easy Availability:</strong> <span class="pro-description">Available in most areas</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Runner Boxer Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Terrible Build Quality:</strong> <span class="con-description">Worst build quality in market</span></li>
<li class="con-item"><strong class="con-label">Extremely Unreliable:</strong> <span class="con-description">Constant breakdowns and issues</span></li>
<li class="con-item"><strong class="con-label">Zero Resale Value:</strong> <span class="con-description">No one wants to buy used Runner</span></li>
<li class="con-item"><strong class="con-label">No Brand Value:</strong> <span class="con-description">No reputation or trust</span></li>
<li class="con-item"><strong class="con-label">Poor Performance:</strong> <span class="con-description">Very weak and unreliable engine</span></li>
<li class="con-item"><strong class="con-label">Safety Concerns:</strong> <span class="con-description">Poor brakes, lighting, and structure</span></li>
<li class="con-item"><strong class="con-label">High Repair Costs:</strong> <span class="con-description">Frequent repairs negate savings</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Runner Boxer in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Absolutely no one - NOT RECOMMENDED</li>
<li class="best-for-item">Only if literally no other option exists</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Runner Boxer?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Anyone valuing reliability</li>
<li class="not-recommended-item">Daily transportation needs</li>
<li class="not-recommended-item">Safety-conscious riders</li>
<li class="not-recommended-item">Anyone with ৳7-12,000 more budget for better quality</li>
<li class="not-recommended-item">Essentially everyone</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Runner Boxer Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳85,000 - ৳92,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,500 - ৳4,000 per month (high repairs)</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,700-2,000/month for 30 km daily at 55 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 1,500 km or 1-2 months (frequent issues)</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳600 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Runner Boxer Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent if running</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- Extremely unreliable</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- False economy</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- Very weak</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- Very poor</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- Worst in market</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- No features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- Very crude</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Runner Boxer Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Should I buy Runner Boxer?</h3>
<p class="faq-answer">Absolutely NOT. Even ৳7-12,000 more for Honda CD 70 or Hero HF Deluxe is dramatically better investment.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Runner Boxer reliable?</h3>
<p class="faq-answer">No, Runner Boxer is extremely unreliable with constant mechanical failures.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Runner Boxer?</h3>
<p class="faq-answer">Runner Boxer is priced around ৳85,000 to ৳92,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Why is Runner Boxer so cheap?</h3>
<p class="faq-answer">Cheap due to extremely poor quality materials, no quality control, and no brand value.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Runner Boxer in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.0/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Runner Boxer is STRONGLY NOT RECOMMENDED at any price. At ৳88,000 it seems cheap, but terrible build quality, extreme unreliability, zero resale value, and constant repair needs make it the worst value in Bangladesh motorcycle market. The ৳7-12,000 saved versus Honda CD 70 or Hero HF Deluxe will be lost 10x over in repairs and frustration. False economy at its worst - avoid at all costs.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- STRONGLY NOT RECOMMENDED - Avoid completely</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.0,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Runner Boxer review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Runner Boxer (Rating: 3.0/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">রানার বক্সার রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">রানার বক্সার ৳৮৮,০০০ টাকায় সবচেয়ে সস্তা স্থানীয় বাংলাদেশী বাইক, যা একেবারে সর্বনিম্ন মূল্য প্রদান করে। অত্যন্ত খারাপ বিল্ড কোয়ালিটি, অত্যন্ত প্রশ্নবিদ্ধ নির্ভরযোগ্যতা এবং শূন্য ব্র্যান্ড ভ্যালু সহ, এটি এমনকি চরম বাজেট ক্রেতাদের জন্যও সুপারিশ করা হয় না। ভালো বাইকের তুলনায় ৳৭-১২,০০০ টাকা সঞ্চয় মেরামত খরচ এবং নির্ভরযোগ্যতা সমস্যায় বহুগুণ হারিয়ে যায়।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রানার বক্সার এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সবচেয়ে সস্তা বাইক:</strong> <span class="highlight-value">৳৮৮,০০০ টাকায় একেবারে সর্বনিম্ন মূল্য</span></li>
<li class="highlight-item"><strong class="highlight-label">হালকা ওজন:</strong> <span class="highlight-value">খুব হালকা এবং হ্যান্ডল করা সহজ</span></li>
<li class="highlight-item"><strong class="highlight-label">বিস্তৃত প্রাপ্যতা:</strong> <span class="highlight-value">বাংলাদেশ জুড়ে উপলব্ধ</span></li>
<li class="highlight-item"><strong class="highlight-label">সস্তা পার্টস:</strong> <span class="highlight-value">স্বল্প-মূল্যের স্পেয়ার পার্টস</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রানার বক্সার এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">একেবারে সর্বনিম্ন মূল্য:</strong> <span class="pro-description">বাংলাদেশে সবচেয়ে সস্তা মোটরসাইকেল</span></li>
<li class="pro-item"><strong class="pro-label">খুব হালকা ওজন:</strong> <span class="pro-description">যেকোনো জায়গায় চালানো সহজ</span></li>
<li class="pro-item"><strong class="pro-label">সস্তা স্পেয়ার পার্টস:</strong> <span class="pro-description">পার্টস খুব সস্তা</span></li>
<li class="pro-item"><strong class="pro-label">সহজ প্রাপ্যতা:</strong> <span class="pro-description">বেশিরভাগ এলাকায় উপলব্ধ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রানার বক্সার এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">ভয়ানক বিল্ড কোয়ালিটি:</strong> <span class="con-description">বাজারে সবচেয়ে খারাপ বিল্ড কোয়ালিটি</span></li>
<li class="con-item"><strong class="con-label">অত্যন্ত অনির্ভরযোগ্য:</strong> <span class="con-description">ধ্রুবক ভাঙ্গন এবং সমস্যা</span></li>
<li class="con-item"><strong class="con-label">শূন্য রিসেল ভ্যালু:</strong> <span class="con-description">কেউ ব্যবহৃত রানার কিনতে চায় না</span></li>
<li class="con-item"><strong class="con-label">কোনো ব্র্যান্ড ভ্যালু নেই:</strong> <span class="con-description">কোনো খ্যাতি বা বিশ্বাস নেই</span></li>
<li class="con-item"><strong class="con-label">খারাপ পারফরম্যান্স:</strong> <span class="con-description">খুব দুর্বল এবং অনির্ভরযোগ্য ইঞ্জিন</span></li>
<li class="con-item"><strong class="con-label">নিরাপত্তা উদ্বেগ:</strong> <span class="con-description">খারাপ ব্রেক, লাইটিং এবং গঠন</span></li>
<li class="con-item"><strong class="con-label">উচ্চ মেরামত খরচ:</strong> <span class="con-description">ঘন ঘন মেরামত সঞ্চয় বাতিল করে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে রানার বক্সার কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Absolutely no one - NOT RECOMMENDED</li>
<li class="best-for-item">Only if literally no other option exists</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">রানার বক্সার কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Anyone valuing reliability</li>
<li class="not-recommended-item">Daily transportation needs</li>
<li class="not-recommended-item">Safety-conscious riders</li>
<li class="not-recommended-item">Anyone with ৳7-12,000 more budget for better quality</li>
<li class="not-recommended-item">Essentially everyone</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">রানার বক্সার এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৮৫,০০০ - ৳৯২,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৫০০ - ৳৪,০০০ (উচ্চ মেরামত)</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫৫ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৭০০-২,০০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ১,৫০০ কিমি বা ১-২ মাস (ঘন ঘন সমস্যা)</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৬০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">রানার বক্সার পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ভালো যদি চলে</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- অত্যন্ত অনির্ভরযোগ্য</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- মিথ্যা অর্থনীতি</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- খুব দুর্বল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- খুব খারাপ</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- বাজারে সবচেয়ে খারাপ</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- কোনো ফিচার নেই</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- খুব অশোধিত</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">রানার বক্সার সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">আমার কি রানার বক্সার কিনা উচিত?</h3>
<p class="faq-answer">একেবারেই না। হোন্ডা সিডি ৭০ বা হিরো এইচএফ ডিলাক্সের জন্য এমনকি ৳৭-১২,০০০ টাকা বেশি নাটকীয়ভাবে ভালো বিনিয়োগ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রানার বক্সার কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">না, রানার বক্সার ধ্রুবক যান্ত্রিক ব্যর্থতা সহ অত্যন্ত অনির্ভরযোগ্য।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রানার বক্সারের দাম কত?</h3>
<p class="faq-answer">রানার বক্সারের দাম ৳৮৫,০০০ থেকে ৳৯২,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রানার বক্সার কেন এত সস্তা?</h3>
<p class="faq-answer">অত্যন্ত খারাপ মানের ম্যাটেরিয়াল, কোনো মান নিয়ন্ত্রণ এবং কোনো ব্র্যান্ড ভ্যালু না থাকার কারণে সস্তা।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে রানার বক্সার কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.0/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">রানার বক্সার যেকোনো মূল্যে দৃঢ়ভাবে সুপারিশ করা হয় না। ৳৮৮,০০০ টাকায় এটি সস্তা মনে হয়, কিন্তু ভয়ানক বিল্ড কোয়ালিটি, চরম অনির্ভরযোগ্যতা, শূন্য রিসেল ভ্যালু এবং ধ্রুবক মেরামত প্রয়োজন এটিকে বাংলাদেশ মোটরসাইকেল বাজারে সবচেয়ে খারাপ মূল্য করে তোলে। হোন্ডা সিডি ৭০ বা হিরো এইচএফ ডিলাক্সের বিপরীতে ৳৭-১২,০০০ টাকা সাশ্রয় মেরামত এবং হতাশায় ১০ গুণ হারিয়ে যাবে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- দৃঢ়ভাবে সুপারিশ করা হয় না - সম্পূর্ণ এড়িয়ে চলুন</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Runner Boxer already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Runner Boxer\n")

	return nil
}
