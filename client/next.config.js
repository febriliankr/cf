/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: false,
  env: {
    firebaseApiKey: process.env.firebaseApiKey,
    API_URL: process.env.API_URL,
  },
  typescript: {
    // !! WARN !!
    // Dangerously allow production builds to successfully complete even if
    // your project has type errors.
    // !! WARN !!
    ignoreBuildErrors: true,
  },
};
