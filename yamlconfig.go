package main

import "C"
import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/stretchr/stew/objects"
	"gopkg.in/yaml.v3"
)

func mapLoadFromFile(AFileName string) (objects.Map, error) { //objects.Map, error) {
	filename, err := filepath.Abs(AFileName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Filename: " + filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println("File readed")
	var unmarshalled map[string]interface{}
	if err := yaml.Unmarshal([]byte(data), &unmarshalled); err != nil {
		return nil, errors.New("Map: YAML decode failed with: " + err.Error())
	}
	//return unmarshalled, nil
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
	val := C.GoString(AValue)
	m.Set(C.GoString(APath), val)
	data, err := yaml.Marshal(m)
	if err != nil {
		return 1
	}
	err = ioutil.WriteFile(C.GoString(AFileName), data, 0644)
	if err != nil {
		return 1
	}
	return 0
}

func main() {
	// Need a main function to make CGO compile package as C shared library
	// var value string
	// var valueLen int
	// YAMLReadString("example.yaml", "str", "default", &value, &valueLen)
	// fmt.Println(value)
}
