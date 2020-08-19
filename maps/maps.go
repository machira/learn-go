package maps

import "errors"

type Dictionary map[string]string

var NoSuchWordError = errors.New("no such word")

func (d Dictionary) Search(key string) (error, string) {
	s, ok := d[key]

	if !ok {
		return NoSuchWordError, ""
	}
	return nil, s
}

func (d Dictionary) Add(key string, value string) {
	d[key] = value
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
