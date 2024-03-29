CREATE TABLE `memories` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `users_id` int(11) NOT NULL,
    `title` VARCHAR(50) NOT NULL,
    `date` DATE NOT NULL,
    `location` VARCHAR(500) NULL,
    `detail` VARCHAR(200) NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    CONSTRAINT `users_id` FOREIGN KEY `users_id` (`users_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
