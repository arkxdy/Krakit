-- ============================================================
-- SUBJECTS
-- ============================================================

-- name: CreateSubject :one
INSERT INTO subjects (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetSubjects :many
SELECT * FROM subjects;

-- name: GetSubject :one
SELECT * FROM subjects WHERE id = $1;

-- name: DeleteSubject :exec
DELETE FROM subjects WHERE id = $1;

-- ============================================================
-- EXAMS
-- ============================================================

-- name: CreateExam :one
INSERT INTO exams (name, exam_type, duration_minutes, total_marks)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateExam :exec
UPDATE exams
SET name = $2,
    duration_minutes = $3,
    total_marks = $4
WHERE id = $1;

-- name: GetExam :one
SELECT * FROM exams WHERE id = $1;

-- name: ListExams :many
SELECT *
FROM exams
WHERE is_active = TRUE
ORDER BY created_at DESC;

-- name: ListExamsPaginated :many
SELECT *
FROM exams
WHERE is_active = TRUE
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListPublishedExams :many
SELECT e.*
FROM exams e
JOIN exam_settings es ON e.id = es.exam_id
WHERE e.is_active = TRUE
  AND es.is_published = TRUE
ORDER BY e.created_at DESC;

-- name: DisableExam :exec
UPDATE exams
SET is_active = FALSE
WHERE id = $1;

-- name: ListExamsWithStatus :many
SELECT e.*, es.is_published
FROM exams e
LEFT JOIN exam_settings es ON e.id = es.exam_id
WHERE e.is_active = TRUE
ORDER BY e.created_at DESC;

-- ============================================================
-- EXAM SETTINGS
-- ============================================================

-- name: UpsertExamSettings :one
INSERT INTO exam_settings (
  exam_id,
  is_published,
  shuffle_questions,
  shuffle_options,
  allow_section_switch
)
VALUES ($1,$2,$3,$4,$5)
ON CONFLICT (exam_id)
DO UPDATE SET
  is_published = EXCLUDED.is_published,
  shuffle_questions = EXCLUDED.shuffle_questions,
  shuffle_options = EXCLUDED.shuffle_options,
  allow_section_switch = EXCLUDED.allow_section_switch
RETURNING *;

-- name: PublishExam :exec
INSERT INTO exam_settings (exam_id, is_published)
VALUES ($1, TRUE)
ON CONFLICT (exam_id)
DO UPDATE SET is_published = TRUE;

-- name: GetExamSettings :one
SELECT * FROM exam_settings WHERE exam_id = $1;

-- ============================================================
-- SECTIONS
-- ============================================================

-- name: CreateSection :one
INSERT INTO exam_sections (
  exam_id,
  name,
  subject_id,
  time_limit,
  question_count,
  order_index,
  is_switch_allowed
)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: UpdateSection :exec
UPDATE exam_sections
SET name = $2,
    time_limit = $3,
    question_count = $4,
    order_index = $5
WHERE id = $1;

-- name: GetSectionsByExam :many
SELECT *
FROM exam_sections
WHERE exam_id = $1 AND is_active = TRUE
ORDER BY order_index;

-- name: GetExamWithSections :many
SELECT
  e.id AS exam_id,
  e.name AS exam_name,
  e.exam_type,
  e.duration_minutes,
  e.total_marks,

  s.id AS section_id,
  s.name AS section_name,
  s.order_index,
  s.question_count
FROM exams e
LEFT JOIN exam_sections s ON e.id = s.exam_id
WHERE e.id = $1;

-- ============================================================
-- QUESTION SETS
-- ============================================================

-- name: CreateQuestionSet :one
INSERT INTO exam_question_sets (
  exam_id,
  section_id,
  question_set_id,
  question_count,
  difficulty_distribution
)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: UpdateQuestionSet :exec
UPDATE exam_question_sets
SET question_count = $2,
    difficulty_distribution = $3
WHERE id = $1;

-- name: GetQuestionSetsByExam :many
SELECT *
FROM exam_question_sets
WHERE exam_id = $1;

-- name: CountQuestionSetsBySection :one
SELECT COUNT(*) AS total
FROM exam_question_sets
WHERE section_id = $1;

-- ============================================================
-- ATTEMPTS
-- ============================================================

-- name: CreateAttempt :one
INSERT INTO attempts (user_id, exam_id)
VALUES ($1,$2)
RETURNING *;

-- name: GetAttempt :one
SELECT * FROM attempts WHERE id = $1;

-- name: GetActiveAttempt :one
SELECT *
FROM attempts
WHERE user_id = $1
  AND exam_id = $2
  AND status = 'in_progress'
LIMIT 1;

-- name: CompleteAttempt :one
UPDATE attempts
SET status = 'submitted',
    total_score = $2,
    completed_at = NOW()
WHERE id = $1 AND status = 'in_progress'
RETURNING *;

-- name: GetLatestAttempt :one
SELECT *
FROM attempts
WHERE user_id = $1 AND exam_id = $2
ORDER BY started_at DESC
LIMIT 1;

-- name: GetAttemptsByUser :many
SELECT *
FROM attempts
WHERE user_id = $1
ORDER BY started_at DESC;

-- name: GetAttemptsByExam :many
SELECT *
FROM attempts
WHERE exam_id = $1;

-- name: GetAttemptsByExamPaginated :many
SELECT *
FROM attempts
WHERE exam_id = $1
ORDER BY started_at DESC
LIMIT $2 OFFSET $3;

-- name: GetAttemptStatus :one
SELECT status
FROM attempts
WHERE id = $1;

-- ============================================================
-- ATTEMPT QUESTION MAP
-- ============================================================

-- name: BulkInsertAttemptMap :copyfrom
INSERT INTO attempt_question_map (
  attempt_id,
  question_id,
  section_id,
  order_index,
  shuffled_options
)
VALUES ($1,$2,$3,$4,$5);

-- name: GetAttemptQuestions :many
SELECT *
FROM attempt_question_map
WHERE attempt_id = $1
ORDER BY order_index;

-- name: GetAttemptQuestionsBySection :many
SELECT *
FROM attempt_question_map
WHERE attempt_id = $1
  AND section_id = $2
ORDER BY order_index;

-- ============================================================
-- ANSWERS
-- ============================================================

-- name: UpsertAnswer :exec
INSERT INTO answers (
  attempt_id,
  question_id,
  selected_answer,
  question_snapshot
)
VALUES ($1,$2,$3,$4)
ON CONFLICT (attempt_id,question_id)
DO UPDATE SET
  selected_answer = EXCLUDED.selected_answer,
  question_snapshot = EXCLUDED.question_snapshot;

-- name: GetAnswersByAttempt :many
SELECT *
FROM answers
WHERE attempt_id = $1;

-- name: GetAnswerByAttemptAndQuestion :one
SELECT *
FROM answers
WHERE attempt_id = $1 AND question_id = $2;

-- name: EvaluateAnswer :exec
UPDATE answers
SET is_correct = $3,
    marks_awarded = $4
WHERE attempt_id = $1 AND question_id = $2;

-- ============================================================
-- SECTION SCORES
-- ============================================================

-- name: UpsertSectionScore :exec
INSERT INTO attempt_section_scores (
  attempt_id,
  section_id,
  score,
  correct,
  wrong
)
VALUES ($1,$2,$3,$4,$5)
ON CONFLICT (attempt_id,section_id)
DO UPDATE SET
  score = EXCLUDED.score,
  correct = EXCLUDED.correct,
  wrong = EXCLUDED.wrong;

-- name: GetSectionScores :many
SELECT *
FROM attempt_section_scores
WHERE attempt_id = $1;

-- ============================================================
-- RESULTS & ANALYTICS
-- ============================================================

-- name: GetAttemptResult :one
SELECT id, total_score, time_taken_seconds, completed_at
FROM attempts
WHERE id = $1;

-- name: GetAverageScoreByExam :one
SELECT COALESCE(AVG(total_score), 0) AS avg_score
FROM attempts
WHERE exam_id = $1
  AND status = 'submitted';

-- name: CountAttemptsByExam :one
SELECT COUNT(*) AS total_attempts
FROM attempts
WHERE exam_id = $1;