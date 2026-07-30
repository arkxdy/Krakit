import Link from "next/link";
import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";
import { Button } from "@/components/ui/Button";

export const metadata = {
  title: "Pricing - Krakit",
  description: "Choose the plan that fits your exam preparation needs.",
};

const plans = [
  {
    title: "Basic",
    price: "Free",
    description: "Perfect for getting started",
    features: ["Basic study materials", "Limited practice tests", "Community support"],
    href: "/signup",
    variant: "outline",
  },
  {
    title: "Pro",
    price: "$9.99/mo",
    description: "For serious learners",
    features: ["Unlimited practice tests", "Advanced analytics", "Priority support"],
    href: "/signup?plan=pro",
    variant: "primary",
  },
  {
    title: "Premium",
    price: "$19.99/mo",
    description: "Coming soon",
    features: ["Everything in Pro", "Dedicated tutor support", "Premium resources"],
    href: "/signup?plan=premium",
    variant: "outline",
  },
];

export default function PricingPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 flex flex-col text-white">
      <Header />
      <main className="flex-1 container mx-auto px-4 py-20 mt-16">
        <div className="text-center mb-12">
          <h1 className="text-4xl font-bold">Choose Your Plan</h1>
          <p className="mt-4 text-gray-300">
            Select the perfect plan for your learning journey.
          </p>
        </div>

        <div className="grid gap-6 md:grid-cols-3">
          {plans.map((plan) => (
            <div
              key={plan.title}
              className="border border-purple-500/20 rounded-xl shadow-sm bg-gray-900"
            >
              <div className="p-8">
                <h2 className="text-2xl font-semibold text-purple-300">
                  {plan.title}
                </h2>
                <p className="text-gray-400 mt-4">{plan.description}</p>
                <p className="text-4xl font-extrabold text-purple-400 mt-8">
                  {plan.price}
                </p>
                <div className="mt-8 space-y-3">
                  {plan.features.map((feature) => (
                    <div key={feature} className="flex items-center gap-3">
                      <span className="text-purple-400">•</span>
                      <span className="text-gray-300">{feature}</span>
                    </div>
                  ))}
                </div>
                <div className="mt-8">
                  <Link href={plan.href} className="block">
                    <Button variant={plan.variant as "primary" | "outline"} className="w-full">
                      {plan.title === "Premium" ? "Notify me" : "Get Started"}
                    </Button>
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      </main>
      <Footer />
    </div>
  );
}
