package utils

import "reflect"

func GetOrDefault[T comparable](v T, d T) T {
	if IsZeroVal(v) {
		return d
	}
	return v
}

func GetOrDefaultFunc[T comparable](v T, a func() T) T {
	if IsZeroVal(v) {
		return a()
	}
	return v
}

// IsZeroVal check if any type is its zero value
func IsZeroVal(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
