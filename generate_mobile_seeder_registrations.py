#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Generate AddSeeder registration calls for all mobile specification seeders
"""

import os
from pathlib import Path

def generate_seeder_registrations(seeders_dir):
    """Generate AddSeeder calls for all mobile specification seeders"""
    
    # Get all mobile spec seeder files
    files = sorted(Path(seeders_dir).glob("specification_seeder_mobile_*.go"))
    
    registrations = []
    
    for filepath in files:
        filename = filepath.stem  # Get filename without extension
        
        # Convert filename to struct name
        # specification_seeder_mobile_iphone_17_pro_max -> SpecificationSeederMobileIphone17ProMax
        parts = filename.split('_')
        struct_name = ''.join(word.capitalize() for word in parts)
        
        # Generate AddSeeder call
        registrations.append(f"\tmanager.AddSeeder(New{struct_name}())")
    
    return registrations

def main():
    seeders_dir = Path(__file__).parent / "internal" / "infrastructure" / "database" / "seeders"
    
    print("Generating AddSeeder registration calls for mobile specification seeders...")
    
    registrations = generate_seeder_registrations(seeders_dir)
    
    print(f"\nFound {len(registrations)} mobile specification seeders\n")
    print("// Mobile Specification Seeders (Category ID: 79)")
    print("// Add these lines to SetupAllSeeders() in seeder.go:")
    print()
    
    for reg in registrations:
        print(reg)
    
    # Save to file
    output_file = Path(__file__).parent / "mobile_seeder_registrations.txt"
    with open(output_file, 'w', encoding='utf-8') as f:
        f.write("// Mobile Specification Seeders (Category ID: 79)\n")
        f.write("// Add these lines to SetupAllSeeders() in seeder.go:\n\n")
        for reg in registrations:
            f.write(reg + "\n")
    
    print(f"\n\nRegistrations saved to: {output_file}")

if __name__ == "__main__":
    main()
