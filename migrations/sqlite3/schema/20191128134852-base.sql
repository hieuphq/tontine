
-- +migrate Up
CREATE TABLE IF NOT EXISTS groups
(
        id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
        name VARCHAR (100) NOT NULL,
        strategy_percent DOUBLE DEFAULT 0,
        amount DOUBLE DEFAULT 0 NOT NULL,
        currency VARCHAR(5) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        deleted_at DATETIME DEFAULT NULL
);

CREATE INDEX groups_id on groups(id);

CREATE TABLE IF NOT EXISTS investors
(
        id INTEGER PRIMARY KEY NOT NULL,
        name VARCHAR (100) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        deleted_at DATETIME DEFAULT NULL
);

CREATE INDEX investors_id on investors(id);

CREATE TABLE IF NOT EXISTS groups_investors
(
        id INTEGER PRIMARY KEY NOT NULL,
        investor_id INTEGER REFERENCES investors (id) NOT NULL,
        group_id INTEGER REFERENCES groups (id) NOT NULL,
        amount DOUBLE DEFAULT 0 NOT NULL,
        currency VARCHAR(5) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        deleted_at DATETIME DEFAULT NULL
);

CREATE INDEX groups_investors_id on groups_investors(id);

CREATE TABLE IF NOT EXISTS activity_logs
(
        id INTEGER PRIMARY KEY NOT NULL,
        investor_id INTEGER REFERENCES investors (id) NOT NULL,
        group_id INTEGER REFERENCES groups (id) DEFAULT NULL,
        action VARCHAR (50) NOT NULL,
        amount DOUBLE DEFAULT 0,
        currency VARCHAR(5) DEFAULT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        deleted_at DATETIME DEFAULT NULL
);

CREATE INDEX activity_logs_id on activity_logs(id);

CREATE TABLE IF NOT EXISTS group_logs
(
        id INTEGER PRIMARY KEY NOT NULL,
        name VARCHAR(255) DEFAULT '' NOT NULL,
        group_id INTEGER REFERENCES groups (id) NOT NULL,
        amount DOUBLE DEFAULT 0 NOT NULL,
        currency VARCHAR(5) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        deleted_at DATETIME DEFAULT NULL
);

CREATE INDEX group_logs_id on group_logs(id);

CREATE TABLE IF NOT EXISTS group_log_items
(
        id INTEGER PRIMARY KEY NOT NULL,
        group_log_id INTEGER REFERENCES group_logs (id) NOT NULL,
        amount DOUBLE DEFAULT 0 NOT NULL,
        currency VARCHAR(5) NOT NULL,
        investor_id INTEGER REFERENCES investors (id) DEFAULT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
        deleted_at DATETIME DEFAULT NULL
);

CREATE INDEX group_log_items_id on group_log_items(id);

-- +migrate Down
DROP TABLE group_log_items;
DROP TABLE group_logs;
DROP TABLE activity_logs;
DROP TABLE groups_investors;
DROP TABLE investors;
DROP TABLE groups;