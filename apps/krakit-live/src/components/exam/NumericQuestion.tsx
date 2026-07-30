import React from "react";
import type { NumericQuestion } from "@/types/exam.types";

interface NumericQuestionViewProps {
  question: NumericQuestion;
  selected: string[];
  onSelect: (selected: string[]) => void;
}

export const NumericQuestionView: React.FC<NumericQuestionViewProps> = ({
  question,
  selected,
  onSelect,
}) => {
  return (
    <div className="flex flex-col gap-2">
      <input
        type="number"
        value={selected[0] || ""}
        onChange={(e) => onSelect([e.target.value])}
        className="border rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
        placeholder="Enter your answer"
      />
    </div>
  );
};
