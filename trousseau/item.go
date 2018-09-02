package trousseau

type Item struct {
	Config Config            `json:"config"`
	Meta   Meta              `json:"meta"`
	Data   map[string]string `json:"data"`
}

func NewItem(config Config, meta Meta, data map[string]string) *Item {
	return &Item{config, meta, data}
}
