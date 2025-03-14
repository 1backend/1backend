import typescript from "rollup-plugin-typescript2";
import resolve from "@rollup/plugin-node-resolve";

export default {
  input: "src/api.ts", // Only use the main entry point
  output: [
    {
      dir: "dist",
      format: "esm", // Use ESM for simplicity
      entryFileNames: "[name].mjs", // Maintain the file extension
      chunkFileNames: "[name].mjs",
      exports: "named", // Export everything exactly as declared
    },
  ],
  plugins: [
    resolve(), // Ensures proper resolution of dependencies
    typescript({
      tsconfig: "./tsconfig.json", // Use existing tsconfig
    }),
  ],
  external: (id) => !id.startsWith(".") && !id.startsWith("/"), // Treat node_modules as external
};
