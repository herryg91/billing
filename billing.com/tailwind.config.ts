import type { Config } from "tailwindcss";
import daisyui from 'daisyui';

const config: Config = {
  darkMode: ['class', '[data-mode="light"]'],
  mode: 'jit',
  content: [
    "./src/**/*.{js,ts,jsx,tsx}",
    "./src/pages/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
    "./public/**/*.html",
    'node_modules/daisyui/dist/**/*.js', 
    'node_modules/react-daisyui/dist/**/*.js',
  ],
  theme: { 
    extend: { 
        fontFamily: { 
          "default": ['var(--font-inter)', "ui-sans-serif", "system-ui", "sans-serif", "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"],
          "mono": ['var(--font-mono)']
        } 
    }, 
  }, 
  daisyui: {
    themes: ['bumblebee'],
  },
  plugins: [daisyui],
};
export default config;
