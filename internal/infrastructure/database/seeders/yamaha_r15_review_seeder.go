package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaR15ReviewSeeder handles seeding Yamaha R15 product review and translation
type YamahaR15ReviewSeeder struct {
	BaseSeeder
}

// NewYamahaR15ReviewSeeder creates a new YamahaR15ReviewSeeder
func NewYamahaR15ReviewSeeder() *YamahaR15ReviewSeeder {
	return &YamahaR15ReviewSeeder{BaseSeeder: BaseSeeder{name: "Yamaha R15 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaR15ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R15").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha R15 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha R15 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha R15 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha R15 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha R15 is a premium 155cc sports bike priced at ৳5,50,000, offering genuine sports bike experience, excellent build quality, and track-ready performance. With VVA technology, slipper clutch, and aggressive styling, it's the ultimate choice for young enthusiasts seeking authentic sports bike thrills without breaking the bank.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R15 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">VVA Technology:</strong> <span class="highlight-value">Variable Valve Actuation</span></li>
<li class="highlight-item"><strong class="highlight-label">Slipper Clutch:</strong> <span class="highlight-value">Racing technology feature</span></li>
<li class="highlight-item"><strong class="highlight-label">Track Performance:</strong> <span class="highlight-value">19 HP, 136+ km/h top speed</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Design:</strong> <span class="highlight-value">Authentic sports bike styling</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R15 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Authentic Sports Bike:</strong> <span class="pro-description">True sports bike experience</span></li>
<li class="pro-item"><strong class="pro-label">VVA Engine:</strong> <span class="pro-description">Variable Valve Actuation technology</span></li>
<li class="pro-item"><strong class="pro-label">Slipper Clutch:</strong> <span class="pro-description">Racing-grade slipper clutch</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Build:</strong> <span class="pro-description">Premium Yamaha quality</span></li>
<li class="pro-item"><strong class="pro-label">Track Ready:</strong> <span class="pro-description">136+ km/h top speed capability</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Sharp sports bike design</span></li>
<li class="pro-item"><strong class="pro-label">ABS Safety:</strong> <span class="pro-description">Single channel ABS</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Fully digital instrument cluster</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R15 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Expensive:</strong> <span class="con-description">₹5,50,000 premium pricing</span></li>
<li class="con-item"><strong class="con-label">Poor Mileage:</strong> <span class="con-description">40-45 km/l, sports bike economy</span></li>
<li class="con-item"><strong class="con-label">Uncomfortable:</strong> <span class="con-description">Aggressive riding position</span></li>
<li class="con-item"><strong class="con-label">High Maintenance:</strong> <span class="con-description">Expensive parts and service</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Fewer Yamaha service centers</span></li>
<li class="con-item"><strong class="con-label">Not for Beginners:</strong> <span class="con-description">Too aggressive for first bike</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha R15 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Sports bike enthusiasts</li>
<li class="best-for-item">Young performance riders</li>
<li class="best-for-item">Track day participants</li>
<li class="best-for-item">Those wanting genuine sports bike</li>
<li class="best-for-item">Experienced riders seeking thrills</li>
<li class="best-for-item">Weekend canyon carvers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha R15?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Beginners or first-time riders</li>
<li class="not-recommended-item">Daily comfort commuters</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Elderly or comfort-seeking riders</li>
<li class="not-recommended-item">Long-distance touring</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R15 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳5,30,000 - ৳5,70,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳4,500 - ৳6,500 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,100-2,600/month for 30 km daily at 42 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳2,000 - ৳3,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R15 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Sports bike economy</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Yamaha excellence</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Great sports bike value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Excellent 19 HP</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Sports position</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Premium Yamaha</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Advanced features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Stunning sports design</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha R15 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Yamaha R15?</h3>
<p class="faq-answer">Yamaha R15 can reach 136+ km/h with 19 HP VVA engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Yamaha R15?</h3>
<p class="faq-answer">Yamaha R15 is priced between ৳5,30,000 to ৳5,70,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is R15 good for daily commute?</h3>
<p class="faq-answer">R15 is designed for sports performance, not comfort commuting. Consider other bikes for daily use.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does R15 have ABS?</h3>
<p class="faq-answer">Yes, Yamaha R15 comes with single channel ABS for front wheel safety.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R15 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Yamaha R15 at ৳5,50,000 is THE ULTIMATE affordable sports bike for enthusiasts, offering authentic sports bike experience with VVA technology, slipper clutch, 19 HP power, and genuine track-ready performance. With premium Yamaha build quality, aggressive styling, ABS safety, and 136+ km/h capability, it's unmatched in the 150cc sports segment. However, high price, poor mileage (40-45 km/l), uncomfortable sports position, expensive maintenance, and aggressive nature make it suitable only for dedicated sports bike enthusiasts. For those seeking genuine sports bike thrills without superbike costs, R15 is perfection.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For authentic affordable sports bike experience</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.8,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha R15 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha R15 (Rating: 4.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা আর১৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা আর১৫ একটি প্রিমিয়াম ১৫৫সিসি স্পোর্টস বাইক যার মূল্য ৳৫,৫০,০০০ টাকা, যা প্রকৃত স্পোর্টস বাইক অভিজ্ঞতা, চমৎকার বিল্ড কোয়ালিটি এবং ট্র্যাক-রেডি পারফরম্যান্স প্রদান করে। ভিভিএ প্রযুক্তি, স্লিপার ক্লাচ এবং আক্রমণাত্মক স্টাইলিং সহ, এটি ব্যাংক না ভেঙে প্রকৃত স্পোর্টস বাইক রোমাঞ্চ খোঁজা তরুণ উৎসাহীদের জন্য চূড়ান্ত পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা আর১৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ভিভিএ প্রযুক্তি:</strong> <span class="highlight-value">ভেরিয়েবল ভালভ অ্যাক্টিভেশন</span></li>
<li class="highlight-item"><strong class="highlight-label">স্লিপার ক্লাচ:</strong> <span class="highlight-value">রেসিং প্রযুক্তি ফিচার</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্র্যাক পারফরম্যান্স:</strong> <span class="highlight-value">১৯ এইচপি, ১৩৬+ কিমি/ঘণ্টা টপ স্পিড</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম ডিজাইন:</strong> <span class="highlight-value">প্রকৃত স্পোর্টস বাইক স্টাইলিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা আর১৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রকৃত স্পোর্টস বাইক:</strong> <span class="pro-description">সত্যিকারের স্পোর্টস বাইক অভিজ্ঞতা</span></li>
<li class="pro-item"><strong class="pro-label">ভিভিএ ইঞ্জিন:</strong> <span class="pro-description">ভেরিয়েবল ভালভ অ্যাক্টিভেশন প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">স্লিপার ক্লাচ:</strong> <span class="pro-description">রেসিং-গ্রেড স্লিপার ক্লাচ</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড:</strong> <span class="pro-description">প্রিমিয়াম ইয়ামাহা মান</span></li>
<li class="pro-item"><strong class="pro-label">ট্র্যাক রেডি:</strong> <span class="pro-description">১৩৬+ কিমি/ঘণ্টা টপ স্পিড সক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">তীক্ষ্ণ স্পোর্টস বাইক ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">এবিএস নিরাপত্তা:</strong> <span class="pro-description">সিঙ্গেল চ্যানেল এবিএস</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">সম্পূর্ণ ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা আর১৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">ব্যয়বহুল:</strong> <span class="con-description">৳৫,৫০,০০০ প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">খারাপ মাইলেজ:</strong> <span class="con-description">৪০-৪৫ কিমি/লিটার, স্পোর্টস বাইক ইকোনমি</span></li>
<li class="con-item"><strong class="con-label">অস্বস্তিকর:</strong> <span class="con-description">আক্রমণাত্মক রাইডিং অবস্থান</span></li>
<li class="con-item"><strong class="con-label">উচ্চ রক্ষণাবেক্ষণ:</strong> <span class="con-description">ব্যয়বহুল যন্ত্রাংশ এবং সার্ভিস</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">কম ইয়ামাহা সার্ভিস সেন্টার</span></li>
<li class="con-item"><strong class="con-label">শিক্ষানবিসদের জন্য নয়:</strong> <span class="con-description">প্রথম বাইকের জন্য অত্যধিক আক্রমণাত্মক</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা আর১৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Sports bike enthusiasts</li>
<li class="best-for-item">Young performance riders</li>
<li class="best-for-item">Track day participants</li>
<li class="best-for-item">Those wanting genuine sports bike</li>
<li class="best-for-item">Experienced riders seeking thrills</li>
<li class="best-for-item">Weekend canyon carvers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা আর১৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Beginners or first-time riders</li>
<li class="not-recommended-item">Daily comfort commuters</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Elderly or comfort-seeking riders</li>
<li class="not-recommended-item">Long-distance touring</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা আর১৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৫,৩০,০০০ - ৳৫,৭০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৪,৫০০ - ৳৬,৫০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,১০০-২,৬০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳২,০০০ - ৳৩,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা আর১৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- স্পোর্টস বাইক ইকোনমি</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- ইয়ামাহা শ্রেষ্ঠত্ব</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- দুর্দান্ত স্পোর্টস বাইক ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- চমৎকার ১৯ এইচপি</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- স্পোর্টস অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- প্রিমিয়াম ইয়ামাহা</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- উন্নত ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অত্যাশচর্য স্পোর্টস ডিজাইন</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা আর১৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা আর১৫ এর টপ স্পিড কত?</h3>
<p class="faq-answer">ইয়ামাহা আর১৫ ১৯ এইচপি ভিভিএ ইঞ্জিন সহ ১৩৬+ কিমি/ঘণ্টা পৌঁছাতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা আর১৫ এর দাম কত?</h3>
<p class="faq-answer">ইয়ামাহা আর১৫ এর দাম ৳৫,৩০,০০০ থেকে ৳৫,৭০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনিক যাতায়াতের জন্য আর১৫ কি ভালো?</h3>
<p class="faq-answer">আর১৫ স্পোর্টস পারফরম্যান্সের জন্য ডিজাইন করা, আরাম যাতায়াতের জন্য নয়। দৈনিক ব্যবহারের জন্য অন্য বাইক বিবেচনা করুন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">আর১৫ এ কি এবিএস আছে?</h3>
<p class="faq-answer">হ্যাঁ, ইয়ামাহা আর১৫ সামনের চাকার নিরাপত্তার জন্য সিঙ্গেল চ্যানেল এবিএস সহ আসে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা আর১৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.8/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৫,৫০,০০০ টাকায় ইয়ামাহা আর১৫ উৎসাহীদের জন্য চূড়ান্ত সাশ্রয়ী স্পোর্টস বাইক, যা ভিভিএ প্রযুক্তি, স্লিপার ক্লাচ, ১৯ এইচপি পাওয়ার এবং প্রকৃত ট্র্যাক-রেডি পারফরম্যান্স সহ প্রকৃত স্পোর্টস বাইক অভিজ্ঞতা প্রদান করে। প্রিমিয়াম ইয়ামাহা বিল্ড কোয়ালিটি, আক্রমণাত্মক স্টাইলিং, এবিএস নিরাপত্তা এবং ১৩৬+ কিমি/ঘণ্টা সক্ষমতা সহ, এটি ১৫০সিসি স্পোর্টস সেগমেন্টে অতুলনীয়। তবে, উচ্চ মূল্য, খারাপ মাইলেজ (৪০-৪৫ কিমি/লিটার), অস্বস্তিকর স্পোর্টস অবস্থান, ব্যয়বহুল রক্ষণাবেক্ষণ এবং আক্রমণাত্মক প্রকৃতি এটিকে শুধুমাত্র নিবেদিতপ্রাণ স্পোর্টস বাইক উৎসাহীদের জন্য উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রকৃত সাশ্রয়ী স্পোর্টস বাইক অভিজ্ঞতার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha R15 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha R15\n")

	return nil
}
