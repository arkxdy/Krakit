db = db.getSiblingDB("krakit-question");

if (!db.getCollectionNames().includes("questions")) {
  db.createCollection("questions", {
    validator: {
      $jsonSchema: {
        bsonType: "object",
        required: [
          "exam_id",
          "section_id",
          "question_set_id",
          "question_text",
          "type",
          "difficulty"
        ],
        properties: {
          exam_id: { bsonType: "string" },
          section_id: { bsonType: "string" },
          subject_id: { bsonType: "string" },

          question_set_id: { bsonType: "string" },
          passage_id: { bsonType: "objectId" },

          question_text: { bsonType: "string" },
          question_image_url: { bsonType: "string" },

          options: {
            bsonType: "array",
            items: {
              bsonType: "object",
              required: ["key"],
              properties: {
                key: { bsonType: "string" },
                option_text: { bsonType: "string" },
                option_image_url: { bsonType: "string" }
              }
            }
          },

          correct_answer: {},
          type: { enum: ["MCQ", "MSQ", "NUMERIC"] },

          marks: { bsonType: "number" },
          negative_marks: { bsonType: "number" },

          difficulty: { enum: ["easy", "medium", "hard"] },

          topic: { bsonType: "string" },
          is_active: { bsonType: "bool" },

          created_at: { bsonType: "date" }
        }
      }
    }
  });
}

// Indexes
db.questions.createIndex({ question_set_id: 1 });
db.questions.createIndex({ exam_id: 1, section_id: 1 });
db.questions.createIndex({ passage_id: 1 });
db.questions.createIndex({ difficulty: 1 });
db.questions.createIndex({ topic: 1 });

print("✅ questions ready");