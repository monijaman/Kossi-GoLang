-- ============================================================================
-- BANGLALINK COMPARISON - REPORTING QUERIES
-- ============================================================================
-- Use these queries to generate reports and insights
-- ============================================================================

-- ============================================================================
-- 1. MARKET OVERVIEW REPORTS
-- ============================================================================

-- Report 1: Market Ranking with All Metrics
SELECT 
    market_rank,
    operator_name,
    subscribers_millions,
    overall_rating,
    recommendation_score,
    network_quality_rating,
    coverage_rating,
    data_speed_rating,
    pricing_rating,
    package_variety_rating,
    customer_support_rating,
    value_for_money_rating
FROM mobile_operator_comparisons
ORDER BY market_rank;

-- Report 2: Rating Comparison Table
SELECT 
    operator_name as "Operator",
    CONCAT(overall_rating, ' ⭐') as "Overall",
    CONCAT(network_quality_rating, ' ⭐') as "Quality",
    CONCAT(coverage_rating, ' ⭐') as "Coverage",
    CONCAT(data_speed_rating, ' ⭐') as "Speed",
    CONCAT(pricing_rating, ' ⭐') as "Price",
    CONCAT(value_for_money_rating, ' ⭐') as "Value"
FROM mobile_operator_comparisons
ORDER BY overall_rating DESC;

-- Report 3: Technical Specifications Comparison
SELECT 
    operator_name as "Operator",
    avg_data_speed_mbps as "Avg Speed (Mbps)",
    peak_hour_speed_mbps as "Peak Speed (Mbps)",
    network_uptime_percent as "Uptime %",
    latency_ms as "Latency (ms)",
    urban_coverage_percent as "Urban %",
    rural_coverage_percent as "Rural %"
FROM mobile_operator_comparisons
ORDER BY avg_data_speed_mbps DESC;

-- ============================================================================
-- 2. PRICING ANALYSIS REPORTS
-- ============================================================================

-- Report 4: Complete Pricing by Operator and Service
SELECT 
    operator_name as "Operator",
    service_type as "Service",
    package_name as "Package",
    volume as "Volume",
    duration as "Duration",
    CONCAT(price_bdt, ' BDT') as "Price",
    notes as "Notes"
FROM operator_pricing_comparison
ORDER BY operator_name, service_type, price_bdt;

-- Report 5: Cheapest Options per Service Type
SELECT 
    service_type as "Service Type",
    operator_name as "Cheapest Option",
    package_name as "Package",
    volume as "Volume",
    CONCAT(price_bdt, ' BDT') as "Price",
    duration as "Duration"
FROM (
    SELECT 
        *,
        ROW_NUMBER() OVER (PARTITION BY service_type ORDER BY price_bdt) as rn
    FROM operator_pricing_comparison
)
WHERE rn = 1
ORDER BY service_type;

-- Report 6: Data Package Price Comparison
SELECT 
    operator_name as "Operator",
    package_name as "Package",
    volume as "Data Volume",
    CONCAT(price_bdt, ' BDT') as "Price",
    duration as "Valid For",
    ROUND(price_bdt::NUMERIC / EXTRACT(DAY FROM interval '1 day' * CASE 
        WHEN duration = '1 Day' THEN 1 
        WHEN duration = '7 Days' THEN 7
        WHEN duration = '30 Days' THEN 30
        ELSE 1 
    END), 2) as "Price per Day"
FROM operator_pricing_comparison
WHERE service_type = 'data'
ORDER BY operator_name, price_bdt;

-- Report 7: Voice & SMS Pricing Summary
SELECT 
    operator_name as "Operator",
    service_type as "Type",
    package_name as "Package",
    volume as "Volume",
    CONCAT(price_bdt, ' BDT') as "Price"
FROM operator_pricing_comparison
WHERE service_type IN ('voice', 'sms')
ORDER BY operator_name, service_type;

-- ============================================================================
-- 3. FEATURE ANALYSIS REPORTS
-- ============================================================================

