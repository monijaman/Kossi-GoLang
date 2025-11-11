package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGSXR150ABSReviewSeeder handles seeding Suzuki GSX-R150 ABS product review and translation
type SuzukiGSXR150ABSReviewSeeder struct {
	BaseSeeder
}

// NewSuzukiGSXR150ABSReviewSeeder creates a new SuzukiGSXR150ABSReviewSeeder
func NewSuzukiGSXR150ABSReviewSeeder() *SuzukiGSXR150ABSReviewSeeder {
	return &SuzukiGSXR150ABSReviewSeeder{BaseSeeder: BaseSeeder{name: "Suzuki GSX-R150 ABS Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGSXR150ABSReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki GSX-R150 ABS").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki GSX-R150 ABS product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki GSX-R150 ABS product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki GSX-R150 ABS already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki GSX-R150 ABS Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki GSX-R150 ABS is a premium 147cc fully-faired sports bike priced at ৳5,24,950, offering authentic MotoGP-inspired design, excellent build quality, and true sports bike experience. With full fairing, dual channel ABS, and Suzuki reliability, it's the ultimate choice for sports bike enthusiasts seeking genuine supersport styling.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki GSX-R150 ABS Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Full Fairing:</strong> <span class="highlight-value">Authentic supersport styling</span></li>
<li class="highlight-item"><strong class="highlight-label">MotoGP Design:</strong> <span class="highlight-value">GSX-R race replica</span></li>
<li class="highlight-item"><strong class="highlight-label">Dual Channel ABS:</strong> <span class="highlight-value">Advanced braking safety</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Quality:</strong> <span class="highlight-value">Excellent Suzuki build</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki GSX-R150 ABS Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Authentic Supersport:</strong> <span class="pro-description">True fully-faired sports bike</span></li>
<li class="pro-item"><strong class="pro-label">MotoGP Styling:</strong> <span class="pro-description">GSX-R race-inspired design</span></li>
<li class="pro-item"><strong class="pro-label">Dual Channel ABS:</strong> <span class="pro-description">Superior braking with ABS</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Build:</strong> <span class="pro-description">Premium Suzuki quality</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">19.2 HP, peppy engine</span></li>
<li class="pro-item"><strong class="pro-label">Aerodynamic:</strong> <span class="pro-description">Full fairing reduces wind</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Modern instrument cluster</span></li>
<li class="pro-item"><strong class="pro-label">LED Lights:</strong> <span class="pro-description">Full LED lighting system</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki GSX-R150 ABS Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Expensive:</strong> <span class="con-description">₹5,24,950 premium pricing</span></li>
<li class="con-item"><strong class="con-label">Poor Mileage:</strong> <span class="con-description">42-45 km/l, sports bike economy</span></li>
<li class="con-item"><strong class="con-label">Uncomfortable:</strong> <span class="con-description">Aggressive sports position</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Fewer Suzuki centers</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">High service costs</span></li>
<li class="con-item"><strong class="con-label">Not for Beginners:</strong> <span class="con-description">Too aggressive for first bike</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki GSX-R150 ABS in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Serious sports bike enthusiasts</li>
<li class="best-for-item">Track day participants</li>
<li class="best-for-item">Those wanting full fairing</li>
<li class="best-for-item">Performance-oriented riders</li>
<li class="best-for-item">MotoGP styling fans</li>
<li class="best-for-item">Experienced riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki GSX-R150 ABS?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Beginners or first-time riders</li>
<li class="not-recommended-item">Daily comfort commuters</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Elderly or comfort seekers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki GSX-R150 ABS Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳5,10,000 - ৳5,40,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳4,800 - ৳6,500 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,100-2,600/month for 30 km daily at 43 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳2,200 - ৳3,800 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki GSX-R150 ABS Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Sports bike economy</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Suzuki quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Premium but worth it</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Good 19.2 HP</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Racing position</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Excellent Suzuki</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Premium features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- MotoGP replica</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki GSX-R150 ABS Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Suzuki GSX-R150?</h3>
<p class="faq-answer">Suzuki GSX-R150 can reach 135+ km/h with 19.2 HP engine and full fairing.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Suzuki GSX-R150 ABS?</h3>
<p class="faq-answer">Suzuki GSX-R150 ABS is priced between ৳5,10,000 to ৳5,40,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is GSX-R150 better than R15?</h3>
<p class="faq-answer">Both excellent. GSX-R150 has full fairing and dual ABS; R15 has VVA and slipper clutch. Choose based on preference.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is GSX-R150 good for daily commute?</h3>
<p class="faq-answer">Not ideal. GSX-R150 is for sports performance, not comfort commuting.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki GSX-R150 ABS in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Suzuki GSX-R150 ABS at ৳5,24,950 is THE ULTIMATE fully-faired sports bike for serious enthusiasts, offering authentic MotoGP-inspired GSX-R design, full fairing aerodynamics, dual channel ABS, premium Suzuki build quality, and genuine supersport experience. With 19.2 HP, full LED package, and excellent reliability, it's the pinnacle of the 150cc sports segment. However, very high price, poor mileage (42-45 km/l), aggressive uncomfortable riding position, expensive maintenance, and limited service network make it suitable only for dedicated sports bike enthusiasts. For those seeking authentic fully-faired supersport with Japanese reliability, GSX-R150 is perfection.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For authentic fully-faired supersport with MotoGP design</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.7,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Suzuki GSX-R150 ABS review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki GSX-R150 ABS (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি জিএসএক্স-আর১৫০ এবিএস রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি জিএসএক্স-আর১৫০ এবিএস একটি প্রিমিয়াম ১৪৭সিসি ফুল-ফেয়ার্ড স্পোর্টস বাইক যার মূল্য ৳৫,২৪,৯৫০ টাকা, যা প্রকৃত মোটোজিপি-অনুপ্রাণিত ডিজাইন, চমৎকার বিল্ড কোয়ালিটি এবং সত্যিকারের স্পোর্টস বাইক অভিজ্ঞতা প্রদান করে। ফুল ফেয়ারিং, ডুয়াল চ্যানেল এবিএস এবং সুজুকি নির্ভরযোগ্যতা সহ, এটি প্রকৃত সুপারস্পোর্ট স্টাইলিং খোঁজা স্পোর্টস বাইক উৎসাহীদের জন্য চূড়ান্ত পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি জিএসএক্স-আর১৫০ এবিএস এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুল ফেয়ারিং:</strong> <span class="highlight-value">প্রকৃত সুপারস্পোর্ট স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">মোটোজিপি ডিজাইন:</strong> <span class="highlight-value">জিএসএক্স-আর রেস রেপ্লিকা</span></li>
<li class="highlight-item"><strong class="highlight-label">ডুয়াল চ্যানেল এবিএস:</strong> <span class="highlight-value">উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম মান:</strong> <span class="highlight-value">চমৎকার সুজুকি বিল্ড</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি জিএসএক্স-আর১৫০ এবিএস এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রকৃত সুপারস্পোর্ট:</strong> <span class="pro-description">সত্যিকারের ফুল-ফেয়ার্ড স্পোর্টস বাইক</span></li>
<li class="pro-item"><strong class="pro-label">মোটোজিপি স্টাইলিং:</strong> <span class="pro-description">জিএসএক্স-আর রেস-অনুপ্রাণিত ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল চ্যানেল এবিএস:</strong> <span class="pro-description">এবিএস সহ উৎকৃষ্ট ব্রেকিং</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড:</strong> <span class="pro-description">প্রিমিয়াম সুজুকি মান</span></li>
<li class="pro-item"><strong class="pro-label">ভালো পারফরম্যান্স:</strong> <span class="pro-description">১৯.২ এইচপি, প্রাণবন্ত ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">অ্যারোডাইনামিক:</strong> <span class="pro-description">ফুল ফেয়ারিং বাতাস কমায়</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">আধুনিক ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি লাইট:</strong> <span class="pro-description">সম্পূর্ণ এলইডি লাইটিং সিস্টেম</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি জিএসএক্স-আর১৫০ এবিএস এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত ব্যয়বহুল:</strong> <span class="con-description">৳৫,২৪,৯৫০ প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">খারাপ মাইলেজ:</strong> <span class="con-description">৪২-৪৫ কিমি/লিটার, স্পোর্টস বাইক ইকোনমি</span></li>
<li class="con-item"><strong class="con-label">অস্বস্তিকর:</strong> <span class="con-description">আক্রমণাত্মক স্পোর্টস অবস্থান</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">কম সুজুকি সেন্টার</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">উচ্চ সার্ভিস খরচ</span></li>
<li class="con-item"><strong class="con-label">শিক্ষানবিসদের জন্য নয়:</strong> <span class="con-description">প্রথম বাইকের জন্য অত্যধিক আক্রমণাত্মক</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি জিএসএক্স-আর১৫০ এবিএস কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Serious sports bike enthusiasts</li>
<li class="best-for-item">Track day participants</li>
<li class="best-for-item">Those wanting full fairing</li>
<li class="best-for-item">Performance-oriented riders</li>
<li class="best-for-item">MotoGP styling fans</li>
<li class="best-for-item">Experienced riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি জিএসএক্স-আর১৫০ এবিএস কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Beginners or first-time riders</li>
<li class="not-recommended-item">Daily comfort commuters</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Elderly or comfort seekers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি জিএসএক্স-আর১৫০ এবিএস এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৫,১০,০০০ - ৳৫,৪০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৪,৮০০ - ৳৬,৫০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪৩ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,১০০-২,৬০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳২,২০০ - ৳৩,৮০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি জিএসএক্স-আর১৫০ এবিএস পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- স্পোর্টস বাইক ইকোনমি</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- সুজুকি মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- প্রিমিয়াম কিন্তু মূল্যবান</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- ভালো ১৯.২ এইচপি</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- রেসিং অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- চমৎকার সুজুকি</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- প্রিমিয়াম ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- মোটোজিপি রেপ্লিকা</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি জিএসএক্স-আর১৫০ এবিএস সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">সুজুকি জিএসএক্স-আর১৫০ এর টপ স্পিড কত?</h3>
<p class="faq-answer">সুজুকি জিএসএক্স-আর১৫০ ১৯.২ এইচপি ইঞ্জিন এবং ফুল ফেয়ারিং সহ ১৩৫+ কিমি/ঘণ্টা পৌঁছাতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি জিএসএক্স-আর১৫০ এবিএস এর দাম কত?</h3>
<p class="faq-answer">সুজুকি জিএসএক্স-আর১৫০ এবিএস এর দাম ৳৫,১০,০০০ থেকে ৳৫,৪০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">জিএসএক্স-আর১৫০ কি আর১৫ এর চেয়ে ভালো?</h3>
<p class="faq-answer">উভয়ই চমৎকার। জিএসএক্স-আর১৫০ এ ফুল ফেয়ারিং এবং ডুয়াল এবিএস আছে; আর১৫ এ ভিভিএ এবং স্লিপার ক্লাচ আছে। পছন্দ অনুযায়ী বেছে নিন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনিক যাতায়াতের জন্য জিএসএক্স-আর১৫০ কি ভালো?</h3>
<p class="faq-answer">আদর্শ নয়। জিএসএক্স-আর১৫০ স্পোর্টস পারফরম্যান্সের জন্য, আরাম যাতায়াতের জন্য নয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি জিএসএক্স-আর১৫০ এবিএস কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৫,২৪,৯৫০ টাকায় সুজুকি জিএসএক্স-আর১৫০ এবিএস গুরুতর উৎসাহীদের জন্য চূড়ান্ত ফুল-ফেয়ার্ড স্পোর্টস বাইক, যা প্রকৃত মোটোজিপি-অনুপ্রাণিত জিএসএক্স-আর ডিজাইন, ফুল ফেয়ারিং অ্যারোডাইনামিক্স, ডুয়াল চ্যানেল এবিএস, প্রিমিয়াম সুজুকি বিল্ড কোয়ালিটি এবং প্রকৃত সুপারস্পোর্ট অভিজ্ঞতা প্রদান করে। ১৯.২ এইচপি, ফুল এলইডি প্যাকেজ এবং চমৎকার নির্ভরযোগ্যতা সহ, এটি ১৫০সিসি স্পোর্টস সেগমেন্টের শিখর। তবে, অত্যন্ত উচ্চ মূল্য, খারাপ মাইলেজ (৪২-৪৫ কিমি/লিটার), আক্রমণাত্মক অস্বস্তিকর রাইডিং অবস্থান, ব্যয়বহুল রক্ষণাবেক্ষণ এবং সীমিত সার্ভিস নেটওয়ার্ক এটিকে শুধুমাত্র নিবেদিতপ্রাণ স্পোর্টস বাইক উৎসাহীদের জন্য উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- মোটোজিপি ডিজাইন সহ প্রকৃত ফুল-ফেয়ার্ড সুপারস্পোর্টের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki GSX-R150 ABS already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki GSX-R150 ABS\n")

	return nil
}
