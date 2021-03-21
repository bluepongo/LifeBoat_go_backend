CREATE TABLE `role_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(100) NOT NULL COMMENT '游戏角色名称',
  `hp` tinyint(4) NOT NULL COMMENT '角色体力值',
  `suv_score` tinyint(4) NOT NULL COMMENT '角色生存分数',
  `seat` tinyint(4) NOT NULL COMMENT '角色船上位置',
  `describe` varchar(255) NOT NULL COMMENT '角色技能介绍',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色信息表';