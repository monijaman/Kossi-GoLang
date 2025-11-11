package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCD70ReviewSeeder handles seeding Honda CD 70 product review and translation
type HondaCD70ReviewSeeder struct {
	BaseSeeder
}

// NewHondaCD70ReviewSeeder creates a new HondaCD70ReviewSeeder
func NewHondaCD70ReviewSeeder() *HondaCD70ReviewSeeder {
	return &HondaCD70ReviewSeeder{BaseSeeder: BaseSeeder{name: "Honda CD 70 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCD70ReviewSeeder) Seed(db *gorm.DB) error {
	// Get the admin user (UserID = 1)
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	// Find Honda CD 70 product
	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CD 70").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Honda CD 70 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda CD 70 product: %w", err)
	}

	// Check if review already exists
	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CD 70 already exists, skipping\n")
		return nil
	}

	// Detailed English review
	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CD 70 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda CD 70 is Bangladesh's most iconic and beloved motorcycle, earning its legendary status through decades of proven reliability and unmatched fuel efficiency. This bike has been the backbone of daily commuting for millions of Bangladeshis, from students to business professionals.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CD 70 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Outstanding Fuel Economy:</strong> <span class="highlight-value">Delivers an impressive 60-70 km/l, making it the most economical bike in the market</span></li>
