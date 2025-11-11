package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPlatina100ReviewSeeder handles seeding Bajaj Platina 100 product review and translation
type BajajPlatina100ReviewSeeder struct {
	BaseSeeder
}

// NewBajajPlatina100ReviewSeeder creates a new BajajPlatina100ReviewSeeder
func NewBajajPlatina100ReviewSeeder() *BajajPlatina100ReviewSeeder {
	return &BajajPlatina100ReviewSeeder{BaseSeeder: BaseSeeder{name: "Bajaj Platina 100 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPlatina100ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Platina 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Platina 100 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Platina 100 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Platina 100 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Platina 100 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Platina 100 is a comfortable 102cc commuter bike priced at ৳82,000, offering excellent ride comfort, good fuel efficiency, and Bajaj reliability. With ComforTec technology, long seat, and smooth ride quality, it's ideal for riders prioritizing comfort over performance for daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Platina 100 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ComforTec Technology:</strong> <span class="highlight-value">Enhanced comfort features</span></li>
<li class="highlight-item"><strong class="highlight-label">Longest Seat:</strong> <span class="highlight-value">Extra long comfortable seat</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">68+ km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Smooth Ride:</strong> <span class="highlight-value">Superior suspension comfort</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Platina 100 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Comfort:</strong> <span class="pro-description">ComforTec technology for comfort</span></li>
<li class="pro-item"><strong class="pro-label">Long Seat:</strong> <span class="pro-description">Longest seat in segment</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">68+ km/l excellent efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">Refined 102cc engine</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Bajaj service everywhere</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Budget-friendly pricing</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Convenient electric starting</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Platina 100 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">Only 8.1 HP, underpowered</span></li>
<li class="con-item"><strong class="con-label">Basic Build:</strong> <span class="con-description">Plastic body panels</span></li>
<li class="con-item"><strong class="con-label">Outdated Design:</strong> <span class="con-description">Old-fashioned styling</span></li>
<li class="con-item"><strong class="con-label">Limited Features:</strong> <span class="con-description">No modern features</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Bajaj Platina 100 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Comfort-seeking budget buyers</li>
<li class="best-for-item">Daily long-distance commuters</li>
<li class="best-for-item">Elderly riders wanting comfort</li>
<li class="best-for-item">Rural area transportation</li>
<li class="best-for-item">Family backup bike</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Bajaj Platina 100?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Style-conscious young riders</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Image-conscious buyers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Platina 100 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳78,000 - ৳86,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,900 - ৳2,500 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,300-1,600/month for 30 km daily at 68 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳600 - ৳900 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Platina 100 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent 68+ km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Bajaj reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Great value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Low power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Exceptional comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Outdated styling</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Bajaj Platina 100 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Bajaj Platina 100?</h3>
<p class="faq-answer">Bajaj Platina 100 delivers excellent 68+ km/l with 102cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Bajaj Platina 100?</h3>
<p class="faq-answer">Bajaj Platina 100 is priced between ৳78,000 to ৳86,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Platina 100 comfortable for long rides?</h3>
<p class="faq-answer">Yes, Platina 100 is designed for comfort with ComforTec technology and longest seat.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: CT 100 or Platina 100?</h3>
<p class="faq-answer">Platina 100 offers better comfort and features but costs ৳7,000 more than CT 100.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Platina 100 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.0/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Bajaj Platina 100 at ৳82,000 is the ULTIMATE comfort choice for budget riders, offering exceptional ride comfort with ComforTec technology, longest seat in segment, excellent 68+ km/l mileage, and smooth ride quality. While power (8.1 HP) is basic and design outdated, the superior comfort features, good fuel economy, affordable price, widespread service network, and Bajaj reliability make it perfect for comfort-seeking daily commuters. Ideal for those prioritizing riding comfort and fuel efficiency over performance and style in the budget segment.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For comfort-focused budget commuting</span></p>
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
		return fmt.Errorf("error creating Bajaj Platina 100 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Platina 100 (Rating: 4.0/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ প্লাটিনা ১০০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ প্লাটিনা ১০০ একটি আরামদায়ক ১০২সিসি কমিউটার বাইক যার মূল্য ৳৮২,০০০ টাকা, যা চমৎকার রাইড কমফোর্ট, ভালো জ্বালানি দক্ষতা এবং বাজাজ নির্ভরযোগ্যতা প্রদান করে। কমফোর্টেক প্রযুক্তি, লম্বা সিট এবং মসৃণ রাইড কোয়ালিটি সহ, এটি দৈনিক যাতায়াতের জন্য পারফরম্যান্সের চেয়ে আরামকে অগ্রাধিকার দেওয়া রাইডারদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ প্লাটিনা ১০০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">কমফোর্টেক প্রযুক্তি:</strong> <span class="highlight-value">উন্নত আরাম বৈশিষ্ট্য</span></li>
<li class="highlight-item"><strong class="highlight-label">দীর্ঘতম সিট:</strong> <span class="highlight-value">অতিরিক্ত লম্বা আরামদায়ক সিট</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৬৮+ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">মসৃণ রাইড:</strong> <span class="highlight-value">উচ্চতর সাসপেনশন আরাম</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ প্লাটিনা ১০০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার আরাম:</strong> <span class="pro-description">আরামের জন্য কমফোর্টেক প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">লম্বা সিট:</strong> <span class="pro-description">সেগমেন্টে দীর্ঘতম সিট</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৬৮+ কিমি/লিটার চমৎকার দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">পরিশীলিত ১০২সিসি ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">ব্যাপক সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সর্বত্র বাজাজ সার্ভিস</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">বাজেট-বান্ধব মূল্য</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">সুবিধাজনক ইলেকট্রিক স্টার্টিং</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ প্লাটিনা ১০০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">কম পাওয়ার:</strong> <span class="con-description">শুধু ৮.১ এইচপি, কম শক্তিশালী</span></li>
<li class="con-item"><strong class="con-label">প্রাথমিক বিল্ড:</strong> <span class="con-description">প্লাস্টিকের বডি প্যানেল</span></li>
<li class="con-item"><strong class="con-label">পুরানো ডিজাইন:</strong> <span class="con-description">পুরনো ধাঁচের স্টাইলিং</span></li>
<li class="con-item"><strong class="con-label">সীমিত ফিচার:</strong> <span class="con-description">কোনো আধুনিক ফিচার নেই</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বাজাজ প্লাটিনা ১০০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Comfort-seeking budget buyers</li>
<li class="best-for-item">Daily long-distance commuters</li>
<li class="best-for-item">Elderly riders wanting comfort</li>
<li class="best-for-item">Rural area transportation</li>
<li class="best-for-item">Family backup bike</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বাজাজ প্লাটিনা ১০০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Style-conscious young riders</li>
<li class="not-recommended-item">Those wanting modern features</li>
<li class="not-recommended-item">Image-conscious buyers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বাজাজ প্লাটিনা ১০০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৭৮,০০০ - ৳৮৬,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৯০০ - ৳২,৫০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৬৮ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৩০০-১,৬০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৬০০ - ৳৯০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বাজাজ প্লাটিনা ১০০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার ৬৮+ কিমি/লিটার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- বাজাজ নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- দুর্দান্ত ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- কম পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- ব্যতিক্রমী আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- প্রাথমিক বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- প্রাথমিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- পুরানো স্টাইলিং</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বাজাজ প্লাটিনা ১০০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">বাজাজ প্লাটিনা ১০০ এর মাইলেজ কত?</h3>
<p class="faq-answer">বাজাজ প্লাটিনা ১০০ ১০২সিসি ইঞ্জিন সহ চমৎকার ৬৮+ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাজাজ প্লাটিনা ১০০ এর দাম কত?</h3>
<p class="faq-answer">বাজাজ প্লাটিনা ১০০ এর দাম ৳৭৮,০০০ থেকে ৳৮৬,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">লম্বা রাইডের জন্য প্লাটিনা ১০০ কি আরামদায়ক?</h3>
<p class="faq-answer">হ্যাঁ, প্লাটিনা ১০০ কমফোর্টেক প্রযুক্তি এবং দীর্ঘতম সিট সহ আরামের জন্য ডিজাইন করা।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: সিটি ১০০ নাকি প্লাটিনা ১০০?</h3>
<p class="faq-answer">প্লাটিনা ১০০ ভালো আরাম এবং ফিচার প্রদান করে কিন্তু সিটি ১০০ এর চেয়ে ৳৭,০০০ টাকা বেশি খরচ।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বাজাজ প্লাটিনা ১০০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.0/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৮২,০০০ টাকায় বাজাজ প্লাটিনা ১০০ বাজেট রাইডারদের জন্য চূড়ান্ত আরামের পছন্দ, যা কমফোর্টেক প্রযুক্তি, সেগমেন্টে দীর্ঘতম সিট, চমৎকার ৬৮+ কিমি/লিটার মাইলেজ এবং মসৃণ রাইড কোয়ালিটি সহ ব্যতিক্রমী রাইড কমফোর্ট প্রদান করে। যদিও পাওয়ার (৮.১ এইচপি) প্রাথমিক এবং ডিজাইন পুরানো, উচ্চতর আরাম বৈশিষ্ট্য, ভালো জ্বালানি অর্থনীতি, সাশ্রয়ী মূল্য, ব্যাপক সার্ভিস নেটওয়ার্ক এবং বাজাজ নির্ভরযোগ্যতা এটিকে আরাম-সন্ধানী দৈনিক যাত্রীদের জন্য নিখুঁত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আরাম-কেন্দ্রিক বাজেট যাতায়াতের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Platina 100 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Platina 100\n")

	return nil
}
