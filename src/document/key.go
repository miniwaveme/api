package document

type Key struct {
    Id          string      `json:"id"`
    AppKey      string      `json:"app_key"`
    AppSecret   string      `json:"app_secret"`
    Roles       []string    `json:"roles"`
}