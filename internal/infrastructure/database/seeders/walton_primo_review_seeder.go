package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// WaltonPrimoReviewSeeder handles seeding Walton Primo product review and translation
type WaltonPrimoReviewSeeder struct {
	BaseSeeder
}

// NewWaltonPrimoReviewSeeder creates a new WaltonPrimoReviewSeeder
func NewWaltonPrimoReviewSeeder() *WaltonPrimoReviewSeeder {
	return &WaltonPrimoReviewSeeder{BaseSeeder: BaseSeeder{name: "Walton Primo Review"}}
}

// Seed implements the Seeder interface
func (s *WaltonPrimoReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Walton Primo").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Walton Primo product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Walton Primo product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Walton Primo already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Walton Primo Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Walton Primo is a local Bangladeshi brand 100cc bike priced at ৳98,000, offering budget pricing with slightly better quality than Runner bikes. However, reliability concerns, limited service network, and poor resale value make it hard to recommend when Honda CD 70 or Hero HF Deluxe cost only marginally more.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Walton Primo Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Local Brand:</strong> <span class="highlight-value">Bangladeshi Walton brand</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Pricing:</strong> <span class="highlight-value">Affordable at ৳98,000</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">55-60 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Better than Runner:</strong> <span class="highlight-value">Slightly better quality</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Walton Primo Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">Budget-friendly at ৳98,000</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Economy:</strong> <span class="pro-description">55-60 km/l mileage</span></li>
<li class="pro-item"><strong class="pro-label">Local Service:</strong> <span class="pro-description">Walton service centers available</span></li>
<li class="pro-item"><strong class="pro-label">Better than Runner:</strong> <span class="pro-description">Slightly better build quality</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight:</strong> <span class="pro-description">Easy to handle</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Walton Primo Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Questionable Reliability:</strong> <span class="con-description">Not as reliable as Japanese brands</span></li>
<li class="con-item"><strong class="con-label">Poor Resale Value:</strong> <span class="con-description">Local brands have low resale</span></li>
<li class="con-item"><strong class="con-label">Basic Build Quality:</strong> <span class="con-description">Lower quality than Honda/Hero</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Not as widespread as Hero</span></li>
<li class="con-item"><strong class="con-label">No Brand Heritage:</strong> <span class="con-description">New to motorcycle market</span></li>
<li class="con-item"><strong class="con-label">Spare Parts Concerns:</strong> <span class="con-description">Long-term availability uncertain</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Walton Primo in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Buyers supporting local brands</li>
<li class="best-for-item">Those on tight budget but avoiding Runner</li>
<li class="best-for-item">Short-distance city commuting</li>
<li class="best-for-item">Areas with Walton service centers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Walton Primo?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Those prioritizing reliability</li>
<li class="not-recommended-item">Long-term ownership plans</li>
<li class="not-recommended-item">Buyers needing resale value</li>
<li class="not-recommended-item">Anyone who can afford Honda CD 70 or Hero HF Deluxe</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Walton Primo Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳95,000 - ৳1,02,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,000 - ৳2,800 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,600-1,900/month for 30 km daily at 57 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 2,500 km or 2-3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳500 - ৳900 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Walton Primo Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Questionable</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Better options exist</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic performance</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Below average</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very basic</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Simple design</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Walton Primo Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Is Walton Primo reliable?</h3>
<p class="faq-answer">Moderately reliable but not as proven as Honda or Hero bikes.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Should I buy Walton Primo?</h3>
<p class="faq-answer">Only if budget is very tight and you want to support local brands. Honda CD 70 or Hero HF Deluxe are better investments for slightly more money.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Walton Primo?</h3>
<p class="faq-answer">Walton Primo delivers 55-60 km/l fuel efficiency.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Walton Primo?</h3>
<p class="faq-answer">Walton Primo is priced between ৳95,000 to ৳1,02,000.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Walton Primo in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Walton Primo at ৳98,000 is better than Runner bikes but still not recommended when Honda CD 70 (৳85,000) or Hero HF Deluxe (৳1,30,000) are available. While supporting local brands is admirable, questionable reliability, poor resale value, and limited proven track record make it risky. For just ৳30-35,000 more, Hero HF Deluxe offers legendary reliability and proper service network. Only consider if strongly preferring local brands.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- Only for local brand supporters with budget constraints</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Walton Primo review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Walton Primo (Rating: 3.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ওয়ালটন প্রিমো রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ওয়ালটন প্রিমো একটি স্থানীয় বাংলাদেশী ব্র্যান্ড ১০০সিসি বাইক যার মূল্য ৳৯৮,০০০ টাকা, যা রানার বাইকের চেয়ে সামান্য ভালো মানের সাথে বাজেট মূল্য প্রদান করে। তবে, নির্ভরযোগ্যতা উদ্বেগ, সীমিত সার্ভিস নেটওয়ার্ক এবং খারাপ রিসেল ভ্যালু এটিকে সুপারিশ করা কঠিন করে তোলে যখন হোন্ডা সিডি ৭০ বা হিরো এইচএফ ডিলাক্স শুধুমাত্র সামান্য বেশি খরচ হয়।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ওয়ালটন প্রিমো এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">স্থানীয় ব্র্যান্ড:</strong> <span class="highlight-value">বাংলাদেশী ওয়ালটন ব্র্যান্ড</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট মূল্য:</strong> <span class="highlight-value">৳৯৮,০০০ টাকায় সাশ্রয়ী</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো মাইলেজ:</strong> <span class="highlight-value">৫৫-৬০ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
<li class="highlight-item"><strong class="highlight-label">রানারের চেয়ে ভালো:</strong> <span class="highlight-value">সামান্য ভালো মান</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ওয়ালটন প্রিমো এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">৳৯৮,০০০ টাকায় বাজেট-বান্ধব</span></li>
<li class="pro-item"><strong class="pro-label">ভালো জ্বালানি সাশ্রয়:</strong> <span class="pro-description">৫৫-৬০ কিমি/লিটার মাইলেজ</span></li>
<li class="pro-item"><strong class="pro-label">স্থানীয় সার্ভিস:</strong> <span class="pro-description">ওয়ালটন সার্ভিস সেন্টার উপলব্ধ</span></li>
<li class="pro-item"><strong class="pro-label">রানারের চেয়ে ভালো:</strong> <span class="pro-description">সামান্য ভালো বিল্ড কোয়ালিটি</span></li>
<li class="pro-item"><strong class="pro-label">হালকা ওজন:</strong> <span class="pro-description">হ্যান্ডল করা সহজ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ওয়ালটন প্রিমো এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রশ্নবিদ্ধ নির্ভরযোগ্যতা:</strong> <span class="con-description">জাপানি ব্র্যান্ডের মতো নির্ভরযোগ্য নয়</span></li>
<li class="con-item"><strong class="con-label">খারাপ রিসেল ভ্যালু:</strong> <span class="con-description">স্থানীয় ব্র্যান্ডের কম রিসেল আছে</span></li>
<li class="con-item"><strong class="con-label">বেসিক বিল্ড কোয়ালিটি:</strong> <span class="con-description">হোন্ডা/হিরোর চেয়ে কম মান</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস নেটওয়ার্ক:</strong> <span class="con-description">হিরোর মতো ব্যাপক নয়</span></li>
<li class="con-item"><strong class="con-label">কোনো ব্র্যান্ড ঐতিহ্য নেই:</strong> <span class="con-description">মোটরসাইকেল বাজারে নতুন</span></li>
<li class="con-item"><strong class="con-label">স্পেয়ার পার্টস উদ্বেগ:</strong> <span class="con-description">দীর্ঘমেয়াদী প্রাপ্যতা অনিশ্চিত</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ওয়ালটন প্রিমো কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Buyers supporting local brands</li>
<li class="best-for-item">Those on tight budget but avoiding Runner</li>
<li class="best-for-item">Short-distance city commuting</li>
<li class="best-for-item">Areas with Walton service centers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ওয়ালটন প্রিমো কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Those prioritizing reliability</li>
<li class="not-recommended-item">Long-term ownership plans</li>
<li class="not-recommended-item">Buyers needing resale value</li>
<li class="not-recommended-item">Anyone who can afford Honda CD 70 or Hero HF Deluxe</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ওয়ালটন প্রিমো এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৯৫,০০০ - ৳১,০২,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,০০০ - ৳২,৮০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳১,৬০০-১,৯০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ২,৫০০ কিমি বা ২-৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৫০০ - ৳৯০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ওয়ালটন প্রিমো পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- প্রশ্নবিদ্ধ</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- ভালো অপশন আছে</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- বেসিক পারফরম্যান্স</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- বেসিক আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- গড়ের নিচে</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- খুব বেসিক</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- সাধারণ ডিজাইন</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ওয়ালটন প্রিমো সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ওয়ালটন প্রিমো কি নির্ভরযোগ্য?</h3>
<p class="faq-answer">মাঝারি নির্ভরযোগ্য কিন্তু হোন্ডা বা হিরো বাইকের মতো প্রমাণিত নয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">আমার কি ওয়ালটন প্রিমো কিনা উচিত?</h3>
<p class="faq-answer">শুধুমাত্র যদি বাজেট খুব টাইট হয় এবং আপনি স্থানীয় ব্র্যান্ডকে সমর্থন করতে চান। হোন্ডা সিডি ৭০ বা হিরো এইচএফ ডিলাক্স সামান্য বেশি টাকায় ভালো বিনিয়োগ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ওয়ালটন প্রিমোর মাইলেজ কত?</h3>
<p class="faq-answer">ওয়ালটন প্রিমো ৫৫-৬০ কিমি/লিটার জ্বালানি দক্ষতা প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ওয়ালটন প্রিমোর দাম কত?</h3>
<p class="faq-answer">ওয়ালটন প্রিমোর দাম ৳৯৫,০০০ থেকে ৳১,০২,০০০ টাকার মধ্যে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ওয়ালটন প্রিমো কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">3.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৯৮,০০০ টাকায় ওয়ালটন প্রিমো রানার বাইকের চেয়ে ভালো কিন্তু এখনও সুপারিশ করা হয় না যখন হোন্ডা সিডি ৭০ (৳৮৫,০০০) বা হিরো এইচএফ ডিলাক্স (৳১,৩০,০০০) উপলব্ধ। যদিও স্থানীয় ব্র্যান্ডকে সমর্থন করা প্রশংসনীয়, প্রশ্নবিদ্ধ নির্ভরযোগ্যতা, খারাপ রিসেল ভ্যালু এবং সীমিত প্রমাণিত ট্র্যাক রেকর্ড এটিকে ঝুঁকিপূর্ণ করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শুধুমাত্র বাজেট সীমাবদ্ধতা সহ স্থানীয় ব্র্যান্ড সমর্থকদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Walton Primo already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Walton Primo\n")

	return nil
}
