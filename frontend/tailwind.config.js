module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true,
  },
  purge: [
    './app/**/*.tsx',
    './app/**/*.ts',
    './app/**/*.css',
    './templates/**/*.html',
  ],
  theme: {
    extend: {},
  },
  variants: {},
  plugins: [],
}
