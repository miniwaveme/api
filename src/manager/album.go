package manager

import (
	"github.com/miniwaveme/api/src/db"
	"github.com/miniwaveme/api/src/document"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"image"
)

const AlbumCollection = "album"

func FindAlbumById(oid string) (*document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)
	album := &document.Album{}

	err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid)}).One(&album)

	return album, err
}

func FindAlbums(skip int, pageSize int) ([]document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)

	var albums = []document.Album{}
	err := col.Find(bson.M{}).Limit(pageSize).Skip(skip).Sort("name").All(&albums)

	return albums, err
}

func CreateAlbum(name string, year int, artist *document.Artist, image image.Image) (*document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)
	oid := bson.NewObjectId()
	album := &document.Album{Id: oid, Name: name, Year: year, ArtistRef: artist.Id, Artist: *artist}

	if image != nil {
		albumCover, err := CreateImage(image)
		if err != nil {
			return album, err
		}
		album.Cover = *albumCover
	}

	err := col.Insert(album)

	return album, err
}

func UpdateAlbum(oid string, name string, year int, artist *document.Artist, image image.Image) (*document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)
	album := &document.Album{}

	update := bson.M{}
	if name != "" {
		update["name"] = name
	}

	if year != 0 {
		update["year"] = year
	}

	if artist.Id != "" {
		update["artist"] = artist.Id
	}

	if image != nil {
		albumCover, err := CreateImage(image)
		if err != nil {
			return album, err
		}
		update["cover"] = *albumCover
	}

	change := mgo.Change{
		Update:    bson.M{"$set": update},
		ReturnNew: true,
	}

	_, err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid)}).Apply(change, &album)

	album.Artist = *artist

	return album, err
}

func RemoveAlbum(oid string) error {
	col := db.DS().DefaultDB().C(AlbumCollection)
	err := col.RemoveId(bson.ObjectIdHex(oid))

	return err
}

func AddAlbumTrack(oid string, number int, name string) (*document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)
	update := bson.M{"number": number, "name": name}

	album := &document.Album{}
	change := mgo.Change{
		Update:    bson.M{"$push": bson.M{"tracks": update}},
		ReturnNew: true,
	}

	_, err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid)}).Apply(change, &album)

	artist, err := FindArtistById(album.ArtistRef.Hex())
	album.Artist = *artist

	return album, err
}

func UpdateAlbumTrack(oid string, number int, name string) (*document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)

	album := &document.Album{}
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"tracks.$.name": name}},
		ReturnNew: true,
	}

	_, err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid), "tracks.number": number}).Apply(change, &album)

	artist, err := FindArtistById(album.ArtistRef.Hex())
	album.Artist = *artist

	return album, err
}

func RemoveAlbumTrack(oid string, trackNumber int) (*document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)
	album := &document.Album{}
	change := mgo.Change{
		Update:    bson.M{"$pull": bson.M{"tracks": bson.M{"number": trackNumber}}},
		ReturnNew: true,
	}

	_, err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid)}).Apply(change, &album)

	return album, err
}

func AddTrackArtist(oid string, number int, artist *document.Artist) (*document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)

	album := &document.Album{}
	change := mgo.Change{
		Update:    bson.M{"$push": bson.M{"tracks.$.artists": artist.Id}},
		ReturnNew: true,
	}

	_, err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid), "tracks.number": number}).Apply(change, &album)

	trackArtist, err := FindArtistById(album.ArtistRef.Hex())
	album.Artist = *trackArtist

	return album, err
}

func RemoveTrackArtist(oid string, trackNumber int, artist *document.Artist) (*document.Album, error) {
	col := db.DS().DefaultDB().C(AlbumCollection)

	album := &document.Album{}
	change := mgo.Change{
		Update:    bson.M{"$pull": bson.M{"tracks.$.artists": artist.Id}},
		ReturnNew: true,
	}

	_, err := col.Find(bson.M{"_id": bson.ObjectIdHex(oid), "tracks.number": trackNumber}).Apply(change, &album)
	trackArtist, err := FindArtistById(album.ArtistRef.Hex())
	album.Artist = *trackArtist

	return album, err
}
