package test

import "testing"

type Data[T any] struct {
	Name string
	Data T
}

func ParameterizedTest[T any](t *testing.T, data []*Data[T], testFunc func(t *testing.T, data T)) {
	for _, item := range data {
		t.Run(item.Name, func(t *testing.T) { testFunc(t, item.Data) })
	}
}
