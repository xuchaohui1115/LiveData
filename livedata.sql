-- 创建主播基础信息表
DROP TABLE if exists `hosts`;
CREATE TABLE `hosts` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(50)  CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `password` VARCHAR(50)  CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码（建议加密存储）',
    `nickname` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
    `phone` VARCHAR(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
    `email` VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电子邮件',
    `profile_picture` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像URL',
    `bio` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '个人简介',
    `followers_count` int NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    `created_at` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间（Unix 时间戳）',
    `updated_at` BIGINT NOT NULL DEFAULT '0' COMMENT '更新时间（Unix 时间戳）',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='主播基础信息表';

-- 创建主播违规记录表
DROP TABLE IF EXISTS `violations`;
CREATE TABLE `violations` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `host_nickname` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '主播昵称',
    `host_id` BIGINT NOT NULL DEFAULT 0 COMMENT '主播ID',  -- 使用 BIGINT 以防 ID 超出 INT 范围
    `security_code` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '安全码',
    `account_status` ENUM('active', 'suspended', 'banned') NOT NULL DEFAULT 'active' COMMENT '账号状态',
    `violation_number` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '违规编号',
    `violation_time` BIGINT NOT NULL DEFAULT 0 COMMENT '违规时间（Unix 时间戳）',
    `violation_reason` TEXT NOT NULL COMMENT '违规原因',
    `violation_impact` TEXT NOT NULL COMMENT '违规影响',
    `violation_live_room_id` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '违规对象是直播间的ID',
    `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix 时间戳）',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='违规记录表';

-- 创建直播间基础信息表
DROP TABLE IF EXISTS `live_details`;
CREATE TABLE `live_details` (
  `live_room_id` BIGINT NOT NULL PRIMARY KEY COMMENT '直播间ID，主键',
  `live_title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '直播标题',
  `live_start_time` BIGINT NOT NULL DEFAULT 0 COMMENT '开播时间（Unix 时间戳）',
  `host_id` BIGINT NOT NULL DEFAULT 0 COMMENT '主播ID，用于关联主播',
  `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix 时间戳)',
  INDEX idx_host_id (`host_id`) COMMENT '主播ID索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='直播间详情表';

-- 创建直播管理数据详情表
DROP TABLE IF EXISTS `live_data`;
CREATE TABLE `live_data` (
   `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
   `host_nickname` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '主播昵称',
   `tb_account` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '淘宝账号',
   `mcn_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'MCN名称',
   `phone` VARCHAR(15) DEFAULT '' COMMENT '手机号码',
   `contact_person` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '对接人',
   `users_count` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '在线人数',
   `host_id` BIGINT NOT NULL DEFAULT 0 COMMENT '主播ID',
   `uid` BIGINT NOT NULL DEFAULT 0 COMMENT 'uid',
   `live_status` ENUM('ongoing', 'completed', 'paused') NOT NULL COMMENT '直播状态',
   `host_status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active' COMMENT '主播状态',
   `login_status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active' COMMENT '登录状态',
   `remarks` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
   `live_duration` INT NOT NULL DEFAULT 0 COMMENT '直播时长（以分钟为单位）',
   `payment_amount` DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '支付金额',
   `view_count` INT NOT NULL DEFAULT 0 COMMENT '观看次数',
   `click_through_rate` DECIMAL(5, 2) NOT NULL DEFAULT 0.00 COMMENT '点击率',
   `conversion_rate` DECIMAL(5, 2) NOT NULL DEFAULT 0.00 COMMENT '成交率',
   `transaction_count` INT NOT NULL DEFAULT 0 COMMENT '成交人数',
   `transaction_items` INT NOT NULL DEFAULT 0 COMMENT '成交件数',
   `yesterday_sales` DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '昨日成交',
   `all_sales` DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '总成交',
   `month_sales` DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '月销售',
   `pause_count` INT NOT NULL DEFAULT 0 COMMENT '暂停次数',
   `pause_record_id` BIGINT NOT NULL DEFAULT 0 COMMENT '暂停记录表ID，用于关联',
   `live_room_id` BIGINT NOT NULL DEFAULT 0 COMMENT '直播间ID，主键',
   `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix 时间戳）',
   `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '修改时间（Unix 时间戳）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='直播数据表';

-- 创建直播暂停记录表
DROP TABLE IF EXISTS `pause_records`;
CREATE TABLE `pause_records` (
   `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
   `live_data_id` BIGINT NOT NULL DEFAULT 0 COMMENT '直播数据表的ID，用于关联',
   `host_id` BIGINT NOT NULL DEFAULT 0 COMMENT '主播ID，用于关联',
   `pause_time` BIGINT NOT NULL DEFAULT 0 COMMENT '暂停时间（Unix 时间戳）',
   `resume_time` BIGINT NOT NULL DEFAULT 0 COMMENT '恢复时间（Unix 时间戳）',
   `duration` INT NOT NULL DEFAULT 0 COMMENT '暂停时长，以分钟为单位',
   `reason` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '暂停理由',
   `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix 时间戳）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='暂停记录表';

-- 创建超级管理员表
DROP TABLE IF EXISTS `super_admins`;
CREATE TABLE `super_admins` (
   `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
   `username` VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名，唯一',
   `password` VARCHAR(255) NOT NULL COMMENT '密码（建议加密存储）',
   `nickname` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '昵称',
   `phone` VARCHAR(15) DEFAULT '' COMMENT '手机号码',
   `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active' COMMENT '状态',
   `created_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间（Unix 时间戳）',
   `updated_at` BIGINT NOT NULL DEFAULT 0 COMMENT '修改时间（Unix 时间戳）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='超级管理员表';

-- 创建热浪引擎热门商品表
-- (请在此处添加热浪引擎热门商品表的创建语句)

-- 创建商品数据详情表
-- (请在此处添加商品数据详情表的创建语句)

-- 创建分销商表
-- (请在此处添加分销商表的创建语句)
