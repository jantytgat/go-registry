INSERT INTO certificates_management_types (name)
VALUES ('acme');

CREATE TABLE IF NOT EXISTS certificates_acme_services
(
    id        INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name      TEXT    NOT NULL,
    url       TEXT    NOT NULL,
    tenant_id INTEGER NOT NULL
        CONSTRAINT f_tenant_id
            REFERENCES tenants (id)
            ON DELETE CASCADE,

    CONSTRAINT u_name_tenant_id
        UNIQUE (tenant_id, name)
            ON CONFLICT FAIL,

    CONSTRAINT u_url_tenant_id
        UNIQUE (tenant_id, url)
            ON CONFLICT FAIL
);

INSERT INTO certificates_acme_services (name, url, tenant_id)
SELECT 'letsencrypt_production', 'https://acme-v02.api.letsencrypt.org/directory', id
FROM tenants;

INSERT INTO certificates_acme_services (name, url, tenant_id)
SELECT 'letsencrypt_staging', 'https://acme-staging-v02.api.letsencrypt.org/directory', id
FROM tenants;

CREATE TABLE IF NOT EXISTS certificates_acme_users
(
    id        INTEGER NOT NULL
        CONSTRAINT p_id
            PRIMARY KEY AUTOINCREMENT,
    name      TEXT    NOT NULL,
    email     TEXT    NOT NULL,
    kid       TEXT,
    hmac      TEXT,
    tenant_id INTEGER NOT NULL
        CONSTRAINT f_tenant_id
            REFERENCES tenants (id)
            ON DELETE CASCADE,

    CONSTRAINT u_name_tenant_id
        UNIQUE (tenant_id, name)
            ON CONFLICT FAIL
);
