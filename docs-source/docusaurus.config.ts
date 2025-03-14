// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

import type * as Preset from "@docusaurus/preset-classic";
import type { Config } from "@docusaurus/types";
import type * as Plugin from "@docusaurus/types/src/plugin";
import type * as OpenApiPlugin from "docusaurus-plugin-openapi-docs";

const config: Config = {
  title: "OpenOrch",
  tagline:
    "A language-agnostic microservices framework for building AI applications.",
  url: "https://openorch.org",
  baseUrl: "/",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  favicon: "img/favicon.ico",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "openorch", // Usually your GitHub org/user name.
  projectName: "openorch", // Usually your repo name.

  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: require.resolve("./sidebars.ts"),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/openorch/openorch/tree/main/docs-source/",
          docItemComponent: "@theme/ApiItem", // Derived from docusaurus-theme-openapi
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/openorch/openorch/tree/main/docs-source/",
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
        content: "OpenOrch Documentation",
      },
      {
        name: "description",
        content: "OpenOrch API, Tutorials, Snippets and more",
      },
      { name: "twitter:card", content: "summary_large_image" },
    ],
    navbar: {
      title: "OpenOrch",
      logo: {
        alt: "OpenOrch Logo",
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
          to: "/docs/category/openorch-api",
        },
        {
          href: "https://github.com/openorch/openorch",
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
              label: "openorch.org",
              href: "https://openorch.org",
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
              href: "https://discordapp.com/invite/eRXyzeXEvM",
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
              href: "https://github.com/openorch/openorch",
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} OpenOrch, Inc. Built with Docusaurus.`,
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
          openorch: {
            specPath: "examples/openorch.yaml",
            outputDir: "docs/openorch",
            downloadUrl:
              "https://raw.githubusercontent.com/openorch/openorch/main/server/docs/swagger.yaml",
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
