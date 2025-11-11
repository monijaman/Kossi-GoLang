package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxerReviewSeeder handles seeding Suzuki Gixxer product review and translation
type SuzukiGixxerReviewSeeder struct {
	BaseSeeder
}

// NewSuzukiGixxerReviewSeeder creates a new SuzukiGixxerReviewSeeder
func NewSuzukiGixxerReviewSeeder() *SuzukiGixxerReviewSeeder {
	return &SuzukiGixxerReviewSeeder{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxerReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Gixxer product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Gixxer product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Gixxer is a sporty 155cc naked bike priced at ৳2,85,000, offering sharp handling, powerful engine, and aggressive styling. With rear disc brake and fuel injection, it's an exciting choice for performance-oriented riders seeking thrilling rides at a competitive price.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Excellent cornering and stability</span></li>
<li class="highlight-item"><strong class="highlight-label">Powerful Engine:</strong> <span class="highlight-value">155cc with 14.1 PS power</span></li>
<li class="highlight-item"><strong class="highlight-label">Dual Disc Brakes:</strong> <span class="highlight-value">Front and rear disc brakes</span></li>
<li class="highlight-item"><strong class="highlight-label">Sporty Design:</strong> <span class="highlight-value">Aggressive naked bike styling</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Handling:</strong> <span class="pro-description">Best-in-class handling and cornering</span></li>
<li class="pro-item"><strong class="pro-label">Powerful Performance:</strong> <span class="pro-description">155cc produces 14.1 PS, peppy acceleration</span></li>
<li class="pro-item"><strong class="pro-label">Dual Disc Brakes:</strong> <span class="pro-description">Front and rear disc for excellent stopping</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Muscular naked bike design looks sporty</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Solid Japanese build standards</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Injection:</strong> <span class="pro-description">FI ensures smooth power delivery</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Fully digital LCD meter with gear indicator</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Sporty yet comfortable riding position</span></li>
<li class="pro-item"><strong class="pro-label">MRF Tyres:</strong> <span class="pro-description">Premium MRF tyres provide great grip</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Average Mileage:</strong> <span class="con-description">38-42 km/l not great for fuel economy</span></li>
<li class="con-item"><strong class="con-label">Single Channel ABS:</strong> <span class="con-description">ABS only on front wheel</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Suzuki service centers limited</span></li>
<li class="con-item"><strong class="con-label">Firm Suspension:</strong> <span class="con-description">Stiff setup uncomfortable on bad roads</span></li>
<li class="con-item"><strong class="con-label">Lower Resale Value:</strong> <span class="con-description">Suzuki bikes depreciate faster</span></li>
<li class="con-item"><strong class="con-label">Vibrations at High Speed:</strong> <span class="con-description">Noticeable vibrations above 80 km/h</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Gixxer in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance-focused riders</li>
<li class="best-for-item">Those who love spirited riding</li>
<li class="best-for-item">Corner carvers and twisty road enthusiasts</li>
<li class="best-for-item">Young riders wanting sporty bike</li>
<li class="best-for-item">City riders with occasional highway use</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Gixxer?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Mileage-focused buyers</li>
<li class="not-recommended-item">Long-distance highway tourers</li>
<li class="not-recommended-item">Those needing maximum comfort</li>
<li class="not-recommended-item">Areas with limited Suzuki service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,80,000 - ৳2,90,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳3,500 - ৳4,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,500-2,900/month for 30 km daily at 40 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 4 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,000 - ৳1,300 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Below average mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good Suzuki quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good performance value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Sporty but adequate</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Solid build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- FI + dual disc + digital</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Sporty aggressive</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Gixxer Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Suzuki Gixxer?</h3>
<p class="faq-answer">Suzuki Gixxer delivers 38-42 km/l with spirited riding.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Suzuki Gixxer in Bangladesh?</h3>
<p class="faq-answer">Suzuki Gixxer is priced between ৳2,80,000 to ৳2,90,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Suzuki Gixxer have ABS?</h3>
<p class="faq-answer">Yes, Suzuki Gixxer has single channel ABS on the front wheel.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Suzuki Gixxer?</h3>
<p class="faq-answer">Suzuki Gixxer can reach approximately 125-130 km/h.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Suzuki Gixxer good for daily use?</h3>
<p class="faq-answer">Yes, it's excellent for daily use with sporty character and good handling.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Suzuki Gixxer or Yamaha FZS?</h3>
<p class="faq-answer">Gixxer has better handling and performance, while FZS has better features and styling.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Suzuki Gixxer have fuel injection?</h3>
<p class="faq-answer">Yes, Suzuki Gixxer comes with fuel injection system.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the engine capacity of Suzuki Gixxer?</h3>
<p class="faq-answer">Suzuki Gixxer has a 155cc single-cylinder air-cooled engine producing 14.1 PS.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Suzuki Gixxer is the handling king in the 150-160cc segment, offering thrilling performance and sharp dynamics at ৳2,85,000. Best for riders who prioritize sporty riding experience over fuel efficiency. While mileage is average and service network limited, the exceptional handling and powerful engine make it a compelling choice for performance enthusiasts. However, Honda and Yamaha offer better overall package for daily commuters.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance enthusiasts loving twisty roads</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Suzuki Gixxer review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি গিক্সার একটি স্পোর্টি ১৫৫সিসি নেকেড বাইক যার মূল্য ৳২,৮৫,০০০ টাকা, যা তীক্ষ্ণ হ্যান্ডলিং, শক্তিশালী ইঞ্জিন এবং আগ্রাসী স্টাইলিং প্রদান করে। রিয়ার ডিস্ক ব্রেক এবং ফুয়েল ইনজেকশন সহ, এটি প্রতিযোগিতামূলক মূল্যে রোমাঞ্চকর রাইড খুঁজছেন পারফরম্যান্স-ভিত্তিক রাইডারদের জন্য একটি উত্তেজনাপূর্ণ পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="highlight-value">চমৎকার কর্নারিং এবং স্থিতিশীলতা</span></li>
<li class="highlight-item"><strong class="highlight-label">শক্তিশালী ইঞ্জিন:</strong> <span class="highlight-value">১৪.১ পিএস পাওয়ার সহ ১৫৫সিসি</span></li>
<li class="highlight-item"><strong class="highlight-label">ডুয়াল ডিস্ক ব্রেক:</strong> <span class="highlight-value">সামনে এবং পিছনে ডিস্ক ব্রেক</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্টি ডিজাইন:</strong> <span class="highlight-value">আগ্রাসী নেকেড বাইক স্টাইলিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার হ্যান্ডলিং:</strong> <span class="pro-description">ক্লাসে সেরা হ্যান্ডলিং এবং কর্নারিং</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী পারফরম্যান্স:</strong> <span class="pro-description">১৫৫সিসি ১৪.১ পিএস উৎপাদন করে, দ্রুত এক্সিলারেশন</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল ডিস্ক ব্রেক:</strong> <span class="pro-description">চমৎকার স্টপিংয়ের জন্য সামনে এবং পিছনে ডিস্ক</span></li>
<li class="pro-item"><strong class="pro-label">আগ্রাসী স্টাইলিং:</strong> <span class="pro-description">পেশীবহুল নেকেড বাইক ডিজাইন স্পোর্টি দেখায়</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">শক্ত জাপানি বিল্ড মান</span></li>
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন:</strong> <span class="pro-description">এফআই মসৃণ পাওয়ার ডেলিভারি নিশ্চিত করে</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">গিয়ার ইন্ডিকেটর সহ সম্পূর্ণ ডিজিটাল এলসিডি মিটার</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এর্গোনমিক্স:</strong> <span class="pro-description">স্পোর্টি তবুও আরামদায়ক রাইডিং অবস্থান</span></li>
<li class="pro-item"><strong class="pro-label">এমআরএফ টায়ার:</strong> <span class="pro-description">প্রিমিয়াম এমআরএফ টায়ার দুর্দান্ত গ্রিপ প্রদান করে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">গড় মাইলেজ:</strong> <span class="con-description">জ্বালানি সাশ্রয়ের জন্য ৩৮-৪২ কিমি/লিটার ভালো নয়</span></li>
<li class="con-item"><strong class="con-label">সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="con-description">শুধুমাত্র সামনের চাকায় এবিএস</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">সুজুকি সার্ভিস সেন্টার সীমিত</span></li>
<li class="con-item"><strong class="con-label">শক্ত সাসপেনশন:</strong> <span class="con-description">খারাপ রাস্তায় শক্ত সেটআপ অস্বস্তিকর</span></li>
<li class="con-item"><strong class="con-label">কম রিসেল ভ্যালু:</strong> <span class="con-description">সুজুকি বাইক দ্রুত অবমূল্যায়ন করে</span></li>
<li class="con-item"><strong class="con-label">উচ্চ গতিতে কম্পন:</strong> <span class="con-description">৮০ কিমি/ঘন্টার উপরে লক্ষণীয় কম্পন</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি গিক্সার কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance-focused riders</li>
<li class="best-for-item">Those who love spirited riding</li>
<li class="best-for-item">Corner carvers and twisty road enthusiasts</li>
<li class="best-for-item">Young riders wanting sporty bike</li>
<li class="best-for-item">City riders with occasional highway use</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি গিক্সার কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Mileage-focused buyers</li>
<li class="not-recommended-item">Long-distance highway tourers</li>
<li class="not-recommended-item">Those needing maximum comfort</li>
<li class="not-recommended-item">Areas with limited Suzuki service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি গিক্সার এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৮০,০০০ - ৳২,৯০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৩,৫০০ - ৳৪,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪০ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,৫০০-২,৯০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৪ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,০০০ - ৳১,৩০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি গিক্সার পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড়ের নিচে মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো সুজুকি মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো পারফরম্যান্স মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- স্পোর্টি কিন্তু পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- শক্ত বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- এফআই + ডুয়াল ডিস্ক + ডিজিটাল</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- স্পোর্টি আগ্রাসী</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি গিক্সার সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">সুজুকি গিক্সার এর মাইলেজ কত?</h3>
<p class="faq-answer">সুজুকি গিক্সার উচ্ছৃঙ্খল রাইডিং সহ ৩৮-৪২ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে সুজুকি গিক্সার এর দাম কত?</h3>
<p class="faq-answer">সুজুকি গিক্সার এর দাম ৳২,৮০,০০০ থেকে ৳২,৯০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি গিক্সারে কি এবিএস আছে?</h3>
<p class="faq-answer">হ্যাঁ, সুজুকি গিক্সারে সামনের চাকায় সিঙ্গেল চ্যানেল এবিএস আছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি গিক্সার এর সর্বোচ্চ গতি কত?</h3>
<p class="faq-answer">সুজুকি গিক্সার প্রায় ১২৫-১৩০ কিমি/ঘন্টা পৌঁছাতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন ব্যবহারের জন্য সুজুকি গিক্সার কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, এটি স্পোর্টি চরিত্র এবং ভালো হ্যান্ডলিং সহ দৈনন্দিন ব্যবহারের জন্য চমৎকার।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: সুজুকি গিক্সার নাকি ইয়ামাহা এফজেডএস?</h3>
<p class="faq-answer">গিক্সারের ভালো হ্যান্ডলিং এবং পারফরম্যান্স আছে, যেখানে এফজেডএসে ভালো ফিচার এবং স্টাইলিং আছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি গিক্সারে কি ফুয়েল ইনজেকশন আছে?</h3>
<p class="faq-answer">হ্যাঁ, সুজুকি গিক্সার ফুয়েল ইনজেকশন সিস্টেম সহ আসে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সুজুকি গিক্সার এর ইঞ্জিন ক্ষমতা কত?</h3>
<p class="faq-answer">সুজুকি গিক্সারে একটি ১৫৫সিসি সিঙ্গেল-সিলিন্ডার এয়ার-কুলড ইঞ্জিন আছে যা ১৪.১ পিএস উৎপাদন করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি গিক্সার কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">সুজুকি গিক্সার ১৫০-১৬০সিসি সেগমেন্টে হ্যান্ডলিং কিং, ৳২,৮৫,০০০ টাকায় রোমাঞ্চকর পারফরম্যান্স এবং তীক্ষ্ণ ডাইনামিক্স প্রদান করে। রাইডারদের জন্য সেরা যারা জ্বালানি দক্ষতার চেয়ে স্পোর্টি রাইডিং অভিজ্ঞতাকে অগ্রাধিকার দেন। যদিও মাইলেজ গড় এবং সার্ভিস নেটওয়ার্ক সীমিত, ব্যতিক্রমী হ্যান্ডলিং এবং শক্তিশালী ইঞ্জিন এটিকে পারফরম্যান্স উৎসাহীদের জন্য একটি আকর্ষণীয় পছন্দ করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- টুইস্টি রাস্তা পছন্দ করেন পারফরম্যান্স উৎসাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer\n")

	return nil
}
