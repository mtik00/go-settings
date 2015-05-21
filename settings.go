// This package is used as an interface to various application settings stored
// in a JSON file or string.
package go_settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type Settings struct {
	_map map[string]interface{}
}

func (c Settings) Get(item string) interface{} {
	return c._map[item]
}

// removeComments parses a byte array and replaces all text following
// double-forward-slashes with a space character.
func removeComments(bytes []byte) []byte {
	bytes = regexp.MustCompile("(//.*)").ReplaceAll(bytes, []byte{32})
	return bytes
}

// Create a new Settings object from a byte array
func NewSettings(bytes []byte) Settings {

	// Create a generic interface
	var f interface{}

	new_bytes := removeComments(bytes)

	// Unmarshal the bytes of the JSON data into the generic interface
	json.Unmarshal(new_bytes, &f)

	// Create the object with a string map() of the unmarshalled data
	return Settings{f.(map[string]interface{})}
}

// Create a new Settings object from a JSON file
func NewSettingsFromFile(filename string) Settings {
	file, _ := os.Open(filename)
	b, _ := ioutil.ReadAll(file)

	new_bytes := removeComments(b)
	return NewSettings(new_bytes)
}
