package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaSp125Review handles seeding Honda SP125 product review and translation
type HondaSp125Review struct {
	BaseSeeder
}

// NewHondaSp125ReviewSeeder creates a new HondaSp125Review
func NewHondaSp125ReviewSeeder() *HondaSp125Review {
	return &HondaSp125Review{BaseSeeder: BaseSeeder{name: "Honda SP125 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaSp125Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda SP125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda SP125 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda SP125 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda SP125 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda SP125 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda SP125 is a sporty 125cc commuter motorcycle featuring youthful design, fuel injection technology, LED headlight with DRL, digital-analog instrument cluster, and practical features. Priced at ৳1,65,000, it offers excellent value with sporty styling, superior fuel efficiency, modern LED lighting, Honda reliability, and affordable running costs making it ideal for young commuters seeking sporty aesthetics with economical operation in the 125cc segment.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda SP125 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">LED Headlight with DRL:</strong> <span class="highlight-value">Modern LED lighting with daytime running lights</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Injection:</strong> <span class="highlight-value">PGM-Fi for better efficiency and performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Sporty Styling:</strong> <span class="highlight-value">Youthful design with sporty graphics</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital-Analog Display:</strong> <span class="highlight-value">Modern instrument cluster with comprehensive info</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda SP125 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">LED Lighting with DRL:</strong> <span class="pro-description">Modern LED headlight with DRL enhances visibility and style</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Fuel Efficiency:</strong> <span class="pro-description">Achieves 55-60 km/l with FI technology</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Appeal:</strong> <span class="pro-description">Youthful styling attracts young urban riders</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Pricing:</strong> <span class="pro-description">Great features at competitive ৳1,65,000 price point</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven Honda quality and dependability</span></li>
<li class="pro-item"><strong class="pro-label">Low Running Costs:</strong> <span class="pro-description">125cc ensures economical daily operation</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda SP125 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Power:</strong> <span class="con-description">125cc provides basic performance for commuting only</span></li>
<li class="con-item"><strong class="con-label">Drum Rear Brake:</strong> <span class="con-description">Rear drum brake instead of disc</span></li>
<li class="con-item"><strong class="con-label">No ABS:</strong> <span class="con-description">Missing anti-lock braking safety feature</span></li>
<li class="con-item"><strong class="con-label">Basic Suspension:</strong> <span class="con-description">Simple setup can feel harsh on rough roads</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Fit and finish adequate but not premium</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda SP125 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young daily commuters</li>
<li class="best-for-item">First-time motorcycle buyers</li>
<li class="best-for-item">Budget-conscious students</li>
<li class="best-for-item">City and short highway rides</li>
<li class="best-for-item">Those prioritizing fuel economy</li>
<li class="best-for-item">Economical daily transportation</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda SP125?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Those requiring more power</li>
<li class="not-recommended-item">ABS safety requirement</li>
<li class="not-recommended-item">Premium feature seekers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda SP125 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,65,000 - Excellent for sporty 125cc commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very Low - ৳4-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳4-5 per km (55-60 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳4,000-6,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda SP125 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Basic 125cc performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Outstanding fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good commuter comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Sporty youthful design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent value for money</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good Honda quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda SP125 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Is 125cc enough for daily commuting?</h3>
<p class="faq-answer">Yes, the SP125's 125cc is perfectly adequate for city commuting and short highway trips. It provides good balance of performance and exceptional fuel economy of 55-60 km/l, making it ideal for economical daily use.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it compare to Honda Shine?</h3>
<p class="faq-answer">The SP125 offers more sporty styling, LED headlight with DRL, and modern features compared to the conservative Shine. Both share similar reliability and fuel economy, but SP125 appeals more to younger riders seeking style.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What's the real-world mileage?</h3>
<p class="faq-answer">Real-world fuel economy ranges from 55-60 km/l with careful riding, dropping to 50-52 km/l in heavy traffic. The FI system ensures consistent efficiency making it one of the most economical 125cc bikes.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda SP125 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,65,000, the Honda SP125 offers excellent value as a sporty 125cc commuter with LED headlight and DRL, fuel injection efficiency (55-60 km/l), youthful styling, and Honda reliability. It excels at economical daily commuting with outstanding fuel economy and low running costs. However, limited 125cc power, drum rear brake, no ABS, basic suspension, and average build quality make it best suited for young daily commuters, first-time buyers, students, and budget-conscious riders who prioritize fuel economy, affordability, and sporty styling over performance and premium features.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For outstanding fuel economy with sporty styling in affordable 125cc</span></p>
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
		return fmt.Errorf("error creating Honda SP125 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda SP125 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা এসপি১২৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা এসপি১২৫ হল একটি স্পোর্টি ১২৫সিসি কম্যুটার মোটরসাইকেল যেখানে তরুণ ডিজাইন, ফুয়েল ইনজেকশন প্রযুক্তি, ডিআরএল সহ এলইডি হেডলাইট, ডিজিটাল-অ্যানালগ ইনস্ট্রুমেন্ট ক্লাস্টার এবং ব্যবহারিক বৈশিষ্ট্য রয়েছে। ৳১,৬৫,০০০ টাকায় মূল্যায়িত, এটি স্পোর্টি স্টাইলিং, উন্নত জ্বালানি দক্ষতা, আধুনিক এলইডি লাইটিং, হোন্ডা নির্ভরযোগ্যতা এবং ১২৫সিসি সেগমেন্টে সাশ্রয়ী অপারেশন সহ স্পোর্টি নান্দনিকতা খোঁজা তরুণ যাত্রীদের জন্য আদর্শ সাশ্রয়ী চলমান খরচ সহ চমৎকার মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা এসপি১২৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ডিআরএল সহ এলইডি হেডলাইট:</strong> <span class="highlight-value">দিনের বেলা চলমান আলো সহ আধুনিক এলইডি লাইটিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন:</strong> <span class="highlight-value">ভাল দক্ষতা এবং পারফরমেন্সের জন্য পিজিএম-এফআই</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্টি স্টাইলিং:</strong> <span class="highlight-value">স্পোর্টি গ্রাফিক্স সহ তরুণ ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল-অ্যানালগ ডিসপ্লে:</strong> <span class="highlight-value">ব্যাপক তথ্য সহ আধুনিক ইনস্ট্রুমেন্ট ক্লাস্টার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা এসপি১২৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ডিআরএল সহ এলইডি লাইটিং:</strong> <span class="pro-description">ডিআরএল সহ আধুনিক এলইডি হেডলাইট দৃশ্যমানতা এবং স্টাইল বাড়ায়</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার জ্বালানি দক্ষতা:</strong> <span class="pro-description">এফআই প্রযুক্তি দিয়ে ৫৫-৬০ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি আবেদন:</strong> <span class="pro-description">তরুণ স্টাইলিং তরুণ শহুরে রাইডারদের আকর্ষণ করে</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">প্রতিযোগিতামূলক ৳১,৬৫,০০০ মূল্য পয়েন্টে দুর্দান্ত বৈশিষ্ট্য</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত হোন্ডা গুণমান এবং নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">কম চলমান খরচ:</strong> <span class="pro-description">১২৫সিসি সাশ্রয়ী দৈনিক অপারেশন নিশ্চিত করে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা এসপি১২৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত শক্তি:</strong> <span class="con-description">১২৫সিসি শুধুমাত্র যাতায়াতের জন্য মৌলিক পারফরমেন্স প্রদান করে</span></li>
<li class="con-item"><strong class="con-label">ড্রাম পিছনের ব্রেক:</strong> <span class="con-description">ডিস্কের পরিবর্তে পিছনের ড্রাম ব্রেক</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস নেই:</strong> <span class="con-description">অ্যান্টি-লক ব্রেকিং নিরাপত্তা বৈশিষ্ট্য অনুপস্থিত</span></li>
<li class="con-item"><strong class="con-label">মৌলিক সাসপেনশন:</strong> <span class="con-description">সাধারণ সেটআপ রুক্ষ রাস্তায় কঠোর মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">ফিট এবং ফিনিশ পর্যাপ্ত কিন্তু প্রিমিয়াম নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা এসপি১২৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young daily commuters</li>
<li class="best-for-item">First-time motorcycle buyers</li>
<li class="best-for-item">Budget-conscious students</li>
<li class="best-for-item">City and short highway rides</li>
<li class="best-for-item">Those prioritizing fuel economy</li>
<li class="best-for-item">Economical daily transportation</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা এসপি১২৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Those requiring more power</li>
<li class="not-recommended-item">ABS safety requirement</li>
<li class="not-recommended-item">Premium feature seekers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা এসপি১২৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৬৫,০০০ - স্পোর্টি ১২৫সিসি কম্যুটারের জন্য চমৎকার</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">অত্যন্ত কম - ৳৪-৬ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৪-৫ প্রতি কিমি (৫৫-৬০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৪,০০০-৬,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা এসপি১২৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- মৌলিক ১২৫সিসি পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- অসাধারণ জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল কম্যুটার আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- স্পোর্টি তরুণ ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার ভ্যালু ফর মানি</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ভাল হোন্ডা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা এসপি১২৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">দৈনিক যাতায়াতের জন্য ১২৫সিসি কি যথেষ্ট?</h3>
<p class="faq-answer">হ্যাঁ, এসপি১২৫ এর ১২৫সিসি শহুরে যাতায়াত এবং সংক্ষিপ্ত হাইওয়ে ভ্রমণের জন্য পুরোপুরি পর্যাপ্ত। এটি পারফরমেন্স এবং ৫৫-৬০ কিমি/লিটার ব্যতিক্রমী জ্বালানি অর্থনীতির ভাল ভারসাম্য প্রদান করে, এটি সাশ্রয়ী দৈনিক ব্যবহারের জন্য আদর্শ করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি হোন্ডা শাইনের সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">এসপি১২৫ রক্ষণশীল শাইনের তুলনায় আরও স্পোর্টি স্টাইলিং, ডিআরএল সহ এলইডি হেডলাইট এবং আধুনিক বৈশিষ্ট্য প্রদান করে। উভয়ই অনুরূপ নির্ভরযোগ্যতা এবং জ্বালানি অর্থনীতি শেয়ার করে, কিন্তু এসপি১২৫ স্টাইল খোঁজা তরুণ রাইডারদের কাছে বেশি আবেদন করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাস্তব-বিশ্ব মাইলেজ কত?</h3>
<p class="faq-answer">বাস্তব-বিশ্ব জ্বালানি অর্থনীতি সাবধানে রাইডিং সহ ৫৫-৬০ কিমি/লিটার পরিসীমা, ভারী ট্রাফিকে ৫০-৫২ কিমি/লিটারে নেমে যায়। এফআই সিস্টেম সামঞ্জস্যপূর্ণ দক্ষতা নিশ্চিত করে এটিকে সবচেয়ে সাশ্রয়ী ১২৫সিসি বাইকগুলির মধ্যে একটি করে তোলে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা এসপি১২৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳১,৬৫,০০০ টাকায় হোন্ডা এসপি১২৫ ডিআরএল সহ এলইডি হেডলাইট, ফুয়েল ইনজেকশন দক্ষতা (৫৫-৬০ কিমি/লিটার), তরুণ স্টাইলিং এবং হোন্ডা নির্ভরযোগ্যতা সহ একটি স্পোর্টি ১২৫সিসি কম্যুটার হিসাবে চমৎকার মূল্য প্রদান করে। এটি অসাধারণ জ্বালানি অর্থনীতি এবং কম চলমান খরচ সহ সাশ্রয়ী দৈনিক যাতায়াতে শ্রেষ্ঠ। তবে, সীমিত ১২৫সিসি শক্তি, ড্রাম পিছনের ব্রেক, এবিএস নেই, মৌলিক সাসপেনশন এবং গড় বিল্ড কোয়ালিটি এটিকে তরুণ দৈনিক যাত্রী, প্রথমবার ক্রেতা, ছাত্র এবং পারফরমেন্স এবং প্রিমিয়াম বৈশিষ্ট্যের চেয়ে জ্বালানি অর্থনীতি, সাশ্রয়ীতা এবং স্পোর্টি স্টাইলিংকে অগ্রাধিকার দেওয়া বাজেট-সচেতন রাইডারদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সাশ্রয়ী ১২৫সিসিতে স্পোর্টি স্টাইলিং সহ অসাধারণ জ্বালানি অর্থনীতির জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda SP125 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda SP125\n")

	return nil
}
