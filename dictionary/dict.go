package main

import "errors"

// Dictionary is dictionary
type Dictionary map[string]string

// Search is search from map
func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", errors.New("there is no test case")
	}

	return value, nil
}

func main() {
}
