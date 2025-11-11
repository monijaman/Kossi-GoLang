package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaPcx150Review handles seeding Honda PCX 150 product review and translation
type HondaPcx150Review struct {
	BaseSeeder
}

// NewHondaPcx150ReviewSeeder creates a new HondaPcx150Review
func NewHondaPcx150ReviewSeeder() *HondaPcx150Review {
	return &HondaPcx150Review{BaseSeeder: BaseSeeder{name: "Honda PCX 150 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaPcx150Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda PCX 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda PCX 150 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda PCX 150 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda PCX 150 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda PCX 150 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda PCX 150 is a premium 150cc automatic scooter featuring sophisticated European-inspired styling, advanced eSP+ engine technology, spacious under-seat storage, full LED lighting, and luxurious features. Priced at ৳5,10,000, it offers premium scooter experience with smooth automatic transmission, excellent fuel efficiency, comfortable ergonomics, smart key system, and Honda's legendary reliability for urban riders seeking sophisticated, practical, and refined automatic transportation.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda PCX 150 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">eSP+ Engine Technology:</strong> <span class="highlight-value">Advanced engine for efficiency and smooth power</span></li>
<li class="highlight-item"><strong class="highlight-label">Smart Key System:</strong> <span class="highlight-value">Keyless operation with answer-back feature</span></li>
<li class="highlight-item"><strong class="highlight-label">Full LED Lighting:</strong> <span class="highlight-value">Premium LED headlight, taillight, and indicators</span></li>
<li class="highlight-item"><strong class="highlight-label">Spacious Storage:</strong> <span class="highlight-value">Large under-seat storage for full-face helmet</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda PCX 150 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Premium Sophistication:</strong> <span class="pro-description">European-inspired styling with luxurious feel and finish</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Fuel Efficiency:</strong> <span class="pro-description">Achieves 50-55 km/l with eSP+ engine technology</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Automatic Transmission:</strong> <span class="pro-description">V-Matic transmission provides effortless riding experience</span></li>
<li class="pro-item"><strong class="pro-label">Smart Key Convenience:</strong> <span class="pro-description">Keyless operation with remote answer-back system</span></li>
<li class="pro-item"><strong class="pro-label">Practical Storage:</strong> <span class="pro-description">Spacious under-seat storage accommodates full-face helmet</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Well-tuned suspension and ergonomics for daily comfort</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda PCX 150 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very High Price:</strong> <span class="con-description">৳5,10,000 is extremely expensive for 150cc scooter segment</span></li>
<li class="con-item"><strong class="con-label">Limited Availability:</strong> <span class="con-description">Not widely available, limited Honda dealership network</span></li>
<li class="con-item"><strong class="con-label">Expensive Parts:</strong> <span class="con-description">Premium components mean high replacement costs</span></li>
<li class="con-item"><strong class="con-label">Complex Electronics:</strong> <span class="con-description">Smart key and electronics may be costly to repair</span></li>
<li class="con-item"><strong class="con-label">Overkill for Basic Commuting:</strong> <span class="con-description">Premium features unnecessary for simple daily use</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda PCX 150 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Premium scooter enthusiasts</li>
<li class="best-for-item">Urban professionals seeking sophistication</li>
<li class="best-for-item">Those prioritizing convenience and comfort</li>
<li class="best-for-item">Riders wanting automatic transmission luxury</li>
<li class="best-for-item">Style-conscious mature riders</li>
<li class="best-for-item">Daily commuters with premium budget</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda PCX 150?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those seeking basic economical transportation</li>
<li class="not-recommended-item">Riders in areas with limited service network</li>
<li class="not-recommended-item">Performance-oriented riders</li>
<li class="not-recommended-item">Those uncomfortable with electronics complexity</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda PCX 150 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳5,10,000 - Premium for 150cc automatic scooter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low to Moderate - ৳6-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳5-7 per km (50-55 km/l with eSP+ efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳8,000-12,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda PCX 150 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Good scooter performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Outstanding comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Premium European styling</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- Premium pricing</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Exceptional build quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda PCX 150 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Is the smart key system reliable?</h3>
<p class="faq-answer">Yes, the smart key system is reliable and adds great convenience with keyless operation and answer-back feature. However, battery replacement can be expensive, and having a backup key is recommended.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it compare to Honda Dio or Activa?</h3>
<p class="faq-answer">The PCX 150 is significantly more premium with 150cc engine, smart key, full LED lighting, and superior build quality. It costs about ৳4 lakh more but offers a completely different level of sophistication and features compared to entry-level Dio/Activa.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for long-distance touring?</h3>
<p class="faq-answer">Yes, the PCX 150 is excellent for touring with comfortable ergonomics, good fuel range, spacious storage, and smooth automatic transmission. The 150cc engine provides adequate highway cruising capability.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda PCX 150 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳5,10,000, the Honda PCX 150 represents the pinnacle of premium automatic scooters with sophisticated European styling, eSP+ engine efficiency (50-55 km/l), smart key convenience, full LED lighting, and exceptional build quality. It excels at providing luxurious automated transportation with outstanding comfort and practical storage. However, the very high price for scooter segment, limited availability, expensive parts, complex electronics, and premium features being overkill for basic commuting make it suitable only for affluent buyers, premium scooter enthusiasts, and urban professionals who can afford the luxury and appreciate the sophisticated refinement over basic economy.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium automatic sophistication with smart key convenience</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Honda PCX 150 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda PCX 150 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা পিসিএক্স ১৫০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা পিসিএক্স ১৫০ হল একটি প্রিমিয়াম ১৫০সিসি অটোমেটিক স্কুটার যেখানে পরিশীলিত ইউরোপীয়-অনুপ্রাণিত স্টাইলিং, উন্নত ইএসপি+ ইঞ্জিন প্রযুক্তি, প্রশস্ত আন্ডার-সিট স্টোরেজ, সম্পূর্ণ এলইডি লাইটিং এবং বিলাসবহুল বৈশিষ্ট্য রয়েছে। ৳৫,১০,০০০ টাকায় মূল্যায়িত, এটি মসৃণ অটোমেটিক ট্রান্সমিশন, চমৎকার জ্বালানি দক্ষতা, আরামদায়ক এরগোনমিক্স, স্মার্ট কী সিস্টেম এবং পরিশীলিত, ব্যবহারিক এবং পরিমার্জিত অটোমেটিক পরিবহন খোঁজা শহুরে রাইডারদের জন্য হোন্ডার কিংবদন্তি নির্ভরযোগ্যতা সহ প্রিমিয়াম স্কুটার অভিজ্ঞতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা পিসিএক্স ১৫০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ইএসপি+ ইঞ্জিন প্রযুক্তি:</strong> <span class="highlight-value">দক্ষতা এবং মসৃণ শক্তির জন্য উন্নত ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">স্মার্ট কী সিস্টেম:</strong> <span class="highlight-value">আনসার-ব্যাক বৈশিষ্ট্য সহ কীলেস অপারেশন</span></li>
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ এলইডি লাইটিং:</strong> <span class="highlight-value">প্রিমিয়াম এলইডি হেডলাইট, টেইললাইট এবং ইন্ডিকেটর</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রশস্ত স্টোরেজ:</strong> <span class="highlight-value">ফুল-ফেস হেলমেটের জন্য বড় আন্ডার-সিট স্টোরেজ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা পিসিএক্স ১৫০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম পরিশীলিততা:</strong> <span class="pro-description">বিলাসবহুল অনুভূতি এবং ফিনিশ সহ ইউরোপীয়-অনুপ্রাণিত স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার জ্বালানি দক্ষতা:</strong> <span class="pro-description">ইএসপি+ ইঞ্জিন প্রযুক্তি দিয়ে ৫০-৫৫ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ অটোমেটিক ট্রান্সমিশন:</strong> <span class="pro-description">ভি-ম্যাটিক ট্রান্সমিশন অনায়াস রাইডিং অভিজ্ঞতা প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">স্মার্ট কী সুবিধা:</strong> <span class="pro-description">রিমোট আনসার-ব্যাক সিস্টেম সহ কীলেস অপারেশন</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক স্টোরেজ:</strong> <span class="pro-description">প্রশস্ত আন্ডার-সিট স্টোরেজ ফুল-ফেস হেলমেট ধারণ করে</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">দৈনিক আরামের জন্য ভাল-টিউনড সাসপেনশন এবং এরগোনমিক্স</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা পিসিএক্স ১৫০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত উচ্চ মূল্য:</strong> <span class="con-description">৳৫,১০,০০০ ১৫০সিসি স্কুটার সেগমেন্টের জন্য অত্যন্ত ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সীমিত প্রাপ্যতা:</strong> <span class="con-description">ব্যাপকভাবে উপলব্ধ নয়, সীমিত হোন্ডা ডিলারশিপ নেটওয়ার্ক</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল যন্ত্রাংশ:</strong> <span class="con-description">প্রিমিয়াম কম্পোনেন্ট মানে উচ্চ প্রতিস্থাপন খরচ</span></li>
<li class="con-item"><strong class="con-label">জটিল ইলেকট্রনিক্স:</strong> <span class="con-description">স্মার্ট কী এবং ইলেকট্রনিক্স মেরামত করতে ব্যয়বহুল হতে পারে</span></li>
<li class="con-item"><strong class="con-label">মৌলিক যাতায়াতের জন্য অতিরিক্ত:</strong> <span class="con-description">সহজ দৈনিক ব্যবহারের জন্য প্রিমিয়াম বৈশিষ্ট্য অপ্রয়োজনীয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা পিসিএক্স ১৫০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Premium scooter enthusiasts</li>
<li class="best-for-item">Urban professionals seeking sophistication</li>
<li class="best-for-item">Those prioritizing convenience and comfort</li>
<li class="best-for-item">Riders wanting automatic transmission luxury</li>
<li class="best-for-item">Style-conscious mature riders</li>
<li class="best-for-item">Daily commuters with premium budget</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা পিসিএক্স ১৫০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those seeking basic economical transportation</li>
<li class="not-recommended-item">Riders in areas with limited service network</li>
<li class="not-recommended-item">Performance-oriented riders</li>
<li class="not-recommended-item">Those uncomfortable with electronics complexity</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা পিসিএক্স ১৫০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৫,১০,০০০ - ১৫০সিসি অটোমেটিক স্কুটারের জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম থেকে মাঝারি - ৳৬-৮ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৫-৭ প্রতি কিমি (ইএসপি+ দক্ষতা সহ ৫০-৫৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৮,০০০-১২,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা পিসিএক্স ১৫০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- ভাল স্কুটার পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- অসাধারণ আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- প্রিমিয়াম ইউরোপীয় স্টাইলিং</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- প্রিমিয়াম মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- ব্যতিক্রমী বিল্ড কোয়ালিটি</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা পিসিএক্স ১৫০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">স্মার্ট কী সিস্টেম কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">হ্যাঁ, স্মার্ট কী সিস্টেম নির্ভরযোগ্য এবং কীলেস অপারেশন এবং আনসার-ব্যাক বৈশিষ্ট্য দিয়ে দুর্দান্ত সুবিধা যোগ করে। তবে, ব্যাটারি প্রতিস্থাপন ব্যয়বহুল হতে পারে এবং ব্যাকআপ কী রাখার পরামর্শ দেওয়া হয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি হোন্ডা ডিও বা অ্যাক্টিভার সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">পিসিএক্স ১৫০ ১৫০সিসি ইঞ্জিন, স্মার্ট কী, সম্পূর্ণ এলইডি লাইটিং এবং উন্নত বিল্ড কোয়ালিটি সহ উল্লেখযোগ্যভাবে আরও প্রিমিয়াম। এটি প্রায় ৳৪ লক্ষ বেশি খরচ করে কিন্তু এন্ট্রি-লেভেল ডিও/অ্যাক্টিভার তুলনায় সম্পূর্ণ ভিন্ন স্তরের পরিশীলিততা এবং বৈশিষ্ট্য প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি দীর্ঘ-দূরত্ব ট্যুরিংয়ের জন্য উপযুক্ত?</h3>
<p class="faq-answer">হ্যাঁ, পিসিএক্স ১৫০ আরামদায়ক এরগোনমিক্স, ভাল জ্বালানি রেঞ্জ, প্রশস্ত স্টোরেজ এবং মসৃণ অটোমেটিক ট্রান্সমিশন সহ ট্যুরিংয়ের জন্য চমৎকার। ১৫০সিসি ইঞ্জিন পর্যাপ্ত হাইওয়ে ক্রুজিং ক্ষমতা প্রদান করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা পিসিএক্স ১৫০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৫,১০,০০০ টাকায় হোন্ডা পিসিএক্স ১৫০ পরিশীলিত ইউরোপীয় স্টাইলিং, ইএসপি+ ইঞ্জিন দক্ষতা (৫০-৫৫ কিমি/লিটার), স্মার্ট কী সুবিধা, সম্পূর্ণ এলইডি লাইটিং এবং ব্যতিক্রমী বিল্ড কোয়ালিটি সহ প্রিমিয়াম অটোমেটিক স্কুটারের শিখর উপস্থাপন করে। এটি অসাধারণ আরাম এবং ব্যবহারিক স্টোরেজ সহ বিলাসবহুল স্বয়ংক্রিয় পরিবহন প্রদানে শ্রেষ্ঠ। তবে, স্কুটার সেগমেন্টের জন্য অত্যন্ত উচ্চ মূল্য, সীমিত প্রাপ্যতা, ব্যয়বহুল যন্ত্রাংশ, জটিল ইলেকট্রনিক্স এবং মৌলিক যাতায়াতের জন্য প্রিমিয়াম বৈশিষ্ট্য অতিরিক্ত হওয়া এটিকে শুধুমাত্র সচ্ছল ক্রেতা, প্রিমিয়াম স্কুটার উৎসাহী এবং শহুরে পেশাদারদের জন্য উপযুক্ত করে তোলে যারা বিলাসিতা বহন করতে পারে এবং মৌলিক অর্থনীতির চেয়ে পরিশীলিত পরিমার্জনের প্রশংসা করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্মার্ট কী সুবিধা সহ প্রিমিয়াম অটোমেটিক পরিশীলিততার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda PCX 150 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda PCX 150\n")

	return nil
}
