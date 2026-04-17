# 📘 OpenAPI 3.0 Specification (Exam System)

```yaml
openapi: 3.0.3
info:
  title: Exam Platform API
  version: 1.0.0
  description: API for Exam, Attempt, Questions, and Admin management

servers:
  - url: https://api.example.com

tags:
  - name: Admin
  - name: User
  - name: Internal

paths:

  # ============================
  # ADMIN APIs
  # ============================

  /v1/admin/exams:
    post:
      tags: [Admin]
      summary: Create exam
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [name, exam_type, duration_minutes, total_marks]
              properties:
                name:
                  type: string
                exam_type:
                  type: string
                  enum: [cat, gate, gmat, gre, upsc]
                duration_minutes:
                  type: integer
                total_marks:
                  type: integer
      responses:
        "200":
          description: Exam created

  /v1/admin/exams/{exam_id}/sections:
    post:
      tags: [Admin]
      summary: Add section
      parameters:
        - name: exam_id
          in: path
          required: true
          schema:
            type: string
      requestBody:
        content:
          application/json:
            schema:
              type: object
              required: [name, order_index]
              properties:
                name:
                  type: string
                subject_id:
                  type: string
                time_limit:
                  type: integer
                question_count:
                  type: integer
                order_index:
                  type: integer
      responses:
        "200":
          description: Section added

  /v1/admin/exams/{exam_id}/question-sets:
    post:
      tags: [Admin]
      summary: Attach question set
      parameters:
        - name: exam_id
          in: path
          required: true
          schema:
            type: string
      requestBody:
        content:
          application/json:
            schema:
              type: object
              required: [section_id, question_set_id, question_count]
              properties:
                section_id:
                  type: string
                question_set_id:
                  type: string
                question_count:
                  type: integer
      responses:
        "200":
          description: Mapping created

  /v1/admin/passages:
    post:
      tags: [Admin]
      summary: Create passage
      requestBody:
        content:
          application/json:
            schema:
              type: object
              required: [exam_id, section_id, passage_text]
              properties:
                exam_id:
                  type: string
                section_id:
                  type: string
                subject_id:
                  type: string
                passage_text:
                  type: string
      responses:
        "200":
          description: Passage created

  /v1/admin/questions:
    post:
      tags: [Admin]
      summary: Create question
      requestBody:
        content:
          application/json:
            schema:
              type: object
              required:
                [exam_id, section_id, question_set_id, question_text, type]
              properties:
                exam_id:
                  type: string
                section_id:
                  type: string
                subject_id:
                  type: string
                question_set_id:
                  type: string
                passage_id:
                  type: string
                question_text:
                  type: string
                options:
                  type: array
                  items:
                    type: object
                correct_answer: {}
                type:
                  type: string
                  enum: [MCQ, MSQ, NUMERIC]
                marks:
                  type: number
                negative_marks:
                  type: number
      responses:
        "200":
          description: Question created

  /v1/admin/exams/{exam_id}/publish:
    post:
      tags: [Admin]
      summary: Publish exam
      parameters:
        - name: exam_id
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Exam published

  # ============================
  # USER APIs
  # ============================

  /v1/exams/{exam_id}:
    get:
      tags: [User]
      summary: Get exam
      parameters:
        - name: exam_id
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Exam details

  /v1/attempts:
    post:
      tags: [User]
      summary: Start attempt
      requestBody:
        content:
          application/json:
            schema:
              type: object
              required: [exam_id]
              properties:
                exam_id:
                  type: string
      responses:
        "200":
          description: Attempt created

  /v1/attempts/{attempt_id}/questions:
    get:
      tags: [User]
      summary: Get attempt questions
      parameters:
        - name: attempt_id
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Questions fetched

  /v1/attempts/{attempt_id}/answers:
    put:
      tags: [User]
      summary: Save answer
      parameters:
        - name: attempt_id
          in: path
          required: true
          schema:
            type: string
      requestBody:
        content:
          application/json:
            schema:
              type: object
              required: [question_id, selected_answer]
              properties:
                question_id:
                  type: string
                selected_answer:
                  type: array
                  items:
                    type: string
      responses:
        "200":
          description: Answer saved

  /v1/attempts/{attempt_id}/submit:
    post:
      tags: [User]
      summary: Submit attempt
      parameters:
        - name: attempt_id
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Attempt submitted

  # ============================
  # INTERNAL (MONGO FETCH)
  # ============================

  /v1/internal/questions:
    get:
      tags: [Internal]
      summary: Fetch questions
      parameters:
        - name: question_set_id
          in: query
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Questions list

  /v1/internal/questions-with-passages:
    post:
      tags: [Internal]
      summary: Fetch questions with passages
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                question_set_id:
                  type: string
      responses:
        "200":
          description: Questions + passages

components:
  schemas:
    Error:
      type: object
      properties:
        message:
          type: string
```

---

# ✅ Notes

* Use **JWT auth middleware** for Admin/User separation
* Add **rate limiting** for `/attempts`
* Add **idempotency** for submit API
* Internal APIs should be **private (no public exposure)**

---
