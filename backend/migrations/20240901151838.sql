-- Create "students" table
CREATE TABLE `students` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` longtext NOT NULL,
  `email` longtext NULL,
  `identifier` varchar(191) NOT NULL,
  `password_hash` longtext NULL,
  `ask_new_password` bool NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_students_deleted_at` (`deleted_at`),
  UNIQUE INDEX `idx_students_identifier` (`identifier`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "sessions" table
CREATE TABLE `sessions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `active` bool NOT NULL,
  `meta_data` longtext NULL,
  `student_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_students_sessions` (`student_id`),
  INDEX `idx_sessions_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_students_sessions` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
