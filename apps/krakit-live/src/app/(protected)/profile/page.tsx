import { Button } from "@/components/ui/Button";

export const metadata = {
  title: "Profile - Krakit",
  description: "Manage your profile and account settings.",
};

export default function ProfilePage() {
  return (
    <div className="max-w-4xl mx-auto space-y-10 text-white">
      <div>
        <h1 className="text-4xl font-bold mb-2">Profile Settings</h1>
        <p className="text-gray-400">Update your personal information and notification preferences.</p>
      </div>

      <div className="space-y-8">
        <div className="bg-gray-800/60 backdrop-blur-sm rounded-3xl p-8">
          <h2 className="text-2xl font-semibold text-white mb-6">Personal Information</h2>
          <div className="space-y-6">
            {[
              { label: "Full Name", value: "John Doe", id: "name" },
              { label: "Email Address", value: "john.doe@example.com", id: "email" },
              { label: "Phone Number", value: "+1 (555) 123-4567", id: "phone" },
            ].map((field) => (
              <div key={field.id}>
                <label className="block text-sm font-medium text-gray-300 mb-2">{field.label}</label>
                <input
                  defaultValue={field.value}
                  className="w-full rounded-2xl bg-gray-900/80 border border-gray-700 px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-primary-500"
                />
              </div>
            ))}
          </div>
        </div>

        <div className="bg-gray-800/60 backdrop-blur-sm rounded-3xl p-8">
          <h2 className="text-2xl font-semibold text-white mb-6">Notification Preferences</h2>
          <div className="space-y-4">
            <div className="flex items-center justify-between rounded-2xl bg-gray-900/80 p-4">
              <div>
                <h3 className="text-white font-medium">Email Notifications</h3>
                <p className="text-gray-400 text-sm">Receive updates about your exam progress.</p>
              </div>
              <button className="rounded-full bg-primary-600 px-4 py-2 text-sm font-semibold">Enabled</button>
            </div>
            <div className="flex items-center justify-between rounded-2xl bg-gray-900/80 p-4">
              <div>
                <h3 className="text-white font-medium">SMS Notifications</h3>
                <p className="text-gray-400 text-sm">Get reminders for upcoming exams.</p>
              </div>
              <button className="rounded-full bg-gray-700 px-4 py-2 text-sm font-semibold text-gray-300">Disabled</button>
            </div>
          </div>
        </div>

        <div className="flex justify-end">
          <Button variant="primary" className="px-6 py-3">
            Save Changes
          </Button>
        </div>
      </div>
    </div>
  );
}
