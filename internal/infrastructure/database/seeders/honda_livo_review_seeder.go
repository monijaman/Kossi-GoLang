package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaLivoReviewSeeder handles seeding Honda Livo product review and translation
type HondaLivoReviewSeeder struct {
	BaseSeeder
}

// NewHondaLivoReviewSeeder creates a new HondaLivoReviewSeeder
func NewHondaLivoReviewSeeder() *HondaLivoReviewSeeder {
	return &HondaLivoReviewSeeder{BaseSeeder: BaseSeeder{name: "Honda Livo Review"}}
}

// Seed implements the Seeder interface
func (s *HondaLivoReviewSeeder) Seed(db *gorm.DB) error {
	// Get the admin user (UserID = 1)
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	// Find Honda Livo product
	var product models.ProductModel
	if err := db.Where("name = ?", "Honda Livo").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda Livo product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda Livo product: %w", err)
	}

	// Check if review already exists
	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda Livo already exists, skipping\n")
		return nil
	}

	// Detailed English review
	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda Livo Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda Livo is a stylish and fuel-efficient commuter motorcycle that perfectly balances modern design with Honda's legendary reliability. Priced at ৳1,50,000, it targets riders seeking a premium 110cc bike with excellent mileage and contemporary features.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda Livo Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Excellent Fuel Economy:</strong> <span class="highlight-value">Delivers 55-60 km/l making it economical for daily use</span></li>
