package document

type Image struct {
	Id   string `json:"id" bson:"_id"`
	Path string `json:"path" bson:"path"`
	Url  string `json:"url" bson:"url"`
}
