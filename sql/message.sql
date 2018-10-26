DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `articleId` int(11) DEFAULT NULL COMMENT '文章id',
  `uid` int(11) DEFAULT NULL COMMENT '用户uid',
  `context` text DEFAULT NULL COMMENT '内容',
  `sendTime` datetime NOT NULL COMMENT '发送时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;