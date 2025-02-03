SELECT *
FROM environments
WHERE tenant_id == (SELECT tenants.id
                    FROM tenants
                             JOIN organizations ON tenants.organization_id = organizations.id
                    WHERE tenants.name == ?
                      AND organizations.name == ?)