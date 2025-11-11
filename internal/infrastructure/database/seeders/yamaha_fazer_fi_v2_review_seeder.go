package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaFazerFiV2Review handles seeding Yamaha Fazer FI V2 product review and translation
type YamahaFazerFiV2Review struct {
	BaseSeeder
}

// NewYamahaFazerFiV2ReviewSeeder creates a new YamahaFazerFiV2Review
func NewYamahaFazerFiV2ReviewSeeder() *YamahaFazerFiV2Review {
	return &YamahaFazerFiV2Review{BaseSeeder: BaseSeeder{name: "Yamaha Fazer FI V2 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaFazerFiV2Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha Fazer FI V2").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha Fazer FI V2 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha Fazer FI V2 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha Fazer FI V2 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha Fazer FI V2 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha Fazer FI V2 is an updated 150cc touring-focused motorcycle featuring fuel injection technology, semi-fairing design, and comfortable ergonomics. Priced at ৳2,70,000, it combines practical touring capabilities with modern FI efficiency, wind protection, reliable build quality, and Yamaha's proven engineering for riders seeking comfortable long-distance commuting and weekend touring.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha Fazer FI V2 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection System:</strong> <span class="highlight-value">Modern FI technology for better efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Semi-Fairing Design:</strong> <span class="highlight-value">Wind protection for comfortable touring</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Ergonomics:</strong> <span class="highlight-value">Upright riding position for long rides</span></li>
<li class="highlight-item"><strong class="highlight-label">Updated Features:</strong> <span class="highlight-value">V2 improvements in electronics and styling</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha Fazer FI V2 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Fuel Injection Benefits:</strong> <span class="pro-description">Better fuel economy, instant starts, and consistent performance</span></li>
<li class="pro-item"><strong class="pro-label">Touring Comfort:</strong> <span class="pro-description">Semi-fairing provides good wind protection for long rides</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Riding Position:</strong> <span class="pro-description">Upright ergonomics reduce fatigue on long journeys</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Yamaha's proven reliability and component quality</span></li>
<li class="pro-item"><strong class="pro-label">Practical Daily Use:</strong> <span class="pro-description">Balanced for both commuting and weekend touring</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">Achieves 45-50 km/l with FI technology</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha Fazer FI V2 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Dated Design:</strong> <span class="con-description">Styling may feel outdated compared to newer models</span></li>
<li class="con-item"><strong class="con-label">Moderate Performance:</strong> <span class="con-description">150cc engine provides adequate but not exciting power</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">Semi-fairing adds weight, affecting city maneuverability</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha dealerships compared to local brands</span></li>
<li class="con-item"><strong class="con-label">No ABS Available:</strong> <span class="con-description">Missing modern safety features like anti-lock braking</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha Fazer FI V2 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Long-distance commuters</li>
<li class="best-for-item">Weekend touring enthusiasts</li>
<li class="best-for-item">Riders prioritizing comfort over performance</li>
<li class="best-for-item">Those needing wind protection</li>
<li class="best-for-item">Daily highway riders</li>
<li class="best-for-item">Mature riders seeking practical motorcycle</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha Fazer FI V2?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Style-conscious younger riders</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Those seeking modern aggressive styling</li>
<li class="not-recommended-item">Riders wanting ABS safety features</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha Fazer FI V2 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,70,000 - Moderate for 150cc FI with semi-fairing</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low to Moderate - ৳7-9 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳6-8 per km (45-50 km/l with FI efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳5,000-8,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha Fazer FI V2 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Adequate for touring</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent touring comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- Dated design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good for touring needs</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Reliable Yamaha build</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha Fazer FI V2 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How effective is the semi-fairing for wind protection?</h3>
<p class="faq-answer">The semi-fairing provides good wind protection for the upper body and reduces fatigue on highway rides, though not as comprehensive as full fairing motorcycles.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for daily city commuting?</h3>
<p class="faq-answer">Yes, but the semi-fairing adds weight that may make it less agile in heavy traffic. It's better suited for highway commuting and touring rather than congested city riding.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What are the V2 improvements?</h3>
<p class="faq-answer">The V2 version features updated electronics, improved fuel injection mapping, refined styling elements, and better build quality compared to the original Fazer FI.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha Fazer FI V2 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,70,000, the Yamaha Fazer FI V2 offers a practical touring-focused package with fuel injection efficiency, semi-fairing wind protection, and comfortable ergonomics. It excels at long-distance commuting and weekend touring with excellent fuel economy of 45-50 km/l. However, the dated styling, moderate performance, heavy weight for city use, lack of ABS, and limited service network make it suitable primarily for mature riders and touring enthusiasts who prioritize comfort and fuel efficiency over modern styling and sporty performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For comfortable touring with FI efficiency and wind protection</span></p>
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
		return fmt.Errorf("error creating Yamaha Fazer FI V2 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha Fazer FI V2 (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা ফেজার এফআই ভি২ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা ফেজার এফআই ভি২ হল একটি আপডেটেড ১৫০সিসি ট্যুরিং-ফোকাসড মোটরসাইকেল যেখানে ফুয়েল ইনজেকশন প্রযুক্তি, সেমি-ফেয়ারিং ডিজাইন এবং আরামদায়ক এরগোনমিক্স রয়েছে। ৳২,৭০,০০০ টাকায় মূল্যায়িত, এটি আধুনিক এফআই দক্ষতা, বায়ু সুরক্ষা, নির্ভরযোগ্য বিল্ড কোয়ালিটি এবং আরামদায়ক দীর্ঘ-দূরত্ব যাতায়াত ও সাপ্তাহিক ভ্রমণ খোঁজা রাইডারদের জন্য ইয়ামাহার প্রমাণিত ইঞ্জিনিয়ারিং সহ ব্যবহারিক ট্যুরিং ক্ষমতাকে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা ফেজার এফআই ভি২ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন সিস্টেম:</strong> <span class="highlight-value">ভাল দক্ষতার জন্য আধুনিক এফআই প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">সেমি-ফেয়ারিং ডিজাইন:</strong> <span class="highlight-value">আরামদায়ক ট্যুরিংয়ের জন্য বায়ু সুরক্ষা</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক এরগোনমিক্স:</strong> <span class="highlight-value">দীর্ঘ রাইডের জন্য সোজা রাইডিং পজিশন</span></li>
<li class="highlight-item"><strong class="highlight-label">আপডেটেড ফিচার:</strong> <span class="highlight-value">ইলেকট্রনিক্স এবং স্টাইলিংয়ে ভি২ উন্নতি</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা ফেজার এফআই ভি২ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন সুবিধা:</strong> <span class="pro-description">ভাল জ্বালানি অর্থনীতি, তাৎক্ষণিক স্টার্ট এবং সামঞ্জস্যপূর্ণ পারফরমেন্স</span></li>
<li class="pro-item"><strong class="pro-label">ট্যুরিং আরাম:</strong> <span class="pro-description">সেমি-ফেয়ারিং দীর্ঘ রাইডের জন্য ভাল বায়ু সুরক্ষা প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইডিং পজিশন:</strong> <span class="pro-description">সোজা এরগোনমিক্স দীর্ঘ যাত্রায় ক্লান্তি হ্রাস করে</span></li>
<li class="pro-item"><strong class="pro-label">ভাল বিল্ড কোয়ালিটি:</strong> <span class="pro-description">ইয়ামাহার প্রমাণিত নির্ভরযোগ্যতা এবং কম্পোনেন্ট গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক দৈনিক ব্যবহার:</strong> <span class="pro-description">যাতায়াত এবং সাপ্তাহিক ভ্রমণ উভয়ের জন্য ভারসাম্যপূর্ণ</span></li>
<li class="pro-item"><strong class="pro-label">ভাল জ্বালানি সাশ্রয়:</strong> <span class="pro-description">এফআই প্রযুক্তি দিয়ে ৪৫-৫০ কিমি/লিটার অর্জন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা ফেজার এফআই ভি২ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">পুরাতন ডিজাইন:</strong> <span class="con-description">নতুন মডেলের তুলনায় স্টাইলিং পুরাতন মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">মাঝারি পারফরমেন্স:</strong> <span class="con-description">১৫০সিসি ইঞ্জিন পর্যাপ্ত কিন্তু উত্তেজনাপূর্ণ নয় এমন শক্তি প্রদান করে</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">সেমি-ফেয়ারিং ওজন যোগ করে, শহুরে চালনা প্রভাবিত করে</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">স্থানীয় ব্র্যান্ডের তুলনায় কম ইয়ামাহা ডিলারশিপ</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস উপলব্ধ নেই:</strong> <span class="con-description">অ্যান্টি-লক ব্রেকিংয়ের মতো আধুনিক নিরাপত্তা বৈশিষ্ট্যের অভাব</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা ফেজার এফআই ভি২ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Long-distance commuters</li>
<li class="best-for-item">Weekend touring enthusiasts</li>
<li class="best-for-item">Riders prioritizing comfort over performance</li>
<li class="best-for-item">Those needing wind protection</li>
<li class="best-for-item">Daily highway riders</li>
<li class="best-for-item">Mature riders seeking practical motorcycle</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা ফেজার এফআই ভি২ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Style-conscious younger riders</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Those seeking modern aggressive styling</li>
<li class="not-recommended-item">Riders wanting ABS safety features</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা ফেজার এফআই ভি২ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৭০,০০০ - সেমি-ফেয়ারিং সহ ১৫০সিসি এফআই এর জন্য মধ্যম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম থেকে মাঝারি - ৳৭-৯ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৬-৮ প্রতি কিমি (এফআই দক্ষতা সহ ৪৫-৫০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৫,০০০-৮,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা ফেজার এফআই ভি২ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- ট্যুরিংয়ের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- চমৎকার জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার ট্যুরিং আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- পুরাতন ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ট্যুরিং চাহিদার জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- নির্ভরযোগ্য ইয়ামাহা বিল্ড</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা ফেজার এফআই ভি২ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">বায়ু সুরক্ষার জন্য সেমি-ফেয়ারিং কতটা কার্যকর?</h3>
<p class="faq-answer">সেমি-ফেয়ারিং উপরের শরীরের জন্য ভাল বায়ু সুরক্ষা প্রদান করে এবং হাইওয়ে রাইডে ক্লান্তি হ্রাস করে, যদিও সম্পূর্ণ ফেয়ারিং মোটরসাইকেলের মতো ব্যাপক নয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি দৈনিক শহুরে যাতায়াতের জন্য উপযুক্ত?</h3>
<p class="faq-answer">হ্যাঁ, কিন্তু সেমি-ফেয়ারিং ওজন যোগ করে যা ভারী ট্রাফিকে এটিকে কম চটপটে করতে পারে। এটি জনাকীর্ণ শহুরে রাইডিংয়ের চেয়ে হাইওয়ে যাতায়াত এবং ভ্রমণের জন্য ভাল উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ভি২ উন্নতিগুলো কী?</h3>
<p class="faq-answer">ভি২ সংস্করণে মূল ফেজার এফআই এর তুলনায় আপডেটেড ইলেকট্রনিক্স, উন্নত ফুয়েল ইনজেকশন ম্যাপিং, পরিমার্জিত স্টাইলিং উপাদান এবং ভাল বিল্ড কোয়ালিটি রয়েছে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা ফেজার এফআই ভি২ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.1/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,৭০,০০০ টাকায় ইয়ামাহা ফেজার এফআই ভি২ ফুয়েল ইনজেকশন দক্ষতা, সেমি-ফেয়ারিং বায়ু সুরক্ষা এবং আরামদায়ক এরগোনমিক্স সহ একটি ব্যবহারিক ট্যুরিং-ফোকাসড প্যাকেজ প্রদান করে। এটি ৪৫-৫০ কিমি/লিটার চমৎকার জ্বালানি অর্থনীতি সহ দীর্ঘ-দূরত্ব যাতায়াত এবং সাপ্তাহিক ভ্রমণে শ্রেষ্ঠ। তবে, পুরাতন স্টাইলিং, মাঝারি পারফরমেন্স, শহুরে ব্যবহারের জন্য ভারী ওজন, এবিএসের অভাব এবং সীমিত সেবা নেটওয়ার্ক এটিকে প্রধানত পরিণত রাইডার এবং ট্যুরিং উৎসাহীদের জন্য উপযুক্ত করে তোলে যারা আধুনিক স্টাইলিং এবং স্পোর্টি পারফরমেন্সের চেয়ে আরাম এবং জ্বালানি দক্ষতাকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- এফআই দক্ষতা এবং বায়ু সুরক্ষা সহ আরামদায়ক ভ্রমণের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha Fazer FI V2 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha Fazer FI V2\n")

	return nil
}
