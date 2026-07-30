import { useQuery } from "@tanstack/react-query";
import { getExams } from "@/api/exam";
import type { GetExamsParams } from "@/types/exam.types";

export const useExams = (params?: GetExamsParams) => {
  return useQuery({
    queryKey: ["exams", params], // Include params in query key for proper caching
    queryFn: () => getExams(params),
    staleTime: 5 * 60 * 1000, // 5 minutes cache
  });
};
