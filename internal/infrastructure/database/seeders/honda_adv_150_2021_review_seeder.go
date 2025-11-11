package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaAdv150Review handles seeding Honda ADV 150 2021 product review and translation
type HondaAdv150Review struct {
	BaseSeeder
}

// NewHondaAdv150ReviewSeeder creates a new HondaAdv150Review
func NewHondaAdv150ReviewSeeder() *HondaAdv150Review {
	return &HondaAdv150Review{BaseSeeder: BaseSeeder{name: "Honda ADV 150 2021 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaAdv150Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda ADV 150 2021").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda ADV 150 2021 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda ADV 150 2021 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda ADV 150 2021 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda ADV 150 2021 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda ADV 150 2021 is an adventure scooter featuring rugged styling with tall windscreen, long-travel suspension for rough roads, automatic CVT transmission, and versatile capabilities. Priced at ৳4,75,000, it combines automatic scooter convenience with adventure aesthetics, offering excellent ground clearance, wind protection, practical storage, Honda reliability, and dual-purpose versatility for riders seeking an adventure-styled automatic with touring comfort and rough road capability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda ADV 150 2021 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Adventure Scooter Design:</strong> <span class="highlight-value">Rugged styling with tall windscreen and protection</span></li>
<li class="highlight-item"><strong class="highlight-label">Automatic CVT:</strong> <span class="highlight-value">Effortless automatic transmission convenience</span></li>
<li class="highlight-item"><strong class="highlight-label">Long-Travel Suspension:</strong> <span class="highlight-value">Enhanced suspension for rough road comfort</span></li>
<li class="highlight-item"><strong class="highlight-label">Full LED Lighting:</strong> <span class="highlight-value">Modern LED headlight and taillight</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda ADV 150 2021 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Automatic Convenience:</strong> <span class="pro-description">CVT transmission provides effortless riding without gear shifting</span></li>
<li class="pro-item"><strong class="pro-label">Adventure Aesthetics:</strong> <span class="pro-description">Rugged styling with windscreen creates unique adventure scooter appeal</span></li>
<li class="pro-item"><strong class="pro-label">Good Ground Clearance:</strong> <span class="pro-description">Higher clearance than regular scooters handles rough roads better</span></li>
<li class="pro-item"><strong class="pro-label">Wind Protection:</strong> <span class="pro-description">Tall windscreen provides good touring wind protection</span></li>
<li class="pro-item"><strong class="pro-label">Practical Storage:</strong> <span class="pro-description">Spacious under-seat storage for daily needs</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven Honda engineering and dependability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda ADV 150 2021 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price:</strong> <span class="con-description">৳4,75,000 is expensive for automatic scooter segment</span></li>
<li class="con-item"><strong class="con-label">Limited Real Off-Road Capability:</strong> <span class="con-description">Adventure styling doesn't equal serious off-road performance</span></li>
<li class="con-item"><strong class="con-label">Moderate Performance:</strong> <span class="con-description">150cc scooter engine provides adequate but not exciting power</span></li>
<li class="con-item"><strong class="con-label">Expensive Parts:</strong> <span class="con-description">Premium adventure scooter means costly maintenance</span></li>
<li class="con-item"><strong class="con-label">Limited Availability:</strong> <span class="con-description">Not widely available in all markets</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda ADV 150 2021 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Adventure scooter enthusiasts</li>
<li class="best-for-item">Riders wanting automatic convenience with rugged style</li>
<li class="best-for-item">Touring and weekend exploration</li>
<li class="best-for-item">Those dealing with rough road conditions</li>
<li class="best-for-item">Urban riders seeking unique adventure aesthetics</li>
<li class="best-for-item">Mature riders prioritizing comfort and convenience</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda ADV 150 2021?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Serious off-road riding</li>
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Those seeking basic economical scooter</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda ADV 150 2021 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,75,000 - Premium for adventure scooter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low to Moderate - ৳7-9 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳6-8 per km (45-50 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳9,000-13,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda ADV 150 2021 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good scooter performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent touring comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Unique adventure scooter styling</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Premium pricing justified</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent Honda quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda ADV 150 2021 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How effective is the windscreen?</h3>
<p class="faq-answer">The tall windscreen provides good wind protection for the upper body during highway riding, reducing fatigue on longer journeys. It's adjustable and helps significantly with touring comfort.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Can it handle off-road riding?</h3>
<p class="faq-answer">While it has adventure styling and better ground clearance than regular scooters, the ADV 150 is not designed for serious off-road riding. It can handle rough roads and light gravel but is best suited for paved and semi-paved touring.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for daily commuting?</h3>
<p class="faq-answer">Yes, the ADV 150 is excellent for daily commuting with automatic convenience, practical storage, comfortable ergonomics, and good ground clearance for rough roads. The wind protection also helps with highway commuting.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda ADV 150 2021 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,75,000, the Honda ADV 150 2021 delivers a unique adventure scooter experience combining automatic CVT convenience with rugged styling, wind protection, good ground clearance, practical storage, and Honda reliability. It excels at touring and rough road riding with excellent comfort (45-50 km/l fuel efficiency). However, the high price for scooter segment, limited real off-road capability despite adventure styling, moderate performance, expensive parts, and limited availability make it best suited for adventure scooter enthusiasts, touring riders, and those seeking automatic convenience with unique adventure aesthetics over basic economy.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For adventure scooter styling with automatic convenience and touring comfort</span></p>
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
		return fmt.Errorf("error creating Honda ADV 150 2021 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda ADV 150 2021 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা এডিভি ১৫০ ২০২১ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা এডিভি ১৫০ ২০২১ হল একটি অ্যাডভেঞ্চার স্কুটার যেখানে লম্বা উইন্ডস্ক্রিন সহ রুক্ষ স্টাইলিং, রুক্ষ রাস্তার জন্য লং-ট্র্যাভেল সাসপেনশন, অটোমেটিক সিভিটি ট্রান্সমিশন এবং বহুমুখী ক্ষমতা রয়েছে। ৳৪,৭৫,০০০ টাকায় মূল্যায়িত, এটি অ্যাডভেঞ্চার নান্দনিকতার সাথে অটোমেটিক স্কুটার সুবিধা একত্রিত করে, চমৎকার গ্রাউন্ড ক্লিয়ারেন্স, বায়ু সুরক্ষা, ব্যবহারিক স্টোরেজ, হোন্ডা নির্ভরযোগ্যতা এবং ট্যুরিং আরাম এবং রুক্ষ রাস্তা ক্ষমতা সহ অ্যাডভেঞ্চার-স্টাইলড অটোমেটিক খোঁজা রাইডারদের জন্য দ্বৈত-উদ্দেশ্য বহুমুখিতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা এডিভি ১৫০ ২০২১ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">অ্যাডভেঞ্চার স্কুটার ডিজাইন:</strong> <span class="highlight-value">লম্বা উইন্ডস্ক্রিন এবং সুরক্ষা সহ রুক্ষ স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">অটোমেটিক সিভিটি:</strong> <span class="highlight-value">অনায়াস অটোমেটিক ট্রান্সমিশন সুবিধা</span></li>
<li class="highlight-item"><strong class="highlight-label">লং-ট্র্যাভেল সাসপেনশন:</strong> <span class="highlight-value">রুক্ষ রাস্তা আরামের জন্য উন্নত সাসপেনশন</span></li>
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ এলইডি লাইটিং:</strong> <span class="highlight-value">আধুনিক এলইডি হেডলাইট এবং টেইললাইট</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা এডিভি ১৫০ ২০২১ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অটোমেটিক সুবিধা:</strong> <span class="pro-description">সিভিটি ট্রান্সমিশন গিয়ার শিফটিং ছাড়াই অনায়াস রাইডিং প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">অ্যাডভেঞ্চার নান্দনিকতা:</strong> <span class="pro-description">উইন্ডস্ক্রিন সহ রুক্ষ স্টাইলিং অনন্য অ্যাডভেঞ্চার স্কুটার আবেদন তৈরি করে</span></li>
<li class="pro-item"><strong class="pro-label">ভাল গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="pro-description">নিয়মিত স্কুটারের চেয়ে উচ্চ ক্লিয়ারেন্স রুক্ষ রাস্তা ভালভাবে পরিচালনা করে</span></li>
<li class="pro-item"><strong class="pro-label">বায়ু সুরক্ষা:</strong> <span class="pro-description">লম্বা উইন্ডস্ক্রিন ভাল ট্যুরিং বায়ু সুরক্ষা প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক স্টোরেজ:</strong> <span class="pro-description">দৈনিক চাহিদার জন্য প্রশস্ত আন্ডার-সিট স্টোরেজ</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত হোন্ডা ইঞ্জিনিয়ারিং এবং নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা এডিভি ১৫০ ২০২১ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">৳৪,৭৫,০০০ অটোমেটিক স্কুটার সেগমেন্টের জন্য ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সীমিত প্রকৃত অফ-রোড ক্ষমতা:</strong> <span class="con-description">অ্যাডভেঞ্চার স্টাইলিং গুরুতর অফ-রোড পারফরমেন্সের সমান নয়</span></li>
<li class="con-item"><strong class="con-label">মাঝারি পারফরমেন্স:</strong> <span class="con-description">১৫০সিসি স্কুটার ইঞ্জিন পর্যাপ্ত কিন্তু উত্তেজনাপূর্ণ নয় এমন শক্তি প্রদান করে</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল যন্ত্রাংশ:</strong> <span class="con-description">প্রিমিয়াম অ্যাডভেঞ্চার স্কুটার মানে ব্যয়বহুল রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">সীমিত প্রাপ্যতা:</strong> <span class="con-description">সব বাজারে ব্যাপকভাবে উপলব্ধ নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা এডিভি ১৫০ ২০২১ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Adventure scooter enthusiasts</li>
<li class="best-for-item">Riders wanting automatic convenience with rugged style</li>
<li class="best-for-item">Touring and weekend exploration</li>
<li class="best-for-item">Those dealing with rough road conditions</li>
<li class="best-for-item">Urban riders seeking unique adventure aesthetics</li>
<li class="best-for-item">Mature riders prioritizing comfort and convenience</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা এডিভি ১৫০ ২০২১ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Serious off-road riding</li>
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Those seeking basic economical scooter</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা এডিভি ১৫০ ২০২১ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৪,৭৫,০০০ - অ্যাডভেঞ্চার স্কুটারের জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম থেকে মাঝারি - ৳৭-৯ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৬-৮ প্রতি কিমি (৪৫-৫০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৯,০০০-১৩,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা এডিভি ১৫০ ২০২১ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল স্কুটার পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- ভাল জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার ট্যুরিং আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- অনন্য অ্যাডভেঞ্চার স্কুটার স্টাইলিং</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- প্রিমিয়াম মূল্য ন্যায্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার হোন্ডা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা এডিভি ১৫০ ২০২১ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">উইন্ডস্ক্রিন কতটা কার্যকর?</h3>
<p class="faq-answer">লম্বা উইন্ডস্ক্রিন হাইওয়ে রাইডিংয়ের সময় উপরের শরীরের জন্য ভাল বায়ু সুরক্ষা প্রদান করে, দীর্ঘ যাত্রায় ক্লান্তি হ্রাস করে। এটি সামঞ্জস্যযোগ্য এবং ট্যুরিং আরামে উল্লেখযোগ্যভাবে সাহায্য করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি অফ-রোড রাইডিং পরিচালনা করতে পারে?</h3>
<p class="faq-answer">যদিও এতে অ্যাডভেঞ্চার স্টাইলিং এবং নিয়মিত স্কুটারের চেয়ে ভাল গ্রাউন্ড ক্লিয়ারেন্স রয়েছে, এডিভি ১৫০ গুরুতর অফ-রোড রাইডিংয়ের জন্য ডিজাইন করা হয়নি। এটি রুক্ষ রাস্তা এবং হালকা নুড়ি পরিচালনা করতে পারে তবে পাকা এবং আধা-পাকা ভ্রমণের জন্য সবচেয়ে উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি দৈনিক যাতায়াতের জন্য উপযুক্ত?</h3>
<p class="faq-answer">হ্যাঁ, এডিভি ১৫০ অটোমেটিক সুবিধা, ব্যবহারিক স্টোরেজ, আরামদায়ক এরগোনমিক্স এবং রুক্ষ রাস্তার জন্য ভাল গ্রাউন্ড ক্লিয়ারেন্স সহ দৈনিক যাতায়াতের জন্য চমৎকার। বায়ু সুরক্ষা হাইওয়ে যাতায়াতেও সাহায্য করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা এডিভি ১৫০ ২০২১ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৪,৭৫,০০০ টাকায় হোন্ডা এডিভি ১৫০ ২০২১ অটোমেটিক সিভিটি সুবিধা রুক্ষ স্টাইলিং, বায়ু সুরক্ষা, ভাল গ্রাউন্ড ক্লিয়ারেন্স, ব্যবহারিক স্টোরেজ এবং হোন্ডা নির্ভরযোগ্যতার সাথে একত্রিত করে একটি অনন্য অ্যাডভেঞ্চার স্কুটার অভিজ্ঞতা প্রদান করে। এটি চমৎকার আরাম (৪৫-৫০ কিমি/লিটার জ্বালানি দক্ষতা) সহ ট্যুরিং এবং রুক্ষ রাস্তা রাইডিংয়ে শ্রেষ্ঠ। তবে, স্কুটার সেগমেন্টের জন্য উচ্চ মূল্য, অ্যাডভেঞ্চার স্টাইলিং সত্ত্বেও সীমিত প্রকৃত অফ-রোড ক্ষমতা, মাঝারি পারফরমেন্স, ব্যয়বহুল যন্ত্রাংশ এবং সীমিত প্রাপ্যতা এটিকে অ্যাডভেঞ্চার স্কুটার উৎসাহী, ট্যুরিং রাইডার এবং মৌলিক অর্থনীতির চেয়ে অনন্য অ্যাডভেঞ্চার নান্দনিকতার সাথে অটোমেটিক সুবিধা খোঁজাদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অটোমেটিক সুবিধা এবং ট্যুরিং আরাম সহ অ্যাডভেঞ্চার স্কুটার স্টাইলিংয়ের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda ADV 150 2021 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda ADV 150 2021\n")

	return nil
}
