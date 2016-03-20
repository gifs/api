/*
 * node upload.js *.gif
 * node upload.js ~/Desktop/tricky.gif ~/Pictures/Videos/*.mp4 ~/Desktop/boondocksAirPlane.mp4
 */
const gifs = require('gifs-api');
const fs   = require('fs');
const path = require('path');
const spinner = require('cli-spinner').Spinner;

function main(argv) {
  if (argv.length < 1) {
    console.log('expecting atleast one path to upload.');
    process.exit(-1);
  }

  argv.forEach(function(arg) {
    var readStream = fs.createReadStream(arg);
    var spinObj = new spinner(arg);
    spinObj.setSpinnerString('┤┘┴└├┌┬┐');
    spinObj.start();
    var params = {
      file: readStream,
      title: path.basename(arg),
      tags: ["uploaded file", "uploaded example"],
    };

    gifs.upload(params, function(status, response) {
	spinObj.stop(true);
	console.log('\nfilepath', arg, 'status', status, '\nresponse', response);
	readStream.close();
    });
  });
}


if (process.argv.length >= 2) {
  main(process.argv.slice(2));
}
