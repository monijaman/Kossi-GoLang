package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCbr150RReview handles seeding New Honda CBR 150R product review and translation
type HondaCbr150RReview struct {
	BaseSeeder
}

// NewHondaCbr150RReviewSeeder creates a new HondaCbr150RReview
func NewHondaCbr150RReviewSeeder() *HondaCbr150RReview {
	return &HondaCbr150RReview{BaseSeeder: BaseSeeder{name: "New Honda CBR 150R Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCbr150RReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "New Honda CBR 150R").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("New Honda CBR 150R product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding New Honda CBR 150R product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for New Honda CBR 150R already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">New Honda CBR 150R Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The New Honda CBR 150R is a premium 150cc fully-faired sports motorcycle featuring aggressive supersport styling, aerodynamic full fairing, liquid-cooled engine, and advanced features. Priced at ৳6,00,000, it offers true sports bike experience with track-inspired design, wind protection, Honda's legendary reliability, excellent build quality, and performance-oriented engineering for enthusiasts seeking a genuine faired sports motorcycle in the 150cc segment.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">New Honda CBR 150R Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Full Fairing Design:</strong> <span class="highlight-value">Aerodynamic supersport styling with complete fairing</span></li>
<li class="highlight-item"><strong class="highlight-label">Liquid-Cooled Engine:</strong> <span class="highlight-value">Advanced cooling for consistent performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Track-Inspired Design:</strong> <span class="highlight-value">Aggressive CBR racing heritage styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Features:</strong> <span class="highlight-value">LED lighting and digital instrumentation</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">New Honda CBR 150R Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Genuine Sports Bike Design:</strong> <span class="pro-description">Full fairing provides true supersport aesthetics and presence</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Aerodynamics:</strong> <span class="pro-description">Full fairing reduces wind resistance for highway stability</span></li>
<li class="pro-item"><strong class="pro-label">Liquid-Cooled Advantage:</strong> <span class="pro-description">Better heat management for consistent performance</span></li>
<li class="pro-item"><strong class="pro-label">Honda Build Quality:</strong> <span class="pro-description">Exceptional fit, finish, and legendary reliability</span></li>
<li class="pro-item"><strong class="pro-label">CBR Heritage:</strong> <span class="pro-description">Part of Honda's iconic CBR racing lineage</span></li>
<li class="pro-item"><strong class="pro-label">Advanced Features:</strong> <span class="pro-description">LED lighting and fully digital instrumentation</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">New Honda CBR 150R Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price Point:</strong> <span class="con-description">৳6,00,000 is expensive for 150cc segment</span></li>
<li class="con-item"><strong class="con-label">Aggressive Riding Position:</strong> <span class="con-description">Forward-leaning ergonomics can be uncomfortable for daily commuting</span></li>
<li class="con-item"><strong class="con-label">Limited City Maneuverability:</strong> <span class="con-description">Full fairing adds weight and bulk for tight traffic</span></li>
<li class="con-item"><strong class="con-label">High Service Costs:</strong> <span class="con-description">Premium bike requires expensive maintenance</span></li>
<li class="con-item"><strong class="con-label">No ABS Option:</strong> <span class="con-description">Missing modern safety feature at premium price</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy New Honda CBR 150R in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Sports bike enthusiasts</li>
<li class="best-for-item">Riders seeking full-faired design</li>
<li class="best-for-item">Highway and long-distance riders</li>
<li class="best-for-item">Young performance-oriented riders</li>
<li class="best-for-item">Those wanting CBR brand prestige</li>
<li class="best-for-item">Track day and spirited riding enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy New Honda CBR 150R?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Daily city commuters in heavy traffic</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Riders needing upright comfortable position</li>
<li class="not-recommended-item">Those requiring ABS safety feature</li>
<li class="not-recommended-item">Beginner riders unfamiliar with sports ergonomics</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">New Honda CBR 150R Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳6,00,000 - Premium for 150cc faired sports bike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-10 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳7-9 per km (40-45 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳10,000-15,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">New Honda CBR 150R Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good sports performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Moderate fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- Sporty aggressive position</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Excellent supersport design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Premium pricing justified</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent Honda quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">New Honda CBR 150R Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Is it suitable for daily commuting?</h3>
<p class="faq-answer">While capable, the aggressive riding position and full fairing make it less ideal for daily city commuting in heavy traffic. It's better suited for highway riding and weekend spirited rides.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it compare to the R15?</h3>
<p class="faq-answer">The CBR 150R offers Honda's legendary reliability and build quality at lower price than R15, but the R15 provides more advanced features like VVA, slipper clutch, and better performance. CBR focuses on reliability and classic styling.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What are the maintenance costs?</h3>
<p class="faq-answer">Annual maintenance costs range from ৳10,000-15,000 with service intervals every 4,000 km. Liquid-cooled engine and premium components mean slightly higher costs than air-cooled bikes.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy New Honda CBR 150R in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳6,00,000, the New Honda CBR 150R delivers a genuine full-faired sports bike experience with supersport styling, liquid-cooled engine, excellent build quality, and Honda reliability. It excels at highway riding and spirited weekend rides with aerodynamic wind protection and CBR heritage prestige. However, the high price for 150cc, aggressive riding position uncomfortable for daily commuting, limited city maneuverability, high service costs, and lack of ABS make it best suited for sports bike enthusiasts and highway riders who prioritize design, reliability, and Honda brand heritage over daily practicality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For genuine full-faired sports bike experience with Honda reliability</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.4,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating New Honda CBR 150R review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for New Honda CBR 150R (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">নিউ হোন্ডা সিবিআর ১৫০আর রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">নিউ হোন্ডা সিবিআর ১৫০আর হল একটি প্রিমিয়াম ১৫০সিসি সম্পূর্ণ-ফেয়ারড স্পোর্টস মোটরসাইকেল যেখানে আক্রমণাত্মক সুপারস্পোর্ট স্টাইলিং, অ্যারোডাইনামিক সম্পূর্ণ ফেয়ারিং, লিকুইড-কুলড ইঞ্জিন এবং উন্নত বৈশিষ্ট্য রয়েছে। ৳৬,০০,০০০ টাকায় মূল্যায়িত, এটি ট্র্যাক-অনুপ্রাণিত ডিজাইন, বায়ু সুরক্ষা, হোন্ডার কিংবদন্তি নির্ভরযোগ্যতা, চমৎকার বিল্ড কোয়ালিটি এবং ১৫০সিসি সেগমেন্টে প্রকৃত ফেয়ারড স্পোর্টস মোটরসাইকেল খোঁজা উৎসাহীদের জন্য পারফরমেন্স-ওরিয়েন্টেড ইঞ্জিনিয়ারিং সহ সত্যিকারের স্পোর্টস বাইক অভিজ্ঞতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">নিউ হোন্ডা সিবিআর ১৫০আর এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ ফেয়ারিং ডিজাইন:</strong> <span class="highlight-value">সম্পূর্ণ ফেয়ারিং সহ অ্যারোডাইনামিক সুপারস্পোর্ট স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">লিকুইড-কুলড ইঞ্জিন:</strong> <span class="highlight-value">সামঞ্জস্যপূর্ণ পারফরমেন্সের জন্য উন্নত কুলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্র্যাক-অনুপ্রাণিত ডিজাইন:</strong> <span class="highlight-value">আক্রমণাত্মক সিবিআর রেসিং ঐতিহ্য স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম বৈশিষ্ট্য:</strong> <span class="highlight-value">এলইডি লাইটিং এবং ডিজিটাল ইনস্ট্রুমেন্টেশন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">নিউ হোন্ডা সিবিআর ১৫০আর এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রকৃত স্পোর্টস বাইক ডিজাইন:</strong> <span class="pro-description">সম্পূর্ণ ফেয়ারিং সত্যিকারের সুপারস্পোর্ট নান্দনিকতা এবং উপস্থিতি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার অ্যারোডাইনামিক্স:</strong> <span class="pro-description">সম্পূর্ণ ফেয়ারিং হাইওয়ে স্থিতিশীলতার জন্য বায়ু প্রতিরোধ হ্রাস করে</span></li>
<li class="pro-item"><strong class="pro-label">লিকুইড-কুলড সুবিধা:</strong> <span class="pro-description">সামঞ্জস্যপূর্ণ পারফরমেন্সের জন্য ভাল তাপ ব্যবস্থাপনা</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা বিল্ড কোয়ালিটি:</strong> <span class="pro-description">ব্যতিক্রমী ফিট, ফিনিশ এবং কিংবদন্তি নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">সিবিআর ঐতিহ্য:</strong> <span class="pro-description">হোন্ডার আইকনিক সিবিআর রেসিং বংশের অংশ</span></li>
<li class="pro-item"><strong class="pro-label">উন্নত বৈশিষ্ট্য:</strong> <span class="pro-description">এলইডি লাইটিং এবং সম্পূর্ণ ডিজিটাল ইনস্ট্রুমেন্টেশন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">নিউ হোন্ডা সিবিআর ১৫০আর এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য পয়েন্ট:</strong> <span class="con-description">৳৬,০০,০০০ ১৫০সিসি সেগমেন্টের জন্য ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">আক্রমণাত্মক রাইডিং পজিশন:</strong> <span class="con-description">সামনের দিকে ঝুঁকে পড়া এরগোনমিক্স দৈনিক যাতায়াতের জন্য অস্বস্তিকর হতে পারে</span></li>
<li class="con-item"><strong class="con-label">সীমিত শহুরে চালনা:</strong> <span class="con-description">সম্পূর্ণ ফেয়ারিং আঁটসাঁট ট্রাফিকের জন্য ওজন এবং বড়তা যোগ করে</span></li>
<li class="con-item"><strong class="con-label">উচ্চ সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম বাইকের ব্যয়বহুল রক্ষণাবেক্ষণ প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস বিকল্প নেই:</strong> <span class="con-description">প্রিমিয়াম মূল্যে আধুনিক নিরাপত্তা বৈশিষ্ট্য অনুপস্থিত</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে নিউ হোন্ডা সিবিআর ১৫০আর কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Sports bike enthusiasts</li>
<li class="best-for-item">Riders seeking full-faired design</li>
<li class="best-for-item">Highway and long-distance riders</li>
<li class="best-for-item">Young performance-oriented riders</li>
<li class="best-for-item">Those wanting CBR brand prestige</li>
<li class="best-for-item">Track day and spirited riding enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">নিউ হোন্ডা সিবিআর ১৫০আর কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Daily city commuters in heavy traffic</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Riders needing upright comfortable position</li>
<li class="not-recommended-item">Those requiring ABS safety feature</li>
<li class="not-recommended-item">Beginner riders unfamiliar with sports ergonomics</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">নিউ হোন্ডা সিবিআর ১৫০আর এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৬,০০,০০০ - ১৫০সিসি ফেয়ারড স্পোর্টস বাইকের জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি - ৳৮-১০ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৭-৯ প্রতি কিমি (৪০-৪৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১০,০০০-১৫,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">নিউ হোন্ডা সিবিআর ১৫০আর পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ভাল স্পোর্টস পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- মাঝারি জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.5</span> <span class="rating-note">- স্পোর্টি আক্রমণাত্মক পজিশন</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- চমৎকার সুপারস্পোর্ট ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- প্রিমিয়াম মূল্য ন্যায্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার হোন্ডা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">নিউ হোন্ডা সিবিআর ১৫০আর সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এটি কি দৈনিক যাতায়াতের জন্য উপযুক্ত?</h3>
<p class="faq-answer">সক্ষম হলেও, আক্রমণাত্মক রাইডিং পজিশন এবং সম্পূর্ণ ফেয়ারিং এটিকে ভারী ট্রাফিকে দৈনিক শহুরে যাতায়াতের জন্য কম আদর্শ করে তোলে। এটি হাইওয়ে রাইডিং এবং সাপ্তাহিক উৎসাহী রাইডের জন্য ভাল উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি আর১৫ এর সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">সিবিআর ১৫০আর আর১৫ এর চেয়ে কম মূল্যে হোন্ডার কিংবদন্তি নির্ভরযোগ্যতা এবং বিল্ড কোয়ালিটি প্রদান করে, কিন্তু আর১৫ ভিভিএ, স্লিপার ক্লাচ এবং ভাল পারফরমেন্সের মতো আরও উন্নত বৈশিষ্ট্য প্রদান করে। সিবিআর নির্ভরযোগ্যতা এবং ক্লাসিক স্টাইলিংয়ে ফোকাস করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রক্ষণাবেক্ষণ খরচ কত?</h3>
<p class="faq-answer">বার্ষিক রক্ষণাবেক্ষণ খরচ প্রতি ৪,০০০ কিমিতে সেবা ব্যবধান সহ ৳১০,০০০-১৫,০০০ পরিসীমা। লিকুইড-কুলড ইঞ্জিন এবং প্রিমিয়াম কম্পোনেন্ট মানে এয়ার-কুলড বাইকের চেয়ে সামান্য বেশি খরচ।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে নিউ হোন্ডা সিবিআর ১৫০আর কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৬,০০,০০০ টাকায় নিউ হোন্ডা সিবিআর ১৫০আর সুপারস্পোর্ট স্টাইলিং, লিকুইড-কুলড ইঞ্জিন, চমৎকার বিল্ড কোয়ালিটি এবং হোন্ডা নির্ভরযোগ্যতা সহ একটি প্রকৃত সম্পূর্ণ-ফেয়ারড স্পোর্টস বাইক অভিজ্ঞতা প্রদান করে। এটি অ্যারোডাইনামিক বায়ু সুরক্ষা এবং সিবিআর ঐতিহ্য প্রতিপত্তি সহ হাইওয়ে রাইডিং এবং উৎসাহী সাপ্তাহিক রাইডে শ্রেষ্ঠ। তবে, ১৫০সিসির জন্য উচ্চ মূল্য, দৈনিক যাতায়াতের জন্য অস্বস্তিকর আক্রমণাত্মক রাইডিং পজিশন, সীমিত শহুরে চালনা, উচ্চ সেবা খরচ এবং এবিএসের অভাব এটিকে স্পোর্টস বাইক উৎসাহী এবং হাইওয়ে রাইডারদের জন্য সবচেয়ে উপযুক্ত করে তোলে যারা দৈনিক ব্যবহারযোগ্যতার চেয়ে ডিজাইন, নির্ভরযোগ্যতা এবং হোন্ডা ব্র্যান্ড ঐতিহ্যকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- হোন্ডা নির্ভরযোগ্যতা সহ প্রকৃত সম্পূর্ণ-ফেয়ারড স্পোর্টস বাইক অভিজ্ঞতার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for New Honda CBR 150R already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for New Honda CBR 150R\n")

	return nil
}
