import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";

export const metadata = {
  title: "Contact Us - Krakit",
  description: "Get in touch with Krakit support for exam preparation assistance.",
};

export default function ContactPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex flex-col text-white">
      <Header />
      <div className="flex-1 container mx-auto px-4 py-16 mt-24">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-12">
            <h1 className="text-4xl font-bold mb-4">Contact Us</h1>
            <p className="text-gray-400">
              Need help with your exam preparation? Our support team is here to
              assist you.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
              <form className="space-y-6">
                <div>
                  <label
                    htmlFor="name"
                    className="block text-sm font-medium text-gray-300 mb-2"
                  >
                    Name
                  </label>
                  <input
                    type="text"
                    id="name"
                    className="w-full px-4 py-2 bg-gray-700/50 border border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 text-white"
                    placeholder="Your name"
                  />
                </div>

                <div>
                  <label
                    htmlFor="email"
                    className="block text-sm font-medium text-gray-300 mb-2"
                  >
                    Email
                  </label>
                  <input
                    type="email"
                    id="email"
                    className="w-full px-4 py-2 bg-gray-700/50 border border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 text-white"
                    placeholder="you@example.com"
                  />
                </div>

                <div>
                  <label
                    htmlFor="subject"
                    className="block text-sm font-medium text-gray-300 mb-2"
                  >
                    Subject
                  </label>
                  <select
                    id="subject"
                    className="w-full px-4 py-2 bg-gray-700/50 border border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 text-white"
                  >
                    <option value="">Select a subject</option>
                    <option value="technical">Technical Support</option>
                    <option value="exam">Exam Related</option>
                    <option value="billing">Billing & Subscription</option>
                    <option value="other">Other</option>
                  </select>
                </div>

                <div>
                  <label
                    htmlFor="message"
                    className="block text-sm font-medium text-gray-300 mb-2"
                  >
                    Message
                  </label>
                  <textarea
                    id="message"
                    rows={4}
                    className="w-full px-4 py-2 bg-gray-700/50 border border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 text-white"
                    placeholder="How can we help you?"
                  />
                </div>

                <button
                  type="submit"
                  className="w-full py-3 px-4 bg-primary-600 hover:bg-primary-700 text-white font-medium rounded-lg transition-colors"
                >
                  Send Message
                </button>
              </form>
            </div>

            <div className="space-y-8">
              <div className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
                <h3 className="text-xl font-semibold text-white mb-4">
                  Support Information
                </h3>
                <div className="space-y-4">
                  <div className="flex items-start space-x-4">
                    <div className="text-primary-400 text-xl">📧</div>
                    <div>
                      <p className="text-white font-medium">Email Support</p>
                      <p className="text-gray-400">support@krakit.com</p>
                    </div>
                  </div>
                  <div className="flex items-start space-x-4">
                    <div className="text-primary-400 text-xl">💬</div>
                    <div>
                      <p className="text-white font-medium">Live Chat</p>
                      <p className="text-gray-400">Available 24/7</p>
                    </div>
                  </div>
                  <div className="flex items-start space-x-4">
                    <div className="text-primary-400 text-xl">📞</div>
                    <div>
                      <p className="text-white font-medium">Phone Support</p>
                      <p className="text-gray-400">+1 (555) 123-4567</p>
                    </div>
                  </div>
                </div>
              </div>

              <div className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-8">
                <h3 className="text-xl font-semibold text-white mb-4">
                  Support Hours
                </h3>
                <div className="space-y-2">
                  <p className="text-gray-400">Monday - Friday: 8:00 AM - 8:00 PM</p>
                  <p className="text-gray-400">Saturday: 9:00 AM - 5:00 PM</p>
                  <p className="text-gray-400">Sunday: 10:00 AM - 4:00 PM</p>
                  <p className="text-primary-400 mt-4">Premium users get 24/7 support</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}
