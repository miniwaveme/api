package document

type Track struct {
	Id       string   `json:"id" bson:"_id"`
	Number   int      `json:"number" bson:"number"`
	Name     string   `json:"name" bson:"name"`
	Artists  []Artist `json:"artists" bson:"artists"`
	Duration int      `json:"duration" bson:"duration"`
	Bpm      int      `json:"bpm" bson:"bpm"`
	Album    Album    `json:"album" bson:"album"`
	Path     string   `json:"path" bson:"path"`
	Url      string   `json:"url" bson:"url"`
}
