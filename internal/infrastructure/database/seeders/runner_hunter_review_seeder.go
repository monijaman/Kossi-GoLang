package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RunnerHunterReviewSeeder handles seeding Runner Hunter product review and translation
type RunnerHunterReviewSeeder struct {
	BaseSeeder
}

// NewRunnerHunterReviewSeeder creates a new RunnerHunterReviewSeeder
func NewRunnerHunterReviewSeeder() *RunnerHunterReviewSeeder {
	return &RunnerHunterReviewSeeder{BaseSeeder: BaseSeeder{name: "Runner Hunter Review"}}
}

// Seed implements the Seeder interface
func (s *RunnerHunterReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Runner Hunter").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Runner Hunter product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Runner Hunter product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Runner Hunter already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Runner Hunter Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Runner Hunter is a local Bangladeshi brand 100cc motorcycle priced at ৳95,000, offering rock-bottom pricing for extreme budget buyers. With very basic features and questionable build quality, it's only suitable for those with absolutely minimal budget who need basic transportation, though reliability concerns make it a risky purchase.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Runner Hunter Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Extremely Cheap:</strong> <span class="highlight-value">One of cheapest bikes at ৳95,000</span></li>
<li class="highlight-item"><strong class="highlight-label">Local Brand:</strong> <span class="highlight-value">Bangladeshi motorcycle brand</span></li>
<li class="highlight-item"><strong class="highlight-label">Decent Mileage:</strong> <span class="highlight-value">55-60 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Easy Availability:</strong> <span class="highlight-value">Available across Bangladesh</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Runner Hunter Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Rock Bottom Price:</strong> <span class="pro-description">Extremely affordable at ৳95,000</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">55-60 km/l fuel economy</span></li>
<li class="pro-item"><strong class="pro-label">Wide Availability:</strong> <span class="pro-description">Available in most areas</span></li>
<li class="pro-item"><strong class="pro-label">Cheap Parts:</strong> <span class="pro-description">Low-cost spare parts</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight:</strong> <span class="pro-description">Easy to handle</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Runner Hunter Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Poor Build Quality:</strong> <span class="con-description">Extremely cheap materials and construction</span></li>
<li class="con-item"><strong class="con-label">Questionable Reliability:</strong> <span class="con-description">Frequent breakdowns reported</span></li>
<li class="con-item"><strong class="con-label">No Brand Value:</strong> <span class="con-description">No resale value whatsoever</span></li>
<li class="con-item"><strong class="con-label">Very Basic Features:</strong> <span class="con-description">Absolutely no features</span></li>
<li class="con-item"><strong class="con-label">Poor Performance:</strong> <span class="con-description">Weak and unreliable engine</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Poor service network quality</span></li>
<li class="con-item"><strong class="con-label">Outdated Design:</strong> <span class="con-description">Very old-fashioned looks</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Runner Hunter in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Absolutely minimal budget (no other option)</li>
<li class="best-for-item">Very short distance use only</li>
<li class="best-for-item">Backup/third bike for household</li>
<li class="best-for-item">Rural areas with no alternatives</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Runner Hunter?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Anyone who can afford ৳5-10,000 more for better quality</li>
<li class="not-recommended-item">Daily reliable transportation needs</li>
<li class="not-recommended-item">Riders valuing safety and reliability</li>
<li class="not-recommended-item">Long-term ownership</li>
<li class="not-recommended-item">Anyone with choice</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Runner Hunter Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳90,000 - ৳1,00,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,000 - ৳3,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,600-1,900/month for 30 km daily at 57 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 2,000 km or 2 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳500 - ৳1,000 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Runner Hunter Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very unreliable</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Cheap but risky</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very weak</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Poor comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- Terrible quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- No features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- Very outdated</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Runner Hunter Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Is Runner Hunter reliable?</h3>
<p class="faq-answer">No, Runner bikes have poor reliability record with frequent mechanical issues.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Runner Hunter?</h3>
<p class="faq-answer">Runner Hunter claims 55-60 km/l but actual mileage varies widely.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Should I buy Runner Hunter?</h3>
<p class="faq-answer">Only if absolutely no budget for better options like Honda CD 70 or Hero HF Deluxe.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Runner Hunter?</h3>
<p class="faq-answer">Runner Hunter is priced around ৳90,000 to ৳1,00,000.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Runner Hunter in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Runner Hunter is NOT RECOMMENDED despite ৳95,000 price tag. While cheap and fuel efficient (55-60 km/l), the extremely poor build quality, questionable reliability, zero resale value, and frequent breakdown issues make it a false economy. Spending just ৳5-10,000 more for Honda CD 70 or Hero HF Deluxe provides dramatically better quality, reliability, and resale value. Only consider if absolutely no other budget option exists.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- NOT RECOMMENDED - Only for desperate budget buyers</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.2,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Runner Hunter review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Runner Hunter (Rating: 3.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">রানার হান্টার রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">রানার হান্টার একটি স্থানীয় বাংলাদেশী ব্র্যান্ড ১০০সিসি মোটরসাইকেল যার মূল্য ৳৯৫,০০০ টাকা, যা চরম বাজেট ক্রেতাদের জন্য সর্বনিম্ন মূল্য প্রদান করে। খুব বেসিক ফিচার এবং প্রশ্নবিদ্ধ বিল্ড কোয়ালিটি সহ, এটি শুধুমাত্র তাদের জন্য উপযুক্ত যাদের একেবারে ন্যূনতম বাজেট আছে এবং মৌলিক পরিবহন প্রয়োজন, যদিও নির্ভরযোগ্যতা উদ্বেগ এটিকে ঝুঁকিপূর্ণ ক্রয় করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রানার হান্টার এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">অত্যন্ত সস্তা:</strong> <span class="highlight-value">৳৯৫,০০০ টাকায় সবচেয়ে সস্তা বাইকগুলির একটি</span></li>
<li class="highlight-item"><strong class="highlight-label">স্থানীয় ব্র্যান্ড:</strong> <span class="highlight-value">বাংলাদেশী মোটরসাইকেল ব্র্যান্ড</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৫৫-৬০ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">সহজ প্রাপ্যতা:</strong> <span class="highlight-value">বাংলাদেশ জুড়ে উপলব্ধ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রানার হান্টার এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">সর্বনিম্ন মূল্য:</strong> <span class="pro-description">৳৯৫,০০০ টাকায় অত্যন্ত সাশ্রয়ী</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৫৫-৬০ কিমি/লিটার জ্বালানি সাশ্রয়</span></li>
<li class="pro-item"><strong class="pro-label">বিস্তৃত প্রাপ্যতা:</strong> <span class="pro-description">বেশিরভাগ এলাকায় উপলব্ধ</span></li>
<li class="pro-item"><strong class="pro-label">সস্তা পার্টস:</strong> <span class="pro-description">স্বল্প-মূল্যের স্পেয়ার পার্টস</span></li>
<li class="pro-item"><strong class="pro-label">হালকা ওজন:</strong> <span class="pro-description">হ্যান্ডল করা সহজ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রানার হান্টার এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">খুব খারাপ বিল্ড কোয়ালিটি:</strong> <span class="con-description">অত্যন্ত সস্তা ম্যাটেরিয়াল এবং নির্মাণ</span></li>
<li class="con-item"><strong class="con-label">প্রশ্নবিদ্ধ নির্ভরযোগ্যতা:</strong> <span class="con-description">ঘন ঘন ভাঙা রিপোর্ট করা হয়েছে</span></li>
<li class="con-item"><strong class="con-label">কোনো ব্র্যান্ড ভ্যালু নেই:</strong> <span class="con-description">কোনো রিসেল ভ্যালু নেই</span></li>
<li class="con-item"><strong class="con-label">খুব বেসিক ফিচার:</strong> <span class="con-description">একেবারে কোনো ফিচার নেই</span></li>
<li class="con-item"><strong class="con-label">খারাপ পারফরম্যান্স:</strong> <span class="con-description">দুর্বল এবং অনির্ভরযোগ্য ইঞ্জিন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">খারাপ সার্ভিস নেটওয়ার্ক মান</span></li>
<li class="con-item"><strong class="con-label">পুরানো ডিজাইন:</strong> <span class="con-description">খুব পুরানো লুক</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে রানার হান্টার কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Absolutely minimal budget (no other option)</li>
<li class="best-for-item">Very short distance use only</li>
<li class="best-for-item">Backup/third bike for household</li>
<li class="best-for-item">Rural areas with no alternatives</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">রানার হান্টার কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Anyone who can afford ৳5-10,000 more for better quality</li>
<li class="not-recommended-item">Daily reliable transportation needs</li>
<li class="not-recommended-item">Riders valuing safety and reliability</li>
<li class="not-recommended-item">Long-term ownership</li>
<li class="not-recommended-item">Anyone with choice</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">রানার হান্টার এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৯০,০০০ - ৳১,০০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,০০০ - ৳৩,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৬০০-১,৯০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ২,০০০ কিমি বা ২ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৫০০ - ৳১,০০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">রানার হান্টার পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব অনির্ভরযোগ্য</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- সস্তা কিন্তু ঝুঁকিপূর্ণ</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব দুর্বল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খারাপ আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- ভয়ানক মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(1/5)</span> <span class="rating-note">- কোনো ফিচার নেই</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- খুব পুরানো</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">রানার হান্টার সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">রানার হান্টার কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">না, রানার বাইকের ঘন ঘন যান্ত্রিক সমস্যা সহ খারাপ নির্ভরযোগ্যতা রেকর্ড আছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রানার হান্টারের মাইলেজ কত?</h3>
<p class="faq-answer">রানার হান্টার ৫৫-৬০ কিমি/লিটার দাবি করে কিন্তু প্রকৃত মাইলেজ ব্যাপকভাবে পরিবর্তিত হয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">আমার কি রানার হান্টার কিনা উচিত?</h3>
<p class="faq-answer">শুধুমাত্র যদি হোন্ডা সিডি ৭০ বা হিরো এইচএফ ডিলাক্সের মতো ভালো অপশনের জন্য একেবারে কোনো বাজেট না থাকে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রানার হান্টারের দাম কত?</h3>
<p class="faq-answer">রানার হান্টারের দাম ৳৯০,০০০ থেকে ৳১,০০,০০০ টাকার মধ্যে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে রানার হান্টার কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">রানার হান্টার ৳৯৫,০০০ টাকা দামের ট্যাগ সত্ত্বেও সুপারিশ করা হয় না। যদিও সস্তা এবং জ্বালানি দক্ষ (৫৫-৬০ কিমি/লিটার), অত্যন্ত খারাপ বিল্ড কোয়ালিটি, প্রশ্নবিদ্ধ নির্ভরযোগ্যতা, শূন্য রিসেল ভ্যালু এবং ঘন ঘন ভাঙ্গন সমস্যা এটিকে মিথ্যা অর্থনীতি করে তোলে। হোন্ডা সিডি ৭০ বা হিরো এইচএফ ডিলাক্সের জন্য মাত্র ৳৫-১০,০০০ টাকা বেশি খরচ করা নাটকীয়ভাবে ভালো মান, নির্ভরযোগ্যতা এবং রিসেল ভ্যালু প্রদান করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সুপারিশ করা হয় না - শুধুমাত্র মরিয়া বাজেট ক্রেতাদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Runner Hunter already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Runner Hunter\n")

	return nil
}
