const sections = [
  { id: "s1", name: "Mathematics", questions: 20, time: 30 },
  { id: "s2", name: "Physics", questions: 20, time: 30 },
  { id: "s3", name: "Chemistry", questions: 20, time: 30 },
];

export default function ExamSectionsPage() {
  return (
    <div className="max-w-7xl mx-auto space-y-10 text-white">
      <div>
        <h1 className="text-4xl font-bold mb-2">Exam Sections</h1>
        <p className="text-gray-400">Choose the section you want to begin with for this exam.</p>
      </div>

      <div className="grid gap-6 md:grid-cols-3">
        {sections.map((section) => (
          <div key={section.id} className="rounded-3xl bg-gray-800/60 backdrop-blur-sm p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">{section.name}</h2>
            <p className="text-gray-400 mb-2">Questions: {section.questions}</p>
            <p className="text-gray-400 mb-6">Time limit: {section.time} minutes</p>
            <a
              href="/exams/start"
              className="inline-flex w-full justify-center rounded-2xl bg-purple-600 px-4 py-3 text-sm font-semibold text-white shadow-sm transition hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500"
            >
              Start Section
            </a>
          </div>
        ))}
      </div>
    </div>
  );
}
