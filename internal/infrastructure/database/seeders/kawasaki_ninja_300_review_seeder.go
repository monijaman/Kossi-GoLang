package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// KawasakiNinja300ReviewSeeder handles seeding Kawasaki Ninja 300 product review and translation
type KawasakiNinja300ReviewSeeder struct {
	BaseSeeder
}

// NewKawasakiNinja300ReviewSeeder creates a new KawasakiNinja300ReviewSeeder
func NewKawasakiNinja300ReviewSeeder() *KawasakiNinja300ReviewSeeder {
	return &KawasakiNinja300ReviewSeeder{BaseSeeder: BaseSeeder{name: "Kawasaki Ninja 300 Review"}}
}

// Seed implements the Seeder interface
func (s *KawasakiNinja300ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Kawasaki Ninja 300").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Kawasaki Ninja 300 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Kawasaki Ninja 300 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Kawasaki Ninja 300 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Kawasaki Ninja 300 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Kawasaki Ninja 300 is a premium 296cc sports bike priced at ৳12,00,000, offering true supersport experience, exceptional build quality, and track-ready performance. With 39 HP power, twin-cylinder engine, and legendary Ninja styling, it's the ultimate choice for serious sports bike enthusiasts seeking authentic supersport thrills.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Kawasaki Ninja 300 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Twin Cylinder Power:</strong> <span class="highlight-value">39 HP parallel twin engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Supersport Design:</strong> <span class="highlight-value">Authentic Ninja styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Features:</strong> <span class="highlight-value">ABS, slipper clutch, digital console</span></li>
<li class="highlight-item"><strong class="highlight-label">Track Ready:</strong> <span class="highlight-value">165+ km/h top speed</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Kawasaki Ninja 300 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Incredible Power:</strong> <span class="pro-description">39 HP twin cylinder beast</span></li>
<li class="pro-item"><strong class="pro-label">Authentic Supersport:</strong> <span class="pro-description">True sports bike experience</span></li>
<li class="pro-item"><strong class="pro-label">Legendary Build:</strong> <span class="pro-description">Japanese premium quality</span></li>
<li class="pro-item"><strong class="pro-label">Slipper Clutch:</strong> <span class="pro-description">Advanced racing technology</span></li>
<li class="pro-item"><strong class="pro-label">Dual Channel ABS:</strong> <span class="pro-description">Superior braking safety</span></li>
<li class="pro-item"><strong class="pro-label">Track Performance:</strong> <span class="pro-description">Race-ready out of box</span></li>
<li class="pro-item"><strong class="pro-label">Iconic Ninja Looks:</strong> <span class="pro-description">Unmistakable supersport styling</span></li>
<li class="pro-item"><strong class="pro-label">High Top Speed:</strong> <span class="pro-description">165+ km/h capability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Kawasaki Ninja 300 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Extremely Expensive:</strong> <span class="con-description">₹12,00,000 premium pricing</span></li>
<li class="con-item"><strong class="con-label">Very Poor Mileage:</strong> <span class="con-description">25-30 km/l, very thirsty</span></li>
<li class="con-item"><strong class="con-label">Very Limited Service:</strong> <span class="con-description">Almost no Kawasaki centers</span></li>
<li class="con-item"><strong class="con-label">Extremely Expensive Parts:</strong> <span class="con-description">Sky-high maintenance costs</span></li>
<li class="con-item"><strong class="con-label">Very Uncomfortable:</strong> <span class="con-description">Extreme sports position</span></li>
<li class="con-item"><strong class="con-label">Only for Experts:</strong> <span class="con-description">Way too powerful for beginners</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Kawasaki Ninja 300 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Expert riders with deep pockets</li>
<li class="best-for-item">Track day enthusiasts</li>
<li class="best-for-item">Supersport purists</li>
<li class="best-for-item">Those wanting authentic Ninja experience</li>
<li class="best-for-item">Performance over everything riders</li>
<li class="best-for-item">Show-off premium buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Kawasaki Ninja 300?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Beginners or intermediate riders</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily commuters</li>
<li class="not-recommended-item">Those prioritizing comfort</li>
<li class="not-recommended-item">Areas without Kawasaki service</li>
<li class="not-recommended-item">Riders wanting practical bike</li>
<li class="not-recommended-item">Fuel economy seekers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Kawasaki Ninja 300 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳11,80,000 - ৳12,20,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳8,000 - ৳12,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳4,500-6,000/month for 30 km daily at 27 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳5,000 - ৳8,000 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Kawasaki Ninja 300 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very poor mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Kawasaki quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Premium but authentic</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Exceptional</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Racing position</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Japanese excellence</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Premium supersport</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Iconic Ninja</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Kawasaki Ninja 300 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Kawasaki Ninja 300?</h3>
<p class="faq-answer">Kawasaki Ninja 300 can reach 165+ km/h with 39 HP twin cylinder.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Kawasaki Ninja 300?</h3>
<p class="faq-answer">Kawasaki Ninja 300 is priced between ৳11,80,000 to ৳12,20,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Ninja 300 good for beginners?</h3>
<p class="faq-answer">Absolutely not! Ninja 300 is way too powerful for beginners. Start with 150cc bikes.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Ninja 300?</h3>
<p class="faq-answer">Ninja 300 delivers 25-30 km/l with twin cylinder performance engine.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Kawasaki Ninja 300 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Kawasaki Ninja 300 at ৳12,00,000 is the ULTIMATE supersport for serious enthusiasts with deep pockets, offering authentic 39 HP twin cylinder power, legendary Ninja pedigree, and true track-ready performance. With slipper clutch, dual channel ABS, and 165+ km/h capability, it's the real deal for experienced riders seeking genuine supersport thrills. However, extremely high price (3.5x more than Duke 200), very poor mileage (25-30 km/l), sky-high maintenance costs, racing-only comfort, and almost non-existent service network make it suitable only for wealthy expert riders who prioritize authentic supersport experience over everything else. If money is no object and you're an expert rider, this is supersport perfection.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For authentic supersport experience with unlimited budget</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.7,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Kawasaki Ninja 300 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Kawasaki Ninja 300 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">কাওয়াসাকি নিনজা ৩০০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">কাওয়াসাকি নিনজা ৩০০ একটি প্রিমিয়াম ২৯৬সিসি স্পোর্টস বাইক যার মূল্য ৳১২,০০,০০০ টাকা, যা সত্যিকারের সুপারস্পোর্ট অভিজ্ঞতা, ব্যতিক্রমী বিল্ড কোয়ালিটি এবং ট্র্যাক-রেডি পারফরম্যান্স প্রদান করে। ৩৯ এইচপি পাওয়ার, টুইন-সিলিন্ডার ইঞ্জিন এবং কিংবদন্তী নিনজা স্টাইলিং সহ, এটি প্রকৃত সুপারস্পোর্ট রোমাঞ্চ খোঁজা গুরুতর স্পোর্টস বাইক উৎসাহীদের জন্য চূড়ান্ত পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কাওয়াসাকি নিনজা ৩০০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">টুইন সিলিন্ডার পাওয়ার:</strong> <span class="highlight-value">৩৯ এইচপি প্যারালেল টুইন ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">সুপারস্পোর্ট ডিজাইন:</strong> <span class="highlight-value">প্রকৃত নিনজা স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম ফিচার:</strong> <span class="highlight-value">এবিএস, স্লিপার ক্লাচ, ডিজিটাল কনসোল</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্র্যাক রেডি:</strong> <span class="highlight-value">১৬৫+ কিমি/ঘণ্টা টপ স্পিড</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কাওয়াসাকি নিনজা ৩০০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অবিশ্বাস্য পাওয়ার:</strong> <span class="pro-description">৩৯ এইচপি টুইন সিলিন্ডার দানব</span></li>
<li class="pro-item"><strong class="pro-label">প্রকৃত সুপারস্পোর্ট:</strong> <span class="pro-description">সত্যিকারের স্পোর্টস বাইক অভিজ্ঞতা</span></li>
<li class="pro-item"><strong class="pro-label">কিংবদন্তী বিল্ড:</strong> <span class="pro-description">জাপানি প্রিমিয়াম মান</span></li>
<li class="pro-item"><strong class="pro-label">স্লিপার ক্লাচ:</strong> <span class="pro-description">উন্নত রেসিং প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল চ্যানেল এবিএস:</strong> <span class="pro-description">উৎকৃষ্ট ব্রেকিং নিরাপত্তা</span></li>
<li class="pro-item"><strong class="pro-label">ট্র্যাক পারফরম্যান্স:</strong> <span class="pro-description">বক্সের বাইরে রেস-রেডি</span></li>
<li class="pro-item"><strong class="pro-label">আইকনিক নিনজা লুক:</strong> <span class="pro-description">অবিস্মরণীয় সুপারস্পোর্ট স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">উচ্চ টপ স্পিড:</strong> <span class="pro-description">১৬৫+ কিমি/ঘণ্টা সক্ষমতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কাওয়াসাকি নিনজা ৩০০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত ব্যয়বহুল:</strong> <span class="con-description">৳১২,০০,০০০ প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">অত্যন্ত খারাপ মাইলেজ:</strong> <span class="con-description">২৫-৩০ কিমি/লিটার, অত্যন্ত জ্বালানি ক্ষুধার্ত</span></li>
<li class="con-item"><strong class="con-label">অত্যন্ত সীমিত সার্ভিস:</strong> <span class="con-description">প্রায় কোনো কাওয়াসাকি সেন্টার নেই</span></li>
<li class="con-item"><strong class="con-label">অত্যন্ত ব্যয়বহুল যন্ত্রাংশ:</strong> <span class="con-description">আকাশছোঁয়া রক্ষণাবেক্ষণ খরচ</span></li>
<li class="con-item"><strong class="con-label">অত্যন্ত অস্বস্তিকর:</strong> <span class="con-description">চরম স্পোর্টস অবস্থান</span></li>
<li class="con-item"><strong class="con-label">শুধু বিশেষজ্ঞদের জন্য:</strong> <span class="con-description">শিক্ষানবিসদের জন্য অত্যধিক শক্তিশালী</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে কাওয়াসাকি নিনজা ৩০০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Expert riders with deep pockets</li>
<li class="best-for-item">Track day enthusiasts</li>
<li class="best-for-item">Supersport purists</li>
<li class="best-for-item">Those wanting authentic Ninja experience</li>
<li class="best-for-item">Performance over everything riders</li>
<li class="best-for-item">Show-off premium buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">কাওয়াসাকি নিনজা ৩০০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Beginners or intermediate riders</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Daily commuters</li>
<li class="not-recommended-item">Those prioritizing comfort</li>
<li class="not-recommended-item">Areas without Kawasaki service</li>
<li class="not-recommended-item">Riders wanting practical bike</li>
<li class="not-recommended-item">Fuel economy seekers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">কাওয়াসাকি নিনজা ৩০০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১১,৮০,০০০ - ৳১২,২০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৮,০০০ - ৳১২,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">২৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳৪,৫০০-৬,০০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৫,০০০ - ৳৮,০০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">কাওয়াসাকি নিনজা ৩০০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- অত্যন্ত খারাপ মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- কাওয়াসাকি মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- প্রিমিয়াম কিন্তু প্রকৃত</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- ব্যতিক্রমী</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- রেসিং অবস্থান</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- জাপানি শ্রেষ্ঠত্ব</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- প্রিমিয়াম সুপারস্পোর্ট</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- আইকনিক নিনজা</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">কাওয়াসাকি নিনজা ৩০০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">কাওয়াসাকি নিনজা ৩০০ এর টপ স্পিড কত?</h3>
<p class="faq-answer">কাওয়াসাকি নিনজা ৩০০ ৩৯ এইচপি টুইন সিলিন্ডার সহ ১৬৫+ কিমি/ঘণ্টা পৌঁছাতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কাওয়াসাকি নিনজা ৩০০ এর দাম কত?</h3>
<p class="faq-answer">কাওয়াসাকি নিনজা ৩০০ এর দাম ৳১১,৮০,০০০ থেকে ৳১২,২০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">শিক্ষানবিসদের জন্য নিনজা ৩০০ কি ভালো?</h3>
<p class="faq-answer">একেবারেই না! নিনজা ৩০০ শিক্ষানবিসদের জন্য অত্যধিক শক্তিশালী। ১৫০সিসি বাইক দিয়ে শুরু করুন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">নিনজা ৩০০ এর মাইলেজ কত?</h3>
<p class="faq-answer">নিনজা ৩০০ টুইন সিলিন্ডার পারফরম্যান্স ইঞ্জিন সহ ২৫-৩০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে কাওয়াসাকি নিনজা ৩০০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳১২,০০,০০০ টাকায় কাওয়াসাকি নিনজা ৩০০ গভীর পকেটের সাথে গুরুতর উৎসাহীদের জন্য চূড়ান্ত সুপারস্পোর্ট, যা প্রকৃত ৩৯ এইচপি টুইন সিলিন্ডার পাওয়ার, কিংবদন্তী নিনজা ঐতিহ্য এবং সত্যিকারের ট্র্যাক-রেডি পারফরম্যান্স প্রদান করে। স্লিপার ক্লাচ, ডুয়াল চ্যানেল এবিএস এবং ১৬৫+ কিমি/ঘণ্টা সক্ষমতা সহ, এটি প্রকৃত সুপারস্পোর্ট রোমাঞ্চ খোঁজা অভিজ্ঞ রাইডারদের জন্য আসল জিনিস। তবে, অত্যন্ত উচ্চ মূল্য (ডিউক ২০০ এর চেয়ে ৩.৫ গুণ বেশি), অত্যন্ত খারাপ মাইলেজ (২৫-৩০ কিমি/লিটার), আকাশছোঁয়া রক্ষণাবেক্ষণ খরচ, রেসিং-শুধু আরাম এবং প্রায় অস্তিত্বহীন সার্ভিস নেটওয়ার্ক এটিকে শুধুমাত্র ধনী বিশেষজ্ঞ রাইডারদের জন্য উপযুক্ত করে যারা অন্য সবকিছুর চেয়ে প্রকৃত সুপারস্পোর্ট অভিজ্ঞতাকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সীমাহীন বাজেট সহ প্রকৃত সুপারস্পোর্ট অভিজ্ঞতার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Kawasaki Ninja 300 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Kawasaki Ninja 300\n")

	return nil
}
