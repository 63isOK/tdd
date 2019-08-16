package main

// Dictionary is dictionary
type Dictionary map[string]string

// Search is search from map
func (d Dictionary) Search(key string) string {
	return d[key]
}

func main() {
}
