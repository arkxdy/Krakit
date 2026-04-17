INSERT INTO permissions (name, description) VALUES
('user.read', 'Read users'),
('user.write', 'Create or update users'),
('exam.create', 'Create exams'),
('exam.publish', 'Publish exams'),
('exam.review', 'Review exams')
ON CONFLICT (name) DO NOTHING;