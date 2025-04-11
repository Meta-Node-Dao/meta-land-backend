/*
Navicat MySQL Data Transfer

Source Server         : test
Source Server Version : 50737
Source Host           : 10.168.171.107:3306
Source Database       : comunion

Target Server Type    : MYSQL
Target Server Version : 50737
File Encoding         : 65001

Date: 2022-05-25 15:46:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `comer`
-- ----------------------------
DROP TABLE IF EXISTS `comer`;
CREATE TABLE `comer` (
                         `id` bigint(20) NOT NULL,
                         `address` char(42) DEFAULT NULL COMMENT 'comer could save some useful info on block chain with this address',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `comer_address_uindex` (`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for `comer_account`
-- ----------------------------
DROP TABLE IF EXISTS `comer_account`;
CREATE TABLE `comer_account` (
                                 `id` bigint(20) NOT NULL,
                                 `comer_id` bigint(20) NOT NULL COMMENT 'comer unique identifier',
                                 `oin` varchar(100) NOT NULL COMMENT 'comer outer account unique identifier, wallet will be public key and Oauth is the OauthID',
                                 `is_primary` tinyint(1) NOT NULL COMMENT 'comer use this account as primay account',
                                 `nick` varchar(50) NOT NULL COMMENT 'comer nick name',
                                 `avatar` varchar(255) NOT NULL COMMENT 'avatar link address',
                                 `type` int(11) NOT NULL COMMENT '1 for github  2 for google 3 for twitter 4 for facebook 5 for likedin',
                                 `is_linked` tinyint(1) NOT NULL COMMENT '0 for unlink 1 for linked',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `comer_account_oin_uindex` (`oin`),
                                 KEY `comer_account_comer_id_index` (`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for `comer_profile`
-- ----------------------------
DROP TABLE IF EXISTS `comer_profile`;
CREATE TABLE `comer_profile` (
                                 `id` bigint(20) NOT NULL,
                                 `comer_id` bigint(20) NOT NULL,
                                 `name` varchar(50) NOT NULL COMMENT 'name',
                                 `avatar` varchar(200) NOT NULL COMMENT 'avatar',
                                 `location` char(42) NOT NULL DEFAULT '' COMMENT 'location city',
                                 `time_zone` varchar(50) DEFAULT NULL COMMENT 'time zone: UTC-09:30',
                                 `website` varchar(50) DEFAULT '' COMMENT 'website',
                                 `email` varchar(100) DEFAULT NULL COMMENT 'email',
                                 `twitter` varchar(100) DEFAULT NULL COMMENT 'twitter',
                                 `discord` varchar(100) DEFAULT NULL COMMENT 'discord',
                                 `telegram` varchar(100) DEFAULT NULL COMMENT 'telegram',
                                 `medium` varchar(100) DEFAULT NULL COMMENT 'medium',
                                 `bio` text COMMENT 'bio',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `comer_profile_comer_id_uindex` (`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for `image`
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image` (
                         `id` bigint(20) NOT NULL,
                         `category` varchar(20) NOT NULL,
                         `name` varchar(64) NOT NULL COMMENT 'name',
                         `url` varchar(200) NOT NULL COMMENT 'url',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `image_category_name_uindex` (`category`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for `startup`
-- ----------------------------
DROP TABLE IF EXISTS `startup`;
CREATE TABLE `startup` (
                           `id` bigint(20) NOT NULL,
                           `comer_id` bigint(20) NOT NULL COMMENT 'comer_id',
                           `name` varchar(100) NOT NULL COMMENT 'name',
                           `mode` smallint(6) NOT NULL COMMENT '0:NONE, 1:ESG, 2:NGO, 3:DAO, 4:COM',
                           `logo` varchar(200) NOT NULL COMMENT 'logo',
                           `mission` varchar(100) NOT NULL COMMENT 'logo',
                           `token_contract_address` char(42) NOT NULL COMMENT 'token contract address',
                           `overview` text NOT NULL COMMENT 'overview',
                           `tx_hash` varchar(200) DEFAULT NULL,
                           `is_set` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is set',
                           `kyc` varchar(200) DEFAULT NULL COMMENT 'KYC',
                           `contract_audit` varchar(200) DEFAULT NULL COMMENT 'contract audit',
                           `website` varchar(200) DEFAULT NULL COMMENT 'website',
                           `discord` varchar(200) DEFAULT NULL COMMENT 'discord',
                           `twitter` varchar(200) DEFAULT NULL COMMENT 'twitter',
                           `telegram` varchar(200) DEFAULT NULL COMMENT 'telegram',
                           `docs` varchar(200) DEFAULT NULL COMMENT 'docs',
                           `launch_network` int(11) DEFAULT NULL COMMENT 'chain id',
                           `token_name` varchar(100) DEFAULT NULL COMMENT 'token name',
                           `token_symbol` varchar(50) DEFAULT NULL COMMENT 'token symbol',
                           `total_supply` bigint(20) DEFAULT NULL COMMENT 'total supply',
                           `presale_start` datetime DEFAULT NULL COMMENT 'presale start date',
                           `presale_end` datetime DEFAULT NULL COMMENT 'presale end date',
                           `launch_date` datetime DEFAULT NULL COMMENT 'launch_date',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `startup_name_uindex` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for `startup_follow_rel`
-- ----------------------------
DROP TABLE IF EXISTS `startup_follow_rel`;
CREATE TABLE `startup_follow_rel` (
                                      `id` bigint(20) NOT NULL,
                                      `comer_id` bigint(20) NOT NULL COMMENT 'comer_id',
                                      `startup_id` bigint(20) NOT NULL COMMENT 'startup_id',
                                      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      PRIMARY KEY (`id`),
                                      UNIQUE KEY `startup_followed_comer_id_startup_id_uindex` (`comer_id`,`startup_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for `startup_team_member_rel`
-- ----------------------------
DROP TABLE IF EXISTS `startup_team_member_rel`;
CREATE TABLE `startup_team_member_rel` (
                                           `id` bigint(20) NOT NULL,
                                           `comer_id` bigint(20) NOT NULL COMMENT 'comer_id',
                                           `startup_id` bigint(20) NOT NULL COMMENT 'startup_id',
                                           `position` text NOT NULL COMMENT 'title',
                                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                           PRIMARY KEY (`id`),
                                           UNIQUE KEY `startup_team_rel_comer_id_startup_id_uindex` (`comer_id`,`startup_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------


-- ----------------------------
-- Table structure for `startup_wallet`
-- ----------------------------
DROP TABLE IF EXISTS `startup_wallet`;
CREATE TABLE `startup_wallet` (
                                  `id` bigint(20) NOT NULL,
                                  `comer_id` bigint(20) NOT NULL COMMENT 'comer_id',
                                  `startup_id` bigint(20) NOT NULL COMMENT 'startup_id',
                                  `wallet_name` varchar(100) NOT NULL COMMENT 'wallet name',
                                  `wallet_address` char(42) NOT NULL COMMENT 'wallet address',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `startup_wallet_startup_id_wallet_address_uindex` (`startup_id`,`wallet_address`),
                                  UNIQUE KEY `startup_wallet_startup_id_wallet_name_uindex` (`startup_id`,`wallet_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for `tag`
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
                       `id` bigint(20) NOT NULL AUTO_INCREMENT,
                       `name` varchar(64) NOT NULL COMMENT 'name',
                       `is_index` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is index',
                       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                       `category` varchar(20) NOT NULL,
                       PRIMARY KEY (`id`),
                       UNIQUE KEY `tag_category_name_uindex` (`name`,`category`),
                       UNIQUE KEY `tag_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=114112701034497 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for `tag_target_rel`
-- ----------------------------
DROP TABLE IF EXISTS `tag_target_rel`;
CREATE TABLE `tag_target_rel` (
                                  `id` bigint(20) NOT NULL,
                                  `target` varchar(20) NOT NULL COMMENT 'comerSkill,startup',
                                  `target_id` bigint(20) NOT NULL COMMENT 'target id',
                                  `tag_id` bigint(20) NOT NULL COMMENT 'skill id',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `comer_id_skill_id_uindex` (`target`,`target_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

