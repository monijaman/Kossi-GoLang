package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiIntruderFiABSReview handles seeding Suzuki Intruder FI ABS product review and translation
type SuzukiIntruderFiABSReview struct {
	BaseSeeder
}

// NewSuzukiIntruderFiABSReviewSeeder creates a new SuzukiIntruderFiABSReview
func NewSuzukiIntruderFiABSReviewSeeder() *SuzukiIntruderFiABSReview {
	return &SuzukiIntruderFiABSReview{BaseSeeder: BaseSeeder{name: "Suzuki Intruder FI ABS Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiIntruderFiABSReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Intruder FI ABS").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Intruder FI ABS product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Intruder FI ABS product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Intruder FI ABS already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Intruder FI ABS Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Intruder FI ABS is a 155cc cruiser motorcycle featuring fuel injection technology and single-channel ABS braking system. Priced at ৳3,19,950, it offers a unique cruiser styling with low seat height, forward controls, relaxed riding position, and modern FI efficiency combined with ABS safety for comfortable cruising and daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Intruder FI ABS Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Cruiser Styling:</strong> <span class="highlight-value">Unique low-slung cruiser design with forward controls</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Injection System:</strong> <span class="highlight-value">Modern FI technology for better efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Single Channel ABS:</strong> <span class="highlight-value">Front wheel ABS for enhanced braking safety</span></li>
<li class="highlight-item"><strong class="highlight-label">Low Seat Height:</strong> <span class="highlight-value">Comfortable seating for riders of all heights</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Intruder FI ABS Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Unique Cruiser Design:</strong> <span class="pro-description">Distinctive styling that stands out from typical motorcycles</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Riding Position:</strong> <span class="pro-description">Relaxed cruiser posture ideal for long rides</span></li>
<li class="pro-item"><strong class="pro-label">ABS Safety Feature:</strong> <span class="pro-description">Anti-lock braking system prevents wheel lock-up</span></li>
<li class="pro-item"><strong class="pro-label">Low Seat Height:</strong> <span class="pro-description">Easy to handle for shorter riders and beginners</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Injection Efficiency:</strong> <span class="pro-description">Better fuel economy and consistent performance</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Suzuki's proven Japanese engineering standards</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Intruder FI ABS Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Ground Clearance:</strong> <span class="con-description">Low design may scrape on uneven roads</span></li>
<li class="con-item"><strong class="con-label">Single Channel ABS Only:</strong> <span class="con-description">Lacks rear wheel ABS for complete protection</span></li>
<li class="con-item"><strong class="con-label">Highway Performance:</strong> <span class="con-description">155cc engine limits high-speed cruising capability</span></li>
<li class="con-item"><strong class="con-label">Higher Maintenance:</strong> <span class="con-description">Fuel injection system requires specialized service</span></li>
<li class="con-item"><strong class="con-label">Limited Storage:</strong> <span class="con-description">Cruiser design lacks practical storage options</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Intruder FI ABS in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Riders seeking unique cruiser styling</li>
<li class="best-for-item">Daily commuters wanting comfort</li>
<li class="best-for-item">Weekend leisure riding enthusiasts</li>
<li class="best-for-item">Shorter riders preferring low seat height</li>
<li class="best-for-item">Those prioritizing ABS safety</li>
<li class="best-for-item">Style-conscious motorcycle enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Intruder FI ABS?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway speed enthusiasts</li>
<li class="not-recommended-item">Off-road or rough terrain riding</li>
<li class="not-recommended-item">Riders needing practical storage</li>
<li class="not-recommended-item">Power-hungry performance seekers</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Intruder FI ABS Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,19,950 - Premium for 155cc cruiser with ABS</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-10 per km with FI efficiency</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳8-10 per km (40-45 km/l with FI technology)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳5,500-7,500 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Intruder FI ABS Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Adequate for cruising</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good with FI system</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent cruiser comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Outstanding cruiser design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- Premium pricing</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good Suzuki build</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Intruder FI ABS Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What makes the Intruder different from other Suzuki bikes?</h3>
<p class="faq-answer">The Intruder features a unique cruiser design with low-slung styling, forward foot controls, relaxed riding position, and distinctive aesthetic that differentiates it from standard motorcycles.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the low seat height suitable for short riders?</h3>
<p class="faq-answer">Yes, the Intruder's low seat height makes it very accessible for shorter riders and provides confidence when stopping or maneuvering at low speeds.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does single-channel ABS work?</h3>
<p class="faq-answer">Single-channel ABS monitors only the front wheel and prevents it from locking during hard braking. While not as comprehensive as dual-channel, it still significantly improves braking safety.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Intruder FI ABS in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,19,950, the Suzuki Intruder FI ABS offers a unique cruiser experience in the 155cc segment with distinctive styling, comfortable ergonomics, ABS safety, and fuel injection efficiency. The low seat height, relaxed riding position, and eye-catching design make it appealing for style-conscious riders and daily commuters seeking comfort. However, the premium pricing for limited performance, single-channel ABS only, poor ground clearance, specialized maintenance requirements, and limited practical utility make it suitable primarily for riders who prioritize style and comfort over performance and value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For unique cruiser styling with comfort and ABS safety</span></p>
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
		return fmt.Errorf("error creating Suzuki Intruder FI ABS review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Intruder FI ABS (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি ইনট্রুডার এফআই এবিএস রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি ইনট্রুডার এফআই এবিএস হল একটি ১৫৫সিসি ক্রুজার মোটরসাইকেল যেখানে ফুয়েল ইনজেকশন প্রযুক্তি এবং সিঙ্গেল-চ্যানেল এবিএস ব্রেকিং সিস্টেম রয়েছে। ৳৩,১৯,৯৫০ টাকায় মূল্যায়িত, এটি নিম্ন সিট উচ্চতা, ফরোয়ার্ড কন্ট্রোল, শিথিল রাইডিং পজিশন এবং আরামদায়ক ক্রুজিং ও দৈনিক যাতায়াতের জন্য এবিএস নিরাপত্তার সাথে মিলিত আধুনিক এফআই দক্ষতা সহ একটি অনন্য ক্রুজার স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি ইনট্রুডার এফআই এবিএস এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ক্রুজার স্টাইলিং:</strong> <span class="highlight-value">ফরোয়ার্ড কন্ট্রোল সহ অনন্য লো-স্লাং ক্রুজার ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">ফুয়েল ইনজেকশন সিস্টেম:</strong> <span class="highlight-value">ভাল দক্ষতার জন্য আধুনিক এফআই প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="highlight-value">উন্নত ব্রেকিং নিরাপত্তার জন্য সামনের চাকার এবিএস</span></li>
<li class="highlight-item"><strong class="highlight-label">নিম্ন সিট উচ্চতা:</strong> <span class="highlight-value">সব উচ্চতার রাইডারদের জন্য আরামদায়ক বসা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি ইনট্রুডার এফআই এবিএস এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অনন্য ক্রুজার ডিজাইন:</strong> <span class="pro-description">সাধারণ মোটরসাইকেল থেকে আলাদা স্বতন্ত্র স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইডিং পজিশন:</strong> <span class="pro-description">দীর্ঘ রাইডের জন্য আদর্শ শিথিল ক্রুজার ভঙ্গি</span></li>
<li class="pro-item"><strong class="pro-label">এবিএস নিরাপত্তা ফিচার:</strong> <span class="pro-description">অ্যান্টি-লক ব্রেকিং সিস্টেম হইল লক-আপ প্রতিরোধ করে</span></li>
<li class="pro-item"><strong class="pro-label">নিম্ন সিট উচ্চতা:</strong> <span class="pro-description">খাটো রাইডার এবং শিক্ষানবিশদের জন্য সহজে পরিচালনা</span></li>
<li class="pro-item"><strong class="pro-label">ফুয়েল ইনজেকশন দক্ষতা:</strong> <span class="pro-description">ভাল জ্বালানি অর্থনীতি এবং সামঞ্জস্যপূর্ণ পারফরমেন্স</span></li>
<li class="pro-item"><strong class="pro-label">ভাল বিল্ড কোয়ালিটি:</strong> <span class="pro-description">সুজুকির প্রমাণিত জাপানি ইঞ্জিনিয়ারিং মান</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি ইনট্রুডার এফআই এবিএস এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="con-description">নিচু ডিজাইন অসমান রাস্তায় ঘষতে পারে</span></li>
<li class="con-item"><strong class="con-label">কেবলমাত্র সিঙ্গেল চ্যানেল এবিএস:</strong> <span class="con-description">সম্পূর্ণ সুরক্ষার জন্য পিছনের চাকার এবিএসের অভাব</span></li>
<li class="con-item"><strong class="con-label">হাইওয়ে পারফরমেন্স:</strong> <span class="con-description">১৫৫সিসি ইঞ্জিন উচ্চ-গতির ক্রুজিং ক্ষমতা সীমিত করে</span></li>
<li class="con-item"><strong class="con-label">উচ্চ রক্ষণাবেক্ষণ:</strong> <span class="con-description">ফুয়েল ইনজেকশন সিস্টেমের বিশেষায়িত সেবা প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত স্টোরেজ:</strong> <span class="con-description">ক্রুজার ডিজাইনে ব্যবহারিক স্টোরেজ অপশনের অভাব</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি ইনট্রুডার এফআই এবিএস কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Riders seeking unique cruiser styling</li>
<li class="best-for-item">Daily commuters wanting comfort</li>
<li class="best-for-item">Weekend leisure riding enthusiasts</li>
<li class="best-for-item">Shorter riders preferring low seat height</li>
<li class="best-for-item">Those prioritizing ABS safety</li>
<li class="best-for-item">Style-conscious motorcycle enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি ইনট্রুডার এফআই এবিএস কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway speed enthusiasts</li>
<li class="not-recommended-item">Off-road or rough terrain riding</li>
<li class="not-recommended-item">Riders needing practical storage</li>
<li class="not-recommended-item">Power-hungry performance seekers</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি ইনট্রুডার এফআই এবিএস এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৩,১৯,৯৫০ - এবিএস সহ ১৫৫সিসি ক্রুজারের জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মধ্যম - এফআই দক্ষতা সহ ৳৮-১০ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৮-১০ প্রতি কিমি (এফআই প্রযুক্তি সহ ৪০-৪৫ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৫,৫০০-৭,৫০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি ইনট্রুডার এফআই এবিএস পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- ক্রুজিংয়ের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- এফআই সিস্টেমের সাথে ভাল</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার ক্রুজার আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- অসাধারণ ক্রুজার ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.7</span> <span class="rating-note">- প্রিমিয়াম মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ভাল সুজুকি বিল্ড</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি ইনট্রুডার এফআই এবিএস সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ইনট্রুডার অন্যান্য সুজুকি বাইক থেকে কী আলাদা করে?</h3>
<p class="faq-answer">ইনট্রুডার একটি অনন্য ক্রুজার ডিজাইন ফিচার করে যেখানে লো-স্লাং স্টাইলিং, ফরোয়ার্ড ফুট কন্ট্রোল, শিথিল রাইডিং পজিশন এবং স্বতন্ত্র নান্দনিকতা রয়েছে যা এটিকে সাধারণ মোটরসাইকেল থেকে আলাদা করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">নিম্ন সিট উচ্চতা খাটো রাইডারদের জন্য উপযুক্ত?</h3>
<p class="faq-answer">হ্যাঁ, ইনট্রুডারের নিম্ন সিট উচ্চতা এটিকে খাটো রাইডারদের জন্য অত্যন্ত অভিগম্য করে তোলে এবং থামানো বা কম গতিতে চালনার সময় আত্মবিশ্বাস প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সিঙ্গেল-চ্যানেল এবিএস কীভাবে কাজ করে?</h3>
<p class="faq-answer">সিঙ্গেল-চ্যানেল এবিএস শুধুমাত্র সামনের চাকা পর্যবেক্ষণ করে এবং কঠিন ব্রেকিংয়ের সময় এটিকে লক হওয়া থেকে প্রতিরোধ করে। ডুয়াল-চ্যানেলের মতো ব্যাপক না হলেও, এটি এখনও উল্লেখযোগ্যভাবে ব্রেকিং নিরাপত্তা উন্নত করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি ইনট্রুডার এফআই এবিএস কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৩,১৯,৯৫০ টাকায় সুজুকি ইনট্রুডার এফআই এবিএস স্বতন্ত্র স্টাইলিং, আরামদায়ক এরগোনমিক্স, এবিএস নিরাপত্তা এবং ফুয়েল ইনজেকশন দক্ষতা সহ ১৫৫সিসি সেগমেন্টে একটি অনন্য ক্রুজার অভিজ্ঞতা প্রদান করে। নিম্ন সিট উচ্চতা, শিথিল রাইডিং পজিশন এবং চোখ ধাঁধানো ডিজাইন এটিকে স্টাইল-সচেতন রাইডার এবং আরাম খোঁজা দৈনিক যাত্রীদের জন্য আকর্ষণীয় করে তোলে। তবে, সীমিত পারফরমেন্সের জন্য প্রিমিয়াম মূল্য, শুধুমাত্র সিঙ্গেল-চ্যানেল এবিএস, খারাপ গ্রাউন্ড ক্লিয়ারেন্স, বিশেষায়িত রক্ষণাবেক্ষণের প্রয়োজনীয়তা এবং সীমিত ব্যবহারিক উপযোগিতা এটিকে প্রধানত সেই রাইডারদের জন্য উপযুক্ত করে তোলে যারা পারফরমেন্স এবং ভ্যালুর চেয়ে স্টাইল এবং আরামকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আরাম এবং এবিএস নিরাপত্তা সহ অনন্য ক্রুজার স্টাইলিংয়ের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Intruder FI ABS already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Intruder FI ABS\n")

	return nil
}
