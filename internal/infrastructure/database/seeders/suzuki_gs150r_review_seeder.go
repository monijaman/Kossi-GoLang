package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGS150RReviewSeeder handles seeding Suzuki GS150R product review and translation
type SuzukiGS150RReviewSeeder struct {
	BaseSeeder
}

// NewSuzukiGS150RReviewSeeder creates a new SuzukiGS150RReviewSeeder
func NewSuzukiGS150RReviewSeeder() *SuzukiGS150RReviewSeeder {
	return &SuzukiGS150RReviewSeeder{BaseSeeder: BaseSeeder{name: "Suzuki GS150R Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGS150RReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki GS150R").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki GS150R product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki GS150R product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki GS150R already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki GS150R Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki GS150R is a sporty 147cc commuter bike priced at ৳2,10,000, offering good performance, decent styling, and Suzuki reliability. With digital console, electric start, and balanced performance, it's ideal for riders seeking a sporty yet practical motorcycle for daily urban commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki GS150R Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Good Power:</strong> <span class="highlight-value">13.8 HP from 147cc</span></li>
<li class="highlight-item"><strong class="highlight-label">Sporty Design:</strong> <span class="highlight-value">Aggressive commuter styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Console:</strong> <span class="highlight-value">Modern instrument cluster</span></li>
<li class="highlight-item"><strong class="highlight-label">Suzuki Quality:</strong> <span class="highlight-value">Japanese reliability</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki GS150R Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">13.8 HP, decent power</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Styling:</strong> <span class="pro-description">Aggressive design elements</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Reliability:</strong> <span class="pro-description">Japanese build quality</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Modern instrument display</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Convenient starting</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">45-50 km/l efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Balanced ergonomics</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki GS150R Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">More expensive than competitors</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Fewer Suzuki centers</span></li>
<li class="con-item"><strong class="con-label">Average Build:</strong> <span class="con-description">Not premium build quality</span></li>
<li class="con-item"><strong class="con-label">Outdated Features:</strong> <span class="con-description">No ABS or modern features</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki GS150R in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban commuters</li>
<li class="best-for-item">Those wanting 150cc performance</li>
<li class="best-for-item">Suzuki brand fans</li>
<li class="best-for-item">Balanced performance seekers</li>
<li class="best-for-item">Daily office riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki GS150R?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget buyers</li>
<li class="not-recommended-item">Premium feature seekers</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Performance enthusiasts</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki GS150R Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,05,000 - ৳2,15,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,800 - ৳3,600 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,900-2,400/month for 30 km daily at 47 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 5 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,000 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki GS150R Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good 47 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Suzuki quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good 150cc power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Sporty looks</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki GS150R Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Suzuki GS150R?</h3>
<p class="faq-answer">Suzuki GS150R delivers 45-50 km/l with 147cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Suzuki GS150R?</h3>
<p class="faq-answer">Suzuki GS150R is priced between ৳2,05,000 to ৳2,15,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is GS150R good for daily commute?</h3>
<p class="faq-answer">Yes, GS150R offers balanced performance and comfort for daily urban commuting.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does GS150R have kick start?</h3>
<p class="faq-answer">GS150R comes with electric start, kick start availability varies by model.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki GS150R in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Suzuki GS150R at ৳2,10,000 is a decent choice for balanced commuting, offering good 147cc performance (13.8 HP), sporty styling, digital console, and Suzuki reliability. With 45-50 km/l mileage and comfortable ergonomics, it serves daily urban needs well. However, higher pricing compared to competitors, limited service network, average build quality, and lack of modern features like ABS limit its appeal. Ideal for those specifically preferring Suzuki brand and seeking balanced 150cc performance for city commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For balanced Suzuki 150cc performance</span></p>
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
		return fmt.Errorf("error creating Suzuki GS150R review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki GS150R (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি জিএস১৫০আর রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি জিএস১৫০আর একটি স্পোর্টি ১৪৭সিসি কমিউটার বাইক যার মূল্য ৳২,১০,০০০ টাকা, যা ভালো পারফরম্যান্স, মোটামুটি স্টাইলিং এবং সুজুকি নির্ভরযোগ্যতা প্রদান করে। ডিজিটাল কনসোল, ইলেকট্রিক স্টার্ট এবং সুষম পারফরম্যান্স সহ, এটি দৈনিক শহুরে যাতায়াতের জন্য স্পোর্টি তবুও ব্যবহারিক মোটরসাইকেল খোঁজা রাইডারদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি জিএস১৫০আর এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ভালো পাওয়ার:</strong> <span class="highlight-value">১৪৭সিসি থেকে ১৩.৮ এইচপি</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্টি ডিজাইন:</strong> <span class="highlight-value">আক্রমণাত্মক কমিউটার স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল কনসোল:</strong> <span class="highlight-value">আধুনিক ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
<li class="highlight-item"><strong class="highlight-label">সুজুকি মান:</strong> <span class="highlight-value">জাপানি নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি জিএস১৫০আর এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ভালো পারফরম্যান্স:</strong> <span class="pro-description">১৩.৮ এইচপি, মোটামুটি পাওয়ার</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি স্টাইলিং:</strong> <span class="pro-description">আক্রমণাত্মক ডিজাইন উপাদান</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি নির্ভরযোগ্যতা:</strong> <span class="pro-description">জাপানি বিল্ড কোয়ালিটি</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">আধুনিক ইন্সট্রুমেন্ট ডিসপ্লে</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">সুবিধাজনক স্টার্টিং</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৪৫-৫০ কিমি/লিটার দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">সুষম এরগোনমিক্স</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি জিএস১৫০আর এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বেশি দাম:</strong> <span class="con-description">প্রতিযোগীদের চেয়ে ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">কম সুজুকি সেন্টার</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড:</strong> <span class="con-description">প্রিমিয়াম বিল্ড কোয়ালিটি নয়</span></li>
<li class="con-item"><strong class="con-label">পুরানো ফিচার:</strong> <span class="con-description">কোনো এবিএস বা আধুনিক ফিচার নেই</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি জিএস১৫০আর কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban commuters</li>
<li class="best-for-item">Those wanting 150cc performance</li>
<li class="best-for-item">Suzuki brand fans</li>
<li class="best-for-item">Balanced performance seekers</li>
<li class="best-for-item">Daily office riders</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি জিএস১৫০আর কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget buyers</li>
<li class="not-recommended-item">Premium feature seekers</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Performance enthusiasts</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি জিএস১৫০আর এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,০৫,০০০ - ৳২,১৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৮০০ - ৳৩,৬০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৯০০-২,৪০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৫ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,০০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি জিএস১৫০আর পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ৪৭ কিমি/লিটার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- সুজুকি মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ১৫০সিসি পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- প্রাথমিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- স্পোর্টি লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি জিএস১৫০আর সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">সুজুকি জিএস১৫০আর এর মাইলেজ কত?</h3>
<p class="faq-answer">সুজুকি জিএস১৫০আর ১৪৭সিসি ইঞ্জিন সহ ৪৫-৫০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি জিএস১৫০আর এর দাম কত?</h3>
<p class="faq-answer">সুজুকি জিএস১৫০আর এর দাম ৳২,০৫,০০০ থেকে ৳২,১৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনিক যাতায়াতের জন্য জিএস১৫০আর কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, জিএস১৫০আর দৈনিক শহুরে যাতায়াতের জন্য সুষম পারফরম্যান্স এবং আরাম প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">জিএস১৫০আর এ কি কিক স্টার্ট আছে?</h3>
<p class="faq-answer">জিএস১৫০আর ইলেকট্রিক স্টার্ট সহ আসে, কিক স্টার্ট উপলব্ধতা মডেল অনুযায়ী পরিবর্তিত হয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি জিএস১৫০আর কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,১০,০০০ টাকায় সুজুকি জিএস১৫০আর সুষম যাতায়াতের জন্য একটি মোটামুটি পছন্দ, যা ভালো ১৪৭সিসি পারফরম্যান্স (১৩.৮ এইচপি), স্পোর্টি স্টাইলিং, ডিজিটাল কনসোল এবং সুজুকি নির্ভরযোগ্যতা প্রদান করে। ৪৫-৫০ কিমি/লিটার মাইলেজ এবং আরামদায়ক এরগোনমিক্স সহ, এটি দৈনিক শহুরে প্রয়োজন ভালভাবে পূরণ করে। তবে, প্রতিযোগীদের তুলনায় বেশি মূল্য, সীমিত সার্ভিস নেটওয়ার্ক, গড় বিল্ড কোয়ালিটি এবং এবিএসের মতো আধুনিক ফিচারের অভাব এর আবেদন সীমিত করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সুষম সুজুকি ১৫০সিসি পারফরম্যান্সের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki GS150R already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki GS150R\n")

	return nil
}
