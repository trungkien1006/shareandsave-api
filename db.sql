CREATE TABLE `user` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `email` VARCHAR(255) UNIQUE,
  `password` VARCHAR(255),
  `avatar` LONGTEXT,
  `active` BOOLEAN,
  `fullname` VARCHAR(64),
  `phone_number` VARCHAR(16) UNIQUE,
  `address` TEXT,
  `status` TINYINT,
  `good_point` INT DEFAULT 0,
  `major` VARCHAR(64),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `item` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255),
  `description` TEXT,
  `image` LONGTEXT,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `post` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `author_id` INT,
  `item_id` INT,
  `title` VARCHAR(255),
  `description` TEXT,
  `status` TINYINT,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `warehouse` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `sku` VARCHAR(255) UNIQUE,
  `quantity` INT,
  `description` TEXT,
  `classify` VARCHAR(12),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `item_warehouse` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `item_id` INT,
  `warehouse_id` INT,
  `code` VARCHAR(255),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `request` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT,
  `request_type` INT,
  `description` TEXT,
  `is_anonymous` BOOLEAN,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP,
  `status` TINYINT,
  `item` INT,
  `item_id` INT,
  `post_id` INT,
  `appointment_time` DATETIME,
  `appointment_location` VARCHAR(255)
);

CREATE TABLE `role` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(64),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `admin` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `email` VARCHAR(255),
  `password` VARCHAR(255),
  `fullname` VARCHAR(64),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP,
  `status` TINYINT,
  `role_id` INT
);

CREATE TABLE `permission` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255),
  `code` VARCHAR(255),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `role_permission` (
  `role_id` INT,
  `permission_id` INT
);

CREATE TABLE `notification` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `sender_id` INT,
  `receiver_id` INT,
  `title` VARCHAR(255),
  `type` VARCHAR(64),
  `target_type` VARCHAR(32),
  `target_id` INT,
  `content` VARCHAR(255),
  `is_read` BOOLEAN,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `import_invoice` (
  `id` VARCHAR(255) PRIMARY KEY,
  `invoice_num` INT,
  `admin_id` INT,
  `sender_id` INT,
  `item_type` VARCHAR(32),
  `send_date` DATETIME,
  `description` TEXT,
  `is_lock` BOOLEAN,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `export_invoice` (
  `id` VARCHAR(255) PRIMARY KEY,
  `invoice_num` INT,
  `admin_id` INT,
  `receiver_id` INT,
  `item_type` VARCHAR(32),
  `receive_date` DATETIME,
  `description` TEXT,
  `is_lock` BOOLEAN,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `item_import_invoice` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `invoice_id` VARCHAR(255),
  `item_id` INT,
  `sku` VARCHAR(255) UNIQUE,
  `quantity` TINYINT,
  `stock_place` VARCHAR(255) DEFAULT null,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `item_export_invoice` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `invoice_id` VARCHAR(255),
  `item_id` INT,
  `sku` VARCHAR(255) UNIQUE,
  `quantity` TINYINT,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

CREATE TABLE `image` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `target` VARCHAR(32),
  `target_id` INT,
  `image` LONGTEXT,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP
);

ALTER TABLE `post` ADD FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);

ALTER TABLE `post` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `item_warehouse` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `item_warehouse` ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

ALTER TABLE `request` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `request` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `request` ADD FOREIGN KEY (`post_id`) REFERENCES `post` (`id`);

ALTER TABLE `admin` ADD FOREIGN KEY (`role_id`) REFERENCES `role` (`id`);

ALTER TABLE `role_permission` ADD FOREIGN KEY (`role_id`) REFERENCES `role` (`id`);

ALTER TABLE `role_permission` ADD FOREIGN KEY (`permission_id`) REFERENCES `permission` (`id`);

ALTER TABLE `notification` ADD FOREIGN KEY (`sender_id`) REFERENCES `user` (`id`);

ALTER TABLE `notification` ADD FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`);

ALTER TABLE `import_invoice` ADD FOREIGN KEY (`admin_id`) REFERENCES `admin` (`id`);

ALTER TABLE `export_invoice` ADD FOREIGN KEY (`admin_id`) REFERENCES `admin` (`id`);

ALTER TABLE `export_invoice` ADD FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`);

ALTER TABLE `item_import_invoice` ADD FOREIGN KEY (`invoice_id`) REFERENCES `import_invoice` (`id`);

ALTER TABLE `item_import_invoice` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `item_export_invoice` ADD FOREIGN KEY (`invoice_id`) REFERENCES `export_invoice` (`id`);

ALTER TABLE `item_export_invoice` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);
