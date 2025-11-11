package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCrf150LReview handles seeding Honda CRF 150L product review and translation
type HondaCrf150LReview struct {
	BaseSeeder
}

// NewHondaCrf150LReviewSeeder creates a new HondaCrf150LReview
func NewHondaCrf150LReviewSeeder() *HondaCrf150LReview {
	return &HondaCrf150LReview{BaseSeeder: BaseSeeder{name: "Honda CRF 150L Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCrf150LReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CRF 150L").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda CRF 150L product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda CRF 150L product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CRF 150L already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CRF 150L Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda CRF 150L is a dual-sport enduro motorcycle featuring off-road capability with competition-inspired CRF styling, long-travel suspension, high ground clearance, and lightweight design. Priced at ৳4,95,000, it offers genuine off-road performance with rally-inspired aesthetics, excellent ground clearance for rough terrain, capable suspension, Honda reliability, and versatile dual-purpose capability for riders seeking authentic trail riding and adventure exploration beyond paved roads.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CRF 150L Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">CRF Rally Heritage:</strong> <span class="highlight-value">Competition-inspired design from CRF racing lineage</span></li>
<li class="highlight-item"><strong class="highlight-label">Long-Travel Suspension:</strong> <span class="highlight-value">265mm front and 245mm rear travel for off-road</span></li>
<li class="highlight-item"><strong class="highlight-label">Exceptional Ground Clearance:</strong> <span class="highlight-value">285mm clearance for serious off-road terrain</span></li>
<li class="highlight-item"><strong class="highlight-label">Lightweight Design:</strong> <span class="highlight-value">Light construction for agile trail handling</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CRF 150L Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Genuine Off-Road Capability:</strong> <span class="pro-description">Long-travel suspension and high clearance handle serious trails</span></li>
<li class="pro-item"><strong class="pro-label">Exceptional Ground Clearance:</strong> <span class="pro-description">285mm clearance conquers obstacles and rough terrain easily</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight and Agile:</strong> <span class="pro-description">Light construction makes trail riding and maneuvering easier</span></li>
<li class="pro-item"><strong class="pro-label">Rally-Inspired Styling:</strong> <span class="pro-description">Authentic CRF competition aesthetics and presence</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven Honda engineering for dependable performance</span></li>
<li class="pro-item"><strong class="pro-label">Versatile Dual-Purpose:</strong> <span class="pro-description">Capable on both roads and off-road trails</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CRF 150L Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Tall Seat Height:</strong> <span class="con-description">870mm seat height extremely challenging for shorter riders</span></li>
<li class="con-item"><strong class="con-label">Uncomfortable for Street Use:</strong> <span class="con-description">Off-road focus makes street riding less comfortable</span></li>
<li class="con-item"><strong class="con-label">Premium Pricing:</strong> <span class="con-description">৳4,95,000 is expensive for 150cc dual-sport</span></li>
<li class="con-item"><strong class="con-label">Limited City Use:</strong> <span class="con-description">Tall design and off-road focus impractical for urban commuting</span></li>
<li class="con-item"><strong class="con-label">Hard Seat:</strong> <span class="con-description">Firm off-road seat uncomfortable for long road rides</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda CRF 150L in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Off-road trail riding enthusiasts</li>
<li class="best-for-item">Adventure riders seeking genuine capability</li>
<li class="best-for-item">Tall riders comfortable with high seat height</li>
<li class="best-for-item">Weekend trail exploration</li>
<li class="best-for-item">Dual-sport riders wanting real off-road performance</li>
<li class="best-for-item">CRF rally heritage fans</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda CRF 150L?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Shorter riders (under 5'8")</li>
<li class="not-recommended-item">Daily city commuters</li>
<li class="not-recommended-item">Those seeking comfortable street riding</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Riders without access to off-road trails</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CRF 150L Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,95,000 - Premium for 150cc dual-sport enduro</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-11 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳7-9 per km (35-40 km/l off-road focused)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 4 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳12,000-16,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CRF 150L Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Excellent off-road performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- Moderate fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.3</span> <span class="rating-note">- Off-road focused, firm ride</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Authentic CRF rally styling</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Premium for dual-sport</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent Honda quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda CRF 150L Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Can beginners ride the CRF 150L?</h3>
<p class="faq-answer">The CRF 150L's very tall seat height (870mm) makes it challenging for beginners, especially shorter riders. It's better suited for experienced riders comfortable with tall enduro bikes and off-road riding techniques.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for daily commuting?</h3>
<p class="faq-answer">The CRF 150L is not ideal for daily city commuting due to its tall height, firm off-road seat, and enduro-focused setup. It's better suited for weekend trail riding and adventure exploration rather than urban traffic.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How capable is it on serious off-road trails?</h3>
<p class="faq-answer">The CRF 150L is genuinely capable on off-road trails with 285mm ground clearance, long-travel suspension (265mm front/245mm rear), and lightweight design. It can handle serious trail riding, though larger displacement bikes offer more power for extreme terrain.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CRF 150L in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,95,000, the Honda CRF 150L offers genuine dual-sport enduro capability with authentic CRF rally heritage, exceptional ground clearance (285mm), long-travel suspension, lightweight design, and Honda reliability. It excels at off-road trail riding and adventure exploration with real capability beyond paved roads. However, the very tall seat height (870mm) extremely challenging for shorter riders, uncomfortable street use, premium pricing, limited city practicality, and firm seat make it suitable only for tall experienced riders, off-road enthusiasts, and adventure seekers who prioritize genuine trail capability over street comfort.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For genuine off-road trail capability with CRF rally heritage</span></p>
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
		return fmt.Errorf("error creating Honda CRF 150L review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CRF 150L (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা সিআরএফ ১৫০এল রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা সিআরএফ ১৫০এল হল একটি ডুয়াল-স্পোর্ট এন্ডুরো মোটরসাইকেল যেখানে প্রতিযোগিতা-অনুপ্রাণিত সিআরএফ স্টাইলিং, লং-ট্র্যাভেল সাসপেনশন, উচ্চ গ্রাউন্ড ক্লিয়ারেন্স এবং লাইটওয়েট ডিজাইন সহ অফ-রোড ক্ষমতা রয়েছে। ৳৪,৯৫,০০০ টাকায় মূল্যায়িত, এটি র‍্যালি-অনুপ্রাণিত নান্দনিকতা, রুক্ষ ভূখণ্ডের জন্য চমৎকার গ্রাউন্ড ক্লিয়ারেন্স, সক্ষম সাসপেনশন, হোন্ডা নির্ভরযোগ্যতা এবং পাকা রাস্তার বাইরে প্রকৃত ট্রেইল রাইডিং এবং অ্যাডভেঞ্চার অন্বেষণ খোঁজা রাইডারদের জন্য বহুমুখী দ্বৈত-উদ্দেশ্য ক্ষমতা সহ প্রকৃত অফ-রোড পারফরমেন্স প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা সিআরএফ ১৫০এল এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সিআরএফ র‍্যালি ঐতিহ্য:</strong> <span class="highlight-value">সিআরএফ রেসিং বংশ থেকে প্রতিযোগিতা-অনুপ্রাণিত ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">লং-ট্র্যাভেল সাসপেনশন:</strong> <span class="highlight-value">অফ-রোডের জন্য ২৬৫মিমি সামনে এবং ২৪৫মিমি পিছনে ট্র্যাভেল</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যতিক্রমী গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="highlight-value">গুরুতর অফ-রোড ভূখণ্ডের জন্য ২৮৫মিমি ক্লিয়ারেন্স</span></li>
<li class="highlight-item"><strong class="highlight-label">লাইটওয়েট ডিজাইন:</strong> <span class="highlight-value">চটপটে ট্রেইল হ্যান্ডলিংয়ের জন্য হালকা নির্মাণ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা সিআরএফ ১৫০এল এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রকৃত অফ-রোড ক্ষমতা:</strong> <span class="pro-description">লং-ট্র্যাভেল সাসপেনশন এবং উচ্চ ক্লিয়ারেন্স গুরুতর ট্রেইল পরিচালনা করে</span></li>
<li class="pro-item"><strong class="pro-label">ব্যতিক্রমী গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="pro-description">২৮৫মিমি ক্লিয়ারেন্স সহজে বাধা এবং রুক্ষ ভূখণ্ড জয় করে</span></li>
<li class="pro-item"><strong class="pro-label">লাইটওয়েট এবং চটপটে:</strong> <span class="pro-description">হালকা নির্মাণ ট্রেইল রাইডিং এবং চালনা সহজ করে</span></li>
<li class="pro-item"><strong class="pro-label">র‍্যালি-অনুপ্রাণিত স্টাইলিং:</strong> <span class="pro-description">প্রকৃত সিআরএফ প্রতিযোগিতা নান্দনিকতা এবং উপস্থিতি</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">নির্ভরযোগ্য পারফরমেন্সের জন্য প্রমাণিত হোন্ডা ইঞ্জিনিয়ারিং</span></li>
<li class="pro-item"><strong class="pro-label">বহুমুখী দ্বৈত-উদ্দেশ্য:</strong> <span class="pro-description">রাস্তা এবং অফ-রোড ট্রেইল উভয়েই সক্ষম</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা সিআরএফ ১৫০এল এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অত্যন্ত লম্বা সিট উচ্চতা:</strong> <span class="con-description">৮৭০মিমি সিট উচ্চতা খাটো রাইডারদের জন্য অত্যন্ত চ্যালেঞ্জিং</span></li>
<li class="con-item"><strong class="con-label">রাস্তা ব্যবহারের জন্য অস্বস্তিকর:</strong> <span class="con-description">অফ-রোড ফোকাস রাস্তা রাইডিং কম আরামদায়ক করে</span></li>
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳৪,৯৫,০০০ ১৫০সিসি ডুয়াল-স্পোর্টের জন্য ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সীমিত শহুরে ব্যবহার:</strong> <span class="con-description">লম্বা ডিজাইন এবং অফ-রোড ফোকাস শহুরে যাতায়াতের জন্য অব্যবহারিক</span></li>
<li class="con-item"><strong class="con-label">শক্ত সিট:</strong> <span class="con-description">শক্ত অফ-রোড সিট দীর্ঘ রাস্তা রাইডের জন্য অস্বস্তিকর</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা সিআরএফ ১৫০এল কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Off-road trail riding enthusiasts</li>
<li class="best-for-item">Adventure riders seeking genuine capability</li>
<li class="best-for-item">Tall riders comfortable with high seat height</li>
<li class="best-for-item">Weekend trail exploration</li>
<li class="best-for-item">Dual-sport riders wanting real off-road performance</li>
<li class="best-for-item">CRF rally heritage fans</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা সিআরএফ ১৫০এল কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Shorter riders (under 5'8")</li>
<li class="not-recommended-item">Daily city commuters</li>
<li class="not-recommended-item">Those seeking comfortable street riding</li>
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Riders without access to off-road trails</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা সিআরএফ ১৫০এল এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৪,৯৫,০০০ - ১৫০সিসি ডুয়াল-স্পোর্ট এন্ডুরোর জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি - ৳৮-১১ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৭-৯ প্রতি কিমি (অফ-রোড ফোকাসড ৩৫-৪০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৪ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১২,০০০-১৬,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা সিআরএফ ১৫০এল পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- চমৎকার অফ-রোড পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.6</span> <span class="rating-note">- মাঝারি জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.3</span> <span class="rating-note">- অফ-রোড ফোকাসড, শক্ত রাইড</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- প্রকৃত সিআরএফ র‍্যালি স্টাইলিং</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- ডুয়াল-স্পোর্টের জন্য প্রিমিয়াম</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার হোন্ডা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা সিআরএফ ১৫০এল সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">নতুনরা কি সিআরএফ ১৫০এল চালাতে পারে?</h3>
<p class="faq-answer">সিআরএফ ১৫০এল এর অত্যন্ত লম্বা সিট উচ্চতা (৮৭০মিমি) এটিকে নতুনদের জন্য চ্যালেঞ্জিং করে তোলে, বিশেষত খাটো রাইডারদের। এটি লম্বা এন্ডুরো বাইক এবং অফ-রোড রাইডিং কৌশলে আরামদায়ক অভিজ্ঞ রাইডারদের জন্য ভাল উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি দৈনিক যাতায়াতের জন্য উপযুক্ত?</h3>
<p class="faq-answer">সিআরএফ ১৫০এল এর লম্বা উচ্চতা, শক্ত অফ-রোড সিট এবং এন্ডুরো-ফোকাসড সেটআপের কারণে দৈনিক শহুরে যাতায়াতের জন্য আদর্শ নয়। এটি শহুরে ট্রাফিকের চেয়ে সাপ্তাহিক ট্রেইল রাইডিং এবং অ্যাডভেঞ্চার অন্বেষণের জন্য ভাল উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">গুরুতর অফ-রোড ট্রেইলে এটি কতটা সক্ষম?</h3>
<p class="faq-answer">সিআরএফ ১৫০এল ২৮৫মিমি গ্রাউন্ড ক্লিয়ারেন্স, লং-ট্র্যাভেল সাসপেনশন (২৬৫মিমি সামনে/২৪৫মিমি পিছনে) এবং লাইটওয়েট ডিজাইন সহ অফ-রোড ট্রেইলে প্রকৃতপক্ষে সক্ষম। এটি গুরুতর ট্রেইল রাইডিং পরিচালনা করতে পারে, যদিও বৃহত্তর ডিসপ্লেসমেন্ট বাইক চরম ভূখণ্ডের জন্য আরও শক্তি প্রদান করে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা সিআরএফ ১৫০এল কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৪,৯৫,০০০ টাকায় হোন্ডা সিআরএফ ১৫০এল প্রকৃত সিআরএফ র‍্যালি ঐতিহ্য, ব্যতিক্রমী গ্রাউন্ড ক্লিয়ারেন্স (২৮৫মিমি), লং-ট্র্যাভেল সাসপেনশন, লাইটওয়েট ডিজাইন এবং হোন্ডা নির্ভরযোগ্যতা সহ প্রকৃত ডুয়াল-স্পোর্ট এন্ডুরো ক্ষমতা প্রদান করে। এটি পাকা রাস্তার বাইরে প্রকৃত ক্ষমতা সহ অফ-রোড ট্রেইল রাইডিং এবং অ্যাডভেঞ্চার অন্বেষণে শ্রেষ্ঠ। তবে, খাটো রাইডারদের জন্য অত্যন্ত চ্যালেঞ্জিং অত্যন্ত লম্বা সিট উচ্চতা (৮৭০মিমি), অস্বস্তিকর রাস্তা ব্যবহার, প্রিমিয়াম মূল্য, সীমিত শহুরে ব্যবহারযোগ্যতা এবং শক্ত সিট এটিকে শুধুমাত্র লম্বা অভিজ্ঞ রাইডার, অফ-রোড উৎসাহী এবং অ্যাডভেঞ্চার অনুসন্ধানকারীদের জন্য উপযুক্ত করে তোলে যারা রাস্তা আরামের চেয়ে প্রকৃত ট্রেইল ক্ষমতাকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সিআরএফ র‍্যালি ঐতিহ্য সহ প্রকৃত অফ-রোড ট্রেইল ক্ষমতার জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CRF 150L already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CRF 150L\n")

	return nil
}
