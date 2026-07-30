import { redirect } from "next/navigation";

export default function RegisterPage() {
  redirect("/signup");

  return (
    <div className="min-h-screen bg-gray-950 flex items-center justify-center text-white px-4">
      <div className="max-w-md bg-gray-900/80 rounded-3xl border border-gray-700 p-10 text-center">
        <h1 className="text-3xl font-bold mb-4">Redirecting...</h1>
        <p className="text-gray-400">Taking you to the signup page now.</p>
      </div>
    </div>
  );
}
