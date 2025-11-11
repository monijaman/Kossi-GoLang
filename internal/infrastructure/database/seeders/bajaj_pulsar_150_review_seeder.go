package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsar150ReviewSeeder handles seeding Bajaj Pulsar 150 product review and translation
type BajajPulsar150ReviewSeeder struct {
	BaseSeeder
}

// NewBajajPulsar150ReviewSeeder creates a new BajajPulsar150ReviewSeeder
func NewBajajPulsar150ReviewSeeder() *BajajPulsar150ReviewSeeder {
	return &BajajPulsar150ReviewSeeder{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar 150 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsar150ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Pulsar 150 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Pulsar 150 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar 150 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar 150 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Pulsar 150 is an iconic 150cc sports commuter priced at ৳2,20,000. With its muscular design, twin-spark engine, and comfortable ergonomics, it has been Bangladesh's favorite sports bike for over a decade. Offers great value with proven reliability and widespread service network.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar 150 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Twin Spark Engine:</strong> <span class="highlight-value">DTS-i technology for efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Iconic Design:</strong> <span class="highlight-value">Muscular sports bike styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Value:</strong> <span class="highlight-value">Feature-packed at competitive price</span></li>
<li class="highlight-item"><strong class="highlight-label">Wide Service Network:</strong> <span class="highlight-value">Bajaj service centers everywhere</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar 150 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">DTS-i Twin Spark Engine:</strong> <span class="pro-description">Efficient twin-plug technology</span></li>
<li class="pro-item"><strong class="pro-label">Great Value for Money:</strong> <span class="pro-description">₳2,20,000 for 150cc with good features</span></li>
<li class="pro-item"><strong class="pro-label">Muscular Styling:</strong> <span class="pro-description">Aggressive sporty design with split seats</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Upright position, comfortable for long rides</span></li>
<li class="pro-item"><strong class="pro-label">Dual Disc Brakes:</strong> <span class="pro-description">Front and rear disc for strong braking</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Bajaj service centers available everywhere</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">45-50 km/l with carburetor</span></li>
<li class="pro-item"><strong class="pro-label">Nitrox Rear Suspension:</strong> <span class="pro-description">Gas-charged rear suspension for comfort</span></li>
<li class="pro-item"><strong class="pro-label">Strong Brand Value:</strong> <span class="pro-description">Pulsar brand has strong reputation</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar 150 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Carburetor Only:</strong> <span class="con-description">No fuel injection, less efficient</span></li>
<li class="con-item"><strong class="con-label">Vibrations:</strong> <span class="con-description">Noticeable vibrations at high RPM</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Plastic quality could be better</span></li>
<li class="con-item"><strong class="con-label">Outdated Features:</strong> <span class="con-description">No LED lights, basic digital console</span></li>
<li class="con-item"><strong class="con-label">Single Channel ABS:</strong> <span class="con-description">ABS only on front wheel</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">150 kg makes it heavy in traffic</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Bajaj Pulsar 150 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Value-conscious sporty bike buyers</li>
<li class="best-for-item">Daily commuters wanting sporty style</li>
<li class="best-for-item">Those needing wide service network</li>
<li class="best-for-item">First-time sports bike buyers</li>
<li class="best-for-item">Riders in remote areas</li>
<li class="best-for-item">Long-distance comfortable touring</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Bajaj Pulsar 150?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Those wanting latest technology</li>
<li class="not-recommended-item">Riders seeking best-in-class handling</li>
<li class="not-recommended-item">Those prioritizing build quality</li>
<li class="not-recommended-item">Lightweight bike preference</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar 150 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,15,000 - ৳2,25,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,800 - ৳3,300 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,000-2,400/month for 30 km daily at 47 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳800 - ৳1,100 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar 150 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Proven Pulsar engine</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent performance</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Iconic muscular</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Bajaj Pulsar 150 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Bajaj Pulsar 150?</h3>
<p class="faq-answer">Bajaj Pulsar 150 delivers 45-50 km/l with DTS-i engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Bajaj Pulsar 150 in Bangladesh?</h3>
<p class="faq-answer">Bajaj Pulsar 150 is priced between ৳2,15,000 to ৳2,25,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Bajaj Pulsar 150 have ABS?</h3>
<p class="faq-answer">Yes, Bajaj Pulsar 150 has single channel ABS on the front wheel.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is DTS-i technology in Pulsar?</h3>
<p class="faq-answer">DTS-i (Digital Twin Spark Ignition) uses two spark plugs for better combustion and efficiency.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Bajaj Pulsar 150 good for daily use?</h3>
<p class="faq-answer">Yes, it's excellent for daily use with comfortable ergonomics and good mileage.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Pulsar 150 or Yamaha FZS?</h3>
<p class="faq-answer">Pulsar 150 offers better value and service network, while FZS has better build quality and features.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Pulsar 150?</h3>
<p class="faq-answer">Bajaj Pulsar 150 can reach approximately 115-120 km/h.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Pulsar 150 have fuel injection?</h3>
<p class="faq-answer">No, current Pulsar 150 uses carburetor, not fuel injection.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar 150 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Bajaj Pulsar 150 remains Bangladesh's most popular sports bike for good reason. At ৳2,20,000, it offers unbeatable value with proven DTS-i engine, comfortable ergonomics, and widespread service network. While it lacks modern features like fuel injection and LED lights, the reliability, comfort, and affordability make it the best all-rounder in the 150cc segment. Perfect for riders wanting sporty style without premium pricing.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For value-conscious sports bike buyers</span></p>
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
		return fmt.Errorf("error creating Bajaj Pulsar 150 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar 150 (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার ১৫০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ পালসার ১৫০ একটি আইকনিক ১৫০সিসি স্পোর্টস কমিউটার যার মূল্য ৳২,২০,০০০ টাকা। এর পেশীবহুল ডিজাইন, টুইন-স্পার্ক ইঞ্জিন এবং আরামদায়ক এর্গোনমিক্স সহ, এটি এক দশকেরও বেশি সময় ধরে বাংলাদেশের প্রিয় স্পোর্টস বাইক হয়ে আছে। প্রমাণিত নির্ভরযোগ্যতা এবং ব্যাপক সার্ভিস নেটওয়ার্ক সহ দুর্দান্ত মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার ১৫০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">টুইন স্পার্ক ইঞ্জিন:</strong> <span class="highlight-value">দক্ষতার জন্য ডিটিএস-আই প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">আইকনিক ডিজাইন:</strong> <span class="highlight-value">পেশীবহুল স্পোর্টস বাইক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মূল্য:</strong> <span class="highlight-value">প্রতিযোগিতামূলক মূল্যে ফিচার-প্যাকড</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যাপক সার্ভিস নেটওয়ার্ক:</strong> <span class="highlight-value">সর্বত্র বাজাজ সার্ভিস সেন্টার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার ১৫০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ডিটিএস-আই টুইন স্পার্ক ইঞ্জিন:</strong> <span class="pro-description">দক্ষ টুইন-প্লাগ প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত ভ্যালু ফর মানি:</strong> <span class="pro-description">ভালো ফিচার সহ ১৫০সিসির জন্য ৳২,২০,০০০ টাকা</span></li>
<li class="pro-item"><strong class="pro-label">পেশীবহুল স্টাইলিং:</strong> <span class="pro-description">স্প্লিট সিট সহ আগ্রাসী স্পোর্টি ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এর্গোনমিক্স:</strong> <span class="pro-description">সোজা অবস্থান, লম্বা রাইডের জন্য আরামদায়ক</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল ডিস্ক ব্রেক:</strong> <span class="pro-description">শক্তিশালী ব্রেকিংয়ের জন্য সামনে এবং পিছনে ডিস্ক</span></li>
<li class="pro-item"><strong class="pro-label">ব্যাপক সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">বাজাজ সার্ভিস সেন্টার সর্বত্র উপলব্ধ</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">কার্বুরেটর সহ ৪৫-৫০ কিমি/লিটার</span></li>
<li class="pro-item"><strong class="pro-label">নাইট্রক্স রিয়ার সাসপেনশন:</strong> <span class="pro-description">আরামের জন্য গ্যাস-চার্জড রিয়ার সাসপেনশন</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী ব্র্যান্ড মূল্য:</strong> <span class="pro-description">পালসার ব্র্যান্ডের শক্তিশালী সুনাম আছে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার ১৫০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">শুধুমাত্র কার্বুরেটর:</strong> <span class="con-description">কোন ফুয়েল ইনজেকশন নেই, কম দক্ষ</span></li>
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">উচ্চ আরপিএমে লক্ষণীয় কম্পন</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">প্লাস্টিকের গুণমান ভালো হতে পারত</span></li>
<li class="con-item"><strong class="con-label">পুরানো ফিচার:</strong> <span class="con-description">কোন এলইডি লাইট নেই, বেসিক ডিজিটাল কনসোল</span></li>
<li class="con-item"><strong class="con-label">সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="con-description">শুধুমাত্র সামনের চাকায় এবিএস</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">১৫০ কেজি ট্রাফিকে ভারী করে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বাজাজ পালসার ১৫০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Value-conscious sporty bike buyers</li>
<li class="best-for-item">Daily commuters wanting sporty style</li>
<li class="best-for-item">Those needing wide service network</li>
<li class="best-for-item">First-time sports bike buyers</li>
<li class="best-for-item">Riders in remote areas</li>
<li class="best-for-item">Long-distance comfortable touring</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বাজাজ পালসার ১৫০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Those wanting latest technology</li>
<li class="not-recommended-item">Riders seeking best-in-class handling</li>
<li class="not-recommended-item">Those prioritizing build quality</li>
<li class="not-recommended-item">Lightweight bike preference</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বাজাজ পালসার ১৫০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,১৫,০০০ - ৳২,২৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৮০০ - ৳৩,৩০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,০০০-২,৪০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৮০০ - ৳১,১০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বাজাজ পালসার ১৫০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- প্রমাণিত পালসার ইঞ্জিন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ভালো পারফরম্যান্স</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় গুণমান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- বেসিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আইকনিক পেশীবহুল</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বাজাজ পালসার ১৫০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">বাজাজ পালসার ১৫০ এর মাইলেজ কত?</h3>
<p class="faq-answer">বাজাজ পালসার ১৫০ ডিটিএস-আই ইঞ্জিন সহ ৪৫-৫০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে বাজাজ পালসার ১৫০ এর দাম কত?</h3>
<p class="faq-answer">বাজাজ পালসার ১৫০ এর দাম ৳২,১৫,০০০ থেকে ৳২,২৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাজাজ পালসার ১৫০এ কি এবিএস আছে?</h3>
<p class="faq-answer">হ্যাঁ, বাজাজ পালসার ১৫০এ সামনের চাকায় সিঙ্গেল চ্যানেল এবিএস আছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">পালসারে ডিটিএস-আই প্রযুক্তি কি?</h3>
<p class="faq-answer">ডিটিএস-আই (ডিজিটাল টুইন স্পার্ক ইগনিশন) ভালো দহন এবং দক্ষতার জন্য দুটি স্পার্ক প্লাগ ব্যবহার করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য বাজাজ পালসার ১৫০ কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, এটি আরামদায়ক এর্গোনমিক্স এবং ভালো মাইলেজ সহ দৈনন্দিন ব্যবহারের জন্য চমৎকার।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: পালসার ১৫০ নাকি ইয়ামাহা এফজেডএস?</h3>
<p class="faq-answer">পালসার ১৫০ ভালো মূল্য এবং সার্ভিস নেটওয়ার্ক দেয়, যেখানে এফজেডএসে ভালো বিল্ড কোয়ালিটি এবং ফিচার আছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">পালসার ১৫০ এর সর্বোচ্চ গতি কত?</h3>
<p class="faq-answer">বাজাজ পালসার ১৫০ প্রায় ১১৫-১২০ কিমি/ঘন্টা পৌঁছাতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">পালসার ১৫০এ কি ফুয়েল ইনজেকশন আছে?</h3>
<p class="faq-answer">না, বর্তমান পালসার ১৫০ কার্বুরেটর ব্যবহার করে, ফুয়েল ইনজেকশন নয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বাজাজ পালসার ১৫০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.1/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">বাজাজ পালসার ১৫০ ভালো কারণেই বাংলাদেশের সবচেয়ে জনপ্রিয় স্পোর্টস বাইক রয়ে গেছে। ৳২,২০,০০০ টাকায়, এটি প্রমাণিত ডিটিএস-আই ইঞ্জিন, আরামদায়ক এর্গোনমিক্স এবং ব্যাপক সার্ভিস নেটওয়ার্ক সহ অপরাজেয় মূল্য প্রদান করে। যদিও এতে ফুয়েল ইনজেকশন এবং এলইডি লাইটের মতো আধুনিক ফিচার নেই, নির্ভরযোগ্যতা, আরাম এবং সাশ্রয়ী মূল্য এটিকে ১৫০সিসি সেগমেন্টে সেরা অল-রাউন্ডার করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- মূল্য-সচেতন স্পোর্টস বাইক ক্রেতাদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar 150 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar 150\n")

	return nil
}
