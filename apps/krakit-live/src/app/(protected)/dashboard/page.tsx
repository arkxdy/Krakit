import { Button } from "@/components/ui/Button";
import Link from "next/link";

export default function DashboardPage() {
  return (
    <div className="max-w-7xl mx-auto space-y-10">
      <div className="grid gap-6 sm:grid-cols-2 xl:grid-cols-3">
        <div className="bg-gray-800/60 backdrop-blur-sm rounded-3xl p-8">
          <h2 className="text-2xl font-semibold text-white mb-4">Recent Exams</h2>
          <div className="space-y-4">
            <div className="bg-gray-900/70 rounded-2xl p-4">
              <p className="text-white font-medium">Mathematics Mock Test</p>
              <p className="text-gray-400 text-sm">Score: 85% • 2 days ago</p>
            </div>
            <div className="bg-gray-900/70 rounded-2xl p-4">
              <p className="text-white font-medium">Physics Practice Test</p>
              <p className="text-gray-400 text-sm">Score: 78% • 5 days ago</p>
            </div>
          </div>
        </div>

        <div className="bg-gray-800/60 backdrop-blur-sm rounded-3xl p-8">
          <h2 className="text-2xl font-semibold text-white mb-4">Upcoming Exams</h2>
          <div className="space-y-4">
            <div className="bg-gray-900/70 rounded-2xl p-4">
              <p className="text-white font-medium">Chemistry Mock Test</p>
              <p className="text-gray-400 text-sm">Scheduled: Tomorrow</p>
            </div>
            <div className="bg-gray-900/70 rounded-2xl p-4">
              <p className="text-white font-medium">Biology Practice Test</p>
              <p className="text-gray-400 text-sm">Scheduled: Next week</p>
            </div>
          </div>
        </div>

        <div className="bg-gray-800/60 backdrop-blur-sm rounded-3xl p-8">
          <h2 className="text-2xl font-semibold text-white mb-4">Performance</h2>
          <div className="space-y-4">
            <div className="rounded-2xl bg-gray-900/70 p-4">
              <p className="text-gray-400 text-sm">Overall Progress</p>
              <p className="text-white text-3xl font-semibold">75%</p>
            </div>
            <div className="rounded-2xl bg-gray-900/70 p-4">
              <p className="text-gray-400 text-sm">Average Score</p>
              <p className="text-white text-3xl font-semibold">82%</p>
            </div>
          </div>
        </div>
      </div>

      <div className="bg-gray-800/60 backdrop-blur-sm rounded-3xl p-8">
        <div className="flex flex-col gap-6 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <h1 className="text-4xl font-bold text-white">Dashboard</h1>
            <p className="text-gray-400 mt-2">Access your exams, materials, progress, and results.</p>
          </div>
          <div className="flex flex-wrap gap-4">
            <Link href="/exams">
              <Button variant="primary">View Exams</Button>
            </Link>
            <Link href="/materials">
              <Button variant="outline">Study Materials</Button>
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
