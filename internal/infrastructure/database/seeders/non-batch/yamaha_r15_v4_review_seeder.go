package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaR15V4Review handles seeding Yamaha R15 V4 product review and translation
type YamahaR15V4Review struct {
	BaseSeeder
}

// NewYamahaR15V4ReviewSeeder creates a new YamahaR15V4Review
func NewYamahaR15V4ReviewSeeder() *YamahaR15V4Review {
	return &YamahaR15V4Review{BaseSeeder: BaseSeeder{name: "Yamaha R15 V4 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaR15V4Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R15 V4").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha R15 V4 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha R15 V4 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha R15 V4 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha R15 V4 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha R15 V4 is a premium 155cc fully-faired sports motorcycle featuring VVA technology, slipper clutch, and dual-channel ABS. Priced at ৳6,10,000, it offers an excellent balance of performance and usability with advanced engine technology, sporty design, premium build quality, and Yamaha's renowned R-series heritage for enthusiasts seeking authentic sports bike experience.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R15 V4 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">VVA Technology:</strong> <span class="highlight-value">Variable Valve Actuation for optimal power delivery</span></li>
<li class="highlight-item"><strong class="highlight-label">Slipper Clutch:</strong> <span class="highlight-value">Anti-hop clutch for smoother downshifts</span></li>
<li class="highlight-item"><strong class="highlight-label">Dual Channel ABS:</strong> <span class="highlight-value">Enhanced braking safety on both wheels</span></li>
<li class="highlight-item"><strong class="highlight-label">Sporty Ergonomics:</strong> <span class="highlight-value">Balanced riding position for performance and comfort</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R15 V4 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">VVA Engine Performance:</strong> <span class="pro-description">Variable valve timing delivers excellent power across rev range</span></li>
<li class="pro-item"><strong class="pro-label">Comprehensive Safety Features:</strong> <span class="pro-description">Dual-channel ABS and slipper clutch enhance riding safety</span></li>
<li class="pro-item"><strong class="pro-label">Premium Build Quality:</strong> <span class="pro-description">High-quality components and excellent fit-finish</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Yet Practical:</strong> <span class="pro-description">Good balance between performance and daily usability</span></li>
<li class="pro-item"><strong class="pro-label">Aerodynamic Design:</strong> <span class="pro-description">Full fairing with excellent wind protection</span></li>
<li class="pro-item"><strong class="pro-label">Brand Heritage:</strong> <span class="pro-description">Yamaha R-series reputation and engineering excellence</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R15 V4 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Purchase Price:</strong> <span class="con-description">Premium pricing compared to other 155cc motorcycles</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">Premium features require specialized service and costly parts</span></li>
<li class="con-item"><strong class="con-label">Sport-Focused Ergonomics:</strong> <span class="con-description">Riding position may be aggressive for long commutes</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha dealerships and trained technicians</span></li>
<li class="con-item"><strong class="con-label">Fuel Economy Trade-off:</strong> <span class="con-description">Performance focus results in moderate fuel efficiency</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha R15 V4 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Sports bike enthusiasts</li>
<li class="best-for-item">Weekend performance riders</li>
<li class="best-for-item">Experienced riders seeking premium features</li>
<li class="best-for-item">Those wanting authentic R-series experience</li>
<li class="best-for-item">Riders prioritizing safety with dual ABS</li>
<li class="best-for-item">Performance-oriented daily commuters</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha R15 V4?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Beginner or novice riders</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Those seeking maximum fuel efficiency</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R15 V4 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳6,10,000 - Premium for 155cc with VVA and dual ABS</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳12-16 per km with premium maintenance</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳10-13 per km (34-40 km/l with VVA efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳10,000-17,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R15 V4 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent VVA performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Good for performance bike</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- Sporty but manageable</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Premium sports design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Fair for features offered</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent Yamaha quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha R15 V4 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does the V4 compare to previous R15 versions?</h3>
<p class="faq-answer">The V4 features upgraded VVA technology, improved aerodynamics, better ergonomics, dual-channel ABS, and slipper clutch, representing significant advancement over previous generations.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the slipper clutch beneficial for street riding?</h3>
<p class="faq-answer">Yes, the slipper clutch prevents rear wheel hop during aggressive downshifting, providing smoother engine braking and better stability, beneficial for both street and track riding.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What maintenance should be expected?</h3>
<p class="faq-answer">Regular services every 3,500 km, quality engine oil, air filter cleaning, chain maintenance, and occasional valve adjustments. Premium parts and specialized service are recommended.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R15 V4 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳6,10,000, the Yamaha R15 V4 delivers exceptional value in the premium 155cc sports segment with advanced VVA technology, dual-channel ABS, slipper clutch, and superior build quality. It offers an excellent balance of performance and practicality with Yamaha's proven R-series DNA. However, the high purchase price, expensive maintenance, sport-focused ergonomics, limited service network, and moderate fuel economy make it suitable primarily for performance enthusiasts and experienced riders who appreciate premium features and are willing to pay for authentic sports bike experience.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium VVA performance with comprehensive safety features</span></p>
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
		return fmt.Errorf("error creating Yamaha R15 V4 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha R15 V4 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা আর১৫ ভি৪ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা আর১৫ ভি৪ হল একটি প্রিমিয়াম ১৫৫সিসি ফুল-ফেয়ার্ড স্পোর্টস মোটরসাইকেল যেখানে ভিভিএ প্রযুক্তি, স্লিপার ক্লাচ এবং ডুয়াল-চ্যানেল এবিএস রয়েছে। ৳৬,১০,০০০ টাকায় মূল্যায়িত, এটি উন্নত ইঞ্জিন প্রযুক্তি, স্পোর্টি ডিজাইন, প্রিমিয়াম বিল্ড কোয়ালিটি এবং খাঁটি স্পোর্টস বাইক অভিজ্ঞতা খোঁজা উৎসাহীদের জন্য ইয়ামাহার বিখ্যাত আর-সিরিজ ঐতিহ্য সহ পারফরমেন্স এবং ব্যবহারযোগ্যতার একটি চমৎকার ভারসাম্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা আর১৫ ভি৪ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ভিভিএ প্রযুক্তি:</strong> <span class="highlight-value">সর্বোত্তম পাওয়ার ডেলিভারির জন্য ভেরিয়েবল ভালভ অ্যাকচুয়েশন</span></li>
<li class="highlight-item"><strong class="highlight-label">স্লিপার ক্লাচ:</strong> <span class="highlight-value">মসৃণ ডাউনশিফটের জন্য অ্যান্টি-হপ ক্লাচ</span></li>
<li class="highlight-item"><strong class="highlight-label">ডুয়াল চ্যানেল এবিএস:</strong> <span class="highlight-value">উভয় চাকায় উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্টি এরগোনমিক্স:</strong> <span class="highlight-value">পারফরমেন্স এবং আরামের জন্য ভারসাম্যপূর্ণ রাইডিং পজিশন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা আর১৫ ভি৪ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ভিভিএ ইঞ্জিন পারফরমেন্স:</strong> <span class="pro-description">ভেরিয়েবল ভালভ টাইমিং রেভ রেঞ্জ জুড়ে চমৎকার পাওয়ার প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ব্যাপক নিরাপত্তা বৈশিষ্ট্য:</strong> <span class="pro-description">ডুয়াল-চ্যানেল এবিএস এবং স্লিপার ক্লাচ রাইডিং নিরাপত্তা বাড়ায়</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বিল্ড কোয়ালিটি:</strong> <span class="pro-description">উচ্চ-মানের কম্পোনেন্ট এবং চমৎকার ফিট-ফিনিশ</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি অথচ ব্যবহারিক:</strong> <span class="pro-description">পারফরমেন্স এবং দৈনিক ব্যবহারযোগ্যতার মধ্যে ভাল ভারসাম্য</span></li>
<li class="pro-item"><strong class="pro-label">অ্যারোডাইনামিক ডিজাইন:</strong> <span class="pro-description">চমৎকার বায়ু সুরক্ষা সহ সম্পূর্ণ ফেয়ারিং</span></li>
<li class="pro-item"><strong class="pro-label">ব্র্যান্ড ঐতিহ্য:</strong> <span class="pro-description">ইয়ামাহা আর-সিরিজ খ্যাতি এবং ইঞ্জিনিয়ারিং শ্রেষ্ঠত্ব</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা আর১৫ ভি৪ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ ক্রয় মূল্য:</strong> <span class="con-description">অন্যান্য ১৫৫সিসি মোটরসাইকেলের তুলনায় প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">প্রিমিয়াম ফিচারের বিশেষায়িত সেবা এবং ব্যয়বহুল যন্ত্রাংশ প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">স্পোর্ট-ফোকাসড এরগোনমিক্স:</strong> <span class="con-description">রাইডিং পজিশন দীর্ঘ যাতায়াতের জন্য আক্রমণাত্মক হতে পারে</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">কম ইয়ামাহা ডিলারশিপ এবং প্রশিক্ষিত প্রযুক্তিবিদ</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি সাশ্রয়ের ট্রেড-অফ:</strong> <span class="con-description">পারফরমেন্স ফোকাস মাঝারি জ্বালানি দক্ষতার ফলাফল</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা আর১৫ ভি৪ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Sports bike enthusiasts</li>
<li class="best-for-item">Weekend performance riders</li>
<li class="best-for-item">Experienced riders seeking premium features</li>
<li class="best-for-item">Those wanting authentic R-series experience</li>
<li class="best-for-item">Riders prioritizing safety with dual ABS</li>
<li class="best-for-item">Performance-oriented daily commuters</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা আর১৫ ভি৪ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Beginner or novice riders</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Those seeking maximum fuel efficiency</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা আর১৫ ভি৪ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৬,১০,০০০ - ভিভিএ এবং ডুয়াল এবিএস সহ ১৫৫সিসির জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">উচ্চ - প্রিমিয়াম রক্ষণাবেক্ষণ সহ ৳১২-১৬ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১০-১৩ প্রতি কিমি (ভিভিএ দক্ষতা সহ ৩৪-৪০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১০,০০০-১৭,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা আর১৫ ভি৪ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার ভিভিএ পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- পারফরমেন্স বাইকের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- স্পোর্টি কিন্তু পরিচালনাযোগ্য</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- প্রিমিয়াম স্পোর্টস ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- প্রদত্ত বৈশিষ্ট্যের জন্য ন্যায্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার ইয়ামাহা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা আর১৫ ভি৪ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ভি৪ পূর্ববর্তী আর১৫ সংস্করণের সাথে কেমন তুলনা করে?</h3>
<p class="faq-answer">ভি৪ এ আপগ্রেডেড ভিভিএ প্রযুক্তি, উন্নত অ্যারোডাইনামিক্স, ভাল এরগোনমিক্স, ডুয়াল-চ্যানেল এবিএস এবং স্লিপার ক্লাচ রয়েছে, যা পূর্ববর্তী প্রজন্মের তুলনায় উল্লেখযোগ্য অগ্রগতির প্রতিনিধিত্ব করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">স্ট্রিট রাইডিংয়ের জন্য স্লিপার ক্লাচ কি উপকারী?</h3>
<p class="faq-answer">হ্যাঁ, স্লিপার ক্লাচ আক্রমণাত্মক ডাউনশিফটিংয়ের সময় পিছনের চাকার হপ প্রতিরোধ করে, মসৃণ ইঞ্জিন ব্রেকিং এবং ভাল স্থিতিশীলতা প্রদান করে, যা স্ট্রিট এবং ট্র্যাক উভয় রাইডিংয়ের জন্য উপকারী।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কী রক্ষণাবেক্ষণ আশা করা উচিত?</h3>
<p class="faq-answer">প্রতি ৩,৫০০ কিমিতে নিয়মিত সেবা, মানসম্পন্ন ইঞ্জিন তেল, এয়ার ফিল্টার পরিষ্কার, চেইন রক্ষণাবেক্ষণ এবং মাঝে মাঝে ভালভ সামঞ্জস্য। প্রিমিয়াম যন্ত্রাংশ এবং বিশেষায়িত সেবা সুপারিশ করা হয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা আর১৫ ভি৪ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৬,১০,০০০ টাকায় ইয়ামাহা আর১৫ ভি৪ উন্নত ভিভিএ প্রযুক্তি, ডুয়াল-চ্যানেল এবিএস, স্লিপার ক্লাচ এবং উন্নত বিল্ড কোয়ালিটি সহ প্রিমিয়াম ১৫৫সিসি স্পোর্টস সেগমেন্টে ব্যতিক্রমী ভ্যালু প্রদান করে। এটি ইয়ামাহার প্রমাণিত আর-সিরিজ ডিএনএ সহ পারফরমেন্স এবং ব্যবহারিকতার একটি চমৎকার ভারসাম্য প্রদান করে। তবে, উচ্চ ক্রয় মূল্য, ব্যয়বহুল রক্ষণাবেক্ষণ, স্পোর্ট-ফোকাসড এরগোনমিক্স, সীমিত সেবা নেটওয়ার্ক এবং মাঝারি জ্বালানি অর্থনীতি এটিকে প্রধানত পারফরমেন্স উৎসাহী এবং অভিজ্ঞ রাইডারদের জন্য উপযুক্ত করে তোলে যারা প্রিমিয়াম বৈশিষ্ট্যের প্রশংসা করেন এবং খাঁটি স্পোর্টস বাইক অভিজ্ঞতার জন্য অর্থ প্রদান করতে ইচ্ছুক।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্যাপক নিরাপত্তা বৈশিষ্ট্য সহ প্রিমিয়াম ভিভিএ পারফরমেন্সের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha R15 V4 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha R15 V4\n")

	return nil
}
