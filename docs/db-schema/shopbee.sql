-- Active: 1691472943542@@20.205.213.44@3306
-- DROP DATABASE --
DROP DATABASE IF EXISTS `shopbee`;

-- CREATE DATABASE --
CREATE DATABASE `shopbee`;

USE `shopbee`;

-- DROP EXISTED TABLE --
DROP TABLE IF EXISTS `images`;

DROP TABLE IF EXISTS `discounts`;

DROP TABLE IF EXISTS `wish_lists`;

DROP TABLE IF EXISTS `payments`;

DROP TABLE IF EXISTS `order_details`;

DROP TABLE IF EXISTS `orders`;

DROP TABLE IF EXISTS `carts`;

DROP TABLE IF EXISTS `shop_follows`;

DROP TABLE IF EXISTS `product_ratings`;

DROP TABLE IF EXISTS `request_upgrades`;

DROP TABLE IF EXISTS `products`;

DROP TABLE IF EXISTS `categories`;

DROP TABLE IF EXISTS `users`;

-- TABLE STRUCTURE --
CREATE TABLE `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `email` VARCHAR(50) UNIQUE NOT NULL,
    `fullname` VARCHAR(50) NOT NULL,
    `password` VARCHAR(50) NOT NULL,
    `salt` VARCHAR(50),
    `addr` VARCHAR(200),
    `phone` VARCHAR(20) UNIQUE,
    `role` ENUM (
        'buyer',
        'retailer',
        'admin',
        'moderator'
    ) NOT NULL DEFAULT "buyer",
    `avatar` JSON DEFAULT(
        '{"url": "http://d1851nciml9u0m.cloudfront.net/user/default-1691832193326062897.png",
        "width": 860,
        "height": 900,
        "cloud_name": "s3",
        "extension": ".png" }'
    ),
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `shop_follows` (
    `user_id` INT NOT NULL,
    `shop_id` INT NOT NULL,
    PRIMARY KEY (`user_id`, `shop_id`)
);

CREATE TABLE `categories` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(100) UNIQUE NOT NULL,
    `image` JSON,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `products` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `shop_id` INT NOT NULL,
    `category_id` INT NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `price` FLOAT NOT NULL,
    `description` TEXT,
    `quantity` INT NOT NULL,
    `condition` ENUM ('new', 'used') NOT NULL DEFAULT "new",
    `image` JSON,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `product_ratings` (
    `user_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `point` INT NOT NULL DEFAULT 5,
    `comment` TEXT,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    PRIMARY KEY (`user_id`, `product_id`)
);

CREATE TABLE `wish_lists` (
    `user_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    PRIMARY KEY (`user_id`, `product_id`)
);

CREATE TABLE `carts` (
    `user_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `product_quantity` INT DEFAULT 1,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    PRIMARY KEY (`user_id`, `product_id`)
);

