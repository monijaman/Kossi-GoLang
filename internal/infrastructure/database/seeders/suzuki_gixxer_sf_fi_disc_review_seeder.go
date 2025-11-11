package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxerSFFiDiscReview handles seeding Suzuki Gixxer SF FI Disc product review and translation
type SuzukiGixxerSFFiDiscReview struct {
	BaseSeeder
}

// NewSuzukiGixxerSFFiDiscReviewSeeder creates a new SuzukiGixxerSFFiDiscReview
func NewSuzukiGixxerSFFiDiscReviewSeeder() *SuzukiGixxerSFFiDiscReview {
	return &SuzukiGixxerSFFiDiscReview{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer SF FI Disc Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxerSFFiDiscReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer SF FI Disc").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Gixxer SF FI Disc product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Gixxer SF FI Disc product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer SF FI Disc already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer SF FI Disc Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Gixxer SF FI Disc is a fully-faired 155cc sports motorcycle featuring fuel injection technology and disc brake system. Priced at ৳3,19,950, it combines modern fuel injection efficiency with aerodynamic full fairing design, disc brake safety, and Suzuki's proven reliability in an affordable sports package for daily commuting and weekend touring.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer SF FI Disc Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection Technology:</strong> <span class="highlight-value">Modern FI system for better efficiency and performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Full Fairing Design:</strong> <span class="highlight-value">Aerodynamic protection and sporty appearance</span></li>
<li class="highlight-item"><strong class="highlight-label">Front Disc Brake:</strong> <span class="highlight-value">Enhanced braking performance and safety</span></li>
<li class="highlight-item"><strong class="highlight-label">155cc Engine:</strong> <span class="highlight-value">Reliable SOHC engine with good fuel economy</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer SF FI Disc Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Fuel Injection Benefits:</strong> <span class="pro-description">Better fuel efficiency, smoother performance, and reduced emissions</span></li>
<li class="pro-item"><strong class="pro-label">Full Fairing Advantage:</strong> <span class="pro-description">Excellent wind protection and aerodynamic styling</span></li>
<li class="pro-item"><strong class="pro-label">Disc Brake Safety:</strong> <span class="pro-description">Superior braking performance compared to drum brakes</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">Achieves 40-45 km/l with fuel injection technology</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Reliability:</strong> <span class="pro-description">Proven Japanese engineering and build quality</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Sports Package:</strong> <span class="pro-description">Full fairing sports bike at reasonable price point</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer SF FI Disc Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Power Output:</strong> <span class="con-description">155cc engine may feel underpowered for highway use</span></li>
<li class="con-item"><strong class="con-label">No ABS Feature:</strong> <span class="con-description">Missing anti-lock braking system for enhanced safety</span></li>
<li class="con-item"><strong class="con-label">Heavy City Riding:</strong> <span class="con-description">Full fairing makes it less maneuverable in traffic</span></li>
<li class="con-item"><strong class="con-label">Maintenance Cost:</strong> <span class="con-description">Fuel injection system requires specialized service</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Suzuki dealerships compared to other brands</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Gixxer SF FI Disc in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters wanting sports bike styling</li>
<li class="best-for-item">Riders upgrading from 125cc to 155cc segment</li>
<li class="best-for-item">Those prioritizing fuel injection technology</li>
<li class="best-for-item">Weekend touring enthusiasts</li>
<li class="best-for-item">Riders wanting wind protection with fairing</li>
<li class="best-for-item">Suzuki brand loyalists</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Gixxer SF FI Disc?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Power-hungry riders seeking performance</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Budget buyers seeking maximum value</li>
<li class="not-recommended-item">Riders requiring ABS safety features</li>
<li class="not-recommended-item">Those needing extensive service support</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer SF FI Disc Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,19,950 - Moderate for 155cc FI with fairing</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-10 per km with good fuel economy</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳8-10 per km (40-45 km/l at current fuel prices)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months for FI system</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳5,000-7,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer SF FI Disc Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Adequate for daily use</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Decent touring position</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Attractive sports design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Fair for features offered</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Good Suzuki standards</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Gixxer SF FI Disc Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What are the benefits of fuel injection over carburetor?</h3>
<p class="faq-answer">Fuel injection provides better fuel efficiency, smoother performance, instant start, reduced emissions, and consistent performance across different weather conditions compared to carburetor systems.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does the disc brake compare to drum brake?</h3>
<p class="faq-answer">Disc brakes provide better stopping power, less fade under heavy use, better heat dissipation, and more consistent performance in wet conditions compared to drum brakes.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is this bike suitable for highway riding?</h3>
<p class="faq-answer">Yes, the full fairing provides good wind protection and the bike can cruise comfortably at 80-90 km/h. However, acceleration may feel limited for aggressive highway riding due to 155cc displacement.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer SF FI Disc in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.0/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,19,950, the Suzuki Gixxer SF FI Disc offers a balanced combination of fuel injection technology, full fairing protection, disc brake safety, and Suzuki reliability in an affordable 155cc sports package. The FI system ensures good fuel efficiency of 40-45 km/l while the full fairing provides adequate wind protection for touring. However, limited power output for highway performance, lack of ABS, heavy handling in city traffic, and specialized FI maintenance requirements make it suitable primarily for daily commuters and weekend riders who prioritize fuel efficiency and wind protection over outright performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For daily commuting with sports styling and FI efficiency</span></p>
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
		return fmt.Errorf("error creating Suzuki Gixxer SF FI Disc review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer SF FI Disc (Rating: 4.0/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার এসএফ এফআই ডিস্ক রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি গিক্সার এসএফ এফআই ডিস্ক হল একটি ফুল-ফেয়ার্ড ১৫৫সিসি স্পোর্টস মোটরসাইকেল যেখানে ফুয়েল ইনজেকশন প্রযুক্তি এবং ডিস্ক ব্রেক সিস্টেম রয়েছে। ৳৩,১৯,৯৫০ টাকায় মূল্যায়িত, এটি আধুনিক ফুয়েল ইনজেকশন দক্ষতাকে অ্যারোডাইনামিক ফুল ফেয়ারিং ডিজাইন, ডিস্ক ব্রেক নিরাপত্তা এবং দৈনিক যাতায়াত ও সাপ্তাহিক ট্যুরিংয়ের জন্য একটি সাশ্রয়ী স্পোর্টস প্যাকেজে সুজুকির প্রমাণিত নির্ভরযোগ্যতার সাথে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার এসএফ এফআই ডিস্ক এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন প্রযুক্তি:</strong> <span class="highlight-value">ভাল দক্ষতা এবং পারফরমেন্সের জন্য আধুনিক এফআই সিস্টেম</span></li>
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ ফেয়ারিং ডিজাইন:</strong> <span class="highlight-value">অ্যারোডাইনামিক সুরক্ষা এবং স্পোর্টি উপস্থিতি</span></li>
<li class="highlight-item"><strong class="highlight-label">সামনের ডিস্ক ব্রেক:</strong> <span class="highlight-value">উন্নত ব্রেকিং পারফরমেন্স এবং নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">১৫৫সিসি ইঞ্জিন:</strong> <span class="highlight-value">ভাল জ্বালানি সাশ্রয় সহ নির্ভরযোগ্য এসওএইচসি ইঞ্জিন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার এসএফ এফআই ডিস্ক এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন সুবিধা:</strong> <span class="pro-description">ভাল জ্বালানি দক্ষতা, মসৃণ পারফরমেন্স এবং কম নির্গমন</span></li>
<li class="pro-item"><strong class="pro-label">সম্পূর্ণ ফেয়ারিং সুবিধা:</strong> <span class="pro-description">চমৎকার বাতাস সুরক্ষা এবং অ্যারোডাইনামিক স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">ডিস্ক ব্রেক নিরাপত্তা:</strong> <span class="pro-description">ড্রাম ব্রেকের তুলনায় উন্নত ব্রেকিং পারফরমেন্স</span></li>
<li class="pro-item"><strong class="pro-label">ভাল জ্বালানি সাশ্রয়:</strong> <span class="pro-description">ফুয়েল ইনজেকশন প্রযুক্তি দিয়ে ৪০-৪৫ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত জাপানি ইঞ্জিনিয়ারিং এবং বিল্ড কোয়ালিটি</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী স্পোর্টস প্যাকেজ:</strong> <span class="pro-description">যুক্তিসঙ্গত মূল্যে ফুল ফেয়ারিং স্পোর্টস বাইক</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার এসএফ এফআই ডিস্ক এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত শক্তি আউটপুট:</strong> <span class="con-description">হাইওয়ে ব্যবহারের জন্য ১৫৫সিসি ইঞ্জিন কম শক্তিশালী মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস ফিচার নেই:</strong> <span class="con-description">উন্নত নিরাপত্তার জন্য অ্যান্টি-লক ব্রেকিং সিস্টেমের অভাব</span></li>
<li class="con-item"><strong class="con-label">ভারী শহুরে রাইডিং:</strong> <span class="con-description">সম্পূর্ণ ফেয়ারিং ট্রাফিকে এটিকে কম চালিত করে তোলে</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">ফুয়েল ইনজেকশন সিস্টেমের বিশেষায়িত সেবা প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">অন্যান্য ব্র্যান্ডের তুলনায় কম সুজুকি ডিলারশিপ</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি গিক্সার এসএফ এফআই ডিস্ক কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters wanting sports bike styling</li>
<li class="best-for-item">Riders upgrading from 125cc to 155cc segment</li>
<li class="best-for-item">Those prioritizing fuel injection technology</li>
<li class="best-for-item">Weekend touring enthusiasts</li>
<li class="best-for-item">Riders wanting wind protection with fairing</li>
<li class="best-for-item">Suzuki brand loyalists</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি গিক্সার এসএফ এফআই ডিস্ক কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Power-hungry riders seeking performance</li>
<li class="not-recommended-item">Heavy city traffic navigation</li>
<li class="not-recommended-item">Budget buyers seeking maximum value</li>
<li class="not-recommended-item">Riders requiring ABS safety features</li>
<li class="not-recommended-item">Those needing extensive service support</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি গিক্সার এসএফ এফআই ডিস্ক এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৩,১৯,৯৫০ - ফেয়ারিং সহ ১৫৫সিসি এফআই এর জন্য মধ্যম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মধ্যম - ভাল জ্বালানি সাশ্রয় সহ ৳৮-১০ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৮-১০ প্রতি কিমি (বর্তমান জ্বালানি দামে ৪০-৪৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">এফআই সিস্টেমের জন্য প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৫,০০০-৭,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি গিক্সার এসএফ এফআই ডিস্ক পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- দৈনিক ব্যবহারের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- ভাল জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- শোভন ট্যুরিং অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- আকর্ষণীয় স্পোর্টস ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- প্রদত্ত ফিচারের জন্য ন্যায্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- ভাল সুজুকি মান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি গিক্সার এসএফ এফআই ডিস্ক সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">কার্বুরেটরের চেয়ে ফুয়েল ইনজেকশনের সুবিধা কী?</h3>
<p class="faq-answer">ফুয়েল ইনজেকশন কার্বুরেটর সিস্টেমের তুলনায় ভাল জ্বালানি দক্ষতা, মসৃণ পারফরমেন্স, তাৎক্ষণিক স্টার্ট, কম নির্গমন এবং বিভিন্ন আবহাওয়ার অবস্থায় সামঞ্জস্যপূর্ণ পারফরমেন্স প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ডিস্ক ব্রেক ড্রাম ব্রেকের সাথে কেমন তুলনা করে?</h3>
<p class="faq-answer">ডিস্ক ব্রেক ড্রাম ব্রেকের তুলনায় ভাল স্টপিং পাওয়ার, ভারী ব্যবহারে কম ফেড, ভাল তাপ নিঃসরণ এবং ভেজা অবস্থায় আরো সামঞ্জস্যপূর্ণ পারফরমেন্স প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এই বাইকটি হাইওয়ে রাইডিংয়ের জন্য উপযুক্ত?</h3>
<p class="faq-answer">হ্যাঁ, সম্পূর্ণ ফেয়ারিং ভাল বাতাস সুরক্ষা প্রদান করে এবং বাইকটি ৮০-৯০ কিমি/ঘণ্টায় আরামদায়কভাবে ক্রুজ করতে পারে। তবে, ১৫৫সিসি ডিসপ্লেসমেন্টের কারণে আক্রমণাত্মক হাইওয়ে রাইডিংয়ের জন্য ত্বরণ সীমিত মনে হতে পারে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি গিক্সার এসএফ এফআই ডিস্ক কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.0/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৩,১৯,৯৫০ টাকায় সুজুকি গিক্সার এসএফ এফআই ডিস্ক একটি সাশ্রয়ী ১৫৫সিসি স্পোর্টস প্যাকেজে ফুয়েল ইনজেকশন প্রযুক্তি, ফুল ফেয়ারিং সুরক্ষা, ডিস্ক ব্রেক নিরাপত্তা এবং সুজুকি নির্ভরযোগ্যতার একটি সুষম সমন্বয় প্রদান করে। এফআই সিস্টেম ৪০-৪৫ কিমি/লিটার ভাল জ্বালানি দক্ষতা নিশ্চিত করে যখন ফুল ফেয়ারিং ট্যুরিংয়ের জন্য পর্যাপ্ত বাতাস সুরক্ষা প্রদান করে। তবে, হাইওয়ে পারফরমেন্সের জন্য সীমিত শক্তি আউটপুট, এবিএসের অভাব, শহুরে ট্রাফিকে ভারী হ্যান্ডলিং এবং বিশেষায়িত এফআই রক্ষণাবেক্ষণের প্রয়োজনীয়তা এটিকে প্রধানত দৈনিক যাত্রী এবং সাপ্তাহিক রাইডারদের জন্য উপযুক্ত করে তোলে যারা সরাসরি পারফরমেন্সের চেয়ে জ্বালানি দক্ষতা এবং বাতাস সুরক্ষাকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্পোর্টস স্টাইলিং এবং এফআই দক্ষতা সহ দৈনিক যাতায়াতের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer SF FI Disc already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer SF FI Disc\n")

	return nil
}
