#!/bin/bash

# Script to automatically comment out all failing seeders
# This will iteratively run seeding and comment out products that don't exist

SEEDER_FILE="internal/infrastructure/database/seeders/seeder.go"
MAX_ITERATIONS=50
ITERATION=0

echo "🚀 Starting automated seeder commenting process..."
echo "📁 Seeder file: $SEEDER_FILE"
echo ""

while [ $ITERATION -lt $MAX_ITERATIONS ]; do
    ITERATION=$((ITERATION + 1))
    echo "🔄 Iteration $ITERATION/$MAX_ITERATIONS"
    
    # Run seeding and capture output
    OUTPUT=$(go run ./cmd/migrate/main.go -seed 2>&1)
    
    # Check if seeding completed successfully
    if echo "$OUTPUT" | grep -q "All seeders completed successfully"; then
        echo "✅ SUCCESS! All seeders completed successfully!"
        echo ""
        echo "📊 Summary:"
        echo "   - Total iterations: $ITERATION"
        echo "   - All failing seeders have been commented out"
        exit 0
    fi
    
    # Extract the failing seeder name
    FAIL_LINE=$(echo "$OUTPUT" | grep "failed to run" | head -1)
    
    if [ -z "$FAIL_LINE" ]; then
        echo "⚠️  No more failing seeders found, but completion message not detected"
        echo "Last output:"
        echo "$OUTPUT" | tail -10
        exit 1
    fi
    
    # Extract seeder name (e.g., "YamahaYBR125FiReview" from "failed to run YamahaYBR125FiReview seeder")
    SEEDER_NAME=$(echo "$FAIL_LINE" | grep -oP "failed to run \K[^ ]+(?= seeder)")
    
    if [ -z "$SEEDER_NAME" ]; then
        echo "❌ Could not extract seeder name from: $FAIL_LINE"
        exit 1
    fi
    
    echo "   ❌ Found failing seeder: $SEEDER_NAME"
    
    # Extract the product name from the error message
    PRODUCT_NAME=$(echo "$OUTPUT" | grep -oP "$SEEDER_NAME.*?: \K[^,]+" | head -1)
    echo "   📦 Product not found: $PRODUCT_NAME"
    
    # Create the pattern to search for in seeder.go
    PATTERN="manager\.AddSeeder\(New${SEEDER_NAME}Seeder\(\)\)"
    
    # Check if this seeder is already commented
    if grep -q "// .*$PATTERN" "$SEEDER_FILE"; then
        echo "   ℹ️  Seeder already commented, checking for other instances..."
    fi
    
    # Count occurrences
    OCCURRENCES=$(grep -c "$PATTERN" "$SEEDER_FILE" || echo "0")
    
    if [ "$OCCURRENCES" -eq "0" ]; then
        echo "   ⚠️  Seeder not found in $SEEDER_FILE, might be a different pattern"
        # Try alternative patterns
        ALT_PATTERN="manager\.AddSeeder\(New${SEEDER_NAME}\(\)\)"
        OCCURRENCES=$(grep -c "$ALT_PATTERN" "$SEEDER_FILE" || echo "0")
        if [ "$OCCURRENCES" -gt "0" ]; then
            PATTERN="$ALT_PATTERN"
            echo "   ✓ Found with alternative pattern"
        fi
    fi
    
    echo "   📝 Found $OCCURRENCES occurrence(s) of $SEEDER_NAME"
    
    # Comment out all occurrences
    if [ "$OCCURRENCES" -gt "0" ]; then
        sed -i "s/\(\s*\)manager\.AddSeeder(New${SEEDER_NAME}Seeder())/\1\/\/ manager.AddSeeder(New${SEEDER_NAME}Seeder()) \/\/ Product not found: $PRODUCT_NAME/" "$SEEDER_FILE"
        
        # Also try the variant without "Seeder" suffix in constructor name
        sed -i "s/\(\s*\)manager\.AddSeeder(New${SEEDER_NAME}())/\1\/\/ manager.AddSeeder(New${SEEDER_NAME}()) \/\/ Product not found: $PRODUCT_NAME/" "$SEEDER_FILE"
        
        echo "   ✅ Commented out $SEEDER_NAME"
    else
        echo "   ⚠️  Could not find seeder in file, manual intervention may be needed"
    fi
    
    echo ""
    sleep 1
done

echo "⚠️  Maximum iterations ($MAX_ITERATIONS) reached"
echo "Some seeders may still need commenting. Please review manually."
exit 1
