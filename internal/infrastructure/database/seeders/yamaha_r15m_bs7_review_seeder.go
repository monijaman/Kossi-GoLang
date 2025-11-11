package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaR15MBS7Review handles seeding Yamaha R15M BS7 product review and translation
type YamahaR15MBS7Review struct {
	BaseSeeder
}

// NewYamahaR15MBS7ReviewSeeder creates a new YamahaR15MBS7Review
func NewYamahaR15MBS7ReviewSeeder() *YamahaR15MBS7Review {
	return &YamahaR15MBS7Review{BaseSeeder: BaseSeeder{name: "Yamaha R15M BS7 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaR15MBS7Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R15M BS7").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha R15M BS7 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha R15M BS7 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha R15M BS7 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha R15M BS7 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha R15M BS7 is the flagship 155cc track-focused sports motorcycle featuring VVA technology, upside-down forks, quickshifter, and BS7 emission compliance. Priced at ৳6,78,000, it offers professional racing features including track-tuned suspension, aerodynamic design, advanced electronics, and Yamaha's renowned R-series DNA for serious track enthusiasts and performance riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R15M BS7 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">VVA Technology:</strong> <span class="highlight-value">Variable Valve Actuation for optimized performance</span></li>
<li class="highlight-item"><strong class="highlight-label">USD Forks:</strong> <span class="highlight-value">Upside-down front suspension for track performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Quickshifter:</strong> <span class="highlight-value">Clutchless upshifts for seamless acceleration</span></li>
<li class="highlight-item"><strong class="highlight-label">Track-Focused Design:</strong> <span class="highlight-value">Aerodynamic bodywork and racing ergonomics</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R15M BS7 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Professional Racing Features:</strong> <span class="pro-description">Quickshifter, USD forks, and track-tuned suspension setup</span></li>
<li class="pro-item"><strong class="pro-label">VVA Engine Technology:</strong> <span class="pro-description">Variable valve timing for optimal power delivery across RPM range</span></li>
<li class="pro-item"><strong class="pro-label">Exceptional Build Quality:</strong> <span class="pro-description">Premium components and Yamaha's renowned engineering standards</span></li>
<li class="pro-item"><strong class="pro-label">Track Performance:</strong> <span class="pro-description">Designed specifically for track riding and racing applications</span></li>
<li class="pro-item"><strong class="pro-label">Advanced Electronics:</strong> <span class="pro-description">Digital instrumentation with lap timer and performance metrics</span></li>
<li class="pro-item"><strong class="pro-label">BS7 Emission Standards:</strong> <span class="pro-description">Latest emission compliance with environmental friendliness</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R15M BS7 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Pricing:</strong> <span class="con-description">Significantly expensive compared to standard 155cc motorcycles</span></li>
<li class="con-item"><strong class="con-label">Track-Focused Ergonomics:</strong> <span class="con-description">Aggressive riding position may be uncomfortable for daily use</span></li>
<li class="con-item"><strong class="con-label">High Maintenance Cost:</strong> <span class="con-description">Premium components require specialized service and expensive parts</span></li>
<li class="con-item"><strong class="con-label">Limited Daily Practicality:</strong> <span class="con-description">Track-oriented features not ideal for commuting or touring</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha dealerships and specialized technicians available</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha R15M BS7 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Serious track riders and racers</li>
<li class="best-for-item">Performance enthusiasts seeking premium features</li>
<li class="best-for-item">Experienced riders wanting professional setup</li>
<li class="best-for-item">Those prioritizing track performance over comfort</li>
<li class="best-for-item">Motorcycle collectors and enthusiasts</li>
<li class="best-for-item">Weekend track day participants</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha R15M BS7?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily commuting in city traffic</li>
<li class="not-recommended-item">Beginner or novice riders</li>
<li class="not-recommended-item">Touring and long-distance riding</li>
<li class="not-recommended-item">Those needing practical everyday motorcycle</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R15M BS7 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳6,78,000 - Premium for 155cc with pro racing features</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳15-20 per km with premium maintenance</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳12-15 per km (30-35 km/l with performance tuning)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 4 months for track use</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳15,000-25,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R15M BS7 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Exceptional track performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- Performance-focused tuning</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.2</span> <span class="rating-note">- Track-oriented position</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Premium racing aesthetics</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Premium for features offered</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Exceptional build quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha R15M BS7 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What makes the R15M BS7 different from regular R15?</h3>
<p class="faq-answer">The R15M BS7 features upside-down front forks, quickshifter, track-tuned suspension, racing ergonomics, and BS7 emission compliance. These professional racing features distinguish it from the standard R15 variants.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the quickshifter beneficial for track riding?</h3>
<p class="faq-answer">Yes, the quickshifter allows clutchless upshifts without losing momentum, crucial for lap times and racing performance. It provides seamless acceleration and reduces fatigue during extended track sessions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does VVA technology work?</h3>
<p class="faq-answer">VVA (Variable Valve Actuation) adjusts valve timing and lift based on RPM and throttle position, optimizing power delivery throughout the rev range for both low-end torque and high-rev power.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R15M BS7 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳6,78,000, the Yamaha R15M BS7 represents the pinnacle of 155cc track performance with professional racing features including quickshifter, USD forks, VVA technology, and BS7 compliance. It offers exceptional build quality, track-focused engineering, and Yamaha's racing DNA for serious enthusiasts. However, the premium pricing, track-oriented ergonomics, high maintenance costs, limited daily practicality, and specialized service requirements make it suitable only for dedicated track riders, performance enthusiasts, and those who prioritize professional racing features over comfort and value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For professional track performance with racing features</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.7,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha R15M BS7 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha R15M BS7 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা আর১৫এম বিএস৭ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা আর১৫এম বিএস৭ হল ফ্ল্যাগশিপ ১৫৫সিসি ট্র্যাক-ফোকাসড স্পোর্টস মোটরসাইকেল যেখানে ভিভিএ প্রযুক্তি, আপসাইড-ডাউন ফর্কস, কুইকশিফটার এবং বিএস৭ নির্গমন সম্মতি রয়েছে। ৳৬,৭৮,০০০ টাকায় মূল্যায়িত, এটি ট্র্যাক-টিউনড সাসপেনশন, অ্যারোডাইনামিক ডিজাইন, উন্নত ইলেকট্রনিক্স এবং গুরুতর ট্র্যাক উৎসাহী এবং পারফরমেন্স রাইডারদের জন্য ইয়ামাহার বিখ্যাত আর-সিরিজ ডিএনএ সহ পেশাদার রেসিং বৈশিষ্ট্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা আর১৫এম বিএস৭ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ভিভিএ প্রযুক্তি:</strong> <span class="highlight-value">অনুকূলিত পারফরমেন্সের জন্য ভেরিয়েবল ভালভ অ্যাকচুয়েশন</span></li>
<li class="highlight-item"><strong class="highlight-label">ইউএসডি ফর্কস:</strong> <span class="highlight-value">ট্র্যাক পারফরমেন্সের জন্য আপসাইড-ডাউন সামনের সাসপেনশন</span></li>
<li class="highlight-item"><strong class="highlight-label">কুইকশিফটার:</strong> <span class="highlight-value">নিরবচ্ছিন্ন ত্বরণের জন্য ক্লাচলেস আপশিফট</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্র্যাক-ফোকাসড ডিজাইন:</strong> <span class="highlight-value">অ্যারোডাইনামিক বডিওয়ার্ক এবং রেসিং এরগোনমিক্স</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা আর১৫এম বিএস৭ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">পেশাদার রেসিং বৈশিষ্ট্য:</strong> <span class="pro-description">কুইকশিফটার, ইউএসডি ফর্কস এবং ট্র্যাক-টিউনড সাসপেনশন সেটআপ</span></li>
<li class="pro-item"><strong class="pro-label">ভিভিএ ইঞ্জিন প্রযুক্তি:</strong> <span class="pro-description">আরপিএম রেঞ্জ জুড়ে সর্বোত্তম পাওয়ার ডেলিভারির জন্য ভেরিয়েবল ভালভ টাইমিং</span></li>
<li class="pro-item"><strong class="pro-label">ব্যতিক্রমী বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম কম্পোনেন্ট এবং ইয়ামাহার বিখ্যাত ইঞ্জিনিয়ারিং মান</span></li>
<li class="pro-item"><strong class="pro-label">ট্র্যাক পারফরমেন্স:</strong> <span class="pro-description">বিশেষভাবে ট্র্যাক রাইডিং এবং রেসিং অ্যাপ্লিকেশনের জন্য ডিজাইন করা</span></li>
<li class="pro-item"><strong class="pro-label">উন্নত ইলেকট্রনিক্স:</strong> <span class="pro-description">ল্যাপ টাইমার এবং পারফরমেন্স মেট্রিক্স সহ ডিজিটাল ইনস্ট্রুমেন্টেশন</span></li>
<li class="pro-item"><strong class="pro-label">বিএস৭ নির্গমন মান:</strong> <span class="pro-description">পরিবেশ বন্ধুত্ব সহ সর্বশেষ নির্গমন সম্মতি</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা আর১৫এম বিএস৭ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">স্ট্যান্ডার্ড ১৫৫সিসি মোটরসাইকেলের তুলনায় উল্লেখযোগ্যভাবে ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">ট্র্যাক-ফোকাসড এরগোনমিক্স:</strong> <span class="con-description">আক্রমণাত্মক রাইডিং পজিশন দৈনিক ব্যবহারের জন্য অস্বস্তিকর হতে পারে</span></li>
<li class="con-item"><strong class="con-label">উচ্চ রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">প্রিমিয়াম কম্পোনেন্টের বিশেষায়িত সেবা এবং ব্যয়বহুল যন্ত্রাংশ প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত দৈনিক ব্যবহারিকতা:</strong> <span class="con-description">ট্র্যাক-ভিত্তিক বৈশিষ্ট্যগুলি যাতায়াত বা ভ্রমণের জন্য আদর্শ নয়</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">কম ইয়ামাহা ডিলারশিপ এবং বিশেষায়িত প্রযুক্তিবিদ উপলব্ধ</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা আর১৫এম বিএস৭ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Serious track riders and racers</li>
<li class="best-for-item">Performance enthusiasts seeking premium features</li>
<li class="best-for-item">Experienced riders wanting professional setup</li>
<li class="best-for-item">Those prioritizing track performance over comfort</li>
<li class="best-for-item">Motorcycle collectors and enthusiasts</li>
<li class="best-for-item">Weekend track day participants</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা আর১৫এম বিএস৭ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily commuting in city traffic</li>
<li class="not-recommended-item">Beginner or novice riders</li>
<li class="not-recommended-item">Touring and long-distance riding</li>
<li class="not-recommended-item">Those needing practical everyday motorcycle</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা আর১৫এম বিএস৭ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৬,৭৮,০০০ - প্রো রেসিং ফিচার সহ ১৫৫সিসির জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">উচ্চ - প্রিমিয়াম রক্ষণাবেক্ষণ সহ ৳১৫-২০ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১২-১৫ প্রতি কিমি (পারফরমেন্স টিউনিং সহ ৩০-৩৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">ট্র্যাক ব্যবহারের জন্য প্রতি ৩,০০০ কিমি বা ৪ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১৫,০০০-২৫,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা আর১৫এম বিএস৭ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- ব্যতিক্রমী ট্র্যাক পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- পারফরমেন্স-ফোকাসড টিউনিং</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.2</span> <span class="rating-note">- ট্র্যাক-ভিত্তিক অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- প্রিমিয়াম রেসিং নান্দনিকতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- প্রদত্ত বৈশিষ্ট্যের জন্য প্রিমিয়াম</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- ব্যতিক্রমী বিল্ড কোয়ালিটি</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা আর১৫এম বিএস৭ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">আর১৫এম বিএস৭ কে নিয়মিত আর১৫ থেকে কী আলাদা করে?</h3>
<p class="faq-answer">আর১৫এম বিএস৭ এ আপসাইড-ডাউন ফ্রন্ট ফর্কস, কুইকশিফটার, ট্র্যাক-টিউনড সাসপেনশন, রেসিং এরগোনমিক্স এবং বিএস৭ নির্গমন সম্মতি রয়েছে। এই পেশাদার রেসিং বৈশিষ্ট্যগুলি এটিকে স্ট্যান্ডার্ড আর১৫ ভেরিয়েন্ট থেকে আলাদা করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ট্র্যাক রাইডিংয়ের জন্য কুইকশিফটার কি উপকারী?</h3>
<p class="faq-answer">হ্যাঁ, কুইকশিফটার মোমেন্টাম হারানো ছাড়াই ক্লাচলেস আপশিফটের অনুমতি দেয়, যা ল্যাপ টাইম এবং রেসিং পারফরমেন্সের জন্য গুরুত্বপূর্ণ। এটি নিরবচ্ছিন্ন ত্বরণ প্রদান করে এবং বর্ধিত ট্র্যাক সেশনের সময় ক্লান্তি হ্রাস করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ভিভিএ প্রযুক্তি কীভাবে কাজ করে?</h3>
<p class="faq-answer">ভিভিএ (ভেরিয়েবল ভালভ অ্যাকচুয়েশন) আরপিএম এবং থ্রোটল পজিশনের উপর ভিত্তি করে ভালভ টাইমিং এবং লিফট সামঞ্জস্য করে, লো-এন্ড টর্ক এবং হাই-রেভ পাওয়ার উভয়ের জন্য রেভ রেঞ্জ জুড়ে পাওয়ার ডেলিভারি অপ্টিমাইজ করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা আর১৫এম বিএস৭ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৬,৭৮,০০০ টাকায় ইয়ামাহা আর১৫এম বিএস৭ কুইকশিফটার, ইউএসডি ফর্কস, ভিভিএ প্রযুক্তি এবং বিএস৭ সম্মতি সহ পেশাদার রেসিং বৈশিষ্ট্য সহ ১৫৫সিসি ট্র্যাক পারফরমেন্সের শীর্ষস্থানীয় প্রতিনিধিত্ব করে। এটি গুরুতর উৎসাহীদের জন্য ব্যতিক্রমী বিল্ড কোয়ালিটি, ট্র্যাক-ফোকাসড ইঞ্জিনিয়ারিং এবং ইয়ামাহার রেসিং ডিএনএ প্রদান করে। তবে, প্রিমিয়াম মূল্য, ট্র্যাক-ভিত্তিক এরগোনমিক্স, উচ্চ রক্ষণাবেক্ষণ খরচ, সীমিত দৈনিক ব্যবহারিকতা এবং বিশেষায়িত সেবার প্রয়োজনীয়তা এটিকে কেবলমাত্র নিবেদিত ট্র্যাক রাইডার, পারফরমেন্স উৎসাহী এবং যারা আরাম ও মূল্যের চেয়ে পেশাদার রেসিং বৈশিষ্ট্যকে অগ্রাধিকার দেয় তাদের জন্য উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- রেসিং ফিচার সহ পেশাদার ট্র্যাক পারফরমেন্সের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha R15M BS7 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha R15M BS7\n")

	return nil
}
