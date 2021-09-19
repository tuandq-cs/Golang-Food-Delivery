CREATE TABLE `carts` (
                         `user_id` int NOT NULL,
                         `food_id` int NOT NULL,
                         `quantity` int NOT NULL,
                         `status` int NOT NULL DEFAULT '1',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`user_id`,`food_id`),
                         KEY `food_id` (`food_id`)
) ENGINE=InnoDB;

CREATE TABLE `categories` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `name` varchar(100) NOT NULL,
                              `description` text,
                              `icon` json DEFAULT NULL,
                              `status` int NOT NULL DEFAULT '1',
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `cities` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `title` varchar(100) NOT NULL,
                          `status` int NOT NULL DEFAULT '1',
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `food_likes` (
                              `user_id` int NOT NULL,
                              `food_id` int NOT NULL,
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`user_id`,`food_id`),
                              KEY `food_id` (`food_id`)
) ENGINE=InnoDB;

CREATE TABLE `food_ratings` (
                                `id` int NOT NULL AUTO_INCREMENT,
                                `user_id` int NOT NULL,
                                `food_id` int NOT NULL,
                                `point` float DEFAULT '0',
                                `comment` text,
                                `status` int NOT NULL DEFAULT '1',
                                `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                PRIMARY KEY (`id`),
                                KEY `food_id` (`food_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `foods` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `restaurant_id` int NOT NULL,
                         `category_id` int DEFAULT NULL,
                         `name` varchar(255) NOT NULL,
                         `description` text,
                         `price` float NOT NULL,
                         `images` json NOT NULL,
                         `status` int NOT NULL DEFAULT '1',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         KEY `restaurant_id` (`restaurant_id`) USING BTREE,
                         KEY `category_id` (`category_id`) USING BTREE,
                         KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `images` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `file_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `width` int NOT NULL,
                          `height` int NOT NULL,
                          `status` int NOT NULL DEFAULT '1',
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `order_details` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `order_id` int NOT NULL,
                                 `food_origin` json DEFAULT NULL,
                                 `price` float NOT NULL,
                                 `quantity` int NOT NULL,
                                 `discount` float DEFAULT '0',
                                 `status` int NOT NULL DEFAULT '1',
                                 `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `order_trackings` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `order_id` int NOT NULL,
                                   `state` enum('waiting_for_shipper','preparing','on_the_way','delivered','cancel') NOT NULL,
                                   `status` int NOT NULL DEFAULT '1',
                                   `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                   `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                   PRIMARY KEY (`id`),
                                   KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `orders` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `user_id` int NOT NULL,
                          `total_price` float NOT NULL,
                          `shipper_id` int DEFAULT NULL,
                          `status` int NOT NULL DEFAULT '1',
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          KEY `user_id` (`user_id`) USING BTREE,
                          KEY `shipper_id` (`shipper_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `restaurant_foods` (
                                    `restaurant_id` int NOT NULL,
                                    `food_id` int NOT NULL,
                                    `status` int NOT NULL DEFAULT '1',
                                    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`restaurant_id`,`food_id`),
                                    KEY `food_id` (`food_id`)
) ENGINE=InnoDB;

CREATE TABLE `restaurant_likes` (
                                    `restaurant_id` int NOT NULL,
                                    `user_id` int NOT NULL,
                                    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`restaurant_id`,`user_id`),
                                    KEY `user_id` (`user_id`)
) ENGINE=InnoDB;

CREATE TABLE `restaurant_ratings` (
                                      `id` int NOT NULL AUTO_INCREMENT,
                                      `user_id` int NOT NULL,
                                      `restaurant_id` int NOT NULL,
                                      `point` float NOT NULL DEFAULT '0',
                                      `comment` text,
                                      `status` int NOT NULL DEFAULT '1',
                                      `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                      `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      PRIMARY KEY (`id`),
                                      KEY `user_id` (`user_id`) USING BTREE,
                                      KEY `restaurant_id` (`restaurant_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `restaurants` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `owner_id` int NOT NULL,
                               `name` varchar(50) NOT NULL,
                               `addr` varchar(255) NOT NULL,
                               `city_id` int DEFAULT NULL,
                               `lat` double DEFAULT NULL,
                               `lng` double DEFAULT NULL,
                               `cover` json NOT NULL,
                               `logo` json NOT NULL,
                               `shipping_fee_per_km` double DEFAULT '0',
                               `status` int NOT NULL DEFAULT '1',
                               `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               KEY `owner_id` (`owner_id`) USING BTREE,
                               KEY `city_id` (`city_id`) USING BTREE,
                               KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `user_addresses` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `user_id` int NOT NULL,
                                  `city_id` int NOT NULL,
                                  `title` varchar(100) DEFAULT NULL,
                                  `icon` json DEFAULT NULL,
                                  `addr` varchar(255) NOT NULL,
                                  `lat` double DEFAULT NULL,
                                  `lng` double DEFAULT NULL,
                                  `status` int NOT NULL DEFAULT '1',
                                  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`),
                                  KEY `user_id` (`user_id`) USING BTREE,
                                  KEY `city_id` (`city_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `user_device_tokens` (
                                      `id` int unsigned NOT NULL AUTO_INCREMENT,
                                      `user_id` int unsigned DEFAULT NULL,
                                      `is_production` tinyint(1) DEFAULT '0',
                                      `os` enum('ios','android','web') DEFAULT 'ios' COMMENT '1: iOS, 2: Android',
                                      `token` varchar(255) DEFAULT NULL,
                                      `status` smallint unsigned NOT NULL DEFAULT '1',
                                      `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                      PRIMARY KEY (`id`),
                                      KEY `user_id` (`user_id`) USING BTREE,
                                      KEY `os` (`os`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `email` varchar(50) NOT NULL,
                         `fb_id` varchar(50) DEFAULT NULL,
                         `gg_id` varchar(50) DEFAULT NULL,
                         `password` varchar(50) NOT NULL,
                         `salt` varchar(50) DEFAULT NULL,
                         `last_name` varchar(50) NOT NULL,
                         `first_name` varchar(50) NOT NULL,
                         `phone` varchar(20) DEFAULT NULL,
                         `role` enum('user','admin','shipper') NOT NULL DEFAULT 'user',
                         `avatar` json DEFAULT NULL,
                         `status` int NOT NULL DEFAULT '1',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB;