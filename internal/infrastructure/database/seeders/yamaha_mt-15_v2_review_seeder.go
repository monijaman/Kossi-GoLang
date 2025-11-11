package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaMT15V2Review handles seeding Yamaha MT-15 V2 product review and translation
type YamahaMT15V2Review struct {
	BaseSeeder
}

// NewYamahaMT15V2ReviewSeeder creates a new YamahaMT15V2Review
func NewYamahaMT15V2ReviewSeeder() *YamahaMT15V2Review {
	return &YamahaMT15V2Review{BaseSeeder: BaseSeeder{name: "Yamaha MT-15 V2 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaMT15V2Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha MT-15 V2").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha MT-15 V2 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha MT-15 V2 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha MT-15 V2 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha MT-15 V2 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha MT-15 V2 is an updated 155cc naked streetfighter featuring VVA technology, aggressive styling, and enhanced features. Priced at ৳5,33,000, it combines the proven VVA engine with improved aerodynamics, updated electronics, premium color schemes, and Yamaha's Dark Side of Japan design philosophy for riders seeking aggressive street presence and performance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha MT-15 V2 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Dark Side of Japan Design:</strong> <span class="highlight-value">Aggressive streetfighter aesthetics</span></li>
<li class="highlight-item"><strong class="highlight-label">VVA Technology:</strong> <span class="highlight-value">Variable Valve Actuation for optimal power</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Color Options:</strong> <span class="highlight-value">Cyan Storm & Ice Fluo Vermillion schemes</span></li>
<li class="highlight-item"><strong class="highlight-label">Upright Riding Position:</strong> <span class="highlight-value">Comfortable naked bike ergonomics</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha MT-15 V2 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Proven VVA Performance:</strong> <span class="pro-description">Variable valve technology delivers excellent power across rev range</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Streetfighter Looks:</strong> <span class="pro-description">Bold and menacing design that commands attention</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight and Agile:</strong> <span class="pro-description">Easy to handle in city traffic and tight spaces</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Yamaha's proven engineering and component quality</span></li>
<li class="pro-item"><strong class="pro-label">Updated Features:</strong> <span class="pro-description">Improved aerodynamics and enhanced electronics</span></li>
<li class="pro-item"><strong class="pro-label">Street-Friendly Ergonomics:</strong> <span class="pro-description">Upright riding position suitable for daily use</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha MT-15 V2 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Purchase Price:</strong> <span class="con-description">Premium pricing for 155cc naked bike segment</span></li>
<li class="con-item"><strong class="con-label">Limited Wind Protection:</strong> <span class="con-description">Naked design offers no protection from wind</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">VVA technology requires specialized service</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha dealerships and trained technicians</span></li>
<li class="con-item"><strong class="con-label">Moderate Fuel Efficiency:</strong> <span class="con-description">Performance focus results in average fuel economy</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha MT-15 V2 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban streetfighter enthusiasts</li>
<li class="best-for-item">Riders seeking aggressive styling</li>
<li class="best-for-item">Daily commuters wanting performance</li>
<li class="best-for-item">Young riders preferring bold design</li>
<li class="best-for-item">Weekend city riding</li>
<li class="best-for-item">Those appreciating VVA technology</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha MT-15 V2?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Highway cruising enthusiasts</li>
<li class="not-recommended-item">Riders needing wind protection</li>
<li class="not-recommended-item">Those seeking maximum fuel efficiency</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha MT-15 V2 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳5,33,000 - Premium for 155cc streetfighter with VVA</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳11-14 per km with VVA maintenance</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳9-12 per km (36-42 km/l with VVA efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳9,000-15,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha MT-15 V2 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent VVA performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Good for performance bike</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good upright position</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Outstanding streetfighter design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Fair for VVA features</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent Yamaha quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha MT-15 V2 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What improvements does the V2 offer over the original MT-15?</h3>
<p class="faq-answer">The V2 features improved aerodynamics, updated electronics, premium color schemes (Cyan Storm & Ice Fluo Vermillion), and refinements in build quality while maintaining the proven VVA engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does the streetfighter design affect daily riding?</h3>
<p class="faq-answer">The upright riding position and lightweight nature make it comfortable for city riding, though the lack of wind protection may be noticeable at highway speeds.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the premium pricing justified?</h3>
<p class="faq-answer">The pricing reflects VVA technology, premium build quality, updated features, and Yamaha's brand value. It's competitive within the premium 155cc naked segment.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha MT-15 V2 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳5,33,000, the Yamaha MT-15 V2 delivers excellent value in the premium naked streetfighter segment with proven VVA technology, aggressive styling, and updated features. The Dark Side of Japan design philosophy and premium color schemes make it stand out, while the VVA engine provides strong performance. However, the high pricing, expensive maintenance, limited wind protection, and moderate fuel efficiency make it suitable primarily for urban riders and streetfighter enthusiasts who prioritize aggressive styling and VVA performance over economy and practicality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For aggressive streetfighter styling with proven VVA performance</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.4,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha MT-15 V2 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha MT-15 V2 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা এমটি-১৫ ভি২ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা এমটি-১৫ ভি২ হল একটি আপডেটেড ১৫৫সিসি নেকেড স্ট্রিটফাইটার যেখানে ভিভিএ প্রযুক্তি, আক্রমণাত্মক স্টাইলিং এবং উন্নত বৈশিষ্ট্য রয়েছে। ৳৫,৩৩,০০০ টাকায় মূল্যায়িত, এটি উন্নত অ্যারোডাইনামিক্স, আপডেটেড ইলেকট্রনিক্স, প্রিমিয়াম কালার স্কিম এবং আক্রমণাত্মক স্ট্রিট উপস্থিতি ও পারফরমেন্স খোঁজা রাইডারদের জন্য ইয়ামাহার ডার্ক সাইড অফ জাপান ডিজাইন দর্শনের সাথে প্রমাণিত ভিভিএ ইঞ্জিনকে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা এমটি-১৫ ভি২ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ডার্ক সাইড অফ জাপান ডিজাইন:</strong> <span class="highlight-value">আক্রমণাত্মক স্ট্রিটফাইটার নান্দনিকতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ভিভিএ প্রযুক্তি:</strong> <span class="highlight-value">সর্বোত্তম শক্তির জন্য ভেরিয়েবল ভালভ অ্যাকচুয়েশন</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম কালার অপশন:</strong> <span class="highlight-value">সায়ান স্ট্রম এবং আইস ফ্লুও ভার্মিলিয়ন স্কিম</span></li>
<li class="highlight-item"><strong class="highlight-label">সোজা রাইডিং পজিশন:</strong> <span class="highlight-value">আরামদায়ক নেকেড বাইক এরগোনমিক্স</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা এমটি-১৫ ভি২ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রমাণিত ভিভিএ পারফরমেন্স:</strong> <span class="pro-description">ভেরিয়েবল ভালভ প্রযুক্তি রেভ রেঞ্জ জুড়ে চমৎকার পাওয়ার প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্ট্রিটফাইটার চেহারা:</strong> <span class="pro-description">সাহসী এবং ভীতিকর ডিজাইন যা মনোযোগ আকর্ষণ করে</span></li>
<li class="pro-item"><strong class="pro-label">হালকা এবং চটপটে:</strong> <span class="pro-description">শহুরে ট্রাফিক এবং সংকীর্ণ জায়গায় পরিচালনা সহজ</span></li>
<li class="pro-item"><strong class="pro-label">ভাল বিল্ড কোয়ালিটি:</strong> <span class="pro-description">ইয়ামাহার প্রমাণিত ইঞ্জিনিয়ারিং এবং কম্পোনেন্ট গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">আপডেটেড ফিচার:</strong> <span class="pro-description">উন্নত অ্যারোডাইনামিক্স এবং উন্নত ইলেকট্রনিক্স</span></li>
<li class="pro-item"><strong class="pro-label">স্ট্রিট-ফ্রেন্ডলি এরগোনমিক্স:</strong> <span class="pro-description">দৈনিক ব্যবহারের জন্য উপযুক্ত সোজা রাইডিং পজিশন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা এমটি-১৫ ভি২ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ ক্রয় মূল্য:</strong> <span class="con-description">১৫৫সিসি নেকেড বাইক সেগমেন্টের জন্য প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">সীমিত বায়ু সুরক্ষা:</strong> <span class="con-description">নেকেড ডিজাইন বায়ু থেকে কোন সুরক্ষা প্রদান করে না</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">ভিভিএ প্রযুক্তির বিশেষায়িত সেবা প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">কম ইয়ামাহা ডিলারশিপ এবং প্রশিক্ষিত প্রযুক্তিবিদ</span></li>
<li class="con-item"><strong class="con-label">মাঝারি জ্বালানি দক্ষতা:</strong> <span class="con-description">পারফরমেন্স ফোকাস গড় জ্বালানি অর্থনীতির ফলাফল</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা এমটি-১৫ ভি২ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban streetfighter enthusiasts</li>
<li class="best-for-item">Riders seeking aggressive styling</li>
<li class="best-for-item">Daily commuters wanting performance</li>
<li class="best-for-item">Young riders preferring bold design</li>
<li class="best-for-item">Weekend city riding</li>
<li class="best-for-item">Those appreciating VVA technology</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা এমটি-১৫ ভি২ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Highway cruising enthusiasts</li>
<li class="not-recommended-item">Riders needing wind protection</li>
<li class="not-recommended-item">Those seeking maximum fuel efficiency</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা এমটি-১৫ ভি২ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৫,৩৩,০০০ - ভিভিএ সহ ১৫৫সিসি স্ট্রিটফাইটারের জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">উচ্চ - ভিভিএ রক্ষণাবেক্ষণ সহ ৳১১-১৪ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৯-১২ প্রতি কিমি (ভিভিএ দক্ষতা সহ ৩৬-৪২ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৯,০০০-১৫,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা এমটি-১৫ ভি২ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- চমৎকার ভিভিএ পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- পারফরমেন্স বাইকের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল সোজা অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- অসাধারণ স্ট্রিটফাইটার ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভিভিএ বৈশিষ্ট্যের জন্য ন্যায্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- চমৎকার ইয়ামাহা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা এমটি-১৫ ভি২ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">মূল এমটি-১৫ এর তুলনায় ভি২ কী উন্নতি প্রদান করে?</h3>
<p class="faq-answer">ভি২ তে উন্নত অ্যারোডাইনামিক্স, আপডেটেড ইলেকট্রনিক্স, প্রিমিয়াম কালার স্কিম (সায়ান স্ট্রম এবং আইস ফ্লুও ভার্মিলিয়ন) এবং প্রমাণিত ভিভিএ ইঞ্জিন বজায় রেখে বিল্ড কোয়ালিটির উন্নতি রয়েছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">স্ট্রিটফাইটার ডিজাইন দৈনিক রাইডিংয়ে কীভাবে প্রভাব ফেলে?</h3>
<p class="faq-answer">সোজা রাইডিং পজিশন এবং হালকা প্রকৃতি এটিকে শহুরে রাইডিংয়ের জন্য আরামদায়ক করে তোলে, যদিও হাইওয়ে গতিতে বায়ু সুরক্ষার অভাব লক্ষণীয় হতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">প্রিমিয়াম মূল্য কি ন্যায্য?</h3>
<p class="faq-answer">মূল্য ভিভিএ প্রযুক্তি, প্রিমিয়াম বিল্ড কোয়ালিটি, আপডেটেড ফিচার এবং ইয়ামাহার ব্র্যান্ড ভ্যালু প্রতিফলিত করে। প্রিমিয়াম ১৫৫সিসি নেকেড সেগমেন্টে এটি প্রতিযোগিতামূলক।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা এমটি-১৫ ভি২ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৫,৩৩,০০০ টাকায় ইয়ামাহা এমটি-১৫ ভি২ প্রমাণিত ভিভিএ প্রযুক্তি, আক্রমণাত্মক স্টাইলিং এবং আপডেটেড ফিচার সহ প্রিমিয়াম নেকেড স্ট্রিটফাইটার সেগমেন্টে চমৎকার ভ্যালু প্রদান করে। ডার্ক সাইড অফ জাপান ডিজাইন দর্শন এবং প্রিমিয়াম কালার স্কিম এটিকে আলাদা করে তোলে, যখন ভিভিএ ইঞ্জিন শক্তিশালী পারফরমেন্স প্রদান করে। তবে, উচ্চ মূল্য, ব্যয়বহুল রক্ষণাবেক্ষণ, সীমিত বায়ু সুরক্ষা এবং মাঝারি জ্বালানি দক্ষতা এটিকে প্রধানত শহুরে রাইডার এবং স্ট্রিটফাইটার উৎসাহীদের জন্য উপযুক্ত করে তোলে যারা অর্থনীতি এবং ব্যবহারিকতার চেয়ে আক্রমণাত্মক স্টাইলিং এবং ভিভিএ পারফরমেন্সকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রমাণিত ভিভিএ পারফরমেন্স সহ আক্রমণাত্মক স্ট্রিটফাইটার স্টাইলিংয়ের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha MT-15 V2 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha MT-15 V2\n")

	return nil
}
