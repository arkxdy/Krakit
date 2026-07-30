import type {
  BaseQuestion,
  Exam,
  ExamSection,
  ExamsResponse,
  GetExamsParams,
  Subject,
} from "@/types/exam.types";
import { HTTP_CLIENT_URL, PAGE_SIZE } from "@/utils/constant";

const examQuestionData: BaseQuestion[] = [
  {
    id: "q1",
    exam_id: "e2c9e0b1-6a4a-4b9f-9c3d-fb1a77c1e980",
    section_id: "s1",
    subject_id: "sub1",
    question_set_id: "qs1",
    passage_id: "p1",
    question_text: "What is the capital of France?",
    question_image_url: "",
    options: [
      { key: "A", option_text: "Paris", option_image_url: "" },
      { key: "B", option_text: "London", option_image_url: "" },
      { key: "C", option_text: "Berlin", option_image_url: "" },
      { key: "D", option_text: "Madrid", option_image_url: "" },
    ],
    correct_answer: "A",
    difficulty: "easy",
    type: "MCQ",
    negative_marks: 0.25,
    topic: "Geography",
    created_at: "2025-10-25T12:00:00Z",
    passage: {
      id: "p1",
      exam_id: "",
      section_id: "",
      subject_id: "",
      topic: "hehe",
      passage_text:
        "Landing in Australia, the British colonists weren't much impressed with the small-bodied, slender-snooted marsupials called bandicoots. \"Their muzzle, which is much too long, gives them an air exceedingly stupid,\" one naturalist noted in 1805. They nicknamed one type the \"zebra rat\" because of its black-striped rump. Silly-looking or not, though, the zebra rat—the smallest bandicoot, more commonly known today as the western barred bandicoot—exhibited a genius for survival in the harsh outback, where its ancestors had persisted for some 26 million years. Its births were triggered by rainfall in the bone-dry desert. It carried its breath-mint-size babies in a backward-facing pouch so mothers could forage for food and dig shallow, camouflaged shelters. Still, these adaptations did not prepare the western barred bandicoot for the colonial-era transformation of its ecosystem, particularly the onslaught of imported British animals, from cattle and rabbits that damaged delicate desert vegetation to ravenous house cats that soon developed a taste for bandicoots. Several of the dozen-odd bandicoot species went extinct, and by the 1940s the western barred bandicoot, whose original range stretched across much of Australia, persisted only on two predator-free islands in Shark Bay, off Australia's western coast. \"Our isolated fauna had simply not been exposed to these predators,\" says Reece Pedler, an ecologist with the Wild Deserts conservation program. Now Wild Deserts is using descendants of those few thousand island survivors, called Shark Bay bandicoots, in a new effort to seed a mainland bandicoot revival. They've imported 20 bandicoots to a preserve on the edge of the Strzelecki Desert, in the remote interior of New South Wales.",
    },
  },
  {
    id: "q2",
    exam_id: "e2c9e0b1-6a4a-4b9f-9c3d-fb1a77c1e980",
    section_id: "s1",
    subject_id: "sub2",
    question_set_id: "qs1",
    passage_id: "",
    question_text: "Solve: 5 + 7 * 2 = ?",
    question_image_url: "test",
    options: [
      { key: "A", option_text: "19", option_image_url: "" },
      { key: "B", option_text: "24", option_image_url: "" },
      { key: "C", option_text: "12", option_image_url: "" },
      { key: "D", option_text: "17", option_image_url: "" },
    ],
    correct_answer: "D",
    difficulty: "medium",
    type: "MCQ",
    negative_marks: 0.5,
    topic: "Math",
    created_at: "2025-10-25T12:05:00Z",
  },
  {
    id: "q3",
    exam_id: "e2c9e0b1-6a4a-4b9f-9c3d-fb1a77c1e980",
    section_id: "s2",
    subject_id: "sub3",
    question_set_id: "qs2",
    passage_id: "",
    question_text: "Which of the following are prime numbers?",
    question_image_url: "",
    options: [
      { key: "A", option_text: "2", option_image_url: "" },
      { key: "B", option_text: "4", option_image_url: "" },
      { key: "C", option_text: "5", option_image_url: "" },
      { key: "D", option_text: "9", option_image_url: "" },
    ],
    correct_answer: ["A", "C"],
    difficulty: "hard",
    type: "MSQ",
    negative_marks: 1,
    topic: "Math",
    created_at: "2025-10-25T12:10:00Z",
  },
];

export const getExams = async (
  params?: GetExamsParams,
): Promise<ExamsResponse> => {
  const page = params?.page || 1;
  const limit = PAGE_SIZE;

  const response = await fetch(`${HTTP_CLIENT_URL}/api/exam/`);
  const data = await response.json();
  const result: ExamsResponse = {
    exams: data.data,
    total: data.data.length,
    totalPages: Math.ceil(data.data.length / PAGE_SIZE),
  };

  return result;
};

export const getExamSections = async (
  examId: string,
): Promise<ExamSection[]> => {
  const response = await fetch(
    `${HTTP_CLIENT_URL}/api/exam/sections/${examId}`,
  );
  const data = await response.json();
  return data.data;
};

export const getExamQuestions = async (
  examId: string,
): Promise<BaseQuestion[]> => {
  await new Promise((res) => setTimeout(res, 10000));
  return examQuestionData.filter((s: any) => s.exam_id === examId);
};
