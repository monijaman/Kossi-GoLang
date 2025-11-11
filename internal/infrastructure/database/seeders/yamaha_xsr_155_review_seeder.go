package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaXSR155Review handles seeding Yamaha XSR 155 product review and translation
type YamahaXSR155Review struct {
	BaseSeeder
}

// NewYamahaXSR155ReviewSeeder creates a new YamahaXSR155Review
func NewYamahaXSR155ReviewSeeder() *YamahaXSR155Review {
	return &YamahaXSR155Review{BaseSeeder: BaseSeeder{name: "Yamaha XSR 155 Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaXSR155Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha XSR 155").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Yamaha XSR 155 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Yamaha XSR 155 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha XSR 155 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha XSR 155 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Yamaha XSR 155 is a premium 155cc neo-retro motorcycle featuring VVA technology, retro-modern design, and advanced electronics. Priced at ৳5,45,000, it combines classic café racer aesthetics with modern performance including variable valve actuation, digital instrumentation, and premium build quality for riders seeking unique style and authentic Yamaha heritage.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha XSR 155 Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Neo-Retro Design:</strong> <span class="highlight-value">Classic café racer aesthetics with modern touches</span></li>
<li class="highlight-item"><strong class="highlight-label">VVA Technology:</strong> <span class="highlight-value">Variable Valve Actuation for optimal performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Materials:</strong> <span class="highlight-value">High-quality finishes and components</span></li>
<li class="highlight-item"><strong class="highlight-label">Digital Instrumentation:</strong> <span class="highlight-value">Modern LCD display with comprehensive information</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha XSR 155 Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Unique Retro Styling:</strong> <span class="pro-description">Distinctive neo-classic design that stands out from crowd</span></li>
<li class="pro-item"><strong class="pro-label">VVA Performance:</strong> <span class="pro-description">Variable valve actuation provides excellent power delivery</span></li>
<li class="pro-item"><strong class="pro-label">Premium Build Quality:</strong> <span class="pro-description">High-quality materials and excellent fit-finish</span></li>
<li class="pro-item"><strong class="pro-label">Heritage Appeal:</strong> <span class="pro-description">Yamaha's classic design heritage with modern reliability</span></li>
<li class="pro-item"><strong class="pro-label">Customization Potential:</strong> <span class="pro-description">Excellent platform for personalization and modifications</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">Digital display, LED lighting, and contemporary electronics</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha XSR 155 Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Purchase Price:</strong> <span class="con-description">Premium pricing for retro styling and VVA technology</span></li>
<li class="con-item"><strong class="con-label">Limited Practicality:</strong> <span class="con-description">Style-focused design may compromise daily usability</span></li>
<li class="con-item"><strong class="con-label">Expensive Maintenance:</strong> <span class="con-description">Premium components require specialized service and parts</span></li>
<li class="con-item"><strong class="con-label">Limited Service Network:</strong> <span class="con-description">Fewer Yamaha dealerships and trained technicians</span></li>
<li class="con-item"><strong class="con-label">Niche Appeal:</strong> <span class="con-description">Retro styling may not appeal to all riders</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy Yamaha XSR 155 in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Style-conscious riders seeking unique design</li>
<li class="best-for-item">Café racer and retro enthusiasts</li>
<li class="best-for-item">Weekend leisure riders</li>
<li class="best-for-item">Motorcycle collectors and enthusiasts</li>
<li class="best-for-item">Urban riders wanting premium aesthetics</li>
<li class="best-for-item">Those appreciating Yamaha heritage</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy Yamaha XSR 155?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Heavy daily commuting</li>
<li class="not-recommended-item">Performance-focused riders</li>
<li class="not-recommended-item">Those needing maximum practicality</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha XSR 155 Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳5,45,000 - Premium for retro styling and VVA</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳12-15 per km with premium maintenance</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳10-12 per km (35-40 km/l with VVA efficiency)</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,500 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳10,000-16,000 annually</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha XSR 155 Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Good VVA performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- Moderate fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Decent riding position</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Outstanding retro design</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- Premium for uniqueness</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent Yamaha quality</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">Yamaha XSR 155 Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What makes the XSR 155 different from other 155cc bikes?</h3>
<p class="faq-answer">The XSR 155 features unique neo-retro styling, premium materials, VVA technology, and café racer aesthetics that distinguish it from conventional motorcycles in this segment.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for daily commuting?</h3>
<p class="faq-answer">While possible, the XSR 155 is more suited for leisure riding and weekend trips. The retro styling and premium positioning make it better for occasional use rather than heavy daily commuting.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">How is the customization potential?</h3>
<p class="faq-answer">Excellent! The XSR 155's retro design provides a great platform for customization. Many aftermarket parts are available for café racer, scrambler, or personal modifications.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha XSR 155 in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳5,45,000, the Yamaha XSR 155 offers a unique proposition in the 155cc segment with its distinctive neo-retro styling, VVA technology, and premium build quality. It appeals to style-conscious riders seeking individuality and Yamaha heritage. However, the high pricing, limited daily practicality, expensive maintenance, and niche appeal make it suitable primarily for enthusiasts, collectors, and weekend riders who prioritize distinctive aesthetics and premium features over conventional practicality and value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For unique neo-retro styling with premium VVA performance</span></p>
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
		return fmt.Errorf("error creating Yamaha XSR 155 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha XSR 155 (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা এক্সএসআর ১৫৫ রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">ইয়ামাহা এক্সএসআর ১৫৫ হল একটি প্রিমিয়াম ১৫৫সিসি নিও-রেট্রো মোটরসাইকেল যেখানে ভিভিএ প্রযুক্তি, রেট্রো-আধুনিক ডিজাইন এবং উন্নত ইলেকট্রনিক্স রয়েছে। ৳৫,৪৫,০০০ টাকায় মূল্যায়িত, এটি ভেরিয়েবল ভালভ অ্যাকচুয়েশন, ডিজিটাল ইনস্ট্রুমেন্টেশন এবং অনন্য স্টাইল ও খাঁটি ইয়ামাহা ঐতিহ্য খোঁজা রাইডারদের জন্য প্রিমিয়াম বিল্ড কোয়ালিটি সহ আধুনিক পারফরমেন্সের সাথে ক্লাসিক ক্যাফে রেসার নান্দনিকতাকে একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা এক্সএসআর ১৫৫ এর মূল বৈশিষ্ট্য ও সুবিধা</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">নিও-রেট্রো ডিজাইন:</strong> <span class="highlight-value">আধুনিক স্পর্শ সহ ক্লাসিক ক্যাফে রেসার নান্দনিকতা</span></li>
<li class="highlight-item"><strong class="highlight-label">ভিভিএ প্রযুক্তি:</strong> <span class="highlight-value">সর্বোত্তম পারফরমেন্সের জন্য ভেরিয়েবল ভালভ অ্যাকচুয়েশন</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম ম্যাটেরিয়াল:</strong> <span class="highlight-value">উচ্চ-মানের ফিনিশ এবং কম্পোনেন্ট</span></li>
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল ইনস্ট্রুমেন্টেশন:</strong> <span class="highlight-value">ব্যাপক তথ্য সহ আধুনিক এলসিডি ডিসপ্লে</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা এক্সএসআর ১৫৫ এর সুবিধা - কেন কিনবেন এই বাইক?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অনন্য রেট্রো স্টাইলিং:</strong> <span class="pro-description">ভিড় থেকে আলাদা হয়ে দাঁড়ানো স্বতন্ত্র নিও-ক্লাসিক ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">ভিভিএ পারফরমেন্স:</strong> <span class="pro-description">ভেরিয়েবল ভালভ অ্যাকচুয়েশন চমৎকার পাওয়ার ডেলিভারি প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বিল্ড কোয়ালিটি:</strong> <span class="pro-description">উচ্চ-মানের উপকরণ এবং চমৎকার ফিট-ফিনিশ</span></li>
<li class="pro-item"><strong class="pro-label">ঐতিহ্যের আবেদন:</strong> <span class="pro-description">আধুনিক নির্ভরযোগ্যতা সহ ইয়ামাহার ক্লাসিক ডিজাইন ঐতিহ্য</span></li>
<li class="pro-item"><strong class="pro-label">কাস্টমাইজেশন সম্ভাবনা:</strong> <span class="pro-description">ব্যক্তিগতকরণ এবং পরিবর্তনের জন্য চমৎকার প্ল্যাটফর্ম</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক বৈশিষ্ট্য:</strong> <span class="pro-description">ডিজিটাল ডিসপ্লে, এলইডি লাইটিং এবং সমসাময়িক ইলেকট্রনিক্স</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা এক্সএসআর ১৫৫ এর অসুবিধা - সমস্যা কি কি?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ ক্রয় মূল্য:</strong> <span class="con-description">রেট্রো স্টাইলিং এবং ভিভিএ প্রযুক্তির জন্য প্রিমিয়াম মূল্য</span></li>
<li class="con-item"><strong class="con-label">সীমিত ব্যবহারিকতা:</strong> <span class="con-description">স্টাইল-ফোকাসড ডিজাইন দৈনিক ব্যবহারযোগ্যতা ক্ষুণ্ন করতে পারে</span></li>
<li class="con-item"><strong class="con-label">ব্যয়বহুল রক্ষণাবেক্ষণ:</strong> <span class="con-description">প্রিমিয়াম কম্পোনেন্টের বিশেষায়িত সেবা এবং যন্ত্রাংশ প্রয়োজন</span></li>
<li class="con-item"><strong class="con-label">সীমিত সেবা নেটওয়ার্ক:</strong> <span class="con-description">কম ইয়ামাহা ডিলারশিপ এবং প্রশিক্ষিত প্রযুক্তিবিদ</span></li>
<li class="con-item"><strong class="con-label">নিশ আবেদন:</strong> <span class="con-description">রেট্রো স্টাইলিং সব রাইডারের কাছে আবেদন নাও করতে পারে</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">বাংলাদেশে ইয়ামাহা এক্সএসআর ১৫৫ কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">Style-conscious riders seeking unique design</li>
<li class="best-for-item">Café racer and retro enthusiasts</li>
<li class="best-for-item">Weekend leisure riders</li>
<li class="best-for-item">Motorcycle collectors and enthusiasts</li>
<li class="best-for-item">Urban riders wanting premium aesthetics</li>
<li class="best-for-item">Those appreciating Yamaha heritage</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">ইয়ামাহা এক্সএসআর ১৫৫ কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Budget-conscious buyers</li>
<li class="not-recommended-item">Long-distance touring</li>
<li class="not-recommended-item">Heavy daily commuting</li>
<li class="not-recommended-item">Performance-focused riders</li>
<li class="not-recommended-item">Those needing maximum practicality</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">ইয়ামাহা এক্সএসআর ১৫৫ এর দাম ও খরচ বিশ্লেষণ বাংলাদেশে</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">মূল্য সীমা:</strong> <span class="value-amount">৳৫,৪৫,০০০ - রেট্রো স্টাইলিং এবং ভিভিএর জন্য প্রিমিয়াম</span></p>
<p class="value-item"><strong class="value-label">রানিং খরচ:</strong> <span class="value-amount">উচ্চ - প্রিমিয়াম রক্ষণাবেক্ষণ সহ ৳১২-১৫ প্রতি কিমি</span></p>
<p class="value-item"><strong class="value-label">মাসিক জ্বালানি খরচ:</strong> <span class="value-amount">৳১০-১২ প্রতি কিমি (ভিভিএ দক্ষতা সহ ৩৫-৪০ কিমি/লিটার)</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,৫০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">বার্ষিক ৳১০,০০০-১৬,০০০</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">ইয়ামাহা এক্সএসআর ১৫৫ পারফরম্যান্স রেটিং ও স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরমেন্স:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- ভাল ভিভিএ পারফরমেন্স</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">3.8</span> <span class="rating-note">- মাঝারি জ্বালানি দক্ষতা</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- শোভন রাইডিং পজিশন</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- অসাধারণ রেট্রো ডিজাইন</span></li>
<li class="rating-item"><strong class="rating-label">ভ্যালু:</strong> <span class="rating-score">3.9</span> <span class="rating-note">- অনন্যতার জন্য প্রিমিয়াম</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- চমৎকার ইয়ামাহা গুণমান</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">ইয়ামাহা এক্সএসআর ১৫৫ সম্পর্কে প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এক্সএসআর ১৫৫ কে অন্যান্য ১৫৫সিসি বাইক থেকে কী আলাদা করে?</h3>
<p class="faq-answer">এক্সএসআর ১৫৫ এ অনন্য নিও-রেট্রো স্টাইলিং, প্রিমিয়াম ম্যাটেরিয়াল, ভিভিএ প্রযুক্তি এবং ক্যাফে রেসার নান্দনিকতা রয়েছে যা এটিকে এই সেগমেন্টের প্রচলিত মোটরসাইকেল থেকে আলাদা করে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি দৈনিক যাতায়াতের জন্য উপযুক্ত?</h3>
<p class="faq-answer">যদিও সম্ভব, এক্সএসআর ১৫৫ অবসর রাইডিং এবং সাপ্তাহিক ট্রিপের জন্য আরও উপযুক্ত। রেট্রো স্টাইলিং এবং প্রিমিয়াম অবস্থান এটিকে ভারী দৈনিক যাতায়াতের চেয়ে মাঝে মাঝে ব্যবহারের জন্য ভাল করে তোলে।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">কাস্টমাইজেশন সম্ভাবনা কেমন?</h3>
<p class="faq-answer">চমৎকার! এক্সএসআর ১৫৫ এর রেট্রো ডিজাইন কাস্টমাইজেশনের জন্য একটি দুর্দান্ত প্ল্যাটফর্ম প্রদান করে। ক্যাফে রেসার, স্ক্র্যাম্বলার বা ব্যক্তিগত পরিবর্তনের জন্য অনেক আফটার মার্কেট পার্টস উপলব্ধ।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাংলাদেশে ইয়ামাহা এক্সএসআর ১৫৫ কিনবেন কি?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳৫,৪৫,০০০ টাকায় ইয়ামাহা এক্সএসআর ১৫৫ তার স্বতন্ত্র নিও-রেট্রো স্টাইলিং, ভিভিএ প্রযুক্তি এবং প্রিমিয়াম বিল্ড কোয়ালিটি দিয়ে ১৫৫সিসি সেগমেন্টে একটি অনন্য প্রস্তাব প্রদান করে। এটি ব্যক্তিত্ব এবং ইয়ামাহা ঐতিহ্য খোঁজা স্টাইল-সচেতন রাইডারদের কাছে আবেদন করে। তবে, উচ্চ মূল্য, সীমিত দৈনিক ব্যবহারিকতা, ব্যয়বহুল রক্ষণাবেক্ষণ এবং নিশ আবেদন এটিকে প্রধানত উৎসাহী, সংগ্রাহক এবং সাপ্তাহিক রাইডারদের জন্য উপযুক্ত করে তোলে যারা প্রচলিত ব্যবহারিকতা এবং মূল্যের চেয়ে স্বতন্ত্র নান্দনিকতা এবং প্রিমিয়াম বৈশিষ্ট্যকে অগ্রাধিকার দেয়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম ভিভিএ পারফরমেন্স সহ অনন্য নিও-রেট্রো স্টাইলিংয়ের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha XSR 155 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha XSR 155\n")

	return nil
}
