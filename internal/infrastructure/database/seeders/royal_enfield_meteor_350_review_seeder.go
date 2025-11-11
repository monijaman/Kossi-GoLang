package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// RoyalEnfieldMeteor350ReviewSeeder handles seeding Royal Enfield Meteor 350 product review and translation
type RoyalEnfieldMeteor350ReviewSeeder struct {
	BaseSeeder
}

// NewRoyalEnfieldMeteor350ReviewSeeder creates a new RoyalEnfieldMeteor350ReviewSeeder
func NewRoyalEnfieldMeteor350ReviewSeeder() *RoyalEnfieldMeteor350ReviewSeeder {
	return &RoyalEnfieldMeteor350ReviewSeeder{BaseSeeder: BaseSeeder{name: "Royal Enfield Meteor 350 Review"}}
}

// Seed implements the Seeder interface
func (s *RoyalEnfieldMeteor350ReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Royal Enfield Meteor 350").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Royal Enfield Meteor 350 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Royal Enfield Meteor 350 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Royal Enfield Meteor 350 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Royal Enfield Meteor 350 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Royal Enfield Meteor 350 is a modern cruiser priced at ৳6,80,000, offering refined 350cc engine, comfortable ergonomics, and contemporary features. With Tripper navigation, smooth power delivery, and less vibrations than Classic, it's ideal for riders wanting Royal Enfield character with modern comfort and technology.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Royal Enfield Meteor 350 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Modern Cruiser:</strong> <span class="highlight-value">Contemporary design with RE character</span></li>
<li class="highlight-item"><strong class="highlight-label">Tripper Navigation:</strong> <span class="highlight-value">Turn-by-turn navigation display</span></li>
<li class="highlight-item"><strong class="highlight-label">Refined Engine:</strong> <span class="highlight-value">Smooth, less vibration than Classic</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Ergonomics:</strong> <span class="highlight-value">Relaxed cruiser riding position</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Royal Enfield Meteor 350 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Smooth Engine:</strong> <span class="pro-description">Much smoother than Classic 350</span></li>
<li class="pro-item"><strong class="pro-label">Tripper Navigation:</strong> <span class="pro-description">Bluetooth navigation system included</span></li>
<li class="pro-item"><strong class="pro-label">Less Vibrations:</strong> <span class="pro-description">Better refinement than older REs</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">LED lights, USB charging, digital console</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Excellent cruiser ergonomics</span></li>
<li class="pro-item"><strong class="pro-label">Three Variants:</strong> <span class="pro-description">Fireball, Stellar, Supernova options</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Premium materials and finish</span></li>
<li class="pro-item"><strong class="pro-label">RE Heritage:</strong> <span class="pro-description">Royal Enfield brand value</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Royal Enfield Meteor 350 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Very Expensive:</strong> <span class="con-description">At ৳6,80,000 it's premium priced</span></li>
<li class="con-item"><strong class="con-label">Poor Fuel Economy:</strong> <span class="con-description">Only 32-37 km/l mileage</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">191 kg difficult in traffic</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">RE service centers not widespread</span></li>
<li class="con-item"><strong class="con-label">High Maintenance:</strong> <span class="con-description">Expensive parts and service</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Royal Enfield Meteor 350 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Modern cruiser enthusiasts</li>
<li class="best-for-item">Long-distance highway touring</li>
<li class="best-for-item">Riders wanting RE with refinement</li>
<li class="best-for-item">Technology-appreciating riders</li>
<li class="best-for-item">Weekend leisure riding</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Royal Enfield Meteor 350?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Daily city commuting (too heavy)</li>
<li class="not-recommended-item">Budget buyers</li>
<li class="not-recommended-item">Those needing fuel efficiency</li>
<li class="not-recommended-item">First-time riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Royal Enfield Meteor 350 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳6,70,000 - ৳6,90,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳6,200 - ৳8,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳3,300-3,900/month for 30 km daily at 34 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 6,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳2,800 - ৳4,500 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Royal Enfield Meteor 350 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- Below average</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Premium but refined</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good cruising power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Excellent comfort</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Premium build</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Modern features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Modern cruiser look</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Royal Enfield Meteor 350 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Royal Enfield Meteor 350?</h3>
<p class="faq-answer">Royal Enfield Meteor 350 delivers 32-37 km/l in mixed conditions.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is Tripper navigation?</h3>
<p class="faq-answer">Tripper is Bluetooth-connected turn-by-turn navigation display powered by Google Maps.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Meteor 350 or Classic 350?</h3>
<p class="faq-answer">Meteor 350 is more refined with modern features; Classic 350 offers iconic retro styling.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Meteor 350?</h3>
<p class="faq-answer">Royal Enfield Meteor 350 is priced between ৳6,70,000 to ৳6,90,000.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Royal Enfield Meteor 350 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Royal Enfield Meteor 350 is the refined modern cruiser at ৳6,80,000, offering smoother engine, Tripper navigation, and contemporary features while maintaining Royal Enfield character. Significantly less vibration than Classic 350 and better refinement make it ideal for long-distance touring. While expensive and fuel economy is average (32-37 km/l), the excellent comfort, modern technology, and premium build justify the price for riders wanting RE heritage with 21st-century refinement.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For refined modern cruiser with RE character</span></p>
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
		return fmt.Errorf("error creating Royal Enfield Meteor 350 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Royal Enfield Meteor 350 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">রয়াল এনফিল্ড মিটিওর ৩৫০ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">রয়াল এনফিল্ড মিটিওর ৩৫০ একটি আধুনিক ক্রুজার যার মূল্য ৳৬,৮০,০০০ টাকা, যা পরিমার্জিত ৩৫০সিসি ইঞ্জিন, আরামদায়ক এরগোনমিক্স এবং সমসাময়িক ফিচার প্রদান করে। ট্রিপার নেভিগেশন, মসৃণ পাওয়ার ডেলিভারি এবং ক্লাসিকের চেয়ে কম কম্পন সহ, এটি আধুনিক আরাম এবং প্রযুক্তি সহ রয়াল এনফিল্ড চরিত্র চান রাইডারদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রয়াল এনফিল্ড মিটিওর ৩৫০ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">আধুনিক ক্রুজার:</strong> <span class="highlight-value">আরই চরিত্র সহ সমসাময়িক ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্রিপার নেভিগেশন:</strong> <span class="highlight-value">টার্ন-বাই-টার্ন নেভিগেশন ডিসপ্লে</span></li>
<li class="highlight-item"><strong class="highlight-label">পরিমার্জিত ইঞ্জিন:</strong> <span class="highlight-value">ক্লাসিকের চেয়ে মসৃণ, কম কম্পন</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক এরগোনমিক্স:</strong> <span class="highlight-value">স্বাচ্ছন্দ্যময় ক্রুজার রাইডিং পজিশন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রয়াল এনফিল্ড মিটিওর ৩৫০ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">মসৃণ ইঞ্জিন:</strong> <span class="pro-description">ক্লাসিক ৩৫০ এর চেয়ে অনেক মসৃণ</span></li>
<li class="pro-item"><strong class="pro-label">ট্রিপার নেভিগেশন:</strong> <span class="pro-description">ব্লুটুথ নেভিগেশন সিস্টেম অন্তর্ভুক্ত</span></li>
<li class="pro-item"><strong class="pro-label">কম কম্পন:</strong> <span class="pro-description">পুরানো আরই-র চেয়ে ভালো পরিমার্জন</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক ফিচার:</strong> <span class="pro-description">এলইডি লাইট, ইউএসবি চার্জিং, ডিজিটাল কনসোল</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">চমৎকার ক্রুজার এরগোনমিক্স</span></li>
<li class="pro-item"><strong class="pro-label">তিনটি ভ্যারিয়েন্ট:</strong> <span class="pro-description">ফায়ারবল, স্টেলার, সুপারনোভা অপশন</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">প্রিমিয়াম ম্যাটেরিয়াল এবং ফিনিশ</span></li>
<li class="pro-item"><strong class="pro-label">আরই ঐতিহ্য:</strong> <span class="pro-description">রয়াল এনফিল্ড ব্র্যান্ড ভ্যালু</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রয়াল এনফিল্ড মিটিওর ৩৫০ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">খুব দামি:</strong> <span class="con-description">৳৬,৮০,০০০ টাকায় এটি প্রিমিয়াম মূল্যের</span></li>
<li class="con-item"><strong class="con-label">খারাপ জ্বালানি সাশ্রয়:</strong> <span class="con-description">মাত্র ৩২-৩৭ কিমি/লিটার মাইলেজ</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">১৯১ কেজি ট্রাফিকে কঠিন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">আরই সার্ভিস সেন্টার ব্যাপক নয়</span></li>
<li class="con-item"><strong class="con-label">উচ্চ রক্ষণাবেক্ষণ:</strong> <span class="con-description">দামি পার্টস এবং সার্ভিস</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে রয়াল এনফিল্ড মিটিওর ৩৫০ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Modern cruiser enthusiasts</li>
<li class="best-for-item">Long-distance highway touring</li>
<li class="best-for-item">Riders wanting RE with refinement</li>
<li class="best-for-item">Technology-appreciating riders</li>
<li class="best-for-item">Weekend leisure riding</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">রয়াল এনফিল্ড মিটিওর ৩৫০ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Daily city commuting (too heavy)</li>
<li class="not-recommended-item">Budget buyers</li>
<li class="not-recommended-item">Those needing fuel efficiency</li>
<li class="not-recommended-item">First-time riders</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">রয়াল এনফিল্ড মিটিওর ৩৫০ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৬,৭০,০০০ - ৳৬,৯০,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳৬,২০০ - ৳৮,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৩৪ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳৩,৩০০-৩,৯০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৬,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳২,৮০০ - ৳৪,৫০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">রয়াল এনফিল্ড মিটিওর ৩৫০ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3/5)</span> <span class="rating-note">- গড়ের নিচে</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- প্রিমিয়াম কিন্তু পরিমার্জিত</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ক্রুজিং পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- চমৎকার আরাম</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- প্রিমিয়াম বিল্ড</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- আধুনিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- আধুনিক ক্রুজার লুক</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">রয়াল এনফিল্ড মিটিওর ৩৫০ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">রয়াল এনফিল্ড মিটিওর ৩৫০-এর মাইলেজ কত?</h3>
<p class="faq-answer">রয়াল এনফিল্ড মিটিওর ৩৫০ মিশ্র পরিস্থিতিতে ৩২-৩৭ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ট্রিপার নেভিগেশন কী?</h3>
<p class="faq-answer">ট্রিপার হল গুগল ম্যাপস দ্বারা চালিত ব্লুটুথ-সংযুক্ত টার্ন-বাই-টার্ন নেভিগেশন ডিসপ্লে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: মিটিওর ৩৫০ নাকি ক্লাসিক ৩৫০?</h3>
<p class="faq-answer">মিটিওর ৩৫০ আধুনিক ফিচার সহ আরও পরিমার্জিত; ক্লাসিক ৩৫০ আইকনিক রেট্রো স্টাইলিং প্রদান করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">মিটিওর ৩৫০-এর দাম কত?</h3>
<p class="faq-answer">রয়াল এনফিল্ড মিটিওর ৩৫০-এর দাম ৳৬,৭০,০০০ থেকে ৳৬,৯০,০০০ টাকার মধ্যে।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে রয়াল এনফিল্ড মিটিওর ৩৫০ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">রয়াল এনফিল্ড মিটিওর ৩৫০ ৳৬,৮০,০০০ টাকায় পরিমার্জিত আধুনিক ক্রুজার, যা রয়াল এনফিল্ড চরিত্র বজায় রেখে মসৃণ ইঞ্জিন, ট্রিপার নেভিগেশন এবং সমসাময়িক ফিচার প্রদান করে। ক্লাসিক ৩৫০ এর চেয়ে উল্লেখযোগ্যভাবে কম কম্পন এবং ভালো পরিমার্জন এটিকে দীর্ঘ-দূরত্ব ট্যুরিংয়ের জন্য আদর্শ করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আরই চরিত্র সহ পরিমার্জিত আধুনিক ক্রুজারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Royal Enfield Meteor 350 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Royal Enfield Meteor 350\n")

	return nil
}
