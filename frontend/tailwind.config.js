/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
    "./error.vue",
  ],
  theme: {
    extend: {
      colors:{
        'dark':'#1B1B1B',
        'primary':'#78E490',
        'secondary':'#69BEFF',
        'error':'#FF9E86',
        'light':'#E8F0EE',
        'half-dark':'#BCC8C4'
      }
    },
  },
  plugins: [],
}

