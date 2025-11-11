package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// NewHondaXblade160Review handles seeding New Honda XBlade 160 2024 product review and translation
type NewHondaXblade160Review struct {
	BaseSeeder
}

// NewNewHondaXblade160ReviewSeeder creates a new NewHondaXblade160Review
func NewNewHondaXblade160ReviewSeeder() *NewHondaXblade160Review {
	return &NewHondaXblade160Review{BaseSeeder: BaseSeeder{name: "New Honda XBlade 160 2024 Review"}}
}

// Seed implements the Seeder interface
func (s *NewHondaXblade160Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "New Honda XBlade 160 2024").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("New Honda XBlade 160 2024 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding New Honda XBlade 160 2024 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for New Honda XBlade 160 2024 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">New Honda XBlade 160 2024 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The New Honda XBlade 160 2024 is an updated aggressive street fighter motorcycle featuring sharp X-inspired styling, 160cc fuel-injected engine, LED headlight, digital instrumentation, and sporty ergonomics. Priced at ৳1,95,000, it offers excellent value with modern styling, good fuel efficiency, sharp handling, Honda reliability, and practical features making it ideal for young urban riders seeking aggressive street presence in an affordable, feature-rich package.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">New Honda XBlade 160 2024 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">X-Inspired Design:</strong> <span class="highlight-value">Sharp aggressive street fighter styling</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Headlight:</strong> <span class="highlight-value">Modern LED headlight for better visibility</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Display:</strong> <span class="highlight-value">Fully digital instrument cluster</span></li>
<li class="highlight-item"><strong class="highlight-label">160cc PGM-Fi:</strong> <span class="highlight-value">Fuel injection for efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">New Honda XBlade 160 2024 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Sharp X-inspired design stands out in traffic</span></li>
<li class="pro-item"><strong class="pro-label">LED Headlight:</strong> <span class="pro-description">Modern LED provides excellent night visibility</span></li>
<li class="pro-item"><strong class="pro-label">Digital Instrumentation:</strong> <span class="pro-description">Fully digital display with comprehensive information</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">Achieves 45-50 km/l with FI technology</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Nimble chassis provides good cornering dynamics</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Pricing:</strong> <span class="pro-description">Great features at competitive ৳1,95,000 price</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">New Honda XBlade 160 2024 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Moderate Performance:</strong> <span class="con-description">160cc provides adequate but not thrilling power</span></li>
<li class="con-item"><strong class="con-label">Single Disc Brake:</strong> <span class="con-description">Front disc only, rear drum brake</span></li>
<li class="con-item"><strong class="con-label">Firm Suspension:</strong> <span class="con-description">Sporty setup can feel harsh on rough roads</span></li>
<li class="con-item"><strong class="con-label">No ABS:</strong> <span class="con-description">Missing anti-lock braking safety feature</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Fit and finish not as premium as higher models</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy New Honda XBlade 160 2024 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban commuters</li>
<li class="best-for-item">Style-conscious riders on budget</li>
<li class="best-for-item">First-time 160cc buyers</li>
<li class="best-for-item">Daily city riding</li>
<li class="best-for-item">College and office commuters</li>
<li class="best-for-item">Those seeking aggressive street styling</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy New Honda XBlade 160 2024?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Those requiring ABS safety</li>
<li class="not-recommended-item">Dual disc brake seekers</li>
<li class="not-recommended-item">Soft suspension comfort needs</li>
<li class="not-recommended-item">Long-distance touring</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">New Honda XBlade 160 2024 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,95,000 - Excellent value for 160cc street fighter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳6-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳6-7 per km (45-50 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳5,000-7,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">New Honda XBlade 160 2024 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good daily performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Decent commuter comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Sharp aggressive design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent value for money</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good Honda quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">New Honda XBlade 160 2024 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What's new in the 2024 version?</h3>
<p class="faq-answer">The 2024 version features updated graphics, refined styling elements, improved ergonomics, and tweaked engine mapping for slightly better response while maintaining the core X-inspired aggressive design.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Should I get single or double disc variant?</h3>
<p class="faq-answer">If budget allows, the double disc variant (৳30,000 more) offers significantly better braking safety. For city use with moderate speeds, the single disc version at ৳1,95,000 provides good value.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it compare to TVS Apache RTR 160?</h3>
<p class="faq-answer">The XBlade offers better fuel economy and Honda reliability, while Apache RTR 160 provides more sporty performance and features. XBlade is better for practical daily use, Apache for sporty enthusiasts.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy New Honda XBlade 160 2024 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,95,000, the New Honda XBlade 160 2024 offers excellent value as an aggressive street fighter with sharp X-inspired styling, LED headlight, digital display, good fuel economy (45-50 km/l), and Honda reliability. It excels at urban commuting with modern features and eye-catching design. However, moderate performance, single disc brake, firm suspension, lack of ABS, and average build quality make it best suited for young urban riders, style-conscious buyers on budget, and daily commuters who prioritize aggressive street presence and fuel efficiency in an affordable package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For aggressive street fighter styling with modern features at affordable price</span></p>
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
		return fmt.Errorf("error creating New Honda XBlade 160 2024 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for New Honda XBlade 160 2024 (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ হল একটি আপডেটেড আক্রমণাত্মক স্ট্রিট ফাইটার মোটরসাইকেল যেখানে তীক্ষ্ণ এক্স-অনুপ্রাণিত স্টাইলিং, ১৬০সিসি ফুয়েল-ইনজেক্টেড ইঞ্জিন, এলইডি হেডলাইট, ডিজিটাল ইনস্ট্রুমেন্টেশন এবং স্পোর্টি এরগোনমিক্স রয়েছে। ৳১,৯৫,০০০ টাকায় মূল্যায়িত, এটি আধুনিক স্টাইলিং, ভাল জ্বালানি দক্ষতা, তীক্ষ্ণ হ্যান্ডলিং, হোন্ডা নির্ভরযোগ্যতা এবং সাশ্রয়ী, ফিচার-সমৃদ্ধ প্যাকেজে আক্রমণাত্মক স্ট্রিট উপস্থিতি খোঁজা তরুণ শহুরে রাইডারদের জন্য আদর্শ ব্যবহারিক বৈশিষ্ট্য সহ চমৎকার মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">এক্স-অনুপ্রাণিত ডিজাইন:</strong> <span class="highlight-value">তীক্ষ্ণ আক্রমণাত্মক স্ট্রিট ফাইটার স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি হেডলাইট:</strong> <span class="highlight-value">ভাল দৃশ্যমানতার জন্য আধুনিক এলইডি হেডলাইট</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল ডিসপ্লে:</strong> <span class="highlight-value">সম্পূর্ণ ডিজিটাল ইনস্ট্রুমেন্ট ক্লাস্টার</span></li>
<li class="highlight-item"><strong class="highlight-label">১৬০সিসি পিজিএম-এফআই:</strong> <span class="highlight-value">দক্ষতার জন্য ফুয়েল ইনজেকশন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">তীক্ষ্ণ এক্স-অনুপ্রাণিত ডিজাইন ট্রাফিকে আলাদা হয়ে দাঁড়ায়</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি হেডলাইট:</strong> <span class="pro-description">আধুনিক এলইডি চমৎকার রাতের দৃশ্যমানতা প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল ইনস্ট্রুমেন্টেশন:</strong> <span class="pro-description">ব্যাপক তথ্য সহ সম্পূর্ণ ডিজিটাল ডিসপ্লে</span></li>
<li class="pro-item"><strong class="pro-label">ভাল জ্বালানি সাশ্রয়:</strong> <span class="pro-description">এফআই প্রযুক্তি দিয়ে ৪৫-৫০ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="pro-description">চটপটে চ্যাসি ভাল কর্নারিং ডায়নামিক্স প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">প্রতিযোগিতামূলক ৳১,৯৫,০০০ মূল্যে দুর্দান্ত বৈশিষ্ট্য</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মাঝারি পারফরমেন্স:</strong> <span class="con-description">১৬০সিসি পর্যাপ্ত কিন্তু রোমাঞ্চকর নয় এমন শক্তি প্রদান করে</span></li>
<li class="con-item"><strong class="con-label">একক ডিস্ক ব্রেক:</strong> <span class="con-description">শুধুমাত্র সামনের ডিস্ক, পিছনের ড্রাম ব্রেক</span></li>
<li class="con-item"><strong class="con-label">শক্ত সাসপেনশন:</strong> <span class="con-description">স্পোর্টি সেটআপ রুক্ষ রাস্তায় কঠোর মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস নেই:</strong> <span class="con-description">অ্যান্টি-লক ব্রেকিং নিরাপত্তা বৈশিষ্ট্য অনুপস্থিত</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">ফিট এবং ফিনিশ উচ্চতর মডেলের মতো প্রিমিয়াম নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban commuters</li>
<li class="best-for-item">Style-conscious riders on budget</li>
<li class="best-for-item">First-time 160cc buyers</li>
<li class="best-for-item">Daily city riding</li>
<li class="best-for-item">College and office commuters</li>
<li class="best-for-item">Those seeking aggressive street styling</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Those requiring ABS safety</li>
<li class="not-recommended-item">Dual disc brake seekers</li>
<li class="not-recommended-item">Soft suspension comfort needs</li>
<li class="not-recommended-item">Long-distance touring</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৯৫,০০০ - ১৬০সিসি স্ট্রিট ফাইটারের জন্য চমৎকার মূল্য</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম - ৳৬-৮ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৬-৭ প্রতি কিমি (৪৫-৫০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৫,০০০-৭,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল দৈনিক পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- ভাল জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- ভাল কম্যুটার আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- তীক্ষ্ণ আক্রমণাত্মক ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার ভ্যালু ফর মানি</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল হোন্ডা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">২০২৪ সংস্করণে নতুন কী?</h3>
<p class="faq-answer">২০২৪ সংস্করণে আপডেটেড গ্রাফিক্স, পরিমার্জিত স্টাইলিং উপাদান, উন্নত এরগোনমিক্স এবং মূল এক্স-অনুপ্রাণিত আক্রমণাত্মক ডিজাইন বজায় রেখে সামান্য ভাল প্রতিক্রিয়ার জন্য টুইক করা ইঞ্জিন ম্যাপিং রয়েছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">আমি কি একক না ডাবল ডিস্ক ভ্যারিয়েন্ট পাব?</h3>
<p class="faq-answer">বাজেট অনুমতি দিলে, ডাবল ডিস্ক ভ্যারিয়েন্ট (৳৩০,০০০ বেশি) উল্লেখযোগ্যভাবে ভাল ব্রেকিং নিরাপত্তা প্রদান করে। মাঝারি গতির সাথে শহুরে ব্যবহারের জন্য, ৳১,৯৫,০০০ এ একক ডিস্ক সংস্করণ ভাল মূল্য প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি টিভিএস অ্যাপাচি আরটিআর ১৬০ এর সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">এক্সব্লেড ভাল জ্বালানি অর্থনীতি এবং হোন্ডা নির্ভরযোগ্যতা প্রদান করে, যেখানে অ্যাপাচি আরটিআর ১৬০ আরও স্পোর্টি পারফরমেন্স এবং বৈশিষ্ট্য প্রদান করে। এক্সব্লেড ব্যবহারিক দৈনিক ব্যবহারের জন্য ভাল, অ্যাপাচি স্পোর্টি উৎসাহীদের জন্য।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳১,৯৫,০০০ টাকায় নিউ হোন্ডা এক্সব্লেড ১৬০ ২০২৪ তীক্ষ্ণ এক্স-অনুপ্রাণিত স্টাইলিং, এলইডি হেডলাইট, ডিজিটাল ডিসপ্লে, ভাল জ্বালানি অর্থনীতি (৪৫-৫০ কিমি/লিটার) এবং হোন্ডা নির্ভরযোগ্যতা সহ একটি আক্রমণাত্মক স্ট্রিট ফাইটার হিসাবে চমৎকার মূল্য প্রদান করে। এটি আধুনিক বৈশিষ্ট্য এবং চোখ আকর্ষণীয় ডিজাইন সহ শহুরে যাতায়াতে শ্রেষ্ঠ। তবে, মাঝারি পারফরমেন্স, একক ডিস্ক ব্রেক, শক্ত সাসপেনশন, এবিএসের অভাব এবং গড় বিল্ড কোয়ালিটি এটিকে তরুণ শহুরে রাইডার, বাজেটে স্টাইল-সচেতন ক্রেতা এবং সাশ্রয়ী প্যাকেজে আক্রমণাত্মক স্ট্রিট উপস্থিতি এবং জ্বালানি দক্ষতাকে অগ্রাধিকার দেওয়া দৈনিক যাত্রীদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সাশ্রয়ী মূল্যে আধুনিক বৈশিষ্ট্য সহ আক্রমণাত্মক স্ট্রিট ফাইটার স্টাইলিংয়ের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for New Honda XBlade 160 2024 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for New Honda XBlade 160 2024\n")

	return nil
}
