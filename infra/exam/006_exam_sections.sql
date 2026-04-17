CREATE TABLE exam_sections (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  exam_id UUID REFERENCES exams(id) ON DELETE CASCADE,

  name TEXT NOT NULL,
  subject_id UUID REFERENCES subjects(id),

  time_limit INT,
  question_count INT,
  order_index INT,

  is_switch_allowed BOOLEAN DEFAULT TRUE,
  is_active BOOLEAN DEFAULT TRUE, 
  
  CONSTRAINT unique_section_order UNIQUE (exam_id, order_index)
);

CREATE INDEX idx_sections_exam_id ON exam_sections(exam_id);
CREATE INDEX idx_sections_active ON exam_sections(is_active);