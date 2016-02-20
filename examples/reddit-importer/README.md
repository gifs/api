# Reddit Importer

Import media found on one or more sub-reddits or user created multi-reddits.

![](https://j.gifs.com/kRoqvJ.gif)

### Usage:

```
# Get the clone of this repo
git clone git@github.com:gifs/api.git

# Enter the directory
cd api/examples/reddit-importer

# Import /r/reactiongifs
go run main.go /r/reactiongifs

# Import /r/reactiongifs and /r/politicalreactiongifs
go run main.go /r/reactiongifs+politicalreactiongifs

# Import a user created multi-reddit
go run main.go /user/bgny/m/50gif
```
