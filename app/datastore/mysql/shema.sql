
DROP TABLE EXISTS `test_bear`;

CREATE TABLE `test_bear`(
    `id`VARCHAR(191) NOT NULL,
    `latitude` DECIMAL(9,6) NOT NULL
    `longitude` DECIMAL(9,6) NOT NULL
    `city` VARCHAR(255) NOT NULL 
    `address` VARCHAR(255) 
    `date` DATE NOT NULL 
    `time` TIME NOT NULL 
    `details` TEXT 
    PRIMARY KEY (id)
)ENGINE=InnoDB
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