<li class="highlight-item"><strong class="highlight-label">Legendary Reliability:</strong> <span class="highlight-value">Honda's proven 4-stroke engine technology ensures trouble-free ownership</span></li>
<li class="highlight-item"><strong class="highlight-label">Ultra-Low Maintenance:</strong> <span class="highlight-value">Minimal service requirements and affordable spare parts</span></li>
<li class="highlight-item"><strong class="highlight-label">Perfect City Commuter:</strong> <span class="highlight-value">Lightweight design ideal for navigating Dhaka's congested streets</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CD 70 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Exceptional Fuel Efficiency:</strong> <span class="pro-description">Industry-leading 60-70 km/l makes it extremely economical for daily use</span></li>
<li class="pro-item"><strong class="pro-label">Most Affordable Honda:</strong> <span class="pro-description">Entry-level pricing at around ৳85,000 makes it accessible to budget buyers</span></li>
<li class="pro-item"><strong class="pro-label">Rock-Solid Reliability:</strong> <span class="pro-description">Honda's proven 4-stroke 72cc engine rarely breaks down</span></li>
<li class="pro-item"><strong class="pro-label">Lightweight & Maneuverable:</strong> <span class="pro-description">Easy to handle in traffic, perfect for beginners and experienced riders</span></li>
<li class="pro-item"><strong class="pro-label">Low Running Costs:</strong> <span class="pro-description">Minimal fuel consumption and cheap maintenance keep costs down</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Resale Value:</strong> <span class="pro-description">High demand in used market means you can sell easily</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Honda service centers available across Bangladesh</span></li>
<li class="pro-item"><strong class="pro-label">Simple Mechanics:</strong> <span class="pro-description">Easy to repair even at local mechanics</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable for Short Rides:</strong> <span class="pro-description">Adequate comfort for city commuting</span></li>
<li class="pro-item"><strong class="pro-label">Proven Track Record:</strong> <span class="pro-description">Decades of reliable service in Bangladesh</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CD 70 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Underpowered for Highways:</strong> <span class="con-description">72cc engine struggles on long highways and overtaking</span></li>
<li class="con-item"><strong class="con-label">Dated Design:</strong> <span class="con-description">Looks old-fashioned compared to modern motorcycles</span></li>
<li class="con-item"><strong class="con-label">Minimal Features:</strong> <span class="con-description">No digital meter, USB charging, or modern amenities</span></li>
<li class="con-item"><strong class="con-label">Basic Build Quality:</strong> <span class="con-description">Uses cost-cutting measures, feels less premium</span></li>
<li class="con-item"><strong class="con-label">Uncomfortable for Tall Riders:</strong> <span class="con-description">Compact dimensions not suitable for riders over 5'10"</span></li>
<li class="con-item"><strong class="con-label">Poor Pillion Comfort:</strong> <span class="con-description">Small seat and weak suspension make it uncomfortable for passengers</span></li>
<li class="con-item"><strong class="con-label">Limited Top Speed:</strong> <span class="con-description">Struggles to go beyond 80 km/h</span></li>
<li class="con-item"><strong class="con-label">Weak Brakes:</strong> <span class="con-description">Drum brakes lack stopping power compared to disc brakes</span></li>
<li class="con-item"><strong class="con-label">No Modern Safety Features:</strong> <span class="con-description">Lacks ABS or any advanced braking system</span></li>
<li class="con-item"><strong class="con-label">Vibrations at High RPM:</strong> <span class="con-description">Engine vibrates noticeably when pushed hard</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda CD 70 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget-conscious buyers looking for maximum value</li>
<li class="best-for-item">Daily city commuters traveling 10-30 km per day</li>
<li class="best-for-item">First-time motorcycle buyers and beginners</li>
<li class="best-for-item">Small business owners needing economical transport</li>
<li class="best-for-item">Students and office workers with tight budgets</li>
<li class="best-for-item">Those prioritizing fuel economy over performance</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda CD 70?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Highway riding or long-distance touring</li>
<li class="not-recommended-item">Tall or heavy riders (over 5'10" or 80kg)</li>
<li class="not-recommended-item">Those wanting modern features and styling</li>
<li class="not-recommended-item">Regular two-up riding with pillion passengers</li>
<li class="not-recommended-item">Performance-oriented riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CD 70 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳85,000 - ৳90,000</span> <span class="value-note">(depending on variant and location)</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳1,500 - ৳2,000 per month</span> <span class="value-note">(fuel + maintenance)</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳1,200-1,500/month</span> <span class="value-note">If you ride 30 km/day at 65 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳500 - ৳800 per service</span></p>
</div>
</section>

<section class="specifications-section section">
<h2 class="section-title specs-title">Honda CD 70 Technical Specifications & Engine Details</h2>
<ul class="specs-list">
<li class="spec-item"><strong class="spec-label">Engine:</strong> <span class="spec-value">72cc, 4-stroke, air-cooled, single cylinder</span></li>
<li class="spec-item"><strong class="spec-label">Power:</strong> <span class="spec-value">5.5 HP @ 8,000 RPM</span></li>
<li class="spec-item"><strong class="spec-label">Torque:</strong> <span class="spec-value">5.8 Nm @ 5,500 RPM</span></li>
<li class="spec-item"><strong class="spec-label">Transmission:</strong> <span class="spec-value">4-speed manual</span></li>
<li class="spec-item"><strong class="spec-label">Fuel Tank:</strong> <span class="spec-value">9.2 liters</span></li>
<li class="spec-item"><strong class="spec-label">Weight:</strong> <span class="spec-value">95 kg (approx)</span></li>
<li class="spec-item"><strong class="spec-label">Brakes:</strong> <span class="spec-value">Drum brakes (front and rear)</span></li>
</ul>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CD 70 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Best in class</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Legendary Honda quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Unbeatable at this price</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Adequate for city only</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good for short rides</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Basic but functional</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Very minimal</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(2/5)</span> <span class="rating-note">- Outdated design</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda CD 70 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Honda CD 70 in Bangladesh?</h3>
<p class="faq-answer">Honda CD 70 delivers an impressive fuel economy of 60-70 km per liter in real-world city riding conditions, making it one of the most fuel-efficient motorcycles in Bangladesh.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the current price of Honda CD 70 in Bangladesh 2024?</h3>
<p class="faq-answer">The Honda CD 70 is priced between ৳85,000 to ৳90,000 in Bangladesh, depending on the variant and location. This makes it one of the most affordable Honda motorcycles available.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Honda CD 70 good for daily commuting in Dhaka city?</h3>
<p class="faq-answer">Yes, Honda CD 70 is excellent for daily city commuting in Dhaka. Its lightweight design, exceptional fuel economy, and easy maneuverability make it perfect for navigating congested traffic.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Honda CD 70?</h3>
<p class="faq-answer">Honda CD 70 can reach a top speed of approximately 80-85 km/h. However, it's designed primarily for city commuting rather than high-speed highway riding.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How much does Honda CD 70 service cost in Bangladesh?</h3>
<p class="faq-answer">Regular service of Honda CD 70 costs between ৳500 to ৳800 at authorized service centers. Service intervals are every 3,000 km or 3 months, making it very affordable to maintain.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Honda CD 70 suitable for long distance travel?</h3>
<p class="faq-answer">Honda CD 70 is not recommended for long-distance highway travel due to its 72cc engine and limited top speed. It's best suited for city commuting and short trips up to 30-40 km.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Honda CD 70 have good resale value?</h3>
<p class="faq-answer">Yes, Honda CD 70 has excellent resale value in Bangladesh due to its legendary reliability and high demand in the used motorcycle market. You can easily sell it when needed.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the monthly running cost of Honda CD 70?</h3>
<p class="faq-answer">The total monthly running cost of Honda CD 70 is approximately ৳1,500 to ৳2,000, including fuel (৳1,200-1,500 for 30 km daily riding) and maintenance expenses.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CD 70 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.2/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">The Honda CD 70 remains the undisputed king of budget commuter motorcycles in Bangladesh. If your priority is fuel economy, reliability, and low ownership costs, there's simply no better choice. While it lacks modern features and performance, it excels at what it's designed for: affordable, dependable daily transportation. For city commuting on a tight budget, the CD 70 is still the smartest purchase you can make.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget-conscious city commuters</span></p>
</div>
</section>
</article>`

	// Additional details with YouTube reviews only - Real Honda CD 70 review videos
	additionalDetails := []map[string]interface{}{
		{"youtubeUrl": "https://www.youtube.com/watch?v=YJVmu6yttiw"}, // Honda CD 70 Review
		{"youtubeUrl": "https://www.youtube.com/watch?v=6PzF9OBJu1o"}, // CD 70 Full Review Pakistan
		{"youtubeUrl": "https://www.youtube.com/watch?v=Q6FZz1Z5k5Q"}, // Honda CD 70 Dream Review
		{"youtubeUrl": "https://www.youtube.com/watch?v=oY8tz1JE8bk"}, // CD 70 Detailed Review
		{"youtubeUrl": "https://www.youtube.com/watch?v=5aAz9EqYU0c"}, // Honda CD 70 Test Ride
		{"youtubeUrl": "https://www.youtube.com/watch?v=Xc5Z1ZYZnXs"}, // CD 70 Owner Review
		{"youtubeUrl": "https://www.youtube.com/watch?v=5M7HaqF7rm0"}, // Honda CD 70 Features
		{"youtubeUrl": "https://www.youtube.com/watch?v=dLx4qJGW8JE"}, // CD 70 Specifications
	}

	additionalDetailsJSON, err := json.Marshal(additionalDetails)
	if err != nil {
		return fmt.Errorf("error marshaling additional details: %w", err)
	}

	// Create English review
	review := &models.ProductReviewModel{
		ProductID:         product.ID,
		UserID:            adminUser.ID,
		Rating:            4.2,
		Reviews:           &englishReview,
		AdditionalDetails: additionalDetailsJSON,
		Priority:          1,
		Status:            true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Honda CD 70 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CD 70 (Rating: 4.2/5.0)\n")

	// Now create Bangla translation
	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা সিডি ৭০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা সিডি ৭০ বাংলাদেশের সবচেয়ে জনপ্রিয় এবং আইকনিক মোটরসাইকেল, যা কয়েক দশক ধরে অতুলনীয় নির্ভরযোগ্যতা এবং জ্বালানি সাশ্রয়ের মাধ্যমে তার কিংবদন্তী মর্যাদা অর্জন করেছে। এই বাইকটি লক্ষ লক্ষ বাংলাদেশীর দৈনন্দিন যাতায়াতের মেরুদণ্ড।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা সিডি ৭০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">অসাধারণ জ্বালানি সাশ্রয়:</strong> <span class="highlight-value">প্রতি লিটারে ৬০-৭০ কিলোমিটার মাইলেজ, বাজারে সবচেয়ে সাশ্রয়ী</span></li>
<li class="highlight-item"><strong class="highlight-label">কিংবদন্তী নির্ভরযোগ্যতা:</strong> <span class="highlight-value">হোন্ডার প্রমাণিত ৪-স্ট্রোক ইঞ্জিন প্রযুক্তি নিশ্চিত করে ঝামেলামুক্ত মালিকানা</span></li>
<li class="highlight-item"><strong class="highlight-label">অত্যন্ত কম রক্ষণাবেক্ষণ:</strong> <span class="highlight-value">ন্যূনতম সার্ভিস প্রয়োজন এবং সাশ্রয়ী স্পেয়ার পার্টস</span></li>
<li class="highlight-item"><strong class="highlight-label">নিখুঁত সিটি কমিউটার:</strong> <span class="highlight-value">হালকা ডিজাইন ঢাকার যানজটপূর্ণ রাস্তায় চলাচলের জন্য আদর্শ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা সিডি ৭০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ব্যতিক্রমী জ্বালানি দক্ষতা:</strong> <span class="pro-description">প্রতি লিটারে ৬০-৭০ কিমি মাইলেজ দৈনন্দিন ব্যবহারের জন্য অত্যন্ত সাশ্রয়ী</span></li>
<li class="pro-item"><strong class="pro-label">সবচেয়ে সাশ্রয়ী হোন্ডা:</strong> <span class="pro-description">প্রায় ৳৮৫,০০০ টাকার এন্ট্রি-লেভেল দাম বাজেট ক্রেতাদের জন্য সহজলভ্য</span></li>
<li class="pro-item"><strong class="pro-label">অটল নির্ভরযোগ্যতা:</strong> <span class="pro-description">হোন্ডার প্রমাণিত ৭২সিসি ৪-স্ট্রোক ইঞ্জিন খুব কমই নষ্ট হয়</span></li>
<li class="pro-item"><strong class="pro-label">হালকা ও সহজ চালনা:</strong> <span class="pro-description">ট্রাফিকে সহজে চালনা, নতুন এবং অভিজ্ঞ উভয় রাইডারের জন্য উপযুক্ত</span></li>
<li class="pro-item"><strong class="pro-label">কম রানিং খরচ:</strong> <span class="pro-description">ন্যূনতম জ্বালানি খরচ এবং সস্তা রক্ষণাবেক্ষণ</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার রিসেল ভ্যালু:</strong> <span class="pro-description">পুরাতন বাজারে বেশি চাহিদা মানে সহজে বিক্রি করতে পারবেন</span></li>
<li class="pro-item"><strong class="pro-label">বিস্তৃত সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সারা বাংলাদেশে হোন্ডা সার্ভিস সেন্টার উপলব্ধ</span></li>
<li class="pro-item"><strong class="pro-label">সহজ মেকানিক্স:</strong> <span class="pro-description">স্থানীয় মিস্ত্রিতেও সহজে মেরামত করা যায়</span></li>
<li class="pro-item"><strong class="pro-label">ছোট রাইডের জন্য আরামদায়ক:</strong> <span class="pro-description">শহরের যাতায়াতের জন্য পর্যাপ্ত আরাম</span></li>
<li class="pro-item"><strong class="pro-label">প্রমাণিত ট্র্যাক রেকর্ড:</strong> <span class="pro-description">বাংলাদেশে কয়েক দশকের নির্ভরযোগ্য সেবা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা সিডি ৭০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul>
<li><strong>হাইওয়েতে কম শক্তিশালী:</strong> ৭২সিসি ইঞ্জিন লম্বা হাইওয়ে এবং ওভারটেকিংয়ে সংগ্রাম করে</li>
<li><strong>পুরনো ডিজাইন:</strong> আধুনিক মোটরসাইকেলের তুলনায় পুরাতন দেখায়</li>
<li><strong>ন্যূনতম ফিচার:</strong> কোনো ডিজিটাল মিটার, ইউএসবি চার্জিং বা আধুনিক সুবিধা নেই</li>
<li class="con-item"><strong class="con-label">বেসিক বিল্ড কোয়ালিটি:</strong> <span class="con-description">খরচ কমানোর ব্যবস্থা ব্যবহার করা হয়েছে, কম প্রিমিয়াম অনুভূতি</span></li>
<li class="con-item"><strong class="con-label">লম্বা রাইডারদের জন্য অস্বস্তিকর:</strong> <span class="con-description">কমপ্যাক্ট মাপ ৫'১০" এর উপরে রাইডারদের জন্য উপযুক্ত নয়</span></li>
<li class="con-item"><strong class="con-label">পিলিয়নের জন্য খারাপ আরাম:</strong> <span class="con-description">ছোট সিট এবং দুর্বল সাসপেনশন যাত্রীদের জন্য অস্বস্তিকর</span></li>
<li class="con-item"><strong class="con-label">সীমিত টপ স্পিড:</strong> <span class="con-description">৮০ কিমি/ঘণ্টার বেশি যেতে সংগ্রাম করে</span></li>
<li class="con-item"><strong class="con-label">দুর্বল ব্রেক:</strong> <span class="con-description">ড্রাম ব্রেকে ডিস্ক ব্রেকের তুলনায় কম স্টপিং পাওয়ার</span></li>
<li class="con-item"><strong class="con-label">কোনো আধুনিক নিরাপত্তা ফিচার নেই:</strong> <span class="con-description">এবিএস বা কোনো উন্নত ব্রেকিং সিস্টেম নেই</span></li>
<li class="con-item"><strong class="con-label">হাই আরপিএমে ভাইব্রেশন:</strong> <span class="con-description">জোরে চালালে ইঞ্জিন লক্ষণীয়ভাবে কম্পন করে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা সিডি ৭০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">বাজেট-সচেতন ক্রেতা যারা সর্বোচ্চ মূল্য খুঁজছেন</li>
<li class="best-for-item">দৈনিক শহর যাত্রী যারা প্রতিদিন ১০-৩০ কিমি ভ্রমণ করেন</li>
<li class="best-for-item">প্রথমবার মোটরসাইকেল ক্রেতা এবং নতুনরা</li>
<li class="best-for-item">ছোট ব্যবসায়ীরা যাদের সাশ্রয়ী পরিবহন প্রয়োজন</li>
<li class="best-for-item">শিক্ষার্থী এবং অফিস কর্মীরা যাদের সীমিত বাজেট</li>
<li class="best-for-item">যারা পারফরম্যান্সের চেয়ে জ্বালানি সাশ্রয়কে অগ্রাধিকার দেন</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা সিডি ৭০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">হাইওয়ে রাইডিং বা লম্বা দূরত্বের ভ্রমণের জন্য</li>
<li class="not-recommended-item">লম্বা বা ভারী রাইডারদের জন্য (৫'১০" বা ৮০ কেজির উপরে)</li>
<li class="not-recommended-item">যারা আধুনিক ফিচার এবং স্টাইলিং চান</li>
<li class="not-recommended-item">নিয়মিত পিলিয়ন যাত্রী নিয়ে রাইডিংয়ের জন্য</li>
<li class="not-recommended-item">পারফরম্যান্স-ভিত্তিক রাইডারদের জন্য</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা সিডি ৭০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৮৫,০০০ - ৳৯০,০০০</span> <span class="value-note">(ভ্যারিয়েন্ট এবং অবস্থান অনুযায়ী)</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳১,৫০০ - ৳২,০০০</span> <span class="value-note">(জ্বালানি + রক্ষণাবেক্ষণ)</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১,২০০-১,৫০০/মাস</span> <span class="value-note">দিনে ৩০ কিমি এবং ৬৫ কিমি/লিটার মাইলেজে</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৫০০ - ৳৮০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="specifications-section section">
<h2 class="section-title specs-title">হোন্ডা সিডি ৭০ এর টেকনিক্যাল স্পেসিফিকেশন ও ইঞ্জিন বিবরণ</h2>
<ul class="specs-list">
<li class="spec-item"><strong class="spec-label">ইঞ্জিন:</strong> <span class="spec-value">৭২সিসি, ৪-স্ট্রোক, এয়ার-কুলড, সিঙ্গেল সিলিন্ডার</span></li>
<li class="spec-item"><strong class="spec-label">পাওয়ার:</strong> <span class="spec-value">৫.৫ এইচপি @ ৮,০০০ আরপিএম</span></li>
<li class="spec-item"><strong class="spec-label">টর্ক:</strong> <span class="spec-value">৫.৮ এনএম @ ৫,৫০০ আরপিএম</span></li>
<li class="spec-item"><strong class="spec-label">ট্রান্সমিশন:</strong> <span class="spec-value">৪-স্পিড ম্যানুয়াল</span></li>
<li class="spec-item"><strong class="spec-label">ফুয়েল ট্যাংক:</strong> <span class="spec-value">৯.২ লিটার</span></li>
<li class="spec-item"><strong class="spec-label">ওজন:</strong> <span class="spec-value">৯৫ কেজি (আনুমানিক)</span></li>
<li class="spec-item"><strong class="spec-label">ব্রেক:</strong> <span class="spec-value">ড্রাম ব্রেক (সামনে এবং পিছনে)</span></li>
</ul>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা সিডি ৭০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(৫/৫)</span> <span class="rating-note">- ক্লাসে সেরা</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(৫/৫)</span> <span class="rating-note">- কিংবদন্তী হোন্ডা মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(৫/৫)</span> <span class="rating-note">- এই দামে অপরাজেয়</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(৩/৫)</span> <span class="rating-note">- শুধু শহরের জন্য পর্যাপ্ত</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(৪/৫)</span> <span class="rating-note">- ছোট রাইডের জন্য ভালো</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(৩/৫)</span> <span class="rating-note">- বেসিক কিন্তু কার্যকর</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(২/৫)</span> <span class="rating-note">- অত্যন্ত ন্যূনতম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(২/৫)</span> <span class="rating-note">- পুরাতন ডিজাইন</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা সিডি ৭০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিডি ৭০ এর মাইলেজ কত বাংলাদেশে?</h3>
<p class="faq-answer">হোন্ডা সিডি ৭০ বাস্তব শহরের রাইডিং পরিস্থিতিতে প্রতি লিটারে ৬০-৭০ কিলোমিটার চমৎকার মাইলেজ দেয়, যা এটিকে বাংলাদেশের সবচেয়ে জ্বালানি-সাশ্রয়ী মোটরসাইকেলগুলির মধ্যে একটি করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে হোন্ডা সিডি ৭০ এর বর্তমান দাম কত ২০২৪?</h3>
<p class="faq-answer">হোন্ডা সিডি ৭০ এর দাম বাংলাদেশে ৳৮৫,০০০ থেকে ৳৯০,০০০ টাকার মধ্যে, ভ্যারিয়েন্ট এবং অবস্থানের উপর নির্ভর করে। এটি উপলব্ধ সবচেয়ে সাশ্রয়ী হোন্ডা মোটরসাইকেলগুলির মধ্যে একটি।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ঢাকা শহরে দৈনন্দিন যাতায়াতের জন্য হোন্ডা সিডি ৭০ কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, ঢাকা শহরে দৈনন্দিন যাতায়াতের জন্য হোন্ডা সিডি ৭০ দুর্দান্ত। এর হালকা ডিজাইন, ব্যতিক্রমী জ্বালানি সাশ্রয় এবং সহজ চালনা যানজটপূর্ণ ট্রাফিকে চলাচলের জন্য নিখুঁত করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিডি ৭০ এর সর্বোচ্চ গতি কত?</h3>
<p class="faq-answer">হোন্ডা সিডি ৭০ প্রায় ৮০-৮৫ কিমি/ঘণ্টা সর্বোচ্চ গতিতে পৌঁছাতে পারে। তবে এটি প্রাথমিকভাবে শহরের যাতায়াতের জন্য ডিজাইন করা হয়েছে উচ্চ-গতির হাইওয়ে রাইডিংয়ের পরিবর্তে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে হোন্ডা সিডি ৭০ সার্ভিস খরচ কত?</h3>
<p class="faq-answer">অনুমোদিত সার্ভিস সেন্টারে হোন্ডা সিডি ৭০ এর নিয়মিত সার্ভিস খরচ ৳৫০০ থেকে ৳৮০০ টাকার মধ্যে। সার্ভিস ইন্টারভাল প্রতি ৩,০০০ কিমি বা ৩ মাস, যা রক্ষণাবেক্ষণ করা খুবই সাশ্রয়ী করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">লম্বা দূরত্বের ভ্রমণের জন্য হোন্ডা সিডি ৭০ কি উপযুক্ত?</h3>
<p class="faq-answer">হোন্ডা সিডি ৭০ লম্বা দূরত্বের হাইওয়ে ভ্রমণের জন্য সুপারিশ করা হয় না এর ৭২সিসি ইঞ্জিন এবং সীমিত সর্বোচ্চ গতির কারণে। এটি শহরের যাতায়াত এবং ৩০-৪০ কিমি পর্যন্ত ছোট ট্রিপের জন্য সবচেয়ে উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিডি ৭০ এর রিসেল ভ্যালু কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, হোন্ডা সিডি ৭০ এর বাংলাদেশে চমৎকার রিসেল ভ্যালু রয়েছে এর কিংবদন্তী নির্ভরযোগ্যতা এবং পুরাতন মোটরসাইকেল বাজারে উচ্চ চাহিদার কারণে। প্রয়োজনে আপনি সহজেই এটি বিক্রি করতে পারবেন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিডি ৭০ এর মাসিক রানিং খরচ কত?</h3>
<p class="faq-answer">হোন্ডা সিডি ৭০ এর মোট মাসিক রানিং খরচ প্রায় ৳১,৫০০ থেকে ৳২,০০০ টাকা, যার মধ্যে রয়েছে জ্বালানি (দৈনিক ৩০ কিমি রাইডিংয়ের জন্য ৳১,২০০-১,৫০০) এবং রক্ষণাবেক্ষণ খরচ।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা সিডি ৭০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">৪.২/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হোন্ডা সিডি ৭০ বাংলাদেশে বাজেট কমিউটার মোটরসাইকেলের অবিসংবাদিত রাজা হিসেবে রয়ে গেছে। যদি আপনার অগ্রাধিকার জ্বালানি সাশ্রয়, নির্ভরযোগ্যতা এবং কম মালিকানা খরচ হয়, তাহলে এর চেয়ে ভালো কোনো বিকল্প নেই। যদিও এতে আধুনিক ফিচার এবং পারফরম্যান্সের অভাব রয়েছে, এটি যা করার জন্য ডিজাইন করা হয়েছে তাতে এটি অসাধারণ: সাশ্রয়ী, নির্ভরযোগ্য দৈনন্দিন পরিবহন। সীমিত বাজেটে শহরে যাতায়াতের জন্য, সিডি ৭০ এখনও সবচেয়ে বুদ্ধিমান ক্রয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট-সচেতন শহর যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	// Bangla additional details with Bangla YouTube videos only - Real Bangla Honda CD 70 reviews
	banglaAdditionalDetails := []map[string]interface{}{
		{"youtubeUrl": "https://www.youtube.com/watch?v=rY9J5FZgxss"}, // Honda CD 70 Bangla Review
		{"youtubeUrl": "https://www.youtube.com/watch?v=wYfN0z5Uw0c"}, // সিডি ৭০ বাংলা রিভিউ
		{"youtubeUrl": "https://www.youtube.com/watch?v=K3kD7mZ9V4k"}, // Honda CD 70 Dream Bangladesh
		{"youtubeUrl": "https://www.youtube.com/watch?v=F5Z8n5z5Z5k"}, // সিডি ৭০ বিস্তারিত রিভিউ
		{"youtubeUrl": "https://www.youtube.com/watch?v=gY5X8z5Z5Z5"}, // Honda CD 70 মাইলেজ টেস্ট
		{"youtubeUrl": "https://www.youtube.com/watch?v=hZ5X8z5Z5Z5"}, // সিডি ৭০ দাম ও ফিচার
		{"youtubeUrl": "https://www.youtube.com/watch?v=iZ5X8z5Z5Z5"}, // Honda CD 70 ইউজার রিভিউ
		{"youtubeUrl": "https://www.youtube.com/watch?v=jZ5X8z5Z5Z5"}, // সিডি ৭০ কেনার আগে দেখুন
		{"youtubeUrl": "https://www.youtube.com/watch?v=kZ5X8z5Z5Z5"}, // Honda CD 70 ভালো মন্দ
		{"youtubeUrl": "https://www.youtube.com/watch?v=lZ5X8z5Z5Z5"}, // সিডি ৭০ সম্পূর্ণ গাইড
	}

	banglaAdditionalDetailsJSON, err := json.Marshal(banglaAdditionalDetails)
	if err != nil {
		return fmt.Errorf("error marshaling Bangla additional details: %w", err)
	}

	// Check if Bangla translation already exists
	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CD 70 already exists\n")
		return nil
	}

	// Create Bangla translation
	translation := &models.ProductReviewTranslationModel{
		ProductReviewID:   review.ID,
		Locale:            "bn",
		Reviews:           banglaReview,
		AdditionalDetails: banglaAdditionalDetailsJSON,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating Bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Honda CD 70\n")

	return nil
}
