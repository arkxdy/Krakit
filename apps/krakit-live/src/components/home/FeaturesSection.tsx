"use client";

import { useState, useEffect } from "react";

const features = [
  {
    title: "Comprehensive Exam Library",
    description:
      "Access a vast collection of mock exams across various subjects and difficulty levels.",
    icon: "📚",
  },
  {
    title: "Real-time Progress Tracking",
    description:
      "Monitor your performance with detailed analytics and progress reports.",
    icon: "📊",
  },
  {
    title: "Personalized Feedback",
    description:
      "Get detailed explanations and suggestions for improvement after each exam.",
    icon: "💡",
  },
  {
    title: "Study Materials",
    description:
      "Access curated study materials and resources to enhance your preparation.",
    icon: "📝",
  },
];

export function FeaturesSection() {
  const [currentFeature, setCurrentFeature] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentFeature((prev) => (prev + 1) % features.length);
    }, 3000);

    return () => clearInterval(interval);
  }, []);

  return (
    <section className="py-20 bg-gray-900">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-16">
          <h2 className="text-3xl font-bold text-white mb-4">
            Why Choose Krakit?
          </h2>
          <p className="text-gray-400 max-w-2xl mx-auto">
            Our platform offers everything you need to excel in your exams
          </p>
        </div>

        <div className="relative h-64">
          {features.map((feature, index) => (
            <div
              key={feature.title}
              className={`absolute inset-0 transition-opacity duration-500 ${
                index === currentFeature ? "opacity-100" : "opacity-0"
              }`}
            >
              <div className="text-center">
                <span className="text-4xl mb-4 block">{feature.icon}</span>
                <h3 className="text-xl font-semibold text-white mb-2">
                  {feature.title}
                </h3>
                <p className="text-gray-400">{feature.description}</p>
              </div>
            </div>
          ))}
        </div>

        {/* Feature Indicators */}
        <div className="flex justify-center space-x-2 mt-8">
          {features.map((_, index) => (
            <button
              key={index}
              onClick={() => setCurrentFeature(index)}
              className={`w-2 h-2 rounded-full transition-colors ${
                index === currentFeature
                  ? "bg-primary-500"
                  : "bg-gray-600 hover:bg-gray-500"
              }`}
              aria-label={`Go to feature ${index + 1}`}
            />
          ))}
        </div>
      </div>
    </section>
  );
}
