import React from "react";
import type { MSQQuestion, QuestionOption } from "@/types/exam.types";

interface MSQQuestionViewProps {
  question: MSQQuestion;
  selected: string[];
  onSelect: (selected: string[]) => void;
}

export const MSQQuestionView: React.FC<MSQQuestionViewProps> = ({
  question,
  selected,
  onSelect,
}) => {
  const toggleOption = (key: string) => {
    if (selected.includes(key))
      onSelect(selected.filter((k) => k !== key));
    else onSelect([...selected, key]);
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
              type="checkbox"
              checked={isSelected}
              onChange={() => toggleOption(opt.key)}
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
