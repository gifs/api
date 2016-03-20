/*
 * node upload.js *.gif
 * node upload.js ~/Desktop/tricky.gif ~/Pictures/Videos/*.mp4 ~/Desktop/boondocksAirPlane.mp4
 */
const gifs = require('gifs-api');
const fs   = require('fs');

function main(argv) {
  if (argv.length < 1) {
    console.log('expecting atleast one path to upload.');
    process.exit(-1);
  }

  argv.forEach(function(arg) {
    var readStream = fs.createReadStream(arg);
    var params = {
      file: readStream,
      title: arg,
      tags: ["uploaded file", "uploaded example"],
    };

    gifs.upload(params, function(status, response) {
	console.log('filepath', arg, 'status', status, response);
    });
  });
}


if (process.argv.length >= 2) {
  main(process.argv.slice(2));
}
