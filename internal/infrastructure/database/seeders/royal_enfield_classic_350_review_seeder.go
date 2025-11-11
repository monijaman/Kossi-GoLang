package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RoyalEnfieldClassic350ReviewSeeder handles seeding Royal Enfield Classic 350 product review and translation
type RoyalEnfieldClassic350ReviewSeeder struct {
	BaseSeeder
}

// NewRoyalEnfieldClassic350ReviewSeeder creates a new RoyalEnfieldClassic350ReviewSeeder
func NewRoyalEnfieldClassic350ReviewSeeder() *RoyalEnfieldClassic350ReviewSeeder {
	return &RoyalEnfieldClassic350ReviewSeeder{BaseSeeder: BaseSeeder{name: "Royal Enfield Classic 350 Review"}}
}

// Seed implements the Seeder interface
func (s *RoyalEnfieldClassic350ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Royal Enfield Classic 350").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Royal Enfield Classic 350 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Royal Enfield Classic 350 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Royal Enfield Classic 350 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Royal Enfield Classic 350 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Royal Enfield Classic 350 is an iconic retro cruiser priced at ৳6,50,000, offering timeless vintage styling, thumping 350cc single-cylinder engine, and legendary Royal Enfield character. With premium build, comfortable cruising ability, and strong brand heritage, it's the ultimate choice for enthusiasts seeking classic motorcycle experience.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Royal Enfield Classic 350 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Iconic Styling:</strong> <span class="highlight-value">Timeless retro classic design</span></li>
<li class="highlight-item"><strong class="highlight-label">Thumping Engine:</strong> <span class="highlight-value">Legendary 350cc single-cylinder</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Build:</strong> <span class="highlight-value">Excellent build quality and finish</span></li>
<li class="highlight-item"><strong class="highlight-label">Cruising Comfort:</strong> <span class="highlight-value">Relaxed riding position</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Royal Enfield Classic 350 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Iconic Brand Heritage:</strong> <span class="pro-description">Royal Enfield's legendary status</span></li>
<li class="pro-item"><strong class="pro-label">Timeless Design:</strong> <span class="pro-description">Classic retro styling never goes out of fashion</span></li>
<li class="pro-item"><strong class="pro-label">Thumping Sound:</strong> <span class="pro-description">Distinctive exhaust note</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Build Quality:</strong> <span class="pro-description">Premium materials and solid construction</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Cruising:</strong> <span class="pro-description">Upright riding position, plush seat</span></li>
<li class="pro-item"><strong class="pro-label">Strong Community:</strong> <span class="pro-description">Huge Royal Enfield enthusiast community</span></li>
<li class="pro-item"><strong class="pro-label">Good Torque:</strong> <span class="pro-description">27 Nm torque for relaxed cruising</span></li>
<li class="pro-item"><strong class="pro-label">Great Resale Value:</strong> <span class="pro-description">Royal Enfields hold value well</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Royal Enfield Classic 350 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Expensive:</strong> <span class="con-description">At ৳6,50,000 it's premium priced</span></li>
<li class="con-item"><strong class="con-label">Poor Fuel Economy:</strong> <span class="con-description">Only 30-35 km/l mileage</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">192 kg makes it difficult to handle</span></li>
<li class="con-item"><strong class="con-label">Vibrations:</strong> <span class="con-description">Single-cylinder causes vibrations</span></li>
<li class="con-item"><strong class="con-label">High Maintenance Cost:</strong> <span class="con-description">Expensive parts and service</span></li>
<li class="con-item"><strong class="con-label">Limited Service Centers:</strong> <span class="con-description">RE service not widespread</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Royal Enfield Classic 350 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Enthusiasts seeking classic motorcycle character</li>
<li class="best-for-item">Weekend touring and leisure riding</li>
<li class="best-for-item">Those prioritizing style and heritage</li>
<li class="best-for-item">Riders wanting strong community connection</li>
<li class="best-for-item">Highway cruising at relaxed pace</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Royal Enfield Classic 350?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily city commuters (too heavy)</li>
<li class="not-recommended-item">Those needing fuel efficiency</li>
<li class="not-recommended-item">Performance-oriented riders</li>
<li class="not-recommended-item">First-time bike buyers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Royal Enfield Classic 350 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳6,40,000 - ৳6,60,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳6,000 - ৳8,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳3,500-4,200/month for 30 km daily at 32 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳2,500 - ৳4,000 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Royal Enfield Classic 350 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Poor mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Expensive but unique</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Adequate cruising power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Premium build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Basic features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Iconic timeless design</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Royal Enfield Classic 350 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Royal Enfield Classic 350?</h3>
<p class="faq-answer">Royal Enfield Classic 350 delivers 30-35 km/l in real-world conditions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Classic 350 in Bangladesh?</h3>
<p class="faq-answer">Royal Enfield Classic 350 is priced around ৳6,40,000 to ৳6,60,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Royal Enfield Classic 350 good for daily use?</h3>
<p class="faq-answer">Not ideal for daily city commuting due to heavy weight and poor fuel economy. Better for weekend touring.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Classic 350 worth the price?</h3>
<p class="faq-answer">Yes, for enthusiasts valuing iconic styling, brand heritage, and classic motorcycle experience. Not for practical buyers.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Royal Enfield Classic 350 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Royal Enfield Classic 350 is the ultimate choice for motorcycle enthusiasts at ৳6,50,000, offering iconic retro styling, legendary brand heritage, and authentic classic riding experience. While fuel economy is poor (30-35 km/l), weight is heavy, and maintenance costs are high, the timeless design, thumping engine character, and strong community make it irreplaceable for those seeking style and heritage over practicality. Perfect for weekend touring and leisure riding, not for daily commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For enthusiasts seeking classic heritage and style</span></p>
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
		return fmt.Errorf("error creating Royal Enfield Classic 350 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Royal Enfield Classic 350 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">রয়াল এনফিল্ড ক্লাসিক ৩৫০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">রয়াল এনফিল্ড ক্লাসিক ৩৫০ একটি আইকনিক রেট্রো ক্রুজার যার মূল্য ৳৬,৫০,০০০ টাকা, যা নিরবধি ভিনটেজ স্টাইলিং, থাম্পিং ৩৫০সিসি সিঙ্গেল-সিলিন্ডার ইঞ্জিন এবং কিংবদন্তী রয়াল এনফিল্ড চরিত্র প্রদান করে। প্রিমিয়াম বিল্ড, আরামদায়ক ক্রুজিং ক্ষমতা এবং শক্তিশালী ব্র্যান্ড ঐতিহ্য সহ, এটি ক্লাসিক মোটরসাইকেল অভিজ্ঞতা খোঁজা উৎসাহীদের জন্য চূড়ান্ত পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রয়াল এনফিল্ড ক্লাসিক ৩৫০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">আইকনিক স্টাইলিং:</strong> <span class="highlight-value">নিরবধি রেট্রো ক্লাসিক ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">থাম্পিং ইঞ্জিন:</strong> <span class="highlight-value">কিংবদন্তী ৩৫০সিসি সিঙ্গেল-সিলিন্ডার</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম বিল্ড:</strong> <span class="highlight-value">চমৎকার বিল্ড কোয়ালিটি এবং ফিনিশ</span></li>
<li class="highlight-item"><strong class="highlight-label">ক্রুজিং আরাম:</strong> <span class="highlight-value">স্বাচ্ছন্দ্যময় রাইডিং পজিশন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রয়াল এনফিল্ড ক্লাসিক ৩৫০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আইকনিক ব্র্যান্ড ঐতিহ্য:</strong> <span class="pro-description">রয়াল এনফিল্ডের কিংবদন্তী মর্যাদা</span></li>
<li class="pro-item"><strong class="pro-label">নিরবধি ডিজাইন:</strong> <span class="pro-description">ক্লাসিক রেট্রো স্টাইলিং কখনও ফ্যাশনের বাইরে যায় না</span></li>
<li class="pro-item"><strong class="pro-label">থাম্পিং সাউন্ড:</strong> <span class="pro-description">স্বতন্ত্র এক্সজস্ট নোট</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম ম্যাটেরিয়াল এবং শক্ত নির্মাণ</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক ক্রুজিং:</strong> <span class="pro-description">সোজা রাইডিং পজিশন, আরামদায়ক সিট</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী কমিউনিটি:</strong> <span class="pro-description">বিশাল রয়াল এনফিল্ড উৎসাহী কমিউনিটি</span></li>
<li class="pro-item"><strong class="pro-label">ভালো টর্ক:</strong> <span class="pro-description">স্বাচ্ছন্দ্যময় ক্রুজিংয়ের জন্য ২৭ এনএম টর্ক</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত রিসেল ভ্যালু:</strong> <span class="pro-description">রয়াল এনফিল্ড ভালো মূল্য ধরে রাখে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রয়াল এনফিল্ড ক্লাসিক ৩৫০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">খুব দামি:</strong> <span class="con-description">৳৬,৫০,০০০ টাকায় এটি প্রিমিয়াম মূল্যের</span></li>
<li class="con-item"><strong class="con-label">খারাপ জ্বালানি সাশ্রয়:</strong> <span class="con-description">মাত্র ৩০-৩৫ কিমি/লিটার মাইলেজ</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">১৯২ কেজি এটিকে পরিচালনা করা কঠিন করে</span></li>
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">সিঙ্গেল-সিলিন্ডার কম্পন সৃষ্টি করে</span></li>
<li class="con-item"><strong class="con-label">উচ্চ রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">দামি পার্টস এবং সার্ভিস</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস সেন্টার:</strong> <span class="con-description">আরই সার্ভিস ব্যাপক নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে রয়াল এনফিল্ড ক্লাসিক ৩৫০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Enthusiasts seeking classic motorcycle character</li>
<li class="best-for-item">Weekend touring and leisure riding</li>
<li class="best-for-item">Those prioritizing style and heritage</li>
<li class="best-for-item">Riders wanting strong community connection</li>
<li class="best-for-item">Highway cruising at relaxed pace</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">রয়াল এনফিল্ড ক্লাসিক ৩৫০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily city commuters (too heavy)</li>
<li class="not-recommended-item">Those needing fuel efficiency</li>
<li class="not-recommended-item">Performance-oriented riders</li>
<li class="not-recommended-item">First-time bike buyers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">রয়াল এনফিল্ড ক্লাসিক ৩৫০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৬,৪০,০০০ - ৳৬,৬০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৬,০০০ - ৳৮,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৩২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳৩,৫০০-৪,২০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳২,৫০০ - ৳৪,০০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">রয়াল এনফিল্ড ক্লাসিক ৩৫০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- খারাপ মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- দামি কিন্তু অনন্য</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- পর্যাপ্ত ক্রুজিং পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- প্রিমিয়াম বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- বেসিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- আইকনিক নিরবধি ডিজাইন</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">রয়াল এনফিল্ড ক্লাসিক ৩৫০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">রয়াল এনফিল্ড ক্লাসিক ৩৫০-এর মাইলেজ কত?</h3>
<p class="faq-answer">রয়াল এনফিল্ড ক্লাসিক ৩৫০ বাস্তব পরিস্থিতিতে ৩০-৩৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে ক্লাসিক ৩৫০-এর দাম কত?</h3>
<p class="faq-answer">রয়াল এনফিল্ড ক্লাসিক ৩৫০-এর দাম ৳৬,৪০,০০০ থেকে ৳৬,৬০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য রয়াল এনফিল্ড ক্লাসিক ৩৫০ কি ভালো?</h3>
<p class="faq-answer">ভারী ওজন এবং খারাপ জ্বালানি সাশ্রয়ের কারণে দৈনিক শহর যাতায়াতের জন্য আদর্শ নয়। সপ্তাহান্তে ট্যুরিংয়ের জন্য ভালো।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ক্লাসিক ৩৫০ কি দামের যোগ্য?</h3>
<p class="faq-answer">হ্যাঁ, আইকনিক স্টাইলিং, ব্র্যান্ড ঐতিহ্য এবং ক্লাসিক মোটরসাইকেল অভিজ্ঞতাকে মূল্য দেওয়া উৎসাহীদের জন্য। ব্যবহারিক ক্রেতাদের জন্য নয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে রয়াল এনফিল্ড ক্লাসিক ৩৫০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">রয়াল এনফিল্ড ক্লাসিক ৩৫০ ৳৬,৫০,০০০ টাকায় মোটরসাইকেল উৎসাহীদের জন্য চূড়ান্ত পছন্দ, যা আইকনিক রেট্রো স্টাইলিং, কিংবদন্তী ব্র্যান্ড ঐতিহ্য এবং খাঁটি ক্লাসিক রাইডিং অভিজ্ঞতা প্রদান করে। যদিও জ্বালানি সাশ্রয় খারাপ (৩০-৩৫ কিমি/লিটার), ওজন ভারী এবং রক্ষণাবেক্ষণ খরচ বেশি, নিরবধি ডিজাইন, থাম্পিং ইঞ্জিন চরিত্র এবং শক্তিশালী কমিউনিটি এটিকে ব্যবহারিকতার উপর স্টাইল এবং ঐতিহ্য খোঁজা ব্যক্তিদের জন্য অপরিবর্তনীয় করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্লাসিক ঐতিহ্য এবং স্টাইল খোঁজা উৎসাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Royal Enfield Classic 350 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Royal Enfield Classic 350\n")

	return nil
}
