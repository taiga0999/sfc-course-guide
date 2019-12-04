const express = require('express');
const history = require('connect-history-api-fallback');
const path = require('path');

express()
  .use(history())
  .use(express.static(path.resolve(`${__dirname}/../dist`)))
  .get('/', (_, res) => res.sendFile('index.html'))
  .listen(3000, () => console.log('Example app listening on port 3000!'));
