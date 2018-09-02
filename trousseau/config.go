package trousseau

type Config struct {
	Description string            `json:"description"`
	Annotations map[string]string `json:"annotations"`
}

func NewConfig(desc string, annotations map[string]string) *Config {
	return &Config{desc, annotations}
}
