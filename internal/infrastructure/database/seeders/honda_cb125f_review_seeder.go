package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCB125FReviewSeeder handles seeding Honda CB125F product review and translation
type HondaCB125FReviewSeeder struct {
	BaseSeeder
}

// NewHondaCB125FReviewSeeder creates a new HondaCB125FReviewSeeder
func NewHondaCB125FReviewSeeder() *HondaCB125FReviewSeeder {
	return &HondaCB125FReviewSeeder{BaseSeeder: BaseSeeder{name: "Honda CB125F Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCB125FReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB125F").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda CB125F product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Honda CB125F product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CB125F already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CB125F Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Honda CB125F is a powerful 125cc commuter motorcycle offering excellent balance between performance and fuel efficiency. Priced at ৳1,55,000, it's perfect for riders who need more power than 110cc bikes while maintaining good mileage and Honda's renowned reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB125F Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Powerful 125cc Engine:</strong> <span class="highlight-value">11 HP engine provides excellent acceleration and highway capability</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Fuel Economy:</strong> <span class="highlight-value">45-50 km/l despite the larger engine</span></li>
<li class="highlight-item"><strong class="highlight-label">5-Speed Gearbox:</strong> <span class="highlight-value">Smooth transmission for better performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Build Quality:</strong> <span class="highlight-value">Solid construction with Honda standards</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB125F Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Performance:</strong> <span class="pro-description">125cc engine delivers 11 HP for confident riding</span></li>
<li class="pro-item"><strong class="pro-label">Highway Capable:</strong> <span class="pro-description">Can easily maintain 100+ km/h on highways</span></li>
<li class="pro-item"><strong class="pro-label">Decent Mileage:</strong> <span class="pro-description">45-50 km/l is good for a 125cc bike</span></li>
<li class="pro-item"><strong class="pro-label">5-Speed Gearbox:</strong> <span class="pro-description">Smooth shifting and better fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Upright seating position reduces fatigue</span></li>
<li class="pro-item"><strong class="pro-label">Electric Start:</strong> <span class="pro-description">Convenient electric starter with kick backup</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Premium materials and solid construction</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven engine technology with minimal issues</span></li>
<li class="pro-item"><strong class="pro-label">Wide Service Network:</strong> <span class="pro-description">Honda service centers across Bangladesh</span></li>
<li class="pro-item"><strong class="pro-label">Strong Resale Value:</strong> <span class="pro-description">Honda brand ensures good used market demand</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB125F Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">₳1,55,000 is expensive for a commuter bike</span></li>
<li class="con-item"><strong class="con-label">Lower Mileage:</strong> <span class="con-description">45-50 km/l is less than 110cc bikes</span></li>
<li class="con-item"><strong class="con-label">Drum Brakes Only:</strong> <span class="con-description">No disc brake option available</span></li>
<li class="con-item"><strong class="con-label">Basic Features:</strong> <span class="con-description">Lacks modern amenities like USB charging</span></li>
<li class="con-item"><strong class="con-label">Plain Styling:</strong> <span class="con-description">Conservative design may seem boring</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">Heavier than 110cc bikes, harder to maneuver</span></li>
<li class="con-item"><strong class="con-label">Higher Service Cost:</strong> <span class="con-description">125cc maintenance costs more than smaller bikes</span></li>
<li class="con-item"><strong class="con-label">Analog Meter:</strong> <span class="con-description">Traditional instrument cluster, not digital</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Honda CB125F in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Riders needing more power for highway commuting</li>
<li class="best-for-item">Long-distance daily commuters (50+ km)</li>
<li class="best-for-item">Those who frequently ride with pillion passengers</li>
<li class="best-for-item">Riders prioritizing performance over maximum mileage</li>
<li class="best-for-item">Highway travelers seeking reliable 125cc option</li>
<li class="best-for-item">Honda enthusiasts wanting proven 125cc technology</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Honda CB125F?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget buyers seeking maximum fuel economy</li>
<li class="not-recommended-item">City-only riders (110cc sufficient)</li>
<li class="not-recommended-item">Those wanting sporty styling and disc brakes</li>
<li class="not-recommended-item">Riders seeking modern digital features</li>
<li class="not-recommended-item">First-time riders (may find it heavy)</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB125F Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,50,000 - ৳1,60,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,500 - ৳3,000 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,000-2,500/month</span> <span class="value-note">For 30 km daily at 47 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳1,000 - ৳1,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB125F Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good for 125cc</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Honda quality</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Pricey but capable</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Strong for commuter</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good ergonomics</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Premium Honda build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(2.5/5)</span> <span class="rating-note">- Basic equipment</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Conservative look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Honda CB125F Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Honda CB125F?</h3>
<p class="faq-answer">Honda CB125F delivers 45-50 km/l in real-world conditions, which is excellent for a 125cc motorcycle with good performance.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Honda CB125F in Bangladesh 2024?</h3>
<p class="faq-answer">Honda CB125F is priced between ৳1,50,000 to ৳1,60,000 in Bangladesh.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Honda CB125F have disc brake?</h3>
<p class="faq-answer">No, Honda CB125F comes with drum brakes on both wheels. No disc brake variant is available.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Honda CB125F good for highway riding?</h3>
<p class="faq-answer">Yes, Honda CB125F is excellent for highway riding with its 125cc engine capable of maintaining 100+ km/h comfortably and 5-speed gearbox.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the top speed of Honda CB125F?</h3>
<p class="faq-answer">Honda CB125F can reach a top speed of approximately 105-110 km/h, making it suitable for highway travel.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How much does Honda CB125F service cost?</h3>
<p class="faq-answer">Regular service costs ৳1,000 to ৳1,500 at authorized Honda centers. Service interval is every 3,000 km or 3 months.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is CB125F better than Livo?</h3>
<p class="faq-answer">CB125F offers more power (125cc vs 110cc) and better highway performance but lower mileage. Livo is more fuel-efficient and stylish. Choose based on your riding needs.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Does Honda CB125F have electric start?</h3>
<p class="faq-answer">Yes, Honda CB125F comes with both electric start and kick start as backup for reliability.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB125F in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">3.9/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Honda CB125F is the ideal choice for riders needing a powerful, reliable 125cc commuter for highway and long-distance travel. While it costs more and returns lower mileage than 110cc bikes, the extra power, 5-speed gearbox, and highway capability make it worthwhile for serious commuters. The lack of disc brakes and modern features is disappointing, but Honda's legendary reliability and strong performance compensate. Best for riders doing 50+ km daily with frequent highway use.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For long-distance highway commuters</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    3.9,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Honda CB125F review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CB125F (Rating: 3.9/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা সিবি১২৫এফ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">হোন্ডা সিবি১২৫এফ একটি শক্তিশালী ১২৫সিসি কমিউটার মোটরসাইকেল যা পারফরম্যান্স এবং জ্বালানি দক্ষতার মধ্যে চমৎকার ভারসাম্য প্রদান করে। ৳১,৫৫,০০০ টাকায় মূল্যবান, এটি এমন রাইডারদের জন্য নিখুঁত যাদের ১১০সিসি বাইকের চেয়ে বেশি শক্তি প্রয়োজন এবং ভালো মাইলেজ এবং হোন্ডার বিখ্যাত নির্ভরযোগ্যতা বজায় রাখতে চান।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা সিবি১২৫এফ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">শক্তিশালী ১২৫সিসি ইঞ্জিন:</strong> <span class="highlight-value">১১ এইচপি ইঞ্জিন চমৎকার এক্সিলারেশন এবং হাইওয়ে ক্ষমতা প্রদান করে</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো জ্বালানি সাশ্রয়:</strong> <span class="highlight-value">বড় ইঞ্জিন সত্ত্বেও ৪৫-৫০ কিমি/লিটার</span></li>
<li class="highlight-item"><strong class="highlight-label">৫-স্পিড গিয়ারবক্স:</strong> <span class="highlight-value">ভালো পারফরম্যান্সের জন্য মসৃণ ট্রান্সমিশন</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম বিল্ড কোয়ালিটি:</strong> <span class="highlight-value">হোন্ডা মান সহ কঠিন নির্মাণ</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা সিবি১২৫এফ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী পারফরম্যান্স:</strong> <span class="pro-description">১২৫সিসি ইঞ্জিন আত্মবিশ্বাসী রাইডিংয়ের জন্য ১১ এইচপি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">হাইওয়ে সক্ষম:</strong> <span class="pro-description">হাইওয়েতে ১০০+ কিমি/ঘণ্টা সহজেই বজায় রাখতে পারে</span></li>
<li class="pro-item"><strong class="pro-label">ভালো মাইলেজ:</strong> <span class="pro-description">১২৫সিসি বাইকের জন্য ৪৫-৫০ কিমি/লিটার ভালো</span></li>
<li class="pro-item"><strong class="pro-label">৫-স্পিড গিয়ারবক্স:</strong> <span class="pro-description">মসৃণ শিফটিং এবং ভালো জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এরগোনমিক্স:</strong> <span class="pro-description">সোজা বসার অবস্থান ক্লান্তি হ্রাস করে</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রিক স্টার্ট:</strong> <span class="pro-description">কিক ব্যাকআপ সহ সুবিধাজনক ইলেকট্রিক স্টার্টার</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম উপকরণ এবং কঠিন নির্মাণ</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">ন্যূনতম সমস্যা সহ প্রমাণিত ইঞ্জিন প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">বিস্তৃত সার্ভিস নেটওয়ার্ক:</strong> <span class="pro-description">সারা বাংলাদেশে হোন্ডা সার্ভিস সেন্টার</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী রিসেল ভ্যালু:</strong> <span class="pro-description">হোন্ডা ব্র্যান্ড ভালো পুরাতন বাজার চাহিদা নিশ্চিত করে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা সিবি১২৫এফ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">কমিউটার বাইকের জন্য ৳১,৫৫,০০০ টাকা ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">কম মাইলেজ:</strong> <span class="con-description">৪৫-৫০ কিমি/লিটার ১১০সিসি বাইকের চেয়ে কম</span></li>
<li class="con-item"><strong class="con-label">শুধু ড্রাম ব্রেক:</strong> <span class="con-description">কোনো ডিস্ক ব্রেক বিকল্প উপলব্ধ নেই</span></li>
<li class="con-item"><strong class="con-label">বেসিক ফিচার:</strong> <span class="con-description">ইউএসবি চার্জিংয়ের মতো আধুনিক সুবিধার অভাব</span></li>
<li class="con-item"><strong class="con-label">সাধারণ স্টাইলিং:</strong> <span class="con-description">রক্ষণশীল ডিজাইন বিরক্তিকর মনে হতে পারে</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">১১০সিসি বাইকের চেয়ে ভারী, চালনা করা কঠিন</span></li>
<li class="con-item"><strong class="con-label">উচ্চ সার্ভিস খরচ:</strong> <span class="con-description">১২৫সিসি রক্ষণাবেক্ষণ ছোট বাইকের চেয়ে বেশি খরচ হয়</span></li>
<li class="con-item"><strong class="con-label">অ্যানালগ মিটার:</strong> <span class="con-description">ঐতিহ্যবাহী ইন্সট্রুমেন্ট ক্লাস্টার, ডিজিটাল নয়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে হোন্ডা সিবি১২৫এফ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">হাইওয়ে যাতায়াতের জন্য আরো শক্তি প্রয়োজন এমন রাইডাররা</li>
<li class="best-for-item">লম্বা দূরত্বের দৈনিক যাত্রীরা (৫০+ কিমি)</li>
<li class="best-for-item">যারা ঘন ঘন পিলিয়ন যাত্রী নিয়ে রাইড করেন</li>
<li class="best-for-item">সর্বোচ্চ মাইলেজের চেয়ে পারফরম্যান্সকে অগ্রাধিকার দেন এমন রাইডাররা</li>
<li class="best-for-item">নির্ভরযোগ্য ১২৫সিসি বিকল্প খুঁজছেন হাইওয়ে ভ্রমণকারীরা</li>
<li class="best-for-item">প্রমাণিত ১২৫সিসি প্রযুক্তি চান হোন্ডা উৎসাহীরা</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">হোন্ডা সিবি১২৫এফ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">সর্বোচ্চ জ্বালানি সাশ্রয় খুঁজছেন বাজেট ক্রেতারা</li>
<li class="not-recommended-item">শুধুমাত্র-শহরের রাইডাররা (১১০সিসি যথেষ্ট)</li>
<li class="not-recommended-item">যারা স্পোর্টি স্টাইলিং এবং ডিস্ক ব্রেক চান</li>
<li class="not-recommended-item">আধুনিক ডিজিটাল ফিচার খুঁজছেন রাইডাররা</li>
<li class="not-recommended-item">প্রথমবার রাইডাররা (ভারী মনে হতে পারে)</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">হোন্ডা সিবি১২৫এফ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৫০,০০০ - ৳১,৬০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৫০০ - ৳৩,০০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳২,০০০-২,৫০০/মাস</span> <span class="value-note">দিনে ৩০ কিমি এবং ৪৭ কিমি/লিটার মাইলেজে</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳১,০০০ - ৳১,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">হোন্ডা সিবি১২৫এফ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(৪/৫)</span> <span class="rating-note">- ১২৫সিসির জন্য ভালো</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(৫/৫)</span> <span class="rating-note">- হোন্ডা মান</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(৩.৫/৫)</span> <span class="rating-note">- দামী কিন্তু সক্ষম</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(৪.৫/৫)</span> <span class="rating-note">- কমিউটারের জন্য শক্তিশালী</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(৪/৫)</span> <span class="rating-note">- ভালো এরগোনমিক্স</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(৪.৫/৫)</span> <span class="rating-note">- প্রিমিয়াম হোন্ডা বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(২.৫/৫)</span> <span class="rating-note">- বেসিক সরঞ্জাম</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(৩/৫)</span> <span class="rating-note">- রক্ষণশীল চেহারা</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">হোন্ডা সিবি১২৫এফ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি১২৫এফ এর মাইলেজ কত?</h3>
<p class="faq-answer">হোন্ডা সিবি১২৫এফ বাস্তব পরিস্থিতিতে প্রতি লিটারে ৪৫-৫০ কিলোমিটার মাইলেজ দেয়, যা ভালো পারফরম্যান্স সহ একটি ১২৫সিসি মোটরসাইকেলের জন্য চমৎকার।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">বাংলাদেশে হোন্ডা সিবি১২৫এফ এর দাম কত ২০২৪?</h3>
<p class="faq-answer">হোন্ডা সিবি১২৫এফ এর দাম বাংলাদেশে ৳১,৫০,০০০ থেকে ৳১,৬০,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি১২৫এফ এ কি ডিস্ক ব্রেক আছে?</h3>
<p class="faq-answer">না, হোন্ডা সিবি১২৫এফ উভয় চাকায় ড্রাম ব্রেক সহ আসে। কোনো ডিস্ক ব্রেক ভ্যারিয়েন্ট উপলব্ধ নেই।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হাইওয়ে রাইডিংয়ের জন্য হোন্ডা সিবি১২৫এফ কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, হোন্ডা সিবি১২৫এফ হাইওয়ে রাইডিংয়ের জন্য চমৎকার এর ১২৫সিসি ইঞ্জিন সহ যা ১০০+ কিমি/ঘণ্টা আরামদায়কভাবে বজায় রাখতে সক্ষম এবং ৫-স্পিড গিয়ারবক্স।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি১২৫এফ এর সর্বোচ্চ গতি কত?</h3>
<p class="faq-answer">হোন্ডা সিবি১২৫এফ প্রায় ১০৫-১১০ কিমি/ঘণ্টা সর্বোচ্চ গতিতে পৌঁছাতে পারে, যা এটিকে হাইওয়ে ভ্রমণের জন্য উপযুক্ত করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি১২৫এফ সার্ভিস খরচ কত?</h3>
<p class="faq-answer">অনুমোদিত হোন্ডা সেন্টারে নিয়মিত সার্ভিস খরচ ৳১,০০০ থেকে ৳১,৫০০ টাকা। সার্ভিস ইন্টারভাল প্রতি ৩,০০০ কিমি বা ৩ মাস।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সিবি১২৫এফ কি লাইভোর চেয়ে ভালো?</h3>
<p class="faq-answer">সিবি১২৫এফ আরো শক্তি (১২৫সিসি বনাম ১১০সিসি) এবং ভালো হাইওয়ে পারফরম্যান্স প্রদান করে কিন্তু কম মাইলেজ। লাইভো আরো জ্বালানি-দক্ষ এবং স্টাইলিশ। আপনার রাইডিং প্রয়োজনের ভিত্তিতে নির্বাচন করুন।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">হোন্ডা সিবি১২৫এফ এ কি ইলেকট্রিক স্টার্ট আছে?</h3>
<p class="faq-answer">হ্যাঁ, হোন্ডা সিবি১২৫এফ নির্ভরযোগ্যতার জন্য ব্যাকআপ হিসাবে ইলেকট্রিক স্টার্ট এবং কিক স্টার্ট উভয়ই সহ আসে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে হোন্ডা সিবি১২৫এফ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">৩.৯/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">হোন্ডা সিবি১২৫এফ হাইওয়ে এবং লম্বা দূরত্বের ভ্রমণের জন্য একটি শক্তিশালী, নির্ভরযোগ্য ১২৫সিসি কমিউটার প্রয়োজন এমন রাইডারদের জন্য আদর্শ পছন্দ। যদিও এটি ১১০সিসি বাইকের চেয়ে বেশি খরচ করে এবং কম মাইলেজ দেয়, অতিরিক্ত শক্তি, ৫-স্পিড গিয়ারবক্স এবং হাইওয়ে ক্ষমতা গুরুতর যাত্রীদের জন্য এটি সার্থক করে তোলে। ডিস্ক ব্রেক এবং আধুনিক ফিচারের অভাব হতাশাজনক, কিন্তু হোন্ডার কিংবদন্তী নির্ভরযোগ্যতা এবং শক্তিশালী পারফরম্যান্স ক্ষতিপূরণ দেয়। ঘন ঘন হাইওয়ে ব্যবহার সহ দৈনিক ৫০+ কিমি করেন এমন রাইডারদের জন্য সেরা।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- লম্বা দূরত্বের হাইওয়ে যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CB125F already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CB125F\n")

	return nil
}
