package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaDioReviewSeeder handles seeding Honda Dio product review and translation
type HondaDioReviewSeeder struct {
	BaseSeeder
}

// NewHondaDioReviewSeeder creates a new HondaDioReviewSeeder
func NewHondaDioReviewSeeder() *HondaDioReviewSeeder {
	return &HondaDioReviewSeeder{BaseSeeder: BaseSeeder{name: "Honda Dio Review"}}
}

// Seed implements the Seeder interface
func (s *HondaDioReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda Dio").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda Dio product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda Dio product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda Dio already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda Dio Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda Dio is a stylish 110cc scooter priced at ৳1,40,000, perfect for urban commuting. With automatic transmission, under-seat storage, and Honda's renowned reliability, it's an ideal choice for riders seeking convenience and comfort in city traffic.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda Dio Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Automatic Transmission:</strong> <span class="highlight-value">CVT automatic, no gear shifting</span></li>
<li class="highlight-item"><strong class="highlight-label">Under-Seat Storage:</strong> <span class="highlight-value">Large storage space under seat</span></li>
<li class="highlight-item"><strong class="highlight-label">Honda Reliability:</strong> <span class="highlight-value">Proven Japanese quality</span></li>
<li class="highlight-item"><strong class="highlight-label">Stylish Design:</strong> <span class="highlight-value">Modern sporty scooter look</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda Dio Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Automatic CVT:</strong> <span class="pro-description">No clutch or gears, very easy to ride</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Build Quality:</strong> <span class="pro-description">Premium Honda quality and durability</span></li>
<li class="pro-item"><strong class="pro-label">Good Storage Space:</strong> <span class="pro-description">Can fit full-face helmet under seat</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">110cc Honda engine is refined and quiet</span></li>
<li class="pro-item"><strong class="pro-label">Great Maneuverability:</strong> <span class="pro-description">Small turning radius, easy in traffic</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">50-55 km/l fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seat:</strong> <span class="pro-description">Wide flat seat for comfort</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Honda service centers everywhere</span></li>
<li class="pro-item"><strong class="pro-label">Strong Resale Value:</strong> <span class="pro-description">Honda scooters hold value well</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda Dio Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">৳1,40,000 expensive for 110cc scooter</span></li>
<li class="con-item"><strong class="con-label">Small Wheels:</strong> <span class="con-description">10-inch wheels struggle on bad roads</span></li>
<li class="con-item"><strong class="con-label">Drum Brakes Only:</strong> <span class="con-description">No disc brake option available</span></li>
<li class="con-item"><strong class="con-label">Limited Power:</strong> <span class="con-description">110cc not great for highway riding</span></li>
<li class="con-item"><strong class="con-label">Lightweight Issues:</strong> <span class="con-description">Feels unstable in strong winds</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda Dio in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban daily commuters</li>
<li class="best-for-item">Students and college goers</li>
<li class="best-for-item">Women riders seeking easy handling</li>
<li class="best-for-item">First-time scooter buyers</li>
<li class="best-for-item">Short-distance city travel</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda Dio?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway and long-distance riders</li>
<li class="not-recommended-item">Riders on very rough roads</li>
<li class="not-recommended-item">Budget-constrained buyers</li>
<li class="not-recommended-item">Those needing powerful acceleration</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda Dio Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,35,000 - ৳1,45,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,000 - ৳2,500 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,400-1,700/month for 30 km daily at 52 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳700 - ৳1,000 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda Dio Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Top Honda quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Premium but worth it</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Adequate for city</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Excellent quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Sporty modern look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda Dio Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Honda Dio?</h3>
<p class="faq-answer">Honda Dio delivers 50-55 km/l in city conditions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Honda Dio in Bangladesh?</h3>
<p class="faq-answer">Honda Dio is priced between ৳1,35,000 to ৳1,45,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Honda Dio good for ladies?</h3>
<p class="faq-answer">Yes, Honda Dio is perfect for ladies with automatic transmission and easy handling.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Honda Dio have disc brake?</h3>
<p class="faq-answer">No, Honda Dio comes with drum brakes on both wheels.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda Dio in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Honda Dio is the perfect urban scooter for those prioritizing convenience and reliability. At ৳1,40,000, it's premium-priced but offers unmatched Honda quality, automatic transmission, and excellent mileage. Best for city commuters, students, and women riders who value ease of use over raw performance. Strong resale value and wide service network make it a smart long-term investment despite the higher initial cost.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For urban commuters seeking convenience</span></p>
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
		return fmt.Errorf("error creating Honda Dio review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda Dio (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হন্ডা ডায়ো রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হন্ডা ডায়ো একটি স্টাইলিশ ১১০সিসি স্কুটার যার মূল্য ৳১,৪০,০০০ টাকা, শহুরে যাতায়াতের জন্য নিখুঁত। অটোমেটিক ট্রান্সমিশন, আন্ডার-সিট স্টোরেজ এবং হোন্ডার বিখ্যাত নির্ভরযোগ্যতা সহ, এটি শহরের ট্রাফিকে সুবিধা এবং আরাম খুঁজছেন রাইডারদের জন্য একটি আদর্শ পছন্দ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হন্ডা ডায়ো এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">অটোমেটিক ট্রান্সমিশন:</strong> <span class="highlight-value">সিভিটি অটোমেটিক, কোন গিয়ার শিফটিং নেই</span></li>
<li class="highlight-item"><strong class="highlight-label">আন্ডার-সিট স্টোরেজ:</strong> <span class="highlight-value">সিটের নিচে বড় স্টোরেজ স্পেস</span></li>
<li class="highlight-item"><strong class="highlight-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="highlight-value">প্রমাণিত জাপানি মান</span></li>
<li class="highlight-item"><strong class="highlight-label">স্টাইলিশ ডিজাইন:</strong> <span class="highlight-value">আধুনিক স্পোর্টি স্কুটার লুক</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হন্ডা ডায়ো এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অটোমেটিক সিভিটি:</strong> <span class="pro-description">কোন ক্লাচ বা গিয়ার নেই, রাইড করা খুব সহজ</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম হোন্ডা মান এবং স্থায়িত্ব</span></li>
<li class="pro-item"><strong class="pro-label">ভালো স্টোরেজ স্পেস:</strong> <span class="pro-description">সিটের নিচে ফুল-ফেস হেলমেট ফিট করতে পারে</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">১১০সিসি হোন্ডা ইঞ্জিন পরিশীলিত এবং শান্ত</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত চলাফেরা:</strong> <span class="pro-description">ছোট টার্নিং ব্যাসার্ধ, ট্রাফিকে সহজ</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">৫০-৫৫ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক সিট:</strong> <span class="pro-description">আরামের জন্য চওড়া ফ্ল্যাট সিট</span></li>
<li class="pro-item"><strong class="pro-label">ব্যাপক সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সর্বত্র হোন্ডা সার্ভিস সেন্টার</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী রিসেল ভ্যালু:</strong> <span class="pro-description">হোন্ডা স্কুটার ভালো মূল্য ধরে রাখে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হন্ডা ডায়ো এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">১১০সিসি স্কুটারের জন্য ৳১,৪০,০০০ টাকা ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">ছোট চাকা:</strong> <span class="con-description">১০-ইঞ্চি চাকা খারাপ রাস্তায় সংগ্রাম করে</span></li>
<li class="con-item"><strong class="con-label">শুধুমাত্র ড্রাম ব্রেক:</strong> <span class="con-description">কোন ডিস্ক ব্রেক বিকল্প নেই</span></li>
<li class="con-item"><strong class="con-label">সীমিত পাওয়ার:</strong> <span class="con-description">হাইওয়ে রাইডিংয়ের জন্য ১১০সিসি ভালো নয়</span></li>
<li class="con-item"><strong class="con-label">হালকা ওজনের সমস্যা:</strong> <span class="con-description">শক্তিশালী বাতাসে অস্থিতিশীল মনে হয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হন্ডা ডায়ো কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban daily commuters</li>
<li class="best-for-item">Students and college goers</li>
<li class="best-for-item">Women riders seeking easy handling</li>
<li class="best-for-item">First-time scooter buyers</li>
<li class="best-for-item">Short-distance city travel</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হন্ডা ডায়ো কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway and long-distance riders</li>
<li class="not-recommended-item">Riders on very rough roads</li>
<li class="not-recommended-item">Budget-constrained buyers</li>
<li class="not-recommended-item">Those needing powerful acceleration</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হন্ডা ডায়ো এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৩৫,০০০ - ৳১,৪৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,০০০ - ৳২,৫০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৪০০-১,৭০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৭০০ - ৳১,০০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হন্ডা ডায়ো পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- শীর্ষ হোন্ডা মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- প্রিমিয়াম কিন্তু মূল্যবান</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- শহরের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- খুব আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- চমৎকার মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- স্পোর্টি আধুনিক লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হন্ডা ডায়ো সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হন্ডা ডায়োর মাইলেজ কত?</h3>
<p class="faq-answer">হন্ডা ডায়ো শহরের পরিস্থিতিতে ৫০-৫৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে হন্ডা ডায়োর দাম কত?</h3>
<p class="faq-answer">হন্ডা ডায়োর দাম ৳১,৩৫,০০০ থেকে ৳১,৪৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">মহিলাদের জন্য হন্ডা ডায়ো কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, হন্ডা ডায়ো অটোমেটিক ট্রান্সমিশন এবং সহজ হ্যান্ডলিং সহ মহিলাদের জন্য নিখুঁত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হন্ডা ডায়োতে কি ডিস্ক ব্রেক আছে?</h3>
<p class="faq-answer">না, হন্ডা ডায়ো উভয় চাকায় ড্রাম ব্রেক সহ আসে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হন্ডা ডায়ো কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হন্ডা ডায়ো সুবিধা এবং নির্ভরযোগ্যতাকে অগ্রাধিকার দেন তাদের জন্য নিখুঁত শহুরে স্কুটার। ৳১,৪০,০০০ টাকায়, এটি প্রিমিয়াম-মূল্য কিন্তু অতুলনীয় হোন্ডা মান, অটোমেটিক ট্রান্সমিশন এবং চমৎকার মাইলেজ প্রদান করে। শহর যাত্রী, ছাত্র এবং মহিলা রাইডারদের জন্য সেরা যারা কাঁচা পারফরম্যান্সের চেয়ে ব্যবহারের সহজতাকে মূল্য দেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সুবিধা খুঁজছেন শহর যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda Dio already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda Dio\n")

	return nil
}
