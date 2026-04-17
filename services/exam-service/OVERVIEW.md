# 🧠 Exam System – Production Ready Design (Postgres + Mongo + Kafka)

---

# 📌 Overview

This system is designed for:

* CAT / GATE / similar exams
* High concurrency attempts
* Scalable question storage
* Clean admin + user flows

### Tech Stack

* **Postgres** → structure + attempts + scoring
* **MongoDB** → questions + passages
* **Kafka** → async processing (evaluation, analytics)

---

# 🧱 1. POSTGRES SCHEMA

## Extensions & Enums

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$ BEGIN
  CREATE TYPE exam_type AS ENUM ('cat','gate','gmat','gre','upsc');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;
```

---

## Subjects

```sql
CREATE TABLE subjects (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT UNIQUE NOT NULL,
  description TEXT
);
```

---

## Exams

```sql
CREATE TABLE exams (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  exam_type exam_type NOT NULL,
  duration_minutes INT NOT NULL,
  total_marks INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Exam Settings

```sql
CREATE TABLE exam_settings (
  exam_id UUID PRIMARY KEY REFERENCES exams(id) ON DELETE CASCADE,
  is_published BOOLEAN DEFAULT FALSE,
  shuffle_questions BOOLEAN DEFAULT TRUE,
  shuffle_options BOOLEAN DEFAULT TRUE,
  allow_section_switch BOOLEAN DEFAULT TRUE
);
```

---

## Sections

```sql
CREATE TABLE exam_sections (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  exam_id UUID REFERENCES exams(id) ON DELETE CASCADE,

  name TEXT NOT NULL,
  subject_id UUID REFERENCES subjects(id),

  time_limit INT,
  question_count INT,
  order_index INT,

  is_switch_allowed BOOLEAN DEFAULT TRUE
);
```

---

## Question Set Mapping

```sql
CREATE TABLE exam_question_sets (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  exam_id UUID REFERENCES exams(id) ON DELETE CASCADE,
  section_id UUID REFERENCES exam_sections(id),

  question_set_id TEXT NOT NULL,
  question_count INT NOT NULL,
  difficulty_distribution JSONB
);
```

---

## Attempts

```sql
CREATE TABLE attempts (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL,
  exam_id UUID REFERENCES exams(id),

  status TEXT DEFAULT 'in_progress',

  total_score FLOAT DEFAULT 0,
  time_taken_seconds INT,

  started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  completed_at TIMESTAMP
);
```

---

## Attempt Question Map

```sql
CREATE TABLE attempt_question_map (
  attempt_id UUID,
  question_id TEXT,
  section_id UUID,
  order_index INT,
  shuffled_options JSONB,
  PRIMARY KEY (attempt_id, question_id)
);
```

---

## Answers

```sql
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
```

---

## Section Scores

```sql
CREATE TABLE attempt_section_scores (
  attempt_id UUID,
  section_id UUID,

  score FLOAT,
  correct INT,
  wrong INT,

  PRIMARY KEY (attempt_id, section_id)
);
```

---

# ⚡ 2. SQLC QUERIES

## attempts.sql

```sql
-- name: CreateAttempt :one
INSERT INTO attempts (user_id, exam_id)
VALUES ($1,$2) RETURNING *;

-- name: GetAttempt :one
SELECT * FROM attempts WHERE id=$1;

-- name: CompleteAttempt :exec
UPDATE attempts
SET status='submitted',
    total_score=$2,
    completed_at=NOW()
WHERE id=$1;
```

---

## sections.sql

```sql
-- name: GetSectionsByExam :many
SELECT * FROM exam_sections
WHERE exam_id=$1 ORDER BY order_index;
```

---

## question_sets.sql

```sql
-- name: GetQuestionSetsByExam :many
SELECT * FROM exam_question_sets
WHERE exam_id=$1;
```

---

## attempt_map.sql

```sql
-- name: BulkInsertAttemptMap :copyfrom
INSERT INTO attempt_question_map (
  attempt_id, question_id, section_id, order_index, shuffled_options
) VALUES ($1,$2,$3,$4,$5);

-- name: GetAttemptQuestions :many
SELECT * FROM attempt_question_map
WHERE attempt_id=$1 ORDER BY order_index;
```

---

## answers.sql

```sql
-- name: UpsertAnswer :exec
INSERT INTO answers (
  attempt_id, question_id, selected_answer, question_snapshot
)
VALUES ($1,$2,$3,$4)
ON CONFLICT (attempt_id,question_id)
DO UPDATE SET selected_answer=EXCLUDED.selected_answer;

-- name: GetAnswersByAttempt :many
SELECT * FROM answers WHERE attempt_id=$1;
```

---

## evaluation.sql

```sql
-- name: EvaluateAnswer :exec
UPDATE answers
SET is_correct=$3,
    marks_awarded=$4
WHERE attempt_id=$1 AND question_id=$2;

-- name: UpsertSectionScore :exec
INSERT INTO attempt_section_scores (
  attempt_id, section_id, score, correct, wrong
)
VALUES ($1,$2,$3,$4,$5)
ON CONFLICT (attempt_id,section_id)
DO UPDATE SET
  score=EXCLUDED.score,
  correct=EXCLUDED.correct,
  wrong=EXCLUDED.wrong;
```

---

# 🍃 3. MONGODB DESIGN

## Passages

```js
db.createCollection("passages", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["exam_id","section_id","passage_text"],
      properties: {
        exam_id: { bsonType: "string" },
        section_id: { bsonType: "string" },
        subject_id: { bsonType: "string" },
        passage_text: { bsonType: "string" },
        topic: { bsonType: "string" },
        is_active: { bsonType: "bool" },
        created_at: { bsonType: "date" }
      }
    }
  }
});
```

---

## Questions

```js
db.createCollection("questions", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["exam_id","section_id","question_set_id","question_text","type"],
      properties: {
        exam_id: { bsonType: "string" },
        section_id: { bsonType: "string" },
        subject_id: { bsonType: "string" },
        question_set_id: { bsonType: "string" },

        passage_id: { bsonType: "objectId" },

        question_text: { bsonType: "string" },
        options: { bsonType: "array" },

        correct_answer: {},
        type: { enum: ["MCQ","MSQ","NUMERIC"] },

        marks: { bsonType: "number" },
        negative_marks: { bsonType: "number" },

        difficulty: { enum: ["easy","medium","hard"] },
        is_active: { bsonType: "bool" }
      }
    }
  }
});
```

---

## Indexes

```js
db.questions.createIndex({ question_set_id: 1 });
db.questions.createIndex({ exam_id: 1, section_id: 1 });

db.passages.createIndex({ exam_id: 1, section_id: 1 });
```

---

# 🌐 4. API CONTRACTS

---

## 👨‍💼 ADMIN

### Create Exam

```
POST /v1/admin/exams
```

### Add Section

```
POST /v1/admin/exams/{id}/sections
```

### Attach Question Set

```
POST /v1/admin/exams/{id}/question-sets
```

### Create Passage

```
POST /v1/admin/passages
```

### Create Question

```
POST /v1/admin/questions
```

### Bulk Upload Questions

```
POST /v1/admin/questions/bulk
```

### Publish Exam

```
POST /v1/admin/exams/{id}/publish
```

---

## 👨‍🎓 USER

### Get Exam

```
GET /v1/exams/{id}
```

### Start Attempt

```
POST /v1/attempts
```

### Get Questions

```
GET /v1/attempts/{id}/questions
```

### Save Answer

```
PUT /v1/attempts/{id}/answers
```

### Submit Attempt

```
POST /v1/attempts/{id}/submit
```

---

## 🔍 INTERNAL (Mongo Fetch)

### Fetch Questions

```
GET /v1/internal/questions?question_set_id=...
```

### Fetch Questions + Passages

```
POST /v1/internal/questions-with-passages
```

---

# 🔄 5. FLOWS

---

## 🧑‍💼 ADMIN FLOW

```
Create Exam
   ↓
Create Sections
   ↓
Create Passages (Mongo)
   ↓
Create Questions (Mongo)
   ↓
Attach Question Sets
   ↓
Publish Exam
```

---

## 👨‍🎓 USER FLOW

```
Start Attempt
   ↓
Fetch question_set_id (Postgres)
   ↓
Fetch questions (Mongo)
   ↓
Shuffle + store mapping
   ↓
Answer Questions
   ↓
Submit Attempt
   ↓
Evaluate
   ↓
Return Result
```

---

# ⚡ 6. KAFKA INTEGRATION

---

## Use Kafka For:

### Attempt Submission

```
attempt_submitted → Kafka → evaluation worker
```

### Analytics

```
answer_saved → Kafka
```

### Bulk Upload

```
question_bulk_upload → Kafka → worker inserts Mongo
```

---

## Do NOT use Kafka for:

* fetching questions
* saving answers
* starting attempts

---

# 🚀 FINAL ARCHITECTURE

```
Backend API
   ├── Postgres (core data)
   ├── MongoDB (questions)
   └── Kafka (async processing)
```

---

# ✅ FINAL STATUS

You now have:

* Full DB design (Postgres + Mongo)
* sqlc-ready queries
* Admin + User APIs
* Internal Mongo APIs
* Kafka integration points
* End-to-end flow

---

# 🔥 Next Steps (Optional)

* Add Redis caching
* Add leaderboard service
* Add proctoring/anti-cheat
* Add difficulty-based question selection

---
