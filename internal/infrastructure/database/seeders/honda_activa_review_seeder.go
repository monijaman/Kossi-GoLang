package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaActivaReviewSeeder handles seeding Honda Activa product review and translation
type HondaActivaReviewSeeder struct {
	BaseSeeder
}

// NewHondaActivaReviewSeeder creates a new HondaActivaReviewSeeder
func NewHondaActivaReviewSeeder() *HondaActivaReviewSeeder {
	return &HondaActivaReviewSeeder{BaseSeeder: BaseSeeder{name: "Honda Activa Review"}}
}

// Seed implements the Seeder interface
func (s *HondaActivaReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda Activa").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda Activa product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda Activa product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda Activa already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda Activa Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda Activa is India's best-selling scooter priced at ৳1,65,000, offering legendary Honda reliability, excellent build quality, and practical features. With smooth 110cc engine, spacious storage, comfortable ride, and widespread service network, it's the gold standard for family scooters in Bangladesh.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda Activa Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Best-Selling Scooter:</strong> <span class="highlight-value">India's #1 scooter</span></li>
<li class="highlight-item"><strong class="highlight-label">Honda Reliability:</strong> <span class="highlight-value">Legendary dependability</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Build:</strong> <span class="highlight-value">Premium Honda quality</span></li>
<li class="highlight-item"><strong class="highlight-label">Family Friendly:</strong> <span class="highlight-value">Perfect for family use</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda Activa Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Legendary Reliability:</strong> <span class="pro-description">Honda's proven bulletproof reliability</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Build Quality:</strong> <span class="pro-description">Premium materials and solid construction</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">Refined 110cc with great performance</span></li>
<li class="pro-item"><strong class="pro-label">Spacious Storage:</strong> <span class="pro-description">18L under-seat storage</span></li>
<li class="pro-item"><strong class="pro-label">Great Resale Value:</strong> <span class="pro-description">Best resale value in scooter segment</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Honda service everywhere</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Excellent comfort for family</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">45-50 km/l fuel efficiency</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda Activa Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">Premium priced at ৳1,65,000</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">No LED lights or digital console</span></li>
<li class="con-item"><strong class="con-label">Conservative Styling:</strong> <span class="con-description">Simple, not exciting design</span></li>
<li class="con-item"><strong class="con-label">Lower Power:</strong> <span class="con-description">Not as powerful as 125cc scooters</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda Activa in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Family daily transportation</li>
<li class="best-for-item">Grocery shopping and errands</li>
<li class="best-for-item">School/office commuting</li>
<li class="best-for-item">Buyers prioritizing reliability</li>
<li class="best-for-item">Those needing best resale value</li>
<li class="best-for-item">First-time scooter buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda Activa?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Tight budget buyers</li>
<li class="not-recommended-item">Young riders wanting sporty look</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda Activa Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,60,000 - ৳1,70,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,400 - ৳3,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,000-2,400/month for 30 km daily at 47 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳700 - ৳1,100 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda Activa Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Legendary reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent long-term value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Adequate power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Best in class</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Simple family look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda Activa Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Honda Activa?</h3>
<p class="faq-answer">Honda Activa delivers 45-50 km/l in real-world conditions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Honda Activa?</h3>
<p class="faq-answer">Honda Activa is priced between ৳1,60,000 to ৳1,70,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Honda Activa good for family use?</h3>
<p class="faq-answer">Yes, Activa is the best family scooter with legendary reliability and comfort.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Activa or Dio?</h3>
<p class="faq-answer">Activa is more reliable and family-oriented; Dio is sportier and cheaper.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda Activa in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Honda Activa is the ultimate family scooter at ৳1,65,000, offering India's best-selling scooter reliability, excellent Honda build quality, and practical features. While premium priced and lacking modern features, the legendary reliability, widespread service network, best-in-class resale value, and comfortable ride make it worth every penny for family buyers. The gold standard that justifies its premium over competitors.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For best family scooter with legendary reliability</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Honda Activa review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda Activa (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হন্ডা অ্যাকটিভা রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হন্ডা অ্যাকটিভা ভারতের সবচেয়ে বেশি বিক্রি হওয়া স্কুটার যার মূল্য ৳১,৬৫,০০০ টাকা, যা কিংবদন্তী হন্ডা নির্ভরযোগ্যতা, চমৎকার বিল্ড কোয়ালিটি এবং ব্যবহারিক ফিচার প্রদান করে। মসৃণ ১১০সিসি ইঞ্জিন, প্রশস্ত স্টোরেজ, আরামদায়ক রাইড এবং ব্যাপক সার্ভিস নেটওয়ার্ক সহ, এটি বাংলাদেশে পারিবারিক স্কুটারের স্বর্ণ মান।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা অ্যাকটিভা এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সবচেয়ে বেশি বিক্রি হওয়া স্কুটার:</strong> <span class="highlight-value">ভারতের #১ স্কুটার</span></li>
<li class="highlight-item"><strong class="highlight-label">হন্ডা নির্ভরযোগ্যতা:</strong> <span class="highlight-value">কিংবদন্তী নির্ভরযোগ্যতা</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার বিল্ড:</strong> <span class="highlight-value">প্রিমিয়াম হন্ডা মান</span></li>
<li class="highlight-item"><strong class="highlight-label">পরিবার বান্ধব:</strong> <span class="highlight-value">পারিবারিক ব্যবহারের জন্য নিখুঁত</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা অ্যাকটিভা এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">কিংবদন্তী নির্ভরযোগ্যতা:</strong> <span class="pro-description">হন্ডার প্রমাণিত বুলেটপ্রুফ নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম ম্যাটেরিয়াল এবং শক্ত নির্মাণ</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">দুর্দান্ত পারফরম্যান্স সহ পরিমার্জিত ১১০সিসি</span></li>
<li class="pro-item"><strong class="pro-label">প্রশস্ত স্টোরেজ:</strong> <span class="pro-description">১৮ লিটার সিটের নিচে স্টোরেজ</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত রিসেল ভ্যালু:</strong> <span class="pro-description">স্কুটার সেগমেন্টে সেরা রিসেল ভ্যালু</span></li>
<li class="pro-item"><strong class="pro-label">বিস্তৃত সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সর্বত্র হন্ডা সার্ভিস</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">পরিবারের জন্য চমৎকার আরাম</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৪৫-৫০ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা অ্যাকটিভা এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বেশি দাম:</strong> <span class="con-description">৳১,৬৫,০০০ টাকায় প্রিমিয়াম মূল্যের</span></li>
<li class="con-item"><strong class="con-label">বেসিক ফিচার:</strong> <span class="con-description">কোনো এলইডি লাইট বা ডিজিটাল কনসোল নেই</span></li>
<li class="con-item"><strong class="con-label">রক্ষণশীল স্টাইলিং:</strong> <span class="con-description">সাধারণ, উত্তেজনাপূর্ণ ডিজাইন নয়</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">১২৫সিসি স্কুটারের মতো শক্তিশালী নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হন্ডা অ্যাকটিভা কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Family daily transportation</li>
<li class="best-for-item">Grocery shopping and errands</li>
<li class="best-for-item">School/office commuting</li>
<li class="best-for-item">Buyers prioritizing reliability</li>
<li class="best-for-item">Those needing best resale value</li>
<li class="best-for-item">First-time scooter buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হন্ডা অ্যাকটিভা কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Tight budget buyers</li>
<li class="not-recommended-item">Young riders wanting sporty look</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হন্ডা অ্যাকটিভা এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৬০,০০০ - ৳১,৭০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৪০০ - ৳৩,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,০০০-২,৪০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৭০০ - ৳১,১০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হন্ডা অ্যাকটিভা পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- কিংবদন্তী নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার দীর্ঘমেয়াদী মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- পর্যাপ্ত পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- শ্রেণীতে সেরা</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- বেসিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- সাধারণ পারিবারিক লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হন্ডা অ্যাকটিভা সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হন্ডা অ্যাকটিভার মাইলেজ কত?</h3>
<p class="faq-answer">হন্ডা অ্যাকটিভা বাস্তব পরিস্থিতিতে ৪৫-৫০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হন্ডা অ্যাকটিভার দাম কত?</h3>
<p class="faq-answer">হন্ডা অ্যাকটিভার দাম ৳১,৬০,০০০ থেকে ৳১,৭০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">পারিবারিক ব্যবহারের জন্য হন্ডা অ্যাকটিভা কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, অ্যাকটিভা কিংবদন্তী নির্ভরযোগ্যতা এবং আরাম সহ সেরা পারিবারিক স্কুটার।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: অ্যাকটিভা নাকি ডায়ো?</h3>
<p class="faq-answer">অ্যাকটিভা আরও নির্ভরযোগ্য এবং পরিবার-ভিত্তিক; ডায়ো আরও স্পোর্টি এবং সস্তা।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হন্ডা অ্যাকটিভা কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হন্ডা অ্যাকটিভা ৳১,৬৫,০০০ টাকায় চূড়ান্ত পারিবারিক স্কুটার, যা ভারতের সবচেয়ে বেশি বিক্রি হওয়া স্কুটার নির্ভরযোগ্যতা, চমৎকার হন্ডা বিল্ড কোয়ালিটি এবং ব্যবহারিক ফিচার প্রদান করে। যদিও প্রিমিয়াম মূল্যের এবং আধুনিক ফিচারের অভাব, কিংবদন্তী নির্ভরযোগ্যতা, ব্যাপক সার্ভিস নেটওয়ার্ক, শ্রেণীতে সেরা রিসেল ভ্যালু এবং আরামদায়ক রাইড এটিকে পারিবারিক ক্রেতাদের জন্য প্রতিটি টাকার যোগ্য করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- কিংবদন্তী নির্ভরযোগ্যতা সহ সেরা পারিবারিক স্কুটারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda Activa already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda Activa\n")

	return nil
}
