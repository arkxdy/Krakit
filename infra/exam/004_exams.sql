CREATE TABLE exams (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  exam_type exam_type NOT NULL,
  duration_minutes INT NOT NULL,
  total_marks INT NOT NULL,
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_exams_created_at ON exams(created_at DESC);
CREATE INDEX idx_exams_active ON exams(is_active);