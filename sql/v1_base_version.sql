/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : lifeboat

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 27/05/2021 11:03:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for role_info
-- ----------------------------
DROP TABLE IF EXISTS `role_info`;
CREATE TABLE `role_info` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(100) NOT NULL COMMENT '游戏角色名称',
  `hp` tinyint NOT NULL COMMENT '角色体力值',
  `suv_score` tinyint NOT NULL COMMENT '角色生存分数',
  `seat` tinyint NOT NULL COMMENT '角色船上位置',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色技能介绍',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色信息表';

-- ----------------------------
-- Records of role_info
-- ----------------------------
BEGIN;
INSERT INTO `role_info` VALUES (1, 'Rolen', 4, 8, 1, 'Double your Jewelry Score');
INSERT INTO `role_info` VALUES (2, 'Steve', 5, 7, 2, 'Double your Picture Score');
INSERT INTO `role_info` VALUES (3, 'Captain', 7, 5, 3, 'Double your Money Score');
INSERT INTO `role_info` VALUES (4, 'Mate', 8, 4, 4, 'Just strength');
INSERT INTO `role_info` VALUES (5, 'France', 6, 6, 5, 'Excellent swimming skills, falling into the sea will not cause damage');
INSERT INTO `role_info` VALUES (6, 'Child', 3, 9, 6, 'When you steal something from someone, they can\'t refuse');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
