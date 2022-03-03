package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("word already exists")
var ErrWordDoesNotExist = errors.New("word does not exist")

func (d Dictionary) Search(word string) (string, error) {
	data, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return data, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	if err == nil {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	if err != nil {
		return ErrWordDoesNotExist
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
