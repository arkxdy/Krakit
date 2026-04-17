CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,

  first_name VARCHAR(255),
  last_name VARCHAR(255),
  full_name VARCHAR(255),

  role role_type NOT NULL DEFAULT 'candidate',
  plan user_plan_type NOT NULL DEFAULT 'free',

  is_active BOOLEAN NOT NULL DEFAULT TRUE,

  last_login_at TIMESTAMPTZ,

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_active ON users(is_active);