package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajDiscoverReviewSeeder handles seeding Bajaj Discover product review and translation
type BajajDiscoverReviewSeeder struct {
	BaseSeeder
}

// NewBajajDiscoverReviewSeeder creates a new BajajDiscoverReviewSeeder
func NewBajajDiscoverReviewSeeder() *BajajDiscoverReviewSeeder {
	return &BajajDiscoverReviewSeeder{BaseSeeder: BaseSeeder{name: "Bajaj Discover Review"}}
}

// Seed implements the Seeder interface
func (s *BajajDiscoverReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Discover").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Discover product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Discover product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Discover already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Discover Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Discover is a practical 125cc commuter motorcycle priced at ৳1,55,000. Known for excellent fuel efficiency and comfortable ergonomics, it's designed for daily commuters who prioritize mileage over style. With Bajaj's wide service network, it offers practical transportation at reasonable cost.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Discover Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">60-65 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Ride:</strong> <span class="highlight-value">Soft seat and upright position</span></li>
<li class="highlight-item"><strong class="highlight-label">Low Maintenance:</strong> <span class="highlight-value">Simple design, cheap parts</span></li>
<li class="highlight-item"><strong class="highlight-label">Wide Service Network:</strong> <span class="highlight-value">Bajaj centers everywhere</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Discover Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Outstanding Mileage:</strong> <span class="pro-description">60-65 km/l among best in segment</span></li>
<li class="pro-item"><strong class="pro-label">Very Comfortable:</strong> <span class="pro-description">Soft seat, upright position for long rides</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">At ৳1,55,000 good value for features</span></li>
<li class="pro-item"><strong class="pro-label">Low Running Cost:</strong> <span class="pro-description">Excellent mileage keeps costs minimal</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Bajaj service available everywhere</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Convenient self-start system</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Discover Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Boring Styling:</strong> <span class="con-description">Very plain and uninspiring design</span></li>
<li class="con-item"><strong class="con-label">Poor Build Quality:</strong> <span class="con-description">Plastic quality is below average</span></li>
<li class="con-item"><strong class="con-label">Weak Performance:</strong> <span class="con-description">125cc engine feels sluggish</span></li>
<li class="con-item"><strong class="con-label">Drum Brakes Only:</strong> <span class="con-description">No disc brake option</span></li>
<li class="con-item"><strong class="con-label">No Modern Features:</strong> <span class="con-description">Lacks any modern technology</span></li>
<li class="con-item"><strong class="con-label">Low Resale Value:</strong> <span class="con-description">Bajaj Discover doesn't hold value</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Bajaj Discover in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Mileage-focused commuters</li>
<li class="best-for-item">Long-distance daily travelers</li>
<li class="best-for-item">Budget-conscious buyers</li>
<li class="best-for-item">Delivery and commercial use</li>
<li class="best-for-item">Those prioritizing comfort over style</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Bajaj Discover?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious riders</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Young buyers seeking sporty look</li>
<li class="not-recommended-item">Those prioritizing resale value</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Discover Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,50,000 - ৳1,60,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,900 - ৳2,400 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,400-1,800/month for 30 km daily at 62 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳600 - ৳900 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Discover Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Outstanding mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good for mileage</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Weak performance</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Poor quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very basic</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Boring design</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Bajaj Discover Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Bajaj Discover?</h3>
<p class="faq-answer">Bajaj Discover delivers outstanding 60-65 km/l.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Bajaj Discover?</h3>
<p class="faq-answer">Bajaj Discover is priced between ৳1,50,000 to ৳1,60,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Bajaj Discover good for long rides?</h3>
<p class="faq-answer">Yes, it's very comfortable with excellent mileage for long commutes.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better for mileage: Discover or CD 70?</h3>
<p class="faq-answer">Discover has better mileage (60-65 km/l) vs CD 70 (60 km/l).</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Discover in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Bajaj Discover is the mileage king for those who don't care about style. At ৳1,55,000 with 60-65 km/l efficiency, it offers lowest running costs in the segment. Very comfortable for long daily commutes. However, boring design, poor build quality, and weak performance make it unsuitable for style-conscious buyers. Best for pure commuting focus, but Honda bikes offer better overall package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For pure mileage-focused commuters</span></p>
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
		return fmt.Errorf("error creating Bajaj Discover review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Discover (Rating: 3.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ ডিসকভার রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ ডিসকভার একটি ব্যবহারিক ১২৫সিসি কমিউটার মোটরসাইকেল যার মূল্য ৳১,৫৫,০০০ টাকা। চমৎকার জ্বালানি দক্ষতা এবং আরামদায়ক এর্গোনমিক্সের জন্য পরিচিত, এটি স্টাইলের চেয়ে মাইলেজকে অগ্রাধিকার দেন দৈনিক যাত্রীদের জন্য ডিজাইন করা হয়েছে। বাজাজের ব্যাপক সার্ভিস নেটওয়ার্ক সহ, এটি যুক্তিসঙ্গত খরচে ব্যবহারিক পরিবহন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ ডিসকভার এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">৬০-৬৫ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক রাইড:</strong> <span class="highlight-value">নরম সিট এবং সোজা অবস্থান</span></li>
<li class="highlight-item"><strong class="highlight-label">কম রক্ষণাবেক্ষণ:</strong> <span class="highlight-value">সাধারণ ডিজাইন, সস্তা পার্টস</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যাপক সার্ভিস নেটওয়ার্ক:</strong> <span class="highlight-value">সর্বত্র বাজাজ সেন্টার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ ডিসকভার এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অসাধারণ মাইলেজ:</strong> <span class="pro-description">৬০-৬৫ কিমি/লিটার সেগমেন্টে সেরাদের মধ্যে</span></li>
<li class="pro-item"><strong class="pro-label">খুব আরামদায়ক:</strong> <span class="pro-description">লম্বা রাইডের জন্য নরম সিট, সোজা অবস্থান</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">ফিচারের জন্য ৳১,৫৫,০০০ টাকায় ভালো মূল্য</span></li>
<li class="pro-item"><strong class="pro-label">কম চলমান খরচ:</strong> <span class="pro-description">চমৎকার মাইলেজ খরচ ন্যূনতম রাখে</span></li>
<li class="pro-item"><strong class="pro-label">ব্যাপক সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সর্বত্র বাজাজ সার্ভিস উপলব্ধ</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">সুবিধাজনক সেল্ফ-স্টার্ট সিস্টেম</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ ডিসকভার এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বিরক্তিকর স্টাইলিং:</strong> <span class="con-description">খুব সাধারণ এবং অনুপ্রেরণাহীন ডিজাইন</span></li>
<li class="con-item"><strong class="con-label">খারাপ বিল্ড কোয়ালিটি:</strong> <span class="con-description">প্লাস্টিকের গুণমান গড়ের নিচে</span></li>
<li class="con-item"><strong class="con-label">দুর্বল পারফরম্যান্স:</strong> <span class="con-description">১২৫সিসি ইঞ্জিন অলস মনে হয়</span></li>
<li class="con-item"><strong class="con-label">শুধুমাত্র ড্রাম ব্রেক:</strong> <span class="con-description">কোন ডিস্ক ব্রেক বিকল্প নেই</span></li>
<li class="con-item"><strong class="con-label">কোন আধুনিক ফিচার নেই:</strong> <span class="con-description">কোন আধুনিক প্রযুক্তি নেই</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">বাজাজ ডিসকভার মূল্য ধরে রাখে না</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বাজাজ ডিসকভার কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Mileage-focused commuters</li>
<li class="best-for-item">Long-distance daily travelers</li>
<li class="best-for-item">Budget-conscious buyers</li>
<li class="best-for-item">Delivery and commercial use</li>
<li class="best-for-item">Those prioritizing comfort over style</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বাজাজ ডিসকভার কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious riders</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Young buyers seeking sporty look</li>
<li class="not-recommended-item">Those prioritizing resale value</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বাজাজ ডিসকভার এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৫০,০০০ - ৳১,৬০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৯০০ - ৳২,৪০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৬২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৪০০-১,৮০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৬০০ - ৳৯০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বাজাজ ডিসকভার পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অসাধারণ মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- মাইলেজের জন্য ভালো</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- দুর্বল পারফরম্যান্স</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- খারাপ মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব বেসিক</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- বিরক্তিকর ডিজাইন</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বাজাজ ডিসকভার সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">বাজাজ ডিসকভারের মাইলেজ কত?</h3>
<p class="faq-answer">বাজাজ ডিসকভার অসাধারণ ৬০-৬৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাজাজ ডিসকভারের দাম কত?</h3>
<p class="faq-answer">বাজাজ ডিসকভারের দাম ৳১,৫০,০০০ থেকে ৳১,৬০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">লম্বা রাইডের জন্য বাজাজ ডিসকভার কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, এটি লম্বা যাতায়াতের জন্য চমৎকার মাইলেজ সহ খুব আরামদায়ক।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">মাইলেজের জন্য কোনটি ভালো: ডিসকভার নাকি সিডি ৭০?</h3>
<p class="faq-answer">সিডি ৭০ (৬০ কিমি/লিটার) এর বিপরীতে ডিসকভারে ভালো মাইলেজ (৬০-৬৫ কিমি/লিটার) আছে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বাজাজ ডিসকভার কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.8/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">বাজাজ ডিসকভার সেই সকলের জন্য মাইলেজ কিং যারা স্টাইল নিয়ে চিন্তা করেন না। ৬০-৬৫ কিমি/লিটার দক্ষতা সহ ৳১,৫৫,০০০ টাকায়, এটি সেগমেন্টে সর্বনিম্ন চলমান খরচ প্রদান করে। দীর্ঘ দৈনিক যাতায়াতের জন্য খুব আরামদায়ক। তবে, বিরক্তিকর ডিজাইন, খারাপ বিল্ড কোয়ালিটি এবং দুর্বল পারফরম্যান্স এটিকে স্টাইল-সচেতন ক্রেতাদের জন্য অনুপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- খাঁটি মাইলেজ-কেন্দ্রিক যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Discover already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Discover\n")

	return nil
}
