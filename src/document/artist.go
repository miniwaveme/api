package document

type Artist struct {
	Id     string  `json:"id" bson:"_id"`
	Name   string  `json:"name" bson:"name"`
	Images []Image `json:"images" bson:"images"`
}
