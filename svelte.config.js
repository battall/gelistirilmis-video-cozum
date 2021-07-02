import adapter from "@sveltejs/adapter-static";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    // hydrate the <div id="app"> element in src/app.html
    target: "#app",
    adapter: adapter(),
  },
};

export default config;
