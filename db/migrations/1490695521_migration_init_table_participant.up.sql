CREATE TABLE `jpcccol_db`.`participant` (
  `email` varchar(200) NOT NULL DEFAULT '',
  `name` varchar(200) DEFAULT NULL,
  `dob` varchar(20) DEFAULT NULL,
  `date` varchar(200) DEFAULT NULL,
  `phone_number` varchar(30) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`email`),
  FULLTEXT KEY `NAME_IDX` (`name`),
  FULLTEXT KEY `DATE_IDX` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
