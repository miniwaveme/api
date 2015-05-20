package document

type Key struct {
	Id        string   `json:"id" bson:"_id"`
	AppKey    string   `json:"app_key" bson:"app_key"`
	AppSecret string   `json:"app_secret" bson:"app_secret"`
	Roles     []string `json:"roles" bson:"roles"`
}
