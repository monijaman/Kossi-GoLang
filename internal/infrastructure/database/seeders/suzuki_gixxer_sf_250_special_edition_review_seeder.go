package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGixxerSF250SpecialEditionReview handles seeding Suzuki Gixxer SF 250 Special Edition product review and translation
type SuzukiGixxerSF250SpecialEditionReview struct {
	BaseSeeder
}

// NewSuzukiGixxerSF250SpecialEditionReviewSeeder creates a new SuzukiGixxerSF250SpecialEditionReview
func NewSuzukiGixxerSF250SpecialEditionReviewSeeder() *SuzukiGixxerSF250SpecialEditionReview {
	return &SuzukiGixxerSF250SpecialEditionReview{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer SF 250 Special Edition Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGixxerSF250SpecialEditionReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer SF 250 Special Edition").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Suzuki Gixxer SF 250 Special Edition product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Suzuki Gixxer SF 250 Special Edition product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki Gixxer SF 250 Special Edition already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki Gixxer SF 250 Special Edition Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Suzuki Gixxer SF 250 Special Edition is the flagship fully-faired quarter-liter sports motorcycle from Suzuki, featuring a powerful 249cc single-cylinder engine producing 26.5 HP and 22.6 Nm torque. Priced at ৳4,29,950, it combines aggressive full-fairing aerodynamics, premium digital instrumentation, LED lighting package, and signature Suzuki reliability in an exclusive special edition package.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer SF 250 Special Edition Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Powerful Quarter-Liter Engine:</strong> <span class="highlight-value">249cc SOHC engine with 26.5 HP and 22.6 Nm torque</span></li>
<li class="highlight-item"><strong class="highlight-label">Exclusive Special Edition:</strong> <span class="highlight-value">Premium graphics and unique color scheme with special badging</span></li>
<li class="highlight-item"><strong class="highlight-label">Full LED Lighting Package:</strong> <span class="highlight-value">Complete LED headlight, taillight and indicators</span></li>
<li class="highlight-item"><strong class="highlight-label">Advanced Digital Console:</strong> <span class="highlight-value">Multi-function display with gear position indicator</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer SF 250 Special Edition Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Powerful Quarter-Liter Performance:</strong> <span class="pro-description">249cc engine delivers strong 26.5 HP with excellent torque curve</span></li>
<li class="pro-item"><strong class="pro-label">Exclusive Special Edition Appeal:</strong> <span class="pro-description">Premium graphics, unique color scheme and special badging</span></li>
<li class="pro-item"><strong class="pro-label">Full Fairing Aerodynamics:</strong> <span class="pro-description">Excellent wind protection and aggressive supersport styling</span></li>
<li class="pro-item"><strong class="pro-label">Comprehensive LED Package:</strong> <span class="pro-description">Full LED headlight, taillight and indicators for modern appeal</span></li>
<li class="pro-item"><strong class="pro-label">Advanced Digital Console:</strong> <span class="pro-description">Multi-function display with gear indicator and trip information</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Build Quality:</strong> <span class="pro-description">Reliable engineering with excellent fit and finish standards</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer SF 250 Special Edition Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Expensive Pricing:</strong> <span class="con-description">৳4,29,950 is extremely high for 250cc single-cylinder motorcycle</span></li>
<li class="con-item"><strong class="con-label">Limited Availability:</strong> <span class="con-description">Special Edition means restricted production and availability</span></li>
<li class="con-item"><strong class="con-label">No ABS Standard:</strong> <span class="con-description">Missing advanced braking technology at this price point</span></li>
<li class="con-item"><strong class="con-label">Single Cylinder Vibrations:</strong> <span class="con-description">249cc single can be buzzy at higher RPMs</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">Premium positioning means higher service and parts costs</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Suzuki Gixxer SF 250 Special Edition in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Quarter-liter sports bike enthusiasts</li>
<li class="best-for-item">Riders wanting exclusive special edition appeal</li>
<li class="best-for-item">Long-distance touring with sport bike styling</li>
<li class="best-for-item">Highway cruising and weekend rides</li>
<li class="best-for-item">Collectors of limited edition motorcycles</li>
<li class="best-for-item">Experienced riders upgrading from 150cc segment</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Suzuki Gixxer SF 250 Special Edition?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers seeking value for money</li>
<li class="not-recommended-item">New riders or beginners to motorcycling</li>
<li class="not-recommended-item">Heavy city commuting and traffic conditions</li>
<li class="not-recommended-item">Riders prioritizing fuel economy over performance</li>
<li class="not-recommended-item">Those needing extensive service network coverage</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer SF 250 Special Edition Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,29,950 - Extremely Premium for 250cc</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳12-15 per km with premium maintenance</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳12-15 per km (35-40 km/l at current fuel prices)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳8,000-12,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer SF 250 Special Edition Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent power delivery</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Decent for displacement</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good touring ergonomics</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Aggressive supersport design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.2</span> <span class="rating-note">- Overpriced for features</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent build standards</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Suzuki Gixxer SF 250 Special Edition Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What makes the SF 250 Special Edition different from regular SF 250?</h3>
<p class="faq-answer">Special Edition features exclusive graphics, premium color schemes, special badging, and enhanced styling elements not available in the standard model.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is the 249cc engine reliable for long-distance touring?</h3>
<p class="faq-answer">Yes, the SOHC 249cc engine is proven reliable for touring with good torque delivery and comfortable ergonomics, though single-cylinder vibrations may be noticeable on very long rides.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it compare to KTM Duke 250 in terms of performance?</h3>
<p class="faq-answer">Both offer similar power (26.5 HP vs 30 HP), but Duke 250 has slightly more power while Gixxer SF 250 SE offers better wind protection with full fairing and more refined ergonomics.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer SF 250 Special Edition in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,29,950, the Suzuki Gixxer SF 250 Special Edition represents the pinnacle of quarter-liter fully-faired motorcycles with exclusive special edition appeal, powerful 249cc performance, comprehensive LED package, advanced instrumentation, and premium Suzuki build quality. The 26.5 HP engine delivers strong performance while the full fairing provides excellent aerodynamics and wind protection. However, the extremely high pricing, lack of ABS at this price point, single-cylinder vibrations, expensive maintenance, and poor value proposition make it suitable only for dedicated enthusiasts who prioritize exclusivity and special edition collectability over practical value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For exclusive quarter-liter special edition collectors</span></p>
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
		return fmt.Errorf("error creating Suzuki Gixxer SF 250 Special Edition review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki Gixxer SF 250 Special Edition (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন হল সুজুকির ফ্ল্যাগশিপ ফুল-ফেয়ার্ড কোয়ার্টার-লিটার স্পোর্টস মোটরসাইকেল, যেখানে ২৪৯সিসি একক সিলিন্ডার ইঞ্জিন রয়েছে যা ২৬.৫ এইচপি এবং ২২.৬ এনএম টর্ক উৎপাদন করে। ৳৪,২৯,৯৫০ টাকায় মূল্যায়িত, এটি আক্রমণাত্মক ফুল-ফেয়ারিং অ্যারোডাইনামিক্স, প্রিমিয়াম ডিজিটাল ইনস্ট্রুমেন্টেশন, এলইডি লাইটিং প্যাকেজ এবং একচেটিয়া স্পেশাল এডিশন প্যাকেজে সুজুকির স্বাক্ষর নির্ভরযোগ্যতা একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">শক্তিশালী কোয়ার্টার-লিটার ইঞ্জিন:</strong> <span class="highlight-value">২৬.৫ এইচপি এবং ২২.৬ এনএম টর্ক সহ ২৪৯সিসি এসওএইচসি ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">একচেটিয়া স্পেশাল এডিশন:</strong> <span class="highlight-value">বিশেষ ব্যাজিং সহ প্রিমিয়াম গ্রাফিক্স এবং অনন্য রঙের স্কিম</span></li>
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ এলইডি লাইটিং প্যাকেজ:</strong> <span class="highlight-value">সম্পূর্ণ এলইডি হেডলাইট, টেইললাইট এবং ইন্ডিকেটর</span></li>
<li class="highlight-item"><strong class="highlight-label">উন্নত ডিজিটাল কনসোল:</strong> <span class="highlight-value">গিয়ার পজিশন ইন্ডিকেটর সহ মাল্টি-ফাংশন ডিসপ্লে</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী কোয়ার্টার-লিটার পারফরমেন্স:</strong> <span class="pro-description">২৪৯সিসি ইঞ্জিন চমৎকার টর্ক কার্ভ সহ শক্তিশালী ২৬.৫ এইচপি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">একচেটিয়া স্পেশাল এডিশন আবেদন:</strong> <span class="pro-description">প্রিমিয়াম গ্রাফিক্স, অনন্য রঙের স্কিম এবং বিশেষ ব্যাজিং</span></li>
<li class="pro-item"><strong class="pro-label">সম্পূর্ণ ফেয়ারিং অ্যারোডাইনামিক্স:</strong> <span class="pro-description">চমৎকার বাতাস সুরক্ষা এবং আক্রমণাত্মক সুপারস্পোর্ট স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">ব্যাপক এলইডি প্যাকেজ:</strong> <span class="pro-description">আধুনিক আবেদনের জন্য সম্পূর্ণ এলইডি হেডলাইট, টেইললাইট এবং ইন্ডিকেটর</span></li>
<li class="pro-item"><strong class="pro-label">উন্নত ডিজিটাল কনসোল:</strong> <span class="pro-description">গিয়ার ইন্ডিকেটর এবং ট্রিপ তথ্য সহ মাল্টি-ফাংশন ডিসপ্লে</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি বিল্ড কোয়ালিটি:</strong> <span class="pro-description">চমৎকার ফিট এবং ফিনিশ মান সহ নির্ভরযোগ্য ইঞ্জিনিয়ারিং</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত ব্যয়বহুল মূল্য:</strong> <span class="con-description">২৫০সিসি একক সিলিন্ডার মোটরসাইকেলের জন্য ৳৪,২৯,৯৫০ অত্যন্ত বেশি</span></li>
<li class="con-item"><strong class="con-label">সীমিত প্রাপ্যতা:</strong> <span class="con-description">স্পেশাল এডিশন মানে সীমাবদ্ধ উৎপাদন এবং প্রাপ্যতা</span></li>
<li class="con-item"><strong class="con-label">কোনো এবিএস স্ট্যান্ডার্ড:</strong> <span class="con-description">এই মূল্যে উন্নত ব্রেকিং প্রযুক্তির অভাব</span></li>
<li class="con-item"><strong class="con-label">একক সিলিন্ডার কম্পন:</strong> <span class="con-description">উচ্চ আরপিএমে ২৪৯সিসি একক কম্পনশীল হতে পারে</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">প্রিমিয়াম অবস্থান মানে উচ্চ সেবা এবং পার্টস খরচ</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Quarter-liter sports bike enthusiasts</li>
<li class="best-for-item">Riders wanting exclusive special edition appeal</li>
<li class="best-for-item">Long-distance touring with sport bike styling</li>
<li class="best-for-item">Highway cruising and weekend rides</li>
<li class="best-for-item">Collectors of limited edition motorcycles</li>
<li class="best-for-item">Experienced riders upgrading from 150cc segment</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers seeking value for money</li>
<li class="not-recommended-item">New riders or beginners to motorcycling</li>
<li class="not-recommended-item">Heavy city commuting and traffic conditions</li>
<li class="not-recommended-item">Riders prioritizing fuel economy over performance</li>
<li class="not-recommended-item">Those needing extensive service network coverage</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৪,২৯,৯৫০ - ২৫০সিসির জন্য অত্যন্ত প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">উচ্চ - প্রিমিয়াম রক্ষণাবেক্ষণ সহ ৳১২-১৫ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১২-১৫ প্রতি কিমি (বর্তমান জ্বালানি দামে ৩৫-৪০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৮,০০০-১২,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার শক্তি বিতরণ</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- ডিসপ্লেসমেন্টের জন্য ভালো</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ভালো ট্যুরিং এরগোনমিক্স</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- আক্রমণাত্মক সুপারস্পোর্ট ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.2</span> <span class="rating-note">- ফিচারের তুলনায় বেশি দামি</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- চমৎকার বিল্ড মান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এসএফ ২৫০ স্পেশাল এডিশন নিয়মিত এসএফ ২৫০ থেকে কী আলাদা?</h3>
<p class="faq-answer">স্পেশাল এডিশনে রয়েছে একচেটিয়া গ্রাফিক্স, প্রিমিয়াম রঙের স্কিম, বিশেষ ব্যাজিং এবং উন্নত স্টাইলিং উপাদান যা স্ট্যান্ডার্ড মডেলে নেই।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দীর্ঘ দূরত্বের ট্যুরিংয়ের জন্য ২৪৯সিসি ইঞ্জিন কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">হ্যাঁ, এসওএইচসি ২৪৯সিসি ইঞ্জিন ট্যুরিংয়ের জন্য প্রমাণিত নির্ভরযোগ্য যার ভাল টর্ক ডেলিভারি এবং আরামদায়ক এরগোনমিক্স আছে, যদিও একক সিলিন্ডার কম্পন খুব দীর্ঘ রাইডে লক্ষণীয় হতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">পারফরমেন্সের দিক থেকে কেটিএম ডিউক ২৫০ এর সাথে তুলনা কেমন?</h3>
<p class="faq-answer">উভয়ই একই রকম শক্তি প্রদান করে (২৬.৫ এইচপি বনাম ৩০ এইচপি), তবে ডিউক ২৫০ এ সামান্য বেশি শক্তি আছে যখন গিক্সার এসএফ ২৫০ এসই ফুল ফেয়ারিং সহ ভাল বাতাস সুরক্ষা এবং আরো পরিশীলিত এরগোনমিক্স প্রদান করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৪,২৯,৯৫০ টাকায় সুজুকি গিক্সার এসএফ ২৫০ স্পেশাল এডিশন একচেটিয়া স্পেশাল এডিশন আবেদন, শক্তিশালী ২৪৯সিসি পারফরমেন্স, ব্যাপক এলইডি প্যাকেজ, অ্যাডভান্সড ইনস্ট্রুমেন্টেশন এবং প্রিমিয়াম সুজুকি বিল্ড কোয়ালিটি সহ কোয়ার্টার-লিটার ফুল-ফেয়ার্ড মোটরসাইকেলের শিখরকে প্রতিনিধিত্ব করে। ২৬.৫ এইচপি ইঞ্জিন শক্তিশালী পারফরমেন্স প্রদান করে যখন ফুল ফেয়ারিং চমৎকার অ্যারোডাইনামিক্স এবং বাতাস সুরক্ষা প্রদান করে। তবে, অত্যন্ত উচ্চ মূল্য, এই মূল্যে এবিএসের অভাব, একক সিলিন্ডার কম্পন, ব্যয়বহুল রক্ষণাবেক্ষণ এবং খারাপ মূল্য প্রস্তাব এটিকে কেবলমাত্র নিবেদিতপ্রাণ উৎসাহীদের জন্য উপযুক্ত করে তোলে যারা ব্যবহারিক মূল্যের চেয়ে একচেটিয়াতা এবং স্পেশাল এডিশন সংগ্রহযোগ্যতাকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- একচেটিয়া কোয়ার্টার-লিটার স্পেশাল এডিশন সংগ্রাহকদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki Gixxer SF 250 Special Edition already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Suzuki Gixxer SF 250 Special Edition\n")

	return nil
}
