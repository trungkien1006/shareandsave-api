CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE,
  `password` varchar(255),
  `avatar` longtext,
  `active` boolean,
  `full_name` varchar(64),
  `phone_number` varchar(16) UNIQUE,
  `address` text,
  `status` tinyint,
  `good_point` int DEFAULT 0,
  `major` varchar(64),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `item` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) UNIQUE,
  `description` text,
  `image` longtext,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `post` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `author_id` int,
  `item_id` int,
  `title` varchar(255),
  `description` text,
  `status` tinyint,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `warehouse` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `sku` varchar(255) UNIQUE,
  `quantity` int,
  `description` text,
  `classify` varchar(12),
  `stock_place` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `item_warehouse` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `item_id` int,
  `warehouse_id` int,
  `code` varchar(255) UNIQUE,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `send_request` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `type` int,
  `description` text,
  `status` tinyint,
  `reply_message` text,
  `appointment_time` datetime,
  `appointment_location` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `receive_request` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `type` int,
  `description` text,
  `status` tinyint,
  `item_warehouse_id` int DEFAULT null,
  `post_id` int DEFAULT null,
  `reply_message` text,
  `appointment_time` datetime,
  `appointment_location` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `role` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(64) UNIQUE,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `admin` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE,
  `password` varchar(255),
  `full_name` varchar(64),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp,
  `status` tinyint,
  `role_id` int
);

CREATE TABLE `permission` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) UNIQUE,
  `code` varchar(255) UNIQUE,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `role_permission` (
  `role_id` int,
  `permission_id` int
);

CREATE TABLE `notification` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `sender_id` int,
  `receiver_id` int,
  `title` varchar(255),
  `type` varchar(64),
  `target_type` varchar(32),
  `target_id` int,
  `content` varchar(255),
  `is_read` boolean,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `import_invoice` (
  `id` varchar(255) PRIMARY KEY,
  `invoice_num` int UNIQUE,
  `admin_id` int,
  `sender_id` int,
  `item_type` varchar(32),
  `send_date` datetime,
  `description` text,
  `is_lock` boolean,
  `is_anonymous` boolean,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `export_invoice` (
  `id` varchar(255) PRIMARY KEY,
  `invoice_num` int UNIQUE,
  `admin_id` int,
  `receiver_id` int,
  `item_type` varchar(32),
  `receive_date` datetime,
  `description` text,
  `is_lock` boolean,
  `is_anonymous` boolean,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `item_import_invoice` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `invoice_id` varchar(255),
  `item_id` int,
  `sku` varchar(255) UNIQUE,
  `quantity` tinyint,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `item_export_invoice` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `invoice_id` varchar(255),
  `item_warehouse_id` int,
  `sku` varchar(255) UNIQUE,
  `quantity` tinyint,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `image` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `target` varchar(32),
  `target_id` int,
  `image` longtext,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

ALTER TABLE `post` ADD FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);

ALTER TABLE `post` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `item_warehouse` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `item_warehouse` ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

ALTER TABLE `send_request` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `receive_request` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `receive_request` ADD FOREIGN KEY (`item_warehouse_id`) REFERENCES `item_warehouse` (`id`);

ALTER TABLE `receive_request` ADD FOREIGN KEY (`post_id`) REFERENCES `post` (`id`);

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

ALTER TABLE `item_export_invoice` ADD FOREIGN KEY (`item_warehouse_id`) REFERENCES `item_warehouse` (`id`);
