import { redirect } from "next/navigation";
import { getCurrentUser } from "@/lib/auth";
import { PrivateHeader } from "@/components/layout/PrivateHeader";
import { Sidebar } from "@/components/Sidebar";
import { Footer } from "@/components/layout/Footer";
import { SidebarProvider } from "@/components/SidebarProvider";

export default async function ProtectedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const user = await getCurrentUser();

  if (!user) {
    redirect("/login");
  }

  return (
    <SidebarProvider>
      <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 text-white">
        <PrivateHeader user={user!} />
        <Sidebar />
        <div className="transition-all duration-300 ease-in-out pl-0 md:pl-64 pt-24">
          <main className="container mx-auto px-4 py-6">{children}</main>
        </div>
        <Footer />
      </div>
    </SidebarProvider>
  );
}
