package datastore

import ()

type Storage struct {
	Urls map[string]string
}

func NewStorage() *Storage {

	return &Storage{Urls: make(map[string]string, 0)}

}
