package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaFazerReviewSeeder handles seeding Yamaha Fazer product review and translation
type YamahaFazerReviewSeeder struct {
	BaseSeeder
}

// NewYamahaFazerReviewSeeder creates a new YamahaFazerReviewSeeder
func NewYamahaFazerReviewSeeder() *YamahaFazerReviewSeeder {
	return &YamahaFazerReviewSeeder{BaseSeeder: BaseSeeder{name: "Yamaha Fazer Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaFazerReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha Fazer").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha Fazer product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha Fazer product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha Fazer already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha Fazer Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha Fazer is a semi-faired 150cc motorcycle priced at ৳2,70,000, offering a blend of touring comfort and sporty performance. With fuel injection, LED lighting, and wind protection from the half fairing, it's ideal for riders seeking long-distance capability with style.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha Fazer Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Semi-Fairing Design:</strong> <span class="highlight-value">Half fairing provides wind protection</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Injection:</strong> <span class="highlight-value">FI technology for efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Full LED headlamp and indicators</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Touring:</strong> <span class="highlight-value">Upright position for long rides</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha Fazer Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Semi-Fairing Protection:</strong> <span class="pro-description">Half fairing reduces wind blast on highways</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Injection System:</strong> <span class="pro-description">Better fuel efficiency and smooth power</span></li>
<li class="pro-item"><strong class="pro-label">LED Lighting:</strong> <span class="pro-description">Full LED headlight and tail light</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Upright seating position for touring</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Solid Yamaha Japanese quality</span></li>
<li class="pro-item"><strong class="pro-label">Dual Disc Brakes:</strong> <span class="pro-description">Front and rear disc for safety</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Fully digital meter with gear indicator</span></li>
<li class="pro-item"><strong class="pro-label">Highway Stability:</strong> <span class="pro-description">Fairing improves stability at speed</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha Fazer Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price:</strong> <span class="con-description">৳2,70,000 expensive for 150cc</span></li>
<li class="con-item"><strong class="con-label">Average Mileage:</strong> <span class="con-description">40-44 km/l not great for FI</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha centers than Honda</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">Fairing adds weight, hard in traffic</span></li>
<li class="con-item"><strong class="con-label">Single Channel ABS:</strong> <span class="con-description">ABS only on front wheel</span></li>
<li class="con-item"><strong class="con-label">Lower Resale Value:</strong> <span class="con-description">Yamaha bikes depreciate faster</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha Fazer in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Long-distance touring riders</li>
<li class="best-for-item">Highway commuters</li>
<li class="best-for-item">Those wanting wind protection</li>
<li class="best-for-item">Riders prioritizing comfort</li>
<li class="best-for-item">Weekend touring enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha Fazer?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">City-only daily commuters</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those needing maximum mileage</li>
<li class="not-recommended-item">Riders in areas with limited Yamaha service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha Fazer Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,65,000 - ৳2,75,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳3,300 - ৳3,800 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,200-2,600/month for 30 km daily at 42 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 4 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,100 - ৳1,400 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha Fazer Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Good Yamaha quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Expensive but capable</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good FI performance</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent touring comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Solid quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- FI + LED + fairing</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Touring sporty look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha Fazer Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Yamaha Fazer?</h3>
<p class="faq-answer">Yamaha Fazer delivers 40-44 km/l with fuel injection.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Yamaha Fazer in Bangladesh?</h3>
<p class="faq-answer">Yamaha Fazer is priced between ৳2,65,000 to ৳2,75,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Yamaha Fazer good for long rides?</h3>
<p class="faq-answer">Yes, Yamaha Fazer is excellent for long rides with semi-fairing and comfortable ergonomics.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Yamaha Fazer have ABS?</h3>
<p class="faq-answer">Yes, Yamaha Fazer has single channel ABS on the front wheel.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha Fazer in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Yamaha Fazer is a capable touring-oriented 150cc bike that balances comfort and performance. At ৳2,70,000, it offers semi-fairing protection, fuel injection, and LED lights - features that justify the premium over naked bikes. Best for riders who frequently do highway rides and value wind protection. However, high price, average mileage, and limited service network make it less practical than alternatives for pure city use.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For highway touring enthusiasts</span></p>
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
		return fmt.Errorf("error creating Yamaha Fazer review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha Fazer (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা ফেইজার রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা ফেইজার একটি সেমি-ফেয়ারড ১৫০সিসি মোটরসাইকেল যার মূল্য ৳২,৭০,০০০ টাকা, যা ট্যুরিং আরাম এবং স্পোর্টি পারফরম্যান্সের মিশ্রণ প্রদান করে। ফুয়েল ইনজেকশন, এলইডি লাইটিং এবং হাফ ফেয়ারিং থেকে উইন্ড প্রোটেকশন সহ, এটি স্টাইল সহ লম্বা-দূরত্বের ক্ষমতা খুঁজছেন রাইডারদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা ফেইজার এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সেমি-ফেয়ারিং ডিজাইন:</strong> <span class="highlight-value">হাফ ফেয়ারিং উইন্ড প্রোটেকশন প্রদান করে</span></li>
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন:</strong> <span class="highlight-value">দক্ষতার জন্য এফআই প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">ফুল এলইডি হেডল্যাম্প এবং ইন্ডিকেটর</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক ট্যুরিং:</strong> <span class="highlight-value">লম্বা রাইডের জন্য সোজা অবস্থান</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা ফেইজার এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">সেমি-ফেয়ারিং প্রোটেকশন:</strong> <span class="pro-description">হাফ ফেয়ারিং হাইওয়েতে উইন্ড ব্লাস্ট কমায়</span></li>
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন সিস্টেম:</strong> <span class="pro-description">ভালো জ্বালানি দক্ষতা এবং মসৃণ পাওয়ার</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি লাইটিং:</strong> <span class="pro-description">ফুল এলইডি হেডলাইট এবং টেইল লাইট</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এর্গোনমিক্স:</strong> <span class="pro-description">ট্যুরিংয়ের জন্য সোজা বসার অবস্থান</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">শক্ত ইয়ামাহা জাপানি মান</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল ডিস্ক ব্রেক:</strong> <span class="pro-description">নিরাপত্তার জন্য সামনে এবং পিছনে ডিস্ক</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">গিয়ার ইন্ডিকেটর সহ সম্পূর্ণ ডিজিটাল মিটার</span></li>
<li class="pro-item"><strong class="pro-label">হাইওয়ে স্থিতিশীলতা:</strong> <span class="pro-description">ফেয়ারিং গতিতে স্থিতিশীলতা উন্নত করে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা ফেইজার এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">১৫০সিসির জন্য ৳২,৭০,০০০ টাকা ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">গড় মাইলেজ:</strong> <span class="con-description">এফআই-এর জন্য ৪০-৪৪ কিমি/লিটার ভালো নয়</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">হোন্ডার চেয়ে কম ইয়ামাহা সেন্টার</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">ফেয়ারিং ওজন যোগ করে, ট্রাফিকে কঠিন</span></li>
<li class="con-item"><strong class="con-label">সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="con-description">শুধুমাত্র সামনের চাকায় এবিএস</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">ইয়ামাহা বাইক দ্রুত অবমূল্যায়ন করে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা ফেইজার কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Long-distance touring riders</li>
<li class="best-for-item">Highway commuters</li>
<li class="best-for-item">Those wanting wind protection</li>
<li class="best-for-item">Riders prioritizing comfort</li>
<li class="best-for-item">Weekend touring enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা ফেইজার কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">City-only daily commuters</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those needing maximum mileage</li>
<li class="not-recommended-item">Riders in areas with limited Yamaha service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা ফেইজার এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৬৫,০০০ - ৳২,৭৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৩,৩০০ - ৳৩,৮০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,২০০-২,৬০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৪ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,১০০ - ৳১,৪০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা ফেইজার পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- ভালো ইয়ামাহা মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ব্যয়বহুল কিন্তু সক্ষম</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো এফআই পারফরম্যান্স</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার ট্যুরিং আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- শক্ত মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- এফআই + এলইডি + ফেয়ারিং</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ট্যুরিং স্পোর্টি লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা ফেইজার সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা ফেইজারের মাইলেজ কত?</h3>
<p class="faq-answer">ইয়ামাহা ফেইজার ফুয়েল ইনজেকশন সহ ৪০-৪৪ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে ইয়ামাহা ফেইজারের দাম কত?</h3>
<p class="faq-answer">ইয়ামাহা ফেইজারের দাম ৳২,৬৫,০০০ থেকে ৳২,৭৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">লম্বা রাইডের জন্য ইয়ামাহা ফেইজার কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, ইয়ামাহা ফেইজার সেমি-ফেয়ারিং এবং আরামদায়ক এর্গোনমিক্স সহ লম্বা রাইডের জন্য চমৎকার।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা ফেইজারে কি এবিএস আছে?</h3>
<p class="faq-answer">হ্যাঁ, ইয়ামাহা ফেইজারে সামনের চাকায় সিঙ্গেল চ্যানেল এবিএস আছে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা ফেইজার কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.1/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">ইয়ামাহা ফেইজার একটি সক্ষম ট্যুরিং-ভিত্তিক ১৫০সিসি বাইক যা আরাম এবং পারফরম্যান্স ভারসাম্য রাখে। ৳২,৭০,০০০ টাকায়, এটি সেমি-ফেয়ারিং প্রোটেকশন, ফুয়েল ইনজেকশন এবং এলইডি লাইট প্রদান করে - ফিচার যা নেকেড বাইকের চেয়ে প্রিমিয়ামকে ন্যায্যতা দেয়। রাইডারদের জন্য সেরা যারা প্রায়ই হাইওয়ে রাইড করেন এবং উইন্ড প্রোটেকশনকে মূল্য দেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- হাইওয়ে ট্যুরিং উৎসাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha Fazer already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha Fazer\n")

	return nil
}
