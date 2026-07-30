import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";

export const metadata = {
  title: "Privacy Policy - Krakit",
  description: "Learn how Krakit handles your data and privacy.",
};

export default function PrivacyPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex flex-col text-white">
      <Header />
      <div className="flex-1 container mx-auto px-4 py-16 mt-24">
        <div className="max-w-4xl mx-auto space-y-8">
          <div className="text-center mb-12">
            <h1 className="text-4xl font-bold mb-4">Privacy Policy</h1>
            <p className="text-gray-400">Last updated: {new Date().toLocaleDateString()}</p>
          </div>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">Information We Collect</h2>
            <div className="space-y-4 text-gray-400">
              <p className="text-white">
                We collect information that you provide directly to us, including:
              </p>
              <ul className="list-disc list-inside space-y-2">
                <li>Account information (name, email address, password)</li>
                <li>Profile information (educational background, preferences)</li>
                <li>Exam results and progress data</li>
                <li>Communication preferences</li>
              </ul>
            </div>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">How We Use Your Information</h2>
            <div className="space-y-4 text-gray-400">
              <p className="text-white">We use the information we collect to:</p>
              <ul className="list-disc list-inside space-y-2">
                <li>Provide and maintain our services</li>
                <li>Personalize your learning experience</li>
                <li>Track your progress and provide feedback</li>
                <li>Communicate with you about our services</li>
                <li>Improve our platform and develop new features</li>
              </ul>
            </div>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">Data Security</h2>
            <div className="space-y-4 text-gray-400">
              <p className="text-white">
                We implement appropriate security measures to protect your personal information:
              </p>
              <ul className="list-disc list-inside space-y-2">
                <li>Encryption of sensitive data</li>
                <li>Regular security assessments</li>
                <li>Access controls and authentication</li>
                <li>Secure data storage and transmission</li>
              </ul>
            </div>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">Your Rights</h2>
            <div className="space-y-4 text-gray-400">
              <p className="text-white">You have the right to:</p>
              <ul className="list-disc list-inside space-y-2">
                <li>Access your personal information</li>
                <li>Correct inaccurate data</li>
                <li>Request deletion of your data</li>
                <li>Opt-out of marketing communications</li>
                <li>Export your data</li>
              </ul>
            </div>
          </section>

          <section className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">Contact Us</h2>
            <div className="space-y-4 text-gray-400">
              <p className="text-white">
                If you have any questions about this Privacy Policy, please contact us at:
              </p>
              <p className="text-primary-400">privacy@krakit.com</p>
            </div>
          </section>
        </div>
      </div>
      <Footer />
    </div>
  );
}
