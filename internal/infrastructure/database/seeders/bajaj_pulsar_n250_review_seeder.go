package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarN250Review handles seeding Bajaj Pulsar N250 product review and translation
type BajajPulsarN250Review struct {
	BaseSeeder
}

// NewBajajPulsarN250ReviewSeeder creates a new BajajPulsarN250Review
func NewBajajPulsarN250ReviewSeeder() *BajajPulsarN250Review {
	return &BajajPulsarN250Review{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar N250 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarN250Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar N250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Pulsar N250 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Pulsar N250 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar N250 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar N250 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Pulsar N250 is a naked 250cc street fighter featuring aggressive muscular styling, oil-cooled engine producing 24.5 PS, dual-channel ABS, LED lighting, and sporty ergonomics. Priced at ৳3,30,100, it offers exceptional value with strong performance, sharp handling, modern features, fuel efficiency, and Pulsar reliability making it ideal for riders seeking affordable 250cc naked sports performance with practical versatility.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar N250 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Muscular Naked Design:</strong> <span class="highlight-value">Aggressive street fighter styling with exposed engine</span></li>
<li class="highlight-item"><strong class="highlight-label">250cc Performance:</strong> <span class="highlight-value">24.5 PS power for strong acceleration</span></li>
<li class="highlight-item"><strong class="highlight-label">Dual-Channel ABS:</strong> <span class="highlight-value">Advanced braking safety system</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Full LED headlight and DRL</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar N250 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Performance:</strong> <span class="pro-description">24.5 PS provides thrilling acceleration and highway capability</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Lightweight naked design enhances agility</span></li>
<li class="pro-item"><strong class="pro-label">Dual-Channel ABS:</strong> <span class="pro-description">Superior braking safety on both wheels</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">LED lighting, digital display, USB charging</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Value:</strong> <span class="pro-description">250cc with ABS at competitive ৳3,30,100</span></li>
<li class="pro-item"><strong class="pro-label">Versatile Use:</strong> <span class="pro-description">Good for both city riding and touring</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar N250 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Vibrations:</strong> <span class="con-description">Oil-cooled engine shows vibrations at high RPM</span></li>
<li class="con-item"><strong class="con-label">No Wind Protection:</strong> <span class="con-description">Naked design lacks fairing for highway touring</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Fit and finish not as refined as premium brands</span></li>
<li class="con-item"><strong class="con-label">Firm Seat:</strong> <span class="con-description">Can be uncomfortable on long rides</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Bajaj service centers in some areas</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Bajaj Pulsar N250 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance enthusiasts on budget</li>
<li class="best-for-item">Urban street fighters</li>
<li class="best-for-item">First-time 250cc buyers</li>
<li class="best-for-item">Daily sporty commuting</li>
<li class="best-for-item">Weekend spirited riding</li>
<li class="best-for-item">Those seeking power and agility</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Bajaj Pulsar N250?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Long-distance touring without wind protection</li>
<li class="not-recommended-item">Premium build quality seekers</li>
<li class="not-recommended-item">Riders sensitive to vibrations</li>
<li class="not-recommended-item">Those requiring soft seat comfort</li>
<li class="not-recommended-item">Heavy city traffic only</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar N250 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,30,100 - Outstanding for naked 250cc with ABS</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-10 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳7-9 per km (35-40 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳7,000-11,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar N250 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Strong 250cc power</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Good fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Sporty but firm</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Aggressive street fighter</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Exceptional value</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- Decent Bajaj quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Bajaj Pulsar N250 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">Should I get N250 or F250?</h3>
<p class="faq-answer">N250 is better for city agility, sporty riding, and lower price (৳35,000 less). F250 is better for highway touring with wind protection. Choose based on primary use - city/sporty vs highway/touring.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How does it perform in city traffic?</h3>
<p class="faq-answer">The N250 handles city traffic well with lighter weight and better maneuverability than faired F250. Good low-end torque makes stop-and-go riding comfortable, though 250cc power can feel excessive in congested traffic.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is maintenance expensive?</h3>
<p class="faq-answer">Maintenance costs are moderate for a 250cc bike at ৳7,000-11,000 annually. Bajaj parts are reasonably priced and widely available, though not as cheap as 150cc commuters.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar N250 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,30,100, the Bajaj Pulsar N250 delivers exceptional value as a powerful naked 250cc street fighter with strong performance (24.5 PS), dual-channel ABS safety, sharp handling, modern features, and good fuel economy (35-40 km/l). It excels at sporty urban riding and spirited weekend rides with muscular styling and agility. However, engine vibrations, lack of wind protection, average build quality, firm seat, and limited service network make it best suited for performance enthusiasts, urban street riders, and sporty commuters who prioritize power, agility, and value over touring comfort and premium refinement.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For exceptional naked 250cc performance with ABS at outstanding value</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Bajaj Pulsar N250 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar N250 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার এন২৫০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ পালসার এন২৫০ হল একটি নেকেড ২৫০সিসি স্ট্রিট ফাইটার যেখানে আক্রমণাত্মক মাসকুলার স্টাইলিং, ২৪.৫ পিএস উৎপাদনকারী অয়েল-কুলড ইঞ্জিন, ডুয়াল-চ্যানেল এবিএস, এলইডি লাইটিং এবং স্পোর্টি এরগোনমিক্স রয়েছে। ৳৩,৩০,১০০ টাকায় মূল্যায়িত, এটি শক্তিশালী পারফরমেন্স, তীক্ষ্ণ হ্যান্ডলিং, আধুনিক বৈশিষ্ট্য, জ্বালানি দক্ষতা এবং পালসার নির্ভরযোগ্যতা সহ ব্যতিক্রমী মূল্য প্রদান করে যা ব্যবহারিক বহুমুখিতা সহ সাশ্রয়ী ২৫০সিসি নেকেড স্পোর্টস পারফরমেন্স খোঁজা রাইডারদের জন্য আদর্শ করে তোলে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার এন২৫০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">মাসকুলার নেকেড ডিজাইন:</strong> <span class="highlight-value">উন্মুক্ত ইঞ্জিন সহ আক্রমণাত্মক স্ট্রিট ফাইটার স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">২৫০সিসি পারফরমেন্স:</strong> <span class="highlight-value">শক্তিশালী ত্বরণের জন্য ২৪.৫ পিএস শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">ডুয়াল-চ্যানেল এবিএস:</strong> <span class="highlight-value">উন্নত ব্রেকিং নিরাপত্তা ব্যবস্থা</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">সম্পূর্ণ এলইডি হেডলাইট এবং ডিআরএল</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার এন২৫০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী পারফরমেন্স:</strong> <span class="pro-description">২৪.৫ পিএস রোমাঞ্চকর ত্বরণ এবং হাইওয়ে ক্ষমতা প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="pro-description">লাইটওয়েট নেকেড ডিজাইন চপলতা বাড়ায়</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল-চ্যানেল এবিএস:</strong> <span class="pro-description">উভয় চাকায় উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক বৈশিষ্ট্য:</strong> <span class="pro-description">এলইডি লাইটিং, ডিজিটাল ডিসপ্লে, ইউএসবি চার্জিং</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মূল্য:</strong> <span class="pro-description">প্রতিযোগিতামূলক ৳৩,৩০,১০০ এ এবিএস সহ ২৫০সিসি</span></li>
<li class="pro-item"><strong class="pro-label">বহুমুখী ব্যবহার:</strong> <span class="pro-description">শহুরে রাইডিং এবং ট্যুরিং উভয়ের জন্য ভাল</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার এন২৫০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">অয়েল-কুলড ইঞ্জিন উচ্চ আরপিএমে কম্পন দেখায়</span></li>
<li class="con-item"><strong class="con-label">বায়ু সুরক্ষা নেই:</strong> <span class="con-description">নেকেড ডিজাইনে হাইওয়ে ট্যুরিংয়ের জন্য ফেয়ারিং নেই</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">ফিট এবং ফিনিশ প্রিমিয়াম ব্র্যান্ডের মতো পরিমার্জিত নয়</span></li>
<li class="con-item"><strong class="con-label">শক্ত সিট:</strong> <span class="con-description">দীর্ঘ রাইডে অস্বস্তিকর হতে পারে</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">কিছু এলাকায় কম বাজাজ সেবা কেন্দ্র</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বাজাজ পালসার এন২৫০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Performance enthusiasts on budget</li>
<li class="best-for-item">Urban street fighters</li>
<li class="best-for-item">First-time 250cc buyers</li>
<li class="best-for-item">Daily sporty commuting</li>
<li class="best-for-item">Weekend spirited riding</li>
<li class="best-for-item">Those seeking power and agility</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বাজাজ পালসার এন২৫০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Long-distance touring without wind protection</li>
<li class="not-recommended-item">Premium build quality seekers</li>
<li class="not-recommended-item">Riders sensitive to vibrations</li>
<li class="not-recommended-item">Those requiring soft seat comfort</li>
<li class="not-recommended-item">Heavy city traffic only</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বাজাজ পালসার এন২৫০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৩,৩০,১০০ - এবিএস সহ নেকেড ২৫০সিসির জন্য অসাধারণ</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি - ৳৮-১০ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৭-৯ প্রতি কিমি (৩৫-৪০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৭,০০০-১১,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বাজাজ পালসার এন২৫০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- শক্তিশালী ২৫০সিসি শক্তি</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- ভাল জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- স্পোর্টি কিন্তু শক্ত</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- আক্রমণাত্মক স্ট্রিট ফাইটার</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- ব্যতিক্রমী মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.0</span> <span class="rating-note">- ভাল বাজাজ গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বাজাজ পালসার এন২৫০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">আমি কি এন২৫০ না এফ২৫০ নেব?</h3>
<p class="faq-answer">এন২৫০ শহুরে চপলতা, স্পোর্টি রাইডিং এবং কম দাম (৳৩৫,০০০ কম) জন্য ভাল। এফ২৫০ বায়ু সুরক্ষা সহ হাইওয়ে ট্যুরিংয়ের জন্য ভাল। প্রাথমিক ব্যবহারের উপর ভিত্তি করে বেছে নিন - শহর/স্পোর্টি বনাম হাইওয়ে/ট্যুরিং।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি শহুরে ট্রাফিকে কীভাবে পারফর্ম করে?</h3>
<p class="faq-answer">এন২৫০ ফেয়ারড এফ২৫০ এর চেয়ে হালকা ওজন এবং ভাল চালনার সাথে শহুরে ট্রাফিক ভালভাবে পরিচালনা করে। ভাল লো-এন্ড টর্ক স্টপ-অ্যান্ড-গো রাইডিং আরামদায়ক করে তোলে, যদিও জনাকীর্ণ ট্রাফিকে ২৫০সিসি শক্তি অতিরিক্ত মনে হতে পারে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রক্ষণাবেক্ষণ কি ব্যয়বহুল?</h3>
<p class="faq-answer">একটি ২৫০সিসি বাইকের জন্য রক্ষণাবেক্ষণ খরচ বার্ষিক ৳৭,০০০-১১,০০০ মাঝারি। বাজাজ যন্ত্রাংশ যুক্তিসঙ্গত মূল্যের এবং ব্যাপকভাবে উপলব্ধ, যদিও ১৫০সিসি কম্যুটারের মতো সস্তা নয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বাজাজ পালসার এন২৫০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৩,৩০,১০০ টাকায় বাজাজ পালসার এন২৫০ শক্তিশালী পারফরমেন্স (২৪.৫ পিএস), ডুয়াল-চ্যানেল এবিএস নিরাপত্তা, তীক্ষ্ণ হ্যান্ডলিং, আধুনিক বৈশিষ্ট্য এবং ভাল জ্বালানি অর্থনীতি (৩৫-৪০ কিমি/লিটার) সহ একটি শক্তিশালী নেকেড ২৫০সিসি স্ট্রিট ফাইটার হিসাবে ব্যতিক্রমী মূল্য প্রদান করে। এটি মাসকুলার স্টাইলিং এবং চপলতা সহ স্পোর্টি শহুরে রাইডিং এবং উৎসাহী সাপ্তাহিক রাইডে শ্রেষ্ঠ। তবে, ইঞ্জিন কম্পন, বায়ু সুরক্ষার অভাব, গড় বিল্ড কোয়ালিটি, শক্ত সিট এবং সীমিত সেবা নেটওয়ার্ক এটিকে পারফরমেন্স উৎসাহী, শহুরে স্ট্রিট রাইডার এবং ট্যুরিং আরাম এবং প্রিমিয়াম পরিমার্জনের চেয়ে শক্তি, চপলতা এবং মূল্যকে অগ্রাধিকার দেওয়া স্পোর্টি যাত্রীদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অসাধারণ মূল্যে এবিএস সহ ব্যতিক্রমী নেকেড ২৫০সিসি পারফরমেন্সের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar N250 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar N250\n")

	return nil
}
