CREATE DATABASE IF NOT EXISTS bear_sighting_db
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE bear_sighting_db;

CREATE TABLE sightings (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Unique sighting ID',
    latitude DECIMAL(9,6) NOT NULL COMMENT 'Latitude of the sighting location',
    longitude DECIMAL(9,6) NOT NULL COMMENT 'Longitude of the sighting location',
    city VARCHAR(255) NOT NULL COMMENT 'City or municipality of the sighting',
    area VARCHAR(255) COMMENT 'Specific area or district where the sighting occurred',
    date DATE NOT NULL COMMENT 'Date of the sighting',
    time TIME NOT NULL COMMENT 'Time of the sighting',
    situation TEXT NOT NULL COMMENT 'Details about the sighting situation',
    details TEXT COMMENT 'Additional details or description',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    PRIMARY KEY (id)
) ENGINE=InnoDB
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

-- Optional: Add indexes for common queries
CREATE INDEX idx_latitude_longitude ON sightings (latitude, longitude);
CREATE INDEX idx_date ON sightings (date);
CREATE INDEX idx_city ON sightings (city);
