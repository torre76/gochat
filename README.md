# gochat

Thi is a very simple chat server that listen on TCP port 10000 for clients.   
The chat protocol is very simple, clients connect with "telnet" and write single lines of text. On each new line of text, the server will broadcast that line to all other connected clients.

## Project Status

[![Build Status](https://travis-ci.org/torre76/gochat.svg?branch=master)](https://travis-ci.org/torre76/gochat)  [![Coverage Status](https://coveralls.io/repos/github/torre76/gochat/badge.svg?branch=master)](https://coveralls.io/github/torre76/gochat?branch=master)

## How to run this server

This server is written in GOLang, so it has to be installed on your system (at least Go version 1.12.x).

To install dependencies:

```bash
go get -u
```

To run this server

```bash
go run main.go
```

### Running on Docker

If you prefer using [Docker](https://www.docker.com) you can use the image [torre76/gochat:latest](https://hub.docker.com/r/torre76/gochat):

```bash
docker run -p 10000:10000 torre76/gochat:latest
```

if you want to use it on _foreground_ mode or:

```bash
docker run -d -p 10000:10000 torre76/gochat:latest
```

if you prefer a _daemon_ mode.
