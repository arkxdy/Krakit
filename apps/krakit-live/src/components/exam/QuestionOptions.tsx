import type { BaseQuestion } from "@/types/exam.types";

interface QuestionOptionsProps {
  question: BaseQuestion;
  selectedAnswers: string[];
  onAnswerSelect: (selected: string[]) => void;
}

export function QuestionOptions({
  question,
  selectedAnswers,
  onAnswerSelect,
}: QuestionOptionsProps) {
  const handleSingleSelect = (optionId: string) => {
    onAnswerSelect([optionId]);
  };

  const handleMultiSelect = (optionId: string) => {
    const newSelected = selectedAnswers.includes(optionId)
      ? selectedAnswers.filter((id) => id !== optionId)
      : [...selectedAnswers, optionId];
    onAnswerSelect(newSelected);
  };

  return (
    <div className="space-y-3">
      {question.options?.map((option) => (
        <div key={option.key} className="flex items-center">
          {question.type == 'MCQ' ? (
            <input
              type="radio"
              id={option.key}
              name={`question-${question.id}`}
              className="h-4 w-4 text-blue-600 focus:ring-blue-500"
              checked={selectedAnswers.includes(option.key)}
              onChange={() => handleSingleSelect(option.key)}
            />
          ) : (
            <input
              type="checkbox"
              id={option.key}
              name={`question-${question.id}`}
              className="h-4 w-4 text-blue-600 focus:ring-blue-500 rounded"
              checked={selectedAnswers.includes(option.key)}
              onChange={() => handleMultiSelect(option.key)}
            />
          )}
          <label htmlFor={option.key} className="ml-3 block text-gray-700">
            {option.option_text}
          </label>
        </div>
      ))}
    </div>
  );
}
