package document

type Image struct {
	Path string `json:"path" bson:"path"`
	Url  string `json:"url" bson:"url"`
}
