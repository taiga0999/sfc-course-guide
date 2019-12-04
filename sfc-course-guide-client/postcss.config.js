const variables = {
  '--dark-mode-background': 'grey',
};


module.exports = {
  plugins: {
    'postcss-css-variables': { variables },
    precss: {},
    autoprefixer: {},
  },
};
