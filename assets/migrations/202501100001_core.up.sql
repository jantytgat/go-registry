PRAGMA foreign_keys = on;

CREATE TABLE IF NOT EXISTS organizations
(
    id   INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    guid TEXT    NOT NULL
        CONSTRAINT u_guid
            UNIQUE
                ON CONFLICT FAIL,
    name TEXT    NOT NULL
        CONSTRAINT u_name
            UNIQUE
                ON CONFLICT FAIL
);

INSERT INTO organizations (guid, name)
VALUES ('0', 'default');

CREATE TABLE IF NOT EXISTS tenants
(
    id              INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    guid            TEXT    NOT NULL
        CONSTRAINT u_guid
            UNIQUE
                ON CONFLICT FAIL,
    name            TEXT    NOT NULL,
    organization_id INTEGER NOT NULL
        CONSTRAINT f_organization_id
            REFERENCES organizations (id)
            ON DELETE CASCADE,

    CONSTRAINT u_tenant_name
        UNIQUE (organization_id, name)
            ON CONFLICT FAIL
);

INSERT INTO tenants (guid, name, organization_id)
VALUES ('0', 'default', 1);

CREATE TABLE IF NOT EXISTS environments
(
    id        INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    guid      TEXT    NOT NULL
        CONSTRAINT u_guid
            UNIQUE
                ON CONFLICT FAIL,
    name      TEXT    NOT NULL,
    tenant_id INTEGER NOT NULL
        CONSTRAINT f_tenant_id
            REFERENCES tenants (id)
            ON DELETE CASCADE,

    CONSTRAINT u_environment_name
        UNIQUE (tenant_id, name)
            ON CONFLICT FAIL
);

INSERT INTO environments (guid, name, tenant_id)
VALUES ('0', 'default', '1');

CREATE VIEW tenant_environments
AS
SELECT t.id   as tenant_id,
       e.id   as environment_id,
       t.name as tenant_name,
       e.name as environment_name,
       t.guid as tenant_guid,
       e.guid as environment_guid
FROM environments e
         JOIN tenants t ON t.id = e.tenant_id;

CREATE TABLE IF NOT EXISTS credentials
(
    id        INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    guid      TEXT    NOT NULL
        CONSTRAINT u_guid
            UNIQUE
                ON CONFLICT FAIL,
    name      TEXT    NOT NULL,
    tenant_id INTEGER NOT NULL
        CONSTRAINT f_tenant_id
            REFERENCES tenants (id)
            ON DELETE CASCADE,

    CONSTRAINT u_credential_name
        UNIQUE (tenant_id, name)
            ON CONFLICT FAIL
);

CREATE TABLE IF NOT EXISTS credential_fields
(
    id            INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name          TEXT    NOT NULL,
    value         TEXT    NOT NULL,
    credential_id INTEGER NOT NULL
        CONSTRAINT f_credential_id
            REFERENCES credentials (id)
            ON DELETE CASCADE,

    CONSTRAINT u_credential_field_name
        UNIQUE (credential_id, name)
            ON CONFLICT FAIL
);

CREATE TABLE IF NOT EXISTS connection_types
(
    id   INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    guid TEXT    NOT NULL,
    name TEXT    NOT NULL
        CONSTRAINT u_name
            UNIQUE
                ON CONFLICT FAIL
);

CREATE TABLE IF NOT EXISTS connections
(
    id                 INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    guid               TEXT    NOT NULL,
    name               TEXT    NOT NULL,
    environment_id     INTEGER NOT NULL
        CONSTRAINT f_environment_id
            REFERENCES environments (id),
    connection_type_id INTEGER NOT NULL
        CONSTRAINT f_connection_type_id
            REFERENCES connection_types (id)
            ON DELETE CASCADE,

    CONSTRAINT u_connection_name
        UNIQUE (environment_id, name)
            ON CONFLICT FAIL
);

CREATE TABLE IF NOT EXISTS connection_credentials
(
    id            INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name          TEXT    NOT NULL,
    connection_id INTEGER NOT NULL
        CONSTRAINT f_connections_id
            REFERENCES connections (id)
            ON DELETE CASCADE,

    credential_id INTEGER NOT NULL
        CONSTRAINT f_credential_id
            REFERENCES credentials (id)
            ON DELETE CASCADE,

    CONSTRAINT u_connection_credential
        UNIQUE (connection_id, credential_id)
            ON CONFLICT FAIL
);