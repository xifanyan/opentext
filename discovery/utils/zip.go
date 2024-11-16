package utils

import "fmt"

// Zip creates a map by zipping together the provided keys and values slices.
// It returns an error if the lengths of the keys and values slices are not equal.
func Zip(keys, vals []string) (map[string]string, error) {
	if len(keys) != len(vals) {
		return nil, fmt.Errorf("lengths not equal: keys=%d vals=%d", len(keys), len(vals))
	}
	m := make(map[string]string)
	for i := range keys {
		m[keys[i]] = vals[i]
	}
	return m, nil
}
