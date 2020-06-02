/*
 Navicat Premium Data Transfer

 Source Server         : docker_mysql
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 192.168.99.100:3306
 Source Schema         : swblogdb

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 02/06/2020 16:43:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_articleb
-- ----------------------------
DROP TABLE IF EXISTS `t_articleb`;
CREATE TABLE `t_articleb`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `userid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户id',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `classify` int(11) NOT NULL COMMENT '分类',
  `up` int(11) NOT NULL DEFAULT 0 COMMENT '置顶',
  `tag` int(11) NOT NULL COMMENT '标签',
  `click` int(255) NOT NULL DEFAULT 0 COMMENT '点击次数',
  `ulike` int(255) NOT NULL DEFAULT 0 COMMENT '点赞次数',
  `createtime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_articleb
-- ----------------------------
INSERT INTO `t_articleb` VALUES (1, '系统生成的第一篇文章', 'c999a2f041c84dc1b5970bb973c1da74', '系统生成的第一篇文章，这仅仅是测试，你可以从后台删除或者从数据库直接删除', 1, 0, 3, 1, 7, '2020-04-29 03:13:10');
INSERT INTO `t_articleb` VALUES (2, '系统生成的第二篇文章', 'c999a2f041c84dc1b5970bb973c1da74', '系统生成的第二篇文章，哈哈哈我随便测测', 1, 0, 4, 2, 19, '2020-05-16 07:43:47');
INSERT INTO `t_articleb` VALUES (3, '系统生成的第3篇文章', 'c999a2f041c84dc1b5970bb973c1da74', '系统生成的第3篇文章系统生成的第3篇文章系统生成的第3篇文章系统生成的第3篇文章系统生成的第3篇文章', 1, 0, 3, 1, 19, '2020-05-16 07:43:43');

-- ----------------------------
-- Table structure for t_classifyb
-- ----------------------------
DROP TABLE IF EXISTS `t_classifyb`;
CREATE TABLE `t_classifyb`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL,
  `userid` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `brief` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_classifyb
-- ----------------------------
INSERT INTO `t_classifyb` VALUES (1, 0, 'c999a2f041c84dc1b5970bb973c1da74', 'C_Sharp', 'iconfont icon-weiruan', ' ');
INSERT INTO `t_classifyb` VALUES (2, 0, 'c999a2f041c84dc1b5970bb973c1da74', 'Golang', 'iconfont icon-guge', ' ');
INSERT INTO `t_classifyb` VALUES (3, 1, 'c999a2f041c84dc1b5970bb973c1da74', 'WPF', ' ', ' ');
INSERT INTO `t_classifyb` VALUES (4, 1, 'c999a2f041c84dc1b5970bb973c1da74', '.NetCore', ' ', ' ');
INSERT INTO `t_classifyb` VALUES (5, 2, 'c999a2f041c84dc1b5970bb973c1da74', '语法', ' ', ' ');
INSERT INTO `t_classifyb` VALUES (6, 2, 'c999a2f041c84dc1b5970bb973c1da74', 'Gin', ' ', ' ');
INSERT INTO `t_classifyb` VALUES (14, 0, 'c999a2f041c84dc1b5970bb973c1da74', '杂七杂八', 'layui-icon-theme', '放一些其他的');
INSERT INTO `t_classifyb` VALUES (23, 14, 'c999a2f041c84dc1b5970bb973c1da74', '特殊的标签', '', '666');

-- ----------------------------
-- Table structure for t_lubopicb
-- ----------------------------
DROP TABLE IF EXISTS `t_lubopicb`;
CREATE TABLE `t_lubopicb`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `md5` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `webdir` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `createtime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_lubopicb
-- ----------------------------
INSERT INTO `t_lubopicb` VALUES (1, 'c999a2f041c84dc1b5970bb973c1da74', 'default1.jpg', 'e205890895285f631a21cd6661439b07', '/static/user/default1.jpg', '2020-05-27 07:48:42');
INSERT INTO `t_lubopicb` VALUES (2, 'c999a2f041c84dc1b5970bb973c1da74', 'default2.JPG', '3ad65e65c10d2b976803c402ae6d1434', '/static/user/default2.JPG', '2020-05-27 07:50:45');

-- ----------------------------
-- Table structure for t_swsystemb
-- ----------------------------
DROP TABLE IF EXISTS `t_swsystemb`;
CREATE TABLE `t_swsystemb`  (
  `id` int(255) NOT NULL,
  `allsee` int(255) NOT NULL COMMENT '总访问量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_swsystemb
-- ----------------------------
INSERT INTO `t_swsystemb` VALUES (1, 56);

-- ----------------------------
-- Table structure for t_userb
-- ----------------------------
DROP TABLE IF EXISTS `t_userb`;
CREATE TABLE `t_userb`  (
  `id` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `loginname` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `brief` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `txurl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `txmd5` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `luboid` int(11) NULL DEFAULT NULL,
  `createtime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updatatime` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_userb
-- ----------------------------
INSERT INTO `t_userb` VALUES ('c999a2f041c84dc1b5970bb973c1da74', '李艳朋', 'liyanpeng', 'c56d0e9a7ccec67b4ea131655038d604', '不想平凡，奈何太懒 T_T', 't115liyanpeng@126.com', '/static/user/user_tx.PNG', 'aecf6ca601b112754f790252f5b69149', NULL, '2020-04-29 10:39:01', '2020-05-29 09:51:53');

SET FOREIGN_KEY_CHECKS = 1;
