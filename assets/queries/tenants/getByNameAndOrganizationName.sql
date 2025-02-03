SELECT *
FROM tenants
         JOIN organizations ON tenants.organization_id = organizations.id
WHERE tenants.name == ?
  AND organizations.name == ?
