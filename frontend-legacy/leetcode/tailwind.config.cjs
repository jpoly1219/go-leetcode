module.exports = {
  mode: 'jit',
  purge: ['./src/**/*.svelte'],
  plugins: [require('@tailwindcss/typography')],
  theme: {
    extend: {
      gridTemplateRows: {
        '16': 'repeat(16, minmax(0, 1fr))',
      },
      gridRow: {
        'span-15': 'span 15 / span 15',
      }
    }
  }
}