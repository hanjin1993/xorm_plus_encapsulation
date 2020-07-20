/*
SQLyog Ultimate v11.11 (64 bit)
MySQL - 5.5.53 : Database - go_order
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`go_order` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `go_order`;

/*Table structure for table `go_goods` */

DROP TABLE IF EXISTS `go_goods`;

CREATE TABLE `go_goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `good_name` varchar(50) NOT NULL DEFAULT '',
  `store` int(11) NOT NULL DEFAULT '0',
  `good_sn` char(50) NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `sale_num` int(11) unsigned NOT NULL DEFAULT '0',
  `is_on_sale` tinyint(1) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `go_goods` */

insert  into `go_goods`(`id`,`good_name`,`store`,`good_sn`,`create_time`,`sale_num`,`is_on_sale`) values (1,'测试商品1',10,'N000001',0,0,1),(2,'测试商品2',20,'N000002',0,0,1),(3,'测试商品3',30,'N000003',0,0,1),(4,'测试商品4',40,'N000004',0,0,0);

/*Table structure for table `go_order` */

DROP TABLE IF EXISTS `go_order`;

CREATE TABLE `go_order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `order_sn` char(30) NOT NULL DEFAULT '',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `address` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `go_order` */

insert  into `go_order`(`id`,`order_sn`,`user_id`,`create_time`,`status`,`address`) values (1,'OR1234567',1,0,0,''),(2,'OR23456',2,0,0,''),(3,'OR34567',1,0,0,'');

/*Table structure for table `go_order_goods` */

DROP TABLE IF EXISTS `go_order_goods`;

CREATE TABLE `go_order_goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(11) unsigned NOT NULL DEFAULT '0',
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `num` int(11) unsigned NOT NULL DEFAULT '0',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `go_order_goods` */

insert  into `go_order_goods`(`id`,`order_id`,`goods_id`,`num`,`create_time`) values (1,1,1,10,0),(2,2,2,10,0),(3,1,3,1,0),(4,3,4,5,0);

/*Table structure for table `go_users` */

DROP TABLE IF EXISTS `go_users`;

CREATE TABLE `go_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` char(20) NOT NULL DEFAULT '',
  `password` varchar(50) NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `go_users` */

insert  into `go_users`(`id`,`user_name`,`password`,`create_time`,`sex`) values (1,'test','123456',0,1),(2,'hello world','123456',1594968771,0),(3,'hello world','123456',1594969020,0);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
