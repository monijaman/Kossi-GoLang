package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCb150RExmotionReview handles seeding Honda CB150R Exmotion product review and translation
type HondaCb150RExmotionReview struct {
	BaseSeeder
}

// NewHondaCb150RExmotionReviewSeeder creates a new HondaCb150RExmotionReview
func NewHondaCb150RExmotionReviewSeeder() *HondaCb150RExmotionReview {
	return &HondaCb150RExmotionReview{BaseSeeder: BaseSeeder{name: "Honda CB150R Exmotion Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCb150RExmotionReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB150R Exmotion").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda CB150R Exmotion product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda CB150R Exmotion product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CB150R Exmotion already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CB150R Exmotion Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda CB150R Exmotion is a premium 150cc naked sports motorcycle featuring aggressive neo-sports café styling, advanced features including LED lighting and digital instrumentation, upside-down front forks, and exceptional build quality. Priced at ৳7,55,000, it offers premium positioning with cutting-edge design, superior handling dynamics, Honda's legendary reliability, and sophisticated engineering for enthusiasts seeking a stylish, feature-rich, high-quality naked sports bike.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB150R Exmotion Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Neo-Sports Café Design:</strong> <span class="highlight-value">Aggressive premium styling with sharp lines</span></li>
<li class="highlight-item"><strong class="highlight-label">USD Front Forks:</strong> <span class="highlight-value">Upside-down forks for superior handling</span></li>
<li class="highlight-item"><strong class="highlight-label">Full LED Lighting:</strong> <span class="highlight-value">Premium LED headlight and taillight</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Instrumentation:</strong> <span class="highlight-value">Fully digital LCD meter with comprehensive info</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB150R Exmotion Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Stunning Design:</strong> <span class="pro-description">Neo-sports café styling is aggressive, modern, and premium</span></li>
<li class="pro-item"><strong class="pro-label">Superior Handling:</strong> <span class="pro-description">USD forks and chassis provide excellent cornering dynamics</span></li>
<li class="pro-item"><strong class="pro-label">Premium Build Quality:</strong> <span class="pro-description">Exceptional fit and finish with quality components</span></li>
<li class="pro-item"><strong class="pro-label">Advanced Features:</strong> <span class="pro-description">Full LED lighting, digital display, and modern technology</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Legendary Honda engineering and long-term dependability</span></li>
<li class="pro-item"><strong class="pro-label">Exclusive Positioning:</strong> <span class="pro-description">Premium price creates exclusivity and prestige</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB150R Exmotion Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very High Price:</strong> <span class="con-description">৳7,55,000 is significantly expensive for 150cc segment</span></li>
<li class="con-item"><strong class="con-label">Limited Availability:</strong> <span class="con-description">Not widely available, limited Honda dealership network</span></li>
<li class="con-item"><strong class="con-label">Moderate Performance:</strong> <span class="con-description">150cc engine doesn't match premium price expectations</span></li>
<li class="con-item"><strong class="con-label">High Service Costs:</strong> <span class="con-description">Premium components mean expensive maintenance</span></li>
<li class="con-item"><strong class="con-label">No ABS Variant:</strong> <span class="con-description">Missing ABS safety feature at this premium price point</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda CB150R Exmotion in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Premium motorcycle enthusiasts</li>
<li class="best-for-item">Style-conscious riders seeking exclusivity</li>
<li class="best-for-item">Those prioritizing design over displacement</li>
<li class="best-for-item">Urban riders wanting premium naked bike</li>
<li class="best-for-item">Honda brand loyalists</li>
<li class="best-for-item">Riders valuing build quality and handling</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda CB150R Exmotion?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Performance-focused riders needing more power</li>
<li class="not-recommended-item">Those seeking practical daily commuter</li>
<li class="not-recommended-item">Riders requiring ABS safety feature</li>
<li class="not-recommended-item">Value-for-money seekers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB150R Exmotion Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳7,55,000 - Premium pricing for 150cc segment</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate to High - ৳9-12 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳7-9 per km (40-45 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳12,000-18,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB150R Exmotion Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Good but not exceptional</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Moderate fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good urban ergonomics</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Outstanding premium design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- Premium pricing</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Exceptional build quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda CB150R Exmotion Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Why is it so expensive compared to other 150cc bikes?</h3>
<p class="faq-answer">The CB150R Exmotion commands premium pricing due to its neo-sports café design, USD forks, full LED lighting, exceptional build quality, premium components, and exclusive positioning in the market.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the performance worth the premium price?</h3>
<p class="faq-answer">The performance is adequate but not exceptional for the price. The value lies in design, handling, build quality, and exclusivity rather than raw power - it's a style and quality statement bike.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it compare to the CBR150R?</h3>
<p class="faq-answer">The CB150R Exmotion is more expensive and style-focused with naked design and USD forks, while CBR150R offers full fairing and more aggressive riding position at lower price, targeting different rider preferences.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB150R Exmotion in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳7,55,000, the Honda CB150R Exmotion is a premium 150cc naked sports motorcycle targeting enthusiasts who prioritize design, build quality, and exclusivity over displacement and value. It excels with stunning neo-sports café styling, superior handling from USD forks, exceptional build quality, advanced features, and Honda reliability. However, the very high price for 150cc, moderate performance, limited availability, high service costs, and lack of ABS make it suitable only for premium buyers and style-conscious riders who can afford the exclusivity and appreciate the design excellence over raw performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium neo-sports café design with exceptional build quality</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Honda CB150R Exmotion review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CB150R Exmotion (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা সিবি১৫০আর এক্সমোশন রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা সিবি১৫০আর এক্সমোশন হল একটি প্রিমিয়াম ১৫০সিসি নেকেড স্পোর্টস মোটরসাইকেল যেখানে আক্রমণাত্মক নিও-স্পোর্টস ক্যাফে স্টাইলিং, এলইডি লাইটিং এবং ডিজিটাল ইনস্ট্রুমেন্টেশন সহ উন্নত বৈশিষ্ট্য, আপসাইড-ডাউন ফ্রন্ট ফর্ক এবং ব্যতিক্রমী বিল্ড কোয়ালিটি রয়েছে। ৳৭,৫৫,০০০ টাকায় মূল্যায়িত, এটি অত্যাধুনিক ডিজাইন, উন্নত হ্যান্ডলিং ডায়নামিক্স, হোন্ডার কিংবদন্তি নির্ভরযোগ্যতা এবং স্টাইলিশ, ফিচার-সমৃদ্ধ, উচ্চ-মানের নেকেড স্পোর্টস বাইক খোঁজা উৎসাহীদের জন্য অত্যাধুনিক ইঞ্জিনিয়ারিং সহ প্রিমিয়াম পজিশনিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা সিবি১৫০আর এক্সমোশন এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">নিও-স্পোর্টস ক্যাফে ডিজাইন:</strong> <span class="highlight-value">তীক্ষ্ণ রেখা সহ আক্রমণাত্মক প্রিমিয়াম স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ইউএসডি ফ্রন্ট ফর্ক:</strong> <span class="highlight-value">উন্নত হ্যান্ডলিং জন্য আপসাইড-ডাউন ফর্ক</span></li>
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ এলইডি লাইটিং:</strong> <span class="highlight-value">প্রিমিয়াম এলইডি হেডলাইট এবং টেইললাইট</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল ইনস্ট্রুমেন্টেশন:</strong> <span class="highlight-value">ব্যাপক তথ্য সহ সম্পূর্ণ ডিজিটাল এলসিডি মিটার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা সিবি১৫০আর এক্সমোশন এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অত্যাশ্চর্য ডিজাইন:</strong> <span class="pro-description">নিও-স্পোর্টস ক্যাফে স্টাইলিং আক্রমণাত্মক, আধুনিক এবং প্রিমিয়াম</span></li>
<li class="pro-item"><strong class="pro-label">উন্নত হ্যান্ডলিং:</strong> <span class="pro-description">ইউএসডি ফর্ক এবং চ্যাসি চমৎকার কর্নারিং ডায়নামিক্স প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বিল্ড কোয়ালিটি:</strong> <span class="pro-description">গুণমানসম্পন্ন কম্পোনেন্ট সহ ব্যতিক্রমী ফিট এবং ফিনিশ</span></li>
<li class="pro-item"><strong class="pro-label">উন্নত বৈশিষ্ট্য:</strong> <span class="pro-description">সম্পূর্ণ এলইডি লাইটিং, ডিজিটাল ডিসপ্লে এবং আধুনিক প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">কিংবদন্তি হোন্ডা ইঞ্জিনিয়ারিং এবং দীর্ঘমেয়াদী নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">এক্সক্লুসিভ পজিশনিং:</strong> <span class="pro-description">প্রিমিয়াম মূল্য এক্সক্লুসিভিটি এবং প্রতিপত্তি তৈরি করে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা সিবি১৫০আর এক্সমোশন এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত উচ্চ মূল্য:</strong> <span class="con-description">৳৭,৫৫,০০০ ১৫০সিসি সেগমেন্টের জন্য উল্লেখযোগ্যভাবে ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সীমিত প্রাপ্যতা:</strong> <span class="con-description">ব্যাপকভাবে উপলব্ধ নয়, সীমিত হোন্ডা ডিলারশিপ নেটওয়ার্ক</span></li>
<li class="con-item"><strong class="con-label">মাঝারি পারফরমেন্স:</strong> <span class="con-description">১৫০সিসি ইঞ্জিন প্রিমিয়াম মূল্য প্রত্যাশার সাথে মিলে না</span></li>
<li class="con-item"><strong class="con-label">উচ্চ সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম কম্পোনেন্ট মানে ব্যয়বহুল রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস ভ্যারিয়েন্ট নেই:</strong> <span class="con-description">এই প্রিমিয়াম মূল্য পয়েন্টে এবিএস নিরাপত্তা বৈশিষ্ট্য অনুপস্থিত</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা সিবি১৫০আর এক্সমোশন কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Premium motorcycle enthusiasts</li>
<li class="best-for-item">Style-conscious riders seeking exclusivity</li>
<li class="best-for-item">Those prioritizing design over displacement</li>
<li class="best-for-item">Urban riders wanting premium naked bike</li>
<li class="best-for-item">Honda brand loyalists</li>
<li class="best-for-item">Riders valuing build quality and handling</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা সিবি১৫০আর এক্সমোশন কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Performance-focused riders needing more power</li>
<li class="not-recommended-item">Those seeking practical daily commuter</li>
<li class="not-recommended-item">Riders requiring ABS safety feature</li>
<li class="not-recommended-item">Value-for-money seekers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা সিবি১৫০আর এক্সমোশন এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৭,৫৫,০০০ - ১৫০সিসি সেগমেন্টের জন্য প্রিমিয়াম মূল্য</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি থেকে উচ্চ - ৳৯-১২ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৭-৯ প্রতি কিমি (৪০-৪৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১২,০০০-১৮,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা সিবি১৫০আর এক্সমোশন পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- ভাল কিন্তু ব্যতিক্রমী নয়</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- মাঝারি জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল শহুরে এরগোনমিক্স</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- অসাধারণ প্রিমিয়াম ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- প্রিমিয়াম মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- ব্যতিক্রমী বিল্ড কোয়ালিটি</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা সিবি১৫০আর এক্সমোশন সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">অন্যান্য ১৫০সিসি বাইকের তুলনায় এটি এত ব্যয়বহুল কেন?</h3>
<p class="faq-answer">সিবি১৫০আর এক্সমোশন এর নিও-স্পোর্টস ক্যাফে ডিজাইন, ইউএসডি ফর্ক, সম্পূর্ণ এলইডি লাইটিং, ব্যতিক্রমী বিল্ড কোয়ালিটি, প্রিমিয়াম কম্পোনেন্ট এবং বাজারে এক্সক্লুসিভ পজিশনিংয়ের কারণে প্রিমিয়াম মূল্য দাবি করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">পারফরমেন্স কি প্রিমিয়াম মূল্যের যোগ্য?</h3>
<p class="faq-answer">পারফরমেন্স মূল্যের জন্য পর্যাপ্ত কিন্তু ব্যতিক্রমী নয়। মূল্য কাঁচা শক্তির চেয়ে ডিজাইন, হ্যান্ডলিং, বিল্ড কোয়ালিটি এবং এক্সক্লুসিভিটিতে নিহিত - এটি একটি স্টাইল এবং গুণমান বিবৃতি বাইক।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি সিবিআর১৫০আর এর সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">সিবি১৫০আর এক্সমোশন নেকেড ডিজাইন এবং ইউএসডি ফর্ক সহ আরও ব্যয়বহুল এবং স্টাইল-ফোকাসড, যেখানে সিবিআর১৫০আর কম মূল্যে সম্পূর্ণ ফেয়ারিং এবং আরও আক্রমণাত্মক রাইডিং পজিশন প্রদান করে, বিভিন্ন রাইডার পছন্দ লক্ষ্য করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা সিবি১৫০আর এক্সমোশন কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৭,৫৫,০০০ টাকায় হোন্ডা সিবি১৫০আর এক্সমোশন হল একটি প্রিমিয়াম ১৫০সিসি নেকেড স্পোর্টস মোটরসাইকেল যা ডিসপ্লেসমেন্ট এবং মূল্যের চেয়ে ডিজাইন, বিল্ড কোয়ালিটি এবং এক্সক্লুসিভিটিকে অগ্রাধিকার দেওয়া উৎসাহীদের লক্ষ্য করে। এটি অত্যাশ্চর্য নিও-স্পোর্টস ক্যাফে স্টাইলিং, ইউএসডি ফর্ক থেকে উন্নত হ্যান্ডলিং, ব্যতিক্রমী বিল্ড কোয়ালিটি, উন্নত বৈশিষ্ট্য এবং হোন্ডা নির্ভরযোগ্যতা দিয়ে শ্রেষ্ঠ। তবে, ১৫০সিসির জন্য অত্যন্ত উচ্চ মূল্য, মাঝারি পারফরমেন্স, সীমিত প্রাপ্যতা, উচ্চ সেবা খরচ এবং এবিএসের অভাব এটিকে শুধুমাত্র প্রিমিয়াম ক্রেতা এবং স্টাইল-সচেতন রাইডারদের জন্য উপযুক্ত করে তোলে যারা এক্সক্লুসিভিটি বহন করতে পারে এবং কাঁচা পারফরমেন্সের চেয়ে ডিজাইন শ্রেষ্ঠত্বের প্রশংসা করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্যতিক্রমী বিল্ড কোয়ালিটি সহ প্রিমিয়াম নিও-স্পোর্টস ক্যাফে ডিজাইনের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CB150R Exmotion already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CB150R Exmotion\n")

	return nil
}
