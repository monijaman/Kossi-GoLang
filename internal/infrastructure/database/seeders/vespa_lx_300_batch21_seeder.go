package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type VespaLX300Batch21 struct {
	BaseSeeder
}

func NewVespaLX300Batch21Seeder() *VespaLX300Batch21 {
	return &VespaLX300Batch21{BaseSeeder: BaseSeeder{name: "Vespa LX 300 Batch21 Review"}}
}

func (s *VespaLX300Batch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Vespa LX 300").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("vespa lx 300 product not found")
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
<h1 class="review-main-title">Vespa LX 300 Review Bangladesh 2024 - Premium Italian Lifestyle Scooter</h1>
<p class="summary-text">The Vespa LX 300 is a 300cc liquid-cooled premium Italian lifestyle scooter representing timeless design and sophisticated urban mobility. Priced around ৳12,00,000, it delivers smooth power, iconic Italian styling, premium comfort, excellent storage, retro-modern design, and refined riding experience for sophisticated urbanites.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Vespa LX 300 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>300cc Liquid-Cooled:</strong> Premium scooter</li>
<li class="highlight-item"><strong>22 bhp Smooth:</strong> Refined power</li>
<li class="highlight-item"><strong>Iconic Italian Design:</strong> Timeless scooter</li>
<li class="highlight-item"><strong>32+ km/l Efficiency:</strong> Practical economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Vespa LX 300 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Iconic Design:</strong> Legendary Italian heritage</li>
<li class="pro-item"><strong>Smooth 300cc:</strong> Refined power delivery</li>
<li class="pro-item"><strong>Premium Comfort:</strong> Sophisticated riding</li>
<li class="pro-item"><strong>Excellent Storage:</strong> Spacious compartment</li>
<li class="pro-item"><strong>Refined Experience:</strong> Lifestyle scooter</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Vespa LX 300 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Pricing:</strong> ৳12,00,000 luxury scooter</li>
<li class="con-item"><strong>Moderate Weight:</strong> 210kg premium build</li>
<li class="con-item"><strong>Fuel Economy:</strong> 32+ km/l average for class</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Vespa LX 300 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳12,00,000 - Premium Italian luxury</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>300cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>22 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>25 nm refined</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>210kg premium</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>32+ km/l practical</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Vespa LX 300 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Iconic Italian</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Premium refined</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - Smooth 22 bhp</li>
<li class="rating-item"><strong>Heritage:</strong> <span>4.9</span> - Vespa legend</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳12,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Vespa LX 300?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳12,00,000, the LX 300 is the premium choice for sophisticated urbanites seeking iconic Italian heritage, refined smooth performance, premium comfort, excellent storage, and timeless Vespa lifestyle.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium lifestyle urbanites</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.8,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ভেসপা LX 300 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম ইতালিয়ান জীবনযাত্রা স্কুটার</h1>
<p class="summary-text">ভেসপা LX 300 একটি 300cc লিকুইড-কুল্ড প্রিমিয়াম ইতালিয়ান জীবনযাত্রা স্কুটার যা কালজয়ী ডিজাইন এবং পরিশীলিত শহুরে গতিশীলতা প্রতিনিধিত্ব করে। ৳12,00,000 টাকায় মূল্যায়িত, এটি মসৃণ শক্তি এবং আইকনিক ইতালিয়ান স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ভেসপা LX 300 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>300cc লিকুইড-কুল্ড:</strong> প্রিমিয়াম স্কুটার</li>
<li class="highlight-item"><strong>22 bhp মসৃণ:</strong> পরিমার্জিত শক্তি</li>
<li class="highlight-item"><strong>আইকনিক ইতালিয়ান ডিজাইন:</strong> কালজয়ী স্কুটার</li>
<li class="highlight-item"><strong>32+ km/l দক্ষতা:</strong> ব্যবহারিক অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ভেসপা LX 300 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আইকনিক ডিজাইন:</strong> কিংবদন্তী ইতালিয়ান ঐতিহ্য</li>
<li class="pro-item"><strong>মসৃণ 300cc:</strong> পরিমার্জিত শক্তি সরবরাহ</li>
<li class="pro-item"><strong>প্রিমিয়াম আরাম:</strong> পরিশীলিত যাত্রা</li>
<li class="pro-item"><strong>চমৎকার স্টোরেজ:</strong> প্রশস্ত বগি</li>
<li class="pro-item"><strong>পরিমার্জিত অভিজ্ঞতা:</strong> জীবনযাত্রা স্কুটার</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ভেসপা LX 300 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳12,00,000 বিলাসিতা স্কুটার</li>
<li class="con-item"><strong>মধ্যম ওজন:</strong> 210kg প্রিমিয়াম বিল্ড</li>
<li class="con-item"><strong>জ্বালানি অর্থনীতি:</strong> 32+ km/l ক্লাসের জন্য গড়</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ভেসপা LX 300 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳12,00,000 টাকায়, LX 300 পরিশীলিত শহুরেদের জন্য প্রিমিয়াম পছন্দ যারা আইকনিক ইতালিয়ান ঐতিহ্য, পরিমার্জিত মসৃণ কর্মক্ষমতা এবং প্রিমিয়াম আরাম চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম জীবনযাত্রা শহুরেদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Vespa LX 300\n")
	return nil
}

func (s *VespaLX300Batch21) GetName() string {
	return "VespaLX300Batch21"
}
