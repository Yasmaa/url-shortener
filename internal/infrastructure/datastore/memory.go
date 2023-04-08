package datastore

import ()

type Storage struct {
	urls map[string]string
}

func NewStorage() *Storage {

	return &Storage{urls: make(map[string]string, 0)}

}
