INSERT INTO role_permissions (role, permission_id)
SELECT 'admin', id
FROM permissions
WHERE name IN (
  'user.read',
  'user.write',
  'exam.create',
  'exam.publish'
)
ON CONFLICT DO NOTHING;


INSERT INTO role_permissions (role, permission_id)
SELECT 'exam_creator', id
FROM permissions
WHERE name = 'exam.create'
ON CONFLICT DO NOTHING;


INSERT INTO role_permissions (role, permission_id)
SELECT 'reviewer', id
FROM permissions
WHERE name = 'exam.review'
ON CONFLICT DO NOTHING;