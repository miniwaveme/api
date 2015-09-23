#MiniWaveMe [![Build Status](https://travis-ci.org/miniwaveme/api.svg?branch=master)](https://travis-ci.org/miniwaveme/api)

##Introduction

This is a personnal Spotify/Deezer like, implemented in golang. Currently under development

Next steps:
- [api client](https://github.com/miniwaveme/go-miniwaveme-api)
- [web client (AngularJS)](https://github.com/miniwaveme/web-client)
- [desktop client](https://github.com/miniwaveme/desktop-client)

##Installation

This project is easily usable through docker:

```bash
$ make cmp-pull # update/download containers
$ make cmp-up   # run containers
$ make api-run  # run the api
$ make api-run-command COMMAND="script_name" #run a bash script of bin directory
$ make api-run-go-command COMMAND="golang_script_name" #run a golang script of bin directory

# Some other commands are available, type this to see them
$ make help
```

##Commands

###Available golang commands:

```create_master_key```: Create the first app key (a key is necessary to create an app key)

```
$ make api-run-go-command COMMAND="create_master_key"

app_key: 1fe50f00-fcc0-4aa5-9d76-35bf215e92ad | app_secret: 89c685e5-aa97-4dfc-8391-1a8f6aee3d9c

```

###Available bash commands:

*Not yet*


##Documentation

###API Documentation
Find the API documentation [here](http://miniwaveme.github.io/docs)

##Authors
- [mhor](http://github.com/mhor)
- [clebettre](http://github.com/clebettre)

##License

MiniWaveMe is released under the MIT License. See [LICENSE](LICENSE) file for details.