<li class="highlight-item"><strong class="highlight-label">Stylish Modern Design:</strong> <span class="highlight-value">Contemporary styling with LED position lamps and sporty graphics</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Riding:</strong> <span class="highlight-value">Well-cushioned seat and upright ergonomics for city comfort</span></li>
<li class="highlight-item"><strong class="highlight-label">Reliable Honda Engine:</strong> <span class="highlight-value">Proven 110cc engine with smooth power delivery</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda Livo Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Great Fuel Efficiency:</strong> <span class="pro-description">55-60 km/l mileage keeps running costs low</span></li>
<li class="pro-item"><strong class="pro-label">Modern Styling:</strong> <span class="pro-description">Attractive design with LED elements and premium finish</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">Refined 110cc engine with minimal vibrations</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seat:</strong> <span class="pro-description">Well-padded seat suitable for long rides</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Honda's quality standards ensure durability</span></li>
<li class="pro-item"><strong class="pro-label">Digital-Analog Meter:</strong> <span class="pro-description">Modern instrument cluster with trip meter</span></li>
<li class="pro-item"><strong class="pro-label">LED Position Lamps:</strong> <span class="pro-description">Better visibility and contemporary look</span></li>
<li class="pro-item"><strong class="pro-label">Easy Maintenance:</strong> <span class="pro-description">Wide service network and affordable parts</span></li>
<li class="pro-item"><strong class="pro-label">Good Resale Value:</strong> <span class="pro-description">Honda brand ensures strong used market demand</span></li>
<li class="pro-item"><strong class="pro-label">Balanced Performance:</strong> <span class="pro-description">Adequate power for city and short highway rides</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda Livo Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">₳1,50,000 is expensive compared to basic commuters</span></li>
<li class="con-item"><strong class="con-label">No Disc Brake Variant:</strong> <span class="con-description">Only drum brakes available, lacks disc brake option</span></li>
<li class="con-item"><strong class="con-label">Limited Top Speed:</strong> <span class="con-description">110cc engine limits high-speed performance</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">No USB charging or fully digital console</span></li>
<li class="con-item"><strong class="con-label">Average Suspension:</strong> <span class="con-description">Standard suspension struggles on rough roads</span></li>
<li class="con-item"><strong class="con-label">Not for Heavy Riders:</strong> <span class="con-description">110cc engine may feel underpowered for riders over 85kg</span></li>
<li class="con-item"><strong class="con-label">Plastic Quality:</strong> <span class="con-description">Some body panels feel lightweight</span></li>
<li class="con-item"><strong class="con-label">No Engine Kill Switch:</strong> <span class="con-description">Missing this convenient safety feature</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda Livo in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters seeking fuel efficiency with style</li>
<li class="best-for-item">Young professionals wanting a modern-looking bike</li>
<li class="best-for-item">First-time buyers preferring Honda reliability</li>
<li class="best-for-item">Riders traveling 30-50 km daily</li>
<li class="best-for-item">Those prioritizing comfort and mileage</li>
<li class="best-for-item">Budget of ৳1.5 lakh with focus on brand value</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda Livo?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Riders seeking disc brakes and modern features</li>
<li class="not-recommended-item">Those on tight budget (CD 70 is better)</li>
<li class="not-recommended-item">Heavy riders needing more power (125cc+ recommended)</li>
<li class="not-recommended-item">Highway enthusiasts wanting speed</li>
<li class="not-recommended-item">Tech-savvy buyers wanting USB charging and digital features</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda Livo Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,45,000 - ৳1,55,000</span> <span class="value-note">(depending on location)</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,000 - ৳2,500 per month</span> <span class="value-note">(fuel + maintenance)</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,600-2,000/month</span> <span class="value-note">For 30 km daily at 57 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳800 - ৳1,200 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda Livo Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Excellent mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Honda quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good but pricey</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Adequate for commuting</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Very comfortable seat</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good Honda standards</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic but functional</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Modern and attractive</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda Livo Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Honda Livo in Bangladesh?</h3>
<p class="faq-answer">Honda Livo delivers 55-60 km per liter in real-world city riding conditions, making it one of the most fuel-efficient 110cc motorcycles in Bangladesh.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Honda Livo in Bangladesh 2024?</h3>
<p class="faq-answer">Honda Livo is priced between ৳1,45,000 to ৳1,55,000 in Bangladesh, depending on the location and dealer.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Honda Livo have disc brake?</h3>
<p class="faq-answer">No, Honda Livo comes with drum brakes on both front and rear. There is no disc brake variant available in Bangladesh currently.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Honda Livo good for daily commuting?</h3>
<p class="faq-answer">Yes, Honda Livo is excellent for daily commuting with its comfortable seat, smooth engine, good mileage of 55-60 km/l, and reliable Honda build quality.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Honda Livo?</h3>
<p class="faq-answer">Honda Livo can reach a top speed of approximately 95-100 km/h. It's designed primarily for efficient city commuting rather than high-speed performance.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How much does Honda Livo service cost?</h3>
<p class="faq-answer">Regular service of Honda Livo costs between ৳800 to ৳1,200 at authorized Honda service centers. Service intervals are every 3,000 km or 3 months.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Honda Livo better than CD 70?</h3>
<p class="faq-answer">Honda Livo offers more power (110cc vs 72cc), better styling, LED lamps, and superior comfort. However, CD 70 has better mileage and is much cheaper. Choose based on your budget and needs.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Honda Livo have LED headlight?</h3>
<p class="faq-answer">Honda Livo has LED position lamps for daytime visibility but uses a regular halogen bulb for the main headlight, not a full LED headlight system.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda Livo in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Honda Livo is an excellent choice for riders seeking a stylish, fuel-efficient commuter motorcycle with Honda's legendary reliability. While it's priced higher than basic commuters, the modern design, comfortable riding position, and 55-60 km/l mileage justify the premium. The lack of disc brakes is a notable omission, but the overall package delivers strong value for daily commuters who prioritize comfort and brand reputation.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For style-conscious daily commuters</span></p>
</div>
</section>
</article>`

	// Create English review with no additional details
	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.1,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Honda Livo review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda Livo (Rating: 4.1/5.0)\n")

	// Now create Bangla translation
	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা লাইভো রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা লাইভো একটি স্টাইলিশ এবং জ্বালানি-সাশ্রয়ী কমিউটার মোটরসাইকেল যা আধুনিক ডিজাইনের সাথে হোন্ডার কিংবদন্তী নির্ভরযোগ্যতার নিখুঁত সমন্বয়। ৳১,৫০,০০০ টাকায় মূল্যবান, এটি চমৎকার মাইলেজ এবং আধুনিক বৈশিষ্ট্য সহ একটি প্রিমিয়াম ১১০সিসি বাইক খুঁজছেন এমন রাইডারদের লক্ষ্য করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা লাইভো এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">চমৎকার জ্বালানি সাশ্রয়:</strong> <span class="highlight-value">প্রতি লিটারে ৫৫-৬০ কিমি মাইলেজ দৈনন্দিন ব্যবহারের জন্য সাশ্রয়ী</span></li>
<li class="highlight-item"><strong class="highlight-label">স্টাইলিশ আধুনিক ডিজাইন:</strong> <span class="highlight-value">এলইডি পজিশন ল্যাম্প এবং স্পোর্টি গ্রাফিক্স সহ সমসাময়িক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক রাইডিং:</strong> <span class="highlight-value">শহরের আরামের জন্য ভাল-কুশনযুক্ত সিট এবং সোজা এরগোনমিক্স</span></li>
<li class="highlight-item"><strong class="highlight-label">নির্ভরযোগ্য হোন্ডা ইঞ্জিন:</strong> <span class="highlight-value">মসৃণ পাওয়ার ডেলিভারি সহ প্রমাণিত ১১০সিসি ইঞ্জিন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা লাইভো এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">দুর্দান্ত জ্বালানি দক্ষতা:</strong> <span class="pro-description">৫৫-৬০ কিমি/লিটার মাইলেজ রানিং খরচ কম রাখে</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক স্টাইলিং:</strong> <span class="pro-description">এলইডি এলিমেন্ট এবং প্রিমিয়াম ফিনিশ সহ আকর্ষণীয় ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">ন্যূনতম কম্পন সহ পরিশোধিত ১১০সিসি ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক সিট:</strong> <span class="pro-description">লম্বা রাইডের জন্য উপযুক্ত ভাল-প্যাডেড সিট</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">হোন্ডার মানের মান স্থায়িত্ব নিশ্চিত করে</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল-অ্যানালগ মিটার:</strong> <span class="pro-description">ট্রিপ মিটার সহ আধুনিক ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
<li class="pro-item"><strong class="pro-label">এলইডি পজিশন ল্যাম্প:</strong> <span class="pro-description">ভালো দৃশ্যমানতা এবং সমসাময়িক চেহারা</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="pro-description">বিস্তৃত সার্ভিস নেটওয়ার্ক এবং সাশ্রয়ী পার্টস</span></li>
<li class="pro-item"><strong class="pro-label">ভালো রিসেল ভ্যালু:</strong> <span class="pro-description">হোন্ডা ব্র্যান্ড শক্তিশালী পুরাতন বাজার চাহিদা নিশ্চিত করে</span></li>
<li class="pro-item"><strong class="pro-label">সুষম পারফরম্যান্স:</strong> <span class="pro-description">শহর এবং ছোট হাইওয়ে রাইডের জন্য পর্যাপ্ত শক্তি</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা লাইভো এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">৳১,৫০,০০০ টাকা সাধারণ কমিউটারের তুলনায় ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">কোনো ডিস্ক ব্রেক ভ্যারিয়েন্ট নেই:</strong> <span class="con-description">শুধুমাত্র ড্রাম ব্রেক উপলব্ধ, ডিস্ক ব্রেক বিকল্প নেই</span></li>
<li class="con-item"><strong class="con-label">সীমিত সর্বোচ্চ গতি:</strong> <span class="con-description">১১০সিসি ইঞ্জিন উচ্চ-গতির পারফরম্যান্স সীমিত করে</span></li>
<li class="con-item"><strong class="con-label">বেসিক ফিচার:</strong> <span class="con-description">কোনো ইউএসবি চার্জিং বা সম্পূর্ণ ডিজিটাল কনসোল নেই</span></li>
<li class="con-item"><strong class="con-label">গড় সাসপেনশন:</strong> <span class="con-description">স্ট্যান্ডার্ড সাসপেনশন রুক্ষ রাস্তায় সংগ্রাম করে</span></li>
<li class="con-item"><strong class="con-label">ভারী রাইডারদের জন্য নয়:</strong> <span class="con-description">৮৫ কেজির উপরে রাইডারদের জন্য ১১০সিসি ইঞ্জিন দুর্বল অনুভব করতে পারে</span></li>
<li class="con-item"><strong class="con-label">প্লাস্টিকের গুণমান:</strong> <span class="con-description">কিছু বডি প্যানেল হালকা মনে হয়</span></li>
<li class="con-item"><strong class="con-label">কোনো ইঞ্জিন কিল সুইচ নেই:</strong> <span class="con-description">এই সুবিধাজনক নিরাপত্তা বৈশিষ্ট্য অনুপস্থিত</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা লাইভো কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">স্টাইলের সাথে জ্বালানি দক্ষতা খুঁজছেন এমন দৈনিক যাত্রীরা</li>
<li class="best-for-item">আধুনিক চেহারার বাইক চান এমন তরুণ পেশাদাররা</li>
<li class="best-for-item">হোন্ডা নির্ভরযোগ্যতা পছন্দ করেন এমন প্রথমবার ক্রেতারা</li>
<li class="best-for-item">দৈনিক ৩০-৫০ কিমি ভ্রমণকারী রাইডাররা</li>
<li class="best-for-item">যারা আরাম এবং মাইলেজকে অগ্রাধিকার দেন</li>
<li class="best-for-item">ব্র্যান্ড মূল্যের উপর ফোকাস সহ ৳১.৫ লাখ বাজেট</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা লাইভো কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">ডিস্ক ব্রেক এবং আধুনিক ফিচার খুঁজছেন এমন রাইডাররা</li>
<li class="not-recommended-item">সীমিত বাজেটে যারা আছেন (সিডি ৭০ ভালো)</li>
<li class="not-recommended-item">আরো শক্তি প্রয়োজন এমন ভারী রাইডাররা (১২৫সিসি+ সুপারিশকৃত)</li>
<li class="not-recommended-item">গতি চান এমন হাইওয়ে উৎসাহীরা</li>
<li class="not-recommended-item">ইউএসবি চার্জিং এবং ডিজিটাল ফিচার চান প্রযুক্তি-বুদ্ধিমান ক্রেতারা</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা লাইভো এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৪৫,০০০ - ৳১,৫৫,০০০</span> <span class="value-note">(অবস্থান অনুযায়ী)</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,০০০ - ৳২,৫০০</span> <span class="value-note">(জ্বালানি + রক্ষণাবেক্ষণ)</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১,৬০০-২,০০০/মাস</span> <span class="value-note">দিনে ৩০ কিমি এবং ৫৭ কিমি/লিটার মাইলেজে</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৮০০ - ৳১,২০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা লাইভো পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(৪.৫/৫)</span> <span class="rating-note">- চমৎকার মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(৫/৫)</span> <span class="rating-note">- হোন্ডা মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(৪/৫)</span> <span class="rating-note">- ভালো কিন্তু দামী</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(৩.৫/৫)</span> <span class="rating-note">- যাতায়াতের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(৪.৫/৫)</span> <span class="rating-note">- খুবই আরামদায়ক সিট</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(৪/৫)</span> <span class="rating-note">- ভালো হোন্ডা মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(৩/৫)</span> <span class="rating-note">- বেসিক কিন্তু কার্যকর</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(৪.৫/৫)</span> <span class="rating-note">- আধুনিক এবং আকর্ষণীয়</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা লাইভো সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হোন্ডা লাইভো এর মাইলেজ কত বাংলাদেশে?</h3>
<p class="faq-answer">হোন্ডা লাইভো বাস্তব শহরের রাইডিং পরিস্থিতিতে প্রতি লিটারে ৫৫-৬০ কিলোমিটার মাইলেজ দেয়, যা এটিকে বাংলাদেশের সবচেয়ে জ্বালানি-সাশ্রয়ী ১১০সিসি মোটরসাইকেলগুলির মধ্যে একটি করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে হোন্ডা লাইভো এর দাম কত ২০২৪?</h3>
<p class="faq-answer">হোন্ডা লাইভো এর দাম বাংলাদেশে ৳১,৪৫,০০০ থেকে ৳১,৫৫,০০০ টাকার মধ্যে, অবস্থান এবং ডিলারের উপর নির্ভর করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা লাইভোতে কি ডিস্ক ব্রেক আছে?</h3>
<p class="faq-answer">না, হোন্ডা লাইভো সামনে এবং পিছনে উভয়ই ড্রাম ব্রেক সহ আসে। বর্তমানে বাংলাদেশে কোনো ডিস্ক ব্রেক ভ্যারিয়েন্ট উপলব্ধ নেই।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">দৈনন্দিন যাতায়াতের জন্য হোন্ডা লাইভো কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, হোন্ডা লাইভো দৈনন্দিন যাতায়াতের জন্য চমৎকার এর আরামদায়ক সিট, মসৃণ ইঞ্জিন, ৫৫-৬০ কিমি/লিটার ভালো মাইলেজ এবং নির্ভরযোগ্য হোন্ডা বিল্ড কোয়ালিটি সহ।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা লাইভো এর সর্বোচ্চ গতি কত?</h3>
<p class="faq-answer">হোন্ডা লাইভো প্রায় ৯৫-১০০ কিমি/ঘণ্টা সর্বোচ্চ গতিতে পৌঁছাতে পারে। এটি প্রাথমিকভাবে দক্ষ শহরের যাতায়াতের জন্য ডিজাইন করা হয়েছে উচ্চ-গতির পারফরম্যান্সের পরিবর্তে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা লাইভো সার্ভিস খরচ কত?</h3>
<p class="faq-answer">অনুমোদিত হোন্ডা সার্ভিস সেন্টারে হোন্ডা লাইভো এর নিয়মিত সার্ভিস খরচ ৳৮০০ থেকে ৳১,২০০ টাকার মধ্যে। সার্ভিস ইন্টারভাল প্রতি ৩,০০০ কিমি বা ৩ মাস।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা লাইভো কি সিডি ৭০ এর চেয়ে ভালো?</h3>
<p class="faq-answer">হোন্ডা লাইভো আরো শক্তি (১১০সিসি বনাম ৭২সিসি), ভালো স্টাইলিং, এলইডি ল্যাম্প এবং উন্নত আরাম প্রদান করে। তবে, সিডি ৭০ এর ভালো মাইলেজ আছে এবং অনেক সস্তা। আপনার বাজেট এবং প্রয়োজনের উপর ভিত্তি করে নির্বাচন করুন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা লাইভোতে কি এলইডি হেডলাইট আছে?</h3>
<p class="faq-answer">হোন্ডা লাইভোতে দিনের দৃশ্যমানতার জন্য এলইডি পজিশন ল্যাম্প রয়েছে কিন্তু মূল হেডলাইটের জন্য একটি নিয়মিত হ্যালোজেন বাল্ব ব্যবহার করে, সম্পূর্ণ এলইডি হেডলাইট সিস্টেম নয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা লাইভো কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">৪.১/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হোন্ডা লাইভো হোন্ডার কিংবদন্তী নির্ভরযোগ্যতা সহ একটি স্টাইলিশ, জ্বালানি-সাশ্রয়ী কমিউটার মোটরসাইকেল খুঁজছেন এমন রাইডারদের জন্য একটি চমৎকার পছন্দ। যদিও এটি সাধারণ কমিউটারদের চেয়ে বেশি দামে মূল্যবান, আধুনিক ডিজাইন, আরামদায়ক রাইডিং পজিশন এবং ৫৫-৬০ কিমি/লিটার মাইলেজ প্রিমিয়ামকে ন্যায্যতা দেয়। ডিস্ক ব্রেকের অভাব একটি উল্লেখযোগ্য বাদ, কিন্তু সামগ্রিক প্যাকেজ দৈনিক যাত্রীদের জন্য শক্তিশালী মূল্য প্রদান করে যারা আরাম এবং ব্র্যান্ড খ্যাতিকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্টাইল-সচেতন দৈনিক যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	// Check if Bangla translation already exists
	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda Livo already exists\n")
		return nil
	}

	// Create Bangla translation with no additional details
	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating Bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Honda Livo\n")

	return nil
}
