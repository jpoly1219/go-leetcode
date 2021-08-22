module.exports = {
  mode: 'jit',
  purge: ['./src/**/*.svelte'],
  plugins: [require('@tailwindcss/typography')],
  theme: {
    extend: {
      gridTemplateRows: {
        '12': 'repeat(12, minmax(0, 1fr))',
      },
      gridRow: {
        'span-11': 'span 11 / span 11',
      }
    }
  }
}