package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarF250Review handles seeding Bajaj Pulsar F250 product review and translation
type BajajPulsarF250Review struct {
	BaseSeeder
}

// NewBajajPulsarF250ReviewSeeder creates a new BajajPulsarF250Review
func NewBajajPulsarF250ReviewSeeder() *BajajPulsarF250Review {
	return &BajajPulsarF250Review{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar F250 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarF250Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar F250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Pulsar F250 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Pulsar F250 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar F250 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar F250 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Pulsar F250 is a fully-faired 250cc sports tourer featuring aggressive aerodynamic fairing, oil-cooled engine, dual-channel ABS, LED lighting, and touring-focused ergonomics. Priced at ৳3,65,100, it offers excellent value in the 250cc segment with sports bike aesthetics, wind protection, good performance, fuel efficiency, and practical features making it ideal for riders seeking affordable faired sports touring capability with Pulsar's proven reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar F250 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Full Fairing Design:</strong> <span class="highlight-value">Aerodynamic fairing for wind protection and sports aesthetics</span></li>
<li class="highlight-item"><strong class="highlight-label">250cc Oil-Cooled Engine:</strong> <span class="highlight-value">24.5 PS power with good thermal management</span></li>
<li class="highlight-item"><strong class="highlight-label">Dual-Channel ABS:</strong> <span class="highlight-value">Advanced safety with ABS on both wheels</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Full LED headlight and taillight</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar F250 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Value Proposition:</strong> <span class="pro-description">Full fairing, ABS, and 250cc at competitive ৳3,65,100 price</span></li>
<li class="pro-item"><strong class="pro-label">Dual-Channel ABS Safety:</strong> <span class="pro-description">Advanced braking safety on both wheels</span></li>
<li class="pro-item"><strong class="pro-label">Good Performance:</strong> <span class="pro-description">24.5 PS provides spirited performance for touring</span></li>
<li class="pro-item"><strong class="pro-label">Wind Protection:</strong> <span class="pro-description">Full fairing reduces wind fatigue on highways</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Efficient:</strong> <span class="pro-description">Achieves 35-40 km/l for a 250cc bike</span></li>
<li class="pro-item"><strong class="pro-label">Proven Pulsar Platform:</strong> <span class="pro-description">Reliable Bajaj engineering and wide service network</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar F250 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Oil-Cooled Engine:</strong> <span class="con-description">Not as advanced as liquid-cooled competitors</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">162 kg makes city maneuverability challenging</span></li>
<li class="con-item"><strong class="con-label">Aggressive Riding Position:</strong> <span class="con-description">Forward-leaning ergonomics can be tiring</span></li>
<li class="con-item"><strong class="con-label">Average Build Quality:</strong> <span class="con-description">Fit and finish not as refined as premium brands</span></li>
<li class="con-item"><strong class="con-label">Vibrations at High RPM:</strong> <span class="con-description">Oil-cooled engine shows vibrations at higher speeds</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Bajaj Pulsar F250 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget-conscious sports touring enthusiasts</li>
<li class="best-for-item">Long-distance highway commuters</li>
<li class="best-for-item">First-time 250cc faired bike buyers</li>
<li class="best-for-item">Riders seeking wind protection</li>
<li class="best-for-item">Weekend touring riders</li>
<li class="best-for-item">Those wanting ABS safety at affordable price</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Bajaj Pulsar F250?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">City traffic navigation</li>
<li class="not-recommended-item">Performance purists seeking high refinement</li>
<li class="not-recommended-item">Shorter riders uncomfortable with sports ergonomics</li>
<li class="not-recommended-item">Premium build quality seekers</li>
<li class="not-recommended-item">Heavy daily city commuting</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar F250 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,65,100 - Excellent for faired 250cc with ABS</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-11 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳7-9 per km (35-40 km/l)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 5,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳8,000-12,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar F250 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good 250cc performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Decent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Good for touring</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Aggressive faired design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Outstanding value for money</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Adequate Bajaj quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Bajaj Pulsar F250 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">How does it compare to N250 naked version?</h3>
<p class="faq-answer">The F250 offers full fairing for wind protection and sports aesthetics at ৳35,000 premium over N250. F250 is better for highway touring, N250 for city riding and maneuverability.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is oil cooling sufficient for 250cc?</h3>
<p class="faq-answer">Yes, oil cooling is adequate for the F250's performance level and touring use. While liquid cooling offers better heat management, oil cooling keeps costs down and provides sufficient cooling for normal riding.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What's the top speed?</h3>
<p class="faq-answer">The Pulsar F250 can achieve approximately 135-140 km/h top speed. The full fairing helps with aerodynamics, allowing comfortable highway cruising at 100-110 km/h.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar F250 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,65,100, the Bajaj Pulsar F250 offers outstanding value as an affordable faired 250cc sports tourer with full aerodynamic fairing, dual-channel ABS safety, good performance (24.5 PS), wind protection, and decent fuel economy (35-40 km/l). It excels at highway touring and long-distance riding with proven Pulsar reliability. However, oil-cooled engine limitations, heavy weight (162 kg), aggressive riding position, average build quality, and high-RPM vibrations make it best suited for budget touring enthusiasts, highway commuters, and riders seeking affordable faired sports capability over city agility and premium refinement.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For affordable faired 250cc sports touring with ABS safety</span></p>
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
		return fmt.Errorf("error creating Bajaj Pulsar F250 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar F250 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার এফ২৫০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ পালসার এফ২৫০ হল একটি সম্পূর্ণ-ফেয়ারড ২৫০সিসি স্পোর্টস ট্যুরার যেখানে আক্রমণাত্মক অ্যারোডাইনামিক ফেয়ারিং, অয়েল-কুলড ইঞ্জিন, ডুয়াল-চ্যানেল এবিএস, এলইডি লাইটিং এবং ট্যুরিং-ফোকাসড এরগোনমিক্স রয়েছে। ৳৩,৬৫,১০০ টাকায় মূল্যায়িত, এটি স্পোর্টস বাইক নান্দনিকতা, বায়ু সুরক্ষা, ভাল পারফরমেন্স, জ্বালানি দক্ষতা এবং পালসারের প্রমাণিত নির্ভরযোগ্যতা সহ সাশ্রয়ী ফেয়ারড স্পোর্টস ট্যুরিং ক্ষমতা খোঁজা রাইডারদের জন্য আদর্শ ব্যবহারিক বৈশিষ্ট্য সহ ২৫০সিসি সেগমেন্টে চমৎকার মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার এফ২৫০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সম্পূর্ণ ফেয়ারিং ডিজাইন:</strong> <span class="highlight-value">বায়ু সুরক্ষা এবং স্পোর্টস নান্দনিকতার জন্য অ্যারোডাইনামিক ফেয়ারিং</span></li>
<li class="highlight-item"><strong class="highlight-label">২৫০সিসি অয়েল-কুলড ইঞ্জিন:</strong> <span class="highlight-value">ভাল তাপ ব্যবস্থাপনা সহ ২৪.৫ পিএস শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">ডুয়াল-চ্যানেল এবিএস:</strong> <span class="highlight-value">উভয় চাকায় এবিএস সহ উন্নত নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">সম্পূর্ণ এলইডি হেডলাইট এবং টেইললাইট</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার এফ২৫০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার মূল্য প্রস্তাব:</strong> <span class="pro-description">প্রতিযোগিতামূলক ৳৩,৬৫,১০০ মূল্যে সম্পূর্ণ ফেয়ারিং, এবিএস এবং ২৫০সিসি</span></li>
<li class="pro-item"><strong class="pro-label">ডুয়াল-চ্যানেল এবিএস নিরাপত্তা:</strong> <span class="pro-description">উভয় চাকায় উন্নত ব্রেকিং নিরাপত্তা</span></li>
<li class="pro-item"><strong class="pro-label">ভাল পারফরমেন্স:</strong> <span class="pro-description">২৪.৫ পিএস ট্যুরিংয়ের জন্য উৎসাহী পারফরমেন্স প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">বায়ু সুরক্ষা:</strong> <span class="pro-description">সম্পূর্ণ ফেয়ারিং হাইওয়েতে বায়ু ক্লান্তি হ্রাস করে</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি দক্ষ:</strong> <span class="pro-description">একটি ২৫০সিসি বাইকের জন্য ৩৫-৪০ কিমি/লিটার অর্জন</span></li>
<li class="pro-item"><strong class="pro-label">প্রমাণিত পালসার প্ল্যাটফর্ম:</strong> <span class="pro-description">নির্ভরযোগ্য বাজাজ ইঞ্জিনিয়ারিং এবং বিস্তৃত সেবা নেটওয়ার্ক</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার এফ২৫০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অয়েল-কুলড ইঞ্জিন:</strong> <span class="con-description">লিকুইড-কুলড প্রতিযোগীদের মতো উন্নত নয়</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">১৬২ কেজি শহুরে চালনা চ্যালেঞ্জিং করে</span></li>
<li class="con-item"><strong class="con-label">আক্রমণাত্মক রাইডিং পজিশন:</strong> <span class="con-description">সামনের দিকে ঝুঁকে পড়া এরগোনমিক্স ক্লান্তিকর হতে পারে</span></li>
<li class="con-item"><strong class="con-label">গড় বিল্ড কোয়ালিটি:</strong> <span class="con-description">ফিট এবং ফিনিশ প্রিমিয়াম ব্র্যান্ডের মতো পরিমার্জিত নয়</span></li>
<li class="con-item"><strong class="con-label">উচ্চ আরপিএমে কম্পন:</strong> <span class="con-description">অয়েল-কুলড ইঞ্জিন উচ্চ গতিতে কম্পন দেখায়</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে বাজাজ পালসার এফ২৫০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Budget-conscious sports touring enthusiasts</li>
<li class="best-for-item">Long-distance highway commuters</li>
<li class="best-for-item">First-time 250cc faired bike buyers</li>
<li class="best-for-item">Riders seeking wind protection</li>
<li class="best-for-item">Weekend touring riders</li>
<li class="best-for-item">Those wanting ABS safety at affordable price</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">বাজাজ পালসার এফ২৫০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">City traffic navigation</li>
<li class="not-recommended-item">Performance purists seeking high refinement</li>
<li class="not-recommended-item">Shorter riders uncomfortable with sports ergonomics</li>
<li class="not-recommended-item">Premium build quality seekers</li>
<li class="not-recommended-item">Heavy daily city commuting</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">বাজাজ পালসার এফ২৫০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৩,৬৫,১০০ - এবিএস সহ ফেয়ারড ২৫০সিসির জন্য চমৎকার</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাঝারি - ৳৮-১১ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳৭-৯ প্রতি কিমি (৩৫-৪০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৫,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳৮,০০০-১২,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">বাজাজ পালসার এফ২৫০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- ভাল ২৫০সিসি পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- ভাল জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- ট্যুরিংয়ের জন্য ভাল</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- আক্রমণাত্মক ফেয়ারড ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- অসাধারণ ভ্যালু ফর মানি</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- পর্যাপ্ত বাজাজ গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">বাজাজ পালসার এফ২৫০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এটি এন২৫০ নেকেড সংস্করণের সাথে কীভাবে তুলনা করে?</h3>
<p class="faq-answer">এফ২৫০ এন২৫০ এর চেয়ে ৳৩৫,০০০ প্রিমিয়ামে বায়ু সুরক্ষা এবং স্পোর্টস নান্দনিকতার জন্য সম্পূর্ণ ফেয়ারিং প্রদান করে। এফ২৫০ হাইওয়ে ট্যুরিংয়ের জন্য ভাল, এন২৫০ শহুরে রাইডিং এবং চালনার জন্য।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">২৫০সিসির জন্য অয়েল কুলিং কি যথেষ্ট?</h3>
<p class="faq-answer">হ্যাঁ, অয়েল কুলিং এফ২৫০ এর পারফরমেন্স লেভেল এবং ট্যুরিং ব্যবহারের জন্য পর্যাপ্ত। যদিও লিকুইড কুলিং ভাল তাপ ব্যবস্থাপনা প্রদান করে, অয়েল কুলিং খরচ কম রাখে এবং স্বাভাবিক রাইডিংয়ের জন্য যথেষ্ট কুলিং প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">সর্বোচ্চ গতি কত?</h3>
<p class="faq-answer">পালসার এফ২৫০ প্রায় ১৩৫-১৪০ কিমি/ঘণ্টা সর্বোচ্চ গতি অর্জন করতে পারে। সম্পূর্ণ ফেয়ারিং অ্যারোডাইনামিক্সে সাহায্য করে, ১০০-১১০ কিমি/ঘণ্টায় আরামদায়ক হাইওয়ে ক্রুজিং অনুমতি দেয়।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে বাজাজ পালসার এফ২৫০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৩,৬৫,১০০ টাকায় বাজাজ পালসার এফ২৫০ সম্পূর্ণ অ্যারোডাইনামিক ফেয়ারিং, ডুয়াল-চ্যানেল এবিএস নিরাপত্তা, ভাল পারফরমেন্স (২৪.৫ পিএস), বায়ু সুরক্ষা এবং ভাল জ্বালানি অর্থনীতি (৩৫-৪০ কিমি/লিটার) সহ একটি সাশ্রয়ী ফেয়ারড ২৫০সিসি স্পোর্টস ট্যুরার হিসাবে অসাধারণ মূল্য প্রদান করে। এটি প্রমাণিত পালসার নির্ভরযোগ্যতা সহ হাইওয়ে ট্যুরিং এবং দীর্ঘ-দূরত্ব রাইডিংয়ে শ্রেষ্ঠ। তবে, অয়েল-কুলড ইঞ্জিন সীমাবদ্ধতা, ভারী ওজন (১৬২ কেজি), আক্রমণাত্মক রাইডিং পজিশন, গড় বিল্ড কোয়ালিটি এবং উচ্চ-আরপিএম কম্পন এটিকে বাজেট ট্যুরিং উৎসাহী, হাইওয়ে যাত্রী এবং শহুরে চপলতা এবং প্রিমিয়াম পরিমার্জনের চেয়ে সাশ্রয়ী ফেয়ারড স্পোর্টস ক্ষমতা খোঁজা রাইডারদের জন্য সবচেয়ে উপযুক্ত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- এবিএস নিরাপত্তা সহ সাশ্রয়ী ফেয়ারড ২৫০সিসি স্পোর্টস ট্যুরিংয়ের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar F250 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar F250\n")

	return nil
}
