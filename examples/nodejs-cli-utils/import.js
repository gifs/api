/*
 * node import.js $(cat urls.txt)
 * node import.js https://vine.co/v/iq6xX5YUdIt https://twitter.com/twitter/status/699999528619905024
 */

const gifs = require('gifs-api');

function main(argv) {
  if (argv.length < 1) {
    console.log('expecting atleast one URL to import from.');
    process.exit(-1);
  }

  argv.forEach(function(sourceURL) {
    var params = {
      source: sourceURL,
      author: process.env.LOGNAME,
      tags: ['uploads', 'api-example'],
    };

    gifs.import(params, function(status, response) {
      console.log('url', sourceURL, 'status', status, 'response', response);
    });
  });
}

if (process.argv.length >= 2) {
  main(process.argv.slice(2));
}
