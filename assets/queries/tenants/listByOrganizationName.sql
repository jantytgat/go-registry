SELECT tenants.*
FROM tenants
         JOIN organizations ON tenants.organization_id = organizations.id
WHERE organizations.name == ?