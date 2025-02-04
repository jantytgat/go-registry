INSERT INTO tenants (guid, name, organization_id)
SELECT ?, ?, id
FROM organizations
WHERE organizations.name == ?
RETURNING id