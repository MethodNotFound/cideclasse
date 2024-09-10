const esbuild = require("esbuild");
const postCssPlugin = require('esbuild-style-plugin')

let config = {
  entryPoints: ["./src/javascript/main.js", "./src/style/style.css"],
  outdir: 'public',
  bundle: true,
  minify: false,
  plugins: [
    postCssPlugin({
      postcss: {
        plugins: [require('tailwindcss'), require('autoprefixer')],
      },
    }),
  ]
};

async function watch() {
  let ctx = await esbuild.context(config);
  await ctx.watch();
  console.log('Watching...');
}

if(!process.env.NOT_WATCH) {
  watch();
} else {
  esbuild.build(config)
}


