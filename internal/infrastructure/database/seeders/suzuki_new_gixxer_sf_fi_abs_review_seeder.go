package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiNewGixxerSFFiABSReview handles seeding Suzuki New Gixxer SF Fi ABS product review and translation
type SuzukiNewGixxerSFFiABSReview struct {
	BaseSeeder
}

// NewSuzukiNewGixxerSFFiABSReviewSeeder creates a new SuzukiNewGixxerSFFiABSReview
func NewSuzukiNewGixxerSFFiABSReviewSeeder() *SuzukiNewGixxerSFFiABSReview {
	return &SuzukiNewGixxerSFFiABSReview{BaseSeeder: BaseSeeder{name: "Suzuki New Gixxer SF Fi ABS Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiNewGixxerSFFiABSReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki New Gixxer SF Fi ABS").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki New Gixxer SF Fi ABS product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki New Gixxer SF Fi ABS product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki New Gixxer SF Fi ABS already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki New Gixxer SF Fi ABS Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki New Gixxer SF Fi ABS is a fully-faired 155cc sports motorcycle featuring fuel injection, single-channel ABS, and modern styling. Priced at ৳3,49,950, it combines the aerodynamic advantages of full fairing with Suzuki's reliable 155cc engine, digital instrumentation, and essential safety features in an affordable sports package.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki New Gixxer SF Fi ABS Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection System:</strong> <span class="highlight-value">155cc fuel-injected engine with reliable SOHC technology</span></li>
<li class="highlight-item"><strong class="highlight-label">Full Fairing Design:</strong> <span class="highlight-value">Excellent wind protection and aerodynamics</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS Safety Feature:</strong> <span class="highlight-value">Single-channel ABS for enhanced braking safety</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Instrumentation:</strong> <span class="highlight-value">Modern cluster with comprehensive information display</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki New Gixxer SF Fi ABS Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Fuel Injection Efficiency:</strong> <span class="pro-description">FI system ensures better fuel efficiency and consistent performance</span></li>
<li class="pro-item"><strong class="pro-label">ABS Safety Feature:</strong> <span class="pro-description">Single-channel ABS provides essential braking safety</span></li>
<li class="pro-item"><strong class="pro-label">Full Fairing Benefits:</strong> <span class="pro-description">Excellent wind protection and aerodynamic efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Build Quality:</strong> <span class="pro-description">Reliable engineering with good build standards</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">Impressive mileage of 45-50 km/l with fuel injection</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Touring Position:</strong> <span class="pro-description">Well-balanced ergonomics for longer rides</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki New Gixxer SF Fi ABS Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price Point:</strong> <span class="con-description">৳3,49,950 is expensive for 155cc motorcycle</span></li>
<li class="con-item"><strong class="con-label">Limited Power Output:</strong> <span class="con-description">155cc engine lacks punch for highway performance</span></li>
<li class="con-item"><strong class="con-label">Single Channel ABS Only:</strong> <span class="con-description">Not dual-channel ABS which is safer</span></li>
<li class="con-item"><strong class="con-label">Heavy for City Use:</strong> <span class="con-description">Full fairing makes it cumbersome in traffic</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">Premium pricing extends to service costs</span></li>
<li class="con-item"><strong class="con-label">Poor Value Proposition:</strong> <span class="con-description">Overpriced compared to competitors with more power</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki New Gixxer SF Fi ABS in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">155cc fully-faired bike enthusiasts</li>
<li class="best-for-item">Daily commuting with weekend touring</li>
<li class="best-for-item">Riders wanting ABS safety at lower displacement</li>
<li class="best-for-item">Fuel-efficient highway commuting</li>
<li class="best-for-item">Upgrading from 125cc to 155cc segment</li>
<li class="best-for-item">Those prioritizing wind protection</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki New Gixxer SF Fi ABS?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Power-hungry riders seeking performance</li>
<li class="not-recommended-item">Budget buyers seeking maximum value</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Riders wanting dual-channel ABS</li>
<li class="not-recommended-item">Those needing extensive service support</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki New Gixxer SF Fi ABS Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,49,950 - Premium for 155cc with ABS</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳9-11 per km with good fuel efficiency</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳9-11 per km (45-50 km/l at current fuel prices)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳6,000-8,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki New Gixxer SF Fi ABS Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Adequate for displacement</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good touring comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Modern sports design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.4</span> <span class="rating-note">- Overpriced for 155cc</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good Suzuki standards</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki New Gixxer SF Fi ABS Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does the fuel injection system benefit the Gixxer SF Fi?</h3>
<p class="faq-answer">Fuel injection provides better fuel efficiency (45-50 km/l), consistent performance across different conditions, reduced emissions, and smoother engine operation compared to carburetor systems.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is single-channel ABS sufficient for safety?</h3>
<p class="faq-answer">Single-channel ABS on the front wheel provides good safety enhancement, but dual-channel ABS covering both wheels would offer better protection, especially in emergency braking scenarios.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How is the highway performance with 155cc engine?</h3>
<p class="faq-answer">Adequate for highway cruising at 80-90 km/h with good fuel efficiency. However, acceleration and top-end power may feel limited compared to larger displacement motorcycles.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki New Gixxer SF Fi ABS in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,49,950, the Suzuki New Gixxer SF Fi ABS offers a well-rounded fully-faired 155cc package with fuel injection efficiency, ABS safety, full fairing protection, and Suzuki reliability. The FI system ensures good fuel economy of 45-50 km/l while the full fairing provides excellent wind protection for touring. However, the high pricing for 155cc displacement, limited power output, single-channel ABS instead of dual-channel, expensive maintenance, and poor value proposition compared to more powerful alternatives make it suitable only for riders who prioritize fuel efficiency, ABS safety, and full-fairing protection over outright performance and value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For fuel-efficient fully-faired commuting with ABS safety</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.1,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Suzuki New Gixxer SF Fi ABS review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki New Gixxer SF Fi ABS (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস হল একটি ফুল-ফেয়ার্ড ১৫৫সিসি স্পোর্টস মোটরসাইকেল যেখানে ফুয়েল ইনজেকশন, সিঙ্গেল-চ্যানেল এবিএস এবং আধুনিক স্টাইলিং রয়েছে। ৳৩,৪৯,৯৫০ টাকায় মূল্যায়িত, এটি ফুল ফেয়ারিংয়ের অ্যারোডাইনামিক সুবিধাগুলি সুজুকির নির্ভরযোগ্য ১৫৫সিসি ইঞ্জিন, ডিজিটাল ইনস্ট্রুমেন্টেশন এবং একটি সাশ্রয়ী স্পোর্টস প্যাকেজে অত্যাবশ্যকীয় নিরাপত্তা বৈশিষ্ট্যগুলির সাথে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন সিস্টেম:</strong> <span class="highlight-value">নির্ভরযোগ্য এসওএইচসি প্রযুক্তি সহ ১৫৫সিসি ফুয়েল-ইনজেক্টেড ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ ফেয়ারিং ডিজাইন:</strong> <span class="highlight-value">চমৎকার বাতাস সুরক্ষা এবং অ্যারোডাইনামিক্স</span></li>
<li class="highlight-item"><strong class="highlight-label">এবিএস নিরাপত্তা বৈশিষ্ট্য:</strong> <span class="highlight-value">উন্নত ব্রেকিং নিরাপত্তার জন্য সিঙ্গেল-চ্যানেল এবিএস</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল ইনস্ট্রুমেন্টেশন:</strong> <span class="highlight-value">ব্যাপক তথ্য প্রদর্শন সহ আধুনিক ক্লাস্টার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন দক্ষতা:</strong> <span class="pro-description">এফআই সিস্টেম ভাল জ্বালানি দক্ষতা এবং সামঞ্জস্যপূর্ণ পারফরমেন্স নিশ্চিত করে</span></li>
<li class="pro-item"><strong class="pro-label">এবিএস নিরাপত্তা বৈশিষ্ট্য:</strong> <span class="pro-description">সিঙ্গেল-চ্যানেল এবিএস অত্যাবশ্যক ব্রেকিং নিরাপত্তা প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">সম্পূর্ণ ফেয়ারিং সুবিধা:</strong> <span class="pro-description">চমৎকার বাতাস সুরক্ষা এবং অ্যারোডাইনামিক দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি বিল্ড কোয়ালিটি:</strong> <span class="pro-description">ভালো বিল্ড মান সহ নির্ভরযোগ্য ইঞ্জিনিয়ারিং</span></li>
<li class="pro-item"><strong class="pro-label">ভালো জ্বালানি সাশ্রয়:</strong> <span class="pro-description">ফুয়েল ইনজেকশন সহ ৪৫-৫০ কিমি/লিটার চিত্তাকর্ষক মাইলেজ</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক ট্যুরিং অবস্থান:</strong> <span class="pro-description">দীর্ঘ রাইডের জন্য সুসংগত এরগোনমিক্স</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য পয়েন্ট:</strong> <span class="con-description">১৫৫সিসি মোটরসাইকেলের জন্য ৳৩,৪৯,৯৫০ ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সীমিত শক্তি আউটপুট:</strong> <span class="con-description">হাইওয়ে পারফরমেন্সের জন্য ১৫৫সিসি ইঞ্জিনের গুণ অভাব</span></li>
<li class="con-item"><strong class="con-label">কেবল সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="con-description">ডুয়াল-চ্যানেল এবিএস নয় যা আরো নিরাপদ</span></li>
<li class="con-item"><strong class="con-label">শহুরে ব্যবহারের জন্য ভারী:</strong> <span class="con-description">সম্পূর্ণ ফেয়ারিং ট্রাফিকে এটিকে কষ্টকর করে তোলে</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">প্রিমিয়াম মূল্য সেবা খরচ পর্যন্ত বিস্তৃত</span></li>
<li class="con-item"><strong class="con-label">খারাপ মূল্য প্রস্তাব:</strong> <span class="con-description">আরো শক্তিশালী প্রতিযোগীদের তুলনায় অতিরিক্ত দামি</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">155cc fully-faired bike enthusiasts</li>
<li class="best-for-item">Daily commuting with weekend touring</li>
<li class="best-for-item">Riders wanting ABS safety at lower displacement</li>
<li class="best-for-item">Fuel-efficient highway commuting</li>
<li class="best-for-item">Upgrading from 125cc to 155cc segment</li>
<li class="best-for-item">Those prioritizing wind protection</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Power-hungry riders seeking performance</li>
<li class="not-recommended-item">Budget buyers seeking maximum value</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Riders wanting dual-channel ABS</li>
<li class="not-recommended-item">Those needing extensive service support</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৩,৪৯,৯৫০ - এবিএস সহ ১৫৫সিসির জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মধ্যম - ভালো জ্বালানি দক্ষতা সহ ৳৯-১১ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৯-১১ প্রতি কিমি (বর্তমান জ্বালানি দামে ৪৫-৫০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৬,০০০-৮,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- ডিসপ্লেসমেন্টের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভালো ট্যুরিং আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- আধুনিক স্পোর্টস ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.4</span> <span class="rating-note">- ১৫৫সিসির জন্য বেশি দামি</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ভালো সুজুকি মান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ফুয়েল ইনজেকশন সিস্টেম গিক্সার এসএফ এফআই এর কীভাবে উপকার করে?</h3>
<p class="faq-answer">ফুয়েল ইনজেকশন কার্বুরেটর সিস্টেমের তুলনায় ভাল জ্বালানি দক্ষতা (৪৫-৫০ কিমি/লিটার), বিভিন্ন অবস্থায় সামঞ্জস্যপূর্ণ পারফরমেন্স, হ্রাসকৃত নির্গমন এবং মসৃণ ইঞ্জিন অপারেশন প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">নিরাপত্তার জন্য সিঙ্গেল-চ্যানেল এবিএস কি যথেষ্ট?</h3>
<p class="faq-answer">সামনের চাকায় সিঙ্গেল-চ্যানেল এবিএস ভাল নিরাপত্তা উন্নতি প্রদান করে, কিন্তু উভয় চাকা কভার করে ডুয়াল-চ্যানেল এবিএস বিশেষত জরুরি ব্রেকিং পরিস্থিতিতে ভাল সুরক্ষা প্রদান করবে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">১৫৫সিসি ইঞ্জিনের সাথে হাইওয়ে পারফরমেন্স কেমন?</h3>
<p class="faq-answer">ভাল জ্বালানি দক্ষতা সহ ৮০-৯০ কিমি/ঘণ্টায় হাইওয়ে ক্রুজিংয়ের জন্য পর্যাপ্ত। তবে, বড় ডিসপ্লেসমেন্ট মোটরসাইকেলের তুলনায় ত্বরণ এবং টপ-এন্ড পাওয়ার সীমিত মনে হতে পারে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.1/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৩,৪৯,৯৫০ টাকায় সুজুকি নিউ গিক্সার এসএফ এফআই এবিএস ফুয়েল ইনজেকশন দক্ষতা, এবিএস নিরাপত্তা, ফুল ফেয়ারিং সুরক্ষা এবং সুজুকি নির্ভরযোগ্যতা সহ একটি সুসংগত ফুল-ফেয়ার্ড ১৫৫সিসি প্যাকেজ প্রদান করে। এফআই সিস্টেম ৪৫-৫০ কিমি/লিটার ভাল জ্বালানি অর্থনীতি নিশ্চিত করে যখন ফুল ফেয়ারিং ট্যুরিংয়ের জন্য চমৎকার বাতাস সুরক্ষা প্রদান করে। তবে, ১৫৫সিসি ডিসপ্লেসমেন্টের জন্য উচ্চ মূল্য, সীমিত পাওয়ার আউটপুট, ডুয়াল-চ্যানেলের পরিবর্তে সিঙ্গেল-চ্যানেল এবিএস, ব্যয়বহুল রক্ষণাবেক্ষণ এবং আরো শক্তিশালী বিকল্পের তুলনায় খারাপ মূল্য প্রস্তাব এটিকে কেবলমাত্র সেই রাইডারদের জন্য উপযুক্ত করে তোলে যারা সরাসরি পারফরমেন্স এবং মূল্যের চেয়ে জ্বালানি দক্ষতা, এবিএস নিরাপত্তা এবং ফুল-ফেয়ারিং সুরক্ষাকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- এবিএস নিরাপত্তা সহ জ্বালানি-দক্ষ ফুল-ফেয়ার্ড যাতায়াতের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki New Gixxer SF Fi ABS already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki New Gixxer SF Fi ABS\n")

	return nil
}
