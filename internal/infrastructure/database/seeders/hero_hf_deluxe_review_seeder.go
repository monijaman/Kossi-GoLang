package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroHFDeluxeReviewSeeder handles seeding Hero HF Deluxe product review and translation
type HeroHFDeluxeReviewSeeder struct {
	BaseSeeder
}

// NewHeroHFDeluxeReviewSeeder creates a new HeroHFDeluxeReviewSeeder
func NewHeroHFDeluxeReviewSeeder() *HeroHFDeluxeReviewSeeder {
	return &HeroHFDeluxeReviewSeeder{BaseSeeder: BaseSeeder{name: "Hero HF Deluxe Review"}}
}

// Seed implements the Seeder interface
func (s *HeroHFDeluxeReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero HF Deluxe").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Hero HF Deluxe product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Hero HF Deluxe product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero HF Deluxe already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero HF Deluxe Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Hero HF Deluxe is the most affordable Hero motorcycle at ৳1,30,000, offering legendary reliability and outstanding 70+ km/l mileage. With simple, proven design and India's widest service network, it's the perfect ultra-budget commuter for those prioritizing lowest ownership costs with Hero quality.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero HF Deluxe Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Best Mileage:</strong> <span class="highlight-value">70-75 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Rock-Solid Reliability:</strong> <span class="highlight-value">Proven Hero reliability</span></li>
<li class="highlight-item"><strong class="highlight-label">Lowest Price:</strong> <span class="highlight-value">Most affordable Hero at ৳1,30,000</span></li>
<li class="highlight-item"><strong class="highlight-label">Everywhere Service:</strong> <span class="highlight-value">Hero service in every corner</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero HF Deluxe Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Outstanding Mileage:</strong> <span class="pro-description">70-75 km/l - best in segment</span></li>
<li class="pro-item"><strong class="pro-label">Legendary Reliability:</strong> <span class="pro-description">Hero's bulletproof reliability</span></li>
<li class="pro-item"><strong class="pro-label">Most Affordable Hero:</strong> <span class="pro-description">Cheapest Hero motorcycle</span></li>
<li class="pro-item"><strong class="pro-label">Lowest Running Cost:</strong> <span class="pro-description">Minimal fuel and maintenance costs</span></li>
<li class="pro-item"><strong class="pro-label">Widest Service Network:</strong> <span class="pro-description">Hero service everywhere in Bangladesh</span></li>
<li class="pro-item"><strong class="pro-label">Great Resale Value:</strong> <span class="pro-description">High demand ensures good resale</span></li>
<li class="pro-item"><strong class="pro-label">Simple Maintenance:</strong> <span class="pro-description">Easy and cheap to maintain</span></li>
<li class="pro-item"><strong class="pro-label">Kick Start Option:</strong> <span class="pro-description">Available with kick start variant</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero HF Deluxe Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Basic Features:</strong> <span class="con-description">Absolutely no modern features</span></li>
<li class="con-item"><strong class="con-label">Outdated Looks:</strong> <span class="con-description">Very old-fashioned styling</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">97cc is underpowered</span></li>
<li class="con-item"><strong class="con-label">Basic Build Quality:</strong> <span class="con-description">Uses cheaper materials</span></li>
<li class="con-item"><strong class="con-label">No Self Start:</strong> <span class="con-description">Base variant has only kick start</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Hero HF Deluxe in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extreme budget buyers needing Hero quality</li>
<li class="best-for-item">Those prioritizing lowest running costs</li>
<li class="best-for-item">Daily short-distance commuters</li>
<li class="best-for-item">Buyers needing maximum fuel economy</li>
<li class="best-for-item">Rural areas with limited service options</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Hero HF Deluxe?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious buyers</li>
<li class="not-recommended-item">Those wanting any modern features</li>
<li class="not-recommended-item">Highway or long-distance riding</li>
<li class="not-recommended-item">Performance requirements</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero HF Deluxe Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,25,000 - ৳1,35,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,700 - ৳2,100 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,300-1,500/month for 30 km daily at 72 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳450 - ৳700 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero HF Deluxe Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Outstanding mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Legendary reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Best value in segment</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Underpowered</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Basic comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- No features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very outdated</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Hero HF Deluxe Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Hero HF Deluxe?</h3>
<p class="faq-answer">Hero HF Deluxe delivers exceptional 70-75 km/l mileage.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Hero HF Deluxe?</h3>
<p class="faq-answer">Hero HF Deluxe is priced between ৳1,25,000 to ৳1,35,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Hero HF Deluxe reliable?</h3>
<p class="faq-answer">Yes, extremely reliable with Hero's proven engineering and wide service network.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: HF Deluxe or Honda CD 70?</h3>
<p class="faq-answer">Both excellent; HF Deluxe has better mileage and service network, CD 70 has superior build quality.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero HF Deluxe in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Hero HF Deluxe is the ultimate budget commuter for Bangladesh at ৳1,30,000, offering best-in-class 70-75 km/l mileage, legendary Hero reliability, and lowest ownership costs. While features are minimal and styling is outdated, the exceptional fuel economy, bulletproof reliability, and India's widest service network make it perfect for budget-conscious daily commuters. Best choice for those prioritizing running costs over everything else.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For lowest ownership costs with Hero reliability</span></p>
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
		return fmt.Errorf("error creating Hero HF Deluxe review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero HF Deluxe (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো এইচএফ ডিলাক্স রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হিরো এইচএফ ডিলাক্স ৳১,৩০,০০০ টাকায় সবচেয়ে সাশ্রয়ী হিরো মোটরসাইকেল, যা কিংবদন্তী নির্ভরযোগ্যতা এবং অসাধারণ ৭০+ কিমি/লিটার মাইলেজ প্রদান করে। সহজ, প্রমাণিত ডিজাইন এবং ভারতের সবচেয়ে বিস্তৃত সার্ভিস নেটওয়ার্ক সহ, এটি হিরো কোয়ালিটি সহ সর্বনিম্ন মালিকানা খরচকে অগ্রাধিকার দেওয়া ব্যক্তিদের জন্য নিখুঁত আল্ট্রা-বাজেট কমিউটার।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো এইচএফ ডিলাক্স এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সেরা মাইলেজ:</strong> <span class="highlight-value">৭০-৭৫ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">শক্ত নির্ভরযোগ্যতা:</strong> <span class="highlight-value">প্রমাণিত হিরো নির্ভরযোগ্যতা</span></li>
<li class="highlight-item"><strong class="highlight-label">সর্বনিম্ন মূল্য:</strong> <span class="highlight-value">৳১,৩০,০০০ টাকায় সবচেয়ে সাশ্রয়ী হিরো</span></li>
<li class="highlight-item"><strong class="highlight-label">সর্বত্র সার্ভিস:</strong> <span class="highlight-value">প্রতি কোণে হিরো সার্ভিস</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো এইচএফ ডিলাক্স এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অসাধারণ মাইলেজ:</strong> <span class="pro-description">৭০-৭৫ কিমি/লিটার - সেগমেন্টে সেরা</span></li>
<li class="pro-item"><strong class="pro-label">কিংবদন্তী নির্ভরযোগ্যতা:</strong> <span class="pro-description">হিরোর বুলেটপ্রুফ নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">সবচেয়ে সাশ্রয়ী হিরো:</strong> <span class="pro-description">সবচেয়ে সস্তা হিরো মোটরসাইকেল</span></li>
<li class="pro-item"><strong class="pro-label">সর্বনিম্ন চলমান খরচ:</strong> <span class="pro-description">ন্যূনতম জ্বালানি এবং রক্ষণাবেক্ষণ খরচ</span></li>
<li class="pro-item"><strong class="pro-label">সবচেয়ে বিস্তৃত সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">বাংলাদেশে সর্বত্র হিরো সার্ভিস</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত রিসেল ভ্যালু:</strong> <span class="pro-description">উচ্চ চাহিদা ভালো রিসেল নিশ্চিত করে</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="pro-description">রক্ষণাবেক্ষণ সহজ এবং সস্তা</span></li>
<li class="pro-item"><strong class="pro-label">কিক স্টার্ট অপশন:</strong> <span class="pro-description">কিক স্টার্ট ভ্যারিয়েন্টে উপলব্ধ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো এইচএফ ডিলাক্স এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">খুব বেসিক ফিচার:</strong> <span class="con-description">একেবারে কোনো আধুনিক ফিচার নেই</span></li>
<li class="con-item"><strong class="con-label">পুরানো লুক:</strong> <span class="con-description">খুব পুরানো স্টাইলিং</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">৯৭সিসি কম শক্তিশালী</span></li>
<li class="con-item"><strong class="con-label">বেসিক বিল্ড কোয়ালিটি:</strong> <span class="con-description">সস্তা ম্যাটেরিয়াল ব্যবহার করে</span></li>
<li class="con-item"><strong class="con-label">কোনো সেল্ফ স্টার্ট নেই:</strong> <span class="con-description">বেস ভ্যারিয়েন্টে শুধুমাত্র কিক স্টার্ট আছে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হিরো এইচএফ ডিলাক্স কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Extreme budget buyers needing Hero quality</li>
<li class="best-for-item">Those prioritizing lowest running costs</li>
<li class="best-for-item">Daily short-distance commuters</li>
<li class="best-for-item">Buyers needing maximum fuel economy</li>
<li class="best-for-item">Rural areas with limited service options</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হিরো এইচএফ ডিলাক্স কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious buyers</li>
<li class="not-recommended-item">Those wanting any modern features</li>
<li class="not-recommended-item">Highway or long-distance riding</li>
<li class="not-recommended-item">Performance requirements</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হিরো এইচএফ ডিলাক্স এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,২৫,০০০ - ৳১,৩৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৭০০ - ৳২,১০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৭২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৩০০-১,৫০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৪৫০ - ৳৭০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হিরো এইচএফ ডিলাক্স পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অসাধারণ মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- কিংবদন্তী নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- সেগমেন্টে সেরা ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- কম শক্তিশালী</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- বেসিক আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- বেসিক মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(1.5/5)</span> <span class="rating-note">- কোনো ফিচার নেই</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব পুরানো</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হিরো এইচএফ ডিলাক্স সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হিরো এইচএফ ডিলাক্সের মাইলেজ কত?</h3>
<p class="faq-answer">হিরো এইচএফ ডিলাক্স ব্যতিক্রমী ৭০-৭৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হিরো এইচএফ ডিলাক্সের দাম কত?</h3>
<p class="faq-answer">হিরো এইচএফ ডিলাক্সের দাম ৳১,২৫,০০০ থেকে ৳১,৩৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হিরো এইচএফ ডিলাক্স কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">হ্যাঁ, হিরোর প্রমাণিত ইঞ্জিনিয়ারিং এবং বিস্তৃত সার্ভিস নেটওয়ার্ক সহ অত্যন্ত নির্ভরযোগ্য।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: এইচএফ ডিলাক্স নাকি হোন্ডা সিডি ৭০?</h3>
<p class="faq-answer">উভয়ই চমৎকার; এইচএফ ডিলাক্সের ভালো মাইলেজ এবং সার্ভিস নেটওয়ার্ক আছে, সিডি ৭০-এর উন্নত বিল্ড কোয়ালিটি আছে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হিরো এইচএফ ডিলাক্স কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হিরো এইচএফ ডিলাক্স ৳১,৩০,০০০ টাকায় বাংলাদেশের জন্য চূড়ান্ত বাজেট কমিউটার, যা শ্রেণীতে সেরা ৭০-৭৫ কিমি/লিটার মাইলেজ, কিংবদন্তী হিরো নির্ভরযোগ্যতা এবং সর্বনিম্ন মালিকানা খরচ প্রদান করে। যদিও ফিচার ন্যূনতম এবং স্টাইলিং পুরানো, ব্যতিক্রমী জ্বালানি সাশ্রয়, বুলেটপ্রুফ নির্ভরযোগ্যতা এবং ভারতের সবচেয়ে বিস্তৃত সার্ভিস নেটওয়ার্ক এটিকে বাজেট-সচেতন দৈনিক যাত্রীদের জন্য নিখুঁত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- হিরো নির্ভরযোগ্যতা সহ সর্বনিম্ন মালিকানা খরচের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero HF Deluxe already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero HF Deluxe\n")

	return nil
}
