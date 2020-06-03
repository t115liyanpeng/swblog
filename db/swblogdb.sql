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

 Date: 03/06/2020 18:38:53
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
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_articleb
-- ----------------------------
INSERT INTO `t_articleb` VALUES (1, '系统生成的第一篇文章', 'c999a2f041c84dc1b5970bb973c1da74', '系统生成的第一篇文章，这仅仅是测试，你可以从后台删除或者从数据库直接删除', 1, 0, 3, 1, 7, '2020-04-29 03:13:10');
INSERT INTO `t_articleb` VALUES (2, '系统生成的第二篇文章', 'c999a2f041c84dc1b5970bb973c1da74', '系统生成的第二篇文章，哈哈哈我随便测测', 1, 0, 4, 2, 19, '2020-05-16 07:43:47');
INSERT INTO `t_articleb` VALUES (3, '系统生成的第3篇文章', 'c999a2f041c84dc1b5970bb973c1da74', '系统生成的第3篇文章系统生成的第3篇文章系统生成的第3篇文章系统生成的第3篇文章系统生成的第3篇文章', 1, 0, 3, 1, 19, '2020-05-16 07:43:43');
INSERT INTO `t_articleb` VALUES (11, '测试markdown转换', 'c999a2f041c84dc1b5970bb973c1da74', '还是老生常谈，之前自己搭建golang开发环境的时候 使用vscode 但是 因为众所周知的原因（被墙的原因），插件总是安装失败，之前自己安装的时候总是 稀里糊涂的就弄好了，\n今天要彻底的弄明白一下。之前的博文都是转载的，后来 自己看的时候都看不明白了，索性就都删了自己写一篇。\n开头安装vscode 就不重新写了，重点是安装 安装失败的golang插件，首先我们 按照提示Install all 结果得到的是（如下）：\n\nInstalling github.com/mdempsky/gocode FAILED\n\nInstalling github.com/uudashr/gopkgs/v2/cmd/gopkgs FAILED\n\nInstalling github.com/ramya-rao-a/go-outline FAILED\n\nInstalling github.com/acroca/go-symbols FAILED\n\nInstalling golang.org/x/tools/cmd/guru FAILED\n\nInstalling golang.org/x/tools/cmd/gorename FAILED\n\nInstalling github.com/cweill/gotests/... FAILED\n\nInstalling github.com/fatih/gomodifytags FAILED\n\nInstalling github.com/josharian/impl FAILED\n\nInstalling github.com/davidrjenni/reftools/cmd/fillstruct FAILED\n\nInstalling github.com/haya14busa/goplay/cmd/goplay FAILED\n\nInstalling github.com/godoctor/godoctor FAILED\n\nInstalling github.com/go-delve/delve/cmd/dlv FAILED\n\nInstalling github.com/stamblerre/gocode FAILED\n\nInstalling github.com/rogpeppe/godef FAILED\n\nInstalling github.com/sqs/goreturns FAILED\n\nInstalling golang.org/x/lint/golint FAILED\n\n全都是失败，可能是因为公司网不太好，我自己在家安装的时候 还是有一部分是可以安装成功的，接下来就是重点了（这里我们以“go-outline”这个插件为例子）我电脑是windows10 我们这里就以win10\n为例，linux mac 都差不多。\n\n首先要安装好git 原因不多讲\n这里我们直接 go install github.com/ramya-rao-a/go-outline \n结果会发现安装失败，不要气馁，我们看报错提示\n![](https://i.loli.net/2020/04/07/eqmcL4j2SavuRoh.png)\n这里我们看到了提示“不能找到 github.com/ramya-rao-a/go-outline ”这个包，那么我们就去我们的gopath下看一下\n\n果然在gopath的src里不仅没有这个包，而且连github.com这个目录都没有，二话不说 我们直接建立好github.com的目录同时 我们需要在刚刚建立好的github.com的目录下建立名为ramya-rao-a的目录，\n因为一会儿我们要用git下载的时候git不会建立这个目录，如果不建立这个目录应该也是没问题的，但是go install 的时候需要注意改变一下路径，为了方便我就直接将ramya-rao-a这个目录建立好，接下\n来 我们使用git 下载代码\n在刚刚建立的的ramya-rao-a目录下使用git工具输入\ngit clone https://github.com/ramya-rao-a/go-outline\n\n会如图\n![](https://i.loli.net/2020/04/07/M65daZ1UqIuWjBJ.jpg)\n此时，这个插件的代码我们已经下载下来了，\n我们回到src路径下 使用go install github.com/ramya-rao-a/go-outline  命令来安装\n \n发现没有提示 失败 这个时候我们再次检查 bin 目录下已经存在go-outline.exe 说明这个插件已经安装好了以此类推 其他的插件我们也可以这样安装好。\n ![](https://i.loli.net/2020/04/07/q9Ee3YR8xfWpuIs.png)\n最后 需要注意的是 类似“ golang.org/x/tools/cmd/guru”这样的包 （golang.org）开头的  由于某些原因（被墙）一般在 GitHub 上都有官方的镜像仓库对应 ，这时我们需要 手动去github.com上去找相应的 镜像\n使用git 将响应的代码 下载下来使用 go install 的方法安装。\n\n希望我写的能对你有用\n  \n\n', 2, 0, 5, 4, 1, '2020-06-02 08:47:03');

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
INSERT INTO `t_swsystemb` VALUES (1, 64);

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
