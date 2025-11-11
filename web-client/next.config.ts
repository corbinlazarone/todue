import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
};

const requiredEnvVars = ["NEXT_PUBLIC_API_URL"];

requiredEnvVars.forEach((envVar) => {
  if (!process.env[envVar]) {
    throw new Error(`Missing env var: ${envVar}`);
  }
});

export default nextConfig;
