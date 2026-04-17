DO $$ BEGIN
  CREATE TYPE exam_type AS ENUM ('cat','gate','gmat','gre','upsc');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN 
  CREATE TYPE attempt_status AS ENUM ('in_progress','submitted','failed','cancelled'); 
EXCEPTION WHEN duplicate_object THEN NULL; END $$;