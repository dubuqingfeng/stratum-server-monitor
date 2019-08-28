CREATE TABLE `ss_servers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `region` varchar(50) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `pool` varchar(255) DEFAULT NULL,
  `addresses` varchar(255) NOT NULL DEFAULT '[]',
  `description` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `region` (`region`),
  KEY `name` (`name`),
  KEY `description` (`description`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

CREATE TABLE `ss_params` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `server_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL DEFAULT '',
  `coin` varchar(20) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `type` varchar(255) NOT NULL DEFAULT '',
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `coinbase` varchar(255) NOT NULL DEFAULT '',
  `coinbase_tags` varchar(255) NOT NULL DEFAULT '',
  `extra` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `server_id` (`server_id`),
  KEY `description` (`description`),
  KEY `coin` (`coin`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;