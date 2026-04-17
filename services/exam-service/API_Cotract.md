# 📘 Exam Platform API Contract

## Base URL

```
/api/v1
```

## Auth

* Bearer Token (JWT)

```
Authorization: Bearer <token>
```

---

# 👨‍💼 ADMIN APIs

---

## 🧾 Subjects

### Create Subject

```
POST /admin/subjects
```

**Body**

```json
{
  "name": "Quantitative Aptitude",
  "description": "Math section"
}
```

**Response**

```json
{
  "id": "uuid",
  "name": "Quantitative Aptitude",
  "description": "Math section"
}
```

---

### List Subjects

```
GET /admin/subjects
```

---

## 📝 Exams

### Create Exam

```
POST /admin/exams
```

**Body**

```json
{
  "name": "CAT 2026 Mock 1",
  "exam_type": "cat",
  "duration_minutes": 120,
  "total_marks": 300
}
```

---

### Update Exam

```
PUT /admin/exams/{exam_id}
```

---

### List Exams

```
GET /admin/exams
```

---

### Disable Exam

```
DELETE /admin/exams/{exam_id}
```

---

## ⚙️ Exam Settings

### Create / Update Settings

```
POST /admin/exams/{exam_id}/settings
```

```json
{
  "shuffle_questions": true,
  "shuffle_options": true,
  "allow_section_switch": false,
  "is_published": false
}
```

---

### Publish Exam

```
POST /admin/exams/{exam_id}/publish
```

---

## 📚 Sections

### Create Section

```
POST /admin/exams/{exam_id}/sections
```

```json
{
  "name": "VARC",
  "subject_id": "uuid",
  "time_limit": 40,
  "question_count": 24,
  "order_index": 1,
  "is_switch_allowed": false
}
```

---

### Update Section

```
PUT /admin/sections/{section_id}
```

---

### List Sections

```
GET /admin/exams/{exam_id}/sections
```

---

## 🧩 Question Sets (Mongo linked)

### Create Question Set Mapping

```
POST /admin/sections/{section_id}/question-sets
```

```json
{
  "question_set_id": "set_123",
  "question_count": 20,
  "difficulty_distribution": {
    "easy": 5,
    "medium": 10,
    "hard": 5
  }
}
```

---

### List Question Sets

```
GET /admin/exams/{exam_id}/question-sets
```

---

## 📊 Analytics

### Attempts for Exam

```
GET /admin/exams/{exam_id}/attempts
```

---

### Average Score

```
GET /admin/exams/{exam_id}/analytics
```

**Response**

```json
{
  "avg_score": 142.5,
  "total_attempts": 1200
}
```

---

# 👨‍🎓 USER APIs

---

## 📋 Exams

### List Available Exams

```
GET /exams
```

---

### Get Exam Details

```
GET /exams/{exam_id}
```

---

## 🚀 Attempt Flow

---

### Start Attempt

```
POST /attempts
```

```json
{
  "exam_id": "uuid"
}
```

**Response**

```json
{
  "attempt_id": "uuid",
  "status": "in_progress"
}
```

---

### Resume Attempt

```
GET /attempts/{exam_id}/active
```

---

## 📦 Fetch Questions

👉 Backend will:

* fetch structure (Postgres)
* fetch questions (Mongo)
* shuffle
* store mapping

---

### Get Attempt Questions

```
GET /attempts/{attempt_id}/questions
```

**Response**

```json
{
  "sections": [
    {
      "section_id": "uuid",
      "name": "VARC",
      "questions": [
        {
          "question_id": "q1",
          "question_text": "...",
          "options": [
            { "key": "A", "text": "..." }
          ]
        }
      ]
    }
  ]
}
```

---

### Get Section Questions

```
GET /attempts/{attempt_id}/sections/{section_id}
```

---

## ✍️ Answering

---

### Save Answer (Auto-save)

```
POST /attempts/{attempt_id}/answers
```

```json
{
  "question_id": "q1",
  "selected_answer": ["A"]
}
```

---

### Get All Answers

```
GET /attempts/{attempt_id}/answers
```

---

## ✅ Submit Attempt

```
POST /attempts/{attempt_id}/submit
```

**Response**

```json
{
  "score": 156,
  "correct": 42,
  "wrong": 10
}
```

---

## 📈 Results

---

### Get Result

```
GET /attempts/{attempt_id}/result
```

```json
{
  "total_score": 156,
  "sections": [
    {
      "section_id": "uuid",
      "score": 50,
      "correct": 15,
      "wrong": 5
    }
  ]
}
```

---

## 📜 History

---

### User Attempts

```
GET /users/me/attempts
```

---

# 🧠 Mongo APIs (Questions & Passages)

---

## ➕ Create Question

```
POST /admin/questions
```

```json
{
  "exam_id": "uuid",
  "section_id": "uuid",
  "question_set_id": "set_123",
  "question_text": "...",
  "options": [...],
  "correct_answer": "A",
  "difficulty": "medium",
  "type": "MCQ"
}
```

---

## 📥 Fetch Questions by Set

```
GET /questions?question_set_id=set_123
```

---

## 📚 Passages

### Create Passage

```
POST /admin/passages
```

---

### Get Passage

```
GET /passages/{id}
```

---

# 🔐 Error Format

```json
{
  "error": {
    "code": "INVALID_REQUEST",
    "message": "Something went wrong"
  }
}
```

---

# ⚡ Notes for Frontend

* All answers are **auto-saved**
* Timer must be handled **on FE + BE validation**
* Questions order is **fixed per attempt**
* Do not cache questions globally
* Resume flow must always check active attempt

---

# 🚀 Flow Summary

### User

```
List Exams → Start Attempt → Fetch Questions →
Auto-save Answers → Submit → View Result
```

### Admin

```
Create Subject → Create Exam → Add Sections →
Map Question Sets → Publish Exam
```
