# Bug Fix: Incorrect Error Handling in CreateReview

## The Problem

The original code had backwards logic:

```go
existingReview, err := repo.GetByProductAndUser(ctx, productID, userID)
if err == nil && existingReview != nil {
    return nil, errors.New("user has already reviewed this product")
}
```

**Why it was wrong:**

- When **NO review exists**: GORM returns `err = gorm.ErrRecordNotFound` (err is NOT nil)
- When **review exists**: GORM returns `err = nil` and `existingReview != nil`

The condition `if err == nil` means "if there was NO error", which only happens when a review IS found!

But the code was checking `err == nil` expecting it to mean "no review found" - which is backwards! 🔄

## The Symptoms

You were getting "user has already reviewed this product" error even when:

- Product ID 3 exists ✓
- User ID 1 exists ✓
- **No review exists for product 3 by user 1** ✗

The error was confusing because the logic was inverted.

## The Fix

```go
existingReview, err := repo.GetByProductAndUser(ctx, productID, userID)
if err != nil && err.Error() != "record not found" {
    // Some other database error occurred
    return nil, err
}
if existingReview != nil {
    // Review was found - user already reviewed this product
    return nil, errors.New("user has already reviewed this product")
}
// No review found - proceed with creation
```

**Now it works correctly:**

1. If there's an error (other than "not found") → return the error
2. If `existingReview != nil` → review exists, return duplicate error
3. Otherwise → no review exists, continue creating new review

## Testing

After restarting the server, you should now be able to:

- ✅ Create review for product ID 3 (if it exists and user hasn't reviewed it)
- ✅ Get proper "product not found" error for product ID 2
- ✅ Get "already reviewed" error only when user HAS reviewed that product

## Root Cause

GORM's `.First()` method behavior:

- **Record found**: Returns `err = nil` and populates the model
- **Record NOT found**: Returns `err = gorm.ErrRecordNotFound`
- **Other DB errors**: Returns the specific error

The original code misunderstood this pattern and checked for the opposite condition!
