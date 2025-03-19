import type { NextConfig } from "next";

const env = {
  API_URL: process.env.API_URL
}

const nextConfig: NextConfig = {
  /* config options here */
  env: env
};

export default nextConfig;
