const benefits = [
  {
    title: "Personalized Learning",
    description:
      "Get recommendations based on your performance and learning style.",
    icon: "🎯",
  },
  {
    title: "Detailed Analytics",
    description:
      "Understand your strengths and weaknesses with comprehensive analytics.",
    icon: "📊",
  },
  {
    title: "Mobile Friendly",
    description: "Study on the go with our mobile-responsive platform.",
    icon: "📱",
  },
  {
    title: "Regular Updates",
    description:
      "Access new content and features with regular platform updates.",
    icon: "🔄",
  },
];

export function SecondaryContent() {
  return (
    <section className="py-20 bg-gray-900">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-16">
          <h2 className="text-3xl font-bold text-white mb-4">
            Additional Benefits
          </h2>
          <p className="text-gray-400 max-w-2xl mx-auto">
            Discover more ways our platform can help you achieve your goals
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {benefits.map((benefit) => (
            <div
              key={benefit.title}
              className="bg-gray-800/50 backdrop-blur-sm rounded-xl p-6"
            >
              <span className="text-3xl mb-4 block">{benefit.icon}</span>
              <h3 className="text-lg font-semibold text-white mb-2">
                {benefit.title}
              </h3>
              <p className="text-gray-400">{benefit.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
