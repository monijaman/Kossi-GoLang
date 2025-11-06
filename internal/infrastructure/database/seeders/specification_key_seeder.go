package seeders

import (
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationKeySeeder handles seeding specification keys
type SpecificationKeySeeder struct {
	BaseSeeder
}

// NewSpecificationKeySeeder creates a new specification key seeder
func NewSpecificationKeySeeder() *SpecificationKeySeeder {
	return &SpecificationKeySeeder{
		BaseSeeder: BaseSeeder{name: "Specification Keys"},
	}
}

// Seed implements the Seeder interface
func (sks *SpecificationKeySeeder) Seed(db *gorm.DB) error {
	keys := []string{
		// Mobile & Electronics
		"2G Bands",
		"3.5mm Audio Jack",
		"3G Bands",
		"4G Bands",
		"5G Bands",
		"Alarm Feature",
		"Alert Types",
		"Announcement Date",
		"Audio Quality",
		"Available Colors",
		"Battery Capacity",
		"Battery Type",
		"Bluetooth Version",
		"Build Material",
		"Call Records",
		"Camera Features",
		"Card Slot Type",
		"Charging Speed",
		"Chipset",
		"Clock Feature",
		"CPU Type",
		"Device Size",
		"Device Status",
		"Device Type",
		"Dimensions",
		"Display Characteristics",
		"Display Type",
		"Dual or Triple Camera Setup",
		"Dual SIM",
		"EDGE (Enhanced Data Rates for GSM Evolution)",
		"Front Camera Resolution",
		"Front Camera Video Resolution",
		"GPU Type",
		"GPRS (General Packet Radio Service)",
		"Infrared Port",
		"Internal Memory Capacity",
		"Java Support",
		"Loudspeaker Quality",
		"Main Camera Resolution",
		"Main Camera Video Resolution",
		"Messaging Features",
		"Model Variants",
		"Music Playback Duration",
		"Network Speed",
		"New Battery Capacity",
		"NFC Support",
		"No Support",
		"Old Battery Capacity",
		"Operating System",
		"Penta Camera Setup",
		"Performance Metrics",
		"Phonebook Capacity",
		"Physical Keyboard",
		"Positioning System",
		"Preinstalled Games",
		"Price",
		"Processor Speed",
		"Quad Camera Setup",
		"Radio Support",
		"RAM",
		"Refresh Rate",
		"Resolution",
		"SAR (Specific Absorption Rate)",
		"SAR (Specific Absorption Rate) EU",
		"Screen Protection",
		"Screen Size",
		"Sensors",
		"SIM Card Type",
		"Special Features",
		"Standby Time",
		"Storage Capacity",
		"Supported Languages",
		"Talk Time Duration",
		"Technology",
		"Triple Camera Setup",
		"USB Type",
		"Video Recording",
		"Water Resistance",
		"Web Browser",
		"Weight",
		"WiFi",
		"Wireless Charging",

		// Automotive (Cars & Motorbikes)
		"Acceleration 0-100 km/h",
		"ABS (Anti-lock Braking System)",
		"Airbags",
		"Air Conditioning",
		"All-Wheel Drive",
		"Bluetooth Connectivity",
		"Body Type",
		"Boot Space",
		"Brake Type",
		"CC (Cubic Capacity)",
		"Climate Control",
		"Clutch Type",
		"Cooling System",
		"Cruise Control",
		"Cylinder Configuration",
		"Displacement",
		"Drive Type",
		"Emission Standard",
		"Engine Type",
		"Fuel Capacity",
		"Fuel System",
		"Fuel Tank Capacity",
		"Fuel Type",
		"Gearbox",
		"Ground Clearance",
		"Headlight Type",
		"Horsepower",
		"Ignition Type",
		"Infotainment System",
		"Kerb Weight",
		"Length",
		"Max Power",
		"Max Torque",
		"Mileage",
		"Number of Cylinders",
		"Number of Gears",
		"Number of Seats",
		"Parking Sensors",
		"Power Steering",
		"Rear Camera",
		"Seating Capacity",
		"Starting System",
		"Sunroof",
		"Suspension Type",
		"Top Speed",
		"Torque",
		"Touchscreen",
		"Transmission",
		"Tyre Size",
		"Tyre Type",
		"Valve Configuration",
		"Valve Per Cylinder",
		"Wheelbase",
		"Width",

		// Home Appliances
		"Annual Energy Consumption",
		"Capacity",
		"Compressor Type",
		"Control Type",
		"Defrost Type",
		"Door Type",
		"Energy Efficiency Rating",
		"Energy Star Rating",
		"Filter Type",
		"Freezer Capacity",
		"Installation Type",
		"Inverter Technology",
		"Material",
		"Noise Level",
		"Number of Burners",
		"Number of Compartments",
		"Number of Doors",
		"Number of Shelves",
		"Power Consumption",
		"Refrigerator Capacity",
		"Spin Speed",
		"Temperature Control",
		"Timer",
		"Voltage",
		"Warranty Period",
		"Wash Programs",
		"Water Consumption",

		// TV & Display
		"Aspect Ratio",
		"Contrast Ratio",
		"HDR Support",
		"HDMI Ports",
		"Panel Type",
		"Response Time",
		"Smart TV",
		"Sound Output",
		"Speaker Type",
		"USB Ports",
		"Viewing Angle",

		// Audio Equipment
		"Audio Formats",
		"Driver Size",
		"Frequency Response",
		"Impedance",
		"Noise Cancellation",
		"Playback Time",
		"Sensitivity",
		"Signal to Noise Ratio",
		"Wired/Wireless",

		// Camera & Photography
		"Aperture",
		"Auto Focus",
		"Digital Zoom",
		"Flash Type",
		"Image Stabilization",
		"ISO Range",
		"Lens Mount",
		"Maximum Aperture",
		"Megapixels",
		"Optical Zoom",
		"Sensor Size",
		"Sensor Type",
		"Shutter Speed",
		"Video Format",
		"Viewfinder Type",

		// Laptop & Computer
		"Backlit Keyboard",
		"Graphics Card",
		"Hard Drive Type",
		"Optical Drive",
		"Port Types",
		"Processor Brand",
		"Processor Generation",
		"Processor Model",
		"Screen Resolution",
		"SSD Capacity",
		"Touchscreen Display",

		// Kitchen Appliances
		"Bowl Capacity",
		"Grinder Jars",
		"Heating Element",
		"Number of Speed Settings",
		"Oven Capacity",
		"Power Rating",
		"Pressure Settings",
		"Safety Features",
		"Temperature Range",

		// Watches
		"Band Material",
		"Band Width",
		"Case Diameter",
		"Case Material",
		"Case Thickness",
		"Clasp Type",
		"Crystal Type",
		"Dial Color",
		"Display Style",
		"Movement Type",
		"Strap Color",
		"Watch Type",
		"Water Resistant Depth",

		// Clothing & Fashion
		"Care Instructions",
		"Closure Type",
		"Collar Style",
		"Fabric",
		"Fit Type",
		"Heel Height",
		"Heel Type",
		"Lining Material",
		"Neck Type",
		"Occasion",
		"Outer Material",
		"Pattern",
		"Season",
		"Sleeve Length",
		"Sleeve Type",
		"Sole Material",
		"Style",
		"Upper Material",

		// Furniture
		"Assembly Required",
		"Finish Type",
		"Frame Material",
		"Number of Drawers",
		"Primary Color",
		"Shape",
		"Upholstery Material",

		// Groceries & Food
		"Allergen Information",
		"Brand Origin",
		"Calories",
		"Carbohydrates",
		"Country of Origin",
		"Dietary Type",
		"Expiry Date",
		"Fat Content",
		"Flavor",
		"Ingredients",
		"Manufacturing Date",
		"Net Quantity",
		"Nutritional Information",
		"Organic",
		"Package Contents",
		"Package Type",
		"Protein",
		"Serving Size",
		"Shelf Life",
		"Sugar Content",
		"Veg/Non-Veg",

		// Health & Beauty
		"Application Area",
		"Benefits",
		"Formulation",
		"Gender",
		"Hair Type",
		"Scent",
		"Skin Type",
		"SPF",
		"Usage Instructions",
		"Volume",

		// Books
		"Author",
		"Binding",
		"Edition",
		"ISBN",
		"Language",
		"Number of Pages",
		"Publication Date",
		"Publisher",

		// Bags & Luggage
		"Closure",
		"Compartments",
		"Handle Type",
		"Lock Type",
		"Number of Wheels",
		"Strap Type",

		// Toys & Games
		"Age Group",
		"Battery Required",
		"Educational Value",
		"Number of Players",
		"Skill Development",

		// Sports Equipment
		"Frame Size",
		"Grip Size",
		"String Tension",

		// General
		"Brand",
		"Color",
		"Condition",
		"Country of Manufacture",
		"Included Components",
		"Manufacturer",
		"Model Name",
		"Model Number",
		"Product Description",
		"Product Dimensions",
		"Product Weight",
		"Size",
		"SKU",
		"Warranty",
	}

	for _, key := range keys {
		// Create SpecificationKey entity
		specKey := &entities.SpecificationKey{
			SpecificationKey: key,
		}

		// Check if specification key already exists
		var existingModel models.SpecificationKeyModel
		result := db.Where("specification_key = ?", key).First(&existingModel)

		if result.Error == gorm.ErrRecordNotFound {
			// Convert entity to model for database insertion
			model := &models.SpecificationKeyModel{
				SpecificationKey: specKey.SpecificationKey,
			}

			if err := db.Create(model).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
