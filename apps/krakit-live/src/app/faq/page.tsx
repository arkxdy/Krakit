import Link from "next/link";
import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";

export const metadata = {
  title: "FAQ - Krakit",
  description: "Frequently asked questions about Krakit's exam platform.",
};

const faqs = [
  {
    question: "What is Krakit?",
    answer:
      "Krakit is a comprehensive exam preparation platform that helps learners practice with mock exams, track progress, and improve performance.",
  },
  {
    question: "How do I get started?",
    answer:
      "Create an account, log in, and begin practicing with a wide range of exams and learning resources.",
  },
  {
    question: "What types of exams are available?",
    answer:
      "Krakit supports a variety of subjects and exam formats with customizable mock tests.",
  },
  {
    question: "Can I retake exams?",
    answer:
      "Yes, you can retake exams multiple times to improve your score and track progress.",
  },
  {
    question: "How are results tracked?",
    answer:
      "Krakit provides analytics and progress metrics to help you understand strengths and improvement areas.",
  },
];

export default function FAQPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex flex-col text-white">
      <Header />
      <div className="flex-1 container mx-auto px-4 py-16 mt-24">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-12">
            <h1 className="text-4xl font-bold mb-4">Frequently Asked Questions</h1>
            <p className="text-gray-400">
              Find answers to common questions about Krakit.
            </p>
          </div>

          <div className="space-y-8">
            {faqs.map((faq) => (
              <section
                key={faq.question}
                className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8"
              >
                <h2 className="text-2xl font-semibold text-white mb-4">
                  {faq.question}
                </h2>
                <p className="text-gray-400">{faq.answer}</p>
              </section>
            ))}
          </div>

          <div className="mt-16 text-center">
            <h2 className="text-2xl font-semibold text-white mb-4">
              Still have questions?
            </h2>
            <p className="text-gray-400 mb-8">
              Can't find the answer you're looking for? Get in touch with our support team.
            </p>
            <Link href="/contact" className="inline-block">
              <button className="px-8 py-3 bg-primary-600 rounded-lg text-white hover:bg-primary-700 transition-colors">
                Contact Us
              </button>
            </Link>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}
