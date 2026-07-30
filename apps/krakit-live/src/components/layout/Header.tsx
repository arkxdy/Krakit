import Link from "next/link";
import { Button } from "@/components/ui/Button";
import { krakitIcon } from "@/utils/constant";

export function Header() {
  return (
    <header
      className="fixed top-0 left-0 right-0 z-50 bg-gray-900/80 backdrop-blur-sm border-b border-gray-800"
      style={{ minHeight: "64px" }}
    >
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo */}
          <Link href="/" className="flex items-center space-x-2">
            <span className="text-2xl">
              <img
                src={krakitIcon}
                alt="Krakit Logo"
                className="w-6 h-6"
                width="24"
                height="24"
              />
            </span>
            <span className="text-xl font-bold text-white">Krakit</span>
          </Link>

          {/* Navigation */}
          <nav className="hidden md:flex items-center space-x-8">
            <Link
              href="/exams"
              className="text-gray-300 hover:text-white transition-colors"
            >
              Exams
            </Link>
            <Link
              href="/materials"
              className="text-gray-300 hover:text-white transition-colors"
            >
              Study Materials
            </Link>
            <Link
              href="/progress"
              className="text-gray-300 hover:text-white transition-colors"
            >
              Progress
            </Link>
          </nav>

          {/* Auth Buttons */}
          <div className="flex items-center space-x-4">
            <Link
              href="/login"
              className="text-gray-300 hover:text-white transition-colors"
            >
              Log in
            </Link>
            <Link href="/signup">
              <Button variant="primary" className="px-4 py-2">
                Sign up
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </header>
  );
}
