package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsar125ReviewSeeder handles seeding Bajaj Pulsar 125 product review and translation
type BajajPulsar125ReviewSeeder struct {
	BaseSeeder
}

// NewBajajPulsar125ReviewSeeder creates a new BajajPulsar125ReviewSeeder
func NewBajajPulsar125ReviewSeeder() *BajajPulsar125ReviewSeeder {
	return &BajajPulsar125ReviewSeeder{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar 125 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsar125ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Pulsar 125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Pulsar 125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar 125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar 125 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Pulsar 125 is a sporty 125cc motorcycle priced at ৳1,70,000, bringing Pulsar's muscular styling to the commuter segment. With split seats, sporty graphics, and decent performance, it's perfect for young riders wanting Pulsar DNA without the 150cc price tag.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar 125 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Pulsar Styling:</strong> <span class="highlight-value">Iconic Pulsar muscular design</span></li>
<li class="highlight-item"><strong class="highlight-label">Split Seats:</strong> <span class="highlight-value">Sporty split seat design</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">50-55 km/l efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Disc Brake:</strong> <span class="highlight-value">Front disc brake standard</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar 125 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Pulsar Brand Value:</strong> <span class="pro-description">Gets iconic Pulsar styling and brand</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Design:</strong> <span class="pro-description">Muscular styling with split seats</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">Peppy 125cc engine with good power</span></li>
<li class="pro-item"><strong class="pro-label">Disc Brake:</strong> <span class="pro-description">Front disc brake for better stopping</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">50-55 km/l better than 150cc</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Bajaj service centers widely available</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Semi-digital instrument cluster</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar 125 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Plastic quality could be better</span></li>
<li class="con-item"><strong class="con-label">Vibrations:</strong> <span class="con-description">Engine vibrates at high RPM</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">No LED lights or modern tech</span></li>
<li class="con-item"><strong class="con-label">Limited Power:</strong> <span class="con-description">125cc not enough for highway</span></li>
<li class="con-item"><strong class="con-label">Firm Seat:</strong> <span class="con-description">Split seat less comfortable than regular</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Bajaj Pulsar 125 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young riders wanting sporty style</li>
<li class="best-for-item">Daily city commuters</li>
<li class="best-for-item">Students and college goers</li>
<li class="best-for-item">Those wanting Pulsar brand on budget</li>
<li class="best-for-item">First-time sporty bike buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Bajaj Pulsar 125?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway long-distance riders</li>
<li class="not-recommended-item">Those needing maximum comfort</li>
<li class="not-recommended-item">Riders wanting premium features</li>
<li class="not-recommended-item">Those prioritizing build quality</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar 125 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,65,000 - ৳1,75,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,300 - ৳2,800 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,700-2,100/month for 30 km daily at 52 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳700 - ৳1,000 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar 125 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Reliable Bajaj</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent for 125cc</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Sporty Pulsar look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Bajaj Pulsar 125 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Bajaj Pulsar 125?</h3>
<p class="faq-answer">Bajaj Pulsar 125 delivers 50-55 km/l.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Pulsar 125?</h3>
<p class="faq-answer">Pulsar 125 is priced between ৳1,65,000 to ৳1,75,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Pulsar 125 or 150?</h3>
<p class="faq-answer">Pulsar 150 has more power, while 125 offers better mileage and lower price.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Pulsar 125 good for daily use?</h3>
<p class="faq-answer">Yes, it's excellent for daily city commuting with sporty style.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar 125 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.0/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Bajaj Pulsar 125 brings iconic Pulsar styling to the 125cc segment at ৳1,70,000. Perfect for young riders wanting sporty looks with commuter efficiency. The muscular design, disc brake, and decent performance make it appealing. However, average build quality and basic features are compromises. Good choice for style-conscious buyers on budget, but Honda CB Shine offers better overall quality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For young style-conscious commuters</span></p>
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
		return fmt.Errorf("error creating Bajaj Pulsar 125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar 125 (Rating: 4.0/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার ১২৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ পালসার ১২৫ একটি স্পোর্টি ১২৫সিসি মোটরসাইকেল যার মূল্য ৳১,৭০,০০০ টাকা, যা পালসারের পেশীবহুল স্টাইলিং কমিউটার সেগমেন্টে নিয়ে আসে। স্প্লিট সিট, স্পোর্টি গ্রাফিক্স এবং ভালো পারফরম্যান্স সহ, এটি ১৫০সিসি মূল্য ট্যাগ ছাড়াই পালসার ডিএনএ চান তরুণ রাইডারদের জন্য নিখুঁত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার ১২৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">পালসার স্টাইলিং:</strong> <span class="highlight-value">আইকনিক পালসার পেশীবহুল ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">স্প্লিট সিট:</strong> <span class="highlight-value">স্পোর্টি স্প্লিট সিট ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৫০-৫৫ কিমি/লিটার দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিস্ক ব্রেক:</strong> <span class="highlight-value">ফ্রন্ট ডিস্ক ব্রেক স্ট্যান্ডার্ড</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার ১২৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">পালসার ব্র্যান্ড ভ্যালু:</strong> <span class="pro-description">আইকনিক পালসার স্টাইলিং এবং ব্র্যান্ড পায়</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি ডিজাইন:</strong> <span class="pro-description">স্প্লিট সিট সহ পেশীবহুল স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">ভালো পারফরম্যান্স:</strong> <span class="pro-description">ভালো পাওয়ার সহ উদ্যমী ১২৫সিসি ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">ডিস্ক ব্রেক:</strong> <span class="pro-description">ভালো স্টপিংয়ের জন্য ফ্রন্ট ডিস্ক ব্রেক</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">১৫০সিসির চেয়ে ৫০-৫৫ কিমি/লিটার ভালো</span></li>
<li class="pro-item"><strong class="pro-label">ব্যাপক সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">বাজাজ সার্ভিস সেন্টার ব্যাপকভাবে উপলব্ধ</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">সেমি-ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার ১২৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">প্লাস্টিকের গুণমান ভালো হতে পারত</span></li>
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">উচ্চ আরপিএমে ইঞ্জিন কম্পন করে</span></li>
<li class="con-item"><strong class="con-label">বেসিক ফিচার:</strong> <span class="con-description">কোন এলইডি লাইট বা আধুনিক প্রযুক্তি নেই</span></li>
<li class="con-item"><strong class="con-label">সীমিত পাওয়ার:</strong> <span class="con-description">হাইওয়ের জন্য ১২৫সিসি যথেষ্ট নয়</span></li>
<li class="con-item"><strong class="con-label">শক্ত সিট:</strong> <span class="con-description">স্প্লিট সিট নিয়মিতের চেয়ে কম আরামদায়ক</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বাজাজ পালসার ১২৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young riders wanting sporty style</li>
<li class="best-for-item">Daily city commuters</li>
<li class="best-for-item">Students and college goers</li>
<li class="best-for-item">Those wanting Pulsar brand on budget</li>
<li class="best-for-item">First-time sporty bike buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বাজাজ পালসার ১২৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway long-distance riders</li>
<li class="not-recommended-item">Those needing maximum comfort</li>
<li class="not-recommended-item">Riders wanting premium features</li>
<li class="not-recommended-item">Those prioritizing build quality</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বাজাজ পালসার ১২৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৬৫,০০০ - ৳১,৭৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৩০০ - ৳২,৮০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৭০০-২,১০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৭০০ - ৳১,০০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বাজাজ পালসার ১২৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- নির্ভরযোগ্য বাজাজ</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- ১২৫সিসির জন্য ভালো</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- বেসিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- স্পোর্টি পালসার লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বাজাজ পালসার ১২৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">বাজাজ পালসার ১২৫ এর মাইলেজ কত?</h3>
<p class="faq-answer">বাজাজ পালসার ১২৫ ৫০-৫৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">পালসার ১২৫ এর দাম কত?</h3>
<p class="faq-answer">পালসার ১২৫ এর দাম ৳১,৬৫,০০০ থেকে ৳১,৭৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: পালসার ১২৫ নাকি ১৫০?</h3>
<p class="faq-answer">পালসার ১৫০ এ বেশি পাওয়ার আছে, যেখানে ১২৫ ভালো মাইলেজ এবং কম দাম দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য পালসার ১২৫ কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, এটি স্পোর্টি স্টাইল সহ দৈনিক শহর যাতায়াতের জন্য চমৎকার।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বাজাজ পালসার ১২৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.0/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">বাজাজ পালসার ১২৫ আইকনিক পালসার স্টাইলিং ১২৫সিসি সেগমেন্টে ৳১,৭০,০০০ টাকায় নিয়ে আসে। কমিউটার দক্ষতা সহ স্পোর্টি লুক চান তরুণ রাইডারদের জন্য নিখুঁত। পেশীবহুল ডিজাইন, ডিস্ক ব্রেক এবং ভালো পারফরম্যান্স এটিকে আকর্ষণীয় করে তোলে। তবে, গড় বিল্ড কোয়ালিটি এবং বেসিক ফিচার সমঝোতা।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- তরুণ স্টাইল-সচেতন যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar 125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar 125\n")

	return nil
}
