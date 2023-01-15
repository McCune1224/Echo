/** @type {import('tailwindcss').Config} */
module.exports = {
   
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: 
    {
      backgroundImage: {
        "authBackground": "url('https://i.ibb.co/y6fqVLK/103342526-p0.png')"
      }
    },
  },
  plugins: [],
}
