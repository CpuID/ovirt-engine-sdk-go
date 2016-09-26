package slice

import ()

// Credit - http://stackoverflow.com/questions/15323767/does-golang-have-if-x-in-construct-similar-to-python
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
