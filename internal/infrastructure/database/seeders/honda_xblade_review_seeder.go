package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaXBladeReviewSeeder handles seeding Honda XBlade product review and translation
type HondaXBladeReviewSeeder struct {
	BaseSeeder
}

// NewHondaXBladeReviewSeeder creates a new HondaXBladeReviewSeeder
func NewHondaXBladeReviewSeeder() *HondaXBladeReviewSeeder {
	return &HondaXBladeReviewSeeder{BaseSeeder: BaseSeeder{name: "Honda XBlade Review"}}
}

// Seed implements the Seeder interface
func (s *HondaXBladeReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda XBlade").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda XBlade product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda XBlade product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda XBlade already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda XBlade Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda XBlade is a sporty 163cc commuter bike priced at ৳2,20,000, offering aggressive styling, good performance, and Honda reliability. With LED headlight, digital console, and comfortable ergonomics, it's ideal for young riders seeking a sporty yet practical motorcycle for daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda XBlade Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Sharp sporty design</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Headlight:</strong> <span class="highlight-value">Modern LED lighting</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Power:</strong> <span class="highlight-value">13.9 HP from 163cc</span></li>
<li class="highlight-item"><strong class="highlight-label">Honda Quality:</strong> <span class="highlight-value">Proven Honda reliability</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda XBlade Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Sporty Design:</strong> <span class="pro-description">Aggressive and modern styling</span></li>
<li class="pro-item"><strong class="pro-label">LED Headlight:</strong> <span class="pro-description">Bright LED headlight</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">13.9 HP, peppy engine</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven Honda build quality</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Fully digital instrument cluster</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Good ergonomics for daily use</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">45-50 km/l fuel efficiency</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda XBlade Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">₹40,000 more than 150cc bikes</span></li>
<li class="con-item"><strong class="con-label">Single Cylinder:</strong> <span class="con-description">Still only single cylinder</span></li>
<li class="con-item"><strong class="con-label">Limited Features:</strong> <span class="con-description">No ABS or advanced features</span></li>
<li class="con-item"><strong class="con-label">Competition:</strong> <span class="con-description">Fierce competition in segment</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda XBlade in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban commuters</li>
<li class="best-for-item">Those wanting sporty 160cc bike</li>
<li class="best-for-item">Honda brand loyalists</li>
<li class="best-for-item">Daily office commuters</li>
<li class="best-for-item">First-time bike buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda XBlade?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget buyers</li>
<li class="not-recommended-item">Those wanting advanced features</li>
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Long-distance touring</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda XBlade Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,15,000 - ৳2,25,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,800 - ৳3,500 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,900-2,400/month for 30 km daily at 47 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳900 - ৳1,400 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda XBlade Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Honda reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Slightly expensive</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good 160cc power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good Honda quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Basic features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very sporty</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda XBlade Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Honda XBlade?</h3>
<p class="faq-answer">Honda XBlade delivers 45-50 km/l with 163cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Honda XBlade?</h3>
<p class="faq-answer">Honda XBlade is priced between ৳2,15,000 to ৳2,25,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is XBlade good for daily commute?</h3>
<p class="faq-answer">Yes, XBlade offers good balance of sporty looks and practical commuting.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does XBlade have ABS?</h3>
<p class="faq-answer">No, Honda XBlade doesn't come with ABS feature.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda XBlade in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Honda XBlade at ৳2,20,000 is a solid choice for young commuters wanting sporty styling with Honda reliability. With aggressive design, LED headlight, good 163cc performance (13.9 HP), and excellent mileage (45-50 km/l), it strikes a nice balance between looks and practicality. While slightly expensive for the segment and lacking ABS, the Honda brand value, widespread service network, and proven reliability make it worthy consideration. Ideal for those prioritizing sporty aesthetics with trusted Honda dependability for daily urban commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sporty commuter with Honda reliability</span></p>
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
		return fmt.Errorf("error creating Honda XBlade review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda XBlade (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হন্ডা এক্সব্লেড রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হন্ডা এক্সব্লেড একটি স্পোর্টি ১৬৩সিসি কমিউটার বাইক যার মূল্য ৳২,২০,০০০ টাকা, যা আক্রমণাত্মক স্টাইলিং, ভালো পারফরম্যান্স এবং হন্ডা নির্ভরযোগ্যতা প্রদান করে। এলইডি হেডলাইট, ডিজিটাল কনসোল এবং আরামদায়ক এরগোনমিক্স সহ, এটি দৈনিক যাতায়াতের জন্য স্পোর্টি তবুও ব্যবহারিক মোটরসাইকেল খোঁজা তরুণ রাইডারদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা এক্সব্লেড এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="highlight-value">তীক্ষ্ণ স্পোর্টি ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি হেডলাইট:</strong> <span class="highlight-value">আধুনিক এলইডি লাইটিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো পাওয়ার:</strong> <span class="highlight-value">১৬৩সিসি থেকে ১৩.৯ এইচপি</span></li>
<li class="highlight-item"><strong class="highlight-label">হন্ডা মান:</strong> <span class="highlight-value">প্রমাণিত হন্ডা নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা এক্সব্লেড এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">স্পোর্টি ডিজাইন:</strong> <span class="pro-description">আক্রমণাত্মক এবং আধুনিক স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি হেডলাইট:</strong> <span class="pro-description">উজ্জ্বল এলইডি হেডলাইট</span></li>
<li class="pro-item"><strong class="pro-label">ভালো পারফরম্যান্স:</strong> <span class="pro-description">১৩.৯ এইচপি, প্রাণবন্ত ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">হন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত হন্ডা বিল্ড কোয়ালিটি</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">সম্পূর্ণ ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">দৈনিক ব্যবহারের জন্য ভালো এরগোনমিক্স</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৪৫-৫০ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা এক্সব্লেড এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বেশি দাম:</strong> <span class="con-description">১৫০সিসি বাইকের চেয়ে ৳৪০,০০০ টাকা বেশি</span></li>
<li class="con-item"><strong class="con-label">সিঙ্গেল সিলিন্ডার:</strong> <span class="con-description">এখনও শুধু সিঙ্গেল সিলিন্ডার</span></li>
<li class="con-item"><strong class="con-label">সীমিত ফিচার:</strong> <span class="con-description">কোনো এবিএস বা উন্নত ফিচার নেই</span></li>
<li class="con-item"><strong class="con-label">প্রতিযোগিতা:</strong> <span class="con-description">সেগমেন্টে তীব্র প্রতিযোগিতা</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হন্ডা এক্সব্লেড কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban commuters</li>
<li class="best-for-item">Those wanting sporty 160cc bike</li>
<li class="best-for-item">Honda brand loyalists</li>
<li class="best-for-item">Daily office commuters</li>
<li class="best-for-item">First-time bike buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হন্ডা এক্সব্লেড কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget buyers</li>
<li class="not-recommended-item">Those wanting advanced features</li>
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Long-distance touring</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হন্ডা এক্সব্লেড এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,১৫,০০০ - ৳২,২৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৮০০ - ৳৩,৫০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৯০০-২,৪০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৯০০ - ৳১,৪০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হন্ডা এক্সব্লেড পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- হন্ডা নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- সামান্য ব্যয়বহুল</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ১৬০সিসি পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো হন্ডা মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- প্রাথমিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব স্পোর্টি</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হন্ডা এক্সব্লেড সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হন্ডা এক্সব্লেডের মাইলেজ কত?</h3>
<p class="faq-answer">হন্ডা এক্সব্লেড ১৬৩সিসি ইঞ্জিন সহ ৪৫-৫০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হন্ডা এক্সব্লেডের দাম কত?</h3>
<p class="faq-answer">হন্ডা এক্সব্লেডের দাম ৳২,১৫,০০০ থেকে ৳২,২৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনিক যাতায়াতের জন্য এক্সব্লেড কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, এক্সব্লেড স্পোর্টি লুক এবং ব্যবহারিক যাতায়াতের ভালো ভারসাম্য প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এক্সব্লেডে কি এবিএস আছে?</h3>
<p class="faq-answer">না, হন্ডা এক্সব্লেড এবিএস ফিচার সহ আসে না।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হন্ডা এক্সব্লেড কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.1/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,২০,০০০ টাকায় হন্ডা এক্সব্লেড হন্ডা নির্ভরযোগ্যতা সহ স্পোর্টি স্টাইলিং চাওয়া তরুণ কমিউটারদের জন্য একটি শক্ত পছন্দ। আক্রমণাত্মক ডিজাইন, এলইডি হেডলাইট, ভালো ১৬৩সিসি পারফরম্যান্স (১৩.৯ এইচপি) এবং চমৎকার মাইলেজ (৪৫-৫০ কিমি/লিটার) সহ, এটি লুক এবং ব্যবহারিকতার মধ্যে একটি চমৎকার ভারসাম্য তৈরি করে। যদিও সেগমেন্টের জন্য সামান্য ব্যয়বহুল এবং এবিএস অনুপস্থিত, হন্ডা ব্র্যান্ড ভ্যালু, ব্যাপক সার্ভিস নেটওয়ার্ক এবং প্রমাণিত নির্ভরযোগ্যতা এটিকে বিবেচনার যোগ্য করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- হন্ডা নির্ভরযোগ্যতা সহ স্পোর্টি কমিউটারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda XBlade already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda XBlade\n")

	return nil
}
