# 📘 openapi.yaml

```yaml
openapi: 3.0.3
info:
  title: Exam Platform API
  version: 1.0.0
  description: APIs for Admin and User exam system

servers:
  - url: /api/v1

security:
  - bearerAuth: []

paths:

  # =========================
  # ADMIN - SUBJECTS
  # =========================
  /admin/subjects:
    post:
      summary: Create subject
      tags: [Admin]
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/CreateSubject'
      responses:
        '200':
          description: Created
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Subject'

    get:
      summary: List subjects
      tags: [Admin]
      responses:
        '200':
          description: List
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Subject'

  # =========================
  # ADMIN - EXAMS
  # =========================
  /admin/exams:
    post:
      summary: Create exam
      tags: [Admin]
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/CreateExam'
      responses:
        '200':
          description: Created

    get:
      summary: List exams
      tags: [Admin]
      responses:
        '200':
          description: List exams

  /admin/exams/{exam_id}:
    put:
      summary: Update exam
      tags: [Admin]
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
              $ref: '#/components/schemas/CreateExam'
      responses:
        '200':
          description: Updated

    delete:
      summary: Disable exam
      tags: [Admin]
      parameters:
        - name: exam_id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Disabled

  # =========================
  # SETTINGS
  # =========================
  /admin/exams/{exam_id}/settings:
    post:
      summary: Create or update settings
      tags: [Admin]
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
              $ref: '#/components/schemas/ExamSettings'
      responses:
        '200':
          description: Saved

  /admin/exams/{exam_id}/publish:
    post:
      summary: Publish exam
      tags: [Admin]
      parameters:
        - name: exam_id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Published

  # =========================
  # SECTIONS
  # =========================
  /admin/exams/{exam_id}/sections:
    post:
      summary: Create section
      tags: [Admin]
      parameters:
        - name: exam_id
          in: path
          required: true
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/CreateSection'
      responses:
        '200':
          description: Created

    get:
      summary: List sections
      tags: [Admin]
      parameters:
        - name: exam_id
          in: path
          required: true
      responses:
        '200':
          description: List

  /admin/sections/{section_id}:
    put:
      summary: Update section
      tags: [Admin]
      parameters:
        - name: section_id
          in: path
          required: true
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/CreateSection'
      responses:
        '200':
          description: Updated

  # =========================
  # QUESTION SETS
  # =========================
  /admin/sections/{section_id}/question-sets:
    post:
      summary: Create question set mapping
      tags: [Admin]
      parameters:
        - name: section_id
          in: path
          required: true
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/CreateQuestionSet'
      responses:
        '200':
          description: Created

  /admin/exams/{exam_id}/question-sets:
    get:
      summary: List question sets
      tags: [Admin]
      parameters:
        - name: exam_id
          in: path
          required: true
      responses:
        '200':
          description: List

  # =========================
  # QUESTIONS (MONGO)
  # =========================
  /admin/questions:
    post:
      summary: Create question
      tags: [Admin]
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Question'
      responses:
        '200':
          description: Created

  /questions:
    get:
      summary: Fetch questions by set
      tags: [User]
      parameters:
        - name: question_set_id
          in: query
          required: true
          schema:
            type: string
      responses:
        '200':
          description: List

  # =========================
  # PASSAGES
  # =========================
  /admin/passages:
    post:
      summary: Create passage
      tags: [Admin]
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Passage'
      responses:
        '200':
          description: Created

  # =========================
  # USER EXAMS
  # =========================
  /exams:
    get:
      summary: List available exams
      tags: [User]

  /exams/{exam_id}:
    get:
      summary: Get exam details
      tags: [User]
      parameters:
        - name: exam_id
          in: path
          required: true

  # =========================
  # ATTEMPTS
  # =========================
  /attempts:
    post:
      summary: Start attempt
      tags: [User]
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/CreateAttempt'
      responses:
        '200':
          description: Started

  /attempts/{attempt_id}/questions:
    get:
      summary: Get attempt questions
      tags: [User]
      parameters:
        - name: attempt_id
          in: path
          required: true

  /attempts/{attempt_id}/answers:
    post:
      summary: Save answer
      tags: [User]
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Answer'
    get:
      summary: Get answers
      tags: [User]

  /attempts/{attempt_id}/submit:
    post:
      summary: Submit attempt
      tags: [User]

  /attempts/{attempt_id}/result:
    get:
      summary: Get result
      tags: [User]

components:

  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer

  schemas:

    CreateSubject:
      type: object
      properties:
        name:
          type: string
        description:
          type: string

    Subject:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        description:
          type: string

    CreateExam:
      type: object
      properties:
        name:
          type: string
        exam_type:
          type: string
        duration_minutes:
          type: integer
        total_marks:
          type: integer

    ExamSettings:
      type: object
      properties:
        shuffle_questions:
          type: boolean
        shuffle_options:
          type: boolean
        allow_section_switch:
          type: boolean
        is_published:
          type: boolean

    CreateSection:
      type: object
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

    CreateQuestionSet:
      type: object
      properties:
        question_set_id:
          type: string
        question_count:
          type: integer
        difficulty_distribution:
          type: object

    CreateAttempt:
      type: object
      properties:
        exam_id:
          type: string

    Answer:
      type: object
      properties:
        question_id:
          type: string
        selected_answer:
          type: object

    Question:
      type: object
      properties:
        question_text:
          type: string
        options:
          type: array
          items:
            type: object
        correct_answer:
          type: string

    Passage:
      type: object
      properties:
        passage_text:
          type: string
```
