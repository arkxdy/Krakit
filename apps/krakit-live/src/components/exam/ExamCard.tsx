import type { Exam, Subject } from "@/types/exam.types";

interface ExamCardProps {
  exam: Exam;
  subjects?: Subject[]
}

export function ExamCard({ exam }: ExamCardProps) {
  return (
    <div className="card-container bg-gray-800/50 backdrop-blur-sm rounded-xl p-6">
      <div className="flex items-center space-x-4 mb-4">
        <div className="w-10 h-10 flex items-center justify-center text-primary-400 text-2xl">
          {getSubjectIcon('Mathematics')}
        </div>
        <div>
          <h3 className="text-lg font-semibold text-white">{exam.name}</h3>
          <p className="text-gray-400 text-sm">{exam.description}</p>
        </div>
      </div>
      <p className="text-gray-300 mb-4">{exam.description}</p>
      <div className="flex justify-between items-center text-sm mb-4">
        <span className="text-gray-400">{exam.exam_type}</span>
        <span className="text-gray-400">{exam.duration_minutes} minutes</span>
      </div>
      <button className="w-full px-4 py-2 bg-primary-600 hover:bg-primary-700 text-white rounded-lg transition-colors">
        Start Exam
      </button>
    </div>
  );
}

function getSubjectIcon(subject: string): string {
  switch (subject) {
    case 'Mathematics':
      return '📐';
    case 'Physics':
      return '⚡';
    case 'Chemistry':
      return '🧪';
    case 'Biology':
      return '🔬';
    default:
      return '📚';
  }
}
