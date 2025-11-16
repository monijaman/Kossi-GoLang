package migrations

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UpdateBanglalinkReviewsWithComparison updates product reviews with HTML comparison content
func UpdateBanglalinkReviewsWithComparison(db *gorm.DB) error {
	fmt.Println("\n========== UPDATING BANGLALINK REVIEWS WITH HTML COMPARISON ==========")

	// English review HTML content
	englishReview := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Banglalink Review</title>
</head>
<body>
    <article class="banglalink-review">
        <header>
            <h1>Banglalink: Best Value Network in Bangladesh</h1>
            <p class="rating"><strong>Overall Rating: 4.1/5 ⭐⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>Executive Summary</h2>
            <p>Banglalink is Bangladesh's second-largest mobile operator with 35+ million subscribers, offering the best value for money with competitive pricing, solid 4G coverage, and diverse service offerings. Ranked #2 in the market with a 4.1/5 overall rating, Banglalink is the ideal choice for budget-conscious users, students, and professionals seeking reliable connectivity without premium pricing.</p>
        </section>

        <section class="network-quality">
            <h2>Network Quality &amp; Coverage</h2>
            <p><strong>Rating: 4.0/5</strong></p>
            <ul>
                <li>Urban Coverage: 99% with consistent 4G availability</li>
                <li>Rural Coverage: 60% with ongoing expansion</li>
                <li>Average Speed: 14 Mbps download, 8 Mbps peak hours</li>
                <li>Network Uptime: 97.5% reliability</li>
                <li>Latency: 25ms average (good for gaming)</li>
            </ul>
        </section>

        <section class="market-comparison">
            <h2>Market Comparison</h2>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Operator</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Rank</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Rating</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Best For</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Position vs Banglalink</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Grameenphone</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#1</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">4.65/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Best Quality</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Premium network, 14% more expensive</td>
                    </tr>
                    <tr style="background-color: #fff3cd;">
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Banglalink</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#2</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>4.1/5 ⭐</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Best Value</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>CURRENT - Best price-to-quality ratio</strong></td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Robi</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#3</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">4.05/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Good Alternative</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Similar quality, slightly cheaper data</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Airtel</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#4</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">3.05/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Regional Option</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Growing network, limited coverage</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Teletalk</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#5</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">2.85/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Budget Only</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Cheapest but outdated infrastructure</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pricing">
            <h2>Pricing Comparison</h2>
            <p><strong>Rating: 4.5/5 (Most Competitive)</strong></p>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Service</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Banglalink</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Grameenphone</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Robi</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Airtel</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Teletalk</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>100MB Daily</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">50 BDT ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">80 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">60 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">50 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">40 BDT</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>1GB Weekly</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">120 BDT ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">180 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">140 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">130 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">100 BDT</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pros">
            <h2>Pros (11 Major Advantages)</h2>
            <ul>
                <li>✅ <strong>Affordable call rates</strong> - Competitive pricing for local and international calls</li>
                <li>✅ <strong>Attractive data bundles</strong> - Regular promotions with bonus offers</li>
                <li>✅ <strong>Good 4G coverage</strong> - 99% urban and 60% rural coverage</li>
                <li>✅ <strong>User-friendly app</strong> - Intuitive account management interface</li>
                <li>✅ <strong>Digital payment support</strong> - Integrated with multiple payment platforms</li>
                <li>✅ <strong>Emergency balance</strong> - Feature for urgent calls when balance low</li>
                <li>✅ <strong>Flexible plans</strong> - Prepaid and postpaid options available</li>
                <li>✅ <strong>Frequent promotions</strong> - Regular special offers and bonuses</li>
                <li>✅ <strong>Social media bundles</strong> - Unlimited access to popular platforms</li>
                <li>✅ <strong>Money services</strong> - Comprehensive digital payment ecosystem</li>
                <li>✅ <strong>Multiple support channels</strong> - Phone, app, USSD, and social media</li>
            </ul>
        </section>

        <section class="cons">
            <h2>Cons (7 Main Limitations)</h2>
            <ul>
                <li>⚠️ <strong>Peak hour congestion</strong> - Network slowdowns during 6-10 PM</li>
                <li>⚠️ <strong>Limited rural coverage</strong> - Only 60% coverage in villages</li>
                <li>⚠️ <strong>Long customer wait times</strong> - Support queues during peak periods</li>
                <li>⚠️ <strong>Higher roaming rates</strong> - More expensive than competitors internationally</li>
                <li>⚠️ <strong>Speed throttling</strong> - Occasional limitations on some packages</li>
                <li>⚠️ <strong>Complex SIM activation</strong> - KYC process could be simplified</li>
                <li>⚠️ <strong>Restrictive validity</strong> - Some packages have strict expiry periods</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>Verdict</h2>
            <p><strong>✅ HIGHLY RECOMMENDED - Best Value for Money</strong></p>
            <p>Banglalink offers the best value-to-quality ratio in Bangladesh's mobile market. As the #2 operator with 35+ million subscribers, it provides reliable 4G coverage in urban areas, competitive pricing that undercuts Grameenphone by 14%, and a growing digital services ecosystem.</p>
        </section>
    </article>
</body>
</html>`

	// Bangla review HTML content
	bengaliReview := `<!DOCTYPE html>
<html lang="bn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>বাংলালিংক রিভিউ</title>
</head>
<body>
    <article class="banglalink-review-bn">
        <header>
            <h1>বাংলালিংক: বাংলাদেশের সেরা মূল্যের নেটওয়ার্ক</h1>
            <p class="rating"><strong>সামগ্রিক রেটিং: ৪.১/৫ ⭐⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>সংক্ষিপ্ত পর্যালোচনা</h2>
            <p>বাংলালিংক বাংলাদেশের দ্বিতীয় বৃহত্তম মোবাইল অপারেটর যা ৩৫+ মিলিয়ন গ্রাহকদের সেবা প্রদান করে। প্রতিযোগিতামূলক মূল্য, শক্তিশালী ৪জি কভারেজ এবং বৈচিত্র্যময় সেবার মাধ্যমে সেরা মূল্যের মূল্য প্রদান করে। বাজারে ২নম্বর স্থানে ৪.১/৫ রেটিং সহ, বাংলালিংক বাজেট-সচেতন ব্যবহারকারী, শিক্ষার্থী এবং নির্ভরযোগ্য সংযোগ খোঁজেন এমন পেশাদারদের জন্য আদর্শ।</p>
        </section>

        <section class="network-quality">
            <h2>নেটওয়ার্ক গুণমান এবং কভারেজ</h2>
            <p><strong>রেটিং: ৪.০/৫</strong></p>
            <ul>
                <li>শহরের কভারেজ: ৯৯% ধারাবাহিক ৪জি উপলব্ধতা</li>
                <li>গ্রামীণ কভারেজ: ৬০% এবং ক্রমাগত সম্প্রসারণ</li>
                <li>গড় গতি: ১४ এমবিপিএস ডাউনলোড, ८ এমবিপিএস পিক আওয়ার্স</li>
                <li>নেটওয়ার্ক আপটাইম: ९७.५% নির্ভরযোগ্যতা</li>
                <li>লেটেন্সি: २५এমএস গড় (গেমিংয়ের জন্য ভাল)</li>
            </ul>
        </section>

        <section class="market-comparison">
            <h2>বাজার তুলনা</h2>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">অপারেটর</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">র্যাঙ্ক</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">রেটিং</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">সেরা</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">বাংলালিংকের বিপরীতে</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>গ্রামীণফোন</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">नम्बर १</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">४.६५/५ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">সেরা মান</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">প্রিমিয়াম নেটওয়ার্ক, १४% বেশি ব्যয়บহুल</td>
                    </tr>
                    <tr style="background-color: #fff3cd;">
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>বাংলালিংক</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">नम्बर २</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>४.१/५ ⭐</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>সেরা মূল्य</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>বর্তমান - সেরা मूल्य-थॉ-मान अनुपat</strong></td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="verdict">
            <h2>সিদ্ধান্ত</h2>
            <p><strong>✅ অत्यंत सुपारिशकृत - সেरে मूল्यের मূल्य</strong></p>
            <p>बाংলालिंक बाংলादेशीर मोबाइल बाजारे सेरे मूल्य-थॉ-मान अनुपात प्रदान करে। ३५+ मिलिјन ग्राहक साथ द्वितीय बृहत्तम अपरेटोर हिसाबे, एटा शहुरे एलाकार मेँ निर्भरयोग्य ४जि कभरेज, ग्रामीणफोनेर चेये १४% कम मূल्य प्रदान करे।</p>
        </section>
    </article>
</body>
</html>`

	// Additional details JSON
	additionalDetails := datatypes.JSONMap{
		"market_rank":               2,
		"subscribers_millions":      35,
		"overall_rating":            4.1,
		"network_quality_rating":    4.0,
		"coverage_rating":           4.0,
		"data_speed_rating":         4.0,
		"pricing_rating":            4.5,
		"package_variety_rating":    4.0,
		"customer_support_rating":   3.5,
		"value_for_money_rating":    4.25,
		"avg_data_speed_mbps":       14,
		"peak_hour_speed_mbps":      8,
		"network_uptime_percent":    97.5,
		"latency_ms":                25,
		"urban_coverage_percent":    99.0,
		"rural_coverage_percent":    60.0,
		"hotline_number":            "111",
		"website_url":               "https://www.banglalink.com.bd",
		"verdict":                   "Highly Recommended - Best Value for Money",
		"recommendation_score":      8,
	}

	detailsJSON, err := json.Marshal(additionalDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Update English review
	result := db.Model(&map[string]interface{}{}).
		Table("product_reviews").
		Where("id = ? AND product_id = (SELECT id FROM products WHERE slug = ?)", 240, "banglalink").
		Updates(map[string]interface{}{
			"reviews":               englishReview,
			"additional_details":    string(detailsJSON),
			"updated_at":            gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating English review: %v", result.Error)
		return result.Error
	}

	log.Printf("Updated English review: %d rows affected", result.RowsAffected)

	// Update Bangla translation
	bengaliDetails := datatypes.JSONMap{
		"language":             "Bengali",
		"market_rank":          2,
		"subscribers_millions": 35,
		"overall_rating":       4.1,
		"verdict":              "অত्यंत सुपারिशकृत - সेरे मूल्यের मूल्य",
		"recommendation_score": 8,
	}

	bengaliJSON, err := json.Marshal(bengaliDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal Bengali JSON: %w", err)
	}

	result = db.Model(&map[string]interface{}{}).
		Table("product_review_translations").
		Where("product_review_id = ? AND locale = ?", 240, "bn").
		Updates(map[string]interface{}{
			"reviews":            bengaliReview,
			"additional_details": string(bengaliJSON),
			"updated_at":         gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating Bangla translation: %v", result.Error)
		return result.Error
	}

	log.Printf("Updated Bangla translation: %d rows affected", result.RowsAffected)

	fmt.Println("✅ Banglalink reviews updated successfully!")
	fmt.Printf("English review: %d rows affected\n", result.RowsAffected)

	return nil
}