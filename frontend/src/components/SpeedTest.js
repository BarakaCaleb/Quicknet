// SpeedTest.js
import { useState } from "react";
import { pingServer, downloadSpeedTest, uploadSpeedTest } from "../libs/api";
import Button from "./Button";
import ResultCard from "./ResultCard";

const SpeedTest = () => {
  const [ping, setPing] = useState(null);
  const [downloadSpeed, setDownloadSpeed] = useState(null);
  const [uploadSpeed, setUploadSpeed] = useState(null);

  const handlePingTest = async () => {
    const result = await pingServer();
    setPing(result);
  };

  const handleDownloadTest = async () => {
    const timeTaken = await downloadSpeedTest();
    setDownloadSpeed((50 / timeTaken).toFixed(2));
  };

  const handleUploadTest = async () => {
    const file = new Blob(["a".repeat(5 * 1024 * 1024)]); // 5MB file
    const timeTaken = await uploadSpeedTest(file);
    setUploadSpeed((5 / timeTaken).toFixed(2));
  };

  return (
    <div className="space-y-6">
      <div className="flex space-x-4">
        <Button onClick={handlePingTest} label="Test Ping" />
        <Button onClick={handleDownloadTest} label="Test Download" />
        <Button onClick={handleUploadTest} label="Test Upload" />
      </div>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <ResultCard title="Ping" value={ping || "--"} unit="ms" />
        <ResultCard title="Download Speed" value={downloadSpeed || "--"} unit="Mbps" />
        <ResultCard title="Upload Speed" value={uploadSpeed || "--"} unit="Mbps" />
      </div>
    </div>
  );
};

export default SpeedTest;
