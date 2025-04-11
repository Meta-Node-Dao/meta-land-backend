alter table tag
    add category varchar(20)  not null;

 #//drop index tag_name_uindex on tag;

create unique index tag_category_name_uindex
	on tag (name, category);

CREATE TABLE image
(
    id         BIGINT                               NOT NULL
        PRIMARY KEY,
    category varchar(20) NOT NULL,
    name       VARCHAR(64)                          NOT NULL COMMENT 'name',
    url        VARCHAR(200)                          NOT NULL COMMENT 'url',
    created_at DATETIME   DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME   DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    is_deleted TINYINT(1) DEFAULT 0                 NOT NULL COMMENT 'Is Deleted',
    CONSTRAINT image_category_name_uindex
        UNIQUE (category,name)
)
    CHARSET = utf8mb4;

