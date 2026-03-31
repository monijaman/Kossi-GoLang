# User Role System Implementation

## Overview
The authentication system now supports three user types/roles:
- **admin**: Full access to all features, including delete operations
- **reviewer**: Can review and moderate content, but cannot delete
- **guest**: Limited access, can only view their own profile and saved items

## What Was Implemented

### 1. Database Schema
The `users` table has a `type` column:
```sql
type VARCHAR(50) NOT NULL DEFAULT 'guest'
```

Valid values: `admin`, `reviewer`, `guest`

### 2. Backend Changes

#### User Entity (`internal/domain/entities/core.go`)
```go
type User struct {
    ID              uint
    Name            string
    Email           string
    EmailVerifiedAt *time.Time
    Password        string
    RememberToken   *string
    Type            string  // <-- Role field
    CreatedAt       time.Time
    UpdatedAt       time.Time
}
```

#### Registration
Users can now register with a specific type:
```bash
POST /api/register
{
  "username": "admin1",
  "email": "admin1@example.com",
  "password": "password123",
  "type": "admin"  // Optional, defaults to "guest"
}
```

#### Login Response
Login now returns the user's type:
```json
{
  "message": "Login successful",
  "email": "admin99@example.com",
  "name": "admin99",
  "type": "admin",
  "token": "eyJhbGc...",
  "refresh_token": "eyJhbGc..."
}
```

### 3. Role-Based Authorization Middleware

Created `internal/interface/middleware/role.go` with:

#### `RequireRole` Middleware
Protects endpoints by checking user roles:
```go
middleware.RequireRole(userRepo, "admin")(handlerFunc)
// or allow multiple roles:
middleware.RequireRole(userRepo, "admin", "reviewer")(handlerFunc)
```

#### Helper Functions
- `GetUserTypeFromContext(r *http.Request)`: Extract user type from request context

### 4. Example Admin Routes

Created `internal/interface/handler/admin/routes.go` with example protected endpoints:

#### Admin-Only Endpoint
```go
// DELETE /admin/products/{id} - Only admins can delete
mux.HandleFunc("/admin/products/", 
    middleware.JWTAuthMiddleware(
        middleware.RequireRole(userRepo, "admin")(deleteProductHandler)
    )
)
```

#### Admin Stats Endpoint
```go
// GET /admin/stats - Only admins
mux.HandleFunc("/admin/stats", 
    middleware.JWTAuthMiddleware(
        middleware.RequireRole(userRepo, "admin")(statsHandler)
    )
)
```

#### Reviewer Endpoint
```go
// GET /admin/reviews/pending - Admins and reviewers
mux.HandleFunc("/admin/reviews/pending", 
    middleware.JWTAuthMiddleware(
        middleware.RequireRole(userRepo, "admin", "reviewer")(pendingReviewsHandler)
    )
)
```

## How to Use

### 1. Register Users with Roles
```bash
# Register an admin
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin1",
    "email": "admin1@example.com",
    "password": "password123",
    "type": "admin"
  }'

# Register a reviewer
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "reviewer1",
    "email": "reviewer1@example.com",
    "password": "password123",
    "type": "reviewer"
  }'

# Register a guest (default)
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "guest1",
    "email": "guest1@example.com",
    "password": "password123"
  }'
```

### 2. Login and Get Token
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin1@example.com",
    "password": "password123"
  }'

# Response includes type:
# {
#   "email": "admin1@example.com",
#   "name": "admin1",
#   "type": "admin",
#   "token": "eyJhbGc...",
#   "refresh_token": "eyJhbGc..."
# }
```

### 3. Access Protected Endpoints
```bash
# Admin accessing admin-only endpoint
curl -X GET http://localhost:8080/admin/stats \
  -H "Authorization: Bearer <admin_token>"

# Reviewer trying to access admin-only endpoint (will be denied)
curl -X GET http://localhost:8080/admin/stats \
  -H "Authorization: Bearer <reviewer_token>"
# Response: {"error": "Access denied. Required role: admin", "your_role": "reviewer"}
```

### 4. Protect Your Own Endpoints

To add role protection to any endpoint:

```go
// Example: Delete product (admin only)
mux.HandleFunc("/products/delete", 
    middleware.JWTAuthMiddleware(
        middleware.RequireRole(userRepo, "admin")(func(w http.ResponseWriter, r *http.Request) {
            // Your delete logic here
            // Only admins will reach this code
        })
    )
)

// Example: Moderate reviews (admin and reviewer)
mux.HandleFunc("/reviews/moderate", 
    middleware.JWTAuthMiddleware(
        middleware.RequireRole(userRepo, "admin", "reviewer")(func(w http.ResponseWriter, r *http.Request) {
            // Your moderation logic here
            // Admins and reviewers can access this
        })
    )
)
```

## Frontend Integration

### Store User Type on Login
```typescript
// After successful login
const response = await fetch('/api/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ email, password })
});

const data = await response.json();

// Store in localStorage or state management
localStorage.setItem('userType', data.type);
localStorage.setItem('token', data.token);
```

### Conditional UI Rendering
```tsx
const userType = localStorage.getItem('userType');

// Show delete button only for admins
{userType === 'admin' && (
  <button onClick={handleDelete}>Delete</button>
)}

// Show edit button for admins and reviewers
{(userType === 'admin' || userType === 'reviewer') && (
  <button onClick={handleEdit}>Edit</button>
)}
```

## Testing

### Test Admin Access
```bash
# 1. Login as admin
ADMIN_TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin99@example.com","password":"password123"}' \
  | grep -o '"token":"[^"]*' | cut -d'"' -f4)

# 2. Access admin endpoint
curl -X GET http://localhost:8080/admin/stats \
  -H "Authorization: Bearer $ADMIN_TOKEN"
```

### Test Reviewer Access (Should Fail)
```bash
# 1. Login as reviewer
REVIEWER_TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"rev1@example.com","password":"password123"}' \
  | grep -o '"token":"[^"]*' | cut -d'"' -f4)

# 2. Try to access admin endpoint (will be denied)
curl -X GET http://localhost:8080/admin/stats \
  -H "Authorization: Bearer $REVIEWER_TOKEN"
```

## Next Steps

1. **Add Role Checks to Existing Endpoints**: Go through your existing handlers and add `RequireRole` middleware where needed
2. **Frontend Implementation**: Update your React components to show/hide features based on user type
3. **Database Migration**: If you have existing users, you may want to update their types:
   ```sql
   UPDATE users SET type = 'admin' WHERE email = 'your-admin@example.com';
   ```

## Files Modified/Created

### Created:
- `internal/interface/middleware/role.go` - Role-based authorization middleware
- `internal/interface/handler/admin/routes.go` - Example admin routes

### Modified:
- `internal/usecase/auth/login.go` - Return user entity from Login
- `internal/interface/handler/auth.go` - Include type in login response
- `internal/interface/handler/auth/handler.go` - Include type in login response and accept type in registration
- `cmd/app/main.go` - Register admin routes
