CREATE TABLE answers (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

  attempt_id UUID REFERENCES attempts(id) ON DELETE CASCADE,
  question_id TEXT NOT NULL,

  selected_answer JSONB,
  is_correct BOOLEAN,
  marks_awarded FLOAT DEFAULT 0,

  question_snapshot JSONB,

  UNIQUE (attempt_id, question_id)
);

CREATE INDEX idx_answers_attempt ON answers(attempt_id);
CREATE INDEX idx_answers_question ON answers(question_id);