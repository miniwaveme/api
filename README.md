#MiniWaveMe [![Build Status](https://travis-ci.org/miniwaveme/api.svg?branch=master)](https://travis-ci.org/miniwaveme/api)

##Introduction

This is a personnal Spotify/Deezer like, implemented in golang. Currently under development

Next steps:
- [web client (AngularJS)](https://github.com/miniwaveme/web-client)
- [desktop client](https://github.com/miniwaveme/desktop-client)

##Installation

This project is easily usable through docker:

```bash
$ make cmp-pull # update/download containers
$ make cmp-up   # run containers
$ make api-run  # run the api

# Some other commands are available, type this to see them
$ make help
```

##Documentation

###API Documentation
Find the API documentation [here](http://miniwaveme.github.io/docs)

##Authors
- [mhor](http://github.com/mhor)
- [clebettre](http://github.com/clebettre)

##License

MiniWaveMe is released under the MIT License. See [LICENSE](LICENSE) file for details.