-- Report 8: Features by Operator
SELECT 
    operator_name as "Operator",
    feature_category as "Category",
    feature_name as "Feature",
    details as "Details"
FROM operator_features
WHERE is_available = true
ORDER BY operator_name, feature_category, feature_name;

-- Report 9: Feature Availability Summary
SELECT 
    operator_name as "Operator",
    COUNT(DISTINCT feature_category) as "Categories",
    COUNT(*) as "Total Features",
    COUNT(CASE WHEN feature_category = 'data' THEN 1 END) as "Data Features",
    COUNT(CASE WHEN feature_category = 'social_media' THEN 1 END) as "Social Media",
    COUNT(CASE WHEN feature_category = 'voice' THEN 1 END) as "Voice Features",
    COUNT(CASE WHEN feature_category = 'digital' THEN 1 END) as "Digital Services",
    COUNT(CASE WHEN feature_category = 'support' THEN 1 END) as "Support Options"
FROM operator_features
WHERE is_available = true
GROUP BY operator_name
ORDER BY total_features DESC;

-- Report 10: Feature Comparison Matrix (Wide View)
WITH feature_pivot AS (
    SELECT 
        feature_category,
        feature_name,
        SUM(CASE WHEN operator_name = 'Banglalink' AND is_available THEN 1 ELSE 0 END) as banglalink,
        SUM(CASE WHEN operator_name = 'Grameenphone' AND is_available THEN 1 ELSE 0 END) as grameenphone,
        SUM(CASE WHEN operator_name = 'Robi' AND is_available THEN 1 ELSE 0 END) as robi,
        SUM(CASE WHEN operator_name = 'Airtel' AND is_available THEN 1 ELSE 0 END) as airtel,
        SUM(CASE WHEN operator_name = 'Teletalk' AND is_available THEN 1 ELSE 0 END) as teletalk
    FROM operator_features
    GROUP BY feature_category, feature_name
)
SELECT 
    feature_category as "Category",
    feature_name as "Feature",
    CASE WHEN banglalink > 0 THEN '✓' ELSE '✗' END as "Banglalink",
    CASE WHEN grameenphone > 0 THEN '✓' ELSE '✗' END as "Grameenphone",
    CASE WHEN robi > 0 THEN '✓' ELSE '✗' END as "Robi",
    CASE WHEN airtel > 0 THEN '✓' ELSE '✗' END as "Airtel",
    CASE WHEN teletalk > 0 THEN '✓' ELSE '✗' END as "Teletalk"
FROM feature_pivot
ORDER BY feature_category, feature_name;

-- ============================================================================
-- 4. PROS & CONS ANALYSIS REPORTS
-- ============================================================================

-- Report 11: Pros by Operator (Ranked)
SELECT 
    operator_name as "Operator",
    item_text as "Advantage",
    priority,
    CASE priority WHEN 1 THEN '★★★ Major' WHEN 2 THEN '★★ Notable' ELSE '★ Minor' END as "Importance"
FROM operator_pros_cons
WHERE type = 'pro'
ORDER BY operator_name, priority, item_text;

-- Report 12: Cons by Operator (Ranked)
SELECT 
    operator_name as "Operator",
    item_text as "Limitation",
    priority,
    CASE priority WHEN 1 THEN '★★★ Serious' WHEN 2 THEN '★★ Notable' ELSE '★ Minor' END as "Severity"
FROM operator_pros_cons
WHERE type = 'con'
ORDER BY operator_name, priority, item_text;

-- Report 13: Operator Strength/Weakness Summary
WITH strength_analysis AS (
    SELECT 
        operator_name,
        COUNT(CASE WHEN type = 'pro' AND priority = 1 THEN 1 END) as major_strengths,
        COUNT(CASE WHEN type = 'pro' THEN 1 END) as total_strengths,
        COUNT(CASE WHEN type = 'con' AND priority = 1 THEN 1 END) as major_weaknesses,
        COUNT(CASE WHEN type = 'con' THEN 1 END) as total_weaknesses
    FROM operator_pros_cons
    GROUP BY operator_name
)
SELECT 
    operator_name as "Operator",
    major_strengths as "Major Strengths",
    total_strengths as "Total Pros",
    major_weaknesses as "Major Weaknesses",
    total_weaknesses as "Total Cons",
    ROUND((total_strengths - total_weaknesses)::NUMERIC / (total_strengths + total_weaknesses) * 100, 1) as "Score Ratio %"
