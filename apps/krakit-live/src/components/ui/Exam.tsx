/* File: components/Exam.tsx */
import { useState } from "react";
import { Input } from "./Input";
import { MultiSelect } from "./MultiSelect";

interface Option {
  label: string;
  value: string;
}

interface Question {
  id: string;
  type: "text" | "single" | "multi";
  label: string;
  options?: Option[];
  maxSelections?: number;
}

interface ExamProps {
  questions: Question[];
}

export function Exam({ questions }: ExamProps) {
  const [answers, setAnswers] = useState<Record<string, string | string[]>>({});

  const handleChange = (id: string, value: string | string[]) => {
    setAnswers((prev) => ({ ...prev, [id]: value }));
  };

  return (
    <div className="space-y-6 p-4 bg-black text-white">
      {questions.map((q) => (
        <div key={q.id} className="space-y-2">
          <p className="text-purple-400 font-medium">{q.label}</p>
          {q.type === "text" && (
            <Input
              value={(answers[q.id] as string) || ""}
              onChange={(e) => handleChange(q.id, e.target.value)}
            />
          )}
          {(q.type === "single" || q.type === "multi") && q.options && (
            <MultiSelect
              options={q.options}
              value={answers[q.id] || (q.type === "multi" ? [] : "")}
              onChange={(val) => handleChange(q.id, val)}
              mode={q.type === "multi" ? "multi" : "single"}
              type={q.type === "multi" ? "checkbox" : "radio"}
              maxSelections={q.maxSelections}
            />
          )}
        </div>
      ))}
    </div>
  );
}
