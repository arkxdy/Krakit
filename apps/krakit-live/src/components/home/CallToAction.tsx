import Link from "next/link";
import { Button } from "@/components/ui/Button";

export function CallToAction() {
  return (
    <section className="py-20 bg-primary-900">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <h2 className="text-3xl font-bold text-white mb-4">
            Ready to Start Your Journey?
          </h2>
          <p className="text-gray-300 mb-8 max-w-2xl mx-auto">
            Take your first step towards success by trying one of our mock
            exams. Get instant feedback and start improving today.
          </p>
          <div className="flex justify-center space-x-4">
            <Link href="/exams">
              <Button
                variant="white"
                className="px-8 py-3 text-lg font-semibold"
              >
                Take a Test
              </Button>
            </Link>
            <Link href="/signup">
              <Button
                variant="primary"
                className="px-8 py-3 text-lg font-semibold"
              >
                Sign Up Now
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
}
