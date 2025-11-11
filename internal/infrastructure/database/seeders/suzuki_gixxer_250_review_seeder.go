package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxer250Review handles seeding Suzuki Gixxer 250 product review and translation
type SuzukiGixxer250Review struct {
	BaseSeeder
}

// NewSuzukiGixxer250ReviewSeeder creates a new SuzukiGixxer250Review
func NewSuzukiGixxer250ReviewSeeder() *SuzukiGixxer250Review {
	return &SuzukiGixxer250Review{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer 250 Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxer250Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer 250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Gixxer 250 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Gixxer 250 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer 250 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer 250 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Gixxer 250 is a naked street fighter motorcycle powered by a 249cc single-cylinder SOHC engine producing 26.5 HP and 22.6 Nm torque. Priced at ৳3,79,950, it combines aggressive streetfighter styling, digital instrumentation, LED lighting, and Suzuki's proven reliability in the premium quarter-liter naked segment.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer 250 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Powerful Quarter-Liter Engine:</strong> <span class="highlight-value">249cc SOHC engine with 26.5 HP and responsive performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Naked Design:</strong> <span class="highlight-value">Streetfighter styling with muscular fuel tank</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting Package:</strong> <span class="highlight-value">Modern LED headlight and indicators</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Instrument Cluster:</strong> <span class="highlight-value">Advanced display with comprehensive information</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer 250 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Quarter-Liter Performance:</strong> <span class="pro-description">249cc engine provides excellent power and torque delivery</span></li>
<li class="pro-item"><strong class="pro-label">Naked Bike Agility:</strong> <span class="pro-description">Lightweight and nimble handling for city and twisty roads</span></li>
<li class="pro-item"><strong class="pro-label">Modern LED Lighting:</strong> <span class="pro-description">Full LED headlight and indicators for contemporary appeal</span></li>
<li class="pro-item"><strong class="pro-label">Digital Instrumentation:</strong> <span class="pro-description">Comprehensive display with gear indicator and trip functions</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Upright riding position suitable for daily commuting</span></li>
<li class="pro-item"><strong class="pro-label">Better City Maneuverability:</strong> <span class="pro-description">Naked design makes traffic navigation easier</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer 250 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Pricing:</strong> <span class="con-description">৳3,79,950 is expensive for 250cc single-cylinder naked bike</span></li>
<li class="con-item"><strong class="con-label">No ABS Available:</strong> <span class="con-description">Missing critical safety feature at this premium price point</span></li>
<li class="con-item"><strong class="con-label">Single Cylinder Vibrations:</strong> <span class="con-description">Engine buzz noticeable at higher RPMs</span></li>
<li class="con-item"><strong class="con-label">No Wind Protection:</strong> <span class="con-description">Naked design offers no wind shielding for highway rides</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">Premium positioning means higher service costs</span></li>
<li class="con-item"><strong class="con-label">Poor Value Proposition:</strong> <span class="con-description">Overpriced compared to competitors with more features</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Gixxer 250 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban commuting and city riding</li>
<li class="best-for-item">Quarter-liter naked bike enthusiasts</li>
<li class="best-for-item">Riders wanting agile handling characteristics</li>
<li class="best-for-item">Daily commuting with weekend touring</li>
<li class="best-for-item">Experienced riders upgrading from 150cc naked bikes</li>
<li class="best-for-item">Those preferring Suzuki brand reliability</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Gixxer 250?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers seeking value</li>
<li class="not-recommended-item">Highway touring without wind protection</li>
<li class="not-recommended-item">Riders requiring ABS safety features</li>
<li class="not-recommended-item">New or inexperienced motorcyclists</li>
<li class="not-recommended-item">Those needing extensive service network</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer 250 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,79,950 - Premium for Quarter-Liter Naked</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate to High - ৳11-13 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳11-13 per km (38-42 km/l at current fuel prices)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳7,000-10,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer 250 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Excellent power and torque</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good fuel economy</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Comfortable city ergonomics</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Aggressive streetfighter look</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.2</span> <span class="rating-note">- Overpriced for segment</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Solid Suzuki build</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Gixxer 250 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How is the Gixxer 250 for city commuting?</h3>
<p class="faq-answer">Excellent for city use with upright ergonomics, good maneuverability, and decent fuel economy of 38-42 km/l. Naked design makes traffic navigation easy.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the difference between Gixxer 250 and Gixxer SF 250?</h3>
<p class="faq-answer">Gixxer 250 is naked with better city maneuverability and lower price, while SF 250 has full fairing for highway wind protection but costs ৳50,000 more.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the 249cc engine reliable for long rides?</h3>
<p class="faq-answer">Yes, the SOHC engine is reliable with good torque delivery. However, single-cylinder vibrations may become noticeable on very long highway rides without wind protection.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer 250 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,79,950, the Suzuki Gixxer 250 offers solid quarter-liter naked bike performance with good city agility, modern LED lighting, digital instrumentation, and reliable Suzuki engineering. The 249cc engine provides decent power while the naked design ensures easy maneuverability in urban conditions. However, the high pricing without ABS, single-cylinder vibrations, expensive maintenance, lack of wind protection for highway use, and poor value proposition compared to competitors make it suitable only for Suzuki loyalists who prioritize brand heritage and naked bike styling over practical value considerations.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For urban quarter-liter naked bike enthusiasts</span></p>
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
		return fmt.Errorf("error creating Suzuki Gixxer 250 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer 250 (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার ২৫০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি গিক্সার ২৫০ হল একটি নেকেড স্ট্রিট ফাইটার মোটরসাইকেল যা ২৪৯সিসি একক সিলিন্ডার এসওএইচসি ইঞ্জিন দ্বারা চালিত যা ২৬.৫ এইচপি এবং ২২.৬ এনএম টর্ক উৎপাদন করে। ৳৩,৭৯,৯৫০ টাকায় মূল্যায়িত, এটি আক্রমণাত্মক স্ট্রিটফাইটার স্টাইলিং, ডিজিটাল ইনস্ট্রুমেন্টেশন, এলইডি লাইটিং এবং প্রিমিয়াম কোয়ার্টার-লিটার নেকেড সেগমেন্টে সুজুকির প্রমাণিত নির্ভরযোগ্যতা একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার ২৫০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">শক্তিশালী কোয়ার্টার-লিটার ইঞ্জিন:</strong> <span class="highlight-value">২৬.৫ এইচপি এবং প্রতিক্রিয়াশীল পারফরমেন্স সহ ২৪৯সিসি এসওএইচসি ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক নেকেড ডিজাইন:</strong> <span class="highlight-value">পেশীবহুল ফুয়েল ট্যাঙ্ক সহ স্ট্রিটফাইটার স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং প্যাকেজ:</strong> <span class="highlight-value">আধুনিক এলইডি হেডলাইট এবং ইন্ডিকেটর</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল ইনস্ট্রুমেন্ট ক্লাস্টার:</strong> <span class="highlight-value">ব্যাপক তথ্য সহ উন্নত ডিসপ্লে</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার ২৫০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী কোয়ার্টার-লিটার পারফরমেন্স:</strong> <span class="pro-description">২৪৯সিসি ইঞ্জিন চমৎকার শক্তি এবং টর্ক ডেলিভারি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">নেকেড বাইক চপলতা:</strong> <span class="pro-description">শহর এবং আঁকাবাঁকা রাস্তার জন্য হালকা এবং ক্ষিপ্র হ্যান্ডলিং</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক এলইডি লাইটিং:</strong> <span class="pro-description">সমসাময়িক আবেদনের জন্য সম্পূর্ণ এলইডি হেডলাইট এবং ইন্ডিকেটর</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল ইনস্ট্রুমেন্টেশন:</strong> <span class="pro-description">গিয়ার ইন্ডিকেটর এবং ট্রিপ ফাংশন সহ ব্যাপক ডিসপ্লে</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এরগোনমিক্স:</strong> <span class="pro-description">দৈনিক যাতায়াতের জন্য উপযুক্ত খাড়া রাইডিং অবস্থান</span></li>
<li class="pro-item"><strong class="pro-label">উন্নত শহুরে চালনা:</strong> <span class="pro-description">নেকেড ডিজাইন ট্রাফিক নেভিগেশনকে সহজ করে তোলে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার ২৫০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">২৫০সিসি একক সিলিন্ডার নেকেড বাইকের জন্য ৳৩,৭৯,৯৫০ ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস উপলব্ধ নেই:</strong> <span class="con-description">এই প্রিমিয়াম মূল্যে গুরুত্বপূর্ণ নিরাপত্তা বৈশিষ্ট্যের অভাব</span></li>
<li class="con-item"><strong class="con-label">একক সিলিন্ডার কম্পন:</strong> <span class="con-description">উচ্চ আরপিএমে ইঞ্জিন গুঞ্জন লক্ষণীয়</span></li>
<li class="con-item"><strong class="con-label">কোনো বাতাস সুরক্ষা নেই:</strong> <span class="con-description">নেকেড ডিজাইন হাইওয়ে রাইডের জন্য কোনো বাতাস রক্ষা প্রদান করে না</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">প্রিমিয়াম অবস্থান মানে উচ্চ সেবা খরচ</span></li>
<li class="con-item"><strong class="con-label">খারাপ মূল্য প্রস্তাব:</strong> <span class="con-description">আরো ফিচার সহ প্রতিযোগীদের তুলনায় অতিরিক্ত দামি</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি গিক্সার ২৫০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban commuting and city riding</li>
<li class="best-for-item">Quarter-liter naked bike enthusiasts</li>
<li class="best-for-item">Riders wanting agile handling characteristics</li>
<li class="best-for-item">Daily commuting with weekend touring</li>
<li class="best-for-item">Experienced riders upgrading from 150cc naked bikes</li>
<li class="best-for-item">Those preferring Suzuki brand reliability</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি গিক্সার ২৫০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers seeking value</li>
<li class="not-recommended-item">Highway touring without wind protection</li>
<li class="not-recommended-item">Riders requiring ABS safety features</li>
<li class="not-recommended-item">New or inexperienced motorcyclists</li>
<li class="not-recommended-item">Those needing extensive service network</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি গিক্সার ২৫০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৩,৭৯,৯৫০ - কোয়ার্টার-লিটার নেকেডের জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মধ্যম থেকে উচ্চ - ৳১১-১৩ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১১-১৩ প্রতি কিমি (বর্তমান জ্বালানি দামে ৩৮-৪২ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৭,০০০-১০,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি গিক্সার ২৫০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- চমৎকার শক্তি এবং টর্ক</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভালো জ্বালানি সাশ্রয়</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- আরামদায়ক শহুরে এরগোনমিক্স</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- আক্রমণাত্মক স্ট্রিটফাইটার চেহারা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.2</span> <span class="rating-note">- সেগমেন্টের জন্য বেশি দামি</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- কঠিন সুজুকি নির্মাণ</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি গিক্সার ২৫০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">শহুরে যাতায়াতের জন্য গিক্সার ২৫০ কেমন?</h3>
<p class="faq-answer">খাড়া এরগোনমিক্স, ভাল চালচলন এবং ৩৮-৪২ কিমি/লিটার শোভন জ্বালানি অর্থনীতি সহ শহুরে ব্যবহারের জন্য চমৎকার। নেকেড ডিজাইন ট্রাফিক নেভিগেশন সহজ করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">গিক্সার ২৫০ এবং গিক্সার এসএফ ২৫০ এর মধ্যে পার্থক্য কী?</h3>
<p class="faq-answer">গিক্সার ২৫০ নেকেড যার ভাল শহুরে চালচলন এবং কম দাম, যখন এসএফ ২৫০ এ হাইওয়ে বাতাস সুরক্ষার জন্য ফুল ফেয়ারিং আছে কিন্তু ৳৫০,০০০ বেশি খরচ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দীর্ঘ রাইডের জন্য ২৪৯সিসি ইঞ্জিন কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">হ্যাঁ, এসওএইচসি ইঞ্জিন ভাল টর্ক ডেলিভারি সহ নির্ভরযোগ্য। তবে, বাতাস সুরক্ষা ছাড়াই খুব দীর্ঘ হাইওয়ে রাইডে একক সিলিন্ডার কম্পন লক্ষণীয় হতে পারে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি গিক্সার ২৫০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৩,৭৯,৯৫০ টাকায় সুজুকি গিক্সার ২৫০ ভাল শহুরে চপলতা, আধুনিক এলইডি লাইটিং, ডিজিটাল ইনস্ট্রুমেন্টেশন এবং নির্ভরযোগ্য সুজুকি ইঞ্জিনিয়ারিং সহ কঠিন কোয়ার্টার-লিটার নেকেড বাইক পারফরমেন্স প্রদান করে। ২৪৯সিসি ইঞ্জিন শোভন শক্তি প্রদান করে যখন নেকেড ডিজাইন শহুরে অবস্থায় সহজ চালচলন নিশ্চিত করে। তবে, এবিএস ছাড়া উচ্চ মূল্য, একক সিলিন্ডার কম্পন, ব্যয়বহুল রক্ষণাবেক্ষণ, হাইওয়ে ব্যবহারের জন্য বাতাস সুরক্ষার অভাব এবং প্রতিযোগীদের তুলনায় খারাপ মূল্য প্রস্তাব এটিকে কেবলমাত্র সুজুকি অনুগতদের জন্য উপযুক্ত করে তোলে যারা ব্যবহারিক মূল্য বিবেচনার চেয়ে ব্র্যান্ড ঐতিহ্য এবং নেকেড বাইক স্টাইলিংকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শহুরে কোয়ার্টার-লিটার নেকেড বাইক উৎসাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer 250 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer 250\n")

	return nil
}
