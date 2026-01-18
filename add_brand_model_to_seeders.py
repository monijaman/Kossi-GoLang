#!/usr/bin/env python3
"""
Script to add brand and model specifications to car seeder files
"""

import os
import re

# Define all cars with their brand and model names
cars = [
    # Audi
    ("audi_a3_seeder.go", "Audi", "A3"),
    ("audi_a4_seeder.go", "Audi", "A4"),
    ("audi_a6_seeder.go", "Audi", "A6"),
    ("audi_a8_seeder.go", "Audi", "A8"),
    ("audi_q3_seeder.go", "Audi", "Q3"),
    ("audi_q5_seeder.go", "Audi", "Q5"),
    ("audi_q7_seeder.go", "Audi", "Q7"),
    ("audi_q8_seeder.go", "Audi", "Q8"),
    ("audi_tt_seeder.go", "Audi", "TT"),
    
    # BMW
    ("bmw_3_series_seeder.go", "BMW", "3 Series"),
    ("bmw_5_series_seeder.go", "BMW", "5 Series"),
    ("bmw_7_series_seeder.go", "BMW", "7 Series"),
    ("bmw_i3_seeder.go", "BMW", "i3"),
    ("bmw_i8_seeder.go", "BMW", "i8"),
    ("bmw_x1_seeder.go", "BMW", "X1"),
    ("bmw_x3_seeder.go", "BMW", "X3"),
    ("bmw_x5_seeder.go", "BMW", "X5"),
    ("bmw_x6_seeder.go", "BMW", "X6"),
    ("bmw_z4_seeder.go", "BMW", "Z4"),
    
    # Changan
    ("changan_alsvin_seeder.go", "Changan", "Alsvin"),
    ("changan_cs35_seeder.go", "Changan", "CS35"),
    ("changan_cs55_seeder.go", "Changan", "CS55"),
    ("changan_cs75_seeder.go", "Changan", "CS75"),
    
    # Chery
    ("chery_qq_seeder.go", "Chery", "QQ"),
    ("chery_tiggo3_seeder.go", "Chery", "Tiggo 3"),
    ("chery_tiggo7_seeder.go", "Chery", "Tiggo 7"),
    ("chery_tiggo8_seeder.go", "Chery", "Tiggo 8"),
    
    # Chevrolet
    ("chevrolet_aveo_seeder.go", "Chevrolet", "Aveo"),
    ("chevrolet_captiva_seeder.go", "Chevrolet", "Captiva"),
    ("chevrolet_cruze_seeder.go", "Chevrolet", "Cruze"),
    
    # Daihatsu
    ("daihatsu_cast_seeder.go", "Daihatsu", "Cast"),
    ("daihatsu_mira_seeder.go", "Daihatsu", "Mira"),
    ("daihatsu_move_seeder.go", "Daihatsu", "Move"),
    ("daihatsu_sirion_seeder.go", "Daihatsu", "Sirion"),
    ("daihatsu_terios_seeder.go", "Daihatsu", "Terios"),
    
    # Ford
    ("ford_ecosport_seeder.go", "Ford", "EcoSport"),
    ("ford_endeavour_seeder.go", "Ford", "Endeavour"),
    ("ford_fiesta_seeder.go", "Ford", "Fiesta"),
    ("ford_focus_seeder.go", "Ford", "Focus"),
    ("ford_ranger_seeder.go", "Ford", "Ranger"),
    
    # Geely
    ("geely_coolray_seeder.go", "Geely", "Coolray"),
    ("geely_emgrand_7_seeder.go", "Geely", "Emgrand 7"),
    ("geely_emgrand_x7_seeder.go", "Geely", "Emgrand X7"),
    ("geely_panda_seeder.go", "Geely", "Panda"),
    
    # GMC
    ("gmc_sierra_seeder.go", "GMC", "Sierra"),
    ("gmc_terrain_seeder.go", "GMC", "Terrain"),
    ("gmc_yukon_seeder.go", "GMC", "Yukon"),
    
    # Haval
    ("haval_h2_seeder.go", "Haval", "H2"),
    ("haval_h6_seeder.go", "Haval", "H6"),
    ("haval_jolion_seeder.go", "Haval", "Jolion"),
    
    # Honda
    ("honda_accord_seeder.go", "Honda", "Accord"),
    ("honda_amaze_seeder.go", "Honda", "Amaze"),
    ("honda_br_v_seeder.go", "Honda", "BR-V"),
    ("honda_city_seeder.go", "Honda", "City"),
    ("honda_civic_seeder.go", "Honda", "Civic"),
    ("honda_cr_v_seeder.go", "Honda", "CR-V"),
    ("honda_fit_seeder.go", "Honda", "Fit"),
    ("honda_jazz_seeder.go", "Honda", "Jazz"),
    
    # Hyundai
    ("hyundai_accent_seeder.go", "Hyundai", "Accent"),
    ("hyundai_creta_seeder.go", "Hyundai", "Creta"),
    ("hyundai_elantra_seeder.go", "Hyundai", "Elantra"),
    ("hyundai_i10_seeder.go", "Hyundai", "i10"),
    ("hyundai_i20_seeder.go", "Hyundai", "i20"),
    ("hyundai_santa_fe_seeder.go", "Hyundai", "Santa Fe"),
    ("hyundai_sonata_seeder.go", "Hyundai", "Sonata"),
    ("hyundai_tucson_seeder.go", "Hyundai", "Tucson"),
    ("hyundai_venue_seeder.go", "Hyundai", "Venue"),
    
    # Isuzu
    ("isuzu_d_max_seeder.go", "Isuzu", "D-Max"),
    ("isuzu_mu_x_seeder.go", "Isuzu", "MU-X"),
    
    # Jaguar
    ("jaguar_e_pace_seeder.go", "Jaguar", "E-PACE"),
    ("jaguar_f_pace_seeder.go", "Jaguar", "F-PACE"),
    ("jaguar_xe_seeder.go", "Jaguar", "XE"),
    ("jaguar_xf_seeder.go", "Jaguar", "XF"),
    ("jaguar_xj_seeder.go", "Jaguar", "XJ"),
    
    # Jeep
    ("jeep_cherokee_seeder.go", "Jeep", "Cherokee"),
    ("jeep_compass_seeder.go", "Jeep", "Compass"),
    ("jeep_grand_cherokee_seeder.go", "Jeep", "Grand Cherokee"),
    ("jeep_wrangler_seeder.go", "Jeep", "Wrangler"),
    
    # Kia
    ("kia_carnival_seeder.go", "Kia", "Carnival"),
    ("kia_picanto_seeder.go", "Kia", "Picanto"),
    ("kia_seltos_seeder.go", "Kia", "Seltos"),
    ("kia_sonet_seeder.go", "Kia", "Sonet"),
    ("kia_sorento_seeder.go", "Kia", "Sorento"),
    ("kia_sportage_seeder.go", "Kia", "Sportage"),
    ("kia_stonic_seeder.go", "Kia", "Stonic"),
    
    # Lexus
    ("lexus_es_seeder.go", "Lexus", "ES"),
    ("lexus_lx_seeder.go", "Lexus", "LX"),
    ("lexus_nx_seeder.go", "Lexus", "NX"),
    ("lexus_rx_seeder.go", "Lexus", "RX"),
    ("lexus_ux_seeder.go", "Lexus", "UX"),
    
    # Mahindra
    ("mahindra_scorpio_seeder.go", "Mahindra", "Scorpio"),
    ("mahindra_thar_seeder.go", "Mahindra", "Thar"),
    ("mahindra_xuv500_seeder.go", "Mahindra", "XUV500"),
    
    # Mazda
    ("mazda_bt50_seeder.go", "Mazda", "BT-50"),
    ("mazda_cx3_seeder.go", "Mazda", "CX-3"),
    ("mazda_cx30_seeder.go", "Mazda", "CX-30"),
    ("mazda_cx5_seeder.go", "Mazda", "CX-5"),
    ("mazda_mazda2_seeder.go", "Mazda", "Mazda2"),
    ("mazda_mazda3_seeder.go", "Mazda", "Mazda3"),
    ("mazda_mazda6_seeder.go", "Mazda", "Mazda6"),
    
    # Mercedes-Benz
    ("mercedes_benz_a_class_seeder.go", "Mercedes-Benz", "A-Class"),
    ("mercedes_benz_c_class_seeder.go", "Mercedes-Benz", "C-Class"),
    ("mercedes_benz_e_class_seeder.go", "Mercedes-Benz", "E-Class"),
    ("mercedes_benz_g_class_seeder.go", "Mercedes-Benz", "G-Class"),
    ("mercedes_benz_gla_seeder.go", "Mercedes-Benz", "GLA"),
    ("mercedes_benz_glc_seeder.go", "Mercedes-Benz", "GLC"),
    ("mercedes_benz_gle_seeder.go", "Mercedes-Benz", "GLE"),
    ("mercedes_benz_gls_seeder.go", "Mercedes-Benz", "GLS"),
    ("mercedes_benz_s_class_seeder.go", "Mercedes-Benz", "S-Class"),
    ("mercedes_benz_v_class_seeder.go", "Mercedes-Benz", "V-Class"),
    
    # MG
    ("mg_5_seeder.go", "MG", "5"),
    ("mg_6_seeder.go", "MG", "6"),
    
    # Mitsubishi
    ("mitsubishi_asx_seeder.go", "Mitsubishi", "ASX"),
    ("mitsubishi_lancer_seeder.go", "Mitsubishi", "Lancer"),
    ("mitsubishi_mirage_seeder.go", "Mitsubishi", "Mirage"),
    ("mitsubishi_montero_seeder.go", "Mitsubishi", "Montero"),
    ("mitsubishi_outlander_seeder.go", "Mitsubishi", "Outlander"),
    ("mitsubishi_pajero_seeder.go", "Mitsubishi", "Pajero"),
    ("mitsubishi_triton_seeder.go", "Mitsubishi", "Triton"),
    
    # Nissan
    ("nissan_juke_seeder.go", "Nissan", "Juke"),
    ("nissan_leaf_seeder.go", "Nissan", "Leaf"),
    ("nissan_march_seeder.go", "Nissan", "March"),
    ("nissan_navara_seeder.go", "Nissan", "Navara"),
    ("nissan_note_seeder.go", "Nissan", "Note"),
    ("nissan_sylphy_seeder.go", "Nissan", "Sylphy"),
    ("nissan_teana_seeder.go", "Nissan", "Teana"),
    ("nissan_x_trail_seeder.go", "Nissan", "X-Trail"),
    
    # Proton
    ("proton_persona_seeder.go", "Proton", "Persona"),
    ("proton_saga_seeder.go", "Proton", "Saga"),
    ("proton_x50_seeder.go", "Proton", "X50"),
    ("proton_x70_seeder.go", "Proton", "X70"),
    
    # Range Rover
    ("range_rover_autobiography_seeder.go", "Range Rover", "Autobiography"),
    ("range_rover_evoque_seeder.go", "Range Rover", "Evoque"),
    ("range_rover_sport_seeder.go", "Range Rover", "Sport"),
    ("range_rover_velar_seeder.go", "Range Rover", "Velar"),
    
    # SsangYong
    ("ssangyong_korando_seeder.go", "SsangYong", "Korando"),
    ("ssangyong_rexton_seeder.go", "SsangYong", "Rexton"),
    ("ssangyong_tivoli_seeder.go", "SsangYong", "Tivoli"),
    
    # Subaru
    ("subaru_forester_seeder.go", "Subaru", "Forester"),
    ("subaru_impreza_seeder.go", "Subaru", "Impreza"),
    ("subaru_legacy_seeder.go", "Subaru", "Legacy"),
    ("subaru_outback_seeder.go", "Subaru", "Outback"),
    ("subaru_xv_seeder.go", "Subaru", "XV"),
    
    # Suzuki
    ("suzuki_alto_seeder.go", "Suzuki", "Alto"),
    ("suzuki_baleno_seeder.go", "Suzuki", "Baleno"),
    ("suzuki_celerio_seeder.go", "Suzuki", "Celerio"),
    ("suzuki_ertiga_seeder.go", "Suzuki", "Ertiga"),
    ("suzuki_jimny_seeder.go", "Suzuki", "Jimny"),
    ("suzuki_swift_seeder.go", "Suzuki", "Swift"),
    ("suzuki_vitara_brezza_seeder.go", "Suzuki", "Vitara Brezza"),
    ("suzuki_wagon_r_seeder.go", "Suzuki", "Wagon R"),
    
    # Tata
    ("tata_altroz_seeder.go", "Tata", "Altroz"),
    ("tata_harrier_seeder.go", "Tata", "Harrier"),
    ("tata_nexon_seeder.go", "Tata", "Nexon"),
    ("tata_safari_seeder.go", "Tata", "Safari"),
    ("tata_tiago_seeder.go", "Tata", "Tiago"),
    ("tata_tigor_seeder.go", "Tata", "Tigor"),
    
    # Toyota
    ("toyota_allion_seeder.go", "Toyota", "Allion"),
    ("toyota_aqua_seeder.go", "Toyota", "Aqua"),
    ("toyota_axio_seeder.go", "Toyota", "Axio"),
    ("toyota_corolla_axio_seeder.go", "Toyota", "Corolla Axio"),
    ("toyota_corolla_seeder.go", "Toyota", "Corolla"),
    ("toyota_fortuner_seeder.go", "Toyota", "Fortuner"),
    ("toyota_harrier_seeder.go", "Toyota", "Harrier"),
    ("toyota_hiace_seeder.go", "Toyota", "Hiace"),
    ("toyota_land_cruiser_seeder.go", "Toyota", "Land Cruiser"),
    ("toyota_noah_seeder.go", "Toyota", "Noah"),
    ("toyota_passo_seeder.go", "Toyota", "Passo"),
    ("toyota_premio_seeder.go", "Toyota", "Premio"),
    ("toyota_prius_seeder.go", "Toyota", "Prius"),
    ("toyota_probox_seeder.go", "Toyota", "Probox"),
    ("toyota_rav4_seeder.go", "Toyota", "RAV4"),
    ("toyota_sienta_seeder.go", "Toyota", "Sienta"),
    ("toyota_vitz_seeder.go", "Toyota", "Vitz"),
    ("toyota_voxy_seeder.go", "Toyota", "Voxy"),
    
    # Volkswagen
    ("volkswagen_golf_seeder.go", "Volkswagen", "Golf"),
    ("volkswagen_passat_seeder.go", "Volkswagen", "Passat"),
    ("volkswagen_polo_seeder.go", "Volkswagen", "Polo"),
    ("volkswagen_tiguan_seeder.go", "Volkswagen", "Tiguan"),
    ("volkswagen_touareg_seeder.go", "Volkswagen", "Touareg"),
    ("volkswagen_vento_seeder.go", "Volkswagen", "Vento"),
    
    # Volvo
    ("volvo_s60_seeder.go", "Volvo", "S60"),
    ("volvo_s90_seeder.go", "Volvo", "S90"),
    ("volvo_xc40_seeder.go", "Volvo", "XC40"),
    ("volvo_xc60_seeder.go", "Volvo", "XC60"),
    ("volvo_xc90_seeder.go", "Volvo", "XC90"),
]