FROM strength_analysis
ORDER BY major_strengths DESC;

-- ============================================================================
-- 5. DECISION-MAKING REPORTS
-- ============================================================================

-- Report 14: Best Operator by Use Case
SELECT 
    'Best Overall' as "Scenario",
    operator_name as "Recommendation",
    overall_rating as "Rating",
    recommendation_text as "Why"
FROM mobile_operator_comparisons
WHERE overall_rating = (SELECT MAX(overall_rating) FROM mobile_operator_comparisons)
UNION ALL
SELECT 
    'Best for Students',
    operator_name,
    overall_rating,
    recommendation_text
FROM mobile_operator_comparisons
WHERE best_for_students = true
UNION ALL
SELECT 
    'Best for Professionals',
    operator_name,
    overall_rating,
    recommendation_text
FROM mobile_operator_comparisons
WHERE best_for_professionals = true
UNION ALL
SELECT 
    'Best for Budget Users',
    operator_name,
    overall_rating,
    recommendation_text
FROM mobile_operator_comparisons
WHERE best_for_budget_users = true
UNION ALL
SELECT 
    'Best for Heavy Users',
    operator_name,
    overall_rating,
    recommendation_text
FROM mobile_operator_comparisons
WHERE best_for_heavy_users = true
UNION ALL
SELECT 
    'Best for Business',
    operator_name,
    overall_rating,
    recommendation_text
FROM mobile_operator_comparisons
WHERE best_for_business = true;

-- Report 15: Quality vs Price Analysis
SELECT 
    operator_name as "Operator",
    overall_rating as "Overall Rating",
    pricing_rating as "Price Rating (higher = cheaper)",
    value_for_money_rating as "Value Rating",
    CASE 
        WHEN overall_rating > 4.2 AND pricing_rating > 3.5 THEN 'Excellent Value'
        WHEN overall_rating > 4.0 AND pricing_rating > 3.0 THEN 'Good Value'
        WHEN overall_rating > 3.5 AND pricing_rating > 2.5 THEN 'Fair Value'
        ELSE 'Poor Value'
    END as "Assessment"
FROM mobile_operator_comparisons
ORDER BY value_for_money_rating DESC;

-- Report 16: Customer Service & Support Comparison
SELECT 
    operator_name as "Operator",
    customer_support_rating as "Support Rating",
    hotline_number as "Hotline",
    website_url as "Website",
    COUNT(*) as support_channels
FROM mobile_operator_comparisons
LEFT JOIN operator_features ON mobile_operator_comparisons.operator_name = operator_features.operator_name
WHERE operator_features.feature_category = 'support'
GROUP BY operator_name, customer_support_rating, hotline_number, website_url
ORDER BY customer_support_rating DESC;

-- ============================================================================
-- 6. MARKET ANALYSIS REPORTS
-- ============================================================================

-- Report 17: Market Share and Position
SELECT 
    market_rank as "Position",
    operator_name as "Operator",
    subscribers_millions as "Subscribers (M)",
    ROUND(100.0 * subscribers_millions / SUM(subscribers_millions) OVER(), 1) as "Market Share %",
    overall_rating as "Rating",
    recommendation_score as "Recommendation (0-10)"
FROM mobile_operator_comparisons
ORDER BY market_rank;

-- Report 18: Network Quality Tier Analysis
SELECT 
    CASE 
        WHEN overall_rating >= 4.5 THEN 'Premium Tier'
        WHEN overall_rating >= 4.0 THEN 'Mid-Premium'
        WHEN overall_rating >= 3.5 THEN 'Mid-Range'
        WHEN overall_rating >= 3.0 THEN 'Budget'
        ELSE 'Economy'
    END as "Quality Tier",
    operator_name as "Operator",
    overall_rating as "Rating",
    subscribers_millions as "Subscribers (M)",
    avg_data_speed_mbps as "Avg Speed (Mbps)"
