package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajAvenger160AbsReview handles seeding Bajaj Avenger 160 ABS product review and translation
type BajajAvenger160AbsReview struct {
	BaseSeeder
}

// NewBajajAvenger160AbsReviewSeeder creates a new BajajAvenger160AbsReview
func NewBajajAvenger160AbsReviewSeeder() *BajajAvenger160AbsReview {
	return &BajajAvenger160AbsReview{BaseSeeder: BaseSeeder{name: "Bajaj Avenger 160 ABS Review"}}
}

// Seed implements the Seeder interface
func (s *BajajAvenger160AbsReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Avenger 160 ABS").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Avenger 160 ABS product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Avenger 160 ABS product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Avenger 160 ABS already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Avenger 160 ABS Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Avenger 160 ABS is India's most affordable cruiser motorcycle featuring relaxed cruising ergonomics, 160cc fuel-injected engine, single-channel ABS, low seat height, and classic cruiser styling. Priced at ৳2,74,000, it offers unique cruiser experience with comfortable riding position, good fuel efficiency, ABS safety, distinctive styling, and Bajaj reliability making it ideal for riders seeking affordable cruiser aesthetics with comfortable long-distance capability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Avenger 160 ABS Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Cruiser Ergonomics:</strong> <span class="highlight-value">Relaxed riding position with forward footpegs</span></li>
<li class="highlight-item"><strong class="highlight-label">Single-Channel ABS:</strong> <span class="highlight-value">Front ABS for enhanced braking safety</span></li>
<li class="highlight-item"><strong class="highlight-label">Low Seat Height:</strong> <span class="highlight-value">737mm seat height for easy accessibility</span></li>
<li class="highlight-item"><strong class="highlight-label">Classic Cruiser Style:</strong> <span class="highlight-value">Distinctive cruiser aesthetics</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Avenger 160 ABS Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Comfortable Cruiser Position:</strong> <span class="pro-description">Relaxed ergonomics reduce fatigue on long rides</span></li>
<li class="pro-item"><strong class="pro-label">ABS Safety:</strong> <span class="pro-description">Single-channel ABS on front wheel enhances braking safety</span></li>
<li class="pro-item"><strong class="pro-label">Low Seat Accessibility:</strong> <span class="pro-description">737mm seat height suits riders of all heights</span></li>
<li class="pro-item"><strong class="pro-label">Unique Cruiser Style:</strong> <span class="pro-description">Distinctive cruiser aesthetics stand out</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">Achieves 40-45 km/l despite cruiser design</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Cruiser Entry:</strong> <span class="pro-description">Most affordable cruiser motorcycle option</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Avenger 160 ABS Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Performance:</strong> <span class="con-description">160cc cruiser provides basic power only</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">155 kg makes maneuvering difficult in traffic</span></li>
<li class="con-item"><strong class="con-label">Only Front ABS:</strong> <span class="con-description">Single-channel ABS, no rear ABS</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Limited electronics and modern amenities</span></li>
<li class="con-item"><strong class="con-label">Not for City Traffic:</strong> <span class="con-description">Cruiser design impractical for congested traffic</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Bajaj Avenger 160 ABS in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Cruiser styling enthusiasts</li>
<li class="best-for-item">Long-distance comfortable touring</li>
<li class="best-for-item">Relaxed highway cruising</li>
<li class="best-for-item">Shorter riders (low seat height)</li>
<li class="best-for-item">Weekend leisure riding</li>
<li class="best-for-item">Those seeking unique motorcycle style</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Bajaj Avenger 160 ABS?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">City traffic navigation</li>
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Sporty aggressive riding</li>
<li class="not-recommended-item">Heavy daily commuting</li>
<li class="not-recommended-item">Those needing nimble handling</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Avenger 160 ABS Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,74,000 - Good for affordable cruiser with ABS</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳7-9 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳6-8 per km (40-45 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳6,000-9,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Avenger 160 ABS Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- Basic cruiser performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good for cruiser</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent cruiser comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Distinctive cruiser design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good cruiser value</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Decent Bajaj quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Bajaj Avenger 160 ABS Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Is it comfortable for long rides?</h3>
<p class="faq-answer">Yes, the Avenger 160 excels at long-distance touring with relaxed cruiser ergonomics, comfortable seat, and laid-back riding position that reduces fatigue significantly compared to sport bikes.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Can it handle city traffic?</h3>
<p class="faq-answer">The Avenger 160's heavy weight (155 kg) and cruiser design make city traffic navigation challenging. It's better suited for highways and open roads rather than congested urban environments.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is 160cc enough for a cruiser?</h3>
<p class="faq-answer">The 160cc engine provides adequate power for relaxed cruising at 80-90 km/h. It's sufficient for highway touring and comfortable riding, though performance enthusiasts may find it underpowered for a cruiser.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Avenger 160 ABS in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,74,000, the Bajaj Avenger 160 ABS offers the most affordable cruiser motorcycle experience with comfortable cruiser ergonomics, single-channel ABS safety, low seat height (737mm), distinctive styling, and good fuel economy (40-45 km/l). It excels at relaxed highway cruising and long-distance touring with excellent comfort. However, limited 160cc performance, heavy weight (155 kg), only front ABS, basic features, and impracticality for city traffic make it best suited for cruiser styling enthusiasts, touring riders, and those seeking comfortable weekend leisure riding over daily urban commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For affordable cruiser comfort with ABS safety at low seat height</span></p>
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
		return fmt.Errorf("error creating Bajaj Avenger 160 ABS review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Avenger 160 ABS (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস হল ভারতের সবচেয়ে সাশ্রয়ী ক্রুজার মোটরসাইকেল যেখানে শিথিল ক্রুজিং এরগোনমিক্স, ১৬০সিসি ফুয়েল-ইনজেক্টেড ইঞ্জিন, একক-চ্যানেল এবিএস, কম সিট উচ্চতা এবং ক্লাসিক ক্রুজার স্টাইলিং রয়েছে। ৳২,৭৪,০০০ টাকায় মূল্যায়িত, এটি আরামদায়ক রাইডিং পজিশন, ভাল জ্বালানি দক্ষতা, এবিএস নিরাপত্তা, স্বতন্ত্র স্টাইলিং এবং বাজাজ নির্ভরযোগ্যতা সহ অনন্য ক্রুজার অভিজ্ঞতা প্রদান করে যা আরামদায়ক দীর্ঘ-দূরত্ব ক্ষমতা সহ সাশ্রয়ী ক্রুজার নান্দনিকতা খোঁজা রাইডারদের জন্য আদর্শ করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ক্রুজার এরগোনমিক্স:</strong> <span class="highlight-value">সামনের ফুটপেগ সহ শিথিল রাইডিং পজিশন</span></li>
<li class="highlight-item"><strong class="highlight-label">একক-চ্যানেল এবিএস:</strong> <span class="highlight-value">উন্নত ব্রেকিং নিরাপত্তার জন্য সামনের এবিএস</span></li>
<li class="highlight-item"><strong class="highlight-label">কম সিট উচ্চতা:</strong> <span class="highlight-value">সহজ প্রবেশযোগ্যতার জন্য ৭৩৭মিমি সিট উচ্চতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ক্লাসিক ক্রুজার স্টাইল:</strong> <span class="highlight-value">স্বতন্ত্র ক্রুজার নান্দনিকতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আরামদায়ক ক্রুজার পজিশন:</strong> <span class="pro-description">শিথিল এরগোনমিক্স দীর্ঘ রাইডে ক্লান্তি হ্রাস করে</span></li>
<li class="pro-item"><strong class="pro-label">এবিএস নিরাপত্তা:</strong> <span class="pro-description">সামনের চাকায় একক-চ্যানেল এবিএস ব্রেকিং নিরাপত্তা বাড়ায়</span></li>
<li class="pro-item"><strong class="pro-label">কম সিট প্রবেশযোগ্যতা:</strong> <span class="pro-description">৭৩৭মিমি সিট উচ্চতা সব উচ্চতার রাইডারদের জন্য উপযুক্ত</span></li>
<li class="pro-item"><strong class="pro-label">অনন্য ক্রুজার স্টাইল:</strong> <span class="pro-description">স্বতন্ত্র ক্রুজার নান্দনিকতা আলাদা হয়ে দাঁড়ায়</span></li>
<li class="pro-item"><strong class="pro-label">ভাল জ্বালানি সাশ্রয়:</strong> <span class="pro-description">ক্রুজার ডিজাইন সত্ত্বেও ৪০-৪৫ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী ক্রুজার এন্ট্রি:</strong> <span class="pro-description">সবচেয়ে সাশ্রয়ী ক্রুজার মোটরসাইকেল বিকল্প</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত পারফরমেন্স:</strong> <span class="con-description">১৬০সিসি ক্রুজার শুধুমাত্র মৌলিক শক্তি প্রদান করে</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">১৫৫ কেজি ট্রাফিকে চালনা কঠিন করে</span></li>
<li class="con-item"><strong class="con-label">শুধুমাত্র সামনের এবিএস:</strong> <span class="con-description">একক-চ্যানেল এবিএস, পিছনের এবিএস নেই</span></li>
<li class="con-item"><strong class="con-label">মৌলিক বৈশিষ্ট্য:</strong> <span class="con-description">সীমিত ইলেকট্রনিক্স এবং আধুনিক সুবিধা</span></li>
<li class="con-item"><strong class="con-label">শহুরে ট্রাফিকের জন্য নয়:</strong> <span class="con-description">ক্রুজার ডিজাইন জনাকীর্ণ ট্রাফিকের জন্য অব্যবহারিক</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Cruiser styling enthusiasts</li>
<li class="best-for-item">Long-distance comfortable touring</li>
<li class="best-for-item">Relaxed highway cruising</li>
<li class="best-for-item">Shorter riders (low seat height)</li>
<li class="best-for-item">Weekend leisure riding</li>
<li class="best-for-item">Those seeking unique motorcycle style</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">City traffic navigation</li>
<li class="not-recommended-item">Performance enthusiasts</li>
<li class="not-recommended-item">Sporty aggressive riding</li>
<li class="not-recommended-item">Heavy daily commuting</li>
<li class="not-recommended-item">Those needing nimble handling</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৭৪,০০০ - এবিএস সহ সাশ্রয়ী ক্রুজারের জন্য ভাল</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি - ৳৭-৯ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৬-৮ প্রতি কিমি (৪০-৪৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৬,০০০-৯,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- মৌলিক ক্রুজার পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ক্রুজারের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার ক্রুজার আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- স্বতন্ত্র ক্রুজার ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- ভাল ক্রুজার মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল বাজাজ গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এটি কি দীর্ঘ রাইডের জন্য আরামদায়?</h3>
<p class="faq-answer">হ্যাঁ, অ্যাভেঞ্জার ১৬০ শিথিল ক্রুজার এরগোনমিক্স, আরামদায়ক সিট এবং শান্ত রাইডিং পজিশন সহ দীর্ঘ-দূরত্ব ট্যুরিংয়ে শ্রেষ্ঠ যা স্পোর্ট বাইকের তুলনায় ক্লান্তি উল্লেখযোগ্যভাবে হ্রাস করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি শহুরে ট্রাফিক পরিচালনা করতে পারে?</h3>
<p class="faq-answer">অ্যাভেঞ্জার ১৬০ এর ভারী ওজন (১৫৫ কেজি) এবং ক্রুজার ডিজাইন শহুরে ট্রাফিক নেভিগেশন চ্যালেঞ্জিং করে তোলে। এটি জনাকীর্ণ শহুরে পরিবেশের পরিবর্তে হাইওয়ে এবং খোলা রাস্তার জন্য ভাল উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">একটি ক্রুজারের জন্য ১৬০সিসি কি যথেষ্ট?</h3>
<p class="faq-answer">১৬০সিসি ইঞ্জিন ৮০-৯০ কিমি/ঘণ্টায় শিথিল ক্রুজিংয়ের জন্য পর্যাপ্ত শক্তি প্রদান করে। এটি হাইওয়ে ট্যুরিং এবং আরামদায়ক রাইডিংয়ের জন্য যথেষ্ট, যদিও পারফরমেন্স উৎসাহীরা এটিকে একটি ক্রুজারের জন্য কম ক্ষমতাসম্পন্ন মনে করতে পারে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,৭৪,০০০ টাকায় বাজাজ অ্যাভেঞ্জার ১৬০ এবিএস আরামদায়ক ক্রুজার এরগোনমিক্স, একক-চ্যানেল এবিএস নিরাপত্তা, কম সিট উচ্চতা (৭৩৭মিমি), স্বতন্ত্র স্টাইলিং এবং ভাল জ্বালানি অর্থনীতি (৪০-৪৫ কিমি/লিটার) সহ সবচেয়ে সাশ্রয়ী ক্রুজার মোটরসাইকেল অভিজ্ঞতা প্রদান করে। এটি চমৎকার আরাম সহ শিথিল হাইওয়ে ক্রুজিং এবং দীর্ঘ-দূরত্ব ট্যুরিংয়ে শ্রেষ্ঠ। তবে, সীমিত ১৬০সিসি পারফরমেন্স, ভারী ওজন (১৫৫ কেজি), শুধুমাত্র সামনের এবিএস, মৌলিক বৈশিষ্ট্য এবং শহুরে ট্রাফিকের জন্য অব্যবহারিকতা এটিকে ক্রুজার স্টাইলিং উৎসাহী, ট্যুরিং রাইডার এবং দৈনিক শহুরে যাতায়াতের চেয়ে আরামদায়ক সাপ্তাহিক অবসর রাইডিং খোঁজাদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- কম সিট উচ্চতায় এবিএস নিরাপত্তা সহ সাশ্রয়ী ক্রুজার আরামের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Avenger 160 ABS already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Avenger 160 ABS\n")

	return nil
}
