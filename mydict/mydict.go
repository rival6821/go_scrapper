package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("NOT FOUND")
var errWordExists = errors.New("THAT WORD ALREADY EXIST")
var errCantUpdate = errors.New("CANT UPDATE")
var errCantDelete = errors.New("CANT DELETE")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add for a word
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}

// Update for a word
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete for a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
