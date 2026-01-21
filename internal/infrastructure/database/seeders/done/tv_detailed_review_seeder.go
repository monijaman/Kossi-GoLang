package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"time"

	"gorm.io/gorm"
)

// TVDetailedReviewSeeder handles seeding detailed, SEO-friendly TV product reviews
type TVDetailedReviewSeeder struct {
	BaseSeeder
}

// NewTVDetailedReviewSeeder creates a new TV detailed review seeder
func NewTVDetailedReviewSeeder() *TVDetailedReviewSeeder {
	return &TVDetailedReviewSeeder{
		BaseSeeder: BaseSeeder{name: "TV Detailed Reviews"},
	}
}

// Seed implements the Seeder interface
func (tdrs *TVDetailedReviewSeeder) Seed(db *gorm.DB) error {
	// Get TV category
	var tvCategory models.CategoryModel
	if err := db.Where("id = ?", 124).First(&tvCategory).Error; err != nil {
		return nil
	}

	detailedReviews := []struct {
		ProductName string
		HTMLContent string
		Verdict     string
		SEOKeywords []string
	}{
		{
			ProductName: "Samsung 55\" 4K Smart TV",
			HTMLContent: `<div class="tv-review-content">
<h2>Samsung 55" 4K Smart TV - Complete Review & Detailed Analysis</h2>

<h3>Introduction</h3>
<p>The Samsung 55" 4K Smart TV stands as a benchmark for mid-range premium televisions in Bangladesh. This model combines Samsung's legendary QLED technology with an impressive array of smart features, making it an ideal choice for families and entertainment enthusiasts who demand quality without breaking the bank.</p>

<h3>Display Technology & Picture Quality</h3>
<p>At the heart of this TV lies Samsung's cutting-edge QLED (Quantum Dot LED) panel technology. The 4K UHD resolution (3840 x 2160 pixels) delivers exceptional clarity with four times the pixels of standard Full HD. The vibrant colors produced by quantum dots are particularly impressive, offering 99% DCI-P3 color space coverage.</p>

<ul>
<li><strong>QLED Technology:</strong> Quantum dots produce 40% brighter images with deeper blacks than traditional LED panels</li>
<li><strong>60Hz Native Refresh Rate:</strong> Smooth motion handling ideal for movies and standard content</li>
<li><strong>120Hz Motion Handling:</strong> AI upscaling technology smooths motion in live sports and fast-paced content</li>
<li><strong>HDR Support:</strong> Full support for HDR10, HDR10+, and HLG formats</li>
<li><strong>400+ Nits Peak Brightness:</strong> Excellent brightness for well-lit rooms</li>
</ul>

<h3>Smart Features & Operating System</h3>
<p>The Tizen OS provides an intuitive, responsive interface with excellent app ecosystem support. Netflix, Amazon Prime Video, YouTube, and Disney+ are pre-installed, and the voice control compatibility with Bixby and Alexa adds convenience to daily use.</p>

<ul>
<li>Quick app loading with SSD-optimized performance</li>
<li>Seamless casting from smartphones and laptops</li>
<li>Voice assistant integration for hands-free operation</li>
<li>Regular OS updates with new features and security patches</li>
<li>Intuitive menu navigation suitable for all age groups</li>
</ul>

<h3>Gaming Performance</h3>
<p>Gamers will appreciate the multiple HDMI 2.1 ports supporting 4K@120Hz gaming. The low input lag (<5ms) makes this TV excellent for competitive gaming on PS5 and Xbox Series X. Game Mode automatically detects and optimizes for gaming content.</p>

<h3>Connectivity & Ports</h3>
<ul>
<li>4 x HDMI 2.1 ports (supports 4K@120Hz)</li>
<li>WiFi 6 (802.11ax) dual-band connectivity</li>
<li>Bluetooth 5.0 for wireless peripherals</li>
<li>2 x USB 3.0 ports for external storage</li>
<li>Optical S/PDIF audio output</li>
<li>Ethernet RJ45 for stable wired connection</li>
</ul>

<h3>Audio System</h3>
<p>The 2.2.2 channel speaker system with 60W output provides immersive sound with Dolby Atmos support. While not exceptional for audiophiles, it's more than adequate for casual viewing and gaming.</p>

<h3>Design & Build Quality</h3>
<p>Premium aluminum and glass construction gives this TV a sophisticated appearance. The ultra-thin bezels maximize the viewing area, and the adjustable stand offers height flexibility for wall mounting or TV stands.</p>

<h3>Pros - Why You Should Buy</h3>
<ul>
<li>Exceptional QLED display technology with vibrant colors</li>
<li>Multiple HDMI 2.1 ports ideal for gaming setup</li>
<li>Comprehensive smart features with excellent app support</li>
<li>Premium build quality and elegant design</li>
<li>Strong 4K upscaling performance</li>
<li>Excellent warranty and service support in Bangladesh</li>
<li>Voice assistant compatibility</li>
<li>Great motion handling for sports</li>
</ul>

<h3>Cons - Areas for Improvement</h3>
<ul>
<li>Can be pricey in premium market segment</li>
<li>Tizen OS updates sometimes slower than competitors</li>
<li>Potential for image retention with static content on QLED</li>
<li>Fan noise during extended high brightness usage</li>
<li>Less bright than high-end OLED models</li>
</ul>

<h3>Price & Value in Bangladesh</h3>
<p>Available at approximately ৳85,000 - ৳105,000 depending on retailer and current promotions. This represents excellent value for a premium 4K TV with QLED technology and comprehensive smart features.</p>

<h3>Best For</h3>
<p>Movie enthusiasts, gamers transitioning to next-gen consoles, sports fans, and anyone seeking a balanced premium TV experience. Perfect for living rooms where both entertainment and gaming are priorities.</p>

<h3>Overall Rating Breakdown</h3>
<ul>
<li>Picture Quality: 4.7/5</li>
<li>Color Accuracy: 4.6/5</li>
<li>Gaming Features: 4.7/5</li>
<li>Smart Features: 4.5/5</li>
<li>Audio Quality: 4.2/5</li>
<li>Build Quality: 4.6/5</li>
<li>Value for Money: 4.4/5</li>
<li><strong>Overall: 4.5/5</strong></li>
</ul>

<h3>Final Verdict</h3>
<p>The Samsung 55" 4K Smart TV is an excellent all-rounder that doesn't compromise on quality. Whether you're a movie buff, gaming enthusiast, or casual viewer, this TV delivers outstanding value. The QLED technology ensures vibrant, accurate colors, while the gaming features make it perfect for next-gen consoles. Highly recommended for anyone seeking premium entertainment in Bangladesh.</p>

<h3>Comparison with Competitors</h3>
<p>Compared to LG's QNED series, Samsung offers better color performance. Against budget alternatives like Xiaomi, it provides superior motion handling and gaming features. While OLED TVs from LG offer perfect blacks, Samsung's QLED provides better brightness and value for money in this price range.</p>

<h3>Where to Buy in Bangladesh</h3>
<p>Available at major retailers including Abans, Basundhara City electronics shops, and online platforms. Free delivery available in Dhaka with 2-3 days delivery time. Extended warranty options available up to 5 years.</p>
</div>`,
			Verdict:     "Premium 4K QLED with excellent gaming features and balanced performance",
			SEOKeywords: []string{"Samsung 55 inch 4K TV Bangladesh", "QLED TV review", "55 inch TV price BD", "4K gaming TV Bangladesh", "best 4K TV under 100k", "Samsung smart TV Bangladesh"},
		},
		{
			ProductName: "LG 55\" OLED TV",
			HTMLContent: `<div class="tv-review-content">
<h2>LG 55" OLED TV - Premium Display Technology Review</h2>

<h3>Introduction</h3>
<p>The LG 55" OLED TV represents the pinnacle of display technology currently available in Bangladesh. OLED (Organic Light Emitting Diode) technology produces perfect blacks and infinite contrast, setting a new standard for picture quality that traditional LED and QLED TVs cannot match.</p>

<h3>Display Technology - The Game Changer</h3>
<p>OLED technology fundamentally differs from LED-based TVs. Each pixel produces its own light independently, meaning pixels can be turned completely off for true blacks. This results in infinite contrast ratios and exceptional picture quality.</p>

<ul>
<li><strong>Perfect Black Levels:</strong> Each pixel turns off completely for true blacks, not dark gray</li>
<li><strong>Infinite Contrast Ratio:</strong> Unmatched contrast between bright and dark areas</li>
<li><strong>120Hz Native Refresh Rate:</strong> Exceptional motion handling for sports and gaming</li>
<li><strong>800+ Nits Peak Brightness (OLED Evo):</strong> Bright enough for any room despite OLED reputation</li>
<li><strong>Instant Pixel Response:</strong> <0.1ms response time for perfect motion</li>
<li><strong>98.5% DCI-P3 Color Gamut:</strong> Professional-level color accuracy</li>
</ul>

<h3>Smart Features & WebOS</h3>
<p>LG's webOS 24 is arguably the best TV operating system available. AI-powered features learn your preferences and optimize content automatically. The interface is smooth, responsive, and packed with useful features.</p>

<ul>
<li>AI Picture Pro for automatic content optimization</li>
<li>Game Optimizer with VRR support</li>
<li>Magic Remote with voice control and motion sensing</li>
<li>WiFi 6E for ultra-fast streaming</li>
<li>AI Sound Pro for immersive audio</li>
</ul>

<h3>Gaming & Sports Performance</h3>
<p>The instant pixel response time and 120Hz refresh rate make this TV exceptional for gaming. HDMI 2.1 on all ports supports 4K@120Hz gaming. Sports viewing is phenomenal thanks to perfect motion handling and infinite contrast.</p>

<h3>Design & Build Quality</h3>
<p>Premium construction with slim bezels and gallery-style stand design. The picture quality is so impressive that mounting it on a wall like artwork is a popular choice among OLED owners.</p>

<h3>Pros - Premium Features</h3>
<ul>
<li>Perfect black levels unmatched by any other technology</li>
<li>Infinite contrast ratio for cinematic experience</li>
<li>Instant pixel response for perfect gaming</li>
<li>Superior webOS interface</li>
<li>Wide viewing angles with no color shift</li>
<li>Exceptional motion handling for sports</li>
<li>Professional-grade color accuracy</li>
<li>Pixel-level HDR processing</li>
</ul>

<h3>Cons - Considerations</h3>
<ul>
<li>Premium pricing - significantly more expensive than QLED</li>
<li>Potential for image burn-in with static content (though rare with modern protections)</li>
<li>Less bright in HDR peaks compared to QLED</li>
<li>Limited stock availability in Bangladesh</li>
<li>May require recalibration after warranty expires</li>
</ul>

<h3>OLED Protection Features</h3>
<p>LG includes multiple protection technologies to prevent burn-in: Pixel Shift, Screensaver Detection, Logo Detection, and periodic Refresher cycles. These make modern OLED TVs very safe for regular viewing.</p>

<h3>Price & Value in Bangladesh</h3>
<p>Available at approximately ৳180,000 - ৳220,000. While expensive, the picture quality justifies the investment for those serious about their viewing experience.</p>

<h3>Best For</h3>
<p>Professional users, movie enthusiasts, discerning gamers, home theater enthusiasts, and anyone wanting the absolute best picture quality regardless of budget.</p>

<h3>Power Consumption</h3>
<p>OLED TVs are surprisingly efficient despite perfect blacks. Average consumption around 120W, with standby power <0.3W.</p>

<h3>Overall Rating Breakdown</h3>
<ul>
<li>Picture Quality: 5.0/5 ⭐</li>
<li>Contrast: 5.0/5 ⭐</li>
<li>Color Accuracy: 4.9/5</li>
<li>Gaming Performance: 4.8/5</li>
<li>Motion Handling: 4.9/5</li>
<li>Build Quality: 4.7/5</li>
<li>Software: 4.8/5</li>
<li><strong>Overall: 4.6/5</strong></li>
</ul>

<h3>Final Verdict</h3>
<p>The LG 55" OLED TV is the gold standard for display technology. While expensive, it offers unparalleled picture quality with perfect blacks and infinite contrast. For those willing to invest in the best viewing experience, this is the ultimate choice. Not just a TV, but a work of art.</p>

<h3>Warranty & Support</h3>
<p>2-3 years standard warranty with extended options available. Excellent service network throughout Bangladesh with quick response times.</p>
</div>`,
			Verdict:     "Gold standard OLED with perfect blacks and infinite contrast - ultimate viewing experience",
			SEOKeywords: []string{"LG OLED TV Bangladesh", "55 inch OLED price BD", "best OLED TV 2024", "LG OLED review", "premium 4K TV Bangladesh", "perfect black TV"},
		},
		{
			ProductName: "Xiaomi 43\" Full HD Smart TV",
			HTMLContent: `<div class="tv-review-content">
<h2>Xiaomi 43" Full HD Smart TV - Best Value Budget Option</h2>

<h3>Introduction</h3>
<p>The Xiaomi 43" Full HD Smart TV represents the best value-for-money option in the budget segment. At its price point, it offers an impressive feature set including a full smart TV platform, good picture quality, and reliable performance for everyday viewing.</p>

<h3>Display Specifications</h3>
<ul>
<li><strong>43-inch Screen:</strong> Perfect middle ground for most living rooms</li>
<li><strong>Full HD Resolution (1920x1080):</strong> Adequate for casual viewing and streaming</li>
<li><strong>IPS LED Panel:</strong> Wide 178-degree viewing angles</li>
<li><strong>250 Nits Brightness:</strong> Suitable for average room lighting</li>
<li><strong>60Hz Refresh Rate:</strong> Standard motion handling</li>
<li><strong>16.7 Million Colors:</strong> Good color reproduction for the price</li>
</ul>

<h3>Smart Features & Android TV</h3>
<p>Running Android TV 10.0, this Xiaomi TV provides access to Google's massive app ecosystem. Pre-installed apps include Netflix, Amazon Prime Video, YouTube, Hotstar, and many others. The interface is familiar to smartphone users.</p>

<ul>
<li>Google Assistant voice control built-in</li>
<li>Chromecast support for easy screen mirroring</li>
<li>Regular security updates from Google</li>
<li>Easy app installation and management</li>
<li>Google account integration</li>
</ul>

<h3>Connectivity & Ports</h3>
<ul>
<li>WiFi 5 (802.11ac) dual-band - 50-100 Mbps speeds</li>
<li>Bluetooth 4.2</li>
<li>3 x HDMI 2.0 ports</li>
<li>2 x USB 2.0 ports</li>
<li>Ethernet RJ45</li>
<li>3.5mm audio jack</li>
</ul>

<h3>Audio System</h3>
<p>Dual 8W stereo speakers (16W total) provide decent audio for the TV's price. Dolby Audio support adds some immersion, though serious audio enthusiasts may want external speakers.</p>

<h3>Picture Quality Performance</h3>
<p>The Full HD resolution is sharp for typical viewing distances. Color reproduction is good, though not exceptional. Brightness is adequate for most rooms. The IPS panel ensures consistent colors from side angles.</p>

<h3>Pros - Why It's Great Value</h3>
<ul>
<li>Extremely affordable price point</li>
<li>Full Android TV experience with massive app support</li>
<li>Good picture quality for the price</li>
<li>Wide IPS viewing angles</li>
<li>Google Assistant voice control</li>
<li>Compact 43-inch size for most rooms</li>
<li>Reliable performance for everyday use</li>
<li>Good warranty and local support</li>
</ul>

<h3>Cons - Limitations</h3>
<ul>
<li>Only Full HD, not 4K resolution</li>
<li>Limited audio output power</li>
<li>Not ideal for gaming due to 60Hz only</li>
<li>Lower brightness for very bright rooms</li>
<li>HDMI 2.0 only, not HDMI 2.1</li>
<li>Slower processor compared to premium models</li>
<li>No advanced picture processing</li>
</ul>

<h3>Best For</h3>
<p>Budget-conscious buyers, small families, bedroom viewing, dorm rooms, offices, and anyone seeking a reliable smart TV without premium pricing. Perfect as a secondary TV.</p>

<h3>Price & Value in Bangladesh</h3>
<p>Available at approximately ৳22,000 - ৳26,000, making it the best value for a full-featured smart TV in Bangladesh. Excellent price-to-feature ratio.</p>

<h3>Power Consumption</h3>
<p>Average consumption around 45W, with maximum 75W. Annual estimated cost approximately ৳1,200.</p>

<h3>Overall Rating Breakdown</h3>
<ul>
<li>Picture Quality: 3.9/5</li>
<li>Smart Features: 4.3/5</li>
<li>Value for Money: 4.5/5 ⭐</li>
<li>App Ecosystem: 4.4/5</li>
<li>Performance: 4.0/5</li>
<li>Audio Quality: 3.5/5</li>
<li><strong>Overall: 4.2/5</strong></li>
</ul>

<h3>Final Verdict</h3>
<p>The Xiaomi 43" Full HD Smart TV is unbeatable value. While not 4K, it offers full smart TV features, good picture quality, and reliable performance. Highly recommended for budget buyers and first-time smart TV users. Best entry-level smart TV in Bangladesh.</p>

<h3>Where to Buy</h3>
<p>Available at major electronics retailers and online platforms. Fast delivery available in Dhaka with 2-year standard warranty.</p>
</div>`,
			Verdict:     "Best value budget smart TV with excellent Android ecosystem",
			SEOKeywords: []string{"Xiaomi 43 inch TV Bangladesh", "cheapest smart TV BD", "Xiaomi Full HD TV price", "budget smart TV 2024", "best value TV Bangladesh", "affordable 43 inch TV"},
		},
	}

	var firstUser models.UserModel
	if err := db.First(&firstUser).Error; err != nil {
		return nil
	}

	for _, reviewData := range detailedReviews {
		var product models.ProductModel
		if err := db.Where("name = ?", reviewData.ProductName).First(&product).Error; err != nil {
			fmt.Printf("   ⚠️  পণ্য '%s' পাওয়া যায়নি\n", reviewData.ProductName)
			continue
		}

		var existingReview models.ProductReviewModel
		result := db.Where("product_id = ?", product.ID).First(&existingReview)

		if result.Error == nil {
			// Update review with detailed HTML content
			existingReview.Reviews = &reviewData.HTMLContent
			existingReview.UpdatedAt = time.Now()

			// Add additional details with SEO keywords
			additionalDetails := map[string]interface{}{
				"seo_keywords": reviewData.SEOKeywords,
				"verdict":      reviewData.Verdict,
				"content_type": "detailed_html",
			}
			detailsJSON, _ := json.Marshal(additionalDetails)
			existingReview.AdditionalDetails = detailsJSON

			if err := db.Save(&existingReview).Error; err != nil {
				fmt.Printf("   ❌ '%s' আপডেট করতে ব্যর্থ\n", reviewData.ProductName)
				continue
			}

			fmt.Printf("   ✅ '%s' এর বিস্তারিত রিভিউ যোগ করা হয়েছে\n", reviewData.ProductName)
		}
	}

	return nil
}
