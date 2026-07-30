import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export default async function LogoutPage() {
  const cookiesStore = await cookies();
  cookiesStore.delete({ name: "token", path: "/" });
  redirect("/login");

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex items-center justify-center text-white px-4">
      <div className="max-w-md text-center bg-gray-900/70 backdrop-blur-sm rounded-2xl p-10 border border-gray-700">
        <h1 className="text-3xl font-bold mb-4">Signing out...</h1>
        <p className="text-gray-400">You are being redirected to the login page.</p>
      </div>
    </div>
  );
}
