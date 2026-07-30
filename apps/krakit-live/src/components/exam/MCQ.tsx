import React from "react";
import type { MCQQuestion } from "@/types/exam.types";
import type { QuestionOption } from "@/types/exam.types";

interface MCQQuestionViewProps {
  question: MCQQuestion;
  selected: string[];
  onSelect: (selected: string[]) => void;
}

export const MCQQuestionView: React.FC<MCQQuestionViewProps> = ({
  question,
  selected,
  onSelect,
}) => {
  const handleSelect = (key: string) => {
    onSelect([key]);
  };

  return (
    <div className="space-y-3">
      {question.options?.map((opt: QuestionOption) => {
        const isSelected = selected.includes(opt.key);
        return (
          <label
            key={opt.key}
            className={`flex flex-col sm:flex-row items-start sm:items-center gap-3 p-3 border rounded-md cursor-pointer hover:bg-gray-50 transition-all ${
              isSelected ? "border-blue-500 bg-blue-50" : "border-gray-200"
            }`}
          >
            <input
              type="radio"
              name={question.id}
              checked={isSelected}
              onChange={() => handleSelect(opt.key)}
              className="mt-1 sm:mt-0"
            />

            <div className="flex flex-col gap-2 w-full">
              {opt.option_text && (
                <span className="text-gray-800 text-sm sm:text-base">
                  {opt.option_text}
                </span>
              )}

              {opt.option_image_url && (
                <img
                  src={opt.option_image_url}
                  alt={opt.option_text || "option image"}
                  className="max-h-40 object-contain rounded-md border"
                />
              )}
            </div>
          </label>
        );
      })}
    </div>
  );
};
