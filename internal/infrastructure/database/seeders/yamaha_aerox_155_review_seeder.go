package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaAerox155Review handles seeding Yamaha Aerox 155 product review and translation
type YamahaAerox155Review struct {
	BaseSeeder
}

// NewYamahaAerox155ReviewSeeder creates a new YamahaAerox155Review
func NewYamahaAerox155ReviewSeeder() *YamahaAerox155Review {
	return &YamahaAerox155Review{BaseSeeder: BaseSeeder{name: "Yamaha Aerox 155 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaAerox155Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha Aerox 155").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha Aerox 155 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha Aerox 155 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha Aerox 155 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha Aerox 155 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha Aerox 155 is a premium 155cc sports maxi-scooter featuring VVA technology, aerodynamic design, and advanced features. Priced at ৳5,25,000, it combines the convenience of automatic transmission with sports bike performance including variable valve actuation, sporty ergonomics, large storage capacity, and Yamaha's proven engineering for riders seeking premium scooter experience with motorcycle-level performance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha Aerox 155 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Sports Maxi-Scooter Design:</strong> <span class="highlight-value">Aggressive aerodynamic styling with sports appeal</span></li>
<li class="highlight-item"><strong class="highlight-label">VVA Technology:</strong> <span class="highlight-value">Variable valve actuation for optimal performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Automatic CVT Transmission:</strong> <span class="highlight-value">Seamless power delivery without clutch</span></li>
<li class="highlight-item"><strong class="highlight-label">Large Storage Capacity:</strong> <span class="highlight-value">Practical underseat storage for daily use</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha Aerox 155 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">VVA Engine Performance:</strong> <span class="pro-description">Variable valve technology delivers strong performance for a scooter</span></li>
<li class="pro-item"><strong class="pro-label">Automatic Convenience:</strong> <span class="pro-description">CVT transmission eliminates clutch operation and gear shifting</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Styling:</strong> <span class="pro-description">Aggressive maxi-scooter design with motorcycle-inspired aesthetics</span></li>
<li class="pro-item"><strong class="pro-label">Practical Storage:</strong> <span class="pro-description">Large underseat compartment for helmets and belongings</span></li>
<li class="pro-item"><strong class="pro-label">Weather Protection:</strong> <span class="pro-description">Better protection from elements compared to naked motorcycles</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Efficiency:</strong> <span class="pro-description">Good fuel economy with VVA and CVT optimization</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha Aerox 155 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Purchase Price:</strong> <span class="con-description">Premium pricing for a 155cc scooter</span></li>
<li class="con-item"><strong class="con-label">Limited Sporty Feel:</strong> <span class="con-description">CVT transmission lacks the engagement of manual gearbox</span></li>
<li class="con-item"><strong class="con-label">Maintenance Complexity:</strong> <span class="con-description">VVA technology and CVT require specialized service</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer dealerships experienced with premium scooter service</span></li>
<li class="con-item"><strong class="con-label">Highway Limitations:</strong> <span class="con-description">Scooter format limits high-speed touring capability</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha Aerox 155 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban commuters seeking premium scooter</li>
<li class="best-for-item">Riders wanting automatic convenience</li>
<li class="best-for-item">Those needing practical storage</li>
<li class="best-for-item">Daily city navigation</li>
<li class="best-for-item">First-time riders preferring automatic</li>
<li class="best-for-item">Style-conscious scooter enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha Aerox 155?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance motorcycle enthusiasts</li>
<li class="not-recommended-item">Long-distance highway touring</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those preferring manual transmission</li>
<li class="not-recommended-item">Off-road or rough terrain riding</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha Aerox 155 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳5,25,000 - Premium for 155cc maxi-scooter with VVA</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate to High - ৳9-12 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳7-9 per km (40-45 km/l with VVA and CVT)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 4,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳8,000-14,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha Aerox 155 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Good for a scooter</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Efficient VVA and CVT</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent scooter comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Sporty maxi-scooter design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Premium scooter pricing</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good Yamaha build</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha Aerox 155 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does the Aerox 155 compare to traditional motorcycles?</h3>
<p class="faq-answer">The Aerox 155 offers automatic convenience, better weather protection, and practical storage, but lacks the sporty engagement and high-speed capability of traditional motorcycles.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What are the advantages of VVA in a scooter?</h3>
<p class="faq-answer">VVA technology in the Aerox 155 provides better power delivery, improved fuel efficiency, and enhanced performance compared to conventional scooter engines, while maintaining automatic convenience.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for highway riding?</h3>
<p class="faq-answer">The Aerox 155 can handle highway speeds reasonably well, but it's primarily designed for urban commuting. Long highway rides may be less comfortable compared to motorcycles.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha Aerox 155 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳5,25,000, the Yamaha Aerox 155 offers a unique proposition as a premium sports maxi-scooter with VVA technology, automatic convenience, and practical features. It combines motorcycle-level performance with scooter practicality, making it ideal for urban commuting and style-conscious riders. However, the high pricing for a scooter, maintenance complexity, limited sporty engagement, and highway limitations make it suitable primarily for daily urban commuters who prioritize automatic convenience, storage, and premium styling over traditional motorcycle dynamics.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium automatic convenience with VVA performance</span></p>
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
		return fmt.Errorf("error creating Yamaha Aerox 155 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha Aerox 155 (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা এরক্স ১৫৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা এরক্স ১৫৫ হল একটি প্রিমিয়াম ১৫৫সিসি স্পোর্টস ম্যাক্সি-স্কুটার যেখানে ভিভিএ প্রযুক্তি, অ্যারোডাইনামিক ডিজাইন এবং উন্নত বৈশিষ্ট্য রয়েছে। ৳৫,২৫,০০০ টাকায় মূল্যায়িত, এটি ভেরিয়েবল ভালভ অ্যাকচুয়েশন, স্পোর্টি এরগোনমিক্স, বড় স্টোরেজ ক্ষমতা এবং মোটরসাইকেল-লেভেল পারফরমেন্স সহ প্রিমিয়াম স্কুটার অভিজ্ঞতা খোঁজা রাইডারদের জন্য ইয়ামাহার প্রমাণিত ইঞ্জিনিয়ারিং সহ স্পোর্টস বাইক পারফরমেন্সের সাথে অটোমেটিক ট্রান্সমিশনের সুবিধা একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা এরক্স ১৫৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">স্পোর্টস ম্যাক্সি-স্কুটার ডিজাইন:</strong> <span class="highlight-value">স্পোর্টস আবেদন সহ আক্রমণাত্মক অ্যারোডাইনামিক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ভিভিএ প্রযুক্তি:</strong> <span class="highlight-value">সর্বোত্তম পারফরমেন্সের জন্য ভেরিয়েবল ভালভ অ্যাকচুয়েশন</span></li>
<li class="highlight-item"><strong class="highlight-label">অটোমেটিক সিভিটি ট্রান্সমিশন:</strong> <span class="highlight-value">ক্লাচ ছাড়াই নিরবচ্ছিন্ন পাওয়ার ডেলিভারি</span></li>
<li class="highlight-item"><strong class="highlight-label">বড় স্টোরেজ ক্ষমতা:</strong> <span class="highlight-value">দৈনিক ব্যবহারের জন্য ব্যবহারিক আন্ডারসিট স্টোরেজ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা এরক্স ১৫৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ভিভিএ ইঞ্জিন পারফরমেন্স:</strong> <span class="pro-description">ভেরিয়েবল ভালভ প্রযুক্তি একটি স্কুটারের জন্য শক্তিশালী পারফরমেন্স প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">অটোমেটিক সুবিধা:</strong> <span class="pro-description">সিভিটি ট্রান্সমিশন ক্লাচ অপারেশন এবং গিয়ার শিফটিং দূর করে</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি স্টাইলিং:</strong> <span class="pro-description">মোটরসাইকেল-অনুপ্রাণিত নান্দনিকতা সহ আক্রমণাত্মক ম্যাক্সি-স্কুটার ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক স্টোরেজ:</strong> <span class="pro-description">হেলমেট এবং জিনিসপত্রের জন্য বড় আন্ডারসিট কম্পার্টমেন্ট</span></li>
<li class="pro-item"><strong class="pro-label">আবহাওয়া সুরক্ষা:</strong> <span class="pro-description">নেকেড মোটরসাইকেলের তুলনায় প্রাকৃতিক উপাদান থেকে ভাল সুরক্ষা</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি দক্ষতা:</strong> <span class="pro-description">ভিভিএ এবং সিভিটি অপ্টিমাইজেশনের সাথে ভাল জ্বালানি অর্থনীতি</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা এরক্স ১৫৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ ক্রয় মূল্য:</strong> <span class="con-description">একটি ১৫৫সিসি স্কুটারের জন্য প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">সীমিত স্পোর্টি অনুভূতি:</strong> <span class="con-description">সিভিটি ট্রান্সমিশনে ম্যানুয়াল গিয়ারবক্সের এনগেজমেন্টের অভাব</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ জটিলতা:</strong> <span class="con-description">ভিভিএ প্রযুক্তি এবং সিভিটির বিশেষায়িত সেবা প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">প্রিমিয়াম স্কুটার সেবায় অভিজ্ঞ কম ডিলারশিপ</span></li>
<li class="con-item"><strong class="con-label">হাইওয়ে সীমাবদ্ধতা:</strong> <span class="con-description">স্কুটার ফরম্যাট উচ্চ-গতির ভ্রমণ ক্ষমতা সীমিত করে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা এরক্স ১৫৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Urban commuters seeking premium scooter</li>
<li class="best-for-item">Riders wanting automatic convenience</li>
<li class="best-for-item">Those needing practical storage</li>
<li class="best-for-item">Daily city navigation</li>
<li class="best-for-item">First-time riders preferring automatic</li>
<li class="best-for-item">Style-conscious scooter enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা এরক্স ১৫৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Performance motorcycle enthusiasts</li>
<li class="not-recommended-item">Long-distance highway touring</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Those preferring manual transmission</li>
<li class="not-recommended-item">Off-road or rough terrain riding</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা এরক্স ১৫৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৫,২৫,০০০ - ভিভিএ সহ ১৫৫সিসি ম্যাক্সি-স্কুটারের জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি থেকে উচ্চ - ৳৯-১২ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৭-৯ প্রতি কিমি (ভিভিএ এবং সিভিটি সহ ৪০-৪৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৪,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৮,০০০-১৪,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা এরক্স ১৫৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- একটি স্কুটারের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- দক্ষ ভিভিএ এবং সিভিটি</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- চমৎকার স্কুটার আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- স্পোর্টি ম্যাক্সি-স্কুটার ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- প্রিমিয়াম স্কুটার মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- ভাল ইয়ামাহা বিল্ড</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা এরক্স ১৫৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এরক্স ১৫৫ প্রথাগত মোটরসাইকেলের সাথে কেমন তুলনা করে?</h3>
<p class="faq-answer">এরক্স ১৫৫ অটোমেটিক সুবিধা, ভাল আবহাওয়া সুরক্ষা এবং ব্যবহারিক স্টোরেজ প্রদান করে, কিন্তু প্রথাগত মোটরসাইকেলের স্পোর্টি এনগেজমেন্ট এবং উচ্চ-গতির ক্ষমতার অভাব রয়েছে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">একটি স্কুটারে ভিভিএর সুবিধা কী?</h3>
<p class="faq-answer">এরক্স ১৫৫ এ ভিভিএ প্রযুক্তি অটোমেটিক সুবিধা বজায় রেখে প্রচলিত স্কুটার ইঞ্জিনের তুলনায় ভাল পাওয়ার ডেলিভারি, উন্নত জ্বালানি দক্ষতা এবং উন্নত পারফরমেন্স প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি হাইওয়ে রাইডিংয়ের জন্য উপযুক্ত?</h3>
<p class="faq-answer">এরক্স ১৫৫ হাইওয়ে গতি যুক্তিসঙ্গতভাবে ভালভাবে পরিচালনা করতে পারে, কিন্তু এটি প্রাথমিকভাবে শহুরে যাতায়াতের জন্য ডিজাইন করা। দীর্ঘ হাইওয়ে রাইড মোটরসাইকেলের তুলনায় কম আরামদায়ক হতে পারে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা এরক্স ১৫৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৫,২৫,০০০ টাকায় ইয়ামাহা এরক্স ১৫৫ ভিভিএ প্রযুক্তি, অটোমেটিক সুবিধা এবং ব্যবহারিক বৈশিষ্ট্য সহ একটি প্রিমিয়াম স্পোর্টস ম্যাক্সি-স্কুটার হিসেবে একটি অনন্য প্রস্তাব প্রদান করে। এটি স্কুটার ব্যবহারিকতার সাথে মোটরসাইকেল-লেভেল পারফরমেন্স একত্রিত করে, এটি শহুরে যাতায়াত এবং স্টাইল-সচেতন রাইডারদের জন্য আদর্শ করে তোলে। তবে, একটি স্কুটারের জন্য উচ্চ মূল্য, রক্ষণাবেক্ষণ জটিলতা, সীমিত স্পোর্টি এনগেজমেন্ট এবং হাইওয়ে সীমাবদ্ধতা এটিকে প্রধানত দৈনিক শহুরে যাত্রীদের জন্য উপযুক্ত করে তোলে যারা প্রথাগত মোটরসাইকেল ডায়নামিক্সের চেয়ে অটোমেটিক সুবিধা, স্টোরেজ এবং প্রিমিয়াম স্টাইলিংকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ভিভিএ পারফরমেন্স সহ প্রিমিয়াম অটোমেটিক সুবিধার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha Aerox 155 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha Aerox 155\n")

	return nil
}
