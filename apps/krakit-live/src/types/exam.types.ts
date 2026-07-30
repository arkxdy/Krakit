export interface GetExamsParams {
  page?: number;
}

export interface ExamsResponse {
  exams: Exam[];
  total: number;
  totalPages: number;
}

export type ExamType = "cat" | "gate" | "gmat" | "gre" | "upsc";

export interface Exam {
  id: string;
  name: string;
  description: string;
  exam_type: ExamType;
  duration_minutes: number;
  total_marks: number;
  created_at: string;
  updated_at: string;
}

export interface ExamSection {
  id: string;
  exam_id: string;
  subject_id: string;
  question_set_id: string;
  weightage: number;
  created_at: string;
  subject?: Subject;
}

export interface Subject {
  id: string;
  name: string;
  description?: string;
  num_questions: number;
  time_limit: number; // in minutes
  marking_type: string;
  positive_marks: number;
  negative_marks: number;
  created_at: string;
  updated_at: string;
}

export interface QuestionOption {
  key: string;
  option_text?: string;
  option_image_url?: string;
}

export type QuestionType = "MCQ" | "MSQ" | "Numeric";

export interface BaseQuestion {
  id: string;
  exam_id: string;
  section_id: string;
  subject_id?: string;
  question_set_id: string;
  passage_id?: string;
  question_text: string;
  question_image_url?: string;
  options?: QuestionOption[];
  correct_answer: string | string[] | number;
  difficulty?: "easy" | "medium" | "hard";
  type?: QuestionType;
  negative_marks?: number;
  topic?: string;
  created_at?: string;
  passage?: Passage
}

export interface MCQQuestion extends BaseQuestion {
  type: "MCQ";
  correctAnswer: string;
}

export interface MSQQuestion extends BaseQuestion {
  type: "MSQ";
  correctAnswer: string[];
}

export interface NumericQuestion extends BaseQuestion {
  type: "Numeric";
  correctAnswer: number;
}

interface Passage {
  id: string,
  exam_id: string,
  section_id: string,
  subject_id: string,
  passage_text: string,
  topic?: string,
  created_at?: string
}