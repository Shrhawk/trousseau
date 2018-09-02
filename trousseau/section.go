package trousseau

type Section struct {
	Items    Item     `json:"items"`
	Sections *Section `json:"sections"`
}

func NewSection(items *Item, sections *Section) *Section {
	return &Section{*items, sections}
}
