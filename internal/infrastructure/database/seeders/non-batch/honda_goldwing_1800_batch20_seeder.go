package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HondaGoldwing1800ReviewBatch20 struct {
	BaseSeeder
}

func NewHondaGoldwing1800ReviewBatch20Seeder() *HondaGoldwing1800ReviewBatch20 {
	return &HondaGoldwing1800ReviewBatch20{BaseSeeder: BaseSeeder{name: "Honda Goldwing 1800 Batch20 Review"}}
}

func (s *HondaGoldwing1800ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda Goldwing 1800").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda goldwing 1800 product not found")
		}
		return fmt.Errorf("error finding product: %w", err)
	}

	var existingReview models.ProductReviewModel
	if err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error; err == nil {
		fmt.Printf("   ℹ️  Review already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda Goldwing 1800 Review Bangladesh 2024 - Touring Luxury Legend</h1>
<p class="summary-text">The Honda Goldwing 1800 is an 1833cc liquid-cooled touring masterpiece representing the ultimate in motorcycle luxury touring. Priced around ৳42,00,000, it delivers 125 bhp smooth power, ultimate comfort seating, advanced electronics, exceptional long-distance capability, and touring luxury soul for serious long-distance explorers seeking the definitive luxury tourer.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda Goldwing 1800 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>1833cc Liquid-Cooled:</strong> Touring masterpiece</li>
<li class="highlight-item"><strong>125 bhp:</strong> Smooth powerful</li>
<li class="highlight-item"><strong>Ultimate Comfort:</strong> Luxury touring</li>
<li class="highlight-item"><strong>Advanced Technology:</strong> Modern systems</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda Goldwing 1800 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Ultimate Comfort:</strong> Luxury seating</li>
<li class="pro-item"><strong>Advanced Technology:</strong> Modern systems</li>
<li class="pro-item"><strong>Smooth Engine:</strong> 125 bhp refined</li>
<li class="pro-item"><strong>Long-Distance:</strong> Exceptional range</li>
<li class="pro-item"><strong>Build Quality:</strong> Japanese premium</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda Goldwing 1800 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Fuel Economy:</strong> 12-15 km/l premium</li>
<li class="con-item"><strong>Heavy Weight:</strong> 383kg very heavy</li>
<li class="con-item"><strong>Ultra-Premium Price:</strong> ৳42,00,000</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda Goldwing 1800 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳42,00,000 - Ultimate luxury tourer</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>1833cc liquid-cooled horizontal-six</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>125 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>167 nm impressive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>383kg very heavy</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>12-15 km/l premium touring</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda Goldwing 1800 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Comfort:</strong> <span>5.0</span> - Ultimate luxury</li>
<li class="rating-item"><strong>Technology:</strong> <span>4.9</span> - Advanced systems</li>
<li class="rating-item"><strong>Long-Distance:</strong> <span>5.0</span> - Exceptional capability</li>
<li class="rating-item"><strong>Build Quality:</strong> <span>4.9</span> - Japanese premium</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳42,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda Goldwing 1800?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.9/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳42,00,000, the Goldwing 1800 is the ultimate luxury choice for serious long-distance explorers seeking ultimate comfort, advanced technology, smooth 125 bhp power, exceptional long-distance capability, and definitive touring luxury for the ultimate motorcycle adventure.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For serious long-distance tourers</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.9,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা গোল্ডওয়িং 1800 রিভিউ বাংলাদেশ 2024 - ট্যুরিং লাক্সারি লিজেন্ড</h1>
<p class="summary-text">হোন্ডা গোল্ডওয়িং 1800 একটি 1833cc লিকুইড-কুল্ড ট্যুরিং মাস্টারপিস যা মোটরসাইকেল লাক্সারি ট্যুরিং এর চূড়ান্ত প্রতিনিধিত্ব করে। ৳42,00,000 টাকায় মূল্যায়িত, এটি 125 bhp মসৃণ শক্তি এবং চূড়ান্ত আরাম আসন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা গোল্ডওয়িং 1800 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>1833cc লিকুইড-কুল্ড:</strong> ট্যুরিং মাস্টারপিস</li>
<li class="highlight-item"><strong>125 bhp:</strong> মসৃণ শক্তিশালী</li>
<li class="highlight-item"><strong>চূড়ান্ত আরাম:</strong> লাক্সারি ট্যুরিং</li>
<li class="highlight-item"><strong>উন্নত প্রযুক্তি:</strong> আধুনিক সিস্টেম</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা গোল্ডওয়িং 1800 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>চূড়ান্ত আরাম:</strong> লাক্সারি আসন</li>
<li class="pro-item"><strong>উন্নত প্রযুক্তি:</strong> আধুনিক সিস্টেম</li>
<li class="pro-item"><strong>মসৃণ ইঞ্জিন:</strong> 125 bhp পরিমার্জিত</li>
<li class="pro-item"><strong>দীর্ঘ-দূরত্ব:</strong> ব্যতিক্রমী পরিসীমা</li>
<li class="pro-item"><strong>বিল্ড কোয়ালিটি:</strong> জাপানি প্রিমিয়াম</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা গোল্ডওয়িং 1800 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>জ্বালানি অর্থনীতি:</strong> 12-15 km/l প্রিমিয়াম</li>
<li class="con-item"><strong>ভারী ওজন:</strong> 383kg অত্যন্ত ভারী</li>
<li class="con-item"><strong>অতি-প্রিমিয়াম মূল্য:</strong> ৳42,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা গোল্ডওয়িং 1800 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.9/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳42,00,000 টাকায়, গোল্ডওয়িং 1800 গুরুতর দীর্ঘ-দূরত্ব অন্বেষণকারীদের জন্য চূড়ান্ত লাক্সারি পছন্দ যারা চূড়ান্ত আরাম, উন্নত প্রযুক্তি এবং সংজ্ঞায়িত ট্যুরিং লাক্সারি খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- গুরুতর দীর্ঘ-দূরত্ব ট্যুরারদের জন্য</span></p>
</div>
</section>
</article>`

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating translation: %w", err)
	}

	fmt.Printf("   ✅ Created Honda Goldwing 1800\n")
	return nil
}

func (s *HondaGoldwing1800ReviewBatch20) GetName() string {
	return "HondaGoldwing1800ReviewBatch20"
}
