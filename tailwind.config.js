/** @type {import('tailwindcss').Config} */

// "./templates/**/*.{html,js,templ}",
module.exports = {
  content: ["./build/**/*.html"],
  darkMode: "class",
  theme: {
    fontFamily: {
      poppins: ["Poppins", "sans-serif"],
      lobster: ["Lobster", "sans-serif"],
      "lobster-two": ['"Lobster Two"', "sans-serif"],
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
