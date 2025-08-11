# Laravel-Compatible Product API

## Overview

The Go server now supports Laravel-compatible product filtering that matches the exact API format used in the original Laravel PHP application. This ensures seamless frontend integration.

## Endpoint

**GET** `/products`

The endpoint automatically detects the request format:

- **Laravel format**: When `locale`, `page`, or `sortby` parameters are present
- **Legacy format**: When using `limit`, `offset` parameters without Laravel-specific ones

## Laravel-Compatible Parameters

| Parameter    | Type    | Description                    | Example                              |
| ------------ | ------- | ------------------------------ | ------------------------------------ |
| `locale`     | string  | Language locale                | `en`, `tr`                           |
| `page`       | integer | Page number (1-based)          | `1`, `2`, `3`                        |
| `limit`      | integer | Items per page                 | `10`, `20`, `50`                     |
| `category`   | string  | Category slug for filtering    | `smartphones`, `laptops`             |
| `brand`      | string  | Brand slug(s), comma-separated | `apple`, `samsung,apple`             |
| `priceRange` | string  | Price range filter             | `100-500`, `1000-`                   |
| `searchterm` | string  | Search in name/description     | `iPhone`, `laptop`                   |
| `sortby`     | string  | Sort order                     | `popular`, `price_asc`, `price_desc` |

## Response Format

### Laravel-Compatible Response

```json
{
  "data": [
    {
      "id": 1,
      "name": "iPhone 15 Pro",
      "description": "Latest iPhone",
      "slug": "iphone-15-pro",
      "price": 1199.99,
      "category_id": 1,
      "brand_id": 1,
      "views_count": 5,
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
    }
  ],
  "meta": {
    "current_page": 1,
    "per_page": 10,
    "total": 25,
    "last_page": 3,
    "from": 1,
    "to": 10,
    "has_next_page": true,
    "has_prev_page": false
  },
  "filters": {
    "locale": "en",
    "category": "smartphones",
    "brand": "apple",
    "price_range": "1000-1500",
    "search_term": "iPhone",
    "sort_by": "price_desc"
  }
}
```

### Legacy Response (Backward Compatibility)

```json
{
  "products": [...],
  "count": 10,
  "limit": 10,
  "offset": 0,
  "total": 25
}
```

## Example Requests

### Basic Laravel-Compatible Request

```bash
curl "http://localhost:8080/products?locale=en&page=1&limit=10"
```

### Filtered Request

```bash
curl "http://localhost:8080/products?locale=en&page=1&limit=10&category=smartphones&brand=apple&searchterm=iPhone&sortby=price_desc"
```

### Legacy Request (Backward Compatibility)

```bash
curl "http://localhost:8080/products?limit=10&offset=0"
```

## Sorting Options

| Sort Value   | Description                              |
| ------------ | ---------------------------------------- |
| `popular`    | Sort by views count (most viewed first)  |
| `price_asc`  | Sort by price ascending (lowest first)   |
| `price_desc` | Sort by price descending (highest first) |
| (empty)      | Default order (by ID)                    |

## Implementation Details

### Handler Function

- `GetFilteredProductsHandler` - New Laravel-compatible handler
- `ListProductsHandler` - Legacy handler for backward compatibility

### Repository Method

- `GetWithFilters(ctx, filters)` - New method supporting complex filtering

### Route Logic

The `/products` endpoint automatically routes to the appropriate handler based on query parameters:

```go
if r.URL.Query().Get("locale") != "" || r.URL.Query().Get("page") != "" || r.URL.Query().Get("sortby") != "" {
    GetFilteredProductsHandler(w, r, productRepo)
} else {
    ListProductsHandler(w, r, productRepo)
}
```

## Database Features

### Filtering Capabilities

- **Text Search**: ILIKE search across product names and descriptions
- **Category Filtering**: JOIN with categories table using slug
- **Brand Filtering**: Support for multiple brand slugs (comma-separated)
- **Price Range**: Support for min-max price filtering (format: "min-max" or "min-")

### Sorting Implementation

- **Popular**: ORDER BY views_count DESC, id ASC
- **Price Ascending**: ORDER BY price ASC, id ASC
- **Price Descending**: ORDER BY price DESC, id ASC
- **Default**: ORDER BY id ASC

## Postman Collection

A new request "Laravel Compatible Products (Filtered)" has been added to the Postman collection with all the Laravel-compatible parameters pre-configured for easy testing.
