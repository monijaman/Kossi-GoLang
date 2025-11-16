-- ============================================================================
-- UPDATE PRODUCT REVIEWS WITH BANGLALINK COMPARISON DATA (PROPER HTML)
-- ============================================================================
-- Format: Clean, semantic HTML with proper structure
-- ============================================================================

-- ============================================================================
-- UPDATE ENGLISH REVIEW (ID: 240)
-- ============================================================================

UPDATE product_reviews
SET
    reviews = '<!DOCTYPE html>
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
            <p>Banglalink is Bangladesh''s second-largest mobile operator with 35+ million subscribers, offering the best value for money with competitive pricing, solid 4G coverage, and diverse service offerings. Ranked #2 in the market with a 4.1/5 overall rating, Banglalink is the ideal choice for budget-conscious users, students, and professionals seeking reliable connectivity without premium pricing.</p>
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
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Unlimited Monthly</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">999 BDT ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">1500 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">1200 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">999 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">800 BDT</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Local Call (per min)</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">0.50 BDT ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">1.50 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">1.0 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">0.50 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">0.30 BDT</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="packages">
            <h2>Packages &amp; Bundles</h2>
            <p><strong>Rating: 4.0/5 (Good Variety)</strong></p>
            <ul>
                <li><strong>Data Services:</strong> 20+ daily, weekly, and monthly bundles</li>
                <li><strong>Social Media:</strong> WhatsApp, Facebook, YouTube bundles</li>
                <li><strong>Voice Calls:</strong> Local, on-net, international options</li>
                <li><strong>OTT Services:</strong> Netflix, YouTube, streaming packages</li>
                <li><strong>Night Packages:</strong> Discounted rates 12 AM - 6 AM</li>
                <li><strong>Roaming:</strong> Day passes for 50+ countries</li>
            </ul>
        </section>

        <section class="digital-services">
            <h2>Digital Services</h2>
            <p><strong>Rating: 4.0/5</strong></p>
            <ul>
                <li><strong>Banglalink Money:</strong> Mobile wallet with 2 million+ users</li>
                <li><strong>Money Transfer:</strong> Fast P2P transfers to other users</li>
                <li><strong>Bill Payments:</strong> Electricity, water, internet bills</li>
                <li><strong>Merchant Payments:</strong> Business payment solutions</li>
                <li><strong>Easy Recharge:</strong> Multiple methods including bKash</li>
            </ul>
        </section>

        <section class="customer-support">
            <h2>Customer Support</h2>
            <p><strong>Rating: 3.5/5</strong></p>
            <ul>
                <li>Hotline: 111 (24/7 available)</li>
                <li>Mobile App: Highly rated user interface</li>
                <li>USSD Codes: Quick balance and usage checks (*123#)</li>
                <li>Social Media: Facebook and Twitter support</li>
                <li><em>Issue:</em> Long wait times during peak hours</li>
            </ul>
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
            <p><strong>✅ HIGHLY RECOMMENDED for:</strong></p>
            <ul>
                <li>Students seeking budget-friendly options</li>
                <li>Professionals wanting reliable connectivity</li>
                <li>Users prioritizing value for money</li>
                <li>Business owners needing digital payment solutions</li>
                <li>Heavy data users with good data packages</li>
                <li>Anyone switching from expensive alternatives</li>
            </ul>
        </section>

        <section class="final-assessment">
            <h2>Final Assessment</h2>
            <p>Banglalink offers the best value-to-quality ratio in Bangladesh''s mobile market. As the #2 operator with 35+ million subscribers, it provides reliable 4G coverage in urban areas, competitive pricing that undercuts Grameenphone by 14%, and a growing digital services ecosystem. While it has limitations in rural coverage and occasional peak-hour congestion, these are offset by its affordable packages and user-friendly services. For most users, Banglalink represents the ideal balance between quality and cost.</p>
        </section>

        <section class="competitor-comparison">
            <h2>Comparison with Competitors</h2>
            <p><strong>vs Grameenphone:</strong> GP offers better network quality and rural coverage but charges 14% more. Banglalink is the smarter choice unless you absolutely need premium coverage.</p>
            <p><strong>vs Robi:</strong> Similar ratings but Banglalink has slightly better pricing on unlimited plans and more features.</p>
            <p><strong>vs Airtel:</strong> Banglalink offers 25% better network quality with comparable pricing.</p>
            <p><strong>vs Teletalk:</strong> Teletalk is only cheaper if you rarely use internet. Banglalink offers significantly better service.</p>
        </section>
    </article>
</body>
</html>',
    additional_details = jsonb_build_object(
        ''market_rank'', 2,
        ''subscribers_millions'', 35,
        ''overall_rating'', 4.1,
        ''network_quality_rating'', 4.0,
        ''coverage_rating'', 4.0,
        ''data_speed_rating'', 4.0,
        ''pricing_rating'', 4.5,
        ''package_variety_rating'', 4.0,
        ''customer_support_rating'', 3.5,
        ''value_for_money_rating'', 4.25,
        ''avg_data_speed_mbps'', 14,
        ''peak_hour_speed_mbps'', 8,
        ''network_uptime_percent'', 97.5,
        ''latency_ms'', 25,
        ''urban_coverage_percent'', 99.0,
        ''rural_coverage_percent'', 60.0,
        ''hotline_number'', ''111'',
        ''website_url'', ''https://www.banglalink.com.bd'',
        ''verdict'', ''Highly Recommended - Best Value for Money'',
        ''recommendation_score'', 8
    ),
    updated_at = CURRENT_TIMESTAMP
WHERE id = 240 AND product_id = (SELECT id FROM products WHERE slug = ''banglalink'');

-- ============================================================================
-- UPDATE BENGALI TRANSLATION (ID: 240)
-- ============================================================================

UPDATE product_review_translations
SET
    reviews = '<!DOCTYPE html>
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
            <p><strong>রেটিং: ৪.०/৫</strong></p>
            <ul>
                <li>শহরের কভারেজ: ৯৯% ধারাবাহিক ৪জি উপলব্ধতা</li>
                <li>গ্রামীণ কভারেজ: ৬०% এবং ক্রমাগত সম্প্রসারণ</li>
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
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१नম्बर</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">४.६५/५ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">সেরা মান</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">প্রিমিয়াম নেটওয়ার্ক, १४% বেশি ব্যয়বহুল</td>
                    </tr>
                    <tr style="background-color: #fff3cd;">
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>বাংলালিংক</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">२नम्बर</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>४.१/५ ⭐</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>সেরা মূল্য</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>বর্তমান - সেরা মূল্য-থেকে-মান অনুপাত</strong></td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>রোবি</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">३नम्बर</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">४.०५/५ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">ভাল বিকল্প</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">অনুরূপ মান, সামান্য সাশ্রয়ী ডেটা</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>এয়ারটেল</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">४नम्बर</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">३.०५/५ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">আঞ্চলিক বিকল্প</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">ক্রমবর্ধমান নেটওয়ার্ক, সীমিত কভারেজ</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>টেলিটক</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">५नम्बर</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">२.८५/५ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">শুধু বাজেট</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">সবচেয়ে সস্তা কিন্তু পুরানো অবকাঠামো</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pricing">
            <h2>মূল্য তুলনা</h2>
            <p><strong>রেটিং: ४.५/५ (সবচেয়ে প্রতিযোগিতামূলক)</strong></p>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">সেবা</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">বাংলালিংক</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">গ্রামীণফোন</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">রোবি</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">এয়ারটেল</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">টেলিটক</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>១००एमবি दैनिक</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">५० টাকা ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">८० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">६० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">५० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">४० টাকা</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>१जिबी साप्ताहिक</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">१२० টাকা ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१८० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१४० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१३० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१०० টাকা</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>सीमाहीन मासिक</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">९९९ টাকা ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१५०० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१२०० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">९९९ টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">८०० টাকা</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="packages">
            <h2>প্যাকেজ এবং বান্ডেল</h2>
            <p><strong>রেটিং: ४.०/५ (ভাল বৈচিত্র্য)</strong></p>
            <ul>
                <li><strong>ডেটা সেবা:</strong> २०+ দৈনিক, সাপ্তাহিক এবং মাসিক বান্ডেল</li>
                <li><strong>সোশ্যাল মিডিয়া:</strong> হোয়াটসঅ্যাপ, ফেসবুক, ইউটিউব বান্ডেল</li>
                <li><strong>ভয়েস কল:</strong> স্থানীয়, নেটওয়ার্ক-অন এবং আন্তর্জাতিক বিকল্প</li>
                <li><strong>ওটিটি সেবা:</strong> নেটফ্লিক্স, ইউটিউব, স্ট্রিমিং প্যাকেজ</li>
                <li><strong>রাত্রি প্যাকেজ:</strong> রাত १२টা - ६টায় ছাড়যুক্ত হার</li>
                <li><strong>রোমিং:</strong> ५०+ দেশে ডে পাস</li>
            </ul>
        </section>

        <section class="digital-services">
            <h2>ডিজিটাল সেবা</h2>
            <p><strong>রেটিং: ४.०/५</strong></p>
            <ul>
                <li><strong>বাংলালিংক মানি:</strong> २ মিলিয়ন+ ব্যবহারকারীর মোবাইল ওয়ালেট</li>
                <li><strong>অর্থ স্থানান্তর:</strong> অন্যান্য ব্যবহারকারীদের দ্রুত পি२পি স্থানান্তর</li>
                <li><strong>বিল পেমেন্ট:</strong> বিদ্যুৎ, জল, ইন্টারনেট বিল</li>
                <li><strong>বণিক পেমেন্ট:</strong> ব্যবসায়িক পেমেন্ট সমাধান</li>
                <li><strong>সহজ রিচার্জ:</strong> বিকাশ সহ একাধিক পদ্ধতি</li>
            </ul>
        </section>

        <section class="customer-support">
            <h2>গ্রাহক সহায়তা</h2>
            <p><strong>রেটিং: ३.५/५</strong></p>
            <ul>
                <li>হটলাইন: १११ (२४/७ উপলব्ध)</li>
                <li>মোবাইল অ্যাপ: উচ্চ-রেটেড ব্যবহারকারী ইন্টারফেস</li>
                <li>ইউএসএসডি কোড: দ্রুত ব্যালেন্স এবং ব্যবহার পরীক্ষা (*१२३#)</li>
                <li>সোশ্যাল মিডিয়া: ফেসবুক এবং টুইটার সহায়তা</li>
                <li><em>সমস্যা:</em> পিক আওয়ার्সে দীর्ঘ অপেক्ষা সময়</li>
            </ul>
        </section>

        <section class="pros">
            <h2>সুবিধা (११টি প्रধान সুবিधा)</h2>
            <ul>
                <li>✅ <strong>সাশ्रय়ী কল হার</strong> - स्थानीय और आंतर्राष्ट्रीय कॉलের लिए प्रतिस्पर्धी मूल्य</li>
                <li>✅ <strong>আকর্ষণীয় ডেটা বান্ডেল</strong> - নিয়মিত প্রচার এবং বোনাস অফার</li>
                <li>✅ <strong>ভাল ४जि कभरेज</strong> - ९९% शहरी और ६०% ग्रामीण कभरेज</li>
                <li>✅ <strong>ব्যবহारকारী-বান्ধব অ্যাপ</strong> - স्वज्ञानी অ्যাकाউंट ব्यवস्थापना ইন्টারফেস</li>
                <li>✅ <strong>ডিजिটल পेমেন্ट সহায়তা</strong> - একাधिक पेमेंट प्लेटфॉर्मের साथ integrated</li>
                <li>✅ <strong>জরুरी ব्যালেন्स</strong> - ব্যালেন্স कम থাকলে জরুरी कল के लिए फ़ीচर</li>
                <li>✅ <strong>नमনीय पরিকल्पনা</strong> - প्রিपेইड और पोस्टपेड बिकल्प उपलब्ध</li>
                <li>✅ <strong>नियमित प्रचার</strong> - নिয়মিত बिशेष অফर এবং बोनस</li>
                <li>✅ <strong>सोशल मीडिया বान্ডেল</strong> - जनप्रिय प्लेटفॉर्मগুলिতে सीमाहीन অ্যাক्সेস</li>
                <li>✅ <strong>অর্थ সেवा</strong> - বিस्तृत डिजिटल पेमेंट ইকोसিস्टেम</li>
                <li>✅ <strong>एकाधिक सहayता চ্যানেল</strong> - ফোন, অ্যাপ, ইউএসএসডি এবং সোশ্যাल মিডিয়া</li>
            </ul>
        </section>

        <section class="cons">
            <h2>অসুবिधा (७টি প্रधान सीमাबद्धता)</h2>
            <ul>
                <li>⚠️ <strong>पीक আওয়ার ভीड়</strong> - संध्या ६-१० टा मेँ नेटवर्क ধीrगতि</li>
                <li>⚠️ <strong>सीमित ग्रामीण কভरেজ</strong> - গ्रामे शুধु ६०% कभरेज</li>
                <li>⚠️ <strong>দीर्ঘ গ्राहক অপेक्ষा सময়</strong> - पीक समय में सहায্यता सरणी</li>
                <li>⚠️ <strong>उच्चตर रোমिंग हार</strong> - আন्तর्জাতिকভাবে প्रতिद्वંद्वીدের চेये बেशи ব्यय्बहुल</li>
                <li>⚠️ <strong>गति थ्रॉटलिंग</strong> - कुछ प्याकेजमेँ मीजे मीजे सीमaबद्धता</li>
                <li>⚠️ <strong>जटिल सिम सक्रिय्करण</strong> - केЕयूसी प्रक्रिया सहल कर्या जे सकता है</li>
                <li>⚠️ <strong>सीमाबद्ध बैधता</strong> - कुछ प्याकेजের कठोर समाप्तीक तारीख आहै</li>
            </ul>
        </section>

        <section class="final-assessment">
            <h2>চূড়ान्ত মूल्यायन</h2>
            <p>बाংলালिंक बाংলादेशीर मोबाइल बाजारे सेरे मूल्य-थॉ-मान अनुपात प्रदान करे। ३५+ मিलिјन ग्राहक साथ द्वितीय बृहत्तम अपरेटोर हिसाबे, एTা शহुरे एलाकार मेँ निर्भरयोग्य ४जि कभरेज, ग्रामीणफोनेर चेये १४% कम मूल्य अंद एक क्रमबर्धमान डिজिटल सेवा इकोसिस्टेम प्रदान करे। यधिपि एटार गर्मीण कभरेज अंद माझे माझे पीक-आवर भीड़ेर सीमाbद्धता रायहै, तबे एग्गुली अर्थसाधक प्याकेज अंद ब्यवहारकारी-बान्धव सेवा दारा पूरण हायै। बेशिरभाग ब्यवहारकारीर जानयॆ, बाङ्गलालिंक मान अंद खर्चेर मिडझ आदर्श भारसमय प्रतिनिधित्व करे।</p>
        </section>
    </article>
</body>
</html>',
    additional_details = jsonb_build_object(
        ''language'', ''Bengali'',
        ''market_rank'', 2,
        ''subscribers_millions'', 35,
        ''overall_rating'', 4.1,
        ''verdict'', ''অত्যন्ত সুপারिशকৃत - সেরे मूल्येर मूल्य'',
        ''recommendation_score'', 8
    ),
    updated_at = CURRENT_TIMESTAMP
WHERE product_review_id = 240 AND locale = ''bn'';

-- ============================================================================
-- VERIFICATION
-- ============================================================================

SELECT 
    id,
    product_id,
    rating,
    LENGTH(reviews) as review_length,
    updated_at
FROM product_reviews
WHERE id = 240;

SELECT 
    id,
    product_review_id,
    locale,
    LENGTH(reviews) as review_length,
    updated_at
FROM product_review_translations
WHERE product_review_id = 240 AND locale = 'bn';

-- ============================================================================
-- COMPLETE HTML STRUCTURE WITH PROPER FORMATTING
-- ============================================================================
