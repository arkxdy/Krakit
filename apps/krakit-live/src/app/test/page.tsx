import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";

export const metadata = {
  title: "Test - Krakit",
  description: "Testing page for Krakit mock exam app.",
};

export default function TestPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex flex-col text-white">
      <Header />
      <main className="flex-1 container mx-auto px-4 py-24 mt-24">
        <div className="max-w-3xl mx-auto text-center">
          <h1 className="text-4xl font-bold mb-4">Test Page</h1>
          <p className="text-gray-400">
            This page exists for route parity with the original web-ui source.
          </p>
        </div>
      </main>
      <Footer />
    </div>
  );
}
