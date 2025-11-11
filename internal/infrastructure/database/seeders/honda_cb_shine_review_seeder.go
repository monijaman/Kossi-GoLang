package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCBShineReviewSeeder handles seeding Honda CB Shine product review and translation
type HondaCBShineReviewSeeder struct {
	BaseSeeder
}

// NewHondaCBShineReviewSeeder creates a new HondaCBShineReviewSeeder
func NewHondaCBShineReviewSeeder() *HondaCBShineReviewSeeder {
	return &HondaCBShineReviewSeeder{BaseSeeder: BaseSeeder{name: "Honda CB Shine Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCBShineReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB Shine").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda CB Shine product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda CB Shine product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CB Shine already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CB Shine Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda CB Shine is a premium 125cc motorcycle known for its refined engine, excellent build quality, and comfortable ride. Priced at ৳1,65,000, it offers smooth performance with Honda's legendary reliability, making it ideal for daily commuters who value quality and comfort.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB Shine Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Refined Engine:</strong> <span class="highlight-value">Ultra-smooth 125cc engine with minimal vibrations</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Comfort:</strong> <span class="highlight-value">Best-in-class comfortable seat and riding position</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Mileage:</strong> <span class="highlight-value">Delivers 55-60 km/l fuel efficiency</span></li>
<li class="highlight-item"><strong class="highlight-label">Honda Reliability:</strong> <span class="highlight-value">Proven quality with minimal maintenance</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB Shine Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Ultra-Smooth Engine:</strong> <span class="pro-description">Refined 125cc engine with class-leading smoothness</span></li>
<li class="pro-item"><strong class="pro-label">Best-in-Class Comfort:</strong> <span class="pro-description">Plush seat and upright ergonomics for long rides</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Build Quality:</strong> <span class="pro-description">Premium materials and solid construction</span></li>
<li class="pro-item"><strong class="pro-label">Great Fuel Economy:</strong> <span class="pro-description">55-60 km/l keeps running costs low</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Performance:</strong> <span class="pro-description">Honda's proven engine rarely has issues</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Minimal service requirements save money</span></li>
<li class="pro-item"><strong class="pro-label">Strong Resale Value:</strong> <span class="pro-description">Honda brand ensures good used market demand</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Convenient self-start with kick backup</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Honda centers available across Bangladesh</span></li>
<li class="pro-item"><strong class="pro-label">Good Ground Clearance:</strong> <span class="pro-description">Handles rough roads and speed breakers well</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB Shine Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Expensive Price:</strong> <span class="con-description">₳1,65,000 is premium pricing for 125cc</span></li>
<li class="con-item"><strong class="con-label">Drum Brakes Only:</strong> <span class="con-description">No disc brake option available</span></li>
<li class="con-item"><strong class="con-label">Conservative Styling:</strong> <span class="con-description">Plain design lacks sporty appeal</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">No USB charging or modern amenities</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">Heavier than competitors, less maneuverable</span></li>
<li class="con-item"><strong class="con-label">Analog Meter:</strong> <span class="con-description">Traditional instrument cluster, not digital</span></li>
<li class="con-item"><strong class="con-label">Lower Power:</strong> <span class="con-description">Less powerful than some 125cc rivals</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda CB Shine in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters prioritizing comfort and smoothness</li>
<li class="best-for-item">Mature riders seeking refined, vibration-free experience</li>
<li class="best-for-item">Those willing to pay premium for Honda quality</li>
<li class="best-for-item">Long-distance city commuters (40-60 km daily)</li>
<li class="best-for-item">Riders valuing reliability over sporty performance</li>
<li class="best-for-item">First-time premium bike buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda CB Shine?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers (cheaper options available)</li>
<li class="not-recommended-item">Young riders seeking sporty styling</li>
<li class="not-recommended-item">Those wanting disc brakes and modern features</li>
<li class="not-recommended-item">Performance enthusiasts (not very powerful)</li>
<li class="not-recommended-item">Riders looking for best value for money</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB Shine Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,60,000 - ৳1,70,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,500 - ৳3,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,000-2,400/month for 30 km daily at 57 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,000 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB Shine Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Honda quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Premium pricing</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Smooth but not powerful</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Best-in-class</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Premium Honda build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic equipment</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Conservative look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda CB Shine Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Honda CB Shine?</h3>
<p class="faq-answer">Honda CB Shine delivers 55-60 km/l in real-world conditions, making it one of the most fuel-efficient 125cc bikes.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Honda CB Shine in Bangladesh 2024?</h3>
<p class="faq-answer">Honda CB Shine is priced between ৳1,60,000 to ৳1,70,000 in Bangladesh.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Honda CB Shine have disc brake?</h3>
<p class="faq-answer">No, Honda CB Shine comes with drum brakes on both wheels. No disc brake variant is available in Bangladesh.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Honda CB Shine good for daily commuting?</h3>
<p class="faq-answer">Yes, Honda CB Shine is excellent for daily commuting with its ultra-smooth engine, comfortable seat, and 55-60 km/l mileage.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Honda CB Shine?</h3>
<p class="faq-answer">Honda CB Shine can reach a top speed of approximately 100-105 km/h, suitable for city and occasional highway riding.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How much does Honda CB Shine service cost?</h3>
<p class="faq-answer">Regular service costs ৳1,000 to ৳1,500 at authorized Honda service centers every 3,000 km or 3 months.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is CB Shine better than CB125F?</h3>
<p class="faq-answer">CB Shine offers better comfort and smoother engine but CB125F provides more power. Choose Shine for comfort, CB125F for performance.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Honda CB Shine have electric start?</h3>
<p class="faq-answer">Yes, Honda CB Shine comes with both electric start and kick start as backup for added convenience.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB Shine in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Honda CB Shine is the ultimate choice for riders prioritizing comfort, smoothness, and reliability over everything else. While its premium price of ৳1,65,000 may seem high, the ultra-smooth engine, best-in-class comfort, and Honda's legendary reliability justify the cost for serious commuters. The lack of disc brakes and sporty features won't matter if you value a refined, vibration-free riding experience. Ideal for mature riders doing long daily commutes who can afford to pay for premium quality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For comfort-focused daily commuters</span></p>
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
		return fmt.Errorf("error creating Honda CB Shine review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CB Shine (Rating: 4.2/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা সিবি শাইন রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা সিবি শাইন একটি প্রিমিয়াম ১২৫সিসি মোটরসাইকেল যা এর পরিশোধিত ইঞ্জিন, চমৎকার বিল্ড কোয়ালিটি এবং আরামদায়ক রাইডের জন্য পরিচিত। ৳১,৬৫,০০০ টাকায় মূল্যবান, এটি হোন্ডার কিংবদন্তী নির্ভরযোগ্যতার সাথে মসৃণ পারফরম্যান্স প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা সিবি শাইন এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">পরিশোধিত ইঞ্জিন:</strong> <span class="highlight-value">ন্যূনতম কম্পন সহ অত্যন্ত মসৃণ ১২৫সিসি ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম আরাম:</strong> <span class="highlight-value">ক্লাসের সেরা আরামদায়ক সিট এবং রাইডিং পজিশন</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার মাইলেজ:</strong> <span class="highlight-value">৫৫-৬০ কিমি/লিটার জ্বালানি দক্ষতা প্রদান করে</span></li>
<li class="highlight-item"><strong class="highlight-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="highlight-value">ন্যূনতম রক্ষণাবেক্ষণ সহ প্রমাণিত গুণমান</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা সিবি শাইন এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অত্যন্ত মসৃণ ইঞ্জিন:</strong> <span class="pro-description">ক্লাস-লিডিং স্মুথনেস সহ পরিশোধিত ১২৫সিসি ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">ক্লাসে সেরা আরাম:</strong> <span class="pro-description">লম্বা রাইডের জন্য প্লাশ সিট এবং সোজা এরগোনমিক্স</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম উপকরণ এবং কঠিন নির্মাণ</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত জ্বালানি সাশ্রয়:</strong> <span class="pro-description">৫৫-৬০ কিমি/লিটার রানিং খরচ কম রাখে</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য পারফরম্যান্স:</strong> <span class="pro-description">হোন্ডার প্রমাণিত ইঞ্জিনে খুব কমই সমস্যা হয়</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">ন্যূনতম সার্ভিস প্রয়োজন অর্থ সাশ্রয় করে</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী রিসেল ভ্যালু:</strong> <span class="pro-description">হোন্ডা ব্র্যান্ড ভালো পুরাতন বাজার চাহিদা নিশ্চিত করে</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">কিক ব্যাকআপ সহ সুবিধাজনক সেল্ফ-স্টার্ট</span></li>
<li class="pro-item"><strong class="pro-label">বিস্তৃত সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সারা বাংলাদেশে হোন্ডা সেন্টার উপলব্ধ</span></li>
<li class="pro-item"><strong class="pro-label">ভালো গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="pro-description">রুক্ষ রাস্তা এবং স্পিড ব্রেকার ভালভাবে সামলায়</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা সিবি শাইন এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">ব্যয়বহুল দাম:</strong> <span class="con-description">১২৫সিসির জন্য ৳১,৬৫,০০০ টাকা প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">শুধু ড্রাম ব্রেক:</strong> <span class="con-description">কোনো ডিস্ক ব্রেক বিকল্প উপলব্ধ নেই</span></li>
<li class="con-item"><strong class="con-label">রক্ষণশীল স্টাইলিং:</strong> <span class="con-description">সাধারণ ডিজাইনে স্পোর্টি আবেদনের অভাব</span></li>
<li class="con-item"><strong class="con-label">বেসিক ফিচার:</strong> <span class="con-description">কোনো ইউএসবি চার্জিং বা আধুনিক সুবিধা নেই</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">প্রতিযোগীদের চেয়ে ভারী, কম চালিত</span></li>
<li class="con-item"><strong class="con-label">অ্যানালগ মিটার:</strong> <span class="con-description">ঐতিহ্যবাহী ইন্সট্রুমেন্ট ক্লাস্টার, ডিজিটাল নয়</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">কিছু ১২৫সিসি প্রতিদ্বন্দ্বীর চেয়ে কম শক্তিশালী</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা সিবি শাইন কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters prioritizing comfort and smoothness</li>
<li class="best-for-item">Mature riders seeking refined, vibration-free experience</li>
<li class="best-for-item">Those willing to pay premium for Honda quality</li>
<li class="best-for-item">Long-distance city commuters (40-60 km daily)</li>
<li class="best-for-item">Riders valuing reliability over sporty performance</li>
<li class="best-for-item">First-time premium bike buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা সিবি শাইন কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers (cheaper options available)</li>
<li class="not-recommended-item">Young riders seeking sporty styling</li>
<li class="not-recommended-item">Those wanting disc brakes and modern features</li>
<li class="not-recommended-item">Performance enthusiasts (not very powerful)</li>
<li class="not-recommended-item">Riders looking for best value for money</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা সিবি শাইন এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৬০,০০০ - ৳১,৭০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৫০০ - ৳৩,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৫৭ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,০০০-২,৪০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,০০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা সিবি শাইন পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- চমৎকার মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- হোন্ডা মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- প্রিমিয়াম মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- মসৃণ কিন্তু শক্তিশালী নয়</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- ক্লাসে সেরা</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- প্রিমিয়াম হোন্ডা বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- বেসিক সরঞ্জাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- রক্ষণশীল চেহারা</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা সিবি শাইন সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি শাইন এর মাইলেজ কত?</h3>
<p class="faq-answer">হোন্ডা সিবি শাইন বাস্তব পরিস্থিতিতে প্রতি লিটারে ৫৫-৬০ কিলোমিটার মাইলেজ দেয়, যা এটিকে সবচেয়ে জ্বালানি-দক্ষ ১২৫সিসি বাইকগুলির মধ্যে একটি করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে হোন্ডা সিবি শাইন এর দাম কত ২০২৪?</h3>
<p class="faq-answer">হোন্ডা সিবি শাইন এর দাম বাংলাদেশে ৳১,৬০,০০০ থেকে ৳১,৭০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি শাইনে কি ডিস্ক ব্রেক আছে?</h3>
<p class="faq-answer">না, হোন্ডা সিবি শাইন উভয় চাকায় ড্রাম ব্রেক সহ আসে। বাংলাদেশে কোনো ডিস্ক ব্রেক ভ্যারিয়েন্ট উপলব্ধ নেই।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন যাতায়াতের জন্য হোন্ডা সিবি শাইন কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, হোন্ডা সিবি শাইন দৈনন্দিন যাতায়াতের জন্য চমৎকার এর অত্যন্ত মসৃণ ইঞ্জিন, আরামদায়ক সিট এবং ৫৫-৬০ কিমি/লিটার মাইলেজ সহ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি শাইন এর সর্বোচ্চ গতি কত?</h3>
<p class="faq-answer">হোন্ডা সিবি শাইন প্রায় ১০০-১০৫ কিমি/ঘণ্টা সর্বোচ্চ গতিতে পৌঁছাতে পারে, শহর এবং মাঝে মাঝে হাইওয়ে রাইডিংয়ের জন্য উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি শাইন সার্ভিস খরচ কত?</h3>
<p class="faq-answer">অনুমোদিত হোন্ডা সার্ভিস সেন্টারে প্রতি ৩,০০০ কিমি বা ৩ মাসে নিয়মিত সার্ভিস খরচ ৳১,০০০ থেকে ৳১,৫০০ টাকা।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সিবি শাইন কি সিবি১২৫এফ এর চেয়ে ভালো?</h3>
<p class="faq-answer">সিবি শাইন ভালো আরাম এবং মসৃণ ইঞ্জিন প্রদান করে কিন্তু সিবি১২৫এফ আরো শক্তি প্রদান করে। আরামের জন্য শাইন নির্বাচন করুন, পারফরম্যান্সের জন্য সিবি১২৫এফ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি শাইনে কি ইলেকট্রিক স্টার্ট আছে?</h3>
<p class="faq-answer">হ্যাঁ, হোন্ডা সিবি শাইন অতিরিক্ত সুবিধার জন্য ব্যাকআপ হিসাবে ইলেকট্রিক স্টার্ট এবং কিক স্টার্ট উভয়ই সহ আসে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা সিবি শাইন কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.2/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হোন্ডা সিবি শাইন অন্য সব কিছুর উপরে আরাম, স্মুথনেস এবং নির্ভরযোগ্যতাকে অগ্রাধিকার দেয় এমন রাইডারদের জন্য চূড়ান্ত পছন্দ। যদিও এর ৳১,৬৫,০০০ টাকার প্রিমিয়াম মূল্য বেশি মনে হতে পারে, অত্যন্ত মসৃণ ইঞ্জিন, ক্লাসের সেরা আরাম এবং হোন্ডার কিংবদন্তী নির্ভরযোগ্যতা গুরুতর যাত্রীদের জন্য খরচকে ন্যায্য করে তোলে। আপনি যদি একটি পরিশোধিত, কম্পন-মুক্ত রাইডিং অভিজ্ঞতাকে মূল্য দেন তবে ডিস্ক ব্রেক এবং স্পোর্টি ফিচারের অভাব গুরুত্বপূর্ণ নয়। প্রিমিয়াম গুণমানের জন্য অর্থ প্রদান করতে পারেন এমন লম্বা দৈনিক যাত্রা করা পরিণত রাইডারদের জন্য আদর্শ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আরাম-কেন্দ্রিক দৈনিক যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CB Shine already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CB Shine\n")

	return nil
}
