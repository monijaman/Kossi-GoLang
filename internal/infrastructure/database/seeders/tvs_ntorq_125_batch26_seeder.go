package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type TVSNtorq125Batch26 struct {
	BaseSeeder
}

func NewTVSNtorq125Batch26Seeder() *TVSNtorq125Batch26 {
	return &TVSNtorq125Batch26{BaseSeeder: BaseSeeder{name: "TVS Ntorq 125 Batch26 Review"}}
}

func (s *TVSNtorq125Batch26) Seed(db *gorm.DB) error {
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
<h1 class="review-main-title">TVS Ntorq 125 Review Bangladesh 2024 - Smart Urban Scooter</h1>
<p class="summary-text">The TVS Ntorq 125 is a feature-rich smart urban scooter combining technology and convenience. Priced around ৳3,15,000, it delivers 9.4 bhp power, smartphone connectivity, spacious under-seat storage, exceptional 55+ km/l efficiency, and modern urban commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Ntorq 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Scooter:</strong> Smart urban featured</li>
<li class="highlight-item"><strong>9.4 bhp Power:</strong> Smooth responsive</li>
<li class="highlight-item"><strong>Smartphone Connect:</strong> Tech integrated</li>
<li class="highlight-item"><strong>55+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Ntorq 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Smart Features:</strong> Tech connected modern</li>
<li class="pro-item"><strong>Spacious Storage:</strong> Convenient under-seat</li>
<li class="pro-item"><strong>Smooth Riding:</strong> 9.4 bhp responsive</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 55+ km/l economical</li>
<li class="pro-item"><strong>Modern Design:</strong> Contemporary styling</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Ntorq 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Tech Dependencies:</strong> App connectivity required</li>
<li class="con-item"><strong>Higher Price:</strong> Feature premium</li>
<li class="con-item"><strong>Limited Power:</strong> 9.4 bhp modest</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Ntorq 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,15,000 - Smart scooter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>9.4 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>10 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>133kg compact</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>55+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Ntorq 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Modern contemporary</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Urban spacious</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 9.4 bhp smooth</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 55+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - TVS reliable</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Ntorq 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,15,000, the TVS Ntorq 125 is excellent for tech-savvy urban commuters seeking smart features, convenient storage, smooth 9.4 bhp performance, and exceptional 55+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For tech-savvy urbanites</span></p>
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
<h1 class="review-main-title">TVS Ntorq 125 রিভিউ বাংলাদেশ 2024 - স্মার্ট শহুরে স্কুটার</h1>
<p class="summary-text">TVS Ntorq 125 একটি বৈশিষ্ট্য সমৃদ্ধ স্মার্ট শহুরে স্কুটার যা প্রযুক্তি এবং সুবিধা একত্রিত করে। ৳3,15,000 টাকায় মূল্যায়িত, এটি 9.4 bhp শক্তি এবং স্মার্টফোন সংযোগ বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Ntorq 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc স্কুটার:</strong> স্মার্ট শহুরে বৈশিষ্ট্য</li>
<li class="highlight-item"><strong>9.4 bhp শক্তি:</strong> মসৃণ প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>স্মার্টফোন সংযোগ:</strong> প্রযুক্তি একীভূত</li>
<li class="highlight-item"><strong>55+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Ntorq 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>স্মার্ট বৈশিষ্ট্য:</strong> প্রযুক্তি সংযুক্ত আধুনিক</li>
<li class="pro-item"><strong>বিস্তৃত সঞ্চয়স্থান:</strong> সুবিধাজনক আন্ডার-সিট</li>
<li class="pro-item"><strong>মসৃণ চালনা:</strong> 9.4 bhp প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 55+ km/l সাশ্রয়ী</li>
<li class="pro-item"><strong>আধুনিক ডিজাইন:</strong> সমসাময়িক স্টাইলিং</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Ntorq 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রযুক্তি নির্ভরতা:</strong> অ্যাপ সংযোগ প্রয়োজন</li>
<li class="con-item"><strong>উচ্চতর মূল্য:</strong> বৈশিষ্ট্য প্রিমিয়াম</li>
<li class="con-item"><strong>সীমিত শক্তি:</strong> 9.4 bhp মধ্যম</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: TVS Ntorq 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,15,000 টাকায়, TVS Ntorq 125 প্রযুক্তি-সচেতন শহুরে যাতায়াতকারীদের জন্য চমৎকার যারা স্মার্ট বৈশিষ্ট্য এবং চমৎকার 55+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রযুক্তি-সচেতন শহুরে প্রবাসীদের জন্য</span></p>
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

func (s *TVSNtorq125Batch26) GetName() string {
	return "TVSNtorq125Batch26"
}
