export const metadata = {
  title: "Exam Results - Krakit",
  description: "View your latest exam results and performance analytics.",
};

const results = [
  { title: "Mathematics Mock Test", score: "92%", date: "March 15, 2024" },
  { title: "Physics Practice Test", score: "78%", date: "March 12, 2024" },
  { title: "Chemistry Mock Test", score: "85%", date: "March 10, 2024" },
];

export default function ResultsPage() {
  return (
    <div className="max-w-7xl mx-auto space-y-10 text-white">
      <div>
        <h1 className="text-4xl font-bold mb-2">Exam Results</h1>
        <p className="text-gray-400">Review your performance and compare your latest exam scores.</p>
      </div>

      <div className="grid gap-6 lg:grid-cols-3">
        {results.map((result) => (
          <div key={result.title} className="rounded-3xl bg-gray-800/60 backdrop-blur-sm p-8">
            <h2 className="text-2xl font-semibold text-white mb-3">{result.title}</h2>
            <p className="text-primary-400 text-3xl font-bold mb-4">{result.score}</p>
            <p className="text-gray-400">Completed {result.date}</p>
          </div>
        ))}
      </div>

      <div className="rounded-3xl bg-gray-800/60 backdrop-blur-sm p-8">
        <h2 className="text-2xl font-semibold text-white mb-4">Performance Overview</h2>
        <div className="grid gap-6 md:grid-cols-2">
          <div className="rounded-2xl bg-gray-900/80 p-6">
            <p className="text-gray-400 mb-2">Overall Average</p>
            <p className="text-white text-4xl font-bold">82%</p>
          </div>
          <div className="rounded-2xl bg-gray-900/80 p-6">
            <p className="text-gray-400 mb-2">Highest Score</p>
            <p className="text-white text-4xl font-bold">92%</p>
          </div>
        </div>
      </div>
    </div>
  );
}
