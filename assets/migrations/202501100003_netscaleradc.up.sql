CREATE TABLE IF NOT EXISTS connections_netscaleradc_protocols
(
    id   INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name TEXT    NOT NULL
        CONSTRAINT u_name
            UNIQUE
                ON CONFLICT FAIL
);

INSERT INTO connections_netscaleradc_protocols (name)
VALUES ('nitro'),
       ('ssh');

CREATE TABLE IF NOT EXISTS connections_netscaleradc
(
    id                        INTEGER NOT NULL
        CONSTRAINT f_connection_id
            REFERENCES connections (id)
            ON DELETE CASCADE,
    management_address        TEXT,
    node_addresses            TEXT,
    connection_timeout        INTEGER,
    user_agent                TEXT,
    use_ssl                   INTEGER NOT NULL,
    validateServerCertificate INTEGER NOT NULL
);

INSERT INTO connection_types (name)
VALUES ('netscaleradc');