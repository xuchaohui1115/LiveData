-- 创建主播基础信息表
CREATE TABLE `hosts` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,  -- 主键ID
    `username` VARCHAR(50) DEFAULT '',     -- 用户名
    `password` VARCHAR(255) DEFAULT '',    -- 密码（建议加密存储）
    `nickname` VARCHAR(50) DEFAULT '',     -- 昵称
    `phone` VARCHAR(15) DEFAULT '',        -- 手机号码
    `email` VARCHAR(100) DEFAULT '',       -- 电子邮件
    `profile_picture` VARCHAR(255) DEFAULT '', -- 头像URL
    `bio` TEXT DEFAULT '',                  -- 个人简介
    `followers_count` INT DEFAULT 0,       -- 粉丝数量
    `created_at` BIGINT DEFAULT 0,         -- 创建时间（Unix 时间戳）
    `updated_at` BIGINT DEFAULT 0           -- 更新时间（Unix 时间戳）
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建主播违规记录表
CREATE TABLE `violations` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,            -- 主键ID
    `host_nickname` VARCHAR(50) DEFAULT '',         -- 主播昵称
    `host_id` INT DEFAULT 0,                         -- 主播ID
    `security_code` VARCHAR(50) DEFAULT '',         -- 安全码
    `account_status` ENUM('active', 'suspended', 'banned') DEFAULT 'active', -- 账号状态
    `violation_number` VARCHAR(50) DEFAULT '',      -- 违规编号
    `violation_time` BIGINT DEFAULT 0,               -- 违规时间（Unix 时间戳）
    `violation_reason` TEXT DEFAULT '',              -- 违规原因
    `violation_impact` TEXT DEFAULT '',              -- 违规影响
    `violation_live_room_id` VARCHAR(100) DEFAULT '', -- 违规对象是直播间的ID
    `created_at` BIGINT DEFAULT 0                    -- 创建时间（Unix 时间戳）
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建直播间基础信息表
CREATE TABLE `live_details` (
    `live_room_id` INT PRIMARY KEY,                 -- 直播间ID，主键
    `live_title` VARCHAR(255) DEFAULT '',           -- 直播标题
    `live_start_time` BIGINT DEFAULT 0,             -- 开播时间（Unix 时间戳）
    `host_id` INT DEFAULT 0,                         -- 主播ID，用于关联主播
    `created_at` BIGINT DEFAULT 0                    -- 创建时间（Unix 时间戳）
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建直播管理数据详情表
CREATE TABLE `live_data` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,             -- 主键ID
    `host_nickname` VARCHAR(50) DEFAULT '',         -- 主播昵称
    `host_id` INT DEFAULT 0,                         -- 主播ID
    `live_status` ENUM('ongoing', 'completed', 'paused') NOT NULL, -- 直播状态
    `host_status` ENUM('active', 'inactive') DEFAULT 'active',     -- 主播状态
    `remarks` TEXT DEFAULT '',                        -- 备注
    `live_duration` INT DEFAULT 0,                   -- 直播时长（以分钟为单位）
    `payment_amount` DECIMAL(10, 2) DEFAULT 0.00,    -- 支付金额
    `view_count` INT DEFAULT 0,                      -- 观看次数
    `click_through_rate` DECIMAL(5, 2) DEFAULT 0.00, -- 点击率
    `conversion_rate` DECIMAL(5, 2) DEFAULT 0.00,    -- 成交率
    `transaction_count` INT DEFAULT 0,               -- 成交人数
    `transaction_items` INT DEFAULT 0,               -- 成交件数
    `yesterday_sales` DECIMAL(10, 2) DEFAULT 0.00,   -- 昨日成交
    `pause_count` INT DEFAULT 0,                     -- 暂停次数
    `pause_record_id` INT DEFAULT 0,                 -- 暂停记录表ID，用于关联
    `created_at` BIGINT DEFAULT 0,                   -- 创建时间（Unix 时间戳）
    `updated_at` BIGINT DEFAULT 0                     -- 修改时间（Unix 时间戳）
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建直播暂停记录表
CREATE TABLE `pause_records` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,             -- 主键ID
    `live_data_id` INT DEFAULT 0,                    -- 直播数据表的ID，用于关联
    `host_id` INT DEFAULT 0,                         -- 主播ID，用于关联
    `pause_time` BIGINT DEFAULT 0,                   -- 暂停时间（Unix 时间戳）
    `resume_time` BIGINT DEFAULT 0,                  -- 恢复时间（Unix 时间戳）
    `duration` INT DEFAULT 0,                        -- 暂停时长，以分钟为单位
    `reason` TEXT DEFAULT '',                         -- 暂停理由
    `created_at` BIGINT DEFAULT 0                    -- 创建时间（Unix 时间戳）
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建超级管理员表
CREATE TABLE `super_admins` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,             -- 主键ID
    `username` VARCHAR(50) DEFAULT '',                -- 用户名，唯一
    `password` VARCHAR(255) DEFAULT '',               -- 密码（建议加密存储）
    `nickname` VARCHAR(50) DEFAULT '',                -- 昵称
    `phone` VARCHAR(15) DEFAULT '',                   -- 手机号码
    `status` ENUM('active', 'inactive') DEFAULT 'active', -- 状态
    `created_at` BIGINT DEFAULT 0,                   -- 创建时间（Unix 时间戳）
    `updated_at` BIGINT DEFAULT 0                     -- 修改时间（Unix 时间戳）
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建热浪引擎热门商品表
-- (请在此处添加热浪引擎热门商品表的创建语句)

-- 创建商品数据详情表
-- (请在此处添加商品数据详情表的创建语句)

-- 创建分销商表
-- (请在此处添加分销商表的创建语句)