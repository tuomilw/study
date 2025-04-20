/**
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
DROP TABLE IF EXISTS `students`;
CREATE TABLE `students` (
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name`  varchar(255) NOT NULL DEFAULT '' COMMENT '学生姓名',
    `age`   int(11) NOT NULL DEFAULT 0 COMMENT '学生年龄',
    `grade` varchar(100)  NOT NULL DEFAULT '' COMMENT '学生年级',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='学生信息';

/**
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
 */
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `balance` decimal(18,2) NOT NULL DEFAULT 0.0 COMMENT '账户余额',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='账户表';

DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `from_account_id` decimal(18,2) NOT NULL DEFAULT 0.0 COMMENT '转出账户ID',
    `to_account_id` decimal(18,2) NOT NULL DEFAULT 0.0 COMMENT 'to_account_id',
    `amount` decimal(18,2) NOT NULL DEFAULT 0.0 COMMENT '转账金额',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='交易记录表';

/**
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
 */
DROP TABLE IF EXISTS `employees`;
CREATE TABLE `employees` (
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '员工姓名',
    `department` varchar(255) NOT NULL DEFAULT '' COMMENT '部门名称',
    `salary` decimal(18,2) NOT NULL DEFAULT 0.0 COMMENT '工资',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='员工信息';

/**
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
 */
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '书本名称',
    `author` varchar(255) NOT NULL DEFAULT '' COMMENT '作者',
    `price` decimal(18,2) NOT NULL DEFAULT 0.0 COMMENT '价格',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='书本信息';

/**
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
 */
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `post_counts` int(11) NOT NULL DEFAULT 0 COMMENT '文章数量',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息';

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` bigint(11) NOT NULL DEFAULT 0 COMMENT '用户ID',
    `comment_counts` int(11) NOT NULL DEFAULT 0 COMMENT '评论数',
    `state` varchar(1) NOT NULL DEFAULT '0' COMMENT '评论状态',
    `content` varchar(500) NOT NULL DEFAULT '' COMMENT '文章内容',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章信息';

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `post_id` bigint(11) NOT NULL DEFAULT 0 COMMENT '文章ID',
    `content` varchar(500) NOT NULL DEFAULT '' COMMENT '评论信息',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论信息';