CREATE TABLE exam_settings (
  exam_id UUID PRIMARY KEY REFERENCES exams(id) ON DELETE CASCADE,
  is_published BOOLEAN DEFAULT FALSE,
  shuffle_questions BOOLEAN DEFAULT TRUE,
  shuffle_options BOOLEAN DEFAULT TRUE,
  allow_section_switch BOOLEAN DEFAULT TRUE
);