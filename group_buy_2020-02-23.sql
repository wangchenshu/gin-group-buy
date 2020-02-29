# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.28)
# Database: group_buy
# Generation Time: 2020-02-23 10:46:57 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table cart
# ------------------------------------------------------------

DROP TABLE IF EXISTS `cart`;

CREATE TABLE `cart` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主鍵ID',
  `product_name` varchar(50) NOT NULL COMMENT '商品名稱',
  `line_user_id` varchar(50) NOT NULL COMMENT 'Line User ID',
  `username` varchar(50) NOT NULL COMMENT 'Line User Name',
  `qty` int(10) unsigned NOT NULL COMMENT '商品數量',
  `price` int(10) unsigned NOT NULL COMMENT '商品價格',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改時間',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table chatfuel_cart
# ------------------------------------------------------------

DROP TABLE IF EXISTS `chatfuel_cart`;

CREATE TABLE `chatfuel_cart` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主鍵ID',
  `product_name` varchar(50) NOT NULL COMMENT '商品名稱',
  `messenger_user_id` varchar(100) NOT NULL COMMENT 'Messenger User ID',
  `username` varchar(50) NOT NULL COMMENT 'Messenger User Name',
  `qty` int(10) unsigned NOT NULL COMMENT '商品數量',
  `price` int(10) unsigned NOT NULL COMMENT '商品價格',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改時間',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table fb_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `fb_user`;

CREATE TABLE `fb_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `messenger_user_id` varchar(100) NOT NULL,
  `timezone` int(11) DEFAULT NULL,
  `locale` varchar(10) DEFAULT NULL,
  `gender` varchar(10) DEFAULT NULL,
  `profile_pic_url` varchar(300) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `messenger_user_id` (`messenger_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table line_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `line_user`;

CREATE TABLE `line_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table order
# ------------------------------------------------------------

DROP TABLE IF EXISTS `order`;

CREATE TABLE `order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned NOT NULL,
  `line_user_id` varchar(50) NOT NULL DEFAULT '',
  `qty` int(10) unsigned NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table product
# ------------------------------------------------------------

DROP TABLE IF EXISTS `product`;

CREATE TABLE `product` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `price` int(10) unsigned DEFAULT NULL,
  `link` varchar(255) NOT NULL DEFAULT '',
  `pic_url` varchar(255) DEFAULT NULL,
  `active` tinyint(4) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;

INSERT INTO `product` (`id`, `name`, `price`, `link`, `pic_url`, `active`, `created_at`, `updated_at`)
VALUES
	(1,'艾多美 香烤海苔(小片裝) 1箱 (1箱4盒)',998,'http://www.atomy.com/tw/Home/Product/ProductView?GdsCode=W00904','https://www.atomy.com/tw/shopping/p_img/100/00904_2.jpg',1,'2020-02-22 13:54:58','2020-02-23 07:15:03'),
	(2,'艾多美 幸福堅果',1350,'http://www.atomy.com/tw/Home/Product/ProductView?GdsCode=W009824','https://www.atomy.com/tw/shopping/p_img/100/00982_2.jpg',1,'2020-02-22 13:57:49','2020-02-23 07:15:01');

/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主鍵ID',
  `name` varchar(50) NOT NULL COMMENT '用户名稱',
  `password` varchar(80) NOT NULL COMMENT '用戶密碼',
  `last_token` text COMMENT '登錄時的token',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '用户狀態 -1代表已删除 0代表正常 1代表凍結',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改時間',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
