DROP TABLE IF EXISTS `user`;

CREATE TABLE `user`
(
  `uid`         INT(11)     NOT NULL AUTO_INCREMENT
  COMMENT '用户id',
  `mobile`      VARCHAR(11) NOT NULL
  COMMENT '手机号=用户名',
  `pwd`         VARCHAR(20) NOT NULL
  COMMENT '密码',
  `profession`  VARCHAR(20)          DEFAULT ''
  COMMENT '职位',
  `business`    VARCHAR(20)          DEFAULT ''
  COMMENT '业务范围',
  `corp`        VARCHAR(30)          DEFAULT ''
  COMMENT '公司',
  `create_time` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
  COMMENT '创建时间',
  `update_time` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  COMMENT '更新时间',
  PRIMARY KEY (`uid`),
  UNIQUE (`mobile`)
)
  AUTO_INCREMENT = 101;

DELETE FROM `user`;

INSERT INTO `user` (`mobile`, `pwd`, `profession`, `business`, `corp`) VALUES
  ('13632385282', '123456', '风控', '广东省', '北京市畅游瑞科互联网技术有限公司广州分公司'),
  ('13687229053', '123456', '文员', '广东省', '广州壹糖网络科技有限公司');