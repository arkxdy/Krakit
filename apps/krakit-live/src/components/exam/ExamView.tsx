import { useState, useEffect, useMemo } from "react";
import Accordion from "@/components/ui/Accordian";
import type { BaseQuestion, Exam, ExamSection } from "@/types/exam.types";
import { CountdownTimer } from "./CountdownTimer";
import { QuestionOptions } from "./QuestionOptions";
import { QuestionRenderer } from "./Question";

type ExamViewProps = {
  exam: Exam;
  sections: ExamSection[];
  questions: BaseQuestion[];
};

export default function ExamView({ exam, sections, questions }: ExamViewProps) {
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

  // ✅ Group questions per section
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
  };

  // 🧱 Layout
  return (
    <div className="min-h-screen bg-gray-50 p-2 sm:p-4 md:p-6">
      <div className="max-w-7xl mx-auto">
        <header className="mb-4 sm:mb-6">
          <h1 className="text-xl sm:text-2xl font-bold text-gray-900">
            {exam.name}
          </h1>
          <p className="text-gray-600 text-sm">{exam.description}</p>
        </header>

        {/* Mobile Controls */}
        {isMobile && (
          <div className="bg-white rounded-lg shadow p-4 mb-4 flex justify-between items-center">
            <button
              onClick={goToPrevious}
              disabled={currentSectionIndex === 0 && currentQuestionIndex === 0}
              className="px-3 py-1 bg-gray-100 text-gray-700 rounded-md disabled:opacity-50 text-sm hover:bg-gray-200 transition-colors duration-200"
            >
              Previous
            </button>
            <span className="text-sm font-medium text-gray-500">
              Q{currentQuestionIndex + 1}/{currentSectionQuestions.length}
            </span>
            <button
              onClick={goToNext}
              disabled={
                currentQuestionIndex === currentSectionQuestions.length - 1 &&
                currentSectionIndex === sections.length - 1
              }
              className="px-3 py-1 bg-blue-600 text-white rounded-md disabled:opacity-50 text-sm hover:bg-blue-700 transition-colors duration-200"
            >
              {currentQuestionIndex === currentSectionQuestions.length - 1 &&
              currentSectionIndex === sections.length - 1
                ? "Submit"
                : "Next"}
            </button>
          </div>
        )}

        {/* Two Column Layout */}
        <div className="flex flex-col lg:flex-row gap-4 sm:gap-6">
          {/* LEFT: Question Panel */}
          <div className={`${isMobile ? "order-2" : "flex-1"} space-y-4`}>
            <div className="bg-white rounded-lg shadow p-4 sm:p-6">
              {!isMobile && (
                <div className="flex justify-between items-center mb-4">
                  <span className="text-sm text-gray-500">
                    {currentSection.subject?.name || "Section"} - Question{" "}
                    {currentQuestionIndex + 1}
                  </span>
                  <span className="text-sm text-gray-500">
                    {currentQuestionIndex + 1} / {currentSectionQuestions.length}
                  </span>
                </div>
              )}

              {currentQuestion ? (
                <>
                  <div className="mb-4">
                    <p className="text-base sm:text-lg mb-3">
                      {currentQuestion.question_text}
                    </p>
                    {/* <QuestionOptions
                      question={currentQuestion}
                      selectedAnswers={
                        selectedAnswers[currentQuestion.id] || []
                      }
                      onAnswerSelect={(answers) =>
                        handleAnswerSelect(currentQuestion.id, answers)
                      }
                    /> */}
                    <QuestionRenderer
                      question={currentQuestion}
                      selectedAnswers={selectedAnswers[currentQuestion.id] || []}
                      onSelect={(answers) => handleAnswerSelect(currentQuestion.id, answers)}
                    />
                  </div>

                  {/* Navigation */}
                  {!isMobile && (
                    <div className="flex justify-between border-t pt-3">
                      <button
                        onClick={goToPrevious}
                        disabled={
                          currentSectionIndex === 0 &&
                          currentQuestionIndex === 0
                        }
                        className="px-4 py-2 bg-gray-100 text-gray-700 rounded-md disabled:opacity-50 hover:bg-gray-200"
                      >
                        Previous
                      </button>
                      <div className="flex gap-2">
                        <button
                          onClick={() =>
                            toggleMarkReview(currentQuestion.id)
                          }
                          className={`px-4 py-2 rounded-md transition-colors duration-200 ${
                            questionStatus[currentQuestion.id] ===
                            "marked-review"
                              ? "bg-yellow-100 text-yellow-800 hover:bg-yellow-200"
                              : "bg-gray-100 text-gray-700 hover:bg-gray-200"
                          }`}
                        >
                          {questionStatus[currentQuestion.id] ===
                          "marked-review"
                            ? "Unmark Review"
                            : "Mark for Review"}
                        </button>
                        <button
                          onClick={goToNext}
                          disabled={
                            currentQuestionIndex ===
                              currentSectionQuestions.length - 1 &&
                            currentSectionIndex === sections.length - 1
                          }
                          className="px-4 py-2 bg-blue-600 text-white rounded-md disabled:opacity-50 hover:bg-blue-700"
                        >
                          {currentQuestionIndex ===
                            currentSectionQuestions.length - 1 &&
                          currentSectionIndex === sections.length - 1
                            ? "Submit"
                            : "Next"}
                        </button>
                      </div>
                    </div>
                  )}
                </>
              ) : (
                <p className="text-gray-500 text-sm">No questions found.</p>
              )}
            </div>

            {/* Legend */}
            <div className="bg-white rounded-lg shadow p-4">
              <h2 className="text-base font-semibold mb-3">Question Status</h2>
              <div className="grid grid-cols-2 gap-3 text-sm">
                <Legend color="bg-blue-600" label="Current" />
                <Legend color="bg-gray-300" label="Not Attempted" />
                <Legend color="bg-green-500" label="Answered" />
                <Legend color="bg-yellow-400" label="Marked Review" />
              </div>
            </div>
          </div>

          {/* RIGHT: Timer & Sections */}
          <div className={`${isMobile ? "order-1" : "w-full lg:w-80"} space-y-4`}>
            <div className="bg-white rounded-lg shadow p-4">
              <h2 className="text-base font-semibold mb-3">Time Left</h2>
              <div className="text-2xl text-center font-mono">
                <CountdownTimer
                  minutes={exam.duration_minutes}
                  onTimerEnd={setTimerEnd}
                />
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

function Legend({ color, label }: { color: string; label: string }) {
  return (
    <div className="flex items-center">
      <div className={`w-4 h-4 ${color} rounded-full mr-2`}></div>
      <span className="text-gray-600 text-xs">{label}</span>
    </div>
  );
}
