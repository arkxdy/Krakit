import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";

export const metadata = {
  title: "Terms of Service - Krakit",
  description: "Read our terms and conditions for using Krakit.",
};

export default function TermsPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex flex-col text-white">
      <Header />
      <div className="flex-1 container mx-auto px-4 py-16 mt-24">
        <div className="max-w-4xl mx-auto space-y-8">
          <div className="text-center mb-12">
            <h1 className="text-4xl font-bold mb-4">Terms of Service</h1>
            <p className="text-gray-400">Last updated: {new Date().toLocaleDateString()}</p>
          </div>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">1. Acceptance of Terms</h2>
            <p className="text-gray-400">
              By accessing and using Krakit, you agree to be bound by these Terms of Service and all applicable laws and regulations.
            </p>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">2. Use License</h2>
            <p className="text-gray-400">
              Permission is granted to temporarily access the materials on Krakit for personal, non-commercial transitory viewing only.
            </p>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">3. User Accounts</h2>
            <p className="text-gray-400">
              To access certain features, you must register for an account and maintain the security of your credentials.
            </p>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">4. Content and Conduct</h2>
            <p className="text-gray-400">
              Users must not share exam content, use bots, manipulate results, or engage in any form of academic dishonesty.
            </p>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">5. Disclaimer</h2>
            <p className="text-gray-400">
              Materials on Krakit are provided on an 'as is' basis without warranties of any kind.
            </p>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">6. Contact Information</h2>
            <p className="text-gray-400">
              If you have questions about these Terms, contact legal@krakit.com.
            </p>
          </section>
        </div>
      </div>
      <Footer />
    </div>
  );
}
