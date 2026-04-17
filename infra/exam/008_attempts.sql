CREATE TABLE attempts (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL,
  exam_id UUID REFERENCES exams(id),

  status attempt_status NOT NULL DEFAULT 'in_progress',

  total_score FLOAT DEFAULT 0,
  time_taken_seconds INT,

  started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  completed_at TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_attempts_user_exam ON attempts(user_id, exam_id);
CREATE INDEX idx_attempts_exam ON attempts(exam_id);
CREATE INDEX idx_attempts_user ON attempts(user_id);
CREATE INDEX idx_attempts_status ON attempts(status);

CREATE UNIQUE INDEX uniq_active_attempt ON attempts(user_id, exam_id) WHERE status = 'in_progress';