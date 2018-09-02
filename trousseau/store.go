package trousseau

type Store struct {
	Meta     Meta    `json:"meta"`
	Config   Config  `json:"config"`
	Sections Section `json:"sections"`
}
