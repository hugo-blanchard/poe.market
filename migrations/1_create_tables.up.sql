CREATE TABLE item (
    id TEXT PRIMARY KEY,
    stash_id TEXT NOT NULL,
    itemname TEXT NOT NULL,
    price TEXT NOT NULL
);

CREATE TABLE solditem (
    id TEXT PRIMARY KEY,
    itemname TEXT NOT NULL,
    price TEXT NOT NULL
);