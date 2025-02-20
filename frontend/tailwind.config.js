/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: '#1DA1F2',
        'gray': {
          800: '#1E2732',
          900: '#15202B',
        },
      },
    },
  },
  plugins: [],
}
