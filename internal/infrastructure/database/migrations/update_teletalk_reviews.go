package migrations

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UpdateTeletalkReviewsWithComparison updates product reviews with HTML comparison content
func UpdateTeletalkReviewsWithComparison(db *gorm.DB) error {
	fmt.Println("\n========== UPDATING TELETALK REVIEWS WITH HTML COMPARISON ==========")

	// English review HTML content
	englishReview := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Teletalk Review</title>
</head>
<body>
    <article class="teletalk-review">
        <header>
            <h1>Teletalk: Cheapest Option with Outdated Infrastructure</h1>
            <p class="rating"><strong>Overall Rating: 2.85/5 ⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>Executive Summary</h2>
            <p>Teletalk is Bangladesh's state-owned operator with 5+ million subscribers, offering the cheapest rates but with outdated infrastructure. Ranked #5 in the market with a 2.85/5 overall rating, Teletalk is only suitable for users with extremely limited budgets who can accept poor service quality.</p>
        </section>

        <section class="network-quality">
            <h2>Network Quality &amp; Coverage</h2>
            <p><strong>Rating: 2.5/5</strong></p>
            <ul>
                <li>Urban Coverage: 85% with limited modernization</li>
                <li>Rural Coverage: 30% with minimal expansion</li>
                <li>Average Speed: 5 Mbps download, 2 Mbps peak hours</li>
                <li>Network Uptime: 95% reliability</li>
                <li>Latency: 60ms average</li>
            </ul>
        </section>

        <section class="pros">
            <h2>Pros (4 Main Advantages)</h2>
            <ul>
                <li> <strong>Extremely cheap</strong> - Lowest rates in the market</li>
                <li> <strong>Widest reach</strong> - Available in most government areas</li>
                <li> <strong>Basic services</strong> - Essential calling and SMS covered</li>
                <li> <strong>Government backing</strong> - State-owned stability</li>
            </ul>
        </section>

        <section class="cons">
            <h2>Cons (9 Main Limitations)</h2>
            <ul>
                <li> <strong>Very slow speeds</strong> - Extremely poor data performance</li>
                <li> <strong>Outdated technology</strong> - Legacy infrastructure</li>
                <li> <strong>Poor coverage</strong> - Limited rural availability</li>
                <li> <strong>Unreliable service</strong> - Frequent connection issues</li>
                <li> <strong>No 4G/5G</strong> - Only 2G/3G available</li>
                <li> <strong>Poor customer service</strong> - Limited support options</li>
                <li> <strong>No digital services</strong> - Minimal app ecosystem</li>
                <li> <strong>High latency</strong> - Not suitable for any modern use</li>
                <li> <strong>Slow upgrade cycle</strong> - Minimal infrastructure investment</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>Verdict</h2>
            <p><strong>NOT RECOMMENDED - Budget Only as Last Resort</strong></p>
            <p>Teletalk should only be considered as an absolute last resort for users with no other options or extremely limited budgets. The outdated infrastructure and poor service quality make it unsuitable for modern mobile usage.</p>
        </section>
    </article>
</body>
</html>`

	// Bangla review
	bengaliReview := `<!DOCTYPE html>
<html lang="bn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>টেলিটক রিভিউ</title>
</head>
<body>
    <article class="teletalk-review-bn">
        <header>
            <h1>টেলিটক: সবচেয়ে সস্তা বিকল্প কিন্তু পুরানো অবকাঠামো</h1>
            <p class="rating"><strong>সামগ্রিক রেটিং: ২.৮৫/৫ ⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>সংক্ষিপ্ত পর্যালোচনা</h2>
            <p>টেলিটক বাংলাদেশের সরকারী মালিকানাধীন অপারেটর ৫+ মিলিয়ন গ্রাহক সেবা প্রদান করে সবচেয়ে সস্তা হার সহ কিন্তু পুরানো অবকাঠামো সহ। শুধুমাত্র চরম সীমিত বাজেটের ব্যবহারকারীদের জন্য উপযুক্ত।</p>
        </section>

        <section class="cons">
            <h2>অসুবিধাসমূহ (৯টি প্রধান সীমাবদ্ধতা)</h2>
            <ul>
                <li> <strong>অত্যন্ত ধীর গতি</strong> - অত্যন্ত দুর্বল ডেটা কর্মক্ষমতা</li>
                <li> <strong>পুরানো প্রযুক্তি</strong> - লিগ্যাসি অবকাঠামো</li>
                <li> <strong>দুর্বল কভারেজ</strong> - সীমিত গ্রামীণ উপলব্ধতা</li>
                <li> <strong>অনির্ভরযোগ্য সেবা</strong> - ঘন ঘন সংযোগ সমস্যা</li>
                <li> <strong>কোন ৪জি/৫জি নেই</strong> - শুধুমাত্র ২জি/৩জি উপলব্ধ</li>
                <li> <strong>দুর্বল গ্রাহক সেবা</strong> - সীমিত সহায়তা বিকল্প</li>
                <li> <strong>কোন ডিজিটাল সেবা নেই</strong> - ন্যূনতম অ্যাপ ইকোসিস্টেম</li>
                <li> <strong>উচ্চ লেটেন্সি</strong> - আধুনিক ব্যবহারের জন্য অনুপযুক্ত</li>
                <li> <strong>ধীর আপগ্রেড চক্র</strong> - ন্যূনতম অবকাঠামো বিনিয়োগ</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>চূড়ান্ত মূল্যায়ন</h2>
            <p><strong>সুপারিশ করা হয় না - শেষ অবলম্বন হিসাবে শুধুমাত্র বাজেট</strong></p>
            <p>টেলিটক শুধুমাত্র চরম সীমিত বাজেটের ব্যবহারকারীদের জন্য বিবেচনা করা উচিত।</p>
        </section>
    </article>
</body>
</html>`

	// Additional details JSON
	additionalDetails := datatypes.JSONMap{
		"market_rank":             5,
		"subscribers_millions":    5,
		"overall_rating":          2.85,
		"network_quality_rating":  2.5,
		"coverage_rating":         2.8,
		"data_speed_rating":       2.2,
		"pricing_rating":          4.8,
		"package_variety_rating":  2.0,
		"customer_support_rating": 2.5,
		"value_for_money_rating":  2.5,
		"avg_data_speed_mbps":     5,
		"peak_hour_speed_mbps":    2,
		"network_uptime_percent":  95.0,
		"latency_ms":              60,
		"urban_coverage_percent":  85.0,
		"rural_coverage_percent":  30.0,
		"hotline_number":          "155",
		"website_url":             "https://www.teletalk.com.bd",
		"verdict":                 "Not Recommended - Budget Only as Last Resort",
		"recommendation_score":    3,
	}

	detailsJSON, err := json.Marshal(additionalDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Update English review
	result := db.Model(&map[string]interface{}{}).
		Table("product_reviews").
		Where("id = ? AND product_id = (SELECT id FROM products WHERE slug = ?)", 4, "teletalk").
		Updates(map[string]interface{}{
			"reviews":            englishReview,
			"additional_details": string(detailsJSON),
			"updated_at":         gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating English review: %v", result.Error)
		return result.Error
	}

	// Update Bangla translation
	bengaliDetails := datatypes.JSONMap{
		"language":             "Bengali",
		"market_rank":          5,
		"subscribers_millions": 5,
		"overall_rating":       2.85,
		"verdict":              "সুপারিশ করা হয় না - শেষ অবলম্বন হিসাবে শুধুমাত্র বাজেট",
		"recommendation_score": 3,
	}

	bengaliJSON, err := json.Marshal(bengaliDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal Bengali JSON: %w", err)
	}

	result = db.Model(&map[string]interface{}{}).
		Table("product_review_translations").
		Where("product_review_id = ? AND locale = ?", 4, "bn").
		Updates(map[string]interface{}{
			"reviews":            bengaliReview,
			"additional_details": string(bengaliJSON),
			"updated_at":         gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating Bangla translation: %v", result.Error)
		return result.Error
	}

	fmt.Println("✅ Teletalk reviews updated successfully!")

	return nil
}
