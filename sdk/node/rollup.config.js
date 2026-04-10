import resolve from "@rollup/plugin-node-resolve";

export default {
  input: "dist/index.js",
  output: [
    {
      file: "dist/index.mjs",
      format: "esm",
      exports: "named"
    },
    {
      file: "dist/index.cjs",
      format: "cjs",
      exports: "named"
    }
  ],
  plugins: [
    resolve()
  ],
  external: (id) => !id.startsWith(".") && !id.startsWith("/")
};
