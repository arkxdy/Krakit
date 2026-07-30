import Link from "next/link";
import { Button } from "@/components/ui/Button";

const exams = [
  {
    id: "exam-1",
    name: "Mathematics Mock Test",
    description: "A complete math exam with algebra, calculus, and geometry questions.",
    duration: 120,
    totalMarks: 100,
    type: "MCQ",
  },
  {
    id: "exam-2",
    name: "Physics Practice Test",
    description: "Test your physics knowledge across mechanics, optics, and electricity.",
    duration: 90,
    totalMarks: 100,
    type: "MCQ",
  },
];

export default function ExamPage() {
  return (
    <div className="space-y-10 max-w-7xl mx-auto">
      <div className="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <h1 className="text-4xl font-bold text-white">Available Exams</h1>
          <p className="text-gray-400 mt-2">
            Select an exam to view details and start your practice session.
          </p>
        </div>
        <Link href="/exams/sections">
          <Button variant="primary">Continue Last Exam</Button>
        </Link>
      </div>

      <div className="grid gap-6 md:grid-cols-2">
        {exams.map((exam) => (
          <div key={exam.id} className="bg-gray-800/60 backdrop-blur-sm rounded-3xl p-8 shadow-lg">
            <div className="flex items-center justify-between mb-6">
              <div>
                <h2 className="text-2xl font-semibold text-white">{exam.name}</h2>
                <p className="text-gray-400 mt-2">{exam.description}</p>
              </div>
              <span className="rounded-full bg-primary-500/20 px-4 py-2 text-sm text-primary-200">
                {exam.type}
              </span>
            </div>
            <div className="grid gap-4 sm:grid-cols-3">
              <div className="rounded-2xl bg-gray-900/80 p-4">
                <p className="text-gray-400 text-sm">Duration</p>
                <p className="text-white font-semibold">{exam.duration} mins</p>
              </div>
              <div className="rounded-2xl bg-gray-900/80 p-4">
                <p className="text-gray-400 text-sm">Total Marks</p>
                <p className="text-white font-semibold">{exam.totalMarks}</p>
              </div>
              <div className="rounded-2xl bg-gray-900/80 p-4">
                <p className="text-gray-400 text-sm">Questions</p>
                <p className="text-white font-semibold">50</p>
              </div>
            </div>
            <div className="mt-6 flex justify-end">
              <Link href="/exams/sections">
                <Button variant="primary">View Sections</Button>
              </Link>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
