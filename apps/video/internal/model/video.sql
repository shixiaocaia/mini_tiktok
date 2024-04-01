DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键,视频唯一id',
    `author_id`    bigint(20) NOT NULL COMMENT '视频作者id',
    `play_url`     varchar(255) NOT NULL COMMENT '播放url',
    `cover_url`    varchar(255) NOT NULL COMMENT '封面url',
    `favorite_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '视频的点赞数量',
    `comment_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '视频的评论数量',
    `publish_time` bigint(20)     NOT NULL COMMENT '发布时间戳',
    `title`        varchar(255) DEFAULT NULL COMMENT '视频名称',
    PRIMARY KEY (`id`),
    KEY            `time` (`publish_time`) USING BTREE,
    KEY            `author` (`author_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8 COMMENT='视频表';