import Link from "next/link";
import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";
import { Button } from "@/components/ui/Button";

export const metadata = {
  title: "About Krakit",
  description: "Learn more about Krakit's mission and approach to exam preparation.",
};

export default function AboutPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex flex-col text-white">
      <Header />
      <main className="flex-1 container mx-auto px-4 py-16 mt-24">
        <div className="max-w-4xl mx-auto space-y-12">
          <div className="text-center">
            <h1 className="text-4xl font-bold mb-4">About Krakit</h1>
            <p className="text-gray-400">
              Empowering learners with realistic exam practice, performance analytics, and personalized study tools.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
              <h2 className="text-2xl font-semibold text-white mb-4">Our Mission</h2>
              <p className="text-gray-400">
                At Krakit, we believe every student deserves access to high-quality exam preparation resources. Our mission is to provide a supportive and intelligent platform for learners to build confidence and improve performance.
              </p>
            </div>
            <div className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
              <h2 className="text-2xl font-semibold text-white mb-4">Our Vision</h2>
              <p className="text-gray-400">
                We envision a world where exam preparation is accessible, effective, and engaging for everyone. We combine technology with expert-designed content to make learning more efficient.
              </p>
            </div>
          </div>

          <div className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-6">What Sets Us Apart</h2>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <h3 className="text-xl font-semibold text-white mb-2">Quality Content</h3>
                <p className="text-gray-400">
                  Our exams and study materials are created by experienced educators and subject matter experts.
                </p>
              </div>
              <div>
                <h3 className="text-xl font-semibold text-white mb-2">Personalized Learning</h3>
                <p className="text-gray-400">
                  We offer tailored feedback and recommendations based on your performance and learning style.
                </p>
              </div>
              <div>
                <h3 className="text-xl font-semibold text-white mb-2">Analytics</h3>
                <p className="text-gray-400">
                  Track your progress with detailed analytics and insights to identify areas for improvement.
                </p>
              </div>
            </div>
          </div>

          <div className="text-center">
            <h2 className="text-2xl font-semibold text-white mb-4">Join Our Community</h2>
            <p className="text-gray-400 mb-8">
              Start your exam preparation journey today and see how Krakit can help you reach your goals.
            </p>
            <Link href="/signup">
              <Button variant="primary" className="px-8 py-3">
                Get Started
              </Button>
            </Link>
          </div>
        </div>
      </main>
      <Footer />
    </div>
  );
}
