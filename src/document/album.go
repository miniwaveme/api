package document

type Album struct {
    Id         string       `json:"id"`
	Name       string       `json:"name"`
	Artists    []Artist     `json:"artists"`
    NbTrack    int          `json:"nb_track"`
    Duration   int          `json:"duration"`
    Cover      Image        `json:"cover"`
    Images     []Image      `json:"images"`
}