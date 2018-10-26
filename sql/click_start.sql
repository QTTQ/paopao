DROP TABLE IF EXISTS `click_start`;
CREATE TABLE `click_start` (
`uid` int(20)  DEFAULT NULL COMMENT '用户姓名',
  `articleId` int(20) DEFAULT NULL COMMENT '文章 id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
