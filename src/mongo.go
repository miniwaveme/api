package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/miniwaveme/api/src/document"
	"gopkg.in/mgo.v2"
)

// theses vars should be created dynamically
var (
	dbUrl  = "127.0.0.1:27017"
	dbName = "miniwaveme"
)

func main() {
	// create session
	session, err := mgo.Dial(dbUrl)

	// should return an error if err != nil
	if err != nil {
		log.Error("can't establish the connection to the database")
		return
	}
	defer session.Close()

	// Collection 'track'
	tracks := session.DB(dbName).C("track")

	// insert a new track
	track := document.Track{Number: 1, Name: "Nice track !", Duration: 360}
	err = tracks.Insert(track)

	if err != nil {
		log.Error("can't insert document")
		return
	}

	// find all tracks
	var findTracks []document.Track
	err = tracks.Find(track).All(&findTracks)

	// print result
	fmt.Println(findTracks)
}
