package main

import "C"

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/stretchr/stew/objects"
	"gopkg.in/yaml.v3"
)

func mapLoadFromFile(AFileName string) (objects.Map, error) {
	filename, err := filepath.Abs(AFileName)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var unmarshalled map[string]interface{}
	if err := yaml.Unmarshal([]byte(data), &unmarshalled); err != nil {
		return nil, errors.New("Map: YAML decode failed with: " + err.Error())
	}
	return objects.Map(unmarshalled), nil
}

//export YAMLReadString
func YAMLReadString(AFileName *C.char, APath *C.char, ADefault *C.char) *C.char {
	m, err := mapLoadFromFile(C.GoString(AFileName))
	if err != nil {
		return ADefault
	}
	return C.CString(m.GetStringOrDefault(C.GoString(APath), C.GoString(ADefault)))
}

//export YAMLWriteString
func YAMLWriteString(AFileName *C.char, APath *C.char, AValue *C.char) C.int {
	m, err := mapLoadFromFile(C.GoString(AFileName))
	if err != nil {
		return 1
	}
	m.Set(C.GoString(APath), C.GoString(AValue))
	data, err := yaml.Marshal(m)
	if err != nil {
		return 1
	}
	if err = ioutil.WriteFile(C.GoString(AFileName), data, 0644); err != nil {
		return 1
	}
	return 0
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
