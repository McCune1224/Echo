/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: 
        {
            backgroundImage: {
              "authBackground": "url('https://i.ibb.co/y6fqVLK/103342526-p0.png')"
            }
        },
        dropShadow: {
            "shine-empty": '0 0px 0px rgba(0,0,0,0)',
            'shine-grey': '0 0px 5px rgba(185, 185, 185, 1)',
            'shine-white': '0 0px 5px rgba(255, 255, 255, 1)',
        }
    },
    plugins: [],
}
