--
-- DDL of feature crowdfunding
--
create table crowdfunding
(
    id                    bigint                               not null comment 'crowdfunding id'
        primary key,
    chain_id              bigint                               not null comment 'Chain id',
    tx_hash               varchar(200)                         not null comment 'Tx hash',
    crowdfunding_contract varchar(50)                          null comment 'Crowdfunding contract address',
    startup_id            bigint                               not null comment 'Startup id',
    comer_id              bigint                               not null comment 'Founder''s comer id',
    raise_goal            decimal(38, 18)                      not null comment 'Raise goal total',
    raise_balance         decimal(38, 18)                      not null comment 'Raise token balance',
    sell_token_contract   varchar(50)                          not null comment 'Sell token contract address',
    sell_token_name       varchar(100)                         null comment 'Sell token name',
    sell_token_symbol     varchar(50)                          null comment 'Sell token symbol',
    sell_token_decimals   int                                  null comment 'Sell token decimals',
    sell_token_supply     decimal(38, 18)                      null comment 'Sell token total supply',
    sell_token_deposit    decimal(38, 18)                      not null comment 'Sell token deposit',
    sell_token_balance    decimal(38, 18)                      not null comment 'Sell token balance',
    buy_token_contract    varchar(50)                          not null comment 'Buy token contract address',
    buy_token_name        varchar(100)                         null comment 'Buy token name',
    buy_token_symbol      varchar(50)                          null comment 'Buy token symbol',
    buy_token_decimals    int                                  null comment 'Buy token decimals',
    buy_token_supply      decimal(38, 18)                      null comment 'Buy token total supply',
    team_wallet           varchar(50)                          not null comment 'Team wallet address',
    swap_percent          float                                not null comment 'Swap percent',
    buy_price             decimal(38, 18)                      not null comment 'IBO rate',
    max_buy_amount        decimal(38, 18)                      not null comment 'Maximum buy amount',
    max_sell_percent      float                                not null comment 'Maximum selling percent',
    sell_tax              float                                not null comment 'Selling tax',
    start_time            datetime                             not null comment 'Start time',
    end_time              datetime                             not null comment 'End time',
    poster                varchar(200)                         not null comment 'Poster url',
    youtube               varchar(200)                         null comment 'Youtube link',
    detail                varchar(200)                         null comment 'Detail url',
    description           varchar(520)                         not null comment 'Description content',
    status                tinyint(1) default 0                 not null comment '0:Pending 1:Upcoming 2:Live 3:Ended 4:Cancelled 5:Failure',
    created_at            datetime   default CURRENT_TIMESTAMP not null,
    updated_at            datetime   default CURRENT_TIMESTAMP not null,
    is_deleted            tinyint(1) default 0                 not null comment 'Is deleted',
    constraint chain_tx_uindex
        unique (chain_id, tx_hash)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


create table `crowdfunding_swap`
(
    id                bigint                               not null
        primary key,
    chain_id          bigint                               not null comment 'Chain id',
    tx_hash           varchar(200)                         not null comment 'Tx hash',
    timestamp         datetime                             null,
    status            tinyint(1) default 0                 not null comment '0:Pending 1:Success 2:Failure',
    crowdfunding_id   bigint                               not null comment 'Crowdfunding id',
    comer_id          bigint                               not null comment 'Comer id',
    access            tinyint(1)                           not null comment '1:Invest 2:Withdraw',
    buy_token_symbol  varchar(50)                          not null comment 'Buy token symbol',
    buy_token_amount  decimal(38, 18)                      not null comment 'Buy token amount',
    sell_token_symbol varchar(50)                          not null comment 'Selling token symbol',
    sell_token_amount decimal(38, 18)                      not null comment 'Selling token amount',
    price             decimal(38, 18)                      not null comment 'Swap price',
    created_at        datetime   default CURRENT_TIMESTAMP not null,
    updated_at        datetime   default CURRENT_TIMESTAMP not null,
    constraint chain_tx_uindex
        unique (chain_id, tx_hash)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;;


CREATE TABLE `crowdfunding_ibo_rate`  (
                                         `id` bigint(20) NOT NULL,
                                         `crowdfunding_id` bigint(20) NOT NULL COMMENT 'Crowdfunding id',
                                         `end_time` datetime NOT NULL COMMENT 'End time',
    -- `buy_token_symbol` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Buy token symbol',
                                         `max_buy_amount` decimal(38,18) NOT NULL COMMENT 'Maximum buy amount',
    -- `sell_token_symbol` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Selling token symbol',
                                         `max_sell_percent` float NOT NULL COMMENT 'Maximum sell percent',
                                         `buy_price` decimal(38,18) NOT NULL COMMENT 'IBO rate',
                                         `swap_percent` float NOT NULL COMMENT 'Swap percent',
                                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `crowdfunding_investor` (
                                         `id` bigint(20) NOT NULL,
                                         `crowdfunding_id` bigint(20) NOT NULL COMMENT 'Crowdfunding id',
                                         `comer_id` bigint(20) NOT NULL COMMENT 'Investor'' comer id',
                                         `buy_token_total` decimal(38,18) NOT NULL COMMENT 'Buy token total',
                                         `buy_token_balance` decimal(38,18) NOT NULL COMMENT 'Buy token balance',
                                         `sell_token_total` decimal(38,18) NOT NULL COMMENT 'Selling token total',
                                         `sell_token_balance` decimal(38,18) NOT NULL COMMENT 'Selling token balance',
                                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         PRIMARY KEY (`id`),
                                         UNIQUE KEY `crowdfunding_comer_uindex` (`crowdfunding_id`,`comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;