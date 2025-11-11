package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaR15MReview handles seeding Yamaha R15M product review and translation
type YamahaR15MReview struct {
	BaseSeeder
}

// NewYamahaR15MReviewSeeder creates a new YamahaR15MReview
func NewYamahaR15MReviewSeeder() *YamahaR15MReview {
	return &YamahaR15MReview{BaseSeeder: BaseSeeder{name: "Yamaha R15M Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaR15MReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R15M").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha R15M product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha R15M product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha R15M already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha R15M Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha R15M is a track-oriented 155cc sports motorcycle featuring VVA technology, upside-down forks, quickshifter, and dual-channel ABS. Priced at ৳6,25,000, it combines racing-inspired features with street usability including aerodynamic design, track-tuned suspension, and Yamaha's proven R-series DNA for enthusiasts seeking authentic track performance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R15M Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Track-Focused Design:</strong> <span class="highlight-value">Racing-inspired aerodynamics and ergonomics</span></li>
<li class="highlight-item"><strong class="highlight-label">USD Front Suspension:</strong> <span class="highlight-value">Upside-down forks for enhanced handling</span></li>
<li class="highlight-item"><strong class="highlight-label">Quickshifter Technology:</strong> <span class="highlight-value">Clutchless upshifts for racing performance</span></li>
<li class="highlight-item"><strong class="highlight-label">VVA Engine:</strong> <span class="highlight-value">Variable valve actuation for optimal power</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R15M Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Professional Racing Features:</strong> <span class="pro-description">Quickshifter, USD forks, and track-oriented suspension setup</span></li>
<li class="pro-item"><strong class="pro-label">VVA Performance:</strong> <span class="pro-description">Variable valve actuation optimizes power delivery throughout rev range</span></li>
<li class="pro-item"><strong class="pro-label">Dual Channel ABS:</strong> <span class="pro-description">Enhanced braking safety with ABS on both wheels</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Build Quality:</strong> <span class="pro-description">Premium components and Yamaha's proven engineering</span></li>
<li class="pro-item"><strong class="pro-label">Aerodynamic Design:</strong> <span class="pro-description">Race-inspired bodywork with excellent wind protection</span></li>
<li class="pro-item"><strong class="pro-label">Track Performance:</strong> <span class="pro-description">Optimized for circuit riding and track day events</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R15M Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Purchase Price:</strong> <span class="con-description">Premium pricing limits accessibility for average buyers</span></li>
<li class="con-item"><strong class="con-label">Track-Focused Ergonomics:</strong> <span class="con-description">Aggressive riding position may be uncomfortable for daily use</span></li>
<li class="con-item"><strong class="con-label">Limited Daily Practicality:</strong> <span class="con-description">Track orientation compromises comfort for city riding</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">Premium components require specialized service and parts</span></li>
<li class="con-item"><strong class="con-label">Moderate Fuel Economy:</strong> <span class="con-description">Performance focus results in lower fuel efficiency</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha R15M in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Track day enthusiasts</li>
<li class="best-for-item">Racing and sport riding fans</li>
<li class="best-for-item">Experienced riders seeking performance</li>
<li class="best-for-item">Weekend track riders</li>
<li class="best-for-item">Those prioritizing handling and performance</li>
<li class="best-for-item">Yamaha R-series enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha R15M?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily city commuting</li>
<li class="not-recommended-item">Beginner or novice riders</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Those needing maximum comfort</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R15M Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳6,25,000 - Premium for track-focused 155cc features</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳13-17 per km with premium maintenance</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳10-13 per km (33-38 km/l with performance tuning)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳11,000-18,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R15M Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Excellent track performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- Performance-focused tuning</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.3</span> <span class="rating-note">- Track-oriented position</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Racing-inspired design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good for track features</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Premium Yamaha quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha R15M Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What makes the R15M different from standard R15?</h3>
<p class="faq-answer">The R15M features upside-down front forks, quickshifter, track-tuned suspension, racing ergonomics, and specialized components designed specifically for track performance and racing applications.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the quickshifter worth it for track riding?</h3>
<p class="faq-answer">Yes, the quickshifter allows seamless upshifts without clutch operation, maintaining momentum and improving lap times. It's particularly beneficial for track riding and racing scenarios.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How do USD forks improve handling?</h3>
<p class="faq-answer">USD (upside-down) forks provide greater rigidity and better feedback compared to conventional forks, improving handling precision, stability during braking, and overall cornering performance.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R15M in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳6,25,000, the Yamaha R15M offers authentic track performance with professional features including quickshifter, USD forks, VVA technology, and dual-channel ABS in the 155cc segment. It delivers exceptional handling, build quality, and racing DNA for serious track enthusiasts. However, the premium pricing, track-focused ergonomics, limited daily practicality, expensive maintenance, and moderate fuel economy make it suitable primarily for dedicated track riders and performance enthusiasts who prioritize racing features over comfort and everyday usability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For authentic track performance with professional features</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha R15M review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha R15M (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা আর১৫এম রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা আর১৫এম হল একটি ট্র্যাক-ভিত্তিক ১৫৫সিসি স্পোর্টস মোটরসাইকেল যেখানে ভিভিএ প্রযুক্তি, আপসাইড-ডাউন ফর্কস, কুইকশিফটার এবং ডুয়াল-চ্যানেল এবিএস রয়েছে। ৳৬,২৫,০০০ টাকায় মূল্যায়িত, এটি অ্যারোডাইনামিক ডিজাইন, ট্র্যাক-টিউনড সাসপেনশন এবং খাঁটি ট্র্যাক পারফরমেন্স খোঁজা উৎসাহীদের জন্য ইয়ামাহার প্রমাণিত আর-সিরিজ ডিএনএ সহ স্ট্রিট ব্যবহারযোগ্যতার সাথে রেসিং-অনুপ্রাণিত বৈশিষ্ট্যকে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা আর১৫এম এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ট্র্যাক-ফোকাসড ডিজাইন:</strong> <span class="highlight-value">রেসিং-অনুপ্রাণিত অ্যারোডাইনামিক্স এবং এরগোনমিক্স</span></li>
<li class="highlight-item"><strong class="highlight-label">ইউএসডি ফ্রন্ট সাসপেনশন:</strong> <span class="highlight-value">উন্নত হ্যান্ডলিংয়ের জন্য আপসাইড-ডাউন ফর্কস</span></li>
<li class="highlight-item"><strong class="highlight-label">কুইকশিফটার প্রযুক্তি:</strong> <span class="highlight-value">রেসিং পারফরমেন্সের জন্য ক্লাচলেস আপশিফট</span></li>
<li class="highlight-item"><strong class="highlight-label">ভিভিএ ইঞ্জিন:</strong> <span class="highlight-value">সর্বোত্তম শক্তির জন্য ভেরিয়েবল ভালভ অ্যাকচুয়েশন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা আর১৫এম এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">পেশাদার রেসিং বৈশিষ্ট্য:</strong> <span class="pro-description">কুইকশিফটার, ইউএসডি ফর্কস এবং ট্র্যাক-ভিত্তিক সাসপেনশন সেটআপ</span></li>
<li class="pro-item"><strong class="pro-label">ভিভিএ পারফরমেন্স:</strong> <span class="pro-description">ভেরিয়েবল ভালভ অ্যাকচুয়েশন রেভ রেঞ্জ জুড়ে পাওয়ার ডেলিভারি অপ্টিমাইজ করে</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল চ্যানেল এবিএস:</strong> <span class="pro-description">উভয় চাকায় এবিএস সহ উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম কম্পোনেন্ট এবং ইয়ামাহার প্রমাণিত ইঞ্জিনিয়ারিং</span></li>
<li class="pro-item"><strong class="pro-label">অ্যারোডাইনামিক ডিজাইন:</strong> <span class="pro-description">চমৎকার বায়ু সুরক্ষা সহ রেস-অনুপ্রাণিত বডিওয়ার্ক</span></li>
<li class="pro-item"><strong class="pro-label">ট্র্যাক পারফরমেন্স:</strong> <span class="pro-description">সার্কিট রাইডিং এবং ট্র্যাক ডে ইভেন্টের জন্য অপ্টিমাইজড</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা আর১৫এম এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ ক্রয় মূল্য:</strong> <span class="con-description">প্রিমিয়াম মূল্য গড় ক্রেতাদের জন্য অ্যাক্সেসিবিলিটি সীমিত করে</span></li>
<li class="con-item"><strong class="con-label">ট্র্যাক-ফোকাসড এরগোনমিক্স:</strong> <span class="con-description">আক্রমণাত্মক রাইডিং পজিশন দৈনিক ব্যবহারের জন্য অস্বস্তিকর হতে পারে</span></li>
<li class="con-item"><strong class="con-label">সীমিত দৈনিক ব্যবহারিকতা:</strong> <span class="con-description">ট্র্যাক অভিমুখিতা শহুরে রাইডিংয়ের জন্য আরাম ক্ষুণ্ন করে</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">প্রিমিয়াম কম্পোনেন্টের বিশেষায়িত সেবা এবং যন্ত্রাংশ প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">মাঝারি জ্বালানি সাশ্রয়:</strong> <span class="con-description">পারফরমেন্স ফোকাস কম জ্বালানি দক্ষতার ফলাফল</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা আর১৫এম কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Track day enthusiasts</li>
<li class="best-for-item">Racing and sport riding fans</li>
<li class="best-for-item">Experienced riders seeking performance</li>
<li class="best-for-item">Weekend track riders</li>
<li class="best-for-item">Those prioritizing handling and performance</li>
<li class="best-for-item">Yamaha R-series enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা আর১৫এম কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily city commuting</li>
<li class="not-recommended-item">Beginner or novice riders</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Those needing maximum comfort</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা আর১৫এম এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৬,২৫,০০০ - ট্র্যাক-ফোকাসড ১৫৫সিসি ফিচারের জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">উচ্চ - প্রিমিয়াম রক্ষণাবেক্ষণ সহ ৳১৩-১৭ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১০-১ৃ প্রতি কিমি (পারফরমেন্স টিউনিং সহ ৩৩-৩৮ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১১,০০০-১৮,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা আর১৫এম পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- চমৎকার ট্র্যাক পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- পারফরমেন্স-ফোকাসড টিউনিং</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.3</span> <span class="rating-note">- ট্র্যাক-ভিত্তিক অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- রেসিং-অনুপ্রাণিত ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ট্র্যাক বৈশিষ্ট্যের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- প্রিমিয়াম ইয়ামাহা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা আর১৫এম সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">আর১৫এম কে স্ট্যান্ডার্ড আর১৫ থেকে কী আলাদা করে?</h3>
<p class="faq-answer">আর১৫এম এ আপসাইড-ডাউন ফ্রন্ট ফর্কস, কুইকশিফটার, ট্র্যাক-টিউনড সাসপেনশন, রেসিং এরগোনমিক্স এবং বিশেষভাবে ট্র্যাক পারফরমেন্স এবং রেসিং অ্যাপ্লিকেশনের জন্য ডিজাইন করা বিশেষায়িত কম্পোনেন্ট রয়েছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ট্র্যাক রাইডিংয়ের জন্য কুইকশিফটার কি মূল্যবান?</h3>
<p class="faq-answer">হ্যাঁ, কুইকশিফটার ক্লাচ অপারেশন ছাড়াই নিরবচ্ছিন্ন আপশিফটের অনুমতি দেয়, মোমেন্টাম বজায় রাখে এবং ল্যাপ টাইম উন্নত করে। এটি বিশেষভাবে ট্র্যাক রাইডিং এবং রেসিং পরিস্থিতির জন্য উপকারী।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইউএসডি ফর্কস কীভাবে হ্যান্ডলিং উন্নত করে?</h3>
<p class="faq-answer">ইউএসডি (আপসাইড-ডাউন) ফর্কস প্রচলিত ফর্কসের তুলনায় আরও দৃঢ়তা এবং ভাল ফিডব্যাক প্রদান করে, হ্যান্ডলিং নির্ভুলতা, ব্রেকিংয়ের সময় স্থিতিশীলতা এবং সামগ্রিক কর্নারিং পারফরমেন্স উন্নত করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা আর১৫এম কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৬,২৫,০০০ টাকায় ইয়ামাহা আর১৫এম ১৫৫সিসি সেগমেন্টে কুইকশিফটার, ইউএসডি ফর্কস, ভিভিএ প্রযুক্তি এবং ডুয়াল-চ্যানেল এবিএস সহ পেশাদার বৈশিষ্ট্য সহ খাঁটি ট্র্যাক পারফরমেন্স প্রদান করে। এটি গুরুতর ট্র্যাক উৎসাহীদের জন্য ব্যতিক্রমী হ্যান্ডলিং, বিল্ড কোয়ালিটি এবং রেসিং ডিএনএ প্রদান করে। তবে, প্রিমিয়াম মূল্য, ট্র্যাক-ফোকাসড এরগোনমিক্স, সীমিত দৈনিক ব্যবহারিকতা, ব্যয়বহুল রক্ষণাবেক্ষণ এবং মাঝারি জ্বালানি অর্থনীতি এটিকে প্রধানত নিবেদিত ট্র্যাক রাইডার এবং পারফরমেন্স উৎসাহীদের জন্য উপযুক্ত করে তোলে যারা আরাম এবং দৈনন্দিন ব্যবহারযোগ্যতার চেয়ে রেসিং বৈশিষ্ট্যকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পেশাদার বৈশিষ্ট্য সহ খাঁটি ট্র্যাক পারফরমেন্সের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha R15M already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha R15M\n")

	return nil
}
