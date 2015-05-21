package lib

import (
	"encoding/json"
	"fmt"
	// "strings"
	"io/ioutil"
	"os"
	"regexp"
)

type Config struct {
	_map map[string]interface{}
}

func (c Config) Get(item string) interface{} {
	return c._map[item]
}

func removeComments(bytes []byte) []byte {
	bytes = regexp.MustCompile("(//.*)").ReplaceAll(bytes, []byte{32})
	return bytes
}

func NewConfig(bytes []byte) Config {

	// Create a generic interface
	var f interface{}

	// Unmarshal the bytes of the JSON data into the generic interface
	json.Unmarshal(bytes, &f)

	// Create the object with a string map() of the unmarshalled data
	return Config{f.(map[string]interface{})}
}

func NewConfigFromFile(filename string) Config {
	file, _ := os.Open(filename)
	b, _ := ioutil.ReadAll(file)

	newb := removeComments(b)
	return NewConfig(newb)
}

func main_settings() {
	c := NewConfigFromFile("config.json")
	fmt.Println(c.Get("logger.enabled"))
	fmt.Println(c.Get("logger.debug"))
	fmt.Println(c.Get("users"))
}
