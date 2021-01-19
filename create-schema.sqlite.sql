PRAGMA read_uncommitted = 1;

CREATE TABLE IF NOT EXISTS `member` (
  `id` varchar(40) NOT NULL,
  `merchant_id` varchar(40) NOT NULL,
  `email` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  UNIQUE (`id`),
  UNIQUE (`merchant_id`,`email`)
);

CREATE TABLE IF NOT EXISTS `merchant` (
  `id` varchar(40) NOT NULL,
  `name` varchar(100) NOT NULL,
  `image` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  UNIQUE (`id`)
);
