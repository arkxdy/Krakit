CREATE TABLE exam_question_sets (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  exam_id UUID REFERENCES exams(id) ON DELETE CASCADE,
  section_id UUID REFERENCES exam_sections(id),

  question_set_id TEXT NOT NULL,
  question_count INT NOT NULL,
  difficulty_distribution JSONB
);

CREATE INDEX idx_question_sets_exam ON exam_question_sets(exam_id);
CREATE INDEX idx_question_sets_section ON exam_question_sets(section_id);