import { Button } from "@/components/ui/Button";

export const metadata = {
  title: "Study Materials - Krakit",
  description: "Access comprehensive study materials and resources.",
};

const subjects = [
  {
    name: "Mathematics",
    materials: [
      { title: "Calculus Fundamentals", description: "Comprehensive guide to differential and integral calculus." },
      { title: "Algebra Practice Problems", description: "Set of 100 practice problems with solutions." },
    ],
  },
  {
    name: "Physics",
    materials: [
      { title: "Mechanics Study Guide", description: "Detailed notes on classical mechanics and motion." },
      { title: "Electromagnetism Notes", description: "Comprehensive coverage of electromagnetic theory." },
    ],
  },
  {
    name: "Chemistry",
    materials: [
      { title: "Organic Chemistry Guide", description: "Complete guide to organic compounds and reactions." },
      { title: "Chemical Bonding Notes", description: "Detailed explanation of chemical bonding types." },
    ],
  },
];

export default function MaterialsPage() {
  return (
    <div className="max-w-7xl mx-auto space-y-10">
      <div className="text-white">
        <h1 className="text-4xl font-bold mb-3">Study Materials</h1>
        <p className="text-gray-400">Explore curated resources and download study content for all major subjects.</p>
      </div>

      <div className="grid gap-6 lg:grid-cols-3">
        {subjects.map((subject) => (
          <div key={subject.name} className="bg-gray-800/60 backdrop-blur-sm rounded-3xl p-8">
            <h2 className="text-2xl font-semibold text-white mb-4">{subject.name}</h2>
            <div className="space-y-4">
              {subject.materials.map((material) => (
                <div key={material.title} className="rounded-2xl bg-gray-900/80 p-4">
                  <h3 className="text-lg font-semibold text-white">{material.title}</h3>
                  <p className="text-gray-400 mt-2">{material.description}</p>
                  <div className="mt-4">
                    <Button variant="outline" className="px-4 py-2 text-sm">
                      Download PDF
                    </Button>
                  </div>
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
