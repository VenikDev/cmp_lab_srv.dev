module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'main-border': '#000000',
        'primary': '#0097b3',
        'main-text': '#000000',
        'soft-peach': '#ffe4d0',
        'pastel-sky': '#DEF7FE',
        'lilac': '#d9e6ff'
      },
    },
    screens: {
      'sm': '640px',
      'md': '768px',
      'lg': '1024px',
      'xl': '1024px',
      '2xl': '1024px',
    }
  },
  plugins: [],
}
