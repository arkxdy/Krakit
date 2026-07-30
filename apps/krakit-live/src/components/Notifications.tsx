"use client";

interface NotificationsProps {
  isOpen: boolean;
  onClose: () => void;
}

export function Notifications({ isOpen, onClose }: NotificationsProps) {
  if (!isOpen) return null;
  return (
    <div className="fixed inset-0 z-50 overflow-y-auto">
      <div className="flex items-center justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
        {/* Background overlay */}
        <div
          className="fixed inset-0 transition-opacity bg-black/90"
          onClick={onClose}
        />

        {/* Modal panel */}
        <div className="relative z-10 inline-block w-full max-w-md p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-gray-900 shadow-xl rounded-2xl border border-purple-500/20">
          {/* Header */}
          <div className="flex items-center justify-between mb-4">
            <h3 className="text-lg font-medium text-purple-400">
              Notifications
            </h3>
            <button
              onClick={onClose}
              className="text-purple-400 hover:text-purple-300 transition-colors"
            >
              <svg
                className="w-6 h-6"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </div>

          {/* Notifications list */}
          <div className="space-y-4">
            {/* Example notifications - replace with real data */}
            <div className="p-4 bg-gray-800/50 rounded-lg border border-purple-500/10 hover:border-purple-500/30 transition-colors">
              <p className="text-sm text-purple-200">
                New exam results are available
              </p>
              <p className="text-xs text-purple-400/70 mt-1">2 hours ago</p>
            </div>
            <div className="p-4 bg-gray-800/50 rounded-lg border border-purple-500/10 hover:border-purple-500/30 transition-colors">
              <p className="text-sm text-purple-200">Study material updated</p>
              <p className="text-xs text-purple-400/70 mt-1">5 hours ago</p>
            </div>
            <div className="p-4 bg-gray-800/50 rounded-lg border border-purple-500/10 hover:border-purple-500/30 transition-colors">
              <p className="text-sm text-purple-200">
                New practice test available
              </p>
              <p className="text-xs text-purple-400/70 mt-1">1 day ago</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