CREATE TABLE `payments` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `payment_status` ENUM ('paid', 'pending') NOT NULL DEFAULT "pending",
    `payment_method` ENUM ('cod', 'card') NOT NULL DEFAULT "cod",
    `amount` FLOAT NOT NULL,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `orders` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `shop_id` INT NOT NULL,
    `payment_id` INT NOT NULL,
    `total_price` FLOAT NOT NULL,
    `shipping_addr` VARCHAR(200) NOT NULL,
    `order_status` ENUM (
        'pending',
        'confirm',
        'delivering',
        'completed'
    ) NOT NULL DEFAULT "pending",
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `order_details` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `order_id` INT NOT NULL,
    `product_origin` JSON NOT NULL,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `request_upgrades` (
    `user_id` INT PRIMARY KEY,
    `request_status` ENUM (
        'pending',
        'accepted',
        'denied'
    ) NOT NULL DEFAULT "pending",
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `report_accounts` (
    `user_id` INT NOT NULL,
    `shop_id` INT NOT NULL,
    `reason` TEXT,
    `report_status` ENUM ('pending', 'done') NOT NULL DEFAULT "pending",
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    PRIMARY KEY (`user_id`, `shop_id`)
);

CREATE TABLE `images` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `url` TEXT NOT NULL,
    `width` INT NOT NULL,
    `height` INT NOT NULL,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `discounts` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `code` VARCHAR(10) NOT NULL,
    `percentage` INT NOT NULL,
    `max_use` INT NOT NULL,
    `use_count` INT,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

--  INDEXING --
CREATE UNIQUE INDEX `users_index_0` ON `users` (`email`);

CREATE UNIQUE INDEX `categories_index_1` ON `categories` (`name`);

CREATE UNIQUE INDEX `discounts_index_2` ON `discounts` (`code`);

-- FOREIGN KEY --
ALTER TABLE
    `shop_follows`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `shop_follows`
ADD
    FOREIGN KEY (`shop_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `products`
ADD
    FOREIGN KEY (`shop_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `products`
ADD
    FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);

ALTER TABLE
    `product_ratings`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `product_ratings`
ADD
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE
    `wish_lists`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `wish_lists`
ADD
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE
    `carts`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `carts`
ADD
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE
    `orders`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `orders`
ADD
    FOREIGN KEY (`payment_id`) REFERENCES `payments` (`id`);

ALTER TABLE
    `order_details`
ADD
    FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE
    `payments`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `request_upgrades`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `report_accounts`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE
    `report_accounts`
ADD
    FOREIGN KEY (`shop_id`) REFERENCES `users` (`id`);

-- USER DATA INSERT --
INSERT INTO
    `users`
VALUES
    (
        1,
        'ptan21@clc.fitus.edu.vn',
        'Phan Thanh An',
        '2ef5a37dbe517c9d4dd2b246e8ff37d3',
        'oTAAdyawGUcHuyHWKCUjlzpfxfLvULDzQSWAxJpLzXxDuEqKCj',
        '227 Nguyễn Văn Cừ, P4, Q5, TP.HCM',
        '0325364474',
        'admin',
        '{\"url\": \"http://d1851nciml9u0m.cloudfront.net/user/default-1691832193326062897.png\", \"width\": 860, \"height\": 900, \"extension\": \".png\", \"cloud_name\": \"s3\"}',
        1,
        '2023-08-12 18:22:11',
        '2023-08-12 18:22:11'
    ),
    (
        2,
        'nhhuu21@clc.fitus.edu.vn',
        'Nguyễn Hi Hữu',
        '2ef5a37dbe517c9d4dd2b246e8ff37d3',
        'oTAAdyawGUcHuyHWKCUjlzpfxfLvULDzQSWAxJpLzXxDuEqKCj',
        '227 Nguyễn Văn Cừ, P4, Q5, TP.HCM',
        '0325364475',
        'admin',
        '{\"url\": \"http://d1851nciml9u0m.cloudfront.net/user/default-1691832193326062897.png\", \"width\": 860, \"height\": 900, \"extension\": \".png\", \"cloud_name\": \"s3\"}',
        1,
        '2023-08-12 18:22:11',
        '2023-08-12 18:22:11'
    ),
    (
        3,
        'phgbao21@clc.fitus.edu.vn',
        'Phạm Hồng Gia Bảo',
        '2ef5a37dbe517c9d4dd2b246e8ff37d3',
        'oTAAdyawGUcHuyHWKCUjlzpfxfLvULDzQSWAxJpLzXxDuEqKCj',
        '227 Nguyễn Văn Cừ, P4, Q5, TP.HCM',
        '0325364476',
        'admin',
        '{\"url\": \"http://d1851nciml9u0m.cloudfront.net/user/default-1691832193326062897.png\", \"width\": 860, \"height\": 900, \"extension\": \".png\", \"cloud_name\": \"s3\"}',
        1,
        '2023-08-12 18:22:11',
        '2023-08-12 18:22:11'
    ),
    (
        4,
        'ntthuan20@clc.fitus.edu.vn',
        'Nguyễn Thành Thuuận',
        '2ef5a37dbe517c9d4dd2b246e8ff37d3',
        'oTAAdyawGUcHuyHWKCUjlzpfxfLvULDzQSWAxJpLzXxDuEqKCj',
        '227 Nguyễn Văn Cừ, P4, Q5, TP.HCM',
        '0325364478',
        'admin',
        '{\"url\": \"http://d1851nciml9u0m.cloudfront.net/user/default-1691832193326062897.png\", \"width\": 860, \"height\": 900, \"extension\": \".png\", \"cloud_name\": \"s3\"}',
        1,
        '2023-08-12 18:22:11',
        '2023-08-12 18:22:11'
    );

-- CATEGORY DATA INSERT --
INSERT INTO
    `categories` (id, name, image)
VALUES
    (
        1,
        'Information Technology',
        '{
            \"id\": 0, 
            \"url\": \"http://d1851nciml9u0m.cloudfront.net/category/it-1690871711814317800.jpg\", 
            \"width\": 507, 
            \"height\": 690, 
            \"extension\": \".jpg\", 
            \"cloud_name\": \"s3\"
        }'
    ),
    (
        2,
        'Physic',
        '{
            \"id\": 0, 
            \"url\": \"http://d1851nciml9u0m.cloudfront.net/category/phys-1690871780988147100.jpg\", 
            \"width\": 800, 
            \"height\": 1207, 
            \"extension\": \".jpg\", 
            \"cloud_name\": \"s3\"
        }'
    ),
    (
        3,
        'Philosophy',
        '{
            \"id\": 0, 
            \"url\": \"http://d1851nciml9u0m.cloudfront.net/category/phil-1690871806137092000.jpg\", 
            \"width\": 800, 
            \"height\": 800, 
            \"extension\": \".jpg\", 
            \"cloud_name\": \"s3\"
        }'
    ),
    (
        4,
        'Chemistry ',
        '{
            \"id\": 0, 
            \"url\": \"http://d1851nciml9u0m.cloudfront.net/category/chem-1690871829236975500.jpg\", 
            \"width\": 900, 
            \"height\": 900, 
            \"extension\": \".jpg\", 
            \"cloud_name\": \"s3\"
        }'
    ),
    (
        5,
        'Electronics',
        '{
            \"id\": 0, 
            \"url\": \"http://d1851nciml9u0m.cloudfront.net/category/elec-1690871848792890800.jpg\", 
            \"width\": 545, 
            \"height\": 763, 
            \"extension\": \".jpg\", 
            \"cloud_name\": \"s3\"
        }'
    ),
    (
        6,
        'Economics',
        '{
            \"id\": 0, 
            \"url\": \"http://d1851nciml9u0m.cloudfront.net/category/eco-1690871880562690000.jpg\", 
            \"width\": 1803, 
            \"height\": 2814, 
            \"extension\": \".jpg\", 
            \"cloud_name\": \"s3\"
        }'
    ),
    (
        7,
        'Biology ',
        '{
            \"id\": 0, 
            \"url\": \"http://d1851nciml9u0m.cloudfront.net/category/bio-1690871901477291100.jpg\", 
            \"width\": 770, 
            \"height\": 1031, 
            \"extension\": \".jpg\", 
            \"cloud_name\": \"s3\"
        }'
    ),
    (
        8,
        'Mathematics',
        '{
            \"id\": 0, 
            \"url\": \"http://d1851nciml9u0m.cloudfront.net/category/math-1690871922855668600.jpg\", 
            \"width\": 800, 
            \"height\": 1143, 
            \"extension\": \".jpg\", 
            \"cloud_name\": \"s3\"
        }'
    );