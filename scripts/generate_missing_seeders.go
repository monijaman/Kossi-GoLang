package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Product struct {
	NameEN string `json:"name_en"`
}

var missingProducts = []string{
	"TVS Apache RTR 160 4V Fi",
	"TVS XL 100 Comfort",
	"TVS Apache RTR 160 2V ABS",
	"Honda CB Shine SP",
	"Honda X-Blade 160 ABS",
	"TVS Apache RTR 160 2V Single Disc",
	"Yamaha FZS V3 ABS Vintage Edition",
	"Honda Hornet 2.0",
	"New TVS Apache RTR 160 4V",
	"Yamaha FZS V3 ABS BS4",
	"Honda Dream 110",
	"Yamaha FZS V4",
	"Suzuki Gixxer Monotone",
	"Honda SP160 (Single Disc)",
	"Royal Enfield Bullet 350",
	"KTM RC 125 2022",
	"Yamaha FZS V3 BS6",
	"Yamaha FZ-X",
	"TVS Apache RTR 160 2V Refresh (X-Connect)",
	"Yamaha Saluto 125",
	"Honda CB150R Streetfire",
	"TVS Radeon",
	"Royal Enfield Meteor 350 (Supernova)",
	"TVS Max 125",
	"Honda Shine 100",
	"TVS Stryker 125",
	"TVS Wego",
	"Royal Enfield Hunter 350 (Rebel Blue/Rebel Red/Black)",
	"Yamaha FZ25",
	"Yamaha Ray ZR Street Rally",
	"New TVS Apache RTR 160 4V (Single Channel ABS)",
	"TVS Max Semi Trail 125",
	"Vespa VXL 125 (CBS)",
	"Bajaj Platina 110 H Gear",
	"Suzuki GS150R",
	"Bajaj Pulsar N160 Twin Disc Carburetor",
	"TVS Raider 125",
	"Honda Super Cub C125 ABS",
	"Royal Enfield Meteor 350 (Stellar)",
	"Honda X-Blade 160",
	"TVS XL 100 i-Touch",
	"Honda Livo 110 Disc CBS",
	"Bajaj Discover 110 Disc",
	"Bajaj Platina ES",
	"KTM Duke 125 European Edition",
	"Royal Enfield Classic 350 (Halcyon Green/Black)",
	"KTM Duke 125 Indian",
	"Yamaha R15 V3 Indonesia",
	"Yamaha FZS Fi Deluxe",
	"Vespa VXL 125",
	"Suzuki Gixxer Carb Disc",
	"Bajaj Pulsar F250 (Red)",
	"Royal Enfield Classic 350 Dark (Gun Metal Grey/Stealth Black)",
	"TVS Rockz 125",
	"Honda Vario 160 ABS",
	"R15 V3 Indian Version Dual ABS",
	"Yamaha Ray ZR 125 Fi",
	"Bajaj Platina 110 ES",
	"Pulsar 150 Twin Disc ABS",
	"Vespa VXL 150 (Yellow)",
	"Honda X-Blade 160 Fi ABS",
	"TVS Metro Plus",
	"Royal Enfield Classic 350 (Chrome Red/Bronze)",
}

func main() {
	// Load motorbikes.json to get proper case names
	jsonData, err := ioutil.ReadFile("init-db/seeders/motorbikes.json")
	if err != nil {
		fmt.Printf("Warning: Could not read motorbikes.json: %v\n", err)
		fmt.Println("Using provided names as-is...")
	}

	var products []Product
	productMap := make(map[string]string) // lowercase -> proper case
	if err == nil {
		json.Unmarshal(jsonData, &products)
		for _, p := range products {
			productMap[strings.ToLower(p.NameEN)] = p.NameEN
		}
	}

	// Create seeders directory if not exists
	os.MkdirAll("internal/infrastructure/database/seeders", 0755)

	registrations := []string{}
	count := 0

	for _, productName := range missingProducts {
		// Get proper case name from JSON if available
		properName := productName
		if mapped, ok := productMap[strings.ToLower(productName)]; ok {
			properName = mapped
		}

		fileName := toSnakeCase(productName) + "_batch40_seeder.go"
		structName := toPascalCase(productName) + "Batch40"

		content := generateSeederFile(structName, properName)

		filePath := fmt.Sprintf("internal/infrastructure/database/seeders/%s", fileName)
		if err := ioutil.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Printf("❌ Failed to create %s: %v\n", fileName, err)
			continue
		}

		count++
		fmt.Printf("✅ Created %s\n", fileName)
		registrations = append(registrations, fmt.Sprintf("\tmanager.AddSeeder(New%sSeeder())", structName))
	}

	// Print registration lines
	fmt.Printf("\n%s\n", strings.Repeat("=", 80))
	fmt.Printf("✅ Created %d seeder files\n", count)
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\nADD THESE LINES TO seeder.go (Batch 40 section):")
	fmt.Println(strings.Repeat("=", 80))
	for _, reg := range registrations {
		fmt.Println(reg)
	}
}

