CREATE TABLE `ss_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `height` bigint(20) NOT NULL DEFAULT '0',
  `old_height` bigint(20) NOT NULL DEFAULT '0',
  `stratum_server_url` varchar(255) NOT NULL DEFAULT '',
  `type` varchar(255) NOT NULL DEFAULT '',
  `prev_hash` varchar(255) DEFAULT NULL,
  `username` varchar(50) NOT NULL DEFAULT '',
  `coin_type` varchar(50) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `notified_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `height` (`height`),
  KEY `stratum_server_url` (`stratum_server_url`),
  KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;