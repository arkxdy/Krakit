import Link from "next/link";

const mainFeatures = [
  {
    title: "Mock Exams",
    description:
      "Practice with our extensive collection of mock exams designed to simulate real exam conditions.",
    icon: "📝",
    link: "/exams",
  },
  {
    title: "Study Materials",
    description:
      "Access comprehensive study materials and resources to enhance your preparation.",
    icon: "📚",
    link: "/materials",
  },
  {
    title: "Progress Tracking",
    description:
      "Monitor your performance and track your improvement over time with detailed analytics.",
    icon: "📊",
    link: "/progress",
  },
];

export function PrimaryContent() {
  return (
    <section className="py-20 bg-gray-800">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-16">
          <h2 className="text-3xl font-bold text-white mb-4">
            Everything You Need to Succeed
          </h2>
          <p className="text-gray-400 max-w-2xl mx-auto">
            Our platform provides all the tools and resources you need to excel
            in your exams
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {mainFeatures.map((feature) => (
            <div
              key={feature.title}
              className="bg-gray-900 rounded-xl p-6 hover:bg-gray-800 transition-colors"
            >
              <span className="text-4xl mb-4 block">{feature.icon}</span>
              <h3 className="text-xl font-semibold text-white mb-2">
                {feature.title}
              </h3>
              <p className="text-gray-400 mb-4">{feature.description}</p>
              <Link
                href={feature.link}
                className="text-primary-400 hover:text-primary-300 font-medium"
              >
                Learn more →
              </Link>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
