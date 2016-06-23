package document

type Image struct {
	Original string `json:"original" bson:"original"`
	Small    string `json:"small" bson:"small"`
	Medium   string `json:"medium" bson:"medium"`
	Large    string `json:"large" bson:"large"`
}
