package trousseau

type Item struct {
	Meta  Meta   `json:"meta"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
