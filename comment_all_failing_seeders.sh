#!/bin/bash

# List of all seeder function names to comment out (extracted from validation)
SEEDERS_TO_COMMENT=(
    "NewBajajRX100Review"
    "NewHeroSplendorPlusDTSiReview"
    "NewHoneyWellMaxBatch24"
    "NewKTMDuke125Batch23"
    "NewSymBike100Batch22"
    "NewHeroMaster125Review"
    "NewTVSApacheRTR160TwoWheelerReview"
    "NewTVSApacheRTR200Review"
    "NewTVSApacheRTR200Batch23"
    "NewBajajDiscover150Review"
    "NewBajajDominar400Review"
    "NewCebiBike125Batch22"
    "NewHondaCB110Review"
    "NewSuzukiGS150RBatch25"
    "NewSuzukiGSXS1000ReviewBatch16"
    "NewTVSJupiter125Review"
    "NewHeroHFDeluxeIXReviewBatch18"
    "NewTVSScootypBatch25"
    "NewYamahaR3Batch24"
    "NewYamahaYBR125FiReview"
    "NewHondaX1XBatch25"
    "NewVespaLX125Review"
    "NewVespaSprintBatch24"
    "NewYamahaFZ300ReviewBatch18"
    "NewBajajAvenger160ReviewBatch19"
    "NewBajajPlatina110Review"
    "NewBajajPulsarNS200ReviewBatch18"
    "NewHeroPulseBatch24"
    "NewHeroMotocorpExtaReview"
    "NewSuzukiGSX150RBatch21"
    "NewSuzukiVStrom650ReviewBatch18"
    "NewUniversalMotorBike100Batch21"
    "NewBajajAvenger200DTSiReview"
    "NewHeroHF100Batch22"
    "NewHondaActiva6Batch23"
    "NewHondaActiva6GReviewBatch19"
    "NewHondaActiva6GBatch26"
    "NewKawasakiNinja650ReviewBatch16"
    "NewKawasakiNinja650Batch25"
    "NewYamahaFZSFiDTSReview"
    "NewYamahaSZ150Batch21"
    "NewBajajDualDo125Batch24"
    "NewBajajPulsarAS200Review"
    "NewHondaCB350ReviewBatch16"
    "NewBajajPulsar200TwinDiscReview"
    "NewHondaCD110DreamReview"
    "NewKawasakiNinja250Batch21"
    "NewRoyalEnfieldSignetBatch24"
    "NewYamahaFZS150Batch25"
    "NewYamahaR15V3Batch23"
    "NewHondaCB200Batch21"
    "NewSuzukiGixxer155Review"
    "NewSuzukiGixxer155Batch22"
    "NewSuzukiGixxer155Batch26"
    "NewTVSApache150Batch25"
    "NewBajajPlatinaBatch22"
    "NewBajajPulsar220FBatch23"
    "NewHondaCRF450RallyReviewBatch18"
    "NewKawasakiNinja400ReviewBatch19"
    "NewKawasakiNinja400Batch22"
    "NewUniversalMotorbike75Batch22"
    "NewYamahaXSR160ReviewBatch16"
    "NewYamahaYZF150Batch22"
    "NewPiggiBike125Batch21"
    "NewBajajDominarD400Batch26"
    "NewBenelliTNT600ReviewBatch16"
    "NewHondaCBTriggerReview"
    "NewSuzukiStorm125Batch23"
    "NewUniversalMotorbike110Batch23"
    "NewBajajPulsarN160Review"
    "NewKTMDuke390Review"
    "NewKTMDuke390Batch26"
    "NewYamahaXMax300ReviewBatch18"
    "NewBajajAvenger220Review"
    "NewBajajAvenger220Batch24"
    "NewHeroHFDeluxeIXReview"
    "NewHondaCB500FReviewBatch19"
    "NewHondaCG125Batch22"
    "NewHondaCG125Batch26"
    "NewSuzukiIntruder150Review"
    "NewBajajPulsarAS250Review"
    "NewHeroMotocorpXPulseReview"
    "NewKawasakiNinja1000SXReviewBatch18"
    "NewVespaLX300Batch21"
)

SEEDER_FILE="internal/infrastructure/database/seeders/seeder.go"

echo "🔧 Starting to comment out ${#SEEDERS_TO_COMMENT[@]} failing seeders..."
echo ""

# Create backup
cp "$SEEDER_FILE" "${SEEDER_FILE}.backup"
echo "✅ Backup created: ${SEEDER_FILE}.backup"
echo ""

COMMENTED_COUNT=0

for seeder in "${SEEDERS_TO_COMMENT[@]}"; do
    # Check if seeder exists and is not already commented
    if grep -q "^\s*manager\.AddSeeder($seeder()" "$SEEDER_FILE"; then
        echo "🔄 Commenting out: $seeder"
        
        # Comment out the line (matching with or without Seeder suffix)
        sed -i "s/\(\s*\)manager\.AddSeeder($seeder())/\1\/\/ manager.AddSeeder($seeder()) \/\/ Product not found/" "$SEEDER_FILE"
        
        COMMENTED_COUNT=$((COMMENTED_COUNT + 1))
    else
        # Check if already commented
        if grep -q "// .*manager\.AddSeeder($seeder()" "$SEEDER_FILE"; then
            echo "   ℹ️  Already commented: $seeder"
        else
            echo "   ⚠️  Not found: $seeder"
        fi
    fi
done

echo ""
echo "=========================================="
echo "✅ Completed!"
echo "   Newly commented: $COMMENTED_COUNT"
echo "   Total to comment: ${#SEEDERS_TO_COMMENT[@]}"
echo "=========================================="
echo ""
echo "🔄 Now running seeding to verify..."
