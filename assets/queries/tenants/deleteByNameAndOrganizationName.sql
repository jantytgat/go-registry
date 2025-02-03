DELETE
FROM tenants
WHERE name == ?
  AND organization_id == (SELECT id
                          FROM organizations
                          WHERE name == ?)