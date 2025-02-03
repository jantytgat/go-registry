SELECT *
FROM environments
         JOIN tenants on environments.tenant_id = tenants.id
WHERE environments.name == ?
  AND environments.tenant_id == (SELECT tenants.id
                                 FROM tenants
                                          JOIN organizations ON tenants.organization_id = organizations.id
                                 WHERE tenants.name == ?
                                   AND organizations.name == ?)