FROM mobile_operator_comparisons
ORDER BY overall_rating DESC;

-- ============================================================================
-- 7. CUSTOM COMPARISON QUERIES
-- ============================================================================

-- Query 1: Compare Any Two Operators
-- Usage: Replace 'Banglalink' and 'Grameenphone' with any operator names
SELECT 
    'Metric' as metric,
    'Banglalink' as operator1,
    'Grameenphone' as operator2
UNION ALL
SELECT 
    'Overall Rating',
    (SELECT CONCAT(overall_rating, ' ⭐') FROM mobile_operator_comparisons WHERE operator_name = 'Banglalink'),
    (SELECT CONCAT(overall_rating, ' ⭐') FROM mobile_operator_comparisons WHERE operator_name = 'Grameenphone')
UNION ALL
SELECT 
    'Market Rank',
    (SELECT CONCAT('#', market_rank) FROM mobile_operator_comparisons WHERE operator_name = 'Banglalink'),
    (SELECT CONCAT('#', market_rank) FROM mobile_operator_comparisons WHERE operator_name = 'Grameenphone')
UNION ALL
SELECT 
    'Avg Data Speed',
    (SELECT CONCAT(avg_data_speed_mbps, ' Mbps') FROM mobile_operator_comparisons WHERE operator_name = 'Banglalink'),
    (SELECT CONCAT(avg_data_speed_mbps, ' Mbps') FROM mobile_operator_comparisons WHERE operator_name = 'Grameenphone')
UNION ALL
SELECT 
    'Coverage (Urban)',
    (SELECT CONCAT(urban_coverage_percent, '%') FROM mobile_operator_comparisons WHERE operator_name = 'Banglalink'),
    (SELECT CONCAT(urban_coverage_percent, '%') FROM mobile_operator_comparisons WHERE operator_name = 'Grameenphone')
UNION ALL
SELECT 
    'Network Uptime',
    (SELECT CONCAT(network_uptime_percent, '%') FROM mobile_operator_comparisons WHERE operator_name = 'Banglalink'),
    (SELECT CONCAT(network_uptime_percent, '%') FROM mobile_operator_comparisons WHERE operator_name = 'Grameenphone')
UNION ALL
SELECT 
    'Cheapest Data Plan',
    (SELECT CONCAT(MIN(price_bdt), ' BDT') FROM operator_pricing_comparison WHERE operator_name = 'Banglalink' AND service_type = 'data'),
    (SELECT CONCAT(MIN(price_bdt), ' BDT') FROM operator_pricing_comparison WHERE operator_name = 'Grameenphone' AND service_type = 'data');

-- Query 2: Find Similar Operators
SELECT DISTINCT
    o1.operator_name as "Operator 1",
    o2.operator_name as "Operator 2",
    ABS(o1.overall_rating - o2.overall_rating) as rating_difference,
    o1.overall_rating,
    o2.overall_rating
FROM mobile_operator_comparisons o1
JOIN mobile_operator_comparisons o2 ON ABS(o1.overall_rating - o2.overall_rating) < 0.5
WHERE o1.operator_name < o2.operator_name
ORDER BY rating_difference;

-- Query 3: Package Variety by Operator
SELECT 
    operator_name as "Operator",
    service_type as "Service Type",
    COUNT(DISTINCT package_name) as "Package Count",
    COUNT(DISTINCT volume) as "Volume Options",
    MIN(price_bdt) as "Min Price (BDT)",
    MAX(price_bdt) as "Max Price (BDT)",
    AVG(price_bdt)::INT as "Avg Price (BDT)"
FROM operator_pricing_comparison
GROUP BY operator_name, service_type
ORDER BY operator_name, service_type;

-- ============================================================================
-- END OF REPORTING QUERIES
-- ============================================================================
