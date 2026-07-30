import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";
import { HeroSection } from "@/components/home/HeroSection";
import { FeaturesSection } from "@/components/home/FeaturesSection";
import { PrimaryContent } from "@/components/home/PrimaryContent";
import { SecondaryContent } from "@/components/home/SecondaryContent";
import { CallToAction } from "@/components/home/CallToAction";

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-primary-900 to-gray-900 text-white">
      <Header />
      <main className="pt-24">
        <HeroSection />
        <FeaturesSection />
        <PrimaryContent />
        <SecondaryContent />
        <CallToAction />
      </main>
      <Footer />
    </div>
  );
}
