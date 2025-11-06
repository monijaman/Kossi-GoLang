package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationKeyTranslationSeeder handles seeding specification key translations
type SpecificationKeyTranslationSeeder struct {
	BaseSeeder
}

// NewSpecificationKeyTranslationSeeder creates a new specification key translation seeder
func NewSpecificationKeyTranslationSeeder() *SpecificationKeyTranslationSeeder {
	return &SpecificationKeyTranslationSeeder{
		BaseSeeder: BaseSeeder{name: "Specification Key Translations"},
	}
}

// Seed implements the Seeder interface
func (skts *SpecificationKeyTranslationSeeder) Seed(db *gorm.DB) error {
	translations := map[string]string{
		// Mobile & Electronics
		"2G Bands":                    "২জি ব্যান্ড",
		"3.5mm Audio Jack":            "৩.৫মিমি অডিও জ্যাক",
		"3G Bands":                    "৩জি ব্যান্ড",
		"4G Bands":                    "৪জি ব্যান্ড",
		"5G Bands":                    "৫জি ব্যান্ড",
		"Alarm Feature":               "অ্যালার্ম বৈশিষ্ট্য",
		"Alert Types":                 "সতর্কতা প্রকার",
		"Announcement Date":           "ঘোষণা তারিখ",
		"Audio Quality":               "অডিও গুণমান",
		"Available Colors":            "উপলব্ধ রং",
		"Battery Capacity":            "ব্যাটারির ধারণক্ষমতা",
		"Battery Type":                "ব্যাটারির ধরন",
		"Bluetooth Version":           "ব্লুটুথ সংস্করণ",
		"Build Material":              "নির্মাণের উপকরণ",
		"Call Records":                "কল রেকর্ড",
		"Camera Features":             "ক্যামেরার বৈশিষ্ট্য",
		"Card Slot Type":              "কার্ড স্লট টাইপ",
		"Charging Speed":              "চার্জিং গতি",
		"Chipset":                     "চিপসেট",
		"Clock Feature":               "ঘড়ির বৈশিষ্ট্য",
		"CPU Type":                    "সিপিইউ টাইপ",
		"Device Size":                 "ডিভাইসের আকার",
		"Device Status":               "ডিভাইসের অবস্থা",
		"Device Type":                 "ডিভাইসের ধরন",
		"Dimensions":                  "আয়তন",
		"Display Characteristics":     "ডিসপ্লের বৈশিষ্ট্য",
		"Display Type":                "ডিসপ্লে টাইপ",
		"Dual or Triple Camera Setup": "ডুয়াল বা ট্রিপল ক্যামেরা সেটআপ",
		"Dual SIM":                    "ডুয়াল সিম",
		"EDGE (Enhanced Data Rates for GSM Evolution)": "এডিজি (জিএসএম বিবর্তনের জন্য উন্নত ডেটা হার)",
		"Front Camera Resolution":                      "ফ্রন্ট ক্যামেরার রেজুলেশন",
		"Front Camera Video Resolution":                "ফ্রন্ট ক্যামেরার ভিডিও রেজুলেশন",
		"GPU Type":                                     "জিপিইউ টাইপ",
		"GPRS (General Packet Radio Service)":          "জিপিআরএস (সাধারণ প্যাকেট রেডিও পরিষেবা)",
		"Infrared Port":                                "ইনফ্রারেড পোর্ট",
		"Internal Memory Capacity":                     "আভ্যন্তরীণ মেমোরির ধারণক্ষমতা",
		"Java Support":                                 "জাভা সমর্থন",
		"Loudspeaker Quality":                          "লাউডস্পিকার গুণমান",
		"Main Camera Resolution":                       "মেইন ক্যামেরার রেজুলেশন",
		"Main Camera Video Resolution":                 "মেইন ক্যামেরার ভিডিও রেজুলেশন",
		"Messaging Features":                           "মেসেজিং বৈশিষ্ট্য",
		"Model Variants":                               "মডেল ভ্যারিয়েন্ট",
		"Music Playback Duration":                      "সঙ্গীত প্লেব্যাক সময়কাল",
		"Network Speed":                                "নেটওয়ার্ক স্পিড",
		"New Battery Capacity":                         "নতুন ব্যাটারির ধারণক্ষমতা",
		"NFC Support":                                  "এনএফসি সমর্থন",
		"No Support":                                   "কোনো সমর্থন নেই",
		"Old Battery Capacity":                         "পুরাতন ব্যাটারির ধারণক্ষমতা",
		"Operating System":                             "অপারেটিং সিস্টেম",
		"Penta Camera Setup":                           "পেন্টা ক্যামেরা সেটআপ",
		"Performance Metrics":                          "কার্যক্ষমতার পরিমাপক",
		"Phonebook Capacity":                           "ফোনবুক ধারণক্ষমতা",
		"Physical Keyboard":                            "শারীরিক কীবোর্ড",
		"Positioning System":                           "অবস্থান নির্ধারণের ব্যবস্থা",
		"Preinstalled Games":                           "পূর্ব-স্থাপিত গেম",
		"Price":                                        "মূল্য",
		"Processor Speed":                              "প্রসেসরের গতি",
		"Quad Camera Setup":                            "কোয়াড ক্যামেরা সেটআপ",
		"Radio Support":                                "রেডিও সমর্থন",
		"RAM":                                          "র‍্যাম",
		"Refresh Rate":                                 "রিফ্রেশ রেট",
		"Resolution":                                   "রেজুলেশন",
		"SAR (Specific Absorption Rate)":               "এসএআর (বিশেষ শোষণ হার)",
		"SAR (Specific Absorption Rate) EU":            "এসএআর (বিশেষ শোষণ হার) ইউরোপীয় ইউনিয়ন",
		"Screen Protection":                            "স্ক্রীন সুরক্ষা",
		"Screen Size":                                  "স্ক্রীনের আকার",
		"Sensors":                                      "সেন্সর",
		"SIM Card Type":                                "সিম কার্ডের ধরন",
		"Special Features":                             "বিশেষ বৈশিষ্ট্য",
		"Standby Time":                                 "স্ট্যান্ডবাই সময়",
		"Storage Capacity":                             "সংরক্ষণ ক্ষমতা",
		"Supported Languages":                          "সমর্থিত ভাষাসমূহ",
		"Talk Time Duration":                           "কথা বলার সময়কাল",
		"Technology":                                   "প্রযুক্তি",
		"Triple Camera Setup":                          "ট্রিপল ক্যামেরা সেটআপ",
		"USB Type":                                     "ইউএসবি টাইপ",
		"Video Recording":                              "ভিডিও রেকর্ডিং",
		"Water Resistance":                             "জল প্রতিরোধ",
		"Web Browser":                                  "ওয়েব ব্রাউজার",
		"Weight":                                       "ওজন",
		"WiFi":                                         "ওয়াইফাই",
		"Wireless Charging":                            "ওয়্যারলেস চার্জিং",

		// Automotive (Cars & Motorbikes)
		"Acceleration 0-100 km/h":        "ত্বরণ ০-১০০ কিমি/ঘণ্টা",
		"ABS (Anti-lock Braking System)": "এবিএস (অ্যান্টি-লক ব্রেকিং সিস্টেম)",
		"Airbags":                        "এয়ারব্যাগ",
		"Air Conditioning":               "এয়ার কন্ডিশনিং",
		"All-Wheel Drive":                "সর্ব-চাকা চালনা",
		"Bluetooth Connectivity":         "ব্লুটুথ সংযোগ",
		"Body Type":                      "বডি টাইপ",
		"Boot Space":                     "বুট স্পেস",
		"Brake Type":                     "ব্রেকের ধরন",
		"CC (Cubic Capacity)":            "সিসি (ঘন ক্ষমতা)",
		"Climate Control":                "জলবায়ু নিয়ন্ত্রণ",
		"Clutch Type":                    "ক্লাচের ধরন",
		"Cooling System":                 "শীতলীকরণ ব্যবস্থা",
		"Cruise Control":                 "ক্রুজ নিয়ন্ত্রণ",
		"Cylinder Configuration":         "সিলিন্ডার কনফিগারেশন",
		"Displacement":                   "ডিসপ্লেসমেন্ট",
		"Drive Type":                     "ড্রাইভের ধরন",
		"Emission Standard":              "নির্গমন মান",
		"Engine Type":                    "ইঞ্জিনের ধরন",
		"Fuel Capacity":                  "জ্বালানি ধারণক্ষমতা",
		"Fuel System":                    "জ্বালানি ব্যবস্থা",
		"Fuel Tank Capacity":             "জ্বালানি ট্যাংক ধারণক্ষমতা",
		"Fuel Type":                      "জ্বালানির ধরন",
		"Gearbox":                        "গিয়ারবক্স",
		"Ground Clearance":               "গ্রাউন্ড ক্লিয়ারেন্স",
		"Headlight Type":                 "হেডলাইট টাইপ",
		"Horsepower":                     "অশ্বশক্তি",
		"Ignition Type":                  "ইগনিশন টাইপ",
		"Infotainment System":            "ইনফোটেইনমেন্ট সিস্টেম",
		"Kerb Weight":                    "কার্ব ওজন",
		"Length":                         "দৈর্ঘ্য",
		"Max Power":                      "সর্বোচ্চ শক্তি",
		"Max Torque":                     "সর্বোচ্চ টর্ক",
		"Mileage":                        "মাইলেজ",
		"Number of Cylinders":            "সিলিন্ডারের সংখ্যা",
		"Number of Gears":                "গিয়ারের সংখ্যা",
		"Number of Seats":                "আসনের সংখ্যা",
		"Parking Sensors":                "পার্কিং সেন্সর",
		"Power Steering":                 "পাওয়ার স্টিয়ারিং",
		"Rear Camera":                    "পিছনের ক্যামেরা",
		"Seating Capacity":               "বসার ক্ষমতা",
		"Starting System":                "স্টার্টিং সিস্টেম",
		"Sunroof":                        "সানরুফ",
		"Suspension Type":                "সাসপেনশন টাইপ",
		"Top Speed":                      "সর্বোচ্চ গতি",
		"Torque":                         "টর্ক",
		"Touchscreen":                    "টাচস্ক্রীন",
		"Transmission":                   "ট্রান্সমিশন",
		"Tyre Size":                      "টায়ারের আকার",
		"Tyre Type":                      "টায়ারের ধরন",
		"Valve Configuration":            "ভালভ কনফিগারেশন",
		"Valve Per Cylinder":             "প্রতি সিলিন্ডারে ভালভ",
		"Wheelbase":                      "হুইলবেস",
		"Width":                          "প্রস্থ",

		// Home Appliances
		"Annual Energy Consumption": "বার্ষিক শক্তি খরচ",
		"Capacity":                  "ক্ষমতা",
		"Compressor Type":           "কম্প্রেসারের ধরন",
		"Control Type":              "নিয়ন্ত্রণের ধরন",
		"Defrost Type":              "ডিফ্রস্ট টাইপ",
		"Door Type":                 "দরজার ধরন",
		"Energy Efficiency Rating":  "শক্তি দক্ষতা রেটিং",
		"Energy Star Rating":        "এনার্জি স্টার রেটিং",
		"Filter Type":               "ফিল্টারের ধরন",
		"Freezer Capacity":          "ফ্রিজার ক্ষমতা",
		"Installation Type":         "স্থাপনের ধরন",
		"Inverter Technology":       "ইনভার্টার প্রযুক্তি",
		"Material":                  "উপাদান",
		"Noise Level":               "শব্দের মাত্রা",
		"Number of Burners":         "বার্নারের সংখ্যা",
		"Number of Compartments":    "কম্পার্টমেন্টের সংখ্যা",
		"Number of Doors":           "দরজার সংখ্যা",
		"Number of Shelves":         "তাকের সংখ্যা",
		"Power Consumption":         "বিদ্যুৎ খরচ",
		"Refrigerator Capacity":     "রেফ্রিজারেটর ক্ষমতা",
		"Spin Speed":                "স্পিন গতি",
		"Temperature Control":       "তাপমাত্রা নিয়ন্ত্রণ",
		"Timer":                     "টাইমার",
		"Voltage":                   "ভোল্টেজ",
		"Warranty Period":           "ওয়ারেন্টি সময়কাল",
		"Wash Programs":             "ওয়াশ প্রোগ্রাম",
		"Water Consumption":         "জল খরচ",

		// TV & Display
		"Aspect Ratio":   "আস্পেক্ট রেশিও",
		"Contrast Ratio": "কনট্রাস্ট রেশিও",
		"HDR Support":    "এইচডিআর সমর্থন",
		"HDMI Ports":     "এইচডিএমআই পোর্ট",
		"Panel Type":     "প্যানেল টাইপ",
		"Response Time":  "প্রতিক্রিয়া সময়",
		"Smart TV":       "স্মার্ট টিভি",
		"Sound Output":   "সাউন্ড আউটপুট",
		"Speaker Type":   "স্পিকারের ধরন",
		"USB Ports":      "ইউএসবি পোর্ট",
		"Viewing Angle":  "দেখার কোণ",

		// Audio Equipment
		"Audio Formats":         "অডিও ফরম্যাট",
		"Driver Size":           "ড্রাইভারের আকার",
		"Frequency Response":    "ফ্রিকোয়েন্সি প্রতিক্রিয়া",
		"Impedance":             "প্রতিবন্ধকতা",
		"Noise Cancellation":    "শব্দ বাতিলকরণ",
		"Playback Time":         "প্লেব্যাক সময়",
		"Sensitivity":           "সংবেদনশীলতা",
		"Signal to Noise Ratio": "সিগন্যাল টু নয়েজ রেশিও",
		"Wired/Wireless":        "তারযুক্ত/তারবিহীন",

		// Camera & Photography
		"Aperture":            "অ্যাপারচার",
		"Auto Focus":          "অটো ফোকাস",
		"Digital Zoom":        "ডিজিটাল জুম",
		"Flash Type":          "ফ্ল্যাশ টাইপ",
		"Image Stabilization": "ইমেজ স্ট্যাবিলাইজেশন",
		"ISO Range":           "আইএসও রেঞ্জ",
		"Lens Mount":          "লেন্স মাউন্ট",
		"Maximum Aperture":    "সর্বোচ্চ অ্যাপারচার",
		"Megapixels":          "মেগাপিক্সেল",
		"Optical Zoom":        "অপটিক্যাল জুম",
		"Sensor Size":         "সেন্সর আকার",
		"Sensor Type":         "সেন্সর টাইপ",
		"Shutter Speed":       "শাটার স্পিড",
		"Video Format":        "ভিডিও ফরম্যাট",
		"Viewfinder Type":     "ভিউফাইন্ডার টাইপ",

		// Laptop & Computer
		"Backlit Keyboard":     "ব্যাকলিট কীবোর্ড",
		"Graphics Card":        "গ্রাফিক্স কার্ড",
		"Hard Drive Type":      "হার্ড ড্রাইভ টাইপ",
		"Optical Drive":        "অপটিক্যাল ড্রাইভ",
		"Port Types":           "পোর্ট টাইপ",
		"Processor Brand":      "প্রসেসর ব্র্যান্ড",
		"Processor Generation": "প্রসেসর প্রজন্ম",
		"Processor Model":      "প্রসেসর মডেল",
		"Screen Resolution":    "স্ক্রীন রেজুলেশন",
		"SSD Capacity":         "এসএসডি ক্ষমতা",
		"Touchscreen Display":  "টাচস্ক্রীন ডিসপ্লে",

		// Kitchen Appliances
		"Bowl Capacity":            "বাটি ক্ষমতা",
		"Grinder Jars":             "গ্রাইন্ডার জার",
		"Heating Element":          "হিটিং এলিমেন্ট",
		"Number of Speed Settings": "গতি সেটিংয়ের সংখ্যা",
		"Oven Capacity":            "ওভেন ক্ষমতা",
		"Power Rating":             "পাওয়ার রেটিং",
		"Pressure Settings":        "চাপ সেটিং",
		"Safety Features":          "সুরক্ষা বৈশিষ্ট্য",
		"Temperature Range":        "তাপমাত্রা পরিসীমা",

		// Watches
		"Band Material":         "ব্যান্ড উপাদান",
		"Band Width":            "ব্যান্ড প্রস্থ",
		"Case Diameter":         "কেস ব্যাস",
		"Case Material":         "কেস উপাদান",
		"Case Thickness":        "কেসের পুরুত্ব",
		"Clasp Type":            "ক্ল্যাস্প টাইপ",
		"Crystal Type":          "ক্রিস্টাল টাইপ",
		"Dial Color":            "ডায়াল রঙ",
		"Display Style":         "ডিসপ্লে স্টাইল",
		"Movement Type":         "মুভমেন্ট টাইপ",
		"Strap Color":           "স্ট্র্যাপ রঙ",
		"Watch Type":            "ঘড়ির ধরন",
		"Water Resistant Depth": "জল প্রতিরোধী গভীরতা",

		// Clothing & Fashion
		"Care Instructions": "যত্নের নির্দেশাবলী",
		"Closure Type":      "বন্ধের ধরন",
		"Collar Style":      "কলার স্টাইল",
		"Fabric":            "ফ্যাব্রিক",
		"Fit Type":          "ফিট টাইপ",
		"Heel Height":       "হিলের উচ্চতা",
		"Heel Type":         "হিলের ধরন",
		"Lining Material":   "আস্তরণ উপাদান",
		"Neck Type":         "গলার ধরন",
		"Occasion":          "উপলক্ষ",
		"Outer Material":    "বাইরের উপাদান",
		"Pattern":           "প্যাটার্ন",
		"Season":            "ঋতু",
		"Sleeve Length":     "হাতার দৈর্ঘ্য",
		"Sleeve Type":       "হাতার ধরন",
		"Sole Material":     "সোল উপাদান",
		"Style":             "স্টাইল",
		"Upper Material":    "উপরের উপাদান",

		// Furniture
		"Assembly Required":   "সমাবেশ প্রয়োজন",
		"Finish Type":         "ফিনিশ টাইপ",
		"Frame Material":      "ফ্রেম উপাদান",
		"Number of Drawers":   "ড্রয়ারের সংখ্যা",
		"Primary Color":       "প্রাথমিক রঙ",
		"Shape":               "আকৃতি",
		"Upholstery Material": "আপহোলস্টারি উপাদান",

		// Groceries & Food
		"Allergen Information":    "অ্যালার্জেন তথ্য",
		"Brand Origin":            "ব্র্যান্ড উৎপত্তি",
		"Calories":                "ক্যালোরি",
		"Carbohydrates":           "কার্বোহাইড্রেট",
		"Country of Origin":       "উৎপত্তি দেশ",
		"Dietary Type":            "খাদ্যের ধরন",
		"Expiry Date":             "মেয়াদ শেষের তারিখ",
		"Fat Content":             "চর্বি উপাদান",
		"Flavor":                  "স্বাদ",
		"Ingredients":             "উপাদান",
		"Manufacturing Date":      "উৎপাদন তারিখ",
		"Net Quantity":            "নেট পরিমাণ",
		"Nutritional Information": "পুষ্টি তথ্য",
		"Organic":                 "জৈব",
		"Package Contents":        "প্যাকেজ বিষয়বস্তু",
		"Package Type":            "প্যাকেজ টাইপ",
		"Protein":                 "প্রোটিন",
		"Serving Size":            "পরিবেশন আকার",
		"Shelf Life":              "শেলফ লাইফ",
		"Sugar Content":           "চিনির পরিমাণ",
		"Veg/Non-Veg":             "নিরামিষ/আমিষ",

		// Health & Beauty
		"Application Area":   "প্রয়োগের এলাকা",
		"Benefits":           "উপকারিতা",
		"Formulation":        "ফরমুলেশন",
		"Gender":             "লিঙ্গ",
		"Hair Type":          "চুলের ধরন",
		"Scent":              "সুগন্ধ",
		"Skin Type":          "ত্বকের ধরন",
		"SPF":                "এসপিএফ",
		"Usage Instructions": "ব্যবহারের নির্দেশাবলী",
		"Volume":             "ভলিউম",

		// Books
		"Author":           "লেখক",
		"Binding":          "বাইন্ডিং",
		"Edition":          "সংস্করণ",
		"ISBN":             "আইএসবিএন",
		"Language":         "ভাষা",
		"Number of Pages":  "পৃষ্ঠার সংখ্যা",
		"Publication Date": "প্রকাশনার তারিখ",
		"Publisher":        "প্রকাশক",

		// Bags & Luggage
		"Closure":          "ক্লোজার",
		"Compartments":     "কম্পার্টমেন্ট",
		"Handle Type":      "হ্যান্ডেল টাইপ",
		"Lock Type":        "লকের ধরন",
		"Number of Wheels": "চাকার সংখ্যা",
		"Strap Type":       "স্ট্র্যাপ টাইপ",

		// Toys & Games
		"Age Group":         "বয়স গোষ্ঠী",
		"Battery Required":  "ব্যাটারি প্রয়োজন",
		"Educational Value": "শিক্ষাগত মূল্য",
		"Number of Players": "খেলোয়াড়ের সংখ্যা",
		"Skill Development": "দক্ষতা উন্নয়ন",

		// Sports Equipment
		"Frame Size":     "ফ্রেম আকার",
		"Grip Size":      "গ্রিপ আকার",
		"String Tension": "স্ট্রিং টেনশন",

		// General
		"Brand":                  "ব্র্যান্ড",
		"Color":                  "রঙ",
		"Condition":              "অবস্থা",
		"Country of Manufacture": "উৎপাদন দেশ",
		"Included Components":    "অন্তর্ভুক্ত উপাদান",
		"Manufacturer":           "প্রস্তুতকারক",
		"Model Name":             "মডেল নাম",
		"Model Number":           "মডেল নম্বর",
		"Product Description":    "পণ্যের বিবরণ",
		"Product Dimensions":     "পণ্যের আয়তন",
		"Product Weight":         "পণ্যের ওজন",
		"Size":                   "আকার",
		"SKU":                    "এসকেইউ",
		"Warranty":               "ওয়ারেন্টি",
	}

	for englishKey, bengaliTranslation := range translations {
		// Find the specification key by its English name
		var specKeyModel models.SpecificationKeyModel
		result := db.Where("specification_key = ?", englishKey).First(&specKeyModel)

		if result.Error == nil {
			// Check if translation already exists
			var existingTranslation models.SpecificationKeyTranslationModel
			translationResult := db.Where("specification_key_id = ? AND locale = ?", specKeyModel.ID, "bn").First(&existingTranslation)

			if translationResult.Error == gorm.ErrRecordNotFound {
				// Create new translation
				translation := &models.SpecificationKeyTranslationModel{
					SpecificationKeyID: specKeyModel.ID,
					Locale:             "bn",
					SpecificationKey:   bengaliTranslation,
				}

				if err := db.Create(translation).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
