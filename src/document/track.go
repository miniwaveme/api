package document

type Track struct {
	Number   int      `json:"number" bson:"number"`
	Name     string   `json:"name" bson:"name"`
	Artists  []Artist `json:"artists" bson:"artists"`
	Duration int      `json:"duration" bson:"duration"`
	Bpm      int      `json:"bpm" bson:"bpm"`
	Path     string   `json:"path" bson:"path"`
	Url      string   `json:"url" bson:"url"`
}
