DROP DATABASE IF EXISTS `test`;
CREATE DATABASE `test`;
USE `test`;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`   int(11)     NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name` varchar(40) NOT NULL COMMENT '名字',
    `age`  int         NOT NULL DEFAULT 0 COMMENT '年龄',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT 'users';