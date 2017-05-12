package parser

type File struct {
	Path   []interface{} `json:"path"`
	Length int           `json:"length"`
}

type BitTorrent struct {
	InfoHash string `json:"infohash"`
	Name     string `json:"name"`
	Length   int    `json:"length,omitempty"`

	Files []File `json:"files,omitempty"`
}
