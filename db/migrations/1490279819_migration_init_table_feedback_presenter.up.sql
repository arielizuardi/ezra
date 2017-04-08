CREATE TABLE `feedback_presenter` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '',
  `class_id` VARCHAR(45) NOT NULL COMMENT '',
  `session_id` INT NOT NULL COMMENT '',
  `presenter_id` INT NOT NULL COMMENT '',
  `participant_email` VARCHAR(200) NULL COMMENT '',
  `fields` JSON NULL COMMENT '',
  `created_at` DATETIME NULL COMMENT '',
  `updated_at` DATETIME NULL COMMENT '',
  `deleted_at` DATETIME NULL COMMENT '',
  PRIMARY KEY (`id`)  COMMENT '',
  INDEX `CLASS_SESSION_IDX` (`class_id` ASC, `session_id` ASC)  COMMENT '');
