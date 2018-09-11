package trousseau

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type KV interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Store struct {
	Meta     *Meta     `json:"meta"`
	Items    []Item    `json:"items"`
	Sections []Section `json:"sections"`

	separator string
}

func Open(filename string) (*Store, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	store := NewStore()
	err = json.Unmarshal(b, &store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) WriteTo(filename string) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Set(key, value string) error {
	s.parseKey(key)
	return nil
}

func (s *Store) Get(key string) (string, error) {
	return "", nil
}

func (s *Store) Delete(key string) error {
	return nil
}

// func (s *Store) GetSection(key string) *Section {
// 	return nil
// }

func (s *Store) AddSection(key string) error {
	return nil
}

func (s *Store) DeleteSection(key string) error {
	return nil
}

func (s *Store) parseKey(key string) error {
	components := strings.Split(key, s.separator)
	fmt.Println(components)
	return nil
}

func NewStore() *Store {
	return &Store{
		Meta:      NewMetaWithDefaults(),
		Items:     []Item{},
		Sections:  []Section{},
		separator: "/",
	}
}
