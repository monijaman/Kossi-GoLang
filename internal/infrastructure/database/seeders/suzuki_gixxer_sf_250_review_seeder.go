package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxerSF250Review handles seeding Suzuki Gixxer SF 250 product review and translation
type SuzukiGixxerSF250Review struct {
	BaseSeeder
}

// NewSuzukiGixxerSF250ReviewSeeder creates a new SuzukiGixxerSF250Review
func NewSuzukiGixxerSF250ReviewSeeder() *SuzukiGixxerSF250Review {
	return &SuzukiGixxerSF250Review{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer SF 250 Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxerSF250Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer SF 250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Gixxer SF 250 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Gixxer SF 250 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer SF 250 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer SF 250 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Gixxer SF 250 is a fully-faired quarter-liter sports motorcycle featuring a 249cc single-cylinder SOHC engine producing 26.5 HP and 22.6 Nm torque. At ৳4,29,950, it offers aggressive supersport styling with full fairing aerodynamics, LED lighting, digital instrumentation, and Suzuki's renowned build quality in the premium 250cc segment.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer SF 250 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Quarter-Liter Performance:</strong> <span class="highlight-value">249cc SOHC single-cylinder engine with 26.5 HP performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Full Fairing Design:</strong> <span class="highlight-value">Superior aerodynamics and wind protection</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Advanced LED headlight with signature Suzuki styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Console:</strong> <span class="highlight-value">Instrument cluster with gear position indicator</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer SF 250 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Quarter-Liter Performance:</strong> <span class="pro-description">249cc engine delivers good power and torque for highway riding</span></li>
<li class="pro-item"><strong class="pro-label">Full Fairing Advantage:</strong> <span class="pro-description">Excellent wind protection and aerodynamic efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Modern LED Lighting:</strong> <span class="pro-description">Bright LED headlight with contemporary styling appeal</span></li>
<li class="pro-item"><strong class="pro-label">Digital Instrumentation:</strong> <span class="pro-description">Comprehensive display with gear indicator and trip computer</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Reliability:</strong> <span class="pro-description">Proven engineering quality and dependable performance</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Solid construction with premium fit and finish</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer SF 250 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Extremely High Price:</strong> <span class="con-description">৳4,29,950 is very expensive for single-cylinder 250cc bike</span></li>
<li class="con-item"><strong class="con-label">No ABS Braking:</strong> <span class="con-description">Missing essential safety technology at this premium price</span></li>
<li class="con-item"><strong class="con-label">Single Cylinder Vibrations:</strong> <span class="con-description">Engine buzz becomes noticeable at higher RPMs</span></li>
<li class="con-item"><strong class="con-label">Heavy City Riding:</strong> <span class="con-description">Full fairing makes it cumbersome in traffic conditions</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">High service costs and parts pricing</span></li>
<li class="con-item"><strong class="con-label">Poor Value Proposition:</strong> <span class="con-description">Overpriced compared to competitors offering more features</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Gixxer SF 250 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Quarter-liter sports bike enthusiasts</li>
<li class="best-for-item">Highway touring and long-distance riding</li>
<li class="best-for-item">Riders wanting full-fairing aerodynamics</li>
<li class="best-for-item">Weekend sport riding and touring</li>
<li class="best-for-item">Experienced riders upgrading from 150cc</li>
<li class="best-for-item">Those prioritizing Suzuki build quality</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Gixxer SF 250?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers seeking value</li>
<li class="not-recommended-item">New or inexperienced riders</li>
<li class="not-recommended-item">Heavy city commuting needs</li>
<li class="not-recommended-item">Riders requiring ABS safety features</li>
<li class="not-recommended-item">Those needing extensive service support</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer SF 250 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,29,950 - Very Expensive for 250cc</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳12-15 per km with premium maintenance</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳12-15 per km (35-40 km/l at current fuel prices)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳8,000-12,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer SF 250 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good power delivery</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Average fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Decent touring comfort</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Aggressive supersport design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.0</span> <span class="rating-note">- Poor value for money</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent build quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Gixxer SF 250 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does the Gixxer SF 250 perform on highways?</h3>
<p class="faq-answer">Excellent highway performance with 26.5 HP providing good cruising speeds of 110-120 km/h. Full fairing offers superior wind protection making long rides comfortable.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does it come with ABS as standard?</h3>
<p class="faq-answer">No, the Gixxer SF 250 does not come with ABS as standard, which is a significant drawback at this premium price point of ৳4,29,950.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the maintenance expensive for Gixxer SF 250?</h3>
<p class="faq-answer">Yes, annual maintenance costs ৳8,000-12,000 including engine oil, filters, and routine servicing. Premium positioning means higher parts and service costs.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer SF 250 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,29,950, the Suzuki Gixxer SF 250 offers strong quarter-liter performance with excellent full-fairing aerodynamics, LED lighting, digital instrumentation, and reliable Suzuki engineering. The 249cc engine delivers good highway performance while the full fairing provides superior wind protection for touring. However, the extremely high pricing without ABS, single-cylinder vibrations, expensive maintenance, poor value proposition, and limited service network make it a difficult recommendation except for dedicated Suzuki enthusiasts who prioritize brand loyalty and full-fairing design over practical value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For quarter-liter fully-faired touring enthusiasts</span></p>
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
		return fmt.Errorf("error creating Suzuki Gixxer SF 250 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer SF 250 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার এসএফ ২৫০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি গিক্সার এসএফ ২৫০ হল একটি ফুল-ফেয়ার্ড কোয়ার্টার-লিটার স্পোর্টস মোটরসাইকেল যেখানে ২৪৯সিসি একক সিলিন্ডার এসওএইচসি ইঞ্জিন রয়েছে যা ২৬.৫ এইচপি এবং ২২.৬ এনএম টর্ক উৎপাদন করে। ৳৪,২৯,৯৫০ টাকায়, এটি ফুল ফেয়ারিং অ্যারোডাইনামিক্স, এলইডি লাইটিং, ডিজিটাল ইনস্ট্রুমেন্টেশন এবং প্রিমিয়াম ২৫০সিসি সেগমেন্টে সুজুকির বিখ্যাত বিল্ড কোয়ালিটি সহ আক্রমণাত্মক সুপারস্পোর্ট স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার এসএফ ২৫০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">কোয়ার্টার-লিটার পারফরমেন্স:</strong> <span class="highlight-value">২৬.৫ এইচপি পারফরমেন্স সহ ২৪৯সিসি এসওএইচসি একক সিলিন্ডার ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ ফেয়ারিং ডিজাইন:</strong> <span class="highlight-value">উন্নত অ্যারোডাইনামিক্স এবং বাতাস সুরক্ষা</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">স্বাক্ষর সুজুকি স্টাইলিং সহ উন্নত এলইডি হেডলাইট</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল কনসোল:</strong> <span class="highlight-value">গিয়ার পজিশন ইন্ডিকেটর সহ ইনস্ট্রুমেন্ট ক্লাস্টার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার এসএফ ২৫০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী কোয়ার্টার-লিটার পারফরমেন্স:</strong> <span class="pro-description">হাইওয়ে রাইডিংয়ের জন্য ২৪৯সিসি ইঞ্জিন ভালো শক্তি এবং টর্ক প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">সম্পূর্ণ ফেয়ারিং সুবিধা:</strong> <span class="pro-description">চমৎকার বাতাস সুরক্ষা এবং অ্যারোডাইনামিক দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক এলইডি লাইটিং:</strong> <span class="pro-description">সমসাময়িক স্টাইলিং আবেদন সহ উজ্জ্বল এলইডি হেডলাইট</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল ইনস্ট্রুমেন্টেশন:</strong> <span class="pro-description">গিয়ার ইন্ডিকেটর এবং ট্রিপ কম্পিউটার সহ ব্যাপক ডিসপ্লে</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত ইঞ্জিনিয়ারিং গুণমান এবং নির্ভরযোগ্য পারফরমেন্স</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম ফিট এবং ফিনিশ সহ কঠিন নির্মাণ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার এসএফ ২৫০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত উচ্চ মূল্য:</strong> <span class="con-description">একক সিলিন্ডার ২৫০সিসি বাইকের জন্য ৳৪,২৯,৯৫০ অত্যন্ত ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস ব্রেকিং:</strong> <span class="con-description">এই প্রিমিয়াম মূল্যে অত্যাবশ্যক নিরাপত্তা প্রযুক্তির অভাব</span></li>
<li class="con-item"><strong class="con-label">একক সিলিন্ডার কম্পন:</strong> <span class="con-description">উচ্চ আরপিএমে ইঞ্জিন কম্পন লক্ষণীয় হয়ে ওঠে</span></li>
<li class="con-item"><strong class="con-label">ভারী শহুরে রাইডিং:</strong> <span class="con-description">সম্পূর্ণ ফেয়ারিং ট্রাফিক অবস্থায় এটিকে কষ্টকর করে তোলে</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">উচ্চ সেবা খরচ এবং পার্টসের মূল্য</span></li>
<li class="con-item"><strong class="con-label">খারাপ মূল্য প্রস্তাব:</strong> <span class="con-description">আরো ফিচার প্রদানকারী প্রতিযোগীদের তুলনায় অতিরিক্ত দামি</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি গিক্সার এসএফ ২৫০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Quarter-liter sports bike enthusiasts</li>
<li class="best-for-item">Highway touring and long-distance riding</li>
<li class="best-for-item">Riders wanting full-fairing aerodynamics</li>
<li class="best-for-item">Weekend sport riding and touring</li>
<li class="best-for-item">Experienced riders upgrading from 150cc</li>
<li class="best-for-item">Those prioritizing Suzuki build quality</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি গিক্সার এসএফ ২৫০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers seeking value</li>
<li class="not-recommended-item">New or inexperienced riders</li>
<li class="not-recommended-item">Heavy city commuting needs</li>
<li class="not-recommended-item">Riders requiring ABS safety features</li>
<li class="not-recommended-item">Those needing extensive service support</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি গিক্সার এসএফ ২৫০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৪,২৯,৯৫০ - ২৫০সিসির জন্য অত্যন্ত দামি</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">উচ্চ - প্রিমিয়াম রক্ষণাবেক্ষণ সহ ৳১২-১৫ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১২-১৫ প্রতি কিমি (বর্তমান জ্বালানি দামে ৩৫-৪০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৮,০০০-১২,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি গিক্সার এসএফ ২৫০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- ভালো শক্তি বিতরণ</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- গড় জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- শোভন ট্যুরিং আরাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- আক্রমণাত্মক সুপারস্পোর্ট ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.0</span> <span class="rating-note">- অর্থের জন্য খারাপ মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার বিল্ড গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি গিক্সার এসএফ ২৫০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">গিক্সার এসএফ ২৫০ হাইওয়েতে কেমন পারফরম করে?</h3>
<p class="faq-answer">২৬.৫ এইচপি সহ চমৎকার হাইওয়ে পারফরমেন্স যা ১১০-১২০ কিমি/ঘণ্টার ভাল ক্রুজিং স্পিড প্রদান করে। ফুল ফেয়ারিং উন্নত বাতাস সুরক্ষা প্রদান করে যা দীর্ঘ রাইডকে আরামদায়ক করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি স্ট্যান্ডার্ড হিসেবে এবিএস সহ আসে?</h3>
<p class="faq-answer">না, গিক্সার এসএফ ২৫০ স্ট্যান্ডার্ড হিসেবে এবিএস সহ আসে না, যা ৳৪,২৯,৯৫০ এই প্রিমিয়াম মূল্যে একটি উল্লেখযোগ্য ত্রুটি।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">গিক্সার এসএফ ২৫০ এর রক্ষণাবেক্ষণ কি ব্যয়বহুল?</h3>
<p class="faq-answer">হ্যাঁ, ইঞ্জিন অয়েল, ফিল্টার এবং নিয়মিত সার্ভিসিং সহ বার্ষিক রক্ষণাবেক্ষণ খরচ ৳৮,০০০-১২,০০০। প্রিমিয়াম পজিশনিং মানে উচ্চ পার্টস এবং সেবা খরচ।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি গিক্সার এসএফ ২৫০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৪,২৯,৯৫০ টাকায় সুজুকি গিক্সার এসএফ ২৫০ চমৎকার ফুল-ফেয়ারিং অ্যারোডাইনামিক্স, এলইডি লাইটিং, ডিজিটাল ইনস্ট্রুমেন্টেশন এবং নির্ভরযোগ্য সুজুকি ইঞ্জিনিয়ারিং সহ শক্তিশালী কোয়ার্টার-লিটার পারফরমেন্স প্রদান করে। ২৪৯সিসি ইঞ্জিন ভাল হাইওয়ে পারফরমেন্স প্রদান করে যখন ফুল ফেয়ারিং ট্যুরিংয়ের জন্য উন্নত বাতাস সুরক্ষা প্রদান করে। তবে, এবিএস ছাড়া অত্যন্ত উচ্চ মূল্য, একক সিলিন্ডার কম্পন, ব্যয়বহুল রক্ষণাবেক্ষণ, খারাপ মূল্য প্রস্তাব এবং সীমিত সেবা নেটওয়ার্ক এটিকে নিবেদিতপ্রাণ সুজুকি উৎসাহীদের ব্যতীত কঠিন সুপারিশ করে তোলে যারা ব্যবহারিক মূল্যের চেয়ে ব্র্যান্ড আনুগত্য এবং ফুল-ফেয়ারিং ডিজাইনকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- কোয়ার্টার-লিটার ফুল-ফেয়ার্ড ট্যুরিং উৎসাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer SF 250 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer SF 250\n")

	return nil
}
