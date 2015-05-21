package document

type Album struct {
	Id       string   `json:"id" bson:"_id"`
	Name     string   `json:"name" bson:"name"`
	Artists  []Artist `json:"artists" bson:"artists"`
	NbTrack  int      `json:"nb_track" bson:"nb_track"`
	Duration int      `json:"duration" bson:"duration"`
	Cover    Image    `json:"cover" bson:"cover"`
	Images   []Image  `json:"images" bson:"images"`
}
