package goyml

import (
	"fmt"
	"strconv"
)

// YamlQuery is an object that enables querying of a Go map with a simple
// positional query language.
type YamlQuery struct {
	blob map[string]interface{}
}

// newQuery creates a new YamlQuery obj from an interface{}.
func newQuery(data interface{}) *YamlQuery {
	j := new(YamlQuery)
	j.blob = data.(map[string]interface{})
	return j
}

// interfaceToString converts an interface{} to a string and returns an error if types don't match.
func interfaceToString(val interface{}) (string, error) {
	switch val.(type) {
	case string:
		return val.(string), nil
	}
	return "", fmt.Errorf("Expected string value for String, got \"%v\"\n", val)
}

// interfaceToBool converts an interface{} to a bool and returns an error if types don't match.
func interfaceToBool(val interface{}) (bool, error) {
	switch val.(type) {
	case bool:
		return val.(bool), nil
	}
	return false, fmt.Errorf("Expected boolean value for Bool, got \"%v\"\n", val)
}

// interfaceToFloat converts an interface{} to a float64 and returns an error if types don't match.
func interfaceToFloat(val interface{}) (float64, error) {
	switch val.(type) {
	case float64:
		return val.(float64), nil
	case int:
		return float64(val.(int)), nil
	case string:
		fval, err := strconv.ParseFloat(val.(string), 64)
		if err == nil {
			return fval, nil
		}
	}
	return 0.0, fmt.Errorf("Expected numeric value for Float, got \"%v\"\n", val)
}

// interfaceToInt converts an interface{} to an int and returns an error if types don't match.
func interfaceToInt(val interface{}) (int, error) {
	switch val.(type) {
	case float64:
		return int(val.(float64)), nil
	case string:
		ival, err := strconv.ParseFloat(val.(string), 64)
		if err == nil {
			return int(ival), nil
		}
	case int:
		return val.(int), nil
	}
	return 0, fmt.Errorf("Expected numeric value for Int, got \"%v\"\n", val)
}

// interfaceToObj converts an interface{} to a map[string]interface{} and returns an error if types don't match.
func interfaceToObj(val interface{}) (map[string]interface{}, error) {
	switch val.(type) {
	case map[string]interface{}:
		return val.(map[string]interface{}), nil
	}
	return map[string]interface{}{}, fmt.Errorf("Expected yaml object for Object, got \"%v\"\n", val)
}

// interfaceToArray converts an interface{} to an []interface{} and returns an error if types don't match.
func interfaceToArray(val interface{}) ([]interface{}, error) {
	switch val.(type) {
	case []interface{}:
		return val.([]interface{}), nil
	}
	return []interface{}{}, fmt.Errorf("Expected yaml array for Array, got \"%v\"\n", val)
}

// Recursively query a decoded yaml blob
func rquery(blob interface{}, s ...string) (interface{}, error) {
	var (
		val interface{}
		err error
	)
	val = blob
	for _, q := range s {
		val, err = query(val, q)
		if err != nil {
			return nil, err
		}
	}
	switch val.(type) {
	case nil:
		return nil, fmt.Errorf("Nil value found at %s\n", s[len(s)-1])
	}
	return val, nil
}

// query a yaml blob for a single field or index.  If query is a string, then
// the blob is treated as a yaml object (map[string]interface{}).  If query is
// an integer, the blob is treated as a yaml array ([]interface{}).  Any kind
// of key or index error will result in a nil return value with an error set.
func query(blob interface{}, query string) (interface{}, error) {
	index, err := strconv.Atoi(query)
	// if it's an integer, then we treat the current interface as an array
	if err == nil {
		switch blob.(type) {
		case []interface{}:
		default:
			return nil, fmt.Errorf("Array index on non-array %v\n", blob)
		}
		if len(blob.([]interface{})) > index {
			return blob.([]interface{})[index], nil
		}
		return nil, fmt.Errorf("Array index %d on array %v out of bounds\n", index, blob)
	}

	// blob is likely an object, but verify first
	switch blob.(type) {
	case map[string]interface{}:
	default:
		return nil, fmt.Errorf("Object lookup \"%s\" on non-object %v\n", query, blob)
	}

	val, ok := blob.(map[string]interface{})[query]
	if !ok {
		return nil, fmt.Errorf("Object %v does not contain field %s\n", blob, query)
	}
	return val, nil
}
