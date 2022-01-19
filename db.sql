drop database if exists `second`;

create database `second`;

use `second`;

-- 用户
create table `user`(
   `user_id` int not null AUTO_INCREMENT comment "用户id" ,
   `account` varchar(30) null comment "学号",
   `nickname` varchar(20) null comment "昵称",
   `image` MEDIUMTEXT null comment "头像",
   `qq_account` varchar(30) null comment "QQ号",
   
-- 添加约束
primary key (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

--商品
create table `goods`(
   `goods_id` int not null AUTO_INCREMENT comment "商品id" ,
   `buyer_id` int null comment "买家id" ,
   `seller_id` int null comment "卖家id" ,
   `time` varchar(30) null comment "发布时间",
   `description` varchar(100) null comment "描述",
   `image` MEDIUMTEXT null comment "图片",
   `tag_ids` varchar(30) null comment "标签串",
   
-- 添加约束
primary key (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

--标签
create table `tag`(
   `tag_id` int not null AUTO_INCREMENT comment "标签id" ,
   `owner_id` int null comment "用户id" ,
   `content` varchar(100) null comment "内容描述",
   
-- 添加约束
primary key (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

--收藏
create table `collection`(
   `collection_id` int not null AUTO_INCREMENT comment "收藏id" ,
   `owner_id` int null comment "用户id" ,
   `goods_id` int null comment "商品id" ,
   
-- 添加约束
primary key (`collection_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

--意见反馈
create table `feedback`(
   `feedback_id` int not null AUTO_INCREMENT comment "意见反馈项id" ,
   `owner_id` int null comment "用户id" ,
   `content` varchar(100) null comment "反馈内容",

-- 添加约束
primary key (`feedback_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

