package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RoadPrinceWegoReviewSeeder handles seeding Road Prince Wego product review and translation
type RoadPrinceWegoReviewSeeder struct {
	BaseSeeder
}

// NewRoadPrinceWegoReviewSeeder creates a new RoadPrinceWegoReviewSeeder
func NewRoadPrinceWegoReviewSeeder() *RoadPrinceWegoReviewSeeder {
	return &RoadPrinceWegoReviewSeeder{BaseSeeder: BaseSeeder{name: "Road Prince Wego Review"}}
}

// Seed implements the Seeder interface
func (s *RoadPrinceWegoReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Road Prince Wego").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Road Prince Wego product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Road Prince Wego product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Road Prince Wego already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Road Prince Wego Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Road Prince Wego is an ultra-budget 110cc scooter priced at ৳90,000, offering basic scooter transportation at rock-bottom price. With simple design, decent mileage, and Chinese build, it's for extremely budget-conscious buyers seeking the cheapest scooter option despite quality compromises.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Road Prince Wego Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Rock Bottom Price:</strong> <span class="highlight-value">Cheapest scooter at ৳90,000</span></li>
<li class="highlight-item"><strong class="highlight-label">Scooter Convenience:</strong> <span class="highlight-value">Automatic transmission</span></li>
<li class="highlight-item"><strong class="highlight-label">Basic Mileage:</strong> <span class="highlight-value">50+ km/l efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Entry Level:</strong> <span class="highlight-value">Cheapest scooter entry</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Road Prince Wego Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Extremely Cheap:</strong> <span class="pro-description">Cheapest scooter available</span></li>
<li class="pro-item"><strong class="pro-label">Automatic Convenience:</strong> <span class="pro-description">No gear shifting needed</span></li>
<li class="pro-item"><strong class="pro-label">Decent Mileage:</strong> <span class="pro-description">50+ km/l fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Easy to Ride:</strong> <span class="pro-description">Simple scooter operation</span></li>
<li class="pro-item"><strong class="pro-label">Local Brand:</strong> <span class="pro-description">Local assembly and support</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Road Prince Wego Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Poor Build:</strong> <span class="con-description">Cheap Chinese quality</span></li>
<li class="con-item"><strong class="con-label">Unreliable:</strong> <span class="con-description">Frequent breakdowns expected</span></li>
<li class="con-item"><strong class="con-label">Poor Resale:</strong> <span class="con-description">Almost no resale value</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Poor service network</span></li>
<li class="con-item"><strong class="con-label">Safety Concerns:</strong> <span class="con-description">Poor braking and handling</span></li>
<li class="con-item"><strong class="con-label">No Brand Value:</strong> <span class="con-description">Unknown brand credibility</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Road Prince Wego in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extremely budget buyers</li>
<li class="best-for-item">First scooter experience</li>
<li class="best-for-item">Short-distance local use</li>
<li class="best-for-item">Backup transportation</li>
<li class="best-for-item">Those needing cheapest option</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Road Prince Wego?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Quality-conscious buyers</li>
<li class="not-recommended-item">Long-term reliable use</li>
<li class="not-recommended-item">Safety-prioritizing riders</li>
<li class="not-recommended-item">Brand image conscious buyers</li>
<li class="not-recommended-item">Those wanting good resale</li>
<li class="not-recommended-item">Daily heavy usage</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Road Prince Wego Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳85,000 - ৳95,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,500 - ৳3,500 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,600-2,000/month for 30 km daily at 50 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 2,000 km or 2 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳800 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Road Prince Wego Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent 50+ km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Poor reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Cheap but poor quality</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic 110cc power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very poor build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Minimal features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic styling</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Road Prince Wego Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Road Prince Wego?</h3>
<p class="faq-answer">Road Prince Wego delivers around 50+ km/l with 110cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Road Prince Wego?</h3>
<p class="faq-answer">Road Prince Wego is priced between ৳85,000 to ৳95,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Road Prince Wego reliable?</h3>
<p class="faq-answer">No, Road Prince Wego has poor build quality and reliability issues. Consider Honda/Yamaha instead.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Should I buy Road Prince Wego?</h3>
<p class="faq-answer">Only if you have extremely tight budget and need cheapest scooter. Otherwise, save more for Honda Activa.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Road Prince Wego in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Road Prince Wego at ৳90,000 is ONLY for extremely budget-conscious buyers needing the cheapest scooter option. While it offers automatic convenience, decent 50+ km/l mileage, and rock-bottom pricing, the very poor Chinese build quality, unreliability, safety concerns, poor resale value, and limited service support make it a risky choice. Frequent breakdowns, high maintenance headaches, and poor long-term value are expected. STRONGLY RECOMMEND saving ৳70,000 more for Honda Activa instead for much better reliability, safety, and long-term value. Only consider if budget is extremely tight.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For cheapest scooter option only (NOT RECOMMENDED)</span></p>
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
		return fmt.Errorf("error creating Road Prince Wego review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Road Prince Wego (Rating: 3.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">Road Prince Wego রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">Road Prince Wego একটি অতি-বাজেট ১১০সিসি স্কুটার যার মূল্য ৳৯০,০০০ টাকা, যা তলানি দামে প্রাথমিক স্কুটার পরিবহন প্রদান করে। সরল ডিজাইন, মোটামুটি মাইলেজ এবং চীনা বিল্ড সহ, এটি মানের সমঝোতা সত্ত্বেও সবচেয়ে সস্তা স্কুটার অপশন খোঁজা অত্যন্ত বাজেট-সচেতন ক্রেতাদের জন্য।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Road Prince Wego এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">তলানি দাম:</strong> <span class="highlight-value">৳৯০,০০০ টাকায় সবচেয়ে সস্তা স্কুটার</span></li>
<li class="highlight-item"><strong class="highlight-label">স্কুটার সুবিধা:</strong> <span class="highlight-value">অটোমেটিক ট্রান্সমিশন</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রাথমিক মাইলেজ:</strong> <span class="highlight-value">৫০+ কিমি/লিটার দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রবেশ স্তর:</strong> <span class="highlight-value">সবচেয়ে সস্তা স্কুটার প্রবেশ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Road Prince Wego এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অত্যন্ত সস্তা:</strong> <span class="pro-description">উপলব্ধ সবচেয়ে সস্তা স্কুটার</span></li>
<li class="pro-item"><strong class="pro-label">অটোমেটিক সুবিধা:</strong> <span class="pro-description">কোনো গিয়ার পরিবর্তনের প্রয়োজন নেই</span></li>
<li class="pro-item"><strong class="pro-label">মোটামুটি মাইলেজ:</strong> <span class="pro-description">৫০+ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">রাইড করা সহজ:</strong> <span class="pro-description">সরল স্কুটার অপারেশন</span></li>
<li class="pro-item"><strong class="pro-label">স্থানীয় ব্র্যান্ড:</strong> <span class="pro-description">স্থানীয় সমাবেশ এবং সহায়তা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Road Prince Wego এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত খারাপ বিল্ড:</strong> <span class="con-description">সস্তা চীনা মান</span></li>
<li class="con-item"><strong class="con-label">অনির্ভরযোগ্য:</strong> <span class="con-description">ঘন ঘন ভাঙ্গন প্রত্যাশিত</span></li>
<li class="con-item"><strong class="con-label">খারাপ রিসেল:</strong> <span class="con-description">প্রায় কোনো রিসেল ভ্যালু নেই</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">খারাপ সার্ভিস নেটওয়ার্ক</span></li>
<li class="con-item"><strong class="con-label">নিরাপত্তা উদ্বেগ:</strong> <span class="con-description">খারাপ ব্রেকিং এবং হ্যান্ডলিং</span></li>
<li class="con-item"><strong class="con-label">কোনো ব্র্যান্ড ভ্যালু নেই:</strong> <span class="con-description">অজানা ব্র্যান্ড বিশ্বাসযোগ্যতা</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে Road Prince Wego কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extremely budget buyers</li>
<li class="best-for-item">First scooter experience</li>
<li class="best-for-item">Short-distance local use</li>
<li class="best-for-item">Backup transportation</li>
<li class="best-for-item">Those needing cheapest option</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Road Prince Wego কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Quality-conscious buyers</li>
<li class="not-recommended-item">Long-term reliable use</li>
<li class="not-recommended-item">Safety-prioritizing riders</li>
<li class="not-recommended-item">Brand image conscious buyers</li>
<li class="not-recommended-item">Those wanting good resale</li>
<li class="not-recommended-item">Daily heavy usage</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Road Prince Wego এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৮৫,০০০ - ৳৯৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৫০০ - ৳৩,৫০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫০ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৬০০-২,০০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ২,০০০ কিমি বা ২ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৮০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Road Prince Wego পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- মোটামুটি ৫০+ কিমি/লিটার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খারাপ নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- সস্তা কিন্তু খারাপ মান</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- প্রাথমিক ১১০সিসি পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- প্রাথমিক আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- অত্যন্ত খারাপ বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- ন্যূনতম ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- প্রাথমিক স্টাইলিং</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Road Prince Wego সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Road Prince Wego এর মাইলেজ কত?</h3>
<p class="faq-answer">Road Prince Wego ১১০সিসি ইঞ্জিন সহ প্রায় ৫০+ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Road Prince Wego এর দাম কত?</h3>
<p class="faq-answer">Road Prince Wego এর দাম ৳৮৫,০০০ থেকে ৳৯৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Road Prince Wego কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">না, Road Prince Wego এর খারাপ বিল্ড কোয়ালিটি এবং নির্ভরযোগ্যতার সমস্যা রয়েছে। বরং হন্ডা/ইয়ামাহা বিবেচনা করুন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">আমি কি Road Prince Wego কিনব?</h3>
<p class="faq-answer">শুধুমাত্র যদি আপনার অত্যন্ত সীমিত বাজেট থাকে এবং সবচেয়ে সস্তা স্কুটার প্রয়োজন। অন্যথায়, হন্ডা অ্যাকটিভার জন্য আরও সঞ্চয় করুন।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে Road Prince Wego কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৯০,০০০ টাকায় Road Prince Wego শুধুমাত্র সবচেয়ে সস্তা স্কুটার অপশনের প্রয়োজনে অত্যন্ত বাজেট-সচেতন ক্রেতাদের জন্য। যদিও এটি অটোমেটিক সুবিধা, মোটামুটি ৫০+ কিমি/লিটার মাইলেজ এবং তলানি মূল্য প্রদান করে, অত্যন্ত খারাপ চীনা বিল্ড কোয়ালিটি, অনির্ভরযোগ্যতা, নিরাপত্তা উদ্বেগ, খারাপ রিসেল ভ্যালু এবং সীমিত সার্ভিস সহায়তা এটিকে একটি ঝুঁকিপূর্ণ পছন্দ করে তোলে। ঘন ঘন ভাঙ্গন, উচ্চ রক্ষণাবেক্ষণ মাথাব্যথা এবং খারাপ দীর্ঘমেয়াদী মূল্য প্রত্যাশিত। অনেক ভালো নির্ভরযোগ্যতা, নিরাপত্তা এবং দীর্ঘমেয়াদী মূল্যের জন্য বরং হন্ডা অ্যাকটিভার জন্য ৳৭০,০০০ টাকা বেশি সঞ্চয় করার দৃঢ় সুপারিশ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শুধু সবচেয়ে সস্তা স্কুটার অপশনের জন্য (সুপারিশ করা হয় না)</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Road Prince Wego already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Road Prince Wego\n")

	return nil
}
