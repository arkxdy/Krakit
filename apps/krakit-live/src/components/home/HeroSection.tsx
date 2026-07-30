import Link from "next/link";
import { Button } from "@/components/ui/Button";

export function HeroSection() {
  return (
    <section className="relative pt-32 pb-20 overflow-hidden">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <h1 className="text-4xl md:text-6xl font-bold text-white mb-6">
            Master Your Exams with{" "}
            <span className="text-primary-400">Krakit</span>
          </h1>
          <p className="text-xl text-gray-300 mb-8 max-w-3xl mx-auto">
            Practice with our comprehensive collection of mock exams, track your
            progress, and improve your scores with detailed analytics and
            personalized feedback.
          </p>
          <div className="flex justify-center space-x-4">
            <Link href="/signup">
              <Button
                variant="primary"
                className="px-8 py-3 text-lg font-semibold"
              >
                Get Started
              </Button>
            </Link>
            <Link href="/exams">
              <Button
                variant="secondary"
                className="px-8 py-3 text-lg font-semibold"
              >
                Browse Exams
              </Button>
            </Link>
          </div>
        </div>
      </div>

      {/* Background Gradient */}
      <div className="absolute inset-0 -z-10">
        <div className="absolute inset-0 bg-gradient-to-b from-gray-900 via-gray-900 to-primary-900/20" />
      </div>
    </section>
  );
}
