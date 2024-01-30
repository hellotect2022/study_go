package mydict

import "errors"

// type 에는 method 를 추가할 수 있음
type Dictionary map[string]string

var errNotFound = errors.New("Not found")
var errCantUpdate = errors.New("Cant update non-existing")
var errCantDelete = errors.New("Cant delete non-existing")
var errWordExists = errors.New("Key alredy Exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err == errNotFound {
		d[key] = value
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[key] = value
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		return errCantDelete
	case nil:
		delete(d, key)
	}
	return nil
}
