DELETE
FROM certificates_management_types
WHERE name == 'acme';

DROP TABLE IF EXISTS certificates_acme_services;
DROP TABLE IF EXISTS certificates_acme_users
