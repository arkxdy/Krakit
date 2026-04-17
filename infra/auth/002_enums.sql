CREATE TYPE role_type AS ENUM (
  'super_admin',
  'admin',
  'exam_creator',
  'reviewer',
  'proctor',
  'support',
  'candidate',
  'viewer'
);

CREATE TYPE user_plan_type AS ENUM (
  'free',
  'premium',
  'subscription'
);

CREATE TYPE platform_type AS ENUM (
  'web',
  'mobile',
  'ios'
);