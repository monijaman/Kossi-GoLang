package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RoadPrinceClassicReviewSeeder handles seeding Road Prince Classic product review and translation
type RoadPrinceClassicReviewSeeder struct {
	BaseSeeder
}

// NewRoadPrinceClassicReviewSeeder creates a new RoadPrinceClassicReviewSeeder
func NewRoadPrinceClassicReviewSeeder() *RoadPrinceClassicReviewSeeder {
	return &RoadPrinceClassicReviewSeeder{BaseSeeder: BaseSeeder{name: "Road Prince Classic Review"}}
}

// Seed implements the Seeder interface
func (s *RoadPrinceClassicReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Road Prince Classic").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Road Prince Classic product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Road Prince Classic product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Road Prince Classic already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Road Prince Classic Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Road Prince Classic is a budget 100cc motorcycle priced at ৳1,15,000, offering basic commuter transportation with Chinese build quality. With simple design, decent mileage, and low price, it's for extremely budget-conscious buyers willing to compromise on quality and reliability for affordability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Road Prince Classic Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Low Price:</strong> <span class="highlight-value">Budget ৳1,15,000 pricing</span></li>
<li class="highlight-item"><strong class="highlight-label">Basic Mileage:</strong> <span class="highlight-value">60+ km/l efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Simple Design:</strong> <span class="highlight-value">Basic commuter styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Entry Level:</strong> <span class="highlight-value">Cheapest motorcycle option</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Road Prince Classic Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Very Cheap:</strong> <span class="pro-description">Budget-friendly pricing</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">60+ km/l fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Easy to Ride:</strong> <span class="pro-description">Simple operation</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance Cost:</strong> <span class="pro-description">Cheap spare parts</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Road Prince Classic Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Poor Build:</strong> <span class="con-description">Cheap Chinese quality</span></li>
<li class="con-item"><strong class="con-label">Unreliable:</strong> <span class="con-description">Frequent breakdowns</span></li>
<li class="con-item"><strong class="con-label">Poor Resale:</strong> <span class="con-description">Almost no resale value</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Poor service network</span></li>
<li class="con-item"><strong class="con-label">Safety Issues:</strong> <span class="con-description">Poor braking and handling</span></li>
<li class="con-item"><strong class="con-label">No Brand Value:</strong> <span class="con-description">Unknown brand</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Road Prince Classic in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extremely budget buyers</li>
<li class="best-for-item">First motorcycle on tight budget</li>
<li class="best-for-item">Short-distance use only</li>
<li class="best-for-item">Rural area basic transport</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Road Prince Classic?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Quality-conscious buyers</li>
<li class="not-recommended-item">Long-term reliable use</li>
<li class="not-recommended-item">Safety-prioritizing riders</li>
<li class="not-recommended-item">Brand conscious buyers</li>
<li class="not-recommended-item">Daily heavy usage</li>
<li class="not-recommended-item">Highway riding</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Road Prince Classic Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,10,000 - ৳1,20,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,200 - ৳3,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,400-1,700/month for 30 km daily at 60 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 2,000 km or 2 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳600 - ৳1,000 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Road Prince Classic Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good 60+ km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very poor</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Cheap but risky</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very basic</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- Very poor</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- Minimal</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Basic</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Road Prince Classic Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Road Prince Classic?</h3>
<p class="faq-answer">Road Prince Classic delivers around 60+ km/l with 100cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Road Prince Classic?</h3>
<p class="faq-answer">Road Prince Classic is priced between ৳1,10,000 to ৳1,20,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Road Prince Classic reliable?</h3>
<p class="faq-answer">No, Road Prince has poor build quality and reliability. Consider saving for Hero/Bajaj instead.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Should I buy Road Prince Classic?</h3>
<p class="faq-answer">Only if extremely budget tight. Better save ৳30,000 more for Bajaj CT 100 or Hero HF Deluxe.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Road Prince Classic in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.0/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Road Prince Classic at ৳1,15,000 is ONLY for extremely budget-tight buyers needing the absolute cheapest motorcycle. While it offers 60+ km/l mileage and rock-bottom pricing, the very poor Chinese build quality, unreliability, safety concerns, poor resale value, and limited service support make it a very risky choice. Frequent breakdowns and high maintenance headaches are expected. STRONGLY RECOMMEND saving ৳30,000 more for Bajaj CT 100 or Hero HF Deluxe instead for much better Japanese reliability, safety, and long-term value. Only consider if you absolutely cannot afford anything better.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For absolute cheapest motorcycle only (NOT RECOMMENDED)</span></p>
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
		return fmt.Errorf("error creating Road Prince Classic review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Road Prince Classic (Rating: 3.0/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">রোড প্রিন্স ক্লাসিক রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">রোড প্রিন্স ক্লাসিক একটি বাজেট ১০০সিসি মোটরসাইকেল যার মূল্য ৳১,১৫,০০০ টাকা, যা চীনা বিল্ড কোয়ালিটি সহ প্রাথমিক কমিউটার পরিবহন প্রদান করে। সরল ডিজাইন, মোটামুটি মাইলেজ এবং কম দাম সহ, এটি সাশ্রয়ের জন্য মান এবং নির্ভরযোগ্যতায় সমঝোতা করতে ইচ্ছুক অত্যন্ত বাজেট-সচেতন ক্রেতাদের জন্য।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রোড প্রিন্স ক্লাসিক এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">কম দাম:</strong> <span class="highlight-value">বাজেট ৳১,১৫,০০০ টাকা মূল্য</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রাথমিক মাইলেজ:</strong> <span class="highlight-value">৬০+ কিমি/লিটার দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">সরল ডিজাইন:</strong> <span class="highlight-value">প্রাথমিক কমিউটার স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রবেশ স্তর:</strong> <span class="highlight-value">সবচেয়ে সস্তা মোটরসাইকেল অপশন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রোড প্রিন্স ক্লাসিক এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">খুব সস্তা:</strong> <span class="pro-description">বাজেট-বান্ধব মূল্য</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৬০+ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">রাইড করা সহজ:</strong> <span class="pro-description">সরল অপারেশন</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ খরচ:</strong> <span class="pro-description">সস্তা খুচরা যন্ত্রাংশ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রোড প্রিন্স ক্লাসিক এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত খারাপ বিল্ড:</strong> <span class="con-description">সস্তা চীনা মান</span></li>
<li class="con-item"><strong class="con-label">অনির্ভরযোগ্য:</strong> <span class="con-description">ঘন ঘন ভাঙ্গন</span></li>
<li class="con-item"><strong class="con-label">খারাপ রিসেল:</strong> <span class="con-description">প্রায় কোনো রিসেল ভ্যালু নেই</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">খারাপ সার্ভিস নেটওয়ার্ক</span></li>
<li class="con-item"><strong class="con-label">নিরাপত্তা সমস্যা:</strong> <span class="con-description">খারাপ ব্রেকিং এবং হ্যান্ডলিং</span></li>
<li class="con-item"><strong class="con-label">কোনো ব্র্যান্ড ভ্যালু নেই:</strong> <span class="con-description">অজানা ব্র্যান্ড</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে রোড প্রিন্স ক্লাসিক কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extremely budget buyers</li>
<li class="best-for-item">First motorcycle on tight budget</li>
<li class="best-for-item">Short-distance use only</li>
<li class="best-for-item">Rural area basic transport</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">রোড প্রিন্স ক্লাসিক কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Quality-conscious buyers</li>
<li class="not-recommended-item">Long-term reliable use</li>
<li class="not-recommended-item">Safety-prioritizing riders</li>
<li class="not-recommended-item">Brand conscious buyers</li>
<li class="not-recommended-item">Daily heavy usage</li>
<li class="not-recommended-item">Highway riding</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">রোড প্রিন্স ক্লাসিক এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,১০,০০০ - ৳১,২০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,২০০ - ৳৩,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৬০ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৪০০-১,৭০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ২,০০০ কিমি বা ২ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৬০০ - ৳১,০০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">রোড প্রিন্স ক্লাসিক পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ৬০+ কিমি/লিটার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- অত্যন্ত খারাপ</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- সস্তা কিন্তু ঝুঁকিপূর্ণ</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- অত্যন্ত প্রাথমিক</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- প্রাথমিক আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- অত্যন্ত খারাপ</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- ন্যূনতম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- প্রাথমিক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">রোড প্রিন্স ক্লাসিক সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">রোড প্রিন্স ক্লাসিকের মাইলেজ কত?</h3>
<p class="faq-answer">রোড প্রিন্স ক্লাসিক ১০০সিসি ইঞ্জিন সহ প্রায় ৬০+ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রোড প্রিন্স ক্লাসিকের দাম কত?</h3>
<p class="faq-answer">রোড প্রিন্স ক্লাসিকের দাম ৳১,১০,০০০ থেকে ৳১,২০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রোড প্রিন্স ক্লাসিক কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">না, রোড প্রিন্সের খারাপ বিল্ড কোয়ালিটি এবং নির্ভরযোগ্যতা। বরং হিরো/বাজাজের জন্য সঞ্চয় করুন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">আমি কি রোড প্রিন্স ক্লাসিক কিনব?</h3>
<p class="faq-answer">শুধুমাত্র যদি অত্যন্ত বাজেট সীমিত। বাজাজ সিটি ১০০ বা হিরো এইচএফ ডিলাক্সের জন্য ৳৩০,০০০ টাকা বেশি সঞ্চয় করা ভালো।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে রোড প্রিন্স ক্লাসিক কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.0/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳১,১৫,০০০ টাকায় রোড প্রিন্স ক্লাসিক শুধুমাত্র একেবারে সবচেয়ে সস্তা মোটরসাইকেলের প্রয়োজনে অত্যন্ত বাজেট-সীমিত ক্রেতাদের জন্য। যদিও এটি ৬০+ কিমি/লিটার মাইলেজ এবং তলানি মূল্য প্রদান করে, অত্যন্ত খারাপ চীনা বিল্ড কোয়ালিটি, অনির্ভরযোগ্যতা, নিরাপত্তা উদ্বেগ, খারাপ রিসেল ভ্যালু এবং সীমিত সার্ভিস সহায়তা এটিকে একটি অত্যন্ত ঝুঁকিপূর্ণ পছন্দ করে তোলে। ঘন ঘন ভাঙ্গন এবং উচ্চ রক্ষণাবেক্ষণ মাথাব্যথা প্রত্যাশিত। অনেক ভালো জাপানি নির্ভরযোগ্যতা, নিরাপত্তা এবং দীর্ঘমেয়াদী মূল্যের জন্য বরং বাজাজ সিটি ১০০ বা হিরো এইচএফ ডিলাক্সের জন্য ৳৩০,০০০ টাকা বেশি সঞ্চয় করার দৃঢ় সুপারিশ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- একেবারে সবচেয়ে সস্তা মোটরসাইকেলের জন্য শুধু (সুপারিশ করা হয় না)</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Road Prince Classic already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Road Prince Classic\n")

	return nil
}
