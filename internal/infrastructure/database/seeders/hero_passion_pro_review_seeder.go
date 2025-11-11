package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroPassionProReviewSeeder handles seeding Hero Passion Pro product review and translation
type HeroPassionProReviewSeeder struct {
	BaseSeeder
}

// NewHeroPassionProReviewSeeder creates a new HeroPassionProReviewSeeder
func NewHeroPassionProReviewSeeder() *HeroPassionProReviewSeeder {
	return &HeroPassionProReviewSeeder{BaseSeeder: BaseSeeder{name: "Hero Passion Pro Review"}}
}

// Seed implements the Seeder interface
func (s *HeroPassionProReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Passion Pro").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Hero Passion Pro product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Hero Passion Pro product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero Passion Pro already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero Passion Pro Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Hero Passion Pro is a refined 110cc commuter priced at ৳1,50,000, offering better styling and features than Splendor Plus. With drum-disc brake combo, tubeless tires, digital console, and smooth i3S engine, it balances Hero reliability with modern touches for discerning commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Passion Pro Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Hero Reliability:</strong> <span class="highlight-value">Proven Hero MotoCorp quality</span></li>
<li class="highlight-item"><strong class="highlight-label">i3S Technology:</strong> <span class="highlight-value">Idle start-stop system</span></li>
<li class="highlight-item"><strong class="highlight-label">Disc Brake:</strong> <span class="highlight-value">Front disc for better braking</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Console:</strong> <span class="highlight-value">Modern digital instrument cluster</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Passion Pro Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Reliability:</strong> <span class="pro-description">Hero's proven reliability and service network</span></li>
<li class="pro-item"><strong class="pro-label">Great Fuel Economy:</strong> <span class="pro-description">60-65 km/l with i3S technology</span></li>
<li class="pro-item"><strong class="pro-label">i3S Start-Stop:</strong> <span class="pro-description">Auto engine cut-off at signals saves fuel</span></li>
<li class="pro-item"><strong class="pro-label">Front Disc Brake:</strong> <span class="pro-description">Superior braking performance</span></li>
<li class="pro-item"><strong class="pro-label">Tubeless Tires:</strong> <span class="pro-description">Safer and more convenient</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Modern instrument cluster with more info</span></li>
<li class="pro-item"><strong class="pro-label">Better Styling:</strong> <span class="pro-description">More modern look than Splendor</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Hero service centers everywhere</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Passion Pro Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">₹10,000 more than Splendor Plus</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Still uses cheaper materials</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">110cc is not very powerful</span></li>
<li class="con-item"><strong class="con-label">i3S Not Always Smooth:</strong> <span class="con-description">Start-stop can be jerky initially</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Hero Passion Pro in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters wanting modern features</li>
<li class="best-for-item">Those prioritizing fuel economy with comfort</li>
<li class="best-for-item">Buyers needing Hero reliability with better looks</li>
<li class="best-for-item">City riding with frequent stops (i3S benefit)</li>
<li class="best-for-item">First-time bike buyers wanting trusted brand</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Hero Passion Pro?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Extreme budget buyers (Splendor Plus is cheaper)</li>
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Highway touring requirements</li>
<li class="not-recommended-item">Those wanting premium build quality</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Passion Pro Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,45,000 - ৳1,55,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,000 - ৳2,400 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,500-1,800/month for 30 km daily at 62 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳600 - ৳900 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Passion Pro Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent with i3S</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Hero reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Adequate power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent styling</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Hero Passion Pro Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Hero Passion Pro?</h3>
<p class="faq-answer">Hero Passion Pro delivers 60-65 km/l with i3S technology.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is Hero i3S technology?</h3>
<p class="faq-answer">i3S is idle start-stop system that automatically cuts engine at red lights to save fuel.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Hero Passion Pro?</h3>
<p class="faq-answer">Hero Passion Pro is priced between ৳1,45,000 to ৳1,55,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Passion Pro or Splendor Plus?</h3>
<p class="faq-answer">Passion Pro offers better features (disc brake, i3S, digital console), Splendor Plus is cheaper and simpler.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Passion Pro in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Hero Passion Pro is an excellent modern commuter at ৳1,50,000, offering Hero's legendary reliability with contemporary features like i3S start-stop, disc brake, and digital console. The 60-65 km/l mileage with i3S technology makes it ideal for city riding with frequent stops. While ৳10,000 more than Splendor Plus, the additional features justify the premium for buyers wanting a more refined riding experience with Hero dependability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For modern features with Hero reliability</span></p>
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
		return fmt.Errorf("error creating Hero Passion Pro review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero Passion Pro (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো প্যাশন প্রো রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হিরো প্যাশন প্রো একটি পরিমার্জিত ১১০সিসি কমিউটার যার মূল্য ৳১,৫০,০০০ টাকা, যা স্প্লেন্ডার প্লাসের চেয়ে ভালো স্টাইলিং এবং ফিচার প্রদান করে। ড্রাম-ডিস্ক ব্রেক কম্বো, টিউবলেস টায়ার, ডিজিটাল কনসোল এবং মসৃণ আই৩এস ইঞ্জিন সহ, এটি বিচক্ষণ যাত্রীদের জন্য হিরো নির্ভরযোগ্যতাকে আধুনিক স্পর্শের সাথে ভারসাম্য রাখে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো প্যাশন প্রো এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">হিরো নির্ভরযোগ্যতা:</strong> <span class="highlight-value">প্রমাণিত হিরো মোটোকর্প মান</span></li>
<li class="highlight-item"><strong class="highlight-label">আই৩এস প্রযুক্তি:</strong> <span class="highlight-value">আইডল স্টার্ট-স্টপ সিস্টেম</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিস্ক ব্রেক:</strong> <span class="highlight-value">ভালো ব্রেকিংয়ের জন্য সামনের ডিস্ক</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল কনসোল:</strong> <span class="highlight-value">আধুনিক ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো প্যাশন প্রো এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার নির্ভরযোগ্যতা:</strong> <span class="pro-description">হিরোর প্রমাণিত নির্ভরযোগ্যতা এবং সার্ভিস নেটওয়ার্ক</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত জ্বালানি সাশ্রয়:</strong> <span class="pro-description">আই৩এস প্রযুক্তি সহ ৬০-৬৫ কিমি/লিটার</span></li>
<li class="pro-item"><strong class="pro-label">আই৩এস স্টার্ট-স্টপ:</strong> <span class="pro-description">সিগন্যালে অটো ইঞ্জিন কাট-অফ জ্বালানি বাঁচায়</span></li>
<li class="pro-item"><strong class="pro-label">সামনের ডিস্ক ব্রেক:</strong> <span class="pro-description">উন্নত ব্রেকিং পারফরম্যান্স</span></li>
<li class="pro-item"><strong class="pro-label">টিউবলেস টায়ার:</strong> <span class="pro-description">নিরাপদ এবং আরও সুবিধাজনক</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">আরও তথ্য সহ আধুনিক ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
<li class="pro-item"><strong class="pro-label">ভালো স্টাইলিং:</strong> <span class="pro-description">স্প্লেন্ডারের চেয়ে আধুনিক লুক</span></li>
<li class="pro-item"><strong class="pro-label">বিস্তৃত সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সর্বত্র হিরো সার্ভিস সেন্টার</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো প্যাশন প্রো এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বেশি দাম:</strong> <span class="con-description">স্প্লেন্ডার প্লাসের চেয়ে ৳১০,০০০ টাকা বেশি</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">এখনও সস্তা ম্যাটেরিয়াল ব্যবহার করে</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">১১০সিসি খুব শক্তিশালী নয়</span></li>
<li class="con-item"><strong class="con-label">আই৩এস সবসময় মসৃণ নয়:</strong> <span class="con-description">স্টার্ট-স্টপ প্রাথমিকভাবে ঝাঁকুনিপূর্ণ হতে পারে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হিরো প্যাশন প্রো কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters wanting modern features</li>
<li class="best-for-item">Those prioritizing fuel economy with comfort</li>
<li class="best-for-item">Buyers needing Hero reliability with better looks</li>
<li class="best-for-item">City riding with frequent stops (i3S benefit)</li>
<li class="best-for-item">First-time bike buyers wanting trusted brand</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হিরো প্যাশন প্রো কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Extreme budget buyers (Splendor Plus is cheaper)</li>
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Highway touring requirements</li>
<li class="not-recommended-item">Those wanting premium build quality</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হিরো প্যাশন প্রো এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৪৫,০০০ - ৳১,৫৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,০০০ - ৳২,৪০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৬২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৫০০-১,৮০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৬০০ - ৳৯০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হিরো প্যাশন প্রো পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- আই৩এস সহ চমৎকার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- হিরো নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- পর্যাপ্ত পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ভালো স্টাইলিং</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হিরো প্যাশন প্রো সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হিরো প্যাশন প্রোর মাইলেজ কত?</h3>
<p class="faq-answer">হিরো প্যাশন প্রো আই৩এস প্রযুক্তি সহ ৬০-৬৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হিরো আই৩এস প্রযুক্তি কী?</h3>
<p class="faq-answer">আই৩এস হল আইডল স্টার্ট-স্টপ সিস্টেম যা জ্বালানি বাঁচাতে লাল লাইটে স্বয়ংক্রিয়ভাবে ইঞ্জিন বন্ধ করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হিরো প্যাশন প্রোর দাম কত?</h3>
<p class="faq-answer">হিরো প্যাশন প্রোর দাম ৳১,৪৫,০০০ থেকে ৳১,৫৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: প্যাশন প্রো নাকি স্প্লেন্ডার প্লাস?</h3>
<p class="faq-answer">প্যাশন প্রো ভালো ফিচার (ডিস্ক ব্রেক, আই৩এস, ডিজিটাল কনসোল) প্রদান করে, স্প্লেন্ডার প্লাস সস্তা এবং সহজ।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হিরো প্যাশন প্রো কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.1/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হিরো প্যাশন প্রো ৳১,৫০,০০০ টাকায় একটি চমৎকার আধুনিক কমিউটার, যা আই৩এস স্টার্ট-স্টপ, ডিস্ক ব্রেক এবং ডিজিটাল কনসোলের মতো সমসাময়িক ফিচার সহ হিরোর কিংবদন্তী নির্ভরযোগ্যতা প্রদান করে। আই৩এস প্রযুক্তি সহ ৬০-৬৫ কিমি/লিটার মাইলেজ এটিকে ঘন ঘন থামা সহ শহর রাইডিংয়ের জন্য আদর্শ করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- হিরো নির্ভরযোগ্যতা সহ আধুনিক ফিচারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero Passion Pro already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero Passion Pro\n")

	return nil
}
