package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxerFiDiscReview handles seeding Suzuki Gixxer FI Disc product review and translation
type SuzukiGixxerFiDiscReview struct {
	BaseSeeder
}

// NewSuzukiGixxerFiDiscReviewSeeder creates a new SuzukiGixxerFiDiscReview
func NewSuzukiGixxerFiDiscReviewSeeder() *SuzukiGixxerFiDiscReview {
	return &SuzukiGixxerFiDiscReview{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer FI Disc Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxerFiDiscReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer FI Disc").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Gixxer FI Disc product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Gixxer FI Disc product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer FI Disc already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer FI Disc Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Gixxer FI Disc is an entry-level 155cc naked motorcycle featuring fuel injection technology and front disc brake system. Priced at ৳2,49,950, it offers modern FI efficiency, disc brake safety, aggressive naked styling, and Suzuki reliability in an affordable package for daily commuting and urban riding.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer FI Disc Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Injection Technology:</strong> <span class="highlight-value">Modern FI system for better fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Front Disc Brake:</strong> <span class="highlight-value">Enhanced braking performance and safety</span></li>
<li class="highlight-item"><strong class="highlight-label">Naked Streetfighter Design:</strong> <span class="highlight-value">Sharp and aggressive styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Affordable Entry Price:</strong> <span class="highlight-value">Budget-friendly FI motorcycle</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer FI Disc Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Competitive Pricing:</strong> <span class="pro-description">Most affordable Suzuki with fuel injection technology</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Injection Benefits:</strong> <span class="pro-description">Better fuel economy, instant starts, and consistent performance</span></li>
<li class="pro-item"><strong class="pro-label">Disc Brake Safety:</strong> <span class="pro-description">Superior stopping power compared to drum brakes</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight Handling:</strong> <span class="pro-description">Easy to maneuver in city traffic and parking</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">Achieves 45-50 km/l with FI technology</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Reliability:</strong> <span class="pro-description">Proven Japanese engineering quality</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer FI Disc Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">No ABS Feature:</strong> <span class="con-description">Missing anti-lock braking system for enhanced safety</span></li>
<li class="con-item"><strong class="con-label">Rear Drum Brake:</strong> <span class="con-description">Basic drum brake setup at rear instead of disc</span></li>
<li class="con-item"><strong class="con-label">Limited Power Output:</strong> <span class="con-description">155cc engine may feel underpowered for highway use</span></li>
<li class="con-item"><strong class="con-label">No Wind Protection:</strong> <span class="con-description">Naked design offers no protection from wind</span></li>
<li class="con-item"><strong class="con-label">Maintenance Cost:</strong> <span class="con-description">Fuel injection requires specialized service</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Gixxer FI Disc in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget-conscious buyers seeking FI technology</li>
<li class="best-for-item">Daily urban commuters</li>
<li class="best-for-item">First-time motorcycle buyers</li>
<li class="best-for-item">Riders upgrading from smaller displacement bikes</li>
<li class="best-for-item">Those prioritizing fuel efficiency</li>
<li class="best-for-item">Young riders seeking modern features</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Gixxer FI Disc?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway touring enthusiasts</li>
<li class="not-recommended-item">Riders requiring ABS safety</li>
<li class="not-recommended-item">Power-hungry performance seekers</li>
<li class="not-recommended-item">Those needing wind protection</li>
<li class="not-recommended-item">Riders preferring drum brake simplicity</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer FI Disc Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,49,950 - Excellent value for FI technology</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳6-8 per km with excellent fuel economy</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳6-8 per km (45-50 km/l with FI efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳4,000-6,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer FI Disc Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Adequate for city use</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Good for short rides</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Attractive naked design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Excellent for price</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good Suzuki build</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Gixxer FI Disc Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does this compare to carburetor variants?</h3>
<p class="faq-answer">The FI variant offers better fuel efficiency (45-50 km/l vs 40-45 km/l), instant starts in all weather, smoother performance, and reduced emissions compared to carburetor models.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the lack of ABS a major concern?</h3>
<p class="faq-answer">While ABS would enhance safety, the front disc brake still provides good stopping power. Careful riding and proper braking technique can compensate for the lack of ABS in this price segment.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the expected maintenance cost?</h3>
<p class="faq-answer">Annual maintenance cost ranges from ৳4,000-6,000 including regular services, oil changes, and minor repairs. FI system may require occasional specialist attention but is generally reliable.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer FI Disc in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.9/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,49,950, the Suzuki Gixxer FI Disc offers exceptional value as the most affordable Suzuki motorcycle with fuel injection technology. The FI system ensures excellent fuel economy of 45-50 km/l, instant starts, and consistent performance, while the front disc brake provides adequate safety. However, the lack of ABS, rear drum brake, limited power for highway use, no wind protection, and specialized FI maintenance requirements make it suitable primarily for budget-conscious urban commuters and first-time buyers who prioritize fuel efficiency and modern features over advanced safety and performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget-friendly FI efficiency with disc brake safety</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.9,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Suzuki Gixxer FI Disc review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer FI Disc (Rating: 3.9/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার এফআই ডিস্ক রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি গিক্সার এফআই ডিস্ক হল একটি এন্ট্রি-লেভেল ১৫৫সিসি নেকেড মোটরসাইকেল যেখানে ফুয়েল ইনজেকশন প্রযুক্তি এবং সামনের ডিস্ক ব্রেক সিস্টেম রয়েছে। ৳২,৪৯,৯৫০ টাকায় মূল্যায়িত, এটি দৈনিক যাতায়াত এবং শহুরে রাইডিংয়ের জন্য একটি সাশ্রয়ী প্যাকেজে আধুনিক এফআই দক্ষতা, ডিস্ক ব্রেক নিরাপত্তা, আক্রমণাত্মক নেকেড স্টাইলিং এবং সুজুকি নির্ভরযোগ্যতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার এফআই ডিস্ক এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন প্রযুক্তি:</strong> <span class="highlight-value">ভাল জ্বালানি দক্ষতার জন্য আধুনিক এফআই সিস্টেম</span></li>
<li class="highlight-item"><strong class="highlight-label">সামনের ডিস্ক ব্রেক:</strong> <span class="highlight-value">উন্নত ব্রেকিং পারফরমেন্স এবং নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">নেকেড স্ট্রিটফাইটার ডিজাইন:</strong> <span class="highlight-value">তীক্ষ্ণ এবং আক্রমণাত্মক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">সাশ্রয়ী প্রবেশ মূল্য:</strong> <span class="highlight-value">বাজেট-বান্ধব এফআই মোটরসাইকেল</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার এফআই ডিস্ক এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রতিযোগিতামূলক মূল্য:</strong> <span class="pro-description">ফুয়েল ইনজেকশন প্রযুক্তি সহ সবচেয়ে সাশ্রয়ী সুজুকি</span></li>
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন সুবিধা:</strong> <span class="pro-description">ভাল জ্বালানি অর্থনীতি, তাৎক্ষণিক স্টার্ট এবং সামঞ্জস্যপূর্ণ পারফরমেন্স</span></li>
<li class="pro-item"><strong class="pro-label">ডিস্ক ব্রেক নিরাপত্তা:</strong> <span class="pro-description">ড্রাম ব্রেকের তুলনায় উন্নত স্টপিং পাওয়ার</span></li>
<li class="pro-item"><strong class="pro-label">হালকা হ্যান্ডলিং:</strong> <span class="pro-description">শহুরে ট্রাফিক এবং পার্কিংয়ে চালনা সহজ</span></li>
<li class="pro-item"><strong class="pro-label">ভাল জ্বালানি সাশ্রয়:</strong> <span class="pro-description">এফআই প্রযুক্তি দিয়ে ৪৫-৫০ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত জাপানি ইঞ্জিনিয়ারিং গুণমান</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার এফআই ডিস্ক এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">কোনো এবিএস ফিচার নেই:</strong> <span class="con-description">উন্নত নিরাপত্তার জন্য অ্যান্টি-লক ব্রেকিং সিস্টেমের অভাব</span></li>
<li class="con-item"><strong class="con-label">পিছনের ড্রাম ব্রেক:</strong> <span class="con-description">ডিস্কের পরিবর্তে পিছনে বেসিক ড্রাম ব্রেক সেটাপ</span></li>
<li class="con-item"><strong class="con-label">সীমিত শক্তি আউটপুট:</strong> <span class="con-description">হাইওয়ে ব্যবহারের জন্য ১৫৫সিসি ইঞ্জিন কম শক্তিশালী মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">কোন বায়ু সুরক্ষা নেই:</strong> <span class="con-description">নেকেড ডিজাইন বায়ু থেকে কোন সুরক্ষা প্রদান করে না</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">ফুয়েল ইনজেকশনের বিশেষায়িত সেবা প্রয়োজন</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি গিক্সার এফআই ডিস্ক কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget-conscious buyers seeking FI technology</li>
<li class="best-for-item">Daily urban commuters</li>
<li class="best-for-item">First-time motorcycle buyers</li>
<li class="best-for-item">Riders upgrading from smaller displacement bikes</li>
<li class="best-for-item">Those prioritizing fuel efficiency</li>
<li class="best-for-item">Young riders seeking modern features</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি গিক্সার এফআই ডিস্ক কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway touring enthusiasts</li>
<li class="not-recommended-item">Riders requiring ABS safety</li>
<li class="not-recommended-item">Power-hungry performance seekers</li>
<li class="not-recommended-item">Those needing wind protection</li>
<li class="not-recommended-item">Riders preferring drum brake simplicity</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি গিক্সার এফআই ডিস্ক এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৪৯,৯৫০ - এফআই প্রযুক্তির জন্য চমৎকার ভ্যালু</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম - চমৎকার জ্বালানি সাশ্রয় সহ ৳৬-৮ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৬-৮ প্রতি কিমি (এফআই দক্ষতা সহ ৪৫-৫০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৪,০০০-৬,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি গিক্সার এফআই ডিস্ক পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- শহুরে ব্যবহারের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- ছোট রাইডের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- আকর্ষণীয় নেকেড ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- দামের জন্য চমৎকার</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল সুজুকি বিল্ড</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি গিক্সার এফআই ডিস্ক সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এটি কার্বুরেটর ভেরিয়েন্টের সাথে কেমন তুলনা করে?</h3>
<p class="faq-answer">এফআই ভেরিয়েন্ট কার্বুরেটর মডেলের তুলনায় ভাল জ্বালানি দক্ষতা (৪৫-৫০ কিমি/লিটার বনাম ৪০-৪৫ কিমি/লিটার), সব আবহাওয়ায় তাৎক্ষণিক স্টার্ট, মসৃণ পারফরমেন্স এবং কম নির্গমন প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এবিএসের অভাব কি একটি বড় উদ্বেগ?</h3>
<p class="faq-answer">এবিএস নিরাপত্তা বাড়াতে পারলেও, সামনের ডিস্ক ব্রেক এখনও ভাল স্টপিং পাওয়ার প্রদান করে। এই দামের সেগমেন্টে সতর্ক রাইডিং এবং সঠিক ব্রেকিং কৌশল এবিএসের অভাব পূরণ করতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">প্রত্যাশিত রক্ষণাবেক্ষণ খরচ কত?</h3>
<p class="faq-answer">নিয়মিত সেবা, তেল পরিবর্তন এবং ছোটখাট মেরামত সহ বার্ষিক রক্ষণাবেক্ষণ খরচ ৳৪,০০০-৬,০০০ এর মধ্যে। এফআই সিস্টেমের মাঝে মাঝে বিশেষজ্ঞ মনোযোগের প্রয়োজন হতে পারে কিন্তু সাধারণত নির্ভরযোগ্য।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি গিক্সার এফআই ডিস্ক কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.9/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,৪৯,৯৫০ টাকায় সুজুকি গিক্সার এফআই ডিস্ক ফুয়েল ইনজেকশন প্রযুক্তি সহ সবচেয়ে সাশ্রয়ী সুজুকি মোটরসাইকেল হিসেবে ব্যতিক্রমী ভ্যালু প্রদান করে। এফআই সিস্টেম ৪৫-৫০ কিমি/লিটার চমৎকার জ্বালানি অর্থনীতি, তাৎক্ষণিক স্টার্ট এবং সামঞ্জস্যপূর্ণ পারফরমেন্স নিশ্চিত করে, যখন সামনের ডিস্ক ব্রেক পর্যাপ্ত নিরাপত্তা প্রদান করে। তবে, এবিএসের অভাব, পিছনের ড্রাম ব্রেক, হাইওয়ে ব্যবহারের জন্য সীমিত শক্তি, বায়ু সুরক্ষার অভাব এবং বিশেষায়িত এফআই রক্ষণাবেক্ষণের প্রয়োজনীয়তা এটিকে প্রধানত বাজেট-সচেতন শহুরে যাত্রী এবং প্রথমবার ক্রেতাদের জন্য উপযুক্ত করে তোলে যারা উন্নত নিরাপত্তা এবং পারফরমেন্সের চেয়ে জ্বালানি দক্ষতা এবং আধুনিক বৈশিষ্ট্যকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ডিস্ক ব্রেক নিরাপত্তা সহ বাজেট-বান্ধব এফআই দক্ষতার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer FI Disc already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer FI Disc\n")

	return nil
}
