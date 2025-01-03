DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
            CREATE DOMAIN UINT256 AS NUMERIC
                CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
        ELSE
            ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
            ALTER DOMAIN UINT256 ADD
                CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
        END IF;
    END
$$;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" cascade;
-- CREATE EXTENSION "uuid-ossp";

CREATE TABLE IF NOT EXISTS template_block_headers
(
    guid        text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    hash        VARCHAR NOT NULL,
    parent_hash VARCHAR NOT NULL,
    number      UINT256 NOT NULL,
    timestamp   INTEGER NOT NULL,
    rlp_bytes   VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS template_block_headers_timestamp ON template_block_headers (timestamp);
CREATE INDEX IF NOT EXISTS template_block_headers_number ON template_block_headers (number);

CREATE TABLE IF NOT EXISTS template_contract_events
(
    guid             text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    block_hash       VARCHAR NOT NULL,
    contract_address VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    log_index        INTEGER NOT NULL,
    block_number     UINT256 NOT NULL,
    event_signature  VARCHAR NOT NULL,
    timestamp        INTEGER NOT NULL,
    rlp_bytes        VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS template_contract_events_number ON template_contract_events (block_number);
CREATE INDEX IF NOT EXISTS template_contract_events_timestamp ON template_contract_events (timestamp);
CREATE INDEX IF NOT EXISTS template_contract_events_block_hash ON template_contract_events (block_hash);
CREATE INDEX IF NOT EXISTS template_contract_events_event_signature ON template_contract_events (event_signature);
CREATE INDEX IF NOT EXISTS template_contract_events_contract_address ON template_contract_events (contract_address);
CREATE INDEX IF NOT EXISTS template_contract_events_transaction_hash ON template_contract_events (transaction_hash);


CREATE TABLE IF NOT EXISTS block_listener
(
    guid         text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    block_number "public"."uint256"                           DEFAULT 0,
    chain_id     varchar,
    created      int4,
    updated      int4,
    CONSTRAINT "block_listener_pkey" PRIMARY KEY (guid),
    CONSTRAINT "block_listener_created_check" CHECK (created > 0),
    CONSTRAINT "block_listener_updated_check" CHECK (updated > 0)
);


CREATE TABLE IF NOT EXISTS runes_order
(
    guid                    text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    ticker                  VARCHAR,
    symbol                  VARCHAR,
    token_address           VARCHAR,
    token_decimal           int8,
    status                  int2,
    amount                  VARCHAR,
    price                   VARCHAR,
    total                   VARCHAR,
    seller                  varchar,
    buyer                   varchar,
    timestamp               UINT256,
    order_id                UINT256,
    market_contract_address varchar,
    order_type              int2,
    block_number            UINT256,
    tx_hash                 varchar,
    order_tx_hash           varchar
);
CREATE INDEX IF NOT EXISTS runes_order_order_tx_hash ON runes_order (order_tx_hash);
CREATE INDEX IF NOT EXISTS runes_order_tx_hash ON runes_order (tx_hash);
CREATE INDEX IF NOT EXISTS runes_order_seller ON runes_order (seller);
CREATE INDEX IF NOT EXISTS runes_order_buyer ON runes_order (buyer);
CREATE INDEX IF NOT EXISTS runes_order_order_id ON runes_order (order_id);

CREATE TABLE IF NOT EXISTS runes_listed
(
    guid                    text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    ticker                  VARCHAR,
    symbol                  VARCHAR,
    token_address           VARCHAR,
    token_decimal           int8,
    amount                  VARCHAR,
    price                   VARCHAR,
    total                   VARCHAR,
    seller                  varchar,
    order_id                UINT256,
    status                  int2,
    create_time             UINT256,
    update_time             UINT256,
    market_contract_address varchar,
    block_number            UINT256,
    tx_hash                 varchar
);

CREATE INDEX IF NOT EXISTS runes_listed_tx_hash ON runes_listed (tx_hash);
CREATE INDEX IF NOT EXISTS runes_listed_seller ON runes_listed (seller);
CREATE INDEX IF NOT EXISTS runes_listed_create_time ON runes_listed (create_time);
CREATE INDEX IF NOT EXISTS runes_listed_update_time ON runes_listed (update_time);

CREATE INDEX IF NOT EXISTS runes_listed_order_id ON runes_listed (order_id);

CREATE TABLE IF NOT EXISTS runes_activity
(
    guid                    text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    ticker                  VARCHAR,
    symbol                  VARCHAR,
    token_address           VARCHAR,
    token_decimal           int8,
    amount                  VARCHAR,
    price                   VARCHAR,
    total                   VARCHAR,
    seller                  varchar,
    buyer                   varchar,
    order_id                UINT256,
    activity_type           int2,
    timestamp               UINT256,
    market_contract_address varchar,
    block_number            UINT256,
    tx_hash                 varchar
);

CREATE INDEX IF NOT EXISTS runes_activity_tx_hash ON runes_activity (tx_hash);
CREATE INDEX IF NOT EXISTS runes_activity_seller ON runes_activity (seller);
CREATE INDEX IF NOT EXISTS runes_activity_buyer ON runes_activity (buyer);
CREATE INDEX IF NOT EXISTS runes_activity_timestamp ON runes_activity (timestamp);
CREATE INDEX IF NOT EXISTS runes_activity_order_id ON runes_activity (order_id);
