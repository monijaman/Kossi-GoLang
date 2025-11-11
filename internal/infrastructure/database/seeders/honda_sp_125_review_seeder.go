package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaSP125ReviewSeeder handles seeding Honda SP 125 product review and translation
type HondaSP125ReviewSeeder struct {
	BaseSeeder
}

// NewHondaSP125ReviewSeeder creates a new HondaSP125ReviewSeeder
func NewHondaSP125ReviewSeeder() *HondaSP125ReviewSeeder {
	return &HondaSP125ReviewSeeder{BaseSeeder: BaseSeeder{name: "Honda SP 125 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaSP125ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda SP 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda SP 125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda SP 125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda SP 125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda SP 125 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda SP 125 is a premium 124cc commuter bike priced at ৳1,60,000, offering sporty design, excellent fuel efficiency, and Honda reliability. With LED lights, digital console, and comfortable ergonomics, it's ideal for daily commuters seeking a perfect balance of style, economy, and dependability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda SP 125 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Premium Styling:</strong> <span class="highlight-value">Sporty and elegant design</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Package:</strong> <span class="highlight-value">LED headlight and position lamp</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">65+ km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Honda Trust:</strong> <span class="highlight-value">Legendary Honda reliability</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda SP 125 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Exceptional Mileage:</strong> <span class="pro-description">65+ km/l best-in-class</span></li>
<li class="pro-item"><strong class="pro-label">Premium Design:</strong> <span class="pro-description">Sporty and premium styling</span></li>
<li class="pro-item"><strong class="pro-label">LED Lighting:</strong> <span class="pro-description">Full LED headlight system</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven Honda build quality</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Excellent riding comfort</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Modern instrument cluster</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Honda's low maintenance costs</span></li>
<li class="pro-item"><strong class="pro-label">Great Resale:</strong> <span class="pro-description">Best resale value in segment</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda SP 125 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Lower Power:</strong> <span class="con-description">Only 10.7 HP, less powerful</span></li>
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">₹15,000 more than Shine</span></li>
<li class="con-item"><strong class="con-label">No Kick Start:</strong> <span class="con-description">Only electric start</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">No ABS or advanced features</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda SP 125 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily office commuters</li>
<li class="best-for-item">Fuel economy seekers</li>
<li class="best-for-item">Honda brand loyalists</li>
<li class="best-for-item">Style-conscious buyers</li>
<li class="best-for-item">Urban family riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda SP 125?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Budget-tight buyers</li>
<li class="not-recommended-item">Those wanting kick start</li>
<li class="not-recommended-item">Power-hungry riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda SP 125 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,55,000 - ৳1,65,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,200 - ৳2,800 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,400-1,800/month for 30 km daily at 65 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳800 - ৳1,200 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda SP 125 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Exceptional 65+ km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Honda excellence</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Adequate power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Honda premium</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Premium styling</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda SP 125 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Honda SP 125?</h3>
<p class="faq-answer">Honda SP 125 delivers exceptional 65+ km/l with 124cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Honda SP 125?</h3>
<p class="faq-answer">Honda SP 125 is priced between ৳1,55,000 to ৳1,65,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is SP 125 better than CB Shine?</h3>
<p class="faq-answer">SP 125 offers better styling, LED lights, and features but costs ৳15,000 more.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does SP 125 have kick start?</h3>
<p class="faq-answer">No, Honda SP 125 comes only with electric start feature.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda SP 125 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Honda SP 125 at ৳1,60,000 is the PERFECT daily commuter for economy-focused riders, offering exceptional 65+ km/l mileage, premium styling with LED lights, legendary Honda reliability, and comfortable ride quality. While power (10.7 HP) is modest and price is premium (₹15,000 more than Shine), the outstanding fuel efficiency, low maintenance costs, best-in-class resale value, and Honda's proven dependability make it ideal for cost-conscious daily commuters. The perfect balance of style, economy, and reliability for urban family riders prioritizing long-term value over raw performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium economy commuter with Honda reliability</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.4,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Honda SP 125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda SP 125 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হন্ডা এসপি ১২৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হন্ডা এসপি ১২৫ একটি প্রিমিয়াম ১২৪সিসি কমিউটার বাইক যার মূল্য ৳১,৬০,০০০ টাকা, যা স্পোর্টি ডিজাইন, চমৎকার জ্বালানি দক্ষতা এবং হন্ডা নির্ভরযোগ্যতা প্রদান করে। এলইডি লাইট, ডিজিটাল কনসোল এবং আরামদায়ক এরগোনমিক্স সহ, এটি স্টাইল, অর্থনীতি এবং নির্ভরযোগ্যতার নিখুঁত ভারসাম্য খোঁজা দৈনিক যাত্রীদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা এসপি ১২৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম স্টাইলিং:</strong> <span class="highlight-value">স্পোর্টি এবং মার্জিত ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি প্যাকেজ:</strong> <span class="highlight-value">এলইডি হেডলাইট এবং পজিশন ল্যাম্প</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">৬৫+ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">হন্ডা বিশ্বাস:</strong> <span class="highlight-value">কিংবদন্তী হন্ডা নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা এসপি ১২৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ব্যতিক্রমী মাইলেজ:</strong> <span class="pro-description">৬৫+ কিমি/লিটার শ্রেণীতে সেরা</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম ডিজাইন:</strong> <span class="pro-description">স্পোর্টি এবং প্রিমিয়াম স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি লাইটিং:</strong> <span class="pro-description">সম্পূর্ণ এলইডি হেডলাইট সিস্টেম</span></li>
<li class="pro-item"><strong class="pro-label">হন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত হন্ডা বিল্ড কোয়ালিটি</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">চমৎকার রাইডিং আরাম</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">আধুনিক ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">হন্ডার কম রক্ষণাবেক্ষণ খরচ</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত রিসেল:</strong> <span class="pro-description">সেগমেন্টে সেরা রিসেল ভ্যালু</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা এসপি ১২৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">কম পাওয়ার:</strong> <span class="con-description">শুধু ১০.৭ এইচপি, কম শক্তিশালী</span></li>
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">শাইনের চেয়ে ৳১৫,০০০ টাকা বেশি</span></li>
<li class="con-item"><strong class="con-label">কোনো কিক স্টার্ট নেই:</strong> <span class="con-description">শুধু ইলেকট্রিক স্টার্ট</span></li>
<li class="con-item"><strong class="con-label">প্রাথমিক ফিচার:</strong> <span class="con-description">কোনো এবিএস বা উন্নত ফিচার নেই</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হন্ডা এসপি ১২৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily office commuters</li>
<li class="best-for-item">Fuel economy seekers</li>
<li class="best-for-item">Honda brand loyalists</li>
<li class="best-for-item">Style-conscious buyers</li>
<li class="best-for-item">Urban family riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হন্ডা এসপি ১২৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Budget-tight buyers</li>
<li class="not-recommended-item">Those wanting kick start</li>
<li class="not-recommended-item">Power-hungry riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হন্ডা এসপি ১২৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৫৫,০০০ - ৳১,৬৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,২০০ - ৳২,৮০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৬৫ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৪০০-১,৮০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৮০০ - ৳১,২০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হন্ডা এসপি ১২৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- ব্যতিক্রমী ৬৫+ কিমি/লিটার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- হন্ডা শ্রেষ্ঠত্ব</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- পর্যাপ্ত পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- হন্ডা প্রিমিয়াম</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- প্রিমিয়াম স্টাইলিং</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হন্ডা এসপি ১২৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হন্ডা এসপি ১২৫ এর মাইলেজ কত?</h3>
<p class="faq-answer">হন্ডা এসপি ১২৫ ১২৪সিসি ইঞ্জিন সহ ব্যতিক্রমী ৬৫+ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হন্ডা এসপি ১২৫ এর দাম কত?</h3>
<p class="faq-answer">হন্ডা এসপি ১২৫ এর দাম ৳১,৫৫,০০০ থেকে ৳১,৬৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এসপি ১২৫ কি সিবি শাইনের চেয়ে ভালো?</h3>
<p class="faq-answer">এসপি ১২৫ ভালো স্টাইলিং, এলইডি লাইট এবং ফিচার প্রদান করে কিন্তু ৳১৫,০০০ টাকা বেশি খরচ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এসপি ১২৫ এ কি কিক স্টার্ট আছে?</h3>
<p class="faq-answer">না, হন্ডা এসপি ১২৫ শুধুমাত্র ইলেকট্রিক স্টার্ট ফিচার সহ আসে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হন্ডা এসপি ১২৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳১,৬০,০০০ টাকায় হন্ডা এসপি ১২৫ অর্থনীতি-কেন্দ্রিক রাইডারদের জন্য নিখুঁত দৈনিক কমিউটার, যা ব্যতিক্রমী ৬৫+ কিমি/লিটার মাইলেজ, এলইডি লাইট সহ প্রিমিয়াম স্টাইলিং, কিংবদন্তী হন্ডা নির্ভরযোগ্যতা এবং আরামদায়ক রাইড কোয়ালিটি প্রদান করে। যদিও পাওয়ার (১০.৭ এইচপি) মাঝারি এবং দাম প্রিমিয়াম (শাইনের চেয়ে ৳১৫,০০০ টাকা বেশি), অসাধারণ জ্বালানি দক্ষতা, কম রক্ষণাবেক্ষণ খরচ, শ্রেণীতে সেরা রিসেল ভ্যালু এবং হন্ডার প্রমাণিত নির্ভরযোগ্যতা এটিকে খরচ-সচেতন দৈনিক যাত্রীদের জন্য আদর্শ করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- হন্ডা নির্ভরযোগ্যতা সহ প্রিমিয়াম ইকোনমি কমিউটারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda SP 125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda SP 125\n")

	return nil
}
