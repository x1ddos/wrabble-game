# Wrabble game

This is a [Scrabble](http://en.wikipedia.org/wiki/Scrabble)-like game
written in Polymer (frontend) and Go (backend) for App Engine.

Give it a go at [wrabble-game.appspot.com](https://wrabble-game.appspot.com)


## Dev setup

Prerequisites

* [Bower](http://bower.io/)
* [Go](http://golang.org/doc/install)
* [App Engine for Go SDK](https://cloud.google.com/appengine/downloads)

Setup

1. Install frontend components: `cd main && bower install`
2. Install backend Go package deps: `cd .. && goapp get ./wrabble/...`
3. Start the dev app server with `goapp serve main/app.yaml admin/admin.yaml dispatch.yaml`

### Bootstrap

You'll need a dictionary for words verification. This can be any text file 
with one word per line.

1. Navigate to [localhost:8080/admin/import](http://localhost:8080/admin/import)
2. Set dict name to `wordlist`
3. Choose a dict file to upload & hit "Submit".

The admin module will start processing uploaded dictionary file.
It might take a while.

Navigate to [localhost:8080](http://localhost:8080) either from desktop 
or mobile and start playing.
