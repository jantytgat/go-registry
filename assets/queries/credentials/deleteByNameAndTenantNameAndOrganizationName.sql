DELETE
FROM environments
WHERE name == ?
  AND tenant_id == (SELECT tenants.id
                    FROM tenants
                             JOIN organizations ON tenants.organization_id = organizations.id
                    WHERE tenants.name == ?
                      AND organizations.name == ?)