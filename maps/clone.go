package maps

func Clone[TKey comparable, TValue any](source map[TKey]TValue) map[TKey]TValue {
	destination := make(map[TKey]TValue, len(source))

	for key, value := range source {
		destination[key] = value
	}

	return destination
}
