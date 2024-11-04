// HomePage.js
import SpeedTest from "../components/SpeedTest";

const HomePage = () => {
  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center">
      <h1 className="text-4xl font-bold mb-6">Quicknet</h1>
      <SpeedTest />
    </div>
  );
};

export default HomePage;
