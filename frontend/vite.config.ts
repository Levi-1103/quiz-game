import { svelte } from "@sveltejs/vite-plugin-svelte";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    tailwindcss(),
    svelte(),
  ],
  // eslint-disable-next-line node/prefer-global/process
  resolve: process.env.VITEST
    ? {
        conditions: ["browser"],
      }
    : undefined,
});
