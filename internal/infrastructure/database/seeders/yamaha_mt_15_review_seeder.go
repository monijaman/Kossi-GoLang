package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaMT15Review handles seeding Yamaha MT 15 product review and translation
type YamahaMT15Review struct {
	BaseSeeder
}

// NewYamahaMT15ReviewSeeder creates a new YamahaMT15Review
func NewYamahaMT15ReviewSeeder() *YamahaMT15Review {
	return &YamahaMT15Review{BaseSeeder: BaseSeeder{name: "Yamaha MT 15 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaMT15Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha MT 15").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha MT 15 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha MT 15 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha MT 15 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha MT 15 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha MT 15 is a 155cc naked streetfighter featuring VVA technology, aggressive Dark Side of Japan styling, and premium build quality. Priced at ৳4,60,000, it offers an accessible entry into Yamaha's MT series with variable valve actuation, upright ergonomics, excellent handling, and proven reliability for riders seeking authentic streetfighter experience in an affordable package.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha MT 15 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Dark Side of Japan Design:</strong> <span class="highlight-value">Menacing streetfighter aesthetics</span></li>
<li class="highlight-item"><strong class="highlight-label">VVA Technology:</strong> <span class="highlight-value">Variable Valve Actuation for optimal power</span></li>
<li class="highlight-item"><strong class="highlight-label">Upright Ergonomics:</strong> <span class="highlight-value">Comfortable riding position for daily use</span></li>
<li class="highlight-item"><strong class="highlight-label">Lightweight Handling:</strong> <span class="highlight-value">Agile and responsive in city traffic</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha MT 15 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Proven VVA Performance:</strong> <span class="pro-description">Variable valve technology delivers strong performance throughout rev range</span></li>
<li class="pro-item"><strong class="pro-label">Iconic Streetfighter Styling:</strong> <span class="pro-description">Aggressive MT-series design language that commands attention</span></li>
<li class="pro-item"><strong class="pro-label">Accessible Pricing:</strong> <span class="pro-description">More affordable entry into Yamaha's premium VVA technology</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Handling:</strong> <span class="pro-description">Lightweight and agile for urban riding and traffic navigation</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Yamaha's proven engineering and component reliability</span></li>
<li class="pro-item"><strong class="pro-label">Daily Usability:</strong> <span class="pro-description">Comfortable ergonomics suitable for regular commuting</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha MT 15 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Wind Protection:</strong> <span class="con-description">Naked design offers no protection from wind at highway speeds</span></li>
<li class="con-item"><strong class="con-label">Moderate Fuel Efficiency:</strong> <span class="con-description">Performance focus results in average fuel economy</span></li>
<li class="con-item"><strong class="con-label">Premium Maintenance:</strong> <span class="con-description">VVA technology requires specialized service knowledge</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha dealerships compared to local brands</span></li>
<li class="con-item"><strong class="con-label">No ABS Available:</strong> <span class="con-description">Basic braking system without anti-lock technology</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha MT 15 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Entry-level performance enthusiasts</li>
<li class="best-for-item">Urban commuters seeking style</li>
<li class="best-for-item">Riders wanting VVA technology affordably</li>
<li class="best-for-item">Daily city riding</li>
<li class="best-for-item">Young riders preferring aggressive styling</li>
<li class="best-for-item">Those upgrading from smaller displacement bikes</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha MT 15?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Long-distance highway touring</li>
<li class="not-recommended-item">Riders requiring ABS safety</li>
<li class="not-recommended-item">Those needing wind protection</li>
<li class="not-recommended-item">Budget buyers seeking maximum value</li>
<li class="not-recommended-item">Riders preferring conservative styling</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha MT 15 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,60,000 - Good value for VVA technology in MT series</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳9-12 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳8-10 per km (37-42 km/l with VVA efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳7,500-12,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha MT 15 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Excellent VVA performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Good for performance</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Good upright position</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Iconic MT design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good for VVA features</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Reliable Yamaha build</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha MT 15 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does the MT 15 compare to other 155cc naked bikes?</h3>
<p class="faq-answer">The MT 15 stands out with VVA technology, aggressive MT-series styling, and Yamaha build quality. It offers better performance than conventional 155cc bikes but at a premium price.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is VVA technology worth the extra cost?</h3>
<p class="faq-answer">VVA provides better power delivery, improved fuel efficiency, and enhanced performance across the rev range. For performance-minded riders, it's a worthwhile investment.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What maintenance should be expected?</h3>
<p class="faq-answer">Regular services every 3,500 km, quality engine oil for VVA system, air filter maintenance, and chain lubrication. VVA requires occasional specialized attention but is generally reliable.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha MT 15 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,60,000, the Yamaha MT 15 offers excellent value as an accessible entry into Yamaha's premium VVA technology and iconic MT-series design. It provides strong performance, aggressive styling, and reliable build quality at a more affordable price point than higher-end MT models. However, the lack of ABS, limited wind protection, moderate fuel efficiency, and premium maintenance requirements make it suitable primarily for urban riders and performance enthusiasts who prioritize VVA technology and streetfighter aesthetics over comfort features and maximum value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For accessible VVA performance with iconic MT styling</span></p>
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
		return fmt.Errorf("error creating Yamaha MT 15 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha MT 15 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা এমটি ১৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা এমটি ১৫ হল একটি ১৫৫সিসি নেকেড স্ট্রিটফাইটার যেখানে ভিভিএ প্রযুক্তি, আক্রমণাত্মক ডার্ক সাইড অফ জাপান স্টাইলিং এবং প্রিমিয়াম বিল্ড কোয়ালিটি রয়েছে। ৳৪,৬০,০০০ টাকায় মূল্যায়িত, এটি ভেরিয়েবল ভালভ অ্যাকচুয়েশন, সোজা এরগোনমিক্স, চমৎকার হ্যান্ডলিং এবং একটি সাশ্রয়ী প্যাকেজে খাঁটি স্ট্রিটফাইটার অভিজ্ঞতা খোঁজা রাইডারদের জন্য প্রমাণিত নির্ভরযোগ্যতা সহ ইয়ামাহার এমটি সিরিজে একটি অ্যাক্সেসযোগ্য প্রবেশ প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা এমটি ১৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ডার্ক সাইড অফ জাপান ডিজাইন:</strong> <span class="highlight-value">ভীতিকর স্ট্রিটফাইটার নান্দনিকতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ভিভিএ প্রযুক্তি:</strong> <span class="highlight-value">সর্বোত্তম শক্তির জন্য ভেরিয়েবল ভালভ অ্যাকচুয়েশন</span></li>
<li class="highlight-item"><strong class="highlight-label">সোজা এরগোনমিক্স:</strong> <span class="highlight-value">দৈনিক ব্যবহারের জন্য আরামদায়ক রাইডিং পজিশন</span></li>
<li class="highlight-item"><strong class="highlight-label">হালকা হ্যান্ডলিং:</strong> <span class="highlight-value">শহুরে ট্রাফিকে চটপটে এবং প্রতিক্রিয়াশীল</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা এমটি ১৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রমাণিত ভিভিএ পারফরমেন্স:</strong> <span class="pro-description">ভেরিয়েবল ভালভ প্রযুক্তি রেভ রেঞ্জ জুড়ে শক্তিশালী পারফরমেন্স প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">আইকনিক স্ট্রিটফাইটার স্টাইলিং:</strong> <span class="pro-description">আক্রমণাত্মক এমটি-সিরিজ ডিজাইন ভাষা যা মনোযোগ আকর্ষণ করে</span></li>
<li class="pro-item"><strong class="pro-label">অ্যাক্সেসযোগ্য মূল্য:</strong> <span class="pro-description">ইয়ামাহার প্রিমিয়াম ভিভিএ প্রযুক্তিতে আরও সাশ্রয়ী প্রবেশ</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার হ্যান্ডলিং:</strong> <span class="pro-description">শহুরে রাইডিং এবং ট্রাফিক নেভিগেশনের জন্য হালকা এবং চটপটে</span></li>
<li class="pro-item"><strong class="pro-label">ভাল বিল্ড কোয়ালিটি:</strong> <span class="pro-description">ইয়ামাহার প্রমাণিত ইঞ্জিনিয়ারিং এবং কম্পোনেন্ট নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">দৈনিক ব্যবহারযোগ্যতা:</strong> <span class="pro-description">নিয়মিত যাতায়াতের জন্য উপযুক্ত আরামদায়ক এরগোনমিক্স</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা এমটি ১৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত বায়ু সুরক্ষা:</strong> <span class="con-description">নেকেড ডিজাইন হাইওয়ে গতিতে বায়ু থেকে কোন সুরক্ষা প্রদান করে না</span></li>
<li class="con-item"><strong class="con-label">মাঝারি জ্বালানি দক্ষতা:</strong> <span class="con-description">পারফরমেন্স ফোকাস গড় জ্বালানি অর্থনীতির ফলাফল</span></li>
<li class="con-item"><strong class="con-label">প্রিমিয়াম রক্ষণাবেক্ষণ:</strong> <span class="con-description">ভিভিএ প্রযুক্তির বিশেষায়িত সেবা জ্ঞান প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">স্থানীয় ব্র্যান্ডের তুলনায় কম ইয়ামাহা ডিলারশিপ</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস উপলব্ধ নেই:</strong> <span class="con-description">অ্যান্টি-লক প্রযুক্তি ছাড়া বেসিক ব্রেকিং সিস্টেম</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা এমটি ১৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Entry-level performance enthusiasts</li>
<li class="best-for-item">Urban commuters seeking style</li>
<li class="best-for-item">Riders wanting VVA technology affordably</li>
<li class="best-for-item">Daily city riding</li>
<li class="best-for-item">Young riders preferring aggressive styling</li>
<li class="best-for-item">Those upgrading from smaller displacement bikes</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা এমটি ১৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Long-distance highway touring</li>
<li class="not-recommended-item">Riders requiring ABS safety</li>
<li class="not-recommended-item">Those needing wind protection</li>
<li class="not-recommended-item">Budget buyers seeking maximum value</li>
<li class="not-recommended-item">Riders preferring conservative styling</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা এমটি ১৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৪,৬০,০০০ - এমটি সিরিজে ভিভিএ প্রযুক্তির জন্য ভাল ভ্যালু</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি - ৳৯-১২ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৮-১০ প্রতি কিমি (ভিভিএ দক্ষতা সহ ৩৭-৪২ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৭,৫০০-১২,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা এমটি ১৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- চমৎকার ভিভিএ পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- পারফরমেন্সের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- ভাল সোজা অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- আইকনিক এমটি ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ভিভিএ বৈশিষ্ট্যের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- নির্ভরযোগ্য ইয়ামাহা বিল্ড</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা এমটি ১৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এমটি ১৫ অন্যান্য ১৫৫সিসি নেকেড বাইকের সাথে কেমন তুলনা করে?</h3>
<p class="faq-answer">এমটি ১৫ ভিভিএ প্রযুক্তি, আক্রমণাত্মক এমটি-সিরিজ স্টাইলিং এবং ইয়ামাহা বিল্ড কোয়ালিটি দিয়ে আলাদা হয়ে দাঁড়ায়। এটি প্রচলিত ১৫৫সিসি বাইকের চেয়ে ভাল পারফরমেন্স প্রদান করে তবে প্রিমিয়াম দামে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ভিভিএ প্রযুক্তি কি অতিরিক্ত খরচের যোগ্য?</h3>
<p class="faq-answer">ভিভিএ ভাল পাওয়ার ডেলিভারি, উন্নত জ্বালানি দক্ষতা এবং রেভ রেঞ্জ জুড়ে উন্নত পারফরমেন্স প্রদান করে। পারফরমেন্স-মনস্ক রাইডারদের জন্য, এটি একটি সার্থক বিনিয়োগ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কী রক্ষণাবেক্ষণ আশা করা উচিত?</h3>
<p class="faq-answer">প্রতি ৩,৫০০ কিমিতে নিয়মিত সেবা, ভিভিএ সিস্টেমের জন্য মানসম্পন্ন ইঞ্জিন তেল, এয়ার ফিল্টার রক্ষণাবেক্ষণ এবং চেইন লুব্রিকেশন। ভিভিএর মাঝে মাঝে বিশেষায়িত মনোযোগ প্রয়োজন কিন্তু সাধারণত নির্ভরযোগ্য।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা এমটি ১৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৪,৬০,০০০ টাকায় ইয়ামাহা এমটি ১৫ ইয়ামাহার প্রিমিয়াম ভিভিএ প্রযুক্তি এবং আইকনিক এমটি-সিরিজ ডিজাইনে একটি অ্যাক্সেসযোগ্য প্রবেশ হিসেবে চমৎকার ভ্যালু প্রদান করে। এটি উচ্চতর এমটি মডেলের চেয়ে আরও সাশ্রয়ী মূল্যে শক্তিশালী পারফরমেন্স, আক্রমণাত্মক স্টাইলিং এবং নির্ভরযোগ্য বিল্ড কোয়ালিটি প্রদান করে। তবে, এবিএসের অভাব, সীমিত বায়ু সুরক্ষা, মাঝারি জ্বালানি দক্ষতা এবং প্রিমিয়াম রক্ষণাবেক্ষণের প্রয়োজনীয়তা এটিকে প্রধানত শহুরে রাইডার এবং পারফরমেন্স উৎসাহীদের জন্য উপযুক্ত করে তোলে যারা আরাম বৈশিষ্ট্য এবং সর্বোচ্চ ভ্যালুর চেয়ে ভিভিএ প্রযুক্তি এবং স্ট্রিটফাইটার নান্দনিকতাকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আইকনিক এমটি স্টাইলিং সহ অ্যাক্সেসযোগ্য ভিভিএ পারফরমেন্সের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha MT 15 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha MT 15\n")

	return nil
}
