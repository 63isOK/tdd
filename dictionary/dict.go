package main

import "errors"

// error
var (
	ErrorNotFound      = errors.New("there is no test case")
	ErrorAlreadyExists = errors.New("already exists")
)

// Dictionary is dictionary
type Dictionary map[string]string

// Search is search from map
func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}

	return value, nil
}

// Add is add to map
func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; ok {
		return ErrorAlreadyExists
	}

	d[key] = value

	return nil
}

func main() {
}
