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

CREATE TABLE `chain_contract` (
  `id` bigint NOT NULL,
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'Chain ID',
  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Chain contract address',
  `project` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1 Startup, 2 Bounty, 3 Crowdfunding, 4 Gover',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1工厂合约、2子合约',
  `version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'contract version',
  `abi` text COLLATE utf8mb4_general_ci NOT NULL COMMENT 'abi json',
  `created_tx_hash` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'created tx hash',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Is deleted',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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

ALTER TABLE `comunion`.`startup` 
CHANGE COLUMN `is_set` `on_chain` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'whether it is on the chain' AFTER `tx_hash`