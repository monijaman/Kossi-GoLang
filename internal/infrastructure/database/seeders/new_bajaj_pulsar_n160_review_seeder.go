package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// NewBajajPulsarN160Review handles seeding New Bajaj Pulsar N160 product review and translation
type NewBajajPulsarN160Review struct {
	BaseSeeder
}

// NewNewBajajPulsarN160ReviewSeeder creates a new NewBajajPulsarN160Review
func NewNewBajajPulsarN160ReviewSeeder() *NewBajajPulsarN160Review {
	return &NewBajajPulsarN160Review{BaseSeeder: BaseSeeder{name: "New Bajaj Pulsar N160 Review"}}
}

// Seed implements the Seeder interface
func (s *NewBajajPulsarN160Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "New Bajaj Pulsar N160").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("New Bajaj Pulsar N160 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding New Bajaj Pulsar N160 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for New Bajaj Pulsar N160 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">New Bajaj Pulsar N160 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The New Bajaj Pulsar N160 is an updated naked street fighter featuring aggressive muscular styling, 160cc oil-cooled engine, dual-channel ABS, LED lighting with infinity LED DRL, and sporty ergonomics. Priced at ৳2,73,600, it offers excellent value with modern features, sharp handling, dual-channel ABS safety, good performance, and Pulsar reliability making it ideal for young riders seeking affordable naked sports styling with practical versatility.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">New Bajaj Pulsar N160 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Dual-Channel ABS:</strong> <span class="highlight-value">Advanced ABS safety on both wheels</span></li>
<li class="highlight-item"><strong class="highlight-label">Infinity LED DRL:</strong> <span class="highlight-value">Distinctive LED lighting signature</span></li>
<li class="highlight-item"><strong class="highlight-label">Muscular Naked Styling:</strong> <span class="highlight-value">Aggressive street fighter design</span></li>
<li class="highlight-item"><strong class="highlight-label">Oil-Cooled Engine:</strong> <span class="highlight-value">160cc with enhanced cooling</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">New Bajaj Pulsar N160 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Dual-Channel ABS Safety:</strong> <span class="pro-description">Superior braking safety on both wheels rare in 160cc segment</span></li>
<li class="pro-item"><strong class="pro-label">Modern LED Lighting:</strong> <span class="pro-description">Infinity LED DRL creates distinctive visual identity</span></li>
<li class="pro-item"><strong class="pro-label">Muscular Styling:</strong> <span class="pro-description">Aggressive street fighter aesthetics attract young riders</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">Oil-cooled 160cc provides spirited performance</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Value:</strong> <span class="pro-description">Dual-channel ABS at ৳2,73,600 is competitive</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Lightweight naked design enhances agility</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">New Bajaj Pulsar N160 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Oil-Cooled Limitations:</strong> <span class="con-description">Shows vibrations at higher RPM</span></li>
<li class="con-item"><strong class="con-label">Firm Suspension:</strong> <span class="con-description">Sporty setup can feel harsh on rough roads</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Fit and finish not premium level</span></li>
<li class="con-item"><strong class="con-label">Limited 160cc Power:</strong> <span class="con-description">May feel underpowered for performance enthusiasts</span></li>
<li class="con-item"><strong class="con-label">Basic Digital Display:</strong> <span class="con-description">Instrument cluster lacks advanced features</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy New Bajaj Pulsar N160 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban riders</li>
<li class="best-for-item">Safety-conscious buyers wanting dual ABS</li>
<li class="best-for-item">Daily sporty commuting</li>
<li class="best-for-item">First-time naked bike buyers</li>
<li class="best-for-item">Budget performance enthusiasts</li>
<li class="best-for-item">College and office commuters</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy New Bajaj Pulsar N160?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Premium build quality seekers</li>
<li class="not-recommended-item">High-performance requirements</li>
<li class="not-recommended-item">Soft suspension comfort needs</li>
<li class="not-recommended-item">Long-distance touring primary use</li>
<li class="not-recommended-item">Those sensitive to vibrations</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">New Bajaj Pulsar N160 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,73,600 - Excellent for 160cc with dual-channel ABS</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low to Moderate - ৳6-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳5-7 per km (45-50 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳6,000-9,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">New Bajaj Pulsar N160 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good 160cc performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Decent sporty comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Aggressive naked design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Outstanding value with ABS</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good Bajaj quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">New Bajaj Pulsar N160 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Is dual-channel ABS worth the extra cost?</h3>
<p class="faq-answer">Yes, dual-channel ABS provides significantly better braking safety on both wheels, especially in emergency situations and wet conditions. It's a valuable safety feature rare in the 160cc segment at this price.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it compare to traditional Pulsar 150?</h3>
<p class="faq-answer">The N160 offers more displacement (160cc vs 150cc), dual-channel ABS, modern LED lighting with infinity DRL, and updated styling. It's more premium and feature-rich than traditional Pulsar 150 models.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it good for daily city commuting?</h3>
<p class="faq-answer">Yes, the N160 excels at city commuting with good fuel economy (45-50 km/l), nimble handling, adequate power, and dual-channel ABS safety making it ideal for daily urban use.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy New Bajaj Pulsar N160 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,73,600, the New Bajaj Pulsar N160 offers outstanding value as an affordable naked 160cc street fighter with rare dual-channel ABS safety, modern infinity LED DRL lighting, muscular styling, good performance, and excellent fuel economy (45-50 km/l). It excels at sporty urban commuting with sharp handling and agility. However, oil-cooled engine vibrations, firm suspension, average build quality, limited 160cc power for enthusiasts, and basic display make it best suited for young urban riders, safety-conscious buyers, and daily commuters prioritizing dual ABS safety, modern styling, and value over premium refinement.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For outstanding value 160cc with rare dual-channel ABS safety</span></p>
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
		return fmt.Errorf("error creating New Bajaj Pulsar N160 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for New Bajaj Pulsar N160 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">নিউ বাজাজ পালসার এন১৬০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">নিউ বাজাজ পালসার এন১৬০ হল একটি আপডেটেড নেকেড স্ট্রিট ফাইটার যেখানে আক্রমণাত্মক মাসকুলার স্টাইলিং, ১৬০সিসি অয়েল-কুলড ইঞ্জিন, ডুয়াল-চ্যানেল এবিএস, ইনফিনিটি এলইডি ডিআরএল সহ এলইডি লাইটিং এবং স্পোর্টি এরগোনমিক্স রয়েছে। ৳২,৭৩,৬০০ টাকায় মূল্যায়িত, এটি আধুনিক বৈশিষ্ট্য, তীক্ষ্ণ হ্যান্ডলিং, ডুয়াল-চ্যানেল এবিএস নিরাপত্তা, ভাল পারফরমেন্স এবং পালসার নির্ভরযোগ্যতা সহ চমৎকার মূল্য প্রদান করে যা ব্যবহারিক বহুমুখিতা সহ সাশ্রয়ী নেকেড স্পোর্টস স্টাইলিং খোঁজা তরুণ রাইডারদের জন্য আদর্শ করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">নিউ বাজাজ পালসার এন১৬০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ডুয়াল-চ্যানেল এবিএস:</strong> <span class="highlight-value">উভয় চাকায় উন্নত এবিএস নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">ইনফিনিটি এলইডি ডিআরএল:</strong> <span class="highlight-value">স্বতন্ত্র এলইডি লাইটিং স্বাক্ষর</span></li>
<li class="highlight-item"><strong class="highlight-label">মাসকুলার নেকেড স্টাইলিং:</strong> <span class="highlight-value">আক্রমণাত্মক স্ট্রিট ফাইটার ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">অয়েল-কুলড ইঞ্জিন:</strong> <span class="highlight-value">উন্নত কুলিং সহ ১৬০সিসি</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">নিউ বাজাজ পালসার এন১৬০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ডুয়াল-চ্যানেল এবিএস নিরাপত্তা:</strong> <span class="pro-description">উভয় চাকায় উন্নত ব্রেকিং নিরাপত্তা ১৬০সিসি সেগমেন্টে বিরল</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক এলইডি লাইটিং:</strong> <span class="pro-description">ইনফিনিটি এলইডি ডিআরএল স্বতন্ত্র ভিজ্যুয়াল পরিচয় তৈরি করে</span></li>
<li class="pro-item"><strong class="pro-label">মাসকুলার স্টাইলিং:</strong> <span class="pro-description">আক্রমণাত্মক স্ট্রিট ফাইটার নান্দনিকতা তরুণ রাইডারদের আকর্ষণ করে</span></li>
<li class="pro-item"><strong class="pro-label">ভাল পারফরমেন্স:</strong> <span class="pro-description">অয়েল-কুলড ১৬০সিসি উৎসাহী পারফরমেন্স প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মূল্য:</strong> <span class="pro-description">৳২,৭৩,৬০০ এ ডুয়াল-চ্যানেল এবিএস প্রতিযোগিতামূলক</span></li>
<li class="pro-item"><strong class="pro-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="pro-description">লাইটওয়েট নেকেড ডিজাইন চপলতা বাড়ায়</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">নিউ বাজাজ পালসার এন১৬০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অয়েল-কুলড সীমাবদ্ধতা:</strong> <span class="con-description">উচ্চ আরপিএমে কম্পন দেখায়</span></li>
<li class="con-item"><strong class="con-label">শক্ত সাসপেনশন:</strong> <span class="con-description">স্পোর্টি সেটআপ রুক্ষ রাস্তায় কঠোর মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">ফিট এবং ফিনিশ প্রিমিয়াম লেভেল নয়</span></li>
<li class="con-item"><strong class="con-label">সীমিত ১৬০সিসি শক্তি:</strong> <span class="con-description">পারফরমেন্স উৎসাহীদের জন্য কম ক্ষমতাসম্পন্ন মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">মৌলিক ডিজিটাল ডিসপ্লে:</strong> <span class="con-description">ইনস্ট্রুমেন্ট ক্লাস্টারে উন্নত বৈশিষ্ট্যের অভাব</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে নিউ বাজাজ পালসার এন১৬০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban riders</li>
<li class="best-for-item">Safety-conscious buyers wanting dual ABS</li>
<li class="best-for-item">Daily sporty commuting</li>
<li class="best-for-item">First-time naked bike buyers</li>
<li class="best-for-item">Budget performance enthusiasts</li>
<li class="best-for-item">College and office commuters</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">নিউ বাজাজ পালসার এন১৬০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Premium build quality seekers</li>
<li class="not-recommended-item">High-performance requirements</li>
<li class="not-recommended-item">Soft suspension comfort needs</li>
<li class="not-recommended-item">Long-distance touring primary use</li>
<li class="not-recommended-item">Those sensitive to vibrations</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">নিউ বাজাজ পালসার এন১৬০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳২,৭৩,৬০০ - ডুয়াল-চ্যানেল এবিএস সহ ১৬০সিসির জন্য চমৎকার</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">কম থেকে মাঝারি - ৳৬-৮ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৫-৭ প্রতি কিমি (৪৫-৫০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৬,০০০-৯,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">নিউ বাজাজ পালসার এন১৬০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ভাল ১৬০সিসি পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- চমৎকার জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- ভাল স্পোর্টি আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- আক্রমণাত্মক নেকেড ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- এবিএস সহ অসাধারণ মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল বাজাজ গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">নিউ বাজাজ পালসার এন১৬০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">অতিরিক্ত খরচের জন্য ডুয়াল-চ্যানেল এবিএস কি মূল্যবান?</h3>
<p class="faq-answer">হ্যাঁ, ডুয়াল-চ্যানেল এবিএস উভয় চাকায় উল্লেখযোগ্যভাবে ভাল ব্রেকিং নিরাপত্তা প্রদান করে, বিশেষত জরুরি পরিস্থিতি এবং ভেজা অবস্থায়। এটি এই মূল্যে ১৬০সিসি সেগমেন্টে বিরল একটি মূল্যবান নিরাপত্তা বৈশিষ্ট্য।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি ঐতিহ্যবাহী পালসার ১৫০ এর সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">এন১৬০ আরও ডিসপ্লেসমেন্ট (১৬০সিসি বনাম ১৫০সিসি), ডুয়াল-চ্যানেল এবিএস, ইনফিনিটি ডিআরএল সহ আধুনিক এলইডি লাইটিং এবং আপডেটেড স্টাইলিং প্রদান করে। এটি ঐতিহ্যবাহী পালসার ১৫০ মডেলের চেয়ে আরও প্রিমিয়াম এবং ফিচার-সমৃদ্ধ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি দৈনিক শহুরে যাতায়াতের জন্য ভাল?</h3>
<p class="faq-answer">হ্যাঁ, এন১৬০ ভাল জ্বালানি অর্থনীতি (৪৫-৫০ কিমি/লিটার), চটপটে হ্যান্ডলিং, পর্যাপ্ত শক্তি এবং ডুয়াল-চ্যানেল এবিএস নিরাপত্তা সহ শহুরে যাতায়াতে শ্রেষ্ঠ যা এটিকে দৈনিক শহুরে ব্যবহারের জন্য আদর্শ করে তোলে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে নিউ বাজাজ পালসার এন১৬০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳২,৭৩,৬০০ টাকায় নিউ বাজাজ পালসার এন১৬০ বিরল ডুয়াল-চ্যানেল এবিএস নিরাপত্তা, আধুনিক ইনফিনিটি এলইডি ডিআরএল লাইটিং, মাসকুলার স্টাইলিং, ভাল পারফরমেন্স এবং চমৎকার জ্বালানি অর্থনীতি (৪৫-৫০ কিমি/লিটার) সহ একটি সাশ্রয়ী নেকেড ১৬০সিসি স্ট্রিট ফাইটার হিসাবে অসাধারণ মূল্য প্রদান করে। এটি তীক্ষ্ণ হ্যান্ডলিং এবং চপলতা সহ স্পোর্টি শহুরে যাতায়াতে শ্রেষ্ঠ। তবে, অয়েল-কুলড ইঞ্জিন কম্পন, শক্ত সাসপেনশন, গড় বিল্ড কোয়ালিটি, উৎসাহীদের জন্য সীমিত ১৬০সিসি শক্তি এবং মৌলিক ডিসপ্লে এটিকে তরুণ শহুরে রাইডার, নিরাপত্তা-সচেতন ক্রেতা এবং প্রিমিয়াম পরিমার্জনের চেয়ে ডুয়াল এবিএস নিরাপত্তা, আধুনিক স্টাইলিং এবং মূল্যকে অগ্রাধিকার দেওয়া দৈনিক যাত্রীদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বিরল ডুয়াল-চ্যানেল এবিএস নিরাপত্তা সহ অসাধারণ মূল্য ১৬০সিসির জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for New Bajaj Pulsar N160 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for New Bajaj Pulsar N160\n")

	return nil
}
