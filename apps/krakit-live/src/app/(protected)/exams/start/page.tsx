"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function ExamStartPage() {
  const router = useRouter();

  useEffect(() => {
    const timer = setTimeout(() => {
      router.push("/exams");
    }, 1200);

    return () => clearTimeout(timer);
  }, [router]);

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 text-white px-4">
      <div className="max-w-md rounded-3xl bg-gray-900/90 border border-gray-700 p-10 text-center backdrop-blur-xl">
        <h1 className="text-3xl font-bold mb-4">Preparing your exam...</h1>
        <p className="text-gray-400">We are setting up your exam session. You will be redirected shortly.</p>
      </div>
    </div>
  );
}
