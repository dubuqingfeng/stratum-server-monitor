/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table btc_pool_ss_heights
# ------------------------------------------------------------

DROP TABLE IF EXISTS `btc_pool_ss_heights`;

CREATE TABLE `btc_pool_ss_heights` (
  `height` bigint(20) NOT NULL DEFAULT '0',
  `stratum_server_url` varchar(255) NOT NULL DEFAULT '',
  `prev_hash` varchar(255) NOT NULL DEFAULT '',
  `type` varchar(255) NOT NULL DEFAULT '',
  `username` varchar(50) NOT NULL DEFAULT '',
  `coin_type` varchar(50) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `notified_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  UNIQUE KEY `stratum_coin_url` (`stratum_server_url`,`type`,`coin_type`),
  KEY `height` (`height`),
  KEY `stratum_server_url` (`stratum_server_url`),
  KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

# Dump of table btc_pool_ss_params
# ------------------------------------------------------------

DROP TABLE IF EXISTS `btc_pool_ss_params`;

CREATE TABLE `btc_pool_ss_params` (
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

LOCK TABLES `btc_pool_ss_params` WRITE;
/*!40000 ALTER TABLE `btc_pool_ss_params` DISABLE KEYS */;

INSERT INTO `btc_pool_ss_params` (`id`, `server_id`, `name`, `coin`, `description`, `type`, `username`, `password`, `coinbase`, `coinbase_tags`, `extra`, `created_at`, `updated_at`)
VALUES
	(1,1,'BTC.com北京BCH测试','bch','','','dubuqingfeng','md=2048','76a914da5b5f7945660ea237a205c3b24ba47c1dfbf93d88ac','{\"nmc\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(2,1,'BTC.com北京测试','btc','','','11fdsf','md=2048','001497cfc76442fe717f2a3f0cc9c175f7561b661997','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(3,2,'火币南方测试','btc','','','zqw','md=2048','76a914e582933875bedfdc448473c00b474f8f053a467588ac','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(4,3,'火币北方测试','btc','','','yzd','md=2048','76a914e582933875bedfdc448473c00b474f8f053a467588ac','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(5,4,'F2Pool','btc','','','dubuqingfeng.002','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(6,5,'Poolin','btc','','','youlinf.001','123','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(7,6,'Antpool','btc','','','testbcc.01','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(8,7,'BTC.top','btc','','','dubuqingfeng.213','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(9,8,'viabtc','btc','','','dubuqingfeng.87','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(10,9,'SlushPool','btc','','','dubuqingfeng.123','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(11,10,'BTC.com深圳测试','btc','','','xxvdfw.145','md=2048','001497cfc76442fe717f2a3f0cc9c175f7561b661997','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(12,11,'AntpoolDCR','dcr','','','DskcowNj5yinRLCchjfgnnt4mMJt27irBf1','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(13,12,'BTC.com北京DCR','dcr','','','dubuqingfeng.1234','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(14,13,'BTC.com北京LTC','ltc','','','testltc.41','','','{\"mm\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(15,14,'BTC.com深圳LTC','ltc','','','qjq93yj.001','','','{\"mm\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(16,15,'AntpoolLTC','ltc','','','apltc.10x32','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(17,16,'PoolinLTC','ltc','','','testtestltc.001','123','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(18,17,'F2poolLTC','ltc','','','dubuqingfeng.001','213123','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(19,18,'BTC.com欧洲测试','btc','','','testeueuueuueu.001','md=2048','001497cfc76442fe717f2a3f0cc9c175f7561b661997','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(20,19,'BTC.com美国测试','btc','','','meiguotest.001','md=2048','001497cfc76442fe717f2a3f0cc9c175f7561b661997','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(21,20,'BTC.com美国DCR','dcr','','','testtesttestgx.001','md=2048','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(22,21,'Zhizhu北方BTC','btc','','','zhizhu.001','md=2048','a9144f0dd555d6f74016b8c57b265b17615dc092fa7687','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(23,22,'Zhizhu南方BTC','btc','','','hangzhou.001','md=2048','a9144f0dd555d6f74016b8c57b265b17615dc092fa7687','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(24,23,'火币美国测试','btc','','','qzs.0012','md=2048','76a914e582933875bedfdc448473c00b474f8f053a467588ac','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(25,24,'火币欧洲测试','btc','','','qwr.0012','md=2048','76a914e582933875bedfdc448473c00b474f8f053a467588ac','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(26,27,'NovaBlock','btc','','','dubuqingfeng.001','123','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(27,28,'WayiLTC','ltc','','','dubuqingfeng.001','123','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(29,30,'火币通用测试','btc','','','xay.0012','md=2048','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(30,31,'火币香港测试','btc','','','q23.0012','md=2048','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(31,10,'BTC.com南方BCH测试','bch','','','7db2mqk','md=2048','76a914da5b5f7945660ea237a205c3b24ba47c1dfbf93d88ac','{\"nmc\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(32,18,'BTC.com欧洲BCH测试','bch','','','6xefmrz.001','md=2048','76a914da5b5f7945660ea237a205c3b24ba47c1dfbf93d88ac','{\"nmc\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(33,19,'BTC.com美国BCH测试','bch','','','m3nsz7x.001','md=2048','76a914da5b5f7945660ea237a205c3b24ba47c1dfbf93d88ac','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(34,25,'BTC.com欧洲LTC','ltc','','','etjpmvq.001','md=2048','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(35,26,'BTC.com美国LTC','ltc','','','j74brke.001','md=2048','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(36,3,'火币北方BCH测试','bch','','','gsf','md=2048','','{\"nmc\":\"fabe6d6d\",\"vcash\":\"b9e11b6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(37,2,'火币南方BCH测试','bch','','','qgxb','md=2048','','{\"nmc\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(38,33,'火币北方LTC测试','ltc','','','qsxv','md=2048','','{\"mm\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(39,32,'火币南方LTC测试','ltc','','','zwr','md=2048','','{\"mm\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(40,34,'火币香港LTC测试','ltc','','','usltc','md=2048','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(41,35,'火币美国LTC测试','ltc','','','usltchaha','md=2048','','{\"mm\":\"fabe6d6d\"}','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(42,36,'火币北京DCR','dcr','','','qxvb.1234','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(43,37,'PoolinDCR','dcr','','','dcredtest.1234','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(44,39,'PoolinBCH','bch','','','testbch.1234','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(45,38,'火币南方DCR','dcr','','','dcrsz.1234','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(46,40,'火币美国DCR','dcr','','','dcrus.1234','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(47,5,'Poolin','btc','','','btcbf01.001','123','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(48,41,'BTC.com北京CKB','ckb','','','ckb.41','','','','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(49,42,'BTC.comBitdeerBTC','btc','','','Bitdeertest001123.test','','001497cfc76442fe717f2a3f0cc9c175f7561b661997','','','2019-07-16 14:17:55','2019-07-16 14:17:55');

/*!40000 ALTER TABLE `btc_pool_ss_params` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table btc_pool_ss_servers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `btc_pool_ss_servers`;

CREATE TABLE `btc_pool_ss_servers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `region` varchar(50) NOT NULL DEFAULT '',
  `coin_type` varchar(50) DEFAULT NULL,
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

LOCK TABLES `btc_pool_ss_servers` WRITE;
/*!40000 ALTER TABLE `btc_pool_ss_servers` DISABLE KEYS */;

INSERT INTO `btc_pool_ss_servers` (`id`, `region`, `coin_type`, `name`, `pool`, `addresses`, `description`, `created_at`, `updated_at`)
VALUES
	(1,'bj','btc','BTC.comBTC北京公共服务器','BTC.com','[\"cn.ss.btc.com:1800\",\"cn.ss.btc.com:443\",\"cn.ss.btc.com:25\"]','{\"pool\":\"BTC.com\"}','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(2,'hz','btc','火币BTC南方公共服务器','HuobiPool','[\"bs.huobipool.com:1800\",\"bs.huobipool.com:443\"]','{\"pool\":\"HuobiPool\"}','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(3,'bj','btc','火币BTC北方公共服务器','HuobiPool','[\"bn.huobipool.com:1800\",\"bn.huobipool.com:443\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(4,'btc','btc','F2poolBTC公共服务器','F2Pool','[\"btc.f2pool.com:1314\",\"btc.f2pool.com:3333\",\"btc.f2pool.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(5,'btc','btc','PoolinBTC公共服务器','Poolin','[\"btc.ss.poolin.com:443\",\"btc.ss.poolin.com:1883\",\"btc.ss.poolin.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(6,'btc','btc','AntpoolBTC公共服务器','Antpool','[\"stratum.antpool.com:3333\",\"stratum.antpool.com:443\",\"stratum.antpool.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(7,'btc','btc','BTC.topBTC公共服务器','BTC.top','[\"stratum.btc.top:8888\",\"bak.btc.top:3333\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(8,'btc','btc','ViabtcBTC公共服务器','Viabtc','[\"btc.viabtc.com:3333\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(9,'btc','btc','SlushPoolBTC亚洲公共服务器','SlushPool','[\"cn.stratum.slushpool.com:3333\",\"cn.stratum.slushpool.com:443\",\"cn02.stratum.slushpool.com:443\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(10,'sz','btc','BTC.comBTC深圳公共服务器','BTC.com','[\"sz.ss.btc.com:1800\",\"sz.ss.btc.com:443\",\"sz.ss.btc.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(11,'dcr','dcr','AntpoolDCR公共服务器','Antpool','[\"dcr.antpool.com:9002\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(12,'dcr','dcr','BTC.comDCR北京公共服务器','BTC.com','[\"cn-dcr.ss.btc.com:1800\",\"cn-dcr.ss.btc.com:443\",\"cn-dcr.ss.btc.com:25\",\"cn-dcr.ss.btc.com:1801\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(13,'bj','ltc','BTC.comLTC北京公共服务器','BTC.com','[\"cn-ltc.ss.btc.com:1800\",\"cn-ltc.ss.btc.com:443\",\"cn-ltc.ss.btc.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(14,'sz','ltc','BTC.comLTC深圳公共服务器','BTC.com','[\"sz-ltc.ss.btc.com:1800\",\"sz-ltc.ss.btc.com:443\",\"sz-ltc.ss.btc.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(15,'ltc','ltc','AntpoolLTC公共服务器','Antpool','[\"stratum-ltc.antpool.com:8888\",\"stratum-ltc.antpool.com:443\",\"stratum-ltc.antpool.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(16,'ltc','ltc','PoolinLTC公共服务器','Poolin','[\"ltc.ss.poolin.com:443\",\"ltc.ss.poolin.com:1883\",\"ltc.ss.poolin.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(17,'ltc','ltc','F2poolLTC公共服务器','F2Pool','[\"ltc.f2pool.com:8888\",\"ltc.f2pool.com:5200\",\"ltc.f2pool.com:3335\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(18,'eu','btc','BTC.comBTC欧洲公共服务器','BTC.com','[\"eu.ss.btc.com:1800\",\"eu.ss.btc.com:443\",\"eu.ss.btc.com:25\",\"47.91.65.135:443\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(19,'us','btc','BTC.comBTC美国公共服务器','BTC.com','[\"us.ss.btc.com:1800\",\"us.ss.btc.com:443\",\"us.ss.btc.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(20,'dcr','dcr','BTC.comDCR美国公共服务器','BTC.com','[\"us-dcr.ss.btc.com:1800\",\"us-dcr.ss.btc.com:443\",\"us-dcr.ss.btc.com:25\",\"us-dcr.ss.btc.com:1801\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(21,'btc','btc','蜘蛛BTC北方公共服务器','Zhizhu','[\"cn.btc.pool.zhizhu.top:1800\",\"cn.btc.pool.zhizhu.top:443\",\"cn.btc.pool.zhizhu.top:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(22,'btc','btc','蜘蛛BTC南方公共服务器','Zhizhu','[\"cs.btc.pool.zhizhu.top:1800\",\"cs.btc.pool.zhizhu.top:443\",\"cs.btc.pool.zhizhu.top:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(23,'us','btc','火币BTC美国公共服务器','HuobiPool','[\"bm.huobipool.com:1800\",\"bm.huobipool.com:443\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(24,'eu','btc','火币BTC欧洲公共服务器','HuobiPool','[\"bu.huobipool.com:1800\",\"bu.huobipool.com:443\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(25,'eu','ltc','BTC.comLTC欧洲公共服务器','BTC.com','[\"eu-ltc.ss.btc.com:1800\",\"eu-ltc.ss.btc.com:443\",\"eu-ltc.ss.btc.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(26,'us','ltc','BTC.comLTC美国公共服务器','BTC.com','[\"us-ltc.ss.btc.com:1800\",\"us-ltc.ss.btc.com:443\",\"us-ltc.ss.btc.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(27,'btc','btc','NovaBlockBTC公共服务器','NovaBlock','[\"btc.s.novablock.com:443\",\"btc.s.novablock.com:1883\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(28,'ltc','ltc','WayiLTC公共服务器','Wayi','[\"ltc.easy2mine.com:8888\",\"ltc.easy2mine.com:443\",\"ltc.easy2mine.com:13313\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(29,'btc','btc','OkexPoolBTC公共服务器','OKexPool','[\"stratum.okpool.top:3333\",\"stratum.okpool.top:443\",\"stratum.okpool.top:23\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(30,'btc','btc','火币通用BTC公共服务器','HuobiPool','[\"stratum.poolhb.com:8888\",\"bak.poolhb.com:8888\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(31,'hk','btc','火币香港BTC公共服务器','HuobiPool','[\"hk.huobipool.com:8888\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(32,'hz','ltc','火币LTC南方公共服务器','HuobiPool','[\"ls.huobipool.com:1800\",\"ls.huobipool.com:443\",\"ls.huobipool.com:3333\"]','{\"pool\":\"HuobiPool\"}','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(33,'bj','ltc','火币LTC北方公共服务器','HuobiPool','[\"ln.huobipool.com:1800\",\"ln.huobipool.com:443\",\"ln.huobipool.com:3333\"]','{\"pool\":\"HuobiPool\"}','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(34,'hk','ltc','火币香港LTC公共服务器','HuobiPool','[\"hkltc.huobipool.com:8888\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(35,'us','ltc','火币LTC美国公共服务器','HuobiPool','[\"lm.huobipool.com:1800\",\"lm.huobipool.com:443\",\"lm.huobipool.com:3333\"]','{\"pool\":\"HuobiPool\"}','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(36,'bj','dcr','火币DCR北方公共服务器','HuobiPool','[\"dn.huobipool.com:1800\",\"dn.huobipool.com:443\",\"dn.huobipool.com:3333\"]','{\"pool\":\"HuobiPool\"}','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(37,'dcr','dcr','PoolinDCR公共服务器','Poolin','[\"dcr.ss.poolin.com:443\",\"dcr.ss.poolin.com:1883\",\"dcr.ss.poolin.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(38,'sz','dcr','火币DCR南方公共服务器','HuobiPool','[\"ds.huobipool.com:1800\",\"ds.huobipool.com:443\",\"ds.huobipool.com:3333\"]','{\"pool\":\"HuobiPool\"}','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(39,'bch','bch','PoolinBCH公共服务器','Poolin','[\"bch.ss.poolin.com:443\",\"bch.ss.poolin.com:1883\",\"bch.ss.poolin.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(40,'us','dcr','火币DCR美国公共服务器','HuobiPool','[\"dm.huobipool.com:1800\",\"dm.huobipool.com:443\",\"dm.huobipool.com:3333\"]','{\"pool\":\"HuobiPool\"}','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(41,'bj','ckb','BTC.comCKB北京公共服务器','BTC.com','[\"cn-ckb.ss.btc.com:1800\",\"cn-ckb.ss.btc.com:443\",\"cn-ckb.ss.btc.com:25\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55'),
	(42,'btc','btc','小鹿BTC北方测试公共服务器','Bitdeer','[\":1800\"]','','2019-07-16 14:17:55','2019-07-16 14:17:55');

/*!40000 ALTER TABLE `btc_pool_ss_servers` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
