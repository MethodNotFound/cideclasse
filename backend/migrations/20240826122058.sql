-- Create "students" table
CREATE TABLE `students` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` longtext NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_students_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "logins" table
CREATE TABLE `logins` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `identificator` varchar(191) NOT NULL,
  `password_hash` longtext NOT NULL,
  `student_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_logins_student` (`student_id`),
  INDEX `idx_logins_deleted_at` (`deleted_at`),
  UNIQUE INDEX `idx_logins_identificator` (`identificator`),
  CONSTRAINT `fk_logins_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
