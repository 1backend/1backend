// DOES NOT WORK YET
// FIX THIS
//
// This custom webpack plugin is here so hundreds of assets won't be generated
// Before:
// Untracked files:
//   (use "git add <file>..." to include in what will be committed)
// 	docs/assets/js/0058b4c6.chunk.js
// 	docs/assets/js/00c16cfb.chunk.js
// 	docs/assets/js/01a85c17.chunk.js
// 	docs/assets/js/03713c66.chunk.js
// 	docs/assets/js/04dd0eaf.chunk.js
// 	docs/assets/js/05856483.chunk.js
// 	docs/assets/js/05db76fe.chunk.js
// 	docs/assets/js/0b7de69a.chunk.js
// 	docs/assets/js/0cc63ed5.chunk.js
//  ... dozens more
// After:
// Untracked files:
//   (use "git add <file>..." to include in what will be committed)
// 	docs/assets/js/main.12a9e92c.js
// 	docs/assets/js/main.12a9e92c.js.LICENSE.txt
// 	docs/lunr-index-1732692963242.json
// 	docs/search-doc-1732692963242.json
//
// There are still a lot of modified files probably because the main.js etc files are included in there.
// Still an improvement without breaking cache busting.
function smallerDiffsWebpackPlugin(context, options) {
  return {
    name: "smaller-diffs-webpack-plugin",
    configureWebpack(config, isServer) {
      if (!isServer) {
        return {
          optimization: {
            runtimeChunk: false,
            moduleIds: "named",
            splitChunks: {
              chunks: "all",
              cacheGroups: {
                default: false,
              },
            },
          },
          output: {
            filename: "assets/js/[name].js",
            chunkFilename: "assets/js/[name].js",
          },
        };
      }
      return {};
    },
  };
}

module.exports = smallerDiffsWebpackPlugin;
