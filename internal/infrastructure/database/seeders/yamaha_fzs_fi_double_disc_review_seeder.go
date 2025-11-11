package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaFzsFiDoubleDiscReview handles seeding Yamaha FZS FI Double Disc product review and translation
type YamahaFzsFiDoubleDiscReview struct {
	BaseSeeder
}

// NewYamahaFzsFiDoubleDiscReviewSeeder creates a new YamahaFzsFiDoubleDiscReview
func NewYamahaFzsFiDoubleDiscReviewSeeder() *YamahaFzsFiDoubleDiscReview {
	return &YamahaFzsFiDoubleDiscReview{BaseSeeder: BaseSeeder{name: "Yamaha FZS FI Double Disc Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaFzsFiDoubleDiscReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha FZS FI Double Disc").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha FZS FI Double Disc product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha FZS FI Double Disc product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha FZS FI Double Disc already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha FZS FI Double Disc Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha FZS FI Double Disc is an entry-level 150cc motorcycle featuring fuel injection technology, dual disc brakes, aggressive street styling, and practical features. Priced at ৳2,36,500, it offers enhanced braking safety with front and rear disc brakes, modern FI efficiency, sporty naked design, and Yamaha's proven reliability for urban commuters seeking safety, style, and performance in an affordable package.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha FZS FI Double Disc Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Dual Disc Brakes:</strong> <span class="highlight-value">Enhanced braking safety with front and rear discs</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Injection System:</strong> <span class="highlight-value">Modern FI for better efficiency and performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Affordable Pricing:</strong> <span class="highlight-value">Entry-level pricing with premium features</span></li>
<li class="highlight-item"><strong class="highlight-label">Muscular Styling:</strong> <span class="highlight-value">Aggressive street naked design appeal</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha FZS FI Double Disc Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Superior Braking Safety:</strong> <span class="pro-description">Dual disc brakes provide excellent stopping power in all conditions</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Injection Benefits:</strong> <span class="pro-description">Better fuel economy, instant starts, and smooth power delivery</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Muscular FZ series design attracts younger buyers</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Premium Features:</strong> <span class="pro-description">Gets FI and double disc at entry-level pricing</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Yamaha's proven reliability and component quality</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Fuel Economy:</strong> <span class="pro-description">Achieves 50-55 km/l with FI technology</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha FZS FI Double Disc Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Moderate Performance:</strong> <span class="con-description">150cc engine provides adequate but not thrilling power</span></li>
<li class="con-item"><strong class="con-label">Firm Suspension:</strong> <span class="con-description">May feel harsh on rough roads and potholes</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Limited electronic features compared to premium models</span></li>
<li class="con-item"><strong class="con-label">No ABS Option:</strong> <span class="con-description">Misses advanced anti-lock braking safety feature</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha dealerships compared to local brands</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha FZS FI Double Disc in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">First-time motorcycle buyers</li>
<li class="best-for-item">Safety-conscious urban commuters</li>
<li class="best-for-item">Young riders seeking aggressive styling</li>
<li class="best-for-item">Daily city and highway riders</li>
<li class="best-for-item">Budget-conscious buyers wanting dual disc safety</li>
<li class="best-for-item">Those prioritizing braking performance</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha FZS FI Double Disc?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts seeking high power</li>
<li class="not-recommended-item">Riders needing soft suspension comfort</li>
<li class="not-recommended-item">Those wanting advanced electronics</li>
<li class="not-recommended-item">ABS safety feature requirement</li>
<li class="not-recommended-item">Long-distance touring riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha FZS FI Double Disc Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,36,500 - Excellent for FI with dual disc brakes</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳6-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳5-7 per km (50-55 km/l with FI efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳5,000-8,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha FZS FI Double Disc Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Adequate for daily use</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- Firm suspension setup</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Aggressive muscular design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Great features for price</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Reliable Yamaha build</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha FZS FI Double Disc Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How much safer are dual disc brakes compared to single disc?</h3>
<p class="faq-answer">Dual disc brakes provide significantly better stopping power, especially in emergency braking situations and wet conditions, distributing braking force more evenly and reducing stopping distances.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for daily city commuting?</h3>
<p class="faq-answer">Yes, the FZS FI Double Disc is excellent for city commuting with good fuel economy, manageable size, superior braking safety, and practical features making it ideal for urban traffic.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What's the actual fuel economy in real-world conditions?</h3>
<p class="faq-answer">Real-world fuel economy ranges from 50-55 km/l depending on riding style and conditions, with the FI technology ensuring consistent efficiency in various weather conditions.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha FZS FI Double Disc in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,36,500, the Yamaha FZS FI Double Disc represents excellent value with dual disc brake safety, fuel injection efficiency, aggressive styling, and proven Yamaha reliability. It excels at urban commuting with exceptional fuel economy of 50-55 km/l and superior braking performance. The affordable price point makes premium features accessible. However, moderate performance, firm suspension, basic features, lack of ABS, and limited service network make it best suited for first-time buyers, safety-conscious commuters, and urban riders prioritizing braking safety and fuel efficiency in an affordable, stylish package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For superior braking safety with FI efficiency at affordable price</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.2,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha FZS FI Double Disc review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha FZS FI Double Disc (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক হল একটি এন্ট্রি-লেভেল ১৫০সিসি মোটরসাইকেল যেখানে ফুয়েল ইনজেকশন প্রযুক্তি, ডুয়াল ডিস্ক ব্রেক, আক্রমণাত্মক স্ট্রিট স্টাইলিং এবং ব্যবহারিক বৈশিষ্ট্য রয়েছে। ৳২,৩৬,৫০০ টাকায় মূল্যায়িত, এটি সামনে এবং পিছনের ডিস্ক ব্রেক দিয়ে উন্নত ব্রেকিং নিরাপত্তা, আধুনিক এফআই দক্ষতা, স্পোর্টি নেকেড ডিজাইন এবং সাশ্রয়ী প্যাকেজে নিরাপত্তা, স্টাইল এবং পারফরমেন্স খোঁজা শহুরে যাত্রীদের জন্য ইয়ামাহার প্রমাণিত নির্ভরযোগ্যতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ডুয়াল ডিস্ক ব্রেক:</strong> <span class="highlight-value">সামনে এবং পিছনের ডিস্ক দিয়ে উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন সিস্টেম:</strong> <span class="highlight-value">ভাল দক্ষতা এবং পারফরমেন্সের জন্য আধুনিক এফআই</span></li>
<li class="highlight-item"><strong class="highlight-label">সাশ্রয়ী মূল্য:</strong> <span class="highlight-value">প্রিমিয়াম বৈশিষ্ট্য সহ এন্ট্রি-লেভেল মূল্য</span></li>
<li class="highlight-item"><strong class="highlight-label">মাসকুলার স্টাইলিং:</strong> <span class="highlight-value">আক্রমণাত্মক স্ট্রিট নেকেড ডিজাইন আবেদন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">উন্নত ব্রেকিং নিরাপত্তা:</strong> <span class="pro-description">ডুয়াল ডিস্ক ব্রেক সব অবস্থায় চমৎকার থামার শক্তি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন সুবিধা:</strong> <span class="pro-description">ভাল জ্বালানি অর্থনীতি, তাৎক্ষণিক স্টার্ট এবং মসৃণ পাওয়ার ডেলিভারি</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">মাসকুলার এফজেড সিরিজ ডিজাইন তরুণ ক্রেতাদের আকর্ষণ করে</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী প্রিমিয়াম বৈশিষ্ট্য:</strong> <span class="pro-description">এন্ট্রি-লেভেল মূল্যে এফআই এবং ডাবল ডিস্ক পায়</span></li>
<li class="pro-item"><strong class="pro-label">ভাল বিল্ড কোয়ালিটি:</strong> <span class="pro-description">ইয়ামাহার প্রমাণিত নির্ভরযোগ্যতা এবং কম্পোনেন্ট গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার জ্বালানি সাশ্রয়:</strong> <span class="pro-description">এফআই প্রযুক্তি দিয়ে ৫০-৫৫ কিমি/লিটার অর্জন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মাঝারি পারফরমেন্স:</strong> <span class="con-description">১৫০সিসি ইঞ্জিন পর্যাপ্ত কিন্তু রোমাঞ্চকর নয় এমন শক্তি প্রদান করে</span></li>
<li class="con-item"><strong class="con-label">শক্ত সাসপেনশন:</strong> <span class="con-description">রুক্ষ রাস্তা এবং গর্তে কঠোর মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">প্রিমিয়াম মডেলের তুলনায় সীমিত ইলেকট্রনিক বৈশিষ্ট্য</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস বিকল্প নেই:</strong> <span class="con-description">অ্যান্টি-লক ব্রেকিং নিরাপত্তা বৈশিষ্ট্য অনুপস্থিত</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">স্থানীয় ব্র্যান্ডের তুলনায় কম ইয়ামাহা ডিলারশিপ</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">First-time motorcycle buyers</li>
<li class="best-for-item">Safety-conscious urban commuters</li>
<li class="best-for-item">Young riders seeking aggressive styling</li>
<li class="best-for-item">Daily city and highway riders</li>
<li class="best-for-item">Budget-conscious buyers wanting dual disc safety</li>
<li class="best-for-item">Those prioritizing braking performance</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts seeking high power</li>
<li class="not-recommended-item">Riders needing soft suspension comfort</li>
<li class="not-recommended-item">Those wanting advanced electronics</li>
<li class="not-recommended-item">ABS safety feature requirement</li>
<li class="not-recommended-item">Long-distance touring riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৩৬,৫০০ - ডুয়াল ডিস্ক ব্রেক সহ এফআই এর জন্য চমৎকার</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম - ৳৬-৮ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৫-৭ প্রতি কিমি (এফআই দক্ষতা সহ ৫০-৫৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৫,০০০-৮,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- দৈনিক ব্যবহারের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- শক্ত সাসপেনশন সেটআপ</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- আক্রমণাত্মক মাসকুলার ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- মূল্যের জন্য দুর্দান্ত বৈশিষ্ট্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- নির্ভরযোগ্য ইয়ামাহা বিল্ড</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">একক ডিস্কের তুলনায় ডুয়াল ডিস্ক ব্রেক কতটা নিরাপদ?</h3>
<p class="faq-answer">ডুয়াল ডিস্ক ব্রেক উল্লেখযোগ্যভাবে ভাল থামার শক্তি প্রদান করে, বিশেষত জরুরি ব্রেকিং পরিস্থিতি এবং ভেজা অবস্থায়, ব্রেকিং ফোর্স আরও সমানভাবে বিতরণ করে এবং থামার দূরত্ব হ্রাস করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি দৈনিক শহুরে যাতায়াতের জন্য উপযুক্ত?</h3>
<p class="faq-answer">হ্যাঁ, এফজেডএস এফআই ডাবল ডিস্ক ভাল জ্বালানি অর্থনীতি, পরিচালনাযোগ্য আকার, উন্নত ব্রেকিং নিরাপত্তা এবং ব্যবহারিক বৈশিষ্ট্য সহ শহুরে যাতায়াতের জন্য চমৎকার যা শহুরে ট্রাফিকের জন্য আদর্শ করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাস্তব-বিশ্ব অবস্থায় প্রকৃত জ্বালানি অর্থনীতি কত?</h3>
<p class="faq-answer">বাস্তব-বিশ্ব জ্বালানি অর্থনীতি রাইডিং স্টাইল এবং অবস্থার উপর নির্ভর করে ৫০-৫৫ কিমি/লিটার, এফআই প্রযুক্তি বিভিন্ন আবহাওয়া অবস্থায় সামঞ্জস্যপূর্ণ দক্ষতা নিশ্চিত করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,৩৬,৫০০ টাকায় ইয়ামাহা এফজেডএস এফআই ডাবল ডিস্ক ডুয়াল ডিস্ক ব্রেক নিরাপত্তা, ফুয়েল ইনজেকশন দক্ষতা, আক্রমণাত্মক স্টাইলিং এবং প্রমাণিত ইয়ামাহা নির্ভরযোগ্যতা সহ চমৎকার মূল্য উপস্থাপন করে। এটি ৫০-৫৫ কিমি/লিটার ব্যতিক্রমী জ্বালানি অর্থনীতি এবং উন্নত ব্রেকিং পারফরমেন্স সহ শহুরে যাতায়াতে শ্রেষ্ঠ। সাশ্রয়ী মূল্য পয়েন্ট প্রিমিয়াম বৈশিষ্ট্যগুলি অ্যাক্সেসযোগ্য করে তোলে। তবে, মাঝারি পারফরমেন্স, শক্ত সাসপেনশন, মৌলিক বৈশিষ্ট্য, এবিএসের অভাব এবং সীমিত সেবা নেটওয়ার্ক এটিকে প্রথমবার ক্রেতা, নিরাপত্তা-সচেতন যাত্রী এবং সাশ্রয়ী, স্টাইলিশ প্যাকেজে ব্রেকিং নিরাপত্তা এবং জ্বালানি দক্ষতাকে অগ্রাধিকার দেওয়া শহুরে রাইডারদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সাশ্রয়ী মূল্যে এফআই দক্ষতা সহ উন্নত ব্রেকিং নিরাপত্তার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha FZS FI Double Disc already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha FZS FI Double Disc\n")

	return nil
}
