/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
    "./pages/**/*.{js,ts,jsx,tsx}", // Optional if you're using a pages/ folder
  ],
  theme: {
    extend: {
      colors: {
        brand: '#f20574', // Example custom color
      },
      borderRadius: {
        'xl': '1rem', // You can customize rounding further here
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'), // Optional plugin for better form elements
    require('@tailwindcss/typography'), // Optional for blog/article content
  ],
  corePlugins: {
    preflight: false, // Disable base reset if you're handling it yourself
  },
  safelist: [
    'flex',
    'items-center',
    'gap-4',
    'bg-slate-100',
    'p-2',
    'rounded-xl',
    // Add more dynamically generated or runtime classes here
  ],
};
