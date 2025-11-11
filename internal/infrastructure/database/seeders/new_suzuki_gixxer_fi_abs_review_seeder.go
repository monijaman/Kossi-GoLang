package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// NewSuzukiGixxerFiABSReview handles seeding New Suzuki Gixxer Fi ABS product review and translation
type NewSuzukiGixxerFiABSReview struct {
	BaseSeeder
}

// NewNewSuzukiGixxerFiABSReviewSeeder creates a new NewSuzukiGixxerFiABSReview
func NewNewSuzukiGixxerFiABSReviewSeeder() *NewSuzukiGixxerFiABSReview {
	return &NewSuzukiGixxerFiABSReview{BaseSeeder: BaseSeeder{name: "New Suzuki Gixxer Fi ABS Review"}}
}

// Seed implements the Seeder interface
func (s *NewSuzukiGixxerFiABSReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "New Suzuki Gixxer Fi ABS").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("New Suzuki Gixxer Fi ABS product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding New Suzuki Gixxer Fi ABS product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for New Suzuki Gixxer Fi ABS already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">New Suzuki Gixxer Fi ABS Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The New Suzuki Gixxer Fi ABS is an updated 155cc naked motorcycle featuring fuel injection technology and single-channel ABS braking system. Priced at ৳2,79,950, it combines modern FI efficiency with enhanced safety features, aggressive naked styling, and Suzuki reliability for daily commuting and spirited riding.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">New Suzuki Gixxer Fi ABS Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection System:</strong> <span class="highlight-value">Advanced FI technology for optimal performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Single Channel ABS:</strong> <span class="highlight-value">Front wheel ABS for enhanced safety</span></li>
<li class="highlight-item"><strong class="highlight-label">Naked Streetfighter Design:</strong> <span class="highlight-value">Aggressive styling with minimal bodywork</span></li>
<li class="highlight-item"><strong class="highlight-label">155cc SOHC Engine:</strong> <span class="highlight-value">Reliable and fuel-efficient powerplant</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">New Suzuki Gixxer Fi ABS Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Fuel Injection Benefits:</strong> <span class="pro-description">Better fuel economy, instant starts, and cleaner emissions</span></li>
<li class="pro-item"><strong class="pro-label">ABS Safety Feature:</strong> <span class="pro-description">Anti-lock braking prevents wheel lockup during emergency stops</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight and Agile:</strong> <span class="pro-description">Easy to handle in city traffic and tight spaces</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">Achieves 45-50 km/l with fuel injection technology</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Sharp and modern naked bike design appeals to younger riders</span></li>
<li class="pro-item"><strong class="pro-label">Reasonable Pricing:</strong> <span class="pro-description">Competitive pricing for FI and ABS features</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">New Suzuki Gixxer Fi ABS Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Single Channel ABS Only:</strong> <span class="con-description">Lacks rear wheel ABS for complete braking safety</span></li>
<li class="con-item"><strong class="con-label">Limited Power Output:</strong> <span class="con-description">155cc engine may feel underpowered for highway riding</span></li>
<li class="con-item"><strong class="con-label">No Wind Protection:</strong> <span class="con-description">Naked design offers no protection from wind at higher speeds</span></li>
<li class="con-item"><strong class="con-label">Maintenance Cost:</strong> <span class="con-description">Fuel injection system requires specialized service support</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Suzuki service centers compared to local brands</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy New Suzuki Gixxer Fi ABS in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily urban commuters</li>
<li class="best-for-item">Young riders seeking modern features</li>
<li class="best-for-item">Those prioritizing fuel injection efficiency</li>
<li class="best-for-item">Riders wanting ABS safety in this price range</li>
<li class="best-for-item">Motorcycle enthusiasts preferring naked styling</li>
<li class="best-for-item">First-time buyers upgrading from smaller bikes</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy New Suzuki Gixxer Fi ABS?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway touring enthusiasts</li>
<li class="not-recommended-item">Riders needing wind protection</li>
<li class="not-recommended-item">Power-hungry performance seekers</li>
<li class="not-recommended-item">Those requiring extensive service network</li>
<li class="not-recommended-item">Budget buyers seeking maximum value</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">New Suzuki Gixxer Fi ABS Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,79,950 - Good value for FI and ABS features</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low to moderate - ৳7-9 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳7-9 per km (45-50 km/l with FI efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳4,500-6,500 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">New Suzuki Gixxer Fi ABS Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Good for city riding</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent with FI</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Decent for naked bike</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Appealing naked design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Fair for features</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Good Suzuki standards</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">New Suzuki Gixxer Fi ABS Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does the fuel injection system improve performance?</h3>
<p class="faq-answer">Fuel injection provides precise fuel delivery, better atomization, improved combustion efficiency, reduced emissions, and consistent performance across different weather conditions compared to carburetor systems.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is single-channel ABS effective for safety?</h3>
<p class="faq-answer">Single-channel ABS on the front wheel provides significant safety improvement by preventing front wheel lockup during emergency braking. While dual-channel would be better, single-channel ABS is still very beneficial.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How is the naked bike design for city riding?</h3>
<p class="faq-answer">Naked bikes are excellent for city riding due to their lightweight nature, upright riding position, easy maneuverability in traffic, and good visibility. However, they lack wind protection for highway use.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy New Suzuki Gixxer Fi ABS in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,79,950, the New Suzuki Gixxer Fi ABS offers a compelling package of fuel injection efficiency, ABS safety, naked bike agility, and modern styling in the affordable 155cc segment. The FI system ensures excellent fuel economy of 45-50 km/l while the ABS adds crucial safety enhancement. However, single-channel ABS limitation, modest power output, lack of wind protection, specialized FI maintenance requirements, and limited Suzuki service network make it suitable primarily for daily urban commuters and young riders who prioritize modern features and efficiency over outright performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For modern FI efficiency with ABS safety in naked styling</span></p>
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
		return fmt.Errorf("error creating New Suzuki Gixxer Fi ABS review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for New Suzuki Gixxer Fi ABS (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">নিউ সুজুকি গিক্সার এফআই এবিএস রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">নিউ সুজুকি গিক্সার এফআই এবিএস হল একটি আপডেটেড ১৫৫সিসি নেকেড মোটরসাইকেল যেখানে ফুয়েল ইনজেকশন প্রযুক্তি এবং সিঙ্গেল-চ্যানেল এবিএস ব্রেকিং সিস্টেম রয়েছে। ৳২,৭৯,৯৫০ টাকায় মূল্যায়িত, এটি দৈনিক যাতায়াত এবং উৎসাহী রাইডিংয়ের জন্য উন্নত নিরাপত্তা বৈশিষ্ট্য, আক্রমণাত্মক নেকেড স্টাইলিং এবং সুজুকি নির্ভরযোগ্যতার সাথে আধুনিক এফআই দক্ষতাকে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">নিউ সুজুকি গিক্সার এফআই এবিএস এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন সিস্টেম:</strong> <span class="highlight-value">সর্বোত্তম পারফরমেন্সের জন্য উন্নত এফআই প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="highlight-value">উন্নত নিরাপত্তার জন্য সামনের চাকার এবিএস</span></li>
<li class="highlight-item"><strong class="highlight-label">নেকেড স্ট্রিটফাইটার ডিজাইন:</strong> <span class="highlight-value">ন্যূনতম বডিওয়ার্ক সহ আক্রমণাত্মক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">১৫৫সিসি এসওএইচসি ইঞ্জিন:</strong> <span class="highlight-value">নির্ভরযোগ্য এবং জ্বালানি-দক্ষ পাওয়ারপ্ল্যান্ট</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">নিউ সুজুকি গিক্সার এফআই এবিএস এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন সুবিধা:</strong> <span class="pro-description">ভাল জ্বালানি অর্থনীতি, তাৎক্ষণিক স্টার্ট এবং পরিষ্কার নির্গমন</span></li>
<li class="pro-item"><strong class="pro-label">এবিএস নিরাপত্তা বৈশিষ্ট্য:</strong> <span class="pro-description">অ্যান্টি-লক ব্রেকিং জরুরি থামানোর সময় হুইল লকআপ প্রতিরোধ করে</span></li>
<li class="pro-item"><strong class="pro-label">হালকা এবং চটপটে:</strong> <span class="pro-description">শহুরে ট্রাফিক এবং সংকীর্ণ জায়গায় পরিচালনা সহজ</span></li>
<li class="pro-item"><strong class="pro-label">ভাল জ্বালানি সাশ্রয়:</strong> <span class="pro-description">ফুয়েল ইনজেকশন প্রযুক্তি দিয়ে ৪৫-৫০ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">তীক্ষ্ণ এবং আধুনিক নেকেড বাইক ডিজাইন তরুণ রাইডারদের আকর্ষণ করে</span></li>
<li class="pro-item"><strong class="pro-label">যুক্তিসঙ্গত মূল্য:</strong> <span class="pro-description">এফআই এবং এবিএস ফিচারের জন্য প্রতিযোগিতামূলক মূল্য</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">নিউ সুজুকি গিক্সার এফআই এবিএস এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">কেবলমাত্র সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="con-description">সম্পূর্ণ ব্রেকিং নিরাপত্তার জন্য পিছনের চাকার এবিএসের অভাব</span></li>
<li class="con-item"><strong class="con-label">সীমিত শক্তি আউটপুট:</strong> <span class="con-description">হাইওয়ে রাইডিংয়ের জন্য ১৫৫সিসি ইঞ্জিন কম শক্তিশালী মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">কোন বায়ু সুরক্ষা নেই:</strong> <span class="con-description">নেকেড ডিজাইন উচ্চ গতিতে বাতাস থেকে কোন সুরক্ষা প্রদান করে না</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">ফুয়েল ইনজেকশন সিস্টেমের বিশেষায়িত সেবা সহায়তা প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">স্থানীয় ব্র্যান্ডের তুলনায় কম সুজুকি সেবা কেন্দ্র</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে নিউ সুজুকি গিক্সার এফআই এবিএস কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily urban commuters</li>
<li class="best-for-item">Young riders seeking modern features</li>
<li class="best-for-item">Those prioritizing fuel injection efficiency</li>
<li class="best-for-item">Riders wanting ABS safety in this price range</li>
<li class="best-for-item">Motorcycle enthusiasts preferring naked styling</li>
<li class="best-for-item">First-time buyers upgrading from smaller bikes</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">নিউ সুজুকি গিক্সার এফআই এবিএস কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway touring enthusiasts</li>
<li class="not-recommended-item">Riders needing wind protection</li>
<li class="not-recommended-item">Power-hungry performance seekers</li>
<li class="not-recommended-item">Those requiring extensive service network</li>
<li class="not-recommended-item">Budget buyers seeking maximum value</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">নিউ সুজুকি গিক্সার এফআই এবিএস এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৭৯,৯৫০ - এফআই এবং এবিএস ফিচারের জন্য ভাল ভ্যালু</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">নিম্ন থেকে মধ্যম - ৳৭-৯ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৭-৯ প্রতি কিমি (এফআই দক্ষতা সহ ৪৫-৫০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৪,৫০০-৬,৫০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">নিউ সুজুকি গিক্সার এফআই এবিএস পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- শহুরে রাইডিংয়ের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- এফআই সহ চমৎকার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- নেকেড বাইকের জন্য শোভন</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- আকর্ষণীয় নেকেড ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ফিচারের জন্য ন্যায্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- ভাল সুজুকি মান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">নিউ সুজুকি গিক্সার এফআই এবিএস সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ফুয়েল ইনজেকশন সিস্টেম কীভাবে পারফরমেন্স উন্নত করে?</h3>
<p class="faq-answer">ফুয়েল ইনজেকশন কার্বুরেটর সিস্টেমের তুলনায় সুনির্দিষ্ট জ্বালানি সরবরাহ, ভাল পরমাণুকরণ, উন্নত দহন দক্ষতা, হ্রাসকৃত নির্গমন এবং বিভিন্ন আবহাওয়ার অবস্থায় সামঞ্জস্যপূর্ণ পারফরমেন্স প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">নিরাপত্তার জন্য সিঙ্গেল-চ্যানেল এবিএস কি কার্যকর?</h3>
<p class="faq-answer">সামনের চাকায় সিঙ্গেল-চ্যানেল এবিএস জরুরি ব্রেকিংয়ের সময় সামনের চাকা লকআপ প্রতিরোধ করে উল্লেখযোগ্য নিরাপত্তা উন্নতি প্রদান করে। ডুয়াল-চ্যানেল ভাল হলেও, সিঙ্গেল-চ্যানেল এবিএস এখনও অত্যন্ত উপকারী।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">শহুরে রাইডিংয়ের জন্য নেকেড বাইক ডিজাইন কেমন?</h3>
<p class="faq-answer">নেকেড বাইকগুলি তাদের হালকা প্রকৃতি, সোজা রাইডিং পজিশন, ট্রাফিকে সহজ চালনা এবং ভাল দৃশ্যমানতার কারণে শহুরে রাইডিংয়ের জন্য চমৎকার। তবে, হাইওয়ে ব্যবহারের জন্য তাদের বায়ু সুরক্ষার অভাব রয়েছে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে নিউ সুজুকি গিক্সার এফআই এবিএস কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.1/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,৭৯,৯৫০ টাকায় নিউ সুজুকি গিক্সার এফআই এবিএস সাশ্রয়ী ১৫৫সিসি সেগমেন্টে ফুয়েল ইনজেকশন দক্ষতা, এবিএস নিরাপত্তা, নেকেড বাইক চটপটেতা এবং আধুনিক স্টাইলিংয়ের একটি আকর্ষণীয় প্যাকেজ প্রদান করে। এফআই সিস্টেম ৪৫-৫০ কিমি/লিটার চমৎকার জ্বালানি অর্থনীতি নিশ্চিত করে যখন এবিএস গুরুত্বপূর্ণ নিরাপত্তা উন্নতি যোগ করে। তবে, সিঙ্গেল-চ্যানেল এবিএস সীমাবদ্ধতা, মাঝারি পাওয়ার আউটপুট, বায়ু সুরক্ষার অভাব, বিশেষায়িত এফআই রক্ষণাবেক্ষণের প্রয়োজনীয়তা এবং সীমিত সুজুকি সেবা নেটওয়ার্ক এটিকে প্রধানত দৈনিক শহুরে যাত্রী এবং তরুণ রাইডারদের জন্য উপযুক্ত করে তোলে যারা সরাসরি পারফরমেন্সের চেয়ে আধুনিক বৈশিষ্ট্য এবং দক্ষতাকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- নেকেড স্টাইলিংয়ে এবিএস নিরাপত্তা সহ আধুনিক এফআই দক্ষতার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for New Suzuki Gixxer Fi ABS already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for New Suzuki Gixxer Fi ABS\n")

	return nil
}
