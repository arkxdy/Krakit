import { useQuery } from "@tanstack/react-query";
import { getExamQuestions } from "@/api/exam";

export const useExamQuestions = (examId: string) => {
  return useQuery({
    queryKey: ["exam_questions", examId],
    queryFn: () => getExamQuestions(examId),
    enabled: !!examId,
  });
};
