import React from "react";
import type { BaseQuestion } from "@/types/exam.types";
import { MCQQuestionView } from "./MCQ";
import { MSQQuestionView } from "./MSQ";
import { NumericQuestionView } from "./NumericQuestion";

interface QuestionRendererProps {
  question: BaseQuestion;
  selectedAnswers: string[];
  onSelect: (answers: string[]) => void;
}

export const QuestionRenderer: React.FC<QuestionRendererProps> = ({
  question,
  selectedAnswers,
  onSelect,
}) => {
  switch (question.type) {
    case "MCQ":
      return (
        <MCQQuestionView
          question={question as any}
          selected={selectedAnswers}
          onSelect={onSelect}
        />
      );
    case "MSQ":
      return (
        <MSQQuestionView
          question={question as any}
          selected={selectedAnswers}
          onSelect={onSelect}
        />
      );
    case "Numeric":
      return (
        <NumericQuestionView
          question={question as any}
          selected={selectedAnswers}
          onSelect={onSelect}
        />
      );
    default:
      return <p className="text-gray-500">Unsupported question type</p>;
  }
};
