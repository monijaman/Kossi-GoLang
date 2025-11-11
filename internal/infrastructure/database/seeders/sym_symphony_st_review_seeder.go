package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SYMSymphonySTReviewSeeder handles seeding SYM Symphony ST product review and translation
type SYMSymphonySTReviewSeeder struct {
	BaseSeeder
}

// NewSYMSymphonySTReviewSeeder creates a new SYMSymphonySTReviewSeeder
func NewSYMSymphonySTReviewSeeder() *SYMSymphonySTReviewSeeder {
	return &SYMSymphonySTReviewSeeder{BaseSeeder: BaseSeeder{name: "SYM Symphony ST Review"}}
}

// Seed implements the Seeder interface
func (s *SYMSymphonySTReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "SYM Symphony ST").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("SYM Symphony ST product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding SYM Symphony ST product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for SYM Symphony ST already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">SYM Symphony ST Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The SYM Symphony ST is a budget 125cc scooter priced at ৳95,000, offering Taiwanese build quality, decent features, and good mileage. With automatic convenience, comfortable ride, and better quality than Chinese brands, it's for budget buyers seeking affordable scooter with better reliability than local options.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">SYM Symphony ST Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Taiwanese Quality:</strong> <span class="highlight-value">Better than Chinese brands</span></li>
<li class="highlight-item"><strong class="highlight-label">125cc Power:</strong> <span class="highlight-value">More power than 110cc</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">55+ km/l efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Price:</strong> <span class="highlight-value">Affordable at ৳95,000</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">SYM Symphony ST Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Better Build Quality:</strong> <span class="pro-description">Taiwanese quality standards</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">55+ km/l fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Automatic Convenience:</strong> <span class="pro-description">No gear shifting needed</span></li>
<li class="pro-item"><strong class="pro-label">125cc Power:</strong> <span class="pro-description">Better than 110cc scooters</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Decent comfort level</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Budget-friendly for scooter</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">SYM Symphony ST Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Very few SYM service centers</span></li>
<li class="con-item"><strong class="con-label">Poor Resale:</strong> <span class="con-description">Low resale value</span></li>
<li class="con-item"><strong class="con-label">Unknown Brand:</strong> <span class="con-description">Not popular in Bangladesh</span></li>
<li class="con-item"><strong class="con-label">Average Features:</strong> <span class="con-description">Basic features only</span></li>
<li class="con-item"><strong class="con-label">Parts Availability:</strong> <span class="con-description">Spare parts hard to find</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy SYM Symphony ST in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget scooter buyers</li>
<li class="best-for-item">Those wanting better than Chinese</li>
<li class="best-for-item">City commuting</li>
<li class="best-for-item">First scooter buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy SYM Symphony ST?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Brand-conscious buyers</li>
<li class="not-recommended-item">Long-term resale value seekers</li>
<li class="not-recommended-item">Areas without SYM service</li>
<li class="not-recommended-item">Those wanting Japanese reliability</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">SYM Symphony ST Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳90,000 - ৳1,00,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,400 - ৳3,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,600-2,000/month for 30 km daily at 55 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 2,500 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳800 - ৳1,300 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">SYM Symphony ST Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good 55+ km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic 125cc power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Better than Chinese</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average styling</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">SYM Symphony ST Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of SYM Symphony ST?</h3>
<p class="faq-answer">SYM Symphony ST delivers 55+ km/l with 125cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of SYM Symphony ST?</h3>
<p class="faq-answer">SYM Symphony ST is priced between ৳90,000 to ৳1,00,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is SYM Symphony ST reliable?</h3>
<p class="faq-answer">SYM offers better quality than Chinese brands but not as reliable as Honda/Yamaha.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Should I buy SYM Symphony ST?</h3>
<p class="faq-answer">Consider if budget-tight and want better than Chinese brands. Otherwise save for Honda Activa.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy SYM Symphony ST in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">SYM Symphony ST at ৳95,000 is a decent budget scooter option offering better Taiwanese build quality than Chinese brands, good 55+ km/l mileage, 125cc power, and automatic convenience. While it's a step up from Road Prince/Chinese options, the limited service network, poor resale value, unknown brand status, and spare parts availability issues limit its appeal. Better than ultra-budget Chinese options but not comparable to Honda/Yamaha reliability. Consider only if budget-constrained and want better quality than Chinese brands, otherwise save ৳65,000 more for Honda Activa's proven reliability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget scooter better than Chinese brands</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating SYM Symphony ST review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for SYM Symphony ST (Rating: 3.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">এসওয়াইএম সিম্ফনি এসটি রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">এসওয়াইএম সিম্ফনি এসটি একটি বাজেট ১২৫সিসি স্কুটার যার মূল্য ৳৯৫,০০০ টাকা, যা তাইওয়ানি বিল্ড কোয়ালিটি, মোটামুটি ফিচার এবং ভালো মাইলেজ প্রদান করে। অটোমেটিক সুবিধা, আরামদায়ক রাইড এবং চীনা ব্র্যান্ডের চেয়ে ভালো মান সহ, এটি স্থানীয় অপশনের চেয়ে ভালো নির্ভরযোগ্যতা সহ সাশ্রয়ী স্কুটার খোঁজা বাজেট ক্রেতাদের জন্য।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">এসওয়াইএম সিম্ফনি এসটি এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">তাইওয়ানি মান:</strong> <span class="highlight-value">চীনা ব্র্যান্ডের চেয়ে ভালো</span></li>
<li class="highlight-item"><strong class="highlight-label">১২৫সিসি পাওয়ার:</strong> <span class="highlight-value">১১০সিসির চেয়ে বেশি পাওয়ার</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৫৫+ কিমি/লিটার দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট মূল্য:</strong> <span class="highlight-value">৳৯৫,০০০ টাকায় সাশ্রয়ী</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">এসওয়াইএম সিম্ফনি এসটি এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">তাইওয়ানি মান স্ট্যান্ডার্ড</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৫৫+ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">অটোমেটিক সুবিধা:</strong> <span class="pro-description">কোনো গিয়ার পরিবর্তনের প্রয়োজন নেই</span></li>
<li class="pro-item"><strong class="pro-label">১২৫সিসি পাওয়ার:</strong> <span class="pro-description">১১০সিসি স্কুটারের চেয়ে ভালো</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">মোটামুটি আরাম স্তর</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">স্কুটারের জন্য বাজেট-বান্ধব</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">এসওয়াইএম সিম্ফনি এসটি এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">খুব কম এসওয়াইএম সার্ভিস সেন্টার</span></li>
<li class="con-item"><strong class="con-label">খারাপ রিসেল:</strong> <span class="con-description">কম রিসেল ভ্যালু</span></li>
<li class="con-item"><strong class="con-label">অজানা ব্র্যান্ড:</strong> <span class="con-description">বাংলাদেশে জনপ্রিয় নয়</span></li>
<li class="con-item"><strong class="con-label">গড় ফিচার:</strong> <span class="con-description">শুধু প্রাথমিক ফিচার</span></li>
<li class="con-item"><strong class="con-label">যন্ত্রাংশ প্রাপ্যতা:</strong> <span class="con-description">খুচরা যন্ত্রাংশ খুঁজে পাওয়া কঠিন</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে এসওয়াইএম সিম্ফনি এসটি কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget scooter buyers</li>
<li class="best-for-item">Those wanting better than Chinese</li>
<li class="best-for-item">City commuting</li>
<li class="best-for-item">First scooter buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">এসওয়াইএম সিম্ফনি এসটি কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Brand-conscious buyers</li>
<li class="not-recommended-item">Long-term resale value seekers</li>
<li class="not-recommended-item">Areas without SYM service</li>
<li class="not-recommended-item">Those wanting Japanese reliability</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">এসওয়াইএম সিম্ফনি এসটি এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৯০,০০০ - ৳১,০০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৪০০ - ৳৩,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫৫ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৬০০-২,০০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ২,৫০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৮০০ - ৳১,৩০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">এসওয়াইএম সিম্ফনি এসটি পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ৫৫+ কিমি/লিটার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- মোটামুটি ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- প্রাথমিক ১২৫সিসি পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- মোটামুটি আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- চীনার চেয়ে ভালো</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- প্রাথমিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় স্টাইলিং</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">এসওয়াইএম সিম্ফনি এসটি সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এসওয়াইএম সিম্ফনি এসটি এর মাইলেজ কত?</h3>
<p class="faq-answer">এসওয়াইএম সিম্ফনি এসটি ১২৫সিসি ইঞ্জিন সহ ৫৫+ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এসওয়াইএম সিম্ফনি এসটি এর দাম কত?</h3>
<p class="faq-answer">এসওয়াইএম সিম্ফনি এসটি এর দাম ৳৯০,০০০ থেকে ৳১,০০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এসওয়াইএম সিম্ফনি এসটি কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">এসওয়াইএম চীনা ব্র্যান্ডের চেয়ে ভালো মান প্রদান করে কিন্তু হন্ডা/ইয়ামাহার মতো নির্ভরযোগ্য নয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">আমি কি এসওয়াইএম সিম্ফনি এসটি কিনব?</h3>
<p class="faq-answer">বিবেচনা করুন যদি বাজেট সীমিত এবং চীনা ব্র্যান্ডের চেয়ে ভালো চান। অন্যথায় হন্ডা অ্যাকটিভার জন্য সঞ্চয় করুন।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে এসওয়াইএম সিম্ফনি এসটি কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৯৫,০০০ টাকায় এসওয়াইএম সিম্ফনি এসটি একটি মোটামুটি বাজেট স্কুটার অপশন যা চীনা ব্র্যান্ডের চেয়ে ভালো তাইওয়ানি বিল্ড কোয়ালিটি, ভালো ৫৫+ কিমি/লিটার মাইলেজ, ১২৫সিসি পাওয়ার এবং অটোমেটিক সুবিধা প্রদান করে। যদিও এটি রোড প্রিন্স/চীনা অপশনের চেয়ে এক ধাপ উপরে, সীমিত সার্ভিস নেটওয়ার্ক, খারাপ রিসেল ভ্যালু, অজানা ব্র্যান্ড স্ট্যাটাস এবং খুচরা যন্ত্রাংশ প্রাপ্যতা সমস্যা এর আবেদন সীমিত করে। অতি-বাজেট চীনা অপশনের চেয়ে ভালো কিন্তু হন্ডা/ইয়ামাহা নির্ভরযোগ্যতার সাথে তুলনীয় নয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চীনা ব্র্যান্ডের চেয়ে ভালো বাজেট স্কুটারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for SYM Symphony ST already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for SYM Symphony ST\n")

	return nil
}
