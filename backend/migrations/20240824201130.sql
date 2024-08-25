-- Create "logins" table
CREATE TABLE `logins` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `identificator` longtext NOT NULL,
  `password_hash` longtext NOT NULL,
  `type` longtext NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_logins_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