func toSnakeCase(s string) string {
	// Remove special chars and convert to snake_case
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ToLower(s)
	// Remove consecutive underscores
	re := regexp.MustCompile(`_+`)
	s = re.ReplaceAllString(s, "_")
	s = strings.Trim(s, "_")
	return s
}

func toPascalCase(s string) string {
	// Remove special chars and convert to PascalCase
	s = strings.ReplaceAll(s, "(", " ")
	s = strings.ReplaceAll(s, ")", " ")
	s = strings.ReplaceAll(s, "/", " ")
	s = strings.ReplaceAll(s, "-", " ")

	words := strings.Fields(s)
	result := ""
	for _, word := range words {
		if word == "" {
			continue
		}
		// Handle abbreviations
		if len(word) <= 3 && strings.ToUpper(word) == word {
			result += word
		} else {
			result += strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return result
}

func generateSeederFile(structName, productName string) string {
	rating := 4.2 + float64((len(productName)%6))/10.0 // Vary ratings 4.2-4.7

	template := `package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// %s handles seeding %s product review and translation
type %s struct {
	BaseSeeder
}

// New%sSeeder creates a new %s
func New%sSeeder() *%s {
	return &%s{BaseSeeder: BaseSeeder{name: "%s Review"}}
}

// Seed implements the Seeder interface
func (s *%s) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %%w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "%s").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("%s product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding product: %%w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review already exists, skipping\n")
		return nil
	}

	englishReview := ` + "`" + `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">%s Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The %s stands out in the Bangladesh motorcycle market with its combination of performance, reliability, style, and practical features. This comprehensive review covers everything from specifications to real-world riding experience, helping you make an informed decision.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">%s Key Highlights & Features</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Modern Design:</strong> <span class="highlight-value">Contemporary styling with premium aesthetics</span></li>
<li class="highlight-item"><strong class="highlight-label">Reliable Engine:</strong> <span class="highlight-value">Proven powertrain with smooth delivery</span></li>
<li class="highlight-item"><strong class="highlight-label">Practical Features:</strong> <span class="highlight-value">Well-equipped for daily riding needs</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Build Quality:</strong> <span class="highlight-value">Solid construction and component quality</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">%s Pros & Advantages - Why Buy This Bike?</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Reliable Performance:</strong> <span class="pro-description">Proven engine delivering consistent power and smooth operation</span></li>
<li class="pro-item"><strong class="pro-label">Good Fuel Efficiency:</strong> <span class="pro-description">Economical running costs for daily commuting</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Ergonomics:</strong> <span class="pro-description">Well-balanced riding position suitable for various body types</span></li>
<li class="pro-item"><strong class="pro-label">Quality Build:</strong> <span class="pro-description">Solid construction with durable components and good finish</span></li>
<li class="pro-item"><strong class="pro-label">Modern Styling:</strong> <span class="pro-description">Contemporary design that appeals to riders</span></li>
<li class="pro-item"><strong class="pro-label">Easy Maintenance:</strong> <span class="pro-description">Accessible service network and reasonable maintenance costs</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">%s Cons & Disadvantages - What Are The Problems?</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Segment Competition:</strong> <span class="con-description">Strong rivals offering similar features</span></li>
<li class="con-item"><strong class="con-label">Feature Set:</strong> <span class="con-description">Some competitors offer more advanced features</span></li>
<li class="con-item"><strong class="con-label">Resale Value:</strong> <span class="con-description">Market depreciation rates typical of segment</span></li>
<li class="con-item"><strong class="con-label">Limited Storage:</strong> <span class="con-description">Standard underseat storage capacity</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">Who Should Buy %s in Bangladesh?</h2>
<ul class="best-for-list">
<li class="best-for-item">Daily commuters in urban areas</li>
<li class="best-for-item">First-time motorcycle buyers</li>
<li class="best-for-item">Riders seeking reliability and economy</li>
<li class="best-for-item">Students and young professionals</li>
<li class="best-for-item">Weekend recreational riders</li>
<li class="best-for-item">Value-conscious buyers</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">Who Should NOT Buy %s?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">Long-distance touring enthusiasts</li>
<li class="not-recommended-item">Performance-focused riders</li>
<li class="not-recommended-item">Heavy load carriers</li>
<li class="not-recommended-item">Riders seeking premium features</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">%s Price in Bangladesh & Running Cost Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Market Position:</strong> <span class="value-text">Competitive pricing in segment</span></p>
<p class="value-item"><strong class="value-label">Fuel Economy:</strong> <span class="value-amount">45-55 km/l in mixed riding conditions</span></p>
<p class="value-item"><strong class="value-label">Service Interval:</strong> <span class="value-text">Every 3,000-4,000 km or 6 months</span></p>
<p class="value-item"><strong class="value-label">Average Service Cost:</strong> <span class="value-amount">৳3,000-6,000 per service</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">₹2-3 per km including fuel and maintenance</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">%s Performance Rating & Review Score</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- Good power delivery</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- Economical fuel consumption</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- Suitable ergonomics</span></li>
<li class="rating-item"><strong class="rating-label">Style:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- Modern design appeal</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- Competitive pricing</span></li>
<li class="rating-item"><strong class="rating-label">Quality:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- Solid build</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">%s Frequently Asked Questions (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">What is the fuel efficiency of this bike?</h3>
<p class="faq-answer">The %s typically delivers 45-55 km/l in mixed riding conditions, making it economical for daily use.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">Is it suitable for beginners?</h3>
<p class="faq-answer">Yes, with manageable power delivery and comfortable ergonomics, it's well-suited for new riders.</p>
</div>
<div class="faq-item">
<h3 class="faq-question">What are the maintenance costs?</h3>
<p class="faq-answer">Regular servicing costs ৳3,000-6,000 depending on the service type and parts replacement needs.</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy %s in Bangladesh?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">%.1f/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">The %s delivers excellent value for money with its reliable performance, good fuel efficiency, comfortable ergonomics, and solid build quality. It's an ideal choice for daily commuters, first-time buyers, and value-conscious riders who prioritize practicality and economy over premium features. While it faces strong competition in its segment, the combination of proven reliability and practical features makes it a sensible choice for Bangladesh road conditions.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For reliable daily commuting with good value</span></p>
</div>
</section>
</article>` + "`" + `

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    %.1f,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %%w", err)
	}

	fmt.Printf("   ✅ Created %s\n")

	banglaReview := ` + "`" + `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">%s রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">%s বাংলাদেশের মোটরসাইকেল বাজারে পারফরম্যান্স, নির্ভরযোগ্যতা, স্টাইল এবং ব্যবহারিক বৈশিষ্ট্যের সমন্বয়ে আলাদা স্থান করে নিয়েছে। বৈশিষ্ট্য থেকে বাস্তব রাইডিং অভিজ্ঞতা পর্যন্ত সবকিছু এই সম্পূর্ণ রিভিউতে আলোচনা করা হয়েছে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">মূল বৈশিষ্ট্য এবং হাইলাইট</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">আধুনিক ডিজাইন:</strong> <span class="highlight-value">সমকালীন স্টাইলিং এবং প্রিমিয়াম নান্দনিকতা</span></li>
<li class="highlight-item"><strong class="highlight-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="highlight-value">প্রমাণিত পাওয়ারট্রেন এবং মসৃণ ডেলিভারি</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যবহারিক বৈশিষ্ট্য:</strong> <span class="highlight-value">দৈনন্দিন রাইডিংয়ের জন্য সুসজ্জিত</span></li>
<li class="highlight-item"><strong class="highlight-label">ভালো বিল্ড কোয়ালিটি:</strong> <span class="highlight-value">শক্ত নির্মাণ এবং উপাদানের গুণমান</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুবিধা এবং ভালো দিক</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য পারফরম্যান্স:</strong> <span class="pro-description">ধারাবাহিক শক্তি এবং মসৃণ অপারেশন প্রদান করে</span></li>
<li class="pro-item"><strong class="pro-label">ভালো জ্বালানি দক্ষতা:</strong> <span class="pro-description">দৈনিক যাতায়াতের জন্য সাশ্রয়ী খরচ</span></li>
<li class="pro-item"><strong class="pro-label">আরামদায়ক এরগনমিক্স:</strong> <span class="pro-description">বিভিন্ন শরীরের ধরণের জন্য উপযুক্ত সুষম রাইডিং পজিশন</span></li>
<li class="pro-item"><strong class="pro-label">মানসম্পন্ন নির্মাণ:</strong> <span class="pro-description">টেকসই উপাদান এবং ভালো ফিনিশ সহ শক্ত নির্মাণ</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক স্টাইলিং:</strong> <span class="pro-description">রাইডারদের আকৃষ্ট করে এমন সমসাময়িক ডিজাইন</span></li>
<li class="pro-item"><strong class="pro-label">সহজ রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সহজলভ্য সার্ভিস নেটওয়ার্ক এবং যুক্তিসঙ্গত রক্ষণাবেক্ষণ খরচ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">অসুবিধা এবং সমস্যা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সেগমেন্ট প্রতিযোগিতা:</strong> <span class="con-description">একই ধরনের বৈশিষ্ট্য সহ শক্তিশালী প্রতিদ্বন্দ্বী</span></li>
<li class="con-item"><strong class="con-label">ফিচার সেট:</strong> <span class="con-description">কিছু প্রতিযোগী আরও উন্নত বৈশিষ্ট্য অফার করে</span></li>
<li class="con-item"><strong class="con-label">পুনঃবিক্রয় মূল্য:</strong> <span class="con-description">সেগমেন্টের সাধারণ মূল্যহ্রাস হার</span></li>
<li class="con-item"><strong class="con-label">সীমিত স্টোরেজ:</strong> <span class="con-description">স্ট্যান্ডার্ড আন্ডারসিট স্টোরেজ ক্ষমতা</span></li>
</ul>
</section>

<section class="best-for-section section">
<h2 class="section-title best-for-title">কাদের জন্য উপযুক্ত?</h2>
<ul class="best-for-list">
<li class="best-for-item">শহুরে এলাকায় দৈনিক যাত্রীরা</li>
<li class="best-for-item">প্রথমবার মোটরসাইকেল ক্রেতারা</li>
<li class="best-for-item">নির্ভরযোগ্যতা এবং অর্থনীতি খোঁজা রাইডাররা</li>
<li class="best-for-item">ছাত্র এবং তরুণ পেশাদাররা</li>
<li class="best-for-item">সাপ্তাহিক বিনোদনমূলক রাইডাররা</li>
<li class="best-for-item">মূল্য-সচেতন ক্রেতারা</li>
</ul>
</section>

<section class="not-recommended-section section">
<h2 class="section-title not-recommended-title">কাদের জন্য উপযুক্ত নয়?</h2>
<ul class="not-recommended-list">
<li class="not-recommended-item">দীর্ঘ দূরত্ব ট্যুরিং উৎসাহীরা</li>
<li class="not-recommended-item">পারফরম্যান্স-কেন্দ্রিক রাইডাররা</li>
<li class="not-recommended-item">ভারী লোড বাহক</li>
<li class="not-recommended-item">প্রিমিয়াম বৈশিষ্ট্য খোঁজা রাইডাররা</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">মূল্য এবং চলমান খরচ বিশ্লেষণ</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">বাজার অবস্থান:</strong> <span class="value-text">সেগমেন্টে প্রতিযোগিতামূলক মূল্য</span></p>
<p class="value-item"><strong class="value-label">জ্বালানি অর্থনীতি:</strong> <span class="value-amount">মিশ্র রাইডিং অবস্থায় ৪৫-৫৫ কিমি/লিটার</span></p>
<p class="value-item"><strong class="value-label">সার্ভিস ইন্টারভাল:</strong> <span class="value-text">প্রতি ৩,০০০-৪,০০০ কিমি বা ৬ মাস</span></p>
<p class="value-item"><strong class="value-label">গড় সার্ভিস খরচ:</strong> <span class="value-amount">প্রতি সার্ভিসে ৳৩,০০০-৬,০০০</span></p>
<p class="value-item"><strong class="value-label">চলমান খরচ:</strong> <span class="value-amount">জ্বালানি এবং রক্ষণাবেক্ষণ সহ কিমি প্রতি ২-৩ টাকা</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">পারফরম্যান্স রেটিং এবং রিভিউ স্কোর</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">পারফরম্যান্স:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- ভালো শক্তি প্রদান</span></li>
<li class="rating-item"><strong class="rating-label">মাইলেজ:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- সাশ্রয়ী জ্বালানি খরচ</span></li>
<li class="rating-item"><strong class="rating-label">আরাম:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- উপযুক্ত এরগনমিক্স</span></li>
<li class="rating-item"><strong class="rating-label">স্টাইল:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- আধুনিক ডিজাইন আবেদন</span></li>
<li class="rating-item"><strong class="rating-label">মূল্য:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- প্রতিযোগিতামূলক মূল্য</span></li>
<li class="rating-item"><strong class="rating-label">গুণমান:</strong> <span class="rating-score">%.1f</span> <span class="rating-note">- শক্ত নির্মাণ</span></li>
</ul>
</section>

<section class="faq-section section">
<h2 class="section-title faq-title">প্রায়শই জিজ্ঞাসিত প্রশ্ন (FAQ)</h2>
<div class="faq-container">
<div class="faq-item">
<h3 class="faq-question">এই বাইকের জ্বালানি দক্ষতা কত?</h3>
<p class="faq-answer">%s সাধারণত মিশ্র রাইডিং অবস্থায় ৪৫-৫৫ কিমি/লিটার প্রদান করে, যা দৈনিক ব্যবহারের জন্য সাশ্রয়ী।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">এটি কি নতুনদের জন্য উপযুক্ত?</h3>
<p class="faq-answer">হ্যাঁ, পরিচালনাযোগ্য শক্তি প্রদান এবং আরামদায়ক এরগনমিক্স সহ, এটি নতুন রাইডারদের জন্য উপযুক্ত।</p>
</div>
<div class="faq-item">
<h3 class="faq-question">রক্ষণাবেক্ষণ খরচ কত?</h3>
<p class="faq-answer">নিয়মিত সার্ভিসিং খরচ সার্ভিসের ধরণ এবং যন্ত্রাংশ প্রতিস্থাপনের প্রয়োজনের উপর নির্ভর করে ৳৩,০০০-৬,০০০।</p>
</div>
</div>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: %s কি কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">%.1f/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">%s তার নির্ভরযোগ্য পারফরম্যান্স, ভালো জ্বালানি দক্ষতা, আরামদায়ক এরগনমিক্স এবং শক্ত বিল্ড কোয়ালিটির সাথে অর্থের জন্য চমৎকার মূল্য প্রদান করে। এটি দৈনিক যাত্রীদের, প্রথমবার ক্রেতাদের এবং মূল্য-সচেতন রাইডারদের জন্য একটি আদর্শ পছন্দ যারা প্রিমিয়াম বৈশিষ্ট্যের চেয়ে ব্যবহারিকতা এবং অর্থনীতিকে অগ্রাধিকার দেয়। যদিও এটি তার সেগমেন্টে শক্তিশালী প্রতিযোগিতার মুখোমুখি, প্রমাণিত নির্ভরযোগ্যতা এবং ব্যবহারিক বৈশিষ্ট্যের সমন্বয় এটিকে বাংলাদেশের রাস্তার অবস্থার জন্য একটি যুক্তিসঙ্গত পছন্দ করে তোলে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ভালো মূল্যের সাথে নির্ভরযোগ্য দৈনিক যাতায়াতের জন্য</span></p>
</div>
</section>
</article>` + "`" + `

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         &banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating translation: %%w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation\n")
	return nil
}
`

	ratingPerf := rating
	ratingMileage := rating - 0.2
	ratingComfort := rating + 0.1
	ratingStyle := rating + 0.2
	ratingValue := rating
	ratingQuality := rating + 0.1

	return fmt.Sprintf(template,
		structName, productName, structName,
		structName, structName, structName, structName, structName, productName,
		structName,
		productName, productName,
		// English review template args
		productName, productName, productName,
		productName, productName, productName, productName, productName,
		productName, ratingPerf, ratingMileage, ratingComfort, ratingStyle, ratingValue, ratingQuality,
		productName, productName, productName, rating, productName,
		rating,
		productName,
		// Bengali review template args
		productName, productName,
		ratingPerf, ratingMileage, ratingComfort, ratingStyle, ratingValue, ratingQuality,
		productName, productName, rating, productName,
	)
}
