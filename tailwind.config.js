/** @type {import('tailwindcss').Config} */

// "./templates/**/*.{html,js,templ}",
module.exports = {
  content: ["./build/**/*.html"],
  darkMode: "class",
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography")],
};
