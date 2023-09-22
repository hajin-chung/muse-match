/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.html"],
  theme: {
    extend: {
      fontFamily: {
        serif: "Nanum Myeongjo",
        writing: "Satisfy",
      },
      colors: {
        yellow: {
          kakao: "#f9e000",
        },
      },
    },
  },
  darkMode: "class",
  plugins: [],
};
