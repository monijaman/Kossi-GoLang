package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// KTMDuke200ReviewSeeder handles seeding KTM Duke 200 product review and translation
type KTMDuke200ReviewSeeder struct {
	BaseSeeder
}

// NewKTMDuke200ReviewSeeder creates a new KTMDuke200ReviewSeeder
func NewKTMDuke200ReviewSeeder() *KTMDuke200ReviewSeeder {
	return &KTMDuke200ReviewSeeder{BaseSeeder: BaseSeeder{name: "KTM Duke 200 Review"}}
}

// Seed implements the Seeder interface
func (s *KTMDuke200ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "KTM Duke 200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("KTM Duke 200 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding KTM Duke 200 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for KTM Duke 200 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">KTM Duke 200 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The KTM Duke 200 is a premium 200cc naked sports bike priced at ৳8,50,000, offering exhilarating performance, aggressive styling, and excellent handling. With 25 HP power, premium features, and KTM racing pedigree, it's the ultimate choice for young enthusiasts seeking adrenaline-pumping performance and street dominance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM Duke 200 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Beast Performance:</strong> <span class="highlight-value">25 HP, 0-100 in 9.8 seconds</span></li>
<li class="highlight-item"><strong class="highlight-label">Racing DNA:</strong> <span class="highlight-value">Austrian racing pedigree</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Features:</strong> <span class="highlight-value">Full LED, ABS, digital console</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Iconic naked street fighter</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM Duke 200 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Explosive Power:</strong> <span class="pro-description">25 HP, fastest in segment</span></li>
<li class="pro-item"><strong class="pro-label">Superb Handling:</strong> <span class="pro-description">Excellent WP suspension</span></li>
<li class="pro-item"><strong class="pro-label">Premium Build:</strong> <span class="pro-description">European quality, trellis frame</span></li>
<li class="pro-item"><strong class="pro-label">Race-Ready:</strong> <span class="pro-description">Track-ready out of the box</span></li>
<li class="pro-item"><strong class="pro-label">Full LED Package:</strong> <span class="pro-description">LED headlight, DRL, taillights</span></li>
<li class="pro-item"><strong class="pro-label">Dual Channel ABS:</strong> <span class="pro-description">Advanced braking safety</span></li>
<li class="pro-item"><strong class="pro-label">Iconic Design:</strong> <span class="pro-description">Unmistakable street presence</span></li>
<li class="pro-item"><strong class="pro-label">Top Speed:</strong> <span class="pro-description">140+ km/h capability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM Duke 200 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Expensive:</strong> <span class="con-description">₹8,50,000 premium pricing</span></li>
<li class="con-item"><strong class="con-label">Poor Mileage:</strong> <span class="con-description">30-35 km/l, thirsty engine</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Very few KTM service centers</span></li>
<li class="con-item"><strong class="con-label">Expensive Parts:</strong> <span class="con-description">High maintenance and parts cost</span></li>
<li class="con-item"><strong class="con-label">Uncomfortable:</strong> <span class="con-description">Aggressive riding position</span></li>
<li class="con-item"><strong class="con-label">Not for Beginners:</strong> <span class="con-description">Too powerful for first bike</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy KTM Duke 200 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Experienced young riders</li>
<li class="best-for-item">Performance enthusiasts</li>
<li class="best-for-item">Track day lovers</li>
<li class="best-for-item">Those wanting premium sports bike</li>
<li class="best-for-item">Street fighter style fans</li>
<li class="best-for-item">Riders prioritizing power over comfort</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy KTM Duke 200?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Beginners or first-time riders</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily commuters needing comfort</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Areas without KTM service</li>
<li class="not-recommended-item">Riders wanting relaxed riding</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">KTM Duke 200 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳8,30,000 - ৳8,70,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳5,000 - ৳7,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳3,500-4,500/month for 30 km daily at 32 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳2,500 - ৳4,000 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">KTM Duke 200 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Poor mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good KTM quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Premium but worth it</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Class-leading</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Aggressive position</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- European quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Premium features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Iconic aggressive</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">KTM Duke 200 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the top speed of KTM Duke 200?</h3>
<p class="faq-answer">KTM Duke 200 can reach 140+ km/h with 25 HP power.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of KTM Duke 200?</h3>
<p class="faq-answer">KTM Duke 200 is priced between ৳8,30,000 to ৳8,70,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is KTM Duke 200 good for beginners?</h3>
<p class="faq-answer">No, Duke 200 is too powerful for beginners. Start with 150cc bikes first.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the mileage of KTM Duke 200?</h3>
<p class="faq-answer">KTM Duke 200 delivers 30-35 km/l with performance-focused engine.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy KTM Duke 200 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">KTM Duke 200 at ৳8,50,000 is THE ultimate choice for performance enthusiasts, offering explosive 25 HP power, racing pedigree, and iconic street fighter styling. With class-leading acceleration (0-100 in 9.8s), superb WP suspension, dual channel ABS, and premium European build quality, it dominates the 200cc segment. However, very high price, poor mileage (30-35 km/l), expensive maintenance, aggressive riding position, and limited service network make it suitable only for experienced riders who prioritize adrenaline and performance over comfort and economy. For those who can afford it and have riding skills, Duke 200 is the ultimate street weapon that delivers unmatched thrills.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultimate performance and street dominance</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating KTM Duke 200 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for KTM Duke 200 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">কেটিএম ডিউক ২০০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">কেটিএম ডিউক ২০০ একটি প্রিমিয়াম ২০০সিসি নেকেড স্পোর্টস বাইক যার মূল্য ৳৮,৫০,০০০ টাকা, যা রোমাঞ্চকর পারফরম্যান্স, আক্রমণাত্মক স্টাইলিং এবং চমৎকার হ্যান্ডলিং প্রদান করে। ২৫ এইচপি পাওয়ার, প্রিমিয়াম ফিচার এবং কেটিএম রেসিং ঐতিহ্য সহ, এটি অ্যাড্রেনালিন-পাম্পিং পারফরম্যান্স এবং রাস্তার আধিপত্য খোঁজা তরুণ উৎসাহীদের জন্য চূড়ান্ত পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কেটিএম ডিউক ২০০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">দানবীয় পারফরম্যান্স:</strong> <span class="highlight-value">২৫ এইচপি, ০-১০০ ৯.৮ সেকেন্ডে</span></li>
<li class="highlight-item"><strong class="highlight-label">রেসিং ডিএনএ:</strong> <span class="highlight-value">অস্ট্রিয়ান রেসিং ঐতিহ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম ফিচার:</strong> <span class="highlight-value">সম্পূর্ণ এলইডি, এবিএস, ডিজিটাল কনসোল</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="highlight-value">আইকনিক নেকেড স্ট্রিট ফাইটার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কেটিএম ডিউক ২০০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">বিস্ফোরক পাওয়ার:</strong> <span class="pro-description">২৫ এইচপি, সেগমেন্টে দ্রুততম</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত হ্যান্ডলিং:</strong> <span class="pro-description">চমৎকার ডব্লিউপি সাসপেনশন</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বিল্ড:</strong> <span class="pro-description">ইউরোপীয় মান, ট্রেলিস ফ্রেম</span></li>
<li class="pro-item"><strong class="pro-label">রেস-রেডি:</strong> <span class="pro-description">বক্সের বাইরে ট্র্যাক-রেডি</span></li>
<li class="pro-item"><strong class="pro-label">সম্পূর্ণ এলইডি প্যাকেজ:</strong> <span class="pro-description">এলইডি হেডলাইট, ডিআরএল, টেইললাইট</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল চ্যানেল এবিএস:</strong> <span class="pro-description">উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="pro-item"><strong class="pro-label">আইকনিক ডিজাইন:</strong> <span class="pro-description">অবিস্মরণীয় রাস্তার উপস্থিতি</span></li>
<li class="pro-item"><strong class="pro-label">টপ স্পিড:</strong> <span class="pro-description">১৪০+ কিমি/ঘণ্টা সক্ষমতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কেটিএম ডিউক ২০০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত ব্যয়বহুল:</strong> <span class="con-description">৳৮,৫০,০০০ প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">খারাপ মাইলেজ:</strong> <span class="con-description">৩০-৩৫ কিমি/লিটার, জ্বালানি ক্ষুধার্ত</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">খুব কম কেটিএম সার্ভিস সেন্টার</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল যন্ত্রাংশ:</strong> <span class="con-description">উচ্চ রক্ষণাবেক্ষণ এবং যন্ত্রাংশ খরচ</span></li>
<li class="con-item"><strong class="con-label">অস্বস্তিকর:</strong> <span class="con-description">আক্রমণাত্মক রাইডিং অবস্থান</span></li>
<li class="con-item"><strong class="con-label">শিক্ষানবিসদের জন্য নয়:</strong> <span class="con-description">প্রথম বাইকের জন্য অত্যধিক শক্তিশালী</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে কেটিএম ডিউক ২০০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Experienced young riders</li>
<li class="best-for-item">Performance enthusiasts</li>
<li class="best-for-item">Track day lovers</li>
<li class="best-for-item">Those wanting premium sports bike</li>
<li class="best-for-item">Street fighter style fans</li>
<li class="best-for-item">Riders prioritizing power over comfort</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">কেটিএম ডিউক ২০০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Beginners or first-time riders</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily commuters needing comfort</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Areas without KTM service</li>
<li class="not-recommended-item">Riders wanting relaxed riding</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">কেটিএম ডিউক ২০০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৮,৩০,০০০ - ৳৮,৭০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৫,০০০ - ৳৭,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৩২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳৩,৫০০-৪,৫০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳২,৫০০ - ৳৪,০০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">কেটিএম ডিউক ২০০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- খারাপ মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো কেটিএম মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- প্রিমিয়াম কিন্তু মূল্যবান</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- শ্রেণীতে নেতৃস্থানীয়</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- আক্রমণাত্মক অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- ইউরোপীয় মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- প্রিমিয়াম ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- আইকনিক আক্রমণাত্মক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">কেটিএম ডিউক ২০০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">কেটিএম ডিউক ২০০ এর টপ স্পিড কত?</h3>
<p class="faq-answer">কেটিএম ডিউক ২০০ ২৫ এইচপি পাওয়ার সহ ১৪০+ কিমি/ঘণ্টা পৌঁছাতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কেটিএম ডিউক ২০০ এর দাম কত?</h3>
<p class="faq-answer">কেটিএম ডিউক ২০০ এর দাম ৳৮,৩০,০০০ থেকে ৳৮,৭০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">শিক্ষানবিসদের জন্য কেটিএম ডিউক ২০০ কি ভালো?</h3>
<p class="faq-answer">না, ডিউক ২০০ শিক্ষানবিসদের জন্য অত্যধিক শক্তিশালী। প্রথমে ১৫০সিসি বাইক দিয়ে শুরু করুন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কেটিএম ডিউক ২০০ এর মাইলেজ কত?</h3>
<p class="faq-answer">কেটিএম ডিউক ২০০ পারফরম্যান্স-কেন্দ্রিক ইঞ্জিন সহ ৩০-৩৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে কেটিএম ডিউক ২০০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৮,৫০,০০০ টাকায় কেটিএম ডিউক ২০০ পারফরম্যান্স উৎসাহীদের জন্য চূড়ান্ত পছন্দ, যা বিস্ফোরক ২৫ এইচপি পাওয়ার, রেসিং ঐতিহ্য এবং আইকনিক স্ট্রিট ফাইটার স্টাইলিং প্রদান করে। শ্রেণীতে নেতৃস্থানীয় ত্বরণ (০-১০০ ৯.৮ সেকেন্ডে), দুর্দান্ত ডব্লিউপি সাসপেনশন, ডুয়াল চ্যানেল এবিএস এবং প্রিমিয়াম ইউরোপীয় বিল্ড কোয়ালিটি সহ, এটি ২০০সিসি সেগমেন্টে আধিপত্য করে। তবে, অত্যন্ত উচ্চ মূল্য, খারাপ মাইলেজ (৩০-৩৫ কিমি/লিটার), ব্যয়বহুল রক্ষণাবেক্ষণ, আক্রমণাত্মক রাইডিং অবস্থান এবং সীমিত সার্ভিস নেটওয়ার্ক এটিকে শুধুমাত্র অভিজ্ঞ রাইডারদের জন্য উপযুক্ত করে যারা আরাম এবং অর্থনীতির চেয়ে অ্যাড্রেনালিন এবং পারফরম্যান্সকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চূড়ান্ত পারফরম্যান্স এবং রাস্তার আধিপত্যের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for KTM Duke 200 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for KTM Duke 200\n")

	return nil
}
