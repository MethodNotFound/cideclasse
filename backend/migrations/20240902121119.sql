-- Create "classes" table
CREATE TABLE `classes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` longtext NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_classes_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Modify "students" table
ALTER TABLE `students` ADD COLUMN `class_id` bigint unsigned NULL, ADD INDEX `fk_classes_students` (`class_id`), ADD CONSTRAINT `fk_classes_students` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
