CREATE TABLE IF NOT EXISTS connections_smtp_protocols
(
    id   INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name TEXT    NOT NULL
        CONSTRAINT u_name
            UNIQUE
                ON CONFLICT FAIL
);

INSERT INTO connections_smtp_protocols (name)
VALUES ('plain'),
       ('ssl'),
       ('tls');

CREATE TABLE IF NOT EXISTS connections_smtp_servers
(
    id               INTEGER NOT NULL
        CONSTRAINT f_connection_id
            REFERENCES connections (id),
    address          TEXT    NOT NULL,
    port             INTEGER,
    smtp_protocol_id INTEGER NOT NULL
        CONSTRAINT f_smtp_protocol_id
            REFERENCES connections_smtp_protocols (id)
            ON DELETE RESTRICT
);

INSERT INTO connection_types (guid, name)
VALUES ('0', 'smtp');