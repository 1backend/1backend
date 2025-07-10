// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

import type * as Preset from "@docusaurus/preset-classic";
import type { Config } from "@docusaurus/types";
import type * as Plugin from "@docusaurus/types/src/plugin";
import type * as OpenApiPlugin from "docusaurus-plugin-openapi-docs";

const config: Config = {
  title: "1Backend",
  tagline:
    "AI-native microservices platform.",
  url: "https://1backend.com",
  baseUrl: "/",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  favicon: "img/favicon.ico",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "1backend", // Usually your GitHub org/user name.
  projectName: "1backend", // Usually your repo name.

  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: require.resolve("./sidebars.ts"),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/1backend/1backend/tree/main/docs-source/",
          docItemComponent: "@theme/ApiItem", // Derived from docusaurus-theme-openapi
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/1backend/1backend/tree/main/docs-source/",
        },
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    docs: {
      sidebar: {
        hideable: true,
      },
    },
    metadata: [
      {
        name: "keywords",
        content:
          "ai, llm, free gpt, gpt, open-source, open source, ai framework, ai server",
      },
      {
        name: "title",
        content: "1Backend Documentation",
      },
      {
        name: "description",
        content: "1Backend API, Tutorials, Snippets and more",
      },
      { name: "twitter:card", content: "summary_large_image" },
    ],
    navbar: {
      title: "1Backend",
      logo: {
        alt: "1Backend Logo",
        src: "img/logo_circled_grey.svg",
      },
      items: [
        // {
        //   type: "doc",
        //   docId: "intro",
        //   position: "left",
        //   label: "Tutorial",
        // },
        {
          label: "Documentation",
          position: "left",
          to: "/docs/intro",
        },
        {
          label: "API",
          position: "left",
          to: "/docs/category/1backend-api",
        },
        {
          href: "https://github.com/1backend/1backend",
          label: "GitHub",
          position: "right",
        },
      ],
    },
    footer: {
      style: "dark",
      links: [
        {
          title: "Docs",
          items: [
            {
              label: "1backend.com",
              href: "https://1backend.com",
            },
            {
              label: "singulatron.com",
              href: "https://singulatron.com",
            },
          ],
        },
        {
          title: "Community",
          items: [
            {
              label: "Discord",
              href: "https://discordapp.com/enroll/eRXyzeXEvM",
            },
            {
              label: "Twitter",
              href: "https://twitter.com/singulatron",
            },
          ],
        },
        {
          title: "More",
          items: [
            {
              label: "GitHub",
              href: "https://github.com/1backend/1backend",
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} 1Backend, Inc. Built with Docusaurus.`,
    },
    prism: {
      prism: {
        additionalLanguages: [
          "ruby",
          "csharp",
          "php",
          "java",
          "powershell",
          "json",
          "bash",
        ],
      },
      languageTabs: [
        {
          highlight: "python",
          language: "python",
          logoClass: "python",
        },
        {
          highlight: "bash",
          language: "curl",
          logoClass: "bash",
        },
        {
          highlight: "csharp",
          language: "csharp",
          logoClass: "csharp",
        },
        {
          highlight: "go",
          language: "go",
          logoClass: "go",
        },
        {
          highlight: "javascript",
          language: "nodejs",
          logoClass: "nodejs",
        },
        {
          highlight: "ruby",
          language: "ruby",
          logoClass: "ruby",
        },
        {
          highlight: "php",
          language: "php",
          logoClass: "php",
        },
        {
          highlight: "java",
          language: "java",
          logoClass: "java",
          variant: "unirest",
        },
        {
          highlight: "powershell",
          language: "powershell",
          logoClass: "powershell",
        },
      ],
    },
  } satisfies Preset.ThemeConfig,

  plugins: [
    // "./smaller-diffs-webpack-plugin",
    [
      "docusaurus-plugin-openapi-docs",
      {
        id: "openapi",
        docsPluginId: "classic",
        config: {
          "1backend": {
            specPath: "examples/1backend.yaml",
            outputDir: "docs/1backend-api",
            downloadUrl:
              "https://raw.githubusercontent.com/1backend/1backend/main/server/docs/swagger.yaml",
            sidebarOptions: {
              groupPathsBy: "tag",
              categoryLinkSource: "tag",
            },
          } satisfies OpenApiPlugin.Options,
        } satisfies Plugin.PluginOptions,
      },
    ],
    require.resolve("docusaurus-lunr-search"),
  ],

  themes: ["docusaurus-theme-openapi-docs"],
};

export default async function createConfig() {
  return config;
}
