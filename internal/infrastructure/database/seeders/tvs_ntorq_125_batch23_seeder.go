package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type TVSNtorq125Batch23 struct {
	BaseSeeder
}

func NewTVSNtorq125Batch23Seeder() *TVSNtorq125Batch23 {
	return &TVSNtorq125Batch23{BaseSeeder: BaseSeeder{name: "TVS Ntorq 125 Batch23 Review"}}
}

func (s *TVSNtorq125Batch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Ntorq 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs ntorq 125 product not found")
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
<h1 class="review-main-title">TVS Ntorq 125 Review Bangladesh 2024 - India's Tech-Forward Youth Scooter</h1>
<p class="summary-text">The TVS Ntorq 125 is India's tech-forward 125cc youth scooter representing modern connected features. Priced around ৳4,25,000, it delivers 9.9 bhp responsive power, modern aggressive styling, smartphone connectivity, good 52+ km/l efficiency, and excellent value for tech-savvy youth seeking connected modern scooter.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Ntorq 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Scooter:</strong> Youth tech-forward</li>
<li class="highlight-item"><strong>9.9 bhp Power:</strong> Responsive agile</li>
<li class="highlight-item"><strong>Smart Connect:</strong> Modern features</li>
<li class="highlight-item"><strong>52+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Ntorq 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Modern Tech Features:</strong> Smartphone connected</li>
<li class="pro-item"><strong>Aggressive Styling:</strong> Youth appeal</li>
<li class="pro-item"><strong>Responsive Power:</strong> 9.9 bhp nimble</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 52+ km/l practical</li>
<li class="pro-item"><strong>TVS Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Ntorq 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳4,25,000 upscale</li>
<li class="con-item"><strong>Tech Dependent:</strong> Feature complexity</li>
<li class="con-item"><strong>Smaller Fuel Tank:</strong> Limited range</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Ntorq 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,25,000 - Tech premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc scooter single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>9.9 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>10 nm agile</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>132kg modern</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>52+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Ntorq 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Technology:</strong> <span>4.8</span> - Smart features</li>
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Modern aggressive</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 9.9 bhp responsive</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 52+ km/l good</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳4,25,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Ntorq 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,25,000, the TVS Ntorq 125 is perfect for tech-savvy youth seeking modern connected scooter, smartphone features, responsive 9.9 bhp power, good 52+ km/l efficiency, and contemporary youth appeal.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For tech-forward youth</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.7,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">টিভিএস এনটর্ক 125 রিভিউ বাংলাদেশ 2024 - ভারতের প্রযুক্তি-সামনের যুব স্কুটার</h1>
<p class="summary-text">টিভিএস এনটর্ক 125 ভারতের প্রযুক্তি-সামনের 125cc যুব স্কুটার যা আধুনিক সংযুক্ত বৈশিষ্ট্য প্রতিনিধিত্ব করে। ৳4,25,000 টাকায় মূল্যায়িত, এটি 9.9 bhp প্রতিক্রিয়াশীল শক্তি এবং আধুনিক আক্রমণাত্মক স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস এনটর্ক 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc স্কুটার:</strong> যুব প্রযুক্তি-সামনের</li>
<li class="highlight-item"><strong>9.9 bhp শক্তি:</strong> প্রতিক্রিয়াশীল চটপটে</li>
<li class="highlight-item"><strong>স্মার্ট সংযোগ:</strong> আধুনিক বৈশিষ্ট্য</li>
<li class="highlight-item"><strong>52+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস এনটর্ক 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আধুনিক প্রযুক্তি বৈশিষ্ট্য:</strong> স্মার্টফোন সংযুক্ত</li>
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> যুব আবেদন</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 9.9 bhp চটপটে</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 52+ km/l ব্যবহারিক</li>
<li class="pro-item"><strong>টিভিএস নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস এনটর্ক 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳4,25,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>প্রযুক্তি নির্ভরশীল:</strong> বৈশিষ্ট্য জটিলতা</li>
<li class="con-item"><strong>ছোটো জ্বালানি ট্যাঙ্ক:</strong> সীমিত পরিসর</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: টিভিএস এনটর্ক 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,25,000 টাকায়, টিভিএস এনটর্ক 125 প্রযুক্তি-সচেতন যুবকদের জন্য নিখুঁত যারা আধুনিক সংযুক্ত স্কুটার, স্মার্টফোন বৈশিষ্ট্য এবং প্রতিক্রিয়াশীল 9.9 bhp শক্তি চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রযুক্তি-সামনের যুবকদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created TVS Ntorq 125\n")
	return nil
}

func (s *TVSNtorq125Batch23) GetName() string {
	return "TVSNtorq125Batch23"
}
