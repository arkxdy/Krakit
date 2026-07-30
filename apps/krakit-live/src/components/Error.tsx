"use client";

import { Footer } from "./layout/Footer";
import { Header } from "./layout/Header";

export function Error(props: { errorMessage: string }) {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex flex-col">
      <Header />
      <div className="flex-1 container mx-auto px-4 py-16 mt-16">
        <div className="max-w-3xl mx-auto text-center">
          <div className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
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
                  d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
                />
              </svg>
            </div>
            <h1 className="text-4xl font-bold text-white mb-4">
              Oops! Something went wrong
            </h1>
            <p className="text-gray-400 mb-8">{props.errorMessage}</p>
            <div className="space-y-4">
              <a
                href="/"
                className="inline-block px-6 py-3 bg-primary-600 text-white font-medium rounded-lg hover:bg-primary-700 transition-colors"
              >
                Return Home
              </a>
              <button
                onClick={() => window.location.reload()}
                className="block w-full px-6 py-3 bg-gray-700 text-white font-medium rounded-lg hover:bg-gray-600 transition-colors"
              >
                Try Again
              </button>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}
