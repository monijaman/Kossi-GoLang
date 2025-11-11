package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCb150XReview handles seeding Honda CB150X product review and translation
type HondaCb150XReview struct {
	BaseSeeder
}

// NewHondaCb150XReviewSeeder creates a new HondaCb150XReview
func NewHondaCb150XReviewSeeder() *HondaCb150XReview {
	return &HondaCb150XReview{BaseSeeder: BaseSeeder{name: "Honda CB150X Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCb150XReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB150X").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda CB150X product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda CB150X product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CB150X already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CB150X Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda CB150X is an adventure-style 150cc motorcycle featuring rugged off-road inspired design, tall suspension with high ground clearance, upright ergonomics, and versatile capabilities. Priced at ৳5,20,000, it combines adventure aesthetics with urban practicality, offering excellent ground clearance for rough roads, comfortable touring position, Honda reliability, and dual-purpose versatility for riders seeking an adventure-styled commuter with go-anywhere capability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB150X Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Adventure Styling:</strong> <span class="highlight-value">Rugged off-road inspired design with rally aesthetics</span></li>
<li class="highlight-item"><strong class="highlight-label">High Ground Clearance:</strong> <span class="highlight-value">220mm clearance for rough roads and obstacles</span></li>
<li class="highlight-item"><strong class="highlight-label">Tall Suspension:</strong> <span class="highlight-value">Long-travel suspension for comfort on rough terrain</span></li>
<li class="highlight-item"><strong class="highlight-label">Upright Position:</strong> <span class="highlight-value">Comfortable touring ergonomics with tall seat</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB150X Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Ground Clearance:</strong> <span class="pro-description">220mm clearance handles potholes, speed bumps, and rough roads effortlessly</span></li>
<li class="pro-item"><strong class="pro-label">Adventure Aesthetics:</strong> <span class="pro-description">Rugged styling stands out and appeals to adventure enthusiasts</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Touring Position:</strong> <span class="pro-description">Upright ergonomics reduce fatigue on long rides</span></li>
<li class="pro-item"><strong class="pro-label">Versatile Capability:</strong> <span class="pro-description">Handles both city commuting and light off-road trails</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven Honda engineering and long-term dependability</span></li>
<li class="pro-item"><strong class="pro-label">Good Visibility:</strong> <span class="pro-description">Tall riding position provides excellent traffic visibility</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB150X Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Tall Seat Height:</strong> <span class="con-description">830mm seat height challenging for shorter riders</span></li>
<li class="con-item"><strong class="con-label">Premium Pricing:</strong> <span class="con-description">৳5,20,000 is expensive for 150cc adventure-styled bike</span></li>
<li class="con-item"><strong class="con-label">Limited Off-Road Capability:</strong> <span class="con-description">Street-biased tires and components limit serious off-roading</span></li>
<li class="con-item"><strong class="con-label">Less City Maneuverability:</strong> <span class="con-description">Tall design makes tight traffic navigation less agile</span></li>
<li class="con-item"><strong class="con-label">Higher Service Costs:</strong> <span class="con-description">Specialized components mean expensive maintenance</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda CB150X in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Adventure styling enthusiasts</li>
<li class="best-for-item">Riders dealing with rough road conditions</li>
<li class="best-for-item">Tall riders comfortable with high seat height</li>
<li class="best-for-item">Weekend touring and exploration</li>
<li class="best-for-item">Those seeking unique adventure aesthetics</li>
<li class="best-for-item">Light off-road trail exploration</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda CB150X?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Shorter riders uncomfortable with tall bikes</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Serious off-road riding</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Those seeking nimble urban commuter</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB150X Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳5,20,000 - Premium for adventure-styled 150cc</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-10 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳7-9 per km (40-45 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳10,000-14,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB150X Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Adequate for adventure touring</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent touring comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Unique adventure styling</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Premium pricing</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent Honda quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda CB150X Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Can it handle serious off-road riding?</h3>
<p class="faq-answer">The CB150X can handle light off-road trails and rough roads, but it's not designed for serious off-roading due to street-biased tires and components. It's best suited for adventure-styled urban riding with occasional light trail exploration.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the tall seat height a problem?</h3>
<p class="faq-answer">The 830mm seat height can be challenging for riders under 5'6" tall. Shorter riders may struggle with flat-footing at stops. However, the seat is relatively narrow, which helps somewhat with reach.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it compare to regular CB150R?</h3>
<p class="faq-answer">The CB150X offers adventure styling, higher ground clearance, taller suspension, and upright ergonomics compared to CB150R's street-focused design. It's about ৳1,65,000 more expensive but provides better rough road capability and touring comfort.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB150X in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳5,20,000, the Honda CB150X offers a unique adventure-styled 150cc package with excellent ground clearance (220mm), comfortable touring ergonomics, rugged aesthetics, and Honda reliability. It excels at handling rough roads, providing touring comfort, and delivering versatile dual-purpose capability. However, the premium pricing, tall seat height (830mm) challenging for shorter riders, limited serious off-road capability, less city maneuverability, and higher service costs make it best suited for adventure styling enthusiasts, tall riders, and those dealing with rough road conditions who prioritize go-anywhere versatility over urban agility.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For adventure styling with excellent ground clearance and touring comfort</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Honda CB150X review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CB150X (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা সিবি১৫০এক্স রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা সিবি১৫০এক্স হল একটি অ্যাডভেঞ্চার-স্টাইল ১৫০সিসি মোটরসাইকেল যেখানে রুক্ষ অফ-রোড অনুপ্রাণিত ডিজাইন, উচ্চ গ্রাউন্ড ক্লিয়ারেন্স সহ লম্বা সাসপেনশন, সোজা এরগোনমিক্স এবং বহুমুখী ক্ষমতা রয়েছে। ৳৫,২০,০০০ টাকায় মূল্যায়িত, এটি শহুরে ব্যবহারিকতার সাথে অ্যাডভেঞ্চার নান্দনিকতা একত্রিত করে, রুক্ষ রাস্তার জন্য চমৎকার গ্রাউন্ড ক্লিয়ারেন্স, আরামদায়ক ট্যুরিং পজিশন, হোন্ডা নির্ভরযোগ্যতা এবং যে কোনো জায়গায় যেতে সক্ষম অ্যাডভেঞ্চার-স্টাইলড কম্যুটার খোঁজা রাইডারদের জন্য দ্বৈত-উদ্দেশ্য বহুমুখিতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা সিবি১৫০এক্স এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">অ্যাডভেঞ্চার স্টাইলিং:</strong> <span class="highlight-value">র‍্যালি নান্দনিকতা সহ রুক্ষ অফ-রোড অনুপ্রাণিত ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">উচ্চ গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="highlight-value">রুক্ষ রাস্তা এবং বাধার জন্য ২২০মিমি ক্লিয়ারেন্স</span></li>
<li class="highlight-item"><strong class="highlight-label">লম্বা সাসপেনশন:</strong> <span class="highlight-value">রুক্ষ ভূখণ্ডে আরামের জন্য লং-ট্র্যাভেল সাসপেনশন</span></li>
<li class="highlight-item"><strong class="highlight-label">সোজা পজিশন:</strong> <span class="highlight-value">লম্বা সিট সহ আরামদায়ক ট্যুরিং এরগোনমিক্স</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা সিবি১৫০এক্স এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="pro-description">২২০মিমি ক্লিয়ারেন্স গর্ত, স্পিড বাম্প এবং রুক্ষ রাস্তা অনায়াসে পরিচালনা করে</span></li>
<li class="pro-item"><strong class="pro-label">অ্যাডভেঞ্চার নান্দনিকতা:</strong> <span class="pro-description">রুক্ষ স্টাইলিং আলাদা হয়ে দাঁড়ায় এবং অ্যাডভেঞ্চার উৎসাহীদের আকর্ষণ করে</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক ট্যুরিং পজিশন:</strong> <span class="pro-description">সোজা এরগোনমিক্স দীর্ঘ রাইডে ক্লান্তি হ্রাস করে</span></li>
<li class="pro-item"><strong class="pro-label">বহুমুখী ক্ষমতা:</strong> <span class="pro-description">শহুরে যাতায়াত এবং হালকা অফ-রোড ট্রেইল উভয়ই পরিচালনা করে</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত হোন্ডা ইঞ্জিনিয়ারিং এবং দীর্ঘমেয়াদী নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">ভাল দৃশ্যমানতা:</strong> <span class="pro-description">লম্বা রাইডিং পজিশন চমৎকার ট্রাফিক দৃশ্যমানতা প্রদান করে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা সিবি১৫০এক্স এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">লম্বা সিট উচ্চতা:</strong> <span class="con-description">৮৩০মিমি সিট উচ্চতা খাটো রাইডারদের জন্য চ্যালেঞ্জিং</span></li>
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳৫,২০,০০০ ১৫০সিসি অ্যাডভেঞ্চার-স্টাইলড বাইকের জন্য ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সীমিত অফ-রোড ক্ষমতা:</strong> <span class="con-description">রাস্তা-পক্ষপাতদুষ্ট টায়ার এবং কম্পোনেন্ট গুরুতর অফ-রোডিং সীমিত করে</span></li>
<li class="con-item"><strong class="con-label">কম শহুরে চালনা:</strong> <span class="con-description">লম্বা ডিজাইন আঁটসাঁট ট্রাফিক নেভিগেশনকে কম চটপটে করে</span></li>
<li class="con-item"><strong class="con-label">উচ্চতর সেবা খরচ:</strong> <span class="con-description">বিশেষায়িত কম্পোনেন্ট মানে ব্যয়বহুল রক্ষণাবেক্ষণ</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা সিবি১৫০এক্স কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Adventure styling enthusiasts</li>
<li class="best-for-item">Riders dealing with rough road conditions</li>
<li class="best-for-item">Tall riders comfortable with high seat height</li>
<li class="best-for-item">Weekend touring and exploration</li>
<li class="best-for-item">Those seeking unique adventure aesthetics</li>
<li class="best-for-item">Light off-road trail exploration</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা সিবি১৫০এক্স কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Shorter riders uncomfortable with tall bikes</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Serious off-road riding</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Those seeking nimble urban commuter</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা সিবি১৫০এক্স এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৫,২০,০০০ - অ্যাডভেঞ্চার-স্টাইলড ১৫০সিসির জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি - ৳৮-১০ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৭-৯ প্রতি কিমি (৪০-৪৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১০,০০০-১৪,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা সিবি১৫০এক্স পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- অ্যাডভেঞ্চার ট্যুরিংয়ের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- চমৎকার ট্যুরিং আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- অনন্য অ্যাডভেঞ্চার স্টাইলিং</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- প্রিমিয়াম মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- চমৎকার হোন্ডা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা সিবি১৫০এক্স সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এটি কি গুরুতর অফ-রোড রাইডিং পরিচালনা করতে পারে?</h3>
<p class="faq-answer">সিবি১৫০এক্স হালকা অফ-রোড ট্রেইল এবং রুক্ষ রাস্তা পরিচালনা করতে পারে, কিন্তু রাস্তা-পক্ষপাতদুষ্ট টায়ার এবং কম্পোনেন্টের কারণে এটি গুরুতর অফ-রোডিংয়ের জন্য ডিজাইন করা হয়নি। এটি মাঝে মাঝে হালকা ট্রেইল অন্বেষণ সহ অ্যাডভেঞ্চার-স্টাইলড শহুরে রাইডিংয়ের জন্য সবচেয়ে উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">লম্বা সিট উচ্চতা কি সমস্যা?</h3>
<p class="faq-answer">৮৩০মিমি সিট উচ্চতা ৫'৬" এর নিচে লম্বা রাইডারদের জন্য চ্যালেঞ্জিং হতে পারে। খাটো রাইডাররা স্টপে ফ্ল্যাট-ফুটিং নিয়ে সংগ্রাম করতে পারে। তবে, সিট অপেক্ষাকৃত সরু, যা পৌঁছাতে কিছুটা সাহায্য করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি নিয়মিত সিবি১৫০আর এর সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">সিবি১৫০এক্স সিবি১৫০আর এর রাস্তা-ফোকাসড ডিজাইনের তুলনায় অ্যাডভেঞ্চার স্টাইলিং, উচ্চতর গ্রাউন্ড ক্লিয়ারেন্স, লম্বা সাসপেনশন এবং সোজা এরগোনমিক্স প্রদান করে। এটি প্রায় ৳১,৬৫,০০০ বেশি ব্যয়বহুল কিন্তু ভাল রুক্ষ রাস্তা ক্ষমতা এবং ট্যুরিং আরাম প্রদান করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা সিবি১৫০এক্স কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৫,২০,০০০ টাকায় হোন্ডা সিবি১৫০এক্স চমৎকার গ্রাউন্ড ক্লিয়ারেন্স (২২০মিমি), আরামদায়ক ট্যুরিং এরগোনমিক্স, রুক্ষ নান্দনিকতা এবং হোন্ডা নির্ভরযোগ্যতা সহ একটি অনন্য অ্যাডভেঞ্চার-স্টাইলড ১৫০সিসি প্যাকেজ প্রদান করে। এটি রুক্ষ রাস্তা পরিচালনা, ট্যুরিং আরাম প্রদান এবং বহুমুখী দ্বৈত-উদ্দেশ্য ক্ষমতা প্রদানে শ্রেষ্ঠ। তবে, প্রিমিয়াম মূল্য, লম্বা সিট উচ্চতা (৮৩০মিমি) খাটো রাইডারদের জন্য চ্যালেঞ্জিং, সীমিত গুরুতর অফ-রোড ক্ষমতা, কম শহুরে চালনা এবং উচ্চতর সেবা খরচ এটিকে অ্যাডভেঞ্চার স্টাইলিং উৎসাহী, লম্বা রাইডার এবং শহুরে চপলতার চেয়ে যে কোনো জায়গায় যেতে সক্ষম বহুমুখিতাকে অগ্রাধিকার দেওয়া রুক্ষ রাস্তা অবস্থার সাথে মোকাবিলা করা ব্যক্তিদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চমৎকার গ্রাউন্ড ক্লিয়ারেন্স এবং ট্যুরিং আরাম সহ অ্যাডভেঞ্চার স্টাইলিংয়ের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CB150X already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CB150X\n")

	return nil
}
