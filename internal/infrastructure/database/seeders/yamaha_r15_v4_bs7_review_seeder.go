package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaR15V4BS7Review handles seeding Yamaha R15 V4 BS7 product review and translation
type YamahaR15V4BS7Review struct {
	BaseSeeder
}

// NewYamahaR15V4BS7ReviewSeeder creates a new YamahaR15V4BS7Review
func NewYamahaR15V4BS7ReviewSeeder() *YamahaR15V4BS7Review {
	return &YamahaR15V4BS7Review{BaseSeeder: BaseSeeder{name: "Yamaha R15 V4 BS7 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaR15V4BS7Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R15 V4 BS7").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha R15 V4 BS7 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha R15 V4 BS7 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha R15 V4 BS7 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha R15 V4 BS7 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha R15 V4 BS7 is a premium 155cc fully-faired sports motorcycle featuring VVA technology, slipper clutch, dual-channel ABS, and BS7 emission compliance. Priced at ৳6,60,000, it combines track-inspired performance with street-friendly features including advanced electronics, aerodynamic design, and Yamaha's proven R-series engineering for enthusiasts seeking premium sports bike experience.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R15 V4 BS7 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">VVA Engine Technology:</strong> <span class="highlight-value">Variable Valve Actuation for optimized power delivery</span></li>
<li class="highlight-item"><strong class="highlight-label">Dual Channel ABS:</strong> <span class="highlight-value">Advanced braking safety for both wheels</span></li>
<li class="highlight-item"><strong class="highlight-label">Slipper Clutch:</strong> <span class="highlight-value">Prevents rear wheel lock-up during aggressive downshifts</span></li>
<li class="highlight-item"><strong class="highlight-label">BS7 Compliance:</strong> <span class="highlight-value">Latest emission standards with cleaner performance</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R15 V4 BS7 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Advanced VVA Technology:</strong> <span class="pro-description">Variable valve timing provides excellent power delivery across RPM range</span></li>
<li class="pro-item"><strong class="pro-label">Dual Channel ABS Safety:</strong> <span class="pro-description">Enhanced braking safety with ABS on both front and rear wheels</span></li>
<li class="pro-item"><strong class="pro-label">Slipper Clutch Benefits:</strong> <span class="pro-description">Smoother downshifts and enhanced cornering stability</span></li>
<li class="pro-item"><strong class="pro-label">Premium Build Quality:</strong> <span class="pro-description">High-quality components and excellent fit-finish</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Sharp, aerodynamic design with full fairing protection</span></li>
<li class="pro-item"><strong class="pro-label">Digital Instrumentation:</strong> <span class="pro-description">Advanced LCD display with multiple riding modes</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R15 V4 BS7 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Purchase Price:</strong> <span class="con-description">Significantly expensive compared to other 155cc motorcycles</span></li>
<li class="con-item"><strong class="con-label">Aggressive Riding Position:</strong> <span class="con-description">Sport-focused ergonomics may be uncomfortable for daily commuting</span></li>
<li class="con-item"><strong class="con-label">High Maintenance Cost:</strong> <span class="con-description">Premium features require specialized service and expensive parts</span></li>
<li class="con-item"><strong class="con-label">Limited Fuel Efficiency:</strong> <span class="con-description">Performance tuning results in moderate fuel economy</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha service centers and specialized technicians</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha R15 V4 BS7 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance enthusiasts seeking premium features</li>
<li class="best-for-item">Weekend sportbike riders</li>
<li class="best-for-item">Experienced riders wanting advanced technology</li>
<li class="best-for-item">Those prioritizing safety with dual-channel ABS</li>
<li class="best-for-item">Style-conscious riders preferring premium brands</li>
<li class="best-for-item">Track day participants seeking street-legal bike</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha R15 V4 BS7?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily city commuting in heavy traffic</li>
<li class="not-recommended-item">Beginner or novice riders</li>
<li class="not-recommended-item">Long-distance touring riders</li>
<li class="not-recommended-item">Those seeking maximum fuel efficiency</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R15 V4 BS7 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳6,60,000 - Premium for 155cc with VVA and dual ABS</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳14-18 per km with premium maintenance</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳11-14 per km (32-38 km/l with VVA efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳12,000-20,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R15 V4 BS7 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Excellent VVA performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- Moderate with performance focus</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.4</span> <span class="rating-note">- Sport-oriented position</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Premium sports design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Good for premium features</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Exceptional Yamaha quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha R15 V4 BS7 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does VVA technology improve performance?</h3>
<p class="faq-answer">VVA adjusts valve timing and lift based on engine speed, optimizing torque at low RPMs and power at high RPMs, providing better overall performance across the entire rev range.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What are the benefits of dual-channel ABS?</h3>
<p class="faq-answer">Dual-channel ABS prevents wheel lock-up on both front and rear wheels independently, providing better braking control, shorter stopping distances, and enhanced safety in emergency situations.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does the slipper clutch benefit riding?</h3>
<p class="faq-answer">The slipper clutch prevents rear wheel lock-up during aggressive downshifting, allowing smoother engine braking and better stability during cornering and track riding.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R15 V4 BS7 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳6,60,000, the Yamaha R15 V4 BS7 offers a compelling package of advanced VVA technology, dual-channel ABS, slipper clutch, and premium build quality in the 155cc sports segment. The VVA engine provides excellent performance across the rev range while the safety features enhance riding confidence. However, the high pricing, sport-focused ergonomics, expensive maintenance, limited fuel efficiency, and specialized service requirements make it suitable primarily for performance enthusiasts and experienced riders who prioritize advanced features and premium quality over comfort and value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For advanced VVA performance with dual ABS safety</span></p>
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
		return fmt.Errorf("error creating Yamaha R15 V4 BS7 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha R15 V4 BS7 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা আর১৫ ভি৪ বিএস৭ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা আর১৫ ভি৪ বিএস৭ হল একটি প্রিমিয়াম ১৫৫সিসি ফুল-ফেয়ার্ড স্পোর্টস মোটরসাইকেল যেখানে ভিভিএ প্রযুক্তি, স্লিপার ক্লাচ, ডুয়াল-চ্যানেল এবিএস এবং বিএস৭ নির্গমন সম্মতি রয়েছে। ৳৬,৬০,০০০ টাকায় মূল্যায়িত, এটি উন্নত ইলেকট্রনিক্স, অ্যারোডাইনামিক ডিজাইন এবং প্রিমিয়াম স্পোর্টস বাইক অভিজ্ঞতা খোঁজা উৎসাহীদের জন্য ইয়ামাহার প্রমাণিত আর-সিরিজ ইঞ্জিনিয়ারিং সহ স্ট্রিট-ফ্রেন্ডলি বৈশিষ্ট্যের সাথে ট্র্যাক-অনুপ্রাণিত পারফরমেন্সকে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা আর১৫ ভি৪ বিএস৭ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ভিভিএ ইঞ্জিন প্রযুক্তি:</strong> <span class="highlight-value">অনুকূলিত পাওয়ার ডেলিভারির জন্য ভেরিয়েবল ভালভ অ্যাকচুয়েশন</span></li>
<li class="highlight-item"><strong class="highlight-label">ডুয়াল চ্যানেল এবিএস:</strong> <span class="highlight-value">উভয় চাকার জন্য উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">স্লিপার ক্লাচ:</strong> <span class="highlight-value">আক্রমণাত্মক ডাউনশিফটের সময় পিছনের চাকা লক-আপ প্রতিরোধ করে</span></li>
<li class="highlight-item"><strong class="highlight-label">বিএস৭ সম্মতি:</strong> <span class="highlight-value">পরিষ্কার পারফরমেন্স সহ সর্বশেষ নির্গমন মান</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা আর১৫ ভি৪ বিএস৭ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">উন্নত ভিভিএ প্রযুক্তি:</strong> <span class="pro-description">ভেরিয়েবল ভালভ টাইমিং আরপিএম রেঞ্জ জুড়ে চমৎকার পাওয়ার ডেলিভারি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল চ্যানেল এবিএস নিরাপত্তা:</strong> <span class="pro-description">সামনে এবং পিছনে উভয় চাকায় এবিএস সহ উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="pro-item"><strong class="pro-label">স্লিপার ক্লাচ সুবিধা:</strong> <span class="pro-description">মসৃণ ডাউনশিফট এবং উন্নত কর্নারিং স্থিতিশীলতা</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বিল্ড কোয়ালিটি:</strong> <span class="pro-description">উচ্চ-মানের কম্পোনেন্ট এবং চমৎকার ফিট-ফিনিশ</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">সম্পূর্ণ ফেয়ারিং সুরক্ষা সহ তীক্ষ্ণ, অ্যারোডাইনামিক ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল ইনস্ট্রুমেন্টেশন:</strong> <span class="pro-description">একাধিক রাইডিং মোড সহ উন্নত এলসিডি ডিসপ্লে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা আর১৫ ভি৪ বিএস৭ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ ক্রয় মূল্য:</strong> <span class="con-description">অন্যান্য ১৫৫সিসি মোটরসাইকেলের তুলনায় উল্লেখযোগ্যভাবে ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">আক্রমণাত্মক রাইডিং পজিশন:</strong> <span class="con-description">স্পোর্ট-ফোকাসড এরগোনমিক্স দৈনিক যাতায়াতের জন্য অস্বস্তিকর হতে পারে</span></li>
<li class="con-item"><strong class="con-label">উচ্চ রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">প্রিমিয়াম ফিচারের বিশেষায়িত সেবা এবং ব্যয়বহুল যন্ত্রাংশ প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত জ্বালানি দক্ষতা:</strong> <span class="con-description">পারফরমেন্স টিউনিং মাঝারি জ্বালানি অর্থনীতির ফলাফল</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">কম ইয়ামাহা সেবা কেন্দ্র এবং বিশেষায়িত প্রযুক্তিবিদ</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা আর১৫ ভি৪ বিএস৭ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance enthusiasts seeking premium features</li>
<li class="best-for-item">Weekend sportbike riders</li>
<li class="best-for-item">Experienced riders wanting advanced technology</li>
<li class="best-for-item">Those prioritizing safety with dual-channel ABS</li>
<li class="best-for-item">Style-conscious riders preferring premium brands</li>
<li class="best-for-item">Track day participants seeking street-legal bike</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা আর১৫ ভি৪ বিএস৭ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily city commuting in heavy traffic</li>
<li class="not-recommended-item">Beginner or novice riders</li>
<li class="not-recommended-item">Long-distance touring riders</li>
<li class="not-recommended-item">Those seeking maximum fuel efficiency</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা আর১৫ ভি৪ বিএস৭ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৬,৬০,০০০ - ভিভিএ এবং ডুয়াল এবিএস সহ ১৫৫সিসির জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">উচ্চ - প্রিমিয়াম রক্ষণাবেক্ষণ সহ ৳১৪-১৮ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১১-১৪ প্রতি কিমি (ভিভিএ দক্ষতা সহ ৩২-৩৮ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১২,০০০-২০,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা আর১৫ ভি৪ বিএস৭ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- চমৎকার ভিভিএ পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- পারফরমেন্স ফোকাসের সাথে মাঝারি</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.4</span> <span class="rating-note">- স্পোর্ট-ভিত্তিক অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- প্রিমিয়াম স্পোর্টস ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- প্রিমিয়াম বৈশিষ্ট্যের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- ব্যতিক্রমী ইয়ামাহা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা আর১৫ ভি৪ বিএস৭ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ভিভিএ প্রযুক্তি কীভাবে পারফরমেন্স উন্নত করে?</h3>
<p class="faq-answer">ভিভিএ ইঞ্জিনের গতির উপর ভিত্তি করে ভালভ টাইমিং এবং লিফট সামঞ্জস্য করে, নিম্ন আরপিএমে টর্ক এবং উচ্চ আরপিএমে পাওয়ার অপ্টিমাইজ করে, সম্পূর্ণ রেভ রেঞ্জ জুড়ে ভাল সামগ্রিক পারফরমেন্স প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ডুয়াল-চ্যানেল এবিএসের সুবিধা কী?</h3>
<p class="faq-answer">ডুয়াল-চ্যানেল এবিএস সামনে এবং পিছনে উভয় চাকায় স্বতন্ত্রভাবে হুইল লক-আপ প্রতিরোধ করে, ভাল ব্রেকিং নিয়ন্ত্রণ, ছোট স্টপিং দূরত্ব এবং জরুরি পরিস্থিতিতে উন্নত নিরাপত্তা প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">স্লিপার ক্লাচ রাইডিংয়ে কীভাবে উপকার করে?</h3>
<p class="faq-answer">স্লিপার ক্লাচ আক্রমণাত্মক ডাউনশিফটিংয়ের সময় পিছনের চাকা লক-আপ প্রতিরোধ করে, মসৃণ ইঞ্জিন ব্রেকিং এবং কর্নারিং ও ট্র্যাক রাইডিংয়ের সময় ভাল স্থিতিশীলতার অনুমতি দেয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা আর১৫ ভি৪ বিএস৭ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৬,৬০,০০০ টাকায় ইয়ামাহা আর১৫ ভি৪ বিএস৭ ১৫৫সিসি স্পোর্টস সেগমেন্টে উন্নত ভিভিএ প্রযুক্তি, ডুয়াল-চ্যানেল এবিএস, স্লিপার ক্লাচ এবং প্রিমিয়াম বিল্ড কোয়ালিটির একটি আকর্ষণীয় প্যাকেজ প্রদান করে। ভিভিএ ইঞ্জিন রেভ রেঞ্জ জুড়ে চমৎকার পারফরমেন্স প্রদান করে যখন নিরাপত্তা বৈশিষ্ট্যগুলি রাইডিং আত্মবিশ্বাস বাড়ায়। তবে, উচ্চ মূল্য, স্পোর্ট-ফোকাসড এরগোনমিক্স, ব্যয়বহুল রক্ষণাবেক্ষণ, সীমিত জ্বালানি দক্ষতা এবং বিশেষায়িত সেবার প্রয়োজনীয়তা এটিকে প্রধানত পারফরমেন্স উৎসাহী এবং অভিজ্ঞ রাইডারদের জন্য উপযুক্ত করে তোলে যারা আরাম ও মূল্যের চেয়ে উন্নত বৈশিষ্ট্য এবং প্রিমিয়াম গুণমানকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ডুয়াল এবিএস নিরাপত্তা সহ উন্নত ভিভিএ পারফরমেন্সের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha R15 V4 BS7 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha R15 V4 BS7\n")

	return nil
}
