-- Modify "classes" table
ALTER TABLE `classes` MODIFY COLUMN `name` varchar(191) NOT NULL, ADD UNIQUE INDEX `idx_classes_name` (`name`);
-- Modify "students" table
ALTER TABLE `students` DROP COLUMN `class_id`, DROP FOREIGN KEY `fk_classes_students`;
-- Create "user_classes" table
CREATE TABLE `user_classes` (
  `student_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`student_id`, `class_id`),
  INDEX `fk_user_classes_class` (`class_id`),
  CONSTRAINT `fk_user_classes_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_user_classes_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
