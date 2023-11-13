package goyml

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// Parse creates a new YamlQuery obj from an slice []byte.
func Parse(data []byte) *YamlQuery {
	d := map[string]interface{}{}
	dec := yaml.NewDecoder(strings.NewReader(string(data)))
	dec.Decode(&d)
	return newQuery(d)
}

// Bool extracts a bool the YamlQuery
func (j *YamlQuery) Bool(s ...string) (bool, error) {
	val, err := rquery(j.blob, s...)
	if err != nil {
		return false, err
	}
	return interfaceToBool(val)
}

// Float extracts a float from the YamlQuery
func (j *YamlQuery) Float(s ...string) (float64, error) {
	val, err := rquery(j.blob, s...)
	if err != nil {
		return 0.0, err
	}
	return interfaceToFloat(val)
}

// Int extracts an int from the YamlQuery
func (j *YamlQuery) Int(s ...string) (int, error) {
	val, err := rquery(j.blob, s...)
	if err != nil {
		return 0, err
	}
	return interfaceToInt(val)
}

// String extracts a string from the YamlQuery
func (j *YamlQuery) String(s ...string) (string, error) {
	val, err := rquery(j.blob, s...)
	if err != nil {
		return "", err
	}
	return interfaceToString(val)
}

// Object extracts a yaml object from the YamlQuery
func (j *YamlQuery) Object(s ...string) (map[string]interface{}, error) {
	val, err := rquery(j.blob, s...)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return interfaceToObj(val)
}

// Array extracts a []interface{} from the YamlQuery
func (j *YamlQuery) Array(s ...string) ([]interface{}, error) {
	val, err := rquery(j.blob, s...)
	if err != nil {
		return []interface{}{}, err
	}
	return interfaceToArray(val)
}

// Interface extracts an interface{} from the YamlQuery
func (j *YamlQuery) Interface(s ...string) (interface{}, error) {
	val, err := rquery(j.blob, s...)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// ArrayOfStrings extracts an array of strings from some yaml
func (j *YamlQuery) ArrayOfStrings(s ...string) ([]string, error) {
	array, err := j.Array(s...)
	if err != nil {
		return []string{}, err
	}
	toReturn := make([]string, len(array))
	for index, val := range array {
		toReturn[index], err = interfaceToString(val)
		if err != nil {
			return toReturn, err
		}
	}
	return toReturn, nil
}

// ArrayOfInts extracts an array of ints from some yaml
func (j *YamlQuery) ArrayOfInts(s ...string) ([]int, error) {
	array, err := j.Array(s...)
	if err != nil {
		return []int{}, err
	}
	toReturn := make([]int, len(array))
	for index, val := range array {
		toReturn[index], err = interfaceToInt(val)
		if err != nil {
			return toReturn, err
		}
	}
	return toReturn, nil
}

// ArrayOfFloats extracts an array of float64s from some yaml
func (j *YamlQuery) ArrayOfFloats(s ...string) ([]float64, error) {
	array, err := j.Array(s...)
	if err != nil {
		return []float64{}, err
	}
	toReturn := make([]float64, len(array))
	for index, val := range array {
		toReturn[index], err = interfaceToFloat(val)
		if err != nil {
			return toReturn, err
		}
	}
	return toReturn, nil
}

// ArrayOfBools extracts an array of bools from some yaml
func (j *YamlQuery) ArrayOfBools(s ...string) ([]bool, error) {
	array, err := j.Array(s...)
	if err != nil {
		return []bool{}, err
	}
	toReturn := make([]bool, len(array))
	for index, val := range array {
		toReturn[index], err = interfaceToBool(val)
		if err != nil {
			return toReturn, err
		}
	}
	return toReturn, nil
}

// ArrayOfObjects extracts an array of map[string]interface{} (objects) from some yaml
func (j *YamlQuery) ArrayOfObjects(s ...string) ([]map[string]interface{}, error) {
	array, err := j.Array(s...)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	toReturn := make([]map[string]interface{}, len(array))
	for index, val := range array {
		toReturn[index], err = interfaceToObj(val)
		if err != nil {
			return toReturn, err
		}
	}
	return toReturn, nil
}

// ArrayOfArrays extracts an array of []interface{} (arrays) from some yaml
func (j *YamlQuery) ArrayOfArrays(s ...string) ([][]interface{}, error) {
	array, err := j.Array(s...)
	if err != nil {
		return [][]interface{}{}, err
	}
	toReturn := make([][]interface{}, len(array))
	for index, val := range array {
		toReturn[index], err = interfaceToArray(val)
		if err != nil {
			return toReturn, err
		}
	}
	return toReturn, nil
}

// Matrix2D is an alias for ArrayOfArrays
func (j *YamlQuery) Matrix2D(s ...string) ([][]interface{}, error) {
	return j.ArrayOfArrays(s...)
}
