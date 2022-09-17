package di

func mapSlice[T any, R any] (source []T, mapper func (T) R) []R {
	results := make([]R, len(source))
	for i, item := range source {
		results[i] = mapper(item)
	}
	return results
}

func cloneSlice[T any](source []T) []T {
	results := make([]T, len(source))
	copy(results, source)
	return results
}

func flattenSlice[T any] (source [][]T) []T {
	results := make([]T, 0)
	for _, item := range source {
		results = append(results, item...)
	}
	return results
}

func bindSlice[T any, R any] (source []T, binder func (T) []R) []R {
	return flattenSlice(mapSlice(source, binder))
}

func filterSlice[T any] (source []T, filter func (T) bool) []T {
	results := make([]T, 0)
	for _, item := range source {
		if filter(item) {
			results = append(results, item)
		}
	}
	return results
}

func findSlice[T any] (source []T, filter func (T) bool) *T {
	for _, item := range source {
		if filter(item) {
			return &item
		}
	}
	return nil
}

func filterSliceIndices[T any] (source []T, filter func (T) bool) []int {
	results := make([]int, 0)
	for i, item := range source {
		if filter(item) {
			results = append(results, i)
		}
	}
	return results
}

func filterSliceAll[T any] (source []T, filter func (T) bool) ([]T, []int) {
	results := make([]T, 0)
	indices := make([]int, 0)
	for i, item := range source {
		if filter(item) {
			results = append(results, item)
			indices = append(indices, i)
		}
	}
	return results, indices
}

func partitionSlice[T any] (source []T, predicate func (T) bool) (trueItems []T, falseItems []T) {
	trueItems = make([]T, 0)
	falseItems = make([]T, 0)
	for _, item := range source {
		if predicate(item) {
			trueItems = append(trueItems, item)
		} else {
			falseItems = append(falseItems, item)
		}
	}
	return trueItems, falseItems
}

func rangeSlice(start int, count int) []int {
	results := make([]int, count)
	for i := 0; i < count; i++ {
		results[i] = start + i
	}
	return results
}

func rangeSliceWithStep(start int, count int, step int) []int {
	results := make([]int, count)
	for i := 0; i < count; i++ {
		results[i] = start + i * step
	}
	return results
}

func rangeMapSlice[T any](start int, count int, mapper func (int) T) []T {
	results := make([]T, count)
	for i := 0; i < count; i++ {
		results[i] = mapper(start + i)
	}
	return results
}
