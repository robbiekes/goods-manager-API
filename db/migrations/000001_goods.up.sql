CREATE TABLE IF NOT EXISTS storages
(
    id      INTEGER      NOT NULL UNIQUE,
    name    VARCHAR(299) NOT NULL,
    allowed BOOLEAN
);

CREATE TABLE IF NOT EXISTS items
(
    id   INTEGER      NOT NULL UNIQUE,
    name VARCHAR(299) NOT NULL,
    size VARCHAR(10)  NOT NULL
);

CREATE TABLE IF NOT EXISTS items_storages
(
    item_id    INTEGER NOT NULL REFERENCES items (id) ON DELETE RESTRICT,
    storage_id INTEGER NOT NULL REFERENCES storages (id) ON DELETE RESTRICT,
    amount     INTEGER,
    reserved   INTEGER
);
