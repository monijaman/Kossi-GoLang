package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiBurgmanStreetReviewSeeder handles seeding Suzuki Burgman Street product review and translation
type SuzukiBurgmanStreetReviewSeeder struct {
	BaseSeeder
}

// NewSuzukiBurgmanStreetReviewSeeder creates a new SuzukiBurgmanStreetReviewSeeder
func NewSuzukiBurgmanStreetReviewSeeder() *SuzukiBurgmanStreetReviewSeeder {
	return &SuzukiBurgmanStreetReviewSeeder{BaseSeeder: BaseSeeder{name: "Suzuki Burgman Street Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiBurgmanStreetReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Burgman Street").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Burgman Street product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Burgman Street product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Burgman Street already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Burgman Street Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Burgman Street is a maxi-style 125cc scooter priced at ৳2,00,000, offering premium sporty design, excellent features, and good performance. With LED lights, digital console, and comfortable ride, it's ideal for young enthusiasts seeking a distinctive sporty scooter that stands out.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Burgman Street Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Maxi-Scooter Design:</strong> <span class="highlight-value">Unique sporty styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Full LED Lights:</strong> <span class="highlight-value">Premium lighting system</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Console:</strong> <span class="highlight-value">Fully digital instrument cluster</span></li>
<li class="highlight-item"><strong class="highlight-label">125cc Power:</strong> <span class="highlight-value">8.7 HP peppy performance</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Burgman Street Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Unique Design:</strong> <span class="pro-description">Eye-catching maxi-scooter looks</span></li>
<li class="pro-item"><strong class="pro-label">Full LED Package:</strong> <span class="pro-description">LED headlight, DRL, taillights</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">Peppy 125cc engine with 8.7 HP</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Plush seat and good suspension</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">Digital console, USB charging</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Reliability:</strong> <span class="pro-description">Japanese proven reliability</span></li>
<li class="pro-item"><strong class="pro-label">Good Storage:</strong> <span class="pro-description">21.5L under-seat storage</span></li>
<li class="pro-item"><strong class="pro-label">Large Wheels:</strong> <span class="pro-description">12-inch front wheel for stability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Burgman Street Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Expensive:</strong> <span class="con-description">₹30,000 more than Activa</span></li>
<li class="con-item"><strong class="con-label">Lower Mileage:</strong> <span class="con-description">38-42 km/l, heavy weight</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Fewer Suzuki service centers</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">108 kg, difficult to maneuver</span></li>
<li class="con-item"><strong class="con-label">Lower Resale:</strong> <span class="con-description">Resale value not as good</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Burgman Street in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young bike enthusiasts</li>
<li class="best-for-item">Those wanting unique sporty scooter</li>
<li class="best-for-item">Riders preferring standout design</li>
<li class="best-for-item">Urban premium buyers</li>
<li class="best-for-item">Feature-loving riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Burgman Street?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Women needing light scooter</li>
<li class="not-recommended-item">Buyers wanting better resale</li>
<li class="not-recommended-item">Areas without Suzuki service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Burgman Street Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,95,000 - ৳2,05,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳3,000 - ৳3,800 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,300-2,800/month for 30 km daily at 40 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳900 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Burgman Street Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Lower mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Suzuki quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Premium priced</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Good 125cc power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Excellent features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Unique maxi-styling</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Burgman Street Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Suzuki Burgman Street?</h3>
<p class="faq-answer">Suzuki Burgman Street delivers 38-42 km/l with 125cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Suzuki Burgman Street?</h3>
<p class="faq-answer">Suzuki Burgman Street is priced between ৳1,95,000 to ৳2,05,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Burgman Street good for daily commute?</h3>
<p class="faq-answer">Yes, with comfortable ride and good features, but heavy and expensive.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Burgman Street or Activa?</h3>
<p class="faq-answer">Burgman offers better features and styling; Activa is more practical with better mileage and resale.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Burgman Street in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Suzuki Burgman Street at ৳2,00,000 is the perfect choice for young enthusiasts wanting a unique sporty maxi-scooter. With eye-catching design, full LED package, excellent features, and good performance, it's ideal for riders prioritizing style and distinctiveness. However, the premium price (₹30,000 more than Activa), lower mileage (38-42 km/l), and heavy weight make it suitable only for those who value unique design over practicality and fuel economy.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For unique sporty maxi-scooter with premium features</span></p>
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
		return fmt.Errorf("error creating Suzuki Burgman Street review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Burgman Street (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি বার্গম্যান স্ট্রিট রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি বার্গম্যান স্ট্রিট একটি ম্যাক্সি-স্টাইল ১২৫সিসি স্কুটার যার মূল্য ৳২,০০,০০০ টাকা, যা প্রিমিয়াম স্পোর্টি ডিজাইন, চমৎকার ফিচার এবং ভালো পারফরম্যান্স প্রদান করে। এলইডি লাইট, ডিজিটাল কনসোল এবং আরামদায়ক রাইড সহ, এটি একটি স্বতন্ত্র স্পোর্টি স্কুটার খোঁজা তরুণ উৎসাহীদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি বার্গম্যান স্ট্রিট এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ম্যাক্সি-স্কুটার ডিজাইন:</strong> <span class="highlight-value">অনন্য স্পোর্টি স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ এলইডি লাইট:</strong> <span class="highlight-value">প্রিমিয়াম লাইটিং সিস্টেম</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল কনসোল:</strong> <span class="highlight-value">সম্পূর্ণ ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
<li class="highlight-item"><strong class="highlight-label">১২৫সিসি পাওয়ার:</strong> <span class="highlight-value">৮.৭ এইচপি প্রাণবন্ত পারফরম্যান্স</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি বার্গম্যান স্ট্রিট এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অনন্য ডিজাইন:</strong> <span class="pro-description">নজরকাড়া ম্যাক্সি-স্কুটার লুক</span></li>
<li class="pro-item"><strong class="pro-label">সম্পূর্ণ এলইডি প্যাকেজ:</strong> <span class="pro-description">এলইডি হেডলাইট, ডিআরএল, টেইললাইট</span></li>
<li class="pro-item"><strong class="pro-label">ভালো পারফরম্যান্স:</strong> <span class="pro-description">প্রাণবন্ত ১২৫সিসি ইঞ্জিন ৮.৭ এইচপি সহ</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">আরামদায়ক সিট এবং ভালো সাসপেনশন</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক ফিচার:</strong> <span class="pro-description">ডিজিটাল কনসোল, ইউএসবি চার্জিং</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি নির্ভরযোগ্যতা:</strong> <span class="pro-description">জাপানি প্রমাণিত নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">ভালো স্টোরেজ:</strong> <span class="pro-description">২১.৫ লিটার সিটের নিচে স্টোরেজ</span></li>
<li class="pro-item"><strong class="pro-label">বড় চাকা:</strong> <span class="pro-description">স্থিতিশীলতার জন্য ১২-ইঞ্চি সামনের চাকা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি বার্গম্যান স্ট্রিট এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত ব্যয়বহুল:</strong> <span class="con-description">অ্যাকটিভার চেয়ে ৳৩০,০০০ টাকা বেশি</span></li>
<li class="con-item"><strong class="con-label">কম মাইলেজ:</strong> <span class="con-description">৩৮-৪২ কিমি/লিটার, ভারী ওজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">কম সুজুকি সার্ভিস সেন্টার</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">১০৮ কেজি, চালনা কঠিন</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল:</strong> <span class="con-description">রিসেল ভ্যালু ততটা ভালো নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি বার্গম্যান স্ট্রিট কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young bike enthusiasts</li>
<li class="best-for-item">Those wanting unique sporty scooter</li>
<li class="best-for-item">Riders preferring standout design</li>
<li class="best-for-item">Urban premium buyers</li>
<li class="best-for-item">Feature-loving riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি বার্গম্যান স্ট্রিট কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Women needing light scooter</li>
<li class="not-recommended-item">Buyers wanting better resale</li>
<li class="not-recommended-item">Areas without Suzuki service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি বার্গম্যান স্ট্রিট এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৯৫,০০০ - ৳২,০৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৩,০০০ - ৳৩,৮০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪০ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,৩০০-২,৮০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৯০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি বার্গম্যান স্ট্রিট পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- কম মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- সুজুকি মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- প্রিমিয়াম মূল্যের</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- ভালো ১২৫সিসি পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- চমৎকার ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অনন্য ম্যাক্সি-স্টাইলিং</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি বার্গম্যান স্ট্রিট সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">সুজুকি বার্গম্যান স্ট্রিটের মাইলেজ কত?</h3>
<p class="faq-answer">সুজুকি বার্গম্যান স্ট্রিট ১২৫সিসি ইঞ্জিন সহ ৩৮-৪২ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি বার্গম্যান স্ট্রিটের দাম কত?</h3>
<p class="faq-answer">সুজুকি বার্গম্যান স্ট্রিটের দাম ৳১,৯৫,০০০ থেকে ৳২,০৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনিক যাতায়াতের জন্য বার্গম্যান স্ট্রিট কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, আরামদায়ক রাইড এবং ভালো ফিচার সহ, তবে ভারী এবং ব্যয়বহুল।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: বার্গম্যান স্ট্রিট নাকি অ্যাকটিভা?</h3>
<p class="faq-answer">বার্গম্যান ভালো ফিচার এবং স্টাইলিং প্রদান করে; অ্যাকটিভা ভালো মাইলেজ এবং রিসেল সহ আরও ব্যবহারিক।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি বার্গম্যান স্ট্রিট কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,০০,০০০ টাকায় সুজুকি বার্গম্যান স্ট্রিট একটি অনন্য স্পোর্টি ম্যাক্সি-স্কুটার চাওয়া তরুণ উৎসাহীদের জন্য নিখুঁত পছন্দ। নজরকাড়া ডিজাইন, সম্পূর্ণ এলইডি প্যাকেজ, চমৎকার ফিচার এবং ভালো পারফরম্যান্স সহ, এটি স্টাইল এবং স্বতন্ত্রতা অগ্রাধিকার দেওয়া রাইডারদের জন্য আদর্শ। তবে, প্রিমিয়াম মূল্য (অ্যাকটিভার চেয়ে ৳৩০,০০০ টাকা বেশি), কম মাইলেজ (৩৮-৪২ কিমি/লিটার) এবং ভারী ওজন এটিকে শুধুমাত্র তাদের জন্য উপযুক্ত করে যারা ব্যবহারিকতা এবং জ্বালানি অর্থনীতির চেয়ে অনন্য ডিজাইনকে মূল্য দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম ফিচার সহ অনন্য স্পোর্টি ম্যাক্সি-স্কুটারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Burgman Street already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Burgman Street\n")

	return nil
}
