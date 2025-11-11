package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSApacheRTR160ReviewSeeder handles seeding TVS Apache RTR 160 product review and translation
type TVSApacheRTR160ReviewSeeder struct {
	BaseSeeder
}

// NewTVSApacheRTR160ReviewSeeder creates a new TVSApacheRTR160ReviewSeeder
func NewTVSApacheRTR160ReviewSeeder() *TVSApacheRTR160ReviewSeeder {
	return &TVSApacheRTR160ReviewSeeder{BaseSeeder: BaseSeeder{name: "TVS Apache RTR 160 Review"}}
}

// Seed implements the Seeder interface
func (s *TVSApacheRTR160ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 160").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("TVS Apache RTR 160 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding TVS Apache RTR 160 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Apache RTR 160 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Apache RTR 160 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The TVS Apache RTR 160 is a performance-focused 160cc motorcycle priced at ৳2,10,000. With race-derived technology, LED lighting, and sporty handling, it offers thrilling performance for enthusiasts. However, limited service network and lower brand value make it a niche choice in Bangladesh market.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 160 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Race Technology:</strong> <span class="highlight-value">Race-derived engine and suspension</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Full LED headlamp and DRL</span></li>
<li class="highlight-item"><strong class="highlight-label">Sporty Performance:</strong> <span class="highlight-value">160cc with race-tuned setup</span></li>
<li class="highlight-item"><strong class="highlight-label">Competitive Price:</strong> <span class="highlight-value">Affordable at ৳2,10,000</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 160 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Race-Derived Technology:</strong> <span class="pro-description">RTR racing pedigree with sharp handling</span></li>
<li class="pro-item"><strong class="pro-label">LED Headlamp & DRL:</strong> <span class="pro-description">Full LED lighting with daytime running lights</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">160cc produces 15.5 PS with race-tuned setup</span></li>
<li class="pro-item"><strong class="pro-label">Competitive Pricing:</strong> <span class="pro-description">₳2,10,000 is value for features offered</span></li>
<li class="pro-item"><strong class="pro-label">Dual Disc Brakes:</strong> <span class="pro-description">Front and rear disc brakes with ABS</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Fully digital LCD cluster with lap timer</span></li>
<li class="pro-item"><strong class="pro-label">Monoshock Suspension:</strong> <span class="pro-description">Race-spec monoshock rear suspension</span></li>
<li class="pro-item"><strong class="pro-label">Remapper Options:</strong> <span class="pro-description">TVR remapper for performance tuning</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 160 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Limited Service Network:</strong> <span class="con-description">TVS service centers very scarce</span></li>
<li class="con-item"><strong class="con-label">Average Mileage:</strong> <span class="con-description">38-42 km/l not great for 160cc</span></li>
<li class="con-item"><strong class="con-label">Low Brand Value:</strong> <span class="con-description">TVS not popular in Bangladesh</span></li>
<li class="con-item"><strong class="con-label">Poor Resale Value:</strong> <span class="con-description">Worst resale among all brands</span></li>
<li class="con-item"><strong class="con-label">Firm Ride Quality:</strong> <span class="con-description">Stiff suspension uncomfortable on bad roads</span></li>
<li class="con-item"><strong class="con-label">Build Quality Issues:</strong> <span class="con-description">Plastic quality and paint not durable</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy TVS Apache RTR 160 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance enthusiasts not worried about resale</li>
<li class="best-for-item">Riders near TVS service centers</li>
<li class="best-for-item">Those wanting race technology at budget</li>
<li class="best-for-item">Tech-savvy riders who like tuning</li>
<li class="best-for-item">Track day enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy TVS Apache RTR 160?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Buyers worried about resale value</li>
<li class="not-recommended-item">Those needing widespread service network</li>
<li class="not-recommended-item">Comfort-focused riders</li>
<li class="not-recommended-item">People in remote areas</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 160 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,05,000 - ৳2,15,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳3,200 - ৳3,800 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,400-2,800/month for 30 km daily at 40 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,000 - ৳1,400 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 160 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Below average</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Good features but poor resale</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent performance</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Firm sporty setup</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- LED + ABS + lap timer</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Aggressive race look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">TVS Apache RTR 160 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of TVS Apache RTR 160?</h3>
<p class="faq-answer">TVS Apache RTR 160 delivers 38-42 km/l with race-tuned engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of TVS Apache RTR 160 in Bangladesh?</h3>
<p class="faq-answer">TVS Apache RTR 160 is priced between ৳2,05,000 to ৳2,15,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does TVS Apache RTR 160 have ABS?</h3>
<p class="faq-answer">Yes, TVS Apache RTR 160 has single channel ABS.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is TVS Apache RTR 160 good for daily use?</h3>
<p class="faq-answer">It's suitable for performance enthusiasts but firm ride may be uncomfortable for some.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 160 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.0/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">TVS Apache RTR 160 offers exciting performance and race technology at ৳2,10,000, making it appealing for performance enthusiasts. The LED lighting, lap timer, and race-tuned setup provide sporty thrills. However, very limited service network, poor resale value, and low brand presence make it a risky choice in Bangladesh. Only recommended for riders near TVS service centers who prioritize performance over practicality. Most buyers should consider Pulsar 150 or Suzuki Gixxer for better overall package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance enthusiasts near service centers</span></p>
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
		return fmt.Errorf("error creating TVS Apache RTR 160 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Apache RTR 160 (Rating: 4.0/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">টিভিএস এপাচি আরটিআর ১৬০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">টিভিএস এপাচি আরটিআর ১৬০ একটি পারফরম্যান্স-কেন্দ্রিক ১৬০সিসি মোটরসাইকেল যার মূল্য ৳২,১০,০০০ টাকা। রেস-উদ্ভূত প্রযুক্তি, এলইডি লাইটিং এবং স্পোর্টি হ্যান্ডলিং সহ, এটি উৎসাহীদের জন্য রোমাঞ্চকর পারফরম্যান্স প্রদান করে। তবে, সীমিত সার্ভিস নেটওয়ার্ক এবং কম ব্র্যান্ড ভ্যালু এটিকে বাংলাদেশ বাজারে একটি বিশেষ পছন্দ করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস এপাচি আরটিআর ১৬০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">রেস প্রযুক্তি:</strong> <span class="highlight-value">রেস-উদ্ভূত ইঞ্জিন এবং সাসপেনশন</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">ফুল এলইডি হেডল্যাম্প এবং ডিআরএল</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্টি পারফরম্যান্স:</strong> <span class="highlight-value">রেস-টিউনড সেটআপ সহ ১৬০সিসি</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রতিযোগিতামূলক মূল্য:</strong> <span class="highlight-value">৳২,১০,০০০ টাকায় সাশ্রয়ী</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস এপাচি আরটিআর ১৬০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">রেস-উদ্ভূত প্রযুক্তি:</strong> <span class="pro-description">তীক্ষ্ণ হ্যান্ডলিং সহ আরটিআর রেসিং পেডিগ্রি</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি হেডল্যাম্প ও ডিআরএল:</strong> <span class="pro-description">ডে-টাইম রানিং লাইট সহ ফুল এলইডি লাইটিং</span></li>
<li class="pro-item"><strong class="pro-label">ভালো পারফরম্যান্স:</strong> <span class="pro-description">১৬০সিসি রেস-টিউনড সেটআপ সহ ১৫.৫ পিএস উৎপাদন করে</span></li>
<li class="pro-item"><strong class="pro-label">প্রতিযোগিতামূলক মূল্য:</strong> <span class="pro-description">প্রদত্ত ফিচারের জন্য ৳২,১০,০০০ টাকা মূল্যবান</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল ডিস্ক ব্রেক:</strong> <span class="pro-description">এবিএস সহ সামনে এবং পিছনে ডিস্ক ব্রেক</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">ল্যাপ টাইমার সহ সম্পূর্ণ ডিজিটাল এলসিডি ক্লাস্টার</span></li>
<li class="pro-item"><strong class="pro-label">মনোশক সাসপেনশন:</strong> <span class="pro-description">রেস-স্পেক মনোশক রিয়ার সাসপেনশন</span></li>
<li class="pro-item"><strong class="pro-label">রিম্যাপার বিকল্প:</strong> <span class="pro-description">পারফরম্যান্স টিউনিংয়ের জন্য টিভিআর রিম্যাপার</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস এপাচি আরটিআর ১৬০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">টিভিএস সার্ভিস সেন্টার খুবই দুর্লভ</span></li>
<li class="con-item"><strong class="con-label">গড় মাইলেজ:</strong> <span class="con-description">১৬০সিসির জন্য ৩৮-৪২ কিমি/লিটার ভালো নয়</span></li>
<li class="con-item"><strong class="con-label">কম ব্র্যান্ড ভ্যালু:</strong> <span class="con-description">টিভিএস বাংলাদেশে জনপ্রিয় নয়</span></li>
<li class="con-item"><strong class="con-label">খারাপ রিসেল ভ্যালু:</strong> <span class="con-description">সব ব্র্যান্ডের মধ্যে সবচেয়ে খারাপ রিসেল</span></li>
<li class="con-item"><strong class="con-label">শক্ত রাইড কোয়ালিটি:</strong> <span class="con-description">খারাপ রাস্তায় শক্ত সাসপেনশন অস্বস্তিকর</span></li>
<li class="con-item"><strong class="con-label">বিল্ড কোয়ালিটি সমস্যা:</strong> <span class="con-description">প্লাস্টিক গুণমান এবং পেইন্ট টেকসই নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে টিভিএস এপাচি আরটিআর ১৬০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance enthusiasts not worried about resale</li>
<li class="best-for-item">Riders near TVS service centers</li>
<li class="best-for-item">Those wanting race technology at budget</li>
<li class="best-for-item">Tech-savvy riders who like tuning</li>
<li class="best-for-item">Track day enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">টিভিএস এপাচি আরটিআর ১৬০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Buyers worried about resale value</li>
<li class="not-recommended-item">Those needing widespread service network</li>
<li class="not-recommended-item">Comfort-focused riders</li>
<li class="not-recommended-item">People in remote areas</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">টিভিএস এপাচি আরটিআর ১৬০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,০৫,০০০ - ৳২,১৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৩,২০০ - ৳৩,৮০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪০ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,৪০০-২,৮০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,০০০ - ৳১,৪০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">টিভিএস এপাচি আরটিআর ১৬০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড়ের নিচে</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ভালো ফিচার কিন্তু খারাপ রিসেল</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার পারফরম্যান্স</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- শক্ত স্পোর্টি সেটআপ</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- এলইডি + এবিএস + ল্যাপ টাইমার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আগ্রাসী রেস লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">টিভিএস এপাচি আরটিআর ১৬০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">টিভিএস এপাচি আরটিআর ১৬০ এর মাইলেজ কত?</h3>
<p class="faq-answer">টিভিএস এপাচি আরটিআর ১৬০ রেস-টিউনড ইঞ্জিন সহ ৩৮-৪২ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে টিভিএস এপাচি আরটিআর ১৬০ এর দাম কত?</h3>
<p class="faq-answer">টিভিএস এপাচি আরটিআর ১৬০ এর দাম ৳২,০৫,০০০ থেকে ৳২,১৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">টিভিএস এপাচি আরটিআর ১৬০এ কি এবিএস আছে?</h3>
<p class="faq-answer">হ্যাঁ, টিভিএস এপাচি আরটিআর ১৬০এ সিঙ্গেল চ্যানেল এবিএস আছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য টিভিএস এপাচি আরটিআর ১৬০ কি ভালো?</h3>
<p class="faq-answer">এটি পারফরম্যান্স উৎসাহীদের জন্য উপযুক্ত কিন্তু শক্ত রাইড কারো কারো জন্য অস্বস্তিকর হতে পারে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে টিভিএস এপাচি আরটিআর ১৬০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.0/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">টিভিএস এপাচি আরটিআর ১৬০ ৳২,১০,০০০ টাকায় উত্তেজনাপূর্ণ পারফরম্যান্স এবং রেস প্রযুক্তি প্রদান করে, যা এটিকে পারফরম্যান্স উৎসাহীদের জন্য আকর্ষণীয় করে তোলে। এলইডি লাইটিং, ল্যাপ টাইমার এবং রেস-টিউনড সেটআপ স্পোর্টি রোমাঞ্চ প্রদান করে। তবে, অত্যন্ত সীমিত সার্ভিস নেটওয়ার্ক, খারাপ রিসেল ভ্যালু এবং কম ব্র্যান্ড উপস্থিতি এটিকে বাংলাদেশে একটি ঝুঁকিপূর্ণ পছন্দ করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সার্ভিস সেন্টারের কাছাকাছি পারফরম্যান্স উৎসাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Apache RTR 160 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Apache RTR 160\n")

	return nil
}
