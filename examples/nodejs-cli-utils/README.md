## nodejs-cli-importer

- Sample programs for using the gifs.com api in node.js programs, for importing and uploading.

## Installing the necessary dependencies:
```shell
$ npm update
```

## Running the samples:
+ import.js
Allows you add media by URL e.g

```shell
$ node import.js $(cat urls.txt)
$ node import.js https://vine.co/v/iq6xX5YUdIt https://twitter.com/twitter/status/699999528619905024 https://www.youtube.com/watch?v=7ByZqiwwWwA
```

+ upload.js
Allows you to add media that is present on your local file system e.g

```shell
$ node upload.js *.gif
$ node upload.js ~/Desktop/tricky.gif ~/Pictures/Videos/*.mp4 ~/Videos/2014Highlights.mp4
```
