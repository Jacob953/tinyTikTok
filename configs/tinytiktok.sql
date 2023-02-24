CREATE DATABASE IF NOT EXISTS `tiny_tiktok_db`;

USE `tiny_tiktok_db`;

-- 创建 user 数据表
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `name` varchar(64) NOT NULL COMMENT '用户名称',
    `password` char(60) NOT NULL COMMENT '用户密码',
    `follow_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '关注总数',
    `follower_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '粉丝总数',
    `is_follow` TINYINT(1) NOT NULL COMMENT '关注状态',
    `avatar` varchar(255) NOT NULL COMMENT '用户头像',
    `background_image` varchar(255) NOT NULL COMMENT '用户个人页顶部大图',
    `signature` varchar(255) NOT NULL COMMENT '个人简介',
    `total_favorited` bigint(20) NOT NULL DEFAULT '0' COMMENT '获赞数量',
    `work_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '作品数量',
    `favorite_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '点赞数量',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COMMENT='用户列表';

-- 创建 video 数据表
DROP TABLE IF EXISTS `video`;
CREATE TABLE IF NOT EXISTS `video` (
    `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `video_id` BIGINT(20) NOT NULL COMMENT '视频唯一标识',
    `author_id` BIGINT(20) NOT NULL COMMENT '作者id',
    `play_url` VARCHAR(255) NOT NULL COMMENT '视频播放地址',
    `cover_url` VARCHAR(255) NOT NULL COMMENT '视频封面地址',
    `favorite_count` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '视频的点赞总数',
    `comment_count` BIGINT(20) NOT NULL DEFAULT '0' COMMENT '视频的评论总数',
    `is_favorite` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '是否已点赞，0-未点赞，1-已点赞',
    `title` VARCHAR(255) NOT NULL COMMENT '视频标题',
    `create_at` BIGINT(20) NOT NULL COMMENT '视频创建时间',
    PRIMARY KEY (`id`),
    INDEX `idx_video_id` (`video_id`),
    CONSTRAINT `fk_policy_user` FOREIGN KEY (author_id) REFERENCES user(user_id)
    ) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COMMENT='视频列表';

-- 创建 comment 数据表
DROP TABLE IF EXISTS `comment`;
CREATE TABLE IF NOT EXISTS `comment` (
   `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
   `comment_id` int(11) NOT NULL DEFAULT '0' COMMENT '视频评论ID',
   `user` varchar(255) NOT NULL DEFAULT '0' COMMENT '评论用户信息',
   `content` varchar(255) NOT NULL DEFAULT '-' COMMENT '评论内容',
   `create_date` datetime NOT NULL DEFAULT '2022-12-25 00:00:00' COMMENT '评论发布日期',
   `video_id` int(11) NOT NULL DEFAULT '0' COMMENT '视频ID',
   PRIMARY KEY (`id`),
   KEY `idx_comment_id` (`comment_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COMMENT='评论列表';

-- 创建 favorite 数据表
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE favorite (
    `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
    `video_id` INT UNSIGNED NOT NULL COMMENT '视频ID',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '视频发布时间',
    PRIMARY KEY (`user_id`, `video_id`),
    UNIQUE KEY (`user_id`, `video_id`)
    )ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COMMENT='点赞列表';
