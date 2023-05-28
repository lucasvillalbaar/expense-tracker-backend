DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS transaction_types;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS currencies;
DROP TABLE IF EXISTS accounts;

CREATE TABLE transaction_types (
    id   VARCHAR(255) PRIMARY KEY,
    type VARCHAR(255),
    name VARCHAR(255) UNIQUE
);

INSERT INTO transaction_types (id, type, name)
VALUES
    ('C', 'Credit', 'Crédito'),
    ('D', 'Debit', 'Débito'),
    ('T', 'Transfer', 'Transferencia');

CREATE TABLE categories (
    id               VARCHAR(255) PRIMARY KEY,
    category_name    VARCHAR(255),
    subcategory_name VARCHAR(255),
    CONSTRAINT unique_category_subcategory UNIQUE (category_name, subcategory_name)
);

INSERT INTO categories (id, category_name, subcategory_name)
VALUES
    ('1', 'Ingreso', 'Salario'),
    ('2', 'Compras', 'Electrónicos'),
    ('3', 'Hogar', 'Luz');

CREATE TABLE currencies (
    id     VARCHAR(255) PRIMARY KEY,
    name   VARCHAR(255) UNIQUE,
    symbol VARCHAR(255)
);

INSERT INTO currencies (id, name, symbol)
VALUES
    ('USD', 'US Dollar', 'u$d'),
    ('ARS', 'Peso Argentino', '$'),
    ('USDT', 'Theter', 'USDT');

CREATE TABLE accounts (
    id       VARCHAR(255) PRIMARY KEY,
    type     VARCHAR(255),
    name     VARCHAR(255) UNIQUE,
    balance  DECIMAL(18, 6),
    currency VARCHAR(255),
    FOREIGN KEY (currency) REFERENCES currencies (id)
);

INSERT INTO accounts (id, type, name, balance, currency)
VALUES
    ('BRUARS', 'Cuenta Bancaria', 'Brubank (ARS)', 0.00, 'ARS'),
    ('BRUUSD', 'Cuenta Bancaria', 'Brubank (USD)', 0.00, 'USD');

CREATE TABLE transactions (
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

CREATE TABLE transfers (
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
