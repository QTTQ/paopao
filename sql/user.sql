DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `uid` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增用户uid',
  `user_name` char(20) NOT NULL COMMENT '用户姓名',
  `pass_word` varchar(20) NOT NULL COMMENT '密码',
  `reg_time` datetime NOT NULL COMMENT '注册时间',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