BASE_DIR = "internal/infrastructure/database/seeders"

def add_brand_model(filepath, brand, model):
    """Add brand and model to specifications map in a seeder file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Try pattern 1: specifications := map[string]string{\n\t\t"engine_type":
        pattern1 = r'(\tspecifications := map\[string\]string\{\n)(\t\t"engine_type":)'
        replacement1 = f'\\1\\t\\t"brand":                       "{brand}",\n\\t\\t"model":                       "{model}",\n\\2'
        new_content = re.sub(pattern1, replacement1, content)
        
        # If pattern1 didn't work, try pattern 2: specs := map[string]string{\n\t\t"
        if new_content == content:
            # Match any key after "specs := map[string]string{\n\t\t"
            pattern2 = r'(\tspecs := map\[string\]string\{\n)(\t\t")'
            replacement2 = f'\\1\\t\\t"Brand":                       "{brand}",\n\\t\\t"Model Name":                  "{model}",\n\\2'
            new_content = re.sub(pattern2, replacement2, content)
        
        if new_content == content:
            print(f"⚠️  No change for {filepath} - pattern not found")
            return False
        
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(new_content)
        
        print(f"[OK] Updated {filepath}")
        return True
    except FileNotFoundError:
        print(f"[FAIL] File not found: {filepath}")
        return False
    except Exception as e:
        print(f"[ERROR] Error processing {filepath}: {e}")
        return False

def main():
    success_count = 0
    fail_count = 0
    
    for filename, brand, model in cars:
        filepath = os.path.join(BASE_DIR, filename)
        if add_brand_model(filepath, brand, model):
            success_count += 1
        else:
            fail_count += 1
    
    print(f"\n{'='*60}")
    print(f"Summary: {success_count} files updated, {fail_count} files failed/skipped")
    print(f"{'='*60}")

if __name__ == "__main__":
    main()
