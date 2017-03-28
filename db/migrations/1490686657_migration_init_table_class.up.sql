CREATE TABLE `jpcccol_db`.`class` (
  `id` varchar(50) NOT NULL DEFAULT '',
  `name` varchar(100) DEFAULT '',
  `batch` int(11) DEFAULT NULL,
  `year` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
