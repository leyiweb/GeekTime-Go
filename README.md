# GeekTime-Go
GeekTime Go训练营仓库

## 课程目录
* 微服务
* 优雅的error处理

## HomeWork
* Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
* A: 
```DDL
CREATE DATABASE geek_time;

CREATE TABLE `service_retry` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `biz_id` int NOT NULL DEFAULT '0' COMMENT '重试任务ID',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '重试任务类型',
  `data` varchar(1024) NOT NULL COMMENT '序列化数据json',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '任务状态: 0未完成 1完成 2超过次数放弃',
  `retry_num` int DEFAULT '0' COMMENT '已重试次数',
  `trace_id` varchar(100) NOT NULL COMMENT '此次任务的trace',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_biz_id` (`biz_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='服务重试表';

INSERT INTO `service_retry` (`biz_id`, `type`, `data`, `status`, `retry_num`, `trace_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (20001, 0, '{}', 100, 0, 'trace:1', '2021-06-05 23:15:48', NULL, NULL);
INSERT INTO `service_retry` (`biz_id`, `type`, `data`, `status`, `retry_num`, `trace_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (20002, 0, '{}', 100, 0, 'trace:2', '2021-06-05 23:15:48', NULL, NULL);
```