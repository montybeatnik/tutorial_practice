-- all
SELECT devices.id, devices.created_at, hostname, loopback,  hardware.vendor, hardware.model, software.version from devices
	JOIN hardware on hardware.id = hardware_id
	JOIN software on software.id = software_id;
-- ByID
SELECT devices.id, devices.created_at, hostname, loopback,  hardware.vendor, hardware.model, software.version from devices
	JOIN hardware on hardware.id = hardware_id
	JOIN software on software.id = software_id
     where devices.id = 5;