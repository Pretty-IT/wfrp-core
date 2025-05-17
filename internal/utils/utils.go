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

func MapMap[K ~int | string, V1, V2 any](input map[K]V1, mapper func(V1) V2) (output map[K]V2) {
	output = make(map[K]V2)

	for key, value := range input {
		output[key] = mapper(value)
	}
	return output
}

func SliceToMap[T any, K ~int](input []T, keyFunc func(T) K) (output map[K]T) {
	output = make(map[K]T)

	for _, item := range input {
		key := keyFunc(item)
		output[key] = item
	}
	return
}

func CopyMap[T any, K ~int](input map[K]T) (output map[K]T) {
	output = make(map[K]T)

	for k, v := range input {
		output[k] = v
	}
	return
}

func CopyList[T any](input []T) (output []T) {
	output = make([]T, len(input))
	for i := range input {
		output[i] = input[i]
	}
	return output
}
