package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaSalutoRXReviewSeeder handles seeding Yamaha Saluto RX product review and translation
type YamahaSalutoRXReviewSeeder struct {
	BaseSeeder
}

// NewYamahaSalutoRXReviewSeeder creates a new YamahaSalutoRXReviewSeeder
func NewYamahaSalutoRXReviewSeeder() *YamahaSalutoRXReviewSeeder {
	return &YamahaSalutoRXReviewSeeder{BaseSeeder: BaseSeeder{name: "Yamaha Saluto RX Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaSalutoRXReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha Saluto RX").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha Saluto RX product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha Saluto RX product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha Saluto RX already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha Saluto RX Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha Saluto RX is an economical 110cc commuter motorcycle offering good fuel efficiency at ৳1,25,000. With disc brake option and modern styling, it provides practical transportation for budget-conscious riders seeking Yamaha reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha Saluto RX Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Affordable Price:</strong> <span class="highlight-value">Budget-friendly at ৳1,25,000</span></li>
<li class="highlight-item"><strong class="highlight-label">Disc Brake Available:</strong> <span class="highlight-value">Front disc brake option for better stopping</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">52-57 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Design:</strong> <span class="highlight-value">Contemporary styling with graphics</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha Saluto RX Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Affordable Pricing:</strong> <span class="pro-description">At ৳1,25,000 it's budget-friendly</span></li>
<li class="pro-item"><strong class="pro-label">Disc Brake Option:</strong> <span class="pro-description">Front disc brake available unlike many competitors</span></li>
<li class="pro-item"><strong class="pro-label">Decent Fuel Economy:</strong> <span class="pro-description">52-57 km/l keeps costs manageable</span></li>
<li class="pro-item"><strong class="pro-label">Yamaha Reliability:</strong> <span class="pro-description">Proven Japanese engineering</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Self-start convenience included</span></li>
<li class="pro-item"><strong class="pro-label">Tubeless Tyres:</strong> <span class="pro-description">Puncture-resistant tubeless tyres</span></li>
<li class="pro-item"><strong class="pro-label">Good Ground Clearance:</strong> <span class="pro-description">Handles bad roads well</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Modern digital-analog meter</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha Saluto RX Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Plastic quality could be better</span></li>
<li class="con-item"><strong class="con-label">Weak Performance:</strong> <span class="con-description">110cc feels underpowered</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha centers than Honda</span></li>
<li class="con-item"><strong class="con-label">Basic Suspension:</strong> <span class="con-description">Struggles on rough roads</span></li>
<li class="con-item"><strong class="con-label">Lower Resale Value:</strong> <span class="con-description">Not as strong as Honda bikes</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha Saluto RX in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget buyers needing disc brake</li>
<li class="best-for-item">Daily short-distance commuters</li>
<li class="best-for-item">First-time bike buyers</li>
<li class="best-for-item">Students and young professionals</li>
<li class="best-for-item">City-only riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha Saluto RX?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Long-distance highway riders</li>
<li class="not-recommended-item">Those needing maximum reliability</li>
<li class="not-recommended-item">Heavy riders needing more power</li>
<li class="not-recommended-item">Riders in areas with limited Yamaha service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha Saluto RX Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,20,000 - ৳1,30,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,800 - ৳2,300 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,400-1,800/month for 30 km daily at 54 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳700 - ৳1,000 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha Saluto RX Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent Yamaha quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Great value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic 110cc power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Adequate for city</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Disc brake + digital meter</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Modern design</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha Saluto RX Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Yamaha Saluto RX?</h3>
<p class="faq-answer">Yamaha Saluto RX delivers 52-57 km/l in real-world conditions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Yamaha Saluto RX in Bangladesh?</h3>
<p class="faq-answer">Yamaha Saluto RX is priced between ৳1,20,000 to ৳1,30,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Yamaha Saluto RX have disc brake?</h3>
<p class="faq-answer">Yes, Yamaha Saluto RX has front disc brake option available.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Yamaha Saluto RX good for daily use?</h3>
<p class="faq-answer">Yes, it's good for daily city commuting with decent mileage and disc brake.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha Saluto RX in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.0/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Yamaha Saluto RX is a practical budget choice offering disc brake and modern features at ৳1,25,000. While build quality and performance are average, it provides good value for money for city commuters. Best for budget-conscious buyers who prioritize disc brake and don't need highway performance. However, Honda's better service network and resale value make their bikes more attractive for long-term ownership.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget buyers wanting disc brake</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.0,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha Saluto RX review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha Saluto RX (Rating: 4.0/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা সালুটো আরএক্স রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা সালুটো আরএক্স একটি সাশ্রয়ী ১১০সিসি কমিউটার মোটরসাইকেল যা ৳১,২৫,০০০ টাকায় ভালো জ্বালানি দক্ষতা প্রদান করে। ডিস্ক ব্রেক বিকল্প এবং আধুনিক স্টাইলিং সহ, এটি ইয়ামাহা নির্ভরযোগ্যতা খুঁজছেন বাজেট-সচেতন রাইডারদের জন্য ব্যবহারিক পরিবহন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা সালুটো আরএক্স এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সাশ্রয়ী মূল্য:</strong> <span class="highlight-value">৳১,২৫,০০০ টাকায় বাজেট-বান্ধব</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিস্ক ব্রেক উপলব্ধ:</strong> <span class="highlight-value">ভালো স্টপিংয়ের জন্য ফ্রন্ট ডিস্ক ব্রেক বিকল্প</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৫২-৫৭ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক ডিজাইন:</strong> <span class="highlight-value">গ্রাফিক্স সহ সমসাময়িক স্টাইলিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা সালুটো আরএক্স এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">৳১,২৫,০০০ টাকায় এটি বাজেট-বান্ধব</span></li>
<li class="pro-item"><strong class="pro-label">ডিস্ক ব্রেক বিকল্প:</strong> <span class="pro-description">অনেক প্রতিযোগীর বিপরীতে ফ্রন্ট ডিস্ক ব্রেক উপলব্ধ</span></li>
<li class="pro-item"><strong class="pro-label">ভালো জ্বালানি সাশ্রয়:</strong> <span class="pro-description">৫২-৫৭ কিমি/লিটার খরচ ব্যবস্থাপনাযোগ্য রাখে</span></li>
<li class="pro-item"><strong class="pro-label">ইয়ামাহা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত জাপানি ইঞ্জিনিয়ারিং</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">সেল্ফ-স্টার্ট সুবিধা অন্তর্ভুক্ত</span></li>
<li class="pro-item"><strong class="pro-label">টিউবলেস টায়ার:</strong> <span class="pro-description">পাংচার-প্রতিরোধী টিউবলেস টায়ার</span></li>
<li class="pro-item"><strong class="pro-label">ভালো গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="pro-description">খারাপ রাস্তা ভালভাবে সামলায়</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">আধুনিক ডিজিটাল-অ্যানালগ মিটার</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা সালুটো আরএক্স এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">প্লাস্টিকের গুণমান ভালো হতে পারত</span></li>
<li class="con-item"><strong class="con-label">দুর্বল পারফরম্যান্স:</strong> <span class="con-description">১১০সিসি শক্তিহীন মনে হয়</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">হোন্ডার চেয়ে কম ইয়ামাহা সেন্টার</span></li>
<li class="con-item"><strong class="con-label">বেসিক সাসপেনশন:</strong> <span class="con-description">রুক্ষ রাস্তায় সংগ্রাম করে</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">হোন্ডা বাইকের মতো শক্তিশালী নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা সালুটো আরএক্স কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget buyers needing disc brake</li>
<li class="best-for-item">Daily short-distance commuters</li>
<li class="best-for-item">First-time bike buyers</li>
<li class="best-for-item">Students and young professionals</li>
<li class="best-for-item">City-only riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা সালুটো আরএক্স কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Long-distance highway riders</li>
<li class="not-recommended-item">Those needing maximum reliability</li>
<li class="not-recommended-item">Heavy riders needing more power</li>
<li class="not-recommended-item">Riders in areas with limited Yamaha service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা সালুটো আরএক্স এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,২০,০০০ - ৳১,৩০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৮০০ - ৳২,৩০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫৪ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৪০০-১,৮০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৭০০ - ৳১,০০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা সালুটো আরএক্স পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ভালো ইয়ামাহা মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- দুর্দান্ত মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- বেসিক ১১০সিসি পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- শহরের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় গুণমান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ডিস্ক ব্রেক + ডিজিটাল মিটার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- আধুনিক ডিজাইন</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা সালুটো আরএক্স সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা সালুটো আরএক্স এর মাইলেজ কত?</h3>
<p class="faq-answer">ইয়ামাহা সালুটো আরএক্স বাস্তব পরিস্থিতিতে ৫২-৫৭ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে ইয়ামাহা সালুটো আরএক্স এর দাম কত?</h3>
<p class="faq-answer">ইয়ামাহা সালুটো আরএক্স এর দাম ৳১,২০,০০০ থেকে ৳১,৩০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা সালুটো আরএক্সে কি ডিস্ক ব্রেক আছে?</h3>
<p class="faq-answer">হ্যাঁ, ইয়ামাহা সালুটো আরএক্সে ফ্রন্ট ডিস্ক ব্রেক বিকল্প উপলব্ধ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য ইয়ামাহা সালুটো আরএক্স কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, এটি ভালো মাইলেজ এবং ডিস্ক ব্রেক সহ দৈনিক শহর যাতায়াতের জন্য ভালো।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা সালুটো আরএক্স কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.0/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">ইয়ামাহা সালুটো আরএক্স একটি ব্যবহারিক বাজেট পছন্দ যা ৳১,২৫,০০০ টাকায় ডিস্ক ব্রেক এবং আধুনিক ফিচার প্রদান করে। যদিও বিল্ড কোয়ালিটি এবং পারফরম্যান্স গড়, এটি শহর যাত্রীদের জন্য ভালো মূল্য প্রদান করে। বাজেট-সচেতন ক্রেতাদের জন্য সেরা যারা ডিস্ক ব্রেককে অগ্রাধিকার দেন এবং হাইওয়ে পারফরম্যান্সের প্রয়োজন নেই।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ডিস্ক ব্রেক চান বাজেট ক্রেতাদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha Saluto RX already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha Saluto RX\n")

	return nil
}
