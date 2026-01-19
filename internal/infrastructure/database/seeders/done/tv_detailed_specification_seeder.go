package seeders

import (
	"encoding/json"
	"kossti/internal/infrastructure/database/models"
	"time"

	"gorm.io/gorm"
)

// TVDetailedSpecificationSeeder handles seeding detailed HTML specifications for TV products
type TVDetailedSpecificationSeeder struct {
	BaseSeeder
}

// NewTVDetailedSpecificationSeeder creates a new TV detailed specification seeder
func NewTVDetailedSpecificationSeeder() *TVDetailedSpecificationSeeder {
	return &TVDetailedSpecificationSeeder{
		BaseSeeder: BaseSeeder{name: "TV Detailed Specifications"},
	}
}

// Seed implements the Seeder interface
func (tdss *TVDetailedSpecificationSeeder) Seed(db *gorm.DB) error {
	// Get TV category
	var tvCategory models.CategoryModel
	if err := db.Where("id = ?", 124).First(&tvCategory).Error; err != nil {
		return nil
	}

	// TV detailed specifications
	tvDetails := map[string]map[string]interface{}{
		"Samsung 55\" 4K Smart TV": {
			"html": `<div class="tv-specification-content">
<h2>Samsung 55" 4K Smart TV - Detailed Specifications & Review</h2>

<h3>Display Technology</h3>
<ul>
<li><strong>Screen Size:</strong> 55 inches (140 cm diagonal)</li>
<li><strong>Resolution:</strong> 4K UHD (3840 x 2160 pixels)</li>
<li><strong>Panel Type:</strong> QLED (Quantum Dot LED)</li>
<li><strong>Refresh Rate:</strong> 60Hz native, 120Hz motion handling</li>
<li><strong>Brightness:</strong> 400+ nits peak brightness</li>
<li><strong>Contrast Ratio:</strong> Exceptional with QLED technology</li>
<li><strong>Color Accuracy:</strong> 99% DCI-P3 color space coverage</li>
<li><strong>Response Time:</strong> <5ms for gaming</li>
</ul>

<h3>Smart Features</h3>
<ul>
<li><strong>Operating System:</strong> Tizen 6.0 or later</li>
<li><strong>Built-in Apps:</strong> Netflix, Amazon Prime Video, YouTube, Disney+</li>
<li><strong>Voice Control:</strong> Bixby, Alexa compatible</li>
<li><strong>AI Upscaling:</strong> AI-based video enhancement technology</li>
<li><strong>Game Mode:</strong> Optimized for gaming consoles</li>
</ul>

<h3>Connectivity & Ports</h3>
<ul>
<li><strong>WiFi:</strong> Dual-band WiFi 6 (802.11ax)</li>
<li><strong>Bluetooth:</strong> 5.0 with multiple device support</li>
<li><strong>HDMI Ports:</strong> 4 x HDMI 2.1 (supports 4K@120Hz)</li>
<li><strong>USB Ports:</strong> 2 x USB 3.0</li>
<li><strong>Ethernet:</strong> RJ45 port for wired connection</li>
<li><strong>Audio Connectors:</strong> Optical S/PDIF output</li>
</ul>

<h3>Audio System</h3>
<ul>
<li><strong>Speaker Configuration:</strong> 2.2.2 channel speaker system</li>
<li><strong>Total Power:</strong> 60W output</li>
<li><strong>Audio Technologies:</strong> Dolby Atmos support</li>
<li><strong>Adaptive Sound:</strong> AI-powered audio enhancement</li>
</ul>

<h3>Design & Build Quality</h3>
<ul>
<li><strong>Dimensions:</strong> 1231 x 710 x 50 mm (without stand)</li>
<li><strong>Weight:</strong> 18 kg (with stand)</li>
<li><strong>Stand Type:</strong> Premium metal stand with adjustable height</li>
<li><strong>Wall Mount:</strong> VESA 300x300 compatible</li>
<li><strong>Build Material:</strong> Premium aluminum and glass combination</li>
<li><strong>Bezels:</strong> Ultra-thin bezels for immersive viewing</li>
</ul>

<h3>Gaming Features</h3>
<ul>
<li><strong>FreeSync Premium:</strong> AMD FreeSync Pro certification</li>
<li><strong>HDMI 2.1:</strong> All ports support HDMI 2.1 for gaming</li>
<li><strong>120Hz @ 4K:</strong> Perfect for next-gen gaming consoles</li>
<li><strong>Auto Game Detection:</strong> Automatic game mode activation</li>
<li><strong>Input Lag:</strong> <10ms input lag for competitive gaming</li>
</ul>

<h3>Picture Quality Features</h3>
<ul>
<li><strong>HDR Support:</strong> HDR10, HDR10+, HLG</li>
<li><strong>Motion Handling:</strong> TruMotion 240Hz equivalent processing</li>
<li><strong>Black Levels:</strong> Deep blacks with quantum dot precision</li>
<li><strong>Anti-Glare:</strong> Matte finish anti-reflective coating</li>
<li><strong>Wide Viewing Angle:</strong> 178-degree wide viewing angles</li>
</ul>

<h3>Pros - Why Choose This TV</h3>
<ul>
<li>Excellent QLED display technology with vibrant colors</li>
<li>Smooth motion handling ideal for sports and gaming</li>
<li>Comprehensive smart features and app ecosystem</li>
<li>Strong 4K upscaling performance</li>
<li>Multiple HDMI 2.1 ports for gaming setup</li>
<li>Good build quality and premium design</li>
<li>Voice control compatibility with popular assistants</li>
<li>Extended warranty and service support</li>
</ul>

<h3>Cons - Areas for Improvement</h3>
<ul>
<li>Can be expensive in premium segment</li>
<li>Tizen OS may have slower app updates</li>
<li>Can suffer from image retention if static content displayed long</li>
<li>Fan noise in some models during high usage</li>
</ul>

<h3>Power Consumption</h3>
<ul>
<li><strong>Average Usage:</strong> 150W</li>
<li><strong>Maximum:</strong> 250W</li>
<li><strong>Standby:</strong> <0.5W</li>
<li><strong>Energy Rating:</strong> A+ efficiency</li>
</ul>

<h3>Warranty & Support</h3>
<ul>
<li><strong>Standard Warranty:</strong> 2 years parts and labor</li>
<li><strong>Screen Warranty:</strong> 3 years against defects</li>
<li><strong>Extended Options:</strong> Available up to 5 years</li>
<li><strong>Service Centers:</strong> Available nationwide in Bangladesh</li>
</ul>

<h3>Price Range in Bangladesh</h3>
<p>Approximately ৳ 85,000 - ৳ 105,000 depending on retailer and promotions</p>

<h3>Best For</h3>
<p>Movie enthusiasts, gamers, sports lovers, and anyone seeking premium 4K entertainment with excellent smart TV features. Perfect for living rooms and entertainment setups.</p>

<h3>Final Verdict</h3>
<p>The Samsung 55" 4K QLED TV represents excellent value for premium 4K viewing. With its impressive display technology, gaming features, and smart capabilities, it's a top choice for mid-to-high budget buyers in Bangladesh.</p>
</div>`,
			"verdict":     "Premium 4K QLED display with excellent gaming features",
			"recommended": true,
			"rating_split": map[string]float32{
				"display_quality": 4.7,
				"color_accuracy":  4.6,
				"gaming":          4.7,
				"smart_features":  4.5,
				"audio":           4.3,
				"build_quality":   4.6,
			},
			"seo_keywords": []string{
				"Samsung 55 inch 4K TV Bangladesh",
				"Best 4K TV in Bangladesh",
				"Samsung QLED TV review",
				"55 inch TV price Bangladesh",
				"4K smart TV Bangladesh",
				"Gaming TV 4K Bangladesh",
			},
			"availability": "In stock at major retailers",
			"delivery":     "Free delivery in Dhaka, 2-3 days",
		},
		"LG 55\" OLED TV": {
			"html": `<div class="tv-specification-content">
<h2>LG 55" OLED TV - Premium Display Technology Review</h2>

<h3>Display Technology</h3>
<ul>
<li><strong>Screen Size:</strong> 55 inches (140 cm diagonal)</li>
<li><strong>Resolution:</strong> 4K UHD (3840 x 2160 pixels)</li>
<li><strong>Panel Type:</strong> OLED (Organic Light Emitting Diode)</li>
<li><strong>Refresh Rate:</strong> 120Hz native support</li>
<li><strong>Brightness:</strong> 800+ nits peak brightness (OLED Evo)</li>
<li><strong>Contrast Ratio:</strong> Perfect black levels with infinite contrast</li>
<li><strong>Color Accuracy:</strong> 98.5% DCI-P3 color gamut</li>
<li><strong>Response Time:</strong> Instant pixel response for perfect motion</li>
<li><strong>Viewing Angle:</strong> Perfect 180-degree viewing angles</li>
</ul>

<h3>OLED Advantages</h3>
<ul>
<li>Perfect blacks with self-emissive pixels</li>
<li>Infinite contrast ratio for cinematic experience</li>
<li>No backlight bleed or halo effects</li>
<li>Instant response time - ideal for gaming</li>
<li>Wide viewing angles without color shift</li>
<li>Superior motion handling</li>
</ul>

<h3>Smart Features</h3>
<ul>
<li><strong>Operating System:</strong> webOS 24 (latest AI-powered OS)</li>
<li><strong>App Support:</strong> Netflix, Amazon Prime, YouTube, Disney+ pre-installed</li>
<li><strong>Voice Control:</strong> Built-in voice assistant, Google & Alexa compatible</li>
<li><strong>AI Upscaling:</strong> AI Picture Pro technology</li>
<li><strong>Game Optimizer:</strong> Advanced gaming features with VRR</li>
</ul>

<h3>Connectivity</h3>
<ul>
<li><strong>WiFi:</strong> WiFi 6E (802.11ax) ultra-fast connection</li>
<li><strong>Bluetooth:</strong> 5.2 with low latency audio</li>
<li><strong>HDMI:</strong> 4 x HDMI 2.1 (all support 4K@120Hz)</li>
<li><strong>USB:</strong> 2 x USB 3.0 ports</li>
<li><strong>Networking:</strong> Ethernet RJ45 connection</li>
<li><strong>Audio Output:</strong> Optical and 3.5mm audio jack</li>
</ul>

<h3>Audio System</h3>
<ul>
<li><strong>Speakers:</strong> 4.2 channel premium speaker configuration</li>
<li><strong>Power Output:</strong> 70W total audio power</li>
<li><strong>Audio Tech:</strong> Dolby Atmos support</li>
<li><strong>AI Audio:</strong> AI Sound Pro for immersive experience</li>
</ul>

<h3>Gaming & Sports Performance</h3>
<ul>
<li><strong>FreeSync Premium Pro:</strong> Certified for AMD graphics</li>
<li><strong>G-Sync Compatible:</strong> Works with NVIDIA graphics</li>
<li><strong>120Hz Gaming:</strong> Full 4K@120Hz for next-gen consoles</li>
<li><strong>Sports Mode:</strong> Optimized motion processing for live sports</li>
<li><strong>Black Frame Insertion:</strong> Enhanced motion clarity</li>
</ul>

<h3>Design & Build</h3>
<ul>
<li><strong>Dimensions:</strong> 1226 x 706 x 45mm (without stand)</li>
<li><strong>Weight:</strong> 19 kg with stand</li>
<li><strong>Design:</strong> Premium slim bezel, gallery stand design</li>
<li><strong>Wall Mount:</strong> VESA 300x300 compatible</li>
<li><strong>Materials:</strong> Premium aluminum, tempered glass</li>
</ul>

<h3>HDR & Color</h3>
<ul>
<li><strong>HDR Support:</strong> HDR10, Dolby Vision, HLG</li>
<li><strong>OLED Evo Technology:</strong> Enhanced brightness and efficiency</li>
<li><strong>ColorVision Engine:</strong> Professional color mapping</li>
<li><strong>HDR Zones:</strong> Pixel-level HDR processing</li>
</ul>

<h3>Pros - Why Choose OLED</h3>
<ul>
<li>Perfect black levels unmatched by any other technology</li>
<li>Incredible contrast and cinematic picture quality</li>
<li>Fast response time perfect for gaming</li>
<li>Premium webOS interface with intuitive navigation</li>
<li>Superior motion handling for sports</li>
<li>Wide viewing angles without quality loss</li>
<li>Excellent AI upscaling technology</li>
<li>Professional-grade color accuracy</li>
</ul>

<h3>Cons - Areas to Consider</h3>
<ul>
<li>Premium pricing - more expensive than QLED</li>
<li>Potential for image burn-in with static content (rare with protections)</li>
<li>Less bright than QLED in HDR peak</li>
<li>Limited stock availability in Bangladesh</li>
</ul>

<h3>OLED Protection Features</h3>
<ul>
<li>Pixel Shift technology prevents burn-in</li>
<li>Screensaver detection automatic activation</li>
<li>Logo detection to avoid static images</li>
<li>Refresher cycle technology</li>
</ul>

<h3>Power Consumption</h3>
<ul>
<li><strong>Average:</strong> 120W</li>
<li><strong>Maximum:</strong> 200W</li>
<li><strong>Standby:</strong> <0.3W</li>
<li><strong>Energy Rating:</strong> A+ efficient</li>
</ul>

<h3>Price Range in Bangladesh</h3>
<p>Approximately ৳ 180,000 - ৳ 220,000 (Premium pricing for OLED technology)</p>

<h3>Best For</h3>
<p>Professional users, movie enthusiasts, gamers demanding the best, home theater enthusiasts, and anyone wanting the absolute best picture quality regardless of budget.</p>

<h3>Final Verdict</h3>
<p>LG 55" OLED TV is the gold standard for display technology. While expensive, it offers unparalleled picture quality with perfect blacks and infinite contrast. Best investment for those wanting the ultimate viewing experience.</p>
</div>`,
			"verdict":     "Gold standard OLED with perfect black levels",
			"recommended": true,
			"rating_split": map[string]float32{
				"display_quality": 5.0,
				"contrast":        5.0,
				"color_accuracy":  4.9,
				"gaming":          4.8,
				"motion_handling": 4.9,
				"build_quality":   4.7,
			},
			"seo_keywords": []string{
				"LG OLED TV Bangladesh",
				"55 inch OLED TV price",
				"Best OLED TV 2024",
				"LG 55 OLED review",
				"Premium 4K OLED TV",
				"Gaming OLED TV Bangladesh",
			},
			"availability": "Limited stock - Pre-order available",
			"delivery":     "Free delivery Dhaka, 5-7 days",
		},
		"Xiaomi 43\" Full HD Smart TV": {
			"html": `<div class="tv-specification-content">
<h2>Xiaomi 43" Full HD Smart TV - Budget-Friendly Smart Entertainment</h2>

<h3>Display Specifications</h3>
<ul>
<li><strong>Screen Size:</strong> 43 inches (109 cm diagonal)</li>
<li><strong>Resolution:</strong> Full HD (1920 x 1080 pixels)</li>
<li><strong>Panel Type:</strong> IPS LED</li>
<li><strong>Refresh Rate:</strong> 60Hz</li>
<li><strong>Brightness:</strong> 250 nits</li>
<li><strong>Viewing Angle:</strong> 178 degrees (IPS technology)</li>
<li><strong>Color Support:</strong> 16.7 million colors</li>
</ul>

<h3>Smart Features</h3>
<ul>
<li><strong>Operating System:</strong> Android TV 10.0</li>
<li><strong>Processor:</strong> Quad-core processor</li>
<li><strong>RAM:</strong> 2GB</li>
<li><strong>Storage:</strong> 8GB internal storage</li>
<li><strong>Apps:</strong> YouTube, Netflix, Prime Video, Hotstar</li>
<li><strong>Voice Control:</strong> Google Assistant built-in</li>
<li><strong>Casting:</strong> Chromecast support</li>
</ul>

<h3>Connectivity</h3>
<ul>
<li><strong>WiFi:</strong> Dual-band WiFi 5 (802.11ac)</li>
<li><strong>Bluetooth:</strong> 4.2</li>
<li><strong>HDMI:</strong> 3 x HDMI 2.0 ports</li>
<li><strong>USB:</strong> 2 x USB 2.0</li>
<li><strong>Ethernet:</strong> RJ45 port</li>
<li><strong>Audio:</strong> 3.5mm jack and optical output</li>
</ul>

<h3>Audio System</h3>
<ul>
<li><strong>Speakers:</strong> 2 x 8W stereo speakers</li>
<li><strong>Total Power:</strong> 16W output</li>
<li><strong>Audio Tech:</strong> Dolby Audio support</li>
<li><strong>Sound Quality:</strong> Good for daily viewing</li>
</ul>

<h3>Design & Dimensions</h3>
<ul>
<li><strong>Dimensions:</strong> 972 x 560 x 70mm (without stand)</li>
<li><strong>Weight:</strong> 9.8 kg with stand</li>
<li><strong>Bezels:</strong> Slim bezel design for modern look</li>
<li><strong>Stand:</strong> Adjustable height stand</li>
<li><strong>Wall Mount:</strong> VESA 200x200 compatible</li>
</ul>

<h3>Picture Processing</h3>
<ul>
<li>Motion Estimation Motion Compensation (MEMC)</li>
<li>Dynamic Contrast enhancement</li>
<li>Color Optimization technology</li>
<li>Image noise reduction</li>
</ul>

<h3>Pros - Value for Money</h3>
<ul>
<li>Extremely affordable price point</li>
<li>Full-featured Android TV experience</li>
<li>Good picture quality for price</li>
<li>Wide app ecosystem</li>
<li>Google Assistant voice control</li>
<li>Compact 43-inch size suitable for most rooms</li>
<li>Reliable performance for everyday use</li>
<li>Good warranty and local support</li>
</ul>

<h3>Cons - Limitations</h3>
<ul>
<li>Only Full HD, not 4K resolution</li>
<li>Limited audio output power</li>
<li>Not ideal for gaming or sports</li>
<li>Lower brightness level</li>
<li>HDMI 2.0 only (not 2.1)</li>
<li>Slower processor compared to premium models</li>
</ul>

<h3>Power Consumption</h3>
<ul>
<li><strong>Average Usage:</strong> 45W</li>
<li><strong>Maximum:</strong> 75W</li>
<li><strong>Standby:</strong> <0.5W</li>
<li><strong>Annual Cost:</strong> Approximately ৳ 1,200</li>
</ul>

<h3>Video Format Support</h3>
<ul>
<li>4K video playback via USB/HDMI (upscaled to Full HD)</li>
<li>Support for MP4, MKV, AVI formats</li>
<li>HEVC codec support</li>
</ul>

<h3>Perfect For</h3>
<ul>
<li>Budget-conscious buyers</li>
<li>Bedrooms and personal use</li>
<li>Daily news and general entertainment</li>
<li>Students' hostels and small apartments</li>
<li>First-time smart TV users</li>
<li>Secondary TV purchase</li>
</ul>

<h3>Price Range in Bangladesh</h3>
<p>Approximately ৳ 22,000 - ৳ 26,000 (Best budget option)</p>

<h3>Warranty & Support</h3>
<ul>
<li><strong>Warranty:</strong> 2 years</li>
<li><strong>Service:</strong> Local authorized service centers</li>
<li><strong>Support:</strong> Online and phone support available</li>
</ul>

<h3>Final Verdict</h3>
<p>Xiaomi 43" Full HD Smart TV is the best budget-friendly option for smart entertainment in Bangladesh. While not 4K, it offers excellent value with full Android TV features, good picture quality, and extensive app support. Highly recommended for budget buyers and first-time smart TV users.</p>
</div>`,
			"verdict":     "Best budget smart TV with great Android ecosystem",
			"recommended": true,
			"rating_split": map[string]float32{
				"display_quality": 3.9,
				"smart_features":  4.3,
				"value_for_money": 4.5,
				"app_ecosystem":   4.4,
				"performance":     4.0,
				"support":         3.8,
			},
			"seo_keywords": []string{
				"Xiaomi 43 inch TV Bangladesh",
				"Cheapest smart TV Bangladesh",
				"Xiaomi Full HD TV price",
				"Budget smart TV Bangladesh 2024",
				"43 inch affordable TV",
				"Best value TV Bangladesh",
			},
			"availability": "Widely available",
			"delivery":     "Next day delivery in Dhaka",
			"price_note":   "Lowest price in Full HD category",
		},
	}

	// Get first user for review
	var firstUser models.UserModel
	if err := db.First(&firstUser).Error; err != nil {
		return nil
	}

	// Update reviews with detailed specifications
	for productName, specs := range tvDetails {
		var product models.ProductModel
		if err := db.Where("name = ?", productName).First(&product).Error; err != nil {
			continue
		}

		// Find existing review
		var review models.ProductReviewModel
		result := db.Where("product_id = ?", product.ID).First(&review)

		if result.Error == nil {
			// Update review with HTML specifications
			htmlSpec := specs["html"].(string)
			review.Reviews = &htmlSpec

			// Marshal additional details
			additionalDetails := make(map[string]interface{})
			for key, value := range specs {
				if key != "html" {
					additionalDetails[key] = value
				}
			}

			detailsJSON, err := json.Marshal(additionalDetails)
			if err != nil {
				continue
			}

			review.AdditionalDetails = detailsJSON
			review.UpdatedAt = time.Now()

			db.Save(&review)
		}
	}

	return nil
}
