CREATE TABLE attempt_section_scores (
  attempt_id UUID,
  section_id UUID,

  score FLOAT,
  correct INT,
  wrong INT,

  PRIMARY KEY (attempt_id, section_id),

  FOREIGN KEY (attempt_id) REFERENCES attempts(id) ON DELETE CASCADE,
  FOREIGN KEY (section_id) REFERENCES exam_sections(id)
);

CREATE INDEX idx_section_scores_attempt ON attempt_section_scores(attempt_id);