export const metadata = {
  title: "Progress - Krakit",
  description: "Track your learning progress and performance.",
};

const stats = [
  { label: "Average Score", value: "82%" },
  { label: "Exams Completed", value: "12/20" },
  { label: "Study Hours", value: "45h" },
];

export default function ProgressPage() {
  return (
    <div className="max-w-7xl mx-auto space-y-10 text-white">
      <div>
        <h1 className="text-4xl font-bold mb-2">Your Progress</h1>
        <p className="text-gray-400">Review your exam performance and recent activity.</p>
      </div>

      <div className="grid gap-6 md:grid-cols-3">
        {stats.map((stat) => (
          <div key={stat.label} className="rounded-3xl bg-gray-800/60 backdrop-blur-sm p-8">
            <p className="text-gray-400">{stat.label}</p>
            <p className="text-4xl font-bold mt-4 text-white">{stat.value}</p>
          </div>
        ))}
      </div>

      <div className="grid gap-6 lg:grid-cols-2">
        <div className="rounded-3xl bg-gray-800/60 backdrop-blur-sm p-8">
          <h2 className="text-2xl font-semibold text-white mb-4">Subject Performance</h2>
          {[
            { subject: "Mathematics", progress: "88%" },
            { subject: "Physics", progress: "75%" },
            { subject: "Chemistry", progress: "82%" },
          ].map((item) => (
            <div key={item.subject} className="mb-6">
              <div className="flex justify-between text-gray-400 mb-2">
                <span>{item.subject}</span>
                <span>{item.progress}</span>
              </div>
              <div className="w-full rounded-full bg-gray-900/70 h-3 overflow-hidden">
                <div className="h-3 bg-primary-500" style={{ width: item.progress }} />
              </div>
            </div>
          ))}
        </div>

        <div className="rounded-3xl bg-gray-800/60 backdrop-blur-sm p-8">
          <h2 className="text-2xl font-semibold text-white mb-4">Recent Activity</h2>
          {[
            { title: "Completed Mathematics Mock Test", detail: "Score: 92% • 2 days ago" },
            { title: "Completed Physics Practice Test", detail: "Score: 78% • 5 days ago" },
            { title: "Completed Chemistry Mock Test", detail: "Score: 85% • 1 week ago" },
          ].map((activity) => (
            <div key={activity.title} className="rounded-2xl bg-gray-900/80 p-4 mb-4">
              <h3 className="text-white font-medium">{activity.title}</h3>
              <p className="text-gray-400 mt-2">{activity.detail}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
