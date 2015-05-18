package document

type Track struct {
    Id          string      `json:"id"`
    Number      int         `json:"number"`
    Name        string      `json:"name"`
    Artists     []Artist    `json:"artists"`
    Duration    int         `json:"duration"`
    Bpm         int         `json:"bpm"`
    Album       Album       `json:"album"`
    Path        string      `json:"path"`
    Url         string      `json:"url"`
}