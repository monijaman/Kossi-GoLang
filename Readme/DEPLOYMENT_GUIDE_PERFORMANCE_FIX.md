# Quick Deployment Guide - Banking Category Performance Fix

## TL;DR

The banking category page timeout was caused by **N+1 database queries**. This fix reduces queries from ~60 to ~3, making the page load in <1 second instead of timing out.

## What Changed

✅ **Backend optimization only** - No frontend changes needed
✅ **Database indexes added** - Automatic during migration
✅ **Query logic improved** - Category/brand filtering now efficient

## Deployment Steps

### Option 1: Using Makefile (Recommended)

```bash
cd gocrit_server

# Apply database migrations
make run-migrate

# Build the server
make build

# Test locally if desired
make run-app
```

### Option 2: Manual Commands

```bash
cd gocrit_server

# Migrate database
go run ./cmd/migrate/main.go

# Build binary
go build -o bin/kossti-server ./cmd/app

# Copy bin/kossti-server to your deployment location
```

### Option 3: Docker

```bash
cd gocrit_server

# Build Docker image with new changes
docker build -t kossti-server:latest .

# Run migration in container
docker run --env-file .env kossti-server:latest go run ./cmd/migrate/main.go

# Deploy the new image
docker-compose up -d kossti-server
```

## Testing After Deployment

### Test 1: Basic Page Load

```
URL: https://kossti.com/en?category=banking
Expected: Page loads in <1 second
```

### Test 2: Other Categories

- Test other categories work (electronics, etc.)
- Verify no errors in logs

### Test 3: Filters

- Test price range filter
- Test brand filter
- Test search functionality

### Test 4: Database Query Count

Enable database query logging and verify:

```
Expected: 2-5 queries per request
Before: 60+ queries (N+1 problem)
After: 3 queries (1 product query + 1-2 lookup queries)
```

## Rollback Instructions (If Needed)

```bash
# Revert to previous code
git checkout HEAD~1

# Rebuild without new changes
go build -o bin/kossti-server ./cmd/app

# Indexes are harmless and don't need to be removed
# but if you want to remove them, run:
# go run ./cmd/migrate/main.go DROP_INDEXES (if supported)
```

## Success Indicators

- ✅ Banking category page loads instantly
- ✅ Others categories still work
- ✅ No 504 Gateway Timeout errors
- ✅ Database CPU usage reduced
- ✅ Network latency reduced

## Monitoring

After deployment, monitor:

1. **Response time** for `/en?category=banking`
2. **Database query count** - should be 2-5 per request
3. **Error logs** - should see no new errors
4. **Server load** - should see CPU/memory reduction

---

**Expected Result**: Banking (and all other category) pages load immediately and reliably.
