package definitions

import (
	"regexp"
	"strings"
)

func mapJSONKeys(from interface{}, mapper func(string) string) (interface{}, error) {
	switch f := from.(type) {
	case map[string]interface{}:
		o := make(map[string]interface{})
		for k, v := range f {
			newK := mapper(k)
			newV, err := mapJSONKeys(v, mapper)
			if err != nil {
				return nil, err
			}
			o[newK] = newV
		}
		return o, nil
	case []interface{}:
		xs := make([]interface{}, 0, len(f))
		for _, x := range f {
			newV, err := mapJSONKeys(x, mapper)
			if err != nil {
				return nil, err
			}
			xs = append(xs, newV)
		}
		return xs, nil
	default:
		// Primitive types
		return from, nil
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
