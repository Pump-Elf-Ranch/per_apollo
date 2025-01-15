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


CREATE TABLE IF NOT EXISTS mint_nft_listed
(
    guid             text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    mint_address     varchar,
    timestamp        UINT256,
    nonce            UINT256,
    contract_address varchar,
    mint_type        varchar,
    block_number     UINT256,
    tx_hash          varchar
);
CREATE INDEX IF NOT EXISTS mint_nft_listed_tx_hash ON mint_nft_listed (tx_hash);
CREATE INDEX IF NOT EXISTS mint_nft_listed_mint_address ON mint_nft_listed (mint_address);
CREATE INDEX IF NOT EXISTS mint_nft_listed_contract_address ON mint_nft_listed (contract_address);
CREATE INDEX IF NOT EXISTS mint_nft_listed_mint_type ON mint_nft_listed (mint_type);

CREATE TABLE IF NOT EXISTS prop_buy_listed
(
    guid             text PRIMARY KEY DEFAULT replace(uuid_generate_v4()::text, '-', ''),
    buyer            varchar,
    item_id          UINT256,
    timestamp        UINT256,
    price            UINT256,
    item_type        varchar,
    contract_address varchar,
    block_number     UINT256,
    is_deposit       int2,
    tx_hash          varchar
);

CREATE INDEX IF NOT EXISTS prop_buy_listed_tx_hash ON prop_buy_listed (tx_hash);
CREATE INDEX IF NOT EXISTS prop_buy_listed_buyer ON prop_buy_listed (buyer);
CREATE INDEX IF NOT EXISTS prop_buy_listed_contract_address ON prop_buy_listed (contract_address);
CREATE INDEX IF NOT EXISTS prop_buy_listed_item_type ON prop_buy_listed (item_type);
CREATE INDEX IF NOT EXISTS prop_buy_listed_item_id ON prop_buy_listed (item_id);

