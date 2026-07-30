"use client";

import Link from "next/link";

export function NotFound() {
  return (
    <div className="min-h-screen bg-gray-900">
      <main className="flex items-center justify-center p-4">
        <div className="w-full max-w-md">
          <div className="bg-gray-800 rounded-xl p-8 shadow-xl border border-gray-700">
            <div className="mb-6">
              <svg
                className="w-16 h-16 text-primary-500 mx-auto"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                />
              </svg>
            </div>
            <h1 className="text-4xl font-bold text-white mb-4 text-center">
              404 - Page Not Found
            </h1>
            <p className="text-gray-400 mb-8 text-center">
              The page you're looking for doesn't exist or has been moved.
            </p>
            <div className="space-y-4">
              <Link href="/">
                <button className="block w-full px-6 py-3 bg-primary-600 text-white font-medium rounded-lg hover:bg-primary-700 transition-colors text-center">
                  Return Home
                </button>
              </Link>
              <button
                onClick={() => window.history.back()}
                className="w-full px-6 py-3 bg-gray-700 text-white font-medium rounded-lg hover:bg-gray-600 transition-colors"
              >
                Go Back
              </button>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}
