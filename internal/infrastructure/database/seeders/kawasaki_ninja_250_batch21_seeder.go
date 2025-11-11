package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type KawasakiNinja250Batch21 struct {
	BaseSeeder
}

func NewKawasakiNinja250Batch21Seeder() *KawasakiNinja250Batch21 {
	return &KawasakiNinja250Batch21{BaseSeeder: BaseSeeder{name: "Kawasaki Ninja 250 Batch21 Review"}}
}

func (s *KawasakiNinja250Batch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Kawasaki Ninja 250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("kawasaki ninja 250 product not found")
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
<h1 class="review-main-title">Kawasaki Ninja 250 Review Bangladesh 2024 - Sporty Performance Beginner Bike</h1>
<p class="summary-text">The Kawasaki Ninja 250 is a 250cc liquid-cooled sporty performance beginner motorcycle representing accessible sportbike performance. Priced around ৳4,45,000, it delivers 26 bhp responsive power, aggressive sportbike styling, responsive handling, practical performance, and enthusiast appeal for confident beginner riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Kawasaki Ninja 250 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>250cc Liquid-Cooled:</strong> Beginner sportbike</li>
<li class="highlight-item"><strong>26 bhp Power:</strong> Responsive performance</li>
<li class="highlight-item"><strong>Ninja Styling:</strong> Sportbike aesthetic</li>
<li class="highlight-item"><strong>38+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Kawasaki Ninja 250 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Responsive 250cc:</strong> Sporty acceleration</li>
<li class="pro-item"><strong>Sportbike Styling:</strong> Aggressive Ninja design</li>
<li class="pro-item"><strong>Responsive Handling:</strong> Sport-focused dynamics</li>
<li class="pro-item"><strong>Manageable Power:</strong> Beginner-friendly 26 bhp</li>
<li class="pro-item"><strong>Enthusiast Appeal:</strong> Kawasaki heritage</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Kawasaki Ninja 250 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Pricing:</strong> ৳4,45,000 imported bike</li>
<li class="con-item"><strong>Sport Seating:</strong> Limited comfort</li>
<li class="con-item"><strong>Fuel Cost:</strong> Premium fuel requirement</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Kawasaki Ninja 250 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,45,000 - Premium beginner sportbike</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>250cc liquid-cooled parallel-twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>26 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>23 nm smooth</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>167kg sport-focused</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>38+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Kawasaki Ninja 250 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 26 bhp responsive</li>
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Sportbike aesthetic</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Sport responsive</li>
<li class="rating-item"><strong>Heritage:</strong> <span>4.8</span> - Kawasaki iconic</li>
<li class="rating-item"><strong>Value:</strong> <span>4.5</span> - ৳4,45,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Kawasaki Ninja 250?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,45,000, the Ninja 250 is perfect for enthusiasts seeking accessible sportbike performance, responsive 26 bhp power, aggressive Ninja styling, sport handling, and Kawasaki heritage for confident beginner riders.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For beginner sportbike enthusiasts</span></p>
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
<h1 class="review-main-title">কাওয়াসাকি নিনজা 250 রিভিউ বাংলাদেশ 2024 - খেলাধুলাপূর্ণ কর্মক্ষমতা শিক্ষানবিস বাইক</h1>
<p class="summary-text">কাওয়াসাকি নিনজা 250 একটি 250cc লিকুইড-কুল্ড খেলাধুলাপূর্ণ কর্মক্ষমতা শিক্ষানবিস মোটরসাইকেল যা সহজলভ্য স্পোর্টবাইক কর্মক্ষমতা প্রতিনিধিত্ব করে। ৳4,45,000 টাকায় মূল্যায়িত, এটি 26 bhp প্রতিক্রিয়াশীল শক্তি এবং আক্রমণাত্মক স্পোর্টবাইক স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কাওয়াসাকি নিনজা 250 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>250cc লিকুইড-কুল্ড:</strong> শিক্ষানবিস স্পোর্টবাইক</li>
<li class="highlight-item"><strong>26 bhp শক্তি:</strong> প্রতিক্রিয়াশীল কর্মক্ষমতা</li>
<li class="highlight-item"><strong>নিনজা স্টাইলিং:</strong> স্পোর্টবাইক নান্দনিক</li>
<li class="highlight-item"><strong>38+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কাওয়াসাকি নিনজা 250 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রতিক্রিয়াশীল 250cc:</strong> খেলাধুলাপূর্ণ ত্বরণ</li>
<li class="pro-item"><strong>স্পোর্টবাইক স্টাইলিং:</strong> আক্রমণাত্মক নিনজা ডিজাইন</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল হ্যান্ডলিং:</strong> খেলাধুলা-ফোকাসড গতিশীলতা</li>
<li class="pro-item"><strong>পরিচালনাযোগ্য শক্তি:</strong> শিক্ষানবিস-বান্ধব 26 bhp</li>
<li class="pro-item"><strong>উত্সাহী আবেদন:</strong> কাওয়াসাকি ঐতিহ্য</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কাওয়াসাকি নিনজা 250 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳4,45,000 আমদানিকৃত বাইক</li>
<li class="con-item"><strong>খেলাধুলা আসন:</strong> সীমিত আরাম</li>
<li class="con-item"><strong>জ্বালানি খরচ:</strong> প্রিমিয়াম জ্বালানি প্রয়োজনীয়তা</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কাওয়াসাকি নিনজা 250 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,45,000 টাকায়, নিনজা 250 উত্সাহীদের জন্য নিখুঁত যারা সহজলভ্য স্পোর্টবাইক কর্মক্ষমতা, প্রতিক্রিয়াশীল 26 bhp শক্তি এবং আক্রমণাত্মক নিনজা স্টাইলিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শিক্ষানবিস স্পোর্টবাইক উত্সাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Kawasaki Ninja 250\n")
	return nil
}

func (s *KawasakiNinja250Batch21) GetName() string {
	return "KawasakiNinja250Batch21"
}
