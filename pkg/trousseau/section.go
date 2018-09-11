package trousseau

type Section struct {
	Name     string    `json:"name"`
	Meta     Meta      `json:"meta"`
	Items    []Item    `json:"items"`
	Sections []Section `json:"sections"`
}

func (s *Section) GetChild(name string) (*Section, error) {
	for _, s := range s.Sections {
		if s.Name == name {
			return &s, nil
		}
	}

	return nil, ErrSectionNotFound
}

// func HasChild(name string) bool {
// 	for _, s := range s.Sections {
// 		if s.Name == name {
// 			return true
// 		}
// 	}

// 	return false
// }

// func AddChild(section *Section) error {
// 	if HasChild(section) {
// 		return ErrSectionAlreadyExists
// 	}

// 	s.Sections = append(s.Sections, section)
// 	return nil
// }

// func DeleteChild(name string) error {
// 	for idx, s := range s.Sections {
// 		if s.Name == name {
// 			s.Sections = append(s.Sections[:idx], s.Sections[idx+1:]...) // Neat in place delete
// 			return nil
// 		}
// 	}

// 	return ErrSectionNotFound
// }

// func (s *Section) GetItem(key string) (*Item, error) {
// 	for _, i := range s.Items {
// 		if i.Key == key {
// 			return &i, nil
// 		}
// 	}

// 	return nil, ErrItemNotFound
// }

// func (s *Section) HasItem(key string) bool {
// 	for _, it := range s.Items {
// 		if it.Name == key {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (s *Section) AddItem(item Item) error {
// 	if HasItem(item) {
// 		return ErrItemAlreadyExists
// 	}

// 	s.Items = append(s.Items, item)
// 	return nil
// }

// func (s *Section) DeleteItem(name string) error {
// 	for idx, s := range s.Items {
// 		if s.Name == name {
// 			s.Items = append(s.Items[:idx], s.Items[idx+1:]...) // Neat in place delete
// 			return nil
// 		}
// 	}

// 	return ErrSectionNotFound
// }
