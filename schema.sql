CREATE DATABASE IF NOT EXISTS the_company_db;

USE the_company_db;

CREATE TABLE IF NOT EXISTS Companies
(
    id          BINARY(16)                                                              NOT NULL,
    name        VARCHAR(15)                                                             NOT NULL UNIQUE,
    description VARCHAR(3000),
    employees   INT                                                                     NOT NULL,
    registered  BOOL                                                                    NOT NULL,
    type        ENUM ('Corporation', 'NonProfit', 'Cooperative', 'Sole Proprietorship') NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;