CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `role_id` int,
  `email` varchar(255) UNIQUE,
  `phone_number` varchar(16) UNIQUE,
  `password` varchar(255),
  `avatar` longtext,
  `active` boolean,
  `full_name` varchar(64),
  `address` text,
  `status` tinyint,
  `good_point` int DEFAULT 0,
  `major` varchar(64),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `category` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `item` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `category_id` int,
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
  `type` int,
  `slug` varchar(255) UNIQUE,
  `title` varchar(255),
  `content` json DEFAULT null,
  `info` json DEFAULT null,
  `description` mediumtext DEFAULT null,
  `image` json DEFAULT null,
  `tag` json DEFAULT null,
  `status` tinyint,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `post_item` (
  `post_id` int,
  `item_id` int,
  `quantity` int,
  `image` longtext
);

CREATE TABLE `interest` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `post_id` int,
  `status` int,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `comment` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `interest_id` int,
  `sender_id` int,
  `receiver_id` int,
  `content` text,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `transaction` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `interest_id` int,
  `sender_id` int,
  `receiver_id` int,
  `status` int,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `transaction_item` (
  `transaction_id` int,
  `item_id` int,
  `quantity` int
);

CREATE TABLE `appointment` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `scheduled_time` datetime,
  `status` int,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `appointment_item_warehouse` (
  `appointment_id` int,
  `item_warehouse_id` int
);

CREATE TABLE `warehouse` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `item_id` int,
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
  `description` text,
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
  `sender_id` int,
  `receiver_id` int,
  `classify` varchar(32),
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
  `sender_id` int,
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
  `quantity` tinyint,
  `description` text,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `item_export_invoice` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `invoice_id` varchar(255),
  `item_warehouse_id` int,
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

CREATE TABLE `setting` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `key` varchar(255),
  `value` mediumtext,
  `created_at` timestamp,
  `updated_at` timestamp,
  `deleted_at` timestamp
);

ALTER TABLE `warehouse` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `user` ADD FOREIGN KEY (`role_id`) REFERENCES `role` (`id`);

ALTER TABLE `item` ADD FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);

ALTER TABLE `post` ADD FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);

ALTER TABLE `post_item` ADD FOREIGN KEY (`post_id`) REFERENCES `post` (`id`);

ALTER TABLE `post_item` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `interest` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `interest` ADD FOREIGN KEY (`post_id`) REFERENCES `post` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`interest_id`) REFERENCES `interest` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`sender_id`) REFERENCES `user` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`interest_id`) REFERENCES `interest` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`sender_id`) REFERENCES `user` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`);

ALTER TABLE `transaction_item` ADD FOREIGN KEY (`transaction_id`) REFERENCES `transaction` (`id`);

ALTER TABLE `transaction_item` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `appointment` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `appointment_item_warehouse` ADD FOREIGN KEY (`appointment_id`) REFERENCES `appointment` (`id`);

ALTER TABLE `appointment_item_warehouse` ADD FOREIGN KEY (`item_warehouse_id`) REFERENCES `item_warehouse` (`id`);

ALTER TABLE `item_warehouse` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `item_warehouse` ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

ALTER TABLE `role_permission` ADD FOREIGN KEY (`role_id`) REFERENCES `role` (`id`);

ALTER TABLE `role_permission` ADD FOREIGN KEY (`permission_id`) REFERENCES `permission` (`id`);

ALTER TABLE `notification` ADD FOREIGN KEY (`sender_id`) REFERENCES `user` (`id`);

ALTER TABLE `notification` ADD FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`);

ALTER TABLE `export_invoice` ADD FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`);

ALTER TABLE `item_import_invoice` ADD FOREIGN KEY (`invoice_id`) REFERENCES `import_invoice` (`id`);

ALTER TABLE `item_import_invoice` ADD FOREIGN KEY (`item_id`) REFERENCES `item` (`id`);

ALTER TABLE `item_export_invoice` ADD FOREIGN KEY (`invoice_id`) REFERENCES `export_invoice` (`id`);

ALTER TABLE `item_export_invoice` ADD FOREIGN KEY (`item_warehouse_id`) REFERENCES `item_warehouse` (`id`);
