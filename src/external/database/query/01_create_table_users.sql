CREATE TABLE IF NOT EXISTS `users`(
    `id` BIGINT UNSIGNED NULL AUTO_INCREMENT,
    `name` CHAR(255) NOT NULL,
    `username` CHAR(255) NOT NULL,
    `email` CHAR(255) NOT NULL,
    `password` CHAR(255) NOT NULL,
    `role` CHAR(255) NOT NULL,
    `create_at` BIGINT UNSIGNED,
    `update_at` BIGINT UNSIGNED,
    `delete_at` BIGINT UNSIGNED,
    CONSTRAINT column_name_can_not_be_empty CHECK(`name` > ''),
    CONSTRAINT column_email_not_valid CHECK(`email` LIKE '%@%.%'),
    CONSTRAINT column_password_not_valid CHECK(`password` > ''),
    PRIMARY KEY(`id`),
    UNIQUE (`username`),
    UNIQUE (`email`),
    INDEX `name` (`name`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci ENGINE=INNODB;