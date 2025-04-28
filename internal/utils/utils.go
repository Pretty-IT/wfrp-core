package utils

import "errors"

func FilterList[T any](input []T, filterFunc func(T) bool) (output []T) {
	for _, item := range input {
		if filterFunc(item) {
			output = append(output, item)
		}
	}
	return
}

func FindFirst[T any](input []T, filterFunc func(T) bool) (T, error) {
	var defaultResult T
	for _, item := range input {
		if filterFunc(item) {
			return item, nil
		}
	}

	return defaultResult, errors.New("item not found")
}

func MapList[T, U any](input []T, mapper func(T) U) (output []U) {
	output = make([]U, len(input))
	for i := range input {
		output[i] = mapper(input[i])
	}
	return output
}

func MapMap[T, U any](input map[string]T, mapper func(T) U) (output map[string]U) {
	output = make(map[string]U)

	for key, value := range input {
		output[key] = mapper(value)
	}
	return output
}

func SliceToMap[T any, K ~int](input []T, keyFunc func(T) K) (output map[K]T) {
	for _, item := range input {
		key := keyFunc(item)
		output[key] = item
	}
	return
}
