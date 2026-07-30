import { useState, useEffect, useMemo } from "react";
import Accordion from "@/components/ui/Accordian";
import type { BaseQuestion, Exam, ExamSection } from "@/types/exam.types";
import { CountdownTimer } from "./CountdownTimer";
import { QuestionOptions } from "./QuestionOptions";

type ExamViewProps = {
  exam: Exam;
  sections: ExamSection[];
  questions: BaseQuestion[];
};

export default function ExamView1({ exam, sections, questions }: ExamViewProps) {
  const [currentSectionIndex, setCurrentSectionIndex] = useState(0);
  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
  const [questionStatus, setQuestionStatus] = useState<
    Record<string, "not-attempted" | "answered" | "marked-review">
  >({});
  const [selectedAnswers, setSelectedAnswers] = useState<Record<string, string[]>>({});
  const [windowSize, setWindowSize] = useState({
    width: typeof window !== "undefined" ? window.innerWidth : 1200,
    height: typeof window !== "undefined" ? window.innerHeight : 800,
  });
  const [timerEnd, setTimerEnd] = useState(false);

  useEffect(() => {
    const handleResize = () => {
      setWindowSize({ width: window.innerWidth, height: window.innerHeight });
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  const isMobile = windowSize.width < 1024;

  const sectionQuestions = useMemo(() => {
    const map: Record<string, BaseQuestion[]> = {};
    for (const section of sections) {
      map[section.id] = questions.filter((q) => q.section_id === section.id);
    }
    return map;
  }, [sections, questions]);

  const currentSection = sections[currentSectionIndex];
  const currentSectionQuestions = sectionQuestions[currentSection.id] || [];
  const currentQuestion = currentSectionQuestions[currentQuestionIndex];

  const handleAnswerSelect = (questionId: string, answers: string[]) => {
    setSelectedAnswers((prev) => ({ ...prev, [questionId]: answers }));
    setQuestionStatus((prev) => ({
      ...prev,
      [questionId]: answers.length > 0 ? "answered" : "not-attempted",
    }));
  };

  const clearAnswer = (questionId: string) => {
    setSelectedAnswers((prev) => ({ ...prev, [questionId]: [] }));
    setQuestionStatus((prev) => ({ ...prev, [questionId]: "not-attempted" }));
  };

  const toggleMarkReview = (questionId: string) => {
    setQuestionStatus((prev) => ({
      ...prev,
      [questionId]:
        prev[questionId] === "marked-review" ? "answered" : "marked-review",
    }));
  };

  const goToNext = () => {
    if (currentQuestionIndex < currentSectionQuestions.length - 1) {
      setCurrentQuestionIndex(currentQuestionIndex + 1);
    } else if (currentSectionIndex < sections.length - 1) {
      setCurrentSectionIndex(currentSectionIndex + 1);
      setCurrentQuestionIndex(0);
    }
    window.scrollTo({ top: 0, behavior: "smooth" });
  };

  const goToPrevious = () => {
    if (currentQuestionIndex > 0) {
      setCurrentQuestionIndex(currentQuestionIndex - 1);
    } else if (currentSectionIndex > 0) {
      const prevSection = sections[currentSectionIndex - 1];
      const prevSectionQuestions = sectionQuestions[prevSection.id] || [];
      setCurrentSectionIndex(currentSectionIndex - 1);
      setCurrentQuestionIndex(prevSectionQuestions.length - 1);
    }
    window.scrollTo({ top: 0, behavior: "smooth" });
  };

  return (
    <div
      className={`min-h-screen bg-gray-50 ${
        isMobile ? "overflow-y-auto pb-24" : "p-2 sm:p-4 md:p-6"
      }`}
    >
      <div className="max-w-7xl mx-auto">
        <header className="mb-4 sm:mb-6 sticky top-0 bg-gray-50 z-20 py-2">
          <h1 className="text-xl sm:text-2xl font-bold text-gray-900">{exam.name}</h1>
          <p className="text-gray-600 text-sm">{exam.description}</p>
        </header>

        {/* --- Layout --- */}
        <div className="flex flex-col lg:flex-row gap-4 sm:gap-6">
          {/* LEFT PANEL */}
          <div className={`${isMobile ? "order-2" : "flex-1"} space-y-4`}>
            <div className="bg-white rounded-lg shadow p-4 sm:p-6">
              {currentQuestion ? (
                <>
                  {/* ✅ Passage with scroll */}
                  {currentQuestion.passage && (
                    <div className="mb-4 p-3 bg-gray-50 border rounded-md text-sm sm:text-base leading-relaxed overflow-y-auto max-h-64 sm:max-h-80">
                      <p
                        dangerouslySetInnerHTML={{
                          __html: currentQuestion.passage.passage_text,
                        }}
                      />
                    </div>
                  )}

                  {/* ✅ Question Text */}
                  <div className="mb-4">
                    <p className="text-base sm:text-lg mb-3 font-medium text-gray-800">
                      {currentQuestion.question_text}
                    </p>

                    {/* ✅ Image Support (no lazy loading) */}
                    {typeof currentQuestion.question_image_url === "string" &&
                      currentQuestion.question_image_url.length > 0 && (
                        <div className="flex flex-wrap gap-3 mb-4 justify-center">
                          <img
                            src={currentQuestion.question_image_url}
                            alt="question image"
                            loading="eager"
                            className={`rounded-lg border object-contain transition-all ${
                              currentQuestion.question_text.length > 150
                                ? "max-w-xs sm:max-w-sm"
                                : "max-w-sm sm:max-w-md"
                            }`}
                          />
                        </div>
                      )}

                    <QuestionOptions
                      question={currentQuestion}
                      selectedAnswers={selectedAnswers[currentQuestion.id] || []}
                      onAnswerSelect={(answers) =>
                        handleAnswerSelect(currentQuestion.id, answers)
                      }
                    />
                  </div>

                  {/* ✅ Action Buttons */}
                  <div className="flex flex-wrap justify-between border-t pt-3 gap-2 sticky bottom-0 bg-white pb-2">
                    <button
                      onClick={() => clearAnswer(currentQuestion.id)}
                      className="px-4 py-2 bg-red-100 text-red-700 rounded-md hover:bg-red-200 text-sm"
                    >
                      Clear
                    </button>
                    <div className="flex gap-2">
                      <button
                        onClick={() => toggleMarkReview(currentQuestion.id)}
                        className={`px-4 py-2 rounded-md text-sm transition-colors ${
                          questionStatus[currentQuestion.id] === "marked-review"
                            ? "bg-yellow-100 text-yellow-800 hover:bg-yellow-200"
                            : "bg-gray-100 text-gray-700 hover:bg-gray-200"
                        }`}
                      >
                        {questionStatus[currentQuestion.id] === "marked-review"
                          ? "Unmark"
                          : "Mark for Review"}
                      </button>
                      <button
                        onClick={goToPrevious}
                        className="px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300 text-sm"
                      >
                        Prev
                      </button>
                      <button
                        onClick={goToNext}
                        className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 text-sm"
                      >
                        Next
                      </button>
                    </div>
                  </div>
                </>
              ) : (
                <p className="text-gray-500 text-sm">No questions found.</p>
              )}
            </div>
          </div>

          {/* RIGHT PANEL */}
          <div className={`${isMobile ? "order-1" : "w-full lg:w-80"} space-y-4`}>
            <div className="bg-white rounded-lg shadow p-4">
              <h2 className="text-base font-semibold mb-3">Time Left</h2>
              <div className="text-2xl text-center font-mono">
                <CountdownTimer minutes={exam.duration_minutes} onTimerEnd={setTimerEnd} />
              </div>
            </div>

            <div className="bg-white rounded-lg shadow p-4">
              <h2 className="text-base font-semibold mb-3">Sections</h2>
              <div className="space-y-2">
                {sections.map((section, idx) => {
                  const sectionQs = sectionQuestions[section.id] || [];
                  return (
                    <Accordion
                      key={section.id}
                      title={section.subject?.name || `Section ${idx + 1}`}
                      isActive={idx === currentSectionIndex}
                      defaultOpen={idx === currentSectionIndex}
                      isMobile={isMobile}
                    >
                      <div
                        className={`grid ${
                          isMobile ? "grid-cols-4" : "grid-cols-5"
                        } gap-2 p-2`}
                      >
                        {sectionQs.map((question, qIdx) => {
                          const status =
                            questionStatus[question.id] || "not-attempted";
                          let color = "bg-gray-300 hover:bg-gray-400";
                          if (
                            idx === currentSectionIndex &&
                            qIdx === currentQuestionIndex
                          )
                            color = "bg-blue-600 text-white";
                          else if (status === "answered")
                            color = "bg-green-500 text-white";
                          else if (status === "marked-review")
                            color = "bg-yellow-400 text-white";

                          return (
                            <button
                              key={question.id}
                              onClick={() => {
                                setCurrentSectionIndex(idx);
                                setCurrentQuestionIndex(qIdx);
                                window.scrollTo({ top: 0, behavior: "smooth" });
                              }}
                              className={`w-7 h-7 sm:w-8 sm:h-8 rounded-full flex items-center justify-center text-xs font-medium ${color}`}
                            >
                              {qIdx + 1}
                            </button>
                          );
                        })}
                      </div>
                    </Accordion>
                  );
                })}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
