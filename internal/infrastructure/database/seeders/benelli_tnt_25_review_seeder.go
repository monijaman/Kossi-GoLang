package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BenelliTNT25ReviewSeeder handles seeding Benelli TNT 25 product review and translation
type BenelliTNT25ReviewSeeder struct {
	BaseSeeder
}

// NewBenelliTNT25ReviewSeeder creates a new BenelliTNT25ReviewSeeder
func NewBenelliTNT25ReviewSeeder() *BenelliTNT25ReviewSeeder {
	return &BenelliTNT25ReviewSeeder{BaseSeeder: BaseSeeder{name: "Benelli TNT 25 Review"}}
}

// Seed implements the Seeder interface
func (s *BenelliTNT25ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Benelli TNT 25").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Benelli TNT 25 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Benelli TNT 25 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Benelli TNT 25 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Benelli TNT 25 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Benelli TNT 25 is a premium 249cc naked street bike priced at ৳4,80,000, offering Italian styling, good power, and distinctive design. With 25.8 HP, liquid cooling, and aggressive looks, it's for enthusiasts seeking unique Italian character without extreme pricing, though service network limitations apply.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Benelli TNT 25 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Italian Design:</strong> <span class="highlight-value">Distinctive Italian styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Power:</strong> <span class="highlight-value">25.8 HP from 249cc twin</span></li>
<li class="highlight-item"><strong class="highlight-label">Liquid Cooled:</strong> <span class="highlight-value">Advanced cooling system</span></li>
<li class="highlight-item"><strong class="highlight-label">Naked Street Style:</strong> <span class="highlight-value">Aggressive naked bike design</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Benelli TNT 25 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Italian Character:</strong> <span class="pro-description">Unique Benelli Italian styling</span></li>
<li class="pro-item"><strong class="pro-label">Twin Cylinder:</strong> <span class="pro-description">25.8 HP parallel twin engine</span></li>
<li class="pro-item"><strong class="pro-label">Liquid Cooling:</strong> <span class="pro-description">Efficient liquid cooling</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Distinctive street fighter looks</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">Peppy 250cc performance</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Upright riding position</span></li>
<li class="pro-item"><strong class="pro-label">ABS Safety:</strong> <span class="pro-description">Dual channel ABS</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Benelli TNT 25 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Expensive:</strong> <span class="con-description">₹4,80,000 premium pricing</span></li>
<li class="con-item"><strong class="con-label">Very Limited Service:</strong> <span class="con-description">Almost no Benelli service centers</span></li>
<li class="con-item"><strong class="con-label">Poor Mileage:</strong> <span class="con-description">35-40 km/l, twin cylinder thirst</span></li>
<li class="con-item"><strong class="con-label">Expensive Parts:</strong> <span class="con-description">High maintenance costs</span></li>
<li class="con-item"><strong class="con-label">Chinese Assembly:</strong> <span class="con-description">Made in China (Italian brand)</span></li>
<li class="con-item"><strong class="con-label">Poor Resale:</strong> <span class="con-description">Low resale value</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Benelli TNT 25 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Italian bike enthusiasts</li>
<li class="best-for-item">Those wanting unique styling</li>
<li class="best-for-item">Naked bike lovers</li>
<li class="best-for-item">Performance seekers on budget</li>
<li class="best-for-item">City riding with style</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Benelli TNT 25?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Service-conscious buyers</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Budget-tight buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Buyers wanting good resale</li>
<li class="not-recommended-item">Areas without Benelli service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Benelli TNT 25 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,60,000 - ৳5,00,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳4,500 - ৳6,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,400-3,000/month for 30 km daily at 37 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳2,000 - ৳3,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Benelli TNT 25 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Average 37 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Decent reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good for Italian bike</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Good 25.8 HP twin</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Comfortable upright</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Chinese assembly</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Unique Italian design</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Benelli TNT 25 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Benelli TNT 25?</h3>
<p class="faq-answer">Benelli TNT 25 delivers 35-40 km/l with 249cc twin cylinder engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Benelli TNT 25?</h3>
<p class="faq-answer">Benelli TNT 25 is priced between ৳4,60,000 to ৳5,00,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Benelli TNT 25 worth buying?</h3>
<p class="faq-answer">Yes, for unique Italian styling and twin cylinder experience, but consider limited service network.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Benelli TNT 25 made in China?</h3>
<p class="faq-answer">Yes, Benelli is Italian brand but bikes are assembled in China with quality control.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Benelli TNT 25 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Benelli TNT 25 at ৳4,80,000 offers UNIQUE Italian character for enthusiasts seeking distinctive styling and twin cylinder experience without extreme pricing. With 25.8 HP parallel twin, liquid cooling, aggressive naked design, dual channel ABS, and comfortable ergonomics, it stands out in the 250cc segment. However, very limited service network, poor mileage (35-40 km/l), expensive maintenance, Chinese assembly quality concerns, and poor resale value are significant drawbacks. Ideal for style-conscious riders who prioritize unique Italian design and twin cylinder character over service convenience and fuel economy. Not for those seeking practicality or reliable service support.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For unique Italian styling and twin cylinder experience</span></p>
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
		return fmt.Errorf("error creating Benelli TNT 25 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Benelli TNT 25 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বেনেলি টিএনটি ২৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বেনেলি টিএনটি ২৫ একটি প্রিমিয়াম ২৪৯সিসি নেকেড স্ট্রিট বাইক যার মূল্য ৳৪,৮০,০০০ টাকা, যা ইতালিয়ান স্টাইলিং, ভালো পাওয়ার এবং স্বতন্ত্র ডিজাইন প্রদান করে। ২৫.৮ এইচপি, লিকুইড কুলিং এবং আক্রমণাত্মক লুক সহ, এটি চরম মূল্য ছাড়াই অনন্য ইতালিয়ান চরিত্র খোঁজা উৎসাহীদের জন্য, যদিও সার্ভিস নেটওয়ার্ক সীমাবদ্ধতা প্রযোজ্য।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বেনেলি টিএনটি ২৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ইতালিয়ান ডিজাইন:</strong> <span class="highlight-value">স্বতন্ত্র ইতালিয়ান স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো পাওয়ার:</strong> <span class="highlight-value">২৪৯সিসি টুইন থেকে ২৫.৮ এইচপি</span></li>
<li class="highlight-item"><strong class="highlight-label">লিকুইড কুলড:</strong> <span class="highlight-value">উন্নত কুলিং সিস্টেম</span></li>
<li class="highlight-item"><strong class="highlight-label">নেকেড স্ট্রিট স্টাইল:</strong> <span class="highlight-value">আক্রমণাত্মক নেকেড বাইক ডিজাইন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বেনেলি টিএনটি ২৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ইতালিয়ান চরিত্র:</strong> <span class="pro-description">অনন্য বেনেলি ইতালিয়ান স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">টুইন সিলিন্ডার:</strong> <span class="pro-description">২৫.৮ এইচপি প্যারালেল টুইন ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">লিকুইড কুলিং:</strong> <span class="pro-description">দক্ষ লিকুইড কুলিং</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">স্বতন্ত্র স্ট্রিট ফাইটার লুক</span></li>
<li class="pro-item"><strong class="pro-label">ভালো পারফরম্যান্স:</strong> <span class="pro-description">প্রাণবন্ত ২৫০সিসি পারফরম্যান্স</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এরগোনমিক্স:</strong> <span class="pro-description">সোজা রাইডিং অবস্থান</span></li>
<li class="pro-item"><strong class="pro-label">এবিএস নিরাপত্তা:</strong> <span class="pro-description">ডুয়াল চ্যানেল এবিএস</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বেনেলি টিএনটি ২৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">ব্যয়বহুল:</strong> <span class="con-description">৳৪,৮০,০০০ প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">অত্যন্ত সীমিত সার্ভিস:</strong> <span class="con-description">প্রায় কোনো বেনেলি সার্ভিস সেন্টার নেই</span></li>
<li class="con-item"><strong class="con-label">খারাপ মাইলেজ:</strong> <span class="con-description">৩৫-৪০ কিমি/লিটার, টুইন সিলিন্ডার তৃষ্ণা</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল যন্ত্রাংশ:</strong> <span class="con-description">উচ্চ রক্ষণাবেক্ষণ খরচ</span></li>
<li class="con-item"><strong class="con-label">চীনা সমাবেশ:</strong> <span class="con-description">চীনে তৈরি (ইতালিয়ান ব্র্যান্ড)</span></li>
<li class="con-item"><strong class="con-label">খারাপ রিসেল:</strong> <span class="con-description">কম রিসেল ভ্যালু</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বেনেলি টিএনটি ২৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Italian bike enthusiasts</li>
<li class="best-for-item">Those wanting unique styling</li>
<li class="best-for-item">Naked bike lovers</li>
<li class="best-for-item">Performance seekers on budget</li>
<li class="best-for-item">City riding with style</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বেনেলি টিএনটি ২৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Service-conscious buyers</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Budget-tight buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Buyers wanting good resale</li>
<li class="not-recommended-item">Areas without Benelli service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বেনেলি টিএনটি ২৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৪,৬০,০০০ - ৳৫,০০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৪,৫০০ - ৳৬,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৩৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,৪০০-৩,০০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳২,০০০ - ৳৩,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বেনেলি টিএনটি ২৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড় ৩৭ কিমি/লিটার</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- মোটামুটি নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ইতালিয়ান বাইকের জন্য ভালো</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- ভালো ২৫.৮ এইচপি টুইন</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আরামদায়ক সোজা</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- চীনা সমাবেশ</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- অনন্য ইতালিয়ান ডিজাইন</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বেনেলি টিএনটি ২৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">বেনেলি টিএনটি ২৫ এর মাইলেজ কত?</h3>
<p class="faq-answer">বেনেলি টিএনটি ২৫ ২৪৯সিসি টুইন সিলিন্ডার ইঞ্জিন সহ ৩৫-৪০ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বেনেলি টিএনটি ২৫ এর দাম কত?</h3>
<p class="faq-answer">বেনেলি টিএনটি ২৫ এর দাম ৳৪,৬০,০০০ থেকে ৳৫,০০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বেনেলি টিএনটি ২৫ কেনার যোগ্য?</h3>
<p class="faq-answer">হ্যাঁ, অনন্য ইতালিয়ান স্টাইলিং এবং টুইন সিলিন্ডার অভিজ্ঞতার জন্য, কিন্তু সীমিত সার্ভিস নেটওয়ার্ক বিবেচনা করুন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বেনেলি টিএনটি ২৫ কি চীনে তৈরি?</h3>
<p class="faq-answer">হ্যাঁ, বেনেলি ইতালিয়ান ব্র্যান্ড কিন্তু বাইকগুলি কোয়ালিটি কন্ট্রোল সহ চীনে সমাবেশ করা হয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বেনেলি টিএনটি ২৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৪,৮০,০০০ টাকায় বেনেলি টিএনটি ২৫ চরম মূল্য ছাড়াই স্বতন্ত্র স্টাইলিং এবং টুইন সিলিন্ডার অভিজ্ঞতা খোঁজা উৎসাহীদের জন্য অনন্য ইতালিয়ান চরিত্র প্রদান করে। ২৫.৮ এইচপি প্যারালেল টুইন, লিকুইড কুলিং, আক্রমণাত্মক নেকেড ডিজাইন, ডুয়াল চ্যানেল এবিএস এবং আরামদায়ক এরগোনমিক্স সহ, এটি ২৫০সিসি সেগমেন্টে আলাদা হয়ে দাঁড়ায়। তবে, অত্যন্ত সীমিত সার্ভিস নেটওয়ার্ক, খারাপ মাইলেজ (৩৫-৪০ কিমি/লিটার), ব্যয়বহুল রক্ষণাবেক্ষণ, চীনা সমাবেশ মানের উদ্বেগ এবং খারাপ রিসেল ভ্যালু উল্লেখযোগ্য ত্রুটি।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অনন্য ইতালিয়ান স্টাইলিং এবং টুইন সিলিন্ডার অভিজ্ঞতার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Benelli TNT 25 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Benelli TNT 25\n")

	return nil
}
