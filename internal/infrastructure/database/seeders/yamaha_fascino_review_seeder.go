package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaFascinoReviewSeeder handles seeding Yamaha Fascino product review and translation
type YamahaFascinoReviewSeeder struct {
	BaseSeeder
}

// NewYamahaFascinoReviewSeeder creates a new YamahaFascinoReviewSeeder
func NewYamahaFascinoReviewSeeder() *YamahaFascinoReviewSeeder {
	return &YamahaFascinoReviewSeeder{BaseSeeder: BaseSeeder{name: "Yamaha Fascino Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaFascinoReviewSeeder) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha Fascino").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha Fascino product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha Fascino product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha Fascino already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha Fascino Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha Fascino is a stylish 125cc scooter priced at ৳1,70,000, offering premium design, good build quality, and modern features. With LED lights, comfortable ride, and Yamaha reliability, it's ideal for young urban riders and women seeking a fashionable yet practical scooter.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha Fascino Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Stylish Design:</strong> <span class="highlight-value">Premium fashionable looks</span></li>
<li class="highlight-item"><strong class="highlight-label">LED Lighting:</strong> <span class="highlight-value">Full LED headlight and taillight</span></li>
<li class="highlight-item"><strong class="highlight-label">125cc Power:</strong> <span class="highlight-value">More powerful than 110cc</span></li>
<li class="highlight-item"><strong class="highlight-label">Yamaha Quality:</strong> <span class="highlight-value">Reliable Japanese engineering</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha Fascino Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Premium Styling:</strong> <span class="pro-description">Beautiful, fashionable design</span></li>
<li class="pro-item"><strong class="pro-label">Full LED Lights:</strong> <span class="pro-description">Bright LED headlight and taillight</span></li>
<li class="pro-item"><strong class="pro-label">Good Power:</strong> <span class="pro-description">125cc offers better performance</span></li>
<li class="pro-item"><strong class="pro-label">Yamaha Reliability:</strong> <span class="pro-description">Proven Japanese reliability</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ride:</strong> <span class="pro-description">Smooth and comfortable</span></li>
<li class="pro-item"><strong class="pro-label">Good Build Quality:</strong> <span class="pro-description">Solid Japanese construction</span></li>
<li class="pro-item"><strong class="pro-label">Digital Console:</strong> <span class="pro-description">Modern instrument cluster</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha Fascino Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Higher Price:</strong> <span class="con-description">₹5,000 more than Activa</span></li>
<li class="con-item"><strong class="con-label">Lower Mileage:</strong> <span class="con-description">40-45 km/l, less than 110cc</span></li>
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Fewer centers than Honda</span></li>
<li class="con-item"><strong class="con-label">Smaller Storage:</strong> <span class="con-description">Less under-seat space</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha Fascino in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban women</li>
<li class="best-for-item">Style-conscious buyers</li>
<li class="best-for-item">College/office commuters</li>
<li class="best-for-item">Those wanting modern features</li>
<li class="best-for-item">City riding enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha Fascino?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Large storage need</li>
<li class="not-recommended-item">Areas without Yamaha service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha Fascino Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,65,000 - ৳1,75,000</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Approximately ৳2,600 - ৳3,200 per month</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2,100-2,600/month for 30 km daily at 42 km/l</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000 km or 3 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳800 - ৳1,200 per service</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha Fascino Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- Average mileage</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Yamaha reliability</span></li>
<li class="rating-item"><strong class="rating-label">Value for Money:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good value</span></li>
<li class="rating-item"><strong class="rating-label">Power & Performance:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good 125cc power</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Comfortable</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- Good quality</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- Modern features</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- Very stylish</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha Fascino Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the mileage of Yamaha Fascino?</h3>
<p class="faq-answer">Yamaha Fascino delivers 40-45 km/l with 125cc engine.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What is the price of Yamaha Fascino?</h3>
<p class="faq-answer">Yamaha Fascino is priced between ৳1,65,000 to ৳1,75,000.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is Yamaha Fascino good for women?</h3>
<p class="faq-answer">Yes, Fascino is designed for women with stylish looks and easy handling.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Which is better: Fascino or Activa?</h3>
<p class="faq-answer">Fascino offers better styling and features; Activa is more reliable with better mileage.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha Fascino in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">Yamaha Fascino at ৳1,70,000 is the best choice for style-conscious buyers, offering premium design, full LED lights, and modern features. While mileage (40-45 km/l) is lower than 110cc scooters and price is slightly higher, the 125cc power, beautiful styling, and Yamaha reliability make it perfect for young urban riders. Ideal for those prioritizing looks and features over fuel economy.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For stylish modern scooter with premium features</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Yamaha Fascino review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha Fascino (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা ফ্যাসিনো রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা ফ্যাসিনো একটি স্টাইলিশ ১২৫সিসি স্কুটার যার মূল্য ৳১,৭০,০০০ টাকা, যা প্রিমিয়াম ডিজাইন, ভালো বিল্ড কোয়ালিটি এবং আধুনিক ফিচার প্রদান করে। এলইডি লাইট, আরামদায়ক রাইড এবং ইয়ামাহা নির্ভরযোগ্যতা সহ, এটি ফ্যাশনেবল তবুও ব্যবহারিক স্কুটার খোঁজা তরুণ শহুরে রাইডার এবং মহিলাদের জন্য আদর্শ।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা ফ্যাসিনো এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">স্টাইলিশ ডিজাইন:</strong> <span class="highlight-value">প্রিমিয়াম ফ্যাশনেবল লুক</span></li>
<li class="highlight-item"><strong class="highlight-label">এলইডি লাইটিং:</strong> <span class="highlight-value">সম্পূর্ণ এলইডি হেডলাইট এবং টেইললাইট</span></li>
<li class="highlight-item"><strong class="highlight-label">১২৫সিসি পাওয়ার:</strong> <span class="highlight-value">১১০সিসির চেয়ে শক্তিশালী</span></li>
<li class="highlight-item"><strong class="highlight-label">ইয়ামাহা মান:</strong> <span class="highlight-value">নির্ভরযোগ্য জাপানি ইঞ্জিনিয়ারিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা ফ্যাসিনো এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম স্টাইলিং:</strong> <span class="pro-description">সুন্দর, ফ্যাশনেবল ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">সম্পূর্ণ এলইডি লাইট:</strong> <span class="pro-description">উজ্জ্বল এলইডি হেডলাইট এবং টেইললাইট</span></li>
<li class="pro-item"><strong class="pro-label">ভালো পাওয়ার:</strong> <span class="pro-description">১২৫সিসি ভালো পারফরম্যান্স প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ইয়ামাহা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত জাপানি নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক রাইড:</strong> <span class="pro-description">মসৃণ এবং আরামদায়ক</span></li>
<li class="pro-item"><strong class="pro-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="pro-description">শক্ত জাপানি নির্মাণ</span></li>
<li class="pro-item"><strong class="pro-label">ডিজিটাল কনসোল:</strong> <span class="pro-description">আধুনিক ইন্সট্রুমেন্ট ক্লাস্টার</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা ফ্যাসিনো এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বেশি দাম:</strong> <span class="con-description">অ্যাকটিভার চেয়ে ৳৫,০০০ টাকা বেশি</span></li>
<li class="con-item"><strong class="con-label">কম মাইলেজ:</strong> <span class="con-description">৪০-৪৫ কিমি/লিটার, ১১০সিসির চেয়ে কম</span></li>
<li class="con-item"><strong class="con-label">সীমিত সার্ভিস:</strong> <span class="con-description">হন্ডার চেয়ে কম সেন্টার</span></li>
<li class="con-item"><strong class="con-label">ছোট স্টোরেজ:</strong> <span class="con-description">সিটের নিচে কম জায়গা</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা ফ্যাসিনো কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Young urban women</li>
<li class="best-for-item">Style-conscious buyers</li>
<li class="best-for-item">College/office commuters</li>
<li class="best-for-item">Those wanting modern features</li>
<li class="best-for-item">City riding enthusiasts</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা ফ্যাসিনো কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget buyers</li>
<li class="not-recommended-item">Those prioritizing mileage</li>
<li class="not-recommended-item">Large storage need</li>
<li class="not-recommended-item">Areas without Yamaha service</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা ফ্যাসিনো এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳১,৬৫,০০০ - ৳১,৭৫,০০০</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">মাসে প্রায় ৳২,৬০০ - ৳৩,২০০</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৪২ কিমি/লিটারে দৈনিক ৩০ কিমির জন্য ৳২,১০০-২,৬০০/মাস</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০ কিমি বা ৩ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">৳৮০০ - ৳১,২০০ প্রতি সার্ভিস</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা ফ্যাসিনো পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">জ্বালানি সাশ্রয়:</strong> <span class="rating-score">(3.5/5)</span> <span class="rating-note">- গড় মাইলেজ</span></li>
<li class="rating-item"><strong class="rating-label">নির্ভরযোগ্যতা:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- ইয়ামাহা নির্ভরযোগ্যতা</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু ফর মানি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ভ্যালু</span></li>
<li class="rating-item"><strong class="rating-label">পাওয়ার ও পারফরম্যান্স:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো ১২৫সিসি পাওয়ার</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- আরামদায়ক</span></li>
<li class="rating-item"><strong class="rating-label">বিল্ড কোয়ালিটি:</strong> <span class="rating-score">(4/5)</span> <span class="rating-note">- ভালো মান</span></li>
<li class="rating-item"><strong class="rating-label">ফিচার:</strong> <span class="rating-score">(4.5/5)</span> <span class="rating-note">- আধুনিক ফিচার</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইলিং:</strong> <span class="rating-score">(5/5)</span> <span class="rating-note">- খুব স্টাইলিশ</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা ফ্যাসিনো সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা ফ্যাসিনোর মাইলেজ কত?</h3>
<p class="faq-answer">ইয়ামাহা ফ্যাসিনো ১২৫সিসি ইঞ্জিন সহ ৪০-৪৫ কিমি/লিটার মাইলেজ দেয়।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">ইয়ামাহা ফ্যাসিনোর দাম কত?</h3>
<p class="faq-answer">ইয়ামাহা ফ্যাসিনোর দাম ৳১,৬৫,০০০ থেকে ৳১,৭৫,০০০ টাকার মধ্যে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">মহিলাদের জন্য ইয়ামাহা ফ্যাসিনো কি ভালো?</h3>
<p class="faq-answer">হ্যাঁ, ফ্যাসিনো স্টাইলিশ লুক এবং সহজ হ্যান্ডলিং সহ মহিলাদের জন্য ডিজাইন করা।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কোনটি ভালো: ফ্যাসিনো নাকি অ্যাকটিভা?</h3>
<p class="faq-answer">ফ্যাসিনো ভালো স্টাইলিং এবং ফিচার প্রদান করে; অ্যাকটিভা ভালো মাইলেজ সহ আরও নির্ভরযোগ্য।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা ফ্যাসিনো কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳১,৭০,০০০ টাকায় ইয়ামাহা ফ্যাসিনো স্টাইল-সচেতন ক্রেতাদের জন্য সেরা পছন্দ, যা প্রিমিয়াম ডিজাইন, সম্পূর্ণ এলইডি লাইট এবং আধুনিক ফিচার প্রদান করে। যদিও মাইলেজ (৪০-৪৫ কিমি/লিটার) ১১০সিসি স্কুটারের চেয়ে কম এবং দাম সামান্য বেশি, ১২৫সিসি পাওয়ার, সুন্দর স্টাইলিং এবং ইয়ামাহা নির্ভরযোগ্যতা এটিকে তরুণ শহুরে রাইডারদের জন্য নিখুঁত করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম ফিচার সহ স্টাইলিশ আধুনিক স্কুটারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha Fascino already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha Fascino\n")

	return nil
}
