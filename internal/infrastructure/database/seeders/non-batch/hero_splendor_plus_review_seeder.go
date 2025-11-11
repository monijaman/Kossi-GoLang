package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroSplendorPlusReviewSeeder handles seeding Hero Splendor Plus product review and translation
type HeroSplendorPlusReviewSeeder struct {
	BaseSeeder
}

// NewHeroSplendorPlusReviewSeeder creates a new HeroSplendorPlusReviewSeeder
func NewHeroSplendorPlusReviewSeeder() *HeroSplendorPlusReviewSeeder {
	return &HeroSplendorPlusReviewSeeder{BaseSeeder: BaseSeeder{name: "Hero Splendor Plus Review"}}
}

// Seed implements the Seeder interface
func (s *HeroSplendorPlusReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Splendor Plus").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Hero Splendor Plus product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Hero Splendor Plus product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero Splendor Plus already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero Splendor Plus Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Hero Splendor Plus is India's best-selling commuter motorcycle priced at ৳1,40,000, offering legendary reliability, excellent fuel efficiency around 70 km/l, and widespread service network. With simple design, low maintenance costs, and Hero's proven technology, it's ideal for budget-conscious daily commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Splendor Plus Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Legendary Reliability:</strong> <span class="highlight-value">India's most trusted commuter</span></li>
<li class="highlight-item"><strong class="highlight-label">Outstanding Mileage:</strong> <span class="highlight-value">65-70 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Wide Service Network:</strong> <span class="highlight-value">Hero service everywhere</span></li>
<li class="highlight-item"><strong class="highlight-label">Low Running Cost:</strong> <span class="highlight-value">Cheapest to maintain</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Splendor Plus Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Exceptional Reliability:</strong> <span class="pro-description">Bulletproof reliability, rarely breaks down</span></li>
<li class="pro-item"><strong class="pro-label">Best-in-Class Mileage:</strong> <span class="pro-description">70 km/l is segment-leading</span></li>
<li class="pro-item"><strong class="pro-label">Extensive Service Network:</strong> <span class="pro-description">Hero service centers everywhere</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance Cost:</strong> <span class="pro-description">Cheapest parts and service</span></li>
<li class="pro-item"><strong class="pro-label">Great Resale Value:</strong> <span class="pro-description">Excellent resale due to popularity</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Soft suspension, good seat</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Convenient self-start available</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Splendor Plus Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Basic Features:</strong> <span class="con-description">No modern features</span></li>
<li class="con-item"><strong class="con-label">Outdated Styling:</strong> <span class="con-description">Looks very old-fashioned</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">97cc is underpowered</span></li>
<li class="con-item"><strong class="con-label">Poor Build Quality:</strong> <span class="con-description">Cheap materials used</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Hero Splendor Plus in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters prioritizing fuel economy</li>
<li class="best-for-item">Those wanting lowest running costs</li>
<li class="best-for-item">Buyers needing maximum reliability</li>
<li class="best-for-item">Areas with limited service availability</li>
<li class="best-for-item">Long-term ownership focused buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Hero Splendor Plus?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious buyers</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Highway riding requirements</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Splendor Plus Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,35,000 - ৳1,45,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,800 - ৳2,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,400-1,600/month for 30 km daily at 68 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳500 - ৳800 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Splendor Plus Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Best mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Legendary reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Underpowered</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very basic</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Outdated</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Hero Splendor Plus Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Hero Splendor Plus?</h3>
<p class="faq-answer">Hero Splendor Plus delivers exceptional 65-70 km/l mileage.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Hero Splendor Plus?</h3>
<p class="faq-answer">Hero Splendor Plus is priced between ৳1,35,000 to ৳1,45,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Hero Splendor Plus reliable?</h3>
<p class="faq-answer">Yes, legendary reliability - India's most trusted commuter bike.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Splendor Plus or Honda CD 70?</h3>
<p class="faq-answer">Both are excellent; Splendor offers better mileage, CD 70 has superior build quality.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Splendor Plus in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Hero Splendor Plus is the ultimate practical commuter at ৳1,40,000, offering legendary reliability, best-in-class 70 km/l mileage, and lowest running costs. While styling is outdated and features are basic, the exceptional reliability, widespread Hero service network, and excellent resale value make it perfect for cost-conscious daily commuters. Best choice for those prioritizing fuel economy and reliability over style.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For maximum reliability and fuel economy</span></p>
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
		return fmt.Errorf("error creating Hero Splendor Plus review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero Splendor Plus (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো স্প্লেন্ডার প্লাস রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হিরো স্প্লেন্ডার প্লাস ভারতের সবচেয়ে বেশি বিক্রি হওয়া কমিউটার মোটরসাইকেল যার মূল্য ৳১,৪০,০০০ টাকা, যা কিংবদন্তী নির্ভরযোগ্যতা, প্রায় ৭০ কিমি/লিটার চমৎকার জ্বালানি দক্ষতা এবং ব্যাপক সার্ভিস নেটওয়ার্ক প্রদান করে। সাধারণ ডিজাইন, কম রক্ষণাবেক্ষণ খরচ এবং হিরোর প্রমাণিত প্রযুক্তি সহ, এটি বাজেট-সচেতন দৈনিক যাত্রীদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো স্প্লেন্ডার প্লাস এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">কিংবদন্তী নির্ভরযোগ্যতা:</strong> <span class="highlight-value">ভারতের সবচেয়ে বিশ্বস্ত কমিউটার</span></li>
<li class="highlight-item"><strong class="highlight-label">অসাধারণ মাইলেজ:</strong> <span class="highlight-value">৬৫-৭০ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">বিস্তৃত সার্ভিস নেটওয়ার্ক:</strong> <span class="highlight-value">সর্বত্র হিরো সার্ভিস</span></li>
<li class="highlight-item"><strong class="highlight-label">কম চলমান খরচ:</strong> <span class="highlight-value">রক্ষণাবেক্ষণ সবচেয়ে সস্তা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো স্প্লেন্ডার প্লাস এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ব্যতিক্রমী নির্ভরযোগ্যতা:</strong> <span class="pro-description">বুলেটপ্রুফ নির্ভরযোগ্যতা, কদাচিৎ নষ্ট হয়</span></li>
<li class="pro-item"><strong class="pro-label">শ্রেণীতে সেরা মাইলেজ:</strong> <span class="pro-description">৭০ কিমি/লিটার সেগমেন্ট-লিডিং</span></li>
<li class="pro-item"><strong class="pro-label">বিস্তৃত সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সর্বত্র হিরো সার্ভিস সেন্টার</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ খরচ:</strong> <span class="pro-description">সবচেয়ে সস্তা পার্টস এবং সার্ভিস</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত রিসেল ভ্যালু:</strong> <span class="pro-description">জনপ্রিয়তার কারণে চমৎকার রিসেল</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">নরম সাসপেনশন, ভালো সিট</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">সুবিধাজনক সেল্ফ-স্টার্ট উপলব্ধ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো স্প্লেন্ডার প্লাস এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">খুব বেসিক ফিচার:</strong> <span class="con-description">কোনো আধুনিক ফিচার নেই</span></li>
<li class="con-item"><strong class="con-label">পুরানো স্টাইলিং:</strong> <span class="con-description">খুব পুরানো দেখায়</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">৯৭সিসি কম শক্তিশালী</span></li>
<li class="con-item"><strong class="con-label">খারাপ বিল্ড কোয়ালিটি:</strong> <span class="con-description">সস্তা ম্যাটেরিয়াল ব্যবহৃত</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হিরো স্প্লেন্ডার প্লাস কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters prioritizing fuel economy</li>
<li class="best-for-item">Those wanting lowest running costs</li>
<li class="best-for-item">Buyers needing maximum reliability</li>
<li class="best-for-item">Areas with limited service availability</li>
<li class="best-for-item">Long-term ownership focused buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হিরো স্প্লেন্ডার প্লাস কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious buyers</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Highway riding requirements</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হিরো স্প্লেন্ডার প্লাস এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৩৫,০০০ - ৳১,৪৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৮০০ - ৳২,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৬৮ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৪০০-১,৬০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৫০০ - ৳৮০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হিরো স্প্লেন্ডার প্লাস পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- সেরা মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- কিংবদন্তী নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- কম শক্তিশালী</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব বেসিক</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- পুরানো</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হিরো স্প্লেন্ডার প্লাস সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হিরো স্প্লেন্ডার প্লাসের মাইলেজ কত?</h3>
<p class="faq-answer">হিরো স্প্লেন্ডার প্লাস ব্যতিক্রমী ৬৫-৭০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হিরো স্প্লেন্ডার প্লাসের দাম কত?</h3>
<p class="faq-answer">হিরো স্প্লেন্ডার প্লাসের দাম ৳১,৩৫,০০০ থেকে ৳১,৪৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হিরো স্প্লেন্ডার প্লাস কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">হ্যাঁ, কিংবদন্তী নির্ভরযোগ্যতা - ভারতের সবচেয়ে বিশ্বস্ত কমিউটার বাইক।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: স্প্লেন্ডার প্লাস নাকি হোন্ডা সিডি ৭০?</h3>
<p class="faq-answer">উভয়ই চমৎকার; স্প্লেন্ডার ভালো মাইলেজ দেয়, সিডি ৭০-এ উন্নত বিল্ড কোয়ালিটি আছে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হিরো স্প্লেন্ডার প্লাস কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হিরো স্প্লেন্ডার প্লাস ৳১,৪০,০০০ টাকায় চূড়ান্ত ব্যবহারিক কমিউটার, যা কিংবদন্তী নির্ভরযোগ্যতা, শ্রেণীতে সেরা ৭০ কিমি/লিটার মাইলেজ এবং সর্বনিম্ন চলমান খরচ প্রদান করে। যদিও স্টাইলিং পুরানো এবং ফিচার বেসিক, ব্যতিক্রমী নির্ভরযোগ্যতা, ব্যাপক হিরো সার্ভিস নেটওয়ার্ক এবং চমৎকার রিসেল ভ্যালু এটিকে খরচ-সচেতন দৈনিক যাত্রীদের জন্য নিখুঁত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সর্বোচ্চ নির্ভরযোগ্যতা এবং জ্বালানি সাশ্রয়ের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero Splendor Plus already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero Splendor Plus\n")

	return nil
}
