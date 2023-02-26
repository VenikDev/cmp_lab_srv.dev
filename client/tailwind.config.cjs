/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        "Eggplant": "#673C4F",
        "ChineseViolet": "#7F557D",
        "PompPower": "#726E97",
        "AirSuperiorityBlue": "#7698B3",
        "CarolinaBlue": "#83B5D1"
      }
    },
  },
  plugins: [],
}
