/*
 Navicat Premium Data Transfer

 Source Server         : metaland
 Source Server Type    : MySQL
 Source Server Version : 80024
 Source Host           : 45.249.209.63:3306
 Source Schema         : metaland

 Target Server Type    : MySQL
 Target Server Version : 80024
 File Encoding         : 65001

 Date: 30/04/2025 13:44:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bounty
-- ----------------------------
DROP TABLE IF EXISTS `bounty`;
CREATE TABLE `bounty` (
                          `id` bigint NOT NULL,
                          `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
                          `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Transcation Hash',
                          `deposit_contract` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Contract Address',
                          `startup_id` bigint NOT NULL DEFAULT '0',
                          `comer_id` bigint NOT NULL DEFAULT '0',
                          `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                          `apply_cutoff_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `discussion_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                          `deposit_token_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                          `applicant_deposit` int NOT NULL DEFAULT '0',
                          `founder_deposit` int NOT NULL DEFAULT '0',
                          `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                          `payment_mode` int NOT NULL DEFAULT '0',
                          `status` tinyint(1) NOT NULL DEFAULT '0',
                          `total_reward_token` int NOT NULL DEFAULT '0',
                          `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_applicant
-- ----------------------------
DROP TABLE IF EXISTS `bounty_applicant`;
CREATE TABLE `bounty_applicant` (
                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                    `bounty_id` bigint DEFAULT NULL,
                                    `comer_id` bigint DEFAULT NULL,
                                    `apply_at` datetime DEFAULT NULL,
                                    `revoke_at` datetime DEFAULT NULL,
                                    `approve_at` datetime DEFAULT NULL,
                                    `quit_at` datetime DEFAULT NULL,
                                    `submit_at` datetime DEFAULT NULL,
                                    `status` tinyint(1) DEFAULT NULL,
                                    `description` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `created_at` datetime DEFAULT NULL,
                                    `updated_at` datetime DEFAULT NULL,
                                    `is_deleted` tinyint(1) DEFAULT NULL,
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_contact
-- ----------------------------
DROP TABLE IF EXISTS `bounty_contact`;
CREATE TABLE `bounty_contact` (
                                  `id` bigint NOT NULL,
                                  `bounty_id` bigint DEFAULT NULL,
                                  `contact_type` tinyint DEFAULT NULL,
                                  `contact_address` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                  `created_at` datetime DEFAULT NULL,
                                  `updated_at` datetime DEFAULT NULL,
                                  `is_deleted` tinyint(1) DEFAULT NULL,
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE KEY `bounty_contact_uindex` (`bounty_id`,`contact_type`,`contact_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_deposit
-- ----------------------------
DROP TABLE IF EXISTS `bounty_deposit`;
CREATE TABLE `bounty_deposit` (
                                  `id` bigint NOT NULL,
                                  `chain_id` bigint DEFAULT NULL,
                                  `tx_hash` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                  `status` tinyint(1) DEFAULT NULL,
                                  `bounty_id` bigint DEFAULT NULL,
                                  `comer_id` bigint DEFAULT NULL,
                                  `access` int DEFAULT NULL,
                                  `token_symbol` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                  `token_amount` int DEFAULT NULL,
                                  `timestamp` datetime DEFAULT NULL,
                                  `created_at` datetime DEFAULT NULL,
                                  `updated_at` datetime DEFAULT NULL,
                                  `is_deleted` tinyint(1) DEFAULT NULL,
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_payment_period
-- ----------------------------
DROP TABLE IF EXISTS `bounty_payment_period`;
CREATE TABLE `bounty_payment_period` (
                                         `id` bigint NOT NULL,
                                         `bounty_id` bigint DEFAULT NULL,
                                         `period_type` tinyint(1) DEFAULT NULL,
                                         `period_amount` bigint DEFAULT NULL,
                                         `hours_per_day` int DEFAULT NULL,
                                         `token1_symbol` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                         `token1_amount` int DEFAULT NULL,
                                         `token2_symbol` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                         `token2_amount` int DEFAULT NULL,
                                         `target` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                         `created_at` datetime DEFAULT NULL,
                                         `updated_at` datetime DEFAULT NULL,
                                         `is_deleted` tinyint(1) DEFAULT NULL,
                                         PRIMARY KEY (`id`) USING BTREE,
                                         UNIQUE KEY `bounty_id_uindex` (`bounty_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bounty_payment_terms
-- ----------------------------
DROP TABLE IF EXISTS `bounty_payment_terms`;
CREATE TABLE `bounty_payment_terms` (
                                        `id` bigint NOT NULL,
                                        `bounty_id` bigint DEFAULT NULL,
                                        `payment_mode` tinyint(1) DEFAULT NULL,
                                        `token1_symbol` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                        `token1_amount` int DEFAULT NULL,
                                        `token2_symbol` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                        `token2_amount` int DEFAULT NULL,
                                        `terms` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                        `seq_num` int DEFAULT NULL,
                                        `status` int DEFAULT NULL,
                                        `created_at` datetime DEFAULT NULL,
                                        `updated_at` datetime DEFAULT NULL,
                                        `is_deleted` tinyint(1) DEFAULT NULL,
                                        PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for chain
-- ----------------------------
DROP TABLE IF EXISTS `chain`;
CREATE TABLE `chain` (
                         `id` bigint NOT NULL,
                         `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
                         `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain name',
                         `logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain logo',
                         `status` tinyint(1) DEFAULT '1' COMMENT '1-normal, 2-disable',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `chain_id` (`chain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for chain_contract
-- ----------------------------
DROP TABLE IF EXISTS `chain_contract`;
CREATE TABLE `chain_contract` (
                                  `id` bigint NOT NULL,
                                  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
                                  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain contract address',
                                  `project` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1 Startup, 2 Bounty, 3 Crowdfunding, 4 Gover',
                                  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1工厂合约、2子合约',
                                  `version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'contract version',
                                  `abi` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'abi json',
                                  `created_tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'created tx hash',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for chain_endpoint
-- ----------------------------
DROP TABLE IF EXISTS `chain_endpoint`;
CREATE TABLE `chain_endpoint` (
                                  `id` bigint NOT NULL,
                                  `protocol` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Communication protocol, 1-rpc 2-wss',
                                  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
                                  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain name',
                                  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1-normal, 2-disable',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for comer
-- ----------------------------
DROP TABLE IF EXISTS `comer`;
CREATE TABLE `comer` (
                         `id` bigint NOT NULL,
                         `address` char(42) DEFAULT NULL COMMENT 'comer could save some useful info on block chain with this address',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `comer_address_uindex` (`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comer_account
-- ----------------------------
DROP TABLE IF EXISTS `comer_account`;
CREATE TABLE `comer_account` (
                                 `id` bigint NOT NULL,
                                 `comer_id` bigint NOT NULL COMMENT 'comer unique identifier',
                                 `oin` varchar(100) NOT NULL COMMENT 'comer outer account unique identifier, wallet will be public key and Oauth is the OauthID',
                                 `is_primary` tinyint(1) NOT NULL COMMENT 'comer use this account as primay account',
                                 `nick` varchar(50) NOT NULL COMMENT 'comer nick name',
                                 `avatar` varchar(255) NOT NULL COMMENT 'avatar link address',
                                 `type` int NOT NULL COMMENT '1 for github  2 for google 3 for twitter 4 for facebook 5 for likedin',
                                 `is_linked` tinyint(1) NOT NULL COMMENT '0 for unlink 1 for linked',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `comer_account_oin_uindex` (`oin`),
                                 KEY `comer_account_comer_id_index` (`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for comer_follow_rel
-- ----------------------------
DROP TABLE IF EXISTS `comer_follow_rel`;
CREATE TABLE `comer_follow_rel` (
                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                    `comer_id` bigint NOT NULL DEFAULT '0',
                                    `target_comer_id` bigint NOT NULL DEFAULT '0',
                                    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for comer_profile
-- ----------------------------
DROP TABLE IF EXISTS `comer_profile`;
CREATE TABLE `comer_profile` (
                                 `id` bigint NOT NULL,
                                 `comer_id` bigint NOT NULL,
                                 `name` varchar(50) NOT NULL COMMENT 'name',
                                 `avatar` varchar(200) NOT NULL COMMENT 'avatar',
                                 `cover` varchar(200) DEFAULT NULL,
                                 `location` char(42) NOT NULL DEFAULT '' COMMENT 'location city',
                                 `time_zone` varchar(50) DEFAULT NULL COMMENT 'time zone: UTC-09:30',
                                 `website` varchar(50) DEFAULT '' COMMENT 'website',
                                 `email` varchar(100) DEFAULT NULL COMMENT 'email',
                                 `twitter` varchar(100) DEFAULT NULL COMMENT 'twitter',
                                 `discord` varchar(100) DEFAULT NULL COMMENT 'discord',
                                 `telegram` varchar(100) DEFAULT NULL COMMENT 'telegram',
                                 `medium` varchar(100) DEFAULT NULL COMMENT 'medium',
                                 `facebook` varchar(256) DEFAULT NULL,
                                 `linktree` varchar(256) DEFAULT NULL,
                                 `bio` text COMMENT 'bio',
                                 `languages` varchar(256) DEFAULT NULL,
                                 `educations` varchar(1024) DEFAULT NULL,
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `comer_profile_comer_id_uindex` (`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for crowdfunding
-- ----------------------------
DROP TABLE IF EXISTS `crowdfunding`;
CREATE TABLE `crowdfunding` (
                                `id` bigint NOT NULL COMMENT 'crowdfunding id',
                                `chain_id` bigint NOT NULL COMMENT 'Chain id',
                                `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Tx hash',
                                `crowdfunding_contract` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Crowdfunding contract address',
                                `startup_id` bigint NOT NULL COMMENT 'Startup id',
                                `comer_id` bigint NOT NULL COMMENT 'Founder''s comer id',
                                `raise_goal` decimal(38,18) NOT NULL COMMENT 'Raise goal total',
                                `raise_balance` decimal(38,18) NOT NULL COMMENT 'Raise token balance',
                                `sell_token_contract` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Sell token contract address',
                                `sell_token_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Sell token name',
                                `sell_token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Sell token symbol',
                                `sell_token_decimals` int DEFAULT NULL COMMENT 'Sell token decimals',
                                `sell_token_supply` decimal(38,18) DEFAULT NULL COMMENT 'Sell token total supply',
                                `sell_token_deposit` decimal(38,18) NOT NULL COMMENT 'Sell token deposit',
                                `sell_token_balance` decimal(38,18) NOT NULL COMMENT 'Sell token balance',
                                `buy_token_contract` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Buy token contract address',
                                `buy_token_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Buy token name',
                                `buy_token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Buy token symbol',
                                `buy_token_decimals` int DEFAULT NULL COMMENT 'Buy token decimals',
                                `buy_token_supply` decimal(38,18) DEFAULT NULL COMMENT 'Buy token total supply',
                                `team_wallet` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Team wallet address',
                                `swap_percent` float NOT NULL COMMENT 'Swap percent',
                                `buy_price` decimal(38,18) NOT NULL COMMENT 'IBO rate',
                                `max_buy_amount` decimal(38,18) NOT NULL COMMENT 'Maximum buy amount',
                                `max_sell_percent` float NOT NULL COMMENT 'Maximum selling percent',
                                `sell_tax` float NOT NULL COMMENT 'Selling tax',
                                `start_time` datetime NOT NULL COMMENT 'Start time',
                                `end_time` datetime NOT NULL COMMENT 'End time',
                                `poster` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Poster url',
                                `youtube` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Youtube link',
                                `detail` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Detail url',
                                `description` varchar(520) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Description content',
                                `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:Pending 1:Upcoming 2:Live 3:Ended 4:Cancelled 5:Failure',
                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for crowdfunding_ibo_rate
-- ----------------------------
DROP TABLE IF EXISTS `crowdfunding_ibo_rate`;
CREATE TABLE `crowdfunding_ibo_rate` (
                                         `id` bigint NOT NULL,
                                         `crowdfunding_id` bigint NOT NULL COMMENT 'Crowdfunding id',
                                         `end_time` datetime NOT NULL COMMENT 'End time',
                                         `max_buy_amount` decimal(38,18) NOT NULL COMMENT 'Maximum buy amount',
                                         `max_sell_percent` float NOT NULL COMMENT 'Maximum sell percent',
                                         `buy_price` decimal(38,18) NOT NULL COMMENT 'IBO rate',
                                         `swap_percent` float NOT NULL COMMENT 'Swap percent',
                                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for crowdfunding_investor
-- ----------------------------
DROP TABLE IF EXISTS `crowdfunding_investor`;
CREATE TABLE `crowdfunding_investor` (
                                         `id` bigint NOT NULL,
                                         `crowdfunding_id` bigint NOT NULL COMMENT 'Crowdfunding id',
                                         `comer_id` bigint NOT NULL COMMENT 'Investor'' comer id',
                                         `buy_token_total` decimal(38,18) NOT NULL COMMENT 'Buy token total',
                                         `buy_token_balance` decimal(38,18) NOT NULL COMMENT 'Buy token balance',
                                         `sell_token_total` decimal(38,18) NOT NULL COMMENT 'Selling token total',
                                         `sell_token_balance` decimal(38,18) NOT NULL COMMENT 'Selling token balance',
                                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         PRIMARY KEY (`id`),
                                         UNIQUE KEY `crowdfunding_comer_uindex` (`crowdfunding_id`,`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for crowdfunding_swap
-- ----------------------------
DROP TABLE IF EXISTS `crowdfunding_swap`;
CREATE TABLE `crowdfunding_swap` (
                                     `id` bigint NOT NULL,
                                     `chain_id` bigint NOT NULL COMMENT 'Chain id',
                                     `tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Tx hash',
                                     `timestamp` datetime DEFAULT NULL,
                                     `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:Pending 1:Success 2:Failure',
                                     `crowdfunding_id` bigint NOT NULL COMMENT 'Crowdfunding id',
                                     `comer_id` bigint NOT NULL COMMENT 'Comer id',
                                     `access` tinyint(1) NOT NULL COMMENT '1:Invest 2:Withdraw',
                                     `buy_token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Buy token symbol',
                                     `buy_token_amount` decimal(38,18) NOT NULL COMMENT 'Buy token amount',
                                     `sell_token_symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Selling token symbol',
                                     `sell_token_amount` decimal(38,18) NOT NULL COMMENT 'Selling token amount',
                                     `price` decimal(38,18) NOT NULL COMMENT 'Swap price',
                                     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`),
                                     UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for dict_data
-- ----------------------------
DROP TABLE IF EXISTS `dict_data`;
CREATE TABLE `dict_data` (
                             `id` int NOT NULL,
                             `startup_id` bigint DEFAULT NULL,
                             `dict_type` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                             `dict_label` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                             `dict_value` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                             `seq_num` int DEFAULT NULL,
                             `status` tinyint(1) DEFAULT NULL COMMENT '1:enabled 2:disabled',
                             `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                             `created_at` datetime DEFAULT NULL,
                             `updated_at` datetime DEFAULT NULL,
                             `is_deleted` tinyint(1) DEFAULT NULL,
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_admin
-- ----------------------------
DROP TABLE IF EXISTS `governance_admin`;
CREATE TABLE `governance_admin` (
                                    `id` int NOT NULL,
                                    `setting_id` bigint DEFAULT NULL,
                                    `wallet_address` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                    `created_at` datetime DEFAULT NULL,
                                    `updated_at` datetime DEFAULT NULL,
                                    `is_deleted` tinyint(1) DEFAULT NULL,
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_choice
-- ----------------------------
DROP TABLE IF EXISTS `governance_choice`;
CREATE TABLE `governance_choice` (
                                     `id` bigint NOT NULL,
                                     `proposal_id` bigint DEFAULT NULL,
                                     `item_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                     `seq_num` tinyint DEFAULT NULL,
                                     `created_at` datetime DEFAULT NULL,
                                     `updated_at` datetime DEFAULT NULL,
                                     `is_deleted` tinyint(1) DEFAULT NULL,
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_proposal
-- ----------------------------
DROP TABLE IF EXISTS `governance_proposal`;
CREATE TABLE `governance_proposal` (
                                       `id` int NOT NULL,
                                       `startup_id` bigint DEFAULT NULL,
                                       `author_comer_id` bigint DEFAULT NULL,
                                       `author_wallet_address` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `chain_id` bigint DEFAULT NULL,
                                       `block_number` bigint DEFAULT NULL,
                                       `release_timestamp` datetime DEFAULT NULL,
                                       `ipfs_hash` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `title` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `description` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `discussion_link` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `vote_system` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `start_time` datetime DEFAULT NULL,
                                       `end_time` datetime DEFAULT NULL,
                                       `status` tinyint(1) DEFAULT NULL COMMENT '0:pending 1:upcoming 2:active 3:ended',
                                       `created_at` datetime DEFAULT NULL,
                                       `updated_at` datetime DEFAULT NULL,
                                       `is_deleted` tinyint(1) DEFAULT NULL,
                                       PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_setting
-- ----------------------------
DROP TABLE IF EXISTS `governance_setting`;
CREATE TABLE `governance_setting` (
                                      `id` bigint NOT NULL,
                                      `startup_id` bigint DEFAULT NULL,
                                      `comer_id` bigint DEFAULT NULL,
                                      `vote_symbol` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                      `allow_member` tinyint(1) DEFAULT NULL COMMENT '0:no  1:yes',
                                      `proposal_threshold` double DEFAULT NULL,
                                      `proposal_validity` double DEFAULT NULL,
                                      `created_at` datetime DEFAULT NULL,
                                      `updated_at` datetime DEFAULT NULL,
                                      `is_deleted` tinyint(1) DEFAULT NULL,
                                      PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_strategy
-- ----------------------------
DROP TABLE IF EXISTS `governance_strategy`;
CREATE TABLE `governance_strategy` (
                                       `id` int NOT NULL,
                                       `setting_id` bigint DEFAULT NULL,
                                       `dict_value` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `strategy_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `chain_id` bigint DEFAULT NULL,
                                       `token_contract_address` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `vote_symbol` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                       `vote_decimals` int DEFAULT NULL,
                                       `token_min_balance` double DEFAULT NULL,
                                       `created_at` datetime DEFAULT NULL,
                                       `updated_at` datetime DEFAULT NULL,
                                       `is_deleted` tinyint(1) DEFAULT NULL,
                                       PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for governance_vote
-- ----------------------------
DROP TABLE IF EXISTS `governance_vote`;
CREATE TABLE `governance_vote` (
                                   `id` bigint NOT NULL,
                                   `proposal_id` bigint DEFAULT NULL,
                                   `voter_comer_id` bigint DEFAULT NULL,
                                   `voter_wallet_address` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `choice_item_id` bigint DEFAULT NULL,
                                   `choice_item_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `votes` double DEFAULT NULL,
                                   `ipfs_hash` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                   `created_at` datetime DEFAULT NULL,
                                   `updated_at` datetime DEFAULT NULL,
                                   `is_deleted` tinyint(1) DEFAULT NULL,
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image` (
                         `id` bigint NOT NULL,
                         `category` varchar(20) NOT NULL,
                         `name` varchar(64) NOT NULL COMMENT 'name',
                         `url` varchar(200) NOT NULL COMMENT 'url',
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `image_category_name_uindex` (`category`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for post_update
-- ----------------------------
DROP TABLE IF EXISTS `post_update`;
CREATE TABLE `post_update` (
                               `id` int NOT NULL,
                               `source_type` tinyint(1) DEFAULT NULL,
                               `source_id` bigint DEFAULT NULL,
                               `comer_id` bigint DEFAULT NULL,
                               `content` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `timestamp` datetime DEFAULT NULL,
                               `created_at` datetime DEFAULT NULL,
                               `updated_at` datetime DEFAULT NULL,
                               `is_deleted` tinyint(1) DEFAULT NULL,
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for startup
-- ----------------------------
DROP TABLE IF EXISTS `startup`;
CREATE TABLE `startup` (
                           `id` bigint NOT NULL,
                           `comer_id` bigint NOT NULL COMMENT 'comer_id',
                           `name` varchar(100) NOT NULL COMMENT 'name',
                           `mode` smallint NOT NULL COMMENT '0:NONE, 1:ESG, 2:NGO, 3:DAO, 4:COM',
                           `logo` varchar(200) NOT NULL COMMENT 'logo',
                           `cover` varchar(200) DEFAULT NULL,
                           `mission` varchar(100) NOT NULL COMMENT 'logo',
                           `token_contract_address` char(42) NOT NULL COMMENT 'token contract address',
                           `overview` text NOT NULL COMMENT 'overview',
                           `tx_hash` varchar(200) DEFAULT NULL,
                           `on_chain` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'whether it is on the chain',
                           `kyc` varchar(200) DEFAULT NULL COMMENT 'KYC',
                           `contract_audit` varchar(200) DEFAULT NULL COMMENT 'contract audit',
                           `website` varchar(200) DEFAULT NULL COMMENT 'website',
                           `discord` varchar(200) DEFAULT NULL COMMENT 'discord',
                           `twitter` varchar(200) DEFAULT NULL COMMENT 'twitter',
                           `telegram` varchar(200) DEFAULT NULL COMMENT 'telegram',
                           `docs` varchar(200) DEFAULT NULL COMMENT 'docs',
                           `email` varchar(180) DEFAULT NULL,
                           `facebook` varchar(180) DEFAULT NULL,
                           `medium` varchar(180) DEFAULT NULL,
                           `linktree` varchar(180) DEFAULT NULL,
                           `launch_network` int DEFAULT NULL COMMENT 'chain id',
                           `token_name` varchar(100) DEFAULT NULL COMMENT 'token name',
                           `token_symbol` varchar(50) DEFAULT NULL COMMENT 'token symbol',
                           `total_supply` bigint DEFAULT NULL COMMENT 'total supply',
                           `presale_start` datetime DEFAULT NULL COMMENT 'presale start date',
                           `presale_end` datetime DEFAULT NULL COMMENT 'presale end date',
                           `launch_date` datetime DEFAULT NULL COMMENT 'launch_date',
                           `tab_sequence` varchar(200) DEFAULT NULL,
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `startup_name_uindex` (`name`),
                           UNIQUE KEY `startup_token_contract_index` (`token_contract_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_follow_rel
-- ----------------------------
DROP TABLE IF EXISTS `startup_follow_rel`;
CREATE TABLE `startup_follow_rel` (
                                      `id` bigint NOT NULL,
                                      `comer_id` bigint NOT NULL COMMENT 'comer_id',
                                      `startup_id` bigint NOT NULL COMMENT 'startup_id',
                                      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      PRIMARY KEY (`id`),
                                      UNIQUE KEY `startup_followed_comer_id_startup_id_uindex` (`comer_id`,`startup_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_group
-- ----------------------------
DROP TABLE IF EXISTS `startup_group`;
CREATE TABLE `startup_group` (
                                 `id` bigint NOT NULL,
                                 `comer_id` bigint NOT NULL COMMENT 'comer_id',
                                 `startup_id` bigint NOT NULL COMMENT 'startup_id',
                                 `name` varchar(200) NOT NULL COMMENT 'group name',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `startup_group_name_unidex` (`name`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_group_member_rel
-- ----------------------------
DROP TABLE IF EXISTS `startup_group_member_rel`;
CREATE TABLE `startup_group_member_rel` (
                                            `id` bigint NOT NULL,
                                            `comer_id` bigint NOT NULL COMMENT 'comer_id',
                                            `startup_id` bigint NOT NULL COMMENT 'startup_id',
                                            `group_id` bigint NOT NULL COMMENT 'group id',
                                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                            PRIMARY KEY (`id`),
                                            UNIQUE KEY `startup_group_comer_id_uindex` (`comer_id`,`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_team_member_rel
-- ----------------------------
DROP TABLE IF EXISTS `startup_team_member_rel`;
CREATE TABLE `startup_team_member_rel` (
                                           `id` bigint NOT NULL,
                                           `comer_id` bigint NOT NULL COMMENT 'comer_id',
                                           `startup_id` bigint NOT NULL COMMENT 'startup_id',
                                           `position` text NOT NULL COMMENT 'title',
                                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                           PRIMARY KEY (`id`),
                                           UNIQUE KEY `startup_team_rel_comer_id_startup_id_uindex` (`comer_id`,`startup_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for startup_wallet
-- ----------------------------
DROP TABLE IF EXISTS `startup_wallet`;
CREATE TABLE `startup_wallet` (
                                  `id` bigint NOT NULL,
                                  `comer_id` bigint NOT NULL COMMENT 'comer_id',
                                  `startup_id` bigint NOT NULL COMMENT 'startup_id',
                                  `wallet_name` varchar(100) NOT NULL COMMENT 'wallet name',
                                  `wallet_address` char(42) NOT NULL COMMENT 'wallet address',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
                       `id` bigint NOT NULL AUTO_INCREMENT,
                       `name` varchar(64) NOT NULL COMMENT 'name',
                       `is_index` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is index',
                       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
                       `category` varchar(20) NOT NULL,
                       PRIMARY KEY (`id`),
                       UNIQUE KEY `tag_category_name_uindex` (`name`,`category`),
                       UNIQUE KEY `tag_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=114112701034523 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for tag_target_rel
-- ----------------------------
DROP TABLE IF EXISTS `tag_target_rel`;
CREATE TABLE `tag_target_rel` (
                                  `id` bigint NOT NULL,
                                  `target` varchar(20) NOT NULL COMMENT 'comerSkill,startup',
                                  `target_id` bigint NOT NULL COMMENT 'target id',
                                  `tag_id` bigint NOT NULL COMMENT 'skill id',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `comer_id_skill_id_uindex` (`target`,`target_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
                               `id` bigint NOT NULL,
                               `chain_id` bigint DEFAULT NULL,
                               `tx_hash` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `timestamp` datetime DEFAULT NULL,
                               `status` tinyint(1) DEFAULT NULL COMMENT '0:Pending 1:Success 2:Failure',
                               `source_type` tinyint(1) DEFAULT NULL,
                               `source_id` bigint DEFAULT NULL,
                               `retry_times` int DEFAULT NULL,
                               `created_at` datetime DEFAULT NULL,
                               `updated_at` datetime DEFAULT NULL,
                               `is_deleted` tinyint(1) DEFAULT NULL,
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE KEY `chain_tx_uindex` (`chain_id`,`tx_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
