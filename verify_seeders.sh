#!/bin/bash

echo "=== Seeder Generation Verification ==="
echo ""
echo "1. Generated Seeder Files:"
echo "   Total files: $(ls -1 internal/infrastructure/database/seeders/specification_seeder_refrigerator_marcel-*.go 2>/dev/null | wc -l)"
echo ""

echo "2. Index File:"
if [ -f internal/infrastructure/database/seeders/specification_seeders_marcel_index.go ]; then
    echo "   ✓ Index file exists"
    echo "   Lines: $(wc -l < internal/infrastructure/database/seeders/specification_seeders_marcel_index.go)"
else
    echo "   ✗ Index file not found"
fi
echo ""

echo "3. Sample Seeder Content:"
SAMPLE_FILE=$(ls -1 internal/infrastructure/database/seeders/specification_seeder_refrigerator_marcel-*.go 2>/dev/null | head -1)
if [ -n "$SAMPLE_FILE" ]; then
    echo "   File: $(basename $SAMPLE_FILE)"
    echo "   Lines: $(wc -l < "$SAMPLE_FILE")"
    echo ""
    echo "   ✓ Package declaration:"
    head -1 "$SAMPLE_FILE"
    echo ""
    echo "   ✓ Type definition:"
    grep "^type " "$SAMPLE_FILE" | head -1
    echo ""
    echo "   ✓ Constructor function:"
    grep "^func New" "$SAMPLE_FILE" | head -1
fi
echo ""

echo "4. Specification Data:"
echo "   CSV file size: $(ls -lh data/marcel_specs.csv 2>/dev/null | awk '{print $5}')"
echo "   CSV line count: $(wc -l < data/marcel_specs.csv 2>/dev/null)"
echo ""

echo "=== Verification Complete ==="
