import type { NextConfig } from "next";

const env = {
  API_URL: process.env.API_URL
}

const nextConfig: NextConfig = {
  /* config options here */
  env: env,
  eslint: {
    ignoreDuringBuilds: true,
  },
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'picsum.photos',
      },
      {
        protocol: 'https',
        hostname: 'user-gin-next.s3.us-east-1.amazonaws.com'
      }
    ]
  }
};

export default nextConfig;
