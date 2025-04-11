alter table `comer_profile`
    add column `cover`    varchar(200) default null after `avatar`,
    add column `facebook` varchar(256) default null after `medium`,
    add column `linktree` varchar(256) default null after `facebook`;


alter table `startup`
    add column `cover` varchar(200) default null after `logo`;

alter table `startup`
    add column `tab_sequence` varchar(200) default null after `launch_date`;

drop table if exists `startup_group`;
create table if not exists `startup_group`
(
    `id`         bigint(20)   NOT NULL,
    `comer_id`   bigint(20)   NOT NULL COMMENT 'comer_id',
    `startup_id` bigint(20)   NOT NULL COMMENT 'startup_id',
    `name`       varchar(200) NOT NULL COMMENT 'group name',
    `created_at` datetime     NOT NULL DEFAULT current_timestamp(),
    `updated_at` datetime     NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`),
    UNIQUE KEY `startup_group_name_unidex` (`name`, `id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


--
drop table if exists `startup_group_member_rel`;
create table if not exists `startup_group_member_rel`
(
    `id`         bigint(20) NOT NULL,
    `comer_id`   bigint(20) NOT NULL COMMENT 'comer_id',
    `startup_id` bigint(20) NOT NULL COMMENT 'startup_id',
    `group_id`   bigint(20) NOT NULL COMMENT 'group id',
    `created_at` datetime   NOT NULL DEFAULT current_timestamp(),
    `updated_at` datetime   NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`),
    UNIQUE KEY `startup_group_comer_id_uindex` (`comer_id`, `group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ff

alter table `startup`
    add column `email`    varchar(180) default null after `docs`,
    add column `facebook` varchar(180) default null after `email`,
    add column `medium`   varchar(180) default null after `facebook`,
    add column `linktree` varchar(180) default null after `medium`;


alter table `comer_profile`
    add column `languages`  varchar(256)  default null after `bio`,
    add column `educations` varchar(1024) default null after `languages`;