import { useQuery } from "@tanstack/react-query";
import { getExamSections } from "@/api/exam";

export const useExamSections = (examId: string) => {
  return useQuery({
    queryKey: ["exam_sections", examId], // 👈 include examId for proper caching
    queryFn: () => getExamSections(examId), // 👈 call with parameter
    enabled: !!examId, // 👈 only run when examId is defined
  });
};
