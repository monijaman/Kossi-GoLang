package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajCT100ReviewSeeder handles seeding Bajaj CT 100 product review and translation
type BajajCT100ReviewSeeder struct {
	BaseSeeder
}

// NewBajajCT100ReviewSeeder creates a new BajajCT100ReviewSeeder
func NewBajajCT100ReviewSeeder() *BajajCT100ReviewSeeder {
	return &BajajCT100ReviewSeeder{BaseSeeder: BaseSeeder{name: "Bajaj CT 100 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajCT100ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj CT 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj CT 100 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj CT 100 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj CT 100 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj CT 100 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj CT 100 is an ultra-affordable 102cc commuter bike priced at ৳75,000, offering basic transportation, decent fuel efficiency, and rock-bottom pricing. With simple design, reliable engine, and minimal features, it's ideal for budget-conscious buyers seeking basic two-wheeler mobility at the lowest possible cost.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj CT 100 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Rock Bottom Price:</strong> <span class="highlight-value">Only ৳75,000 entry level</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">70+ km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Simple & Reliable:</strong> <span class="highlight-value">Basic proven technology</span></li>
<li class="highlight-item"><strong class="highlight-label">Wide Service:</strong> <span class="highlight-value">Bajaj service everywhere</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj CT 100 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Ultra Affordable:</strong> <span class="pro-description">Cheapest bike in market</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">70+ km/l outstanding efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Simple engine, cheap parts</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Bajaj service everywhere</span></li>
<li class="pro-item"><strong class="pro-label">Kick Start:</strong> <span class="pro-description">Both kick and electric start</span></li>
<li class="pro-item"><strong class="pro-label">Proven Engine:</strong> <span class="pro-description">Time-tested 102cc engine</span></li>
<li class="pro-item"><strong class="pro-label">Rural Friendly:</strong> <span class="pro-description">Perfect for rural areas</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj CT 100 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Basic:</strong> <span class="con-description">No modern features</span></li>
<li class="con-item"><strong class="con-label">Poor Build Quality:</strong> <span class="con-description">Cheap plastic and materials</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">Only 8.1 HP, very weak</span></li>
<li class="con-item"><strong class="con-label">Outdated Design:</strong> <span class="con-description">Old-fashioned styling</span></li>
<li class="con-item"><strong class="con-label">Poor Resale:</strong> <span class="con-description">Low resale value</span></li>
<li class="con-item"><strong class="con-label">No Comfort:</strong> <span class="con-description">Basic seat and suspension</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Bajaj CT 100 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">First-time bike buyers on tight budget</li>
<li class="best-for-item">Rural area transportation</li>
<li class="best-for-item">Students and workers</li>
<li class="best-for-item">Backup bike for families</li>
<li class="best-for-item">Those needing basic mobility</li>
<li class="best-for-item">Delivery boys on budget</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Bajaj CT 100?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious buyers</li>
<li class="not-recommended-item">Urban commuters wanting features</li>
<li class="not-recommended-item">Those prioritizing comfort</li>
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Image-conscious riders</li>
<li class="not-recommended-item">Highway riding</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj CT 100 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳70,000 - ৳80,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,800 - ৳2,400 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,200-1,500/month for 30 km daily at 70 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳500 - ৳800 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj CT 100 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Outstanding 70+ km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Simple and reliable</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Unbeatable price</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Very basic power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Budget build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Minimal features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic styling</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Bajaj CT 100 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Bajaj CT 100?</h3>
<p class="faq-answer">Bajaj CT 100 delivers outstanding 70+ km/l with 102cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Bajaj CT 100?</h3>
<p class="faq-answer">Bajaj CT 100 is priced between ৳70,000 to ৳80,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is CT 100 good for first bike?</h3>
<p class="faq-answer">Yes, CT 100 is perfect first bike for budget buyers with excellent mileage.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does CT 100 have electric start?</h3>
<p class="faq-answer">Yes, CT 100 comes with both kick start and electric start options.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj CT 100 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Bajaj CT 100 at ৳75,000 is the ULTIMATE budget choice for basic transportation, offering unbeatable affordability, outstanding 70+ km/l mileage, simple reliable technology, and widespread service support. While it lacks modern features, comfort, and style (basic 8.1 HP power, outdated design, poor build quality), the rock-bottom price, exceptional fuel economy, low maintenance costs, and Bajaj's extensive service network make it perfect for budget-conscious buyers needing basic mobility. Ideal for students, rural riders, first-time buyers, and anyone prioritizing absolute affordability over features and comfort. The cheapest way to own a motorcycle.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultra-budget basic transportation</span></p>
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
		return fmt.Errorf("error creating Bajaj CT 100 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj CT 100 (Rating: 3.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ সিটি ১০০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ সিটি ১০০ একটি অতি-সাশ্রয়ী ১০২সিসি কমিউটার বাইক যার মূল্য ৳৭৫,০০০ টাকা, যা প্রাথমিক পরিবহন, মোটামুটি জ্বালানি দক্ষতা এবং তলানির দাম প্রদান করে। সরল ডিজাইন, নির্ভরযোগ্য ইঞ্জিন এবং ন্যূনতম ফিচার সহ, এটি সর্বনিম্ন সম্ভাব্য খরচে প্রাথমিক দ্বিচাকা চলাচল খোঁজা বাজেট-সচেতন ক্রেতাদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ সিটি ১০০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">তলানি দাম:</strong> <span class="highlight-value">মাত্র ৳৭৫,০০০ টাকা প্রবেশ স্তর</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৭০+ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">সরল ও নির্ভরযোগ্য:</strong> <span class="highlight-value">প্রাথমিক প্রমাণিত প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যাপক সার্ভিস:</strong> <span class="highlight-value">সর্বত্র বাজাজ সার্ভিস</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ সিটি ১০০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অতি সাশ্রয়ী:</strong> <span class="pro-description">বাজারের সবচেয়ে সস্তা বাইক</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">৭০+ কিমি/লিটার অসাধারণ দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সরল ইঞ্জিন, সস্তা যন্ত্রাংশ</span></li>
<li class="pro-item"><strong class="pro-label">ব্যাপক সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সর্বত্র বাজাজ সার্ভিস</span></li>
<li class="pro-item"><strong class="pro-label">কিক স্টার্ট:</strong> <span class="pro-description">কিক এবং ইলেকট্রিক দুটো স্টার্ট</span></li>
<li class="pro-item"><strong class="pro-label">প্রমাণিত ইঞ্জিন:</strong> <span class="pro-description">সময়ের পরীক্ষায় উত্তীর্ণ ১০২সিসি ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">গ্রামীণ বান্ধব:</strong> <span class="pro-description">গ্রামীণ এলাকার জন্য নিখুঁত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ সিটি ১০০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত প্রাথমিক:</strong> <span class="con-description">কোনো আধুনিক ফিচার নেই</span></li>
<li class="con-item"><strong class="con-label">খারাপ বিল্ড কোয়ালিটি:</strong> <span class="con-description">সস্তা প্লাস্টিক এবং উপাদান</span></li>
<li class="con-item"><strong class="con-label">কম পাওয়ার:</strong> <span class="con-description">শুধু ৮.১ এইচপি, অত্যন্ত দুর্বল</span></li>
<li class="con-item"><strong class="con-label">পুরানো ডিজাইন:</strong> <span class="con-description">পুরনো ধাঁচের স্টাইলিং</span></li>
<li class="con-item"><strong class="con-label">খারাপ রিসেল:</strong> <span class="con-description">কম রিসেল ভ্যালু</span></li>
<li class="con-item"><strong class="con-label">কোনো আরাম নেই:</strong> <span class="con-description">প্রাথমিক সিট এবং সাসপেনশন</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বাজাজ সিটি ১০০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">First-time bike buyers on tight budget</li>
<li class="best-for-item">Rural area transportation</li>
<li class="best-for-item">Students and workers</li>
<li class="best-for-item">Backup bike for families</li>
<li class="best-for-item">Those needing basic mobility</li>
<li class="best-for-item">Delivery boys on budget</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বাজাজ সিটি ১০০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Style-conscious buyers</li>
<li class="not-recommended-item">Urban commuters wanting features</li>
<li class="not-recommended-item">Those prioritizing comfort</li>
<li class="not-recommended-item">Performance seekers</li>
<li class="not-recommended-item">Image-conscious riders</li>
<li class="not-recommended-item">Highway riding</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বাজাজ সিটি ১০০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৭০,০০০ - ৳৮০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৮০০ - ৳২,৪০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৭০ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,২০০-১,৫০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৫০০ - ৳৮০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বাজাজ সিটি ১০০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অসাধারণ ৭০+ কিমি/লিটার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- সরল এবং নির্ভরযোগ্য</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অপরাজেয় দাম</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- অত্যন্ত প্রাথমিক পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- প্রাথমিক আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- বাজেট বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- ন্যূনতম ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- প্রাথমিক স্টাইলিং</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বাজাজ সিটি ১০০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">বাজাজ সিটি ১০০ এর মাইলেজ কত?</h3>
<p class="faq-answer">বাজাজ সিটি ১০০ ১০২সিসি ইঞ্জিন সহ অসাধারণ ৭০+ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাজাজ সিটি ১০০ এর দাম কত?</h3>
<p class="faq-answer">বাজাজ সিটি ১০০ এর দাম ৳৭০,০০০ থেকে ৳৮০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">প্রথম বাইকের জন্য সিটি ১০০ কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, সিটি ১০০ চমৎকার মাইলেজ সহ বাজেট ক্রেতাদের জন্য নিখুঁত প্রথম বাইক।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সিটি ১০০ এ কি ইলেকট্রিক স্টার্ট আছে?</h3>
<p class="faq-answer">হ্যাঁ, সিটি ১০০ কিক স্টার্ট এবং ইলেকট্রিক স্টার্ট দুটো অপশন সহ আসে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বাজাজ সিটি ১০০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.8/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৭৫,০০০ টাকায় বাজাজ সিটি ১০০ প্রাথমিক পরিবহনের জন্য চূড়ান্ত বাজেট পছন্দ, যা অপরাজেয় সাশ্রয়ীতা, অসাধারণ ৭০+ কিমি/লিটার মাইলেজ, সরল নির্ভরযোগ্য প্রযুক্তি এবং ব্যাপক সার্ভিস সহায়তা প্রদান করে। যদিও এতে আধুনিক ফিচার, আরাম এবং স্টাইলের অভাব রয়েছে (প্রাথমিক ৮.১ এইচপি পাওয়ার, পুরানো ডিজাইন, খারাপ বিল্ড কোয়ালিটি), তলানি দাম, ব্যতিক্রমী জ্বালানি অর্থনীতি, কম রক্ষণাবেক্ষণ খরচ এবং বাজাজের ব্যাপক সার্ভিস নেটওয়ার্ক এটিকে প্রাথমিক চলাচলের প্রয়োজনে বাজেট-সচেতন ক্রেতাদের জন্য নিখুঁত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অতি-বাজেট প্রাথমিক পরিবহনের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj CT 100 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj CT 100\n")

	return nil
}
