import Link from "next/link";
import { Button } from "@/components/ui/Button";

export function ProUpgrade() {
  return (
    <div className="p-4 border-t border-gray-700/50">
      <div className="bg-gradient-to-r from-primary-600/20 to-primary-500/20 rounded-lg p-4">
        <div className="flex items-center justify-between mb-2">
          <span className="text-sm font-medium text-primary-400">
            Pro Version
          </span>
          <span className="text-xs bg-primary-500/20 text-primary-300 px-2 py-1 rounded-full">
            Upgrade
          </span>
        </div>
        <p className="text-sm text-gray-400 mb-3">
          Get access to all premium features and unlimited mock exams
        </p>
        <Link href="/pricing">
          <Button variant="primary" className="w-full text-sm">
            Upgrade Now
          </Button>
        </Link>
      </div>
    </div>
  );
}
