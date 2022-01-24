drop database if exists `second`;

create database `second`;

use `second`;



-- 用户
create table `users`(
   `id` int not null AUTO_INCREMENT comment "用户id" ,
   `account` varchar(30) null comment "学号",
   `nickname` varchar(20) null comment "昵称",
   `password` varchar(20) null comment "密码",
   `image` varchar(250) null comment "头像",
   `qq_account` varchar(30) null comment "QQ号",
   
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;




-- 商品
create table `goods`(
   `id` int not null AUTO_INCREMENT comment "商品id" ,
   `buyer_id` int null comment "买家id" ,
   `seller_id` int null comment "卖家id" ,
   `time` varchar(30) null comment "发布时间",
   `description` varchar(5000) null comment "描述",
   `images_videos` varchar(250) null comment "图片和视频",
   `tag_ids` varchar(30) null comment "标签串",
   `if_sell` boolean null comment "是否卖出",
   `if_del`  boolean null comment "是否删除",
   
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;




-- 标签
create table `tags`(
   `id` int not null AUTO_INCREMENT comment "标签id" ,
   `owner_id` int null comment "用户id" ,
   `content` varchar(100) null comment "内容描述",
   
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;




-- 收藏
create table `collections`(
   `id` int not null AUTO_INCREMENT comment "收藏id" ,
   `owner_id` int null comment "用户id" ,
   `goods_id` int null comment "商品id" ,
   
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;




-- 意见反馈
create table `feedbacks`(
   `id` int not null AUTO_INCREMENT comment "意见反馈项id" ,
   `owner_id` int null comment "用户id" ,
   `content` varchar(5000) null comment "反馈内容",

-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

