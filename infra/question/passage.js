db = db.getSiblingDB("krakit-question");

if (!db.getCollectionNames().includes("passages")) {
  db.createCollection("passages", {
    validator: {
      $jsonSchema: {
        bsonType: "object",
        required: ["exam_id", "section_id", "passage_text"],
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
}

db.passages.createIndex({ exam_id: 1, section_id: 1 });

print("✅ passages ready");