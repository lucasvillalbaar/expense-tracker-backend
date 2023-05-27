CREATE TABLE IF NOT EXISTS transaction_types (
    id   VARCHAR(255) PRIMARY KEY,
    type VARCHAR(255),
    name VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS categories (
    id               VARCHAR(255),
    category_name    VARCHAR(255),
    subcategory_name VARCHAR(255),
    CONSTRAINT unique_category_subcategory UNIQUE (category_name, subcategory_name)
);

CREATE TABLE IF NOT EXISTS currencies (
    id     VARCHAR(255) PRIMARY KEY,
    name   VARCHAR(255) UNIQUE,
    symbol VARCHAR(255),
);

CREATE TABLE IF NOT EXISTS accounts (
    id       VARCHAR(255) PRIMARY KEY,
    type     VARCHAR(255),
    name     VARCHAR(255) UNIQUE,
    balance  DECIMAL(18, 2),
    currency VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS transactions (
    id              VARCHAR(255) PRIMARY KEY,
    created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
    type            VARCHAR(255),
    category        VARCHAR(255),
    description     VARCHAR(255),
    account         VARCHAR(255),
    original_amount DECIMAL(18, 6),
    currency        VARCHAR(255),
    base_amount     DECIMAL(18, 6),
    FOREIGN KEY (account) REFERENCES accounts (id),
    FOREIGN KEY (category) REFERENCES categories (id)
);

CREATE TABLE IF NOT EXISTS transfers (
    id                  VARCHAR(255) PRIMARY KEY,
    created_at          TIMESTAMP NOT NULL DEFAULT NOW(),
    source_account      VARCHAR(255),
    source_amount       DECIMAL(18, 6),
    source_fee          DECIMAL(18, 6),
    destination_account VARCHAR(255),
    destination_amount  DECIMAL(18, 6),
    destination_fee     DECIMAL(18, 6),
    FOREIGN KEY (source_account) REFERENCES accounts (id),
    FOREIGN KEY (destination_account) REFERENCES accounts (id)
);
