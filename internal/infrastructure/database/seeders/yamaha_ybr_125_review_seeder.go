package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaYBR125ReviewSeeder handles seeding Yamaha YBR 125 product review and translation
type YamahaYBR125ReviewSeeder struct {
	BaseSeeder
}

// NewYamahaYBR125ReviewSeeder creates a new YamahaYBR125ReviewSeeder
func NewYamahaYBR125ReviewSeeder() *YamahaYBR125ReviewSeeder {
	return &YamahaYBR125ReviewSeeder{BaseSeeder: BaseSeeder{name: "Yamaha YBR 125 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaYBR125ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha YBR 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha YBR 125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha YBR 125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha YBR 125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha YBR 125 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha YBR 125 is a classic 125cc commuter motorcycle priced at ৳1,35,000. Known for its reliable engine and comfortable ergonomics, it's a no-frills practical choice for daily commuting, though outdated styling and average features make it less appealing than newer alternatives.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha YBR 125 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Reliable Engine:</strong> <span class="highlight-value">Proven Yamaha 125cc powerplant</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Comfort:</strong> <span class="highlight-value">Comfortable seat for long rides</span></li>
<li class="highlight-item"><strong class="highlight-label">Decent Mileage:</strong> <span class="highlight-value">48-52 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Easy Maintenance:</strong> <span class="highlight-value">Simple design, easy to service</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha YBR 125 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Reliable Yamaha Engine:</strong> <span class="pro-description">Proven 125cc engine with good reliability</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Upright seating and soft seat padding</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">48-52 km/l keeps running costs low</span></li>
<li class="pro-item"><strong class="pro-label">Easy Maintenance:</strong> <span class="pro-description">Simple carburetor design, cheap parts</span></li>
<li class="pro-item"><strong class="pro-label">Good Ground Clearance:</strong> <span class="pro-description">Handles rough roads adequately</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight:</strong> <span class="pro-description">Easy to handle in traffic</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Convenient self-start system</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha YBR 125 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Outdated Styling:</strong> <span class="con-description">Old-fashioned design looks dated</span></li>
<li class="con-item"><strong class="con-label">Drum Brakes Only:</strong> <span class="con-description">No disc brake option available</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">No modern features or technology</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Plastic quality could be better</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha centers than Honda</span></li>
<li class="con-item"><strong class="con-label">Lower Resale Value:</strong> <span class="con-description">Not as strong as Honda bikes</span></li>
<li class="con-item"><strong class="con-label">Weak Performance:</strong> <span class="con-description">Underpowered for highway riding</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha YBR 125 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget-conscious commuters</li>
<li class="best-for-item">Riders wanting simple reliable bike</li>
<li class="best-for-item">City-only daily use</li>
<li class="best-for-item">First-time buyers</li>
<li class="best-for-item">Those not caring about style</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha YBR 125?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious riders</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Highway and long-distance riding</li>
<li class="not-recommended-item">Riders needing disc brake</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha YBR 125 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,30,000 - ৳1,40,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,000 - ৳2,500 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,500-1,900/month for 30 km daily at 50 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳600 - ৳900 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha YBR 125 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Reliable Yamaha</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Overpriced for features</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Weak performance</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Comfortable seat</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very basic</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Outdated design</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha YBR 125 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Yamaha YBR 125?</h3>
<p class="faq-answer">Yamaha YBR 125 delivers 48-52 km/l in real-world conditions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Yamaha YBR 125 in Bangladesh?</h3>
<p class="faq-answer">Yamaha YBR 125 is priced between ৳1,30,000 to ৳1,40,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Yamaha YBR 125 have disc brake?</h3>
<p class="faq-answer">No, Yamaha YBR 125 comes with drum brakes only.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Yamaha YBR 125 good for daily use?</h3>
<p class="faq-answer">Yes, it's reliable for city commuting but lacks modern features.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha YBR 125 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Yamaha YBR 125 is a reliable but outdated 125cc commuter. At ৳1,35,000, it's overpriced considering the drum brakes, basic features, and old styling. While the engine is proven and comfortable, Honda CD 70 or Livo offer better value. Only suitable for riders who specifically want Yamaha reliability and don't care about modern features or styling.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For Yamaha loyalists not caring about looks</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.8,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha YBR 125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha YBR 125 (Rating: 3.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা ওয়াইবিআর ১২৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা ওয়াইবিআর ১২৫ একটি ক্লাসিক ১২৫সিসি কমিউটার মোটরসাইকেল যার মূল্য ৳১,৩৫,০০০ টাকা। এর নির্ভরযোগ্য ইঞ্জিন এবং আরামদায়ক এর্গোনমিক্সের জন্য পরিচিত, এটি দৈনিক যাতায়াতের জন্য একটি বাস্তবসম্মত পছন্দ, যদিও পুরানো স্টাইলিং এবং গড় ফিচার এটিকে নতুন বিকল্পগুলির চেয়ে কম আকর্ষণীয় করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা ওয়াইবিআর ১২৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="highlight-value">প্রমাণিত ইয়ামাহা ১২৫সিসি পাওয়ারপ্ল্যান্ট</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো আরাম:</strong> <span class="highlight-value">লম্বা রাইডের জন্য আরামদায়ক সিট</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৪৮-৫২ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="highlight-value">সাধারণ ডিজাইন, সার্ভিস করা সহজ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা ওয়াইবিআর ১২৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইয়ামাহা ইঞ্জিন:</strong> <span class="pro-description">ভালো নির্ভরযোগ্যতা সহ প্রমাণিত ১২৫সিসি ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এর্গোনমিক্স:</strong> <span class="pro-description">সোজা বসার অবস্থান এবং নরম সিট প্যাডিং</span></li>
<li class="pro-item"><strong class="pro-label">ভালো জ্বালানি সাশ্রয়:</strong> <span class="pro-description">৪৮-৫২ কিমি/লিটার চলমান খরচ কম রাখে</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সাধারণ কার্বুরেটর ডিজাইন, সস্তা পার্টস</span></li>
<li class="pro-item"><strong class="pro-label">ভালো গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="pro-description">রুক্ষ রাস্তা পর্যাপ্তভাবে সামলায়</span></li>
<li class="pro-item"><strong class="pro-label">হালকা ওজন:</strong> <span class="pro-description">ট্রাফিকে সামলানো সহজ</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">সুবিধাজনক সেল্ফ-স্টার্ট সিস্টেম</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা ওয়াইবিআর ১২৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">পুরানো স্টাইলিং:</strong> <span class="con-description">পুরানো ফ্যাশন ডিজাইন পুরানো দেখায়</span></li>
<li class="con-item"><strong class="con-label">শুধুমাত্র ড্রাম ব্রেক:</strong> <span class="con-description">কোন ডিস্ক ব্রেক বিকল্প নেই</span></li>
<li class="con-item"><strong class="con-label">বেসিক ফিচার:</strong> <span class="con-description">কোন আধুনিক ফিচার বা প্রযুক্তি নেই</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">প্লাস্টিকের গুণমান ভালো হতে পারত</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">হোন্ডার চেয়ে কম ইয়ামাহা সেন্টার</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">হোন্ডা বাইকের মতো শক্তিশালী নয়</span></li>
<li class="con-item"><strong class="con-label">দুর্বল পারফরম্যান্স:</strong> <span class="con-description">হাইওয়ে রাইডিংয়ের জন্য শক্তিহীন</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা ওয়াইবিআর ১২৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget-conscious commuters</li>
<li class="best-for-item">Riders wanting simple reliable bike</li>
<li class="best-for-item">City-only daily use</li>
<li class="best-for-item">First-time buyers</li>
<li class="best-for-item">Those not caring about style</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা ওয়াইবিআর ১২৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious riders</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Highway and long-distance riding</li>
<li class="not-recommended-item">Riders needing disc brake</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা ওয়াইবিআর ১২৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৩০,০০০ - ৳১,৪০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,০০০ - ৳২,৫০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫০ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৫০০-১,৯০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৬০০ - ৳৯০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা ওয়াইবিআর ১২৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- নির্ভরযোগ্য ইয়ামাহা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- ফিচারের জন্য বেশি দাম</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- দুর্বল পারফরম্যান্স</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আরামদায়ক সিট</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় গুণমান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব বেসিক</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- পুরানো ডিজাইন</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা ওয়াইবিআর ১২৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা ওয়াইবিআর ১২৫ এর মাইলেজ কত?</h3>
<p class="faq-answer">ইয়ামাহা ওয়াইবিআর ১২৫ বাস্তব পরিস্থিতিতে ৪৮-৫২ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে ইয়ামাহা ওয়াইবিআর ১২৫ এর দাম কত?</h3>
<p class="faq-answer">ইয়ামাহা ওয়াইবিআর ১২৫ এর দাম ৳১,৩০,০০০ থেকে ৳১,৪০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা ওয়াইবিআর ১২৫এ কি ডিস্ক ব্রেক আছে?</h3>
<p class="faq-answer">না, ইয়ামাহা ওয়াইবিআর ১২৫ শুধুমাত্র ড্রাম ব্রেক সহ আসে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য ইয়ামাহা ওয়াইবিআর ১২৫ কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, এটি শহর যাতায়াতের জন্য নির্ভরযোগ্য কিন্তু আধুনিক ফিচার নেই।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা ওয়াইবিআর ১২৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.8/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">ইয়ামাহা ওয়াইবিআর ১২৫ একটি নির্ভরযোগ্য কিন্তু পুরানো ১২৫সিসি কমিউটার। ৳১,৩৫,০০০ টাকায়, ড্রাম ব্রেক, বেসিক ফিচার এবং পুরানো স্টাইলিং বিবেচনা করে এটি বেশি দাম। যদিও ইঞ্জিন প্রমাণিত এবং আরামদায়ক, হোন্ডা সিডি ৭০ বা লিভো ভালো মূল্য দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চেহারা নিয়ে চিন্তা করেন না এমন ইয়ামাহা অনুরাগীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha YBR 125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha YBR 125\n")

	return nil
}
