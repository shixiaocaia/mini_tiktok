DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`               bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
    `user_name`        varchar(255) NOT NULL COMMENT '用户名',
    `password`         varchar(255) NOT NULL COMMENT '用户密码',
    `follow_count`     bigint(20) NOT NULL DEFAULT 0 COMMENT '该用户关注其他用户个数',
    `follower_count`   bigint(20) NOT NULL DEFAULT 0 COMMENT '该用户粉丝个数',
    `total_favorited`  bigint(20) NOT NULL DEFAULT 0 COMMENT '该用户被喜欢的视频数量',
    `favorite_count`   bigint(20) NOT NULL DEFAULT 0 COMMENT '该用户喜欢的视频数量',
    `signature` varchar(1024) COMMENT '签名',
    `avatar`           varchar(1024) COMMENT '用户头像',
    `background_image` varchar(1024) COMMENT '主页背景',
    `work_count`      bigint(20) NOT NULL DEFAULT 0 COMMENT '作品数量',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_name_unique` (`user_name`) -- 这里添加唯一键约束
) ENGINE=InnoDB AUTO_INCREMENT=20044 DEFAULT CHARSET=utf8 COMMENT='用户表';