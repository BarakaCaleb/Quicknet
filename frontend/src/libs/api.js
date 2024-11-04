// api.js
import API_BASE_URL from "../config/apiConfig";

export const pingServer = async () => {
  const response = await fetch(`${API_BASE_URL}/ping`);
  return response.text();
};

export const downloadSpeedTest = async () => {
  const startTime = Date.now();
  await fetch(`${API_BASE_URL}/download`);
  return (Date.now() - startTime) / 1000;
};

export const uploadSpeedTest = async (file) => {
  const formData = new FormData();
  formData.append("file", file);

  const startTime = Date.now();
  await fetch(`${API_BASE_URL}/upload`, {
    method: "POST",
    body: formData,
  });

  return (Date.now() - startTime) / 1000;
};
