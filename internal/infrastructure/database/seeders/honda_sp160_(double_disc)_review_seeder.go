package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaSp160DoubleDiscReview handles seeding Honda SP160 (Double Disc) product review and translation
type HondaSp160DoubleDiscReview struct {
	BaseSeeder
}

// NewHondaSp160DoubleDiscReviewSeeder creates a new HondaSp160DoubleDiscReview
func NewHondaSp160DoubleDiscReviewSeeder() *HondaSp160DoubleDiscReview {
	return &HondaSp160DoubleDiscReview{BaseSeeder: BaseSeeder{name: "Honda SP160 (Double Disc) Review"}}
}

// Seed implements the Seeder interface
func (s *HondaSp160DoubleDiscReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda SP160 (Double Disc)").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda SP160 (Double Disc) product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda SP160 (Double Disc) product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda SP160 (Double Disc) already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda SP160 (Double Disc) Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda SP160 Double Disc is a sporty 160cc commuter motorcycle featuring aggressive styling, dual disc brakes for enhanced safety, fuel injection technology, and sporty ergonomics. Priced at ৳2,25,000, it offers excellent value with superior braking performance, modern FI efficiency, sporty design appeal, Honda reliability, and practical features making it ideal for young riders seeking sporty styling with safety features in an affordable, fuel-efficient package.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda SP160 (Double Disc) Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Dual Disc Brakes:</strong> <span class="highlight-value">Enhanced safety with front and rear disc brakes</span></li>
<li class="highlight-item"><strong class="highlight-label">160cc PGM-Fi Engine:</strong> <span class="highlight-value">Fuel injection for efficiency and performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Sporty Styling:</strong> <span class="highlight-value">Aggressive design with sharp graphics</span></li>
<li class="highlight-item"><strong class="highlight-label">Affordable Pricing:</strong> <span class="highlight-value">Great features at competitive price point</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda SP160 (Double Disc) Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Superior Braking Safety:</strong> <span class="pro-description">Dual disc brakes provide excellent stopping power for enhanced safety</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Injection Benefits:</strong> <span class="pro-description">PGM-Fi ensures better fuel economy and instant starts</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Appeal:</strong> <span class="pro-description">Aggressive styling attracts young riders and enthusiasts</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Value:</strong> <span class="pro-description">160cc with FI and double disc at affordable ৳2,25,000</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven Honda quality and long-term dependability</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">Achieves 50-55 km/l with FI technology</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda SP160 (Double Disc) Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Moderate Performance:</strong> <span class="con-description">160cc provides adequate but not thrilling power delivery</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Limited electronics and modern features</span></li>
<li class="con-item"><strong class="con-label">Firm Suspension:</strong> <span class="con-description">Sporty setup may feel harsh on rough roads</span></li>
<li class="con-item"><strong class="con-label">No ABS Option:</strong> <span class="con-description">Misses advanced anti-lock braking system</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Not as premium as higher-end Honda models</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda SP160 (Double Disc) in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban commuters</li>
<li class="best-for-item">First-time 160cc buyers</li>
<li class="best-for-item">Budget-conscious riders seeking safety</li>
<li class="best-for-item">Daily city and highway use</li>
<li class="best-for-item">Those prioritizing fuel efficiency and braking</li>
<li class="best-for-item">Sporty styling enthusiasts on budget</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda SP160 (Double Disc)?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts seeking high power</li>
<li class="not-recommended-item">Riders wanting premium features</li>
<li class="not-recommended-item">Those requiring ABS safety</li>
<li class="not-recommended-item">Soft suspension comfort seekers</li>
<li class="not-recommended-item">Long-distance touring riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda SP160 (Double Disc) Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,25,000 - Excellent for 160cc FI with dual disc</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳6-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳5-7 per km (50-55 km/l with FI)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳5,000-8,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda SP160 (Double Disc) Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good for daily use</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Decent commuter comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Sporty aggressive design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Outstanding value for money</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Good Honda quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda SP160 (Double Disc) Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does it compare to single disc variant?</h3>
<p class="faq-answer">The double disc variant offers significantly better braking performance and safety, especially in emergency situations and wet conditions. The ৳26,000 price difference is justified for the enhanced braking capability.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for highway riding?</h3>
<p class="faq-answer">Yes, the SP160 handles highway riding well with adequate 160cc power for cruising at 80-90 km/h. The dual disc brakes provide confidence at higher speeds, though it's primarily a city commuter.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What's the real-world fuel economy?</h3>
<p class="faq-answer">Real-world fuel economy ranges from 50-55 km/l with careful riding, dropping to 45-48 km/l in heavy traffic. The PGM-Fi system ensures consistent efficiency across conditions.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda SP160 (Double Disc) in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,25,000, the Honda SP160 Double Disc offers outstanding value as a sporty 160cc commuter with superior dual disc braking safety, fuel injection efficiency (50-55 km/l), aggressive styling, and Honda reliability. It excels at daily urban commuting with excellent fuel economy and enhanced safety features. However, moderate performance, basic features, firm suspension, lack of ABS, and average build quality make it best suited for young urban commuters, first-time 160cc buyers, and budget-conscious riders who prioritize fuel efficiency, braking safety, and sporty styling in an affordable package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For excellent value 160cc with dual disc safety and FI efficiency</span></p>
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
		return fmt.Errorf("error creating Honda SP160 (Double Disc) review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda SP160 (Double Disc) (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা এসপি১৬০ ডাবল ডিস্ক হল একটি স্পোর্টি ১৬০সিসি কম্যুটার মোটরসাইকেল যেখানে আক্রমণাত্মক স্টাইলিং, উন্নত নিরাপত্তার জন্য ডুয়াল ডিস্ক ব্রেক, ফুয়েল ইনজেকশন প্রযুক্তি এবং স্পোর্টি এরগোনমিক্স রয়েছে। ৳২,২৫,০০০ টাকায় মূল্যায়িত, এটি উন্নত ব্রেকিং পারফরমেন্স, আধুনিক এফআই দক্ষতা, স্পোর্টি ডিজাইন আবেদন, হোন্ডা নির্ভরযোগ্যতা এবং সাশ্রয়ী, জ্বালানি-দক্ষ প্যাকেজে নিরাপত্তা বৈশিষ্ট্য সহ স্পোর্টি স্টাইলিং খোঁজা তরুণ রাইডারদের জন্য আদর্শ ব্যবহারিক বৈশিষ্ট্য সহ চমৎকার মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ডুয়াল ডিস্ক ব্রেক:</strong> <span class="highlight-value">সামনে এবং পিছনের ডিস্ক ব্রেক দিয়ে উন্নত নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">১৬০সিসি পিজিএম-এফআই ইঞ্জিন:</strong> <span class="highlight-value">দক্ষতা এবং পারফরমেন্সের জন্য ফুয়েল ইনজেকশন</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্টি স্টাইলিং:</strong> <span class="highlight-value">তীক্ষ্ণ গ্রাফিক্স সহ আক্রমণাত্মক ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">সাশ্রয়ী মূল্য:</strong> <span class="highlight-value">প্রতিযোগিতামূলক মূল্য পয়েন্টে দুর্দান্ত বৈশিষ্ট্য</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">উন্নত ব্রেকিং নিরাপত্তা:</strong> <span class="pro-description">ডুয়াল ডিস্ক ব্রেক উন্নত নিরাপত্তার জন্য চমৎকার থামার শক্তি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন সুবিধা:</strong> <span class="pro-description">পিজিএম-এফআই ভাল জ্বালানি অর্থনীতি এবং তাৎক্ষণিক স্টার্ট নিশ্চিত করে</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি আবেদন:</strong> <span class="pro-description">আক্রমণাত্মক স্টাইলিং তরুণ রাইডার এবং উৎসাহীদের আকর্ষণ করে</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মূল্য:</strong> <span class="pro-description">সাশ্রয়ী ৳২,২৫,০০০ এ এফআই এবং ডাবল ডিস্ক সহ ১৬০সিসি</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত হোন্ডা গুণমান এবং দীর্ঘমেয়াদী নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">ভাল জ্বালানি সাশ্রয়:</strong> <span class="pro-description">এফআই প্রযুক্তি দিয়ে ৫০-৫৫ কিমি/লিটার অর্জন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মাঝারি পারফরমেন্স:</strong> <span class="con-description">১৬০সিসি পর্যাপ্ত কিন্তু রোমাঞ্চকর নয় এমন পাওয়ার ডেলিভারি প্রদান করে</span></li>
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">সীমিত ইলেকট্রনিক্স এবং আধুনিক বৈশিষ্ট্য</span></li>
<li class="con-item"><strong class="con-label">শক্ত সাসপেনশন:</strong> <span class="con-description">স্পোর্টি সেটআপ রুক্ষ রাস্তায় কঠোর মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস বিকল্প নেই:</strong> <span class="con-description">অ্যান্টি-লক ব্রেকিং সিস্টেম অনুপস্থিত</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">উচ্চ-সীমার হোন্ডা মডেলের মতো প্রিমিয়াম নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban commuters</li>
<li class="best-for-item">First-time 160cc buyers</li>
<li class="best-for-item">Budget-conscious riders seeking safety</li>
<li class="best-for-item">Daily city and highway use</li>
<li class="best-for-item">Those prioritizing fuel efficiency and braking</li>
<li class="best-for-item">Sporty styling enthusiasts on budget</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts seeking high power</li>
<li class="not-recommended-item">Riders wanting premium features</li>
<li class="not-recommended-item">Those requiring ABS safety</li>
<li class="not-recommended-item">Soft suspension comfort seekers</li>
<li class="not-recommended-item">Long-distance touring riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,২৫,০০০ - ডুয়াল ডিস্ক সহ ১৬০সিসি এফআই এর জন্য চমৎকার</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম - ৳৬-৮ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৫-৭ প্রতি কিমি (এফআই সহ ৫০-৫৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৫,০০০-৮,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- দৈনিক ব্যবহারের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- ভাল কম্যুটার আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- স্পোর্টি আক্রমণাত্মক ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- অসাধারণ ভ্যালু ফর মানি</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- ভাল হোন্ডা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এটি একক ডিস্ক ভ্যারিয়েন্টের সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">ডাবল ডিস্ক ভ্যারিয়েন্ট উল্লেখযোগ্যভাবে ভাল ব্রেকিং পারফরমেন্স এবং নিরাপত্তা প্রদান করে, বিশেষত জরুরি পরিস্থিতি এবং ভেজা অবস্থায়। উন্নত ব্রেকিং ক্ষমতার জন্য ৳২৬,০০০ মূল্য পার্থক্য ন্যায্য।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি হাইওয়ে রাইডিংয়ের জন্য উপযুক্ত?</h3>
<p class="faq-answer">হ্যাঁ, এসপি১৬০ ৮০-৯০ কিমি/ঘণ্টায় ক্রুজিংয়ের জন্য পর্যাপ্ত ১৬০সিসি শক্তি দিয়ে হাইওয়ে রাইডিং ভালভাবে পরিচালনা করে। ডুয়াল ডিস্ক ব্রেক উচ্চ গতিতে আত্মবিশ্বাস প্রদান করে, যদিও এটি মূলত একটি শহুরে কম্যুটার।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাস্তব-বিশ্ব জ্বালানি অর্থনীতি কত?</h3>
<p class="faq-answer">বাস্তব-বিশ্ব জ্বালানি অর্থনীতি সাবধানে রাইডিং সহ ৫০-৫৫ কিমি/লিটার পরিসীমা, ভারী ট্রাফিকে ৪৫-৪৮ কিমি/লিটারে নেমে যায়। পিজিএম-এফআই সিস্টেম অবস্থা জুড়ে সামঞ্জস্যপূর্ণ দক্ষতা নিশ্চিত করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা এসপি১৬০ (ডাবল ডিস্ক) কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,২৫,০০০ টাকায় হোন্ডা এসপি১৬০ ডাবল ডিস্ক উন্নত ডুয়াল ডিস্ক ব্রেকিং নিরাপত্তা, ফুয়েল ইনজেকশন দক্ষতা (৫০-৫৫ কিমি/লিটার), আক্রমণাত্মক স্টাইলিং এবং হোন্ডা নির্ভরযোগ্যতা সহ একটি স্পোর্টি ১৬০সিসি কম্যুটার হিসাবে অসাধারণ মূল্য প্রদান করে। এটি চমৎকার জ্বালানি অর্থনীতি এবং উন্নত নিরাপত্তা বৈশিষ্ট্য সহ দৈনিক শহুরে যাতায়াতে শ্রেষ্ঠ। তবে, মাঝারি পারফরমেন্স, মৌলিক বৈশিষ্ট্য, শক্ত সাসপেনশন, এবিএসের অভাব এবং গড় বিল্ড কোয়ালিটি এটিকে তরুণ শহুরে যাত্রী, প্রথমবার ১৬০সিসি ক্রেতা এবং সাশ্রয়ী প্যাকেজে জ্বালানি দক্ষতা, ব্রেকিং নিরাপত্তা এবং স্পোর্টি স্টাইলিংকে অগ্রাধিকার দেওয়া বাজেট-সচেতন রাইডারদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ডুয়াল ডিস্ক নিরাপত্তা এবং এফআই দক্ষতা সহ চমৎকার মূল্য ১৬০সিসির জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda SP160 (Double Disc) already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda SP160 (Double Disc)\n")

	return nil
}
