CREATE TABLE attempt_question_map (
  attempt_id UUID,
  question_id TEXT,
  section_id UUID,
  order_index INT,
  shuffled_options JSONB,

  PRIMARY KEY (attempt_id, question_id),

  FOREIGN KEY (attempt_id) REFERENCES attempts(id) ON DELETE CASCADE,
  FOREIGN KEY (section_id) REFERENCES exam_sections(id)
);

CREATE INDEX idx_attempt_map_attempt ON attempt_question_map(attempt_id);
CREATE INDEX idx_attempt_map_section ON attempt_question_map(section_id);