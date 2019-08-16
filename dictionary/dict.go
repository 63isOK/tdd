package main

import "errors"

// error
var (
	ErrorNotFound = errors.New("there is no test case")
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

func main() {
}
