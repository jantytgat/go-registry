CREATE TABLE IF NOT EXISTS certificates_management_types
(
    id   INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name TEXT    NOT NULL
        CONSTRAINT u_name
            UNIQUE
                ON CONFLICT FAIL
);

INSERT INTO certificates_management_types (name)
VALUES ('manual');


CREATE TABLE IF NOT EXISTS certificates_types
(
    id   INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name TEXT    NOT NULL
);

INSERT INTO certificates_types (name)
VALUES ('root ca'),
       ('intermediate ca'),
       ('server certificate');

CREATE TABLE IF NOT EXISTS certificates
(
    id               INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name             TEXT    NOT NULL,
    certificate_type_id INTEGER NOT NULL
        CONSTRAINT f_certificate_type_id
            REFERENCES certificates_types (id)
            ON DELETE CASCADE,
    management_type_id  INTEGER NOT NULL
        CONSTRAINT f_certificate_management_type_id
            REFERENCES certificates_management_types
            ON DELETE CASCADE,
    tenant_id        INTEGER NOT NULL
        CONSTRAINT f_tenant_id
            REFERENCES tenants (id)
            ON DELETE CASCADE,

    CONSTRAINT u_tenant_certificate_name
        UNIQUE (tenant_id, name)
            ON CONFLICT FAIL
);

CREATE TABLE IF NOT EXISTS certificates_tree
(
    certificate_id        INTEGER NOT NULL
        CONSTRAINT p_certificate_id
            REFERENCES certificates,
    parent_certificate_id INTEGER NOT NULL
        CONSTRAINT p_parent_certificate_id
            REFERENCES certificates,

    CONSTRAINT u_parent_child_certificate
        UNIQUE (certificate_id, parent_certificate_id)
            ON CONFLICT FAIL
);