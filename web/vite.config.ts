import { fileURLToPath, URL } from "url";

import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
      "@app": fileURLToPath(new URL("./src/app", import.meta.url)),
      "@entities": fileURLToPath(new URL("./src/entities", import.meta.url)),
      "@features": fileURLToPath(new URL("./src/features", import.meta.url)),
      "@widgets": fileURLToPath(new URL("./src/widgets", import.meta.url)),
      "@pages": fileURLToPath(new URL("./src/pages", import.meta.url)),
      "@shared": fileURLToPath(new URL("./src/shared", import.meta.url)),
      "@assets": fileURLToPath(new URL("./src/assets", import.meta.url)),
      "@mocks": fileURLToPath(new URL("./src/shared/mocks", import.meta.url)),
    },
  },
  server: {
    watch: {
      usePolling: true,
    },
    host: true,
    strictPort: true,
    port: 3000,

    // This is for running the dev enviroment on the server
    allowedHosts: ["prog2900-lorawan.vm.iik.ntnu.no", "localhost", "127.0.0.1"],

    // caddy proxies the api route for doing internal http requests
    proxy: {
      "/api": {
        target: process.env.API_PROXY_TARGET || "http://api-gateway:8080",
        changeOrigin: true,
      },
    },
  },
});
