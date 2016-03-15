package manager

import (
	"github.com/miniwaveme/api/src/db"
	"github.com/miniwaveme/api/src/document"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ArtistCollection = "artist"

func FindArtistById(oid string) (*document.Artist, error) {
	col := db.DS().DefaultDB().C(ArtistCollection)
	artist := &document.Artist{}

	err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid)}).One(&artist)

	return artist, err
}

func FindArtists(skip int, pageSize int) ([]document.Artist, error) {
	col := db.DS().DefaultDB().C(ArtistCollection)

	var artists = []document.Artist{}
	err := col.Find(bson.M{}).Limit(pageSize).Skip(skip).Sort("name").All(&artists)

	return artists, err
}

func CreateArtist(name string) (*document.Artist, error) {
	col := db.DS().DefaultDB().C(ArtistCollection)
	oid := bson.NewObjectId()
	artist := &document.Artist{Id: oid, Name: name}
	err := col.Insert(artist)

	return artist, err
}

func UpdateArtist(oid string, name string) (*document.Artist, error) {
	col := db.DS().DefaultDB().C(ArtistCollection)
	artist := &document.Artist{}

	update := bson.M{}
	if name != "" {
		update["name"] = name
	}

	change := mgo.Change{
		Update:    bson.M{"$set": update},
		ReturnNew: true,
	}

	_, err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid)}).Apply(change, &artist)

	return artist, err
}

func RemoveArtist(oid string) error {
	col := db.DS().DefaultDB().C(ArtistCollection)
	err := col.RemoveId(bson.ObjectIdHex(oid))

	return err
}
