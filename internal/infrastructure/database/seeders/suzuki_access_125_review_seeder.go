package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiAccess125ReviewSeeder handles seeding Suzuki Access 125 product review and translation
type SuzukiAccess125ReviewSeeder struct {
	BaseSeeder
}

// NewSuzukiAccess125ReviewSeeder creates a new SuzukiAccess125ReviewSeeder
func NewSuzukiAccess125ReviewSeeder() *SuzukiAccess125ReviewSeeder {
	return &SuzukiAccess125ReviewSeeder{BaseSeeder: BaseSeeder{name: "Suzuki Access 125 Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiAccess125ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Access 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Access 125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Access 125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Access 125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Access 125 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Access 125 is a premium 125cc scooter priced at ৳1,80,000, offering excellent build quality, spacious storage, and comfortable riding. With modern features like LED lights, digital meter, and smooth engine performance, it's ideal for urban commuters who prioritize comfort and practicality.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Access 125 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Premium Build:</strong> <span class="highlight-value">Excellent build quality and finish</span></li>
<li class="highlight-item"><strong class="highlight-label">Large Storage:</strong> <span class="highlight-value">Spacious under-seat storage</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Modern LED headlight and taillight</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Ride:</strong> <span class="highlight-value">Smooth suspension and seating</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Access 125 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Build Quality:</strong> <span class="pro-description">Premium materials and solid construction</span></li>
<li class="pro-item"><strong class="pro-label">Spacious Storage:</strong> <span class="pro-description">21L under-seat storage fits full-face helmet</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">Refined 125cc engine with good power</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seating:</strong> <span class="pro-description">Wide, well-cushioned seat for long rides</span></li>
<li class="pro-item"><strong class="pro-label">LED Lights:</strong> <span class="pro-description">Bright LED headlight and taillight</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">45-50 km/l fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Fully digital instrument cluster</span></li>
<li class="pro-item"><strong class="pro-label">Easy Handling:</strong> <span class="pro-description">Lightweight and nimble in traffic</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Access 125 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">At ৳1,80,000 it's expensive for 125cc scooter</span></li>
<li class="con-item"><strong class="con-label">Limited Top Speed:</strong> <span class="con-description">Not suitable for highway riding</span></li>
<li class="con-item"><strong class="con-label">No Disc Brake Option:</strong> <span class="con-description">Some variants lack front disc brake</span></li>
<li class="con-item"><strong class="con-label">Average Power:</strong> <span class="con-description">Not the most powerful in segment</span></li>
<li class="con-item"><strong class="con-label">Plastic Quality:</strong> <span class="con-description">Some plastic panels feel cheap</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Access 125 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">City commuting with comfort priority</li>
<li class="best-for-item">Users needing large storage space</li>
<li class="best-for-item">Quality-conscious scooter buyers</li>
<li class="best-for-item">Long daily commutes</li>
<li class="best-for-item">Family use and shopping trips</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Access 125?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Highway touring</li>
<li class="not-recommended-item">Extremely tight budget buyers</li>
<li class="not-recommended-item">Riders preferring motorcycle handling</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Access 125 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,75,000 - ৳1,85,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,500 - ৳3,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,000-2,600/month for 30 km daily at 47 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳800 - ৳1,200 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Access 125 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very reliable</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Adequate power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Modern features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Premium look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Access 125 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Suzuki Access 125?</h3>
<p class="faq-answer">Suzuki Access 125 delivers 45-50 km/l in real-world conditions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Suzuki Access 125?</h3>
<p class="faq-answer">Suzuki Access 125 is priced between ৳1,75,000 to ৳1,85,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Suzuki Access 125 good for daily use?</h3>
<p class="faq-answer">Yes, excellent choice with comfortable ride, good storage, and reliable performance.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Access 125 or Honda Dio?</h3>
<p class="faq-answer">Access 125 offers better build quality and storage, while Dio is more affordable.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Access 125 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Suzuki Access 125 is the premium choice in 125cc scooter segment, offering exceptional build quality, spacious storage, and comfortable riding experience at ৳1,80,000. While it's priced higher than competitors, the superior Suzuki engineering, modern LED lights, digital console, and reliable performance justify the investment. Best for daily commuters who value quality and comfort over budget.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For quality-conscious scooter buyers</span></p>
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
		return fmt.Errorf("error creating Suzuki Access 125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Access 125 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি অ্যাক্সেস ১২৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি অ্যাক্সেস ১২৫ একটি প্রিমিয়াম ১২৫সিসি স্কুটার যার মূল্য ৳১,৮০,০০০ টাকা, যা চমৎকার বিল্ড কোয়ালিটি, প্রশস্ত স্টোরেজ এবং আরামদায়ক রাইডিং প্রদান করে। এলইডি লাইট, ডিজিটাল মিটার এবং মসৃণ ইঞ্জিন পারফরম্যান্সের মতো আধুনিক ফিচার সহ, এটি শহুরে যাত্রীদের জন্য আদর্শ যারা আরাম এবং ব্যবহারিকতাকে অগ্রাধিকার দেন।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি অ্যাক্সেস ১২৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম বিল্ড:</strong> <span class="highlight-value">চমৎকার বিল্ড কোয়ালিটি এবং ফিনিশ</span></li>
<li class="highlight-item"><strong class="highlight-label">বড় স্টোরেজ:</strong> <span class="highlight-value">প্রশস্ত সিটের নিচে স্টোরেজ</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">আধুনিক এলইডি হেডলাইট এবং টেইললাইট</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক রাইড:</strong> <span class="highlight-value">মসৃণ সাসপেনশন এবং সিটিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি অ্যাক্সেস ১২৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম ম্যাটেরিয়াল এবং শক্ত নির্মাণ</span></li>
<li class="pro-item"><strong class="pro-label">প্রশস্ত স্টোরেজ:</strong> <span class="pro-description">২১ লিটার সিটের নিচে স্টোরেজ ফুল-ফেস হেলমেট ফিট করে</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">ভালো পাওয়ার সহ পরিমার্জিত ১২৫সিসি ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক সিটিং:</strong> <span class="pro-description">দীর্ঘ রাইডের জন্য চওড়া, ভালো কুশনযুক্ত সিট</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি লাইট:</strong> <span class="pro-description">উজ্জ্বল এলইডি হেডলাইট এবং টেইললাইট</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৪৫-৫০ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">সম্পূর্ণ ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
<li class="pro-item"><strong class="pro-label">সহজ হ্যান্ডলিং:</strong> <span class="pro-description">ট্রাফিকে হালকা এবং চটপটে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি অ্যাক্সেস ১২৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বেশি দাম:</strong> <span class="con-description">৳১,৮০,০০০ টাকায় ১২৫সিসি স্কুটারের জন্য দামি</span></li>
<li class="con-item"><strong class="con-label">সীমিত টপ স্পিড:</strong> <span class="con-description">হাইওয়ে রাইডিংয়ের জন্য উপযুক্ত নয়</span></li>
<li class="con-item"><strong class="con-label">কোনো ডিস্ক ব্রেক অপশন নেই:</strong> <span class="con-description">কিছু ভ্যারিয়েন্টে সামনের ডিস্ক ব্রেক নেই</span></li>
<li class="con-item"><strong class="con-label">গড় পাওয়ার:</strong> <span class="con-description">সেগমেন্টে সবচেয়ে শক্তিশালী নয়</span></li>
<li class="con-item"><strong class="con-label">প্লাস্টিকের মান:</strong> <span class="con-description">কিছু প্লাস্টিক প্যানেল সস্তা মনে হয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি অ্যাক্সেস ১২৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">City commuting with comfort priority</li>
<li class="best-for-item">Users needing large storage space</li>
<li class="best-for-item">Quality-conscious scooter buyers</li>
<li class="best-for-item">Long daily commutes</li>
<li class="best-for-item">Family use and shopping trips</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি অ্যাক্সেস ১২৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Highway touring</li>
<li class="not-recommended-item">Extremely tight budget buyers</li>
<li class="not-recommended-item">Riders preferring motorcycle handling</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি অ্যাক্সেস ১২৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৭৫,০০০ - ৳১,৮৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৫০০ - ৳৩,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,০০০-২,৬০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৮০০ - ৳১,২০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি অ্যাক্সেস ১২৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব নির্ভরযোগ্য</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- পর্যাপ্ত পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আধুনিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- প্রিমিয়াম লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি অ্যাক্সেস ১২৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">সুজুকি অ্যাক্সেস ১২৫-এর মাইলেজ কত?</h3>
<p class="faq-answer">সুজুকি অ্যাক্সেস ১২৫ বাস্তব পরিস্থিতিতে ৪৫-৫০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি অ্যাক্সেস ১২৫-এর দাম কত?</h3>
<p class="faq-answer">সুজুকি অ্যাক্সেস ১২৫-এর দাম ৳১,৭৫,০০০ থেকে ৳১,৮৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য সুজুকি অ্যাক্সেস ১২৫ কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, আরামদায়ক রাইড, ভালো স্টোরেজ এবং নির্ভরযোগ্য পারফরম্যান্স সহ চমৎকার পছন্দ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: অ্যাক্সেস ১২৫ নাকি হোন্ডা ডায়ো?</h3>
<p class="faq-answer">অ্যাক্সেস ১২৫ ভালো বিল্ড কোয়ালিটি এবং স্টোরেজ প্রদান করে, যখন ডায়ো আরও সাশ্রয়ী।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি অ্যাক্সেস ১২৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">সুজুকি অ্যাক্সেস ১২৫ হল ১২৫সিসি স্কুটার সেগমেন্টে প্রিমিয়াম পছন্দ, যা ৳১,৮০,০০০ টাকায় ব্যতিক্রমী বিল্ড কোয়ালিটি, প্রশস্ত স্টোরেজ এবং আরামদায়ক রাইডিং অভিজ্ঞতা প্রদান করে। যদিও এটি প্রতিযোগীদের তুলনায় বেশি দামের, উন্নত সুজুকি ইঞ্জিনিয়ারিং, আধুনিক এলইডি লাইট, ডিজিটাল কনসোল এবং নির্ভরযোগ্য পারফরম্যান্স বিনিয়োগকে ন্যায্যতা দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- মান-সচেতন স্কুটার ক্রেতাদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Access 125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Access 125\n")

	return nil
}
