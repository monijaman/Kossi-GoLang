-- TV Specifications Only (New additions, excluding existing specs)
-- These are the TV-specific specifications to be inserted
-- Generated to exclude all 491 existing specifications

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Inserting only TV-specific specifications that don't exist in the main list
INSERT INTO "specification_keys" ("specification_key") VALUES
	('Screen Size'),
	('Resolution'),
	('Panel Type'),
	('Refresh Rate'),
	('Brightness (Nits)'),
	('Contrast Ratio'),
	('Color Accuracy'),
	('Viewing Angle'),
	('Response Time'),
	('HDR Support'),
	('Dolby Vision'),
	('Local Dimming'),
	('Backlight Type'),
	('Smart TV OS'),
	('Built-in Apps'),
	('WiFi Connectivity'),
	('Bluetooth Connectivity'),
	('HDMI Ports'),
	('USB Ports'),
	('Audio Output (Watts)'),
	('Speaker Configuration'),
	('Dolby Atmos Support'),
	('Built-in Tuner'),
	('Supported Video Formats'),
	('Supported Audio Formats'),
	('Frame Rate Support'),
	('ALLM Support'),
	('VRR Support'),
	('Game Mode'),
	('Motion Smoothing'),
	('Picture Quality Modes'),
	('Energy Consumption (Watts)'),
	('Power Consumption Standby (Watts)'),
	('Dimensions (W x H x D)'),
	('Weight Without Stand (kg)'),
	('Weight With Stand (kg)'),
	('VESA Mount Pattern'),
	('Warranty Period (Years)'),
	('Warranty Coverage'),
	('Lifespan (Hours)'),
	('Thickness'),
	('Bezel Width'),
	('Remote Type'),
	('Voice Control Support'),
	('Gaming Features'),
	('Sports Mode'),
	('Movie Mode'),
	('Eye Comfort'),
	('Flicker-Free'),
	('Blue Light Filter'),
	('Eco Mode'),
	('Auto Brightness'),
	('Motion Interpolation'),
	('Upscaling Technology'),
	('Color Enhancement'),
	('Price (BDT)'),
	('Availability in Bangladesh'),
	('After-Sales Service'),
	('Return / Exchange Policy');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
