CREATE DATABASE IF NOT EXISTS tiny_tiktok_db;

-- 创建 user 数据表
DROP TABLE IF EXISTS `tiny_tiktok_db`.`user`;
CREATE TABLE IF NOT EXISTS `tiny_tiktok_db`.`user` (
    `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` int(11) NOT NULL COMMENT '用户ID',
    `name` varchar(64) NOT NULL COMMENT '用户名称',
    `password` varchar(64) NOT NULL COMMENT '用户密码',
    `follow_count` int(11) NOT NULL DEFAULT '0' COMMENT '关注总数',
    `follower_count` int(11) NOT NULL DEFAULT '0' COMMENT '粉丝总数',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COMMENT='用户列表';

-- 创建 video 数据表
DROP TABLE IF EXISTS `tiny_tiktok_db`.`video`;
CREATE TABLE IF NOT EXISTS `tiny_tiktok_db`.`video` (
    `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title` varchar(255) NOT NULL DEFAULT 'default_title' COMMENT '视频标题',
    `video_id` int(11) NOT NULL DEFAULT '0' COMMENT '视频ID',
    `author_id` int(11) NOT NULL DEFAULT '0' COMMENT '视频作者ID',
    `play_url` varchar(255) NOT NULL DEFAULT '-' COMMENT '视频播放地址',
    `cover_url` varchar(255) NOT NULL DEFAULT '-' COMMENT '视频封面地址',
    `created_time` datetime NOT NULL DEFAULT '2022-12-25 00:00:00' COMMENT '投稿时间',
    `favorite_count` int(11) NOT NULL DEFAULT '0' COMMENT '视频点赞总数',
    `comment_count` int(11) NOT NULL DEFAULT '0' COMMENT '视频评论总数',
    PRIMARY KEY (`id`),
    KEY `idx_video_id` (`video_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COMMENT='视频列表';