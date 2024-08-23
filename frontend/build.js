const esbuild = require("esbuild");
async function watch() {
  let ctx = await esbuild.context({
    entryPoints: ["./public/javascripts/main.js"],
    minify: false,
    outfile: "./public/javascripts/build.js",
    bundle: true,
    loader: { ".js": "js" },
  });
  await ctx.watch();
  console.log('Watching...');
}

// IMPORTANT: this call MUST NOT have an `await`.
watch();

// If the call above had an `await`, Node would return
// immediately and you would NOT have the watcher
// running. Alternative, you could use an iife[1]:
(async() => {
  // The same code from the `watch` function above.
  // Notice that it also doesn't have an `await` in
  // front of it.
})()
