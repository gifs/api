/*
 * node import.js $(cat urls.txt)
 * node import.js https://vine.co/v/iq6xX5YUdIt https://twitter.com/twitter/status/699999528619905024
 */

const gifs = require('gifs-api');
const spinner = require('cli-spinner').Spinner;

function main(argv) {
  if (argv.length < 1) {
    console.log('expecting atleast one URL to import from.');
    process.exit(-1);
  }

  argv.forEach(function(sourceURL) {
    var spinObj = new spinner(sourceURL);
    spinObj.setSpinnerString('▁▃▄▅▆▇█▇▆▅▄▃');
    spinObj.start();
    var params = {
      source: sourceURL,
      author: process.env.LOGNAME,
      tags: ['uploads', 'api-example'],
    };

    gifs.import(params, function(status, response) {
      spinObj.stop(true);
      console.log('\nurl', sourceURL, 'status', status, '\nresponse', response);
    });
  });
}

if (process.argv.length >= 2) {
  main(process.argv.slice(2));
}
